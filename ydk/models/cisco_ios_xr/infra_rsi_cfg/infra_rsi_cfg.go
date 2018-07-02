// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-rsi package configuration.
// 
// This module contains definitions
// for the following management objects:
//   vrfs: VRF configuration
//   global-af: global af
//   srlg: srlg
//   vrf-groups: vrf groups
//   selective-vrf-download: selective vrf download
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg,
//   Cisco-IOS-XR-snmp-agent-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_rsi_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_rsi_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rsi-cfg vrfs}", reflect.TypeOf(Vrfs{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rsi-cfg:vrfs", reflect.TypeOf(Vrfs{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rsi-cfg global-af}", reflect.TypeOf(GlobalAf{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rsi-cfg:global-af", reflect.TypeOf(GlobalAf{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rsi-cfg srlg}", reflect.TypeOf(Srlg{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rsi-cfg:srlg", reflect.TypeOf(Srlg{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rsi-cfg vrf-groups}", reflect.TypeOf(VrfGroups{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rsi-cfg:vrf-groups", reflect.TypeOf(VrfGroups{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rsi-cfg selective-vrf-download}", reflect.TypeOf(SelectiveVrfDownload{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rsi-cfg:selective-vrf-download", reflect.TypeOf(SelectiveVrfDownload{}))
}

// VrfAddressFamily represents Vrf address family
type VrfAddressFamily string

const (
    // IPv4
    VrfAddressFamily_ipv4 VrfAddressFamily = "ipv4"

    // IPv6
    VrfAddressFamily_ipv6 VrfAddressFamily = "ipv6"
)

// SrlgPriority represents Srlg priority
type SrlgPriority string

const (
    // Critical
    SrlgPriority_critical SrlgPriority = "critical"

    // High
    SrlgPriority_high SrlgPriority = "high"

    // Default
    SrlgPriority_default_ SrlgPriority = "default"

    // Low
    SrlgPriority_low SrlgPriority = "low"

    // Very low
    SrlgPriority_very_low SrlgPriority = "very-low"
)

// VrfSubAddressFamily represents Vrf sub address family
type VrfSubAddressFamily string

const (
    // Unicast
    VrfSubAddressFamily_unicast VrfSubAddressFamily = "unicast"

    // Multicast
    VrfSubAddressFamily_multicast VrfSubAddressFamily = "multicast"

    // Flow spec
    VrfSubAddressFamily_flow_spec VrfSubAddressFamily = "flow-spec"
)

// Vrfs
// VRF configuration
type Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF configuration. The type is slice of Vrfs_Vrf.
    Vrf []*Vrfs_Vrf
}

func (vrfs *Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "Cisco-IOS-XR-infra-rsi-cfg"
    vrfs.EntityData.SegmentPath = "Cisco-IOS-XR-infra-rsi-cfg:vrfs"
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

// Vrfs_Vrf
// VRF configuration
type Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // Fallback VRF. The type is string with length: 1..32.
    FallbackVrf interface{}

    // For disabling remote route filtering for this VRF on core-facing card. The
    // type is interface{}.
    RemoteRouteFilterDisable interface{}

    // VRF global configuration. The type is interface{}.
    Create interface{}

    // Configuration enable of big VRF. The type is interface{}.
    ModeBig interface{}

    // A textual description of the VRF. The type is string with length: 1..244.
    Description interface{}

    // VPN-ID for the VRF.
    VpnId Vrfs_Vrf_VpnId

    // VRF address family configuration.
    Afs Vrfs_Vrf_Afs

    // BGP related VRF GBL config.
    BgpGlobal Vrfs_Vrf_BgpGlobal

    // Multicast host stack configuration.
    MulticastHost Vrfs_Vrf_MulticastHost
}

func (vrf *Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("vpn-id", types.YChild{"VpnId", &vrf.VpnId})
    vrf.EntityData.Children.Append("afs", types.YChild{"Afs", &vrf.Afs})
    vrf.EntityData.Children.Append("Cisco-IOS-XR-ipv4-bgp-cfg:bgp-global", types.YChild{"BgpGlobal", &vrf.BgpGlobal})
    vrf.EntityData.Children.Append("Cisco-IOS-XR-ip-iarm-vrf-cfg:multicast-host", types.YChild{"MulticastHost", &vrf.MulticastHost})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("fallback-vrf", types.YLeaf{"FallbackVrf", vrf.FallbackVrf})
    vrf.EntityData.Leafs.Append("remote-route-filter-disable", types.YLeaf{"RemoteRouteFilterDisable", vrf.RemoteRouteFilterDisable})
    vrf.EntityData.Leafs.Append("create", types.YLeaf{"Create", vrf.Create})
    vrf.EntityData.Leafs.Append("mode-big", types.YLeaf{"ModeBig", vrf.ModeBig})
    vrf.EntityData.Leafs.Append("description", types.YLeaf{"Description", vrf.Description})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Vrfs_Vrf_VpnId
// VPN-ID for the VRF
// This type is a presence type.
type Vrfs_Vrf_VpnId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // OUI of VPNID OUI. The type is interface{} with range: 0..16777215. This
    // attribute is mandatory.
    VpnOui interface{}

    // Index of VPNID Index. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    VpnIndex interface{}
}

func (vpnId *Vrfs_Vrf_VpnId) GetEntityData() *types.CommonEntityData {
    vpnId.EntityData.YFilter = vpnId.YFilter
    vpnId.EntityData.YangName = "vpn-id"
    vpnId.EntityData.BundleName = "cisco_ios_xr"
    vpnId.EntityData.ParentYangName = "vrf"
    vpnId.EntityData.SegmentPath = "vpn-id"
    vpnId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpnId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpnId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpnId.EntityData.Children = types.NewOrderedMap()
    vpnId.EntityData.Leafs = types.NewOrderedMap()
    vpnId.EntityData.Leafs.Append("vpn-oui", types.YLeaf{"VpnOui", vpnId.VpnOui})
    vpnId.EntityData.Leafs.Append("vpn-index", types.YLeaf{"VpnIndex", vpnId.VpnIndex})

    vpnId.EntityData.YListKeys = []string {}

    return &(vpnId.EntityData)
}

// Vrfs_Vrf_Afs
// VRF address family configuration
type Vrfs_Vrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF address family configuration. The type is slice of Vrfs_Vrf_Afs_Af.
    Af []*Vrfs_Vrf_Afs_Af
}

func (afs *Vrfs_Vrf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "vrf"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// Vrfs_Vrf_Afs_Af
// VRF address family configuration
type Vrfs_Vrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address family. The type is VrfAddressFamily.
    AfName interface{}

    // This attribute is a key. Sub-Address family. The type is
    // VrfSubAddressFamily.
    SafName interface{}

    // This attribute is a key. Topology name. The type is string with length:
    // 1..244.
    TopologyName interface{}

    // VRF configuration for a particular address family. The type is interface{}.
    Create interface{}

    // BGP AF VRF config.
    Bgp Vrfs_Vrf_Afs_Af_Bgp

    // Set maximum prefix limits.
    MaximumPrefix Vrfs_Vrf_Afs_Af_MaximumPrefix
}

