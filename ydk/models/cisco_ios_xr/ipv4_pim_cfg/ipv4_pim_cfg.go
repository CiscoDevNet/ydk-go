// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-pim package configuration.
// 
// This module contains definitions
// for the following management objects:
//   pim: PIM configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_pim_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_pim_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-pim-cfg pim}", reflect.TypeOf(Pim{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-pim-cfg:pim", reflect.TypeOf(Pim{}))
}

// PimProtocolMode represents Pim protocol mode
type PimProtocolMode string

const (
    // Sparse Mode
    PimProtocolMode_sm PimProtocolMode = "sm"

    // Bidirectional
    PimProtocolMode_bidir PimProtocolMode = "bidir"
)

// PimMultipath represents Pim multipath
type PimMultipath string

const (
    // Enable PIM multipath
    PimMultipath_enable PimMultipath = "enable"

    // Enable PIM multipath with interface based
    // hashing
    PimMultipath_interface_hash PimMultipath = "interface-hash"

    // Enable PIM multipath with source based hashing
    PimMultipath_source_hash PimMultipath = "source-hash"

    // Enable PIM multipath with source next-hop
    // hashing
    PimMultipath_source_next_hop_hash PimMultipath = "source-next-hop-hash"

    // Enable PIM multipath with source, group based
    // hashing
    PimMultipath_source_group_hash PimMultipath = "source-group-hash"
)

// Pim
// PIM configuration
type Pim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF table.
    Vrfs Pim_Vrfs

    // Default Context.
    DefaultContext Pim_DefaultContext
}

func (pim *Pim) GetEntityData() *types.CommonEntityData {
    pim.EntityData.YFilter = pim.YFilter
    pim.EntityData.YangName = "pim"
    pim.EntityData.BundleName = "cisco_ios_xr"
    pim.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-pim-cfg"
    pim.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-pim-cfg:pim"
    pim.EntityData.AbsolutePath = pim.EntityData.SegmentPath
    pim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pim.EntityData.Children = types.NewOrderedMap()
    pim.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &pim.Vrfs})
    pim.EntityData.Children.Append("default-context", types.YChild{"DefaultContext", &pim.DefaultContext})
    pim.EntityData.Leafs = types.NewOrderedMap()

    pim.EntityData.YListKeys = []string {}

    return &(pim.EntityData)
}

// Pim_Vrfs
// VRF table
type Pim_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is slice of Pim_Vrfs_Vrf.
    Vrf []*Pim_Vrfs_Vrf
}

func (vrfs *Pim_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "pim"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/" + vrfs.EntityData.SegmentPath
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

// Pim_Vrfs_Vrf
// VRF name
type Pim_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // IPV4 commands.
    Ipv4 Pim_Vrfs_Vrf_Ipv4

    // IPV6 commands.
    Ipv6 Pim_Vrfs_Vrf_Ipv6
}

func (vrf *Pim_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &vrf.Ipv4})
    vrf.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &vrf.Ipv6})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4
// IPV4 commands
type Pim_Vrfs_Vrf_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable PIM neighbor checking when receiving PIM messages. The type is
    // interface{}.
    NeighborCheckOnReceive interface{}

    // Generate registers compatible with older IOS versions. The type is
    // interface{}.
    OldRegisterChecksum interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Configure threshold of infinity for switching to SPT on last-hop. The type
    // is string.
    SptThresholdInfinity interface{}

    // PIM neighbor state change logging is turned on if configured. The type is
    // interface{}.
    LogNeighborChanges interface{}

    // Source address to use for register messages. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    RegisterSource interface{}

    // Access-list which specifies unauthorized sources. The type is string with
    // length: 1..64.
    AcceptRegister interface{}

    // Suppress prunes triggered as a result of RPF changes. The type is
    // interface{}.
    SuppressRpfPrunes interface{}

    // Allow SSM ranges to be overridden by more specific ranges. The type is
    // interface{}.
    SsmAllowOverride interface{}

    // Enable equal-cost multipath routing. The type is PimMultipath.
    Multipath interface{}

    // Configure static RP deny range. The type is string with length: 1..64.
    RpStaticDeny interface{}

    // Suppress data registers after initial state setup. The type is interface{}.
    SuppressDataRegisters interface{}

    // Enable PIM neighbor checking when sending join-prunes. The type is
    // interface{}.
    NeighborCheckOnSend interface{}

    // Disable Rendezvous Point discovery through the AutoRP protocol. The type is
    // interface{}.
    AutoRpDisable interface{}

    // Configure Sparse-Mode Rendezvous Point.
    SparseModeRpAddresses Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses

    // Inheritable defaults.
    InheritableDefaults Pim_Vrfs_Vrf_Ipv4_InheritableDefaults

    // Configure RPF options.
    Rpf Pim_Vrfs_Vrf_Ipv4_Rpf

    // Configure PIM State Limits.
    Maximum Pim_Vrfs_Vrf_Ipv4_Maximum

    // Configure expiry timer for S,G routes.
    SgExpiryTimer Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer

    // Enable PIM RPF Vector Proxy's.
    RpfVectorEnable Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable

    // Configure IP Multicast SSM.
    Ssm Pim_Vrfs_Vrf_Ipv4_Ssm

    // Inject Explicit PIM RPF Vector Proxy's.
    Injects Pim_Vrfs_Vrf_Ipv4_Injects

    // Configure Bidirectional PIM Rendezvous Point.
    BidirRpAddresses Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses

    // PIM BSR configuration.
    Bsr Pim_Vrfs_Vrf_Ipv4_Bsr

    // Multicast Only Fast Re-Route.
    Mofrr Pim_Vrfs_Vrf_Ipv4_Mofrr

    // Inject PIM RPF Vector Proxy's.
    Paths Pim_Vrfs_Vrf_Ipv4_Paths

    // Enable allow-rp filtering for SM joins.
    AllowRp Pim_Vrfs_Vrf_Ipv4_AllowRp

    // Configure convergence parameters.
    Convergence Pim_Vrfs_Vrf_Ipv4_Convergence

    // Interface-level Configuration.
    Interfaces Pim_Vrfs_Vrf_Ipv4_Interfaces
}

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "vrf"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("sparse-mode-rp-addresses", types.YChild{"SparseModeRpAddresses", &ipv4.SparseModeRpAddresses})
    ipv4.EntityData.Children.Append("inheritable-defaults", types.YChild{"InheritableDefaults", &ipv4.InheritableDefaults})
    ipv4.EntityData.Children.Append("rpf", types.YChild{"Rpf", &ipv4.Rpf})
    ipv4.EntityData.Children.Append("maximum", types.YChild{"Maximum", &ipv4.Maximum})
    ipv4.EntityData.Children.Append("sg-expiry-timer", types.YChild{"SgExpiryTimer", &ipv4.SgExpiryTimer})
    ipv4.EntityData.Children.Append("rpf-vector-enable", types.YChild{"RpfVectorEnable", &ipv4.RpfVectorEnable})
    ipv4.EntityData.Children.Append("ssm", types.YChild{"Ssm", &ipv4.Ssm})
    ipv4.EntityData.Children.Append("injects", types.YChild{"Injects", &ipv4.Injects})
    ipv4.EntityData.Children.Append("bidir-rp-addresses", types.YChild{"BidirRpAddresses", &ipv4.BidirRpAddresses})
    ipv4.EntityData.Children.Append("bsr", types.YChild{"Bsr", &ipv4.Bsr})
    ipv4.EntityData.Children.Append("mofrr", types.YChild{"Mofrr", &ipv4.Mofrr})
    ipv4.EntityData.Children.Append("paths", types.YChild{"Paths", &ipv4.Paths})
    ipv4.EntityData.Children.Append("allow-rp", types.YChild{"AllowRp", &ipv4.AllowRp})
    ipv4.EntityData.Children.Append("convergence", types.YChild{"Convergence", &ipv4.Convergence})
    ipv4.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &ipv4.Interfaces})
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("neighbor-check-on-receive", types.YLeaf{"NeighborCheckOnReceive", ipv4.NeighborCheckOnReceive})
    ipv4.EntityData.Leafs.Append("old-register-checksum", types.YLeaf{"OldRegisterChecksum", ipv4.OldRegisterChecksum})
    ipv4.EntityData.Leafs.Append("neighbor-filter", types.YLeaf{"NeighborFilter", ipv4.NeighborFilter})
    ipv4.EntityData.Leafs.Append("spt-threshold-infinity", types.YLeaf{"SptThresholdInfinity", ipv4.SptThresholdInfinity})
    ipv4.EntityData.Leafs.Append("log-neighbor-changes", types.YLeaf{"LogNeighborChanges", ipv4.LogNeighborChanges})
    ipv4.EntityData.Leafs.Append("register-source", types.YLeaf{"RegisterSource", ipv4.RegisterSource})
    ipv4.EntityData.Leafs.Append("accept-register", types.YLeaf{"AcceptRegister", ipv4.AcceptRegister})
    ipv4.EntityData.Leafs.Append("suppress-rpf-prunes", types.YLeaf{"SuppressRpfPrunes", ipv4.SuppressRpfPrunes})
    ipv4.EntityData.Leafs.Append("ssm-allow-override", types.YLeaf{"SsmAllowOverride", ipv4.SsmAllowOverride})
    ipv4.EntityData.Leafs.Append("multipath", types.YLeaf{"Multipath", ipv4.Multipath})
    ipv4.EntityData.Leafs.Append("rp-static-deny", types.YLeaf{"RpStaticDeny", ipv4.RpStaticDeny})
    ipv4.EntityData.Leafs.Append("suppress-data-registers", types.YLeaf{"SuppressDataRegisters", ipv4.SuppressDataRegisters})
    ipv4.EntityData.Leafs.Append("neighbor-check-on-send", types.YLeaf{"NeighborCheckOnSend", ipv4.NeighborCheckOnSend})
    ipv4.EntityData.Leafs.Append("auto-rp-disable", types.YLeaf{"AutoRpDisable", ipv4.AutoRpDisable})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses
