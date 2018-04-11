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

// MplsVpnAfi represents Layer 3 VPN Address Family Type
type MplsVpnAfi string

const (
    // VRF IPv4 address family
    MplsVpnAfi_ipv4 MplsVpnAfi = "ipv4"

    // VRF IPv6 address family
    MplsVpnAfi_ipv6 MplsVpnAfi = "ipv6"
)

// L3Vpn
// L3VPN operational data
type L3Vpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Invalid VRF Table (VRFs that are forward referenced).
    InvalidVrfs L3Vpn_InvalidVrfs

    // VRF Table.
    Vrfs L3Vpn_Vrfs
}

func (l3Vpn *L3Vpn) GetEntityData() *types.CommonEntityData {
    l3Vpn.EntityData.YFilter = l3Vpn.YFilter
    l3Vpn.EntityData.YangName = "l3vpn"
    l3Vpn.EntityData.BundleName = "cisco_ios_xr"
    l3Vpn.EntityData.ParentYangName = "Cisco-IOS-XR-mpls-vpn-oper"
    l3Vpn.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-vpn-oper:l3vpn"
    l3Vpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l3Vpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l3Vpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l3Vpn.EntityData.Children = make(map[string]types.YChild)
    l3Vpn.EntityData.Children["invalid-vrfs"] = types.YChild{"InvalidVrfs", &l3Vpn.InvalidVrfs}
    l3Vpn.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &l3Vpn.Vrfs}
    l3Vpn.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l3Vpn.EntityData)
}

// L3Vpn_InvalidVrfs
// Invalid VRF Table (VRFs that are forward
// referenced)
type L3Vpn_InvalidVrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Invalid VRF (VRF that is forward referenced). The type is slice of
    // L3Vpn_InvalidVrfs_InvalidVrf.
    InvalidVrf []L3Vpn_InvalidVrfs_InvalidVrf
}

func (invalidVrfs *L3Vpn_InvalidVrfs) GetEntityData() *types.CommonEntityData {
    invalidVrfs.EntityData.YFilter = invalidVrfs.YFilter
    invalidVrfs.EntityData.YangName = "invalid-vrfs"
    invalidVrfs.EntityData.BundleName = "cisco_ios_xr"
    invalidVrfs.EntityData.ParentYangName = "l3vpn"
    invalidVrfs.EntityData.SegmentPath = "invalid-vrfs"
    invalidVrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    invalidVrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    invalidVrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    invalidVrfs.EntityData.Children = make(map[string]types.YChild)
    invalidVrfs.EntityData.Children["invalid-vrf"] = types.YChild{"InvalidVrf", nil}
    for i := range invalidVrfs.InvalidVrf {
        invalidVrfs.EntityData.Children[types.GetSegmentPath(&invalidVrfs.InvalidVrf[i])] = types.YChild{"InvalidVrf", &invalidVrfs.InvalidVrf[i]}
    }
    invalidVrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(invalidVrfs.EntityData)
}

// L3Vpn_InvalidVrfs_InvalidVrf
// Invalid VRF (VRF that is forward referenced)
type L3Vpn_InvalidVrfs_InvalidVrf struct {
    EntityData types.CommonEntityData
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
    // L3Vpn_InvalidVrfs_InvalidVrf_Interface_.
    Interface_ []L3Vpn_InvalidVrfs_InvalidVrf_Interface

    // AF/SAF information. The type is slice of L3Vpn_InvalidVrfs_InvalidVrf_Af.
    Af []L3Vpn_InvalidVrfs_InvalidVrf_Af
}

func (invalidVrf *L3Vpn_InvalidVrfs_InvalidVrf) GetEntityData() *types.CommonEntityData {
    invalidVrf.EntityData.YFilter = invalidVrf.YFilter
    invalidVrf.EntityData.YangName = "invalid-vrf"
    invalidVrf.EntityData.BundleName = "cisco_ios_xr"
    invalidVrf.EntityData.ParentYangName = "invalid-vrfs"
    invalidVrf.EntityData.SegmentPath = "invalid-vrf" + "[vrf-name='" + fmt.Sprintf("%v", invalidVrf.VrfName) + "']"
    invalidVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    invalidVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    invalidVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    invalidVrf.EntityData.Children = make(map[string]types.YChild)
    invalidVrf.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range invalidVrf.Interface_ {
        invalidVrf.EntityData.Children[types.GetSegmentPath(&invalidVrf.Interface_[i])] = types.YChild{"Interface_", &invalidVrf.Interface_[i]}
    }
    invalidVrf.EntityData.Children["af"] = types.YChild{"Af", nil}
    for i := range invalidVrf.Af {
        invalidVrf.EntityData.Children[types.GetSegmentPath(&invalidVrf.Af[i])] = types.YChild{"Af", &invalidVrf.Af[i]}
    }
    invalidVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    invalidVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", invalidVrf.VrfName}
    invalidVrf.EntityData.Leafs["vrf-name-xr"] = types.YLeaf{"VrfNameXr", invalidVrf.VrfNameXr}
    invalidVrf.EntityData.Leafs["vrf-description"] = types.YLeaf{"VrfDescription", invalidVrf.VrfDescription}
    invalidVrf.EntityData.Leafs["route-distinguisher"] = types.YLeaf{"RouteDistinguisher", invalidVrf.RouteDistinguisher}
    invalidVrf.EntityData.Leafs["is-big-vrf"] = types.YLeaf{"IsBigVrf", invalidVrf.IsBigVrf}
    return &(invalidVrf.EntityData)
}

// L3Vpn_InvalidVrfs_InvalidVrf_Interface
// Interfaces in VRF
type L3Vpn_InvalidVrfs_InvalidVrf_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}
}

