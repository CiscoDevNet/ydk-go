// This module contains a collection of YANG definitions
// for Cisco IOS-XR lib-mpp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   control-plane: Configure control Plane
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package lib_mpp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lib_mpp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lib-mpp-cfg control-plane}", reflect.TypeOf(ControlPlane{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lib-mpp-cfg:control-plane", reflect.TypeOf(ControlPlane{}))
}

// ControlPlane
// Configure control Plane
type ControlPlane struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure management plane protection.
    ManagementPlaneProtection ControlPlane_ManagementPlaneProtection
}

func (controlPlane *ControlPlane) GetEntityData() *types.CommonEntityData {
    controlPlane.EntityData.YFilter = controlPlane.YFilter
    controlPlane.EntityData.YangName = "control-plane"
    controlPlane.EntityData.BundleName = "cisco_ios_xr"
    controlPlane.EntityData.ParentYangName = "Cisco-IOS-XR-lib-mpp-cfg"
    controlPlane.EntityData.SegmentPath = "Cisco-IOS-XR-lib-mpp-cfg:control-plane"
    controlPlane.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controlPlane.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controlPlane.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controlPlane.EntityData.Children = types.NewOrderedMap()
    controlPlane.EntityData.Children.Append("management-plane-protection", types.YChild{"ManagementPlaneProtection", &controlPlane.ManagementPlaneProtection})
    controlPlane.EntityData.Leafs = types.NewOrderedMap()

    controlPlane.EntityData.YListKeys = []string {}

    return &(controlPlane.EntityData)
}

// ControlPlane_ManagementPlaneProtection
// Configure management plane protection
type ControlPlane_ManagementPlaneProtection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outband Configuration.
    Outband ControlPlane_ManagementPlaneProtection_Outband

    // Inband Configuration.
    Inband ControlPlane_ManagementPlaneProtection_Inband

    // MPP Third Party Application Configuration Commands.
    Tpa ControlPlane_ManagementPlaneProtection_Tpa
}

func (managementPlaneProtection *ControlPlane_ManagementPlaneProtection) GetEntityData() *types.CommonEntityData {
    managementPlaneProtection.EntityData.YFilter = managementPlaneProtection.YFilter
    managementPlaneProtection.EntityData.YangName = "management-plane-protection"
    managementPlaneProtection.EntityData.BundleName = "cisco_ios_xr"
    managementPlaneProtection.EntityData.ParentYangName = "control-plane"
    managementPlaneProtection.EntityData.SegmentPath = "management-plane-protection"
    managementPlaneProtection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    managementPlaneProtection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    managementPlaneProtection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    managementPlaneProtection.EntityData.Children = types.NewOrderedMap()
    managementPlaneProtection.EntityData.Children.Append("outband", types.YChild{"Outband", &managementPlaneProtection.Outband})
    managementPlaneProtection.EntityData.Children.Append("inband", types.YChild{"Inband", &managementPlaneProtection.Inband})
    managementPlaneProtection.EntityData.Children.Append("tpa", types.YChild{"Tpa", &managementPlaneProtection.Tpa})
    managementPlaneProtection.EntityData.Leafs = types.NewOrderedMap()

    managementPlaneProtection.EntityData.YListKeys = []string {}

    return &(managementPlaneProtection.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband
// Outband Configuration
type ControlPlane_ManagementPlaneProtection_Outband struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure outband VRF. The type is string.
    OutbandVrf interface{}

    // Configure interfaces.
    InterfaceSelection ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection
}

func (outband *ControlPlane_ManagementPlaneProtection_Outband) GetEntityData() *types.CommonEntityData {
    outband.EntityData.YFilter = outband.YFilter
    outband.EntityData.YangName = "outband"
    outband.EntityData.BundleName = "cisco_ios_xr"
    outband.EntityData.ParentYangName = "management-plane-protection"
    outband.EntityData.SegmentPath = "outband"
    outband.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outband.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outband.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outband.EntityData.Children = types.NewOrderedMap()
    outband.EntityData.Children.Append("interface-selection", types.YChild{"InterfaceSelection", &outband.InterfaceSelection})
    outband.EntityData.Leafs = types.NewOrderedMap()
    outband.EntityData.Leafs.Append("outband-vrf", types.YLeaf{"OutbandVrf", outband.OutbandVrf})

    outband.EntityData.YListKeys = []string {}

    return &(outband.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection
// Configure interfaces
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure a specific interface.
    Interfaces ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces

    // Configure all Inband interfaces.
    AllInterfaces ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces
}

func (interfaceSelection *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection) GetEntityData() *types.CommonEntityData {
    interfaceSelection.EntityData.YFilter = interfaceSelection.YFilter
    interfaceSelection.EntityData.YangName = "interface-selection"
    interfaceSelection.EntityData.BundleName = "cisco_ios_xr"
    interfaceSelection.EntityData.ParentYangName = "outband"
    interfaceSelection.EntityData.SegmentPath = "interface-selection"
    interfaceSelection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSelection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSelection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSelection.EntityData.Children = types.NewOrderedMap()
    interfaceSelection.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &interfaceSelection.Interfaces})
    interfaceSelection.EntityData.Children.Append("all-interfaces", types.YChild{"AllInterfaces", &interfaceSelection.AllInterfaces})
    interfaceSelection.EntityData.Leafs = types.NewOrderedMap()

    interfaceSelection.EntityData.YListKeys = []string {}

    return &(interfaceSelection.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces
// Configure a specific interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specific interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface.
    Interface []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface
}

func (interfaces *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "interface-selection"
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

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface
// Specific interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Configure HTTP on this interface.
    HttpProtocol ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol

    // Configure TFTP on this interface.
    TftpProtocol ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol

    // Configure NETCONF protocol and peer addresses.
    NetconfProtocol ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol

    // Configure XML and peer addresses.
    XrXml ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml

    // Configure SSH protocol and peer addresses.
    SshProtocol ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol

    // Configure SNMP for this interface.
    SnmpProtocol ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol

    // Configure Telnet for this interface.
    TelnetProtocol ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol

    // Configure all protocols on this interface.
    AllProtocols ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols
}

func (self *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("http-protocol", types.YChild{"HttpProtocol", &self.HttpProtocol})
    self.EntityData.Children.Append("tftp-protocol", types.YChild{"TftpProtocol", &self.TftpProtocol})
    self.EntityData.Children.Append("netconf-protocol", types.YChild{"NetconfProtocol", &self.NetconfProtocol})
    self.EntityData.Children.Append("xr-xml", types.YChild{"XrXml", &self.XrXml})
    self.EntityData.Children.Append("ssh-protocol", types.YChild{"SshProtocol", &self.SshProtocol})
    self.EntityData.Children.Append("snmp-protocol", types.YChild{"SnmpProtocol", &self.SnmpProtocol})
    self.EntityData.Children.Append("telnet-protocol", types.YChild{"TelnetProtocol", &self.TelnetProtocol})
    self.EntityData.Children.Append("all-protocols", types.YChild{"AllProtocols", &self.AllProtocols})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol
// Configure HTTP on this interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass
}

func (httpProtocol *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol) GetEntityData() *types.CommonEntityData {
    httpProtocol.EntityData.YFilter = httpProtocol.YFilter
    httpProtocol.EntityData.YangName = "http-protocol"
    httpProtocol.EntityData.BundleName = "cisco_ios_xr"
    httpProtocol.EntityData.ParentYangName = "interface"
    httpProtocol.EntityData.SegmentPath = "http-protocol"
    httpProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    httpProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    httpProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    httpProtocol.EntityData.Children = types.NewOrderedMap()
    httpProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &httpProtocol.PeerClass})
    httpProtocol.EntityData.Leafs = types.NewOrderedMap()

    httpProtocol.EntityData.YListKeys = []string {}

    return &(httpProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "http-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol
// Configure TFTP on this interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass
}