func (af *Vrfs_Vrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name") + types.AddKeyToken(af.SafName, "saf-name") + types.AddKeyToken(af.TopologyName, "topology-name")
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("Cisco-IOS-XR-ipv4-bgp-cfg:bgp", types.YChild{"Bgp", &af.Bgp})
    af.EntityData.Children.Append("Cisco-IOS-XR-ip-rib-cfg:maximum-prefix", types.YChild{"MaximumPrefix", &af.MaximumPrefix})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", af.SafName})
    af.EntityData.Leafs.Append("topology-name", types.YLeaf{"TopologyName", af.TopologyName})
    af.EntityData.Leafs.Append("create", types.YLeaf{"Create", af.Create})

    af.EntityData.YListKeys = []string {"AfName", "SafName", "TopologyName"}

    return &(af.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp
// BGP AF VRF config
type Vrfs_Vrf_Afs_Af_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route policy for export filtering. The type is string.
    ExportRoutePolicy interface{}

    // Route policy for import filtering. The type is string.
    ImportRoutePolicy interface{}

    // TRUE Enable advertising imported paths to PEsFALSE Disable advertising
    // imported paths to PEs. The type is bool.
    ImportVrfOptions interface{}

    // Import Route targets.
    ImportRouteTargets Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets

    // Export Route targets.
    ExportRouteTargets Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets

    // Route policy for vrf to global export filtering.
    VrfToGlobalExportRoutePolicy Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy

    // Export VRF options.
    ExportVrfOptions Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions

    // Route policy for global to vrf import filtering.
    GlobalToVrfImportRoutePolicy Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy
}

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "af"
    bgp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-bgp-cfg:bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Children.Append("import-route-targets", types.YChild{"ImportRouteTargets", &bgp.ImportRouteTargets})
    bgp.EntityData.Children.Append("export-route-targets", types.YChild{"ExportRouteTargets", &bgp.ExportRouteTargets})
    bgp.EntityData.Children.Append("vrf-to-global-export-route-policy", types.YChild{"VrfToGlobalExportRoutePolicy", &bgp.VrfToGlobalExportRoutePolicy})
    bgp.EntityData.Children.Append("export-vrf-options", types.YChild{"ExportVrfOptions", &bgp.ExportVrfOptions})
    bgp.EntityData.Children.Append("global-to-vrf-import-route-policy", types.YChild{"GlobalToVrfImportRoutePolicy", &bgp.GlobalToVrfImportRoutePolicy})
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("export-route-policy", types.YLeaf{"ExportRoutePolicy", bgp.ExportRoutePolicy})
    bgp.EntityData.Leafs.Append("import-route-policy", types.YLeaf{"ImportRoutePolicy", bgp.ImportRoutePolicy})
    bgp.EntityData.Leafs.Append("import-vrf-options", types.YLeaf{"ImportVrfOptions", bgp.ImportVrfOptions})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets
// Import Route targets
type Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route target table.
    RouteTargets Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets
}

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetEntityData() *types.CommonEntityData {
    importRouteTargets.EntityData.YFilter = importRouteTargets.YFilter
    importRouteTargets.EntityData.YangName = "import-route-targets"
    importRouteTargets.EntityData.BundleName = "cisco_ios_xr"
    importRouteTargets.EntityData.ParentYangName = "bgp"
    importRouteTargets.EntityData.SegmentPath = "import-route-targets"
    importRouteTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    importRouteTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    importRouteTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    importRouteTargets.EntityData.Children = types.NewOrderedMap()
    importRouteTargets.EntityData.Children.Append("route-targets", types.YChild{"RouteTargets", &importRouteTargets.RouteTargets})
    importRouteTargets.EntityData.Leafs = types.NewOrderedMap()

    importRouteTargets.EntityData.YListKeys = []string {}

    return &(importRouteTargets.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets
// Route target table
type Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route target. The type is slice of
    // Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget.
    RouteTarget []*Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetEntityData() *types.CommonEntityData {
    routeTargets.EntityData.YFilter = routeTargets.YFilter
    routeTargets.EntityData.YangName = "route-targets"
    routeTargets.EntityData.BundleName = "cisco_ios_xr"
    routeTargets.EntityData.ParentYangName = "import-route-targets"
    routeTargets.EntityData.SegmentPath = "route-targets"
    routeTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTargets.EntityData.Children = types.NewOrderedMap()
    routeTargets.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", nil})
    for i := range routeTargets.RouteTarget {
        routeTargets.EntityData.Children.Append(types.GetSegmentPath(routeTargets.RouteTarget[i]), types.YChild{"RouteTarget", routeTargets.RouteTarget[i]})
    }
    routeTargets.EntityData.Leafs = types.NewOrderedMap()

    routeTargets.EntityData.YListKeys = []string {}

    return &(routeTargets.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget
// Route target
type Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Type of RT. The type is BgpVrfRouteTarget.
    Type interface{}

    // as or four byte as. The type is slice of
    // Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs.
    AsOrFourByteAs []*Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs

    // ipv4 address. The type is slice of
    // Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address.
    Ipv4Address []*Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "route-targets"
    routeTarget.EntityData.SegmentPath = "route-target" + types.AddKeyToken(routeTarget.Type, "type")
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Children.Append("as-or-four-byte-as", types.YChild{"AsOrFourByteAs", nil})
    for i := range routeTarget.AsOrFourByteAs {
        routeTarget.EntityData.Children.Append(types.GetSegmentPath(routeTarget.AsOrFourByteAs[i]), types.YChild{"AsOrFourByteAs", routeTarget.AsOrFourByteAs[i]})
    }
    routeTarget.EntityData.Children.Append("ipv4-address", types.YChild{"Ipv4Address", nil})
    for i := range routeTarget.Ipv4Address {
        routeTarget.EntityData.Children.Append(types.GetSegmentPath(routeTarget.Ipv4Address[i]), types.YChild{"Ipv4Address", routeTarget.Ipv4Address[i]})
    }
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("type", types.YLeaf{"Type", routeTarget.Type})

    routeTarget.EntityData.YListKeys = []string {"Type"}

    return &(routeTarget.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs
// as or four byte as
type Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. AS number. The type is interface{} with range:
    // 0..4294967295.
    AsXx interface{}

    // This attribute is a key. AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // This attribute is a key. AS number Index. The type is interface{} with
    // range: 0..4294967295.
    AsIndex interface{}

    // This attribute is a key. Stitching RT. The type is interface{} with range:
    // 0..1.
    StitchingRt interface{}
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetEntityData() *types.CommonEntityData {
    asOrFourByteAs.EntityData.YFilter = asOrFourByteAs.YFilter
    asOrFourByteAs.EntityData.YangName = "as-or-four-byte-as"
    asOrFourByteAs.EntityData.BundleName = "cisco_ios_xr"
    asOrFourByteAs.EntityData.ParentYangName = "route-target"
    asOrFourByteAs.EntityData.SegmentPath = "as-or-four-byte-as" + types.AddKeyToken(asOrFourByteAs.AsXx, "as-xx") + types.AddKeyToken(asOrFourByteAs.As, "as") + types.AddKeyToken(asOrFourByteAs.AsIndex, "as-index") + types.AddKeyToken(asOrFourByteAs.StitchingRt, "stitching-rt")
    asOrFourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asOrFourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asOrFourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asOrFourByteAs.EntityData.Children = types.NewOrderedMap()
    asOrFourByteAs.EntityData.Leafs = types.NewOrderedMap()
    asOrFourByteAs.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", asOrFourByteAs.AsXx})
    asOrFourByteAs.EntityData.Leafs.Append("as", types.YLeaf{"As", asOrFourByteAs.As})
    asOrFourByteAs.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", asOrFourByteAs.AsIndex})
    asOrFourByteAs.EntityData.Leafs.Append("stitching-rt", types.YLeaf{"StitchingRt", asOrFourByteAs.StitchingRt})

    asOrFourByteAs.EntityData.YListKeys = []string {"AsXx", "As", "AsIndex", "StitchingRt"}

    return &(asOrFourByteAs.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address
// ipv4 address
type Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IP address Index. The type is interface{} with
    // range: 0..65535.
    AddressIndex interface{}

    // This attribute is a key. Stitching RT. The type is interface{} with range:
    // 0..1.
    StitchingRt interface{}
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetEntityData() *types.CommonEntityData {
    ipv4Address.EntityData.YFilter = ipv4Address.YFilter
    ipv4Address.EntityData.YangName = "ipv4-address"
    ipv4Address.EntityData.BundleName = "cisco_ios_xr"
    ipv4Address.EntityData.ParentYangName = "route-target"
    ipv4Address.EntityData.SegmentPath = "ipv4-address" + types.AddKeyToken(ipv4Address.Address, "address") + types.AddKeyToken(ipv4Address.AddressIndex, "address-index") + types.AddKeyToken(ipv4Address.StitchingRt, "stitching-rt")
    ipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Address.EntityData.Children = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4Address.Address})
    ipv4Address.EntityData.Leafs.Append("address-index", types.YLeaf{"AddressIndex", ipv4Address.AddressIndex})
    ipv4Address.EntityData.Leafs.Append("stitching-rt", types.YLeaf{"StitchingRt", ipv4Address.StitchingRt})

    ipv4Address.EntityData.YListKeys = []string {"Address", "AddressIndex", "StitchingRt"}

    return &(ipv4Address.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets
// Export Route targets
type Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route target table.
    RouteTargets Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets
}

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetEntityData() *types.CommonEntityData {
    exportRouteTargets.EntityData.YFilter = exportRouteTargets.YFilter
    exportRouteTargets.EntityData.YangName = "export-route-targets"
    exportRouteTargets.EntityData.BundleName = "cisco_ios_xr"
    exportRouteTargets.EntityData.ParentYangName = "bgp"
    exportRouteTargets.EntityData.SegmentPath = "export-route-targets"
    exportRouteTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exportRouteTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exportRouteTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exportRouteTargets.EntityData.Children = types.NewOrderedMap()
    exportRouteTargets.EntityData.Children.Append("route-targets", types.YChild{"RouteTargets", &exportRouteTargets.RouteTargets})
    exportRouteTargets.EntityData.Leafs = types.NewOrderedMap()

    exportRouteTargets.EntityData.YListKeys = []string {}

    return &(exportRouteTargets.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets
// Route target table
type Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route target. The type is slice of
    // Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget.
    RouteTarget []*Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetEntityData() *types.CommonEntityData {
    routeTargets.EntityData.YFilter = routeTargets.YFilter
    routeTargets.EntityData.YangName = "route-targets"
    routeTargets.EntityData.BundleName = "cisco_ios_xr"
    routeTargets.EntityData.ParentYangName = "export-route-targets"
    routeTargets.EntityData.SegmentPath = "route-targets"
    routeTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTargets.EntityData.Children = types.NewOrderedMap()
    routeTargets.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", nil})
    for i := range routeTargets.RouteTarget {
        routeTargets.EntityData.Children.Append(types.GetSegmentPath(routeTargets.RouteTarget[i]), types.YChild{"RouteTarget", routeTargets.RouteTarget[i]})
    }
    routeTargets.EntityData.Leafs = types.NewOrderedMap()

    routeTargets.EntityData.YListKeys = []string {}

    return &(routeTargets.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget
// Route target
type Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Type of RT. The type is BgpVrfRouteTarget.
    Type interface{}

    // as or four byte as. The type is slice of
    // Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs.
    AsOrFourByteAs []*Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs

    // ipv4 address. The type is slice of
    // Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address.
    Ipv4Address []*Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "route-targets"
    routeTarget.EntityData.SegmentPath = "route-target" + types.AddKeyToken(routeTarget.Type, "type")
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Children.Append("as-or-four-byte-as", types.YChild{"AsOrFourByteAs", nil})
    for i := range routeTarget.AsOrFourByteAs {
        routeTarget.EntityData.Children.Append(types.GetSegmentPath(routeTarget.AsOrFourByteAs[i]), types.YChild{"AsOrFourByteAs", routeTarget.AsOrFourByteAs[i]})
    }
    routeTarget.EntityData.Children.Append("ipv4-address", types.YChild{"Ipv4Address", nil})
    for i := range routeTarget.Ipv4Address {
        routeTarget.EntityData.Children.Append(types.GetSegmentPath(routeTarget.Ipv4Address[i]), types.YChild{"Ipv4Address", routeTarget.Ipv4Address[i]})
    }
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("type", types.YLeaf{"Type", routeTarget.Type})

    routeTarget.EntityData.YListKeys = []string {"Type"}

    return &(routeTarget.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs
// as or four byte as
type Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. AS number. The type is interface{} with range:
    // 0..4294967295.
    AsXx interface{}

    // This attribute is a key. AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // This attribute is a key. AS number Index. The type is interface{} with
    // range: 0..4294967295.
    AsIndex interface{}

    // This attribute is a key. Stitching RT. The type is interface{} with range:
    // 0..1.
    StitchingRt interface{}
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetEntityData() *types.CommonEntityData {
    asOrFourByteAs.EntityData.YFilter = asOrFourByteAs.YFilter
    asOrFourByteAs.EntityData.YangName = "as-or-four-byte-as"
    asOrFourByteAs.EntityData.BundleName = "cisco_ios_xr"
    asOrFourByteAs.EntityData.ParentYangName = "route-target"
    asOrFourByteAs.EntityData.SegmentPath = "as-or-four-byte-as" + types.AddKeyToken(asOrFourByteAs.AsXx, "as-xx") + types.AddKeyToken(asOrFourByteAs.As, "as") + types.AddKeyToken(asOrFourByteAs.AsIndex, "as-index") + types.AddKeyToken(asOrFourByteAs.StitchingRt, "stitching-rt")
    asOrFourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asOrFourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asOrFourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asOrFourByteAs.EntityData.Children = types.NewOrderedMap()
    asOrFourByteAs.EntityData.Leafs = types.NewOrderedMap()
    asOrFourByteAs.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", asOrFourByteAs.AsXx})
    asOrFourByteAs.EntityData.Leafs.Append("as", types.YLeaf{"As", asOrFourByteAs.As})
    asOrFourByteAs.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", asOrFourByteAs.AsIndex})
    asOrFourByteAs.EntityData.Leafs.Append("stitching-rt", types.YLeaf{"StitchingRt", asOrFourByteAs.StitchingRt})

    asOrFourByteAs.EntityData.YListKeys = []string {"AsXx", "As", "AsIndex", "StitchingRt"}

    return &(asOrFourByteAs.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address
// ipv4 address
type Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IP address Index. The type is interface{} with
    // range: 0..65535.
    AddressIndex interface{}

    // This attribute is a key. Stitching RT. The type is interface{} with range:
    // 0..1.
    StitchingRt interface{}
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetEntityData() *types.CommonEntityData {
    ipv4Address.EntityData.YFilter = ipv4Address.YFilter
    ipv4Address.EntityData.YangName = "ipv4-address"
    ipv4Address.EntityData.BundleName = "cisco_ios_xr"
    ipv4Address.EntityData.ParentYangName = "route-target"
    ipv4Address.EntityData.SegmentPath = "ipv4-address" + types.AddKeyToken(ipv4Address.Address, "address") + types.AddKeyToken(ipv4Address.AddressIndex, "address-index") + types.AddKeyToken(ipv4Address.StitchingRt, "stitching-rt")
    ipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Address.EntityData.Children = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4Address.Address})
    ipv4Address.EntityData.Leafs.Append("address-index", types.YLeaf{"AddressIndex", ipv4Address.AddressIndex})
    ipv4Address.EntityData.Leafs.Append("stitching-rt", types.YLeaf{"StitchingRt", ipv4Address.StitchingRt})

    ipv4Address.EntityData.YListKeys = []string {"Address", "AddressIndex", "StitchingRt"}

    return &(ipv4Address.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy
// Route policy for vrf to global export filtering
// This type is a presence type.
type Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Vrf to global export route policy. The type is string. This attribute is
    // mandatory.
    RoutePolicyName interface{}

    // TRUE Enable imported VPN paths to be exported to Default VRF.FALSE Disable
    // imported VPN paths to be exported to Default VRF. The type is bool.
    AllowImportedVpn interface{}
}

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetEntityData() *types.CommonEntityData {
    vrfToGlobalExportRoutePolicy.EntityData.YFilter = vrfToGlobalExportRoutePolicy.YFilter
    vrfToGlobalExportRoutePolicy.EntityData.YangName = "vrf-to-global-export-route-policy"
    vrfToGlobalExportRoutePolicy.EntityData.BundleName = "cisco_ios_xr"
    vrfToGlobalExportRoutePolicy.EntityData.ParentYangName = "bgp"
    vrfToGlobalExportRoutePolicy.EntityData.SegmentPath = "vrf-to-global-export-route-policy"
    vrfToGlobalExportRoutePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfToGlobalExportRoutePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfToGlobalExportRoutePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfToGlobalExportRoutePolicy.EntityData.Children = types.NewOrderedMap()
    vrfToGlobalExportRoutePolicy.EntityData.Leafs = types.NewOrderedMap()
    vrfToGlobalExportRoutePolicy.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", vrfToGlobalExportRoutePolicy.RoutePolicyName})
    vrfToGlobalExportRoutePolicy.EntityData.Leafs.Append("allow-imported-vpn", types.YLeaf{"AllowImportedVpn", vrfToGlobalExportRoutePolicy.AllowImportedVpn})

    vrfToGlobalExportRoutePolicy.EntityData.YListKeys = []string {}

    return &(vrfToGlobalExportRoutePolicy.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions
// Export VRF options
type Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE Enable imported VPN paths to be exported to non-default VRFFALSE
    // Disable imported VPN paths to be exported to non-default VRF. The type is
    // bool.
    AllowImportedVpn interface{}

    // TRUE Use stitchng RTs to import extranet pathsFALSE Use regular RTs to
    // import extranet paths. The type is bool.
    ImportStitchingRt interface{}
}

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetEntityData() *types.CommonEntityData {
    exportVrfOptions.EntityData.YFilter = exportVrfOptions.YFilter
    exportVrfOptions.EntityData.YangName = "export-vrf-options"
    exportVrfOptions.EntityData.BundleName = "cisco_ios_xr"
    exportVrfOptions.EntityData.ParentYangName = "bgp"
    exportVrfOptions.EntityData.SegmentPath = "export-vrf-options"
    exportVrfOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exportVrfOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exportVrfOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exportVrfOptions.EntityData.Children = types.NewOrderedMap()
    exportVrfOptions.EntityData.Leafs = types.NewOrderedMap()
    exportVrfOptions.EntityData.Leafs.Append("allow-imported-vpn", types.YLeaf{"AllowImportedVpn", exportVrfOptions.AllowImportedVpn})
    exportVrfOptions.EntityData.Leafs.Append("import-stitching-rt", types.YLeaf{"ImportStitchingRt", exportVrfOptions.ImportStitchingRt})

    exportVrfOptions.EntityData.YListKeys = []string {}

    return &(exportVrfOptions.EntityData)
}

// Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy
// Route policy for global to vrf import filtering
// This type is a presence type.
type Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Global to vrf import route policy. The type is string. This attribute is
    // mandatory.
    RoutePolicyName interface{}

    // TRUE Enable advertising imported paths to PEsFALSE Disable advertising
    // imported paths to PEs. The type is bool.
    AdvertiseAsVpn interface{}
}

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetEntityData() *types.CommonEntityData {
    globalToVrfImportRoutePolicy.EntityData.YFilter = globalToVrfImportRoutePolicy.YFilter
    globalToVrfImportRoutePolicy.EntityData.YangName = "global-to-vrf-import-route-policy"
    globalToVrfImportRoutePolicy.EntityData.BundleName = "cisco_ios_xr"
    globalToVrfImportRoutePolicy.EntityData.ParentYangName = "bgp"
    globalToVrfImportRoutePolicy.EntityData.SegmentPath = "global-to-vrf-import-route-policy"
    globalToVrfImportRoutePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalToVrfImportRoutePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalToVrfImportRoutePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalToVrfImportRoutePolicy.EntityData.Children = types.NewOrderedMap()
    globalToVrfImportRoutePolicy.EntityData.Leafs = types.NewOrderedMap()
    globalToVrfImportRoutePolicy.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", globalToVrfImportRoutePolicy.RoutePolicyName})
    globalToVrfImportRoutePolicy.EntityData.Leafs.Append("advertise-as-vpn", types.YLeaf{"AdvertiseAsVpn", globalToVrfImportRoutePolicy.AdvertiseAsVpn})

    globalToVrfImportRoutePolicy.EntityData.YListKeys = []string {}

    return &(globalToVrfImportRoutePolicy.EntityData)
}

// Vrfs_Vrf_Afs_Af_MaximumPrefix
// Set maximum prefix limits
// This type is a presence type.
type Vrfs_Vrf_Afs_Af_MaximumPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Set table's maximum prefix limit. The type is interface{} with range:
    // 32..10000000. This attribute is mandatory.
    PrefixLimit interface{}

    // Mid-threshold (% of maximum). The type is interface{} with range: 1..100.
    MidThreshold interface{}
}

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetEntityData() *types.CommonEntityData {
    maximumPrefix.EntityData.YFilter = maximumPrefix.YFilter
    maximumPrefix.EntityData.YangName = "maximum-prefix"
    maximumPrefix.EntityData.BundleName = "cisco_ios_xr"
    maximumPrefix.EntityData.ParentYangName = "af"
    maximumPrefix.EntityData.SegmentPath = "Cisco-IOS-XR-ip-rib-cfg:maximum-prefix"
    maximumPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumPrefix.EntityData.Children = types.NewOrderedMap()
    maximumPrefix.EntityData.Leafs = types.NewOrderedMap()
    maximumPrefix.EntityData.Leafs.Append("prefix-limit", types.YLeaf{"PrefixLimit", maximumPrefix.PrefixLimit})
    maximumPrefix.EntityData.Leafs.Append("mid-threshold", types.YLeaf{"MidThreshold", maximumPrefix.MidThreshold})

    maximumPrefix.EntityData.YListKeys = []string {}

    return &(maximumPrefix.EntityData)
}

// Vrfs_Vrf_BgpGlobal
// BGP related VRF GBL config
type Vrfs_Vrf_BgpGlobal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global Route distinguisher.
    RouteDistinguisher Vrfs_Vrf_BgpGlobal_RouteDistinguisher
}

func (bgpGlobal *Vrfs_Vrf_BgpGlobal) GetEntityData() *types.CommonEntityData {
    bgpGlobal.EntityData.YFilter = bgpGlobal.YFilter
    bgpGlobal.EntityData.YangName = "bgp-global"
    bgpGlobal.EntityData.BundleName = "cisco_ios_xr"
    bgpGlobal.EntityData.ParentYangName = "vrf"
    bgpGlobal.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-bgp-cfg:bgp-global"
    bgpGlobal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpGlobal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpGlobal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpGlobal.EntityData.Children = types.NewOrderedMap()
    bgpGlobal.EntityData.Children.Append("route-distinguisher", types.YChild{"RouteDistinguisher", &bgpGlobal.RouteDistinguisher})
    bgpGlobal.EntityData.Leafs = types.NewOrderedMap()

    bgpGlobal.EntityData.YListKeys = []string {}

    return &(bgpGlobal.EntityData)
}

// Vrfs_Vrf_BgpGlobal_RouteDistinguisher
// Global Route distinguisher
type Vrfs_Vrf_BgpGlobal_RouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of RD. The type is BgpGlobalRouteDistinguisher.
    Type interface{}

    // AS number. The type is interface{} with range: 0..4294967295.
    AsXx interface{}

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // ASN Index. The type is interface{} with range: 0..4294967295.
    AsIndex interface{}

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // IP address index. The type is interface{} with range: 0..65535.
    AddressIndex interface{}
}

func (routeDistinguisher *Vrfs_Vrf_BgpGlobal_RouteDistinguisher) GetEntityData() *types.CommonEntityData {
    routeDistinguisher.EntityData.YFilter = routeDistinguisher.YFilter
    routeDistinguisher.EntityData.YangName = "route-distinguisher"
    routeDistinguisher.EntityData.BundleName = "cisco_ios_xr"
    routeDistinguisher.EntityData.ParentYangName = "bgp-global"
    routeDistinguisher.EntityData.SegmentPath = "route-distinguisher"
    routeDistinguisher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeDistinguisher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeDistinguisher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeDistinguisher.EntityData.Children = types.NewOrderedMap()
    routeDistinguisher.EntityData.Leafs = types.NewOrderedMap()
    routeDistinguisher.EntityData.Leafs.Append("type", types.YLeaf{"Type", routeDistinguisher.Type})
    routeDistinguisher.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", routeDistinguisher.AsXx})
    routeDistinguisher.EntityData.Leafs.Append("as", types.YLeaf{"As", routeDistinguisher.As})
    routeDistinguisher.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", routeDistinguisher.AsIndex})
    routeDistinguisher.EntityData.Leafs.Append("address", types.YLeaf{"Address", routeDistinguisher.Address})
    routeDistinguisher.EntityData.Leafs.Append("address-index", types.YLeaf{"AddressIndex", routeDistinguisher.AddressIndex})

    routeDistinguisher.EntityData.YListKeys = []string {}

    return &(routeDistinguisher.EntityData)
}

// Vrfs_Vrf_MulticastHost
// Multicast host stack configuration
type Vrfs_Vrf_MulticastHost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 configuration.
    Ipv4 Vrfs_Vrf_MulticastHost_Ipv4

    // IPv6 configuration.
    Ipv6 Vrfs_Vrf_MulticastHost_Ipv6
}

func (multicastHost *Vrfs_Vrf_MulticastHost) GetEntityData() *types.CommonEntityData {
    multicastHost.EntityData.YFilter = multicastHost.YFilter
    multicastHost.EntityData.YangName = "multicast-host"
    multicastHost.EntityData.BundleName = "cisco_ios_xr"
    multicastHost.EntityData.ParentYangName = "vrf"
    multicastHost.EntityData.SegmentPath = "Cisco-IOS-XR-ip-iarm-vrf-cfg:multicast-host"
    multicastHost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastHost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastHost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastHost.EntityData.Children = types.NewOrderedMap()
    multicastHost.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &multicastHost.Ipv4})
    multicastHost.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &multicastHost.Ipv6})
    multicastHost.EntityData.Leafs = types.NewOrderedMap()

    multicastHost.EntityData.YListKeys = []string {}

    return &(multicastHost.EntityData)
}

// Vrfs_Vrf_MulticastHost_Ipv4
// IPv4 configuration
type Vrfs_Vrf_MulticastHost_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default multicast host interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}
}

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "multicast-host"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", ipv4.Interface})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Vrfs_Vrf_MulticastHost_Ipv6
// IPv6 configuration
type Vrfs_Vrf_MulticastHost_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default multicast host interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}
}

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "multicast-host"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", ipv6.Interface})

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// GlobalAf
// global af
type GlobalAf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF address family configuration.
    Afs GlobalAf_Afs
}

