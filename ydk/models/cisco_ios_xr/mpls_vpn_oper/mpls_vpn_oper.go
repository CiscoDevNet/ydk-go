// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-vpn package operational data.
// 
// This module contains definitions
// for the following management objects:
//   l3vpn: L3VPN operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mpls-vpn-oper l3vpn}", reflect.TypeOf(L3vpn{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mpls-vpn-oper:l3vpn", reflect.TypeOf(L3vpn{}))
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

// L3vpn
// L3VPN operational data
type L3vpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Invalid VRF Table (VRFs that are forward referenced).
    InvalidVrfs L3vpn_InvalidVrfs

    // VRF Table.
    Vrfs L3vpn_Vrfs
}

func (l3vpn *L3vpn) GetEntityData() *types.CommonEntityData {
    l3vpn.EntityData.YFilter = l3vpn.YFilter
    l3vpn.EntityData.YangName = "l3vpn"
    l3vpn.EntityData.BundleName = "cisco_ios_xr"
    l3vpn.EntityData.ParentYangName = "Cisco-IOS-XR-mpls-vpn-oper"
    l3vpn.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-vpn-oper:l3vpn"
    l3vpn.EntityData.AbsolutePath = l3vpn.EntityData.SegmentPath
    l3vpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l3vpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l3vpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l3vpn.EntityData.Children = types.NewOrderedMap()
    l3vpn.EntityData.Children.Append("invalid-vrfs", types.YChild{"InvalidVrfs", &l3vpn.InvalidVrfs})
    l3vpn.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &l3vpn.Vrfs})
    l3vpn.EntityData.Leafs = types.NewOrderedMap()

    l3vpn.EntityData.YListKeys = []string {}

    return &(l3vpn.EntityData)
}

// L3vpn_InvalidVrfs
// Invalid VRF Table (VRFs that are forward
// referenced)
type L3vpn_InvalidVrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Invalid VRF (VRF that is forward referenced). The type is slice of
    // L3vpn_InvalidVrfs_InvalidVrf.
    InvalidVrf []*L3vpn_InvalidVrfs_InvalidVrf
}

func (invalidVrfs *L3vpn_InvalidVrfs) GetEntityData() *types.CommonEntityData {
    invalidVrfs.EntityData.YFilter = invalidVrfs.YFilter
    invalidVrfs.EntityData.YangName = "invalid-vrfs"
    invalidVrfs.EntityData.BundleName = "cisco_ios_xr"
    invalidVrfs.EntityData.ParentYangName = "l3vpn"
    invalidVrfs.EntityData.SegmentPath = "invalid-vrfs"
    invalidVrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-vpn-oper:l3vpn/" + invalidVrfs.EntityData.SegmentPath
    invalidVrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    invalidVrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    invalidVrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    invalidVrfs.EntityData.Children = types.NewOrderedMap()
    invalidVrfs.EntityData.Children.Append("invalid-vrf", types.YChild{"InvalidVrf", nil})
    for i := range invalidVrfs.InvalidVrf {
        invalidVrfs.EntityData.Children.Append(types.GetSegmentPath(invalidVrfs.InvalidVrf[i]), types.YChild{"InvalidVrf", invalidVrfs.InvalidVrf[i]})
    }
    invalidVrfs.EntityData.Leafs = types.NewOrderedMap()

    invalidVrfs.EntityData.YListKeys = []string {}

    return &(invalidVrfs.EntityData)
}

// L3vpn_InvalidVrfs_InvalidVrf
// Invalid VRF (VRF that is forward referenced)
type L3vpn_InvalidVrfs_InvalidVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    // L3vpn_InvalidVrfs_InvalidVrf_Interface.
    Interface []*L3vpn_InvalidVrfs_InvalidVrf_Interface

    // AF/SAF information. The type is slice of L3vpn_InvalidVrfs_InvalidVrf_Af.
    Af []*L3vpn_InvalidVrfs_InvalidVrf_Af
}

