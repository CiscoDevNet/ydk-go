// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-vpn package operational data.
// 
// This module contains definitions
// for the following management objects:
//   l3vpn: L3VPN operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package mpls_vpn_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_vpn_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mpls-vpn-oper l3vpn}", reflect.TypeOf(L3Vpn{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mpls-vpn-oper:l3vpn", reflect.TypeOf(L3Vpn{}))
}

// MplsVpnRt represents Layer 3 VPN Route Target Type
type MplsVpnRt string

const (
    // VRF Route Target Type Import
    MplsVpnRt_import_ MplsVpnRt = "import"

    // VRF Route Target Type Export
    MplsVpnRt_export MplsVpnRt = "export"

    // VRF Route Target Type Import and Export
    MplsVpnRt_both MplsVpnRt = "both"
)

// MplsVpnAfi represents Layer 3 VPN Address Family Type
type MplsVpnAfi string

const (
    // VRF IPv4 address family
    MplsVpnAfi_ipv4 MplsVpnAfi = "ipv4"

    // VRF IPv6 address family
    MplsVpnAfi_ipv6 MplsVpnAfi = "ipv6"
)

// MplsVpnSafi represents Layer 3 VPN Sub-Address Family Type
type MplsVpnSafi string

const (
    // VRF Unicast sub-address family
    MplsVpnSafi_unicast MplsVpnSafi = "unicast"

    // VRF Multicast sub-address family
    MplsVpnSafi_multicast MplsVpnSafi = "multicast"

    // VRF Flowspec sub-address family
    MplsVpnSafi_flowspec MplsVpnSafi = "flowspec"
)

// L3Vpn
// L3VPN operational data
type L3Vpn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Invalid VRF Table (VRFs that are forward referenced).
    InvalidVrfs L3Vpn_InvalidVrfs

    // VRF Table.
    Vrfs L3Vpn_Vrfs
}

func (l3Vpn *L3Vpn) GetFilter() yfilter.YFilter { return l3Vpn.YFilter }

func (l3Vpn *L3Vpn) SetFilter(yf yfilter.YFilter) { l3Vpn.YFilter = yf }

