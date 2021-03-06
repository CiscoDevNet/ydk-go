// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-iarm-v6 package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ipv6arm: IPv6 Address Repository Manager (IPv6 ARM)
//     operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-iarm-v6-oper ipv6arm}", reflect.TypeOf(Ipv6arm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm", reflect.TypeOf(Ipv6arm{}))
}

// Ipv6arm
// IPv6 Address Repository Manager (IPv6 ARM)
// operational data
type Ipv6arm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default multicast host interface. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    MulticastHostInterface interface{}

    // IPv6 ARM address database information.
    Addresses Ipv6arm_Addresses

    // IPv6 ARM summary information.
    Summary Ipv6arm_Summary

    // IPv6 ARM VRFs summary information.
    VrfSummaries Ipv6arm_VrfSummaries
}

func (ipv6arm *Ipv6arm) GetEntityData() *types.CommonEntityData {
    ipv6arm.EntityData.YFilter = ipv6arm.YFilter
    ipv6arm.EntityData.YangName = "ipv6arm"
    ipv6arm.EntityData.BundleName = "cisco_ios_xr"
    ipv6arm.EntityData.ParentYangName = "Cisco-IOS-XR-ip-iarm-v6-oper"
    ipv6arm.EntityData.SegmentPath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm"
    ipv6arm.EntityData.AbsolutePath = ipv6arm.EntityData.SegmentPath
    ipv6arm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6arm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6arm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6arm.EntityData.Children = types.NewOrderedMap()
    ipv6arm.EntityData.Children.Append("addresses", types.YChild{"Addresses", &ipv6arm.Addresses})
    ipv6arm.EntityData.Children.Append("summary", types.YChild{"Summary", &ipv6arm.Summary})
    ipv6arm.EntityData.Children.Append("vrf-summaries", types.YChild{"VrfSummaries", &ipv6arm.VrfSummaries})
    ipv6arm.EntityData.Leafs = types.NewOrderedMap()
    ipv6arm.EntityData.Leafs.Append("multicast-host-interface", types.YLeaf{"MulticastHostInterface", ipv6arm.MulticastHostInterface})

    ipv6arm.EntityData.YListKeys = []string {}

    return &(ipv6arm.EntityData)
}

// Ipv6arm_Addresses
// IPv6 ARM address database information
type Ipv6arm_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 ARM address database information per VRF.
    Vrfs Ipv6arm_Addresses_Vrfs
}

func (addresses *Ipv6arm_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "ipv6arm"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/" + addresses.EntityData.SegmentPath
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &addresses.Vrfs})
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// Ipv6arm_Addresses_Vrfs
// IPv6 ARM address database information per VRF
type Ipv6arm_Addresses_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 ARM address database information in a VRF. The type is slice of
    // Ipv6arm_Addresses_Vrfs_Vrf.
    Vrf []*Ipv6arm_Addresses_Vrfs_Vrf
}

func (vrfs *Ipv6arm_Addresses_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "addresses"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/addresses/" + vrfs.EntityData.SegmentPath
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

// Ipv6arm_Addresses_Vrfs_Vrf
// IPv6 ARM address database information in a VRF
type Ipv6arm_Addresses_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string.
    VrfName interface{}

    // IPv6 ARM address database information by network.
    Networks Ipv6arm_Addresses_Vrfs_Vrf_Networks

    // IPv6 ARM address database information by interface.
    Interfaces Ipv6arm_Addresses_Vrfs_Vrf_Interfaces
}

func (vrf *Ipv6arm_Addresses_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/addresses/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("networks", types.YChild{"Networks", &vrf.Networks})
    vrf.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &vrf.Interfaces})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Ipv6arm_Addresses_Vrfs_Vrf_Networks
// IPv6 ARM address database information by
// network
type Ipv6arm_Addresses_Vrfs_Vrf_Networks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An IPv6 Address in IPv6 ARM. The type is slice of
    // Ipv6arm_Addresses_Vrfs_Vrf_Networks_Network.
    Network []*Ipv6arm_Addresses_Vrfs_Vrf_Networks_Network
}

func (networks *Ipv6arm_Addresses_Vrfs_Vrf_Networks) GetEntityData() *types.CommonEntityData {
    networks.EntityData.YFilter = networks.YFilter
    networks.EntityData.YangName = "networks"
    networks.EntityData.BundleName = "cisco_ios_xr"
    networks.EntityData.ParentYangName = "vrf"
    networks.EntityData.SegmentPath = "networks"
    networks.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/addresses/vrfs/vrf/" + networks.EntityData.SegmentPath
    networks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networks.EntityData.Children = types.NewOrderedMap()
    networks.EntityData.Children.Append("network", types.YChild{"Network", nil})
    for i := range networks.Network {
        types.SetYListKey(networks.Network[i], i)
        networks.EntityData.Children.Append(types.GetSegmentPath(networks.Network[i]), types.YChild{"Network", networks.Network[i]})
    }
    networks.EntityData.Leafs = types.NewOrderedMap()

    networks.EntityData.YListKeys = []string {}

    return &(networks.EntityData)
}