func (globalAf *GlobalAf) GetEntityData() *types.CommonEntityData {
    globalAf.EntityData.YFilter = globalAf.YFilter
    globalAf.EntityData.YangName = "global-af"
    globalAf.EntityData.BundleName = "cisco_ios_xr"
    globalAf.EntityData.ParentYangName = "Cisco-IOS-XR-infra-rsi-cfg"
    globalAf.EntityData.SegmentPath = "Cisco-IOS-XR-infra-rsi-cfg:global-af"
    globalAf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalAf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalAf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalAf.EntityData.Children = types.NewOrderedMap()
    globalAf.EntityData.Children.Append("afs", types.YChild{"Afs", &globalAf.Afs})
    globalAf.EntityData.Leafs = types.NewOrderedMap()

    globalAf.EntityData.YListKeys = []string {}

    return &(globalAf.EntityData)
}

// GlobalAf_Afs
// VRF address family configuration
type GlobalAf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF address family configuration. The type is slice of GlobalAf_Afs_Af.
    Af []*GlobalAf_Afs_Af
}

func (afs *GlobalAf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "global-af"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// GlobalAf_Afs_Af
// VRF address family configuration
type GlobalAf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address family. The type is VrfAddressFamily.
    AfName interface{}

    // This attribute is a key. Sub-Address family. The type is
    // VrfSubAddressFamily.
    SafName interface{}

    // This attribute is a key. Topology name. The type is string with length:
    // 1..244.
    TopologyName interface{}

    // VRF configuration for a particular address family. The type is interface{}.
    Create interface{}

    // BGP AF VRF config.
    Bgp GlobalAf_Afs_Af_Bgp

    // Set maximum prefix limits.
    MaximumPrefix GlobalAf_Afs_Af_MaximumPrefix
}