func (l3Vpn *L3Vpn) GetGoName(yname string) string {
    if yname == "invalid-vrfs" { return "InvalidVrfs" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (l3Vpn *L3Vpn) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-vpn-oper:l3vpn"
}

func (l3Vpn *L3Vpn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "invalid-vrfs" {
        return &l3Vpn.InvalidVrfs
    }
    if childYangName == "vrfs" {
        return &l3Vpn.Vrfs
    }
    return nil
}

func (l3Vpn *L3Vpn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["invalid-vrfs"] = &l3Vpn.InvalidVrfs
    children["vrfs"] = &l3Vpn.Vrfs
    return children
}

func (l3Vpn *L3Vpn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l3Vpn *L3Vpn) GetBundleName() string { return "cisco_ios_xr" }

func (l3Vpn *L3Vpn) GetYangName() string { return "l3vpn" }

func (l3Vpn *L3Vpn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (l3Vpn *L3Vpn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (l3Vpn *L3Vpn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (l3Vpn *L3Vpn) SetParent(parent types.Entity) { l3Vpn.parent = parent }

func (l3Vpn *L3Vpn) GetParent() types.Entity { return l3Vpn.parent }

func (l3Vpn *L3Vpn) GetParentYangName() string { return "Cisco-IOS-XR-mpls-vpn-oper" }

// L3Vpn_InvalidVrfs
// Invalid VRF Table (VRFs that are forward
// referenced)
type L3Vpn_InvalidVrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Invalid VRF (VRF that is forward referenced). The type is slice of
    // L3Vpn_InvalidVrfs_InvalidVrf.
    InvalidVrf []L3Vpn_InvalidVrfs_InvalidVrf
}

func (invalidVrfs *L3Vpn_InvalidVrfs) GetFilter() yfilter.YFilter { return invalidVrfs.YFilter }

func (invalidVrfs *L3Vpn_InvalidVrfs) SetFilter(yf yfilter.YFilter) { invalidVrfs.YFilter = yf }

func (invalidVrfs *L3Vpn_InvalidVrfs) GetGoName(yname string) string {
    if yname == "invalid-vrf" { return "InvalidVrf" }
    return ""
}

func (invalidVrfs *L3Vpn_InvalidVrfs) GetSegmentPath() string {
    return "invalid-vrfs"
}

func (invalidVrfs *L3Vpn_InvalidVrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "invalid-vrf" {
        for _, c := range invalidVrfs.InvalidVrf {
            if invalidVrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L3Vpn_InvalidVrfs_InvalidVrf{}
        invalidVrfs.InvalidVrf = append(invalidVrfs.InvalidVrf, child)
        return &invalidVrfs.InvalidVrf[len(invalidVrfs.InvalidVrf)-1]
    }
    return nil
}

func (invalidVrfs *L3Vpn_InvalidVrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range invalidVrfs.InvalidVrf {
        children[invalidVrfs.InvalidVrf[i].GetSegmentPath()] = &invalidVrfs.InvalidVrf[i]
    }
    return children
}

func (invalidVrfs *L3Vpn_InvalidVrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (invalidVrfs *L3Vpn_InvalidVrfs) GetBundleName() string { return "cisco_ios_xr" }

func (invalidVrfs *L3Vpn_InvalidVrfs) GetYangName() string { return "invalid-vrfs" }

func (invalidVrfs *L3Vpn_InvalidVrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (invalidVrfs *L3Vpn_InvalidVrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (invalidVrfs *L3Vpn_InvalidVrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (invalidVrfs *L3Vpn_InvalidVrfs) SetParent(parent types.Entity) { invalidVrfs.parent = parent }

func (invalidVrfs *L3Vpn_InvalidVrfs) GetParent() types.Entity { return invalidVrfs.parent }

func (invalidVrfs *L3Vpn_InvalidVrfs) GetParentYangName() string { return "l3vpn" }

// L3Vpn_InvalidVrfs_InvalidVrf
// Invalid VRF (VRF that is forward referenced)
type L3Vpn_InvalidVrfs_InvalidVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Name for an invalid VRF. The type is string.
    VrfName interface{}

    // VRF Name. The type is string.
    VrfNameXr interface{}

    // VRF Description. The type is string.
    VrfDescription interface{}

    // Route Distinguisher. The type is string.
    RouteDistinguisher interface{}

    // VRF mode information. The type is bool.
    IsBigVrf interface{}

    // Interfaces in VRF. The type is slice of
    // L3Vpn_InvalidVrfs_InvalidVrf_Interface.
    Interface []L3Vpn_InvalidVrfs_InvalidVrf_Interface

    // AF/SAF information. The type is slice of L3Vpn_InvalidVrfs_InvalidVrf_Af.
    Af []L3Vpn_InvalidVrfs_InvalidVrf_Af
}

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetFilter() yfilter.YFilter { return invalidVrf.YFilter }

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) SetFilter(yf yfilter.YFilter) { invalidVrf.YFilter = yf }

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "vrf-name-xr" { return "VrfNameXr" }
    if yname == "vrf-description" { return "VrfDescription" }
    if yname == "route-distinguisher" { return "RouteDistinguisher" }
    if yname == "is-big-vrf" { return "IsBigVrf" }
    if yname == "interface" { return "Interface" }
    if yname == "af" { return "Af" }
    return ""
}

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetSegmentPath() string {
    return "invalid-vrf" + "[vrf-name='" + fmt.Sprintf("%v", invalidVrf.VrfName) + "']"
}

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range invalidVrf.Interface {
            if invalidVrf.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L3Vpn_InvalidVrfs_InvalidVrf_Interface{}
        invalidVrf.Interface = append(invalidVrf.Interface, child)
        return &invalidVrf.Interface[len(invalidVrf.Interface)-1]
    }
    if childYangName == "af" {
        for _, c := range invalidVrf.Af {
            if invalidVrf.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L3Vpn_InvalidVrfs_InvalidVrf_Af{}
        invalidVrf.Af = append(invalidVrf.Af, child)
        return &invalidVrf.Af[len(invalidVrf.Af)-1]
    }
    return nil
}

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range invalidVrf.Interface {
        children[invalidVrf.Interface[i].GetSegmentPath()] = &invalidVrf.Interface[i]
    }
    for i := range invalidVrf.Af {
        children[invalidVrf.Af[i].GetSegmentPath()] = &invalidVrf.Af[i]
    }
    return children
}

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = invalidVrf.VrfName
    leafs["vrf-name-xr"] = invalidVrf.VrfNameXr
    leafs["vrf-description"] = invalidVrf.VrfDescription
    leafs["route-distinguisher"] = invalidVrf.RouteDistinguisher
    leafs["is-big-vrf"] = invalidVrf.IsBigVrf
    return leafs
}

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetBundleName() string { return "cisco_ios_xr" }

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetYangName() string { return "invalid-vrf" }

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) SetParent(parent types.Entity) { invalidVrf.parent = parent }

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetParent() types.Entity { return invalidVrf.parent }

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetParentYangName() string { return "invalid-vrfs" }

// L3Vpn_InvalidVrfs_InvalidVrf_Interface
// Interfaces in VRF
type L3Vpn_InvalidVrfs_InvalidVrf_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}
}

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetYangName() string { return "interface" }

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetParent() types.Entity { return self.parent }

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetParentYangName() string { return "invalid-vrf" }