// Ipv6arm_Addresses_Vrfs_Vrf_Networks_Network
// An IPv6 Address in IPv6 ARM
type Ipv6arm_Addresses_Vrfs_Vrf_Networks_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Prefix Length. The type is interface{} with range: 0..128.
    PrefixLength interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    Interface interface{}

    // Interface name. The type is string.
    InterfaceName interface{}

    // Referenced Interface - only valid for an unnumbered interface. The type is
    // string.
    ReferencedInterface interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // Address info.
    AddressXr Ipv6arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr
}

func (network *Ipv6arm_Addresses_Vrfs_Vrf_Networks_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xr"
    network.EntityData.ParentYangName = "networks"
    network.EntityData.SegmentPath = "network" + types.AddNoKeyToken(network)
    network.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/addresses/vrfs/vrf/networks/" + network.EntityData.SegmentPath
    network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Children.Append("address-xr", types.YChild{"AddressXr", &network.AddressXr})
    network.EntityData.Leafs = types.NewOrderedMap()
    network.EntityData.Leafs.Append("address", types.YLeaf{"Address", network.Address})
    network.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", network.PrefixLength})
    network.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", network.Interface})
    network.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", network.InterfaceName})
    network.EntityData.Leafs.Append("referenced-interface", types.YLeaf{"ReferencedInterface", network.ReferencedInterface})
    network.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", network.VrfName})

    network.EntityData.YListKeys = []string {}

    return &(network.EntityData)
}

// Ipv6arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr
// Address info
type Ipv6arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr struct {
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
    Address Ipv6arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address
}

func (addressXr *Ipv6arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr) GetEntityData() *types.CommonEntityData {
    addressXr.EntityData.YFilter = addressXr.YFilter
    addressXr.EntityData.YangName = "address-xr"
    addressXr.EntityData.BundleName = "cisco_ios_xr"
    addressXr.EntityData.ParentYangName = "network"
    addressXr.EntityData.SegmentPath = "address-xr"
    addressXr.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/addresses/vrfs/vrf/networks/network/" + addressXr.EntityData.SegmentPath
    addressXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressXr.EntityData.Children = types.NewOrderedMap()
    addressXr.EntityData.Children.Append("address", types.YChild{"Address", &addressXr.Address})
    addressXr.EntityData.Leafs = types.NewOrderedMap()
    addressXr.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", addressXr.PrefixLength})
    addressXr.EntityData.Leafs.Append("route-tag", types.YLeaf{"RouteTag", addressXr.RouteTag})
    addressXr.EntityData.Leafs.Append("is-primary", types.YLeaf{"IsPrimary", addressXr.IsPrimary})
    addressXr.EntityData.Leafs.Append("is-tentative", types.YLeaf{"IsTentative", addressXr.IsTentative})
    addressXr.EntityData.Leafs.Append("is-prefix-sid", types.YLeaf{"IsPrefixSid", addressXr.IsPrefixSid})
    addressXr.EntityData.Leafs.Append("producer", types.YLeaf{"Producer", addressXr.Producer})

    addressXr.EntityData.YListKeys = []string {}

    return &(addressXr.EntityData)
}

// Ipv6arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address
// Address
type Ipv6arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address struct {
    EntityData types.CommonEntityData
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

func (address *Ipv6arm_Addresses_Vrfs_Vrf_Networks_Network_AddressXr_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "address-xr"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/addresses/vrfs/vrf/networks/network/address-xr/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", address.Afi})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Ipv6arm_Addresses_Vrfs_Vrf_Interfaces
// IPv6 ARM address database information by
// interface
type Ipv6arm_Addresses_Vrfs_Vrf_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An IPv6 address in IPv6 ARM. The type is slice of
    // Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface.
    Interface []*Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface
}

func (interfaces *Ipv6arm_Addresses_Vrfs_Vrf_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "vrf"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/addresses/vrfs/vrf/" + interfaces.EntityData.SegmentPath
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

// Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface
// An IPv6 address in IPv6 ARM
type Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    Interface interface{}

    // Referenced Interface - only valid for an unnumbered interface. The type is
    // string.
    ReferencedInterface interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // Address info. The type is slice of
    // Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address.
    Address []*Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address
}

func (self *Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Interface, "interface")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/addresses/vrfs/vrf/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("address", types.YChild{"Address", nil})
    for i := range self.Address {
        types.SetYListKey(self.Address[i], i)
        self.EntityData.Children.Append(types.GetSegmentPath(self.Address[i]), types.YChild{"Address", self.Address[i]})
    }
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", self.Interface})
    self.EntityData.Leafs.Append("referenced-interface", types.YLeaf{"ReferencedInterface", self.ReferencedInterface})
    self.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", self.VrfName})

    self.EntityData.YListKeys = []string {"Interface"}

    return &(self.EntityData)
}

// Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address
// Address info
type Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    Address Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address
}