func (af *GlobalAf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name") + types.AddKeyToken(af.SafName, "saf-name") + types.AddKeyToken(af.TopologyName, "topology-name")
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("Cisco-IOS-XR-ipv4-bgp-cfg:bgp", types.YChild{"Bgp", &af.Bgp})
    af.EntityData.Children.Append("Cisco-IOS-XR-ip-rib-cfg:maximum-prefix", types.YChild{"MaximumPrefix", &af.MaximumPrefix})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", af.SafName})
    af.EntityData.Leafs.Append("topology-name", types.YLeaf{"TopologyName", af.TopologyName})
    af.EntityData.Leafs.Append("create", types.YLeaf{"Create", af.Create})

    af.EntityData.YListKeys = []string {"AfName", "SafName", "TopologyName"}

    return &(af.EntityData)
}

// GlobalAf_Afs_Af_Bgp
// BGP AF VRF config
type GlobalAf_Afs_Af_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route policy for export filtering. The type is string.
    ExportRoutePolicy interface{}

    // Route policy for import filtering. The type is string.
    ImportRoutePolicy interface{}

    // TRUE Enable advertising imported paths to PEsFALSE Disable advertising
    // imported paths to PEs. The type is bool.
    ImportVrfOptions interface{}

    // Import Route targets.
    ImportRouteTargets GlobalAf_Afs_Af_Bgp_ImportRouteTargets

    // Export Route targets.
    ExportRouteTargets GlobalAf_Afs_Af_Bgp_ExportRouteTargets

    // Route policy for vrf to global export filtering.
    VrfToGlobalExportRoutePolicy GlobalAf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy

    // Export VRF options.
    ExportVrfOptions GlobalAf_Afs_Af_Bgp_ExportVrfOptions

    // Route policy for global to vrf import filtering.
    GlobalToVrfImportRoutePolicy GlobalAf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy
}

func (bgp *GlobalAf_Afs_Af_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "af"
    bgp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-bgp-cfg:bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Children.Append("import-route-targets", types.YChild{"ImportRouteTargets", &bgp.ImportRouteTargets})
    bgp.EntityData.Children.Append("export-route-targets", types.YChild{"ExportRouteTargets", &bgp.ExportRouteTargets})
    bgp.EntityData.Children.Append("vrf-to-global-export-route-policy", types.YChild{"VrfToGlobalExportRoutePolicy", &bgp.VrfToGlobalExportRoutePolicy})
    bgp.EntityData.Children.Append("export-vrf-options", types.YChild{"ExportVrfOptions", &bgp.ExportVrfOptions})
    bgp.EntityData.Children.Append("global-to-vrf-import-route-policy", types.YChild{"GlobalToVrfImportRoutePolicy", &bgp.GlobalToVrfImportRoutePolicy})
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("export-route-policy", types.YLeaf{"ExportRoutePolicy", bgp.ExportRoutePolicy})
    bgp.EntityData.Leafs.Append("import-route-policy", types.YLeaf{"ImportRoutePolicy", bgp.ImportRoutePolicy})
    bgp.EntityData.Leafs.Append("import-vrf-options", types.YLeaf{"ImportVrfOptions", bgp.ImportVrfOptions})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// GlobalAf_Afs_Af_Bgp_ImportRouteTargets
// Import Route targets
type GlobalAf_Afs_Af_Bgp_ImportRouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route target table.
    RouteTargets GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets
}

func (importRouteTargets *GlobalAf_Afs_Af_Bgp_ImportRouteTargets) GetEntityData() *types.CommonEntityData {
    importRouteTargets.EntityData.YFilter = importRouteTargets.YFilter
    importRouteTargets.EntityData.YangName = "import-route-targets"
    importRouteTargets.EntityData.BundleName = "cisco_ios_xr"
    importRouteTargets.EntityData.ParentYangName = "bgp"
    importRouteTargets.EntityData.SegmentPath = "import-route-targets"
    importRouteTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    importRouteTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    importRouteTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    importRouteTargets.EntityData.Children = types.NewOrderedMap()
    importRouteTargets.EntityData.Children.Append("route-targets", types.YChild{"RouteTargets", &importRouteTargets.RouteTargets})
    importRouteTargets.EntityData.Leafs = types.NewOrderedMap()

    importRouteTargets.EntityData.YListKeys = []string {}

    return &(importRouteTargets.EntityData)
}

// GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets
// Route target table
type GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route target. The type is slice of
    // GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget.
    RouteTarget []*GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget
}

func (routeTargets *GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetEntityData() *types.CommonEntityData {
    routeTargets.EntityData.YFilter = routeTargets.YFilter
    routeTargets.EntityData.YangName = "route-targets"
    routeTargets.EntityData.BundleName = "cisco_ios_xr"
    routeTargets.EntityData.ParentYangName = "import-route-targets"
    routeTargets.EntityData.SegmentPath = "route-targets"
    routeTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTargets.EntityData.Children = types.NewOrderedMap()
    routeTargets.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", nil})
    for i := range routeTargets.RouteTarget {
        routeTargets.EntityData.Children.Append(types.GetSegmentPath(routeTargets.RouteTarget[i]), types.YChild{"RouteTarget", routeTargets.RouteTarget[i]})
    }
    routeTargets.EntityData.Leafs = types.NewOrderedMap()

    routeTargets.EntityData.YListKeys = []string {}

    return &(routeTargets.EntityData)
}

// GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget
// Route target
type GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Type of RT. The type is BgpVrfRouteTarget.
    Type interface{}

    // as or four byte as. The type is slice of
    // GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs.
    AsOrFourByteAs []*GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs

    // ipv4 address. The type is slice of
    // GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address.
    Ipv4Address []*GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address
}

func (routeTarget *GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "route-targets"
    routeTarget.EntityData.SegmentPath = "route-target" + types.AddKeyToken(routeTarget.Type, "type")
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Children.Append("as-or-four-byte-as", types.YChild{"AsOrFourByteAs", nil})
    for i := range routeTarget.AsOrFourByteAs {
        routeTarget.EntityData.Children.Append(types.GetSegmentPath(routeTarget.AsOrFourByteAs[i]), types.YChild{"AsOrFourByteAs", routeTarget.AsOrFourByteAs[i]})
    }
    routeTarget.EntityData.Children.Append("ipv4-address", types.YChild{"Ipv4Address", nil})
    for i := range routeTarget.Ipv4Address {
        routeTarget.EntityData.Children.Append(types.GetSegmentPath(routeTarget.Ipv4Address[i]), types.YChild{"Ipv4Address", routeTarget.Ipv4Address[i]})
    }
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("type", types.YLeaf{"Type", routeTarget.Type})

    routeTarget.EntityData.YListKeys = []string {"Type"}

    return &(routeTarget.EntityData)
}

// GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs
// as or four byte as
type GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. AS number. The type is interface{} with range:
    // 0..4294967295.
    AsXx interface{}

    // This attribute is a key. AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // This attribute is a key. AS number Index. The type is interface{} with
    // range: 0..4294967295.
    AsIndex interface{}

    // This attribute is a key. Stitching RT. The type is interface{} with range:
    // 0..1.
    StitchingRt interface{}
}

func (asOrFourByteAs *GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetEntityData() *types.CommonEntityData {
    asOrFourByteAs.EntityData.YFilter = asOrFourByteAs.YFilter
    asOrFourByteAs.EntityData.YangName = "as-or-four-byte-as"
    asOrFourByteAs.EntityData.BundleName = "cisco_ios_xr"
    asOrFourByteAs.EntityData.ParentYangName = "route-target"
    asOrFourByteAs.EntityData.SegmentPath = "as-or-four-byte-as" + types.AddKeyToken(asOrFourByteAs.AsXx, "as-xx") + types.AddKeyToken(asOrFourByteAs.As, "as") + types.AddKeyToken(asOrFourByteAs.AsIndex, "as-index") + types.AddKeyToken(asOrFourByteAs.StitchingRt, "stitching-rt")
    asOrFourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asOrFourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asOrFourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asOrFourByteAs.EntityData.Children = types.NewOrderedMap()
    asOrFourByteAs.EntityData.Leafs = types.NewOrderedMap()
    asOrFourByteAs.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", asOrFourByteAs.AsXx})
    asOrFourByteAs.EntityData.Leafs.Append("as", types.YLeaf{"As", asOrFourByteAs.As})
    asOrFourByteAs.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", asOrFourByteAs.AsIndex})
    asOrFourByteAs.EntityData.Leafs.Append("stitching-rt", types.YLeaf{"StitchingRt", asOrFourByteAs.StitchingRt})

    asOrFourByteAs.EntityData.YListKeys = []string {"AsXx", "As", "AsIndex", "StitchingRt"}

    return &(asOrFourByteAs.EntityData)
}

// GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address
// ipv4 address
type GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IP address Index. The type is interface{} with
    // range: 0..65535.
    AddressIndex interface{}

    // This attribute is a key. Stitching RT. The type is interface{} with range:
    // 0..1.
    StitchingRt interface{}
}

