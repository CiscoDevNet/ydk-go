// This module contains a collection of YANG definitions
// for Cisco IOS-XR eigrp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   eigrp: Configure Neighbor
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package eigrp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package eigrp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-eigrp-cfg eigrp}", reflect.TypeOf(Eigrp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-eigrp-cfg:eigrp", reflect.TypeOf(Eigrp{}))
}

// EigrpMetricVersion represents Eigrp metric version
type EigrpMetricVersion string

const (
    // 64-bit metrics
    EigrpMetricVersion_Y_64bit EigrpMetricVersion = "64bit"

    // 32-bit metrics
    EigrpMetricVersion_Y_32bit EigrpMetricVersion = "32bit"
)

// EigrpSoo represents Eigrp soo
type EigrpSoo string

const (
    // AS:nn format
    EigrpSoo_as EigrpSoo = "as"

    // IPV4Address:nn format
    EigrpSoo_ipv4_address EigrpSoo = "ipv4-address"

    // AS2.AS:nn format
    EigrpSoo_four_byte_as EigrpSoo = "four-byte-as"
)

// EigrpRedistProto represents Eigrp redist proto
type EigrpRedistProto string

const (
    // BGP routes
    EigrpRedistProto_bgp EigrpRedistProto = "bgp"

    // Connected routes
    EigrpRedistProto_connected EigrpRedistProto = "connected"

    // EIGRP routes
    EigrpRedistProto_eigrp EigrpRedistProto = "eigrp"

    // ISIS routes
    EigrpRedistProto_isis EigrpRedistProto = "isis"

    // OSPF routes
    EigrpRedistProto_ospf EigrpRedistProto = "ospf"

    // RIP routes
    EigrpRedistProto_rip EigrpRedistProto = "rip"

    // Static routes
    EigrpRedistProto_static EigrpRedistProto = "static"

    // OSPFv3 routes
    EigrpRedistProto_ospfv3 EigrpRedistProto = "ospfv3"

    // Subscriber routes
    EigrpRedistProto_subscriber EigrpRedistProto = "subscriber"

    // Application routes
    EigrpRedistProto_application EigrpRedistProto = "application"

    // Mobile routes
    EigrpRedistProto_mobile EigrpRedistProto = "mobile"
)

// EigrpStub represents Eigrp stub
type EigrpStub string

const (
    // receive only
    EigrpStub_receive_only EigrpStub = "receive-only"

    // filter some routes
    EigrpStub_filtered EigrpStub = "filtered"
)

// EigrpDelayUnit represents Eigrp delay unit
type EigrpDelayUnit string

const (
    // Delay in 10's of microsecond
    EigrpDelayUnit_ten_microsecond EigrpDelayUnit = "ten-microsecond"

    // Delay in picosecond
    EigrpDelayUnit_picosecond EigrpDelayUnit = "picosecond"
)

// EigrpMet represents Eigrp met
type EigrpMet string

const (
    // Metric maxhops
    EigrpMet_maximum_hops EigrpMet = "maximum-hops"

    // Metric weights
    EigrpMet_weights EigrpMet = "weights"

    // Metric ribscale
    EigrpMet_rib_scale EigrpMet = "rib-scale"

    // Metric version
    EigrpMet_version EigrpMet = "version"
)

// EigrpDir represents Eigrp dir
type EigrpDir string

const (
    // inbound direction
    EigrpDir_in EigrpDir = "in"

    // outbound direction
    EigrpDir_out EigrpDir = "out"
)

// EigrpTimer represents Eigrp timer
type EigrpTimer string

const (
    // Active time
    EigrpTimer_active EigrpTimer = "active"

    // Route-hold time
    EigrpTimer_route_hold EigrpTimer = "route-hold"

    // Signal time
    EigrpTimer_signal EigrpTimer = "signal"

    // Converge time
    EigrpTimer_converge EigrpTimer = "converge"
)

// Eigrp
// Configure Neighbor
type Eigrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process related configuration.
    Processes Eigrp_Processes
}

func (eigrp *Eigrp) GetEntityData() *types.CommonEntityData {
    eigrp.EntityData.YFilter = eigrp.YFilter
    eigrp.EntityData.YangName = "eigrp"
    eigrp.EntityData.BundleName = "cisco_ios_xr"
    eigrp.EntityData.ParentYangName = "Cisco-IOS-XR-eigrp-cfg"
    eigrp.EntityData.SegmentPath = "Cisco-IOS-XR-eigrp-cfg:eigrp"
    eigrp.EntityData.AbsolutePath = eigrp.EntityData.SegmentPath
    eigrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrp.EntityData.Children = types.NewOrderedMap()
    eigrp.EntityData.Children.Append("processes", types.YChild{"Processes", &eigrp.Processes})
    eigrp.EntityData.Leafs = types.NewOrderedMap()

    eigrp.EntityData.YListKeys = []string {}

    return &(eigrp.EntityData)
}

// Eigrp_Processes
// Process related configuration
type Eigrp_Processes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a particular EIGRP process. The type is slice of
    // Eigrp_Processes_Process.
    Process []*Eigrp_Processes_Process
}

func (processes *Eigrp_Processes) GetEntityData() *types.CommonEntityData {
    processes.EntityData.YFilter = processes.YFilter
    processes.EntityData.YangName = "processes"
    processes.EntityData.BundleName = "cisco_ios_xr"
    processes.EntityData.ParentYangName = "eigrp"
    processes.EntityData.SegmentPath = "processes"
    processes.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/" + processes.EntityData.SegmentPath
    processes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processes.EntityData.Children = types.NewOrderedMap()
    processes.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range processes.Process {
        processes.EntityData.Children.Append(types.GetSegmentPath(processes.Process[i]), types.YChild{"Process", processes.Process[i]})
    }
    processes.EntityData.Leafs = types.NewOrderedMap()

    processes.EntityData.YListKeys = []string {}

    return &(processes.EntityData)
}

// Eigrp_Processes_Process
// Configuration for a particular EIGRP process.
type Eigrp_Processes_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. AS number (1 - 65535) or Virutual instance name of
    // the EIGRP process. The type is string with length: 1..32.
    ProcessId interface{}

    // Disable NSF for all address families under all VRF's. It takes precedence
    // over the NSF related configuration in the address family submode. The type
    // is interface{}.
    NsfDisable interface{}

    // List of VRFs.
    Vrfs Eigrp_Processes_Process_Vrfs

    // Default VRF configuration.Deletion of this object also causes deletion of
    // all objectsunder 'Process' associated with this process instance.
    DefaultVrf Eigrp_Processes_Process_DefaultVrf
}

func (process *Eigrp_Processes_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "processes"
    process.EntityData.SegmentPath = "process" + types.AddKeyToken(process.ProcessId, "process-id")
    process.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/" + process.EntityData.SegmentPath
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &process.Vrfs})
    process.EntityData.Children.Append("default-vrf", types.YChild{"DefaultVrf", &process.DefaultVrf})
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", process.ProcessId})
    process.EntityData.Leafs.Append("nsf-disable", types.YLeaf{"NsfDisable", process.NsfDisable})

    process.EntityData.YListKeys = []string {"ProcessId"}

    return &(process.EntityData)
}

// Eigrp_Processes_Process_Vrfs
// List of VRFs
type Eigrp_Processes_Process_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Non-default VRF configuration.Deletion of this object also causes deletion
    // of all objectsunder 'VRF' associated with this VRF instance. The type is
    // slice of Eigrp_Processes_Process_Vrfs_Vrf.
    Vrf []*Eigrp_Processes_Process_Vrfs_Vrf
}

func (vrfs *Eigrp_Processes_Process_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "process"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/" + vrfs.EntityData.SegmentPath
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

// Eigrp_Processes_Process_Vrfs_Vrf
// Non-default VRF configuration.Deletion of
// this object also causes deletion of all
// objectsunder 'VRF' associated with this VRF
// instance.
type Eigrp_Processes_Process_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Enable a non-default VRF under the EIGRP process. The type is interface{}.
    Enable interface{}

    // Address family list in a VRF.
    Afs Eigrp_Processes_Process_Vrfs_Vrf_Afs
}

func (vrf *Eigrp_Processes_Process_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("afs", types.YChild{"Afs", &vrf.Afs})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", vrf.Enable})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs
// Address family list in a VRF
type Eigrp_Processes_Process_Vrfs_Vrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration under an AF in a non-default VRF context.Deletion of this
    // object also causes deletion of all objectsunder 'AF' associated with this
    // AF instance. The type is slice of Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af.
    Af []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af
}

func (afs *Eigrp_Processes_Process_Vrfs_Vrf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "vrf"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/" + afs.EntityData.SegmentPath
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

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af
// Configuration under an AF in a non-default
// VRF context.Deletion of this object also
// causes deletion of all objectsunder 'AF'
// associated with this AF instance.
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address Family. The type is EigrpAf.
    AfName interface{}

    // Enable an Address Family under a non-default VRF. The type is interface{}.
    Enable interface{}

    // Enable Auto Summarization. The type is interface{}.
    AutoSummary interface{}

    // number of paths. The type is interface{} with range: 1..32.
    MaximumPaths interface{}

    // Set router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Enable/Disable neighbor state change warnings . The type is interface{}.
    LogNeighborWarnings interface{}

    // Set the autonomous system of a VRF. The type is interface{} with range:
    // 1..65535.
    AutonomousSystem interface{}

    // Control load balancing variance. The type is interface{} with range:
    // 1..128.
    Variance interface{}

    // Disable NSF for this address family under this VRF. The type is
    // interface{}.
    NsfDisable interface{}

    // Suppress routing updates on all interfaces. The type is interface{}.
    PassiveInterfaceDefault interface{}

    // Enable/Disable logginf of neighbor state changes. The type is interface{}.
    LogNeighborChanges interface{}

    // Maximum number of IP prefixes acceptable in aggregate, from neighbors.
    AllNeighborsMaximumPrefix Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_AllNeighborsMaximumPrefix

    // Maximum number of prefixes redistributed to protocol.
    RedistMaximumPrefix Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_RedistMaximumPrefix

    // List of neighbors with prefix limits.
    NeighborMaximumPrefixes Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_NeighborMaximumPrefixes

    // Maximum number of IP prefixes acceptable in aggregate.
    MaximumPrefix Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_MaximumPrefix

    // EIGRP stub configuration.
    Stub Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Stub

    // List of redistributed protocols.
    Redistributes Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes

    // Inbound and outbound filter policies.
    FilterPolicies Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_FilterPolicies

    // Set metric of redistributed routes.
    DefaultMetric Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_DefaultMetric

    // List of metric change behaviours.
    Metrics Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Metrics

    // List of timer configurations.
    Timers Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Timers

    // Candidate default policy table.
    DefaultAccepts Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_DefaultAccepts

    // List of interfaces.
    Interfaces Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces

    // Set distance for EIGRP routes.
    Distance Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Distance
}

func (af *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("all-neighbors-maximum-prefix", types.YChild{"AllNeighborsMaximumPrefix", &af.AllNeighborsMaximumPrefix})
    af.EntityData.Children.Append("redist-maximum-prefix", types.YChild{"RedistMaximumPrefix", &af.RedistMaximumPrefix})
    af.EntityData.Children.Append("neighbor-maximum-prefixes", types.YChild{"NeighborMaximumPrefixes", &af.NeighborMaximumPrefixes})
    af.EntityData.Children.Append("maximum-prefix", types.YChild{"MaximumPrefix", &af.MaximumPrefix})
    af.EntityData.Children.Append("stub", types.YChild{"Stub", &af.Stub})
    af.EntityData.Children.Append("redistributes", types.YChild{"Redistributes", &af.Redistributes})
    af.EntityData.Children.Append("filter-policies", types.YChild{"FilterPolicies", &af.FilterPolicies})
    af.EntityData.Children.Append("default-metric", types.YChild{"DefaultMetric", &af.DefaultMetric})
    af.EntityData.Children.Append("metrics", types.YChild{"Metrics", &af.Metrics})
    af.EntityData.Children.Append("timers", types.YChild{"Timers", &af.Timers})
    af.EntityData.Children.Append("default-accepts", types.YChild{"DefaultAccepts", &af.DefaultAccepts})
    af.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &af.Interfaces})
    af.EntityData.Children.Append("distance", types.YChild{"Distance", &af.Distance})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", af.Enable})
    af.EntityData.Leafs.Append("auto-summary", types.YLeaf{"AutoSummary", af.AutoSummary})
    af.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", af.MaximumPaths})
    af.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", af.RouterId})
    af.EntityData.Leafs.Append("log-neighbor-warnings", types.YLeaf{"LogNeighborWarnings", af.LogNeighborWarnings})
    af.EntityData.Leafs.Append("autonomous-system", types.YLeaf{"AutonomousSystem", af.AutonomousSystem})
    af.EntityData.Leafs.Append("variance", types.YLeaf{"Variance", af.Variance})
    af.EntityData.Leafs.Append("nsf-disable", types.YLeaf{"NsfDisable", af.NsfDisable})
    af.EntityData.Leafs.Append("passive-interface-default", types.YLeaf{"PassiveInterfaceDefault", af.PassiveInterfaceDefault})
    af.EntityData.Leafs.Append("log-neighbor-changes", types.YLeaf{"LogNeighborChanges", af.LogNeighborChanges})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_AllNeighborsMaximumPrefix
// Maximum number of IP prefixes acceptable
// in aggregate, from neighbors
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_AllNeighborsMaximumPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of IP prefixes for maximum-prefix limit. The type is interface{}
    // with range: 1..4294967295.
    MaxPrefix interface{}

    // Configure threshold percentage for warnings. The type is interface{} with
    // range: 1..100. Units are percentage.
    Threshold interface{}

    // Enable decay penalty to be applied to time period. The type is bool.
    Dampened interface{}

    // Time to reset restart count. The type is interface{} with range:
    // 0..4294967295.
    ResetTime interface{}

    // Shutdown time after hitting max-prefix limit. The type is interface{} with
    // range: 0..4294967295.
    Restart interface{}

    // Restart count after hitting max-prefix limit. The type is interface{} with
    // range: 0..4294967295.
    RestartCount interface{}

    // Only a warning is logged when prefix limit is reached. The type is bool.
    WarningOnly interface{}
}