func (tftpProtocol *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol) GetEntityData() *types.CommonEntityData {
    tftpProtocol.EntityData.YFilter = tftpProtocol.YFilter
    tftpProtocol.EntityData.YangName = "tftp-protocol"
    tftpProtocol.EntityData.BundleName = "cisco_ios_xr"
    tftpProtocol.EntityData.ParentYangName = "interface"
    tftpProtocol.EntityData.SegmentPath = "tftp-protocol"
    tftpProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tftpProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tftpProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tftpProtocol.EntityData.Children = types.NewOrderedMap()
    tftpProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &tftpProtocol.PeerClass})
    tftpProtocol.EntityData.Leafs = types.NewOrderedMap()

    tftpProtocol.EntityData.YListKeys = []string {}

    return &(tftpProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "tftp-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol
// Configure NETCONF protocol and peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass
}

func (netconfProtocol *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol) GetEntityData() *types.CommonEntityData {
    netconfProtocol.EntityData.YFilter = netconfProtocol.YFilter
    netconfProtocol.EntityData.YangName = "netconf-protocol"
    netconfProtocol.EntityData.BundleName = "cisco_ios_xr"
    netconfProtocol.EntityData.ParentYangName = "interface"
    netconfProtocol.EntityData.SegmentPath = "netconf-protocol"
    netconfProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netconfProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netconfProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netconfProtocol.EntityData.Children = types.NewOrderedMap()
    netconfProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &netconfProtocol.PeerClass})
    netconfProtocol.EntityData.Leafs = types.NewOrderedMap()

    netconfProtocol.EntityData.YListKeys = []string {}

    return &(netconfProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "netconf-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml
// Configure XML and peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass
}

func (xrXml *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml) GetEntityData() *types.CommonEntityData {
    xrXml.EntityData.YFilter = xrXml.YFilter
    xrXml.EntityData.YangName = "xr-xml"
    xrXml.EntityData.BundleName = "cisco_ios_xr"
    xrXml.EntityData.ParentYangName = "interface"
    xrXml.EntityData.SegmentPath = "xr-xml"
    xrXml.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xrXml.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xrXml.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xrXml.EntityData.Children = types.NewOrderedMap()
    xrXml.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &xrXml.PeerClass})
    xrXml.EntityData.Leafs = types.NewOrderedMap()

    xrXml.EntityData.YListKeys = []string {}

    return &(xrXml.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "xr-xml"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol
// Configure SSH protocol and peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass
}

func (sshProtocol *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol) GetEntityData() *types.CommonEntityData {
    sshProtocol.EntityData.YFilter = sshProtocol.YFilter
    sshProtocol.EntityData.YangName = "ssh-protocol"
    sshProtocol.EntityData.BundleName = "cisco_ios_xr"
    sshProtocol.EntityData.ParentYangName = "interface"
    sshProtocol.EntityData.SegmentPath = "ssh-protocol"
    sshProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sshProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sshProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sshProtocol.EntityData.Children = types.NewOrderedMap()
    sshProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &sshProtocol.PeerClass})
    sshProtocol.EntityData.Leafs = types.NewOrderedMap()

    sshProtocol.EntityData.YListKeys = []string {}

    return &(sshProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "ssh-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol
// Configure SNMP for this interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass
}

func (snmpProtocol *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol) GetEntityData() *types.CommonEntityData {
    snmpProtocol.EntityData.YFilter = snmpProtocol.YFilter
    snmpProtocol.EntityData.YangName = "snmp-protocol"
    snmpProtocol.EntityData.BundleName = "cisco_ios_xr"
    snmpProtocol.EntityData.ParentYangName = "interface"
    snmpProtocol.EntityData.SegmentPath = "snmp-protocol"
    snmpProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpProtocol.EntityData.Children = types.NewOrderedMap()
    snmpProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &snmpProtocol.PeerClass})
    snmpProtocol.EntityData.Leafs = types.NewOrderedMap()

    snmpProtocol.EntityData.YListKeys = []string {}

    return &(snmpProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "snmp-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol
// Configure Telnet for this interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass
}

func (telnetProtocol *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol) GetEntityData() *types.CommonEntityData {
    telnetProtocol.EntityData.YFilter = telnetProtocol.YFilter
    telnetProtocol.EntityData.YangName = "telnet-protocol"
    telnetProtocol.EntityData.BundleName = "cisco_ios_xr"
    telnetProtocol.EntityData.ParentYangName = "interface"
    telnetProtocol.EntityData.SegmentPath = "telnet-protocol"
    telnetProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    telnetProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    telnetProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    telnetProtocol.EntityData.Children = types.NewOrderedMap()
    telnetProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &telnetProtocol.PeerClass})
    telnetProtocol.EntityData.Leafs = types.NewOrderedMap()

    telnetProtocol.EntityData.YListKeys = []string {}

    return &(telnetProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "telnet-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols
// Configure all protocols on this interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass
}

func (allProtocols *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols) GetEntityData() *types.CommonEntityData {
    allProtocols.EntityData.YFilter = allProtocols.YFilter
    allProtocols.EntityData.YangName = "all-protocols"
    allProtocols.EntityData.BundleName = "cisco_ios_xr"
    allProtocols.EntityData.ParentYangName = "interface"
    allProtocols.EntityData.SegmentPath = "all-protocols"
    allProtocols.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allProtocols.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allProtocols.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allProtocols.EntityData.Children = types.NewOrderedMap()
    allProtocols.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &allProtocols.PeerClass})
    allProtocols.EntityData.Leafs = types.NewOrderedMap()

    allProtocols.EntityData.YListKeys = []string {}

    return &(allProtocols.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "all-protocols"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces
// Configure all Inband interfaces
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure HTTP on this interface.
    HttpProtocol ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol

    // Configure TFTP on this interface.
    TftpProtocol ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol

    // Configure NETCONF protocol and peer addresses.
    NetconfProtocol ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol

    // Configure XML and peer addresses.
    XrXml ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml

    // Configure SSH protocol and peer addresses.
    SshProtocol ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol

    // Configure SNMP for this interface.
    SnmpProtocol ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol

    // Configure Telnet for this interface.
    TelnetProtocol ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol

    // Configure all protocols on this interface.
    AllProtocols ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols
}

