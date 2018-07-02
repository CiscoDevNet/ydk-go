// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-smiap package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ipv6-virtual: IPv6 virtual address for management interfaces
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_smiap_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_smiap_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-smiap-cfg ipv6-virtual}", reflect.TypeOf(Ipv6Virtual{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-smiap-cfg:ipv6-virtual", reflect.TypeOf(Ipv6Virtual{}))
}

// Ipv6Virtual
// IPv6 virtual address for management interfaces
type Ipv6Virtual struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable use as default source address on sourced packets. The type is
    // interface{}.
    UseAsSourceAddress interface{}

    // VRFs for the virtual IPv6 addresses.
    Vrfs Ipv6Virtual_Vrfs
}

func (ipv6Virtual *Ipv6Virtual) GetEntityData() *types.CommonEntityData {
    ipv6Virtual.EntityData.YFilter = ipv6Virtual.YFilter
    ipv6Virtual.EntityData.YangName = "ipv6-virtual"
    ipv6Virtual.EntityData.BundleName = "cisco_ios_xr"
    ipv6Virtual.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-smiap-cfg"
    ipv6Virtual.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-smiap-cfg:ipv6-virtual"
    ipv6Virtual.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Virtual.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Virtual.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Virtual.EntityData.Children = types.NewOrderedMap()
    ipv6Virtual.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &ipv6Virtual.Vrfs})
    ipv6Virtual.EntityData.Leafs = types.NewOrderedMap()
    ipv6Virtual.EntityData.Leafs.Append("use-as-source-address", types.YLeaf{"UseAsSourceAddress", ipv6Virtual.UseAsSourceAddress})

    ipv6Virtual.EntityData.YListKeys = []string {}

    return &(ipv6Virtual.EntityData)
}

// Ipv6Virtual_Vrfs
// VRFs for the virtual IPv6 addresses
type Ipv6Virtual_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRF for a virtual IPv6 address.  Specify 'default' for VRF default. The
    // type is slice of Ipv6Virtual_Vrfs_Vrf.
    Vrf []*Ipv6Virtual_Vrfs_Vrf
}

func (vrfs *Ipv6Virtual_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "ipv6-virtual"
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

// Ipv6Virtual_Vrfs_Vrf
// A VRF for a virtual IPv6 address.  Specify
// 'default' for VRF default
type Ipv6Virtual_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of VRF. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv6 address and mask.
    Address Ipv6Virtual_Vrfs_Vrf_Address
}

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("address", types.YChild{"Address", &vrf.Address})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Ipv6Virtual_Vrfs_Vrf_Address
// IPv6 address and mask
// This type is a presence type.
type Ipv6Virtual_Vrfs_Vrf_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}

    // IPv6 address prefix length. The type is interface{} with range: 0..128.
    // This attribute is mandatory.
    PrefixLength interface{}
}

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "vrf"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address", types.YLeaf{"Address", address.Address})
    address.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", address.PrefixLength})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

