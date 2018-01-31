// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-iarm-v6 package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ipv6arm: IPv6 Address Repository Manager (IPv6 ARM)
//     operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_iarm_v6_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_iarm_v6_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-iarm-v6-oper ipv6arm}", reflect.TypeOf(Ipv6Arm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm", reflect.TypeOf(Ipv6Arm{}))
}

// Ipv6Arm
// IPv6 Address Repository Manager (IPv6 ARM)
// operational data
type Ipv6Arm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default multicast host interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    MulticastHostInterface interface{}

    // IPv6 ARM address database information.
    Addresses Ipv6Arm_Addresses

    // IPv6 ARM summary information.
    Summary Ipv6Arm_Summary

    // IPv6 ARM VRFs summary information.
    VrfSummaries Ipv6Arm_VrfSummaries
}

func (ipv6Arm *Ipv6Arm) GetFilter() yfilter.YFilter { return ipv6Arm.YFilter }

func (ipv6Arm *Ipv6Arm) SetFilter(yf yfilter.YFilter) { ipv6Arm.YFilter = yf }

func (ipv6Arm *Ipv6Arm) GetGoName(yname string) string {
    if yname == "multicast-host-interface" { return "MulticastHostInterface" }
    if yname == "addresses" { return "Addresses" }
    if yname == "summary" { return "Summary" }
    if yname == "vrf-summaries" { return "VrfSummaries" }
    return ""
}

func (ipv6Arm *Ipv6Arm) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm"
}

func (ipv6Arm *Ipv6Arm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addresses" {
        return &ipv6Arm.Addresses
    }
    if childYangName == "summary" {
        return &ipv6Arm.Summary
    }
    if childYangName == "vrf-summaries" {
        return &ipv6Arm.VrfSummaries
    }
    return nil
}

func (ipv6Arm *Ipv6Arm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addresses"] = &ipv6Arm.Addresses
    children["summary"] = &ipv6Arm.Summary
    children["vrf-summaries"] = &ipv6Arm.VrfSummaries
    return children
}

func (ipv6Arm *Ipv6Arm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["multicast-host-interface"] = ipv6Arm.MulticastHostInterface
    return leafs
}

func (ipv6Arm *Ipv6Arm) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Arm *Ipv6Arm) GetYangName() string { return "ipv6arm" }

func (ipv6Arm *Ipv6Arm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Arm *Ipv6Arm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Arm *Ipv6Arm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Arm *Ipv6Arm) SetParent(parent types.Entity) { ipv6Arm.parent = parent }

func (ipv6Arm *Ipv6Arm) GetParent() types.Entity { return ipv6Arm.parent }

func (ipv6Arm *Ipv6Arm) GetParentYangName() string { return "Cisco-IOS-XR-ip-iarm-v6-oper" }

// Ipv6Arm_Addresses
// IPv6 ARM address database information
type Ipv6Arm_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 ARM address database information per VRF.
    Vrfs Ipv6Arm_Addresses_Vrfs
}

func (addresses *Ipv6Arm_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *Ipv6Arm_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *Ipv6Arm_Addresses) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (addresses *Ipv6Arm_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *Ipv6Arm_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &addresses.Vrfs
    }
    return nil
}

func (addresses *Ipv6Arm_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &addresses.Vrfs
    return children
}

func (addresses *Ipv6Arm_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *Ipv6Arm_Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (addresses *Ipv6Arm_Addresses) GetYangName() string { return "addresses" }

func (addresses *Ipv6Arm_Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addresses *Ipv6Arm_Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addresses *Ipv6Arm_Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addresses *Ipv6Arm_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *Ipv6Arm_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *Ipv6Arm_Addresses) GetParentYangName() string { return "ipv6arm" }

// Ipv6Arm_Addresses_Vrfs
// IPv6 ARM address database information per VRF
type Ipv6Arm_Addresses_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 ARM address database information in a VRF. The type is slice of
    // Ipv6Arm_Addresses_Vrfs_Vrf.
    Vrf []Ipv6Arm_Addresses_Vrfs_Vrf
}

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv6Arm_Addresses_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Arm_Addresses_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv6Arm_Addresses_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetParentYangName() string { return "addresses" }

