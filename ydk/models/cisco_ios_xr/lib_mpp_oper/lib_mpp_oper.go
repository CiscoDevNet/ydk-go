// This module contains a collection of YANG definitions
// for Cisco IOS-XR lib-mpp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   management-plane-protection: Management Plane Protection (MPP)
//     operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Management Plane Protection (MPP) outband interface data.
    Outband ManagementPlaneProtection_Outband

    // Management Plane Protection (MPP) inband interface data.
    Inband ManagementPlaneProtection_Inband
}

func (managementPlaneProtection *ManagementPlaneProtection) GetFilter() yfilter.YFilter { return managementPlaneProtection.YFilter }

func (managementPlaneProtection *ManagementPlaneProtection) SetFilter(yf yfilter.YFilter) { managementPlaneProtection.YFilter = yf }

func (managementPlaneProtection *ManagementPlaneProtection) GetGoName(yname string) string {
    if yname == "outband" { return "Outband" }
    if yname == "inband" { return "Inband" }
    return ""
}

func (managementPlaneProtection *ManagementPlaneProtection) GetSegmentPath() string {
    return "Cisco-IOS-XR-lib-mpp-oper:management-plane-protection"
}

func (managementPlaneProtection *ManagementPlaneProtection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "outband" {
        return &managementPlaneProtection.Outband
    }
    if childYangName == "inband" {
        return &managementPlaneProtection.Inband
    }
    return nil
}

func (managementPlaneProtection *ManagementPlaneProtection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["outband"] = &managementPlaneProtection.Outband
    children["inband"] = &managementPlaneProtection.Inband
    return children
}

func (managementPlaneProtection *ManagementPlaneProtection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (managementPlaneProtection *ManagementPlaneProtection) GetBundleName() string { return "cisco_ios_xr" }

func (managementPlaneProtection *ManagementPlaneProtection) GetYangName() string { return "management-plane-protection" }

func (managementPlaneProtection *ManagementPlaneProtection) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (managementPlaneProtection *ManagementPlaneProtection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (managementPlaneProtection *ManagementPlaneProtection) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (managementPlaneProtection *ManagementPlaneProtection) SetParent(parent types.Entity) { managementPlaneProtection.parent = parent }

func (managementPlaneProtection *ManagementPlaneProtection) GetParent() types.Entity { return managementPlaneProtection.parent }

func (managementPlaneProtection *ManagementPlaneProtection) GetParentYangName() string { return "Cisco-IOS-XR-lib-mpp-oper" }

// ManagementPlaneProtection_Outband
// Management Plane Protection (MPP) outband
// interface data
type ManagementPlaneProtection_Outband struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Outband VRF information.
    Vrf ManagementPlaneProtection_Outband_Vrf

    // List of inband/outband interfaces.
    Interfaces ManagementPlaneProtection_Outband_Interfaces
}

func (outband *ManagementPlaneProtection_Outband) GetFilter() yfilter.YFilter { return outband.YFilter }

func (outband *ManagementPlaneProtection_Outband) SetFilter(yf yfilter.YFilter) { outband.YFilter = yf }

func (outband *ManagementPlaneProtection_Outband) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (outband *ManagementPlaneProtection_Outband) GetSegmentPath() string {
    return "outband"
}

func (outband *ManagementPlaneProtection_Outband) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        return &outband.Vrf
    }
    if childYangName == "interfaces" {
        return &outband.Interfaces
    }
    return nil
}

func (outband *ManagementPlaneProtection_Outband) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrf"] = &outband.Vrf
    children["interfaces"] = &outband.Interfaces
    return children
}

func (outband *ManagementPlaneProtection_Outband) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (outband *ManagementPlaneProtection_Outband) GetBundleName() string { return "cisco_ios_xr" }

func (outband *ManagementPlaneProtection_Outband) GetYangName() string { return "outband" }

func (outband *ManagementPlaneProtection_Outband) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outband *ManagementPlaneProtection_Outband) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outband *ManagementPlaneProtection_Outband) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outband *ManagementPlaneProtection_Outband) SetParent(parent types.Entity) { outband.parent = parent }

func (outband *ManagementPlaneProtection_Outband) GetParent() types.Entity { return outband.parent }

func (outband *ManagementPlaneProtection_Outband) GetParentYangName() string { return "management-plane-protection" }