func (allInterfaces *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces) GetEntityData() *types.CommonEntityData {
    allInterfaces.EntityData.YFilter = allInterfaces.YFilter
    allInterfaces.EntityData.YangName = "all-interfaces"
    allInterfaces.EntityData.BundleName = "cisco_ios_xr"
    allInterfaces.EntityData.ParentYangName = "interface-selection"
    allInterfaces.EntityData.SegmentPath = "all-interfaces"
    allInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allInterfaces.EntityData.Children = types.NewOrderedMap()
    allInterfaces.EntityData.Children.Append("http-protocol", types.YChild{"HttpProtocol", &allInterfaces.HttpProtocol})
    allInterfaces.EntityData.Children.Append("tftp-protocol", types.YChild{"TftpProtocol", &allInterfaces.TftpProtocol})
    allInterfaces.EntityData.Children.Append("netconf-protocol", types.YChild{"NetconfProtocol", &allInterfaces.NetconfProtocol})
    allInterfaces.EntityData.Children.Append("xr-xml", types.YChild{"XrXml", &allInterfaces.XrXml})
    allInterfaces.EntityData.Children.Append("ssh-protocol", types.YChild{"SshProtocol", &allInterfaces.SshProtocol})
    allInterfaces.EntityData.Children.Append("snmp-protocol", types.YChild{"SnmpProtocol", &allInterfaces.SnmpProtocol})
    allInterfaces.EntityData.Children.Append("telnet-protocol", types.YChild{"TelnetProtocol", &allInterfaces.TelnetProtocol})
    allInterfaces.EntityData.Children.Append("all-protocols", types.YChild{"AllProtocols", &allInterfaces.AllProtocols})
    allInterfaces.EntityData.Leafs = types.NewOrderedMap()

    allInterfaces.EntityData.YListKeys = []string {}

    return &(allInterfaces.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol
// Configure HTTP on this interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass
}

func (httpProtocol *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol) GetEntityData() *types.CommonEntityData {
    httpProtocol.EntityData.YFilter = httpProtocol.YFilter
    httpProtocol.EntityData.YangName = "http-protocol"
    httpProtocol.EntityData.BundleName = "cisco_ios_xr"
    httpProtocol.EntityData.ParentYangName = "all-interfaces"
    httpProtocol.EntityData.SegmentPath = "http-protocol"
    httpProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    httpProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    httpProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    httpProtocol.EntityData.Children = types.NewOrderedMap()
    httpProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &httpProtocol.PeerClass})
    httpProtocol.EntityData.Leafs = types.NewOrderedMap()

    httpProtocol.EntityData.YListKeys = []string {}

    return &(httpProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "http-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol
// Configure TFTP on this interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass
}

func (tftpProtocol *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol) GetEntityData() *types.CommonEntityData {
    tftpProtocol.EntityData.YFilter = tftpProtocol.YFilter
    tftpProtocol.EntityData.YangName = "tftp-protocol"
    tftpProtocol.EntityData.BundleName = "cisco_ios_xr"
    tftpProtocol.EntityData.ParentYangName = "all-interfaces"
    tftpProtocol.EntityData.SegmentPath = "tftp-protocol"
    tftpProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tftpProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tftpProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tftpProtocol.EntityData.Children = types.NewOrderedMap()
    tftpProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &tftpProtocol.PeerClass})
    tftpProtocol.EntityData.Leafs = types.NewOrderedMap()

    tftpProtocol.EntityData.YListKeys = []string {}

    return &(tftpProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "tftp-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol
// Configure NETCONF protocol and peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass
}

func (netconfProtocol *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol) GetEntityData() *types.CommonEntityData {
    netconfProtocol.EntityData.YFilter = netconfProtocol.YFilter
    netconfProtocol.EntityData.YangName = "netconf-protocol"
    netconfProtocol.EntityData.BundleName = "cisco_ios_xr"
    netconfProtocol.EntityData.ParentYangName = "all-interfaces"
    netconfProtocol.EntityData.SegmentPath = "netconf-protocol"
    netconfProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netconfProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netconfProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netconfProtocol.EntityData.Children = types.NewOrderedMap()
    netconfProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &netconfProtocol.PeerClass})
    netconfProtocol.EntityData.Leafs = types.NewOrderedMap()

    netconfProtocol.EntityData.YListKeys = []string {}

    return &(netconfProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "netconf-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml
// Configure XML and peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass
}

func (xrXml *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml) GetEntityData() *types.CommonEntityData {
    xrXml.EntityData.YFilter = xrXml.YFilter
    xrXml.EntityData.YangName = "xr-xml"
    xrXml.EntityData.BundleName = "cisco_ios_xr"
    xrXml.EntityData.ParentYangName = "all-interfaces"
    xrXml.EntityData.SegmentPath = "xr-xml"
    xrXml.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xrXml.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xrXml.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xrXml.EntityData.Children = types.NewOrderedMap()
    xrXml.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &xrXml.PeerClass})
    xrXml.EntityData.Leafs = types.NewOrderedMap()

    xrXml.EntityData.YListKeys = []string {}

    return &(xrXml.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "xr-xml"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol
// Configure SSH protocol and peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass
}

func (sshProtocol *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol) GetEntityData() *types.CommonEntityData {
    sshProtocol.EntityData.YFilter = sshProtocol.YFilter
    sshProtocol.EntityData.YangName = "ssh-protocol"
    sshProtocol.EntityData.BundleName = "cisco_ios_xr"
    sshProtocol.EntityData.ParentYangName = "all-interfaces"
    sshProtocol.EntityData.SegmentPath = "ssh-protocol"
    sshProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sshProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sshProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sshProtocol.EntityData.Children = types.NewOrderedMap()
    sshProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &sshProtocol.PeerClass})
    sshProtocol.EntityData.Leafs = types.NewOrderedMap()

    sshProtocol.EntityData.YListKeys = []string {}

    return &(sshProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "ssh-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol
// Configure SNMP for this interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass
}

func (snmpProtocol *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol) GetEntityData() *types.CommonEntityData {
    snmpProtocol.EntityData.YFilter = snmpProtocol.YFilter
    snmpProtocol.EntityData.YangName = "snmp-protocol"
    snmpProtocol.EntityData.BundleName = "cisco_ios_xr"
    snmpProtocol.EntityData.ParentYangName = "all-interfaces"
    snmpProtocol.EntityData.SegmentPath = "snmp-protocol"
    snmpProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpProtocol.EntityData.Children = types.NewOrderedMap()
    snmpProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &snmpProtocol.PeerClass})
    snmpProtocol.EntityData.Leafs = types.NewOrderedMap()

    snmpProtocol.EntityData.YListKeys = []string {}

    return &(snmpProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "snmp-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol
// Configure Telnet for this interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass
}

func (telnetProtocol *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol) GetEntityData() *types.CommonEntityData {
    telnetProtocol.EntityData.YFilter = telnetProtocol.YFilter
    telnetProtocol.EntityData.YangName = "telnet-protocol"
    telnetProtocol.EntityData.BundleName = "cisco_ios_xr"
    telnetProtocol.EntityData.ParentYangName = "all-interfaces"
    telnetProtocol.EntityData.SegmentPath = "telnet-protocol"
    telnetProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    telnetProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    telnetProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    telnetProtocol.EntityData.Children = types.NewOrderedMap()
    telnetProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &telnetProtocol.PeerClass})
    telnetProtocol.EntityData.Leafs = types.NewOrderedMap()

    telnetProtocol.EntityData.YListKeys = []string {}

    return &(telnetProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "telnet-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols
// Configure all protocols on this interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass
}

