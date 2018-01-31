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
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable use as default source address on sourced packets. The type is
    // interface{}.
    UseAsSourceAddress interface{}

    // VRFs for the virtual IPv6 addresses.
    Vrfs Ipv6Virtual_Vrfs
}

func (ipv6Virtual *Ipv6Virtual) GetFilter() yfilter.YFilter { return ipv6Virtual.YFilter }

func (ipv6Virtual *Ipv6Virtual) SetFilter(yf yfilter.YFilter) { ipv6Virtual.YFilter = yf }

func (ipv6Virtual *Ipv6Virtual) GetGoName(yname string) string {
    if yname == "use-as-source-address" { return "UseAsSourceAddress" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (ipv6Virtual *Ipv6Virtual) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-smiap-cfg:ipv6-virtual"
}

func (ipv6Virtual *Ipv6Virtual) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &ipv6Virtual.Vrfs
    }
    return nil
}

func (ipv6Virtual *Ipv6Virtual) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &ipv6Virtual.Vrfs
    return children
}

func (ipv6Virtual *Ipv6Virtual) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["use-as-source-address"] = ipv6Virtual.UseAsSourceAddress
    return leafs
}

func (ipv6Virtual *Ipv6Virtual) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Virtual *Ipv6Virtual) GetYangName() string { return "ipv6-virtual" }

func (ipv6Virtual *Ipv6Virtual) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Virtual *Ipv6Virtual) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Virtual *Ipv6Virtual) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Virtual *Ipv6Virtual) SetParent(parent types.Entity) { ipv6Virtual.parent = parent }

func (ipv6Virtual *Ipv6Virtual) GetParent() types.Entity { return ipv6Virtual.parent }

func (ipv6Virtual *Ipv6Virtual) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-smiap-cfg" }

// Ipv6Virtual_Vrfs
// VRFs for the virtual IPv6 addresses
type Ipv6Virtual_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRF for a virtual IPv6 address.  Specify 'default' for VRF default. The
    // type is slice of Ipv6Virtual_Vrfs_Vrf.
    Vrf []Ipv6Virtual_Vrfs_Vrf
}

func (vrfs *Ipv6Virtual_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv6Virtual_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv6Virtual_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv6Virtual_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv6Virtual_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Virtual_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv6Virtual_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv6Virtual_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv6Virtual_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv6Virtual_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv6Virtual_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv6Virtual_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv6Virtual_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv6Virtual_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv6Virtual_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv6Virtual_Vrfs) GetParentYangName() string { return "ipv6-virtual" }

// Ipv6Virtual_Vrfs_Vrf
// A VRF for a virtual IPv6 address.  Specify
// 'default' for VRF default
type Ipv6Virtual_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of VRF. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv6 address and mask.
    Address Ipv6Virtual_Vrfs_Vrf_Address
}

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv6Virtual_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "address" { return "Address" }
    return ""
}

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &vrf.Address
    }
    return nil
}

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &vrf.Address
    return children
}

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv6Virtual_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv6Virtual_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv6Virtual_Vrfs_Vrf_Address
// IPv6 address and mask
// This type is a presence type.
type Ipv6Virtual_Vrfs_Vrf_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}

    // IPv6 address prefix length. The type is interface{} with range: 0..128.
    // This attribute is mandatory.
    PrefixLength interface{}
}

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ipv6Virtual_Vrfs_Vrf_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = address.Address
    leafs["prefix-length"] = address.PrefixLength
    return leafs
}

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetYangName() string { return "address" }

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ipv6Virtual_Vrfs_Vrf_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetParent() types.Entity { return address.parent }

func (address *Ipv6Virtual_Vrfs_Vrf_Address) GetParentYangName() string { return "vrf" }