// ManagementPlaneProtection_Outband_Vrf
// Outband VRF information
type ManagementPlaneProtection_Outband_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Outband VRF name. The type is string.
    VrfName interface{}
}

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *ManagementPlaneProtection_Outband_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetSegmentPath() string {
    return "vrf"
}

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetYangName() string { return "vrf" }

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *ManagementPlaneProtection_Outband_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetParentYangName() string { return "outband" }

// ManagementPlaneProtection_Outband_Interfaces
// List of inband/outband interfaces
type ManagementPlaneProtection_Outband_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPP interface information. The type is slice of
    // ManagementPlaneProtection_Outband_Interfaces_Interface.
    Interface []ManagementPlaneProtection_Outband_Interfaces_Interface
}

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ManagementPlaneProtection_Outband_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetParentYangName() string { return "outband" }

// ManagementPlaneProtection_Outband_Interfaces_Interface
// MPP interface information
type ManagementPlaneProtection_Outband_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name, specify 'all' for all interfaces.
    // The type is string.
    InterfaceName interface{}

    // MPP Interface protocols. The type is slice of
    // ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol.
    Protocol []ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol
}

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "protocol" { return "Protocol" }
    return ""
}

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocol" {
        for _, c := range self.Protocol {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol{}
        self.Protocol = append(self.Protocol, child)
        return &self.Protocol[len(self.Protocol)-1]
    }
    return nil
}

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range self.Protocol {
        children[self.Protocol[i].GetSegmentPath()] = &self.Protocol[i]
    }
    return children
}

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol
// MPP Interface protocols
type ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPP allow. The type is MppAllow.
    Allow interface{}

    // If TRUE, all peers are allowed. The type is bool.
    IsAllPeersAllowed interface{}

    // List of peer addresses. The type is slice of
    // ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress.
    PeerAddress []ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress
}

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetFilter() yfilter.YFilter { return protocol.YFilter }

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) SetFilter(yf yfilter.YFilter) { protocol.YFilter = yf }

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetGoName(yname string) string {
    if yname == "allow" { return "Allow" }
    if yname == "is-all-peers-allowed" { return "IsAllPeersAllowed" }
    if yname == "peer-address" { return "PeerAddress" }
    return ""
}

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetSegmentPath() string {
    return "protocol"
}

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-address" {
        for _, c := range protocol.PeerAddress {
            if protocol.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress{}
        protocol.PeerAddress = append(protocol.PeerAddress, child)
        return &protocol.PeerAddress[len(protocol.PeerAddress)-1]
    }
    return nil
}

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range protocol.PeerAddress {
        children[protocol.PeerAddress[i].GetSegmentPath()] = &protocol.PeerAddress[i]
    }
    return children
}

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow"] = protocol.Allow
    leafs["is-all-peers-allowed"] = protocol.IsAllPeersAllowed
    return leafs
}

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetYangName() string { return "protocol" }

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) SetParent(parent types.Entity) { protocol.parent = parent }

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetParent() types.Entity { return protocol.parent }

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetParentYangName() string { return "interface" }

// ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress
// List of peer addresses
type ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress struct {
    parent types.Entity
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

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetFilter() yfilter.YFilter { return peerAddress.YFilter }

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) SetFilter(yf yfilter.YFilter) { peerAddress.YFilter = yf }

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetSegmentPath() string {
    return "peer-address"
}

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = peerAddress.AfName
    leafs["ipv4-address"] = peerAddress.Ipv4Address
    leafs["ipv6-address"] = peerAddress.Ipv6Address
    return leafs
}

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetBundleName() string { return "cisco_ios_xr" }

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetYangName() string { return "peer-address" }

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) SetParent(parent types.Entity) { peerAddress.parent = parent }

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetParent() types.Entity { return peerAddress.parent }

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetParentYangName() string { return "protocol" }

// ManagementPlaneProtection_Inband
// Management Plane Protection (MPP) inband
// interface data
type ManagementPlaneProtection_Inband struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of inband/outband interfaces.
    Interfaces ManagementPlaneProtection_Inband_Interfaces
}

func (inband *ManagementPlaneProtection_Inband) GetFilter() yfilter.YFilter { return inband.YFilter }

func (inband *ManagementPlaneProtection_Inband) SetFilter(yf yfilter.YFilter) { inband.YFilter = yf }

func (inband *ManagementPlaneProtection_Inband) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (inband *ManagementPlaneProtection_Inband) GetSegmentPath() string {
    return "inband"
}