func (allProtocols *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols) GetEntityData() *types.CommonEntityData {
    allProtocols.EntityData.YFilter = allProtocols.YFilter
    allProtocols.EntityData.YangName = "all-protocols"
    allProtocols.EntityData.BundleName = "cisco_ios_xr"
    allProtocols.EntityData.ParentYangName = "all-interfaces"
    allProtocols.EntityData.SegmentPath = "all-protocols"
    allProtocols.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allProtocols.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allProtocols.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allProtocols.EntityData.Children = types.NewOrderedMap()
    allProtocols.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &allProtocols.PeerClass})
    allProtocols.EntityData.Leafs = types.NewOrderedMap()

    allProtocols.EntityData.YListKeys = []string {}

    return &(allProtocols.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "all-protocols"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Outband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband
// Inband Configuration
type ControlPlane_ManagementPlaneProtection_Inband struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure interfaces.
    InterfaceSelection ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection
}

func (inband *ControlPlane_ManagementPlaneProtection_Inband) GetEntityData() *types.CommonEntityData {
    inband.EntityData.YFilter = inband.YFilter
    inband.EntityData.YangName = "inband"
    inband.EntityData.BundleName = "cisco_ios_xr"
    inband.EntityData.ParentYangName = "management-plane-protection"
    inband.EntityData.SegmentPath = "inband"
    inband.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inband.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inband.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inband.EntityData.Children = types.NewOrderedMap()
    inband.EntityData.Children.Append("interface-selection", types.YChild{"InterfaceSelection", &inband.InterfaceSelection})
    inband.EntityData.Leafs = types.NewOrderedMap()

    inband.EntityData.YListKeys = []string {}

    return &(inband.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection
// Configure interfaces
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure a specific interface.
    Interfaces ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces

    // Configure all Inband interfaces.
    AllInterfaces ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces
}

func (interfaceSelection *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection) GetEntityData() *types.CommonEntityData {
    interfaceSelection.EntityData.YFilter = interfaceSelection.YFilter
    interfaceSelection.EntityData.YangName = "interface-selection"
    interfaceSelection.EntityData.BundleName = "cisco_ios_xr"
    interfaceSelection.EntityData.ParentYangName = "inband"
    interfaceSelection.EntityData.SegmentPath = "interface-selection"
    interfaceSelection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSelection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSelection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSelection.EntityData.Children = types.NewOrderedMap()
    interfaceSelection.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &interfaceSelection.Interfaces})
    interfaceSelection.EntityData.Children.Append("all-interfaces", types.YChild{"AllInterfaces", &interfaceSelection.AllInterfaces})
    interfaceSelection.EntityData.Leafs = types.NewOrderedMap()

    interfaceSelection.EntityData.YListKeys = []string {}

    return &(interfaceSelection.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces
// Configure a specific interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specific interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface.
    Interface []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface
}

func (interfaces *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "interface-selection"
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

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface
// Specific interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Configure HTTP on this interface.
    HttpProtocol ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol

    // Configure TFTP on this interface.
    TftpProtocol ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol

    // Configure NETCONF protocol and peer addresses.
    NetconfProtocol ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol

    // Configure XML and peer addresses.
    XrXml ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml

    // Configure SSH protocol and peer addresses.
    SshProtocol ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol

    // Configure SNMP for this interface.
    SnmpProtocol ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol

    // Configure Telnet for this interface.
    TelnetProtocol ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol

    // Configure all protocols on this interface.
    AllProtocols ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols
}

func (self *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("http-protocol", types.YChild{"HttpProtocol", &self.HttpProtocol})
    self.EntityData.Children.Append("tftp-protocol", types.YChild{"TftpProtocol", &self.TftpProtocol})
    self.EntityData.Children.Append("netconf-protocol", types.YChild{"NetconfProtocol", &self.NetconfProtocol})
    self.EntityData.Children.Append("xr-xml", types.YChild{"XrXml", &self.XrXml})
    self.EntityData.Children.Append("ssh-protocol", types.YChild{"SshProtocol", &self.SshProtocol})
    self.EntityData.Children.Append("snmp-protocol", types.YChild{"SnmpProtocol", &self.SnmpProtocol})
    self.EntityData.Children.Append("telnet-protocol", types.YChild{"TelnetProtocol", &self.TelnetProtocol})
    self.EntityData.Children.Append("all-protocols", types.YChild{"AllProtocols", &self.AllProtocols})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol
// Configure HTTP on this interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass
}

func (httpProtocol *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol) GetEntityData() *types.CommonEntityData {
    httpProtocol.EntityData.YFilter = httpProtocol.YFilter
    httpProtocol.EntityData.YangName = "http-protocol"
    httpProtocol.EntityData.BundleName = "cisco_ios_xr"
    httpProtocol.EntityData.ParentYangName = "interface"
    httpProtocol.EntityData.SegmentPath = "http-protocol"
    httpProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    httpProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    httpProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    httpProtocol.EntityData.Children = types.NewOrderedMap()
    httpProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &httpProtocol.PeerClass})
    httpProtocol.EntityData.Leafs = types.NewOrderedMap()

    httpProtocol.EntityData.YListKeys = []string {}

    return &(httpProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "http-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol
// Configure TFTP on this interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass
}

func (tftpProtocol *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol) GetEntityData() *types.CommonEntityData {
    tftpProtocol.EntityData.YFilter = tftpProtocol.YFilter
    tftpProtocol.EntityData.YangName = "tftp-protocol"
    tftpProtocol.EntityData.BundleName = "cisco_ios_xr"
    tftpProtocol.EntityData.ParentYangName = "interface"
    tftpProtocol.EntityData.SegmentPath = "tftp-protocol"
    tftpProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tftpProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tftpProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tftpProtocol.EntityData.Children = types.NewOrderedMap()
    tftpProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &tftpProtocol.PeerClass})
    tftpProtocol.EntityData.Leafs = types.NewOrderedMap()

    tftpProtocol.EntityData.YListKeys = []string {}

    return &(tftpProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "tftp-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol
// Configure NETCONF protocol and peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass
}

func (netconfProtocol *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol) GetEntityData() *types.CommonEntityData {
    netconfProtocol.EntityData.YFilter = netconfProtocol.YFilter
    netconfProtocol.EntityData.YangName = "netconf-protocol"
    netconfProtocol.EntityData.BundleName = "cisco_ios_xr"
    netconfProtocol.EntityData.ParentYangName = "interface"
    netconfProtocol.EntityData.SegmentPath = "netconf-protocol"
    netconfProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netconfProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netconfProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netconfProtocol.EntityData.Children = types.NewOrderedMap()
    netconfProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &netconfProtocol.PeerClass})
    netconfProtocol.EntityData.Leafs = types.NewOrderedMap()

    netconfProtocol.EntityData.YListKeys = []string {}

    return &(netconfProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "netconf-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml
// Configure XML and peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass
}