func (invalidVrf *L3vpn_InvalidVrfs_InvalidVrf) GetEntityData() *types.CommonEntityData {
    invalidVrf.EntityData.YFilter = invalidVrf.YFilter
    invalidVrf.EntityData.YangName = "invalid-vrf"
    invalidVrf.EntityData.BundleName = "cisco_ios_xr"
    invalidVrf.EntityData.ParentYangName = "invalid-vrfs"
    invalidVrf.EntityData.SegmentPath = "invalid-vrf" + types.AddKeyToken(invalidVrf.VrfName, "vrf-name")
    invalidVrf.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-vpn-oper:l3vpn/invalid-vrfs/" + invalidVrf.EntityData.SegmentPath
    invalidVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    invalidVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    invalidVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    invalidVrf.EntityData.Children = types.NewOrderedMap()
    invalidVrf.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range invalidVrf.Interface {
        types.SetYListKey(invalidVrf.Interface[i], i)
        invalidVrf.EntityData.Children.Append(types.GetSegmentPath(invalidVrf.Interface[i]), types.YChild{"Interface", invalidVrf.Interface[i]})
    }
    invalidVrf.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range invalidVrf.Af {
        types.SetYListKey(invalidVrf.Af[i], i)
        invalidVrf.EntityData.Children.Append(types.GetSegmentPath(invalidVrf.Af[i]), types.YChild{"Af", invalidVrf.Af[i]})
    }
    invalidVrf.EntityData.Leafs = types.NewOrderedMap()
    invalidVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", invalidVrf.VrfName})
    invalidVrf.EntityData.Leafs.Append("vrf-name-xr", types.YLeaf{"VrfNameXr", invalidVrf.VrfNameXr})
    invalidVrf.EntityData.Leafs.Append("vrf-description", types.YLeaf{"VrfDescription", invalidVrf.VrfDescription})
    invalidVrf.EntityData.Leafs.Append("route-distinguisher", types.YLeaf{"RouteDistinguisher", invalidVrf.RouteDistinguisher})
    invalidVrf.EntityData.Leafs.Append("is-big-vrf", types.YLeaf{"IsBigVrf", invalidVrf.IsBigVrf})

    invalidVrf.EntityData.YListKeys = []string {"VrfName"}

    return &(invalidVrf.EntityData)
}

// L3vpn_InvalidVrfs_InvalidVrf_Interface
// Interfaces in VRF
type L3vpn_InvalidVrfs_InvalidVrf_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface Name. The type is string.
    InterfaceName interface{}
}

func (self *L3vpn_InvalidVrfs_InvalidVrf_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "invalid-vrf"
    self.EntityData.SegmentPath = "interface" + types.AddNoKeyToken(self)
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-vpn-oper:l3vpn/invalid-vrfs/invalid-vrf/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// L3vpn_InvalidVrfs_InvalidVrf_Af
// AF/SAF information
type L3vpn_InvalidVrfs_InvalidVrf_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AF name. The type is MplsVpnAfi.
    AfName interface{}

    // SAF name. The type is MplsVpnSafi.
    SafName interface{}

    // Import Route Policy. The type is string.
    ImportRoutePolicy interface{}

    // Export Route Policy. The type is string.
    ExportRoutePolicy interface{}

    // Route Targets. The type is slice of
    // L3vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget.
    RouteTarget []*L3vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget
}

func (af *L3vpn_InvalidVrfs_InvalidVrf_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "invalid-vrf"
    af.EntityData.SegmentPath = "af" + types.AddNoKeyToken(af)
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-vpn-oper:l3vpn/invalid-vrfs/invalid-vrf/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", nil})
    for i := range af.RouteTarget {
        types.SetYListKey(af.RouteTarget[i], i)
        af.EntityData.Children.Append(types.GetSegmentPath(af.RouteTarget[i]), types.YChild{"RouteTarget", af.RouteTarget[i]})
    }
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", af.SafName})
    af.EntityData.Leafs.Append("import-route-policy", types.YLeaf{"ImportRoutePolicy", af.ImportRoutePolicy})
    af.EntityData.Leafs.Append("export-route-policy", types.YLeaf{"ExportRoutePolicy", af.ExportRoutePolicy})

    af.EntityData.YListKeys = []string {}

    return &(af.EntityData)
}

// L3vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget
// Route Targets
type L3vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Route Target Type. The type is MplsVpnRt.
    RouteTargetType interface{}

    // Route Target Value. The type is string.
    RouteTargetValue interface{}

    // AF name. The type is MplsVpnAfi.
    AfName interface{}

    // SAF name. The type is MplsVpnSafi.
    SafName interface{}
}

func (routeTarget *L3vpn_InvalidVrfs_InvalidVrf_Af_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "af"
    routeTarget.EntityData.SegmentPath = "route-target" + types.AddNoKeyToken(routeTarget)
    routeTarget.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-vpn-oper:l3vpn/invalid-vrfs/invalid-vrf/af/" + routeTarget.EntityData.SegmentPath
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("route-target-type", types.YLeaf{"RouteTargetType", routeTarget.RouteTargetType})
    routeTarget.EntityData.Leafs.Append("route-target-value", types.YLeaf{"RouteTargetValue", routeTarget.RouteTargetValue})
    routeTarget.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", routeTarget.AfName})
    routeTarget.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", routeTarget.SafName})

    routeTarget.EntityData.YListKeys = []string {}

    return &(routeTarget.EntityData)
}

// L3vpn_Vrfs
// VRF Table
type L3vpn_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF. The type is slice of L3vpn_Vrfs_Vrf.
    Vrf []*L3vpn_Vrfs_Vrf
}