func (ipv4Address *GlobalAf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetEntityData() *types.CommonEntityData {
    ipv4Address.EntityData.YFilter = ipv4Address.YFilter
    ipv4Address.EntityData.YangName = "ipv4-address"
    ipv4Address.EntityData.BundleName = "cisco_ios_xr"
    ipv4Address.EntityData.ParentYangName = "route-target"
    ipv4Address.EntityData.SegmentPath = "ipv4-address" + types.AddKeyToken(ipv4Address.Address, "address") + types.AddKeyToken(ipv4Address.AddressIndex, "address-index") + types.AddKeyToken(ipv4Address.StitchingRt, "stitching-rt")
    ipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Address.EntityData.Children = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4Address.Address})
    ipv4Address.EntityData.Leafs.Append("address-index", types.YLeaf{"AddressIndex", ipv4Address.AddressIndex})
    ipv4Address.EntityData.Leafs.Append("stitching-rt", types.YLeaf{"StitchingRt", ipv4Address.StitchingRt})

    ipv4Address.EntityData.YListKeys = []string {"Address", "AddressIndex", "StitchingRt"}

    return &(ipv4Address.EntityData)
}

// GlobalAf_Afs_Af_Bgp_ExportRouteTargets
// Export Route targets
type GlobalAf_Afs_Af_Bgp_ExportRouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route target table.
    RouteTargets GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets
}

func (exportRouteTargets *GlobalAf_Afs_Af_Bgp_ExportRouteTargets) GetEntityData() *types.CommonEntityData {
    exportRouteTargets.EntityData.YFilter = exportRouteTargets.YFilter
    exportRouteTargets.EntityData.YangName = "export-route-targets"
    exportRouteTargets.EntityData.BundleName = "cisco_ios_xr"
    exportRouteTargets.EntityData.ParentYangName = "bgp"
    exportRouteTargets.EntityData.SegmentPath = "export-route-targets"
    exportRouteTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exportRouteTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exportRouteTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exportRouteTargets.EntityData.Children = types.NewOrderedMap()
    exportRouteTargets.EntityData.Children.Append("route-targets", types.YChild{"RouteTargets", &exportRouteTargets.RouteTargets})
    exportRouteTargets.EntityData.Leafs = types.NewOrderedMap()

    exportRouteTargets.EntityData.YListKeys = []string {}

    return &(exportRouteTargets.EntityData)
}

// GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets
// Route target table
type GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route target. The type is slice of
    // GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget.
    RouteTarget []*GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget
}

func (routeTargets *GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetEntityData() *types.CommonEntityData {
    routeTargets.EntityData.YFilter = routeTargets.YFilter
    routeTargets.EntityData.YangName = "route-targets"
    routeTargets.EntityData.BundleName = "cisco_ios_xr"
    routeTargets.EntityData.ParentYangName = "export-route-targets"
    routeTargets.EntityData.SegmentPath = "route-targets"
    routeTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTargets.EntityData.Children = types.NewOrderedMap()
    routeTargets.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", nil})
    for i := range routeTargets.RouteTarget {
        routeTargets.EntityData.Children.Append(types.GetSegmentPath(routeTargets.RouteTarget[i]), types.YChild{"RouteTarget", routeTargets.RouteTarget[i]})
    }
    routeTargets.EntityData.Leafs = types.NewOrderedMap()

    routeTargets.EntityData.YListKeys = []string {}

    return &(routeTargets.EntityData)
}

// GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget
// Route target
type GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Type of RT. The type is BgpVrfRouteTarget.
    Type interface{}

    // as or four byte as. The type is slice of
    // GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs.
    AsOrFourByteAs []*GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs

    // ipv4 address. The type is slice of
    // GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address.
    Ipv4Address []*GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address
}

func (routeTarget *GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "route-targets"
    routeTarget.EntityData.SegmentPath = "route-target" + types.AddKeyToken(routeTarget.Type, "type")
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Children.Append("as-or-four-byte-as", types.YChild{"AsOrFourByteAs", nil})
    for i := range routeTarget.AsOrFourByteAs {
        routeTarget.EntityData.Children.Append(types.GetSegmentPath(routeTarget.AsOrFourByteAs[i]), types.YChild{"AsOrFourByteAs", routeTarget.AsOrFourByteAs[i]})
    }
    routeTarget.EntityData.Children.Append("ipv4-address", types.YChild{"Ipv4Address", nil})
    for i := range routeTarget.Ipv4Address {
        routeTarget.EntityData.Children.Append(types.GetSegmentPath(routeTarget.Ipv4Address[i]), types.YChild{"Ipv4Address", routeTarget.Ipv4Address[i]})
    }
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("type", types.YLeaf{"Type", routeTarget.Type})

    routeTarget.EntityData.YListKeys = []string {"Type"}

    return &(routeTarget.EntityData)
}

// GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs
// as or four byte as
type GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. AS number. The type is interface{} with range:
    // 0..4294967295.
    AsXx interface{}

    // This attribute is a key. AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // This attribute is a key. AS number Index. The type is interface{} with
    // range: 0..4294967295.
    AsIndex interface{}

    // This attribute is a key. Stitching RT. The type is interface{} with range:
    // 0..1.
    StitchingRt interface{}
}

func (asOrFourByteAs *GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetEntityData() *types.CommonEntityData {
    asOrFourByteAs.EntityData.YFilter = asOrFourByteAs.YFilter
    asOrFourByteAs.EntityData.YangName = "as-or-four-byte-as"
    asOrFourByteAs.EntityData.BundleName = "cisco_ios_xr"
    asOrFourByteAs.EntityData.ParentYangName = "route-target"
    asOrFourByteAs.EntityData.SegmentPath = "as-or-four-byte-as" + types.AddKeyToken(asOrFourByteAs.AsXx, "as-xx") + types.AddKeyToken(asOrFourByteAs.As, "as") + types.AddKeyToken(asOrFourByteAs.AsIndex, "as-index") + types.AddKeyToken(asOrFourByteAs.StitchingRt, "stitching-rt")
    asOrFourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asOrFourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asOrFourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asOrFourByteAs.EntityData.Children = types.NewOrderedMap()
    asOrFourByteAs.EntityData.Leafs = types.NewOrderedMap()
    asOrFourByteAs.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", asOrFourByteAs.AsXx})
    asOrFourByteAs.EntityData.Leafs.Append("as", types.YLeaf{"As", asOrFourByteAs.As})
    asOrFourByteAs.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", asOrFourByteAs.AsIndex})
    asOrFourByteAs.EntityData.Leafs.Append("stitching-rt", types.YLeaf{"StitchingRt", asOrFourByteAs.StitchingRt})

    asOrFourByteAs.EntityData.YListKeys = []string {"AsXx", "As", "AsIndex", "StitchingRt"}

    return &(asOrFourByteAs.EntityData)
}

// GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address
// ipv4 address
type GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IP address Index. The type is interface{} with
    // range: 0..65535.
    AddressIndex interface{}

    // This attribute is a key. Stitching RT. The type is interface{} with range:
    // 0..1.
    StitchingRt interface{}
}

func (ipv4Address *GlobalAf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetEntityData() *types.CommonEntityData {
    ipv4Address.EntityData.YFilter = ipv4Address.YFilter
    ipv4Address.EntityData.YangName = "ipv4-address"
    ipv4Address.EntityData.BundleName = "cisco_ios_xr"
    ipv4Address.EntityData.ParentYangName = "route-target"
    ipv4Address.EntityData.SegmentPath = "ipv4-address" + types.AddKeyToken(ipv4Address.Address, "address") + types.AddKeyToken(ipv4Address.AddressIndex, "address-index") + types.AddKeyToken(ipv4Address.StitchingRt, "stitching-rt")
    ipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Address.EntityData.Children = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4Address.Address})
    ipv4Address.EntityData.Leafs.Append("address-index", types.YLeaf{"AddressIndex", ipv4Address.AddressIndex})
    ipv4Address.EntityData.Leafs.Append("stitching-rt", types.YLeaf{"StitchingRt", ipv4Address.StitchingRt})

    ipv4Address.EntityData.YListKeys = []string {"Address", "AddressIndex", "StitchingRt"}

    return &(ipv4Address.EntityData)
}

// GlobalAf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy
// Route policy for vrf to global export filtering
// This type is a presence type.
type GlobalAf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Vrf to global export route policy. The type is string. This attribute is
    // mandatory.
    RoutePolicyName interface{}

    // TRUE Enable imported VPN paths to be exported to Default VRF.FALSE Disable
    // imported VPN paths to be exported to Default VRF. The type is bool.
    AllowImportedVpn interface{}
}

func (vrfToGlobalExportRoutePolicy *GlobalAf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetEntityData() *types.CommonEntityData {
    vrfToGlobalExportRoutePolicy.EntityData.YFilter = vrfToGlobalExportRoutePolicy.YFilter
    vrfToGlobalExportRoutePolicy.EntityData.YangName = "vrf-to-global-export-route-policy"
    vrfToGlobalExportRoutePolicy.EntityData.BundleName = "cisco_ios_xr"
    vrfToGlobalExportRoutePolicy.EntityData.ParentYangName = "bgp"
    vrfToGlobalExportRoutePolicy.EntityData.SegmentPath = "vrf-to-global-export-route-policy"
    vrfToGlobalExportRoutePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfToGlobalExportRoutePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfToGlobalExportRoutePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfToGlobalExportRoutePolicy.EntityData.Children = types.NewOrderedMap()
    vrfToGlobalExportRoutePolicy.EntityData.Leafs = types.NewOrderedMap()
    vrfToGlobalExportRoutePolicy.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", vrfToGlobalExportRoutePolicy.RoutePolicyName})
    vrfToGlobalExportRoutePolicy.EntityData.Leafs.Append("allow-imported-vpn", types.YLeaf{"AllowImportedVpn", vrfToGlobalExportRoutePolicy.AllowImportedVpn})

    vrfToGlobalExportRoutePolicy.EntityData.YListKeys = []string {}

    return &(vrfToGlobalExportRoutePolicy.EntityData)
}

// GlobalAf_Afs_Af_Bgp_ExportVrfOptions
// Export VRF options
type GlobalAf_Afs_Af_Bgp_ExportVrfOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE Enable imported VPN paths to be exported to non-default VRFFALSE
    // Disable imported VPN paths to be exported to non-default VRF. The type is
    // bool.
    AllowImportedVpn interface{}

    // TRUE Use stitchng RTs to import extranet pathsFALSE Use regular RTs to
    // import extranet paths. The type is bool.
    ImportStitchingRt interface{}
}