func (xrXml *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml) GetEntityData() *types.CommonEntityData {
    xrXml.EntityData.YFilter = xrXml.YFilter
    xrXml.EntityData.YangName = "xr-xml"
    xrXml.EntityData.BundleName = "cisco_ios_xr"
    xrXml.EntityData.ParentYangName = "interface"
    xrXml.EntityData.SegmentPath = "xr-xml"
    xrXml.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xrXml.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xrXml.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xrXml.EntityData.Children = types.NewOrderedMap()
    xrXml.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &xrXml.PeerClass})
    xrXml.EntityData.Leafs = types.NewOrderedMap()

    xrXml.EntityData.YListKeys = []string {}

    return &(xrXml.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "xr-xml"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol
// Configure SSH protocol and peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass
}

func (sshProtocol *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol) GetEntityData() *types.CommonEntityData {
    sshProtocol.EntityData.YFilter = sshProtocol.YFilter
    sshProtocol.EntityData.YangName = "ssh-protocol"
    sshProtocol.EntityData.BundleName = "cisco_ios_xr"
    sshProtocol.EntityData.ParentYangName = "interface"
    sshProtocol.EntityData.SegmentPath = "ssh-protocol"
    sshProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sshProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sshProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sshProtocol.EntityData.Children = types.NewOrderedMap()
    sshProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &sshProtocol.PeerClass})
    sshProtocol.EntityData.Leafs = types.NewOrderedMap()

    sshProtocol.EntityData.YListKeys = []string {}

    return &(sshProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "ssh-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol
// Configure SNMP for this interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass
}

func (snmpProtocol *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol) GetEntityData() *types.CommonEntityData {
    snmpProtocol.EntityData.YFilter = snmpProtocol.YFilter
    snmpProtocol.EntityData.YangName = "snmp-protocol"
    snmpProtocol.EntityData.BundleName = "cisco_ios_xr"
    snmpProtocol.EntityData.ParentYangName = "interface"
    snmpProtocol.EntityData.SegmentPath = "snmp-protocol"
    snmpProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpProtocol.EntityData.Children = types.NewOrderedMap()
    snmpProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &snmpProtocol.PeerClass})
    snmpProtocol.EntityData.Leafs = types.NewOrderedMap()

    snmpProtocol.EntityData.YListKeys = []string {}

    return &(snmpProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "snmp-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol
// Configure Telnet for this interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass
}

func (telnetProtocol *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol) GetEntityData() *types.CommonEntityData {
    telnetProtocol.EntityData.YFilter = telnetProtocol.YFilter
    telnetProtocol.EntityData.YangName = "telnet-protocol"
    telnetProtocol.EntityData.BundleName = "cisco_ios_xr"
    telnetProtocol.EntityData.ParentYangName = "interface"
    telnetProtocol.EntityData.SegmentPath = "telnet-protocol"
    telnetProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    telnetProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    telnetProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    telnetProtocol.EntityData.Children = types.NewOrderedMap()
    telnetProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &telnetProtocol.PeerClass})
    telnetProtocol.EntityData.Leafs = types.NewOrderedMap()

    telnetProtocol.EntityData.YListKeys = []string {}

    return &(telnetProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "telnet-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols
// Configure all protocols on this interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass
}

func (allProtocols *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols) GetEntityData() *types.CommonEntityData {
    allProtocols.EntityData.YFilter = allProtocols.YFilter
    allProtocols.EntityData.YangName = "all-protocols"
    allProtocols.EntityData.BundleName = "cisco_ios_xr"
    allProtocols.EntityData.ParentYangName = "interface"
    allProtocols.EntityData.SegmentPath = "all-protocols"
    allProtocols.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allProtocols.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allProtocols.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allProtocols.EntityData.Children = types.NewOrderedMap()
    allProtocols.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &allProtocols.PeerClass})
    allProtocols.EntityData.Leafs = types.NewOrderedMap()

    allProtocols.EntityData.YListKeys = []string {}

    return &(allProtocols.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "all-protocols"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_Interfaces_Interface_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces
// Configure all Inband interfaces
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure HTTP on this interface.
    HttpProtocol ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol

    // Configure TFTP on this interface.
    TftpProtocol ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol

    // Configure NETCONF protocol and peer addresses.
    NetconfProtocol ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol

    // Configure XML and peer addresses.
    XrXml ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml

    // Configure SSH protocol and peer addresses.
    SshProtocol ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol

    // Configure SNMP for this interface.
    SnmpProtocol ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol

    // Configure Telnet for this interface.
    TelnetProtocol ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol

    // Configure all protocols on this interface.
    AllProtocols ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols
}

func (allInterfaces *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces) GetEntityData() *types.CommonEntityData {
    allInterfaces.EntityData.YFilter = allInterfaces.YFilter
    allInterfaces.EntityData.YangName = "all-interfaces"
    allInterfaces.EntityData.BundleName = "cisco_ios_xr"
    allInterfaces.EntityData.ParentYangName = "interface-selection"
    allInterfaces.EntityData.SegmentPath = "all-interfaces"
    allInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allInterfaces.EntityData.Children = types.NewOrderedMap()
    allInterfaces.EntityData.Children.Append("http-protocol", types.YChild{"HttpProtocol", &allInterfaces.HttpProtocol})
    allInterfaces.EntityData.Children.Append("tftp-protocol", types.YChild{"TftpProtocol", &allInterfaces.TftpProtocol})
    allInterfaces.EntityData.Children.Append("netconf-protocol", types.YChild{"NetconfProtocol", &allInterfaces.NetconfProtocol})
    allInterfaces.EntityData.Children.Append("xr-xml", types.YChild{"XrXml", &allInterfaces.XrXml})
    allInterfaces.EntityData.Children.Append("ssh-protocol", types.YChild{"SshProtocol", &allInterfaces.SshProtocol})
    allInterfaces.EntityData.Children.Append("snmp-protocol", types.YChild{"SnmpProtocol", &allInterfaces.SnmpProtocol})
    allInterfaces.EntityData.Children.Append("telnet-protocol", types.YChild{"TelnetProtocol", &allInterfaces.TelnetProtocol})
    allInterfaces.EntityData.Children.Append("all-protocols", types.YChild{"AllProtocols", &allInterfaces.AllProtocols})
    allInterfaces.EntityData.Leafs = types.NewOrderedMap()

    allInterfaces.EntityData.YListKeys = []string {}

    return &(allInterfaces.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol
// Configure HTTP on this interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass
}

func (httpProtocol *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol) GetEntityData() *types.CommonEntityData {
    httpProtocol.EntityData.YFilter = httpProtocol.YFilter
    httpProtocol.EntityData.YangName = "http-protocol"
    httpProtocol.EntityData.BundleName = "cisco_ios_xr"
    httpProtocol.EntityData.ParentYangName = "all-interfaces"
    httpProtocol.EntityData.SegmentPath = "http-protocol"
    httpProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    httpProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    httpProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    httpProtocol.EntityData.Children = types.NewOrderedMap()
    httpProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &httpProtocol.PeerClass})
    httpProtocol.EntityData.Leafs = types.NewOrderedMap()

    httpProtocol.EntityData.YListKeys = []string {}

    return &(httpProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "http-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_HttpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol
// Configure TFTP on this interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass
}

