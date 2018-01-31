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
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable use as default source address on sourced packets. The type is
    // interface{}.
    UseAsSourceAddress interface{}

    // VRFs for the virtual IPv4 addresses.
    Vrfs Ipv4Virtual_Vrfs
}

func (ipv4Virtual *Ipv4Virtual) GetFilter() yfilter.YFilter { return ipv4Virtual.YFilter }

func (ipv4Virtual *Ipv4Virtual) SetFilter(yf yfilter.YFilter) { ipv4Virtual.YFilter = yf }

func (ipv4Virtual *Ipv4Virtual) GetGoName(yname string) string {
    if yname == "use-as-source-address" { return "UseAsSourceAddress" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (ipv4Virtual *Ipv4Virtual) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-smiap-cfg:ipv4-virtual"
}

func (ipv4Virtual *Ipv4Virtual) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &ipv4Virtual.Vrfs
    }
    return nil
}

func (ipv4Virtual *Ipv4Virtual) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &ipv4Virtual.Vrfs
    return children
}

func (ipv4Virtual *Ipv4Virtual) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["use-as-source-address"] = ipv4Virtual.UseAsSourceAddress
    return leafs
}

func (ipv4Virtual *Ipv4Virtual) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Virtual *Ipv4Virtual) GetYangName() string { return "ipv4-virtual" }

func (ipv4Virtual *Ipv4Virtual) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Virtual *Ipv4Virtual) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Virtual *Ipv4Virtual) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Virtual *Ipv4Virtual) SetParent(parent types.Entity) { ipv4Virtual.parent = parent }

func (ipv4Virtual *Ipv4Virtual) GetParent() types.Entity { return ipv4Virtual.parent }

func (ipv4Virtual *Ipv4Virtual) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-smiap-cfg" }

// Ipv4Virtual_Vrfs
// VRFs for the virtual IPv4 addresses
type Ipv4Virtual_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRF for a virtual IPv4 address.  Specify 'default' for VRF default. The
    // type is slice of Ipv4Virtual_Vrfs_Vrf.
    Vrf []Ipv4Virtual_Vrfs_Vrf
}

func (vrfs *Ipv4Virtual_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv4Virtual_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv4Virtual_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv4Virtual_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv4Virtual_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Virtual_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv4Virtual_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv4Virtual_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv4Virtual_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv4Virtual_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv4Virtual_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv4Virtual_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv4Virtual_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv4Virtual_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv4Virtual_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv4Virtual_Vrfs) GetParentYangName() string { return "ipv4-virtual" }

// Ipv4Virtual_Vrfs_Vrf
// A VRF for a virtual IPv4 address.  Specify
// 'default' for VRF default
type Ipv4Virtual_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of VRF. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv4 sddress and mask.
    Address Ipv4Virtual_Vrfs_Vrf_Address
}

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv4Virtual_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "address" { return "Address" }
    return ""
}

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &vrf.Address
    }
    return nil
}

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &vrf.Address
    return children
}

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv4Virtual_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv4Virtual_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv4Virtual_Vrfs_Vrf_Address
// IPv4 sddress and mask
// This type is a presence type.
type Ipv4Virtual_Vrfs_Vrf_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}

    // IPv4 address mask. The type is interface{} with range: 0..32. This
    // attribute is mandatory.
    Netmask interface{}
}

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ipv4Virtual_Vrfs_Vrf_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "netmask" { return "Netmask" }
    return ""
}

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = address.Address
    leafs["netmask"] = address.Netmask
    return leafs
}

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetYangName() string { return "address" }

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ipv4Virtual_Vrfs_Vrf_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetParent() types.Entity { return address.parent }

func (address *Ipv4Virtual_Vrfs_Vrf_Address) GetParentYangName() string { return "vrf" }