func (exportVrfOptions *GlobalAf_Afs_Af_Bgp_ExportVrfOptions) GetEntityData() *types.CommonEntityData {
    exportVrfOptions.EntityData.YFilter = exportVrfOptions.YFilter
    exportVrfOptions.EntityData.YangName = "export-vrf-options"
    exportVrfOptions.EntityData.BundleName = "cisco_ios_xr"
    exportVrfOptions.EntityData.ParentYangName = "bgp"
    exportVrfOptions.EntityData.SegmentPath = "export-vrf-options"
    exportVrfOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exportVrfOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exportVrfOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exportVrfOptions.EntityData.Children = types.NewOrderedMap()
    exportVrfOptions.EntityData.Leafs = types.NewOrderedMap()
    exportVrfOptions.EntityData.Leafs.Append("allow-imported-vpn", types.YLeaf{"AllowImportedVpn", exportVrfOptions.AllowImportedVpn})
    exportVrfOptions.EntityData.Leafs.Append("import-stitching-rt", types.YLeaf{"ImportStitchingRt", exportVrfOptions.ImportStitchingRt})

    exportVrfOptions.EntityData.YListKeys = []string {}

    return &(exportVrfOptions.EntityData)
}

// GlobalAf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy
// Route policy for global to vrf import filtering
// This type is a presence type.
type GlobalAf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Global to vrf import route policy. The type is string. This attribute is
    // mandatory.
    RoutePolicyName interface{}

    // TRUE Enable advertising imported paths to PEsFALSE Disable advertising
    // imported paths to PEs. The type is bool.
    AdvertiseAsVpn interface{}
}

func (globalToVrfImportRoutePolicy *GlobalAf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetEntityData() *types.CommonEntityData {
    globalToVrfImportRoutePolicy.EntityData.YFilter = globalToVrfImportRoutePolicy.YFilter
    globalToVrfImportRoutePolicy.EntityData.YangName = "global-to-vrf-import-route-policy"
    globalToVrfImportRoutePolicy.EntityData.BundleName = "cisco_ios_xr"
    globalToVrfImportRoutePolicy.EntityData.ParentYangName = "bgp"
    globalToVrfImportRoutePolicy.EntityData.SegmentPath = "global-to-vrf-import-route-policy"
    globalToVrfImportRoutePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalToVrfImportRoutePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalToVrfImportRoutePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalToVrfImportRoutePolicy.EntityData.Children = types.NewOrderedMap()
    globalToVrfImportRoutePolicy.EntityData.Leafs = types.NewOrderedMap()
    globalToVrfImportRoutePolicy.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", globalToVrfImportRoutePolicy.RoutePolicyName})
    globalToVrfImportRoutePolicy.EntityData.Leafs.Append("advertise-as-vpn", types.YLeaf{"AdvertiseAsVpn", globalToVrfImportRoutePolicy.AdvertiseAsVpn})

    globalToVrfImportRoutePolicy.EntityData.YListKeys = []string {}

    return &(globalToVrfImportRoutePolicy.EntityData)
}

// GlobalAf_Afs_Af_MaximumPrefix
// Set maximum prefix limits
// This type is a presence type.
type GlobalAf_Afs_Af_MaximumPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Set table's maximum prefix limit. The type is interface{} with range:
    // 32..10000000. This attribute is mandatory.
    PrefixLimit interface{}

    // Mid-threshold (% of maximum). The type is interface{} with range: 1..100.
    MidThreshold interface{}
}

func (maximumPrefix *GlobalAf_Afs_Af_MaximumPrefix) GetEntityData() *types.CommonEntityData {
    maximumPrefix.EntityData.YFilter = maximumPrefix.YFilter
    maximumPrefix.EntityData.YangName = "maximum-prefix"
    maximumPrefix.EntityData.BundleName = "cisco_ios_xr"
    maximumPrefix.EntityData.ParentYangName = "af"
    maximumPrefix.EntityData.SegmentPath = "Cisco-IOS-XR-ip-rib-cfg:maximum-prefix"
    maximumPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumPrefix.EntityData.Children = types.NewOrderedMap()
    maximumPrefix.EntityData.Leafs = types.NewOrderedMap()
    maximumPrefix.EntityData.Leafs.Append("prefix-limit", types.YLeaf{"PrefixLimit", maximumPrefix.PrefixLimit})
    maximumPrefix.EntityData.Leafs.Append("mid-threshold", types.YLeaf{"MidThreshold", maximumPrefix.MidThreshold})

    maximumPrefix.EntityData.YListKeys = []string {}

    return &(maximumPrefix.EntityData)
}

// Srlg
// srlg
type Srlg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable SRLG. The type is interface{}.
    Enable interface{}

    // Set of interfaces configured with SRLG.
    Interfaces Srlg_Interfaces

    // Set of SRLG name configuration.
    SrlgNames Srlg_SrlgNames

    // Set of groups configured with SRLG.
    Groups Srlg_Groups

    // Set of inherit nodes configured with SRLG.
    InheritNodes Srlg_InheritNodes
}

func (srlg *Srlg) GetEntityData() *types.CommonEntityData {
    srlg.EntityData.YFilter = srlg.YFilter
    srlg.EntityData.YangName = "srlg"
    srlg.EntityData.BundleName = "cisco_ios_xr"
    srlg.EntityData.ParentYangName = "Cisco-IOS-XR-infra-rsi-cfg"
    srlg.EntityData.SegmentPath = "Cisco-IOS-XR-infra-rsi-cfg:srlg"
    srlg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlg.EntityData.Children = types.NewOrderedMap()
    srlg.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &srlg.Interfaces})
    srlg.EntityData.Children.Append("srlg-names", types.YChild{"SrlgNames", &srlg.SrlgNames})
    srlg.EntityData.Children.Append("groups", types.YChild{"Groups", &srlg.Groups})
    srlg.EntityData.Children.Append("inherit-nodes", types.YChild{"InheritNodes", &srlg.InheritNodes})
    srlg.EntityData.Leafs = types.NewOrderedMap()
    srlg.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", srlg.Enable})

    srlg.EntityData.YListKeys = []string {}

    return &(srlg.EntityData)
}

// Srlg_Interfaces
// Set of interfaces configured with SRLG
type Srlg_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface configurations. The type is slice of Srlg_Interfaces_Interface.
    Interface []*Srlg_Interfaces_Interface
}

func (interfaces *Srlg_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "srlg"
    interfaces.EntityData.SegmentPath = "interfaces"
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

// Srlg_Interfaces_Interface
// Interface configurations
type Srlg_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Enable SRLG interface. The type is interface{}.
    Enable interface{}

    // Include optical configuration for an interface.
    IncludeOptical Srlg_Interfaces_Interface_IncludeOptical

    // Group configuration for an interface.
    InterfaceGroup Srlg_Interfaces_Interface_InterfaceGroup

    // SRLG Value configuration for an interface.
    Values Srlg_Interfaces_Interface_Values

    // SRLG Name configuration for an interface.
    InterfaceSrlgNames Srlg_Interfaces_Interface_InterfaceSrlgNames
}

func (self *Srlg_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("include-optical", types.YChild{"IncludeOptical", &self.IncludeOptical})
    self.EntityData.Children.Append("interface-group", types.YChild{"InterfaceGroup", &self.InterfaceGroup})
    self.EntityData.Children.Append("values", types.YChild{"Values", &self.Values})
    self.EntityData.Children.Append("interface-srlg-names", types.YChild{"InterfaceSrlgNames", &self.InterfaceSrlgNames})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Srlg_Interfaces_Interface_IncludeOptical
// Include optical configuration for an interface
type Srlg_Interfaces_Interface_IncludeOptical struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable SRLG interface include optical. The type is interface{}.
    Enable interface{}

    // Priority for optical domain values. The type is SrlgPriority. The default
    // value is default.
    Priority interface{}
}

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetEntityData() *types.CommonEntityData {
    includeOptical.EntityData.YFilter = includeOptical.YFilter
    includeOptical.EntityData.YangName = "include-optical"
    includeOptical.EntityData.BundleName = "cisco_ios_xr"
    includeOptical.EntityData.ParentYangName = "interface"
    includeOptical.EntityData.SegmentPath = "include-optical"
    includeOptical.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    includeOptical.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    includeOptical.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    includeOptical.EntityData.Children = types.NewOrderedMap()
    includeOptical.EntityData.Leafs = types.NewOrderedMap()
    includeOptical.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", includeOptical.Enable})
    includeOptical.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", includeOptical.Priority})

    includeOptical.EntityData.YListKeys = []string {}

    return &(includeOptical.EntityData)
}

// Srlg_Interfaces_Interface_InterfaceGroup
// Group configuration for an interface
type Srlg_Interfaces_Interface_InterfaceGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable SRLG interface group submode. The type is interface{}.
    Enable interface{}

    // Set of group name under an interface.
    GroupNames Srlg_Interfaces_Interface_InterfaceGroup_GroupNames
}

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetEntityData() *types.CommonEntityData {
    interfaceGroup.EntityData.YFilter = interfaceGroup.YFilter
    interfaceGroup.EntityData.YangName = "interface-group"
    interfaceGroup.EntityData.BundleName = "cisco_ios_xr"
    interfaceGroup.EntityData.ParentYangName = "interface"
    interfaceGroup.EntityData.SegmentPath = "interface-group"
    interfaceGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceGroup.EntityData.Children = types.NewOrderedMap()
    interfaceGroup.EntityData.Children.Append("group-names", types.YChild{"GroupNames", &interfaceGroup.GroupNames})
    interfaceGroup.EntityData.Leafs = types.NewOrderedMap()
    interfaceGroup.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", interfaceGroup.Enable})

    interfaceGroup.EntityData.YListKeys = []string {}

    return &(interfaceGroup.EntityData)
}

// Srlg_Interfaces_Interface_InterfaceGroup_GroupNames
// Set of group name under an interface
type Srlg_Interfaces_Interface_InterfaceGroup_GroupNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group name included under interface. The type is slice of
    // Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName.
    GroupName []*Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName
}

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetEntityData() *types.CommonEntityData {
    groupNames.EntityData.YFilter = groupNames.YFilter
    groupNames.EntityData.YangName = "group-names"
    groupNames.EntityData.BundleName = "cisco_ios_xr"
    groupNames.EntityData.ParentYangName = "interface-group"
    groupNames.EntityData.SegmentPath = "group-names"
    groupNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupNames.EntityData.Children = types.NewOrderedMap()
    groupNames.EntityData.Children.Append("group-name", types.YChild{"GroupName", nil})
    for i := range groupNames.GroupName {
        groupNames.EntityData.Children.Append(types.GetSegmentPath(groupNames.GroupName[i]), types.YChild{"GroupName", groupNames.GroupName[i]})
    }
    groupNames.EntityData.Leafs = types.NewOrderedMap()

    groupNames.EntityData.YListKeys = []string {}

    return &(groupNames.EntityData)
}

// Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName
// Group name included under interface
type Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group name index. The type is interface{} with
    // range: 0..4294967295.
    GroupNameIndex interface{}

    // Group name. The type is string. This attribute is mandatory.
    GroupName interface{}

    // SRLG priority. The type is SrlgPriority. The default value is default.
    SrlgPriority interface{}
}

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetEntityData() *types.CommonEntityData {
    groupName.EntityData.YFilter = groupName.YFilter
    groupName.EntityData.YangName = "group-name"
    groupName.EntityData.BundleName = "cisco_ios_xr"
    groupName.EntityData.ParentYangName = "group-names"
    groupName.EntityData.SegmentPath = "group-name" + types.AddKeyToken(groupName.GroupNameIndex, "group-name-index")
    groupName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupName.EntityData.Children = types.NewOrderedMap()
    groupName.EntityData.Leafs = types.NewOrderedMap()
    groupName.EntityData.Leafs.Append("group-name-index", types.YLeaf{"GroupNameIndex", groupName.GroupNameIndex})
    groupName.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", groupName.GroupName})
    groupName.EntityData.Leafs.Append("srlg-priority", types.YLeaf{"SrlgPriority", groupName.SrlgPriority})

    groupName.EntityData.YListKeys = []string {"GroupNameIndex"}

    return &(groupName.EntityData)
}

// Srlg_Interfaces_Interface_Values
// SRLG Value configuration for an interface
type Srlg_Interfaces_Interface_Values struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG value data. The type is slice of
    // Srlg_Interfaces_Interface_Values_Value.
    Value []*Srlg_Interfaces_Interface_Values_Value
}

func (values *Srlg_Interfaces_Interface_Values) GetEntityData() *types.CommonEntityData {
    values.EntityData.YFilter = values.YFilter
    values.EntityData.YangName = "values"
    values.EntityData.BundleName = "cisco_ios_xr"
    values.EntityData.ParentYangName = "interface"
    values.EntityData.SegmentPath = "values"
    values.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    values.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    values.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    values.EntityData.Children = types.NewOrderedMap()
    values.EntityData.Children.Append("value", types.YChild{"Value", nil})
    for i := range values.Value {
        values.EntityData.Children.Append(types.GetSegmentPath(values.Value[i]), types.YChild{"Value", values.Value[i]})
    }
    values.EntityData.Leafs = types.NewOrderedMap()

    values.EntityData.YListKeys = []string {}

    return &(values.EntityData)
}

// Srlg_Interfaces_Interface_Values_Value
// SRLG value data
type Srlg_Interfaces_Interface_Values_Value struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG index. The type is interface{} with range:
    // 1..65535.
    SrlgIndex interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    SrlgValue interface{}

    // SRLG priority. The type is SrlgPriority. The default value is default.
    SrlgPriority interface{}
}

func (value *Srlg_Interfaces_Interface_Values_Value) GetEntityData() *types.CommonEntityData {
    value.EntityData.YFilter = value.YFilter
    value.EntityData.YangName = "value"
    value.EntityData.BundleName = "cisco_ios_xr"
    value.EntityData.ParentYangName = "values"
    value.EntityData.SegmentPath = "value" + types.AddKeyToken(value.SrlgIndex, "srlg-index")
    value.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    value.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    value.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    value.EntityData.Children = types.NewOrderedMap()
    value.EntityData.Leafs = types.NewOrderedMap()
    value.EntityData.Leafs.Append("srlg-index", types.YLeaf{"SrlgIndex", value.SrlgIndex})
    value.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", value.SrlgValue})
    value.EntityData.Leafs.Append("srlg-priority", types.YLeaf{"SrlgPriority", value.SrlgPriority})

    value.EntityData.YListKeys = []string {"SrlgIndex"}

    return &(value.EntityData)
}

// Srlg_Interfaces_Interface_InterfaceSrlgNames
// SRLG Name configuration for an interface
type Srlg_Interfaces_Interface_InterfaceSrlgNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG name data. The type is slice of
    // Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName.
    InterfaceSrlgName []*Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName
}

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetEntityData() *types.CommonEntityData {
    interfaceSrlgNames.EntityData.YFilter = interfaceSrlgNames.YFilter
    interfaceSrlgNames.EntityData.YangName = "interface-srlg-names"
    interfaceSrlgNames.EntityData.BundleName = "cisco_ios_xr"
    interfaceSrlgNames.EntityData.ParentYangName = "interface"
    interfaceSrlgNames.EntityData.SegmentPath = "interface-srlg-names"
    interfaceSrlgNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSrlgNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSrlgNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSrlgNames.EntityData.Children = types.NewOrderedMap()
    interfaceSrlgNames.EntityData.Children.Append("interface-srlg-name", types.YChild{"InterfaceSrlgName", nil})
    for i := range interfaceSrlgNames.InterfaceSrlgName {
        interfaceSrlgNames.EntityData.Children.Append(types.GetSegmentPath(interfaceSrlgNames.InterfaceSrlgName[i]), types.YChild{"InterfaceSrlgName", interfaceSrlgNames.InterfaceSrlgName[i]})
    }
    interfaceSrlgNames.EntityData.Leafs = types.NewOrderedMap()

    interfaceSrlgNames.EntityData.YListKeys = []string {}

    return &(interfaceSrlgNames.EntityData)
}

// Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName
// SRLG name data
type Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG name. The type is string with length: 1..64.
    SrlgName interface{}
}

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetEntityData() *types.CommonEntityData {
    interfaceSrlgName.EntityData.YFilter = interfaceSrlgName.YFilter
    interfaceSrlgName.EntityData.YangName = "interface-srlg-name"
    interfaceSrlgName.EntityData.BundleName = "cisco_ios_xr"
    interfaceSrlgName.EntityData.ParentYangName = "interface-srlg-names"
    interfaceSrlgName.EntityData.SegmentPath = "interface-srlg-name" + types.AddKeyToken(interfaceSrlgName.SrlgName, "srlg-name")
    interfaceSrlgName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSrlgName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSrlgName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSrlgName.EntityData.Children = types.NewOrderedMap()
    interfaceSrlgName.EntityData.Leafs = types.NewOrderedMap()
    interfaceSrlgName.EntityData.Leafs.Append("srlg-name", types.YLeaf{"SrlgName", interfaceSrlgName.SrlgName})

    interfaceSrlgName.EntityData.YListKeys = []string {"SrlgName"}

    return &(interfaceSrlgName.EntityData)
}

// Srlg_SrlgNames
// Set of SRLG name configuration
type Srlg_SrlgNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG name configuration. The type is slice of Srlg_SrlgNames_SrlgName.
    SrlgName []*Srlg_SrlgNames_SrlgName
}

func (srlgNames *Srlg_SrlgNames) GetEntityData() *types.CommonEntityData {
    srlgNames.EntityData.YFilter = srlgNames.YFilter
    srlgNames.EntityData.YangName = "srlg-names"
    srlgNames.EntityData.BundleName = "cisco_ios_xr"
    srlgNames.EntityData.ParentYangName = "srlg"
    srlgNames.EntityData.SegmentPath = "srlg-names"
    srlgNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgNames.EntityData.Children = types.NewOrderedMap()
    srlgNames.EntityData.Children.Append("srlg-name", types.YChild{"SrlgName", nil})
    for i := range srlgNames.SrlgName {
        srlgNames.EntityData.Children.Append(types.GetSegmentPath(srlgNames.SrlgName[i]), types.YChild{"SrlgName", srlgNames.SrlgName[i]})
    }
    srlgNames.EntityData.Leafs = types.NewOrderedMap()

    srlgNames.EntityData.YListKeys = []string {}

    return &(srlgNames.EntityData)
}

// Srlg_SrlgNames_SrlgName
// SRLG name configuration
type Srlg_SrlgNames_SrlgName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG name. The type is string with length: 1..64.
    SrlgName interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    SrlgValue interface{}
}

func (srlgName *Srlg_SrlgNames_SrlgName) GetEntityData() *types.CommonEntityData {
    srlgName.EntityData.YFilter = srlgName.YFilter
    srlgName.EntityData.YangName = "srlg-name"
    srlgName.EntityData.BundleName = "cisco_ios_xr"
    srlgName.EntityData.ParentYangName = "srlg-names"
    srlgName.EntityData.SegmentPath = "srlg-name" + types.AddKeyToken(srlgName.SrlgName, "srlg-name")
    srlgName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgName.EntityData.Children = types.NewOrderedMap()
    srlgName.EntityData.Leafs = types.NewOrderedMap()
    srlgName.EntityData.Leafs.Append("srlg-name", types.YLeaf{"SrlgName", srlgName.SrlgName})
    srlgName.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", srlgName.SrlgValue})

    srlgName.EntityData.YListKeys = []string {"SrlgName"}

    return &(srlgName.EntityData)
}

// Srlg_Groups
// Set of groups configured with SRLG
type Srlg_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group configurations. The type is slice of Srlg_Groups_Group.
    Group []*Srlg_Groups_Group
}

func (groups *Srlg_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "srlg"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// Srlg_Groups_Group
// Group configurations
type Srlg_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    GroupName interface{}

    // Enable SRLG group. The type is interface{}.
    Enable interface{}

    // Set of SRLG values configured under a group.
    GroupValues Srlg_Groups_Group_GroupValues
}

func (group *Srlg_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.GroupName, "group-name")
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Children.Append("group-values", types.YChild{"GroupValues", &group.GroupValues})
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", group.GroupName})
    group.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", group.Enable})

    group.EntityData.YListKeys = []string {"GroupName"}

    return &(group.EntityData)
}

// Srlg_Groups_Group_GroupValues
// Set of SRLG values configured under a group
type Srlg_Groups_Group_GroupValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group SRLG values with attribute. The type is slice of
    // Srlg_Groups_Group_GroupValues_GroupValue.
    GroupValue []*Srlg_Groups_Group_GroupValues_GroupValue
}

func (groupValues *Srlg_Groups_Group_GroupValues) GetEntityData() *types.CommonEntityData {
    groupValues.EntityData.YFilter = groupValues.YFilter
    groupValues.EntityData.YangName = "group-values"
    groupValues.EntityData.BundleName = "cisco_ios_xr"
    groupValues.EntityData.ParentYangName = "group"
    groupValues.EntityData.SegmentPath = "group-values"
    groupValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupValues.EntityData.Children = types.NewOrderedMap()
    groupValues.EntityData.Children.Append("group-value", types.YChild{"GroupValue", nil})
    for i := range groupValues.GroupValue {
        groupValues.EntityData.Children.Append(types.GetSegmentPath(groupValues.GroupValue[i]), types.YChild{"GroupValue", groupValues.GroupValue[i]})
    }
    groupValues.EntityData.Leafs = types.NewOrderedMap()

    groupValues.EntityData.YListKeys = []string {}

    return &(groupValues.EntityData)
}

// Srlg_Groups_Group_GroupValues_GroupValue
// Group SRLG values with attribute
type Srlg_Groups_Group_GroupValues_GroupValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG index. The type is interface{} with range:
    // 1..65535.
    SrlgIndex interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    SrlgValue interface{}

    // SRLG priority. The type is SrlgPriority. The default value is default.
    SrlgPriority interface{}
}

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetEntityData() *types.CommonEntityData {
    groupValue.EntityData.YFilter = groupValue.YFilter
    groupValue.EntityData.YangName = "group-value"
    groupValue.EntityData.BundleName = "cisco_ios_xr"
    groupValue.EntityData.ParentYangName = "group-values"
    groupValue.EntityData.SegmentPath = "group-value" + types.AddKeyToken(groupValue.SrlgIndex, "srlg-index")
    groupValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupValue.EntityData.Children = types.NewOrderedMap()
    groupValue.EntityData.Leafs = types.NewOrderedMap()
    groupValue.EntityData.Leafs.Append("srlg-index", types.YLeaf{"SrlgIndex", groupValue.SrlgIndex})
    groupValue.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", groupValue.SrlgValue})
    groupValue.EntityData.Leafs.Append("srlg-priority", types.YLeaf{"SrlgPriority", groupValue.SrlgPriority})

    groupValue.EntityData.YListKeys = []string {"SrlgIndex"}

    return &(groupValue.EntityData)
}