// L3Vpn_InvalidVrfs_InvalidVrf_Af
// AF/SAF information
type L3Vpn_InvalidVrfs_InvalidVrf_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AF name. The type is MplsVpnAfi.
    AfName interface{}

    // SAF name. The type is MplsVpnSafi.
    SafName interface{}

    // Import Route Policy. The type is string.
    ImportRoutePolicy interface{}

    // Export Route Policy. The type is string.
    ExportRoutePolicy interface{}

    // Route Targets. The type is slice of
    // L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget.
    RouteTarget []L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget
}

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "import-route-policy" { return "ImportRoutePolicy" }
    if yname == "export-route-policy" { return "ExportRoutePolicy" }
    if yname == "route-target" { return "RouteTarget" }
    return ""
}

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetSegmentPath() string {
    return "af"
}

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-target" {
        for _, c := range af.RouteTarget {
            if af.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget{}
        af.RouteTarget = append(af.RouteTarget, child)
        return &af.RouteTarget[len(af.RouteTarget)-1]
    }
    return nil
}

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range af.RouteTarget {
        children[af.RouteTarget[i].GetSegmentPath()] = &af.RouteTarget[i]
    }
    return children
}

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    leafs["saf-name"] = af.SafName
    leafs["import-route-policy"] = af.ImportRoutePolicy
    leafs["export-route-policy"] = af.ExportRoutePolicy
    return leafs
}

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetYangName() string { return "af" }

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetParent() types.Entity { return af.parent }

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetParentYangName() string { return "invalid-vrf" }

// L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget
// Route Targets
type L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route Target Type. The type is MplsVpnRt.
    RouteTargetType interface{}

    // Route Target Value. The type is string.
    RouteTargetValue interface{}

    // AF name. The type is MplsVpnAfi.
    AfName interface{}

    // SAF name. The type is MplsVpnSafi.
    SafName interface{}
}

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetFilter() yfilter.YFilter { return routeTarget.YFilter }

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) SetFilter(yf yfilter.YFilter) { routeTarget.YFilter = yf }

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetGoName(yname string) string {
    if yname == "route-target-type" { return "RouteTargetType" }
    if yname == "route-target-value" { return "RouteTargetValue" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    return ""
}

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetSegmentPath() string {
    return "route-target"
}

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-target-type"] = routeTarget.RouteTargetType
    leafs["route-target-value"] = routeTarget.RouteTargetValue
    leafs["af-name"] = routeTarget.AfName
    leafs["saf-name"] = routeTarget.SafName
    return leafs
}

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetBundleName() string { return "cisco_ios_xr" }

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetYangName() string { return "route-target" }

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) SetParent(parent types.Entity) { routeTarget.parent = parent }

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetParent() types.Entity { return routeTarget.parent }

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetParentYangName() string { return "af" }