func (allNeighborsMaximumPrefix *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_AllNeighborsMaximumPrefix) GetEntityData() *types.CommonEntityData {
    allNeighborsMaximumPrefix.EntityData.YFilter = allNeighborsMaximumPrefix.YFilter
    allNeighborsMaximumPrefix.EntityData.YangName = "all-neighbors-maximum-prefix"
    allNeighborsMaximumPrefix.EntityData.BundleName = "cisco_ios_xr"
    allNeighborsMaximumPrefix.EntityData.ParentYangName = "af"
    allNeighborsMaximumPrefix.EntityData.SegmentPath = "all-neighbors-maximum-prefix"
    allNeighborsMaximumPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/" + allNeighborsMaximumPrefix.EntityData.SegmentPath
    allNeighborsMaximumPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allNeighborsMaximumPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allNeighborsMaximumPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allNeighborsMaximumPrefix.EntityData.Children = types.NewOrderedMap()
    allNeighborsMaximumPrefix.EntityData.Leafs = types.NewOrderedMap()
    allNeighborsMaximumPrefix.EntityData.Leafs.Append("max-prefix", types.YLeaf{"MaxPrefix", allNeighborsMaximumPrefix.MaxPrefix})
    allNeighborsMaximumPrefix.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", allNeighborsMaximumPrefix.Threshold})
    allNeighborsMaximumPrefix.EntityData.Leafs.Append("dampened", types.YLeaf{"Dampened", allNeighborsMaximumPrefix.Dampened})
    allNeighborsMaximumPrefix.EntityData.Leafs.Append("reset-time", types.YLeaf{"ResetTime", allNeighborsMaximumPrefix.ResetTime})
    allNeighborsMaximumPrefix.EntityData.Leafs.Append("restart", types.YLeaf{"Restart", allNeighborsMaximumPrefix.Restart})
    allNeighborsMaximumPrefix.EntityData.Leafs.Append("restart-count", types.YLeaf{"RestartCount", allNeighborsMaximumPrefix.RestartCount})
    allNeighborsMaximumPrefix.EntityData.Leafs.Append("warning-only", types.YLeaf{"WarningOnly", allNeighborsMaximumPrefix.WarningOnly})

    allNeighborsMaximumPrefix.EntityData.YListKeys = []string {}

    return &(allNeighborsMaximumPrefix.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_RedistMaximumPrefix
// Maximum number of prefixes redistributed
// to protocol
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_RedistMaximumPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of IP prefixes for maximum-prefix limit. The type is interface{}
    // with range: 1..4294967295.
    MaxPrefix interface{}

    // Configure threshold percentage for warnings. The type is interface{} with
    // range: 1..100. Units are percentage.
    Threshold interface{}

    // Enable decay penalty to be applied to time period. The type is interface{}
    // with range: 0..4294967295.
    Dampened interface{}

    // Time to reset restart count. The type is interface{} with range:
    // 0..4294967295.
    ResetTime interface{}

    // Shutdown time after hitting max-prefix limit. The type is interface{} with
    // range: 0..4294967295.
    Restart interface{}

    // Restart count after hitting max-prefix limit. The type is interface{} with
    // range: 0..4294967295.
    RestartCount interface{}

    // Only a warning is logged when prefix limit is reached. The type is bool.
    WarningOnly interface{}
}

func (redistMaximumPrefix *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_RedistMaximumPrefix) GetEntityData() *types.CommonEntityData {
    redistMaximumPrefix.EntityData.YFilter = redistMaximumPrefix.YFilter
    redistMaximumPrefix.EntityData.YangName = "redist-maximum-prefix"
    redistMaximumPrefix.EntityData.BundleName = "cisco_ios_xr"
    redistMaximumPrefix.EntityData.ParentYangName = "af"
    redistMaximumPrefix.EntityData.SegmentPath = "redist-maximum-prefix"
    redistMaximumPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/" + redistMaximumPrefix.EntityData.SegmentPath
    redistMaximumPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistMaximumPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistMaximumPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistMaximumPrefix.EntityData.Children = types.NewOrderedMap()
    redistMaximumPrefix.EntityData.Leafs = types.NewOrderedMap()
    redistMaximumPrefix.EntityData.Leafs.Append("max-prefix", types.YLeaf{"MaxPrefix", redistMaximumPrefix.MaxPrefix})
    redistMaximumPrefix.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", redistMaximumPrefix.Threshold})
    redistMaximumPrefix.EntityData.Leafs.Append("dampened", types.YLeaf{"Dampened", redistMaximumPrefix.Dampened})
    redistMaximumPrefix.EntityData.Leafs.Append("reset-time", types.YLeaf{"ResetTime", redistMaximumPrefix.ResetTime})
    redistMaximumPrefix.EntityData.Leafs.Append("restart", types.YLeaf{"Restart", redistMaximumPrefix.Restart})
    redistMaximumPrefix.EntityData.Leafs.Append("restart-count", types.YLeaf{"RestartCount", redistMaximumPrefix.RestartCount})
    redistMaximumPrefix.EntityData.Leafs.Append("warning-only", types.YLeaf{"WarningOnly", redistMaximumPrefix.WarningOnly})

    redistMaximumPrefix.EntityData.YListKeys = []string {}

    return &(redistMaximumPrefix.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_NeighborMaximumPrefixes
// List of neighbors with prefix limits
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_NeighborMaximumPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor prefix limits configuration. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_NeighborMaximumPrefixes_NeighborMaximumPrefix.
    NeighborMaximumPrefix []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_NeighborMaximumPrefixes_NeighborMaximumPrefix
}

func (neighborMaximumPrefixes *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_NeighborMaximumPrefixes) GetEntityData() *types.CommonEntityData {
    neighborMaximumPrefixes.EntityData.YFilter = neighborMaximumPrefixes.YFilter
    neighborMaximumPrefixes.EntityData.YangName = "neighbor-maximum-prefixes"
    neighborMaximumPrefixes.EntityData.BundleName = "cisco_ios_xr"
    neighborMaximumPrefixes.EntityData.ParentYangName = "af"
    neighborMaximumPrefixes.EntityData.SegmentPath = "neighbor-maximum-prefixes"
    neighborMaximumPrefixes.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/" + neighborMaximumPrefixes.EntityData.SegmentPath
    neighborMaximumPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborMaximumPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborMaximumPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborMaximumPrefixes.EntityData.Children = types.NewOrderedMap()
    neighborMaximumPrefixes.EntityData.Children.Append("neighbor-maximum-prefix", types.YChild{"NeighborMaximumPrefix", nil})
    for i := range neighborMaximumPrefixes.NeighborMaximumPrefix {
        neighborMaximumPrefixes.EntityData.Children.Append(types.GetSegmentPath(neighborMaximumPrefixes.NeighborMaximumPrefix[i]), types.YChild{"NeighborMaximumPrefix", neighborMaximumPrefixes.NeighborMaximumPrefix[i]})
    }
    neighborMaximumPrefixes.EntityData.Leafs = types.NewOrderedMap()

    neighborMaximumPrefixes.EntityData.YListKeys = []string {}

    return &(neighborMaximumPrefixes.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_NeighborMaximumPrefixes_NeighborMaximumPrefix
// Neighbor prefix limits configuration
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_NeighborMaximumPrefixes_NeighborMaximumPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NeighborAddress interface{}

    // Number of IP prefixes for maximum-prefix limit. The type is interface{}
    // with range: 1..4294967295.
    MaxPrefix interface{}

    // Configure threshold percentage for warnings. The type is interface{} with
    // range: 1..100. Units are percentage.
    Threshold interface{}

    // Only a warning is logged when prefix limit is reached. The type is bool.
    WarningOnly interface{}
}

func (neighborMaximumPrefix *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_NeighborMaximumPrefixes_NeighborMaximumPrefix) GetEntityData() *types.CommonEntityData {
    neighborMaximumPrefix.EntityData.YFilter = neighborMaximumPrefix.YFilter
    neighborMaximumPrefix.EntityData.YangName = "neighbor-maximum-prefix"
    neighborMaximumPrefix.EntityData.BundleName = "cisco_ios_xr"
    neighborMaximumPrefix.EntityData.ParentYangName = "neighbor-maximum-prefixes"
    neighborMaximumPrefix.EntityData.SegmentPath = "neighbor-maximum-prefix" + types.AddKeyToken(neighborMaximumPrefix.NeighborAddress, "neighbor-address")
    neighborMaximumPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/neighbor-maximum-prefixes/" + neighborMaximumPrefix.EntityData.SegmentPath
    neighborMaximumPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborMaximumPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborMaximumPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborMaximumPrefix.EntityData.Children = types.NewOrderedMap()
    neighborMaximumPrefix.EntityData.Leafs = types.NewOrderedMap()
    neighborMaximumPrefix.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", neighborMaximumPrefix.NeighborAddress})
    neighborMaximumPrefix.EntityData.Leafs.Append("max-prefix", types.YLeaf{"MaxPrefix", neighborMaximumPrefix.MaxPrefix})
    neighborMaximumPrefix.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", neighborMaximumPrefix.Threshold})
    neighborMaximumPrefix.EntityData.Leafs.Append("warning-only", types.YLeaf{"WarningOnly", neighborMaximumPrefix.WarningOnly})

    neighborMaximumPrefix.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(neighborMaximumPrefix.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_MaximumPrefix
// Maximum number of IP prefixes acceptable
// in aggregate
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_MaximumPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of IP prefixes for maximum-prefix limit. The type is interface{}
    // with range: 1..4294967295.
    MaxPrefix interface{}

    // Configure threshold percentage for warnings. The type is interface{} with
    // range: 1..100. Units are percentage.
    Threshold interface{}

    // Enable decay penalty to be applied to time period. The type is bool.
    Dampened interface{}

    // Time to reset restart count. The type is interface{} with range:
    // 0..4294967295.
    ResetTime interface{}

    // Shutdown time after hitting max-prefix limit. The type is interface{} with
    // range: 0..4294967295.
    Restart interface{}

    // Restart count after hitting max-prefix limit. The type is interface{} with
    // range: 0..4294967295.
    RestartCount interface{}

    // Only a warning is logged when prefix limit is reached. The type is bool.
    WarningOnly interface{}
}

func (maximumPrefix *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_MaximumPrefix) GetEntityData() *types.CommonEntityData {
    maximumPrefix.EntityData.YFilter = maximumPrefix.YFilter
    maximumPrefix.EntityData.YangName = "maximum-prefix"
    maximumPrefix.EntityData.BundleName = "cisco_ios_xr"
    maximumPrefix.EntityData.ParentYangName = "af"
    maximumPrefix.EntityData.SegmentPath = "maximum-prefix"
    maximumPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/" + maximumPrefix.EntityData.SegmentPath
    maximumPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumPrefix.EntityData.Children = types.NewOrderedMap()
    maximumPrefix.EntityData.Leafs = types.NewOrderedMap()
    maximumPrefix.EntityData.Leafs.Append("max-prefix", types.YLeaf{"MaxPrefix", maximumPrefix.MaxPrefix})
    maximumPrefix.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", maximumPrefix.Threshold})
    maximumPrefix.EntityData.Leafs.Append("dampened", types.YLeaf{"Dampened", maximumPrefix.Dampened})
    maximumPrefix.EntityData.Leafs.Append("reset-time", types.YLeaf{"ResetTime", maximumPrefix.ResetTime})
    maximumPrefix.EntityData.Leafs.Append("restart", types.YLeaf{"Restart", maximumPrefix.Restart})
    maximumPrefix.EntityData.Leafs.Append("restart-count", types.YLeaf{"RestartCount", maximumPrefix.RestartCount})
    maximumPrefix.EntityData.Leafs.Append("warning-only", types.YLeaf{"WarningOnly", maximumPrefix.WarningOnly})

    maximumPrefix.EntityData.YListKeys = []string {}

    return &(maximumPrefix.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Stub
// EIGRP stub configuration
// This type is a presence type.
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Stub struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Stub config type. The type is EigrpStub.
    Type interface{}

    // Do advertise connected routes. The type is bool.
    Connected interface{}

    // Do advertise redistributed routes. The type is bool.
    Redistributed interface{}

    // Do advertise static routes. The type is bool.
    Static interface{}

    // Do advertise summary routes. The type is bool.
    Summary interface{}
}

func (stub *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Stub) GetEntityData() *types.CommonEntityData {
    stub.EntityData.YFilter = stub.YFilter
    stub.EntityData.YangName = "stub"
    stub.EntityData.BundleName = "cisco_ios_xr"
    stub.EntityData.ParentYangName = "af"
    stub.EntityData.SegmentPath = "stub"
    stub.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/" + stub.EntityData.SegmentPath
    stub.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stub.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stub.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stub.EntityData.Children = types.NewOrderedMap()
    stub.EntityData.Leafs = types.NewOrderedMap()
    stub.EntityData.Leafs.Append("type", types.YLeaf{"Type", stub.Type})
    stub.EntityData.Leafs.Append("connected", types.YLeaf{"Connected", stub.Connected})
    stub.EntityData.Leafs.Append("redistributed", types.YLeaf{"Redistributed", stub.Redistributed})
    stub.EntityData.Leafs.Append("static", types.YLeaf{"Static", stub.Static})
    stub.EntityData.Leafs.Append("summary", types.YLeaf{"Summary", stub.Summary})

    stub.EntityData.YListKeys = []string {}

    return &(stub.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes
// List of redistributed protocols
// This type is a presence type.
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_Redistribute.
    Redistribute []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_Redistribute

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXx.
    RedistributeAsXx []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXx

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsYy.
    RedistributeAsYy []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsYy

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxAsYy.
    RedistributeAsXxAsYy []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxAsYy

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeTagName.
    RedistributeTagName []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeTagName

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxTagName.
    RedistributeAsXxTagName []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxTagName

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsYyTagName.
    RedistributeAsYyTagName []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsYyTagName

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxAsYyTagName.
    RedistributeAsXxAsYyTagName []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxAsYyTagName
}

func (redistributes *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes) GetEntityData() *types.CommonEntityData {
    redistributes.EntityData.YFilter = redistributes.YFilter
    redistributes.EntityData.YangName = "redistributes"
    redistributes.EntityData.BundleName = "cisco_ios_xr"
    redistributes.EntityData.ParentYangName = "af"
    redistributes.EntityData.SegmentPath = "redistributes"
    redistributes.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/" + redistributes.EntityData.SegmentPath
    redistributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributes.EntityData.Children = types.NewOrderedMap()
    redistributes.EntityData.Children.Append("redistribute", types.YChild{"Redistribute", nil})
    for i := range redistributes.Redistribute {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.Redistribute[i]), types.YChild{"Redistribute", redistributes.Redistribute[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-as-xx", types.YChild{"RedistributeAsXx", nil})
    for i := range redistributes.RedistributeAsXx {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeAsXx[i]), types.YChild{"RedistributeAsXx", redistributes.RedistributeAsXx[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-as-yy", types.YChild{"RedistributeAsYy", nil})
    for i := range redistributes.RedistributeAsYy {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeAsYy[i]), types.YChild{"RedistributeAsYy", redistributes.RedistributeAsYy[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-as-xx-as-yy", types.YChild{"RedistributeAsXxAsYy", nil})
    for i := range redistributes.RedistributeAsXxAsYy {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeAsXxAsYy[i]), types.YChild{"RedistributeAsXxAsYy", redistributes.RedistributeAsXxAsYy[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-tag-name", types.YChild{"RedistributeTagName", nil})
    for i := range redistributes.RedistributeTagName {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeTagName[i]), types.YChild{"RedistributeTagName", redistributes.RedistributeTagName[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-as-xx-tag-name", types.YChild{"RedistributeAsXxTagName", nil})
    for i := range redistributes.RedistributeAsXxTagName {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeAsXxTagName[i]), types.YChild{"RedistributeAsXxTagName", redistributes.RedistributeAsXxTagName[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-as-yy-tag-name", types.YChild{"RedistributeAsYyTagName", nil})
    for i := range redistributes.RedistributeAsYyTagName {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeAsYyTagName[i]), types.YChild{"RedistributeAsYyTagName", redistributes.RedistributeAsYyTagName[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-as-xx-as-yy-tag-name", types.YChild{"RedistributeAsXxAsYyTagName", nil})
    for i := range redistributes.RedistributeAsXxAsYyTagName {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeAsXxAsYyTagName[i]), types.YChild{"RedistributeAsXxAsYyTagName", redistributes.RedistributeAsXxAsYyTagName[i]})
    }
    redistributes.EntityData.Leafs = types.NewOrderedMap()

    redistributes.EntityData.YListKeys = []string {}

    return &(redistributes.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_Redistribute
// Redistribute another protocol
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_Redistribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistribute *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_Redistribute) GetEntityData() *types.CommonEntityData {
    redistribute.EntityData.YFilter = redistribute.YFilter
    redistribute.EntityData.YangName = "redistribute"
    redistribute.EntityData.BundleName = "cisco_ios_xr"
    redistribute.EntityData.ParentYangName = "redistributes"
    redistribute.EntityData.SegmentPath = "redistribute" + types.AddKeyToken(redistribute.ProtocolName, "protocol-name")
    redistribute.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/redistributes/" + redistribute.EntityData.SegmentPath
    redistribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistribute.EntityData.Children = types.NewOrderedMap()
    redistribute.EntityData.Leafs = types.NewOrderedMap()
    redistribute.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistribute.ProtocolName})
    redistribute.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistribute.PolicyName})
    redistribute.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistribute.PolicySpecified})

    redistribute.EntityData.YListKeys = []string {"ProtocolName"}

    return &(redistribute.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXx
// Redistribute another protocol
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Higher sixteen bits of 4-byte BGP AS number. The
    // type is interface{} with range: 0..65535.
    AsXx interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeAsXx *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXx) GetEntityData() *types.CommonEntityData {
    redistributeAsXx.EntityData.YFilter = redistributeAsXx.YFilter
    redistributeAsXx.EntityData.YangName = "redistribute-as-xx"
    redistributeAsXx.EntityData.BundleName = "cisco_ios_xr"
    redistributeAsXx.EntityData.ParentYangName = "redistributes"
    redistributeAsXx.EntityData.SegmentPath = "redistribute-as-xx" + types.AddKeyToken(redistributeAsXx.AsXx, "as-xx") + types.AddKeyToken(redistributeAsXx.ProtocolName, "protocol-name")
    redistributeAsXx.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/redistributes/" + redistributeAsXx.EntityData.SegmentPath
    redistributeAsXx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeAsXx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeAsXx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeAsXx.EntityData.Children = types.NewOrderedMap()
    redistributeAsXx.EntityData.Leafs = types.NewOrderedMap()
    redistributeAsXx.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", redistributeAsXx.AsXx})
    redistributeAsXx.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeAsXx.ProtocolName})
    redistributeAsXx.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeAsXx.PolicyName})
    redistributeAsXx.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeAsXx.PolicySpecified})

    redistributeAsXx.EntityData.YListKeys = []string {"AsXx", "ProtocolName"}

    return &(redistributeAsXx.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsYy
// Redistribute another protocol
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsYy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. 2-byte or 4-byte BGP AS Number. The type is
    // interface{} with range: 0..4294967295.
    AsYy interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeAsYy *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsYy) GetEntityData() *types.CommonEntityData {
    redistributeAsYy.EntityData.YFilter = redistributeAsYy.YFilter
    redistributeAsYy.EntityData.YangName = "redistribute-as-yy"
    redistributeAsYy.EntityData.BundleName = "cisco_ios_xr"
    redistributeAsYy.EntityData.ParentYangName = "redistributes"
    redistributeAsYy.EntityData.SegmentPath = "redistribute-as-yy" + types.AddKeyToken(redistributeAsYy.AsYy, "as-yy") + types.AddKeyToken(redistributeAsYy.ProtocolName, "protocol-name")
    redistributeAsYy.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/redistributes/" + redistributeAsYy.EntityData.SegmentPath
    redistributeAsYy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeAsYy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeAsYy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeAsYy.EntityData.Children = types.NewOrderedMap()
    redistributeAsYy.EntityData.Leafs = types.NewOrderedMap()
    redistributeAsYy.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", redistributeAsYy.AsYy})
    redistributeAsYy.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeAsYy.ProtocolName})
    redistributeAsYy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeAsYy.PolicyName})
    redistributeAsYy.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeAsYy.PolicySpecified})

    redistributeAsYy.EntityData.YListKeys = []string {"AsYy", "ProtocolName"}

    return &(redistributeAsYy.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxAsYy
// Redistribute another protocol
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxAsYy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Higher sixteen bits of 4-byte BGP AS number. The
    // type is interface{} with range: 0..65535.
    AsXx interface{}

    // This attribute is a key. 2-byte or 4-byte BGP AS Number. The type is
    // interface{} with range: 0..4294967295.
    AsYy interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeAsXxAsYy *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxAsYy) GetEntityData() *types.CommonEntityData {
    redistributeAsXxAsYy.EntityData.YFilter = redistributeAsXxAsYy.YFilter
    redistributeAsXxAsYy.EntityData.YangName = "redistribute-as-xx-as-yy"
    redistributeAsXxAsYy.EntityData.BundleName = "cisco_ios_xr"
    redistributeAsXxAsYy.EntityData.ParentYangName = "redistributes"
    redistributeAsXxAsYy.EntityData.SegmentPath = "redistribute-as-xx-as-yy" + types.AddKeyToken(redistributeAsXxAsYy.AsXx, "as-xx") + types.AddKeyToken(redistributeAsXxAsYy.AsYy, "as-yy") + types.AddKeyToken(redistributeAsXxAsYy.ProtocolName, "protocol-name")
    redistributeAsXxAsYy.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/redistributes/" + redistributeAsXxAsYy.EntityData.SegmentPath
    redistributeAsXxAsYy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeAsXxAsYy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeAsXxAsYy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeAsXxAsYy.EntityData.Children = types.NewOrderedMap()
    redistributeAsXxAsYy.EntityData.Leafs = types.NewOrderedMap()
    redistributeAsXxAsYy.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", redistributeAsXxAsYy.AsXx})
    redistributeAsXxAsYy.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", redistributeAsXxAsYy.AsYy})
    redistributeAsXxAsYy.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeAsXxAsYy.ProtocolName})
    redistributeAsXxAsYy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeAsXxAsYy.PolicyName})
    redistributeAsXxAsYy.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeAsXxAsYy.PolicySpecified})

    redistributeAsXxAsYy.EntityData.YListKeys = []string {"AsXx", "AsYy", "ProtocolName"}

    return &(redistributeAsXxAsYy.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeTagName
// Redistribute another protocol
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeTagName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. OSPF/OSPFv3/ISIS/OnePK Application tag name. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TagName interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeTagName *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeTagName) GetEntityData() *types.CommonEntityData {
    redistributeTagName.EntityData.YFilter = redistributeTagName.YFilter
    redistributeTagName.EntityData.YangName = "redistribute-tag-name"
    redistributeTagName.EntityData.BundleName = "cisco_ios_xr"
    redistributeTagName.EntityData.ParentYangName = "redistributes"
    redistributeTagName.EntityData.SegmentPath = "redistribute-tag-name" + types.AddKeyToken(redistributeTagName.TagName, "tag-name") + types.AddKeyToken(redistributeTagName.ProtocolName, "protocol-name")
    redistributeTagName.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/redistributes/" + redistributeTagName.EntityData.SegmentPath
    redistributeTagName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeTagName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeTagName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeTagName.EntityData.Children = types.NewOrderedMap()
    redistributeTagName.EntityData.Leafs = types.NewOrderedMap()
    redistributeTagName.EntityData.Leafs.Append("tag-name", types.YLeaf{"TagName", redistributeTagName.TagName})
    redistributeTagName.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeTagName.ProtocolName})
    redistributeTagName.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeTagName.PolicyName})
    redistributeTagName.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeTagName.PolicySpecified})

    redistributeTagName.EntityData.YListKeys = []string {"TagName", "ProtocolName"}

    return &(redistributeTagName.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxTagName
// Redistribute another protocol
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxTagName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Higher sixteen bits of 4-byte BGP AS number. The
    // type is interface{} with range: 0..65535.
    AsXx interface{}

    // This attribute is a key. OSPF/OSPFv3/ISIS/OnePK Application tag name. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TagName interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeAsXxTagName *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxTagName) GetEntityData() *types.CommonEntityData {
    redistributeAsXxTagName.EntityData.YFilter = redistributeAsXxTagName.YFilter
    redistributeAsXxTagName.EntityData.YangName = "redistribute-as-xx-tag-name"
    redistributeAsXxTagName.EntityData.BundleName = "cisco_ios_xr"
    redistributeAsXxTagName.EntityData.ParentYangName = "redistributes"
    redistributeAsXxTagName.EntityData.SegmentPath = "redistribute-as-xx-tag-name" + types.AddKeyToken(redistributeAsXxTagName.AsXx, "as-xx") + types.AddKeyToken(redistributeAsXxTagName.TagName, "tag-name") + types.AddKeyToken(redistributeAsXxTagName.ProtocolName, "protocol-name")
    redistributeAsXxTagName.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/redistributes/" + redistributeAsXxTagName.EntityData.SegmentPath
    redistributeAsXxTagName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeAsXxTagName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeAsXxTagName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeAsXxTagName.EntityData.Children = types.NewOrderedMap()
    redistributeAsXxTagName.EntityData.Leafs = types.NewOrderedMap()
    redistributeAsXxTagName.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", redistributeAsXxTagName.AsXx})
    redistributeAsXxTagName.EntityData.Leafs.Append("tag-name", types.YLeaf{"TagName", redistributeAsXxTagName.TagName})
    redistributeAsXxTagName.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeAsXxTagName.ProtocolName})
    redistributeAsXxTagName.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeAsXxTagName.PolicyName})
    redistributeAsXxTagName.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeAsXxTagName.PolicySpecified})

    redistributeAsXxTagName.EntityData.YListKeys = []string {"AsXx", "TagName", "ProtocolName"}

    return &(redistributeAsXxTagName.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsYyTagName
// Redistribute another protocol
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsYyTagName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. 2-byte or 4-byte BGP AS Number. The type is
    // interface{} with range: 0..4294967295.
    AsYy interface{}

    // This attribute is a key. OSPF/OSPFv3/ISIS/OnePK Application tag name. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TagName interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeAsYyTagName *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsYyTagName) GetEntityData() *types.CommonEntityData {
    redistributeAsYyTagName.EntityData.YFilter = redistributeAsYyTagName.YFilter
    redistributeAsYyTagName.EntityData.YangName = "redistribute-as-yy-tag-name"
    redistributeAsYyTagName.EntityData.BundleName = "cisco_ios_xr"
    redistributeAsYyTagName.EntityData.ParentYangName = "redistributes"
    redistributeAsYyTagName.EntityData.SegmentPath = "redistribute-as-yy-tag-name" + types.AddKeyToken(redistributeAsYyTagName.AsYy, "as-yy") + types.AddKeyToken(redistributeAsYyTagName.TagName, "tag-name") + types.AddKeyToken(redistributeAsYyTagName.ProtocolName, "protocol-name")
    redistributeAsYyTagName.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/redistributes/" + redistributeAsYyTagName.EntityData.SegmentPath
    redistributeAsYyTagName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeAsYyTagName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeAsYyTagName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeAsYyTagName.EntityData.Children = types.NewOrderedMap()
    redistributeAsYyTagName.EntityData.Leafs = types.NewOrderedMap()
    redistributeAsYyTagName.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", redistributeAsYyTagName.AsYy})
    redistributeAsYyTagName.EntityData.Leafs.Append("tag-name", types.YLeaf{"TagName", redistributeAsYyTagName.TagName})
    redistributeAsYyTagName.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeAsYyTagName.ProtocolName})
    redistributeAsYyTagName.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeAsYyTagName.PolicyName})
    redistributeAsYyTagName.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeAsYyTagName.PolicySpecified})

    redistributeAsYyTagName.EntityData.YListKeys = []string {"AsYy", "TagName", "ProtocolName"}

    return &(redistributeAsYyTagName.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxAsYyTagName
// Redistribute another protocol
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxAsYyTagName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Higher sixteen bits of 4-byte BGP AS number. The
    // type is interface{} with range: 0..65535.
    AsXx interface{}

    // This attribute is a key. 2-byte or 4-byte BGP AS Number. The type is
    // interface{} with range: 0..4294967295.
    AsYy interface{}

    // This attribute is a key. OSPF/OSPFv3/ISIS/OnePK Application tag name. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TagName interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeAsXxAsYyTagName *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Redistributes_RedistributeAsXxAsYyTagName) GetEntityData() *types.CommonEntityData {
    redistributeAsXxAsYyTagName.EntityData.YFilter = redistributeAsXxAsYyTagName.YFilter
    redistributeAsXxAsYyTagName.EntityData.YangName = "redistribute-as-xx-as-yy-tag-name"
    redistributeAsXxAsYyTagName.EntityData.BundleName = "cisco_ios_xr"
    redistributeAsXxAsYyTagName.EntityData.ParentYangName = "redistributes"
    redistributeAsXxAsYyTagName.EntityData.SegmentPath = "redistribute-as-xx-as-yy-tag-name" + types.AddKeyToken(redistributeAsXxAsYyTagName.AsXx, "as-xx") + types.AddKeyToken(redistributeAsXxAsYyTagName.AsYy, "as-yy") + types.AddKeyToken(redistributeAsXxAsYyTagName.TagName, "tag-name") + types.AddKeyToken(redistributeAsXxAsYyTagName.ProtocolName, "protocol-name")
    redistributeAsXxAsYyTagName.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/redistributes/" + redistributeAsXxAsYyTagName.EntityData.SegmentPath
    redistributeAsXxAsYyTagName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeAsXxAsYyTagName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeAsXxAsYyTagName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeAsXxAsYyTagName.EntityData.Children = types.NewOrderedMap()
    redistributeAsXxAsYyTagName.EntityData.Leafs = types.NewOrderedMap()
    redistributeAsXxAsYyTagName.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", redistributeAsXxAsYyTagName.AsXx})
    redistributeAsXxAsYyTagName.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", redistributeAsXxAsYyTagName.AsYy})
    redistributeAsXxAsYyTagName.EntityData.Leafs.Append("tag-name", types.YLeaf{"TagName", redistributeAsXxAsYyTagName.TagName})
    redistributeAsXxAsYyTagName.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeAsXxAsYyTagName.ProtocolName})
    redistributeAsXxAsYyTagName.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeAsXxAsYyTagName.PolicyName})
    redistributeAsXxAsYyTagName.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeAsXxAsYyTagName.PolicySpecified})

    redistributeAsXxAsYyTagName.EntityData.YListKeys = []string {"AsXx", "AsYy", "TagName", "ProtocolName"}

    return &(redistributeAsXxAsYyTagName.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_FilterPolicies
// Inbound and outbound filter policies
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_FilterPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inbound/outbound policies. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_FilterPolicies_FilterPolicy.
    FilterPolicy []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_FilterPolicies_FilterPolicy
}