// Srlg_InheritNodes
// Set of inherit nodes configured with SRLG
type Srlg_InheritNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inherit node configurations. The type is slice of
    // Srlg_InheritNodes_InheritNode.
    InheritNode []*Srlg_InheritNodes_InheritNode
}

func (inheritNodes *Srlg_InheritNodes) GetEntityData() *types.CommonEntityData {
    inheritNodes.EntityData.YFilter = inheritNodes.YFilter
    inheritNodes.EntityData.YangName = "inherit-nodes"
    inheritNodes.EntityData.BundleName = "cisco_ios_xr"
    inheritNodes.EntityData.ParentYangName = "srlg"
    inheritNodes.EntityData.SegmentPath = "inherit-nodes"
    inheritNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritNodes.EntityData.Children = types.NewOrderedMap()
    inheritNodes.EntityData.Children.Append("inherit-node", types.YChild{"InheritNode", nil})
    for i := range inheritNodes.InheritNode {
        inheritNodes.EntityData.Children.Append(types.GetSegmentPath(inheritNodes.InheritNode[i]), types.YChild{"InheritNode", inheritNodes.InheritNode[i]})
    }
    inheritNodes.EntityData.Leafs = types.NewOrderedMap()

    inheritNodes.EntityData.YListKeys = []string {}

    return &(inheritNodes.EntityData)
}

// Srlg_InheritNodes_InheritNode
// Inherit node configurations
type Srlg_InheritNodes_InheritNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The inherit node name. The type is string with
    // pattern: ((([a-zA-Z0-9_]*\d+)|(\*))/){2}(([a-zA-Z0-9_]*\d+)|(\*)).
    InheritNodeName interface{}

    // Enable SRLG inherit node. The type is interface{}.
    Enable interface{}

    // Set of SRLG values configured under an inherit node.
    InheritNodeValues Srlg_InheritNodes_InheritNode_InheritNodeValues
}

func (inheritNode *Srlg_InheritNodes_InheritNode) GetEntityData() *types.CommonEntityData {
    inheritNode.EntityData.YFilter = inheritNode.YFilter
    inheritNode.EntityData.YangName = "inherit-node"
    inheritNode.EntityData.BundleName = "cisco_ios_xr"
    inheritNode.EntityData.ParentYangName = "inherit-nodes"
    inheritNode.EntityData.SegmentPath = "inherit-node" + types.AddKeyToken(inheritNode.InheritNodeName, "inherit-node-name")
    inheritNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritNode.EntityData.Children = types.NewOrderedMap()
    inheritNode.EntityData.Children.Append("inherit-node-values", types.YChild{"InheritNodeValues", &inheritNode.InheritNodeValues})
    inheritNode.EntityData.Leafs = types.NewOrderedMap()
    inheritNode.EntityData.Leafs.Append("inherit-node-name", types.YLeaf{"InheritNodeName", inheritNode.InheritNodeName})
    inheritNode.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", inheritNode.Enable})

    inheritNode.EntityData.YListKeys = []string {"InheritNodeName"}

    return &(inheritNode.EntityData)
}

// Srlg_InheritNodes_InheritNode_InheritNodeValues
// Set of SRLG values configured under an inherit
// node
type Srlg_InheritNodes_InheritNode_InheritNodeValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inherit node SRLG value with attributes. The type is slice of
    // Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue.
    InheritNodeValue []*Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue
}

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetEntityData() *types.CommonEntityData {
    inheritNodeValues.EntityData.YFilter = inheritNodeValues.YFilter
    inheritNodeValues.EntityData.YangName = "inherit-node-values"
    inheritNodeValues.EntityData.BundleName = "cisco_ios_xr"
    inheritNodeValues.EntityData.ParentYangName = "inherit-node"
    inheritNodeValues.EntityData.SegmentPath = "inherit-node-values"
    inheritNodeValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritNodeValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritNodeValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritNodeValues.EntityData.Children = types.NewOrderedMap()
    inheritNodeValues.EntityData.Children.Append("inherit-node-value", types.YChild{"InheritNodeValue", nil})
    for i := range inheritNodeValues.InheritNodeValue {
        inheritNodeValues.EntityData.Children.Append(types.GetSegmentPath(inheritNodeValues.InheritNodeValue[i]), types.YChild{"InheritNodeValue", inheritNodeValues.InheritNodeValue[i]})
    }
    inheritNodeValues.EntityData.Leafs = types.NewOrderedMap()

    inheritNodeValues.EntityData.YListKeys = []string {}

    return &(inheritNodeValues.EntityData)
}

// Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue
// Inherit node SRLG value with attributes
type Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG index. The type is interface{} with range:
    // 1..65535.
    SrlgIndex interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    SrlgValue interface{}

    // SRLG priority. The type is SrlgPriority. The default value is default.
    SrlgPriority interface{}
}

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetEntityData() *types.CommonEntityData {
    inheritNodeValue.EntityData.YFilter = inheritNodeValue.YFilter
    inheritNodeValue.EntityData.YangName = "inherit-node-value"
    inheritNodeValue.EntityData.BundleName = "cisco_ios_xr"
    inheritNodeValue.EntityData.ParentYangName = "inherit-node-values"
    inheritNodeValue.EntityData.SegmentPath = "inherit-node-value" + types.AddKeyToken(inheritNodeValue.SrlgIndex, "srlg-index")
    inheritNodeValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritNodeValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritNodeValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritNodeValue.EntityData.Children = types.NewOrderedMap()
    inheritNodeValue.EntityData.Leafs = types.NewOrderedMap()
    inheritNodeValue.EntityData.Leafs.Append("srlg-index", types.YLeaf{"SrlgIndex", inheritNodeValue.SrlgIndex})
    inheritNodeValue.EntityData.Leafs.Append("srlg-value", types.YLeaf{"SrlgValue", inheritNodeValue.SrlgValue})
    inheritNodeValue.EntityData.Leafs.Append("srlg-priority", types.YLeaf{"SrlgPriority", inheritNodeValue.SrlgPriority})

    inheritNodeValue.EntityData.YListKeys = []string {"SrlgIndex"}

    return &(inheritNodeValue.EntityData)
}

// VrfGroups
// vrf groups
type VrfGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF group configuration. The type is slice of VrfGroups_VrfGroup.
    VrfGroup []*VrfGroups_VrfGroup
}

func (vrfGroups *VrfGroups) GetEntityData() *types.CommonEntityData {
    vrfGroups.EntityData.YFilter = vrfGroups.YFilter
    vrfGroups.EntityData.YangName = "vrf-groups"
    vrfGroups.EntityData.BundleName = "cisco_ios_xr"
    vrfGroups.EntityData.ParentYangName = "Cisco-IOS-XR-infra-rsi-cfg"
    vrfGroups.EntityData.SegmentPath = "Cisco-IOS-XR-infra-rsi-cfg:vrf-groups"
    vrfGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfGroups.EntityData.Children = types.NewOrderedMap()
    vrfGroups.EntityData.Children.Append("vrf-group", types.YChild{"VrfGroup", nil})
    for i := range vrfGroups.VrfGroup {
        vrfGroups.EntityData.Children.Append(types.GetSegmentPath(vrfGroups.VrfGroup[i]), types.YChild{"VrfGroup", vrfGroups.VrfGroup[i]})
    }
    vrfGroups.EntityData.Leafs = types.NewOrderedMap()

    vrfGroups.EntityData.YListKeys = []string {}

    return &(vrfGroups.EntityData)
}

// VrfGroups_VrfGroup
// VRF group configuration
type VrfGroups_VrfGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF group name. The type is string with length:
    // 1..32.
    VrfGroupName interface{}

    // Enable VRF group. The type is interface{}.
    Enable interface{}

    // Set of VRFs configured under a VRF group.
    Vrfs VrfGroups_VrfGroup_Vrfs
}

func (vrfGroup *VrfGroups_VrfGroup) GetEntityData() *types.CommonEntityData {
    vrfGroup.EntityData.YFilter = vrfGroup.YFilter
    vrfGroup.EntityData.YangName = "vrf-group"
    vrfGroup.EntityData.BundleName = "cisco_ios_xr"
    vrfGroup.EntityData.ParentYangName = "vrf-groups"
    vrfGroup.EntityData.SegmentPath = "vrf-group" + types.AddKeyToken(vrfGroup.VrfGroupName, "vrf-group-name")
    vrfGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfGroup.EntityData.Children = types.NewOrderedMap()
    vrfGroup.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &vrfGroup.Vrfs})
    vrfGroup.EntityData.Leafs = types.NewOrderedMap()
    vrfGroup.EntityData.Leafs.Append("vrf-group-name", types.YLeaf{"VrfGroupName", vrfGroup.VrfGroupName})
    vrfGroup.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", vrfGroup.Enable})

    vrfGroup.EntityData.YListKeys = []string {"VrfGroupName"}

    return &(vrfGroup.EntityData)
}

// VrfGroups_VrfGroup_Vrfs
// Set of VRFs configured under a VRF group
type VrfGroups_VrfGroup_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF configuration. The type is slice of VrfGroups_VrfGroup_Vrfs_Vrf.
    Vrf []*VrfGroups_VrfGroup_Vrfs_Vrf
}

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "vrf-group"
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

// VrfGroups_VrfGroup_Vrfs_Vrf
// VRF configuration
type VrfGroups_VrfGroup_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}
}

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// SelectiveVrfDownload
// selective vrf download
type SelectiveVrfDownload struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable selective VRF download. The type is interface{}.
    Disable interface{}
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetEntityData() *types.CommonEntityData {
    selectiveVrfDownload.EntityData.YFilter = selectiveVrfDownload.YFilter
    selectiveVrfDownload.EntityData.YangName = "selective-vrf-download"
    selectiveVrfDownload.EntityData.BundleName = "cisco_ios_xr"
    selectiveVrfDownload.EntityData.ParentYangName = "Cisco-IOS-XR-infra-rsi-cfg"
    selectiveVrfDownload.EntityData.SegmentPath = "Cisco-IOS-XR-infra-rsi-cfg:selective-vrf-download"
    selectiveVrfDownload.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectiveVrfDownload.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectiveVrfDownload.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectiveVrfDownload.EntityData.Children = types.NewOrderedMap()
    selectiveVrfDownload.EntityData.Leafs = types.NewOrderedMap()
    selectiveVrfDownload.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", selectiveVrfDownload.Disable})

    selectiveVrfDownload.EntityData.YListKeys = []string {}

    return &(selectiveVrfDownload.EntityData)
}