func (tftpProtocol *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol) GetEntityData() *types.CommonEntityData {
    tftpProtocol.EntityData.YFilter = tftpProtocol.YFilter
    tftpProtocol.EntityData.YangName = "tftp-protocol"
    tftpProtocol.EntityData.BundleName = "cisco_ios_xr"
    tftpProtocol.EntityData.ParentYangName = "all-interfaces"
    tftpProtocol.EntityData.SegmentPath = "tftp-protocol"
    tftpProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tftpProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tftpProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tftpProtocol.EntityData.Children = types.NewOrderedMap()
    tftpProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &tftpProtocol.PeerClass})
    tftpProtocol.EntityData.Leafs = types.NewOrderedMap()

    tftpProtocol.EntityData.YListKeys = []string {}

    return &(tftpProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "tftp-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TftpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol
// Configure NETCONF protocol and peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass
}

func (netconfProtocol *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol) GetEntityData() *types.CommonEntityData {
    netconfProtocol.EntityData.YFilter = netconfProtocol.YFilter
    netconfProtocol.EntityData.YangName = "netconf-protocol"
    netconfProtocol.EntityData.BundleName = "cisco_ios_xr"
    netconfProtocol.EntityData.ParentYangName = "all-interfaces"
    netconfProtocol.EntityData.SegmentPath = "netconf-protocol"
    netconfProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netconfProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netconfProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netconfProtocol.EntityData.Children = types.NewOrderedMap()
    netconfProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &netconfProtocol.PeerClass})
    netconfProtocol.EntityData.Leafs = types.NewOrderedMap()

    netconfProtocol.EntityData.YListKeys = []string {}

    return &(netconfProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "netconf-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_NetconfProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml
// Configure XML and peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass
}

func (xrXml *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml) GetEntityData() *types.CommonEntityData {
    xrXml.EntityData.YFilter = xrXml.YFilter
    xrXml.EntityData.YangName = "xr-xml"
    xrXml.EntityData.BundleName = "cisco_ios_xr"
    xrXml.EntityData.ParentYangName = "all-interfaces"
    xrXml.EntityData.SegmentPath = "xr-xml"
    xrXml.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xrXml.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xrXml.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xrXml.EntityData.Children = types.NewOrderedMap()
    xrXml.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &xrXml.PeerClass})
    xrXml.EntityData.Leafs = types.NewOrderedMap()

    xrXml.EntityData.YListKeys = []string {}

    return &(xrXml.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "xr-xml"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_XrXml_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol
// Configure SSH protocol and peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass
}

func (sshProtocol *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol) GetEntityData() *types.CommonEntityData {
    sshProtocol.EntityData.YFilter = sshProtocol.YFilter
    sshProtocol.EntityData.YangName = "ssh-protocol"
    sshProtocol.EntityData.BundleName = "cisco_ios_xr"
    sshProtocol.EntityData.ParentYangName = "all-interfaces"
    sshProtocol.EntityData.SegmentPath = "ssh-protocol"
    sshProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sshProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sshProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sshProtocol.EntityData.Children = types.NewOrderedMap()
    sshProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &sshProtocol.PeerClass})
    sshProtocol.EntityData.Leafs = types.NewOrderedMap()

    sshProtocol.EntityData.YListKeys = []string {}

    return &(sshProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "ssh-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SshProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol
// Configure SNMP for this interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass
}

func (snmpProtocol *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol) GetEntityData() *types.CommonEntityData {
    snmpProtocol.EntityData.YFilter = snmpProtocol.YFilter
    snmpProtocol.EntityData.YangName = "snmp-protocol"
    snmpProtocol.EntityData.BundleName = "cisco_ios_xr"
    snmpProtocol.EntityData.ParentYangName = "all-interfaces"
    snmpProtocol.EntityData.SegmentPath = "snmp-protocol"
    snmpProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpProtocol.EntityData.Children = types.NewOrderedMap()
    snmpProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &snmpProtocol.PeerClass})
    snmpProtocol.EntityData.Leafs = types.NewOrderedMap()

    snmpProtocol.EntityData.YListKeys = []string {}

    return &(snmpProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "snmp-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_SnmpProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol
// Configure Telnet for this interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass
}

func (telnetProtocol *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol) GetEntityData() *types.CommonEntityData {
    telnetProtocol.EntityData.YFilter = telnetProtocol.YFilter
    telnetProtocol.EntityData.YangName = "telnet-protocol"
    telnetProtocol.EntityData.BundleName = "cisco_ios_xr"
    telnetProtocol.EntityData.ParentYangName = "all-interfaces"
    telnetProtocol.EntityData.SegmentPath = "telnet-protocol"
    telnetProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    telnetProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    telnetProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    telnetProtocol.EntityData.Children = types.NewOrderedMap()
    telnetProtocol.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &telnetProtocol.PeerClass})
    telnetProtocol.EntityData.Leafs = types.NewOrderedMap()

    telnetProtocol.EntityData.YListKeys = []string {}

    return &(telnetProtocol.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "telnet-protocol"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_TelnetProtocol_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols
// Configure all protocols on this interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    PeerClass ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass
}

func (allProtocols *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols) GetEntityData() *types.CommonEntityData {
    allProtocols.EntityData.YFilter = allProtocols.YFilter
    allProtocols.EntityData.YangName = "all-protocols"
    allProtocols.EntityData.BundleName = "cisco_ios_xr"
    allProtocols.EntityData.ParentYangName = "all-interfaces"
    allProtocols.EntityData.SegmentPath = "all-protocols"
    allProtocols.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allProtocols.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allProtocols.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allProtocols.EntityData.Children = types.NewOrderedMap()
    allProtocols.EntityData.Children.Append("peer-class", types.YChild{"PeerClass", &allProtocols.PeerClass})
    allProtocols.EntityData.Leafs = types.NewOrderedMap()

    allProtocols.EntityData.YListKeys = []string {}

    return &(allProtocols.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Only takes 'True'. The type is interface{}.
    PeerAll interface{}

    // Configure v4 peer addresses.
    PeerV4 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4

    // Configure v6 peer addresses.
    PeerV6 ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6
}

func (peerClass *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass) GetEntityData() *types.CommonEntityData {
    peerClass.EntityData.YFilter = peerClass.YFilter
    peerClass.EntityData.YangName = "peer-class"
    peerClass.EntityData.BundleName = "cisco_ios_xr"
    peerClass.EntityData.ParentYangName = "all-protocols"
    peerClass.EntityData.SegmentPath = "peer-class"
    peerClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerClass.EntityData.Children = types.NewOrderedMap()
    peerClass.EntityData.Children.Append("peer-v4", types.YChild{"PeerV4", &peerClass.PeerV4})
    peerClass.EntityData.Children.Append("peer-v6", types.YChild{"PeerV6", &peerClass.PeerV6})
    peerClass.EntityData.Leafs = types.NewOrderedMap()
    peerClass.EntityData.Leafs.Append("peer-all", types.YLeaf{"PeerAll", peerClass.PeerAll})

    peerClass.EntityData.YListKeys = []string {}

    return &(peerClass.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4
// Configure v4 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes
}

func (peerV4 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4) GetEntityData() *types.CommonEntityData {
    peerV4.EntityData.YFilter = peerV4.YFilter
    peerV4.EntityData.YangName = "peer-v4"
    peerV4.EntityData.BundleName = "cisco_ios_xr"
    peerV4.EntityData.ParentYangName = "peer-class"
    peerV4.EntityData.SegmentPath = "peer-v4"
    peerV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV4.EntityData.Children = types.NewOrderedMap()
    peerV4.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV4.Peers})
    peerV4.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV4.PeerPrefixes})
    peerV4.EntityData.Leafs = types.NewOrderedMap()

    peerV4.EntityData.YListKeys = []string {}

    return &(peerV4.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v4"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v4"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV4_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6
// Configure v6 peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer addresses.
    Peers ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers

    // Configure peer addresses with prefix.
    PeerPrefixes ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes
}