// Configure Sparse-Mode Rendezvous Point
type Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress.
    SparseModeRpAddress []*Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetEntityData() *types.CommonEntityData {
    sparseModeRpAddresses.EntityData.YFilter = sparseModeRpAddresses.YFilter
    sparseModeRpAddresses.EntityData.YangName = "sparse-mode-rp-addresses"
    sparseModeRpAddresses.EntityData.BundleName = "cisco_ios_xr"
    sparseModeRpAddresses.EntityData.ParentYangName = "ipv4"
    sparseModeRpAddresses.EntityData.SegmentPath = "sparse-mode-rp-addresses"
    sparseModeRpAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + sparseModeRpAddresses.EntityData.SegmentPath
    sparseModeRpAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sparseModeRpAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sparseModeRpAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sparseModeRpAddresses.EntityData.Children = types.NewOrderedMap()
    sparseModeRpAddresses.EntityData.Children.Append("sparse-mode-rp-address", types.YChild{"SparseModeRpAddress", nil})
    for i := range sparseModeRpAddresses.SparseModeRpAddress {
        sparseModeRpAddresses.EntityData.Children.Append(types.GetSegmentPath(sparseModeRpAddresses.SparseModeRpAddress[i]), types.YChild{"SparseModeRpAddress", sparseModeRpAddresses.SparseModeRpAddress[i]})
    }
    sparseModeRpAddresses.EntityData.Leafs = types.NewOrderedMap()

    sparseModeRpAddresses.EntityData.YListKeys = []string {}

    return &(sparseModeRpAddresses.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress
// Address of the Rendezvous Point
type Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a  given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetEntityData() *types.CommonEntityData {
    sparseModeRpAddress.EntityData.YFilter = sparseModeRpAddress.YFilter
    sparseModeRpAddress.EntityData.YangName = "sparse-mode-rp-address"
    sparseModeRpAddress.EntityData.BundleName = "cisco_ios_xr"
    sparseModeRpAddress.EntityData.ParentYangName = "sparse-mode-rp-addresses"
    sparseModeRpAddress.EntityData.SegmentPath = "sparse-mode-rp-address" + types.AddKeyToken(sparseModeRpAddress.RpAddress, "rp-address")
    sparseModeRpAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/sparse-mode-rp-addresses/" + sparseModeRpAddress.EntityData.SegmentPath
    sparseModeRpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sparseModeRpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sparseModeRpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sparseModeRpAddress.EntityData.Children = types.NewOrderedMap()
    sparseModeRpAddress.EntityData.Leafs = types.NewOrderedMap()
    sparseModeRpAddress.EntityData.Leafs.Append("rp-address", types.YLeaf{"RpAddress", sparseModeRpAddress.RpAddress})
    sparseModeRpAddress.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", sparseModeRpAddress.AccessListName})
    sparseModeRpAddress.EntityData.Leafs.Append("auto-rp-override", types.YLeaf{"AutoRpOverride", sparseModeRpAddress.AutoRpOverride})

    sparseModeRpAddress.EntityData.YListKeys = []string {"RpAddress"}

    return &(sparseModeRpAddress.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_InheritableDefaults
// Inheritable defaults
type Pim_Vrfs_Vrf_Ipv4_InheritableDefaults struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Convergency timeout in seconds. The type is interface{} with range:
    // 1800..2400. Units are second.
    ConvergenceTimeout interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetEntityData() *types.CommonEntityData {
    inheritableDefaults.EntityData.YFilter = inheritableDefaults.YFilter
    inheritableDefaults.EntityData.YangName = "inheritable-defaults"
    inheritableDefaults.EntityData.BundleName = "cisco_ios_xr"
    inheritableDefaults.EntityData.ParentYangName = "ipv4"
    inheritableDefaults.EntityData.SegmentPath = "inheritable-defaults"
    inheritableDefaults.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + inheritableDefaults.EntityData.SegmentPath
    inheritableDefaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritableDefaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritableDefaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritableDefaults.EntityData.Children = types.NewOrderedMap()
    inheritableDefaults.EntityData.Leafs = types.NewOrderedMap()
    inheritableDefaults.EntityData.Leafs.Append("convergence-timeout", types.YLeaf{"ConvergenceTimeout", inheritableDefaults.ConvergenceTimeout})
    inheritableDefaults.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", inheritableDefaults.HelloInterval})
    inheritableDefaults.EntityData.Leafs.Append("propagation-delay", types.YLeaf{"PropagationDelay", inheritableDefaults.PropagationDelay})
    inheritableDefaults.EntityData.Leafs.Append("dr-priority", types.YLeaf{"DrPriority", inheritableDefaults.DrPriority})
    inheritableDefaults.EntityData.Leafs.Append("join-prune-mtu", types.YLeaf{"JoinPruneMtu", inheritableDefaults.JoinPruneMtu})
    inheritableDefaults.EntityData.Leafs.Append("jp-interval", types.YLeaf{"JpInterval", inheritableDefaults.JpInterval})
    inheritableDefaults.EntityData.Leafs.Append("override-interval", types.YLeaf{"OverrideInterval", inheritableDefaults.OverrideInterval})

    inheritableDefaults.EntityData.YListKeys = []string {}

    return &(inheritableDefaults.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Rpf
// Configure RPF options
type Pim_Vrfs_Vrf_Ipv4_Rpf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route policy to select RPF topology. The type is string with length: 1..64.
    RoutePolicy interface{}
}

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetEntityData() *types.CommonEntityData {
    rpf.EntityData.YFilter = rpf.YFilter
    rpf.EntityData.YangName = "rpf"
    rpf.EntityData.BundleName = "cisco_ios_xr"
    rpf.EntityData.ParentYangName = "ipv4"
    rpf.EntityData.SegmentPath = "rpf"
    rpf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + rpf.EntityData.SegmentPath
    rpf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpf.EntityData.Children = types.NewOrderedMap()
    rpf.EntityData.Leafs = types.NewOrderedMap()
    rpf.EntityData.Leafs.Append("route-policy", types.YLeaf{"RoutePolicy", rpf.RoutePolicy})

    rpf.EntityData.YListKeys = []string {}

    return &(rpf.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Maximum
// Configure PIM State Limits
type Pim_Vrfs_Vrf_Ipv4_Maximum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Override default maximum for number of group mappings from autorp mapping
    // agent.
    GroupMappingsAutoRp Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp

    // Override default maximum and threshold for number of group mappings from
    // BSR.
    BsrGroupMappings Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings

    // Override default maximum for number of sparse-mode source registers.
    RegisterStates Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates

    // Override default maximum for number of route-interfaces.
    RouteInterfaces Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces

    // Override default maximum and threshold for BSR C-RP cache setting.
    BsrCandidateRpCache Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache

    // Override default maximum for number of routes.
    Routes Pim_Vrfs_Vrf_Ipv4_Maximum_Routes
}

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetEntityData() *types.CommonEntityData {
    maximum.EntityData.YFilter = maximum.YFilter
    maximum.EntityData.YangName = "maximum"
    maximum.EntityData.BundleName = "cisco_ios_xr"
    maximum.EntityData.ParentYangName = "ipv4"
    maximum.EntityData.SegmentPath = "maximum"
    maximum.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + maximum.EntityData.SegmentPath
    maximum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximum.EntityData.Children = types.NewOrderedMap()
    maximum.EntityData.Children.Append("group-mappings-auto-rp", types.YChild{"GroupMappingsAutoRp", &maximum.GroupMappingsAutoRp})
    maximum.EntityData.Children.Append("bsr-group-mappings", types.YChild{"BsrGroupMappings", &maximum.BsrGroupMappings})
    maximum.EntityData.Children.Append("register-states", types.YChild{"RegisterStates", &maximum.RegisterStates})
    maximum.EntityData.Children.Append("route-interfaces", types.YChild{"RouteInterfaces", &maximum.RouteInterfaces})
    maximum.EntityData.Children.Append("bsr-candidate-rp-cache", types.YChild{"BsrCandidateRpCache", &maximum.BsrCandidateRpCache})
    maximum.EntityData.Children.Append("routes", types.YChild{"Routes", &maximum.Routes})
    maximum.EntityData.Leafs = types.NewOrderedMap()

    maximum.EntityData.YListKeys = []string {}

    return &(maximum.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp
// Override default maximum for number of group
// mappings from autorp mapping agent
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM group mappings from autorp. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    MaximumGroupRangesAutoRp interface{}

    // Warning threshold number of PIM group mappings from autorp. The type is
    // interface{} with range: 1..10000. The default value is 450.
    ThresholdGroupRangesAutoRp interface{}
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetEntityData() *types.CommonEntityData {
    groupMappingsAutoRp.EntityData.YFilter = groupMappingsAutoRp.YFilter
    groupMappingsAutoRp.EntityData.YangName = "group-mappings-auto-rp"
    groupMappingsAutoRp.EntityData.BundleName = "cisco_ios_xr"
    groupMappingsAutoRp.EntityData.ParentYangName = "maximum"
    groupMappingsAutoRp.EntityData.SegmentPath = "group-mappings-auto-rp"
    groupMappingsAutoRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/maximum/" + groupMappingsAutoRp.EntityData.SegmentPath
    groupMappingsAutoRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupMappingsAutoRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupMappingsAutoRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupMappingsAutoRp.EntityData.Children = types.NewOrderedMap()
    groupMappingsAutoRp.EntityData.Leafs = types.NewOrderedMap()
    groupMappingsAutoRp.EntityData.Leafs.Append("maximum-group-ranges-auto-rp", types.YLeaf{"MaximumGroupRangesAutoRp", groupMappingsAutoRp.MaximumGroupRangesAutoRp})
    groupMappingsAutoRp.EntityData.Leafs.Append("threshold-group-ranges-auto-rp", types.YLeaf{"ThresholdGroupRangesAutoRp", groupMappingsAutoRp.ThresholdGroupRangesAutoRp})

    groupMappingsAutoRp.EntityData.YListKeys = []string {}

    return &(groupMappingsAutoRp.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings
// Override default maximum and threshold for
// number of group mappings from BSR
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM group mappings from BSR. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumGroupRanges interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 500.
    WarningThreshold interface{}
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetEntityData() *types.CommonEntityData {
    bsrGroupMappings.EntityData.YFilter = bsrGroupMappings.YFilter
    bsrGroupMappings.EntityData.YangName = "bsr-group-mappings"
    bsrGroupMappings.EntityData.BundleName = "cisco_ios_xr"
    bsrGroupMappings.EntityData.ParentYangName = "maximum"
    bsrGroupMappings.EntityData.SegmentPath = "bsr-group-mappings"
    bsrGroupMappings.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/maximum/" + bsrGroupMappings.EntityData.SegmentPath
    bsrGroupMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsrGroupMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsrGroupMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsrGroupMappings.EntityData.Children = types.NewOrderedMap()
    bsrGroupMappings.EntityData.Leafs = types.NewOrderedMap()
    bsrGroupMappings.EntityData.Leafs.Append("bsr-maximum-group-ranges", types.YLeaf{"BsrMaximumGroupRanges", bsrGroupMappings.BsrMaximumGroupRanges})
    bsrGroupMappings.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", bsrGroupMappings.WarningThreshold})

    bsrGroupMappings.EntityData.YListKeys = []string {}

    return &(bsrGroupMappings.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates
// Override default maximum for number of
// sparse-mode source registers
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM Sparse-Mode register states. The type is interface{}
    // with range: 0..75000. This attribute is mandatory.
    MaximumRegisterStates interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 0..75000. The default value is 20000.
    WarningThreshold interface{}
}

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetEntityData() *types.CommonEntityData {
    registerStates.EntityData.YFilter = registerStates.YFilter
    registerStates.EntityData.YangName = "register-states"
    registerStates.EntityData.BundleName = "cisco_ios_xr"
    registerStates.EntityData.ParentYangName = "maximum"
    registerStates.EntityData.SegmentPath = "register-states"
    registerStates.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/maximum/" + registerStates.EntityData.SegmentPath
    registerStates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registerStates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registerStates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registerStates.EntityData.Children = types.NewOrderedMap()
    registerStates.EntityData.Leafs = types.NewOrderedMap()
    registerStates.EntityData.Leafs.Append("maximum-register-states", types.YLeaf{"MaximumRegisterStates", registerStates.MaximumRegisterStates})
    registerStates.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", registerStates.WarningThreshold})

    registerStates.EntityData.YListKeys = []string {}

    return &(registerStates.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces
// Override default maximum for number of
// route-interfaces
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM route-interfaces. The type is interface{} with range:
    // 1..1100000. This attribute is mandatory.
    MaximumRouteInterfaces interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000. The default value is 300000.
    WarningThreshold interface{}
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetEntityData() *types.CommonEntityData {
    routeInterfaces.EntityData.YFilter = routeInterfaces.YFilter
    routeInterfaces.EntityData.YangName = "route-interfaces"
    routeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    routeInterfaces.EntityData.ParentYangName = "maximum"
    routeInterfaces.EntityData.SegmentPath = "route-interfaces"
    routeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/maximum/" + routeInterfaces.EntityData.SegmentPath
    routeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeInterfaces.EntityData.Children = types.NewOrderedMap()
    routeInterfaces.EntityData.Leafs = types.NewOrderedMap()
    routeInterfaces.EntityData.Leafs.Append("maximum-route-interfaces", types.YLeaf{"MaximumRouteInterfaces", routeInterfaces.MaximumRouteInterfaces})
    routeInterfaces.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", routeInterfaces.WarningThreshold})

    routeInterfaces.EntityData.YListKeys = []string {}

    return &(routeInterfaces.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache
// Override default maximum and threshold for BSR
// C-RP cache setting
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of BSR C-RP cache setting. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumCandidateRpCache interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 100.
    WarningThreshold interface{}
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetEntityData() *types.CommonEntityData {
    bsrCandidateRpCache.EntityData.YFilter = bsrCandidateRpCache.YFilter
    bsrCandidateRpCache.EntityData.YangName = "bsr-candidate-rp-cache"
    bsrCandidateRpCache.EntityData.BundleName = "cisco_ios_xr"
    bsrCandidateRpCache.EntityData.ParentYangName = "maximum"
    bsrCandidateRpCache.EntityData.SegmentPath = "bsr-candidate-rp-cache"
    bsrCandidateRpCache.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/maximum/" + bsrCandidateRpCache.EntityData.SegmentPath
    bsrCandidateRpCache.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsrCandidateRpCache.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsrCandidateRpCache.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsrCandidateRpCache.EntityData.Children = types.NewOrderedMap()
    bsrCandidateRpCache.EntityData.Leafs = types.NewOrderedMap()
    bsrCandidateRpCache.EntityData.Leafs.Append("bsr-maximum-candidate-rp-cache", types.YLeaf{"BsrMaximumCandidateRpCache", bsrCandidateRpCache.BsrMaximumCandidateRpCache})
    bsrCandidateRpCache.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", bsrCandidateRpCache.WarningThreshold})

    bsrCandidateRpCache.EntityData.YListKeys = []string {}

    return &(bsrCandidateRpCache.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Maximum_Routes
// Override default maximum for number of routes
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Maximum_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM routes. The type is interface{} with range:
    // 1..200000. This attribute is mandatory.
    MaximumRoutes interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..200000. The default value is 100000.
    WarningThreshold interface{}
}

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "maximum"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/maximum/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Leafs = types.NewOrderedMap()
    routes.EntityData.Leafs.Append("maximum-routes", types.YLeaf{"MaximumRoutes", routes.MaximumRoutes})
    routes.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", routes.WarningThreshold})

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer
// Configure expiry timer for S,G routes
type Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // (S,G) expiry time in seconds. The type is interface{} with range:
    // 40..57600. Units are second.
    Interval interface{}

    // Access-list of applicable S,G routes. The type is string with length:
    // 1..64.
    AccessListName interface{}
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetEntityData() *types.CommonEntityData {
    sgExpiryTimer.EntityData.YFilter = sgExpiryTimer.YFilter
    sgExpiryTimer.EntityData.YangName = "sg-expiry-timer"
    sgExpiryTimer.EntityData.BundleName = "cisco_ios_xr"
    sgExpiryTimer.EntityData.ParentYangName = "ipv4"
    sgExpiryTimer.EntityData.SegmentPath = "sg-expiry-timer"
    sgExpiryTimer.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + sgExpiryTimer.EntityData.SegmentPath
    sgExpiryTimer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sgExpiryTimer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sgExpiryTimer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sgExpiryTimer.EntityData.Children = types.NewOrderedMap()
    sgExpiryTimer.EntityData.Leafs = types.NewOrderedMap()
    sgExpiryTimer.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", sgExpiryTimer.Interval})
    sgExpiryTimer.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", sgExpiryTimer.AccessListName})

    sgExpiryTimer.EntityData.YListKeys = []string {}

    return &(sgExpiryTimer.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable
// Enable PIM RPF Vector Proxy's
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // RPF Vector is turned on if configured. The type is interface{}. This
    // attribute is mandatory.
    Enable interface{}

    // Allow RPF Vector origination over eBGP sessions. The type is interface{}.
    AllowEbgp interface{}

    // Disable RPF Vector origination over iBGP sessions. The type is interface{}.
    DisableIbgp interface{}

    // Use RPF Vector RFC compliant encoding. The type is interface{}.
    UseStandardEncoding interface{}
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetEntityData() *types.CommonEntityData {
    rpfVectorEnable.EntityData.YFilter = rpfVectorEnable.YFilter
    rpfVectorEnable.EntityData.YangName = "rpf-vector-enable"
    rpfVectorEnable.EntityData.BundleName = "cisco_ios_xr"
    rpfVectorEnable.EntityData.ParentYangName = "ipv4"
    rpfVectorEnable.EntityData.SegmentPath = "rpf-vector-enable"
    rpfVectorEnable.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + rpfVectorEnable.EntityData.SegmentPath
    rpfVectorEnable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpfVectorEnable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpfVectorEnable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpfVectorEnable.EntityData.Children = types.NewOrderedMap()
    rpfVectorEnable.EntityData.Leafs = types.NewOrderedMap()
    rpfVectorEnable.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", rpfVectorEnable.Enable})
    rpfVectorEnable.EntityData.Leafs.Append("allow-ebgp", types.YLeaf{"AllowEbgp", rpfVectorEnable.AllowEbgp})
    rpfVectorEnable.EntityData.Leafs.Append("disable-ibgp", types.YLeaf{"DisableIbgp", rpfVectorEnable.DisableIbgp})
    rpfVectorEnable.EntityData.Leafs.Append("use-standard-encoding", types.YLeaf{"UseStandardEncoding", rpfVectorEnable.UseStandardEncoding})

    rpfVectorEnable.EntityData.YListKeys = []string {}

    return &(rpfVectorEnable.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Ssm
// Configure IP Multicast SSM
type Pim_Vrfs_Vrf_Ipv4_Ssm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE if SSM is disabled on this router. The type is bool. The default value
    // is false.
    Disable interface{}

    // Access list of groups enabled with SSM. The type is string with length:
    // 1..64.
    Range interface{}
}

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetEntityData() *types.CommonEntityData {
    ssm.EntityData.YFilter = ssm.YFilter
    ssm.EntityData.YangName = "ssm"
    ssm.EntityData.BundleName = "cisco_ios_xr"
    ssm.EntityData.ParentYangName = "ipv4"
    ssm.EntityData.SegmentPath = "ssm"
    ssm.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + ssm.EntityData.SegmentPath
    ssm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssm.EntityData.Children = types.NewOrderedMap()
    ssm.EntityData.Leafs = types.NewOrderedMap()
    ssm.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", ssm.Disable})
    ssm.EntityData.Leafs.Append("range", types.YLeaf{"Range", ssm.Range})

    ssm.EntityData.YListKeys = []string {}

    return &(ssm.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Injects
// Inject Explicit PIM RPF Vector Proxy's
type Pim_Vrfs_Vrf_Ipv4_Injects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inject Explicit PIM RPF Vector Proxy's. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_Injects_Inject.
    Inject []*Pim_Vrfs_Vrf_Ipv4_Injects_Inject
}

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetEntityData() *types.CommonEntityData {
    injects.EntityData.YFilter = injects.YFilter
    injects.EntityData.YangName = "injects"
    injects.EntityData.BundleName = "cisco_ios_xr"
    injects.EntityData.ParentYangName = "ipv4"
    injects.EntityData.SegmentPath = "injects"
    injects.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + injects.EntityData.SegmentPath
    injects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    injects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    injects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    injects.EntityData.Children = types.NewOrderedMap()
    injects.EntityData.Children.Append("inject", types.YChild{"Inject", nil})
    for i := range injects.Inject {
        injects.EntityData.Children.Append(types.GetSegmentPath(injects.Inject[i]), types.YChild{"Inject", injects.Inject[i]})
    }
    injects.EntityData.Leafs = types.NewOrderedMap()

    injects.EntityData.YListKeys = []string {}

    return &(injects.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Injects_Inject
// Inject Explicit PIM RPF Vector Proxy's
type Pim_Vrfs_Vrf_Ipv4_Injects_Inject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Source Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Masklen. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}

    // RPF Proxy Address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RpfProxyAddress []interface{}
}

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetEntityData() *types.CommonEntityData {
    inject.EntityData.YFilter = inject.YFilter
    inject.EntityData.YangName = "inject"
    inject.EntityData.BundleName = "cisco_ios_xr"
    inject.EntityData.ParentYangName = "injects"
    inject.EntityData.SegmentPath = "inject" + types.AddKeyToken(inject.SourceAddress, "source-address") + types.AddKeyToken(inject.PrefixLength, "prefix-length")
    inject.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/injects/" + inject.EntityData.SegmentPath
    inject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inject.EntityData.Children = types.NewOrderedMap()
    inject.EntityData.Leafs = types.NewOrderedMap()
    inject.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", inject.SourceAddress})
    inject.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", inject.PrefixLength})
    inject.EntityData.Leafs.Append("rpf-proxy-address", types.YLeaf{"RpfProxyAddress", inject.RpfProxyAddress})

    inject.EntityData.YListKeys = []string {"SourceAddress", "PrefixLength"}

    return &(inject.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses
// Configure Bidirectional PIM Rendezvous Point
type Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress.
    BidirRpAddress []*Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetEntityData() *types.CommonEntityData {
    bidirRpAddresses.EntityData.YFilter = bidirRpAddresses.YFilter
    bidirRpAddresses.EntityData.YangName = "bidir-rp-addresses"
    bidirRpAddresses.EntityData.BundleName = "cisco_ios_xr"
    bidirRpAddresses.EntityData.ParentYangName = "ipv4"
    bidirRpAddresses.EntityData.SegmentPath = "bidir-rp-addresses"
    bidirRpAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + bidirRpAddresses.EntityData.SegmentPath
    bidirRpAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bidirRpAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bidirRpAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bidirRpAddresses.EntityData.Children = types.NewOrderedMap()
    bidirRpAddresses.EntityData.Children.Append("bidir-rp-address", types.YChild{"BidirRpAddress", nil})
    for i := range bidirRpAddresses.BidirRpAddress {
        bidirRpAddresses.EntityData.Children.Append(types.GetSegmentPath(bidirRpAddresses.BidirRpAddress[i]), types.YChild{"BidirRpAddress", bidirRpAddresses.BidirRpAddress[i]})
    }
    bidirRpAddresses.EntityData.Leafs = types.NewOrderedMap()

    bidirRpAddresses.EntityData.YListKeys = []string {}

    return &(bidirRpAddresses.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress
// Address of the Rendezvous Point
type Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetEntityData() *types.CommonEntityData {
    bidirRpAddress.EntityData.YFilter = bidirRpAddress.YFilter
    bidirRpAddress.EntityData.YangName = "bidir-rp-address"
    bidirRpAddress.EntityData.BundleName = "cisco_ios_xr"
    bidirRpAddress.EntityData.ParentYangName = "bidir-rp-addresses"
    bidirRpAddress.EntityData.SegmentPath = "bidir-rp-address" + types.AddKeyToken(bidirRpAddress.RpAddress, "rp-address")
    bidirRpAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/bidir-rp-addresses/" + bidirRpAddress.EntityData.SegmentPath
    bidirRpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bidirRpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bidirRpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bidirRpAddress.EntityData.Children = types.NewOrderedMap()
    bidirRpAddress.EntityData.Leafs = types.NewOrderedMap()
    bidirRpAddress.EntityData.Leafs.Append("rp-address", types.YLeaf{"RpAddress", bidirRpAddress.RpAddress})
    bidirRpAddress.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", bidirRpAddress.AccessListName})
    bidirRpAddress.EntityData.Leafs.Append("auto-rp-override", types.YLeaf{"AutoRpOverride", bidirRpAddress.AutoRpOverride})

    bidirRpAddress.EntityData.YListKeys = []string {"RpAddress"}

    return &(bidirRpAddress.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Bsr
// PIM BSR configuration
type Pim_Vrfs_Vrf_Ipv4_Bsr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PIM Candidate BSR configuration.
    CandidateBsr Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr

    // PIM RP configuration.
    CandidateRps Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps
}

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetEntityData() *types.CommonEntityData {
    bsr.EntityData.YFilter = bsr.YFilter
    bsr.EntityData.YangName = "bsr"
    bsr.EntityData.BundleName = "cisco_ios_xr"
    bsr.EntityData.ParentYangName = "ipv4"
    bsr.EntityData.SegmentPath = "bsr"
    bsr.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + bsr.EntityData.SegmentPath
    bsr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsr.EntityData.Children = types.NewOrderedMap()
    bsr.EntityData.Children.Append("candidate-bsr", types.YChild{"CandidateBsr", &bsr.CandidateBsr})
    bsr.EntityData.Children.Append("candidate-rps", types.YChild{"CandidateRps", &bsr.CandidateRps})
    bsr.EntityData.Leafs = types.NewOrderedMap()

    bsr.EntityData.YListKeys = []string {}

    return &(bsr.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr
// PIM Candidate BSR configuration
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // BSR Address configured. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    Address interface{}

    // Hash Mask Length for this candidate BSR. The type is interface{} with
    // range: 0..32. The default value is 30.
    PrefixLength interface{}

    // Priority of the Candidate BSR. The type is interface{} with range: 1..255.
    // The default value is 1.
    Priority interface{}
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetEntityData() *types.CommonEntityData {
    candidateBsr.EntityData.YFilter = candidateBsr.YFilter
    candidateBsr.EntityData.YangName = "candidate-bsr"
    candidateBsr.EntityData.BundleName = "cisco_ios_xr"
    candidateBsr.EntityData.ParentYangName = "bsr"
    candidateBsr.EntityData.SegmentPath = "candidate-bsr"
    candidateBsr.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/bsr/" + candidateBsr.EntityData.SegmentPath
    candidateBsr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateBsr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateBsr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateBsr.EntityData.Children = types.NewOrderedMap()
    candidateBsr.EntityData.Leafs = types.NewOrderedMap()
    candidateBsr.EntityData.Leafs.Append("address", types.YLeaf{"Address", candidateBsr.Address})
    candidateBsr.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", candidateBsr.PrefixLength})
    candidateBsr.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", candidateBsr.Priority})

    candidateBsr.EntityData.YListKeys = []string {}

    return &(candidateBsr.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps
// PIM RP configuration
type Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of PIM SM BSR Candidate-RP. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp.
    CandidateRp []*Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetEntityData() *types.CommonEntityData {
    candidateRps.EntityData.YFilter = candidateRps.YFilter
    candidateRps.EntityData.YangName = "candidate-rps"
    candidateRps.EntityData.BundleName = "cisco_ios_xr"
    candidateRps.EntityData.ParentYangName = "bsr"
    candidateRps.EntityData.SegmentPath = "candidate-rps"
    candidateRps.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/bsr/" + candidateRps.EntityData.SegmentPath
    candidateRps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateRps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateRps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateRps.EntityData.Children = types.NewOrderedMap()
    candidateRps.EntityData.Children.Append("candidate-rp", types.YChild{"CandidateRp", nil})
    for i := range candidateRps.CandidateRp {
        candidateRps.EntityData.Children.Append(types.GetSegmentPath(candidateRps.CandidateRp[i]), types.YChild{"CandidateRp", candidateRps.CandidateRp[i]})
    }
    candidateRps.EntityData.Leafs = types.NewOrderedMap()

    candidateRps.EntityData.YListKeys = []string {}

    return &(candidateRps.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp
// Address of PIM SM BSR Candidate-RP
type Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address of Candidate-RP. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. SM or Bidir. The type is PimProtocolMode.
    Mode interface{}

    // Access-list specifying the group range for the Candidate-RP. The type is
    // string with length: 1..64.
    GroupList interface{}

    // Priority of the CRP. The type is interface{} with range: 1..255. The
    // default value is 192.
    Priority interface{}

    // Advertisement interval. The type is interface{} with range: 30..600. The
    // default value is 60.
    Interval interface{}
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetEntityData() *types.CommonEntityData {
    candidateRp.EntityData.YFilter = candidateRp.YFilter
    candidateRp.EntityData.YangName = "candidate-rp"
    candidateRp.EntityData.BundleName = "cisco_ios_xr"
    candidateRp.EntityData.ParentYangName = "candidate-rps"
    candidateRp.EntityData.SegmentPath = "candidate-rp" + types.AddKeyToken(candidateRp.Address, "address") + types.AddKeyToken(candidateRp.Mode, "mode")
    candidateRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/bsr/candidate-rps/" + candidateRp.EntityData.SegmentPath
    candidateRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateRp.EntityData.Children = types.NewOrderedMap()
    candidateRp.EntityData.Leafs = types.NewOrderedMap()
    candidateRp.EntityData.Leafs.Append("address", types.YLeaf{"Address", candidateRp.Address})
    candidateRp.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", candidateRp.Mode})
    candidateRp.EntityData.Leafs.Append("group-list", types.YLeaf{"GroupList", candidateRp.GroupList})
    candidateRp.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", candidateRp.Priority})
    candidateRp.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", candidateRp.Interval})

    candidateRp.EntityData.YListKeys = []string {"Address", "Mode"}

    return &(candidateRp.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Mofrr
// Multicast Only Fast Re-Route
type Pim_Vrfs_Vrf_Ipv4_Mofrr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access-list specifying SG that should do RIB MOFRR. The type is string with
    // length: 1..64.
    Rib interface{}

    // Non-revertive Multicast Only Fast Re-Route. The type is interface{}.
    NonRevertive interface{}

    // Enable Multicast Only FRR. The type is interface{}.
    Enable interface{}

    // Access-list specifying SG that should do FLOW MOFRR. The type is string
    // with length: 1..64.
    Flow interface{}

    // Clone multicast joins.
    CloneJoins Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins

    // Clone multicast traffic.
    CloneSources Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources
}

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetEntityData() *types.CommonEntityData {
    mofrr.EntityData.YFilter = mofrr.YFilter
    mofrr.EntityData.YangName = "mofrr"
    mofrr.EntityData.BundleName = "cisco_ios_xr"
    mofrr.EntityData.ParentYangName = "ipv4"
    mofrr.EntityData.SegmentPath = "mofrr"
    mofrr.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + mofrr.EntityData.SegmentPath
    mofrr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mofrr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mofrr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mofrr.EntityData.Children = types.NewOrderedMap()
    mofrr.EntityData.Children.Append("clone-joins", types.YChild{"CloneJoins", &mofrr.CloneJoins})
    mofrr.EntityData.Children.Append("clone-sources", types.YChild{"CloneSources", &mofrr.CloneSources})
    mofrr.EntityData.Leafs = types.NewOrderedMap()
    mofrr.EntityData.Leafs.Append("rib", types.YLeaf{"Rib", mofrr.Rib})
    mofrr.EntityData.Leafs.Append("non-revertive", types.YLeaf{"NonRevertive", mofrr.NonRevertive})
    mofrr.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", mofrr.Enable})
    mofrr.EntityData.Leafs.Append("flow", types.YLeaf{"Flow", mofrr.Flow})

    mofrr.EntityData.YListKeys = []string {}

    return &(mofrr.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins
// Clone multicast joins
type Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clone S,G joins as S1,G joins and S2,G joins. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin.
    CloneJoin []*Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin
}

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetEntityData() *types.CommonEntityData {
    cloneJoins.EntityData.YFilter = cloneJoins.YFilter
    cloneJoins.EntityData.YangName = "clone-joins"
    cloneJoins.EntityData.BundleName = "cisco_ios_xr"
    cloneJoins.EntityData.ParentYangName = "mofrr"
    cloneJoins.EntityData.SegmentPath = "clone-joins"
    cloneJoins.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/mofrr/" + cloneJoins.EntityData.SegmentPath
    cloneJoins.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cloneJoins.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cloneJoins.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cloneJoins.EntityData.Children = types.NewOrderedMap()
    cloneJoins.EntityData.Children.Append("clone-join", types.YChild{"CloneJoin", nil})
    for i := range cloneJoins.CloneJoin {
        cloneJoins.EntityData.Children.Append(types.GetSegmentPath(cloneJoins.CloneJoin[i]), types.YChild{"CloneJoin", cloneJoins.CloneJoin[i]})
    }
    cloneJoins.EntityData.Leafs = types.NewOrderedMap()

    cloneJoins.EntityData.YListKeys = []string {}

    return &(cloneJoins.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin
// Clone S,G joins as S1,G joins and S2,G joins
type Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Original source address (S). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Source interface{}

    // This attribute is a key. Primary cloned address (S1). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Primary interface{}

    // This attribute is a key. Backup cloned address (S2). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Backup interface{}

    // This attribute is a key. Mask length. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}
}

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetEntityData() *types.CommonEntityData {
    cloneJoin.EntityData.YFilter = cloneJoin.YFilter
    cloneJoin.EntityData.YangName = "clone-join"
    cloneJoin.EntityData.BundleName = "cisco_ios_xr"
    cloneJoin.EntityData.ParentYangName = "clone-joins"
    cloneJoin.EntityData.SegmentPath = "clone-join" + types.AddKeyToken(cloneJoin.Source, "source") + types.AddKeyToken(cloneJoin.Primary, "primary") + types.AddKeyToken(cloneJoin.Backup, "backup") + types.AddKeyToken(cloneJoin.PrefixLength, "prefix-length")
    cloneJoin.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/mofrr/clone-joins/" + cloneJoin.EntityData.SegmentPath
    cloneJoin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cloneJoin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cloneJoin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cloneJoin.EntityData.Children = types.NewOrderedMap()
    cloneJoin.EntityData.Leafs = types.NewOrderedMap()
    cloneJoin.EntityData.Leafs.Append("source", types.YLeaf{"Source", cloneJoin.Source})
    cloneJoin.EntityData.Leafs.Append("primary", types.YLeaf{"Primary", cloneJoin.Primary})
    cloneJoin.EntityData.Leafs.Append("backup", types.YLeaf{"Backup", cloneJoin.Backup})
    cloneJoin.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", cloneJoin.PrefixLength})

    cloneJoin.EntityData.YListKeys = []string {"Source", "Primary", "Backup", "PrefixLength"}

    return &(cloneJoin.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources
// Clone multicast traffic
type Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clone S,G traffic as S1,G traffic and S2,G traffic. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource.
    CloneSource []*Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource
}

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetEntityData() *types.CommonEntityData {
    cloneSources.EntityData.YFilter = cloneSources.YFilter
    cloneSources.EntityData.YangName = "clone-sources"
    cloneSources.EntityData.BundleName = "cisco_ios_xr"
    cloneSources.EntityData.ParentYangName = "mofrr"
    cloneSources.EntityData.SegmentPath = "clone-sources"
    cloneSources.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/mofrr/" + cloneSources.EntityData.SegmentPath
    cloneSources.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cloneSources.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cloneSources.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cloneSources.EntityData.Children = types.NewOrderedMap()
    cloneSources.EntityData.Children.Append("clone-source", types.YChild{"CloneSource", nil})
    for i := range cloneSources.CloneSource {
        cloneSources.EntityData.Children.Append(types.GetSegmentPath(cloneSources.CloneSource[i]), types.YChild{"CloneSource", cloneSources.CloneSource[i]})
    }
    cloneSources.EntityData.Leafs = types.NewOrderedMap()

    cloneSources.EntityData.YListKeys = []string {}

    return &(cloneSources.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource
// Clone S,G traffic as S1,G traffic and S2,G
// traffic
type Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Original source address (S). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Source interface{}

    // This attribute is a key. Primary cloned address (S1). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Primary interface{}

    // This attribute is a key. Backup cloned address (S2). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Backup interface{}

    // This attribute is a key. Mask length. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}
}

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetEntityData() *types.CommonEntityData {
    cloneSource.EntityData.YFilter = cloneSource.YFilter
    cloneSource.EntityData.YangName = "clone-source"
    cloneSource.EntityData.BundleName = "cisco_ios_xr"
    cloneSource.EntityData.ParentYangName = "clone-sources"
    cloneSource.EntityData.SegmentPath = "clone-source" + types.AddKeyToken(cloneSource.Source, "source") + types.AddKeyToken(cloneSource.Primary, "primary") + types.AddKeyToken(cloneSource.Backup, "backup") + types.AddKeyToken(cloneSource.PrefixLength, "prefix-length")
    cloneSource.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/mofrr/clone-sources/" + cloneSource.EntityData.SegmentPath
    cloneSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cloneSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cloneSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cloneSource.EntityData.Children = types.NewOrderedMap()
    cloneSource.EntityData.Leafs = types.NewOrderedMap()
    cloneSource.EntityData.Leafs.Append("source", types.YLeaf{"Source", cloneSource.Source})
    cloneSource.EntityData.Leafs.Append("primary", types.YLeaf{"Primary", cloneSource.Primary})
    cloneSource.EntityData.Leafs.Append("backup", types.YLeaf{"Backup", cloneSource.Backup})
    cloneSource.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", cloneSource.PrefixLength})

    cloneSource.EntityData.YListKeys = []string {"Source", "Primary", "Backup", "PrefixLength"}

    return &(cloneSource.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Paths
// Inject PIM RPF Vector Proxy's
type Pim_Vrfs_Vrf_Ipv4_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inject PIM RPF Vector Proxy's. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_Paths_Path.
    Path []*Pim_Vrfs_Vrf_Ipv4_Paths_Path
}

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "ipv4"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + paths.EntityData.SegmentPath
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = types.NewOrderedMap()
    paths.EntityData.Children.Append("path", types.YChild{"Path", nil})
    for i := range paths.Path {
        paths.EntityData.Children.Append(types.GetSegmentPath(paths.Path[i]), types.YChild{"Path", paths.Path[i]})
    }
    paths.EntityData.Leafs = types.NewOrderedMap()

    paths.EntityData.YListKeys = []string {}

    return &(paths.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Paths_Path
// Inject PIM RPF Vector Proxy's
type Pim_Vrfs_Vrf_Ipv4_Paths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Source Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Masklen. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}

    // RPF Proxy Address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RpfProxyAddress []interface{}
}

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "paths"
    path.EntityData.SegmentPath = "path" + types.AddKeyToken(path.SourceAddress, "source-address") + types.AddKeyToken(path.PrefixLength, "prefix-length")
    path.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/paths/" + path.EntityData.SegmentPath
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", path.SourceAddress})
    path.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", path.PrefixLength})
    path.EntityData.Leafs.Append("rpf-proxy-address", types.YLeaf{"RpfProxyAddress", path.RpfProxyAddress})

    path.EntityData.YListKeys = []string {"SourceAddress", "PrefixLength"}

    return &(path.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_AllowRp
// Enable allow-rp filtering for SM joins
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_AllowRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Access-list specifiying applicable RPs. The type is string with length:
    // 1..64.
    RpListName interface{}

    // Access-list specifiying applicable groups. The type is string with length:
    // 1..64.
    GroupListName interface{}
}

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetEntityData() *types.CommonEntityData {
    allowRp.EntityData.YFilter = allowRp.YFilter
    allowRp.EntityData.YangName = "allow-rp"
    allowRp.EntityData.BundleName = "cisco_ios_xr"
    allowRp.EntityData.ParentYangName = "ipv4"
    allowRp.EntityData.SegmentPath = "allow-rp"
    allowRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + allowRp.EntityData.SegmentPath
    allowRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowRp.EntityData.Children = types.NewOrderedMap()
    allowRp.EntityData.Leafs = types.NewOrderedMap()
    allowRp.EntityData.Leafs.Append("rp-list-name", types.YLeaf{"RpListName", allowRp.RpListName})
    allowRp.EntityData.Leafs.Append("group-list-name", types.YLeaf{"GroupListName", allowRp.GroupListName})

    allowRp.EntityData.YListKeys = []string {}

    return &(allowRp.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Convergence
// Configure convergence parameters
type Pim_Vrfs_Vrf_Ipv4_Convergence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampen first join if RPF path is through one of the downstream neighbor.
    // The type is interface{} with range: 0..15. Units are second.
    RpfConflictJoinDelay interface{}

    // Delay prunes if route join state transitions to not-joined on link down.
    // The type is interface{} with range: 0..60. Units are second.
    LinkDownPruneDelay interface{}
}

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetEntityData() *types.CommonEntityData {
    convergence.EntityData.YFilter = convergence.YFilter
    convergence.EntityData.YangName = "convergence"
    convergence.EntityData.BundleName = "cisco_ios_xr"
    convergence.EntityData.ParentYangName = "ipv4"
    convergence.EntityData.SegmentPath = "convergence"
    convergence.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + convergence.EntityData.SegmentPath
    convergence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    convergence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    convergence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    convergence.EntityData.Children = types.NewOrderedMap()
    convergence.EntityData.Leafs = types.NewOrderedMap()
    convergence.EntityData.Leafs.Append("rpf-conflict-join-delay", types.YLeaf{"RpfConflictJoinDelay", convergence.RpfConflictJoinDelay})
    convergence.EntityData.Leafs.Append("link-down-prune-delay", types.YLeaf{"LinkDownPruneDelay", convergence.LinkDownPruneDelay})

    convergence.EntityData.YListKeys = []string {}

    return &(convergence.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Interfaces
// Interface-level Configuration
type Pim_Vrfs_Vrf_Ipv4_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface.
    Interface []*Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface
}

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ipv4"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/" + interfaces.EntityData.SegmentPath
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

// Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface
// The name of the interface
type Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of interface. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Enter PIM Interface processing. The type is interface{}.
    Enable interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // BSR Border configuration for Interface. The type is bool.
    BsrBorder interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Enable PIM processing on the interface. The type is bool.
    InterfaceEnable interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}

    // Maximum number of allowed routes for this interface.
    MaximumRoutes Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes

    // BFD configuration.
    Bfd Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd
}

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("maximum-routes", types.YChild{"MaximumRoutes", &self.MaximumRoutes})
    self.EntityData.Children.Append("bfd", types.YChild{"Bfd", &self.Bfd})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})
    self.EntityData.Leafs.Append("neighbor-filter", types.YLeaf{"NeighborFilter", self.NeighborFilter})
    self.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", self.HelloInterval})
    self.EntityData.Leafs.Append("bsr-border", types.YLeaf{"BsrBorder", self.BsrBorder})
    self.EntityData.Leafs.Append("propagation-delay", types.YLeaf{"PropagationDelay", self.PropagationDelay})
    self.EntityData.Leafs.Append("dr-priority", types.YLeaf{"DrPriority", self.DrPriority})
    self.EntityData.Leafs.Append("join-prune-mtu", types.YLeaf{"JoinPruneMtu", self.JoinPruneMtu})
    self.EntityData.Leafs.Append("interface-enable", types.YLeaf{"InterfaceEnable", self.InterfaceEnable})
    self.EntityData.Leafs.Append("jp-interval", types.YLeaf{"JpInterval", self.JpInterval})
    self.EntityData.Leafs.Append("override-interval", types.YLeaf{"OverrideInterval", self.OverrideInterval})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes
// Maximum number of allowed routes for this
// interface
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of routes for this interface. The type is interface{} with
    // range: 1..1100000. This attribute is mandatory.
    Maximum interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetEntityData() *types.CommonEntityData {
    maximumRoutes.EntityData.YFilter = maximumRoutes.YFilter
    maximumRoutes.EntityData.YangName = "maximum-routes"
    maximumRoutes.EntityData.BundleName = "cisco_ios_xr"
    maximumRoutes.EntityData.ParentYangName = "interface"
    maximumRoutes.EntityData.SegmentPath = "maximum-routes"
    maximumRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/interfaces/interface/" + maximumRoutes.EntityData.SegmentPath
    maximumRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumRoutes.EntityData.Children = types.NewOrderedMap()
    maximumRoutes.EntityData.Leafs = types.NewOrderedMap()
    maximumRoutes.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", maximumRoutes.Maximum})
    maximumRoutes.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", maximumRoutes.WarningThreshold})
    maximumRoutes.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", maximumRoutes.AccessListName})

    maximumRoutes.EntityData.YListKeys = []string {}

    return &(maximumRoutes.EntityData)
}

// Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd
// BFD configuration
type Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detection multiplier for BFD sessions created by PIM. The type is
    // interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval for BFD sessions created by PIM. The type is interface{}
    // with range: 3..30000. Units are millisecond.
    Interval interface{}

    // TRUE to enable BFD. FALSE to disable and to prevent inheritance from a
    // parent. The type is bool.
    Enable interface{}
}

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "interface"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv4/interfaces/interface/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bfd.Enable})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6
// IPV6 commands
type Pim_Vrfs_Vrf_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable PIM neighbor checking when receiving PIM messages. The type is
    // interface{}.
    NeighborCheckOnReceive interface{}

    // Generate registers compatible with older IOS versions. The type is
    // interface{}.
    OldRegisterChecksum interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Configure threshold of infinity for switching to SPT on last-hop. The type
    // is string.
    SptThresholdInfinity interface{}

    // PIM neighbor state change logging is turned on if configured. The type is
    // interface{}.
    LogNeighborChanges interface{}

    // Source address to use for register messages. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    RegisterSource interface{}

    // Access-list which specifies unauthorized sources. The type is string with
    // length: 1..64.
    AcceptRegister interface{}

    // Set Embedded RP processing support. The type is interface{}.
    EmbeddedRpDisable interface{}

    // Suppress prunes triggered as a result of RPF changes. The type is
    // interface{}.
    SuppressRpfPrunes interface{}

    // Allow SSM ranges to be overridden by more specific ranges. The type is
    // interface{}.
    SsmAllowOverride interface{}

    // Enable equal-cost multipath routing. The type is PimMultipath.
    Multipath interface{}

    // Configure static RP deny range. The type is string with length: 1..64.
    RpStaticDeny interface{}

    // Suppress data registers after initial state setup. The type is interface{}.
    SuppressDataRegisters interface{}

    // Enable PIM neighbor checking when sending join-prunes. The type is
    // interface{}.
    NeighborCheckOnSend interface{}

    // Configure Sparse-Mode Rendezvous Point.
    SparseModeRpAddresses Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses

    // Inheritable defaults.
    InheritableDefaults Pim_Vrfs_Vrf_Ipv6_InheritableDefaults

    // Configure RPF options.
    Rpf Pim_Vrfs_Vrf_Ipv6_Rpf

    // Configure PIM State Limits.
    Maximum Pim_Vrfs_Vrf_Ipv6_Maximum

    // Configure expiry timer for S,G routes.
    SgExpiryTimer Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer

    // Enable PIM RPF Vector Proxy's.
    RpfVectorEnable Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable

    // Configure IP Multicast SSM.
    Ssm Pim_Vrfs_Vrf_Ipv6_Ssm

    // Configure Bidirectional PIM Rendezvous Point.
    BidirRpAddresses Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses

    // PIM BSR configuration.
    Bsr Pim_Vrfs_Vrf_Ipv6_Bsr

    // Enable allow-rp filtering for SM joins.
    AllowRp Pim_Vrfs_Vrf_Ipv6_AllowRp

    // Set Embedded RP processing support.
    EmbeddedRpAddresses Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses

    // Configure convergence parameters.
    Convergence Pim_Vrfs_Vrf_Ipv6_Convergence

    // Interface-level Configuration.
    Interfaces Pim_Vrfs_Vrf_Ipv6_Interfaces
}

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "vrf"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/" + ipv6.EntityData.SegmentPath
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("sparse-mode-rp-addresses", types.YChild{"SparseModeRpAddresses", &ipv6.SparseModeRpAddresses})
    ipv6.EntityData.Children.Append("inheritable-defaults", types.YChild{"InheritableDefaults", &ipv6.InheritableDefaults})
    ipv6.EntityData.Children.Append("rpf", types.YChild{"Rpf", &ipv6.Rpf})
    ipv6.EntityData.Children.Append("maximum", types.YChild{"Maximum", &ipv6.Maximum})
    ipv6.EntityData.Children.Append("sg-expiry-timer", types.YChild{"SgExpiryTimer", &ipv6.SgExpiryTimer})
    ipv6.EntityData.Children.Append("rpf-vector-enable", types.YChild{"RpfVectorEnable", &ipv6.RpfVectorEnable})
    ipv6.EntityData.Children.Append("ssm", types.YChild{"Ssm", &ipv6.Ssm})
    ipv6.EntityData.Children.Append("bidir-rp-addresses", types.YChild{"BidirRpAddresses", &ipv6.BidirRpAddresses})
    ipv6.EntityData.Children.Append("bsr", types.YChild{"Bsr", &ipv6.Bsr})
    ipv6.EntityData.Children.Append("allow-rp", types.YChild{"AllowRp", &ipv6.AllowRp})
    ipv6.EntityData.Children.Append("embedded-rp-addresses", types.YChild{"EmbeddedRpAddresses", &ipv6.EmbeddedRpAddresses})
    ipv6.EntityData.Children.Append("convergence", types.YChild{"Convergence", &ipv6.Convergence})
    ipv6.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &ipv6.Interfaces})
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("neighbor-check-on-receive", types.YLeaf{"NeighborCheckOnReceive", ipv6.NeighborCheckOnReceive})
    ipv6.EntityData.Leafs.Append("old-register-checksum", types.YLeaf{"OldRegisterChecksum", ipv6.OldRegisterChecksum})
    ipv6.EntityData.Leafs.Append("neighbor-filter", types.YLeaf{"NeighborFilter", ipv6.NeighborFilter})
    ipv6.EntityData.Leafs.Append("spt-threshold-infinity", types.YLeaf{"SptThresholdInfinity", ipv6.SptThresholdInfinity})
    ipv6.EntityData.Leafs.Append("log-neighbor-changes", types.YLeaf{"LogNeighborChanges", ipv6.LogNeighborChanges})
    ipv6.EntityData.Leafs.Append("register-source", types.YLeaf{"RegisterSource", ipv6.RegisterSource})
    ipv6.EntityData.Leafs.Append("accept-register", types.YLeaf{"AcceptRegister", ipv6.AcceptRegister})
    ipv6.EntityData.Leafs.Append("embedded-rp-disable", types.YLeaf{"EmbeddedRpDisable", ipv6.EmbeddedRpDisable})
    ipv6.EntityData.Leafs.Append("suppress-rpf-prunes", types.YLeaf{"SuppressRpfPrunes", ipv6.SuppressRpfPrunes})
    ipv6.EntityData.Leafs.Append("ssm-allow-override", types.YLeaf{"SsmAllowOverride", ipv6.SsmAllowOverride})
    ipv6.EntityData.Leafs.Append("multipath", types.YLeaf{"Multipath", ipv6.Multipath})
    ipv6.EntityData.Leafs.Append("rp-static-deny", types.YLeaf{"RpStaticDeny", ipv6.RpStaticDeny})
    ipv6.EntityData.Leafs.Append("suppress-data-registers", types.YLeaf{"SuppressDataRegisters", ipv6.SuppressDataRegisters})
    ipv6.EntityData.Leafs.Append("neighbor-check-on-send", types.YLeaf{"NeighborCheckOnSend", ipv6.NeighborCheckOnSend})

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses
// Configure Sparse-Mode Rendezvous Point
type Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress.
    SparseModeRpAddress []*Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetEntityData() *types.CommonEntityData {
    sparseModeRpAddresses.EntityData.YFilter = sparseModeRpAddresses.YFilter
    sparseModeRpAddresses.EntityData.YangName = "sparse-mode-rp-addresses"
    sparseModeRpAddresses.EntityData.BundleName = "cisco_ios_xr"
    sparseModeRpAddresses.EntityData.ParentYangName = "ipv6"
    sparseModeRpAddresses.EntityData.SegmentPath = "sparse-mode-rp-addresses"
    sparseModeRpAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/" + sparseModeRpAddresses.EntityData.SegmentPath
    sparseModeRpAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sparseModeRpAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sparseModeRpAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sparseModeRpAddresses.EntityData.Children = types.NewOrderedMap()
    sparseModeRpAddresses.EntityData.Children.Append("sparse-mode-rp-address", types.YChild{"SparseModeRpAddress", nil})
    for i := range sparseModeRpAddresses.SparseModeRpAddress {
        sparseModeRpAddresses.EntityData.Children.Append(types.GetSegmentPath(sparseModeRpAddresses.SparseModeRpAddress[i]), types.YChild{"SparseModeRpAddress", sparseModeRpAddresses.SparseModeRpAddress[i]})
    }
    sparseModeRpAddresses.EntityData.Leafs = types.NewOrderedMap()

    sparseModeRpAddresses.EntityData.YListKeys = []string {}

    return &(sparseModeRpAddresses.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress
// Address of the Rendezvous Point
type Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a  given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetEntityData() *types.CommonEntityData {
    sparseModeRpAddress.EntityData.YFilter = sparseModeRpAddress.YFilter
    sparseModeRpAddress.EntityData.YangName = "sparse-mode-rp-address"
    sparseModeRpAddress.EntityData.BundleName = "cisco_ios_xr"
    sparseModeRpAddress.EntityData.ParentYangName = "sparse-mode-rp-addresses"
    sparseModeRpAddress.EntityData.SegmentPath = "sparse-mode-rp-address" + types.AddKeyToken(sparseModeRpAddress.RpAddress, "rp-address")
    sparseModeRpAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/sparse-mode-rp-addresses/" + sparseModeRpAddress.EntityData.SegmentPath
    sparseModeRpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sparseModeRpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sparseModeRpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sparseModeRpAddress.EntityData.Children = types.NewOrderedMap()
    sparseModeRpAddress.EntityData.Leafs = types.NewOrderedMap()
    sparseModeRpAddress.EntityData.Leafs.Append("rp-address", types.YLeaf{"RpAddress", sparseModeRpAddress.RpAddress})
    sparseModeRpAddress.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", sparseModeRpAddress.AccessListName})
    sparseModeRpAddress.EntityData.Leafs.Append("auto-rp-override", types.YLeaf{"AutoRpOverride", sparseModeRpAddress.AutoRpOverride})

    sparseModeRpAddress.EntityData.YListKeys = []string {"RpAddress"}

    return &(sparseModeRpAddress.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_InheritableDefaults
// Inheritable defaults
type Pim_Vrfs_Vrf_Ipv6_InheritableDefaults struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Convergency timeout in seconds. The type is interface{} with range:
    // 1800..2400. Units are second.
    ConvergenceTimeout interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetEntityData() *types.CommonEntityData {
    inheritableDefaults.EntityData.YFilter = inheritableDefaults.YFilter
    inheritableDefaults.EntityData.YangName = "inheritable-defaults"
    inheritableDefaults.EntityData.BundleName = "cisco_ios_xr"
    inheritableDefaults.EntityData.ParentYangName = "ipv6"
    inheritableDefaults.EntityData.SegmentPath = "inheritable-defaults"
    inheritableDefaults.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/" + inheritableDefaults.EntityData.SegmentPath
    inheritableDefaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritableDefaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritableDefaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritableDefaults.EntityData.Children = types.NewOrderedMap()
    inheritableDefaults.EntityData.Leafs = types.NewOrderedMap()
    inheritableDefaults.EntityData.Leafs.Append("convergence-timeout", types.YLeaf{"ConvergenceTimeout", inheritableDefaults.ConvergenceTimeout})
    inheritableDefaults.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", inheritableDefaults.HelloInterval})
    inheritableDefaults.EntityData.Leafs.Append("propagation-delay", types.YLeaf{"PropagationDelay", inheritableDefaults.PropagationDelay})
    inheritableDefaults.EntityData.Leafs.Append("dr-priority", types.YLeaf{"DrPriority", inheritableDefaults.DrPriority})
    inheritableDefaults.EntityData.Leafs.Append("join-prune-mtu", types.YLeaf{"JoinPruneMtu", inheritableDefaults.JoinPruneMtu})
    inheritableDefaults.EntityData.Leafs.Append("jp-interval", types.YLeaf{"JpInterval", inheritableDefaults.JpInterval})
    inheritableDefaults.EntityData.Leafs.Append("override-interval", types.YLeaf{"OverrideInterval", inheritableDefaults.OverrideInterval})

    inheritableDefaults.EntityData.YListKeys = []string {}

    return &(inheritableDefaults.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Rpf
// Configure RPF options
type Pim_Vrfs_Vrf_Ipv6_Rpf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route policy to select RPF topology. The type is string with length: 1..64.
    RoutePolicy interface{}
}

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetEntityData() *types.CommonEntityData {
    rpf.EntityData.YFilter = rpf.YFilter
    rpf.EntityData.YangName = "rpf"
    rpf.EntityData.BundleName = "cisco_ios_xr"
    rpf.EntityData.ParentYangName = "ipv6"
    rpf.EntityData.SegmentPath = "rpf"
    rpf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/" + rpf.EntityData.SegmentPath
    rpf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpf.EntityData.Children = types.NewOrderedMap()
    rpf.EntityData.Leafs = types.NewOrderedMap()
    rpf.EntityData.Leafs.Append("route-policy", types.YLeaf{"RoutePolicy", rpf.RoutePolicy})

    rpf.EntityData.YListKeys = []string {}

    return &(rpf.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Maximum
// Configure PIM State Limits
type Pim_Vrfs_Vrf_Ipv6_Maximum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Override default maximum for number of group mappings from autorp mapping
    // agent.
    GroupMappingsAutoRp Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp

    // Override default maximum and threshold for number of group mappings from
    // BSR.
    BsrGroupMappings Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings

    // Override default maximum for number of sparse-mode source registers.
    RegisterStates Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates

    // Override default maximum for number of route-interfaces.
    RouteInterfaces Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces

    // Override default maximum and threshold for BSR C-RP cache setting.
    BsrCandidateRpCache Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache

    // Override default maximum for number of routes.
    Routes Pim_Vrfs_Vrf_Ipv6_Maximum_Routes
}

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetEntityData() *types.CommonEntityData {
    maximum.EntityData.YFilter = maximum.YFilter
    maximum.EntityData.YangName = "maximum"
    maximum.EntityData.BundleName = "cisco_ios_xr"
    maximum.EntityData.ParentYangName = "ipv6"
    maximum.EntityData.SegmentPath = "maximum"
    maximum.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/" + maximum.EntityData.SegmentPath
    maximum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximum.EntityData.Children = types.NewOrderedMap()
    maximum.EntityData.Children.Append("group-mappings-auto-rp", types.YChild{"GroupMappingsAutoRp", &maximum.GroupMappingsAutoRp})
    maximum.EntityData.Children.Append("bsr-group-mappings", types.YChild{"BsrGroupMappings", &maximum.BsrGroupMappings})
    maximum.EntityData.Children.Append("register-states", types.YChild{"RegisterStates", &maximum.RegisterStates})
    maximum.EntityData.Children.Append("route-interfaces", types.YChild{"RouteInterfaces", &maximum.RouteInterfaces})
    maximum.EntityData.Children.Append("bsr-candidate-rp-cache", types.YChild{"BsrCandidateRpCache", &maximum.BsrCandidateRpCache})
    maximum.EntityData.Children.Append("routes", types.YChild{"Routes", &maximum.Routes})
    maximum.EntityData.Leafs = types.NewOrderedMap()

    maximum.EntityData.YListKeys = []string {}

    return &(maximum.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp
// Override default maximum for number of group
// mappings from autorp mapping agent
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM group mappings from autorp. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    MaximumGroupRangesAutoRp interface{}

    // Warning threshold number of PIM group mappings from autorp. The type is
    // interface{} with range: 1..10000. The default value is 450.
    ThresholdGroupRangesAutoRp interface{}
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetEntityData() *types.CommonEntityData {
    groupMappingsAutoRp.EntityData.YFilter = groupMappingsAutoRp.YFilter
    groupMappingsAutoRp.EntityData.YangName = "group-mappings-auto-rp"
    groupMappingsAutoRp.EntityData.BundleName = "cisco_ios_xr"
    groupMappingsAutoRp.EntityData.ParentYangName = "maximum"
    groupMappingsAutoRp.EntityData.SegmentPath = "group-mappings-auto-rp"
    groupMappingsAutoRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/maximum/" + groupMappingsAutoRp.EntityData.SegmentPath
    groupMappingsAutoRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupMappingsAutoRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupMappingsAutoRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupMappingsAutoRp.EntityData.Children = types.NewOrderedMap()
    groupMappingsAutoRp.EntityData.Leafs = types.NewOrderedMap()
    groupMappingsAutoRp.EntityData.Leafs.Append("maximum-group-ranges-auto-rp", types.YLeaf{"MaximumGroupRangesAutoRp", groupMappingsAutoRp.MaximumGroupRangesAutoRp})
    groupMappingsAutoRp.EntityData.Leafs.Append("threshold-group-ranges-auto-rp", types.YLeaf{"ThresholdGroupRangesAutoRp", groupMappingsAutoRp.ThresholdGroupRangesAutoRp})

    groupMappingsAutoRp.EntityData.YListKeys = []string {}

    return &(groupMappingsAutoRp.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings
// Override default maximum and threshold for
// number of group mappings from BSR
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM group mappings from BSR. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumGroupRanges interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 500.
    WarningThreshold interface{}
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetEntityData() *types.CommonEntityData {
    bsrGroupMappings.EntityData.YFilter = bsrGroupMappings.YFilter
    bsrGroupMappings.EntityData.YangName = "bsr-group-mappings"
    bsrGroupMappings.EntityData.BundleName = "cisco_ios_xr"
    bsrGroupMappings.EntityData.ParentYangName = "maximum"
    bsrGroupMappings.EntityData.SegmentPath = "bsr-group-mappings"
    bsrGroupMappings.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/maximum/" + bsrGroupMappings.EntityData.SegmentPath
    bsrGroupMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsrGroupMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsrGroupMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsrGroupMappings.EntityData.Children = types.NewOrderedMap()
    bsrGroupMappings.EntityData.Leafs = types.NewOrderedMap()
    bsrGroupMappings.EntityData.Leafs.Append("bsr-maximum-group-ranges", types.YLeaf{"BsrMaximumGroupRanges", bsrGroupMappings.BsrMaximumGroupRanges})
    bsrGroupMappings.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", bsrGroupMappings.WarningThreshold})

    bsrGroupMappings.EntityData.YListKeys = []string {}

    return &(bsrGroupMappings.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates
// Override default maximum for number of
// sparse-mode source registers
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM Sparse-Mode register states. The type is interface{}
    // with range: 0..75000. This attribute is mandatory.
    MaximumRegisterStates interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 0..75000. The default value is 20000.
    WarningThreshold interface{}
}

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetEntityData() *types.CommonEntityData {
    registerStates.EntityData.YFilter = registerStates.YFilter
    registerStates.EntityData.YangName = "register-states"
    registerStates.EntityData.BundleName = "cisco_ios_xr"
    registerStates.EntityData.ParentYangName = "maximum"
    registerStates.EntityData.SegmentPath = "register-states"
    registerStates.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/maximum/" + registerStates.EntityData.SegmentPath
    registerStates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registerStates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registerStates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registerStates.EntityData.Children = types.NewOrderedMap()
    registerStates.EntityData.Leafs = types.NewOrderedMap()
    registerStates.EntityData.Leafs.Append("maximum-register-states", types.YLeaf{"MaximumRegisterStates", registerStates.MaximumRegisterStates})
    registerStates.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", registerStates.WarningThreshold})

    registerStates.EntityData.YListKeys = []string {}

    return &(registerStates.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces
// Override default maximum for number of
// route-interfaces
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM route-interfaces. The type is interface{} with range:
    // 1..1100000. This attribute is mandatory.
    MaximumRouteInterfaces interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000. The default value is 300000.
    WarningThreshold interface{}
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetEntityData() *types.CommonEntityData {
    routeInterfaces.EntityData.YFilter = routeInterfaces.YFilter
    routeInterfaces.EntityData.YangName = "route-interfaces"
    routeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    routeInterfaces.EntityData.ParentYangName = "maximum"
    routeInterfaces.EntityData.SegmentPath = "route-interfaces"
    routeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/maximum/" + routeInterfaces.EntityData.SegmentPath
    routeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeInterfaces.EntityData.Children = types.NewOrderedMap()
    routeInterfaces.EntityData.Leafs = types.NewOrderedMap()
    routeInterfaces.EntityData.Leafs.Append("maximum-route-interfaces", types.YLeaf{"MaximumRouteInterfaces", routeInterfaces.MaximumRouteInterfaces})
    routeInterfaces.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", routeInterfaces.WarningThreshold})

    routeInterfaces.EntityData.YListKeys = []string {}

    return &(routeInterfaces.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache
// Override default maximum and threshold for BSR
// C-RP cache setting
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of BSR C-RP cache setting. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumCandidateRpCache interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 100.
    WarningThreshold interface{}
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetEntityData() *types.CommonEntityData {
    bsrCandidateRpCache.EntityData.YFilter = bsrCandidateRpCache.YFilter
    bsrCandidateRpCache.EntityData.YangName = "bsr-candidate-rp-cache"
    bsrCandidateRpCache.EntityData.BundleName = "cisco_ios_xr"
    bsrCandidateRpCache.EntityData.ParentYangName = "maximum"
    bsrCandidateRpCache.EntityData.SegmentPath = "bsr-candidate-rp-cache"
    bsrCandidateRpCache.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/maximum/" + bsrCandidateRpCache.EntityData.SegmentPath
    bsrCandidateRpCache.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsrCandidateRpCache.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsrCandidateRpCache.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsrCandidateRpCache.EntityData.Children = types.NewOrderedMap()
    bsrCandidateRpCache.EntityData.Leafs = types.NewOrderedMap()
    bsrCandidateRpCache.EntityData.Leafs.Append("bsr-maximum-candidate-rp-cache", types.YLeaf{"BsrMaximumCandidateRpCache", bsrCandidateRpCache.BsrMaximumCandidateRpCache})
    bsrCandidateRpCache.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", bsrCandidateRpCache.WarningThreshold})

    bsrCandidateRpCache.EntityData.YListKeys = []string {}

    return &(bsrCandidateRpCache.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Maximum_Routes
// Override default maximum for number of routes
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Maximum_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM routes. The type is interface{} with range:
    // 1..200000. This attribute is mandatory.
    MaximumRoutes interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..200000. The default value is 100000.
    WarningThreshold interface{}
}

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "maximum"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/maximum/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Leafs = types.NewOrderedMap()
    routes.EntityData.Leafs.Append("maximum-routes", types.YLeaf{"MaximumRoutes", routes.MaximumRoutes})
    routes.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", routes.WarningThreshold})

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer
// Configure expiry timer for S,G routes
type Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // (S,G) expiry time in seconds. The type is interface{} with range:
    // 40..57600. Units are second.
    Interval interface{}

    // Access-list of applicable S,G routes. The type is string with length:
    // 1..64.
    AccessListName interface{}
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetEntityData() *types.CommonEntityData {
    sgExpiryTimer.EntityData.YFilter = sgExpiryTimer.YFilter
    sgExpiryTimer.EntityData.YangName = "sg-expiry-timer"
    sgExpiryTimer.EntityData.BundleName = "cisco_ios_xr"
    sgExpiryTimer.EntityData.ParentYangName = "ipv6"
    sgExpiryTimer.EntityData.SegmentPath = "sg-expiry-timer"
    sgExpiryTimer.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/" + sgExpiryTimer.EntityData.SegmentPath
    sgExpiryTimer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sgExpiryTimer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sgExpiryTimer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sgExpiryTimer.EntityData.Children = types.NewOrderedMap()
    sgExpiryTimer.EntityData.Leafs = types.NewOrderedMap()
    sgExpiryTimer.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", sgExpiryTimer.Interval})
    sgExpiryTimer.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", sgExpiryTimer.AccessListName})

    sgExpiryTimer.EntityData.YListKeys = []string {}

    return &(sgExpiryTimer.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable
// Enable PIM RPF Vector Proxy's
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // RPF Vector is turned on if configured. The type is interface{}. This
    // attribute is mandatory.
    Enable interface{}

    // Allow RPF Vector origination over eBGP sessions. The type is interface{}.
    AllowEbgp interface{}

    // Disable RPF Vector origination over iBGP sessions. The type is interface{}.
    DisableIbgp interface{}

    // Use RPF Vector RFC compliant encoding. The type is interface{}.
    UseStandardEncoding interface{}
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetEntityData() *types.CommonEntityData {
    rpfVectorEnable.EntityData.YFilter = rpfVectorEnable.YFilter
    rpfVectorEnable.EntityData.YangName = "rpf-vector-enable"
    rpfVectorEnable.EntityData.BundleName = "cisco_ios_xr"
    rpfVectorEnable.EntityData.ParentYangName = "ipv6"
    rpfVectorEnable.EntityData.SegmentPath = "rpf-vector-enable"
    rpfVectorEnable.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/" + rpfVectorEnable.EntityData.SegmentPath
    rpfVectorEnable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpfVectorEnable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpfVectorEnable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpfVectorEnable.EntityData.Children = types.NewOrderedMap()
    rpfVectorEnable.EntityData.Leafs = types.NewOrderedMap()
    rpfVectorEnable.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", rpfVectorEnable.Enable})
    rpfVectorEnable.EntityData.Leafs.Append("allow-ebgp", types.YLeaf{"AllowEbgp", rpfVectorEnable.AllowEbgp})
    rpfVectorEnable.EntityData.Leafs.Append("disable-ibgp", types.YLeaf{"DisableIbgp", rpfVectorEnable.DisableIbgp})
    rpfVectorEnable.EntityData.Leafs.Append("use-standard-encoding", types.YLeaf{"UseStandardEncoding", rpfVectorEnable.UseStandardEncoding})

    rpfVectorEnable.EntityData.YListKeys = []string {}

    return &(rpfVectorEnable.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Ssm
// Configure IP Multicast SSM
type Pim_Vrfs_Vrf_Ipv6_Ssm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE if SSM is disabled on this router. The type is bool. The default value
    // is false.
    Disable interface{}

    // Access list of groups enabled with SSM. The type is string with length:
    // 1..64.
    Range interface{}
}

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetEntityData() *types.CommonEntityData {
    ssm.EntityData.YFilter = ssm.YFilter
    ssm.EntityData.YangName = "ssm"
    ssm.EntityData.BundleName = "cisco_ios_xr"
    ssm.EntityData.ParentYangName = "ipv6"
    ssm.EntityData.SegmentPath = "ssm"
    ssm.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/" + ssm.EntityData.SegmentPath
    ssm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssm.EntityData.Children = types.NewOrderedMap()
    ssm.EntityData.Leafs = types.NewOrderedMap()
    ssm.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", ssm.Disable})
    ssm.EntityData.Leafs.Append("range", types.YLeaf{"Range", ssm.Range})

    ssm.EntityData.YListKeys = []string {}

    return &(ssm.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses
// Configure Bidirectional PIM Rendezvous Point
type Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress.
    BidirRpAddress []*Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetEntityData() *types.CommonEntityData {
    bidirRpAddresses.EntityData.YFilter = bidirRpAddresses.YFilter
    bidirRpAddresses.EntityData.YangName = "bidir-rp-addresses"
    bidirRpAddresses.EntityData.BundleName = "cisco_ios_xr"
    bidirRpAddresses.EntityData.ParentYangName = "ipv6"
    bidirRpAddresses.EntityData.SegmentPath = "bidir-rp-addresses"
    bidirRpAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/" + bidirRpAddresses.EntityData.SegmentPath
    bidirRpAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bidirRpAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bidirRpAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bidirRpAddresses.EntityData.Children = types.NewOrderedMap()
    bidirRpAddresses.EntityData.Children.Append("bidir-rp-address", types.YChild{"BidirRpAddress", nil})
    for i := range bidirRpAddresses.BidirRpAddress {
        bidirRpAddresses.EntityData.Children.Append(types.GetSegmentPath(bidirRpAddresses.BidirRpAddress[i]), types.YChild{"BidirRpAddress", bidirRpAddresses.BidirRpAddress[i]})
    }
    bidirRpAddresses.EntityData.Leafs = types.NewOrderedMap()

    bidirRpAddresses.EntityData.YListKeys = []string {}

    return &(bidirRpAddresses.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress
// Address of the Rendezvous Point
type Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetEntityData() *types.CommonEntityData {
    bidirRpAddress.EntityData.YFilter = bidirRpAddress.YFilter
    bidirRpAddress.EntityData.YangName = "bidir-rp-address"
    bidirRpAddress.EntityData.BundleName = "cisco_ios_xr"
    bidirRpAddress.EntityData.ParentYangName = "bidir-rp-addresses"
    bidirRpAddress.EntityData.SegmentPath = "bidir-rp-address" + types.AddKeyToken(bidirRpAddress.RpAddress, "rp-address")
    bidirRpAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/bidir-rp-addresses/" + bidirRpAddress.EntityData.SegmentPath
    bidirRpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bidirRpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bidirRpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bidirRpAddress.EntityData.Children = types.NewOrderedMap()
    bidirRpAddress.EntityData.Leafs = types.NewOrderedMap()
    bidirRpAddress.EntityData.Leafs.Append("rp-address", types.YLeaf{"RpAddress", bidirRpAddress.RpAddress})
    bidirRpAddress.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", bidirRpAddress.AccessListName})
    bidirRpAddress.EntityData.Leafs.Append("auto-rp-override", types.YLeaf{"AutoRpOverride", bidirRpAddress.AutoRpOverride})

    bidirRpAddress.EntityData.YListKeys = []string {"RpAddress"}

    return &(bidirRpAddress.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Bsr
// PIM BSR configuration
type Pim_Vrfs_Vrf_Ipv6_Bsr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PIM Candidate BSR configuration.
    CandidateBsr Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr

    // PIM RP configuration.
    CandidateRps Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps
}

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetEntityData() *types.CommonEntityData {
    bsr.EntityData.YFilter = bsr.YFilter
    bsr.EntityData.YangName = "bsr"
    bsr.EntityData.BundleName = "cisco_ios_xr"
    bsr.EntityData.ParentYangName = "ipv6"
    bsr.EntityData.SegmentPath = "bsr"
    bsr.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/" + bsr.EntityData.SegmentPath
    bsr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsr.EntityData.Children = types.NewOrderedMap()
    bsr.EntityData.Children.Append("candidate-bsr", types.YChild{"CandidateBsr", &bsr.CandidateBsr})
    bsr.EntityData.Children.Append("candidate-rps", types.YChild{"CandidateRps", &bsr.CandidateRps})
    bsr.EntityData.Leafs = types.NewOrderedMap()

    bsr.EntityData.YListKeys = []string {}

    return &(bsr.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr
// PIM Candidate BSR configuration
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // BSR Address configured. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}

    // Hash Mask Length for this candidate BSR. The type is interface{} with
    // range: 0..128. The default value is 126.
    PrefixLength interface{}

    // Priority of the Candidate BSR. The type is interface{} with range: 1..255.
    // The default value is 1.
    Priority interface{}
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetEntityData() *types.CommonEntityData {
    candidateBsr.EntityData.YFilter = candidateBsr.YFilter
    candidateBsr.EntityData.YangName = "candidate-bsr"
    candidateBsr.EntityData.BundleName = "cisco_ios_xr"
    candidateBsr.EntityData.ParentYangName = "bsr"
    candidateBsr.EntityData.SegmentPath = "candidate-bsr"
    candidateBsr.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/bsr/" + candidateBsr.EntityData.SegmentPath
    candidateBsr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateBsr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateBsr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateBsr.EntityData.Children = types.NewOrderedMap()
    candidateBsr.EntityData.Leafs = types.NewOrderedMap()
    candidateBsr.EntityData.Leafs.Append("address", types.YLeaf{"Address", candidateBsr.Address})
    candidateBsr.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", candidateBsr.PrefixLength})
    candidateBsr.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", candidateBsr.Priority})

    candidateBsr.EntityData.YListKeys = []string {}

    return &(candidateBsr.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps
// PIM RP configuration
type Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of PIM SM BSR Candidate-RP. The type is slice of
    // Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp.
    CandidateRp []*Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetEntityData() *types.CommonEntityData {
    candidateRps.EntityData.YFilter = candidateRps.YFilter
    candidateRps.EntityData.YangName = "candidate-rps"
    candidateRps.EntityData.BundleName = "cisco_ios_xr"
    candidateRps.EntityData.ParentYangName = "bsr"
    candidateRps.EntityData.SegmentPath = "candidate-rps"
    candidateRps.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/bsr/" + candidateRps.EntityData.SegmentPath
    candidateRps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateRps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateRps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateRps.EntityData.Children = types.NewOrderedMap()
    candidateRps.EntityData.Children.Append("candidate-rp", types.YChild{"CandidateRp", nil})
    for i := range candidateRps.CandidateRp {
        candidateRps.EntityData.Children.Append(types.GetSegmentPath(candidateRps.CandidateRp[i]), types.YChild{"CandidateRp", candidateRps.CandidateRp[i]})
    }
    candidateRps.EntityData.Leafs = types.NewOrderedMap()

    candidateRps.EntityData.YListKeys = []string {}

    return &(candidateRps.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp
// Address of PIM SM BSR Candidate-RP
type Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address of Candidate-RP. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. SM or Bidir. The type is PimProtocolMode.
    Mode interface{}

    // Access-list specifying the group range for the Candidate-RP. The type is
    // string with length: 1..64.
    GroupList interface{}

    // Priority of the CRP. The type is interface{} with range: 1..255. The
    // default value is 192.
    Priority interface{}

    // Advertisement interval. The type is interface{} with range: 30..600. The
    // default value is 60.
    Interval interface{}
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetEntityData() *types.CommonEntityData {
    candidateRp.EntityData.YFilter = candidateRp.YFilter
    candidateRp.EntityData.YangName = "candidate-rp"
    candidateRp.EntityData.BundleName = "cisco_ios_xr"
    candidateRp.EntityData.ParentYangName = "candidate-rps"
    candidateRp.EntityData.SegmentPath = "candidate-rp" + types.AddKeyToken(candidateRp.Address, "address") + types.AddKeyToken(candidateRp.Mode, "mode")
    candidateRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/bsr/candidate-rps/" + candidateRp.EntityData.SegmentPath
    candidateRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateRp.EntityData.Children = types.NewOrderedMap()
    candidateRp.EntityData.Leafs = types.NewOrderedMap()
    candidateRp.EntityData.Leafs.Append("address", types.YLeaf{"Address", candidateRp.Address})
    candidateRp.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", candidateRp.Mode})
    candidateRp.EntityData.Leafs.Append("group-list", types.YLeaf{"GroupList", candidateRp.GroupList})
    candidateRp.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", candidateRp.Priority})
    candidateRp.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", candidateRp.Interval})

    candidateRp.EntityData.YListKeys = []string {"Address", "Mode"}

    return &(candidateRp.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_AllowRp
// Enable allow-rp filtering for SM joins
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_AllowRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Access-list specifiying applicable RPs. The type is string with length:
    // 1..64.
    RpListName interface{}

    // Access-list specifiying applicable groups. The type is string with length:
    // 1..64.
    GroupListName interface{}
}

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetEntityData() *types.CommonEntityData {
    allowRp.EntityData.YFilter = allowRp.YFilter
    allowRp.EntityData.YangName = "allow-rp"
    allowRp.EntityData.BundleName = "cisco_ios_xr"
    allowRp.EntityData.ParentYangName = "ipv6"
    allowRp.EntityData.SegmentPath = "allow-rp"
    allowRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/" + allowRp.EntityData.SegmentPath
    allowRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowRp.EntityData.Children = types.NewOrderedMap()
    allowRp.EntityData.Leafs = types.NewOrderedMap()
    allowRp.EntityData.Leafs.Append("rp-list-name", types.YLeaf{"RpListName", allowRp.RpListName})
    allowRp.EntityData.Leafs.Append("group-list-name", types.YLeaf{"GroupListName", allowRp.GroupListName})

    allowRp.EntityData.YListKeys = []string {}

    return &(allowRp.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses
// Set Embedded RP processing support
type Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set Embedded RP processing support. The type is slice of
    // Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress.
    EmbeddedRpAddress []*Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress
}

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetEntityData() *types.CommonEntityData {
    embeddedRpAddresses.EntityData.YFilter = embeddedRpAddresses.YFilter
    embeddedRpAddresses.EntityData.YangName = "embedded-rp-addresses"
    embeddedRpAddresses.EntityData.BundleName = "cisco_ios_xr"
    embeddedRpAddresses.EntityData.ParentYangName = "ipv6"
    embeddedRpAddresses.EntityData.SegmentPath = "embedded-rp-addresses"
    embeddedRpAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/" + embeddedRpAddresses.EntityData.SegmentPath
    embeddedRpAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    embeddedRpAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    embeddedRpAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    embeddedRpAddresses.EntityData.Children = types.NewOrderedMap()
    embeddedRpAddresses.EntityData.Children.Append("embedded-rp-address", types.YChild{"EmbeddedRpAddress", nil})
    for i := range embeddedRpAddresses.EmbeddedRpAddress {
        embeddedRpAddresses.EntityData.Children.Append(types.GetSegmentPath(embeddedRpAddresses.EmbeddedRpAddress[i]), types.YChild{"EmbeddedRpAddress", embeddedRpAddresses.EmbeddedRpAddress[i]})
    }
    embeddedRpAddresses.EntityData.Leafs = types.NewOrderedMap()

    embeddedRpAddresses.EntityData.YListKeys = []string {}

    return &(embeddedRpAddresses.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress
// Set Embedded RP processing support
type Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. RP address of the Rendezvous Point. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a given RP. The type is string
    // with length: 1..64. This attribute is mandatory.
    AccessListName interface{}
}

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetEntityData() *types.CommonEntityData {
    embeddedRpAddress.EntityData.YFilter = embeddedRpAddress.YFilter
    embeddedRpAddress.EntityData.YangName = "embedded-rp-address"
    embeddedRpAddress.EntityData.BundleName = "cisco_ios_xr"
    embeddedRpAddress.EntityData.ParentYangName = "embedded-rp-addresses"
    embeddedRpAddress.EntityData.SegmentPath = "embedded-rp-address" + types.AddKeyToken(embeddedRpAddress.RpAddress, "rp-address")
    embeddedRpAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/embedded-rp-addresses/" + embeddedRpAddress.EntityData.SegmentPath
    embeddedRpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    embeddedRpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    embeddedRpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    embeddedRpAddress.EntityData.Children = types.NewOrderedMap()
    embeddedRpAddress.EntityData.Leafs = types.NewOrderedMap()
    embeddedRpAddress.EntityData.Leafs.Append("rp-address", types.YLeaf{"RpAddress", embeddedRpAddress.RpAddress})
    embeddedRpAddress.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", embeddedRpAddress.AccessListName})

    embeddedRpAddress.EntityData.YListKeys = []string {"RpAddress"}

    return &(embeddedRpAddress.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Convergence
// Configure convergence parameters
type Pim_Vrfs_Vrf_Ipv6_Convergence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampen first join if RPF path is through one of the downstream neighbor.
    // The type is interface{} with range: 0..15. Units are second.
    RpfConflictJoinDelay interface{}

    // Delay prunes if route join state transitions to not-joined on link down.
    // The type is interface{} with range: 0..60. Units are second.
    LinkDownPruneDelay interface{}
}

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetEntityData() *types.CommonEntityData {
    convergence.EntityData.YFilter = convergence.YFilter
    convergence.EntityData.YangName = "convergence"
    convergence.EntityData.BundleName = "cisco_ios_xr"
    convergence.EntityData.ParentYangName = "ipv6"
    convergence.EntityData.SegmentPath = "convergence"
    convergence.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/" + convergence.EntityData.SegmentPath
    convergence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    convergence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    convergence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    convergence.EntityData.Children = types.NewOrderedMap()
    convergence.EntityData.Leafs = types.NewOrderedMap()
    convergence.EntityData.Leafs.Append("rpf-conflict-join-delay", types.YLeaf{"RpfConflictJoinDelay", convergence.RpfConflictJoinDelay})
    convergence.EntityData.Leafs.Append("link-down-prune-delay", types.YLeaf{"LinkDownPruneDelay", convergence.LinkDownPruneDelay})

    convergence.EntityData.YListKeys = []string {}

    return &(convergence.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Interfaces
// Interface-level Configuration
type Pim_Vrfs_Vrf_Ipv6_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface.
    Interface []*Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface
}

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ipv6"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/" + interfaces.EntityData.SegmentPath
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

// Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface
// The name of the interface
type Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of interface. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Enter PIM Interface processing. The type is interface{}.
    Enable interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // BSR Border configuration for Interface. The type is bool.
    BsrBorder interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Enable PIM processing on the interface. The type is bool.
    InterfaceEnable interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}

    // Maximum number of allowed routes for this interface.
    MaximumRoutes Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes

    // BFD configuration.
    Bfd Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd
}

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("maximum-routes", types.YChild{"MaximumRoutes", &self.MaximumRoutes})
    self.EntityData.Children.Append("bfd", types.YChild{"Bfd", &self.Bfd})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})
    self.EntityData.Leafs.Append("neighbor-filter", types.YLeaf{"NeighborFilter", self.NeighborFilter})
    self.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", self.HelloInterval})
    self.EntityData.Leafs.Append("bsr-border", types.YLeaf{"BsrBorder", self.BsrBorder})
    self.EntityData.Leafs.Append("propagation-delay", types.YLeaf{"PropagationDelay", self.PropagationDelay})
    self.EntityData.Leafs.Append("dr-priority", types.YLeaf{"DrPriority", self.DrPriority})
    self.EntityData.Leafs.Append("join-prune-mtu", types.YLeaf{"JoinPruneMtu", self.JoinPruneMtu})
    self.EntityData.Leafs.Append("interface-enable", types.YLeaf{"InterfaceEnable", self.InterfaceEnable})
    self.EntityData.Leafs.Append("jp-interval", types.YLeaf{"JpInterval", self.JpInterval})
    self.EntityData.Leafs.Append("override-interval", types.YLeaf{"OverrideInterval", self.OverrideInterval})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes
// Maximum number of allowed routes for this
// interface
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of routes for this interface. The type is interface{} with
    // range: 1..1100000. This attribute is mandatory.
    Maximum interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetEntityData() *types.CommonEntityData {
    maximumRoutes.EntityData.YFilter = maximumRoutes.YFilter
    maximumRoutes.EntityData.YangName = "maximum-routes"
    maximumRoutes.EntityData.BundleName = "cisco_ios_xr"
    maximumRoutes.EntityData.ParentYangName = "interface"
    maximumRoutes.EntityData.SegmentPath = "maximum-routes"
    maximumRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/interfaces/interface/" + maximumRoutes.EntityData.SegmentPath
    maximumRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumRoutes.EntityData.Children = types.NewOrderedMap()
    maximumRoutes.EntityData.Leafs = types.NewOrderedMap()
    maximumRoutes.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", maximumRoutes.Maximum})
    maximumRoutes.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", maximumRoutes.WarningThreshold})
    maximumRoutes.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", maximumRoutes.AccessListName})

    maximumRoutes.EntityData.YListKeys = []string {}

    return &(maximumRoutes.EntityData)
}

// Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd
// BFD configuration
type Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detection multiplier for BFD sessions created by PIM. The type is
    // interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval for BFD sessions created by PIM. The type is interface{}
    // with range: 3..30000. Units are millisecond.
    Interval interface{}

    // TRUE to enable BFD. FALSE to disable and to prevent inheritance from a
    // parent. The type is bool.
    Enable interface{}
}

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "interface"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/vrfs/vrf/ipv6/interfaces/interface/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bfd.Enable})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Pim_DefaultContext
// Default Context
type Pim_DefaultContext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPV6 commands.
    Ipv6 Pim_DefaultContext_Ipv6

    // IPV4 commands.
    Ipv4 Pim_DefaultContext_Ipv4
}

func (defaultContext *Pim_DefaultContext) GetEntityData() *types.CommonEntityData {
    defaultContext.EntityData.YFilter = defaultContext.YFilter
    defaultContext.EntityData.YangName = "default-context"
    defaultContext.EntityData.BundleName = "cisco_ios_xr"
    defaultContext.EntityData.ParentYangName = "pim"
    defaultContext.EntityData.SegmentPath = "default-context"
    defaultContext.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/" + defaultContext.EntityData.SegmentPath
    defaultContext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultContext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultContext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultContext.EntityData.Children = types.NewOrderedMap()
    defaultContext.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &defaultContext.Ipv6})
    defaultContext.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &defaultContext.Ipv4})
    defaultContext.EntityData.Leafs = types.NewOrderedMap()

    defaultContext.EntityData.YListKeys = []string {}

    return &(defaultContext.EntityData)
}

// Pim_DefaultContext_Ipv6
// IPV6 commands
type Pim_DefaultContext_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable PIM neighbor checking when receiving PIM messages. The type is
    // interface{}.
    NeighborCheckOnReceive interface{}

    // Generate registers compatible with older IOS versions. The type is
    // interface{}.
    OldRegisterChecksum interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Configure threshold of infinity for switching to SPT on last-hop. The type
    // is string.
    SptThresholdInfinity interface{}

    // PIM neighbor state change logging is turned on if configured. The type is
    // interface{}.
    LogNeighborChanges interface{}

    // Source address to use for register messages. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    RegisterSource interface{}

    // Access-list which specifies unauthorized sources. The type is string with
    // length: 1..64.
    AcceptRegister interface{}

    // Set Embedded RP processing support. The type is interface{}.
    EmbeddedRpDisable interface{}

    // Suppress prunes triggered as a result of RPF changes. The type is
    // interface{}.
    SuppressRpfPrunes interface{}

    // Allow SSM ranges to be overridden by more specific ranges. The type is
    // interface{}.
    SsmAllowOverride interface{}

    // Enable equal-cost multipath routing. The type is PimMultipath.
    Multipath interface{}

    // Configure static RP deny range. The type is string with length: 1..64.
    RpStaticDeny interface{}

    // Suppress data registers after initial state setup. The type is interface{}.
    SuppressDataRegisters interface{}

    // Enable PIM neighbor checking when sending join-prunes. The type is
    // interface{}.
    NeighborCheckOnSend interface{}

    // Interface-level Configuration.
    Interfaces Pim_DefaultContext_Ipv6_Interfaces

    // Configure Sparse-Mode Rendezvous Point.
    SparseModeRpAddresses Pim_DefaultContext_Ipv6_SparseModeRpAddresses

    // Inheritable defaults.
    InheritableDefaults Pim_DefaultContext_Ipv6_InheritableDefaults

    // Configure RPF options.
    Rpf Pim_DefaultContext_Ipv6_Rpf

    // Configure expiry timer for S,G routes.
    SgExpiryTimer Pim_DefaultContext_Ipv6_SgExpiryTimer

    // Enable PIM RPF Vector Proxy's.
    RpfVectorEnable Pim_DefaultContext_Ipv6_RpfVectorEnable

    // Configure Non-stop forwarding (NSF) options.
    Nsf Pim_DefaultContext_Ipv6_Nsf

    // Configure PIM State Limits.
    Maximum Pim_DefaultContext_Ipv6_Maximum

    // Configure IP Multicast SSM.
    Ssm Pim_DefaultContext_Ipv6_Ssm

    // Configure Bidirectional PIM Rendezvous Point.
    BidirRpAddresses Pim_DefaultContext_Ipv6_BidirRpAddresses

    // PIM BSR configuration.
    Bsr Pim_DefaultContext_Ipv6_Bsr

    // Enable allow-rp filtering for SM joins.
    AllowRp Pim_DefaultContext_Ipv6_AllowRp

    // Set Embedded RP processing support.
    EmbeddedRpAddresses Pim_DefaultContext_Ipv6_EmbeddedRpAddresses

    // Configure convergence parameters.
    Convergence Pim_DefaultContext_Ipv6_Convergence
}

func (ipv6 *Pim_DefaultContext_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "default-context"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/" + ipv6.EntityData.SegmentPath
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &ipv6.Interfaces})
    ipv6.EntityData.Children.Append("sparse-mode-rp-addresses", types.YChild{"SparseModeRpAddresses", &ipv6.SparseModeRpAddresses})
    ipv6.EntityData.Children.Append("inheritable-defaults", types.YChild{"InheritableDefaults", &ipv6.InheritableDefaults})
    ipv6.EntityData.Children.Append("rpf", types.YChild{"Rpf", &ipv6.Rpf})
    ipv6.EntityData.Children.Append("sg-expiry-timer", types.YChild{"SgExpiryTimer", &ipv6.SgExpiryTimer})
    ipv6.EntityData.Children.Append("rpf-vector-enable", types.YChild{"RpfVectorEnable", &ipv6.RpfVectorEnable})
    ipv6.EntityData.Children.Append("nsf", types.YChild{"Nsf", &ipv6.Nsf})
    ipv6.EntityData.Children.Append("maximum", types.YChild{"Maximum", &ipv6.Maximum})
    ipv6.EntityData.Children.Append("ssm", types.YChild{"Ssm", &ipv6.Ssm})
    ipv6.EntityData.Children.Append("bidir-rp-addresses", types.YChild{"BidirRpAddresses", &ipv6.BidirRpAddresses})
    ipv6.EntityData.Children.Append("bsr", types.YChild{"Bsr", &ipv6.Bsr})
    ipv6.EntityData.Children.Append("allow-rp", types.YChild{"AllowRp", &ipv6.AllowRp})
    ipv6.EntityData.Children.Append("embedded-rp-addresses", types.YChild{"EmbeddedRpAddresses", &ipv6.EmbeddedRpAddresses})
    ipv6.EntityData.Children.Append("convergence", types.YChild{"Convergence", &ipv6.Convergence})
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("neighbor-check-on-receive", types.YLeaf{"NeighborCheckOnReceive", ipv6.NeighborCheckOnReceive})
    ipv6.EntityData.Leafs.Append("old-register-checksum", types.YLeaf{"OldRegisterChecksum", ipv6.OldRegisterChecksum})
    ipv6.EntityData.Leafs.Append("neighbor-filter", types.YLeaf{"NeighborFilter", ipv6.NeighborFilter})
    ipv6.EntityData.Leafs.Append("spt-threshold-infinity", types.YLeaf{"SptThresholdInfinity", ipv6.SptThresholdInfinity})
    ipv6.EntityData.Leafs.Append("log-neighbor-changes", types.YLeaf{"LogNeighborChanges", ipv6.LogNeighborChanges})
    ipv6.EntityData.Leafs.Append("register-source", types.YLeaf{"RegisterSource", ipv6.RegisterSource})
    ipv6.EntityData.Leafs.Append("accept-register", types.YLeaf{"AcceptRegister", ipv6.AcceptRegister})
    ipv6.EntityData.Leafs.Append("embedded-rp-disable", types.YLeaf{"EmbeddedRpDisable", ipv6.EmbeddedRpDisable})
    ipv6.EntityData.Leafs.Append("suppress-rpf-prunes", types.YLeaf{"SuppressRpfPrunes", ipv6.SuppressRpfPrunes})
    ipv6.EntityData.Leafs.Append("ssm-allow-override", types.YLeaf{"SsmAllowOverride", ipv6.SsmAllowOverride})
    ipv6.EntityData.Leafs.Append("multipath", types.YLeaf{"Multipath", ipv6.Multipath})
    ipv6.EntityData.Leafs.Append("rp-static-deny", types.YLeaf{"RpStaticDeny", ipv6.RpStaticDeny})
    ipv6.EntityData.Leafs.Append("suppress-data-registers", types.YLeaf{"SuppressDataRegisters", ipv6.SuppressDataRegisters})
    ipv6.EntityData.Leafs.Append("neighbor-check-on-send", types.YLeaf{"NeighborCheckOnSend", ipv6.NeighborCheckOnSend})

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Pim_DefaultContext_Ipv6_Interfaces
// Interface-level Configuration
type Pim_DefaultContext_Ipv6_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Pim_DefaultContext_Ipv6_Interfaces_Interface.
    Interface []*Pim_DefaultContext_Ipv6_Interfaces_Interface
}

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ipv6"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + interfaces.EntityData.SegmentPath
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

// Pim_DefaultContext_Ipv6_Interfaces_Interface
// The name of the interface
type Pim_DefaultContext_Ipv6_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of interface. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Enter PIM Interface processing. The type is interface{}.
    Enable interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // BSR Border configuration for Interface. The type is bool.
    BsrBorder interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Enable PIM processing on the interface. The type is bool.
    InterfaceEnable interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}

    // Maximum number of allowed routes for this interface.
    MaximumRoutes Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes

    // BFD configuration.
    Bfd Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd
}

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("maximum-routes", types.YChild{"MaximumRoutes", &self.MaximumRoutes})
    self.EntityData.Children.Append("bfd", types.YChild{"Bfd", &self.Bfd})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})
    self.EntityData.Leafs.Append("neighbor-filter", types.YLeaf{"NeighborFilter", self.NeighborFilter})
    self.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", self.HelloInterval})
    self.EntityData.Leafs.Append("bsr-border", types.YLeaf{"BsrBorder", self.BsrBorder})
    self.EntityData.Leafs.Append("propagation-delay", types.YLeaf{"PropagationDelay", self.PropagationDelay})
    self.EntityData.Leafs.Append("dr-priority", types.YLeaf{"DrPriority", self.DrPriority})
    self.EntityData.Leafs.Append("join-prune-mtu", types.YLeaf{"JoinPruneMtu", self.JoinPruneMtu})
    self.EntityData.Leafs.Append("interface-enable", types.YLeaf{"InterfaceEnable", self.InterfaceEnable})
    self.EntityData.Leafs.Append("jp-interval", types.YLeaf{"JpInterval", self.JpInterval})
    self.EntityData.Leafs.Append("override-interval", types.YLeaf{"OverrideInterval", self.OverrideInterval})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes
// Maximum number of allowed routes for this
// interface
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of routes for this interface. The type is interface{} with
    // range: 1..1100000. This attribute is mandatory.
    Maximum interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetEntityData() *types.CommonEntityData {
    maximumRoutes.EntityData.YFilter = maximumRoutes.YFilter
    maximumRoutes.EntityData.YangName = "maximum-routes"
    maximumRoutes.EntityData.BundleName = "cisco_ios_xr"
    maximumRoutes.EntityData.ParentYangName = "interface"
    maximumRoutes.EntityData.SegmentPath = "maximum-routes"
    maximumRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/interfaces/interface/" + maximumRoutes.EntityData.SegmentPath
    maximumRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumRoutes.EntityData.Children = types.NewOrderedMap()
    maximumRoutes.EntityData.Leafs = types.NewOrderedMap()
    maximumRoutes.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", maximumRoutes.Maximum})
    maximumRoutes.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", maximumRoutes.WarningThreshold})
    maximumRoutes.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", maximumRoutes.AccessListName})

    maximumRoutes.EntityData.YListKeys = []string {}

    return &(maximumRoutes.EntityData)
}

// Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd
// BFD configuration
type Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detection multiplier for BFD sessions created by PIM. The type is
    // interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval for BFD sessions created by PIM. The type is interface{}
    // with range: 3..30000. Units are millisecond.
    Interval interface{}

    // TRUE to enable BFD. FALSE to disable and to prevent inheritance from a
    // parent. The type is bool.
    Enable interface{}
}

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "interface"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/interfaces/interface/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bfd.Enable})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Pim_DefaultContext_Ipv6_SparseModeRpAddresses
// Configure Sparse-Mode Rendezvous Point
type Pim_DefaultContext_Ipv6_SparseModeRpAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress.
    SparseModeRpAddress []*Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetEntityData() *types.CommonEntityData {
    sparseModeRpAddresses.EntityData.YFilter = sparseModeRpAddresses.YFilter
    sparseModeRpAddresses.EntityData.YangName = "sparse-mode-rp-addresses"
    sparseModeRpAddresses.EntityData.BundleName = "cisco_ios_xr"
    sparseModeRpAddresses.EntityData.ParentYangName = "ipv6"
    sparseModeRpAddresses.EntityData.SegmentPath = "sparse-mode-rp-addresses"
    sparseModeRpAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + sparseModeRpAddresses.EntityData.SegmentPath
    sparseModeRpAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sparseModeRpAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sparseModeRpAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sparseModeRpAddresses.EntityData.Children = types.NewOrderedMap()
    sparseModeRpAddresses.EntityData.Children.Append("sparse-mode-rp-address", types.YChild{"SparseModeRpAddress", nil})
    for i := range sparseModeRpAddresses.SparseModeRpAddress {
        sparseModeRpAddresses.EntityData.Children.Append(types.GetSegmentPath(sparseModeRpAddresses.SparseModeRpAddress[i]), types.YChild{"SparseModeRpAddress", sparseModeRpAddresses.SparseModeRpAddress[i]})
    }
    sparseModeRpAddresses.EntityData.Leafs = types.NewOrderedMap()

    sparseModeRpAddresses.EntityData.YListKeys = []string {}

    return &(sparseModeRpAddresses.EntityData)
}

// Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress
// Address of the Rendezvous Point
type Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a  given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetEntityData() *types.CommonEntityData {
    sparseModeRpAddress.EntityData.YFilter = sparseModeRpAddress.YFilter
    sparseModeRpAddress.EntityData.YangName = "sparse-mode-rp-address"
    sparseModeRpAddress.EntityData.BundleName = "cisco_ios_xr"
    sparseModeRpAddress.EntityData.ParentYangName = "sparse-mode-rp-addresses"
    sparseModeRpAddress.EntityData.SegmentPath = "sparse-mode-rp-address" + types.AddKeyToken(sparseModeRpAddress.RpAddress, "rp-address")
    sparseModeRpAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/sparse-mode-rp-addresses/" + sparseModeRpAddress.EntityData.SegmentPath
    sparseModeRpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sparseModeRpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sparseModeRpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sparseModeRpAddress.EntityData.Children = types.NewOrderedMap()
    sparseModeRpAddress.EntityData.Leafs = types.NewOrderedMap()
    sparseModeRpAddress.EntityData.Leafs.Append("rp-address", types.YLeaf{"RpAddress", sparseModeRpAddress.RpAddress})
    sparseModeRpAddress.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", sparseModeRpAddress.AccessListName})
    sparseModeRpAddress.EntityData.Leafs.Append("auto-rp-override", types.YLeaf{"AutoRpOverride", sparseModeRpAddress.AutoRpOverride})

    sparseModeRpAddress.EntityData.YListKeys = []string {"RpAddress"}

    return &(sparseModeRpAddress.EntityData)
}

// Pim_DefaultContext_Ipv6_InheritableDefaults
// Inheritable defaults
type Pim_DefaultContext_Ipv6_InheritableDefaults struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Convergency timeout in seconds. The type is interface{} with range:
    // 1800..2400. Units are second.
    ConvergenceTimeout interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}
}

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetEntityData() *types.CommonEntityData {
    inheritableDefaults.EntityData.YFilter = inheritableDefaults.YFilter
    inheritableDefaults.EntityData.YangName = "inheritable-defaults"
    inheritableDefaults.EntityData.BundleName = "cisco_ios_xr"
    inheritableDefaults.EntityData.ParentYangName = "ipv6"
    inheritableDefaults.EntityData.SegmentPath = "inheritable-defaults"
    inheritableDefaults.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + inheritableDefaults.EntityData.SegmentPath
    inheritableDefaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritableDefaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritableDefaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritableDefaults.EntityData.Children = types.NewOrderedMap()
    inheritableDefaults.EntityData.Leafs = types.NewOrderedMap()
    inheritableDefaults.EntityData.Leafs.Append("convergence-timeout", types.YLeaf{"ConvergenceTimeout", inheritableDefaults.ConvergenceTimeout})
    inheritableDefaults.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", inheritableDefaults.HelloInterval})
    inheritableDefaults.EntityData.Leafs.Append("propagation-delay", types.YLeaf{"PropagationDelay", inheritableDefaults.PropagationDelay})
    inheritableDefaults.EntityData.Leafs.Append("dr-priority", types.YLeaf{"DrPriority", inheritableDefaults.DrPriority})
    inheritableDefaults.EntityData.Leafs.Append("join-prune-mtu", types.YLeaf{"JoinPruneMtu", inheritableDefaults.JoinPruneMtu})
    inheritableDefaults.EntityData.Leafs.Append("jp-interval", types.YLeaf{"JpInterval", inheritableDefaults.JpInterval})
    inheritableDefaults.EntityData.Leafs.Append("override-interval", types.YLeaf{"OverrideInterval", inheritableDefaults.OverrideInterval})

    inheritableDefaults.EntityData.YListKeys = []string {}

    return &(inheritableDefaults.EntityData)
}

// Pim_DefaultContext_Ipv6_Rpf
// Configure RPF options
type Pim_DefaultContext_Ipv6_Rpf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route policy to select RPF topology. The type is string with length: 1..64.
    RoutePolicy interface{}
}

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetEntityData() *types.CommonEntityData {
    rpf.EntityData.YFilter = rpf.YFilter
    rpf.EntityData.YangName = "rpf"
    rpf.EntityData.BundleName = "cisco_ios_xr"
    rpf.EntityData.ParentYangName = "ipv6"
    rpf.EntityData.SegmentPath = "rpf"
    rpf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + rpf.EntityData.SegmentPath
    rpf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpf.EntityData.Children = types.NewOrderedMap()
    rpf.EntityData.Leafs = types.NewOrderedMap()
    rpf.EntityData.Leafs.Append("route-policy", types.YLeaf{"RoutePolicy", rpf.RoutePolicy})

    rpf.EntityData.YListKeys = []string {}

    return &(rpf.EntityData)
}

// Pim_DefaultContext_Ipv6_SgExpiryTimer
// Configure expiry timer for S,G routes
type Pim_DefaultContext_Ipv6_SgExpiryTimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // (S,G) expiry time in seconds. The type is interface{} with range:
    // 40..57600. Units are second.
    Interval interface{}

    // Access-list of applicable S,G routes. The type is string with length:
    // 1..64.
    AccessListName interface{}
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetEntityData() *types.CommonEntityData {
    sgExpiryTimer.EntityData.YFilter = sgExpiryTimer.YFilter
    sgExpiryTimer.EntityData.YangName = "sg-expiry-timer"
    sgExpiryTimer.EntityData.BundleName = "cisco_ios_xr"
    sgExpiryTimer.EntityData.ParentYangName = "ipv6"
    sgExpiryTimer.EntityData.SegmentPath = "sg-expiry-timer"
    sgExpiryTimer.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + sgExpiryTimer.EntityData.SegmentPath
    sgExpiryTimer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sgExpiryTimer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sgExpiryTimer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sgExpiryTimer.EntityData.Children = types.NewOrderedMap()
    sgExpiryTimer.EntityData.Leafs = types.NewOrderedMap()
    sgExpiryTimer.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", sgExpiryTimer.Interval})
    sgExpiryTimer.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", sgExpiryTimer.AccessListName})

    sgExpiryTimer.EntityData.YListKeys = []string {}

    return &(sgExpiryTimer.EntityData)
}

// Pim_DefaultContext_Ipv6_RpfVectorEnable
// Enable PIM RPF Vector Proxy's
// This type is a presence type.
type Pim_DefaultContext_Ipv6_RpfVectorEnable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // RPF Vector is turned on if configured. The type is interface{}. This
    // attribute is mandatory.
    Enable interface{}

    // Allow RPF Vector origination over eBGP sessions. The type is interface{}.
    AllowEbgp interface{}

    // Disable RPF Vector origination over iBGP sessions. The type is interface{}.
    DisableIbgp interface{}

    // Use RPF Vector RFC compliant encoding. The type is interface{}.
    UseStandardEncoding interface{}
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetEntityData() *types.CommonEntityData {
    rpfVectorEnable.EntityData.YFilter = rpfVectorEnable.YFilter
    rpfVectorEnable.EntityData.YangName = "rpf-vector-enable"
    rpfVectorEnable.EntityData.BundleName = "cisco_ios_xr"
    rpfVectorEnable.EntityData.ParentYangName = "ipv6"
    rpfVectorEnable.EntityData.SegmentPath = "rpf-vector-enable"
    rpfVectorEnable.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + rpfVectorEnable.EntityData.SegmentPath
    rpfVectorEnable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpfVectorEnable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpfVectorEnable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpfVectorEnable.EntityData.Children = types.NewOrderedMap()
    rpfVectorEnable.EntityData.Leafs = types.NewOrderedMap()
    rpfVectorEnable.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", rpfVectorEnable.Enable})
    rpfVectorEnable.EntityData.Leafs.Append("allow-ebgp", types.YLeaf{"AllowEbgp", rpfVectorEnable.AllowEbgp})
    rpfVectorEnable.EntityData.Leafs.Append("disable-ibgp", types.YLeaf{"DisableIbgp", rpfVectorEnable.DisableIbgp})
    rpfVectorEnable.EntityData.Leafs.Append("use-standard-encoding", types.YLeaf{"UseStandardEncoding", rpfVectorEnable.UseStandardEncoding})

    rpfVectorEnable.EntityData.YListKeys = []string {}

    return &(rpfVectorEnable.EntityData)
}

// Pim_DefaultContext_Ipv6_Nsf
// Configure Non-stop forwarding (NSF) options
type Pim_DefaultContext_Ipv6_Nsf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Override default maximum lifetime for PIM NSF mode. The type is interface{}
    // with range: 10..600. Units are second.
    Lifetime interface{}
}

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetEntityData() *types.CommonEntityData {
    nsf.EntityData.YFilter = nsf.YFilter
    nsf.EntityData.YangName = "nsf"
    nsf.EntityData.BundleName = "cisco_ios_xr"
    nsf.EntityData.ParentYangName = "ipv6"
    nsf.EntityData.SegmentPath = "nsf"
    nsf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + nsf.EntityData.SegmentPath
    nsf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nsf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nsf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nsf.EntityData.Children = types.NewOrderedMap()
    nsf.EntityData.Leafs = types.NewOrderedMap()
    nsf.EntityData.Leafs.Append("lifetime", types.YLeaf{"Lifetime", nsf.Lifetime})

    nsf.EntityData.YListKeys = []string {}

    return &(nsf.EntityData)
}

// Pim_DefaultContext_Ipv6_Maximum
// Configure PIM State Limits
type Pim_DefaultContext_Ipv6_Maximum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum packet queue size in bytes. The type is interface{} with range:
    // 0..2147483648. Units are byte.
    GlobalLowPriorityPacketQueue interface{}

    // Maximum packet queue size in bytes. The type is interface{} with range:
    // 0..2147483648. Units are byte.
    GlobalHighPriorityPacketQueue interface{}

    // Override default global maximum and threshold for PIM group mapping ranges
    // from BSR.
    BsrGlobalGroupMappings Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings

    // Override default maximum for number of routes.
    GlobalRoutes Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes

    // Maximum for number of group mappings from autorp mapping agent.
    GlobalGroupMappingsAutoRp Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp

    // Override default global maximum and threshold for C-RP set in BSR.
    BsrGlobalCandidateRpCache Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache

    // Override default maximum for number of sparse-mode source registers.
    GlobalRegisterStates Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates

    // Override default maximum for number of route-interfaces.
    GlobalRouteInterfaces Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces

    // Override default maximum for number of group mappings from autorp mapping
    // agent.
    GroupMappingsAutoRp Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp

    // Override default maximum and threshold for number of group mappings from
    // BSR.
    BsrGroupMappings Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings

    // Override default maximum for number of sparse-mode source registers.
    RegisterStates Pim_DefaultContext_Ipv6_Maximum_RegisterStates

    // Override default maximum for number of route-interfaces.
    RouteInterfaces Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces

    // Override default maximum and threshold for BSR C-RP cache setting.
    BsrCandidateRpCache Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache

    // Override default maximum for number of routes.
    Routes Pim_DefaultContext_Ipv6_Maximum_Routes
}

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetEntityData() *types.CommonEntityData {
    maximum.EntityData.YFilter = maximum.YFilter
    maximum.EntityData.YangName = "maximum"
    maximum.EntityData.BundleName = "cisco_ios_xr"
    maximum.EntityData.ParentYangName = "ipv6"
    maximum.EntityData.SegmentPath = "maximum"
    maximum.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + maximum.EntityData.SegmentPath
    maximum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximum.EntityData.Children = types.NewOrderedMap()
    maximum.EntityData.Children.Append("bsr-global-group-mappings", types.YChild{"BsrGlobalGroupMappings", &maximum.BsrGlobalGroupMappings})
    maximum.EntityData.Children.Append("global-routes", types.YChild{"GlobalRoutes", &maximum.GlobalRoutes})
    maximum.EntityData.Children.Append("global-group-mappings-auto-rp", types.YChild{"GlobalGroupMappingsAutoRp", &maximum.GlobalGroupMappingsAutoRp})
    maximum.EntityData.Children.Append("bsr-global-candidate-rp-cache", types.YChild{"BsrGlobalCandidateRpCache", &maximum.BsrGlobalCandidateRpCache})
    maximum.EntityData.Children.Append("global-register-states", types.YChild{"GlobalRegisterStates", &maximum.GlobalRegisterStates})
    maximum.EntityData.Children.Append("global-route-interfaces", types.YChild{"GlobalRouteInterfaces", &maximum.GlobalRouteInterfaces})
    maximum.EntityData.Children.Append("group-mappings-auto-rp", types.YChild{"GroupMappingsAutoRp", &maximum.GroupMappingsAutoRp})
    maximum.EntityData.Children.Append("bsr-group-mappings", types.YChild{"BsrGroupMappings", &maximum.BsrGroupMappings})
    maximum.EntityData.Children.Append("register-states", types.YChild{"RegisterStates", &maximum.RegisterStates})
    maximum.EntityData.Children.Append("route-interfaces", types.YChild{"RouteInterfaces", &maximum.RouteInterfaces})
    maximum.EntityData.Children.Append("bsr-candidate-rp-cache", types.YChild{"BsrCandidateRpCache", &maximum.BsrCandidateRpCache})
    maximum.EntityData.Children.Append("routes", types.YChild{"Routes", &maximum.Routes})
    maximum.EntityData.Leafs = types.NewOrderedMap()
    maximum.EntityData.Leafs.Append("global-low-priority-packet-queue", types.YLeaf{"GlobalLowPriorityPacketQueue", maximum.GlobalLowPriorityPacketQueue})
    maximum.EntityData.Leafs.Append("global-high-priority-packet-queue", types.YLeaf{"GlobalHighPriorityPacketQueue", maximum.GlobalHighPriorityPacketQueue})

    maximum.EntityData.YListKeys = []string {}

    return &(maximum.EntityData)
}

// Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings
// Override default global maximum and threshold
// for PIM group mapping ranges from BSR
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Global Maximum number of PIM group mapping ranges from BSR. The type is
    // interface{} with range: 1..10000. This attribute is mandatory.
    BsrMaximumGlobalGroupMappings interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 500.
    WarningThreshold interface{}
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetEntityData() *types.CommonEntityData {
    bsrGlobalGroupMappings.EntityData.YFilter = bsrGlobalGroupMappings.YFilter
    bsrGlobalGroupMappings.EntityData.YangName = "bsr-global-group-mappings"
    bsrGlobalGroupMappings.EntityData.BundleName = "cisco_ios_xr"
    bsrGlobalGroupMappings.EntityData.ParentYangName = "maximum"
    bsrGlobalGroupMappings.EntityData.SegmentPath = "bsr-global-group-mappings"
    bsrGlobalGroupMappings.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/maximum/" + bsrGlobalGroupMappings.EntityData.SegmentPath
    bsrGlobalGroupMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsrGlobalGroupMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsrGlobalGroupMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsrGlobalGroupMappings.EntityData.Children = types.NewOrderedMap()
    bsrGlobalGroupMappings.EntityData.Leafs = types.NewOrderedMap()
    bsrGlobalGroupMappings.EntityData.Leafs.Append("bsr-maximum-global-group-mappings", types.YLeaf{"BsrMaximumGlobalGroupMappings", bsrGlobalGroupMappings.BsrMaximumGlobalGroupMappings})
    bsrGlobalGroupMappings.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", bsrGlobalGroupMappings.WarningThreshold})

    bsrGlobalGroupMappings.EntityData.YListKeys = []string {}

    return &(bsrGlobalGroupMappings.EntityData)
}

// Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes
// Override default maximum for number of routes
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM routes. The type is interface{} with range:
    // 1..200000. This attribute is mandatory.
    MaximumRoutes interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..200000. The default value is 100000.
    WarningThreshold interface{}
}

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetEntityData() *types.CommonEntityData {
    globalRoutes.EntityData.YFilter = globalRoutes.YFilter
    globalRoutes.EntityData.YangName = "global-routes"
    globalRoutes.EntityData.BundleName = "cisco_ios_xr"
    globalRoutes.EntityData.ParentYangName = "maximum"
    globalRoutes.EntityData.SegmentPath = "global-routes"
    globalRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/maximum/" + globalRoutes.EntityData.SegmentPath
    globalRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalRoutes.EntityData.Children = types.NewOrderedMap()
    globalRoutes.EntityData.Leafs = types.NewOrderedMap()
    globalRoutes.EntityData.Leafs.Append("maximum-routes", types.YLeaf{"MaximumRoutes", globalRoutes.MaximumRoutes})
    globalRoutes.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", globalRoutes.WarningThreshold})

    globalRoutes.EntityData.YListKeys = []string {}

    return &(globalRoutes.EntityData)
}

// Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp
// Maximum for number of group mappings from
// autorp mapping agent
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM group mappings from autorp. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    MaximumGlobalGroupRangesAutoRp interface{}

    // Warning threshold number of PIM group mappings from autorp. The type is
    // interface{} with range: 1..10000. The default value is 450.
    ThresholdGlobalGroupRangesAutoRp interface{}
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetEntityData() *types.CommonEntityData {
    globalGroupMappingsAutoRp.EntityData.YFilter = globalGroupMappingsAutoRp.YFilter
    globalGroupMappingsAutoRp.EntityData.YangName = "global-group-mappings-auto-rp"
    globalGroupMappingsAutoRp.EntityData.BundleName = "cisco_ios_xr"
    globalGroupMappingsAutoRp.EntityData.ParentYangName = "maximum"
    globalGroupMappingsAutoRp.EntityData.SegmentPath = "global-group-mappings-auto-rp"
    globalGroupMappingsAutoRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/maximum/" + globalGroupMappingsAutoRp.EntityData.SegmentPath
    globalGroupMappingsAutoRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalGroupMappingsAutoRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalGroupMappingsAutoRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalGroupMappingsAutoRp.EntityData.Children = types.NewOrderedMap()
    globalGroupMappingsAutoRp.EntityData.Leafs = types.NewOrderedMap()
    globalGroupMappingsAutoRp.EntityData.Leafs.Append("maximum-global-group-ranges-auto-rp", types.YLeaf{"MaximumGlobalGroupRangesAutoRp", globalGroupMappingsAutoRp.MaximumGlobalGroupRangesAutoRp})
    globalGroupMappingsAutoRp.EntityData.Leafs.Append("threshold-global-group-ranges-auto-rp", types.YLeaf{"ThresholdGlobalGroupRangesAutoRp", globalGroupMappingsAutoRp.ThresholdGlobalGroupRangesAutoRp})

    globalGroupMappingsAutoRp.EntityData.YListKeys = []string {}

    return &(globalGroupMappingsAutoRp.EntityData)
}

// Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache
// Override default global maximum and threshold
// for C-RP set in BSR
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Global Maximum number of PIM C-RP Sets from BSR. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    BsrMaximumGlobalCandidateRpCache interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 100.
    WarningThreshold interface{}
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetEntityData() *types.CommonEntityData {
    bsrGlobalCandidateRpCache.EntityData.YFilter = bsrGlobalCandidateRpCache.YFilter
    bsrGlobalCandidateRpCache.EntityData.YangName = "bsr-global-candidate-rp-cache"
    bsrGlobalCandidateRpCache.EntityData.BundleName = "cisco_ios_xr"
    bsrGlobalCandidateRpCache.EntityData.ParentYangName = "maximum"
    bsrGlobalCandidateRpCache.EntityData.SegmentPath = "bsr-global-candidate-rp-cache"
    bsrGlobalCandidateRpCache.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/maximum/" + bsrGlobalCandidateRpCache.EntityData.SegmentPath
    bsrGlobalCandidateRpCache.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsrGlobalCandidateRpCache.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsrGlobalCandidateRpCache.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsrGlobalCandidateRpCache.EntityData.Children = types.NewOrderedMap()
    bsrGlobalCandidateRpCache.EntityData.Leafs = types.NewOrderedMap()
    bsrGlobalCandidateRpCache.EntityData.Leafs.Append("bsr-maximum-global-candidate-rp-cache", types.YLeaf{"BsrMaximumGlobalCandidateRpCache", bsrGlobalCandidateRpCache.BsrMaximumGlobalCandidateRpCache})
    bsrGlobalCandidateRpCache.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", bsrGlobalCandidateRpCache.WarningThreshold})

    bsrGlobalCandidateRpCache.EntityData.YListKeys = []string {}

    return &(bsrGlobalCandidateRpCache.EntityData)
}

// Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates
// Override default maximum for number of
// sparse-mode source registers
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM Sparse-Mode register states. The type is interface{}
    // with range: 0..75000. This attribute is mandatory.
    MaximumRegisterStates interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 0..75000. The default value is 20000.
    WarningThreshold interface{}
}

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetEntityData() *types.CommonEntityData {
    globalRegisterStates.EntityData.YFilter = globalRegisterStates.YFilter
    globalRegisterStates.EntityData.YangName = "global-register-states"
    globalRegisterStates.EntityData.BundleName = "cisco_ios_xr"
    globalRegisterStates.EntityData.ParentYangName = "maximum"
    globalRegisterStates.EntityData.SegmentPath = "global-register-states"
    globalRegisterStates.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/maximum/" + globalRegisterStates.EntityData.SegmentPath
    globalRegisterStates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalRegisterStates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalRegisterStates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalRegisterStates.EntityData.Children = types.NewOrderedMap()
    globalRegisterStates.EntityData.Leafs = types.NewOrderedMap()
    globalRegisterStates.EntityData.Leafs.Append("maximum-register-states", types.YLeaf{"MaximumRegisterStates", globalRegisterStates.MaximumRegisterStates})
    globalRegisterStates.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", globalRegisterStates.WarningThreshold})

    globalRegisterStates.EntityData.YListKeys = []string {}

    return &(globalRegisterStates.EntityData)
}

// Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces
// Override default maximum for number of
// route-interfaces
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM route-interfaces. The type is interface{} with range:
    // 1..1100000. This attribute is mandatory.
    MaximumRouteInterfaces interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000. The default value is 300000.
    WarningThreshold interface{}
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetEntityData() *types.CommonEntityData {
    globalRouteInterfaces.EntityData.YFilter = globalRouteInterfaces.YFilter
    globalRouteInterfaces.EntityData.YangName = "global-route-interfaces"
    globalRouteInterfaces.EntityData.BundleName = "cisco_ios_xr"
    globalRouteInterfaces.EntityData.ParentYangName = "maximum"
    globalRouteInterfaces.EntityData.SegmentPath = "global-route-interfaces"
    globalRouteInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/maximum/" + globalRouteInterfaces.EntityData.SegmentPath
    globalRouteInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalRouteInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalRouteInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalRouteInterfaces.EntityData.Children = types.NewOrderedMap()
    globalRouteInterfaces.EntityData.Leafs = types.NewOrderedMap()
    globalRouteInterfaces.EntityData.Leafs.Append("maximum-route-interfaces", types.YLeaf{"MaximumRouteInterfaces", globalRouteInterfaces.MaximumRouteInterfaces})
    globalRouteInterfaces.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", globalRouteInterfaces.WarningThreshold})

    globalRouteInterfaces.EntityData.YListKeys = []string {}

    return &(globalRouteInterfaces.EntityData)
}

// Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp
// Override default maximum for number of group
// mappings from autorp mapping agent
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM group mappings from autorp. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    MaximumGroupRangesAutoRp interface{}

    // Warning threshold number of PIM group mappings from autorp. The type is
    // interface{} with range: 1..10000. The default value is 450.
    ThresholdGroupRangesAutoRp interface{}
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetEntityData() *types.CommonEntityData {
    groupMappingsAutoRp.EntityData.YFilter = groupMappingsAutoRp.YFilter
    groupMappingsAutoRp.EntityData.YangName = "group-mappings-auto-rp"
    groupMappingsAutoRp.EntityData.BundleName = "cisco_ios_xr"
    groupMappingsAutoRp.EntityData.ParentYangName = "maximum"
    groupMappingsAutoRp.EntityData.SegmentPath = "group-mappings-auto-rp"
    groupMappingsAutoRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/maximum/" + groupMappingsAutoRp.EntityData.SegmentPath
    groupMappingsAutoRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupMappingsAutoRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupMappingsAutoRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupMappingsAutoRp.EntityData.Children = types.NewOrderedMap()
    groupMappingsAutoRp.EntityData.Leafs = types.NewOrderedMap()
    groupMappingsAutoRp.EntityData.Leafs.Append("maximum-group-ranges-auto-rp", types.YLeaf{"MaximumGroupRangesAutoRp", groupMappingsAutoRp.MaximumGroupRangesAutoRp})
    groupMappingsAutoRp.EntityData.Leafs.Append("threshold-group-ranges-auto-rp", types.YLeaf{"ThresholdGroupRangesAutoRp", groupMappingsAutoRp.ThresholdGroupRangesAutoRp})

    groupMappingsAutoRp.EntityData.YListKeys = []string {}

    return &(groupMappingsAutoRp.EntityData)
}

// Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings
// Override default maximum and threshold for
// number of group mappings from BSR
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM group mappings from BSR. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumGroupRanges interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 500.
    WarningThreshold interface{}
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetEntityData() *types.CommonEntityData {
    bsrGroupMappings.EntityData.YFilter = bsrGroupMappings.YFilter
    bsrGroupMappings.EntityData.YangName = "bsr-group-mappings"
    bsrGroupMappings.EntityData.BundleName = "cisco_ios_xr"
    bsrGroupMappings.EntityData.ParentYangName = "maximum"
    bsrGroupMappings.EntityData.SegmentPath = "bsr-group-mappings"
    bsrGroupMappings.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/maximum/" + bsrGroupMappings.EntityData.SegmentPath
    bsrGroupMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsrGroupMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsrGroupMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsrGroupMappings.EntityData.Children = types.NewOrderedMap()
    bsrGroupMappings.EntityData.Leafs = types.NewOrderedMap()
    bsrGroupMappings.EntityData.Leafs.Append("bsr-maximum-group-ranges", types.YLeaf{"BsrMaximumGroupRanges", bsrGroupMappings.BsrMaximumGroupRanges})
    bsrGroupMappings.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", bsrGroupMappings.WarningThreshold})

    bsrGroupMappings.EntityData.YListKeys = []string {}

    return &(bsrGroupMappings.EntityData)
}

// Pim_DefaultContext_Ipv6_Maximum_RegisterStates
// Override default maximum for number of
// sparse-mode source registers
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_RegisterStates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM Sparse-Mode register states. The type is interface{}
    // with range: 0..75000. This attribute is mandatory.
    MaximumRegisterStates interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 0..75000. The default value is 20000.
    WarningThreshold interface{}
}

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetEntityData() *types.CommonEntityData {
    registerStates.EntityData.YFilter = registerStates.YFilter
    registerStates.EntityData.YangName = "register-states"
    registerStates.EntityData.BundleName = "cisco_ios_xr"
    registerStates.EntityData.ParentYangName = "maximum"
    registerStates.EntityData.SegmentPath = "register-states"
    registerStates.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/maximum/" + registerStates.EntityData.SegmentPath
    registerStates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registerStates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registerStates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registerStates.EntityData.Children = types.NewOrderedMap()
    registerStates.EntityData.Leafs = types.NewOrderedMap()
    registerStates.EntityData.Leafs.Append("maximum-register-states", types.YLeaf{"MaximumRegisterStates", registerStates.MaximumRegisterStates})
    registerStates.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", registerStates.WarningThreshold})

    registerStates.EntityData.YListKeys = []string {}

    return &(registerStates.EntityData)
}

// Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces
// Override default maximum for number of
// route-interfaces
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM route-interfaces. The type is interface{} with range:
    // 1..1100000. This attribute is mandatory.
    MaximumRouteInterfaces interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000. The default value is 300000.
    WarningThreshold interface{}
}

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetEntityData() *types.CommonEntityData {
    routeInterfaces.EntityData.YFilter = routeInterfaces.YFilter
    routeInterfaces.EntityData.YangName = "route-interfaces"
    routeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    routeInterfaces.EntityData.ParentYangName = "maximum"
    routeInterfaces.EntityData.SegmentPath = "route-interfaces"
    routeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/maximum/" + routeInterfaces.EntityData.SegmentPath
    routeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeInterfaces.EntityData.Children = types.NewOrderedMap()
    routeInterfaces.EntityData.Leafs = types.NewOrderedMap()
    routeInterfaces.EntityData.Leafs.Append("maximum-route-interfaces", types.YLeaf{"MaximumRouteInterfaces", routeInterfaces.MaximumRouteInterfaces})
    routeInterfaces.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", routeInterfaces.WarningThreshold})

    routeInterfaces.EntityData.YListKeys = []string {}

    return &(routeInterfaces.EntityData)
}

// Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache
// Override default maximum and threshold for BSR
// C-RP cache setting
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of BSR C-RP cache setting. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumCandidateRpCache interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 100.
    WarningThreshold interface{}
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetEntityData() *types.CommonEntityData {
    bsrCandidateRpCache.EntityData.YFilter = bsrCandidateRpCache.YFilter
    bsrCandidateRpCache.EntityData.YangName = "bsr-candidate-rp-cache"
    bsrCandidateRpCache.EntityData.BundleName = "cisco_ios_xr"
    bsrCandidateRpCache.EntityData.ParentYangName = "maximum"
    bsrCandidateRpCache.EntityData.SegmentPath = "bsr-candidate-rp-cache"
    bsrCandidateRpCache.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/maximum/" + bsrCandidateRpCache.EntityData.SegmentPath
    bsrCandidateRpCache.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsrCandidateRpCache.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsrCandidateRpCache.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsrCandidateRpCache.EntityData.Children = types.NewOrderedMap()
    bsrCandidateRpCache.EntityData.Leafs = types.NewOrderedMap()
    bsrCandidateRpCache.EntityData.Leafs.Append("bsr-maximum-candidate-rp-cache", types.YLeaf{"BsrMaximumCandidateRpCache", bsrCandidateRpCache.BsrMaximumCandidateRpCache})
    bsrCandidateRpCache.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", bsrCandidateRpCache.WarningThreshold})

    bsrCandidateRpCache.EntityData.YListKeys = []string {}

    return &(bsrCandidateRpCache.EntityData)
}

// Pim_DefaultContext_Ipv6_Maximum_Routes
// Override default maximum for number of routes
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM routes. The type is interface{} with range:
    // 1..200000. This attribute is mandatory.
    MaximumRoutes interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..200000. The default value is 100000.
    WarningThreshold interface{}
}

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "maximum"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/maximum/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Leafs = types.NewOrderedMap()
    routes.EntityData.Leafs.Append("maximum-routes", types.YLeaf{"MaximumRoutes", routes.MaximumRoutes})
    routes.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", routes.WarningThreshold})

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// Pim_DefaultContext_Ipv6_Ssm
// Configure IP Multicast SSM
type Pim_DefaultContext_Ipv6_Ssm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE if SSM is disabled on this router. The type is bool. The default value
    // is false.
    Disable interface{}

    // Access list of groups enabled with SSM. The type is string with length:
    // 1..64.
    Range interface{}
}

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetEntityData() *types.CommonEntityData {
    ssm.EntityData.YFilter = ssm.YFilter
    ssm.EntityData.YangName = "ssm"
    ssm.EntityData.BundleName = "cisco_ios_xr"
    ssm.EntityData.ParentYangName = "ipv6"
    ssm.EntityData.SegmentPath = "ssm"
    ssm.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + ssm.EntityData.SegmentPath
    ssm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssm.EntityData.Children = types.NewOrderedMap()
    ssm.EntityData.Leafs = types.NewOrderedMap()
    ssm.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", ssm.Disable})
    ssm.EntityData.Leafs.Append("range", types.YLeaf{"Range", ssm.Range})

    ssm.EntityData.YListKeys = []string {}

    return &(ssm.EntityData)
}

// Pim_DefaultContext_Ipv6_BidirRpAddresses
// Configure Bidirectional PIM Rendezvous Point
type Pim_DefaultContext_Ipv6_BidirRpAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress.
    BidirRpAddress []*Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetEntityData() *types.CommonEntityData {
    bidirRpAddresses.EntityData.YFilter = bidirRpAddresses.YFilter
    bidirRpAddresses.EntityData.YangName = "bidir-rp-addresses"
    bidirRpAddresses.EntityData.BundleName = "cisco_ios_xr"
    bidirRpAddresses.EntityData.ParentYangName = "ipv6"
    bidirRpAddresses.EntityData.SegmentPath = "bidir-rp-addresses"
    bidirRpAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + bidirRpAddresses.EntityData.SegmentPath
    bidirRpAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bidirRpAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bidirRpAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bidirRpAddresses.EntityData.Children = types.NewOrderedMap()
    bidirRpAddresses.EntityData.Children.Append("bidir-rp-address", types.YChild{"BidirRpAddress", nil})
    for i := range bidirRpAddresses.BidirRpAddress {
        bidirRpAddresses.EntityData.Children.Append(types.GetSegmentPath(bidirRpAddresses.BidirRpAddress[i]), types.YChild{"BidirRpAddress", bidirRpAddresses.BidirRpAddress[i]})
    }
    bidirRpAddresses.EntityData.Leafs = types.NewOrderedMap()

    bidirRpAddresses.EntityData.YListKeys = []string {}

    return &(bidirRpAddresses.EntityData)
}

// Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress
// Address of the Rendezvous Point
type Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetEntityData() *types.CommonEntityData {
    bidirRpAddress.EntityData.YFilter = bidirRpAddress.YFilter
    bidirRpAddress.EntityData.YangName = "bidir-rp-address"
    bidirRpAddress.EntityData.BundleName = "cisco_ios_xr"
    bidirRpAddress.EntityData.ParentYangName = "bidir-rp-addresses"
    bidirRpAddress.EntityData.SegmentPath = "bidir-rp-address" + types.AddKeyToken(bidirRpAddress.RpAddress, "rp-address")
    bidirRpAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/bidir-rp-addresses/" + bidirRpAddress.EntityData.SegmentPath
    bidirRpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bidirRpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bidirRpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bidirRpAddress.EntityData.Children = types.NewOrderedMap()
    bidirRpAddress.EntityData.Leafs = types.NewOrderedMap()
    bidirRpAddress.EntityData.Leafs.Append("rp-address", types.YLeaf{"RpAddress", bidirRpAddress.RpAddress})
    bidirRpAddress.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", bidirRpAddress.AccessListName})
    bidirRpAddress.EntityData.Leafs.Append("auto-rp-override", types.YLeaf{"AutoRpOverride", bidirRpAddress.AutoRpOverride})

    bidirRpAddress.EntityData.YListKeys = []string {"RpAddress"}

    return &(bidirRpAddress.EntityData)
}

// Pim_DefaultContext_Ipv6_Bsr
// PIM BSR configuration
type Pim_DefaultContext_Ipv6_Bsr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PIM Candidate BSR configuration.
    CandidateBsr Pim_DefaultContext_Ipv6_Bsr_CandidateBsr

    // PIM RP configuration.
    CandidateRps Pim_DefaultContext_Ipv6_Bsr_CandidateRps
}

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetEntityData() *types.CommonEntityData {
    bsr.EntityData.YFilter = bsr.YFilter
    bsr.EntityData.YangName = "bsr"
    bsr.EntityData.BundleName = "cisco_ios_xr"
    bsr.EntityData.ParentYangName = "ipv6"
    bsr.EntityData.SegmentPath = "bsr"
    bsr.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + bsr.EntityData.SegmentPath
    bsr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsr.EntityData.Children = types.NewOrderedMap()
    bsr.EntityData.Children.Append("candidate-bsr", types.YChild{"CandidateBsr", &bsr.CandidateBsr})
    bsr.EntityData.Children.Append("candidate-rps", types.YChild{"CandidateRps", &bsr.CandidateRps})
    bsr.EntityData.Leafs = types.NewOrderedMap()

    bsr.EntityData.YListKeys = []string {}

    return &(bsr.EntityData)
}

// Pim_DefaultContext_Ipv6_Bsr_CandidateBsr
// PIM Candidate BSR configuration
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Bsr_CandidateBsr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // BSR Address configured. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}

    // Hash Mask Length for this candidate BSR. The type is interface{} with
    // range: 0..128. The default value is 126.
    PrefixLength interface{}

    // Priority of the Candidate BSR. The type is interface{} with range: 1..255.
    // The default value is 1.
    Priority interface{}
}

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetEntityData() *types.CommonEntityData {
    candidateBsr.EntityData.YFilter = candidateBsr.YFilter
    candidateBsr.EntityData.YangName = "candidate-bsr"
    candidateBsr.EntityData.BundleName = "cisco_ios_xr"
    candidateBsr.EntityData.ParentYangName = "bsr"
    candidateBsr.EntityData.SegmentPath = "candidate-bsr"
    candidateBsr.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/bsr/" + candidateBsr.EntityData.SegmentPath
    candidateBsr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateBsr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateBsr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateBsr.EntityData.Children = types.NewOrderedMap()
    candidateBsr.EntityData.Leafs = types.NewOrderedMap()
    candidateBsr.EntityData.Leafs.Append("address", types.YLeaf{"Address", candidateBsr.Address})
    candidateBsr.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", candidateBsr.PrefixLength})
    candidateBsr.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", candidateBsr.Priority})

    candidateBsr.EntityData.YListKeys = []string {}

    return &(candidateBsr.EntityData)
}

// Pim_DefaultContext_Ipv6_Bsr_CandidateRps
// PIM RP configuration
type Pim_DefaultContext_Ipv6_Bsr_CandidateRps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of PIM SM BSR Candidate-RP. The type is slice of
    // Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp.
    CandidateRp []*Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp
}

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetEntityData() *types.CommonEntityData {
    candidateRps.EntityData.YFilter = candidateRps.YFilter
    candidateRps.EntityData.YangName = "candidate-rps"
    candidateRps.EntityData.BundleName = "cisco_ios_xr"
    candidateRps.EntityData.ParentYangName = "bsr"
    candidateRps.EntityData.SegmentPath = "candidate-rps"
    candidateRps.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/bsr/" + candidateRps.EntityData.SegmentPath
    candidateRps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateRps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateRps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateRps.EntityData.Children = types.NewOrderedMap()
    candidateRps.EntityData.Children.Append("candidate-rp", types.YChild{"CandidateRp", nil})
    for i := range candidateRps.CandidateRp {
        candidateRps.EntityData.Children.Append(types.GetSegmentPath(candidateRps.CandidateRp[i]), types.YChild{"CandidateRp", candidateRps.CandidateRp[i]})
    }
    candidateRps.EntityData.Leafs = types.NewOrderedMap()

    candidateRps.EntityData.YListKeys = []string {}

    return &(candidateRps.EntityData)
}

// Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp
// Address of PIM SM BSR Candidate-RP
type Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address of Candidate-RP. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. SM or Bidir. The type is PimProtocolMode.
    Mode interface{}

    // Access-list specifying the group range for the Candidate-RP. The type is
    // string with length: 1..64.
    GroupList interface{}

    // Priority of the CRP. The type is interface{} with range: 1..255. The
    // default value is 192.
    Priority interface{}

    // Advertisement interval. The type is interface{} with range: 30..600. The
    // default value is 60.
    Interval interface{}
}

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetEntityData() *types.CommonEntityData {
    candidateRp.EntityData.YFilter = candidateRp.YFilter
    candidateRp.EntityData.YangName = "candidate-rp"
    candidateRp.EntityData.BundleName = "cisco_ios_xr"
    candidateRp.EntityData.ParentYangName = "candidate-rps"
    candidateRp.EntityData.SegmentPath = "candidate-rp" + types.AddKeyToken(candidateRp.Address, "address") + types.AddKeyToken(candidateRp.Mode, "mode")
    candidateRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/bsr/candidate-rps/" + candidateRp.EntityData.SegmentPath
    candidateRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateRp.EntityData.Children = types.NewOrderedMap()
    candidateRp.EntityData.Leafs = types.NewOrderedMap()
    candidateRp.EntityData.Leafs.Append("address", types.YLeaf{"Address", candidateRp.Address})
    candidateRp.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", candidateRp.Mode})
    candidateRp.EntityData.Leafs.Append("group-list", types.YLeaf{"GroupList", candidateRp.GroupList})
    candidateRp.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", candidateRp.Priority})
    candidateRp.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", candidateRp.Interval})

    candidateRp.EntityData.YListKeys = []string {"Address", "Mode"}

    return &(candidateRp.EntityData)
}

// Pim_DefaultContext_Ipv6_AllowRp
// Enable allow-rp filtering for SM joins
// This type is a presence type.
type Pim_DefaultContext_Ipv6_AllowRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Access-list specifiying applicable RPs. The type is string with length:
    // 1..64.
    RpListName interface{}

    // Access-list specifiying applicable groups. The type is string with length:
    // 1..64.
    GroupListName interface{}
}

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetEntityData() *types.CommonEntityData {
    allowRp.EntityData.YFilter = allowRp.YFilter
    allowRp.EntityData.YangName = "allow-rp"
    allowRp.EntityData.BundleName = "cisco_ios_xr"
    allowRp.EntityData.ParentYangName = "ipv6"
    allowRp.EntityData.SegmentPath = "allow-rp"
    allowRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + allowRp.EntityData.SegmentPath
    allowRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowRp.EntityData.Children = types.NewOrderedMap()
    allowRp.EntityData.Leafs = types.NewOrderedMap()
    allowRp.EntityData.Leafs.Append("rp-list-name", types.YLeaf{"RpListName", allowRp.RpListName})
    allowRp.EntityData.Leafs.Append("group-list-name", types.YLeaf{"GroupListName", allowRp.GroupListName})

    allowRp.EntityData.YListKeys = []string {}

    return &(allowRp.EntityData)
}

// Pim_DefaultContext_Ipv6_EmbeddedRpAddresses
// Set Embedded RP processing support
type Pim_DefaultContext_Ipv6_EmbeddedRpAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set Embedded RP processing support. The type is slice of
    // Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress.
    EmbeddedRpAddress []*Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress
}

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetEntityData() *types.CommonEntityData {
    embeddedRpAddresses.EntityData.YFilter = embeddedRpAddresses.YFilter
    embeddedRpAddresses.EntityData.YangName = "embedded-rp-addresses"
    embeddedRpAddresses.EntityData.BundleName = "cisco_ios_xr"
    embeddedRpAddresses.EntityData.ParentYangName = "ipv6"
    embeddedRpAddresses.EntityData.SegmentPath = "embedded-rp-addresses"
    embeddedRpAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + embeddedRpAddresses.EntityData.SegmentPath
    embeddedRpAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    embeddedRpAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    embeddedRpAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    embeddedRpAddresses.EntityData.Children = types.NewOrderedMap()
    embeddedRpAddresses.EntityData.Children.Append("embedded-rp-address", types.YChild{"EmbeddedRpAddress", nil})
    for i := range embeddedRpAddresses.EmbeddedRpAddress {
        embeddedRpAddresses.EntityData.Children.Append(types.GetSegmentPath(embeddedRpAddresses.EmbeddedRpAddress[i]), types.YChild{"EmbeddedRpAddress", embeddedRpAddresses.EmbeddedRpAddress[i]})
    }
    embeddedRpAddresses.EntityData.Leafs = types.NewOrderedMap()

    embeddedRpAddresses.EntityData.YListKeys = []string {}

    return &(embeddedRpAddresses.EntityData)
}

// Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress
// Set Embedded RP processing support
type Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. RP address of the Rendezvous Point. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a given RP. The type is string
    // with length: 1..64. This attribute is mandatory.
    AccessListName interface{}
}

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetEntityData() *types.CommonEntityData {
    embeddedRpAddress.EntityData.YFilter = embeddedRpAddress.YFilter
    embeddedRpAddress.EntityData.YangName = "embedded-rp-address"
    embeddedRpAddress.EntityData.BundleName = "cisco_ios_xr"
    embeddedRpAddress.EntityData.ParentYangName = "embedded-rp-addresses"
    embeddedRpAddress.EntityData.SegmentPath = "embedded-rp-address" + types.AddKeyToken(embeddedRpAddress.RpAddress, "rp-address")
    embeddedRpAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/embedded-rp-addresses/" + embeddedRpAddress.EntityData.SegmentPath
    embeddedRpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    embeddedRpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    embeddedRpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    embeddedRpAddress.EntityData.Children = types.NewOrderedMap()
    embeddedRpAddress.EntityData.Leafs = types.NewOrderedMap()
    embeddedRpAddress.EntityData.Leafs.Append("rp-address", types.YLeaf{"RpAddress", embeddedRpAddress.RpAddress})
    embeddedRpAddress.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", embeddedRpAddress.AccessListName})

    embeddedRpAddress.EntityData.YListKeys = []string {"RpAddress"}

    return &(embeddedRpAddress.EntityData)
}

// Pim_DefaultContext_Ipv6_Convergence
// Configure convergence parameters
type Pim_DefaultContext_Ipv6_Convergence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampen first join if RPF path is through one of the downstream neighbor.
    // The type is interface{} with range: 0..15. Units are second.
    RpfConflictJoinDelay interface{}

    // Delay prunes if route join state transitions to not-joined on link down.
    // The type is interface{} with range: 0..60. Units are second.
    LinkDownPruneDelay interface{}
}

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetEntityData() *types.CommonEntityData {
    convergence.EntityData.YFilter = convergence.YFilter
    convergence.EntityData.YangName = "convergence"
    convergence.EntityData.BundleName = "cisco_ios_xr"
    convergence.EntityData.ParentYangName = "ipv6"
    convergence.EntityData.SegmentPath = "convergence"
    convergence.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv6/" + convergence.EntityData.SegmentPath
    convergence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    convergence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    convergence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    convergence.EntityData.Children = types.NewOrderedMap()
    convergence.EntityData.Leafs = types.NewOrderedMap()
    convergence.EntityData.Leafs.Append("rpf-conflict-join-delay", types.YLeaf{"RpfConflictJoinDelay", convergence.RpfConflictJoinDelay})
    convergence.EntityData.Leafs.Append("link-down-prune-delay", types.YLeaf{"LinkDownPruneDelay", convergence.LinkDownPruneDelay})

    convergence.EntityData.YListKeys = []string {}

    return &(convergence.EntityData)
}

// Pim_DefaultContext_Ipv4
// IPV4 commands
type Pim_DefaultContext_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable PIM neighbor checking when receiving PIM messages. The type is
    // interface{}.
    NeighborCheckOnReceive interface{}

    // Generate registers compatible with older IOS versions. The type is
    // interface{}.
    OldRegisterChecksum interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Configure threshold of infinity for switching to SPT on last-hop. The type
    // is string.
    SptThresholdInfinity interface{}

    // PIM neighbor state change logging is turned on if configured. The type is
    // interface{}.
    LogNeighborChanges interface{}

    // Source address to use for register messages. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    RegisterSource interface{}

    // Access-list which specifies unauthorized sources. The type is string with
    // length: 1..64.
    AcceptRegister interface{}

    // Suppress prunes triggered as a result of RPF changes. The type is
    // interface{}.
    SuppressRpfPrunes interface{}

    // Allow SSM ranges to be overridden by more specific ranges. The type is
    // interface{}.
    SsmAllowOverride interface{}

    // Enable equal-cost multipath routing. The type is PimMultipath.
    Multipath interface{}

    // Configure static RP deny range. The type is string with length: 1..64.
    RpStaticDeny interface{}

    // Suppress data registers after initial state setup. The type is interface{}.
    SuppressDataRegisters interface{}

    // Enable PIM neighbor checking when sending join-prunes. The type is
    // interface{}.
    NeighborCheckOnSend interface{}

    // Disable Rendezvous Point discovery through the AutoRP protocol. The type is
    // interface{}.
    AutoRpDisable interface{}

    // Configure RPF-redirect feature.
    RpfRedirect Pim_DefaultContext_Ipv4_RpfRedirect

    // Interface-level Configuration.
    Interfaces Pim_DefaultContext_Ipv4_Interfaces

    // Configure Candidate-RPs.
    AutoRpCandidateRps Pim_DefaultContext_Ipv4_AutoRpCandidateRps

    // Configure AutoRP Mapping Agent.
    AutoRpMappingAgent Pim_DefaultContext_Ipv4_AutoRpMappingAgent

    // Configure Sparse-Mode Rendezvous Point.
    SparseModeRpAddresses Pim_DefaultContext_Ipv4_SparseModeRpAddresses

    // Inheritable defaults.
    InheritableDefaults Pim_DefaultContext_Ipv4_InheritableDefaults

    // Configure RPF options.
    Rpf Pim_DefaultContext_Ipv4_Rpf

    // Configure expiry timer for S,G routes.
    SgExpiryTimer Pim_DefaultContext_Ipv4_SgExpiryTimer

    // Enable PIM RPF Vector Proxy's.
    RpfVectorEnable Pim_DefaultContext_Ipv4_RpfVectorEnable

    // Configure Non-stop forwarding (NSF) options.
    Nsf Pim_DefaultContext_Ipv4_Nsf

    // Configure PIM State Limits.
    Maximum Pim_DefaultContext_Ipv4_Maximum

    // Configure IP Multicast SSM.
    Ssm Pim_DefaultContext_Ipv4_Ssm

    // Inject Explicit PIM RPF Vector Proxy's.
    Injects Pim_DefaultContext_Ipv4_Injects

    // Configure Bidirectional PIM Rendezvous Point.
    BidirRpAddresses Pim_DefaultContext_Ipv4_BidirRpAddresses

    // PIM BSR configuration.
    Bsr Pim_DefaultContext_Ipv4_Bsr

    // Multicast Only Fast Re-Route.
    Mofrr Pim_DefaultContext_Ipv4_Mofrr

    // Inject PIM RPF Vector Proxy's.
    Paths Pim_DefaultContext_Ipv4_Paths

    // Enable allow-rp filtering for SM joins.
    AllowRp Pim_DefaultContext_Ipv4_AllowRp

    // Configure convergence parameters.
    Convergence Pim_DefaultContext_Ipv4_Convergence
}

func (ipv4 *Pim_DefaultContext_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "default-context"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("rpf-redirect", types.YChild{"RpfRedirect", &ipv4.RpfRedirect})
    ipv4.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &ipv4.Interfaces})
    ipv4.EntityData.Children.Append("auto-rp-candidate-rps", types.YChild{"AutoRpCandidateRps", &ipv4.AutoRpCandidateRps})
    ipv4.EntityData.Children.Append("auto-rp-mapping-agent", types.YChild{"AutoRpMappingAgent", &ipv4.AutoRpMappingAgent})
    ipv4.EntityData.Children.Append("sparse-mode-rp-addresses", types.YChild{"SparseModeRpAddresses", &ipv4.SparseModeRpAddresses})
    ipv4.EntityData.Children.Append("inheritable-defaults", types.YChild{"InheritableDefaults", &ipv4.InheritableDefaults})
    ipv4.EntityData.Children.Append("rpf", types.YChild{"Rpf", &ipv4.Rpf})
    ipv4.EntityData.Children.Append("sg-expiry-timer", types.YChild{"SgExpiryTimer", &ipv4.SgExpiryTimer})
    ipv4.EntityData.Children.Append("rpf-vector-enable", types.YChild{"RpfVectorEnable", &ipv4.RpfVectorEnable})
    ipv4.EntityData.Children.Append("nsf", types.YChild{"Nsf", &ipv4.Nsf})
    ipv4.EntityData.Children.Append("maximum", types.YChild{"Maximum", &ipv4.Maximum})
    ipv4.EntityData.Children.Append("ssm", types.YChild{"Ssm", &ipv4.Ssm})
    ipv4.EntityData.Children.Append("injects", types.YChild{"Injects", &ipv4.Injects})
    ipv4.EntityData.Children.Append("bidir-rp-addresses", types.YChild{"BidirRpAddresses", &ipv4.BidirRpAddresses})
    ipv4.EntityData.Children.Append("bsr", types.YChild{"Bsr", &ipv4.Bsr})
    ipv4.EntityData.Children.Append("mofrr", types.YChild{"Mofrr", &ipv4.Mofrr})
    ipv4.EntityData.Children.Append("paths", types.YChild{"Paths", &ipv4.Paths})
    ipv4.EntityData.Children.Append("allow-rp", types.YChild{"AllowRp", &ipv4.AllowRp})
    ipv4.EntityData.Children.Append("convergence", types.YChild{"Convergence", &ipv4.Convergence})
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("neighbor-check-on-receive", types.YLeaf{"NeighborCheckOnReceive", ipv4.NeighborCheckOnReceive})
    ipv4.EntityData.Leafs.Append("old-register-checksum", types.YLeaf{"OldRegisterChecksum", ipv4.OldRegisterChecksum})
    ipv4.EntityData.Leafs.Append("neighbor-filter", types.YLeaf{"NeighborFilter", ipv4.NeighborFilter})
    ipv4.EntityData.Leafs.Append("spt-threshold-infinity", types.YLeaf{"SptThresholdInfinity", ipv4.SptThresholdInfinity})
    ipv4.EntityData.Leafs.Append("log-neighbor-changes", types.YLeaf{"LogNeighborChanges", ipv4.LogNeighborChanges})
    ipv4.EntityData.Leafs.Append("register-source", types.YLeaf{"RegisterSource", ipv4.RegisterSource})
    ipv4.EntityData.Leafs.Append("accept-register", types.YLeaf{"AcceptRegister", ipv4.AcceptRegister})
    ipv4.EntityData.Leafs.Append("suppress-rpf-prunes", types.YLeaf{"SuppressRpfPrunes", ipv4.SuppressRpfPrunes})
    ipv4.EntityData.Leafs.Append("ssm-allow-override", types.YLeaf{"SsmAllowOverride", ipv4.SsmAllowOverride})
    ipv4.EntityData.Leafs.Append("multipath", types.YLeaf{"Multipath", ipv4.Multipath})
    ipv4.EntityData.Leafs.Append("rp-static-deny", types.YLeaf{"RpStaticDeny", ipv4.RpStaticDeny})
    ipv4.EntityData.Leafs.Append("suppress-data-registers", types.YLeaf{"SuppressDataRegisters", ipv4.SuppressDataRegisters})
    ipv4.EntityData.Leafs.Append("neighbor-check-on-send", types.YLeaf{"NeighborCheckOnSend", ipv4.NeighborCheckOnSend})
    ipv4.EntityData.Leafs.Append("auto-rp-disable", types.YLeaf{"AutoRpDisable", ipv4.AutoRpDisable})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Pim_DefaultContext_Ipv4_RpfRedirect
// Configure RPF-redirect feature
type Pim_DefaultContext_Ipv4_RpfRedirect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route policy to select RPF topology. The type is string with length: 1..64.
    RoutePolicy interface{}
}

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetEntityData() *types.CommonEntityData {
    rpfRedirect.EntityData.YFilter = rpfRedirect.YFilter
    rpfRedirect.EntityData.YangName = "rpf-redirect"
    rpfRedirect.EntityData.BundleName = "cisco_ios_xr"
    rpfRedirect.EntityData.ParentYangName = "ipv4"
    rpfRedirect.EntityData.SegmentPath = "rpf-redirect"
    rpfRedirect.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + rpfRedirect.EntityData.SegmentPath
    rpfRedirect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpfRedirect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpfRedirect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpfRedirect.EntityData.Children = types.NewOrderedMap()
    rpfRedirect.EntityData.Leafs = types.NewOrderedMap()
    rpfRedirect.EntityData.Leafs.Append("route-policy", types.YLeaf{"RoutePolicy", rpfRedirect.RoutePolicy})

    rpfRedirect.EntityData.YListKeys = []string {}

    return &(rpfRedirect.EntityData)
}

// Pim_DefaultContext_Ipv4_Interfaces
// Interface-level Configuration
type Pim_DefaultContext_Ipv4_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Pim_DefaultContext_Ipv4_Interfaces_Interface.
    Interface []*Pim_DefaultContext_Ipv4_Interfaces_Interface
}

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ipv4"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + interfaces.EntityData.SegmentPath
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

// Pim_DefaultContext_Ipv4_Interfaces_Interface
// The name of the interface
type Pim_DefaultContext_Ipv4_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of interface. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Enter PIM Interface processing. The type is interface{}.
    Enable interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // BSR Border configuration for Interface. The type is bool.
    BsrBorder interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Enable PIM processing on the interface. The type is bool.
    InterfaceEnable interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}

    // Configure RPF-redirect bundle for interface. Applicable for IPv4 only.
    RedirectBundle Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle

    // Maximum number of allowed routes for this interface.
    MaximumRoutes Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes

    // BFD configuration.
    Bfd Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd
}

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("redirect-bundle", types.YChild{"RedirectBundle", &self.RedirectBundle})
    self.EntityData.Children.Append("maximum-routes", types.YChild{"MaximumRoutes", &self.MaximumRoutes})
    self.EntityData.Children.Append("bfd", types.YChild{"Bfd", &self.Bfd})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})
    self.EntityData.Leafs.Append("neighbor-filter", types.YLeaf{"NeighborFilter", self.NeighborFilter})
    self.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", self.HelloInterval})
    self.EntityData.Leafs.Append("bsr-border", types.YLeaf{"BsrBorder", self.BsrBorder})
    self.EntityData.Leafs.Append("propagation-delay", types.YLeaf{"PropagationDelay", self.PropagationDelay})
    self.EntityData.Leafs.Append("dr-priority", types.YLeaf{"DrPriority", self.DrPriority})
    self.EntityData.Leafs.Append("join-prune-mtu", types.YLeaf{"JoinPruneMtu", self.JoinPruneMtu})
    self.EntityData.Leafs.Append("interface-enable", types.YLeaf{"InterfaceEnable", self.InterfaceEnable})
    self.EntityData.Leafs.Append("jp-interval", types.YLeaf{"JpInterval", self.JpInterval})
    self.EntityData.Leafs.Append("override-interval", types.YLeaf{"OverrideInterval", self.OverrideInterval})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle
// Configure RPF-redirect bundle for interface.
// Applicable for IPv4 only
type Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bundle name. The type is string with length: 1..32.
    BundleName interface{}

    // Interface bandwidth in Kbps. The type is interface{} with range:
    // 0..100000000. Units are kbit/s.
    InterfaceBandwidth interface{}

    // Threshold bandwidth in Kbps. The type is interface{} with range:
    // 0..100000000. Units are kbit/s.
    ThresholdBandwidth interface{}
}

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetEntityData() *types.CommonEntityData {
    redirectBundle.EntityData.YFilter = redirectBundle.YFilter
    redirectBundle.EntityData.YangName = "redirect-bundle"
    redirectBundle.EntityData.BundleName = "cisco_ios_xr"
    redirectBundle.EntityData.ParentYangName = "interface"
    redirectBundle.EntityData.SegmentPath = "redirect-bundle"
    redirectBundle.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/interfaces/interface/" + redirectBundle.EntityData.SegmentPath
    redirectBundle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redirectBundle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redirectBundle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redirectBundle.EntityData.Children = types.NewOrderedMap()
    redirectBundle.EntityData.Leafs = types.NewOrderedMap()
    redirectBundle.EntityData.Leafs.Append("bundle-name", types.YLeaf{"BundleName", redirectBundle.BundleName})
    redirectBundle.EntityData.Leafs.Append("interface-bandwidth", types.YLeaf{"InterfaceBandwidth", redirectBundle.InterfaceBandwidth})
    redirectBundle.EntityData.Leafs.Append("threshold-bandwidth", types.YLeaf{"ThresholdBandwidth", redirectBundle.ThresholdBandwidth})

    redirectBundle.EntityData.YListKeys = []string {}

    return &(redirectBundle.EntityData)
}

// Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes
// Maximum number of allowed routes for this
// interface
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of routes for this interface. The type is interface{} with
    // range: 1..1100000. This attribute is mandatory.
    Maximum interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetEntityData() *types.CommonEntityData {
    maximumRoutes.EntityData.YFilter = maximumRoutes.YFilter
    maximumRoutes.EntityData.YangName = "maximum-routes"
    maximumRoutes.EntityData.BundleName = "cisco_ios_xr"
    maximumRoutes.EntityData.ParentYangName = "interface"
    maximumRoutes.EntityData.SegmentPath = "maximum-routes"
    maximumRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/interfaces/interface/" + maximumRoutes.EntityData.SegmentPath
    maximumRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumRoutes.EntityData.Children = types.NewOrderedMap()
    maximumRoutes.EntityData.Leafs = types.NewOrderedMap()
    maximumRoutes.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", maximumRoutes.Maximum})
    maximumRoutes.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", maximumRoutes.WarningThreshold})
    maximumRoutes.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", maximumRoutes.AccessListName})

    maximumRoutes.EntityData.YListKeys = []string {}

    return &(maximumRoutes.EntityData)
}

// Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd
// BFD configuration
type Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detection multiplier for BFD sessions created by PIM. The type is
    // interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval for BFD sessions created by PIM. The type is interface{}
    // with range: 3..30000. Units are millisecond.
    Interval interface{}

    // TRUE to enable BFD. FALSE to disable and to prevent inheritance from a
    // parent. The type is bool.
    Enable interface{}
}

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "interface"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/interfaces/interface/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bfd.Enable})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Pim_DefaultContext_Ipv4_AutoRpCandidateRps
// Configure Candidate-RPs
type Pim_DefaultContext_Ipv4_AutoRpCandidateRps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifications for a Candidate-RP. The type is slice of
    // Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp.
    AutoRpCandidateRp []*Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp
}

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetEntityData() *types.CommonEntityData {
    autoRpCandidateRps.EntityData.YFilter = autoRpCandidateRps.YFilter
    autoRpCandidateRps.EntityData.YangName = "auto-rp-candidate-rps"
    autoRpCandidateRps.EntityData.BundleName = "cisco_ios_xr"
    autoRpCandidateRps.EntityData.ParentYangName = "ipv4"
    autoRpCandidateRps.EntityData.SegmentPath = "auto-rp-candidate-rps"
    autoRpCandidateRps.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + autoRpCandidateRps.EntityData.SegmentPath
    autoRpCandidateRps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoRpCandidateRps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoRpCandidateRps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoRpCandidateRps.EntityData.Children = types.NewOrderedMap()
    autoRpCandidateRps.EntityData.Children.Append("auto-rp-candidate-rp", types.YChild{"AutoRpCandidateRp", nil})
    for i := range autoRpCandidateRps.AutoRpCandidateRp {
        autoRpCandidateRps.EntityData.Children.Append(types.GetSegmentPath(autoRpCandidateRps.AutoRpCandidateRp[i]), types.YChild{"AutoRpCandidateRp", autoRpCandidateRps.AutoRpCandidateRp[i]})
    }
    autoRpCandidateRps.EntityData.Leafs = types.NewOrderedMap()

    autoRpCandidateRps.EntityData.YListKeys = []string {}

    return &(autoRpCandidateRps.EntityData)
}

// Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp
// Specifications for a Candidate-RP
type Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface from which Candidate-RP packets will be
    // sourced. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // This attribute is a key. Protocol Mode. The type is AutoRpProtocolMode.
    ProtocolMode interface{}

    // TTL in Hops. The type is interface{} with range: 1..255. This attribute is
    // mandatory.
    Ttl interface{}

    // Access-list specifying the group range for the Candidate-RP. The type is
    // string with length: 1..64. The default value is 224-4.
    AccessListName interface{}

    // Time between announcements <in seconds> . The type is interface{} with
    // range: 1..600. Units are second. The default value is 60.
    AnnouncePeriod interface{}
}

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetEntityData() *types.CommonEntityData {
    autoRpCandidateRp.EntityData.YFilter = autoRpCandidateRp.YFilter
    autoRpCandidateRp.EntityData.YangName = "auto-rp-candidate-rp"
    autoRpCandidateRp.EntityData.BundleName = "cisco_ios_xr"
    autoRpCandidateRp.EntityData.ParentYangName = "auto-rp-candidate-rps"
    autoRpCandidateRp.EntityData.SegmentPath = "auto-rp-candidate-rp" + types.AddKeyToken(autoRpCandidateRp.InterfaceName, "interface-name") + types.AddKeyToken(autoRpCandidateRp.ProtocolMode, "protocol-mode")
    autoRpCandidateRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/auto-rp-candidate-rps/" + autoRpCandidateRp.EntityData.SegmentPath
    autoRpCandidateRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoRpCandidateRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoRpCandidateRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoRpCandidateRp.EntityData.Children = types.NewOrderedMap()
    autoRpCandidateRp.EntityData.Leafs = types.NewOrderedMap()
    autoRpCandidateRp.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", autoRpCandidateRp.InterfaceName})
    autoRpCandidateRp.EntityData.Leafs.Append("protocol-mode", types.YLeaf{"ProtocolMode", autoRpCandidateRp.ProtocolMode})
    autoRpCandidateRp.EntityData.Leafs.Append("ttl", types.YLeaf{"Ttl", autoRpCandidateRp.Ttl})
    autoRpCandidateRp.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", autoRpCandidateRp.AccessListName})
    autoRpCandidateRp.EntityData.Leafs.Append("announce-period", types.YLeaf{"AnnouncePeriod", autoRpCandidateRp.AnnouncePeriod})

    autoRpCandidateRp.EntityData.YListKeys = []string {"InterfaceName", "ProtocolMode"}

    return &(autoRpCandidateRp.EntityData)
}

// Pim_DefaultContext_Ipv4_AutoRpMappingAgent
// Configure AutoRP Mapping Agent
type Pim_DefaultContext_Ipv4_AutoRpMappingAgent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifications for Mapping Agent configured on this box.
    Parameters Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters

    // Mapping Agent cache size limit.
    CacheLimit Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit
}

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetEntityData() *types.CommonEntityData {
    autoRpMappingAgent.EntityData.YFilter = autoRpMappingAgent.YFilter
    autoRpMappingAgent.EntityData.YangName = "auto-rp-mapping-agent"
    autoRpMappingAgent.EntityData.BundleName = "cisco_ios_xr"
    autoRpMappingAgent.EntityData.ParentYangName = "ipv4"
    autoRpMappingAgent.EntityData.SegmentPath = "auto-rp-mapping-agent"
    autoRpMappingAgent.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + autoRpMappingAgent.EntityData.SegmentPath
    autoRpMappingAgent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoRpMappingAgent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoRpMappingAgent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoRpMappingAgent.EntityData.Children = types.NewOrderedMap()
    autoRpMappingAgent.EntityData.Children.Append("parameters", types.YChild{"Parameters", &autoRpMappingAgent.Parameters})
    autoRpMappingAgent.EntityData.Children.Append("cache-limit", types.YChild{"CacheLimit", &autoRpMappingAgent.CacheLimit})
    autoRpMappingAgent.EntityData.Leafs = types.NewOrderedMap()

    autoRpMappingAgent.EntityData.YListKeys = []string {}

    return &(autoRpMappingAgent.EntityData)
}

// Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters
// Specifications for Mapping Agent configured
// on this box
// This type is a presence type.
type Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Interface from which mapping packets will be sourced . The type is string
    // with pattern: [a-zA-Z0-9._/-]+. This attribute is mandatory.
    InterfaceName interface{}

    // TTL in Hops. The type is interface{} with range: 1..255. This attribute is
    // mandatory.
    Ttl interface{}

    // Time between discovery messages <in seconds>. The type is interface{} with
    // range: 1..600. Units are second. The default value is 60.
    AnnouncePeriod interface{}
}

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetEntityData() *types.CommonEntityData {
    parameters.EntityData.YFilter = parameters.YFilter
    parameters.EntityData.YangName = "parameters"
    parameters.EntityData.BundleName = "cisco_ios_xr"
    parameters.EntityData.ParentYangName = "auto-rp-mapping-agent"
    parameters.EntityData.SegmentPath = "parameters"
    parameters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/auto-rp-mapping-agent/" + parameters.EntityData.SegmentPath
    parameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parameters.EntityData.Children = types.NewOrderedMap()
    parameters.EntityData.Leafs = types.NewOrderedMap()
    parameters.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", parameters.InterfaceName})
    parameters.EntityData.Leafs.Append("ttl", types.YLeaf{"Ttl", parameters.Ttl})
    parameters.EntityData.Leafs.Append("announce-period", types.YLeaf{"AnnouncePeriod", parameters.AnnouncePeriod})

    parameters.EntityData.YListKeys = []string {}

    return &(parameters.EntityData)
}

// Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit
// Mapping Agent cache size limit
// This type is a presence type.
type Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of mapping cache entries. The type is interface{} with
    // range: 1..1000. This attribute is mandatory.
    MaximumCacheEntry interface{}

    // Warning threshold number of cache entries. The type is interface{} with
    // range: 1..1000. The default value is 450.
    ThresholdCacheEntry interface{}
}

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetEntityData() *types.CommonEntityData {
    cacheLimit.EntityData.YFilter = cacheLimit.YFilter
    cacheLimit.EntityData.YangName = "cache-limit"
    cacheLimit.EntityData.BundleName = "cisco_ios_xr"
    cacheLimit.EntityData.ParentYangName = "auto-rp-mapping-agent"
    cacheLimit.EntityData.SegmentPath = "cache-limit"
    cacheLimit.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/auto-rp-mapping-agent/" + cacheLimit.EntityData.SegmentPath
    cacheLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cacheLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cacheLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cacheLimit.EntityData.Children = types.NewOrderedMap()
    cacheLimit.EntityData.Leafs = types.NewOrderedMap()
    cacheLimit.EntityData.Leafs.Append("maximum-cache-entry", types.YLeaf{"MaximumCacheEntry", cacheLimit.MaximumCacheEntry})
    cacheLimit.EntityData.Leafs.Append("threshold-cache-entry", types.YLeaf{"ThresholdCacheEntry", cacheLimit.ThresholdCacheEntry})

    cacheLimit.EntityData.YListKeys = []string {}

    return &(cacheLimit.EntityData)
}

// Pim_DefaultContext_Ipv4_SparseModeRpAddresses
// Configure Sparse-Mode Rendezvous Point
type Pim_DefaultContext_Ipv4_SparseModeRpAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress.
    SparseModeRpAddress []*Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetEntityData() *types.CommonEntityData {
    sparseModeRpAddresses.EntityData.YFilter = sparseModeRpAddresses.YFilter
    sparseModeRpAddresses.EntityData.YangName = "sparse-mode-rp-addresses"
    sparseModeRpAddresses.EntityData.BundleName = "cisco_ios_xr"
    sparseModeRpAddresses.EntityData.ParentYangName = "ipv4"
    sparseModeRpAddresses.EntityData.SegmentPath = "sparse-mode-rp-addresses"
    sparseModeRpAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + sparseModeRpAddresses.EntityData.SegmentPath
    sparseModeRpAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sparseModeRpAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sparseModeRpAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sparseModeRpAddresses.EntityData.Children = types.NewOrderedMap()
    sparseModeRpAddresses.EntityData.Children.Append("sparse-mode-rp-address", types.YChild{"SparseModeRpAddress", nil})
    for i := range sparseModeRpAddresses.SparseModeRpAddress {
        sparseModeRpAddresses.EntityData.Children.Append(types.GetSegmentPath(sparseModeRpAddresses.SparseModeRpAddress[i]), types.YChild{"SparseModeRpAddress", sparseModeRpAddresses.SparseModeRpAddress[i]})
    }
    sparseModeRpAddresses.EntityData.Leafs = types.NewOrderedMap()

    sparseModeRpAddresses.EntityData.YListKeys = []string {}

    return &(sparseModeRpAddresses.EntityData)
}

// Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress
// Address of the Rendezvous Point
type Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a  given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetEntityData() *types.CommonEntityData {
    sparseModeRpAddress.EntityData.YFilter = sparseModeRpAddress.YFilter
    sparseModeRpAddress.EntityData.YangName = "sparse-mode-rp-address"
    sparseModeRpAddress.EntityData.BundleName = "cisco_ios_xr"
    sparseModeRpAddress.EntityData.ParentYangName = "sparse-mode-rp-addresses"
    sparseModeRpAddress.EntityData.SegmentPath = "sparse-mode-rp-address" + types.AddKeyToken(sparseModeRpAddress.RpAddress, "rp-address")
    sparseModeRpAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/sparse-mode-rp-addresses/" + sparseModeRpAddress.EntityData.SegmentPath
    sparseModeRpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sparseModeRpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sparseModeRpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sparseModeRpAddress.EntityData.Children = types.NewOrderedMap()
    sparseModeRpAddress.EntityData.Leafs = types.NewOrderedMap()
    sparseModeRpAddress.EntityData.Leafs.Append("rp-address", types.YLeaf{"RpAddress", sparseModeRpAddress.RpAddress})
    sparseModeRpAddress.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", sparseModeRpAddress.AccessListName})
    sparseModeRpAddress.EntityData.Leafs.Append("auto-rp-override", types.YLeaf{"AutoRpOverride", sparseModeRpAddress.AutoRpOverride})

    sparseModeRpAddress.EntityData.YListKeys = []string {"RpAddress"}

    return &(sparseModeRpAddress.EntityData)
}

// Pim_DefaultContext_Ipv4_InheritableDefaults
// Inheritable defaults
type Pim_DefaultContext_Ipv4_InheritableDefaults struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Convergency timeout in seconds. The type is interface{} with range:
    // 1800..2400. Units are second.
    ConvergenceTimeout interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}
}

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetEntityData() *types.CommonEntityData {
    inheritableDefaults.EntityData.YFilter = inheritableDefaults.YFilter
    inheritableDefaults.EntityData.YangName = "inheritable-defaults"
    inheritableDefaults.EntityData.BundleName = "cisco_ios_xr"
    inheritableDefaults.EntityData.ParentYangName = "ipv4"
    inheritableDefaults.EntityData.SegmentPath = "inheritable-defaults"
    inheritableDefaults.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + inheritableDefaults.EntityData.SegmentPath
    inheritableDefaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritableDefaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritableDefaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritableDefaults.EntityData.Children = types.NewOrderedMap()
    inheritableDefaults.EntityData.Leafs = types.NewOrderedMap()
    inheritableDefaults.EntityData.Leafs.Append("convergence-timeout", types.YLeaf{"ConvergenceTimeout", inheritableDefaults.ConvergenceTimeout})
    inheritableDefaults.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", inheritableDefaults.HelloInterval})
    inheritableDefaults.EntityData.Leafs.Append("propagation-delay", types.YLeaf{"PropagationDelay", inheritableDefaults.PropagationDelay})
    inheritableDefaults.EntityData.Leafs.Append("dr-priority", types.YLeaf{"DrPriority", inheritableDefaults.DrPriority})
    inheritableDefaults.EntityData.Leafs.Append("join-prune-mtu", types.YLeaf{"JoinPruneMtu", inheritableDefaults.JoinPruneMtu})
    inheritableDefaults.EntityData.Leafs.Append("jp-interval", types.YLeaf{"JpInterval", inheritableDefaults.JpInterval})
    inheritableDefaults.EntityData.Leafs.Append("override-interval", types.YLeaf{"OverrideInterval", inheritableDefaults.OverrideInterval})

    inheritableDefaults.EntityData.YListKeys = []string {}

    return &(inheritableDefaults.EntityData)
}

// Pim_DefaultContext_Ipv4_Rpf
// Configure RPF options
type Pim_DefaultContext_Ipv4_Rpf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route policy to select RPF topology. The type is string with length: 1..64.
    RoutePolicy interface{}
}

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetEntityData() *types.CommonEntityData {
    rpf.EntityData.YFilter = rpf.YFilter
    rpf.EntityData.YangName = "rpf"
    rpf.EntityData.BundleName = "cisco_ios_xr"
    rpf.EntityData.ParentYangName = "ipv4"
    rpf.EntityData.SegmentPath = "rpf"
    rpf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + rpf.EntityData.SegmentPath
    rpf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpf.EntityData.Children = types.NewOrderedMap()
    rpf.EntityData.Leafs = types.NewOrderedMap()
    rpf.EntityData.Leafs.Append("route-policy", types.YLeaf{"RoutePolicy", rpf.RoutePolicy})

    rpf.EntityData.YListKeys = []string {}

    return &(rpf.EntityData)
}

// Pim_DefaultContext_Ipv4_SgExpiryTimer
// Configure expiry timer for S,G routes
type Pim_DefaultContext_Ipv4_SgExpiryTimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // (S,G) expiry time in seconds. The type is interface{} with range:
    // 40..57600. Units are second.
    Interval interface{}

    // Access-list of applicable S,G routes. The type is string with length:
    // 1..64.
    AccessListName interface{}
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetEntityData() *types.CommonEntityData {
    sgExpiryTimer.EntityData.YFilter = sgExpiryTimer.YFilter
    sgExpiryTimer.EntityData.YangName = "sg-expiry-timer"
    sgExpiryTimer.EntityData.BundleName = "cisco_ios_xr"
    sgExpiryTimer.EntityData.ParentYangName = "ipv4"
    sgExpiryTimer.EntityData.SegmentPath = "sg-expiry-timer"
    sgExpiryTimer.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + sgExpiryTimer.EntityData.SegmentPath
    sgExpiryTimer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sgExpiryTimer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sgExpiryTimer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sgExpiryTimer.EntityData.Children = types.NewOrderedMap()
    sgExpiryTimer.EntityData.Leafs = types.NewOrderedMap()
    sgExpiryTimer.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", sgExpiryTimer.Interval})
    sgExpiryTimer.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", sgExpiryTimer.AccessListName})

    sgExpiryTimer.EntityData.YListKeys = []string {}

    return &(sgExpiryTimer.EntityData)
}

// Pim_DefaultContext_Ipv4_RpfVectorEnable
// Enable PIM RPF Vector Proxy's
// This type is a presence type.
type Pim_DefaultContext_Ipv4_RpfVectorEnable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // RPF Vector is turned on if configured. The type is interface{}. This
    // attribute is mandatory.
    Enable interface{}

    // Allow RPF Vector origination over eBGP sessions. The type is interface{}.
    AllowEbgp interface{}

    // Disable RPF Vector origination over iBGP sessions. The type is interface{}.
    DisableIbgp interface{}

    // Use RPF Vector RFC compliant encoding. The type is interface{}.
    UseStandardEncoding interface{}
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetEntityData() *types.CommonEntityData {
    rpfVectorEnable.EntityData.YFilter = rpfVectorEnable.YFilter
    rpfVectorEnable.EntityData.YangName = "rpf-vector-enable"
    rpfVectorEnable.EntityData.BundleName = "cisco_ios_xr"
    rpfVectorEnable.EntityData.ParentYangName = "ipv4"
    rpfVectorEnable.EntityData.SegmentPath = "rpf-vector-enable"
    rpfVectorEnable.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + rpfVectorEnable.EntityData.SegmentPath
    rpfVectorEnable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpfVectorEnable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpfVectorEnable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpfVectorEnable.EntityData.Children = types.NewOrderedMap()
    rpfVectorEnable.EntityData.Leafs = types.NewOrderedMap()
    rpfVectorEnable.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", rpfVectorEnable.Enable})
    rpfVectorEnable.EntityData.Leafs.Append("allow-ebgp", types.YLeaf{"AllowEbgp", rpfVectorEnable.AllowEbgp})
    rpfVectorEnable.EntityData.Leafs.Append("disable-ibgp", types.YLeaf{"DisableIbgp", rpfVectorEnable.DisableIbgp})
    rpfVectorEnable.EntityData.Leafs.Append("use-standard-encoding", types.YLeaf{"UseStandardEncoding", rpfVectorEnable.UseStandardEncoding})

    rpfVectorEnable.EntityData.YListKeys = []string {}

    return &(rpfVectorEnable.EntityData)
}

// Pim_DefaultContext_Ipv4_Nsf
// Configure Non-stop forwarding (NSF) options
type Pim_DefaultContext_Ipv4_Nsf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Override default maximum lifetime for PIM NSF mode. The type is interface{}
    // with range: 10..600. Units are second.
    Lifetime interface{}
}

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetEntityData() *types.CommonEntityData {
    nsf.EntityData.YFilter = nsf.YFilter
    nsf.EntityData.YangName = "nsf"
    nsf.EntityData.BundleName = "cisco_ios_xr"
    nsf.EntityData.ParentYangName = "ipv4"
    nsf.EntityData.SegmentPath = "nsf"
    nsf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + nsf.EntityData.SegmentPath
    nsf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nsf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nsf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nsf.EntityData.Children = types.NewOrderedMap()
    nsf.EntityData.Leafs = types.NewOrderedMap()
    nsf.EntityData.Leafs.Append("lifetime", types.YLeaf{"Lifetime", nsf.Lifetime})

    nsf.EntityData.YListKeys = []string {}

    return &(nsf.EntityData)
}

// Pim_DefaultContext_Ipv4_Maximum
// Configure PIM State Limits
type Pim_DefaultContext_Ipv4_Maximum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum packet queue size in bytes. The type is interface{} with range:
    // 0..2147483648. Units are byte.
    GlobalLowPriorityPacketQueue interface{}

    // Maximum packet queue size in bytes. The type is interface{} with range:
    // 0..2147483648. Units are byte.
    GlobalHighPriorityPacketQueue interface{}

    // Override default global maximum and threshold for PIM group mapping ranges
    // from BSR.
    BsrGlobalGroupMappings Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings

    // Override default maximum for number of routes.
    GlobalRoutes Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes

    // Maximum for number of group mappings from autorp mapping agent.
    GlobalGroupMappingsAutoRp Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp

    // Override default global maximum and threshold for C-RP set in BSR.
    BsrGlobalCandidateRpCache Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache

    // Override default maximum for number of sparse-mode source registers.
    GlobalRegisterStates Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates

    // Override default maximum for number of route-interfaces.
    GlobalRouteInterfaces Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces

    // Override default maximum for number of group mappings from autorp mapping
    // agent.
    GroupMappingsAutoRp Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp

    // Override default maximum and threshold for number of group mappings from
    // BSR.
    BsrGroupMappings Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings

    // Override default maximum for number of sparse-mode source registers.
    RegisterStates Pim_DefaultContext_Ipv4_Maximum_RegisterStates

    // Override default maximum for number of route-interfaces.
    RouteInterfaces Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces

    // Override default maximum and threshold for BSR C-RP cache setting.
    BsrCandidateRpCache Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache

    // Override default maximum for number of routes.
    Routes Pim_DefaultContext_Ipv4_Maximum_Routes
}

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetEntityData() *types.CommonEntityData {
    maximum.EntityData.YFilter = maximum.YFilter
    maximum.EntityData.YangName = "maximum"
    maximum.EntityData.BundleName = "cisco_ios_xr"
    maximum.EntityData.ParentYangName = "ipv4"
    maximum.EntityData.SegmentPath = "maximum"
    maximum.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + maximum.EntityData.SegmentPath
    maximum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximum.EntityData.Children = types.NewOrderedMap()
    maximum.EntityData.Children.Append("bsr-global-group-mappings", types.YChild{"BsrGlobalGroupMappings", &maximum.BsrGlobalGroupMappings})
    maximum.EntityData.Children.Append("global-routes", types.YChild{"GlobalRoutes", &maximum.GlobalRoutes})
    maximum.EntityData.Children.Append("global-group-mappings-auto-rp", types.YChild{"GlobalGroupMappingsAutoRp", &maximum.GlobalGroupMappingsAutoRp})
    maximum.EntityData.Children.Append("bsr-global-candidate-rp-cache", types.YChild{"BsrGlobalCandidateRpCache", &maximum.BsrGlobalCandidateRpCache})
    maximum.EntityData.Children.Append("global-register-states", types.YChild{"GlobalRegisterStates", &maximum.GlobalRegisterStates})
    maximum.EntityData.Children.Append("global-route-interfaces", types.YChild{"GlobalRouteInterfaces", &maximum.GlobalRouteInterfaces})
    maximum.EntityData.Children.Append("group-mappings-auto-rp", types.YChild{"GroupMappingsAutoRp", &maximum.GroupMappingsAutoRp})
    maximum.EntityData.Children.Append("bsr-group-mappings", types.YChild{"BsrGroupMappings", &maximum.BsrGroupMappings})
    maximum.EntityData.Children.Append("register-states", types.YChild{"RegisterStates", &maximum.RegisterStates})
    maximum.EntityData.Children.Append("route-interfaces", types.YChild{"RouteInterfaces", &maximum.RouteInterfaces})
    maximum.EntityData.Children.Append("bsr-candidate-rp-cache", types.YChild{"BsrCandidateRpCache", &maximum.BsrCandidateRpCache})
    maximum.EntityData.Children.Append("routes", types.YChild{"Routes", &maximum.Routes})
    maximum.EntityData.Leafs = types.NewOrderedMap()
    maximum.EntityData.Leafs.Append("global-low-priority-packet-queue", types.YLeaf{"GlobalLowPriorityPacketQueue", maximum.GlobalLowPriorityPacketQueue})
    maximum.EntityData.Leafs.Append("global-high-priority-packet-queue", types.YLeaf{"GlobalHighPriorityPacketQueue", maximum.GlobalHighPriorityPacketQueue})

    maximum.EntityData.YListKeys = []string {}

    return &(maximum.EntityData)
}

// Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings
// Override default global maximum and threshold
// for PIM group mapping ranges from BSR
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Global Maximum number of PIM group mapping ranges from BSR. The type is
    // interface{} with range: 1..10000. This attribute is mandatory.
    BsrMaximumGlobalGroupMappings interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 500.
    WarningThreshold interface{}
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetEntityData() *types.CommonEntityData {
    bsrGlobalGroupMappings.EntityData.YFilter = bsrGlobalGroupMappings.YFilter
    bsrGlobalGroupMappings.EntityData.YangName = "bsr-global-group-mappings"
    bsrGlobalGroupMappings.EntityData.BundleName = "cisco_ios_xr"
    bsrGlobalGroupMappings.EntityData.ParentYangName = "maximum"
    bsrGlobalGroupMappings.EntityData.SegmentPath = "bsr-global-group-mappings"
    bsrGlobalGroupMappings.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/maximum/" + bsrGlobalGroupMappings.EntityData.SegmentPath
    bsrGlobalGroupMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsrGlobalGroupMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsrGlobalGroupMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsrGlobalGroupMappings.EntityData.Children = types.NewOrderedMap()
    bsrGlobalGroupMappings.EntityData.Leafs = types.NewOrderedMap()
    bsrGlobalGroupMappings.EntityData.Leafs.Append("bsr-maximum-global-group-mappings", types.YLeaf{"BsrMaximumGlobalGroupMappings", bsrGlobalGroupMappings.BsrMaximumGlobalGroupMappings})
    bsrGlobalGroupMappings.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", bsrGlobalGroupMappings.WarningThreshold})

    bsrGlobalGroupMappings.EntityData.YListKeys = []string {}

    return &(bsrGlobalGroupMappings.EntityData)
}

// Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes
// Override default maximum for number of routes
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM routes. The type is interface{} with range:
    // 1..200000. This attribute is mandatory.
    MaximumRoutes interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..200000. The default value is 100000.
    WarningThreshold interface{}
}

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetEntityData() *types.CommonEntityData {
    globalRoutes.EntityData.YFilter = globalRoutes.YFilter
    globalRoutes.EntityData.YangName = "global-routes"
    globalRoutes.EntityData.BundleName = "cisco_ios_xr"
    globalRoutes.EntityData.ParentYangName = "maximum"
    globalRoutes.EntityData.SegmentPath = "global-routes"
    globalRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/maximum/" + globalRoutes.EntityData.SegmentPath
    globalRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalRoutes.EntityData.Children = types.NewOrderedMap()
    globalRoutes.EntityData.Leafs = types.NewOrderedMap()
    globalRoutes.EntityData.Leafs.Append("maximum-routes", types.YLeaf{"MaximumRoutes", globalRoutes.MaximumRoutes})
    globalRoutes.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", globalRoutes.WarningThreshold})

    globalRoutes.EntityData.YListKeys = []string {}

    return &(globalRoutes.EntityData)
}

// Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp
// Maximum for number of group mappings from
// autorp mapping agent
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM group mappings from autorp. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    MaximumGlobalGroupRangesAutoRp interface{}

    // Warning threshold number of PIM group mappings from autorp. The type is
    // interface{} with range: 1..10000. The default value is 450.
    ThresholdGlobalGroupRangesAutoRp interface{}
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetEntityData() *types.CommonEntityData {
    globalGroupMappingsAutoRp.EntityData.YFilter = globalGroupMappingsAutoRp.YFilter
    globalGroupMappingsAutoRp.EntityData.YangName = "global-group-mappings-auto-rp"
    globalGroupMappingsAutoRp.EntityData.BundleName = "cisco_ios_xr"
    globalGroupMappingsAutoRp.EntityData.ParentYangName = "maximum"
    globalGroupMappingsAutoRp.EntityData.SegmentPath = "global-group-mappings-auto-rp"
    globalGroupMappingsAutoRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/maximum/" + globalGroupMappingsAutoRp.EntityData.SegmentPath
    globalGroupMappingsAutoRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalGroupMappingsAutoRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalGroupMappingsAutoRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalGroupMappingsAutoRp.EntityData.Children = types.NewOrderedMap()
    globalGroupMappingsAutoRp.EntityData.Leafs = types.NewOrderedMap()
    globalGroupMappingsAutoRp.EntityData.Leafs.Append("maximum-global-group-ranges-auto-rp", types.YLeaf{"MaximumGlobalGroupRangesAutoRp", globalGroupMappingsAutoRp.MaximumGlobalGroupRangesAutoRp})
    globalGroupMappingsAutoRp.EntityData.Leafs.Append("threshold-global-group-ranges-auto-rp", types.YLeaf{"ThresholdGlobalGroupRangesAutoRp", globalGroupMappingsAutoRp.ThresholdGlobalGroupRangesAutoRp})

    globalGroupMappingsAutoRp.EntityData.YListKeys = []string {}

    return &(globalGroupMappingsAutoRp.EntityData)
}

// Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache
// Override default global maximum and threshold
// for C-RP set in BSR
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Global Maximum number of PIM C-RP Sets from BSR. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    BsrMaximumGlobalCandidateRpCache interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 100.
    WarningThreshold interface{}
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetEntityData() *types.CommonEntityData {
    bsrGlobalCandidateRpCache.EntityData.YFilter = bsrGlobalCandidateRpCache.YFilter
    bsrGlobalCandidateRpCache.EntityData.YangName = "bsr-global-candidate-rp-cache"
    bsrGlobalCandidateRpCache.EntityData.BundleName = "cisco_ios_xr"
    bsrGlobalCandidateRpCache.EntityData.ParentYangName = "maximum"
    bsrGlobalCandidateRpCache.EntityData.SegmentPath = "bsr-global-candidate-rp-cache"
    bsrGlobalCandidateRpCache.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/maximum/" + bsrGlobalCandidateRpCache.EntityData.SegmentPath
    bsrGlobalCandidateRpCache.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsrGlobalCandidateRpCache.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsrGlobalCandidateRpCache.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsrGlobalCandidateRpCache.EntityData.Children = types.NewOrderedMap()
    bsrGlobalCandidateRpCache.EntityData.Leafs = types.NewOrderedMap()
    bsrGlobalCandidateRpCache.EntityData.Leafs.Append("bsr-maximum-global-candidate-rp-cache", types.YLeaf{"BsrMaximumGlobalCandidateRpCache", bsrGlobalCandidateRpCache.BsrMaximumGlobalCandidateRpCache})
    bsrGlobalCandidateRpCache.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", bsrGlobalCandidateRpCache.WarningThreshold})

    bsrGlobalCandidateRpCache.EntityData.YListKeys = []string {}

    return &(bsrGlobalCandidateRpCache.EntityData)
}

// Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates
// Override default maximum for number of
// sparse-mode source registers
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM Sparse-Mode register states. The type is interface{}
    // with range: 0..75000. This attribute is mandatory.
    MaximumRegisterStates interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 0..75000. The default value is 20000.
    WarningThreshold interface{}
}

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetEntityData() *types.CommonEntityData {
    globalRegisterStates.EntityData.YFilter = globalRegisterStates.YFilter
    globalRegisterStates.EntityData.YangName = "global-register-states"
    globalRegisterStates.EntityData.BundleName = "cisco_ios_xr"
    globalRegisterStates.EntityData.ParentYangName = "maximum"
    globalRegisterStates.EntityData.SegmentPath = "global-register-states"
    globalRegisterStates.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/maximum/" + globalRegisterStates.EntityData.SegmentPath
    globalRegisterStates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalRegisterStates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalRegisterStates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalRegisterStates.EntityData.Children = types.NewOrderedMap()
    globalRegisterStates.EntityData.Leafs = types.NewOrderedMap()
    globalRegisterStates.EntityData.Leafs.Append("maximum-register-states", types.YLeaf{"MaximumRegisterStates", globalRegisterStates.MaximumRegisterStates})
    globalRegisterStates.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", globalRegisterStates.WarningThreshold})

    globalRegisterStates.EntityData.YListKeys = []string {}

    return &(globalRegisterStates.EntityData)
}

// Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces
// Override default maximum for number of
// route-interfaces
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM route-interfaces. The type is interface{} with range:
    // 1..1100000. This attribute is mandatory.
    MaximumRouteInterfaces interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000. The default value is 300000.
    WarningThreshold interface{}
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetEntityData() *types.CommonEntityData {
    globalRouteInterfaces.EntityData.YFilter = globalRouteInterfaces.YFilter
    globalRouteInterfaces.EntityData.YangName = "global-route-interfaces"
    globalRouteInterfaces.EntityData.BundleName = "cisco_ios_xr"
    globalRouteInterfaces.EntityData.ParentYangName = "maximum"
    globalRouteInterfaces.EntityData.SegmentPath = "global-route-interfaces"
    globalRouteInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/maximum/" + globalRouteInterfaces.EntityData.SegmentPath
    globalRouteInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalRouteInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalRouteInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalRouteInterfaces.EntityData.Children = types.NewOrderedMap()
    globalRouteInterfaces.EntityData.Leafs = types.NewOrderedMap()
    globalRouteInterfaces.EntityData.Leafs.Append("maximum-route-interfaces", types.YLeaf{"MaximumRouteInterfaces", globalRouteInterfaces.MaximumRouteInterfaces})
    globalRouteInterfaces.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", globalRouteInterfaces.WarningThreshold})

    globalRouteInterfaces.EntityData.YListKeys = []string {}

    return &(globalRouteInterfaces.EntityData)
}

// Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp
// Override default maximum for number of group
// mappings from autorp mapping agent
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM group mappings from autorp. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    MaximumGroupRangesAutoRp interface{}

    // Warning threshold number of PIM group mappings from autorp. The type is
    // interface{} with range: 1..10000. The default value is 450.
    ThresholdGroupRangesAutoRp interface{}
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetEntityData() *types.CommonEntityData {
    groupMappingsAutoRp.EntityData.YFilter = groupMappingsAutoRp.YFilter
    groupMappingsAutoRp.EntityData.YangName = "group-mappings-auto-rp"
    groupMappingsAutoRp.EntityData.BundleName = "cisco_ios_xr"
    groupMappingsAutoRp.EntityData.ParentYangName = "maximum"
    groupMappingsAutoRp.EntityData.SegmentPath = "group-mappings-auto-rp"
    groupMappingsAutoRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/maximum/" + groupMappingsAutoRp.EntityData.SegmentPath
    groupMappingsAutoRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupMappingsAutoRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupMappingsAutoRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupMappingsAutoRp.EntityData.Children = types.NewOrderedMap()
    groupMappingsAutoRp.EntityData.Leafs = types.NewOrderedMap()
    groupMappingsAutoRp.EntityData.Leafs.Append("maximum-group-ranges-auto-rp", types.YLeaf{"MaximumGroupRangesAutoRp", groupMappingsAutoRp.MaximumGroupRangesAutoRp})
    groupMappingsAutoRp.EntityData.Leafs.Append("threshold-group-ranges-auto-rp", types.YLeaf{"ThresholdGroupRangesAutoRp", groupMappingsAutoRp.ThresholdGroupRangesAutoRp})

    groupMappingsAutoRp.EntityData.YListKeys = []string {}

    return &(groupMappingsAutoRp.EntityData)
}

// Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings
// Override default maximum and threshold for
// number of group mappings from BSR
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM group mappings from BSR. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumGroupRanges interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 500.
    WarningThreshold interface{}
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetEntityData() *types.CommonEntityData {
    bsrGroupMappings.EntityData.YFilter = bsrGroupMappings.YFilter
    bsrGroupMappings.EntityData.YangName = "bsr-group-mappings"
    bsrGroupMappings.EntityData.BundleName = "cisco_ios_xr"
    bsrGroupMappings.EntityData.ParentYangName = "maximum"
    bsrGroupMappings.EntityData.SegmentPath = "bsr-group-mappings"
    bsrGroupMappings.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/maximum/" + bsrGroupMappings.EntityData.SegmentPath
    bsrGroupMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsrGroupMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsrGroupMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsrGroupMappings.EntityData.Children = types.NewOrderedMap()
    bsrGroupMappings.EntityData.Leafs = types.NewOrderedMap()
    bsrGroupMappings.EntityData.Leafs.Append("bsr-maximum-group-ranges", types.YLeaf{"BsrMaximumGroupRanges", bsrGroupMappings.BsrMaximumGroupRanges})
    bsrGroupMappings.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", bsrGroupMappings.WarningThreshold})

    bsrGroupMappings.EntityData.YListKeys = []string {}

    return &(bsrGroupMappings.EntityData)
}