// L3Vpn_Vrfs
// VRF Table
type L3Vpn_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF. The type is slice of L3Vpn_Vrfs_Vrf.
    Vrf []L3Vpn_Vrfs_Vrf
}

func (vrfs *L3Vpn_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *L3Vpn_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *L3Vpn_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *L3Vpn_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *L3Vpn_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L3Vpn_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *L3Vpn_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *L3Vpn_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *L3Vpn_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *L3Vpn_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *L3Vpn_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *L3Vpn_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *L3Vpn_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *L3Vpn_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *L3Vpn_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *L3Vpn_Vrfs) GetParentYangName() string { return "l3vpn" }

// L3Vpn_Vrfs_Vrf
// VRF
type L3Vpn_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Name for a VRF. The type is string.
    VrfName interface{}

    // VRF Name. The type is string.
    VrfNameXr interface{}

    // VRF Description. The type is string.
    VrfDescription interface{}

    // Route Distinguisher. The type is string.
    RouteDistinguisher interface{}

    // VRF mode information. The type is bool.
    IsBigVrf interface{}

    // Interfaces in VRF. The type is slice of L3Vpn_Vrfs_Vrf_Interface.
    Interface []L3Vpn_Vrfs_Vrf_Interface

    // AF/SAF information. The type is slice of L3Vpn_Vrfs_Vrf_Af.
    Af []L3Vpn_Vrfs_Vrf_Af
}

func (vrf *L3Vpn_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *L3Vpn_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *L3Vpn_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "vrf-name-xr" { return "VrfNameXr" }
    if yname == "vrf-description" { return "VrfDescription" }
    if yname == "route-distinguisher" { return "RouteDistinguisher" }
    if yname == "is-big-vrf" { return "IsBigVrf" }
    if yname == "interface" { return "Interface" }
    if yname == "af" { return "Af" }
    return ""
}

func (vrf *L3Vpn_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *L3Vpn_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range vrf.Interface {
            if vrf.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L3Vpn_Vrfs_Vrf_Interface{}
        vrf.Interface = append(vrf.Interface, child)
        return &vrf.Interface[len(vrf.Interface)-1]
    }
    if childYangName == "af" {
        for _, c := range vrf.Af {
            if vrf.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L3Vpn_Vrfs_Vrf_Af{}
        vrf.Af = append(vrf.Af, child)
        return &vrf.Af[len(vrf.Af)-1]
    }
    return nil
}

func (vrf *L3Vpn_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrf.Interface {
        children[vrf.Interface[i].GetSegmentPath()] = &vrf.Interface[i]
    }
    for i := range vrf.Af {
        children[vrf.Af[i].GetSegmentPath()] = &vrf.Af[i]
    }
    return children
}

func (vrf *L3Vpn_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["vrf-name-xr"] = vrf.VrfNameXr
    leafs["vrf-description"] = vrf.VrfDescription
    leafs["route-distinguisher"] = vrf.RouteDistinguisher
    leafs["is-big-vrf"] = vrf.IsBigVrf
    return leafs
}

func (vrf *L3Vpn_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *L3Vpn_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *L3Vpn_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *L3Vpn_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *L3Vpn_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *L3Vpn_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *L3Vpn_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *L3Vpn_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// L3Vpn_Vrfs_Vrf_Interface
// Interfaces in VRF
type L3Vpn_Vrfs_Vrf_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}
}

func (self *L3Vpn_Vrfs_Vrf_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *L3Vpn_Vrfs_Vrf_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *L3Vpn_Vrfs_Vrf_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *L3Vpn_Vrfs_Vrf_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *L3Vpn_Vrfs_Vrf_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *L3Vpn_Vrfs_Vrf_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *L3Vpn_Vrfs_Vrf_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *L3Vpn_Vrfs_Vrf_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *L3Vpn_Vrfs_Vrf_Interface) GetYangName() string { return "interface" }

func (self *L3Vpn_Vrfs_Vrf_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *L3Vpn_Vrfs_Vrf_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *L3Vpn_Vrfs_Vrf_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *L3Vpn_Vrfs_Vrf_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *L3Vpn_Vrfs_Vrf_Interface) GetParent() types.Entity { return self.parent }