func (peerV6 *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6) GetEntityData() *types.CommonEntityData {
    peerV6.EntityData.YFilter = peerV6.YFilter
    peerV6.EntityData.YangName = "peer-v6"
    peerV6.EntityData.BundleName = "cisco_ios_xr"
    peerV6.EntityData.ParentYangName = "peer-class"
    peerV6.EntityData.SegmentPath = "peer-v6"
    peerV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerV6.EntityData.Children = types.NewOrderedMap()
    peerV6.EntityData.Children.Append("peers", types.YChild{"Peers", &peerV6.Peers})
    peerV6.EntityData.Children.Append("peer-prefixes", types.YChild{"PeerPrefixes", &peerV6.PeerPrefixes})
    peerV6.EntityData.Leafs = types.NewOrderedMap()

    peerV6.EntityData.YListKeys = []string {}

    return &(peerV6.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers
// Configure peer addresses
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure peer on the interface. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers_Peer.
    Peer []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers_Peer
}

func (peers *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "peer-v6"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers_Peer
// Configure peer on the interface
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (peer *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.Address, "address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("address", types.YLeaf{"Address", peer.Address})

    peer.EntityData.YListKeys = []string {"Address"}

    return &(peer.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes
// Configure peer addresses with prefix
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address (with prefix). The type is slice of
    // ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix.
    PeerPrefix []*ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
}

func (peerPrefixes *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes) GetEntityData() *types.CommonEntityData {
    peerPrefixes.EntityData.YFilter = peerPrefixes.YFilter
    peerPrefixes.EntityData.YangName = "peer-prefixes"
    peerPrefixes.EntityData.BundleName = "cisco_ios_xr"
    peerPrefixes.EntityData.ParentYangName = "peer-v6"
    peerPrefixes.EntityData.SegmentPath = "peer-prefixes"
    peerPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefixes.EntityData.Children = types.NewOrderedMap()
    peerPrefixes.EntityData.Children.Append("peer-prefix", types.YChild{"PeerPrefix", nil})
    for i := range peerPrefixes.PeerPrefix {
        peerPrefixes.EntityData.Children.Append(types.GetSegmentPath(peerPrefixes.PeerPrefix[i]), types.YChild{"PeerPrefix", peerPrefixes.PeerPrefix[i]})
    }
    peerPrefixes.EntityData.Leafs = types.NewOrderedMap()

    peerPrefixes.EntityData.YListKeys = []string {}

    return &(peerPrefixes.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix
// Peer address (with prefix)
type ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix/length. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}
}

func (peerPrefix *ControlPlane_ManagementPlaneProtection_Inband_InterfaceSelection_AllInterfaces_AllProtocols_PeerClass_PeerV6_PeerPrefixes_PeerPrefix) GetEntityData() *types.CommonEntityData {
    peerPrefix.EntityData.YFilter = peerPrefix.YFilter
    peerPrefix.EntityData.YangName = "peer-prefix"
    peerPrefix.EntityData.BundleName = "cisco_ios_xr"
    peerPrefix.EntityData.ParentYangName = "peer-prefixes"
    peerPrefix.EntityData.SegmentPath = "peer-prefix" + types.AddKeyToken(peerPrefix.AddressPrefix, "address-prefix")
    peerPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerPrefix.EntityData.Children = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs = types.NewOrderedMap()
    peerPrefix.EntityData.Leafs.Append("address-prefix", types.YLeaf{"AddressPrefix", peerPrefix.AddressPrefix})

    peerPrefix.EntityData.YListKeys = []string {"AddressPrefix"}

    return &(peerPrefix.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Tpa
// MPP Third Party Application Configuration
// Commands
type ControlPlane_ManagementPlaneProtection_Tpa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF configuration.
    Vrfs ControlPlane_ManagementPlaneProtection_Tpa_Vrfs
}

func (tpa *ControlPlane_ManagementPlaneProtection_Tpa) GetEntityData() *types.CommonEntityData {
    tpa.EntityData.YFilter = tpa.YFilter
    tpa.EntityData.YangName = "tpa"
    tpa.EntityData.BundleName = "cisco_ios_xr"
    tpa.EntityData.ParentYangName = "management-plane-protection"
    tpa.EntityData.SegmentPath = "tpa"
    tpa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tpa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tpa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tpa.EntityData.Children = types.NewOrderedMap()
    tpa.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &tpa.Vrfs})
    tpa.EntityData.Leafs = types.NewOrderedMap()

    tpa.EntityData.YListKeys = []string {}

    return &(tpa.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Tpa_Vrfs
// VRF configuration
type ControlPlane_ManagementPlaneProtection_Tpa_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF configuration. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf.
    Vrf []*ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf
}

func (vrfs *ControlPlane_ManagementPlaneProtection_Tpa_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "tpa"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf
// VRF configuration
type ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // Address family.
    AddressFamily ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily
}

func (vrf *ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("address-family", types.YChild{"AddressFamily", &vrf.AddressFamily})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily
// Address family
type ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 configuration.
    Ipv4Table ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv4Table

    // IPv6 configuration.
    Ipv6Table ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv6Table
}

func (addressFamily *ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily) GetEntityData() *types.CommonEntityData {
    addressFamily.EntityData.YFilter = addressFamily.YFilter
    addressFamily.EntityData.YangName = "address-family"
    addressFamily.EntityData.BundleName = "cisco_ios_xr"
    addressFamily.EntityData.ParentYangName = "vrf"
    addressFamily.EntityData.SegmentPath = "address-family"
    addressFamily.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamily.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamily.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamily.EntityData.Children = types.NewOrderedMap()
    addressFamily.EntityData.Children.Append("ipv4-table", types.YChild{"Ipv4Table", &addressFamily.Ipv4Table})
    addressFamily.EntityData.Children.Append("ipv6-table", types.YChild{"Ipv6Table", &addressFamily.Ipv6Table})
    addressFamily.EntityData.Leafs = types.NewOrderedMap()

    addressFamily.EntityData.YListKeys = []string {}

    return &(addressFamily.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv4Table
// IPv4 configuration
type ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv4Table struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPP TPA control entries. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv4Table_TpaAllow.
    TpaAllow []*ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv4Table_TpaAllow
}

func (ipv4Table *ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv4Table) GetEntityData() *types.CommonEntityData {
    ipv4Table.EntityData.YFilter = ipv4Table.YFilter
    ipv4Table.EntityData.YangName = "ipv4-table"
    ipv4Table.EntityData.BundleName = "cisco_ios_xr"
    ipv4Table.EntityData.ParentYangName = "address-family"
    ipv4Table.EntityData.SegmentPath = "ipv4-table"
    ipv4Table.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Table.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Table.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Table.EntityData.Children = types.NewOrderedMap()
    ipv4Table.EntityData.Children.Append("tpa-allow", types.YChild{"TpaAllow", nil})
    for i := range ipv4Table.TpaAllow {
        ipv4Table.EntityData.Children.Append(types.GetSegmentPath(ipv4Table.TpaAllow[i]), types.YChild{"TpaAllow", ipv4Table.TpaAllow[i]})
    }
    ipv4Table.EntityData.Leafs = types.NewOrderedMap()

    ipv4Table.EntityData.YListKeys = []string {}

    return &(ipv4Table.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv4Table_TpaAllow
// MPP TPA control entries
type ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv4Table_TpaAllow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local port value. The type is interface{} with
    // range: 1..65535.
    LocalPort interface{}

    // This attribute is a key. L4 protocol value. The type is interface{} with
    // range: 1..255.
    Protocol interface{}

    // Interface name for allow command, or 'any'. The type is string with length:
    // 1..32. This attribute is mandatory.
    InterfaceName interface{}

    // IPv4/6 remote-address prefix to match, or 'any'. The type is string with
    // length: 1..32. This attribute is mandatory.
    RemoteAddress interface{}

    // IPv4/6 remote-address prefix length. The type is interface{} with range:
    // 0..128.
    RemoteAddressPrefix interface{}

    // IPv4/6 local-address prefix to match, or 'any'. The type is string with
    // length: 1..32. This attribute is mandatory.
    LocalAddress interface{}

    // IPv4/6 local-address prefix length. The type is interface{} with range:
    // 0..128.
    LocalAddressPrefix interface{}
}

func (tpaAllow *ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv4Table_TpaAllow) GetEntityData() *types.CommonEntityData {
    tpaAllow.EntityData.YFilter = tpaAllow.YFilter
    tpaAllow.EntityData.YangName = "tpa-allow"
    tpaAllow.EntityData.BundleName = "cisco_ios_xr"
    tpaAllow.EntityData.ParentYangName = "ipv4-table"
    tpaAllow.EntityData.SegmentPath = "tpa-allow" + types.AddKeyToken(tpaAllow.LocalPort, "local-port") + types.AddKeyToken(tpaAllow.Protocol, "protocol")
    tpaAllow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tpaAllow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tpaAllow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tpaAllow.EntityData.Children = types.NewOrderedMap()
    tpaAllow.EntityData.Leafs = types.NewOrderedMap()
    tpaAllow.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", tpaAllow.LocalPort})
    tpaAllow.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", tpaAllow.Protocol})
    tpaAllow.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", tpaAllow.InterfaceName})
    tpaAllow.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", tpaAllow.RemoteAddress})
    tpaAllow.EntityData.Leafs.Append("remote-address-prefix", types.YLeaf{"RemoteAddressPrefix", tpaAllow.RemoteAddressPrefix})
    tpaAllow.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", tpaAllow.LocalAddress})
    tpaAllow.EntityData.Leafs.Append("local-address-prefix", types.YLeaf{"LocalAddressPrefix", tpaAllow.LocalAddressPrefix})

    tpaAllow.EntityData.YListKeys = []string {"LocalPort", "Protocol"}

    return &(tpaAllow.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv6Table
// IPv6 configuration
type ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv6Table struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPP TPA control entries. The type is slice of
    // ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv6Table_TpaAllow.
    TpaAllow []*ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv6Table_TpaAllow
}

func (ipv6Table *ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv6Table) GetEntityData() *types.CommonEntityData {
    ipv6Table.EntityData.YFilter = ipv6Table.YFilter
    ipv6Table.EntityData.YangName = "ipv6-table"
    ipv6Table.EntityData.BundleName = "cisco_ios_xr"
    ipv6Table.EntityData.ParentYangName = "address-family"
    ipv6Table.EntityData.SegmentPath = "ipv6-table"
    ipv6Table.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Table.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Table.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Table.EntityData.Children = types.NewOrderedMap()
    ipv6Table.EntityData.Children.Append("tpa-allow", types.YChild{"TpaAllow", nil})
    for i := range ipv6Table.TpaAllow {
        ipv6Table.EntityData.Children.Append(types.GetSegmentPath(ipv6Table.TpaAllow[i]), types.YChild{"TpaAllow", ipv6Table.TpaAllow[i]})
    }
    ipv6Table.EntityData.Leafs = types.NewOrderedMap()

    ipv6Table.EntityData.YListKeys = []string {}

    return &(ipv6Table.EntityData)
}

// ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv6Table_TpaAllow
// MPP TPA control entries
type ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv6Table_TpaAllow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local port value. The type is interface{} with
    // range: 1..65535.
    LocalPort interface{}

    // This attribute is a key. L4 protocol value. The type is interface{} with
    // range: 1..255.
    Protocol interface{}

    // Interface name for allow command, or 'any'. The type is string with length:
    // 1..32. This attribute is mandatory.
    InterfaceName interface{}

    // IPv4/6 remote-address prefix to match, or 'any'. The type is string with
    // length: 1..32. This attribute is mandatory.
    RemoteAddress interface{}

    // IPv4/6 remote-address prefix length. The type is interface{} with range:
    // 0..128.
    RemoteAddressPrefix interface{}

    // IPv4/6 local-address prefix to match, or 'any'. The type is string with
    // length: 1..32. This attribute is mandatory.
    LocalAddress interface{}

    // IPv4/6 local-address prefix length. The type is interface{} with range:
    // 0..128.
    LocalAddressPrefix interface{}
}

func (tpaAllow *ControlPlane_ManagementPlaneProtection_Tpa_Vrfs_Vrf_AddressFamily_Ipv6Table_TpaAllow) GetEntityData() *types.CommonEntityData {
    tpaAllow.EntityData.YFilter = tpaAllow.YFilter
    tpaAllow.EntityData.YangName = "tpa-allow"
    tpaAllow.EntityData.BundleName = "cisco_ios_xr"
    tpaAllow.EntityData.ParentYangName = "ipv6-table"
    tpaAllow.EntityData.SegmentPath = "tpa-allow" + types.AddKeyToken(tpaAllow.LocalPort, "local-port") + types.AddKeyToken(tpaAllow.Protocol, "protocol")
    tpaAllow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tpaAllow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tpaAllow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tpaAllow.EntityData.Children = types.NewOrderedMap()
    tpaAllow.EntityData.Leafs = types.NewOrderedMap()
    tpaAllow.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", tpaAllow.LocalPort})
    tpaAllow.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", tpaAllow.Protocol})
    tpaAllow.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", tpaAllow.InterfaceName})
    tpaAllow.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", tpaAllow.RemoteAddress})
    tpaAllow.EntityData.Leafs.Append("remote-address-prefix", types.YLeaf{"RemoteAddressPrefix", tpaAllow.RemoteAddressPrefix})
    tpaAllow.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", tpaAllow.LocalAddress})
    tpaAllow.EntityData.Leafs.Append("local-address-prefix", types.YLeaf{"LocalAddressPrefix", tpaAllow.LocalAddressPrefix})

    tpaAllow.EntityData.YListKeys = []string {"LocalPort", "Protocol"}

    return &(tpaAllow.EntityData)
}