// Ipv6Arm_Addresses_Vrfs_Vrf
// IPv6 ARM address database information in a VRF
type Ipv6Arm_Addresses_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string.
    VrfName interface{}

    // IPv6 ARM address database information by network.
    Networks Ipv6Arm_Addresses_Vrfs_Vrf_Networks

    // IPv6 ARM address database information by interface.
    Interfaces Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces
}

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "networks" { return "Networks" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "networks" {
        return &vrf.Networks
    }
    if childYangName == "interfaces" {
        return &vrf.Interfaces
    }
    return nil
}

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["networks"] = &vrf.Networks
    children["interfaces"] = &vrf.Interfaces
    return children
}

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv6Arm_Addresses_Vrfs_Vrf_Networks
// IPv6 ARM address database information by
// network
type Ipv6Arm_Addresses_Vrfs_Vrf_Networks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An IPv6 Address in IPv6 ARM. The type is slice of
    // Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network.
    Network []Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network
}

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetFilter() yfilter.YFilter { return networks.YFilter }

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) SetFilter(yf yfilter.YFilter) { networks.YFilter = yf }

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetGoName(yname string) string {
    if yname == "network" { return "Network" }
    return ""
}

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetSegmentPath() string {
    return "networks"
}

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network" {
        for _, c := range networks.Network {
            if networks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network{}
        networks.Network = append(networks.Network, child)
        return &networks.Network[len(networks.Network)-1]
    }
    return nil
}

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networks.Network {
        children[networks.Network[i].GetSegmentPath()] = &networks.Network[i]
    }
    return children
}

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetBundleName() string { return "cisco_ios_xr" }

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetYangName() string { return "networks" }

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) SetParent(parent types.Entity) { networks.parent = parent }

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetParent() types.Entity { return networks.parent }

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetParentYangName() string { return "vrf" }

// Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network
// An IPv6 Address in IPv6 ARM
type Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Prefix Length. The type is interface{} with range: 0..128.
    PrefixLength interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Handle interface{}

    // Interface name. The type is string.
    InterfaceName interface{}

    // Referenced Interface - only valid for an unnumbered interface. The type is
    // string.
    ReferencedInterface interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // Address info.
    AddressXr Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr
}

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetFilter() yfilter.YFilter { return network.YFilter }

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) SetFilter(yf yfilter.YFilter) { network.YFilter = yf }

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "handle" { return "Handle" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "referenced-interface" { return "ReferencedInterface" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "address-xr" { return "AddressXr" }
    return ""
}

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetSegmentPath() string {
    return "network"
}

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-xr" {
        return &network.AddressXr
    }
    return nil
}

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address-xr"] = &network.AddressXr
    return children
}

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = network.Address
    leafs["prefix-length"] = network.PrefixLength
    leafs["handle"] = network.Handle
    leafs["interface-name"] = network.InterfaceName
    leafs["referenced-interface"] = network.ReferencedInterface
    leafs["vrf-name"] = network.VrfName
    return leafs
}

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetBundleName() string { return "cisco_ios_xr" }

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetYangName() string { return "network" }

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) SetParent(parent types.Entity) { network.parent = parent }

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetParent() types.Entity { return network.parent }

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetParentYangName() string { return "networks" }

// Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr
// Address info
type Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // Route Tag of the address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}

    // Is address primary - valid only for IPv4 addresses. The type is bool.
    IsPrimary interface{}

    // Is address valid/tentative - valid only for IPV6 addresses. The type is
    // bool.
    IsTentative interface{}

    // Is prefix_sid valid - valid only for IPV6 addresses. The type is bool.
    IsPrefixSid interface{}

    // Producer Name. The type is string.
    Producer interface{}

    // Address.
    Address Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address
}

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetFilter() yfilter.YFilter { return addressXr.YFilter }

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) SetFilter(yf yfilter.YFilter) { addressXr.YFilter = yf }

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "route-tag" { return "RouteTag" }
    if yname == "is-primary" { return "IsPrimary" }
    if yname == "is-tentative" { return "IsTentative" }
    if yname == "is-prefix-sid" { return "IsPrefixSid" }
    if yname == "producer" { return "Producer" }
    if yname == "address" { return "Address" }
    return ""
}

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetSegmentPath() string {
    return "address-xr"
}

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &addressXr.Address
    }
    return nil
}

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &addressXr.Address
    return children
}

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = addressXr.PrefixLength
    leafs["route-tag"] = addressXr.RouteTag
    leafs["is-primary"] = addressXr.IsPrimary
    leafs["is-tentative"] = addressXr.IsTentative
    leafs["is-prefix-sid"] = addressXr.IsPrefixSid
    leafs["producer"] = addressXr.Producer
    return leafs
}

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetBundleName() string { return "cisco_ios_xr" }

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetYangName() string { return "address-xr" }

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) SetParent(parent types.Entity) { addressXr.parent = parent }

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetParent() types.Entity { return addressXr.parent }

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetParentYangName() string { return "network" }

// Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address
// Address
type Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is interface{} with range: -2147483648..2147483647.
    Afi interface{}

    // IPV4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPV6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = address.Afi
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetYangName() string { return "address" }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetParent() types.Entity { return address.parent }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetParentYangName() string { return "address-xr" }

// Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces
// IPv6 ARM address database information by
// interface
type Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An IPv6 address in IPv6 ARM. The type is slice of
    // Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface.
    Interface []Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface
}

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetParentYangName() string { return "vrf" }

// Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface
// An IPv6 address in IPv6 ARM
type Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}

    // Referenced Interface - only valid for an unnumbered interface. The type is
    // string.
    ReferencedInterface interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // Address info. The type is slice of
    // Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address.
    Address []Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address
}

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "referenced-interface" { return "ReferencedInterface" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "address" { return "Address" }
    return ""
}

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface) + "']"
}

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range self.Address {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address{}
        self.Address = append(self.Address, child)
        return &self.Address[len(self.Address)-1]
    }
    return nil
}

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range self.Address {
        children[self.Address[i].GetSegmentPath()] = &self.Address[i]
    }
    return children
}

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = self.Interface
    leafs["referenced-interface"] = self.ReferencedInterface
    leafs["vrf-name"] = self.VrfName
    return leafs
}

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address
// Address info
type Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // Route Tag of the address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}

    // Is address primary - valid only for IPv4 addresses. The type is bool.
    IsPrimary interface{}

    // Is address valid/tentative - valid only for IPV6 addresses. The type is
    // bool.
    IsTentative interface{}

    // Is prefix_sid valid - valid only for IPV6 addresses. The type is bool.
    IsPrefixSid interface{}

    // Producer Name. The type is string.
    Producer interface{}

    // Address.
    Address Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "route-tag" { return "RouteTag" }
    if yname == "is-primary" { return "IsPrimary" }
    if yname == "is-tentative" { return "IsTentative" }
    if yname == "is-prefix-sid" { return "IsPrefixSid" }
    if yname == "producer" { return "Producer" }
    if yname == "address" { return "Address" }
    return ""
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &address.Address
    }
    return nil
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &address.Address
    return children
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = address.PrefixLength
    leafs["route-tag"] = address.RouteTag
    leafs["is-primary"] = address.IsPrimary
    leafs["is-tentative"] = address.IsTentative
    leafs["is-prefix-sid"] = address.IsPrefixSid
    leafs["producer"] = address.Producer
    return leafs
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetYangName() string { return "address" }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetParent() types.Entity { return address.parent }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetParentYangName() string { return "interface" }

// Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address
// Address
type Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is interface{} with range: -2147483648..2147483647.
    Afi interface{}

    // IPV4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPV6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = address.Afi
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetYangName() string { return "address" }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetParent() types.Entity { return address.parent }

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetParentYangName() string { return "address" }

// Ipv6Arm_Summary
// IPv6 ARM summary information
type Ipv6Arm_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of producers. The type is interface{} with range:
    // -2147483648..2147483647.
    ProducerCount interface{}

    // Number of address conflicts. The type is interface{} with range:
    // -2147483648..2147483647.
    AddressConflictCount interface{}

    // Number of unnumbered interface conflicts. The type is interface{} with
    // range: -2147483648..2147483647.
    UnnumberedConflictCount interface{}

    // IP-ARM DB master version. The type is interface{} with range:
    // 0..4294967295.
    DbMasterVersion interface{}

    // Number of known VRFs. The type is interface{} with range:
    // -2147483648..2147483647.
    VrfCount interface{}
}