func (filterPolicies *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_FilterPolicies) GetEntityData() *types.CommonEntityData {
    filterPolicies.EntityData.YFilter = filterPolicies.YFilter
    filterPolicies.EntityData.YangName = "filter-policies"
    filterPolicies.EntityData.BundleName = "cisco_ios_xr"
    filterPolicies.EntityData.ParentYangName = "af"
    filterPolicies.EntityData.SegmentPath = "filter-policies"
    filterPolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/" + filterPolicies.EntityData.SegmentPath
    filterPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    filterPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    filterPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    filterPolicies.EntityData.Children = types.NewOrderedMap()
    filterPolicies.EntityData.Children.Append("filter-policy", types.YChild{"FilterPolicy", nil})
    for i := range filterPolicies.FilterPolicy {
        filterPolicies.EntityData.Children.Append(types.GetSegmentPath(filterPolicies.FilterPolicy[i]), types.YChild{"FilterPolicy", filterPolicies.FilterPolicy[i]})
    }
    filterPolicies.EntityData.Leafs = types.NewOrderedMap()

    filterPolicies.EntityData.YListKeys = []string {}

    return &(filterPolicies.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_FilterPolicies_FilterPolicy
// Inbound/outbound policies
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_FilterPolicies_FilterPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Direction (in or out). The type is EigrpDir.
    Direction interface{}

    // Policy name. The type is string. This attribute is mandatory.
    PolicyName interface{}
}

func (filterPolicy *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_FilterPolicies_FilterPolicy) GetEntityData() *types.CommonEntityData {
    filterPolicy.EntityData.YFilter = filterPolicy.YFilter
    filterPolicy.EntityData.YangName = "filter-policy"
    filterPolicy.EntityData.BundleName = "cisco_ios_xr"
    filterPolicy.EntityData.ParentYangName = "filter-policies"
    filterPolicy.EntityData.SegmentPath = "filter-policy" + types.AddKeyToken(filterPolicy.Direction, "direction")
    filterPolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/filter-policies/" + filterPolicy.EntityData.SegmentPath
    filterPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    filterPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    filterPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    filterPolicy.EntityData.Children = types.NewOrderedMap()
    filterPolicy.EntityData.Leafs = types.NewOrderedMap()
    filterPolicy.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", filterPolicy.Direction})
    filterPolicy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", filterPolicy.PolicyName})

    filterPolicy.EntityData.YListKeys = []string {"Direction"}

    return &(filterPolicy.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_DefaultMetric
// Set metric of redistributed routes
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_DefaultMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bandwidth in Kbits per second. The type is interface{} with range:
    // 1..4294967295. Units are kbit/s.
    Bandwidth interface{}

    // Delay metric, in 10 microsecond units. The type is interface{} with range:
    // 1..4294967295.
    Delay interface{}

    // Reliability metric where 255 is 100% reliable. The type is interface{} with
    // range: 0..255.
    Reliability interface{}

    // Effective bandwidth metric (Loading) where 255 is 100% loaded. The type is
    // interface{} with range: 1..255.
    Load interface{}

    // Maximum Transmission Unit metric of the path. The type is interface{} with
    // range: 1..65535.
    Mtu interface{}
}

func (defaultMetric *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_DefaultMetric) GetEntityData() *types.CommonEntityData {
    defaultMetric.EntityData.YFilter = defaultMetric.YFilter
    defaultMetric.EntityData.YangName = "default-metric"
    defaultMetric.EntityData.BundleName = "cisco_ios_xr"
    defaultMetric.EntityData.ParentYangName = "af"
    defaultMetric.EntityData.SegmentPath = "default-metric"
    defaultMetric.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/" + defaultMetric.EntityData.SegmentPath
    defaultMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultMetric.EntityData.Children = types.NewOrderedMap()
    defaultMetric.EntityData.Leafs = types.NewOrderedMap()
    defaultMetric.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", defaultMetric.Bandwidth})
    defaultMetric.EntityData.Leafs.Append("delay", types.YLeaf{"Delay", defaultMetric.Delay})
    defaultMetric.EntityData.Leafs.Append("reliability", types.YLeaf{"Reliability", defaultMetric.Reliability})
    defaultMetric.EntityData.Leafs.Append("load", types.YLeaf{"Load", defaultMetric.Load})
    defaultMetric.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", defaultMetric.Mtu})

    defaultMetric.EntityData.YListKeys = []string {}

    return &(defaultMetric.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Metrics
// List of metric change behaviours
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Metrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Modify EIGRP routing metrics and parameters. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Metrics_Metric.
    Metric []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Metrics_Metric
}

func (metrics *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Metrics) GetEntityData() *types.CommonEntityData {
    metrics.EntityData.YFilter = metrics.YFilter
    metrics.EntityData.YangName = "metrics"
    metrics.EntityData.BundleName = "cisco_ios_xr"
    metrics.EntityData.ParentYangName = "af"
    metrics.EntityData.SegmentPath = "metrics"
    metrics.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/" + metrics.EntityData.SegmentPath
    metrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metrics.EntityData.Children = types.NewOrderedMap()
    metrics.EntityData.Children.Append("metric", types.YChild{"Metric", nil})
    for i := range metrics.Metric {
        metrics.EntityData.Children.Append(types.GetSegmentPath(metrics.Metric[i]), types.YChild{"Metric", metrics.Metric[i]})
    }
    metrics.EntityData.Leafs = types.NewOrderedMap()

    metrics.EntityData.YListKeys = []string {}

    return &(metrics.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Metrics_Metric
// Modify EIGRP routing metrics and parameters
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Metrics_Metric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Type of metric change. The type is EigrpMet.
    MetricName interface{}

    // Hop count. The type is interface{} with range: 1..255.
    MaxHops interface{}

    // Type of Service (Only TOS 0 supported). The type is interface{} with range:
    // 0..0.
    Tos interface{}

    // K1. The type is interface{} with range: 0..255.
    K1 interface{}

    // K2. The type is interface{} with range: 0..255.
    K2 interface{}

    // K3. The type is interface{} with range: 0..255.
    K3 interface{}

    // K4. The type is interface{} with range: 0..255.
    K4 interface{}

    // K5. The type is interface{} with range: 0..255.
    K5 interface{}

    // K6. The type is interface{} with range: 0..255.
    K6 interface{}

    // RIB scale. The type is interface{} with range: 1..4294967295.
    RibScale interface{}

    // Metric version. The type is EigrpMetricVersion.
    MetricVersion interface{}
}

func (metric *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Metrics_Metric) GetEntityData() *types.CommonEntityData {
    metric.EntityData.YFilter = metric.YFilter
    metric.EntityData.YangName = "metric"
    metric.EntityData.BundleName = "cisco_ios_xr"
    metric.EntityData.ParentYangName = "metrics"
    metric.EntityData.SegmentPath = "metric" + types.AddKeyToken(metric.MetricName, "metric-name")
    metric.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/metrics/" + metric.EntityData.SegmentPath
    metric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metric.EntityData.Children = types.NewOrderedMap()
    metric.EntityData.Leafs = types.NewOrderedMap()
    metric.EntityData.Leafs.Append("metric-name", types.YLeaf{"MetricName", metric.MetricName})
    metric.EntityData.Leafs.Append("max-hops", types.YLeaf{"MaxHops", metric.MaxHops})
    metric.EntityData.Leafs.Append("tos", types.YLeaf{"Tos", metric.Tos})
    metric.EntityData.Leafs.Append("k1", types.YLeaf{"K1", metric.K1})
    metric.EntityData.Leafs.Append("k2", types.YLeaf{"K2", metric.K2})
    metric.EntityData.Leafs.Append("k3", types.YLeaf{"K3", metric.K3})
    metric.EntityData.Leafs.Append("k4", types.YLeaf{"K4", metric.K4})
    metric.EntityData.Leafs.Append("k5", types.YLeaf{"K5", metric.K5})
    metric.EntityData.Leafs.Append("k6", types.YLeaf{"K6", metric.K6})
    metric.EntityData.Leafs.Append("rib-scale", types.YLeaf{"RibScale", metric.RibScale})
    metric.EntityData.Leafs.Append("metric-version", types.YLeaf{"MetricVersion", metric.MetricVersion})

    metric.EntityData.YListKeys = []string {"MetricName"}

    return &(metric.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Timers
// List of timer configurations
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure EIGRP timers. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Timers_Timer.
    Timer []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Timers_Timer
}

func (timers *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "af"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/" + timers.EntityData.SegmentPath
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Children.Append("timer", types.YChild{"Timer", nil})
    for i := range timers.Timer {
        timers.EntityData.Children.Append(types.GetSegmentPath(timers.Timer[i]), types.YChild{"Timer", timers.Timer[i]})
    }
    timers.EntityData.Leafs = types.NewOrderedMap()

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Timers_Timer
// Configure EIGRP timers
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Timers_Timer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Type of timer. The type is EigrpTimer.
    TimerType interface{}

    // Active Time (in seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    ActiveTime interface{}

    // Hold time (in seconds). The type is interface{} with range: 20..6000. Units
    // are second.
    HoldTime interface{}

    // Signal time (in seconds). The type is interface{} with range: 10..30. Units
    // are second.
    SignalTime interface{}

    // Converge time (in seconds). The type is interface{} with range: 60..5000.
    // Units are second.
    ConvergeTime interface{}
}

func (timer *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Timers_Timer) GetEntityData() *types.CommonEntityData {
    timer.EntityData.YFilter = timer.YFilter
    timer.EntityData.YangName = "timer"
    timer.EntityData.BundleName = "cisco_ios_xr"
    timer.EntityData.ParentYangName = "timers"
    timer.EntityData.SegmentPath = "timer" + types.AddKeyToken(timer.TimerType, "timer-type")
    timer.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/timers/" + timer.EntityData.SegmentPath
    timer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timer.EntityData.Children = types.NewOrderedMap()
    timer.EntityData.Leafs = types.NewOrderedMap()
    timer.EntityData.Leafs.Append("timer-type", types.YLeaf{"TimerType", timer.TimerType})
    timer.EntityData.Leafs.Append("active-time", types.YLeaf{"ActiveTime", timer.ActiveTime})
    timer.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", timer.HoldTime})
    timer.EntityData.Leafs.Append("signal-time", types.YLeaf{"SignalTime", timer.SignalTime})
    timer.EntityData.Leafs.Append("converge-time", types.YLeaf{"ConvergeTime", timer.ConvergeTime})

    timer.EntityData.YListKeys = []string {"TimerType"}

    return &(timer.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_DefaultAccepts
// Candidate default policy table
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_DefaultAccepts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate default behaviour. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_DefaultAccepts_DefaultAccept.
    DefaultAccept []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_DefaultAccepts_DefaultAccept
}

func (defaultAccepts *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_DefaultAccepts) GetEntityData() *types.CommonEntityData {
    defaultAccepts.EntityData.YFilter = defaultAccepts.YFilter
    defaultAccepts.EntityData.YangName = "default-accepts"
    defaultAccepts.EntityData.BundleName = "cisco_ios_xr"
    defaultAccepts.EntityData.ParentYangName = "af"
    defaultAccepts.EntityData.SegmentPath = "default-accepts"
    defaultAccepts.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/" + defaultAccepts.EntityData.SegmentPath
    defaultAccepts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultAccepts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultAccepts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultAccepts.EntityData.Children = types.NewOrderedMap()
    defaultAccepts.EntityData.Children.Append("default-accept", types.YChild{"DefaultAccept", nil})
    for i := range defaultAccepts.DefaultAccept {
        defaultAccepts.EntityData.Children.Append(types.GetSegmentPath(defaultAccepts.DefaultAccept[i]), types.YChild{"DefaultAccept", defaultAccepts.DefaultAccept[i]})
    }
    defaultAccepts.EntityData.Leafs = types.NewOrderedMap()

    defaultAccepts.EntityData.YListKeys = []string {}

    return &(defaultAccepts.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_DefaultAccepts_DefaultAccept
// Candidate default behaviour
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_DefaultAccepts_DefaultAccept struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Direction (in or out). The type is EigrpDir.
    Direction interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (defaultAccept *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_DefaultAccepts_DefaultAccept) GetEntityData() *types.CommonEntityData {
    defaultAccept.EntityData.YFilter = defaultAccept.YFilter
    defaultAccept.EntityData.YangName = "default-accept"
    defaultAccept.EntityData.BundleName = "cisco_ios_xr"
    defaultAccept.EntityData.ParentYangName = "default-accepts"
    defaultAccept.EntityData.SegmentPath = "default-accept" + types.AddKeyToken(defaultAccept.Direction, "direction")
    defaultAccept.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/default-accepts/" + defaultAccept.EntityData.SegmentPath
    defaultAccept.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultAccept.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultAccept.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultAccept.EntityData.Children = types.NewOrderedMap()
    defaultAccept.EntityData.Leafs = types.NewOrderedMap()
    defaultAccept.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", defaultAccept.Direction})
    defaultAccept.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", defaultAccept.PolicyName})
    defaultAccept.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", defaultAccept.PolicySpecified})

    defaultAccept.EntityData.YListKeys = []string {"Direction"}

    return &(defaultAccept.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces
// List of interfaces
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for an Interface.Deletion of this object also causes deletion
    // of all objectsunder 'Interface' associated with this interface instance.
    // The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface.
    Interface []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface
}

func (interfaces *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "af"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/" + interfaces.EntityData.SegmentPath
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

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface
// Configuration for an Interface.Deletion of this
// object also causes deletion of all objectsunder
// 'Interface' associated with this interface
// instance.
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Neighbor hold time (in seconds). The type is interface{} with range:
    // 1..65535. Units are second.
    HoldTime interface{}

    // Bandwidth limit. The type is interface{} with range: 1..999999. Units are
    // percentage.
    BandwidthPercent interface{}

    // Suppress routing updates on an interface. The type is bool.
    PassiveInterface interface{}

    // Interval (in seconds). The type is interface{} with range: 1..65535. Units
    // are second.
    HelloInterval interface{}

    // Disable next-hop-self. The type is interface{}.
    NextHopSelf interface{}

    // Enable Interface configuration. The type is interface{}.
    Enable interface{}

    // Configure split horizon behaviour. The type is interface{}.
    SplitHorizon interface{}

    // Metric.
    InterfaceMetric Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceMetric

    // Remote-Neighbors enabled, default is 65535.
    RemoteNeighbor Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_RemoteNeighbor

    // Configure BFD parameters.
    Bfd Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_Bfd

    // Configure Site-of-origin.
    SiteOfOrigin Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_SiteOfOrigin

    // Authentication configuration.
    Authentication Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_Authentication

    // List of summary addresses under this interface.
    SummaryAddresses Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_SummaryAddresses

    // List of filter policies.
    InterfaceFilterPolicies Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceFilterPolicies

    // List of Neighbors.
    InterfaceStaticNeighbors Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceStaticNeighbors
}

func (self *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("interface-metric", types.YChild{"InterfaceMetric", &self.InterfaceMetric})
    self.EntityData.Children.Append("remote-neighbor", types.YChild{"RemoteNeighbor", &self.RemoteNeighbor})
    self.EntityData.Children.Append("bfd", types.YChild{"Bfd", &self.Bfd})
    self.EntityData.Children.Append("site-of-origin", types.YChild{"SiteOfOrigin", &self.SiteOfOrigin})
    self.EntityData.Children.Append("authentication", types.YChild{"Authentication", &self.Authentication})
    self.EntityData.Children.Append("summary-addresses", types.YChild{"SummaryAddresses", &self.SummaryAddresses})
    self.EntityData.Children.Append("interface-filter-policies", types.YChild{"InterfaceFilterPolicies", &self.InterfaceFilterPolicies})
    self.EntityData.Children.Append("interface-static-neighbors", types.YChild{"InterfaceStaticNeighbors", &self.InterfaceStaticNeighbors})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", self.HoldTime})
    self.EntityData.Leafs.Append("bandwidth-percent", types.YLeaf{"BandwidthPercent", self.BandwidthPercent})
    self.EntityData.Leafs.Append("passive-interface", types.YLeaf{"PassiveInterface", self.PassiveInterface})
    self.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", self.HelloInterval})
    self.EntityData.Leafs.Append("next-hop-self", types.YLeaf{"NextHopSelf", self.NextHopSelf})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})
    self.EntityData.Leafs.Append("split-horizon", types.YLeaf{"SplitHorizon", self.SplitHorizon})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceMetric
// Metric
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bandwidth in Kbits per second. The type is interface{} with range:
    // 1..4294967295. Units are kbit/s.
    Bandwidth interface{}

    // Delay metric, in 10 microsecond units (default) or picosecond units. The
    // type is interface{} with range: 1..4294967295.
    Delay interface{}

    // Delay unit. The type is EigrpDelayUnit.
    DelayUnit interface{}

    // Reliability metric where 255 is 100% reliable. The type is interface{} with
    // range: 0..255.
    Reliability interface{}

    // Effective bandwidth metric (Loading) where 255 is 100% loaded. The type is
    // interface{} with range: 1..255.
    Load interface{}
}

func (interfaceMetric *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceMetric) GetEntityData() *types.CommonEntityData {
    interfaceMetric.EntityData.YFilter = interfaceMetric.YFilter
    interfaceMetric.EntityData.YangName = "interface-metric"
    interfaceMetric.EntityData.BundleName = "cisco_ios_xr"
    interfaceMetric.EntityData.ParentYangName = "interface"
    interfaceMetric.EntityData.SegmentPath = "interface-metric"
    interfaceMetric.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/interfaces/interface/" + interfaceMetric.EntityData.SegmentPath
    interfaceMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMetric.EntityData.Children = types.NewOrderedMap()
    interfaceMetric.EntityData.Leafs = types.NewOrderedMap()
    interfaceMetric.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", interfaceMetric.Bandwidth})
    interfaceMetric.EntityData.Leafs.Append("delay", types.YLeaf{"Delay", interfaceMetric.Delay})
    interfaceMetric.EntityData.Leafs.Append("delay-unit", types.YLeaf{"DelayUnit", interfaceMetric.DelayUnit})
    interfaceMetric.EntityData.Leafs.Append("reliability", types.YLeaf{"Reliability", interfaceMetric.Reliability})
    interfaceMetric.EntityData.Leafs.Append("load", types.YLeaf{"Load", interfaceMetric.Load})

    interfaceMetric.EntityData.YListKeys = []string {}

    return &(interfaceMetric.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_RemoteNeighbor
// Remote-Neighbors enabled, default is 65535
// This type is a presence type.
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_RemoteNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable Remote neighbor unicast-listen. The type is bool. This attribute is
    // mandatory.
    Enable interface{}

    // Policy name. The type is string.
    AllowList interface{}

    // Neighbor count. The type is interface{} with range: 1..65535.
    MaxNeighbors interface{}
}

func (remoteNeighbor *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_RemoteNeighbor) GetEntityData() *types.CommonEntityData {
    remoteNeighbor.EntityData.YFilter = remoteNeighbor.YFilter
    remoteNeighbor.EntityData.YangName = "remote-neighbor"
    remoteNeighbor.EntityData.BundleName = "cisco_ios_xr"
    remoteNeighbor.EntityData.ParentYangName = "interface"
    remoteNeighbor.EntityData.SegmentPath = "remote-neighbor"
    remoteNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/interfaces/interface/" + remoteNeighbor.EntityData.SegmentPath
    remoteNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNeighbor.EntityData.Children = types.NewOrderedMap()
    remoteNeighbor.EntityData.Leafs = types.NewOrderedMap()
    remoteNeighbor.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", remoteNeighbor.Enable})
    remoteNeighbor.EntityData.Leafs.Append("allow-list", types.YLeaf{"AllowList", remoteNeighbor.AllowList})
    remoteNeighbor.EntityData.Leafs.Append("max-neighbors", types.YLeaf{"MaxNeighbors", remoteNeighbor.MaxNeighbors})

    remoteNeighbor.EntityData.YListKeys = []string {}

    return &(remoteNeighbor.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_Bfd
// Configure BFD parameters
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BFD fast detection. The type is bool.
    FastDetect interface{}

    // Detect multiplier. The type is interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 15..3000. Units are millisecond.
    Interval interface{}
}

func (bfd *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "interface"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/interfaces/interface/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("fast-detect", types.YLeaf{"FastDetect", bfd.FastDetect})
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_SiteOfOrigin
// Configure Site-of-origin
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_SiteOfOrigin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SoO type. The type is EigrpSoo.
    Type interface{}

    // Higher sixteen bits of 4-byte BGP AS Number. The type is interface{} with
    // range: 0..65535.
    AsXx interface{}

    // 2-byte or 4-byte BGP AS Number. The type is interface{} with range:
    // 0..4294967295.
    AsYy interface{}

    // AS Number Index. The type is interface{} with range: 0..4294967295.
    Index interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // IPv4 address index. The type is interface{} with range: 0..65535.
    AddressIndex interface{}
}

func (siteOfOrigin *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_SiteOfOrigin) GetEntityData() *types.CommonEntityData {
    siteOfOrigin.EntityData.YFilter = siteOfOrigin.YFilter
    siteOfOrigin.EntityData.YangName = "site-of-origin"
    siteOfOrigin.EntityData.BundleName = "cisco_ios_xr"
    siteOfOrigin.EntityData.ParentYangName = "interface"
    siteOfOrigin.EntityData.SegmentPath = "site-of-origin"
    siteOfOrigin.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/interfaces/interface/" + siteOfOrigin.EntityData.SegmentPath
    siteOfOrigin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    siteOfOrigin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    siteOfOrigin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    siteOfOrigin.EntityData.Children = types.NewOrderedMap()
    siteOfOrigin.EntityData.Leafs = types.NewOrderedMap()
    siteOfOrigin.EntityData.Leafs.Append("type", types.YLeaf{"Type", siteOfOrigin.Type})
    siteOfOrigin.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", siteOfOrigin.AsXx})
    siteOfOrigin.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", siteOfOrigin.AsYy})
    siteOfOrigin.EntityData.Leafs.Append("index", types.YLeaf{"Index", siteOfOrigin.Index})
    siteOfOrigin.EntityData.Leafs.Append("address", types.YLeaf{"Address", siteOfOrigin.Address})
    siteOfOrigin.EntityData.Leafs.Append("address-index", types.YLeaf{"AddressIndex", siteOfOrigin.AddressIndex})

    siteOfOrigin.EntityData.YListKeys = []string {}

    return &(siteOfOrigin.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_Authentication
// Authentication configuration
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authentication keychain configuration. The type is string.
    Keychain interface{}
}

func (authentication *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "interface"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/interfaces/interface/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("keychain", types.YLeaf{"Keychain", authentication.Keychain})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_SummaryAddresses
// List of summary addresses under this interface
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_SummaryAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary address. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_SummaryAddresses_SummaryAddress.
    SummaryAddress []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_SummaryAddresses_SummaryAddress
}

func (summaryAddresses *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_SummaryAddresses) GetEntityData() *types.CommonEntityData {
    summaryAddresses.EntityData.YFilter = summaryAddresses.YFilter
    summaryAddresses.EntityData.YangName = "summary-addresses"
    summaryAddresses.EntityData.BundleName = "cisco_ios_xr"
    summaryAddresses.EntityData.ParentYangName = "interface"
    summaryAddresses.EntityData.SegmentPath = "summary-addresses"
    summaryAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/interfaces/interface/" + summaryAddresses.EntityData.SegmentPath
    summaryAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryAddresses.EntityData.Children = types.NewOrderedMap()
    summaryAddresses.EntityData.Children.Append("summary-address", types.YChild{"SummaryAddress", nil})
    for i := range summaryAddresses.SummaryAddress {
        summaryAddresses.EntityData.Children.Append(types.GetSegmentPath(summaryAddresses.SummaryAddress[i]), types.YChild{"SummaryAddress", summaryAddresses.SummaryAddress[i]})
    }
    summaryAddresses.EntityData.Leafs = types.NewOrderedMap()

    summaryAddresses.EntityData.YListKeys = []string {}

    return &(summaryAddresses.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_SummaryAddresses_SummaryAddress
// Summary address
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_SummaryAddresses_SummaryAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Summary Prefix (address part). The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SummaryAddressAddr interface{}

    // This attribute is a key. Summary Prefix (prefix part). The type is
    // interface{} with range: 0..128.
    SummaryAddressPrefix interface{}

    // Administrative distance. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    Distance interface{}
}

func (summaryAddress *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_SummaryAddresses_SummaryAddress) GetEntityData() *types.CommonEntityData {
    summaryAddress.EntityData.YFilter = summaryAddress.YFilter
    summaryAddress.EntityData.YangName = "summary-address"
    summaryAddress.EntityData.BundleName = "cisco_ios_xr"
    summaryAddress.EntityData.ParentYangName = "summary-addresses"
    summaryAddress.EntityData.SegmentPath = "summary-address" + types.AddKeyToken(summaryAddress.SummaryAddressAddr, "summary-address-addr") + types.AddKeyToken(summaryAddress.SummaryAddressPrefix, "summary-address-prefix")
    summaryAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/interfaces/interface/summary-addresses/" + summaryAddress.EntityData.SegmentPath
    summaryAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryAddress.EntityData.Children = types.NewOrderedMap()
    summaryAddress.EntityData.Leafs = types.NewOrderedMap()
    summaryAddress.EntityData.Leafs.Append("summary-address-addr", types.YLeaf{"SummaryAddressAddr", summaryAddress.SummaryAddressAddr})
    summaryAddress.EntityData.Leafs.Append("summary-address-prefix", types.YLeaf{"SummaryAddressPrefix", summaryAddress.SummaryAddressPrefix})
    summaryAddress.EntityData.Leafs.Append("distance", types.YLeaf{"Distance", summaryAddress.Distance})

    summaryAddress.EntityData.YListKeys = []string {"SummaryAddressAddr", "SummaryAddressPrefix"}

    return &(summaryAddress.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceFilterPolicies
// List of filter policies
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceFilterPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceFilterPolicies_InterfaceFilterPolicy.
    InterfaceFilterPolicy []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceFilterPolicies_InterfaceFilterPolicy
}

func (interfaceFilterPolicies *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceFilterPolicies) GetEntityData() *types.CommonEntityData {
    interfaceFilterPolicies.EntityData.YFilter = interfaceFilterPolicies.YFilter
    interfaceFilterPolicies.EntityData.YangName = "interface-filter-policies"
    interfaceFilterPolicies.EntityData.BundleName = "cisco_ios_xr"
    interfaceFilterPolicies.EntityData.ParentYangName = "interface"
    interfaceFilterPolicies.EntityData.SegmentPath = "interface-filter-policies"
    interfaceFilterPolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/interfaces/interface/" + interfaceFilterPolicies.EntityData.SegmentPath
    interfaceFilterPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFilterPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFilterPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFilterPolicies.EntityData.Children = types.NewOrderedMap()
    interfaceFilterPolicies.EntityData.Children.Append("interface-filter-policy", types.YChild{"InterfaceFilterPolicy", nil})
    for i := range interfaceFilterPolicies.InterfaceFilterPolicy {
        interfaceFilterPolicies.EntityData.Children.Append(types.GetSegmentPath(interfaceFilterPolicies.InterfaceFilterPolicy[i]), types.YChild{"InterfaceFilterPolicy", interfaceFilterPolicies.InterfaceFilterPolicy[i]})
    }
    interfaceFilterPolicies.EntityData.Leafs = types.NewOrderedMap()

    interfaceFilterPolicies.EntityData.YListKeys = []string {}

    return &(interfaceFilterPolicies.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceFilterPolicies_InterfaceFilterPolicy
// none
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceFilterPolicies_InterfaceFilterPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Direction (in or out). The type is EigrpDir.
    Direction interface{}

    // Policy name. The type is string. This attribute is mandatory.
    PolicyName interface{}
}

func (interfaceFilterPolicy *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceFilterPolicies_InterfaceFilterPolicy) GetEntityData() *types.CommonEntityData {
    interfaceFilterPolicy.EntityData.YFilter = interfaceFilterPolicy.YFilter
    interfaceFilterPolicy.EntityData.YangName = "interface-filter-policy"
    interfaceFilterPolicy.EntityData.BundleName = "cisco_ios_xr"
    interfaceFilterPolicy.EntityData.ParentYangName = "interface-filter-policies"
    interfaceFilterPolicy.EntityData.SegmentPath = "interface-filter-policy" + types.AddKeyToken(interfaceFilterPolicy.Direction, "direction")
    interfaceFilterPolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/interfaces/interface/interface-filter-policies/" + interfaceFilterPolicy.EntityData.SegmentPath
    interfaceFilterPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFilterPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFilterPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFilterPolicy.EntityData.Children = types.NewOrderedMap()
    interfaceFilterPolicy.EntityData.Leafs = types.NewOrderedMap()
    interfaceFilterPolicy.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", interfaceFilterPolicy.Direction})
    interfaceFilterPolicy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", interfaceFilterPolicy.PolicyName})

    interfaceFilterPolicy.EntityData.YListKeys = []string {"Direction"}

    return &(interfaceFilterPolicy.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceStaticNeighbors
// List of Neighbors
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceStaticNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Neighbor. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceStaticNeighbors_InterfaceStaticNeighbor.
    InterfaceStaticNeighbor []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceStaticNeighbors_InterfaceStaticNeighbor
}

func (interfaceStaticNeighbors *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceStaticNeighbors) GetEntityData() *types.CommonEntityData {
    interfaceStaticNeighbors.EntityData.YFilter = interfaceStaticNeighbors.YFilter
    interfaceStaticNeighbors.EntityData.YangName = "interface-static-neighbors"
    interfaceStaticNeighbors.EntityData.BundleName = "cisco_ios_xr"
    interfaceStaticNeighbors.EntityData.ParentYangName = "interface"
    interfaceStaticNeighbors.EntityData.SegmentPath = "interface-static-neighbors"
    interfaceStaticNeighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/interfaces/interface/" + interfaceStaticNeighbors.EntityData.SegmentPath
    interfaceStaticNeighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStaticNeighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStaticNeighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStaticNeighbors.EntityData.Children = types.NewOrderedMap()
    interfaceStaticNeighbors.EntityData.Children.Append("interface-static-neighbor", types.YChild{"InterfaceStaticNeighbor", nil})
    for i := range interfaceStaticNeighbors.InterfaceStaticNeighbor {
        interfaceStaticNeighbors.EntityData.Children.Append(types.GetSegmentPath(interfaceStaticNeighbors.InterfaceStaticNeighbor[i]), types.YChild{"InterfaceStaticNeighbor", interfaceStaticNeighbors.InterfaceStaticNeighbor[i]})
    }
    interfaceStaticNeighbors.EntityData.Leafs = types.NewOrderedMap()

    interfaceStaticNeighbors.EntityData.YListKeys = []string {}

    return &(interfaceStaticNeighbors.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceStaticNeighbors_InterfaceStaticNeighbor
// Configure Neighbor
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceStaticNeighbors_InterfaceStaticNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NeighborAddress interface{}

    // Number of hops. The type is interface{} with range: 2..100. This attribute
    // is mandatory.
    MaxHops interface{}
}

func (interfaceStaticNeighbor *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Interfaces_Interface_InterfaceStaticNeighbors_InterfaceStaticNeighbor) GetEntityData() *types.CommonEntityData {
    interfaceStaticNeighbor.EntityData.YFilter = interfaceStaticNeighbor.YFilter
    interfaceStaticNeighbor.EntityData.YangName = "interface-static-neighbor"
    interfaceStaticNeighbor.EntityData.BundleName = "cisco_ios_xr"
    interfaceStaticNeighbor.EntityData.ParentYangName = "interface-static-neighbors"
    interfaceStaticNeighbor.EntityData.SegmentPath = "interface-static-neighbor" + types.AddKeyToken(interfaceStaticNeighbor.NeighborAddress, "neighbor-address")
    interfaceStaticNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/interfaces/interface/interface-static-neighbors/" + interfaceStaticNeighbor.EntityData.SegmentPath
    interfaceStaticNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStaticNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStaticNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStaticNeighbor.EntityData.Children = types.NewOrderedMap()
    interfaceStaticNeighbor.EntityData.Leafs = types.NewOrderedMap()
    interfaceStaticNeighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", interfaceStaticNeighbor.NeighborAddress})
    interfaceStaticNeighbor.EntityData.Leafs.Append("max-hops", types.YLeaf{"MaxHops", interfaceStaticNeighbor.MaxHops})

    interfaceStaticNeighbor.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(interfaceStaticNeighbor.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Distance
// Set distance for EIGRP routes
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Distance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Internal routes' distance. The type is interface{} with range: 1..255.
    InternalDistance interface{}

    // External routes' distance. The type is interface{} with range: 1..255.
    ExternalDistance interface{}
}

func (distance *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Distance) GetEntityData() *types.CommonEntityData {
    distance.EntityData.YFilter = distance.YFilter
    distance.EntityData.YangName = "distance"
    distance.EntityData.BundleName = "cisco_ios_xr"
    distance.EntityData.ParentYangName = "af"
    distance.EntityData.SegmentPath = "distance"
    distance.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/vrfs/vrf/afs/af/" + distance.EntityData.SegmentPath
    distance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distance.EntityData.Children = types.NewOrderedMap()
    distance.EntityData.Leafs = types.NewOrderedMap()
    distance.EntityData.Leafs.Append("internal-distance", types.YLeaf{"InternalDistance", distance.InternalDistance})
    distance.EntityData.Leafs.Append("external-distance", types.YLeaf{"ExternalDistance", distance.ExternalDistance})

    distance.EntityData.YListKeys = []string {}

    return &(distance.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf
// Default VRF configuration.Deletion of this
// object also causes deletion of all
// objectsunder 'Process' associated with this
// process instance.
type Eigrp_Processes_Process_DefaultVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable EIGRP Default VRF configurationDeletion of this object also causes
    // deletion of all objectsunder 'Process' associated with this process
    // instance. The type is interface{}.
    Enable interface{}

    // Address family list in the default context.
    DefaultAfs Eigrp_Processes_Process_DefaultVrf_DefaultAfs
}

func (defaultVrf *Eigrp_Processes_Process_DefaultVrf) GetEntityData() *types.CommonEntityData {
    defaultVrf.EntityData.YFilter = defaultVrf.YFilter
    defaultVrf.EntityData.YangName = "default-vrf"
    defaultVrf.EntityData.BundleName = "cisco_ios_xr"
    defaultVrf.EntityData.ParentYangName = "process"
    defaultVrf.EntityData.SegmentPath = "default-vrf"
    defaultVrf.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/" + defaultVrf.EntityData.SegmentPath
    defaultVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultVrf.EntityData.Children = types.NewOrderedMap()
    defaultVrf.EntityData.Children.Append("default-afs", types.YChild{"DefaultAfs", &defaultVrf.DefaultAfs})
    defaultVrf.EntityData.Leafs = types.NewOrderedMap()
    defaultVrf.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", defaultVrf.Enable})

    defaultVrf.EntityData.YListKeys = []string {}

    return &(defaultVrf.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs
// Address family list in the default context
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration under an AF in the default context.Deletion of this object
    // also causes deletion of all objectsunder 'DefaultAF' associated with this
    // AF instance. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf.
    DefaultAf []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf
}

func (defaultAfs *Eigrp_Processes_Process_DefaultVrf_DefaultAfs) GetEntityData() *types.CommonEntityData {
    defaultAfs.EntityData.YFilter = defaultAfs.YFilter
    defaultAfs.EntityData.YangName = "default-afs"
    defaultAfs.EntityData.BundleName = "cisco_ios_xr"
    defaultAfs.EntityData.ParentYangName = "default-vrf"
    defaultAfs.EntityData.SegmentPath = "default-afs"
    defaultAfs.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/" + defaultAfs.EntityData.SegmentPath
    defaultAfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultAfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultAfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultAfs.EntityData.Children = types.NewOrderedMap()
    defaultAfs.EntityData.Children.Append("default-af", types.YChild{"DefaultAf", nil})
    for i := range defaultAfs.DefaultAf {
        defaultAfs.EntityData.Children.Append(types.GetSegmentPath(defaultAfs.DefaultAf[i]), types.YChild{"DefaultAf", defaultAfs.DefaultAf[i]})
    }
    defaultAfs.EntityData.Leafs = types.NewOrderedMap()

    defaultAfs.EntityData.YListKeys = []string {}

    return &(defaultAfs.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf
// Configuration under an AF in the default
// context.Deletion of this object also causes
// deletion of all objectsunder 'DefaultAF'
// associated with this AF instance.
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address Family. The type is EigrpAf.
    AfName interface{}

    // Enable an Address Family under a default VRF. The type is interface{}.
    Enable interface{}

    // Enable Auto Summarization. The type is interface{}.
    AutoSummary interface{}

    // number of paths. The type is interface{} with range: 1..32.
    MaximumPaths interface{}

    // Set router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Enable/Disable neighbor state change warnings . The type is interface{}.
    LogNeighborWarnings interface{}

    // Set the autonomous system of a VRF. The type is interface{} with range:
    // 1..65535.
    AutonomousSystem interface{}

    // Control load balancing variance. The type is interface{} with range:
    // 1..128.
    Variance interface{}

    // Disable NSF for this address family under this VRF. The type is
    // interface{}.
    NsfDisable interface{}

    // Suppress routing updates on all interfaces. The type is interface{}.
    PassiveInterfaceDefault interface{}

    // Enable/Disable logginf of neighbor state changes. The type is interface{}.
    LogNeighborChanges interface{}

    // EIGRP stub configuration.
    Stub Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Stub

    // List of redistributed protocols.
    Redistributes Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes

    // Inbound and outbound filter policies.
    FilterPolicies Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_FilterPolicies

    // Set metric of redistributed routes.
    DefaultMetric Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_DefaultMetric

    // List of metric change behaviours.
    Metrics Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Metrics

    // List of timer configurations.
    Timers Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Timers

    // Candidate default policy table.
    DefaultAccepts Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_DefaultAccepts

    // List of interfaces.
    Interfaces Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces

    // Set distance for EIGRP routes.
    Distance Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Distance
}

func (defaultAf *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf) GetEntityData() *types.CommonEntityData {
    defaultAf.EntityData.YFilter = defaultAf.YFilter
    defaultAf.EntityData.YangName = "default-af"
    defaultAf.EntityData.BundleName = "cisco_ios_xr"
    defaultAf.EntityData.ParentYangName = "default-afs"
    defaultAf.EntityData.SegmentPath = "default-af" + types.AddKeyToken(defaultAf.AfName, "af-name")
    defaultAf.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/" + defaultAf.EntityData.SegmentPath
    defaultAf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultAf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultAf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultAf.EntityData.Children = types.NewOrderedMap()
    defaultAf.EntityData.Children.Append("stub", types.YChild{"Stub", &defaultAf.Stub})
    defaultAf.EntityData.Children.Append("redistributes", types.YChild{"Redistributes", &defaultAf.Redistributes})
    defaultAf.EntityData.Children.Append("filter-policies", types.YChild{"FilterPolicies", &defaultAf.FilterPolicies})
    defaultAf.EntityData.Children.Append("default-metric", types.YChild{"DefaultMetric", &defaultAf.DefaultMetric})
    defaultAf.EntityData.Children.Append("metrics", types.YChild{"Metrics", &defaultAf.Metrics})
    defaultAf.EntityData.Children.Append("timers", types.YChild{"Timers", &defaultAf.Timers})
    defaultAf.EntityData.Children.Append("default-accepts", types.YChild{"DefaultAccepts", &defaultAf.DefaultAccepts})
    defaultAf.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &defaultAf.Interfaces})
    defaultAf.EntityData.Children.Append("distance", types.YChild{"Distance", &defaultAf.Distance})
    defaultAf.EntityData.Leafs = types.NewOrderedMap()
    defaultAf.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", defaultAf.AfName})
    defaultAf.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", defaultAf.Enable})
    defaultAf.EntityData.Leafs.Append("auto-summary", types.YLeaf{"AutoSummary", defaultAf.AutoSummary})
    defaultAf.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", defaultAf.MaximumPaths})
    defaultAf.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", defaultAf.RouterId})
    defaultAf.EntityData.Leafs.Append("log-neighbor-warnings", types.YLeaf{"LogNeighborWarnings", defaultAf.LogNeighborWarnings})
    defaultAf.EntityData.Leafs.Append("autonomous-system", types.YLeaf{"AutonomousSystem", defaultAf.AutonomousSystem})
    defaultAf.EntityData.Leafs.Append("variance", types.YLeaf{"Variance", defaultAf.Variance})
    defaultAf.EntityData.Leafs.Append("nsf-disable", types.YLeaf{"NsfDisable", defaultAf.NsfDisable})
    defaultAf.EntityData.Leafs.Append("passive-interface-default", types.YLeaf{"PassiveInterfaceDefault", defaultAf.PassiveInterfaceDefault})
    defaultAf.EntityData.Leafs.Append("log-neighbor-changes", types.YLeaf{"LogNeighborChanges", defaultAf.LogNeighborChanges})

    defaultAf.EntityData.YListKeys = []string {"AfName"}

    return &(defaultAf.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Stub
// EIGRP stub configuration
// This type is a presence type.
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Stub struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Stub config type. The type is EigrpStub.
    Type interface{}

    // Do advertise connected routes. The type is bool.
    Connected interface{}

    // Do advertise redistributed routes. The type is bool.
    Redistributed interface{}

    // Do advertise static routes. The type is bool.
    Static interface{}

    // Do advertise summary routes. The type is bool.
    Summary interface{}
}

func (stub *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Stub) GetEntityData() *types.CommonEntityData {
    stub.EntityData.YFilter = stub.YFilter
    stub.EntityData.YangName = "stub"
    stub.EntityData.BundleName = "cisco_ios_xr"
    stub.EntityData.ParentYangName = "default-af"
    stub.EntityData.SegmentPath = "stub"
    stub.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/" + stub.EntityData.SegmentPath
    stub.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stub.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stub.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stub.EntityData.Children = types.NewOrderedMap()
    stub.EntityData.Leafs = types.NewOrderedMap()
    stub.EntityData.Leafs.Append("type", types.YLeaf{"Type", stub.Type})
    stub.EntityData.Leafs.Append("connected", types.YLeaf{"Connected", stub.Connected})
    stub.EntityData.Leafs.Append("redistributed", types.YLeaf{"Redistributed", stub.Redistributed})
    stub.EntityData.Leafs.Append("static", types.YLeaf{"Static", stub.Static})
    stub.EntityData.Leafs.Append("summary", types.YLeaf{"Summary", stub.Summary})

    stub.EntityData.YListKeys = []string {}

    return &(stub.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes
// List of redistributed protocols
// This type is a presence type.
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_Redistribute.
    Redistribute []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_Redistribute

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXx.
    RedistributeAsXx []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXx

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsYy.
    RedistributeAsYy []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsYy

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxAsYy.
    RedistributeAsXxAsYy []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxAsYy

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeTagName.
    RedistributeTagName []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeTagName

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxTagName.
    RedistributeAsXxTagName []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxTagName

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsYyTagName.
    RedistributeAsYyTagName []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsYyTagName

    // Redistribute another protocol. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxAsYyTagName.
    RedistributeAsXxAsYyTagName []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxAsYyTagName
}

func (redistributes *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes) GetEntityData() *types.CommonEntityData {
    redistributes.EntityData.YFilter = redistributes.YFilter
    redistributes.EntityData.YangName = "redistributes"
    redistributes.EntityData.BundleName = "cisco_ios_xr"
    redistributes.EntityData.ParentYangName = "default-af"
    redistributes.EntityData.SegmentPath = "redistributes"
    redistributes.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/" + redistributes.EntityData.SegmentPath
    redistributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributes.EntityData.Children = types.NewOrderedMap()
    redistributes.EntityData.Children.Append("redistribute", types.YChild{"Redistribute", nil})
    for i := range redistributes.Redistribute {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.Redistribute[i]), types.YChild{"Redistribute", redistributes.Redistribute[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-as-xx", types.YChild{"RedistributeAsXx", nil})
    for i := range redistributes.RedistributeAsXx {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeAsXx[i]), types.YChild{"RedistributeAsXx", redistributes.RedistributeAsXx[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-as-yy", types.YChild{"RedistributeAsYy", nil})
    for i := range redistributes.RedistributeAsYy {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeAsYy[i]), types.YChild{"RedistributeAsYy", redistributes.RedistributeAsYy[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-as-xx-as-yy", types.YChild{"RedistributeAsXxAsYy", nil})
    for i := range redistributes.RedistributeAsXxAsYy {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeAsXxAsYy[i]), types.YChild{"RedistributeAsXxAsYy", redistributes.RedistributeAsXxAsYy[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-tag-name", types.YChild{"RedistributeTagName", nil})
    for i := range redistributes.RedistributeTagName {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeTagName[i]), types.YChild{"RedistributeTagName", redistributes.RedistributeTagName[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-as-xx-tag-name", types.YChild{"RedistributeAsXxTagName", nil})
    for i := range redistributes.RedistributeAsXxTagName {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeAsXxTagName[i]), types.YChild{"RedistributeAsXxTagName", redistributes.RedistributeAsXxTagName[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-as-yy-tag-name", types.YChild{"RedistributeAsYyTagName", nil})
    for i := range redistributes.RedistributeAsYyTagName {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeAsYyTagName[i]), types.YChild{"RedistributeAsYyTagName", redistributes.RedistributeAsYyTagName[i]})
    }
    redistributes.EntityData.Children.Append("redistribute-as-xx-as-yy-tag-name", types.YChild{"RedistributeAsXxAsYyTagName", nil})
    for i := range redistributes.RedistributeAsXxAsYyTagName {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.RedistributeAsXxAsYyTagName[i]), types.YChild{"RedistributeAsXxAsYyTagName", redistributes.RedistributeAsXxAsYyTagName[i]})
    }
    redistributes.EntityData.Leafs = types.NewOrderedMap()

    redistributes.EntityData.YListKeys = []string {}

    return &(redistributes.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_Redistribute
// Redistribute another protocol
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_Redistribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistribute *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_Redistribute) GetEntityData() *types.CommonEntityData {
    redistribute.EntityData.YFilter = redistribute.YFilter
    redistribute.EntityData.YangName = "redistribute"
    redistribute.EntityData.BundleName = "cisco_ios_xr"
    redistribute.EntityData.ParentYangName = "redistributes"
    redistribute.EntityData.SegmentPath = "redistribute" + types.AddKeyToken(redistribute.ProtocolName, "protocol-name")
    redistribute.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/redistributes/" + redistribute.EntityData.SegmentPath
    redistribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistribute.EntityData.Children = types.NewOrderedMap()
    redistribute.EntityData.Leafs = types.NewOrderedMap()
    redistribute.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistribute.ProtocolName})
    redistribute.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistribute.PolicyName})
    redistribute.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistribute.PolicySpecified})

    redistribute.EntityData.YListKeys = []string {"ProtocolName"}

    return &(redistribute.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXx
// Redistribute another protocol
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Higher sixteen bits of 4-byte BGP AS number. The
    // type is interface{} with range: 0..65535.
    AsXx interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeAsXx *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXx) GetEntityData() *types.CommonEntityData {
    redistributeAsXx.EntityData.YFilter = redistributeAsXx.YFilter
    redistributeAsXx.EntityData.YangName = "redistribute-as-xx"
    redistributeAsXx.EntityData.BundleName = "cisco_ios_xr"
    redistributeAsXx.EntityData.ParentYangName = "redistributes"
    redistributeAsXx.EntityData.SegmentPath = "redistribute-as-xx" + types.AddKeyToken(redistributeAsXx.AsXx, "as-xx") + types.AddKeyToken(redistributeAsXx.ProtocolName, "protocol-name")
    redistributeAsXx.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/redistributes/" + redistributeAsXx.EntityData.SegmentPath
    redistributeAsXx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeAsXx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeAsXx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeAsXx.EntityData.Children = types.NewOrderedMap()
    redistributeAsXx.EntityData.Leafs = types.NewOrderedMap()
    redistributeAsXx.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", redistributeAsXx.AsXx})
    redistributeAsXx.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeAsXx.ProtocolName})
    redistributeAsXx.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeAsXx.PolicyName})
    redistributeAsXx.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeAsXx.PolicySpecified})

    redistributeAsXx.EntityData.YListKeys = []string {"AsXx", "ProtocolName"}

    return &(redistributeAsXx.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsYy
// Redistribute another protocol
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsYy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. 2-byte or 4-byte BGP AS Number. The type is
    // interface{} with range: 0..4294967295.
    AsYy interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeAsYy *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsYy) GetEntityData() *types.CommonEntityData {
    redistributeAsYy.EntityData.YFilter = redistributeAsYy.YFilter
    redistributeAsYy.EntityData.YangName = "redistribute-as-yy"
    redistributeAsYy.EntityData.BundleName = "cisco_ios_xr"
    redistributeAsYy.EntityData.ParentYangName = "redistributes"
    redistributeAsYy.EntityData.SegmentPath = "redistribute-as-yy" + types.AddKeyToken(redistributeAsYy.AsYy, "as-yy") + types.AddKeyToken(redistributeAsYy.ProtocolName, "protocol-name")
    redistributeAsYy.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/redistributes/" + redistributeAsYy.EntityData.SegmentPath
    redistributeAsYy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeAsYy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeAsYy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeAsYy.EntityData.Children = types.NewOrderedMap()
    redistributeAsYy.EntityData.Leafs = types.NewOrderedMap()
    redistributeAsYy.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", redistributeAsYy.AsYy})
    redistributeAsYy.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeAsYy.ProtocolName})
    redistributeAsYy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeAsYy.PolicyName})
    redistributeAsYy.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeAsYy.PolicySpecified})

    redistributeAsYy.EntityData.YListKeys = []string {"AsYy", "ProtocolName"}

    return &(redistributeAsYy.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxAsYy
// Redistribute another protocol
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxAsYy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Higher sixteen bits of 4-byte BGP AS number. The
    // type is interface{} with range: 0..65535.
    AsXx interface{}

    // This attribute is a key. 2-byte or 4-byte BGP AS Number. The type is
    // interface{} with range: 0..4294967295.
    AsYy interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeAsXxAsYy *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxAsYy) GetEntityData() *types.CommonEntityData {
    redistributeAsXxAsYy.EntityData.YFilter = redistributeAsXxAsYy.YFilter
    redistributeAsXxAsYy.EntityData.YangName = "redistribute-as-xx-as-yy"
    redistributeAsXxAsYy.EntityData.BundleName = "cisco_ios_xr"
    redistributeAsXxAsYy.EntityData.ParentYangName = "redistributes"
    redistributeAsXxAsYy.EntityData.SegmentPath = "redistribute-as-xx-as-yy" + types.AddKeyToken(redistributeAsXxAsYy.AsXx, "as-xx") + types.AddKeyToken(redistributeAsXxAsYy.AsYy, "as-yy") + types.AddKeyToken(redistributeAsXxAsYy.ProtocolName, "protocol-name")
    redistributeAsXxAsYy.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/redistributes/" + redistributeAsXxAsYy.EntityData.SegmentPath
    redistributeAsXxAsYy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeAsXxAsYy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeAsXxAsYy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeAsXxAsYy.EntityData.Children = types.NewOrderedMap()
    redistributeAsXxAsYy.EntityData.Leafs = types.NewOrderedMap()
    redistributeAsXxAsYy.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", redistributeAsXxAsYy.AsXx})
    redistributeAsXxAsYy.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", redistributeAsXxAsYy.AsYy})
    redistributeAsXxAsYy.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeAsXxAsYy.ProtocolName})
    redistributeAsXxAsYy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeAsXxAsYy.PolicyName})
    redistributeAsXxAsYy.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeAsXxAsYy.PolicySpecified})

    redistributeAsXxAsYy.EntityData.YListKeys = []string {"AsXx", "AsYy", "ProtocolName"}

    return &(redistributeAsXxAsYy.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeTagName
// Redistribute another protocol
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeTagName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. OSPF/OSPFv3/ISIS/OnePK Application tag name. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TagName interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeTagName *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeTagName) GetEntityData() *types.CommonEntityData {
    redistributeTagName.EntityData.YFilter = redistributeTagName.YFilter
    redistributeTagName.EntityData.YangName = "redistribute-tag-name"
    redistributeTagName.EntityData.BundleName = "cisco_ios_xr"
    redistributeTagName.EntityData.ParentYangName = "redistributes"
    redistributeTagName.EntityData.SegmentPath = "redistribute-tag-name" + types.AddKeyToken(redistributeTagName.TagName, "tag-name") + types.AddKeyToken(redistributeTagName.ProtocolName, "protocol-name")
    redistributeTagName.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/redistributes/" + redistributeTagName.EntityData.SegmentPath
    redistributeTagName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeTagName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeTagName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeTagName.EntityData.Children = types.NewOrderedMap()
    redistributeTagName.EntityData.Leafs = types.NewOrderedMap()
    redistributeTagName.EntityData.Leafs.Append("tag-name", types.YLeaf{"TagName", redistributeTagName.TagName})
    redistributeTagName.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeTagName.ProtocolName})
    redistributeTagName.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeTagName.PolicyName})
    redistributeTagName.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeTagName.PolicySpecified})

    redistributeTagName.EntityData.YListKeys = []string {"TagName", "ProtocolName"}

    return &(redistributeTagName.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxTagName
// Redistribute another protocol
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxTagName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Higher sixteen bits of 4-byte BGP AS number. The
    // type is interface{} with range: 0..65535.
    AsXx interface{}

    // This attribute is a key. OSPF/OSPFv3/ISIS/OnePK Application tag name. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TagName interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeAsXxTagName *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxTagName) GetEntityData() *types.CommonEntityData {
    redistributeAsXxTagName.EntityData.YFilter = redistributeAsXxTagName.YFilter
    redistributeAsXxTagName.EntityData.YangName = "redistribute-as-xx-tag-name"
    redistributeAsXxTagName.EntityData.BundleName = "cisco_ios_xr"
    redistributeAsXxTagName.EntityData.ParentYangName = "redistributes"
    redistributeAsXxTagName.EntityData.SegmentPath = "redistribute-as-xx-tag-name" + types.AddKeyToken(redistributeAsXxTagName.AsXx, "as-xx") + types.AddKeyToken(redistributeAsXxTagName.TagName, "tag-name") + types.AddKeyToken(redistributeAsXxTagName.ProtocolName, "protocol-name")
    redistributeAsXxTagName.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/redistributes/" + redistributeAsXxTagName.EntityData.SegmentPath
    redistributeAsXxTagName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeAsXxTagName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeAsXxTagName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeAsXxTagName.EntityData.Children = types.NewOrderedMap()
    redistributeAsXxTagName.EntityData.Leafs = types.NewOrderedMap()
    redistributeAsXxTagName.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", redistributeAsXxTagName.AsXx})
    redistributeAsXxTagName.EntityData.Leafs.Append("tag-name", types.YLeaf{"TagName", redistributeAsXxTagName.TagName})
    redistributeAsXxTagName.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeAsXxTagName.ProtocolName})
    redistributeAsXxTagName.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeAsXxTagName.PolicyName})
    redistributeAsXxTagName.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeAsXxTagName.PolicySpecified})

    redistributeAsXxTagName.EntityData.YListKeys = []string {"AsXx", "TagName", "ProtocolName"}

    return &(redistributeAsXxTagName.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsYyTagName
// Redistribute another protocol
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsYyTagName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. 2-byte or 4-byte BGP AS Number. The type is
    // interface{} with range: 0..4294967295.
    AsYy interface{}

    // This attribute is a key. OSPF/OSPFv3/ISIS/OnePK Application tag name. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TagName interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeAsYyTagName *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsYyTagName) GetEntityData() *types.CommonEntityData {
    redistributeAsYyTagName.EntityData.YFilter = redistributeAsYyTagName.YFilter
    redistributeAsYyTagName.EntityData.YangName = "redistribute-as-yy-tag-name"
    redistributeAsYyTagName.EntityData.BundleName = "cisco_ios_xr"
    redistributeAsYyTagName.EntityData.ParentYangName = "redistributes"
    redistributeAsYyTagName.EntityData.SegmentPath = "redistribute-as-yy-tag-name" + types.AddKeyToken(redistributeAsYyTagName.AsYy, "as-yy") + types.AddKeyToken(redistributeAsYyTagName.TagName, "tag-name") + types.AddKeyToken(redistributeAsYyTagName.ProtocolName, "protocol-name")
    redistributeAsYyTagName.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/redistributes/" + redistributeAsYyTagName.EntityData.SegmentPath
    redistributeAsYyTagName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeAsYyTagName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeAsYyTagName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeAsYyTagName.EntityData.Children = types.NewOrderedMap()
    redistributeAsYyTagName.EntityData.Leafs = types.NewOrderedMap()
    redistributeAsYyTagName.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", redistributeAsYyTagName.AsYy})
    redistributeAsYyTagName.EntityData.Leafs.Append("tag-name", types.YLeaf{"TagName", redistributeAsYyTagName.TagName})
    redistributeAsYyTagName.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeAsYyTagName.ProtocolName})
    redistributeAsYyTagName.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeAsYyTagName.PolicyName})
    redistributeAsYyTagName.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeAsYyTagName.PolicySpecified})

    redistributeAsYyTagName.EntityData.YListKeys = []string {"AsYy", "TagName", "ProtocolName"}

    return &(redistributeAsYyTagName.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxAsYyTagName
// Redistribute another protocol
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxAsYyTagName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Higher sixteen bits of 4-byte BGP AS number. The
    // type is interface{} with range: 0..65535.
    AsXx interface{}

    // This attribute is a key. 2-byte or 4-byte BGP AS Number. The type is
    // interface{} with range: 0..4294967295.
    AsYy interface{}

    // This attribute is a key. OSPF/OSPFv3/ISIS/OnePK Application tag name. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TagName interface{}

    // This attribute is a key. Redistributed protocol. The type is
    // EigrpRedistProto.
    ProtocolName interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (redistributeAsXxAsYyTagName *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Redistributes_RedistributeAsXxAsYyTagName) GetEntityData() *types.CommonEntityData {
    redistributeAsXxAsYyTagName.EntityData.YFilter = redistributeAsXxAsYyTagName.YFilter
    redistributeAsXxAsYyTagName.EntityData.YangName = "redistribute-as-xx-as-yy-tag-name"
    redistributeAsXxAsYyTagName.EntityData.BundleName = "cisco_ios_xr"
    redistributeAsXxAsYyTagName.EntityData.ParentYangName = "redistributes"
    redistributeAsXxAsYyTagName.EntityData.SegmentPath = "redistribute-as-xx-as-yy-tag-name" + types.AddKeyToken(redistributeAsXxAsYyTagName.AsXx, "as-xx") + types.AddKeyToken(redistributeAsXxAsYyTagName.AsYy, "as-yy") + types.AddKeyToken(redistributeAsXxAsYyTagName.TagName, "tag-name") + types.AddKeyToken(redistributeAsXxAsYyTagName.ProtocolName, "protocol-name")
    redistributeAsXxAsYyTagName.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/redistributes/" + redistributeAsXxAsYyTagName.EntityData.SegmentPath
    redistributeAsXxAsYyTagName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributeAsXxAsYyTagName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributeAsXxAsYyTagName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributeAsXxAsYyTagName.EntityData.Children = types.NewOrderedMap()
    redistributeAsXxAsYyTagName.EntityData.Leafs = types.NewOrderedMap()
    redistributeAsXxAsYyTagName.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", redistributeAsXxAsYyTagName.AsXx})
    redistributeAsXxAsYyTagName.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", redistributeAsXxAsYyTagName.AsYy})
    redistributeAsXxAsYyTagName.EntityData.Leafs.Append("tag-name", types.YLeaf{"TagName", redistributeAsXxAsYyTagName.TagName})
    redistributeAsXxAsYyTagName.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistributeAsXxAsYyTagName.ProtocolName})
    redistributeAsXxAsYyTagName.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", redistributeAsXxAsYyTagName.PolicyName})
    redistributeAsXxAsYyTagName.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", redistributeAsXxAsYyTagName.PolicySpecified})

    redistributeAsXxAsYyTagName.EntityData.YListKeys = []string {"AsXx", "AsYy", "TagName", "ProtocolName"}

    return &(redistributeAsXxAsYyTagName.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_FilterPolicies
// Inbound and outbound filter policies
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_FilterPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inbound/outbound policies. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_FilterPolicies_FilterPolicy.
    FilterPolicy []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_FilterPolicies_FilterPolicy
}

