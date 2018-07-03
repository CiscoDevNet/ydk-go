// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-smiap package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ipv4-virtual: IPv4 virtual address for management interfaces
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_smiap_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_smiap_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-smiap-cfg ipv4-virtual}", reflect.TypeOf(Ipv4Virtual{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-smiap-cfg:ipv4-virtual", reflect.TypeOf(Ipv4Virtual{}))
}

// Ipv4Virtual
// IPv4 virtual address for management interfaces
type Ipv4Virtual struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable use as default source address on sourced packets. The type is
    // interface{}.
    UseAsSourceAddress interface{}

    // VRFs for the virtual IPv4 addresses.
    Vrfs Ipv4Virtual_Vrfs
}

func (ipv4Virtual *Ipv4Virtual) GetEntityData() *types.CommonEntityData {
    ipv4Virtual.EntityData.YFilter = ipv4Virtual.YFilter
    ipv4Virtual.EntityData.YangName = "ipv4-virtual"
    ipv4Virtual.EntityData.BundleName = "cisco_ios_xr"
    ipv4Virtual.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-smiap-cfg"
    ipv4Virtual.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-smiap-cfg:ipv4-virtual"
    ipv4Virtual.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Virtual.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Virtual.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Virtual.EntityData.Children = types.NewOrderedMap()
    ipv4Virtual.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &ipv4Virtual.Vrfs})
    ipv4Virtual.EntityData.Leafs = types.NewOrderedMap()
    ipv4Virtual.EntityData.Leafs.Append("use-as-source-address", types.YLeaf{"UseAsSourceAddress", ipv4Virtual.UseAsSourceAddress})

    ipv4Virtual.EntityData.YListKeys = []string {}

    return &(ipv4Virtual.EntityData)
}

// Ipv4Virtual_Vrfs
// VRFs for the virtual IPv4 addresses
type Ipv4Virtual_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRF for a virtual IPv4 address.  Specify 'default' for VRF default. The
    // type is slice of Ipv4Virtual_Vrfs_Vrf.
    Vrf []*Ipv4Virtual_Vrfs_Vrf
}

func (vrfs *Ipv4Virtual_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "ipv4-virtual"
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

// Ipv4Virtual_Vrfs_Vrf
// A VRF for a virtual IPv4 address.  Specify
// 'default' for VRF default
type Ipv4Virtual_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of VRF. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv4 sddress and mask.
    Address Ipv4Virtual_Vrfs_Vrf_Address
}

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
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

// Ipv4Virtual_Vrfs_Vrf_Address
// IPv4 sddress and mask
// This type is a presence type.
type Ipv4Virtual_Vrfs_Vrf_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}

    // IPv4 address mask. The type is interface{} with range: 0..32. This
    // attribute is mandatory.
    Netmask interface{}
}

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetEntityData() *types.CommonEntityData {
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
    address.EntityData.Leafs.Append("netmask", types.YLeaf{"Netmask", address.Netmask})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

