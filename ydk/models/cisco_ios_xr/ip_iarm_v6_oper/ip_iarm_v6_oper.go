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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default multicast host interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    MulticastHostInterface interface{}

    // IPv6 ARM address database information.
    Addresses Ipv6Arm_Addresses

    // IPv6 ARM summary information.
    Summary Ipv6Arm_Summary

    // IPv6 ARM VRFs summary information.
    VrfSummaries Ipv6Arm_VrfSummaries
}

func (ipv6Arm *Ipv6Arm) GetEntityData() *types.CommonEntityData {
    ipv6Arm.EntityData.YFilter = ipv6Arm.YFilter
    ipv6Arm.EntityData.YangName = "ipv6arm"
    ipv6Arm.EntityData.BundleName = "cisco_ios_xr"
    ipv6Arm.EntityData.ParentYangName = "Cisco-IOS-XR-ip-iarm-v6-oper"
    ipv6Arm.EntityData.SegmentPath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm"
    ipv6Arm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Arm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Arm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Arm.EntityData.Children = make(map[string]types.YChild)
    ipv6Arm.EntityData.Children["addresses"] = types.YChild{"Addresses", &ipv6Arm.Addresses}
    ipv6Arm.EntityData.Children["summary"] = types.YChild{"Summary", &ipv6Arm.Summary}
    ipv6Arm.EntityData.Children["vrf-summaries"] = types.YChild{"VrfSummaries", &ipv6Arm.VrfSummaries}
    ipv6Arm.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Arm.EntityData.Leafs["multicast-host-interface"] = types.YLeaf{"MulticastHostInterface", ipv6Arm.MulticastHostInterface}
    return &(ipv6Arm.EntityData)
}

// Ipv6Arm_Addresses
// IPv6 ARM address database information
type Ipv6Arm_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 ARM address database information per VRF.
    Vrfs Ipv6Arm_Addresses_Vrfs
}

func (addresses *Ipv6Arm_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "ipv6arm"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = make(map[string]types.YChild)
    addresses.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &addresses.Vrfs}
    addresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addresses.EntityData)
}

// Ipv6Arm_Addresses_Vrfs
// IPv6 ARM address database information per VRF
type Ipv6Arm_Addresses_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 ARM address database information in a VRF. The type is slice of
    // Ipv6Arm_Addresses_Vrfs_Vrf.
    Vrf []Ipv6Arm_Addresses_Vrfs_Vrf
}

func (vrfs *Ipv6Arm_Addresses_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "addresses"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// Ipv6Arm_Addresses_Vrfs_Vrf
// IPv6 ARM address database information in a VRF
type Ipv6Arm_Addresses_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string.
    VrfName interface{}

    // IPv6 ARM address database information by network.
    Networks Ipv6Arm_Addresses_Vrfs_Vrf_Networks

    // IPv6 ARM address database information by interface.
    Interfaces Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces
}

func (vrf *Ipv6Arm_Addresses_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["networks"] = types.YChild{"Networks", &vrf.Networks}
    vrf.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &vrf.Interfaces}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// Ipv6Arm_Addresses_Vrfs_Vrf_Networks
// IPv6 ARM address database information by
// network
type Ipv6Arm_Addresses_Vrfs_Vrf_Networks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An IPv6 Address in IPv6 ARM. The type is slice of
    // Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network.
    Network []Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network
}

func (networks *Ipv6Arm_Addresses_Vrfs_Vrf_Networks) GetEntityData() *types.CommonEntityData {
    networks.EntityData.YFilter = networks.YFilter
    networks.EntityData.YangName = "networks"
    networks.EntityData.BundleName = "cisco_ios_xr"
    networks.EntityData.ParentYangName = "vrf"
    networks.EntityData.SegmentPath = "networks"
    networks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networks.EntityData.Children = make(map[string]types.YChild)
    networks.EntityData.Children["network"] = types.YChild{"Network", nil}
    for i := range networks.Network {
        networks.EntityData.Children[types.GetSegmentPath(&networks.Network[i])] = types.YChild{"Network", &networks.Network[i]}
    }
    networks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(networks.EntityData)
}

// Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network
// An IPv6 Address in IPv6 ARM
type Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Prefix Length. The type is interface{} with range: 0..128.
    PrefixLength interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
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

func (network *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xr"
    network.EntityData.ParentYangName = "networks"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Children["address-xr"] = types.YChild{"AddressXr", &network.AddressXr}
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["address"] = types.YLeaf{"Address", network.Address}
    network.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", network.PrefixLength}
    network.EntityData.Leafs["handle"] = types.YLeaf{"Handle", network.Handle}
    network.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", network.InterfaceName}
    network.EntityData.Leafs["referenced-interface"] = types.YLeaf{"ReferencedInterface", network.ReferencedInterface}
    network.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", network.VrfName}
    return &(network.EntityData)
}

// Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr
// Address info
type Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr struct {
    EntityData types.CommonEntityData
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

func (addressXr *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetEntityData() *types.CommonEntityData {
    addressXr.EntityData.YFilter = addressXr.YFilter
    addressXr.EntityData.YangName = "address-xr"
    addressXr.EntityData.BundleName = "cisco_ios_xr"
    addressXr.EntityData.ParentYangName = "network"
    addressXr.EntityData.SegmentPath = "address-xr"
    addressXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressXr.EntityData.Children = make(map[string]types.YChild)
    addressXr.EntityData.Children["address"] = types.YChild{"Address", &addressXr.Address}
    addressXr.EntityData.Leafs = make(map[string]types.YLeaf)
    addressXr.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", addressXr.PrefixLength}
    addressXr.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", addressXr.RouteTag}
    addressXr.EntityData.Leafs["is-primary"] = types.YLeaf{"IsPrimary", addressXr.IsPrimary}
    addressXr.EntityData.Leafs["is-tentative"] = types.YLeaf{"IsTentative", addressXr.IsTentative}
    addressXr.EntityData.Leafs["is-prefix-sid"] = types.YLeaf{"IsPrefixSid", addressXr.IsPrefixSid}
    addressXr.EntityData.Leafs["producer"] = types.YLeaf{"Producer", addressXr.Producer}
    return &(addressXr.EntityData)
}

// Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address
// Address
type Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is interface{} with range: -2147483648..2147483647.
    Afi interface{}

    // IPV4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPV6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "address-xr"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["afi"] = types.YLeaf{"Afi", address.Afi}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", address.Ipv6Address}
    return &(address.EntityData)
}

// Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces
// IPv6 ARM address database information by
// interface
type Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An IPv6 address in IPv6 ARM. The type is slice of
    // Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_.
    Interface_ []Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface
}

func (interfaces *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "vrf"
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

// Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface
// An IPv6 address in IPv6 ARM
type Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Referenced Interface - only valid for an unnumbered interface. The type is
    // string.
    ReferencedInterface interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // Address info. The type is slice of
    // Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address.
    Address []Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address
}

func (self *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface_) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["address"] = types.YChild{"Address", nil}
    for i := range self.Address {
        self.EntityData.Children[types.GetSegmentPath(&self.Address[i])] = types.YChild{"Address", &self.Address[i]}
    }
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
    self.EntityData.Leafs["referenced-interface"] = types.YLeaf{"ReferencedInterface", self.ReferencedInterface}
    self.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", self.VrfName}
    return &(self.EntityData)
}

// Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address
// Address info
type Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address struct {
    EntityData types.CommonEntityData
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
    Address Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address_
}

func (address *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "interface"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Children["address"] = types.YChild{"Address", &address.Address}
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", address.PrefixLength}
    address.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", address.RouteTag}
    address.EntityData.Leafs["is-primary"] = types.YLeaf{"IsPrimary", address.IsPrimary}
    address.EntityData.Leafs["is-tentative"] = types.YLeaf{"IsTentative", address.IsTentative}
    address.EntityData.Leafs["is-prefix-sid"] = types.YLeaf{"IsPrefixSid", address.IsPrefixSid}
    address.EntityData.Leafs["producer"] = types.YLeaf{"Producer", address.Producer}
    return &(address.EntityData)
}

// Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address_
// Address
type Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is interface{} with range: -2147483648..2147483647.
    Afi interface{}

    // IPV4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPV6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (address_ *Ipv6Arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address_) GetEntityData() *types.CommonEntityData {
    address_.EntityData.YFilter = address_.YFilter
    address_.EntityData.YangName = "address"
    address_.EntityData.BundleName = "cisco_ios_xr"
    address_.EntityData.ParentYangName = "address"
    address_.EntityData.SegmentPath = "address"
    address_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address_.EntityData.Children = make(map[string]types.YChild)
    address_.EntityData.Leafs = make(map[string]types.YLeaf)
    address_.EntityData.Leafs["afi"] = types.YLeaf{"Afi", address_.Afi}
    address_.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address_.Ipv4Address}
    address_.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", address_.Ipv6Address}
    return &(address_.EntityData)
}

// Ipv6Arm_Summary
// IPv6 ARM summary information
type Ipv6Arm_Summary struct {
    EntityData types.CommonEntityData
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

func (summary *Ipv6Arm_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "ipv6arm"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["producer-count"] = types.YLeaf{"ProducerCount", summary.ProducerCount}
    summary.EntityData.Leafs["address-conflict-count"] = types.YLeaf{"AddressConflictCount", summary.AddressConflictCount}
    summary.EntityData.Leafs["unnumbered-conflict-count"] = types.YLeaf{"UnnumberedConflictCount", summary.UnnumberedConflictCount}
    summary.EntityData.Leafs["db-master-version"] = types.YLeaf{"DbMasterVersion", summary.DbMasterVersion}
    summary.EntityData.Leafs["vrf-count"] = types.YLeaf{"VrfCount", summary.VrfCount}
    return &(summary.EntityData)
}

// Ipv6Arm_VrfSummaries
// IPv6 ARM VRFs summary information
type Ipv6Arm_VrfSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 ARM VRF summary information. The type is slice of
    // Ipv6Arm_VrfSummaries_VrfSummary.
    VrfSummary []Ipv6Arm_VrfSummaries_VrfSummary
}

func (vrfSummaries *Ipv6Arm_VrfSummaries) GetEntityData() *types.CommonEntityData {
    vrfSummaries.EntityData.YFilter = vrfSummaries.YFilter
    vrfSummaries.EntityData.YangName = "vrf-summaries"
    vrfSummaries.EntityData.BundleName = "cisco_ios_xr"
    vrfSummaries.EntityData.ParentYangName = "ipv6arm"
    vrfSummaries.EntityData.SegmentPath = "vrf-summaries"
    vrfSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummaries.EntityData.Children = make(map[string]types.YChild)
    vrfSummaries.EntityData.Children["vrf-summary"] = types.YChild{"VrfSummary", nil}
    for i := range vrfSummaries.VrfSummary {
        vrfSummaries.EntityData.Children[types.GetSegmentPath(&vrfSummaries.VrfSummary[i])] = types.YChild{"VrfSummary", &vrfSummaries.VrfSummary[i]}
    }
    vrfSummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfSummaries.EntityData)
}

// Ipv6Arm_VrfSummaries_VrfSummary
// IPv6 ARM VRF summary information
type Ipv6Arm_VrfSummaries_VrfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string.
    VrfName interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // VRF Name. The type is string.
    VrfNameXr interface{}
}

func (vrfSummary *Ipv6Arm_VrfSummaries_VrfSummary) GetEntityData() *types.CommonEntityData {
    vrfSummary.EntityData.YFilter = vrfSummary.YFilter
    vrfSummary.EntityData.YangName = "vrf-summary"
    vrfSummary.EntityData.BundleName = "cisco_ios_xr"
    vrfSummary.EntityData.ParentYangName = "vrf-summaries"
    vrfSummary.EntityData.SegmentPath = "vrf-summary" + "[vrf-name='" + fmt.Sprintf("%v", vrfSummary.VrfName) + "']"
    vrfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummary.EntityData.Children = make(map[string]types.YChild)
    vrfSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    vrfSummary.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrfSummary.VrfName}
    vrfSummary.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", vrfSummary.VrfId}
    vrfSummary.EntityData.Leafs["vrf-name-xr"] = types.YLeaf{"VrfNameXr", vrfSummary.VrfNameXr}
    return &(vrfSummary.EntityData)
}