func (inband *ManagementPlaneProtection_Inband) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &inband.Interfaces
    }
    return nil
}

func (inband *ManagementPlaneProtection_Inband) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &inband.Interfaces
    return children
}

func (inband *ManagementPlaneProtection_Inband) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inband *ManagementPlaneProtection_Inband) GetBundleName() string { return "cisco_ios_xr" }

func (inband *ManagementPlaneProtection_Inband) GetYangName() string { return "inband" }

func (inband *ManagementPlaneProtection_Inband) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inband *ManagementPlaneProtection_Inband) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inband *ManagementPlaneProtection_Inband) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inband *ManagementPlaneProtection_Inband) SetParent(parent types.Entity) { inband.parent = parent }

func (inband *ManagementPlaneProtection_Inband) GetParent() types.Entity { return inband.parent }

func (inband *ManagementPlaneProtection_Inband) GetParentYangName() string { return "management-plane-protection" }

// ManagementPlaneProtection_Inband_Interfaces
// List of inband/outband interfaces
type ManagementPlaneProtection_Inband_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPP interface information. The type is slice of
    // ManagementPlaneProtection_Inband_Interfaces_Interface.
    Interface []ManagementPlaneProtection_Inband_Interfaces_Interface
}

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ManagementPlaneProtection_Inband_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetParentYangName() string { return "inband" }

// ManagementPlaneProtection_Inband_Interfaces_Interface
// MPP interface information
type ManagementPlaneProtection_Inband_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name, specify 'all' for all interfaces.
    // The type is string.
    InterfaceName interface{}

    // MPP Interface protocols. The type is slice of
    // ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol.
    Protocol []ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol
}

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "protocol" { return "Protocol" }
    return ""
}

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocol" {
        for _, c := range self.Protocol {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol{}
        self.Protocol = append(self.Protocol, child)
        return &self.Protocol[len(self.Protocol)-1]
    }
    return nil
}

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range self.Protocol {
        children[self.Protocol[i].GetSegmentPath()] = &self.Protocol[i]
    }
    return children
}

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol
// MPP Interface protocols
type ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPP allow. The type is MppAllow.
    Allow interface{}

    // If TRUE, all peers are allowed. The type is bool.
    IsAllPeersAllowed interface{}

    // List of peer addresses. The type is slice of
    // ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress.
    PeerAddress []ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress
}

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetFilter() yfilter.YFilter { return protocol.YFilter }

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) SetFilter(yf yfilter.YFilter) { protocol.YFilter = yf }

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetGoName(yname string) string {
    if yname == "allow" { return "Allow" }
    if yname == "is-all-peers-allowed" { return "IsAllPeersAllowed" }
    if yname == "peer-address" { return "PeerAddress" }
    return ""
}

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetSegmentPath() string {
    return "protocol"
}

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-address" {
        for _, c := range protocol.PeerAddress {
            if protocol.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress{}
        protocol.PeerAddress = append(protocol.PeerAddress, child)
        return &protocol.PeerAddress[len(protocol.PeerAddress)-1]
    }
    return nil
}

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range protocol.PeerAddress {
        children[protocol.PeerAddress[i].GetSegmentPath()] = &protocol.PeerAddress[i]
    }
    return children
}

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow"] = protocol.Allow
    leafs["is-all-peers-allowed"] = protocol.IsAllPeersAllowed
    return leafs
}

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetYangName() string { return "protocol" }

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) SetParent(parent types.Entity) { protocol.parent = parent }

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetParent() types.Entity { return protocol.parent }

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetParentYangName() string { return "interface" }

// ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress
// List of peer addresses
type ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress struct {
    parent types.Entity
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

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetFilter() yfilter.YFilter { return peerAddress.YFilter }

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) SetFilter(yf yfilter.YFilter) { peerAddress.YFilter = yf }

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetSegmentPath() string {
    return "peer-address"
}

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = peerAddress.AfName
    leafs["ipv4-address"] = peerAddress.Ipv4Address
    leafs["ipv6-address"] = peerAddress.Ipv6Address
    return leafs
}

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetBundleName() string { return "cisco_ios_xr" }

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetYangName() string { return "peer-address" }

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) SetParent(parent types.Entity) { peerAddress.parent = parent }

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetParent() types.Entity { return peerAddress.parent }

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetParentYangName() string { return "protocol" }