func (vrfs *L3vpn_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "l3vpn"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-vpn-oper:l3vpn/" + vrfs.EntityData.SegmentPath
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

// L3vpn_Vrfs_Vrf
// VRF
type L3vpn_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Interfaces in VRF. The type is slice of L3vpn_Vrfs_Vrf_Interface.
    Interface []*L3vpn_Vrfs_Vrf_Interface

    // AF/SAF information. The type is slice of L3vpn_Vrfs_Vrf_Af.
    Af []*L3vpn_Vrfs_Vrf_Af
}

func (vrf *L3vpn_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-vpn-oper:l3vpn/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range vrf.Interface {
        types.SetYListKey(vrf.Interface[i], i)
        vrf.EntityData.Children.Append(types.GetSegmentPath(vrf.Interface[i]), types.YChild{"Interface", vrf.Interface[i]})
    }
    vrf.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range vrf.Af {
        types.SetYListKey(vrf.Af[i], i)
        vrf.EntityData.Children.Append(types.GetSegmentPath(vrf.Af[i]), types.YChild{"Af", vrf.Af[i]})
    }
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("vrf-name-xr", types.YLeaf{"VrfNameXr", vrf.VrfNameXr})
    vrf.EntityData.Leafs.Append("vrf-description", types.YLeaf{"VrfDescription", vrf.VrfDescription})
    vrf.EntityData.Leafs.Append("route-distinguisher", types.YLeaf{"RouteDistinguisher", vrf.RouteDistinguisher})
    vrf.EntityData.Leafs.Append("is-big-vrf", types.YLeaf{"IsBigVrf", vrf.IsBigVrf})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// L3vpn_Vrfs_Vrf_Interface
// Interfaces in VRF
type L3vpn_Vrfs_Vrf_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface Name. The type is string.
    InterfaceName interface{}
}

func (self *L3vpn_Vrfs_Vrf_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "vrf"
    self.EntityData.SegmentPath = "interface" + types.AddNoKeyToken(self)
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-vpn-oper:l3vpn/vrfs/vrf/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// L3vpn_Vrfs_Vrf_Af
// AF/SAF information
type L3vpn_Vrfs_Vrf_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AF name. The type is MplsVpnAfi.
    AfName interface{}

    // SAF name. The type is MplsVpnSafi.
    SafName interface{}

    // Import Route Policy. The type is string.
    ImportRoutePolicy interface{}

    // Export Route Policy. The type is string.
    ExportRoutePolicy interface{}

    // Route Targets. The type is slice of L3vpn_Vrfs_Vrf_Af_RouteTarget.
    RouteTarget []*L3vpn_Vrfs_Vrf_Af_RouteTarget
}

func (af *L3vpn_Vrfs_Vrf_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "vrf"
    af.EntityData.SegmentPath = "af" + types.AddNoKeyToken(af)
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-vpn-oper:l3vpn/vrfs/vrf/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", nil})
    for i := range af.RouteTarget {
        types.SetYListKey(af.RouteTarget[i], i)
        af.EntityData.Children.Append(types.GetSegmentPath(af.RouteTarget[i]), types.YChild{"RouteTarget", af.RouteTarget[i]})
    }
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", af.SafName})
    af.EntityData.Leafs.Append("import-route-policy", types.YLeaf{"ImportRoutePolicy", af.ImportRoutePolicy})
    af.EntityData.Leafs.Append("export-route-policy", types.YLeaf{"ExportRoutePolicy", af.ExportRoutePolicy})

    af.EntityData.YListKeys = []string {}

    return &(af.EntityData)
}

// L3vpn_Vrfs_Vrf_Af_RouteTarget
// Route Targets
type L3vpn_Vrfs_Vrf_Af_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Route Target Type. The type is MplsVpnRt.
    RouteTargetType interface{}

    // Route Target Value. The type is string.
    RouteTargetValue interface{}

    // AF name. The type is MplsVpnAfi.
    AfName interface{}

    // SAF name. The type is MplsVpnSafi.
    SafName interface{}
}

func (routeTarget *L3vpn_Vrfs_Vrf_Af_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "af"
    routeTarget.EntityData.SegmentPath = "route-target" + types.AddNoKeyToken(routeTarget)
    routeTarget.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-vpn-oper:l3vpn/vrfs/vrf/af/" + routeTarget.EntityData.SegmentPath
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("route-target-type", types.YLeaf{"RouteTargetType", routeTarget.RouteTargetType})
    routeTarget.EntityData.Leafs.Append("route-target-value", types.YLeaf{"RouteTargetValue", routeTarget.RouteTargetValue})
    routeTarget.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", routeTarget.AfName})
    routeTarget.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", routeTarget.SafName})

    routeTarget.EntityData.YListKeys = []string {}

    return &(routeTarget.EntityData)
}