func (self *L3Vpn_Vrfs_Vrf_Interface) GetParentYangName() string { return "vrf" }

// L3Vpn_Vrfs_Vrf_Af
// AF/SAF information
type L3Vpn_Vrfs_Vrf_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AF name. The type is MplsVpnAfi.
    AfName interface{}

    // SAF name. The type is MplsVpnSafi.
    SafName interface{}

    // Import Route Policy. The type is string.
    ImportRoutePolicy interface{}

    // Export Route Policy. The type is string.
    ExportRoutePolicy interface{}

    // Route Targets. The type is slice of L3Vpn_Vrfs_Vrf_Af_RouteTarget.
    RouteTarget []L3Vpn_Vrfs_Vrf_Af_RouteTarget
}

func (af *L3Vpn_Vrfs_Vrf_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *L3Vpn_Vrfs_Vrf_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *L3Vpn_Vrfs_Vrf_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "import-route-policy" { return "ImportRoutePolicy" }
    if yname == "export-route-policy" { return "ExportRoutePolicy" }
    if yname == "route-target" { return "RouteTarget" }
    return ""
}

func (af *L3Vpn_Vrfs_Vrf_Af) GetSegmentPath() string {
    return "af"
}

func (af *L3Vpn_Vrfs_Vrf_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-target" {
        for _, c := range af.RouteTarget {
            if af.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L3Vpn_Vrfs_Vrf_Af_RouteTarget{}
        af.RouteTarget = append(af.RouteTarget, child)
        return &af.RouteTarget[len(af.RouteTarget)-1]
    }
    return nil
}

func (af *L3Vpn_Vrfs_Vrf_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range af.RouteTarget {
        children[af.RouteTarget[i].GetSegmentPath()] = &af.RouteTarget[i]
    }
    return children
}

func (af *L3Vpn_Vrfs_Vrf_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    leafs["saf-name"] = af.SafName
    leafs["import-route-policy"] = af.ImportRoutePolicy
    leafs["export-route-policy"] = af.ExportRoutePolicy
    return leafs
}

func (af *L3Vpn_Vrfs_Vrf_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *L3Vpn_Vrfs_Vrf_Af) GetYangName() string { return "af" }

func (af *L3Vpn_Vrfs_Vrf_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *L3Vpn_Vrfs_Vrf_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *L3Vpn_Vrfs_Vrf_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *L3Vpn_Vrfs_Vrf_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *L3Vpn_Vrfs_Vrf_Af) GetParent() types.Entity { return af.parent }

func (af *L3Vpn_Vrfs_Vrf_Af) GetParentYangName() string { return "vrf" }

// L3Vpn_Vrfs_Vrf_Af_RouteTarget
// Route Targets
type L3Vpn_Vrfs_Vrf_Af_RouteTarget struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route Target Type. The type is MplsVpnRt.
    RouteTargetType interface{}

    // Route Target Value. The type is string.
    RouteTargetValue interface{}

    // AF name. The type is MplsVpnAfi.
    AfName interface{}

    // SAF name. The type is MplsVpnSafi.
    SafName interface{}
}

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetFilter() yfilter.YFilter { return routeTarget.YFilter }

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) SetFilter(yf yfilter.YFilter) { routeTarget.YFilter = yf }

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetGoName(yname string) string {
    if yname == "route-target-type" { return "RouteTargetType" }
    if yname == "route-target-value" { return "RouteTargetValue" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    return ""
}

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetSegmentPath() string {
    return "route-target"
}

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-target-type"] = routeTarget.RouteTargetType
    leafs["route-target-value"] = routeTarget.RouteTargetValue
    leafs["af-name"] = routeTarget.AfName
    leafs["saf-name"] = routeTarget.SafName
    return leafs
}

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetBundleName() string { return "cisco_ios_xr" }

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetYangName() string { return "route-target" }

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) SetParent(parent types.Entity) { routeTarget.parent = parent }

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetParent() types.Entity { return routeTarget.parent }

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetParentYangName() string { return "af" }