func (self *L3Vpn_InvalidVrfs_InvalidVrf_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "invalid-vrf"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// L3Vpn_InvalidVrfs_InvalidVrf_Af
// AF/SAF information
type L3Vpn_InvalidVrfs_InvalidVrf_Af struct {
    EntityData types.CommonEntityData
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

func (af *L3Vpn_InvalidVrfs_InvalidVrf_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "invalid-vrf"
    af.EntityData.SegmentPath = "af"
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = make(map[string]types.YChild)
    af.EntityData.Children["route-target"] = types.YChild{"RouteTarget", nil}
    for i := range af.RouteTarget {
        af.EntityData.Children[types.GetSegmentPath(&af.RouteTarget[i])] = types.YChild{"RouteTarget", &af.RouteTarget[i]}
    }
    af.EntityData.Leafs = make(map[string]types.YLeaf)
    af.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", af.AfName}
    af.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", af.SafName}
    af.EntityData.Leafs["import-route-policy"] = types.YLeaf{"ImportRoutePolicy", af.ImportRoutePolicy}
    af.EntityData.Leafs["export-route-policy"] = types.YLeaf{"ExportRoutePolicy", af.ExportRoutePolicy}
    return &(af.EntityData)
}

// L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget
// Route Targets
type L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget struct {
    EntityData types.CommonEntityData
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

func (routeTarget *L3Vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "af"
    routeTarget.EntityData.SegmentPath = "route-target"
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = make(map[string]types.YChild)
    routeTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTarget.EntityData.Leafs["route-target-type"] = types.YLeaf{"RouteTargetType", routeTarget.RouteTargetType}
    routeTarget.EntityData.Leafs["route-target-value"] = types.YLeaf{"RouteTargetValue", routeTarget.RouteTargetValue}
    routeTarget.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", routeTarget.AfName}
    routeTarget.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", routeTarget.SafName}
    return &(routeTarget.EntityData)
}

// L3Vpn_Vrfs
// VRF Table
type L3Vpn_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF. The type is slice of L3Vpn_Vrfs_Vrf.
    Vrf []L3Vpn_Vrfs_Vrf
}

func (vrfs *L3Vpn_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "l3vpn"
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

// L3Vpn_Vrfs_Vrf
// VRF
type L3Vpn_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
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

    // Interfaces in VRF. The type is slice of L3Vpn_Vrfs_Vrf_Interface_.
    Interface_ []L3Vpn_Vrfs_Vrf_Interface

    // AF/SAF information. The type is slice of L3Vpn_Vrfs_Vrf_Af.
    Af []L3Vpn_Vrfs_Vrf_Af
}

func (vrf *L3Vpn_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range vrf.Interface_ {
        vrf.EntityData.Children[types.GetSegmentPath(&vrf.Interface_[i])] = types.YChild{"Interface_", &vrf.Interface_[i]}
    }
    vrf.EntityData.Children["af"] = types.YChild{"Af", nil}
    for i := range vrf.Af {
        vrf.EntityData.Children[types.GetSegmentPath(&vrf.Af[i])] = types.YChild{"Af", &vrf.Af[i]}
    }
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    vrf.EntityData.Leafs["vrf-name-xr"] = types.YLeaf{"VrfNameXr", vrf.VrfNameXr}
    vrf.EntityData.Leafs["vrf-description"] = types.YLeaf{"VrfDescription", vrf.VrfDescription}
    vrf.EntityData.Leafs["route-distinguisher"] = types.YLeaf{"RouteDistinguisher", vrf.RouteDistinguisher}
    vrf.EntityData.Leafs["is-big-vrf"] = types.YLeaf{"IsBigVrf", vrf.IsBigVrf}
    return &(vrf.EntityData)
}

// L3Vpn_Vrfs_Vrf_Interface
// Interfaces in VRF
type L3Vpn_Vrfs_Vrf_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}
}

func (self *L3Vpn_Vrfs_Vrf_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "vrf"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// L3Vpn_Vrfs_Vrf_Af
// AF/SAF information
type L3Vpn_Vrfs_Vrf_Af struct {
    EntityData types.CommonEntityData
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

func (af *L3Vpn_Vrfs_Vrf_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "vrf"
    af.EntityData.SegmentPath = "af"
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = make(map[string]types.YChild)
    af.EntityData.Children["route-target"] = types.YChild{"RouteTarget", nil}
    for i := range af.RouteTarget {
        af.EntityData.Children[types.GetSegmentPath(&af.RouteTarget[i])] = types.YChild{"RouteTarget", &af.RouteTarget[i]}
    }
    af.EntityData.Leafs = make(map[string]types.YLeaf)
    af.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", af.AfName}
    af.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", af.SafName}
    af.EntityData.Leafs["import-route-policy"] = types.YLeaf{"ImportRoutePolicy", af.ImportRoutePolicy}
    af.EntityData.Leafs["export-route-policy"] = types.YLeaf{"ExportRoutePolicy", af.ExportRoutePolicy}
    return &(af.EntityData)
}

// L3Vpn_Vrfs_Vrf_Af_RouteTarget
// Route Targets
type L3Vpn_Vrfs_Vrf_Af_RouteTarget struct {
    EntityData types.CommonEntityData
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

func (routeTarget *L3Vpn_Vrfs_Vrf_Af_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "af"
    routeTarget.EntityData.SegmentPath = "route-target"
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = make(map[string]types.YChild)
    routeTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTarget.EntityData.Leafs["route-target-type"] = types.YLeaf{"RouteTargetType", routeTarget.RouteTargetType}
    routeTarget.EntityData.Leafs["route-target-value"] = types.YLeaf{"RouteTargetValue", routeTarget.RouteTargetValue}
    routeTarget.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", routeTarget.AfName}
    routeTarget.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", routeTarget.SafName}
    return &(routeTarget.EntityData)
}