func (address *Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "interface"
    address.EntityData.SegmentPath = "address" + types.AddNoKeyToken(address)
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/addresses/vrfs/vrf/interfaces/interface/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("address", types.YChild{"Address", &address.Address})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", address.PrefixLength})
    address.EntityData.Leafs.Append("route-tag", types.YLeaf{"RouteTag", address.RouteTag})
    address.EntityData.Leafs.Append("is-primary", types.YLeaf{"IsPrimary", address.IsPrimary})
    address.EntityData.Leafs.Append("is-tentative", types.YLeaf{"IsTentative", address.IsTentative})
    address.EntityData.Leafs.Append("is-prefix-sid", types.YLeaf{"IsPrefixSid", address.IsPrefixSid})
    address.EntityData.Leafs.Append("producer", types.YLeaf{"Producer", address.Producer})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address
// Address
type Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address struct {
    EntityData types.CommonEntityData
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

func (address *Ipv6arm_Addresses_Vrfs_Vrf_Interfaces_Interface_Address_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "address"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/addresses/vrfs/vrf/interfaces/interface/address/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", address.Afi})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Ipv6arm_Summary
// IPv6 ARM summary information
type Ipv6arm_Summary struct {
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

func (summary *Ipv6arm_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "ipv6arm"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("producer-count", types.YLeaf{"ProducerCount", summary.ProducerCount})
    summary.EntityData.Leafs.Append("address-conflict-count", types.YLeaf{"AddressConflictCount", summary.AddressConflictCount})
    summary.EntityData.Leafs.Append("unnumbered-conflict-count", types.YLeaf{"UnnumberedConflictCount", summary.UnnumberedConflictCount})
    summary.EntityData.Leafs.Append("db-master-version", types.YLeaf{"DbMasterVersion", summary.DbMasterVersion})
    summary.EntityData.Leafs.Append("vrf-count", types.YLeaf{"VrfCount", summary.VrfCount})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Ipv6arm_VrfSummaries
// IPv6 ARM VRFs summary information
type Ipv6arm_VrfSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 ARM VRF summary information. The type is slice of
    // Ipv6arm_VrfSummaries_VrfSummary.
    VrfSummary []*Ipv6arm_VrfSummaries_VrfSummary
}

func (vrfSummaries *Ipv6arm_VrfSummaries) GetEntityData() *types.CommonEntityData {
    vrfSummaries.EntityData.YFilter = vrfSummaries.YFilter
    vrfSummaries.EntityData.YangName = "vrf-summaries"
    vrfSummaries.EntityData.BundleName = "cisco_ios_xr"
    vrfSummaries.EntityData.ParentYangName = "ipv6arm"
    vrfSummaries.EntityData.SegmentPath = "vrf-summaries"
    vrfSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/" + vrfSummaries.EntityData.SegmentPath
    vrfSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummaries.EntityData.Children = types.NewOrderedMap()
    vrfSummaries.EntityData.Children.Append("vrf-summary", types.YChild{"VrfSummary", nil})
    for i := range vrfSummaries.VrfSummary {
        vrfSummaries.EntityData.Children.Append(types.GetSegmentPath(vrfSummaries.VrfSummary[i]), types.YChild{"VrfSummary", vrfSummaries.VrfSummary[i]})
    }
    vrfSummaries.EntityData.Leafs = types.NewOrderedMap()

    vrfSummaries.EntityData.YListKeys = []string {}

    return &(vrfSummaries.EntityData)
}

// Ipv6arm_VrfSummaries_VrfSummary
// IPv6 ARM VRF summary information
type Ipv6arm_VrfSummaries_VrfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string.
    VrfName interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // VRF Name. The type is string.
    VrfNameXr interface{}
}

func (vrfSummary *Ipv6arm_VrfSummaries_VrfSummary) GetEntityData() *types.CommonEntityData {
    vrfSummary.EntityData.YFilter = vrfSummary.YFilter
    vrfSummary.EntityData.YangName = "vrf-summary"
    vrfSummary.EntityData.BundleName = "cisco_ios_xr"
    vrfSummary.EntityData.ParentYangName = "vrf-summaries"
    vrfSummary.EntityData.SegmentPath = "vrf-summary" + types.AddKeyToken(vrfSummary.VrfName, "vrf-name")
    vrfSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-iarm-v6-oper:ipv6arm/vrf-summaries/" + vrfSummary.EntityData.SegmentPath
    vrfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummary.EntityData.Children = types.NewOrderedMap()
    vrfSummary.EntityData.Leafs = types.NewOrderedMap()
    vrfSummary.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrfSummary.VrfName})
    vrfSummary.EntityData.Leafs.Append("vrf-id", types.YLeaf{"VrfId", vrfSummary.VrfId})
    vrfSummary.EntityData.Leafs.Append("vrf-name-xr", types.YLeaf{"VrfNameXr", vrfSummary.VrfNameXr})

    vrfSummary.EntityData.YListKeys = []string {"VrfName"}

    return &(vrfSummary.EntityData)
}