func (summary *Ipv6Arm_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Ipv6Arm_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Ipv6Arm_Summary) GetGoName(yname string) string {
    if yname == "producer-count" { return "ProducerCount" }
    if yname == "address-conflict-count" { return "AddressConflictCount" }
    if yname == "unnumbered-conflict-count" { return "UnnumberedConflictCount" }
    if yname == "db-master-version" { return "DbMasterVersion" }
    if yname == "vrf-count" { return "VrfCount" }
    return ""
}

func (summary *Ipv6Arm_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Ipv6Arm_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Ipv6Arm_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Ipv6Arm_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["producer-count"] = summary.ProducerCount
    leafs["address-conflict-count"] = summary.AddressConflictCount
    leafs["unnumbered-conflict-count"] = summary.UnnumberedConflictCount
    leafs["db-master-version"] = summary.DbMasterVersion
    leafs["vrf-count"] = summary.VrfCount
    return leafs
}

func (summary *Ipv6Arm_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Ipv6Arm_Summary) GetYangName() string { return "summary" }

func (summary *Ipv6Arm_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Ipv6Arm_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Ipv6Arm_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Ipv6Arm_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Ipv6Arm_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Ipv6Arm_Summary) GetParentYangName() string { return "ipv6arm" }

// Ipv6Arm_VrfSummaries
// IPv6 ARM VRFs summary information
type Ipv6Arm_VrfSummaries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 ARM VRF summary information. The type is slice of
    // Ipv6Arm_VrfSummaries_VrfSummary.
    VrfSummary []Ipv6Arm_VrfSummaries_VrfSummary
}

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetFilter() yfilter.YFilter { return vrfSummaries.YFilter }

func (vrfSummaries *Ipv6Arm_VrfSummaries) SetFilter(yf yfilter.YFilter) { vrfSummaries.YFilter = yf }

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetGoName(yname string) string {
    if yname == "vrf-summary" { return "VrfSummary" }
    return ""
}

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetSegmentPath() string {
    return "vrf-summaries"
}

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-summary" {
        for _, c := range vrfSummaries.VrfSummary {
            if vrfSummaries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Arm_VrfSummaries_VrfSummary{}
        vrfSummaries.VrfSummary = append(vrfSummaries.VrfSummary, child)
        return &vrfSummaries.VrfSummary[len(vrfSummaries.VrfSummary)-1]
    }
    return nil
}

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfSummaries.VrfSummary {
        children[vrfSummaries.VrfSummary[i].GetSegmentPath()] = &vrfSummaries.VrfSummary[i]
    }
    return children
}

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetBundleName() string { return "cisco_ios_xr" }

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetYangName() string { return "vrf-summaries" }

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfSummaries *Ipv6Arm_VrfSummaries) SetParent(parent types.Entity) { vrfSummaries.parent = parent }

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetParent() types.Entity { return vrfSummaries.parent }

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetParentYangName() string { return "ipv6arm" }

// Ipv6Arm_VrfSummaries_VrfSummary
// IPv6 ARM VRF summary information
type Ipv6Arm_VrfSummaries_VrfSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string.
    VrfName interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // VRF Name. The type is string.
    VrfNameXr interface{}
}

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetFilter() yfilter.YFilter { return vrfSummary.YFilter }

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) SetFilter(yf yfilter.YFilter) { vrfSummary.YFilter = yf }

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "vrf-name-xr" { return "VrfNameXr" }
    return ""
}

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetSegmentPath() string {
    return "vrf-summary" + "[vrf-name='" + fmt.Sprintf("%v", vrfSummary.VrfName) + "']"
}

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrfSummary.VrfName
    leafs["vrf-id"] = vrfSummary.VrfId
    leafs["vrf-name-xr"] = vrfSummary.VrfNameXr
    return leafs
}

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetBundleName() string { return "cisco_ios_xr" }

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetYangName() string { return "vrf-summary" }

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) SetParent(parent types.Entity) { vrfSummary.parent = parent }

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetParent() types.Entity { return vrfSummary.parent }

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetParentYangName() string { return "vrf-summaries" }