// Pim_DefaultContext_Ipv4_Maximum_RegisterStates
// Override default maximum for number of
// sparse-mode source registers
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_RegisterStates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM Sparse-Mode register states. The type is interface{}
    // with range: 0..75000. This attribute is mandatory.
    MaximumRegisterStates interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 0..75000. The default value is 20000.
    WarningThreshold interface{}
}

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetEntityData() *types.CommonEntityData {
    registerStates.EntityData.YFilter = registerStates.YFilter
    registerStates.EntityData.YangName = "register-states"
    registerStates.EntityData.BundleName = "cisco_ios_xr"
    registerStates.EntityData.ParentYangName = "maximum"
    registerStates.EntityData.SegmentPath = "register-states"
    registerStates.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/maximum/" + registerStates.EntityData.SegmentPath
    registerStates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registerStates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registerStates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registerStates.EntityData.Children = types.NewOrderedMap()
    registerStates.EntityData.Leafs = types.NewOrderedMap()
    registerStates.EntityData.Leafs.Append("maximum-register-states", types.YLeaf{"MaximumRegisterStates", registerStates.MaximumRegisterStates})
    registerStates.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", registerStates.WarningThreshold})

    registerStates.EntityData.YListKeys = []string {}

    return &(registerStates.EntityData)
}

// Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces
// Override default maximum for number of
// route-interfaces
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM route-interfaces. The type is interface{} with range:
    // 1..1100000. This attribute is mandatory.
    MaximumRouteInterfaces interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000. The default value is 300000.
    WarningThreshold interface{}
}

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetEntityData() *types.CommonEntityData {
    routeInterfaces.EntityData.YFilter = routeInterfaces.YFilter
    routeInterfaces.EntityData.YangName = "route-interfaces"
    routeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    routeInterfaces.EntityData.ParentYangName = "maximum"
    routeInterfaces.EntityData.SegmentPath = "route-interfaces"
    routeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/maximum/" + routeInterfaces.EntityData.SegmentPath
    routeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeInterfaces.EntityData.Children = types.NewOrderedMap()
    routeInterfaces.EntityData.Leafs = types.NewOrderedMap()
    routeInterfaces.EntityData.Leafs.Append("maximum-route-interfaces", types.YLeaf{"MaximumRouteInterfaces", routeInterfaces.MaximumRouteInterfaces})
    routeInterfaces.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", routeInterfaces.WarningThreshold})

    routeInterfaces.EntityData.YListKeys = []string {}

    return &(routeInterfaces.EntityData)
}

// Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache
// Override default maximum and threshold for BSR
// C-RP cache setting
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of BSR C-RP cache setting. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumCandidateRpCache interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 100.
    WarningThreshold interface{}
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetEntityData() *types.CommonEntityData {
    bsrCandidateRpCache.EntityData.YFilter = bsrCandidateRpCache.YFilter
    bsrCandidateRpCache.EntityData.YangName = "bsr-candidate-rp-cache"
    bsrCandidateRpCache.EntityData.BundleName = "cisco_ios_xr"
    bsrCandidateRpCache.EntityData.ParentYangName = "maximum"
    bsrCandidateRpCache.EntityData.SegmentPath = "bsr-candidate-rp-cache"
    bsrCandidateRpCache.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/maximum/" + bsrCandidateRpCache.EntityData.SegmentPath
    bsrCandidateRpCache.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsrCandidateRpCache.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsrCandidateRpCache.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsrCandidateRpCache.EntityData.Children = types.NewOrderedMap()
    bsrCandidateRpCache.EntityData.Leafs = types.NewOrderedMap()
    bsrCandidateRpCache.EntityData.Leafs.Append("bsr-maximum-candidate-rp-cache", types.YLeaf{"BsrMaximumCandidateRpCache", bsrCandidateRpCache.BsrMaximumCandidateRpCache})
    bsrCandidateRpCache.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", bsrCandidateRpCache.WarningThreshold})

    bsrCandidateRpCache.EntityData.YListKeys = []string {}

    return &(bsrCandidateRpCache.EntityData)
}

// Pim_DefaultContext_Ipv4_Maximum_Routes
// Override default maximum for number of routes
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum number of PIM routes. The type is interface{} with range:
    // 1..200000. This attribute is mandatory.
    MaximumRoutes interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..200000. The default value is 100000.
    WarningThreshold interface{}
}

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "maximum"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/maximum/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Leafs = types.NewOrderedMap()
    routes.EntityData.Leafs.Append("maximum-routes", types.YLeaf{"MaximumRoutes", routes.MaximumRoutes})
    routes.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", routes.WarningThreshold})

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// Pim_DefaultContext_Ipv4_Ssm
// Configure IP Multicast SSM
type Pim_DefaultContext_Ipv4_Ssm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE if SSM is disabled on this router. The type is bool. The default value
    // is false.
    Disable interface{}

    // Access list of groups enabled with SSM. The type is string with length:
    // 1..64.
    Range interface{}
}

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetEntityData() *types.CommonEntityData {
    ssm.EntityData.YFilter = ssm.YFilter
    ssm.EntityData.YangName = "ssm"
    ssm.EntityData.BundleName = "cisco_ios_xr"
    ssm.EntityData.ParentYangName = "ipv4"
    ssm.EntityData.SegmentPath = "ssm"
    ssm.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + ssm.EntityData.SegmentPath
    ssm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssm.EntityData.Children = types.NewOrderedMap()
    ssm.EntityData.Leafs = types.NewOrderedMap()
    ssm.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", ssm.Disable})
    ssm.EntityData.Leafs.Append("range", types.YLeaf{"Range", ssm.Range})

    ssm.EntityData.YListKeys = []string {}

    return &(ssm.EntityData)
}

// Pim_DefaultContext_Ipv4_Injects
// Inject Explicit PIM RPF Vector Proxy's
type Pim_DefaultContext_Ipv4_Injects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inject Explicit PIM RPF Vector Proxy's. The type is slice of
    // Pim_DefaultContext_Ipv4_Injects_Inject.
    Inject []*Pim_DefaultContext_Ipv4_Injects_Inject
}

func (injects *Pim_DefaultContext_Ipv4_Injects) GetEntityData() *types.CommonEntityData {
    injects.EntityData.YFilter = injects.YFilter
    injects.EntityData.YangName = "injects"
    injects.EntityData.BundleName = "cisco_ios_xr"
    injects.EntityData.ParentYangName = "ipv4"
    injects.EntityData.SegmentPath = "injects"
    injects.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + injects.EntityData.SegmentPath
    injects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    injects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    injects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    injects.EntityData.Children = types.NewOrderedMap()
    injects.EntityData.Children.Append("inject", types.YChild{"Inject", nil})
    for i := range injects.Inject {
        injects.EntityData.Children.Append(types.GetSegmentPath(injects.Inject[i]), types.YChild{"Inject", injects.Inject[i]})
    }
    injects.EntityData.Leafs = types.NewOrderedMap()

    injects.EntityData.YListKeys = []string {}

    return &(injects.EntityData)
}

// Pim_DefaultContext_Ipv4_Injects_Inject
// Inject Explicit PIM RPF Vector Proxy's
type Pim_DefaultContext_Ipv4_Injects_Inject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Source Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Masklen. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}

    // RPF Proxy Address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RpfProxyAddress []interface{}
}

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetEntityData() *types.CommonEntityData {
    inject.EntityData.YFilter = inject.YFilter
    inject.EntityData.YangName = "inject"
    inject.EntityData.BundleName = "cisco_ios_xr"
    inject.EntityData.ParentYangName = "injects"
    inject.EntityData.SegmentPath = "inject" + types.AddKeyToken(inject.SourceAddress, "source-address") + types.AddKeyToken(inject.PrefixLength, "prefix-length")
    inject.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/injects/" + inject.EntityData.SegmentPath
    inject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inject.EntityData.Children = types.NewOrderedMap()
    inject.EntityData.Leafs = types.NewOrderedMap()
    inject.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", inject.SourceAddress})
    inject.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", inject.PrefixLength})
    inject.EntityData.Leafs.Append("rpf-proxy-address", types.YLeaf{"RpfProxyAddress", inject.RpfProxyAddress})

    inject.EntityData.YListKeys = []string {"SourceAddress", "PrefixLength"}

    return &(inject.EntityData)
}

// Pim_DefaultContext_Ipv4_BidirRpAddresses
// Configure Bidirectional PIM Rendezvous Point
type Pim_DefaultContext_Ipv4_BidirRpAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress.
    BidirRpAddress []*Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetEntityData() *types.CommonEntityData {
    bidirRpAddresses.EntityData.YFilter = bidirRpAddresses.YFilter
    bidirRpAddresses.EntityData.YangName = "bidir-rp-addresses"
    bidirRpAddresses.EntityData.BundleName = "cisco_ios_xr"
    bidirRpAddresses.EntityData.ParentYangName = "ipv4"
    bidirRpAddresses.EntityData.SegmentPath = "bidir-rp-addresses"
    bidirRpAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + bidirRpAddresses.EntityData.SegmentPath
    bidirRpAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bidirRpAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bidirRpAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bidirRpAddresses.EntityData.Children = types.NewOrderedMap()
    bidirRpAddresses.EntityData.Children.Append("bidir-rp-address", types.YChild{"BidirRpAddress", nil})
    for i := range bidirRpAddresses.BidirRpAddress {
        bidirRpAddresses.EntityData.Children.Append(types.GetSegmentPath(bidirRpAddresses.BidirRpAddress[i]), types.YChild{"BidirRpAddress", bidirRpAddresses.BidirRpAddress[i]})
    }
    bidirRpAddresses.EntityData.Leafs = types.NewOrderedMap()

    bidirRpAddresses.EntityData.YListKeys = []string {}

    return &(bidirRpAddresses.EntityData)
}

// Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress
// Address of the Rendezvous Point
type Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetEntityData() *types.CommonEntityData {
    bidirRpAddress.EntityData.YFilter = bidirRpAddress.YFilter
    bidirRpAddress.EntityData.YangName = "bidir-rp-address"
    bidirRpAddress.EntityData.BundleName = "cisco_ios_xr"
    bidirRpAddress.EntityData.ParentYangName = "bidir-rp-addresses"
    bidirRpAddress.EntityData.SegmentPath = "bidir-rp-address" + types.AddKeyToken(bidirRpAddress.RpAddress, "rp-address")
    bidirRpAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/bidir-rp-addresses/" + bidirRpAddress.EntityData.SegmentPath
    bidirRpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bidirRpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bidirRpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bidirRpAddress.EntityData.Children = types.NewOrderedMap()
    bidirRpAddress.EntityData.Leafs = types.NewOrderedMap()
    bidirRpAddress.EntityData.Leafs.Append("rp-address", types.YLeaf{"RpAddress", bidirRpAddress.RpAddress})
    bidirRpAddress.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", bidirRpAddress.AccessListName})
    bidirRpAddress.EntityData.Leafs.Append("auto-rp-override", types.YLeaf{"AutoRpOverride", bidirRpAddress.AutoRpOverride})

    bidirRpAddress.EntityData.YListKeys = []string {"RpAddress"}

    return &(bidirRpAddress.EntityData)
}

// Pim_DefaultContext_Ipv4_Bsr
// PIM BSR configuration
type Pim_DefaultContext_Ipv4_Bsr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PIM Candidate BSR configuration.
    CandidateBsr Pim_DefaultContext_Ipv4_Bsr_CandidateBsr

    // PIM RP configuration.
    CandidateRps Pim_DefaultContext_Ipv4_Bsr_CandidateRps
}

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetEntityData() *types.CommonEntityData {
    bsr.EntityData.YFilter = bsr.YFilter
    bsr.EntityData.YangName = "bsr"
    bsr.EntityData.BundleName = "cisco_ios_xr"
    bsr.EntityData.ParentYangName = "ipv4"
    bsr.EntityData.SegmentPath = "bsr"
    bsr.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + bsr.EntityData.SegmentPath
    bsr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bsr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bsr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bsr.EntityData.Children = types.NewOrderedMap()
    bsr.EntityData.Children.Append("candidate-bsr", types.YChild{"CandidateBsr", &bsr.CandidateBsr})
    bsr.EntityData.Children.Append("candidate-rps", types.YChild{"CandidateRps", &bsr.CandidateRps})
    bsr.EntityData.Leafs = types.NewOrderedMap()

    bsr.EntityData.YListKeys = []string {}

    return &(bsr.EntityData)
}

// Pim_DefaultContext_Ipv4_Bsr_CandidateBsr
// PIM Candidate BSR configuration
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Bsr_CandidateBsr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // BSR Address configured. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    Address interface{}

    // Hash Mask Length for this candidate BSR. The type is interface{} with
    // range: 0..32. The default value is 30.
    PrefixLength interface{}

    // Priority of the Candidate BSR. The type is interface{} with range: 1..255.
    // The default value is 1.
    Priority interface{}
}

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetEntityData() *types.CommonEntityData {
    candidateBsr.EntityData.YFilter = candidateBsr.YFilter
    candidateBsr.EntityData.YangName = "candidate-bsr"
    candidateBsr.EntityData.BundleName = "cisco_ios_xr"
    candidateBsr.EntityData.ParentYangName = "bsr"
    candidateBsr.EntityData.SegmentPath = "candidate-bsr"
    candidateBsr.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/bsr/" + candidateBsr.EntityData.SegmentPath
    candidateBsr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateBsr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateBsr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateBsr.EntityData.Children = types.NewOrderedMap()
    candidateBsr.EntityData.Leafs = types.NewOrderedMap()
    candidateBsr.EntityData.Leafs.Append("address", types.YLeaf{"Address", candidateBsr.Address})
    candidateBsr.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", candidateBsr.PrefixLength})
    candidateBsr.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", candidateBsr.Priority})

    candidateBsr.EntityData.YListKeys = []string {}

    return &(candidateBsr.EntityData)
}

// Pim_DefaultContext_Ipv4_Bsr_CandidateRps
// PIM RP configuration
type Pim_DefaultContext_Ipv4_Bsr_CandidateRps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of PIM SM BSR Candidate-RP. The type is slice of
    // Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp.
    CandidateRp []*Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp
}

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetEntityData() *types.CommonEntityData {
    candidateRps.EntityData.YFilter = candidateRps.YFilter
    candidateRps.EntityData.YangName = "candidate-rps"
    candidateRps.EntityData.BundleName = "cisco_ios_xr"
    candidateRps.EntityData.ParentYangName = "bsr"
    candidateRps.EntityData.SegmentPath = "candidate-rps"
    candidateRps.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/bsr/" + candidateRps.EntityData.SegmentPath
    candidateRps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateRps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateRps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateRps.EntityData.Children = types.NewOrderedMap()
    candidateRps.EntityData.Children.Append("candidate-rp", types.YChild{"CandidateRp", nil})
    for i := range candidateRps.CandidateRp {
        candidateRps.EntityData.Children.Append(types.GetSegmentPath(candidateRps.CandidateRp[i]), types.YChild{"CandidateRp", candidateRps.CandidateRp[i]})
    }
    candidateRps.EntityData.Leafs = types.NewOrderedMap()

    candidateRps.EntityData.YListKeys = []string {}

    return &(candidateRps.EntityData)
}

// Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp
// Address of PIM SM BSR Candidate-RP
type Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address of Candidate-RP. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. SM or Bidir. The type is PimProtocolMode.
    Mode interface{}

    // Access-list specifying the group range for the Candidate-RP. The type is
    // string with length: 1..64.
    GroupList interface{}

    // Priority of the CRP. The type is interface{} with range: 1..255. The
    // default value is 192.
    Priority interface{}

    // Advertisement interval. The type is interface{} with range: 30..600. The
    // default value is 60.
    Interval interface{}
}

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetEntityData() *types.CommonEntityData {
    candidateRp.EntityData.YFilter = candidateRp.YFilter
    candidateRp.EntityData.YangName = "candidate-rp"
    candidateRp.EntityData.BundleName = "cisco_ios_xr"
    candidateRp.EntityData.ParentYangName = "candidate-rps"
    candidateRp.EntityData.SegmentPath = "candidate-rp" + types.AddKeyToken(candidateRp.Address, "address") + types.AddKeyToken(candidateRp.Mode, "mode")
    candidateRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/bsr/candidate-rps/" + candidateRp.EntityData.SegmentPath
    candidateRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateRp.EntityData.Children = types.NewOrderedMap()
    candidateRp.EntityData.Leafs = types.NewOrderedMap()
    candidateRp.EntityData.Leafs.Append("address", types.YLeaf{"Address", candidateRp.Address})
    candidateRp.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", candidateRp.Mode})
    candidateRp.EntityData.Leafs.Append("group-list", types.YLeaf{"GroupList", candidateRp.GroupList})
    candidateRp.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", candidateRp.Priority})
    candidateRp.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", candidateRp.Interval})

    candidateRp.EntityData.YListKeys = []string {"Address", "Mode"}

    return &(candidateRp.EntityData)
}

// Pim_DefaultContext_Ipv4_Mofrr
// Multicast Only Fast Re-Route
type Pim_DefaultContext_Ipv4_Mofrr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access-list specifying SG that should do RIB MOFRR. The type is string with
    // length: 1..64.
    Rib interface{}

    // Non-revertive Multicast Only Fast Re-Route. The type is interface{}.
    NonRevertive interface{}

    // Enable Multicast Only FRR. The type is interface{}.
    Enable interface{}

    // Access-list specifying SG that should do FLOW MOFRR. The type is string
    // with length: 1..64.
    Flow interface{}

    // Clone multicast joins.
    CloneJoins Pim_DefaultContext_Ipv4_Mofrr_CloneJoins

    // Clone multicast traffic.
    CloneSources Pim_DefaultContext_Ipv4_Mofrr_CloneSources
}

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetEntityData() *types.CommonEntityData {
    mofrr.EntityData.YFilter = mofrr.YFilter
    mofrr.EntityData.YangName = "mofrr"
    mofrr.EntityData.BundleName = "cisco_ios_xr"
    mofrr.EntityData.ParentYangName = "ipv4"
    mofrr.EntityData.SegmentPath = "mofrr"
    mofrr.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + mofrr.EntityData.SegmentPath
    mofrr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mofrr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mofrr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mofrr.EntityData.Children = types.NewOrderedMap()
    mofrr.EntityData.Children.Append("clone-joins", types.YChild{"CloneJoins", &mofrr.CloneJoins})
    mofrr.EntityData.Children.Append("clone-sources", types.YChild{"CloneSources", &mofrr.CloneSources})
    mofrr.EntityData.Leafs = types.NewOrderedMap()
    mofrr.EntityData.Leafs.Append("rib", types.YLeaf{"Rib", mofrr.Rib})
    mofrr.EntityData.Leafs.Append("non-revertive", types.YLeaf{"NonRevertive", mofrr.NonRevertive})
    mofrr.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", mofrr.Enable})
    mofrr.EntityData.Leafs.Append("flow", types.YLeaf{"Flow", mofrr.Flow})

    mofrr.EntityData.YListKeys = []string {}

    return &(mofrr.EntityData)
}

// Pim_DefaultContext_Ipv4_Mofrr_CloneJoins
// Clone multicast joins
type Pim_DefaultContext_Ipv4_Mofrr_CloneJoins struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clone S,G joins as S1,G joins and S2,G joins. The type is slice of
    // Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin.
    CloneJoin []*Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin
}

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetEntityData() *types.CommonEntityData {
    cloneJoins.EntityData.YFilter = cloneJoins.YFilter
    cloneJoins.EntityData.YangName = "clone-joins"
    cloneJoins.EntityData.BundleName = "cisco_ios_xr"
    cloneJoins.EntityData.ParentYangName = "mofrr"
    cloneJoins.EntityData.SegmentPath = "clone-joins"
    cloneJoins.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/mofrr/" + cloneJoins.EntityData.SegmentPath
    cloneJoins.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cloneJoins.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cloneJoins.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cloneJoins.EntityData.Children = types.NewOrderedMap()
    cloneJoins.EntityData.Children.Append("clone-join", types.YChild{"CloneJoin", nil})
    for i := range cloneJoins.CloneJoin {
        cloneJoins.EntityData.Children.Append(types.GetSegmentPath(cloneJoins.CloneJoin[i]), types.YChild{"CloneJoin", cloneJoins.CloneJoin[i]})
    }
    cloneJoins.EntityData.Leafs = types.NewOrderedMap()

    cloneJoins.EntityData.YListKeys = []string {}

    return &(cloneJoins.EntityData)
}

// Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin
// Clone S,G joins as S1,G joins and S2,G joins
type Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Original source address (S). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Source interface{}

    // This attribute is a key. Primary cloned address (S1). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Primary interface{}

    // This attribute is a key. Backup cloned address (S2). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Backup interface{}

    // This attribute is a key. Mask length. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}
}

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetEntityData() *types.CommonEntityData {
    cloneJoin.EntityData.YFilter = cloneJoin.YFilter
    cloneJoin.EntityData.YangName = "clone-join"
    cloneJoin.EntityData.BundleName = "cisco_ios_xr"
    cloneJoin.EntityData.ParentYangName = "clone-joins"
    cloneJoin.EntityData.SegmentPath = "clone-join" + types.AddKeyToken(cloneJoin.Source, "source") + types.AddKeyToken(cloneJoin.Primary, "primary") + types.AddKeyToken(cloneJoin.Backup, "backup") + types.AddKeyToken(cloneJoin.PrefixLength, "prefix-length")
    cloneJoin.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/mofrr/clone-joins/" + cloneJoin.EntityData.SegmentPath
    cloneJoin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cloneJoin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cloneJoin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cloneJoin.EntityData.Children = types.NewOrderedMap()
    cloneJoin.EntityData.Leafs = types.NewOrderedMap()
    cloneJoin.EntityData.Leafs.Append("source", types.YLeaf{"Source", cloneJoin.Source})
    cloneJoin.EntityData.Leafs.Append("primary", types.YLeaf{"Primary", cloneJoin.Primary})
    cloneJoin.EntityData.Leafs.Append("backup", types.YLeaf{"Backup", cloneJoin.Backup})
    cloneJoin.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", cloneJoin.PrefixLength})

    cloneJoin.EntityData.YListKeys = []string {"Source", "Primary", "Backup", "PrefixLength"}

    return &(cloneJoin.EntityData)
}

// Pim_DefaultContext_Ipv4_Mofrr_CloneSources
// Clone multicast traffic
type Pim_DefaultContext_Ipv4_Mofrr_CloneSources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clone S,G traffic as S1,G traffic and S2,G traffic. The type is slice of
    // Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource.
    CloneSource []*Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource
}

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetEntityData() *types.CommonEntityData {
    cloneSources.EntityData.YFilter = cloneSources.YFilter
    cloneSources.EntityData.YangName = "clone-sources"
    cloneSources.EntityData.BundleName = "cisco_ios_xr"
    cloneSources.EntityData.ParentYangName = "mofrr"
    cloneSources.EntityData.SegmentPath = "clone-sources"
    cloneSources.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/mofrr/" + cloneSources.EntityData.SegmentPath
    cloneSources.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cloneSources.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cloneSources.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cloneSources.EntityData.Children = types.NewOrderedMap()
    cloneSources.EntityData.Children.Append("clone-source", types.YChild{"CloneSource", nil})
    for i := range cloneSources.CloneSource {
        cloneSources.EntityData.Children.Append(types.GetSegmentPath(cloneSources.CloneSource[i]), types.YChild{"CloneSource", cloneSources.CloneSource[i]})
    }
    cloneSources.EntityData.Leafs = types.NewOrderedMap()

    cloneSources.EntityData.YListKeys = []string {}

    return &(cloneSources.EntityData)
}

// Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource
// Clone S,G traffic as S1,G traffic and S2,G
// traffic
type Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Original source address (S). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Source interface{}

    // This attribute is a key. Primary cloned address (S1). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Primary interface{}

    // This attribute is a key. Backup cloned address (S2). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Backup interface{}

    // This attribute is a key. Mask length. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}
}

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetEntityData() *types.CommonEntityData {
    cloneSource.EntityData.YFilter = cloneSource.YFilter
    cloneSource.EntityData.YangName = "clone-source"
    cloneSource.EntityData.BundleName = "cisco_ios_xr"
    cloneSource.EntityData.ParentYangName = "clone-sources"
    cloneSource.EntityData.SegmentPath = "clone-source" + types.AddKeyToken(cloneSource.Source, "source") + types.AddKeyToken(cloneSource.Primary, "primary") + types.AddKeyToken(cloneSource.Backup, "backup") + types.AddKeyToken(cloneSource.PrefixLength, "prefix-length")
    cloneSource.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/mofrr/clone-sources/" + cloneSource.EntityData.SegmentPath
    cloneSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cloneSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cloneSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cloneSource.EntityData.Children = types.NewOrderedMap()
    cloneSource.EntityData.Leafs = types.NewOrderedMap()
    cloneSource.EntityData.Leafs.Append("source", types.YLeaf{"Source", cloneSource.Source})
    cloneSource.EntityData.Leafs.Append("primary", types.YLeaf{"Primary", cloneSource.Primary})
    cloneSource.EntityData.Leafs.Append("backup", types.YLeaf{"Backup", cloneSource.Backup})
    cloneSource.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", cloneSource.PrefixLength})

    cloneSource.EntityData.YListKeys = []string {"Source", "Primary", "Backup", "PrefixLength"}

    return &(cloneSource.EntityData)
}

// Pim_DefaultContext_Ipv4_Paths
// Inject PIM RPF Vector Proxy's
type Pim_DefaultContext_Ipv4_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inject PIM RPF Vector Proxy's. The type is slice of
    // Pim_DefaultContext_Ipv4_Paths_Path.
    Path []*Pim_DefaultContext_Ipv4_Paths_Path
}

func (paths *Pim_DefaultContext_Ipv4_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "ipv4"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + paths.EntityData.SegmentPath
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = types.NewOrderedMap()
    paths.EntityData.Children.Append("path", types.YChild{"Path", nil})
    for i := range paths.Path {
        paths.EntityData.Children.Append(types.GetSegmentPath(paths.Path[i]), types.YChild{"Path", paths.Path[i]})
    }
    paths.EntityData.Leafs = types.NewOrderedMap()

    paths.EntityData.YListKeys = []string {}

    return &(paths.EntityData)
}

// Pim_DefaultContext_Ipv4_Paths_Path
// Inject PIM RPF Vector Proxy's
type Pim_DefaultContext_Ipv4_Paths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Source Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Masklen. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}

    // RPF Proxy Address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RpfProxyAddress []interface{}
}

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "paths"
    path.EntityData.SegmentPath = "path" + types.AddKeyToken(path.SourceAddress, "source-address") + types.AddKeyToken(path.PrefixLength, "prefix-length")
    path.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/paths/" + path.EntityData.SegmentPath
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", path.SourceAddress})
    path.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", path.PrefixLength})
    path.EntityData.Leafs.Append("rpf-proxy-address", types.YLeaf{"RpfProxyAddress", path.RpfProxyAddress})

    path.EntityData.YListKeys = []string {"SourceAddress", "PrefixLength"}

    return &(path.EntityData)
}

// Pim_DefaultContext_Ipv4_AllowRp
// Enable allow-rp filtering for SM joins
// This type is a presence type.
type Pim_DefaultContext_Ipv4_AllowRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Access-list specifiying applicable RPs. The type is string with length:
    // 1..64.
    RpListName interface{}

    // Access-list specifiying applicable groups. The type is string with length:
    // 1..64.
    GroupListName interface{}
}

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetEntityData() *types.CommonEntityData {
    allowRp.EntityData.YFilter = allowRp.YFilter
    allowRp.EntityData.YangName = "allow-rp"
    allowRp.EntityData.BundleName = "cisco_ios_xr"
    allowRp.EntityData.ParentYangName = "ipv4"
    allowRp.EntityData.SegmentPath = "allow-rp"
    allowRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + allowRp.EntityData.SegmentPath
    allowRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowRp.EntityData.Children = types.NewOrderedMap()
    allowRp.EntityData.Leafs = types.NewOrderedMap()
    allowRp.EntityData.Leafs.Append("rp-list-name", types.YLeaf{"RpListName", allowRp.RpListName})
    allowRp.EntityData.Leafs.Append("group-list-name", types.YLeaf{"GroupListName", allowRp.GroupListName})

    allowRp.EntityData.YListKeys = []string {}

    return &(allowRp.EntityData)
}

// Pim_DefaultContext_Ipv4_Convergence
// Configure convergence parameters
type Pim_DefaultContext_Ipv4_Convergence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampen first join if RPF path is through one of the downstream neighbor.
    // The type is interface{} with range: 0..15. Units are second.
    RpfConflictJoinDelay interface{}

    // Delay prunes if route join state transitions to not-joined on link down.
    // The type is interface{} with range: 0..60. Units are second.
    LinkDownPruneDelay interface{}
}

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetEntityData() *types.CommonEntityData {
    convergence.EntityData.YFilter = convergence.YFilter
    convergence.EntityData.YangName = "convergence"
    convergence.EntityData.BundleName = "cisco_ios_xr"
    convergence.EntityData.ParentYangName = "ipv4"
    convergence.EntityData.SegmentPath = "convergence"
    convergence.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-pim-cfg:pim/default-context/ipv4/" + convergence.EntityData.SegmentPath
    convergence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    convergence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    convergence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    convergence.EntityData.Children = types.NewOrderedMap()
    convergence.EntityData.Leafs = types.NewOrderedMap()
    convergence.EntityData.Leafs.Append("rpf-conflict-join-delay", types.YLeaf{"RpfConflictJoinDelay", convergence.RpfConflictJoinDelay})
    convergence.EntityData.Leafs.Append("link-down-prune-delay", types.YLeaf{"LinkDownPruneDelay", convergence.LinkDownPruneDelay})

    convergence.EntityData.YListKeys = []string {}

    return &(convergence.EntityData)
}