func (filterPolicies *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_FilterPolicies) GetEntityData() *types.CommonEntityData {
    filterPolicies.EntityData.YFilter = filterPolicies.YFilter
    filterPolicies.EntityData.YangName = "filter-policies"
    filterPolicies.EntityData.BundleName = "cisco_ios_xr"
    filterPolicies.EntityData.ParentYangName = "default-af"
    filterPolicies.EntityData.SegmentPath = "filter-policies"
    filterPolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/" + filterPolicies.EntityData.SegmentPath
    filterPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    filterPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    filterPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    filterPolicies.EntityData.Children = types.NewOrderedMap()
    filterPolicies.EntityData.Children.Append("filter-policy", types.YChild{"FilterPolicy", nil})
    for i := range filterPolicies.FilterPolicy {
        filterPolicies.EntityData.Children.Append(types.GetSegmentPath(filterPolicies.FilterPolicy[i]), types.YChild{"FilterPolicy", filterPolicies.FilterPolicy[i]})
    }
    filterPolicies.EntityData.Leafs = types.NewOrderedMap()

    filterPolicies.EntityData.YListKeys = []string {}

    return &(filterPolicies.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_FilterPolicies_FilterPolicy
// Inbound/outbound policies
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_FilterPolicies_FilterPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Direction (in or out). The type is EigrpDir.
    Direction interface{}

    // Policy name. The type is string. This attribute is mandatory.
    PolicyName interface{}
}

func (filterPolicy *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_FilterPolicies_FilterPolicy) GetEntityData() *types.CommonEntityData {
    filterPolicy.EntityData.YFilter = filterPolicy.YFilter
    filterPolicy.EntityData.YangName = "filter-policy"
    filterPolicy.EntityData.BundleName = "cisco_ios_xr"
    filterPolicy.EntityData.ParentYangName = "filter-policies"
    filterPolicy.EntityData.SegmentPath = "filter-policy" + types.AddKeyToken(filterPolicy.Direction, "direction")
    filterPolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/filter-policies/" + filterPolicy.EntityData.SegmentPath
    filterPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    filterPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    filterPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    filterPolicy.EntityData.Children = types.NewOrderedMap()
    filterPolicy.EntityData.Leafs = types.NewOrderedMap()
    filterPolicy.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", filterPolicy.Direction})
    filterPolicy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", filterPolicy.PolicyName})

    filterPolicy.EntityData.YListKeys = []string {"Direction"}

    return &(filterPolicy.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_DefaultMetric
// Set metric of redistributed routes
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_DefaultMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bandwidth in Kbits per second. The type is interface{} with range:
    // 1..4294967295. Units are kbit/s.
    Bandwidth interface{}

    // Delay metric, in 10 microsecond units. The type is interface{} with range:
    // 1..4294967295.
    Delay interface{}

    // Reliability metric where 255 is 100% reliable. The type is interface{} with
    // range: 0..255.
    Reliability interface{}

    // Effective bandwidth metric (Loading) where 255 is 100% loaded. The type is
    // interface{} with range: 1..255.
    Load interface{}

    // Maximum Transmission Unit metric of the path. The type is interface{} with
    // range: 1..65535.
    Mtu interface{}
}

func (defaultMetric *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_DefaultMetric) GetEntityData() *types.CommonEntityData {
    defaultMetric.EntityData.YFilter = defaultMetric.YFilter
    defaultMetric.EntityData.YangName = "default-metric"
    defaultMetric.EntityData.BundleName = "cisco_ios_xr"
    defaultMetric.EntityData.ParentYangName = "default-af"
    defaultMetric.EntityData.SegmentPath = "default-metric"
    defaultMetric.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/" + defaultMetric.EntityData.SegmentPath
    defaultMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultMetric.EntityData.Children = types.NewOrderedMap()
    defaultMetric.EntityData.Leafs = types.NewOrderedMap()
    defaultMetric.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", defaultMetric.Bandwidth})
    defaultMetric.EntityData.Leafs.Append("delay", types.YLeaf{"Delay", defaultMetric.Delay})
    defaultMetric.EntityData.Leafs.Append("reliability", types.YLeaf{"Reliability", defaultMetric.Reliability})
    defaultMetric.EntityData.Leafs.Append("load", types.YLeaf{"Load", defaultMetric.Load})
    defaultMetric.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", defaultMetric.Mtu})

    defaultMetric.EntityData.YListKeys = []string {}

    return &(defaultMetric.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Metrics
// List of metric change behaviours
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Metrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Modify EIGRP routing metrics and parameters. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Metrics_Metric.
    Metric []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Metrics_Metric
}

func (metrics *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Metrics) GetEntityData() *types.CommonEntityData {
    metrics.EntityData.YFilter = metrics.YFilter
    metrics.EntityData.YangName = "metrics"
    metrics.EntityData.BundleName = "cisco_ios_xr"
    metrics.EntityData.ParentYangName = "default-af"
    metrics.EntityData.SegmentPath = "metrics"
    metrics.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/" + metrics.EntityData.SegmentPath
    metrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metrics.EntityData.Children = types.NewOrderedMap()
    metrics.EntityData.Children.Append("metric", types.YChild{"Metric", nil})
    for i := range metrics.Metric {
        metrics.EntityData.Children.Append(types.GetSegmentPath(metrics.Metric[i]), types.YChild{"Metric", metrics.Metric[i]})
    }
    metrics.EntityData.Leafs = types.NewOrderedMap()

    metrics.EntityData.YListKeys = []string {}

    return &(metrics.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Metrics_Metric
// Modify EIGRP routing metrics and parameters
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Metrics_Metric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Type of metric change. The type is EigrpMet.
    MetricName interface{}

    // Hop count. The type is interface{} with range: 1..255.
    MaxHops interface{}

    // Type of Service (Only TOS 0 supported). The type is interface{} with range:
    // 0..0.
    Tos interface{}

    // K1. The type is interface{} with range: 0..255.
    K1 interface{}

    // K2. The type is interface{} with range: 0..255.
    K2 interface{}

    // K3. The type is interface{} with range: 0..255.
    K3 interface{}

    // K4. The type is interface{} with range: 0..255.
    K4 interface{}

    // K5. The type is interface{} with range: 0..255.
    K5 interface{}

    // K6. The type is interface{} with range: 0..255.
    K6 interface{}

    // RIB scale. The type is interface{} with range: 1..4294967295.
    RibScale interface{}

    // Metric version. The type is EigrpMetricVersion.
    MetricVersion interface{}
}

func (metric *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Metrics_Metric) GetEntityData() *types.CommonEntityData {
    metric.EntityData.YFilter = metric.YFilter
    metric.EntityData.YangName = "metric"
    metric.EntityData.BundleName = "cisco_ios_xr"
    metric.EntityData.ParentYangName = "metrics"
    metric.EntityData.SegmentPath = "metric" + types.AddKeyToken(metric.MetricName, "metric-name")
    metric.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/metrics/" + metric.EntityData.SegmentPath
    metric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metric.EntityData.Children = types.NewOrderedMap()
    metric.EntityData.Leafs = types.NewOrderedMap()
    metric.EntityData.Leafs.Append("metric-name", types.YLeaf{"MetricName", metric.MetricName})
    metric.EntityData.Leafs.Append("max-hops", types.YLeaf{"MaxHops", metric.MaxHops})
    metric.EntityData.Leafs.Append("tos", types.YLeaf{"Tos", metric.Tos})
    metric.EntityData.Leafs.Append("k1", types.YLeaf{"K1", metric.K1})
    metric.EntityData.Leafs.Append("k2", types.YLeaf{"K2", metric.K2})
    metric.EntityData.Leafs.Append("k3", types.YLeaf{"K3", metric.K3})
    metric.EntityData.Leafs.Append("k4", types.YLeaf{"K4", metric.K4})
    metric.EntityData.Leafs.Append("k5", types.YLeaf{"K5", metric.K5})
    metric.EntityData.Leafs.Append("k6", types.YLeaf{"K6", metric.K6})
    metric.EntityData.Leafs.Append("rib-scale", types.YLeaf{"RibScale", metric.RibScale})
    metric.EntityData.Leafs.Append("metric-version", types.YLeaf{"MetricVersion", metric.MetricVersion})

    metric.EntityData.YListKeys = []string {"MetricName"}

    return &(metric.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Timers
// List of timer configurations
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure EIGRP timers. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Timers_Timer.
    Timer []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Timers_Timer
}

func (timers *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "default-af"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/" + timers.EntityData.SegmentPath
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Children.Append("timer", types.YChild{"Timer", nil})
    for i := range timers.Timer {
        timers.EntityData.Children.Append(types.GetSegmentPath(timers.Timer[i]), types.YChild{"Timer", timers.Timer[i]})
    }
    timers.EntityData.Leafs = types.NewOrderedMap()

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Timers_Timer
// Configure EIGRP timers
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Timers_Timer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Type of timer. The type is EigrpTimer.
    TimerType interface{}

    // Active Time (in seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    ActiveTime interface{}

    // Hold time (in seconds). The type is interface{} with range: 20..6000. Units
    // are second.
    HoldTime interface{}

    // Signal time (in seconds). The type is interface{} with range: 10..30. Units
    // are second.
    SignalTime interface{}

    // Converge time (in seconds). The type is interface{} with range: 60..5000.
    // Units are second.
    ConvergeTime interface{}
}

func (timer *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Timers_Timer) GetEntityData() *types.CommonEntityData {
    timer.EntityData.YFilter = timer.YFilter
    timer.EntityData.YangName = "timer"
    timer.EntityData.BundleName = "cisco_ios_xr"
    timer.EntityData.ParentYangName = "timers"
    timer.EntityData.SegmentPath = "timer" + types.AddKeyToken(timer.TimerType, "timer-type")
    timer.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/timers/" + timer.EntityData.SegmentPath
    timer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timer.EntityData.Children = types.NewOrderedMap()
    timer.EntityData.Leafs = types.NewOrderedMap()
    timer.EntityData.Leafs.Append("timer-type", types.YLeaf{"TimerType", timer.TimerType})
    timer.EntityData.Leafs.Append("active-time", types.YLeaf{"ActiveTime", timer.ActiveTime})
    timer.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", timer.HoldTime})
    timer.EntityData.Leafs.Append("signal-time", types.YLeaf{"SignalTime", timer.SignalTime})
    timer.EntityData.Leafs.Append("converge-time", types.YLeaf{"ConvergeTime", timer.ConvergeTime})

    timer.EntityData.YListKeys = []string {"TimerType"}

    return &(timer.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_DefaultAccepts
// Candidate default policy table
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_DefaultAccepts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate default behaviour. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_DefaultAccepts_DefaultAccept.
    DefaultAccept []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_DefaultAccepts_DefaultAccept
}

func (defaultAccepts *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_DefaultAccepts) GetEntityData() *types.CommonEntityData {
    defaultAccepts.EntityData.YFilter = defaultAccepts.YFilter
    defaultAccepts.EntityData.YangName = "default-accepts"
    defaultAccepts.EntityData.BundleName = "cisco_ios_xr"
    defaultAccepts.EntityData.ParentYangName = "default-af"
    defaultAccepts.EntityData.SegmentPath = "default-accepts"
    defaultAccepts.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/" + defaultAccepts.EntityData.SegmentPath
    defaultAccepts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultAccepts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultAccepts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultAccepts.EntityData.Children = types.NewOrderedMap()
    defaultAccepts.EntityData.Children.Append("default-accept", types.YChild{"DefaultAccept", nil})
    for i := range defaultAccepts.DefaultAccept {
        defaultAccepts.EntityData.Children.Append(types.GetSegmentPath(defaultAccepts.DefaultAccept[i]), types.YChild{"DefaultAccept", defaultAccepts.DefaultAccept[i]})
    }
    defaultAccepts.EntityData.Leafs = types.NewOrderedMap()

    defaultAccepts.EntityData.YListKeys = []string {}

    return &(defaultAccepts.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_DefaultAccepts_DefaultAccept
// Candidate default behaviour
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_DefaultAccepts_DefaultAccept struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Direction (in or out). The type is EigrpDir.
    Direction interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // TRUE if Policy has been specified. The type is bool. This attribute is
    // mandatory.
    PolicySpecified interface{}
}

func (defaultAccept *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_DefaultAccepts_DefaultAccept) GetEntityData() *types.CommonEntityData {
    defaultAccept.EntityData.YFilter = defaultAccept.YFilter
    defaultAccept.EntityData.YangName = "default-accept"
    defaultAccept.EntityData.BundleName = "cisco_ios_xr"
    defaultAccept.EntityData.ParentYangName = "default-accepts"
    defaultAccept.EntityData.SegmentPath = "default-accept" + types.AddKeyToken(defaultAccept.Direction, "direction")
    defaultAccept.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/default-accepts/" + defaultAccept.EntityData.SegmentPath
    defaultAccept.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultAccept.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultAccept.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultAccept.EntityData.Children = types.NewOrderedMap()
    defaultAccept.EntityData.Leafs = types.NewOrderedMap()
    defaultAccept.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", defaultAccept.Direction})
    defaultAccept.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", defaultAccept.PolicyName})
    defaultAccept.EntityData.Leafs.Append("policy-specified", types.YLeaf{"PolicySpecified", defaultAccept.PolicySpecified})

    defaultAccept.EntityData.YListKeys = []string {"Direction"}

    return &(defaultAccept.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces
// List of interfaces
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for an Interface.Deletion of this object also causes deletion
    // of all objectsunder 'Interface' associated with this interface instance.
    // The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface.
    Interface []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface
}

func (interfaces *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "default-af"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/" + interfaces.EntityData.SegmentPath
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

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface
// Configuration for an Interface.Deletion of this
// object also causes deletion of all objectsunder
// 'Interface' associated with this interface
// instance.
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Neighbor hold time (in seconds). The type is interface{} with range:
    // 1..65535. Units are second.
    HoldTime interface{}

    // Bandwidth limit. The type is interface{} with range: 1..999999. Units are
    // percentage.
    BandwidthPercent interface{}

    // Suppress routing updates on an interface. The type is bool.
    PassiveInterface interface{}

    // Interval (in seconds). The type is interface{} with range: 1..65535. Units
    // are second.
    HelloInterval interface{}

    // Disable next-hop-self. The type is interface{}.
    NextHopSelf interface{}

    // Enable Interface configuration. The type is interface{}.
    Enable interface{}

    // Configure split horizon behaviour. The type is interface{}.
    SplitHorizon interface{}

    // Metric.
    InterfaceMetric Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceMetric

    // Remote-Neighbors enabled, default is 65535.
    RemoteNeighbor Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_RemoteNeighbor

    // Configure BFD parameters.
    Bfd Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_Bfd

    // Configure Site-of-origin.
    SiteOfOrigin Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_SiteOfOrigin

    // Authentication configuration.
    Authentication Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_Authentication

    // List of summary addresses under this interface.
    SummaryAddresses Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_SummaryAddresses

    // List of filter policies.
    InterfaceFilterPolicies Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceFilterPolicies

    // List of Neighbors.
    InterfaceStaticNeighbors Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceStaticNeighbors
}

func (self *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("interface-metric", types.YChild{"InterfaceMetric", &self.InterfaceMetric})
    self.EntityData.Children.Append("remote-neighbor", types.YChild{"RemoteNeighbor", &self.RemoteNeighbor})
    self.EntityData.Children.Append("bfd", types.YChild{"Bfd", &self.Bfd})
    self.EntityData.Children.Append("site-of-origin", types.YChild{"SiteOfOrigin", &self.SiteOfOrigin})
    self.EntityData.Children.Append("authentication", types.YChild{"Authentication", &self.Authentication})
    self.EntityData.Children.Append("summary-addresses", types.YChild{"SummaryAddresses", &self.SummaryAddresses})
    self.EntityData.Children.Append("interface-filter-policies", types.YChild{"InterfaceFilterPolicies", &self.InterfaceFilterPolicies})
    self.EntityData.Children.Append("interface-static-neighbors", types.YChild{"InterfaceStaticNeighbors", &self.InterfaceStaticNeighbors})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", self.HoldTime})
    self.EntityData.Leafs.Append("bandwidth-percent", types.YLeaf{"BandwidthPercent", self.BandwidthPercent})
    self.EntityData.Leafs.Append("passive-interface", types.YLeaf{"PassiveInterface", self.PassiveInterface})
    self.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", self.HelloInterval})
    self.EntityData.Leafs.Append("next-hop-self", types.YLeaf{"NextHopSelf", self.NextHopSelf})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})
    self.EntityData.Leafs.Append("split-horizon", types.YLeaf{"SplitHorizon", self.SplitHorizon})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceMetric
// Metric
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bandwidth in Kbits per second. The type is interface{} with range:
    // 1..4294967295. Units are kbit/s.
    Bandwidth interface{}

    // Delay metric, in 10 microsecond units (default) or picosecond units. The
    // type is interface{} with range: 1..4294967295.
    Delay interface{}

    // Delay unit. The type is EigrpDelayUnit.
    DelayUnit interface{}

    // Reliability metric where 255 is 100% reliable. The type is interface{} with
    // range: 0..255.
    Reliability interface{}

    // Effective bandwidth metric (Loading) where 255 is 100% loaded. The type is
    // interface{} with range: 1..255.
    Load interface{}
}

func (interfaceMetric *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceMetric) GetEntityData() *types.CommonEntityData {
    interfaceMetric.EntityData.YFilter = interfaceMetric.YFilter
    interfaceMetric.EntityData.YangName = "interface-metric"
    interfaceMetric.EntityData.BundleName = "cisco_ios_xr"
    interfaceMetric.EntityData.ParentYangName = "interface"
    interfaceMetric.EntityData.SegmentPath = "interface-metric"
    interfaceMetric.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/interfaces/interface/" + interfaceMetric.EntityData.SegmentPath
    interfaceMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMetric.EntityData.Children = types.NewOrderedMap()
    interfaceMetric.EntityData.Leafs = types.NewOrderedMap()
    interfaceMetric.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", interfaceMetric.Bandwidth})
    interfaceMetric.EntityData.Leafs.Append("delay", types.YLeaf{"Delay", interfaceMetric.Delay})
    interfaceMetric.EntityData.Leafs.Append("delay-unit", types.YLeaf{"DelayUnit", interfaceMetric.DelayUnit})
    interfaceMetric.EntityData.Leafs.Append("reliability", types.YLeaf{"Reliability", interfaceMetric.Reliability})
    interfaceMetric.EntityData.Leafs.Append("load", types.YLeaf{"Load", interfaceMetric.Load})

    interfaceMetric.EntityData.YListKeys = []string {}

    return &(interfaceMetric.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_RemoteNeighbor
// Remote-Neighbors enabled, default is 65535
// This type is a presence type.
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_RemoteNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable Remote neighbor unicast-listen. The type is bool. This attribute is
    // mandatory.
    Enable interface{}

    // Policy name. The type is string.
    AllowList interface{}

    // Neighbor count. The type is interface{} with range: 1..65535.
    MaxNeighbors interface{}
}

func (remoteNeighbor *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_RemoteNeighbor) GetEntityData() *types.CommonEntityData {
    remoteNeighbor.EntityData.YFilter = remoteNeighbor.YFilter
    remoteNeighbor.EntityData.YangName = "remote-neighbor"
    remoteNeighbor.EntityData.BundleName = "cisco_ios_xr"
    remoteNeighbor.EntityData.ParentYangName = "interface"
    remoteNeighbor.EntityData.SegmentPath = "remote-neighbor"
    remoteNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/interfaces/interface/" + remoteNeighbor.EntityData.SegmentPath
    remoteNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNeighbor.EntityData.Children = types.NewOrderedMap()
    remoteNeighbor.EntityData.Leafs = types.NewOrderedMap()
    remoteNeighbor.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", remoteNeighbor.Enable})
    remoteNeighbor.EntityData.Leafs.Append("allow-list", types.YLeaf{"AllowList", remoteNeighbor.AllowList})
    remoteNeighbor.EntityData.Leafs.Append("max-neighbors", types.YLeaf{"MaxNeighbors", remoteNeighbor.MaxNeighbors})

    remoteNeighbor.EntityData.YListKeys = []string {}

    return &(remoteNeighbor.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_Bfd
// Configure BFD parameters
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BFD fast detection. The type is bool.
    FastDetect interface{}

    // Detect multiplier. The type is interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 15..3000. Units are millisecond.
    Interval interface{}
}

func (bfd *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "interface"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/interfaces/interface/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("fast-detect", types.YLeaf{"FastDetect", bfd.FastDetect})
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_SiteOfOrigin
// Configure Site-of-origin
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_SiteOfOrigin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SoO type. The type is EigrpSoo.
    Type interface{}

    // Higher sixteen bits of 4-byte BGP AS Number. The type is interface{} with
    // range: 0..65535.
    AsXx interface{}

    // 2-byte or 4-byte BGP AS Number. The type is interface{} with range:
    // 0..4294967295.
    AsYy interface{}

    // AS Number Index. The type is interface{} with range: 0..4294967295.
    Index interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // IPv4 address index. The type is interface{} with range: 0..65535.
    AddressIndex interface{}
}

func (siteOfOrigin *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_SiteOfOrigin) GetEntityData() *types.CommonEntityData {
    siteOfOrigin.EntityData.YFilter = siteOfOrigin.YFilter
    siteOfOrigin.EntityData.YangName = "site-of-origin"
    siteOfOrigin.EntityData.BundleName = "cisco_ios_xr"
    siteOfOrigin.EntityData.ParentYangName = "interface"
    siteOfOrigin.EntityData.SegmentPath = "site-of-origin"
    siteOfOrigin.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/interfaces/interface/" + siteOfOrigin.EntityData.SegmentPath
    siteOfOrigin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    siteOfOrigin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    siteOfOrigin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    siteOfOrigin.EntityData.Children = types.NewOrderedMap()
    siteOfOrigin.EntityData.Leafs = types.NewOrderedMap()
    siteOfOrigin.EntityData.Leafs.Append("type", types.YLeaf{"Type", siteOfOrigin.Type})
    siteOfOrigin.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", siteOfOrigin.AsXx})
    siteOfOrigin.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", siteOfOrigin.AsYy})
    siteOfOrigin.EntityData.Leafs.Append("index", types.YLeaf{"Index", siteOfOrigin.Index})
    siteOfOrigin.EntityData.Leafs.Append("address", types.YLeaf{"Address", siteOfOrigin.Address})
    siteOfOrigin.EntityData.Leafs.Append("address-index", types.YLeaf{"AddressIndex", siteOfOrigin.AddressIndex})

    siteOfOrigin.EntityData.YListKeys = []string {}

    return &(siteOfOrigin.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_Authentication
// Authentication configuration
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authentication keychain configuration. The type is string.
    Keychain interface{}
}

func (authentication *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "interface"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/interfaces/interface/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("keychain", types.YLeaf{"Keychain", authentication.Keychain})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_SummaryAddresses
// List of summary addresses under this interface
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_SummaryAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary address. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_SummaryAddresses_SummaryAddress.
    SummaryAddress []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_SummaryAddresses_SummaryAddress
}

func (summaryAddresses *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_SummaryAddresses) GetEntityData() *types.CommonEntityData {
    summaryAddresses.EntityData.YFilter = summaryAddresses.YFilter
    summaryAddresses.EntityData.YangName = "summary-addresses"
    summaryAddresses.EntityData.BundleName = "cisco_ios_xr"
    summaryAddresses.EntityData.ParentYangName = "interface"
    summaryAddresses.EntityData.SegmentPath = "summary-addresses"
    summaryAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/interfaces/interface/" + summaryAddresses.EntityData.SegmentPath
    summaryAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryAddresses.EntityData.Children = types.NewOrderedMap()
    summaryAddresses.EntityData.Children.Append("summary-address", types.YChild{"SummaryAddress", nil})
    for i := range summaryAddresses.SummaryAddress {
        summaryAddresses.EntityData.Children.Append(types.GetSegmentPath(summaryAddresses.SummaryAddress[i]), types.YChild{"SummaryAddress", summaryAddresses.SummaryAddress[i]})
    }
    summaryAddresses.EntityData.Leafs = types.NewOrderedMap()

    summaryAddresses.EntityData.YListKeys = []string {}

    return &(summaryAddresses.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_SummaryAddresses_SummaryAddress
// Summary address
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_SummaryAddresses_SummaryAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Summary Prefix (address part). The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SummaryAddressAddr interface{}

    // This attribute is a key. Summary Prefix (prefix part). The type is
    // interface{} with range: 0..128.
    SummaryAddressPrefix interface{}

    // Administrative distance. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    Distance interface{}
}

func (summaryAddress *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_SummaryAddresses_SummaryAddress) GetEntityData() *types.CommonEntityData {
    summaryAddress.EntityData.YFilter = summaryAddress.YFilter
    summaryAddress.EntityData.YangName = "summary-address"
    summaryAddress.EntityData.BundleName = "cisco_ios_xr"
    summaryAddress.EntityData.ParentYangName = "summary-addresses"
    summaryAddress.EntityData.SegmentPath = "summary-address" + types.AddKeyToken(summaryAddress.SummaryAddressAddr, "summary-address-addr") + types.AddKeyToken(summaryAddress.SummaryAddressPrefix, "summary-address-prefix")
    summaryAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/interfaces/interface/summary-addresses/" + summaryAddress.EntityData.SegmentPath
    summaryAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryAddress.EntityData.Children = types.NewOrderedMap()
    summaryAddress.EntityData.Leafs = types.NewOrderedMap()
    summaryAddress.EntityData.Leafs.Append("summary-address-addr", types.YLeaf{"SummaryAddressAddr", summaryAddress.SummaryAddressAddr})
    summaryAddress.EntityData.Leafs.Append("summary-address-prefix", types.YLeaf{"SummaryAddressPrefix", summaryAddress.SummaryAddressPrefix})
    summaryAddress.EntityData.Leafs.Append("distance", types.YLeaf{"Distance", summaryAddress.Distance})

    summaryAddress.EntityData.YListKeys = []string {"SummaryAddressAddr", "SummaryAddressPrefix"}

    return &(summaryAddress.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceFilterPolicies
// List of filter policies
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceFilterPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceFilterPolicies_InterfaceFilterPolicy.
    InterfaceFilterPolicy []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceFilterPolicies_InterfaceFilterPolicy
}

func (interfaceFilterPolicies *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceFilterPolicies) GetEntityData() *types.CommonEntityData {
    interfaceFilterPolicies.EntityData.YFilter = interfaceFilterPolicies.YFilter
    interfaceFilterPolicies.EntityData.YangName = "interface-filter-policies"
    interfaceFilterPolicies.EntityData.BundleName = "cisco_ios_xr"
    interfaceFilterPolicies.EntityData.ParentYangName = "interface"
    interfaceFilterPolicies.EntityData.SegmentPath = "interface-filter-policies"
    interfaceFilterPolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/interfaces/interface/" + interfaceFilterPolicies.EntityData.SegmentPath
    interfaceFilterPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFilterPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFilterPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFilterPolicies.EntityData.Children = types.NewOrderedMap()
    interfaceFilterPolicies.EntityData.Children.Append("interface-filter-policy", types.YChild{"InterfaceFilterPolicy", nil})
    for i := range interfaceFilterPolicies.InterfaceFilterPolicy {
        interfaceFilterPolicies.EntityData.Children.Append(types.GetSegmentPath(interfaceFilterPolicies.InterfaceFilterPolicy[i]), types.YChild{"InterfaceFilterPolicy", interfaceFilterPolicies.InterfaceFilterPolicy[i]})
    }
    interfaceFilterPolicies.EntityData.Leafs = types.NewOrderedMap()

    interfaceFilterPolicies.EntityData.YListKeys = []string {}

    return &(interfaceFilterPolicies.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceFilterPolicies_InterfaceFilterPolicy
// none
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceFilterPolicies_InterfaceFilterPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Direction (in or out). The type is EigrpDir.
    Direction interface{}

    // Policy name. The type is string. This attribute is mandatory.
    PolicyName interface{}
}

func (interfaceFilterPolicy *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceFilterPolicies_InterfaceFilterPolicy) GetEntityData() *types.CommonEntityData {
    interfaceFilterPolicy.EntityData.YFilter = interfaceFilterPolicy.YFilter
    interfaceFilterPolicy.EntityData.YangName = "interface-filter-policy"
    interfaceFilterPolicy.EntityData.BundleName = "cisco_ios_xr"
    interfaceFilterPolicy.EntityData.ParentYangName = "interface-filter-policies"
    interfaceFilterPolicy.EntityData.SegmentPath = "interface-filter-policy" + types.AddKeyToken(interfaceFilterPolicy.Direction, "direction")
    interfaceFilterPolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/interfaces/interface/interface-filter-policies/" + interfaceFilterPolicy.EntityData.SegmentPath
    interfaceFilterPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFilterPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFilterPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFilterPolicy.EntityData.Children = types.NewOrderedMap()
    interfaceFilterPolicy.EntityData.Leafs = types.NewOrderedMap()
    interfaceFilterPolicy.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", interfaceFilterPolicy.Direction})
    interfaceFilterPolicy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", interfaceFilterPolicy.PolicyName})

    interfaceFilterPolicy.EntityData.YListKeys = []string {"Direction"}

    return &(interfaceFilterPolicy.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceStaticNeighbors
// List of Neighbors
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceStaticNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Neighbor. The type is slice of
    // Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceStaticNeighbors_InterfaceStaticNeighbor.
    InterfaceStaticNeighbor []*Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceStaticNeighbors_InterfaceStaticNeighbor
}

func (interfaceStaticNeighbors *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceStaticNeighbors) GetEntityData() *types.CommonEntityData {
    interfaceStaticNeighbors.EntityData.YFilter = interfaceStaticNeighbors.YFilter
    interfaceStaticNeighbors.EntityData.YangName = "interface-static-neighbors"
    interfaceStaticNeighbors.EntityData.BundleName = "cisco_ios_xr"
    interfaceStaticNeighbors.EntityData.ParentYangName = "interface"
    interfaceStaticNeighbors.EntityData.SegmentPath = "interface-static-neighbors"
    interfaceStaticNeighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/interfaces/interface/" + interfaceStaticNeighbors.EntityData.SegmentPath
    interfaceStaticNeighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStaticNeighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStaticNeighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStaticNeighbors.EntityData.Children = types.NewOrderedMap()
    interfaceStaticNeighbors.EntityData.Children.Append("interface-static-neighbor", types.YChild{"InterfaceStaticNeighbor", nil})
    for i := range interfaceStaticNeighbors.InterfaceStaticNeighbor {
        interfaceStaticNeighbors.EntityData.Children.Append(types.GetSegmentPath(interfaceStaticNeighbors.InterfaceStaticNeighbor[i]), types.YChild{"InterfaceStaticNeighbor", interfaceStaticNeighbors.InterfaceStaticNeighbor[i]})
    }
    interfaceStaticNeighbors.EntityData.Leafs = types.NewOrderedMap()

    interfaceStaticNeighbors.EntityData.YListKeys = []string {}

    return &(interfaceStaticNeighbors.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceStaticNeighbors_InterfaceStaticNeighbor
// Configure Neighbor
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceStaticNeighbors_InterfaceStaticNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NeighborAddress interface{}

    // Number of hops. The type is interface{} with range: 2..100. This attribute
    // is mandatory.
    MaxHops interface{}
}

func (interfaceStaticNeighbor *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Interfaces_Interface_InterfaceStaticNeighbors_InterfaceStaticNeighbor) GetEntityData() *types.CommonEntityData {
    interfaceStaticNeighbor.EntityData.YFilter = interfaceStaticNeighbor.YFilter
    interfaceStaticNeighbor.EntityData.YangName = "interface-static-neighbor"
    interfaceStaticNeighbor.EntityData.BundleName = "cisco_ios_xr"
    interfaceStaticNeighbor.EntityData.ParentYangName = "interface-static-neighbors"
    interfaceStaticNeighbor.EntityData.SegmentPath = "interface-static-neighbor" + types.AddKeyToken(interfaceStaticNeighbor.NeighborAddress, "neighbor-address")
    interfaceStaticNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/interfaces/interface/interface-static-neighbors/" + interfaceStaticNeighbor.EntityData.SegmentPath
    interfaceStaticNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStaticNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStaticNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStaticNeighbor.EntityData.Children = types.NewOrderedMap()
    interfaceStaticNeighbor.EntityData.Leafs = types.NewOrderedMap()
    interfaceStaticNeighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", interfaceStaticNeighbor.NeighborAddress})
    interfaceStaticNeighbor.EntityData.Leafs.Append("max-hops", types.YLeaf{"MaxHops", interfaceStaticNeighbor.MaxHops})

    interfaceStaticNeighbor.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(interfaceStaticNeighbor.EntityData)
}

// Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Distance
// Set distance for EIGRP routes
type Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Distance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Internal routes' distance. The type is interface{} with range: 1..255.
    InternalDistance interface{}

    // External routes' distance. The type is interface{} with range: 1..255.
    ExternalDistance interface{}
}

func (distance *Eigrp_Processes_Process_DefaultVrf_DefaultAfs_DefaultAf_Distance) GetEntityData() *types.CommonEntityData {
    distance.EntityData.YFilter = distance.YFilter
    distance.EntityData.YangName = "distance"
    distance.EntityData.BundleName = "cisco_ios_xr"
    distance.EntityData.ParentYangName = "default-af"
    distance.EntityData.SegmentPath = "distance"
    distance.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-cfg:eigrp/processes/process/default-vrf/default-afs/default-af/" + distance.EntityData.SegmentPath
    distance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distance.EntityData.Children = types.NewOrderedMap()
    distance.EntityData.Leafs = types.NewOrderedMap()
    distance.EntityData.Leafs.Append("internal-distance", types.YLeaf{"InternalDistance", distance.InternalDistance})
    distance.EntityData.Leafs.Append("external-distance", types.YLeaf{"ExternalDistance", distance.ExternalDistance})

    distance.EntityData.YListKeys = []string {}

    return &(distance.EntityData)
}

