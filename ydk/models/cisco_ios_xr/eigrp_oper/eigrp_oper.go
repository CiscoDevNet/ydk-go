// This module contains a collection of YANG definitions
// for Cisco IOS-XR eigrp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   eigrp: EIGRP operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package eigrp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package eigrp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-eigrp-oper eigrp}", reflect.TypeOf(Eigrp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-eigrp-oper:eigrp", reflect.TypeOf(Eigrp{}))
}

// EigrpBdSoo represents EIGRP SoO types
type EigrpBdSoo string

const (
    // No SoO configured
    EigrpBdSoo_none EigrpBdSoo = "none"

    // AS:nn format
    EigrpBdSoo_as EigrpBdSoo = "as"

    // IPV4Address:nn format
    EigrpBdSoo_ipv4_address EigrpBdSoo = "ipv4-address"

    // AS2.AS:nn format
    EigrpBdSoo_four_byte_as EigrpBdSoo = "four-byte-as"
)

// EigrpBdPathRibState represents Eigrp bd path rib state
type EigrpBdPathRibState string

const (
    // Active path
    EigrpBdPathRibState_active_path EigrpBdPathRibState = "active-path"

    // Backup path
    EigrpBdPathRibState_backup_path EigrpBdPathRibState = "backup-path"

    // Path sent to RIB
    EigrpBdPathRibState_path_sent_to_rib EigrpBdPathRibState = "path-sent-to-rib"

    // Path not selected for installation in RIB
    EigrpBdPathRibState_path_not_selected EigrpBdPathRibState = "path-not-selected"

    // Path in error state
    EigrpBdPathRibState_error_state EigrpBdPathRibState = "error-state"
)

// EigrpBdDelayUnit represents EIGRP delay unit
type EigrpBdDelayUnit string

const (
    // No Delay configured
    EigrpBdDelayUnit_none EigrpBdDelayUnit = "none"

    // Delay in 10's of Microseconds
    EigrpBdDelayUnit_ten_microsecond EigrpBdDelayUnit = "ten-microsecond"

    // Delay in Picoseconds
    EigrpBdDelayUnit_picosecond EigrpBdDelayUnit = "picosecond"

    // Delay in Microseconds
    EigrpBdDelayUnit_microsecond EigrpBdDelayUnit = "microsecond"
)

// EigrpBdPathOrigin represents EIGRP path origin
type EigrpBdPathOrigin string

const (
    // connected destination
    EigrpBdPathOrigin_connected EigrpBdPathOrigin = "connected"

    // static redistribution
    EigrpBdPathOrigin_static_redistributed EigrpBdPathOrigin = "static-redistributed"

    // connected redistribution
    EigrpBdPathOrigin_connected_redistributed EigrpBdPathOrigin = "connected-redistributed"

    // subscriber redistribution
    EigrpBdPathOrigin_subscriber_redistributed EigrpBdPathOrigin = "subscriber-redistributed"

    // redistributed destination
    EigrpBdPathOrigin_redistributed EigrpBdPathOrigin = "redistributed"

    // vpnv4 destination
    EigrpBdPathOrigin_vpnv4_sourced EigrpBdPathOrigin = "vpnv4-sourced"

    // vpnv6 destination
    EigrpBdPathOrigin_vpnv6_sourced EigrpBdPathOrigin = "vpnv6-sourced"

    // summary destination
    EigrpBdPathOrigin_summary EigrpBdPathOrigin = "summary"

    // bogus drdb used for sia transmission
    EigrpBdPathOrigin_dummy EigrpBdPathOrigin = "dummy"

    // igrp2 destination
    EigrpBdPathOrigin_eigrp_destination EigrpBdPathOrigin = "eigrp-destination"

    // Number of org types
    EigrpBdPathOrigin_origin_count EigrpBdPathOrigin = "origin-count"
)

// EigrpBdPathSendFlag represents EIGRP path send flag
type EigrpBdPathSendFlag string

const (
    // No packet send pending
    EigrpBdPathSendFlag_no_send_pending EigrpBdPathSendFlag = "no-send-pending"

    // Multicast update pending
    EigrpBdPathSendFlag_multicast_update_pending EigrpBdPathSendFlag = "multicast-update-pending"

    // Multicast query pending
    EigrpBdPathSendFlag_multicast_query_pending EigrpBdPathSendFlag = "multicast-query-pending"

    // Reply pending
    EigrpBdPathSendFlag_reply_pending EigrpBdPathSendFlag = "reply-pending"

    // SIA Query pending
    EigrpBdPathSendFlag_sia_query_pending EigrpBdPathSendFlag = "sia-query-pending"

    // SIA Reply pending
    EigrpBdPathSendFlag_sia_reply_pending EigrpBdPathSendFlag = "sia-reply-pending"
)

// EigrpBdMetricVersion represents EIGRP metric version
type EigrpBdMetricVersion string

const (
    // Metric version is 32 bit
    EigrpBdMetricVersion_metric_version32_bit EigrpBdMetricVersion = "metric-version32-bit"

    // Metric version is 64 bit
    EigrpBdMetricVersion_metric_version64_bit EigrpBdMetricVersion = "metric-version64-bit"
)

// Eigrp
// EIGRP operational data
type Eigrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for an EIGRP process.
    Processes Eigrp_Processes
}

func (eigrp *Eigrp) GetEntityData() *types.CommonEntityData {
    eigrp.EntityData.YFilter = eigrp.YFilter
    eigrp.EntityData.YangName = "eigrp"
    eigrp.EntityData.BundleName = "cisco_ios_xr"
    eigrp.EntityData.ParentYangName = "Cisco-IOS-XR-eigrp-oper"
    eigrp.EntityData.SegmentPath = "Cisco-IOS-XR-eigrp-oper:eigrp"
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
// Operational data for an EIGRP process
type Eigrp_Processes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for an EIGRP process. The type is slice of
    // Eigrp_Processes_Process.
    Process []*Eigrp_Processes_Process
}

func (processes *Eigrp_Processes) GetEntityData() *types.CommonEntityData {
    processes.EntityData.YFilter = processes.YFilter
    processes.EntityData.YangName = "processes"
    processes.EntityData.BundleName = "cisco_ios_xr"
    processes.EntityData.ParentYangName = "eigrp"
    processes.EntityData.SegmentPath = "processes"
    processes.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/" + processes.EntityData.SegmentPath
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
// Operational data for an EIGRP process
type Eigrp_Processes_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. AS number or Virtual instance name of the EIGRP
    // process. The type is string with length: 1..32.
    ProcessId interface{}

    // List of VRFs.
    VrfsXr Eigrp_Processes_Process_VrfsXr

    // List of VRFs.
    Vrfs Eigrp_Processes_Process_Vrfs
}

func (process *Eigrp_Processes_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "processes"
    process.EntityData.SegmentPath = "process" + types.AddKeyToken(process.ProcessId, "process-id")
    process.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/" + process.EntityData.SegmentPath
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Children.Append("vrfs-xr", types.YChild{"VrfsXr", &process.VrfsXr})
    process.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &process.Vrfs})
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", process.ProcessId})

    process.EntityData.YListKeys = []string {"ProcessId"}

    return &(process.EntityData)
}

// Eigrp_Processes_Process_VrfsXr
// List of VRFs
type Eigrp_Processes_Process_VrfsXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRF. The type is slice of Eigrp_Processes_Process_VrfsXr_Vrf.
    Vrf []*Eigrp_Processes_Process_VrfsXr_Vrf
}

func (vrfsXr *Eigrp_Processes_Process_VrfsXr) GetEntityData() *types.CommonEntityData {
    vrfsXr.EntityData.YFilter = vrfsXr.YFilter
    vrfsXr.EntityData.YangName = "vrfs-xr"
    vrfsXr.EntityData.BundleName = "cisco_ios_xr"
    vrfsXr.EntityData.ParentYangName = "process"
    vrfsXr.EntityData.SegmentPath = "vrfs-xr"
    vrfsXr.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/" + vrfsXr.EntityData.SegmentPath
    vrfsXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfsXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfsXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfsXr.EntityData.Children = types.NewOrderedMap()
    vrfsXr.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfsXr.Vrf {
        vrfsXr.EntityData.Children.Append(types.GetSegmentPath(vrfsXr.Vrf[i]), types.YChild{"Vrf", vrfsXr.Vrf[i]})
    }
    vrfsXr.EntityData.Leafs = types.NewOrderedMap()

    vrfsXr.EntityData.YListKeys = []string {}

    return &(vrfsXr.EntityData)
}

// Eigrp_Processes_Process_VrfsXr_Vrf
// A VRF
type Eigrp_Processes_Process_VrfsXr_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // VRF Name. The type is string.
    VrfNameXr interface{}
}

func (vrf *Eigrp_Processes_Process_VrfsXr_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs-xr"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs-xr/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("vrf-name-xr", types.YLeaf{"VrfNameXr", vrf.VrfNameXr})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Eigrp_Processes_Process_Vrfs
// List of VRFs
type Eigrp_Processes_Process_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a VRF. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf.
    Vrf []*Eigrp_Processes_Process_Vrfs_Vrf
}

func (vrfs *Eigrp_Processes_Process_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "process"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/" + vrfs.EntityData.SegmentPath
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
// Operational data for a VRF
type Eigrp_Processes_Process_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // List of Address Families.
    Afs Eigrp_Processes_Process_Vrfs_Vrf_Afs
}

func (vrf *Eigrp_Processes_Process_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("afs", types.YChild{"Afs", &vrf.Afs})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs
// List of Address Families
type Eigrp_Processes_Process_Vrfs_Vrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for one AF. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af.
    Af []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af
}

func (afs *Eigrp_Processes_Process_Vrfs_Vrf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "vrf"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/" + afs.EntityData.SegmentPath
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
// Operational data for one AF
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address Family. The type is EigrpAf.
    AfName interface{}

    // Address family specific protocol information.
    Protocol Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol

    // List of Autonomous Systems.
    Ases Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases
}

func (af *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("protocol", types.YChild{"Protocol", &af.Protocol})
    af.EntityData.Children.Append("ases", types.YChild{"Ases", &af.Ases})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol
// Address family specific protocol
// information
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // AS Number. The type is interface{} with range: 0..4294967295.
    AsNumber interface{}

    // Router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Auto Summarization. The type is bool.
    AutoSummarization interface{}

    // Neighbor changes logged. The type is bool.
    LogNeighborChanges interface{}

    // Neighbor warnings logged. The type is bool.
    LogNeighborWarnings interface{}

    // RIB Table limit has been reached. The type is bool.
    RibTableLimitReached interface{}

    // Outbound Filter Policy. The type is string.
    OutboundFilterPolicy interface{}

    // Inbound Filter Policy. The type is string.
    InboundFilterPolicy interface{}

    // Default Allowed Out. The type is bool.
    OutgoingCandidateDefaultFlagged interface{}

    // Default Allowed Out Policy. The type is string.
    OutgoingCandidateDefaultPolicy interface{}

    // Default Allowed In. The type is bool.
    IncomingCandidateDefaultFlagged interface{}

    // Default Allowed In Policy. The type is string.
    IncomingCandidateDefaultPolicy interface{}

    // Internal Distance. The type is interface{} with range: 0..255.
    InternalDistance interface{}

    // External Distance. The type is interface{} with range: 0..255.
    ExternalDistance interface{}

    // Maximum paths. The type is interface{} with range: 0..255.
    MaximumPaths interface{}

    // Variance. The type is interface{} with range: 0..255.
    Variance interface{}

    // K1 value. The type is interface{} with range: 0..4294967295.
    MetricWeightK1 interface{}

    // K2 value. The type is interface{} with range: 0..4294967295.
    MetricWeightK2 interface{}

    // K3 value. The type is interface{} with range: 0..4294967295.
    MetricWeightK3 interface{}

    // K4 value. The type is interface{} with range: 0..4294967295.
    MetricWeightK4 interface{}

    // K5 value. The type is interface{} with range: 0..4294967295.
    MetricWeightK5 interface{}

    // K6 value. The type is interface{} with range: 0..4294967295.
    MetricWeightK6 interface{}

    // RIB Scale. The type is interface{} with range: 0..4294967295.
    RibScale interface{}

    // Metric Version. The type is EigrpBdMetricVersion.
    MetricVersion interface{}

    // Metric MaxHops configured. The type is interface{} with range:
    // 0..4294967295.
    MetricMaximumHopcount interface{}

    // Default Metric Configured. The type is bool.
    DefaultMetricConfigured interface{}

    // Default Bandwidth. The type is interface{} with range: 0..4294967295.
    DefaultBandwidth interface{}

    // Default Delay. The type is interface{} with range: 0..4294967295.
    DefaultDelay interface{}

    // Default Reliability. The type is interface{} with range: 0..4294967295.
    DefaultReliability interface{}

    // Default Load. The type is interface{} with range: 0..4294967295.
    DefaultLoad interface{}

    // Default MTU. The type is interface{} with range: 0..4294967295.
    DefaultMtu interface{}

    // Stub Configured. The type is bool.
    StubConfigured interface{}

    // Stub Receive-only configured. The type is bool.
    StubReceiveOnly interface{}

    // ConnectedRoutes allowed. The type is bool.
    StubAllowConnectedRoutes interface{}

    // Static Routes allowed. The type is bool.
    StubAllowStaticRoutes interface{}

    // Summary Routes allowed. The type is bool.
    StubAllowSummaryRoutes interface{}

    // Redistributed Routes allowed. The type is bool.
    StubAllowRedistributedRoutes interface{}

    // NSF Enabled. The type is bool.
    NsfEnabled interface{}

    // NSF Route Hold Time. The type is interface{} with range: 0..4294967295.
    NsfRouteHoldTime interface{}

    // NSF Signal Time. The type is interface{} with range: 0..4294967295.
    NsfSignalTime interface{}

    // NSF Converge Time. The type is interface{} with range: 0..4294967295.
    NsfConvergeTime interface{}

    // Is Restart time configured. The type is bool.
    RestartConfigured interface{}

    // Restart time (seconds). The type is interface{} with range: 0..4294967295.
    // Units are second.
    RestartTime interface{}

    // SIA Active Time. The type is interface{} with range: 0..4294967295.
    SiaActiveTime interface{}

    // RIB Protocol ID. The type is interface{} with range: 0..4294967295.
    RibProtocolId interface{}

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // VRF activated by ITAL. The type is bool.
    ItalActivationReceived interface{}

    // VRF activated by EIGRP. The type is bool.
    VrfActivated interface{}

    // VRF information available. The type is bool.
    Up interface{}

    // RIB initialization for VRF. The type is bool.
    RibInitialized interface{}

    // RIB convergence for VRF. The type is bool.
    RibConverged interface{}

    // Reload following RIB Convergence. The type is bool.
    RibConvergedReload interface{}

    // Requested Socket Option for VRF. The type is bool.
    SocketRequest interface{}

    // Setup socket state for VRF. The type is bool.
    SocketSetup interface{}

    // VRF represents default-context. The type is bool.
    DefaultVrf interface{}

    // AF Enabled. The type is bool.
    AfEnabled interface{}

    // Passive-Interface default configured. The type is bool.
    IsPassiveDefault interface{}

    // VRF Configured Items. The type is interface{} with range: 0..4294967295.
    ConfiguredItems interface{}

    // AF Configured Items. The type is interface{} with range: 0..4294967295.
    AfConfiguredItems interface{}

    // IP ARM Router ID. The type is interface{} with range: 0..4294967295.
    IpArmRouterId interface{}

    // IP Address of first UP interface. The type is interface{} with range:
    // 0..4294967295.
    FirstInterfaceUpAddress interface{}

    // DDB NSF in progress indication. The type is interface{} with range:
    // 0..4294967295.
    NsfInProgress interface{}

    // RIB Table convergence indication. The type is interface{} with range:
    // 0..4294967295.
    RibTableConverged interface{}

    // Redistributed Protocols. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol_RedistributedProtocol.
    RedistributedProtocol []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol_RedistributedProtocol

    // Interfaces. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol_Interface.
    Interface []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol_Interface
}

func (protocol *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol) GetEntityData() *types.CommonEntityData {
    protocol.EntityData.YFilter = protocol.YFilter
    protocol.EntityData.YangName = "protocol"
    protocol.EntityData.BundleName = "cisco_ios_xr"
    protocol.EntityData.ParentYangName = "af"
    protocol.EntityData.SegmentPath = "protocol"
    protocol.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/" + protocol.EntityData.SegmentPath
    protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocol.EntityData.Children = types.NewOrderedMap()
    protocol.EntityData.Children.Append("redistributed-protocol", types.YChild{"RedistributedProtocol", nil})
    for i := range protocol.RedistributedProtocol {
        types.SetYListKey(protocol.RedistributedProtocol[i], i)
        protocol.EntityData.Children.Append(types.GetSegmentPath(protocol.RedistributedProtocol[i]), types.YChild{"RedistributedProtocol", protocol.RedistributedProtocol[i]})
    }
    protocol.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range protocol.Interface {
        types.SetYListKey(protocol.Interface[i], i)
        protocol.EntityData.Children.Append(types.GetSegmentPath(protocol.Interface[i]), types.YChild{"Interface", protocol.Interface[i]})
    }
    protocol.EntityData.Leafs = types.NewOrderedMap()
    protocol.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", protocol.Afi})
    protocol.EntityData.Leafs.Append("as-number", types.YLeaf{"AsNumber", protocol.AsNumber})
    protocol.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", protocol.RouterId})
    protocol.EntityData.Leafs.Append("auto-summarization", types.YLeaf{"AutoSummarization", protocol.AutoSummarization})
    protocol.EntityData.Leafs.Append("log-neighbor-changes", types.YLeaf{"LogNeighborChanges", protocol.LogNeighborChanges})
    protocol.EntityData.Leafs.Append("log-neighbor-warnings", types.YLeaf{"LogNeighborWarnings", protocol.LogNeighborWarnings})
    protocol.EntityData.Leafs.Append("rib-table-limit-reached", types.YLeaf{"RibTableLimitReached", protocol.RibTableLimitReached})
    protocol.EntityData.Leafs.Append("outbound-filter-policy", types.YLeaf{"OutboundFilterPolicy", protocol.OutboundFilterPolicy})
    protocol.EntityData.Leafs.Append("inbound-filter-policy", types.YLeaf{"InboundFilterPolicy", protocol.InboundFilterPolicy})
    protocol.EntityData.Leafs.Append("outgoing-candidate-default-flagged", types.YLeaf{"OutgoingCandidateDefaultFlagged", protocol.OutgoingCandidateDefaultFlagged})
    protocol.EntityData.Leafs.Append("outgoing-candidate-default-policy", types.YLeaf{"OutgoingCandidateDefaultPolicy", protocol.OutgoingCandidateDefaultPolicy})
    protocol.EntityData.Leafs.Append("incoming-candidate-default-flagged", types.YLeaf{"IncomingCandidateDefaultFlagged", protocol.IncomingCandidateDefaultFlagged})
    protocol.EntityData.Leafs.Append("incoming-candidate-default-policy", types.YLeaf{"IncomingCandidateDefaultPolicy", protocol.IncomingCandidateDefaultPolicy})
    protocol.EntityData.Leafs.Append("internal-distance", types.YLeaf{"InternalDistance", protocol.InternalDistance})
    protocol.EntityData.Leafs.Append("external-distance", types.YLeaf{"ExternalDistance", protocol.ExternalDistance})
    protocol.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", protocol.MaximumPaths})
    protocol.EntityData.Leafs.Append("variance", types.YLeaf{"Variance", protocol.Variance})
    protocol.EntityData.Leafs.Append("metric-weight-k1", types.YLeaf{"MetricWeightK1", protocol.MetricWeightK1})
    protocol.EntityData.Leafs.Append("metric-weight-k2", types.YLeaf{"MetricWeightK2", protocol.MetricWeightK2})
    protocol.EntityData.Leafs.Append("metric-weight-k3", types.YLeaf{"MetricWeightK3", protocol.MetricWeightK3})
    protocol.EntityData.Leafs.Append("metric-weight-k4", types.YLeaf{"MetricWeightK4", protocol.MetricWeightK4})
    protocol.EntityData.Leafs.Append("metric-weight-k5", types.YLeaf{"MetricWeightK5", protocol.MetricWeightK5})
    protocol.EntityData.Leafs.Append("metric-weight-k6", types.YLeaf{"MetricWeightK6", protocol.MetricWeightK6})
    protocol.EntityData.Leafs.Append("rib-scale", types.YLeaf{"RibScale", protocol.RibScale})
    protocol.EntityData.Leafs.Append("metric-version", types.YLeaf{"MetricVersion", protocol.MetricVersion})
    protocol.EntityData.Leafs.Append("metric-maximum-hopcount", types.YLeaf{"MetricMaximumHopcount", protocol.MetricMaximumHopcount})
    protocol.EntityData.Leafs.Append("default-metric-configured", types.YLeaf{"DefaultMetricConfigured", protocol.DefaultMetricConfigured})
    protocol.EntityData.Leafs.Append("default-bandwidth", types.YLeaf{"DefaultBandwidth", protocol.DefaultBandwidth})
    protocol.EntityData.Leafs.Append("default-delay", types.YLeaf{"DefaultDelay", protocol.DefaultDelay})
    protocol.EntityData.Leafs.Append("default-reliability", types.YLeaf{"DefaultReliability", protocol.DefaultReliability})
    protocol.EntityData.Leafs.Append("default-load", types.YLeaf{"DefaultLoad", protocol.DefaultLoad})
    protocol.EntityData.Leafs.Append("default-mtu", types.YLeaf{"DefaultMtu", protocol.DefaultMtu})
    protocol.EntityData.Leafs.Append("stub-configured", types.YLeaf{"StubConfigured", protocol.StubConfigured})
    protocol.EntityData.Leafs.Append("stub-receive-only", types.YLeaf{"StubReceiveOnly", protocol.StubReceiveOnly})
    protocol.EntityData.Leafs.Append("stub-allow-connected-routes", types.YLeaf{"StubAllowConnectedRoutes", protocol.StubAllowConnectedRoutes})
    protocol.EntityData.Leafs.Append("stub-allow-static-routes", types.YLeaf{"StubAllowStaticRoutes", protocol.StubAllowStaticRoutes})
    protocol.EntityData.Leafs.Append("stub-allow-summary-routes", types.YLeaf{"StubAllowSummaryRoutes", protocol.StubAllowSummaryRoutes})
    protocol.EntityData.Leafs.Append("stub-allow-redistributed-routes", types.YLeaf{"StubAllowRedistributedRoutes", protocol.StubAllowRedistributedRoutes})
    protocol.EntityData.Leafs.Append("nsf-enabled", types.YLeaf{"NsfEnabled", protocol.NsfEnabled})
    protocol.EntityData.Leafs.Append("nsf-route-hold-time", types.YLeaf{"NsfRouteHoldTime", protocol.NsfRouteHoldTime})
    protocol.EntityData.Leafs.Append("nsf-signal-time", types.YLeaf{"NsfSignalTime", protocol.NsfSignalTime})
    protocol.EntityData.Leafs.Append("nsf-converge-time", types.YLeaf{"NsfConvergeTime", protocol.NsfConvergeTime})
    protocol.EntityData.Leafs.Append("restart-configured", types.YLeaf{"RestartConfigured", protocol.RestartConfigured})
    protocol.EntityData.Leafs.Append("restart-time", types.YLeaf{"RestartTime", protocol.RestartTime})
    protocol.EntityData.Leafs.Append("sia-active-time", types.YLeaf{"SiaActiveTime", protocol.SiaActiveTime})
    protocol.EntityData.Leafs.Append("rib-protocol-id", types.YLeaf{"RibProtocolId", protocol.RibProtocolId})
    protocol.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", protocol.TableId})
    protocol.EntityData.Leafs.Append("vrf-id", types.YLeaf{"VrfId", protocol.VrfId})
    protocol.EntityData.Leafs.Append("ital-activation-received", types.YLeaf{"ItalActivationReceived", protocol.ItalActivationReceived})
    protocol.EntityData.Leafs.Append("vrf-activated", types.YLeaf{"VrfActivated", protocol.VrfActivated})
    protocol.EntityData.Leafs.Append("up", types.YLeaf{"Up", protocol.Up})
    protocol.EntityData.Leafs.Append("rib-initialized", types.YLeaf{"RibInitialized", protocol.RibInitialized})
    protocol.EntityData.Leafs.Append("rib-converged", types.YLeaf{"RibConverged", protocol.RibConverged})
    protocol.EntityData.Leafs.Append("rib-converged-reload", types.YLeaf{"RibConvergedReload", protocol.RibConvergedReload})
    protocol.EntityData.Leafs.Append("socket-request", types.YLeaf{"SocketRequest", protocol.SocketRequest})
    protocol.EntityData.Leafs.Append("socket-setup", types.YLeaf{"SocketSetup", protocol.SocketSetup})
    protocol.EntityData.Leafs.Append("default-vrf", types.YLeaf{"DefaultVrf", protocol.DefaultVrf})
    protocol.EntityData.Leafs.Append("af-enabled", types.YLeaf{"AfEnabled", protocol.AfEnabled})
    protocol.EntityData.Leafs.Append("is-passive-default", types.YLeaf{"IsPassiveDefault", protocol.IsPassiveDefault})
    protocol.EntityData.Leafs.Append("configured-items", types.YLeaf{"ConfiguredItems", protocol.ConfiguredItems})
    protocol.EntityData.Leafs.Append("af-configured-items", types.YLeaf{"AfConfiguredItems", protocol.AfConfiguredItems})
    protocol.EntityData.Leafs.Append("ip-arm-router-id", types.YLeaf{"IpArmRouterId", protocol.IpArmRouterId})
    protocol.EntityData.Leafs.Append("first-interface-up-address", types.YLeaf{"FirstInterfaceUpAddress", protocol.FirstInterfaceUpAddress})
    protocol.EntityData.Leafs.Append("nsf-in-progress", types.YLeaf{"NsfInProgress", protocol.NsfInProgress})
    protocol.EntityData.Leafs.Append("rib-table-converged", types.YLeaf{"RibTableConverged", protocol.RibTableConverged})

    protocol.EntityData.YListKeys = []string {}

    return &(protocol.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol_RedistributedProtocol
// Redistributed Protocols
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol_RedistributedProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // Redistributed Protocol. The type is string.
    RedistributedProtocol interface{}

    // Redistributed Protocol tag. The type is string.
    RedistributedProtocolTag interface{}

    // Redistribute Filter policy. The type is string.
    RedristributePolicy interface{}

    // Redistributed Protocol ID. The type is interface{} with range:
    // 0..4294967295.
    RedistributeProtocolId interface{}

    // Redistributed Protocol handle. The type is interface{} with range:
    // 0..4294967295.
    RibHandle interface{}
}

func (redistributedProtocol *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol_RedistributedProtocol) GetEntityData() *types.CommonEntityData {
    redistributedProtocol.EntityData.YFilter = redistributedProtocol.YFilter
    redistributedProtocol.EntityData.YangName = "redistributed-protocol"
    redistributedProtocol.EntityData.BundleName = "cisco_ios_xr"
    redistributedProtocol.EntityData.ParentYangName = "protocol"
    redistributedProtocol.EntityData.SegmentPath = "redistributed-protocol" + types.AddNoKeyToken(redistributedProtocol)
    redistributedProtocol.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/protocol/" + redistributedProtocol.EntityData.SegmentPath
    redistributedProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributedProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributedProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributedProtocol.EntityData.Children = types.NewOrderedMap()
    redistributedProtocol.EntityData.Leafs = types.NewOrderedMap()
    redistributedProtocol.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", redistributedProtocol.Afi})
    redistributedProtocol.EntityData.Leafs.Append("redistributed-protocol", types.YLeaf{"RedistributedProtocol", redistributedProtocol.RedistributedProtocol})
    redistributedProtocol.EntityData.Leafs.Append("redistributed-protocol-tag", types.YLeaf{"RedistributedProtocolTag", redistributedProtocol.RedistributedProtocolTag})
    redistributedProtocol.EntityData.Leafs.Append("redristribute-policy", types.YLeaf{"RedristributePolicy", redistributedProtocol.RedristributePolicy})
    redistributedProtocol.EntityData.Leafs.Append("redistribute-protocol-id", types.YLeaf{"RedistributeProtocolId", redistributedProtocol.RedistributeProtocolId})
    redistributedProtocol.EntityData.Leafs.Append("rib-handle", types.YLeaf{"RibHandle", redistributedProtocol.RibHandle})

    redistributedProtocol.EntityData.YListKeys = []string {}

    return &(redistributedProtocol.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol_Interface
// Interfaces
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // Interface. The type is string.
    Interface interface{}

    // Outbound Filter Policy. The type is string.
    OutboundFilterPolicy interface{}

    // Inbound Filter Policy. The type is string.
    InboundFilterPolicy interface{}

    // Interface is DOWN. The type is bool.
    Inactive interface{}

    // Interface is passive. The type is bool.
    PassiveInterface interface{}
}

func (self *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Protocol_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "protocol"
    self.EntityData.SegmentPath = "interface" + types.AddNoKeyToken(self)
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/protocol/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", self.Afi})
    self.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", self.Interface})
    self.EntityData.Leafs.Append("outbound-filter-policy", types.YLeaf{"OutboundFilterPolicy", self.OutboundFilterPolicy})
    self.EntityData.Leafs.Append("inbound-filter-policy", types.YLeaf{"InboundFilterPolicy", self.InboundFilterPolicy})
    self.EntityData.Leafs.Append("inactive", types.YLeaf{"Inactive", self.Inactive})
    self.EntityData.Leafs.Append("passive-interface", types.YLeaf{"PassiveInterface", self.PassiveInterface})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases
// List of Autonomous Systems
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for one AS. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As.
    As []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As
}

func (ases *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases) GetEntityData() *types.CommonEntityData {
    ases.EntityData.YFilter = ases.YFilter
    ases.EntityData.YangName = "ases"
    ases.EntityData.BundleName = "cisco_ios_xr"
    ases.EntityData.ParentYangName = "af"
    ases.EntityData.SegmentPath = "ases"
    ases.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/" + ases.EntityData.SegmentPath
    ases.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ases.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ases.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ases.EntityData.Children = types.NewOrderedMap()
    ases.EntityData.Children.Append("as", types.YChild{"As", nil})
    for i := range ases.As {
        ases.EntityData.Children.Append(types.GetSegmentPath(ases.As[i]), types.YChild{"As", ases.As[i]})
    }
    ases.EntityData.Leafs = types.NewOrderedMap()

    ases.EntityData.YListKeys = []string {}

    return &(ases.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As
// Operational data for one AS
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. AS number. The type is interface{} with range:
    // 1..65535.
    Asn interface{}

    // EIGRP static neighbors.
    StaticNeighbors Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_StaticNeighbors

    // EIGRP topology table.
    TopologyRoutes Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes

    // Accounting info for one VRF AF.
    EigrpAccounting Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpAccounting

    // Traffic info for one VRF AF.
    EigrpTraffic Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpTraffic

    // Topology summary info for one VRF AF.
    EigrpTopologySummary Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpTopologySummary

    // EIGRP interfaces.
    Interfaces Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces

    // Events for a specific VRF AF.
    EigrpEvents Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpEvents

    // EIGRP neighbors.
    Neighbors Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors

    // Statistics for a specific VRF AF.
    EigrpStatistics Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpStatistics
}

func (as *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As) GetEntityData() *types.CommonEntityData {
    as.EntityData.YFilter = as.YFilter
    as.EntityData.YangName = "as"
    as.EntityData.BundleName = "cisco_ios_xr"
    as.EntityData.ParentYangName = "ases"
    as.EntityData.SegmentPath = "as" + types.AddKeyToken(as.Asn, "asn")
    as.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/" + as.EntityData.SegmentPath
    as.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    as.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    as.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    as.EntityData.Children = types.NewOrderedMap()
    as.EntityData.Children.Append("static-neighbors", types.YChild{"StaticNeighbors", &as.StaticNeighbors})
    as.EntityData.Children.Append("topology-routes", types.YChild{"TopologyRoutes", &as.TopologyRoutes})
    as.EntityData.Children.Append("eigrp-accounting", types.YChild{"EigrpAccounting", &as.EigrpAccounting})
    as.EntityData.Children.Append("eigrp-traffic", types.YChild{"EigrpTraffic", &as.EigrpTraffic})
    as.EntityData.Children.Append("eigrp-topology-summary", types.YChild{"EigrpTopologySummary", &as.EigrpTopologySummary})
    as.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &as.Interfaces})
    as.EntityData.Children.Append("eigrp-events", types.YChild{"EigrpEvents", &as.EigrpEvents})
    as.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &as.Neighbors})
    as.EntityData.Children.Append("eigrp-statistics", types.YChild{"EigrpStatistics", &as.EigrpStatistics})
    as.EntityData.Leafs = types.NewOrderedMap()
    as.EntityData.Leafs.Append("asn", types.YLeaf{"Asn", as.Asn})

    as.EntityData.YListKeys = []string {"Asn"}

    return &(as.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_StaticNeighbors
// EIGRP static neighbors
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_StaticNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information on one static EIGRP neighbor. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_StaticNeighbors_StaticNeighbor.
    StaticNeighbor []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_StaticNeighbors_StaticNeighbor
}

func (staticNeighbors *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_StaticNeighbors) GetEntityData() *types.CommonEntityData {
    staticNeighbors.EntityData.YFilter = staticNeighbors.YFilter
    staticNeighbors.EntityData.YangName = "static-neighbors"
    staticNeighbors.EntityData.BundleName = "cisco_ios_xr"
    staticNeighbors.EntityData.ParentYangName = "as"
    staticNeighbors.EntityData.SegmentPath = "static-neighbors"
    staticNeighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/" + staticNeighbors.EntityData.SegmentPath
    staticNeighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticNeighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticNeighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticNeighbors.EntityData.Children = types.NewOrderedMap()
    staticNeighbors.EntityData.Children.Append("static-neighbor", types.YChild{"StaticNeighbor", nil})
    for i := range staticNeighbors.StaticNeighbor {
        staticNeighbors.EntityData.Children.Append(types.GetSegmentPath(staticNeighbors.StaticNeighbor[i]), types.YChild{"StaticNeighbor", staticNeighbors.StaticNeighbor[i]})
    }
    staticNeighbors.EntityData.Leafs = types.NewOrderedMap()

    staticNeighbors.EntityData.YListKeys = []string {}

    return &(staticNeighbors.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_StaticNeighbors_StaticNeighbor
// Information on one static EIGRP
// neighbor
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_StaticNeighbors_StaticNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor Address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    NeighborAddress interface{}

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // AS Number. The type is interface{} with range: 0..4294967295.
    AsNumber interface{}

    // Interface Name. The type is string.
    InterfaceList interface{}

    // Neighbor address.
    Source Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_StaticNeighbors_StaticNeighbor_Source
}

func (staticNeighbor *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_StaticNeighbors_StaticNeighbor) GetEntityData() *types.CommonEntityData {
    staticNeighbor.EntityData.YFilter = staticNeighbor.YFilter
    staticNeighbor.EntityData.YangName = "static-neighbor"
    staticNeighbor.EntityData.BundleName = "cisco_ios_xr"
    staticNeighbor.EntityData.ParentYangName = "static-neighbors"
    staticNeighbor.EntityData.SegmentPath = "static-neighbor" + types.AddKeyToken(staticNeighbor.NeighborAddress, "neighbor-address")
    staticNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/static-neighbors/" + staticNeighbor.EntityData.SegmentPath
    staticNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticNeighbor.EntityData.Children = types.NewOrderedMap()
    staticNeighbor.EntityData.Children.Append("source", types.YChild{"Source", &staticNeighbor.Source})
    staticNeighbor.EntityData.Leafs = types.NewOrderedMap()
    staticNeighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", staticNeighbor.NeighborAddress})
    staticNeighbor.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", staticNeighbor.Afi})
    staticNeighbor.EntityData.Leafs.Append("as-number", types.YLeaf{"AsNumber", staticNeighbor.AsNumber})
    staticNeighbor.EntityData.Leafs.Append("interface-list", types.YLeaf{"InterfaceList", staticNeighbor.InterfaceList})

    staticNeighbor.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(staticNeighbor.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_StaticNeighbors_StaticNeighbor_Source
// Neighbor address
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_StaticNeighbors_StaticNeighbor_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (source *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_StaticNeighbors_StaticNeighbor_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "static-neighbor"
    source.EntityData.SegmentPath = "source"
    source.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/static-neighbors/static-neighbor/" + source.EntityData.SegmentPath
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = types.NewOrderedMap()
    source.EntityData.Leafs = types.NewOrderedMap()
    source.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", source.Ipv4Address})
    source.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", source.Ipv6Address})

    source.EntityData.YListKeys = []string {}

    return &(source.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes
// EIGRP topology table
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about one EIGRP route. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute.
    TopologyRoute []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute
}

func (topologyRoutes *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes) GetEntityData() *types.CommonEntityData {
    topologyRoutes.EntityData.YFilter = topologyRoutes.YFilter
    topologyRoutes.EntityData.YangName = "topology-routes"
    topologyRoutes.EntityData.BundleName = "cisco_ios_xr"
    topologyRoutes.EntityData.ParentYangName = "as"
    topologyRoutes.EntityData.SegmentPath = "topology-routes"
    topologyRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/" + topologyRoutes.EntityData.SegmentPath
    topologyRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyRoutes.EntityData.Children = types.NewOrderedMap()
    topologyRoutes.EntityData.Children.Append("topology-route", types.YChild{"TopologyRoute", nil})
    for i := range topologyRoutes.TopologyRoute {
        types.SetYListKey(topologyRoutes.TopologyRoute[i], i)
        topologyRoutes.EntityData.Children.Append(types.GetSegmentPath(topologyRoutes.TopologyRoute[i]), types.YChild{"TopologyRoute", topologyRoutes.TopologyRoute[i]})
    }
    topologyRoutes.EntityData.Leafs = types.NewOrderedMap()

    topologyRoutes.EntityData.YListKeys = []string {}

    return &(topologyRoutes.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute
// Information about one EIGRP route
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // IP Prefix. The type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Prefix interface{}

    // IP Prefix length. The type is interface{} with range: 0..128.
    PrefixLength interface{}

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // AS Number. The type is interface{} with range: 0..4294967295.
    AsNumber interface{}

    // Router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Active route. The type is bool.
    Active interface{}

    // Successors. The type is interface{} with range: 0..4294967295.
    Successors interface{}

    // Are there successors. The type is bool.
    SuccessorsPresent interface{}

    // Deprecated. Please migrate to use OldMetric64. The value of this object
    // might wrap since the metric value can go up to (2^48 - 1) in 64-bit metric
    // mode. The type is interface{} with range: 0..4294967295.
    OldMetric interface{}

    // Old metric. The type is interface{} with range: 0..18446744073709551615.
    OldMetric64 interface{}

    // Deprecated. Please migrate to use Metric64. The value of this object might
    // wrap since the metric value can go up to (2^48 - 1) in 64-bit metric mode.
    // The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Metric. The type is interface{} with range: 0..18446744073709551615.
    Metric64 interface{}

    // Metric downloaded to RIB. The type is interface{} with range:
    // 0..4294967295.
    RibMetric interface{}

    // Tag. The type is interface{} with range: 0..4294967295.
    Tag interface{}

    // Send flag. The type is EigrpBdPathSendFlag.
    SendFlag interface{}

    // Transmit thread Serial Number. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmitSerialNumber interface{}

    // Transmit thread refcount. The type is interface{} with range:
    // 0..4294967295.
    TransmitRefcount interface{}

    // Is Transmit thread anchored. The type is bool.
    TransmitAnchored interface{}

    // Reply handles used. The type is interface{} with range: 0..4294967295.
    ReplyHandles interface{}

    // Active time seconds. The type is interface{} with range: 0..4294967295.
    // Units are second.
    ActiveTimeSecs interface{}

    // Active time nanoseconds. The type is interface{} with range: 0..4294967295.
    // Units are nanosecond.
    ActiveTimeNsecs interface{}

    // Origin. The type is interface{} with range: 0..4294967295.
    Origin interface{}

    // Retry count. The type is interface{} with range: 0..4294967295.
    RetryCount interface{}

    // Active stats flag. The type is bool.
    ActiveStats interface{}

    // Active stats min time. The type is interface{} with range: 0..4294967295.
    MinTime interface{}

    // Active stats max time. The type is interface{} with range: 0..4294967295.
    MaxTime interface{}

    // Active stats average time. The type is interface{} with range:
    // 0..4294967295.
    AverageTime interface{}

    // Active stats active count. The type is interface{} with range:
    // 0..4294967295.
    AckCount interface{}

    // Number of replies outstanding. The type is interface{} with range:
    // 0..4294967295.
    Replies interface{}

    // Route is SIA. The type is bool.
    RouteInSia interface{}

    // Reply handles used. The type is interface{} with range: 0..4294967295.
    SiaReplyHandles interface{}

    // IP Prefix/length.
    PrefixXr Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_PrefixXr

    // Paths for this route. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_Paths.
    Paths []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_Paths

    // Peers yet to respond. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_ActivePeer.
    ActivePeer []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_ActivePeer

    // SIA Peers yet to respond. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_SiaPeer.
    SiaPeer []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_SiaPeer
}

func (topologyRoute *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute) GetEntityData() *types.CommonEntityData {
    topologyRoute.EntityData.YFilter = topologyRoute.YFilter
    topologyRoute.EntityData.YangName = "topology-route"
    topologyRoute.EntityData.BundleName = "cisco_ios_xr"
    topologyRoute.EntityData.ParentYangName = "topology-routes"
    topologyRoute.EntityData.SegmentPath = "topology-route" + types.AddNoKeyToken(topologyRoute)
    topologyRoute.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/topology-routes/" + topologyRoute.EntityData.SegmentPath
    topologyRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyRoute.EntityData.Children = types.NewOrderedMap()
    topologyRoute.EntityData.Children.Append("prefix-xr", types.YChild{"PrefixXr", &topologyRoute.PrefixXr})
    topologyRoute.EntityData.Children.Append("paths", types.YChild{"Paths", nil})
    for i := range topologyRoute.Paths {
        types.SetYListKey(topologyRoute.Paths[i], i)
        topologyRoute.EntityData.Children.Append(types.GetSegmentPath(topologyRoute.Paths[i]), types.YChild{"Paths", topologyRoute.Paths[i]})
    }
    topologyRoute.EntityData.Children.Append("active-peer", types.YChild{"ActivePeer", nil})
    for i := range topologyRoute.ActivePeer {
        types.SetYListKey(topologyRoute.ActivePeer[i], i)
        topologyRoute.EntityData.Children.Append(types.GetSegmentPath(topologyRoute.ActivePeer[i]), types.YChild{"ActivePeer", topologyRoute.ActivePeer[i]})
    }
    topologyRoute.EntityData.Children.Append("sia-peer", types.YChild{"SiaPeer", nil})
    for i := range topologyRoute.SiaPeer {
        types.SetYListKey(topologyRoute.SiaPeer[i], i)
        topologyRoute.EntityData.Children.Append(types.GetSegmentPath(topologyRoute.SiaPeer[i]), types.YChild{"SiaPeer", topologyRoute.SiaPeer[i]})
    }
    topologyRoute.EntityData.Leafs = types.NewOrderedMap()
    topologyRoute.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", topologyRoute.Prefix})
    topologyRoute.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", topologyRoute.PrefixLength})
    topologyRoute.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", topologyRoute.Afi})
    topologyRoute.EntityData.Leafs.Append("as-number", types.YLeaf{"AsNumber", topologyRoute.AsNumber})
    topologyRoute.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", topologyRoute.RouterId})
    topologyRoute.EntityData.Leafs.Append("active", types.YLeaf{"Active", topologyRoute.Active})
    topologyRoute.EntityData.Leafs.Append("successors", types.YLeaf{"Successors", topologyRoute.Successors})
    topologyRoute.EntityData.Leafs.Append("successors-present", types.YLeaf{"SuccessorsPresent", topologyRoute.SuccessorsPresent})
    topologyRoute.EntityData.Leafs.Append("old-metric", types.YLeaf{"OldMetric", topologyRoute.OldMetric})
    topologyRoute.EntityData.Leafs.Append("old-metric64", types.YLeaf{"OldMetric64", topologyRoute.OldMetric64})
    topologyRoute.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", topologyRoute.Metric})
    topologyRoute.EntityData.Leafs.Append("metric64", types.YLeaf{"Metric64", topologyRoute.Metric64})
    topologyRoute.EntityData.Leafs.Append("rib-metric", types.YLeaf{"RibMetric", topologyRoute.RibMetric})
    topologyRoute.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", topologyRoute.Tag})
    topologyRoute.EntityData.Leafs.Append("send-flag", types.YLeaf{"SendFlag", topologyRoute.SendFlag})
    topologyRoute.EntityData.Leafs.Append("transmit-serial-number", types.YLeaf{"TransmitSerialNumber", topologyRoute.TransmitSerialNumber})
    topologyRoute.EntityData.Leafs.Append("transmit-refcount", types.YLeaf{"TransmitRefcount", topologyRoute.TransmitRefcount})
    topologyRoute.EntityData.Leafs.Append("transmit-anchored", types.YLeaf{"TransmitAnchored", topologyRoute.TransmitAnchored})
    topologyRoute.EntityData.Leafs.Append("reply-handles", types.YLeaf{"ReplyHandles", topologyRoute.ReplyHandles})
    topologyRoute.EntityData.Leafs.Append("active-time-secs", types.YLeaf{"ActiveTimeSecs", topologyRoute.ActiveTimeSecs})
    topologyRoute.EntityData.Leafs.Append("active-time-nsecs", types.YLeaf{"ActiveTimeNsecs", topologyRoute.ActiveTimeNsecs})
    topologyRoute.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", topologyRoute.Origin})
    topologyRoute.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", topologyRoute.RetryCount})
    topologyRoute.EntityData.Leafs.Append("active-stats", types.YLeaf{"ActiveStats", topologyRoute.ActiveStats})
    topologyRoute.EntityData.Leafs.Append("min-time", types.YLeaf{"MinTime", topologyRoute.MinTime})
    topologyRoute.EntityData.Leafs.Append("max-time", types.YLeaf{"MaxTime", topologyRoute.MaxTime})
    topologyRoute.EntityData.Leafs.Append("average-time", types.YLeaf{"AverageTime", topologyRoute.AverageTime})
    topologyRoute.EntityData.Leafs.Append("ack-count", types.YLeaf{"AckCount", topologyRoute.AckCount})
    topologyRoute.EntityData.Leafs.Append("replies", types.YLeaf{"Replies", topologyRoute.Replies})
    topologyRoute.EntityData.Leafs.Append("route-in-sia", types.YLeaf{"RouteInSia", topologyRoute.RouteInSia})
    topologyRoute.EntityData.Leafs.Append("sia-reply-handles", types.YLeaf{"SiaReplyHandles", topologyRoute.SiaReplyHandles})

    topologyRoute.EntityData.YListKeys = []string {}

    return &(topologyRoute.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_PrefixXr
// IP Prefix/length
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_PrefixXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Prefix. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Prefix interface{}

    // IPv6 Prefix. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}
}

func (prefixXr *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_PrefixXr) GetEntityData() *types.CommonEntityData {
    prefixXr.EntityData.YFilter = prefixXr.YFilter
    prefixXr.EntityData.YangName = "prefix-xr"
    prefixXr.EntityData.BundleName = "cisco_ios_xr"
    prefixXr.EntityData.ParentYangName = "topology-route"
    prefixXr.EntityData.SegmentPath = "prefix-xr"
    prefixXr.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/topology-routes/topology-route/" + prefixXr.EntityData.SegmentPath
    prefixXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixXr.EntityData.Children = types.NewOrderedMap()
    prefixXr.EntityData.Leafs = types.NewOrderedMap()
    prefixXr.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", prefixXr.Ipv4Prefix})
    prefixXr.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", prefixXr.Ipv6Prefix})
    prefixXr.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefixXr.PrefixLength})

    prefixXr.EntityData.YListKeys = []string {}

    return &(prefixXr.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_Paths
// Paths for this route
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // NH flag. The type is bool.
    NextHopPresent interface{}

    // Interface handle. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceHandle interface{}

    // Interface name. The type is string.
    InterfaceName interface{}

    // Origin of route. The type is EigrpBdPathOrigin.
    Origin interface{}

    // Send flag. The type is EigrpBdPathSendFlag.
    SendFlag interface{}

    // Outstanding reply. The type is bool.
    ReplyOutstanding interface{}

    // Deprecated. Please migrate to use Metric64. The value of this object might
    // wrap since the metric value can go up to (2^48 - 1) in 64-bit metric mode.
    // The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Metric. The type is interface{} with range: 0..18446744073709551615.
    Metric64 interface{}

    // Deprecated. Please migrate to use SuccessorMetric64. The value of this
    // object might wrap since the metric value can go up to (2^48 - 1) in 64-bit
    // metric mode. The type is interface{} with range: 0..4294967295.
    SuccessorMetric interface{}

    // Successor metric. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessorMetric64 interface{}

    // Reply status. The type is bool.
    ReplyStatus interface{}

    // SIA status. The type is bool.
    SiaStatus interface{}

    // Transmit thread serial number. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmitSerialNumber interface{}

    // Is Transmit thread anchored. The type is bool.
    Anchored interface{}

    // External. The type is bool.
    ExternalPath interface{}

    // Deprecated. Please migrate to use Bandwidth64. The type is interface{} with
    // range: 0..4294967295.
    Bandwidth interface{}

    // Bandwidth. The type is interface{} with range: 0..18446744073709551615.
    Bandwidth64 interface{}

    // Deprecated. Please migrate to use Delay64. The value of this object might
    // wrap if it is in picosecond units . The type is interface{} with range:
    // 0..4294967295.
    Delay interface{}

    // Delay. The type is interface{} with range: 0..18446744073709551615.
    Delay64 interface{}

    // Delay units. The type is EigrpBdDelayUnit.
    DelayUnit interface{}

    // MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Hopcount. The type is interface{} with range: 0..4294967295.
    HopCount interface{}

    // Reliability. The type is interface{} with range: 0..255.
    Reliability interface{}

    // Load. The type is interface{} with range: 0..255.
    Load interface{}

    // Router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    InternalRouterId interface{}

    // Internal Tag. The type is interface{} with range: 0..255.
    InternalTag interface{}

    // Extended communities present. The type is bool.
    ExtendedCommunitiesPresent interface{}

    // Length of extended communities. The type is interface{} with range:
    // 0..4294967295.
    ExtendedCommunitiesLength interface{}

    // Extended communities. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ExtendedCommunities interface{}

    // External information present. The type is bool.
    ExternalInformationPresent interface{}

    // Router ID. The type is interface{} with range: 0..4294967295.
    ExternalRouterId interface{}

    // Is it this system. The type is bool.
    ExternalThisSystem interface{}

    // AS Number. The type is interface{} with range: 0..4294967295.
    ExternalAs interface{}

    // Protocol ID. The type is string.
    ExternalProtocol interface{}

    // Metric. The type is interface{} with range: 0..4294967295.
    ExternalMetric interface{}

    // Tag. The type is interface{} with range: 0..4294967295.
    ExternalTag interface{}

    // Conditional Default flag. The type is bool.
    ConditionalDefaultPath interface{}

    // State of path in RIB. The type is EigrpBdPathRibState.
    RibState interface{}

    // Nexthop address.
    NextHopAddress Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_Paths_NextHopAddress

    // Source of route.
    Infosource Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_Paths_Infosource
}

func (paths *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "topology-route"
    paths.EntityData.SegmentPath = "paths" + types.AddNoKeyToken(paths)
    paths.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/topology-routes/topology-route/" + paths.EntityData.SegmentPath
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = types.NewOrderedMap()
    paths.EntityData.Children.Append("next-hop-address", types.YChild{"NextHopAddress", &paths.NextHopAddress})
    paths.EntityData.Children.Append("infosource", types.YChild{"Infosource", &paths.Infosource})
    paths.EntityData.Leafs = types.NewOrderedMap()
    paths.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", paths.Afi})
    paths.EntityData.Leafs.Append("next-hop-present", types.YLeaf{"NextHopPresent", paths.NextHopPresent})
    paths.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", paths.InterfaceHandle})
    paths.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", paths.InterfaceName})
    paths.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", paths.Origin})
    paths.EntityData.Leafs.Append("send-flag", types.YLeaf{"SendFlag", paths.SendFlag})
    paths.EntityData.Leafs.Append("reply-outstanding", types.YLeaf{"ReplyOutstanding", paths.ReplyOutstanding})
    paths.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", paths.Metric})
    paths.EntityData.Leafs.Append("metric64", types.YLeaf{"Metric64", paths.Metric64})
    paths.EntityData.Leafs.Append("successor-metric", types.YLeaf{"SuccessorMetric", paths.SuccessorMetric})
    paths.EntityData.Leafs.Append("successor-metric64", types.YLeaf{"SuccessorMetric64", paths.SuccessorMetric64})
    paths.EntityData.Leafs.Append("reply-status", types.YLeaf{"ReplyStatus", paths.ReplyStatus})
    paths.EntityData.Leafs.Append("sia-status", types.YLeaf{"SiaStatus", paths.SiaStatus})
    paths.EntityData.Leafs.Append("transmit-serial-number", types.YLeaf{"TransmitSerialNumber", paths.TransmitSerialNumber})
    paths.EntityData.Leafs.Append("anchored", types.YLeaf{"Anchored", paths.Anchored})
    paths.EntityData.Leafs.Append("external-path", types.YLeaf{"ExternalPath", paths.ExternalPath})
    paths.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", paths.Bandwidth})
    paths.EntityData.Leafs.Append("bandwidth64", types.YLeaf{"Bandwidth64", paths.Bandwidth64})
    paths.EntityData.Leafs.Append("delay", types.YLeaf{"Delay", paths.Delay})
    paths.EntityData.Leafs.Append("delay64", types.YLeaf{"Delay64", paths.Delay64})
    paths.EntityData.Leafs.Append("delay-unit", types.YLeaf{"DelayUnit", paths.DelayUnit})
    paths.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", paths.Mtu})
    paths.EntityData.Leafs.Append("hop-count", types.YLeaf{"HopCount", paths.HopCount})
    paths.EntityData.Leafs.Append("reliability", types.YLeaf{"Reliability", paths.Reliability})
    paths.EntityData.Leafs.Append("load", types.YLeaf{"Load", paths.Load})
    paths.EntityData.Leafs.Append("internal-router-id", types.YLeaf{"InternalRouterId", paths.InternalRouterId})
    paths.EntityData.Leafs.Append("internal-tag", types.YLeaf{"InternalTag", paths.InternalTag})
    paths.EntityData.Leafs.Append("extended-communities-present", types.YLeaf{"ExtendedCommunitiesPresent", paths.ExtendedCommunitiesPresent})
    paths.EntityData.Leafs.Append("extended-communities-length", types.YLeaf{"ExtendedCommunitiesLength", paths.ExtendedCommunitiesLength})
    paths.EntityData.Leafs.Append("extended-communities", types.YLeaf{"ExtendedCommunities", paths.ExtendedCommunities})
    paths.EntityData.Leafs.Append("external-information-present", types.YLeaf{"ExternalInformationPresent", paths.ExternalInformationPresent})
    paths.EntityData.Leafs.Append("external-router-id", types.YLeaf{"ExternalRouterId", paths.ExternalRouterId})
    paths.EntityData.Leafs.Append("external-this-system", types.YLeaf{"ExternalThisSystem", paths.ExternalThisSystem})
    paths.EntityData.Leafs.Append("external-as", types.YLeaf{"ExternalAs", paths.ExternalAs})
    paths.EntityData.Leafs.Append("external-protocol", types.YLeaf{"ExternalProtocol", paths.ExternalProtocol})
    paths.EntityData.Leafs.Append("external-metric", types.YLeaf{"ExternalMetric", paths.ExternalMetric})
    paths.EntityData.Leafs.Append("external-tag", types.YLeaf{"ExternalTag", paths.ExternalTag})
    paths.EntityData.Leafs.Append("conditional-default-path", types.YLeaf{"ConditionalDefaultPath", paths.ConditionalDefaultPath})
    paths.EntityData.Leafs.Append("rib-state", types.YLeaf{"RibState", paths.RibState})

    paths.EntityData.YListKeys = []string {}

    return &(paths.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_Paths_NextHopAddress
// Nexthop address
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_Paths_NextHopAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (nextHopAddress *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_Paths_NextHopAddress) GetEntityData() *types.CommonEntityData {
    nextHopAddress.EntityData.YFilter = nextHopAddress.YFilter
    nextHopAddress.EntityData.YangName = "next-hop-address"
    nextHopAddress.EntityData.BundleName = "cisco_ios_xr"
    nextHopAddress.EntityData.ParentYangName = "paths"
    nextHopAddress.EntityData.SegmentPath = "next-hop-address"
    nextHopAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/topology-routes/topology-route/paths/" + nextHopAddress.EntityData.SegmentPath
    nextHopAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHopAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHopAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHopAddress.EntityData.Children = types.NewOrderedMap()
    nextHopAddress.EntityData.Leafs = types.NewOrderedMap()
    nextHopAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHopAddress.Ipv4Address})
    nextHopAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHopAddress.Ipv6Address})

    nextHopAddress.EntityData.YListKeys = []string {}

    return &(nextHopAddress.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_Paths_Infosource
// Source of route
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_Paths_Infosource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (infosource *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_Paths_Infosource) GetEntityData() *types.CommonEntityData {
    infosource.EntityData.YFilter = infosource.YFilter
    infosource.EntityData.YangName = "infosource"
    infosource.EntityData.BundleName = "cisco_ios_xr"
    infosource.EntityData.ParentYangName = "paths"
    infosource.EntityData.SegmentPath = "infosource"
    infosource.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/topology-routes/topology-route/paths/" + infosource.EntityData.SegmentPath
    infosource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infosource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infosource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infosource.EntityData.Children = types.NewOrderedMap()
    infosource.EntityData.Leafs = types.NewOrderedMap()
    infosource.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", infosource.Ipv4Address})
    infosource.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", infosource.Ipv6Address})

    infosource.EntityData.YListKeys = []string {}

    return &(infosource.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_ActivePeer
// Peers yet to respond
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_ActivePeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Peer available. The type is bool.
    PeerAvailable interface{}

    // Interface name. The type is string.
    InterfaceList interface{}

    // Handle number. The type is interface{} with range: 0..4294967295.
    HandleNumber interface{}

    // Peer Address.
    Source Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_ActivePeer_Source
}

func (activePeer *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_ActivePeer) GetEntityData() *types.CommonEntityData {
    activePeer.EntityData.YFilter = activePeer.YFilter
    activePeer.EntityData.YangName = "active-peer"
    activePeer.EntityData.BundleName = "cisco_ios_xr"
    activePeer.EntityData.ParentYangName = "topology-route"
    activePeer.EntityData.SegmentPath = "active-peer" + types.AddNoKeyToken(activePeer)
    activePeer.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/topology-routes/topology-route/" + activePeer.EntityData.SegmentPath
    activePeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activePeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activePeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activePeer.EntityData.Children = types.NewOrderedMap()
    activePeer.EntityData.Children.Append("source", types.YChild{"Source", &activePeer.Source})
    activePeer.EntityData.Leafs = types.NewOrderedMap()
    activePeer.EntityData.Leafs.Append("peer-available", types.YLeaf{"PeerAvailable", activePeer.PeerAvailable})
    activePeer.EntityData.Leafs.Append("interface-list", types.YLeaf{"InterfaceList", activePeer.InterfaceList})
    activePeer.EntityData.Leafs.Append("handle-number", types.YLeaf{"HandleNumber", activePeer.HandleNumber})

    activePeer.EntityData.YListKeys = []string {}

    return &(activePeer.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_ActivePeer_Source
// Peer Address
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_ActivePeer_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (source *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_ActivePeer_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "active-peer"
    source.EntityData.SegmentPath = "source"
    source.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/topology-routes/topology-route/active-peer/" + source.EntityData.SegmentPath
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = types.NewOrderedMap()
    source.EntityData.Leafs = types.NewOrderedMap()
    source.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", source.Ipv4Address})
    source.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", source.Ipv6Address})

    source.EntityData.YListKeys = []string {}

    return &(source.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_SiaPeer
// SIA Peers yet to respond
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_SiaPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Peer available. The type is bool.
    PeerAvailable interface{}

    // Interface name. The type is string.
    InterfaceList interface{}

    // Handle number. The type is interface{} with range: 0..4294967295.
    HandleNumber interface{}

    // Peer Address.
    Source Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_SiaPeer_Source
}

func (siaPeer *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_SiaPeer) GetEntityData() *types.CommonEntityData {
    siaPeer.EntityData.YFilter = siaPeer.YFilter
    siaPeer.EntityData.YangName = "sia-peer"
    siaPeer.EntityData.BundleName = "cisco_ios_xr"
    siaPeer.EntityData.ParentYangName = "topology-route"
    siaPeer.EntityData.SegmentPath = "sia-peer" + types.AddNoKeyToken(siaPeer)
    siaPeer.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/topology-routes/topology-route/" + siaPeer.EntityData.SegmentPath
    siaPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    siaPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    siaPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    siaPeer.EntityData.Children = types.NewOrderedMap()
    siaPeer.EntityData.Children.Append("source", types.YChild{"Source", &siaPeer.Source})
    siaPeer.EntityData.Leafs = types.NewOrderedMap()
    siaPeer.EntityData.Leafs.Append("peer-available", types.YLeaf{"PeerAvailable", siaPeer.PeerAvailable})
    siaPeer.EntityData.Leafs.Append("interface-list", types.YLeaf{"InterfaceList", siaPeer.InterfaceList})
    siaPeer.EntityData.Leafs.Append("handle-number", types.YLeaf{"HandleNumber", siaPeer.HandleNumber})

    siaPeer.EntityData.YListKeys = []string {}

    return &(siaPeer.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_SiaPeer_Source
// Peer Address
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_SiaPeer_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (source *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_TopologyRoutes_TopologyRoute_SiaPeer_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "sia-peer"
    source.EntityData.SegmentPath = "source"
    source.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/topology-routes/topology-route/sia-peer/" + source.EntityData.SegmentPath
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = types.NewOrderedMap()
    source.EntityData.Leafs = types.NewOrderedMap()
    source.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", source.Ipv4Address})
    source.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", source.Ipv6Address})

    source.EntityData.YListKeys = []string {}

    return &(source.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpAccounting
// Accounting info for one VRF AF
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // AS Number. The type is interface{} with range: 0..4294967295.
    AsNumber interface{}

    // Router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Number of aggregates. The type is interface{} with range: 0..4294967295.
    AggregateCount interface{}

    // Redist state. The type is interface{} with range: -128..127.
    State interface{}

    // Redist prefix count. The type is interface{} with range: 0..4294967295.
    RedistPrefixCount interface{}

    // Restart count. The type is interface{} with range: 0..4294967295.
    RestartCount interface{}

    // Time left. The type is interface{} with range: 0..4294967295.
    TimeLeft interface{}

    // Are there redist'ed prefixes ?. The type is bool.
    RedistPrefixPresent interface{}

    // Peers and their status. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpAccounting_PeerStatistics.
    PeerStatistics []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpAccounting_PeerStatistics
}

func (eigrpAccounting *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpAccounting) GetEntityData() *types.CommonEntityData {
    eigrpAccounting.EntityData.YFilter = eigrpAccounting.YFilter
    eigrpAccounting.EntityData.YangName = "eigrp-accounting"
    eigrpAccounting.EntityData.BundleName = "cisco_ios_xr"
    eigrpAccounting.EntityData.ParentYangName = "as"
    eigrpAccounting.EntityData.SegmentPath = "eigrp-accounting"
    eigrpAccounting.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/" + eigrpAccounting.EntityData.SegmentPath
    eigrpAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrpAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrpAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrpAccounting.EntityData.Children = types.NewOrderedMap()
    eigrpAccounting.EntityData.Children.Append("peer-statistics", types.YChild{"PeerStatistics", nil})
    for i := range eigrpAccounting.PeerStatistics {
        types.SetYListKey(eigrpAccounting.PeerStatistics[i], i)
        eigrpAccounting.EntityData.Children.Append(types.GetSegmentPath(eigrpAccounting.PeerStatistics[i]), types.YChild{"PeerStatistics", eigrpAccounting.PeerStatistics[i]})
    }
    eigrpAccounting.EntityData.Leafs = types.NewOrderedMap()
    eigrpAccounting.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", eigrpAccounting.Afi})
    eigrpAccounting.EntityData.Leafs.Append("as-number", types.YLeaf{"AsNumber", eigrpAccounting.AsNumber})
    eigrpAccounting.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", eigrpAccounting.RouterId})
    eigrpAccounting.EntityData.Leafs.Append("aggregate-count", types.YLeaf{"AggregateCount", eigrpAccounting.AggregateCount})
    eigrpAccounting.EntityData.Leafs.Append("state", types.YLeaf{"State", eigrpAccounting.State})
    eigrpAccounting.EntityData.Leafs.Append("redist-prefix-count", types.YLeaf{"RedistPrefixCount", eigrpAccounting.RedistPrefixCount})
    eigrpAccounting.EntityData.Leafs.Append("restart-count", types.YLeaf{"RestartCount", eigrpAccounting.RestartCount})
    eigrpAccounting.EntityData.Leafs.Append("time-left", types.YLeaf{"TimeLeft", eigrpAccounting.TimeLeft})
    eigrpAccounting.EntityData.Leafs.Append("redist-prefix-present", types.YLeaf{"RedistPrefixPresent", eigrpAccounting.RedistPrefixPresent})

    eigrpAccounting.EntityData.YListKeys = []string {}

    return &(eigrpAccounting.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpAccounting_PeerStatistics
// Peers and their status
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpAccounting_PeerStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // State. The type is interface{} with range: -128..127.
    State interface{}

    // Interface name. The type is string.
    InterfaceList interface{}

    // Peer prefix count. The type is interface{} with range: 0..4294967295.
    PeerPrefixCount interface{}

    // Restart count. The type is interface{} with range: 0..4294967295.
    RestartCount interface{}

    // Time left. The type is interface{} with range: 0..4294967295.
    TimeLeft interface{}

    // Source address.
    Source Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpAccounting_PeerStatistics_Source
}

func (peerStatistics *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpAccounting_PeerStatistics) GetEntityData() *types.CommonEntityData {
    peerStatistics.EntityData.YFilter = peerStatistics.YFilter
    peerStatistics.EntityData.YangName = "peer-statistics"
    peerStatistics.EntityData.BundleName = "cisco_ios_xr"
    peerStatistics.EntityData.ParentYangName = "eigrp-accounting"
    peerStatistics.EntityData.SegmentPath = "peer-statistics" + types.AddNoKeyToken(peerStatistics)
    peerStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/eigrp-accounting/" + peerStatistics.EntityData.SegmentPath
    peerStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerStatistics.EntityData.Children = types.NewOrderedMap()
    peerStatistics.EntityData.Children.Append("source", types.YChild{"Source", &peerStatistics.Source})
    peerStatistics.EntityData.Leafs = types.NewOrderedMap()
    peerStatistics.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", peerStatistics.Afi})
    peerStatistics.EntityData.Leafs.Append("state", types.YLeaf{"State", peerStatistics.State})
    peerStatistics.EntityData.Leafs.Append("interface-list", types.YLeaf{"InterfaceList", peerStatistics.InterfaceList})
    peerStatistics.EntityData.Leafs.Append("peer-prefix-count", types.YLeaf{"PeerPrefixCount", peerStatistics.PeerPrefixCount})
    peerStatistics.EntityData.Leafs.Append("restart-count", types.YLeaf{"RestartCount", peerStatistics.RestartCount})
    peerStatistics.EntityData.Leafs.Append("time-left", types.YLeaf{"TimeLeft", peerStatistics.TimeLeft})

    peerStatistics.EntityData.YListKeys = []string {}

    return &(peerStatistics.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpAccounting_PeerStatistics_Source
// Source address
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpAccounting_PeerStatistics_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (source *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpAccounting_PeerStatistics_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "peer-statistics"
    source.EntityData.SegmentPath = "source"
    source.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/eigrp-accounting/peer-statistics/" + source.EntityData.SegmentPath
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = types.NewOrderedMap()
    source.EntityData.Leafs = types.NewOrderedMap()
    source.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", source.Ipv4Address})
    source.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", source.Ipv6Address})

    source.EntityData.YListKeys = []string {}

    return &(source.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpTraffic
// Traffic info for one VRF AF
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpTraffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // AS Number. The type is interface{} with range: 0..4294967295.
    AsNumber interface{}

    // Hellos sent. The type is interface{} with range: 0..4294967295.
    HellosSent interface{}

    // Hellos received. The type is interface{} with range: 0..4294967295.
    HellosReceived interface{}

    // Updates sent. The type is interface{} with range: 0..4294967295.
    UpdatesSent interface{}

    // Updates received. The type is interface{} with range: 0..4294967295.
    UpdatesReceived interface{}

    // Queries sent. The type is interface{} with range: 0..4294967295.
    QueriesSent interface{}

    // Queries received. The type is interface{} with range: 0..4294967295.
    QueriesReceived interface{}

    // Replies sent. The type is interface{} with range: 0..4294967295.
    RepliesSent interface{}

    // Replies received. The type is interface{} with range: 0..4294967295.
    RepliesReceived interface{}

    // Acks sent. The type is interface{} with range: 0..4294967295.
    AcksSent interface{}

    // Acks received. The type is interface{} with range: 0..4294967295.
    AcksReceived interface{}

    // SIA Queries sent. The type is interface{} with range: 0..4294967295.
    SiaQueriesSent interface{}

    // SIA Queries received. The type is interface{} with range: 0..4294967295.
    SiaQueriesReceived interface{}

    // SIA Replies sent. The type is interface{} with range: 0..4294967295.
    SiaRepliesSent interface{}

    // SIA Replies received. The type is interface{} with range: 0..4294967295.
    SiaRepliesReceived interface{}

    // Maximum queue depth. The type is interface{} with range: 0..4294967295.
    MaxQueueDepth interface{}

    // Queue drops. The type is interface{} with range: 0..4294967295.
    QueueDrops interface{}
}

func (eigrpTraffic *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpTraffic) GetEntityData() *types.CommonEntityData {
    eigrpTraffic.EntityData.YFilter = eigrpTraffic.YFilter
    eigrpTraffic.EntityData.YangName = "eigrp-traffic"
    eigrpTraffic.EntityData.BundleName = "cisco_ios_xr"
    eigrpTraffic.EntityData.ParentYangName = "as"
    eigrpTraffic.EntityData.SegmentPath = "eigrp-traffic"
    eigrpTraffic.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/" + eigrpTraffic.EntityData.SegmentPath
    eigrpTraffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrpTraffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrpTraffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrpTraffic.EntityData.Children = types.NewOrderedMap()
    eigrpTraffic.EntityData.Leafs = types.NewOrderedMap()
    eigrpTraffic.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", eigrpTraffic.Afi})
    eigrpTraffic.EntityData.Leafs.Append("as-number", types.YLeaf{"AsNumber", eigrpTraffic.AsNumber})
    eigrpTraffic.EntityData.Leafs.Append("hellos-sent", types.YLeaf{"HellosSent", eigrpTraffic.HellosSent})
    eigrpTraffic.EntityData.Leafs.Append("hellos-received", types.YLeaf{"HellosReceived", eigrpTraffic.HellosReceived})
    eigrpTraffic.EntityData.Leafs.Append("updates-sent", types.YLeaf{"UpdatesSent", eigrpTraffic.UpdatesSent})
    eigrpTraffic.EntityData.Leafs.Append("updates-received", types.YLeaf{"UpdatesReceived", eigrpTraffic.UpdatesReceived})
    eigrpTraffic.EntityData.Leafs.Append("queries-sent", types.YLeaf{"QueriesSent", eigrpTraffic.QueriesSent})
    eigrpTraffic.EntityData.Leafs.Append("queries-received", types.YLeaf{"QueriesReceived", eigrpTraffic.QueriesReceived})
    eigrpTraffic.EntityData.Leafs.Append("replies-sent", types.YLeaf{"RepliesSent", eigrpTraffic.RepliesSent})
    eigrpTraffic.EntityData.Leafs.Append("replies-received", types.YLeaf{"RepliesReceived", eigrpTraffic.RepliesReceived})
    eigrpTraffic.EntityData.Leafs.Append("acks-sent", types.YLeaf{"AcksSent", eigrpTraffic.AcksSent})
    eigrpTraffic.EntityData.Leafs.Append("acks-received", types.YLeaf{"AcksReceived", eigrpTraffic.AcksReceived})
    eigrpTraffic.EntityData.Leafs.Append("sia-queries-sent", types.YLeaf{"SiaQueriesSent", eigrpTraffic.SiaQueriesSent})
    eigrpTraffic.EntityData.Leafs.Append("sia-queries-received", types.YLeaf{"SiaQueriesReceived", eigrpTraffic.SiaQueriesReceived})
    eigrpTraffic.EntityData.Leafs.Append("sia-replies-sent", types.YLeaf{"SiaRepliesSent", eigrpTraffic.SiaRepliesSent})
    eigrpTraffic.EntityData.Leafs.Append("sia-replies-received", types.YLeaf{"SiaRepliesReceived", eigrpTraffic.SiaRepliesReceived})
    eigrpTraffic.EntityData.Leafs.Append("max-queue-depth", types.YLeaf{"MaxQueueDepth", eigrpTraffic.MaxQueueDepth})
    eigrpTraffic.EntityData.Leafs.Append("queue-drops", types.YLeaf{"QueueDrops", eigrpTraffic.QueueDrops})

    eigrpTraffic.EntityData.YListKeys = []string {}

    return &(eigrpTraffic.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpTopologySummary
// Topology summary info for one VRF AF
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpTopologySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // AS Number. The type is interface{} with range: 0..4294967295.
    AsNumber interface{}

    // Router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Is thread present ?. The type is bool.
    ThreadPresent interface{}

    // Thread serial number. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmitSerialNumber interface{}

    // Next serial number. The type is interface{} with range:
    // 0..18446744073709551615.
    NextSerialNumber interface{}

    // Number of routes. The type is interface{} with range: 0..4294967295.
    RouteCount interface{}

    // Number of paths. The type is interface{} with range: 0..4294967295.
    PathCount interface{}

    // Dummy count. The type is interface{} with range: 0..4294967295.
    DummyCount interface{}

    // DDB Name. The type is string.
    DdbName interface{}

    // Number of interfaces. The type is interface{} with range: 0..4294967295.
    InterfaceCount interface{}

    // Number of handles used. The type is interface{} with range: 0..4294967295.
    HandlesUsed interface{}

    // Number of active interfaces. The type is interface{} with range:
    // 0..4294967295.
    ActiveInterfaceCount interface{}

    // Quiescent interfaces. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpTopologySummary_Quiescent.
    Quiescent []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpTopologySummary_Quiescent
}

func (eigrpTopologySummary *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpTopologySummary) GetEntityData() *types.CommonEntityData {
    eigrpTopologySummary.EntityData.YFilter = eigrpTopologySummary.YFilter
    eigrpTopologySummary.EntityData.YangName = "eigrp-topology-summary"
    eigrpTopologySummary.EntityData.BundleName = "cisco_ios_xr"
    eigrpTopologySummary.EntityData.ParentYangName = "as"
    eigrpTopologySummary.EntityData.SegmentPath = "eigrp-topology-summary"
    eigrpTopologySummary.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/" + eigrpTopologySummary.EntityData.SegmentPath
    eigrpTopologySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrpTopologySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrpTopologySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrpTopologySummary.EntityData.Children = types.NewOrderedMap()
    eigrpTopologySummary.EntityData.Children.Append("quiescent", types.YChild{"Quiescent", nil})
    for i := range eigrpTopologySummary.Quiescent {
        types.SetYListKey(eigrpTopologySummary.Quiescent[i], i)
        eigrpTopologySummary.EntityData.Children.Append(types.GetSegmentPath(eigrpTopologySummary.Quiescent[i]), types.YChild{"Quiescent", eigrpTopologySummary.Quiescent[i]})
    }
    eigrpTopologySummary.EntityData.Leafs = types.NewOrderedMap()
    eigrpTopologySummary.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", eigrpTopologySummary.Afi})
    eigrpTopologySummary.EntityData.Leafs.Append("as-number", types.YLeaf{"AsNumber", eigrpTopologySummary.AsNumber})
    eigrpTopologySummary.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", eigrpTopologySummary.RouterId})
    eigrpTopologySummary.EntityData.Leafs.Append("thread-present", types.YLeaf{"ThreadPresent", eigrpTopologySummary.ThreadPresent})
    eigrpTopologySummary.EntityData.Leafs.Append("transmit-serial-number", types.YLeaf{"TransmitSerialNumber", eigrpTopologySummary.TransmitSerialNumber})
    eigrpTopologySummary.EntityData.Leafs.Append("next-serial-number", types.YLeaf{"NextSerialNumber", eigrpTopologySummary.NextSerialNumber})
    eigrpTopologySummary.EntityData.Leafs.Append("route-count", types.YLeaf{"RouteCount", eigrpTopologySummary.RouteCount})
    eigrpTopologySummary.EntityData.Leafs.Append("path-count", types.YLeaf{"PathCount", eigrpTopologySummary.PathCount})
    eigrpTopologySummary.EntityData.Leafs.Append("dummy-count", types.YLeaf{"DummyCount", eigrpTopologySummary.DummyCount})
    eigrpTopologySummary.EntityData.Leafs.Append("ddb-name", types.YLeaf{"DdbName", eigrpTopologySummary.DdbName})
    eigrpTopologySummary.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", eigrpTopologySummary.InterfaceCount})
    eigrpTopologySummary.EntityData.Leafs.Append("handles-used", types.YLeaf{"HandlesUsed", eigrpTopologySummary.HandlesUsed})
    eigrpTopologySummary.EntityData.Leafs.Append("active-interface-count", types.YLeaf{"ActiveInterfaceCount", eigrpTopologySummary.ActiveInterfaceCount})

    eigrpTopologySummary.EntityData.YListKeys = []string {}

    return &(eigrpTopologySummary.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpTopologySummary_Quiescent
// Quiescent interfaces
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpTopologySummary_Quiescent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface Name. The type is string.
    QuiescentInterfaceList interface{}
}

func (quiescent *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpTopologySummary_Quiescent) GetEntityData() *types.CommonEntityData {
    quiescent.EntityData.YFilter = quiescent.YFilter
    quiescent.EntityData.YangName = "quiescent"
    quiescent.EntityData.BundleName = "cisco_ios_xr"
    quiescent.EntityData.ParentYangName = "eigrp-topology-summary"
    quiescent.EntityData.SegmentPath = "quiescent" + types.AddNoKeyToken(quiescent)
    quiescent.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/eigrp-topology-summary/" + quiescent.EntityData.SegmentPath
    quiescent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    quiescent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    quiescent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    quiescent.EntityData.Children = types.NewOrderedMap()
    quiescent.EntityData.Leafs = types.NewOrderedMap()
    quiescent.EntityData.Leafs.Append("quiescent-interface-list", types.YLeaf{"QuiescentInterfaceList", quiescent.QuiescentInterfaceList})

    quiescent.EntityData.YListKeys = []string {}

    return &(quiescent.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces
// EIGRP interfaces
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information for an EIGRP interface. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces_Interface.
    Interface []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces_Interface
}

func (interfaces *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "as"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/" + interfaces.EntityData.SegmentPath
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

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces_Interface
// Information for an EIGRP interface
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // AS Number. The type is interface{} with range: 0..4294967295.
    AsNumber interface{}

    // Peer Count. The type is interface{} with range: 0..4294967295.
    PeerCount interface{}

    // Classic Peer Count. The type is interface{} with range: 0..4294967295.
    ClassicPeerCount interface{}

    // Wide Peer Count. The type is interface{} with range: 0..4294967295.
    WidePeerCount interface{}

    // Unreliable transmissions. The type is interface{} with range:
    // 0..4294967295.
    UnreliableTransmits interface{}

    // Reliable transmissions. The type is interface{} with range: 0..4294967295.
    ReliableTransmits interface{}

    // Total SRTT. The type is interface{} with range: 0..4294967295.
    TotalSrtt interface{}

    // Send interval for Unreliable transmissions. The type is interface{} with
    // range: 0..4294967295.
    UnreliableSendInterval interface{}

    // Send interval for Reliable transmissions. The type is interface{} with
    // range: 0..4294967295.
    ReliableSendInterval interface{}

    // Last multicast flow delay. The type is interface{} with range:
    // 0..4294967295.
    LastMcFlowDelay interface{}

    // Number of pending routes. The type is interface{} with range:
    // 0..4294967295.
    PendingRoutes interface{}

    // Hello interval. The type is interface{} with range: 0..4294967295.
    HelloInterval interface{}

    // Hold time. The type is interface{} with range: 0..4294967295.
    Holdtime interface{}

    // BFD enabled. The type is bool.
    BfdEnabled interface{}

    // BFD interval. The type is interface{} with range: 0..4294967295.
    BfdInterval interface{}

    // BFD multiplier. The type is interface{} with range: 0..4294967295.
    BfdMultiplier interface{}

    // Is serno present. The type is bool.
    SerialNumberPresent interface{}

    // Thread serial number. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmitSerialNumber interface{}

    // Packetize Timer pending. The type is bool.
    PacketizePending interface{}

    // Unreliable multicasts sent. The type is interface{} with range:
    // 0..4294967295.
    UnreliableMulticastSent interface{}

    // Reliable multicasts sent. The type is interface{} with range:
    // 0..4294967295.
    ReliableMulticastSent interface{}

    // Unreliable unicasts sent. The type is interface{} with range:
    // 0..4294967295.
    UnreliableUnicastSent interface{}

    // Reliable unicasts sent. The type is interface{} with range: 0..4294967295.
    ReliableUnicastSent interface{}

    // Multicast Exceptions sent. The type is interface{} with range:
    // 0..4294967295.
    MulticastExceptionsSent interface{}

    // CR packets sent. The type is interface{} with range: 0..4294967295.
    CrPacketsSent interface{}

    // Suppressed Acks. The type is interface{} with range: 0..4294967295.
    AcksSuppressed interface{}

    // Retransmissions sent. The type is interface{} with range: 0..4294967295.
    RetransmissionsSent interface{}

    // Out-of-sequence received. The type is interface{} with range:
    // 0..4294967295.
    OutOfSequenceReceived interface{}

    // All stub peers. The type is bool.
    StubInterface interface{}

    // Next-hop-self enabled. The type is bool.
    NextHopSelfEnabled interface{}

    // SplitHorizon enabled. The type is bool.
    SplitHorizonEnabled interface{}

    // Interface is passive. The type is bool.
    PassiveInterface interface{}

    // Bandwidth percent. The type is interface{} with range: 0..4294967295. Units
    // are percentage.
    BandwidthPercent interface{}

    // Site of Origin Type. The type is EigrpBdSoo.
    SiteOfOriginType interface{}

    // Site of Origin. The type is string.
    SiteOfOrigin interface{}

    // Authentication Mode. The type is interface{} with range: 0..4294967295.
    AuthMode interface{}

    // Authentication Keychain Name. The type is string.
    AuthKeychain interface{}

    // Authentication key exists. The type is bool.
    AuthKeyExists interface{}

    // Authentication key programmed with MD5 algorithm. The type is bool.
    AuthKeyMd5 interface{}

    // Current active Authentication Key Id. The type is interface{} with range:
    // 0..18446744073709551615.
    AuthKeyId interface{}

    // Total packets received. The type is interface{} with range: 0..4294967295.
    TotalPktRecvd interface{}

    // Packets dropped due to wrong keychain configured. The type is interface{}
    // with range: 0..4294967295.
    PktDropWrongKc interface{}

    // Packets dropped due to missing authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropNoAuth interface{}

    // Packets dropped due to invalid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropInvalidAuth interface{}

    // Packets accepted with valid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktAcceptedValidAuth interface{}

    // Deprecated. Please migrate to use Bandwidth64. The type is interface{} with
    // range: 0..4294967295.
    Bandwidth interface{}

    // Bandwidth. The type is interface{} with range: 0..18446744073709551615.
    Bandwidth64 interface{}

    // Deprecated. Please migrate to use Delay64. The value of this object might
    // wrap if it is in picosecond units. The type is interface{} with range:
    // 0..4294967295.
    Delay interface{}

    // Delay. The type is interface{} with range: 0..18446744073709551615.
    Delay64 interface{}

    // Delay unit. The type is EigrpBdDelayUnit.
    DelayUnit interface{}

    // Reliability. The type is interface{} with range: 0..4294967295.
    Reliability interface{}

    // Load. The type is interface{} with range: 0..4294967295.
    Load interface{}

    // MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Deprecated. Please migrate to use ConfiguredBandwidth64. The type is
    // interface{} with range: 0..4294967295.
    ConfiguredBandwidth interface{}

    // Configured bandwidth. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfiguredBandwidth64 interface{}

    // Deprecated. Please migrate to use ConfiguredDelay64. The value of this
    // object might wrap if it is in picosecond units. The type is interface{}
    // with range: 0..4294967295.
    ConfiguredDelay interface{}

    // Configured delay. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfiguredDelay64 interface{}

    // Configured delay unit. The type is EigrpBdDelayUnit.
    ConfiguredDelayUnit interface{}

    // Configured reliability. The type is interface{} with range: 0..4294967295.
    ConfiguredReliability interface{}

    // Configured load. The type is interface{} with range: 0..4294967295.
    ConfiguredLoad interface{}

    // Bandwidth configured. The type is bool.
    ConfiguredBandwidthFlag interface{}

    // Delay configured. The type is bool.
    ConfiguredDelayFlag interface{}

    // Reliability configured. The type is bool.
    ConfiguredReliabilityFlag interface{}

    // Load configured. The type is bool.
    ConfiguredLoadFlag interface{}

    // Interface is UP. The type is bool.
    Up interface{}

    // Interface type is supported. The type is bool.
    TypeSupported interface{}

    // ITAL Record valid. The type is bool.
    ItalRecordFound interface{}

    // Interface config exists. The type is bool.
    Configured interface{}

    // Requested socket state. The type is bool.
    MulticastEnabled interface{}

    // Setup socket state. The type is bool.
    SocketSetup interface{}

    // Setup LPTS socket state. The type is bool.
    LptsSocketSetup interface{}

    // Primary IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PrimaryIpv4Address interface{}

    // IPv6 LL Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6LinkLocalAddr interface{}

    // Primary prefix length. The type is interface{} with range: 0..4294967295.
    PrimaryPrefixLength interface{}

    // Interface Handle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // IM Interface Type. The type is interface{} with range: 0..4294967295.
    InterfaceType interface{}

    // Interface Configured Items. The type is interface{} with range:
    // 0..4294967295.
    ConfiguredItems interface{}

    // Passive-Interface configured. The type is bool.
    IsPassiveEnabled interface{}

    // Passive-Interface disabled. The type is bool.
    IsPassiveDisabled interface{}

    // Static Neighbors. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces_Interface_StaticNeighbor.
    StaticNeighbor []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces_Interface_StaticNeighbor
}

func (self *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("static-neighbor", types.YChild{"StaticNeighbor", nil})
    for i := range self.StaticNeighbor {
        types.SetYListKey(self.StaticNeighbor[i], i)
        self.EntityData.Children.Append(types.GetSegmentPath(self.StaticNeighbor[i]), types.YChild{"StaticNeighbor", self.StaticNeighbor[i]})
    }
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", self.Afi})
    self.EntityData.Leafs.Append("as-number", types.YLeaf{"AsNumber", self.AsNumber})
    self.EntityData.Leafs.Append("peer-count", types.YLeaf{"PeerCount", self.PeerCount})
    self.EntityData.Leafs.Append("classic-peer-count", types.YLeaf{"ClassicPeerCount", self.ClassicPeerCount})
    self.EntityData.Leafs.Append("wide-peer-count", types.YLeaf{"WidePeerCount", self.WidePeerCount})
    self.EntityData.Leafs.Append("unreliable-transmits", types.YLeaf{"UnreliableTransmits", self.UnreliableTransmits})
    self.EntityData.Leafs.Append("reliable-transmits", types.YLeaf{"ReliableTransmits", self.ReliableTransmits})
    self.EntityData.Leafs.Append("total-srtt", types.YLeaf{"TotalSrtt", self.TotalSrtt})
    self.EntityData.Leafs.Append("unreliable-send-interval", types.YLeaf{"UnreliableSendInterval", self.UnreliableSendInterval})
    self.EntityData.Leafs.Append("reliable-send-interval", types.YLeaf{"ReliableSendInterval", self.ReliableSendInterval})
    self.EntityData.Leafs.Append("last-mc-flow-delay", types.YLeaf{"LastMcFlowDelay", self.LastMcFlowDelay})
    self.EntityData.Leafs.Append("pending-routes", types.YLeaf{"PendingRoutes", self.PendingRoutes})
    self.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", self.HelloInterval})
    self.EntityData.Leafs.Append("holdtime", types.YLeaf{"Holdtime", self.Holdtime})
    self.EntityData.Leafs.Append("bfd-enabled", types.YLeaf{"BfdEnabled", self.BfdEnabled})
    self.EntityData.Leafs.Append("bfd-interval", types.YLeaf{"BfdInterval", self.BfdInterval})
    self.EntityData.Leafs.Append("bfd-multiplier", types.YLeaf{"BfdMultiplier", self.BfdMultiplier})
    self.EntityData.Leafs.Append("serial-number-present", types.YLeaf{"SerialNumberPresent", self.SerialNumberPresent})
    self.EntityData.Leafs.Append("transmit-serial-number", types.YLeaf{"TransmitSerialNumber", self.TransmitSerialNumber})
    self.EntityData.Leafs.Append("packetize-pending", types.YLeaf{"PacketizePending", self.PacketizePending})
    self.EntityData.Leafs.Append("unreliable-multicast-sent", types.YLeaf{"UnreliableMulticastSent", self.UnreliableMulticastSent})
    self.EntityData.Leafs.Append("reliable-multicast-sent", types.YLeaf{"ReliableMulticastSent", self.ReliableMulticastSent})
    self.EntityData.Leafs.Append("unreliable-unicast-sent", types.YLeaf{"UnreliableUnicastSent", self.UnreliableUnicastSent})
    self.EntityData.Leafs.Append("reliable-unicast-sent", types.YLeaf{"ReliableUnicastSent", self.ReliableUnicastSent})
    self.EntityData.Leafs.Append("multicast-exceptions-sent", types.YLeaf{"MulticastExceptionsSent", self.MulticastExceptionsSent})
    self.EntityData.Leafs.Append("cr-packets-sent", types.YLeaf{"CrPacketsSent", self.CrPacketsSent})
    self.EntityData.Leafs.Append("acks-suppressed", types.YLeaf{"AcksSuppressed", self.AcksSuppressed})
    self.EntityData.Leafs.Append("retransmissions-sent", types.YLeaf{"RetransmissionsSent", self.RetransmissionsSent})
    self.EntityData.Leafs.Append("out-of-sequence-received", types.YLeaf{"OutOfSequenceReceived", self.OutOfSequenceReceived})
    self.EntityData.Leafs.Append("stub-interface", types.YLeaf{"StubInterface", self.StubInterface})
    self.EntityData.Leafs.Append("next-hop-self-enabled", types.YLeaf{"NextHopSelfEnabled", self.NextHopSelfEnabled})
    self.EntityData.Leafs.Append("split-horizon-enabled", types.YLeaf{"SplitHorizonEnabled", self.SplitHorizonEnabled})
    self.EntityData.Leafs.Append("passive-interface", types.YLeaf{"PassiveInterface", self.PassiveInterface})
    self.EntityData.Leafs.Append("bandwidth-percent", types.YLeaf{"BandwidthPercent", self.BandwidthPercent})
    self.EntityData.Leafs.Append("site-of-origin-type", types.YLeaf{"SiteOfOriginType", self.SiteOfOriginType})
    self.EntityData.Leafs.Append("site-of-origin", types.YLeaf{"SiteOfOrigin", self.SiteOfOrigin})
    self.EntityData.Leafs.Append("auth-mode", types.YLeaf{"AuthMode", self.AuthMode})
    self.EntityData.Leafs.Append("auth-keychain", types.YLeaf{"AuthKeychain", self.AuthKeychain})
    self.EntityData.Leafs.Append("auth-key-exists", types.YLeaf{"AuthKeyExists", self.AuthKeyExists})
    self.EntityData.Leafs.Append("auth-key-md5", types.YLeaf{"AuthKeyMd5", self.AuthKeyMd5})
    self.EntityData.Leafs.Append("auth-key-id", types.YLeaf{"AuthKeyId", self.AuthKeyId})
    self.EntityData.Leafs.Append("total-pkt-recvd", types.YLeaf{"TotalPktRecvd", self.TotalPktRecvd})
    self.EntityData.Leafs.Append("pkt-drop-wrong-kc", types.YLeaf{"PktDropWrongKc", self.PktDropWrongKc})
    self.EntityData.Leafs.Append("pkt-drop-no-auth", types.YLeaf{"PktDropNoAuth", self.PktDropNoAuth})
    self.EntityData.Leafs.Append("pkt-drop-invalid-auth", types.YLeaf{"PktDropInvalidAuth", self.PktDropInvalidAuth})
    self.EntityData.Leafs.Append("pkt-accepted-valid-auth", types.YLeaf{"PktAcceptedValidAuth", self.PktAcceptedValidAuth})
    self.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", self.Bandwidth})
    self.EntityData.Leafs.Append("bandwidth64", types.YLeaf{"Bandwidth64", self.Bandwidth64})
    self.EntityData.Leafs.Append("delay", types.YLeaf{"Delay", self.Delay})
    self.EntityData.Leafs.Append("delay64", types.YLeaf{"Delay64", self.Delay64})
    self.EntityData.Leafs.Append("delay-unit", types.YLeaf{"DelayUnit", self.DelayUnit})
    self.EntityData.Leafs.Append("reliability", types.YLeaf{"Reliability", self.Reliability})
    self.EntityData.Leafs.Append("load", types.YLeaf{"Load", self.Load})
    self.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", self.Mtu})
    self.EntityData.Leafs.Append("configured-bandwidth", types.YLeaf{"ConfiguredBandwidth", self.ConfiguredBandwidth})
    self.EntityData.Leafs.Append("configured-bandwidth64", types.YLeaf{"ConfiguredBandwidth64", self.ConfiguredBandwidth64})
    self.EntityData.Leafs.Append("configured-delay", types.YLeaf{"ConfiguredDelay", self.ConfiguredDelay})
    self.EntityData.Leafs.Append("configured-delay64", types.YLeaf{"ConfiguredDelay64", self.ConfiguredDelay64})
    self.EntityData.Leafs.Append("configured-delay-unit", types.YLeaf{"ConfiguredDelayUnit", self.ConfiguredDelayUnit})
    self.EntityData.Leafs.Append("configured-reliability", types.YLeaf{"ConfiguredReliability", self.ConfiguredReliability})
    self.EntityData.Leafs.Append("configured-load", types.YLeaf{"ConfiguredLoad", self.ConfiguredLoad})
    self.EntityData.Leafs.Append("configured-bandwidth-flag", types.YLeaf{"ConfiguredBandwidthFlag", self.ConfiguredBandwidthFlag})
    self.EntityData.Leafs.Append("configured-delay-flag", types.YLeaf{"ConfiguredDelayFlag", self.ConfiguredDelayFlag})
    self.EntityData.Leafs.Append("configured-reliability-flag", types.YLeaf{"ConfiguredReliabilityFlag", self.ConfiguredReliabilityFlag})
    self.EntityData.Leafs.Append("configured-load-flag", types.YLeaf{"ConfiguredLoadFlag", self.ConfiguredLoadFlag})
    self.EntityData.Leafs.Append("up", types.YLeaf{"Up", self.Up})
    self.EntityData.Leafs.Append("type-supported", types.YLeaf{"TypeSupported", self.TypeSupported})
    self.EntityData.Leafs.Append("ital-record-found", types.YLeaf{"ItalRecordFound", self.ItalRecordFound})
    self.EntityData.Leafs.Append("configured", types.YLeaf{"Configured", self.Configured})
    self.EntityData.Leafs.Append("multicast-enabled", types.YLeaf{"MulticastEnabled", self.MulticastEnabled})
    self.EntityData.Leafs.Append("socket-setup", types.YLeaf{"SocketSetup", self.SocketSetup})
    self.EntityData.Leafs.Append("lpts-socket-setup", types.YLeaf{"LptsSocketSetup", self.LptsSocketSetup})
    self.EntityData.Leafs.Append("primary-ipv4-address", types.YLeaf{"PrimaryIpv4Address", self.PrimaryIpv4Address})
    self.EntityData.Leafs.Append("ipv6-link-local-addr", types.YLeaf{"Ipv6LinkLocalAddr", self.Ipv6LinkLocalAddr})
    self.EntityData.Leafs.Append("primary-prefix-length", types.YLeaf{"PrimaryPrefixLength", self.PrimaryPrefixLength})
    self.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", self.InterfaceHandle})
    self.EntityData.Leafs.Append("interface-type", types.YLeaf{"InterfaceType", self.InterfaceType})
    self.EntityData.Leafs.Append("configured-items", types.YLeaf{"ConfiguredItems", self.ConfiguredItems})
    self.EntityData.Leafs.Append("is-passive-enabled", types.YLeaf{"IsPassiveEnabled", self.IsPassiveEnabled})
    self.EntityData.Leafs.Append("is-passive-disabled", types.YLeaf{"IsPassiveDisabled", self.IsPassiveDisabled})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces_Interface_StaticNeighbor
// Static Neighbors
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces_Interface_StaticNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (staticNeighbor *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Interfaces_Interface_StaticNeighbor) GetEntityData() *types.CommonEntityData {
    staticNeighbor.EntityData.YFilter = staticNeighbor.YFilter
    staticNeighbor.EntityData.YangName = "static-neighbor"
    staticNeighbor.EntityData.BundleName = "cisco_ios_xr"
    staticNeighbor.EntityData.ParentYangName = "interface"
    staticNeighbor.EntityData.SegmentPath = "static-neighbor" + types.AddNoKeyToken(staticNeighbor)
    staticNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/interfaces/interface/" + staticNeighbor.EntityData.SegmentPath
    staticNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticNeighbor.EntityData.Children = types.NewOrderedMap()
    staticNeighbor.EntityData.Leafs = types.NewOrderedMap()
    staticNeighbor.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", staticNeighbor.Ipv4Address})
    staticNeighbor.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", staticNeighbor.Ipv6Address})

    staticNeighbor.EntityData.YListKeys = []string {}

    return &(staticNeighbor.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpEvents
// Events for a specific VRF AF
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpEvents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // AS Number. The type is interface{} with range: 0..4294967295.
    AsNumber interface{}

    // Current event. The type is interface{} with range: 0..4294967295.
    CurrentEventIndex interface{}

    // Seconds since EIGRP started (absolute). The type is interface{} with range:
    // 0..4294967295. Units are second.
    EigrpStartAbsoluteSeconds interface{}

    // Seconds since EIGRP started (absolute). The type is interface{} with range:
    // 0..4294967295. Units are second.
    EigrpStartAbsoluteNanoseconds interface{}

    // Seconds since EIGRP started (relative). The type is interface{} with range:
    // 0..4294967295. Units are second.
    EigrpStartRelativeSeconds interface{}

    // Seconds since EIGRP started (relative). The type is interface{} with range:
    // 0..4294967295. Units are second.
    EigrpStartRelativeNanoseconds interface{}
}

func (eigrpEvents *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpEvents) GetEntityData() *types.CommonEntityData {
    eigrpEvents.EntityData.YFilter = eigrpEvents.YFilter
    eigrpEvents.EntityData.YangName = "eigrp-events"
    eigrpEvents.EntityData.BundleName = "cisco_ios_xr"
    eigrpEvents.EntityData.ParentYangName = "as"
    eigrpEvents.EntityData.SegmentPath = "eigrp-events"
    eigrpEvents.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/" + eigrpEvents.EntityData.SegmentPath
    eigrpEvents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrpEvents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrpEvents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrpEvents.EntityData.Children = types.NewOrderedMap()
    eigrpEvents.EntityData.Leafs = types.NewOrderedMap()
    eigrpEvents.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", eigrpEvents.Afi})
    eigrpEvents.EntityData.Leafs.Append("as-number", types.YLeaf{"AsNumber", eigrpEvents.AsNumber})
    eigrpEvents.EntityData.Leafs.Append("current-event-index", types.YLeaf{"CurrentEventIndex", eigrpEvents.CurrentEventIndex})
    eigrpEvents.EntityData.Leafs.Append("eigrp-start-absolute-seconds", types.YLeaf{"EigrpStartAbsoluteSeconds", eigrpEvents.EigrpStartAbsoluteSeconds})
    eigrpEvents.EntityData.Leafs.Append("eigrp-start-absolute-nanoseconds", types.YLeaf{"EigrpStartAbsoluteNanoseconds", eigrpEvents.EigrpStartAbsoluteNanoseconds})
    eigrpEvents.EntityData.Leafs.Append("eigrp-start-relative-seconds", types.YLeaf{"EigrpStartRelativeSeconds", eigrpEvents.EigrpStartRelativeSeconds})
    eigrpEvents.EntityData.Leafs.Append("eigrp-start-relative-nanoseconds", types.YLeaf{"EigrpStartRelativeNanoseconds", eigrpEvents.EigrpStartRelativeNanoseconds})

    eigrpEvents.EntityData.YListKeys = []string {}

    return &(eigrpEvents.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors
// EIGRP neighbors
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information on one EIGRP neighbor. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor.
    Neighbor []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor
}

func (neighbors *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "as"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/" + neighbors.EntityData.SegmentPath
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor
// Information on one EIGRP neighbor
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor Address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    NeighborAddress interface{}

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // AS Number. The type is interface{} with range: 0..4294967295.
    AsNumber interface{}

    // Is it a suspended peer. The type is bool.
    PeerSuspended interface{}

    // Peer handle. The type is interface{} with range: 0..4294967295.
    PeerHandle interface{}

    // Interface name. The type is string.
    InterfaceList interface{}

    // Hold time. The type is interface{} with range: 0..4294967295.
    Holdtime interface{}

    // UP time (seconds). The type is interface{} with range: 0..4294967295. Units
    // are second.
    Uptime interface{}

    // Round trip time. The type is interface{} with range: 0..4294967295.
    Srtt interface{}

    // RTO. The type is interface{} with range: 0..4294967295.
    Rto interface{}

    // BFD enabled. The type is bool.
    BfdEnabled interface{}

    // Q counts. The type is interface{} with range: 0..4294967295.
    QueueCount interface{}

    // Last sequence number. The type is interface{} with range: 0..4294967295.
    LastSequenceNumber interface{}

    // Is it a static neighbor. The type is bool.
    StaticNeighbor interface{}

    // Is it a remote ucast neighbor. The type is bool.
    RemoteNeighbor interface{}

    // Hop count of the static peer. The type is interface{} with range: 0..255.
    HopCount interface{}

    // Is Restart time configured. The type is bool.
    RestartConfigured interface{}

    // Restart time (seconds). The type is interface{} with range: 0..4294967295.
    // Units are second.
    RestartTime interface{}

    // Last startup serial number. The type is interface{} with range:
    // 0..18446744073709551615.
    LastStartupSerialNumber interface{}

    // IOS Major version. The type is interface{} with range: 0..255.
    IosMajorVersion interface{}

    // IOS Minor version. The type is interface{} with range: 0..255.
    IosMinorVersion interface{}

    // EIGRP Major version. The type is interface{} with range: 0..255.
    EigrpMajorVersion interface{}

    // EIGRP Major version. The type is interface{} with range: 0..255.
    EigrpMinorVersion interface{}

    // Retransmission count. The type is interface{} with range: 0..4294967295.
    RetransmissionCount interface{}

    // Retry count. The type is interface{} with range: 0..4294967295.
    RetryCount interface{}

    // Need EIGRP Init message. The type is bool.
    NeedInit interface{}

    // Need EIGRP InitAck message. The type is bool.
    NeedInitAck interface{}

    // Reinitialization needed. The type is bool.
    ReinitializationNeeded interface{}

    // Reinit period. The type is interface{} with range: 0..4294967295.
    ReinitStart interface{}

    // Prefix count. The type is interface{} with range: 0..4294967295.
    PeerPrefixCount interface{}

    // Is it stubbed. The type is bool.
    Stubbed interface{}

    // Connected routes accepted. The type is bool.
    AllowConnected interface{}

    // Static routes accepted. The type is bool.
    AllowStatic interface{}

    // Summary routes accepted. The type is bool.
    AllowSummaries interface{}

    // Redist'ed routes accepted. The type is bool.
    AllowRedistributed interface{}

    // Test handle flag. The type is bool.
    TestHandle interface{}

    // Is it stubbed. The type is bool.
    StubbedInterface interface{}

    // Suspension manually reset. The type is bool.
    SuspendedReset interface{}

    // Suspended time left. The type is interface{} with range: 0..4294967295.
    SuspendedTimeLeft interface{}

    // Peer address.
    Source Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor_Source

    // Neighbor Queue. The type is slice of
    // Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor_NeighborQueue.
    NeighborQueue []*Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor_NeighborQueue
}

func (neighbor *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.NeighborAddress, "neighbor-address")
    neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/neighbors/" + neighbor.EntityData.SegmentPath
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("source", types.YChild{"Source", &neighbor.Source})
    neighbor.EntityData.Children.Append("neighbor-queue", types.YChild{"NeighborQueue", nil})
    for i := range neighbor.NeighborQueue {
        types.SetYListKey(neighbor.NeighborQueue[i], i)
        neighbor.EntityData.Children.Append(types.GetSegmentPath(neighbor.NeighborQueue[i]), types.YChild{"NeighborQueue", neighbor.NeighborQueue[i]})
    }
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", neighbor.NeighborAddress})
    neighbor.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", neighbor.Afi})
    neighbor.EntityData.Leafs.Append("as-number", types.YLeaf{"AsNumber", neighbor.AsNumber})
    neighbor.EntityData.Leafs.Append("peer-suspended", types.YLeaf{"PeerSuspended", neighbor.PeerSuspended})
    neighbor.EntityData.Leafs.Append("peer-handle", types.YLeaf{"PeerHandle", neighbor.PeerHandle})
    neighbor.EntityData.Leafs.Append("interface-list", types.YLeaf{"InterfaceList", neighbor.InterfaceList})
    neighbor.EntityData.Leafs.Append("holdtime", types.YLeaf{"Holdtime", neighbor.Holdtime})
    neighbor.EntityData.Leafs.Append("uptime", types.YLeaf{"Uptime", neighbor.Uptime})
    neighbor.EntityData.Leafs.Append("srtt", types.YLeaf{"Srtt", neighbor.Srtt})
    neighbor.EntityData.Leafs.Append("rto", types.YLeaf{"Rto", neighbor.Rto})
    neighbor.EntityData.Leafs.Append("bfd-enabled", types.YLeaf{"BfdEnabled", neighbor.BfdEnabled})
    neighbor.EntityData.Leafs.Append("queue-count", types.YLeaf{"QueueCount", neighbor.QueueCount})
    neighbor.EntityData.Leafs.Append("last-sequence-number", types.YLeaf{"LastSequenceNumber", neighbor.LastSequenceNumber})
    neighbor.EntityData.Leafs.Append("static-neighbor", types.YLeaf{"StaticNeighbor", neighbor.StaticNeighbor})
    neighbor.EntityData.Leafs.Append("remote-neighbor", types.YLeaf{"RemoteNeighbor", neighbor.RemoteNeighbor})
    neighbor.EntityData.Leafs.Append("hop-count", types.YLeaf{"HopCount", neighbor.HopCount})
    neighbor.EntityData.Leafs.Append("restart-configured", types.YLeaf{"RestartConfigured", neighbor.RestartConfigured})
    neighbor.EntityData.Leafs.Append("restart-time", types.YLeaf{"RestartTime", neighbor.RestartTime})
    neighbor.EntityData.Leafs.Append("last-startup-serial-number", types.YLeaf{"LastStartupSerialNumber", neighbor.LastStartupSerialNumber})
    neighbor.EntityData.Leafs.Append("ios-major-version", types.YLeaf{"IosMajorVersion", neighbor.IosMajorVersion})
    neighbor.EntityData.Leafs.Append("ios-minor-version", types.YLeaf{"IosMinorVersion", neighbor.IosMinorVersion})
    neighbor.EntityData.Leafs.Append("eigrp-major-version", types.YLeaf{"EigrpMajorVersion", neighbor.EigrpMajorVersion})
    neighbor.EntityData.Leafs.Append("eigrp-minor-version", types.YLeaf{"EigrpMinorVersion", neighbor.EigrpMinorVersion})
    neighbor.EntityData.Leafs.Append("retransmission-count", types.YLeaf{"RetransmissionCount", neighbor.RetransmissionCount})
    neighbor.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", neighbor.RetryCount})
    neighbor.EntityData.Leafs.Append("need-init", types.YLeaf{"NeedInit", neighbor.NeedInit})
    neighbor.EntityData.Leafs.Append("need-init-ack", types.YLeaf{"NeedInitAck", neighbor.NeedInitAck})
    neighbor.EntityData.Leafs.Append("reinitialization-needed", types.YLeaf{"ReinitializationNeeded", neighbor.ReinitializationNeeded})
    neighbor.EntityData.Leafs.Append("reinit-start", types.YLeaf{"ReinitStart", neighbor.ReinitStart})
    neighbor.EntityData.Leafs.Append("peer-prefix-count", types.YLeaf{"PeerPrefixCount", neighbor.PeerPrefixCount})
    neighbor.EntityData.Leafs.Append("stubbed", types.YLeaf{"Stubbed", neighbor.Stubbed})
    neighbor.EntityData.Leafs.Append("allow-connected", types.YLeaf{"AllowConnected", neighbor.AllowConnected})
    neighbor.EntityData.Leafs.Append("allow-static", types.YLeaf{"AllowStatic", neighbor.AllowStatic})
    neighbor.EntityData.Leafs.Append("allow-summaries", types.YLeaf{"AllowSummaries", neighbor.AllowSummaries})
    neighbor.EntityData.Leafs.Append("allow-redistributed", types.YLeaf{"AllowRedistributed", neighbor.AllowRedistributed})
    neighbor.EntityData.Leafs.Append("test-handle", types.YLeaf{"TestHandle", neighbor.TestHandle})
    neighbor.EntityData.Leafs.Append("stubbed-interface", types.YLeaf{"StubbedInterface", neighbor.StubbedInterface})
    neighbor.EntityData.Leafs.Append("suspended-reset", types.YLeaf{"SuspendedReset", neighbor.SuspendedReset})
    neighbor.EntityData.Leafs.Append("suspended-time-left", types.YLeaf{"SuspendedTimeLeft", neighbor.SuspendedTimeLeft})

    neighbor.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(neighbor.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor_Source
// Peer address
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (source *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "neighbor"
    source.EntityData.SegmentPath = "source"
    source.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/neighbors/neighbor/" + source.EntityData.SegmentPath
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = types.NewOrderedMap()
    source.EntityData.Leafs = types.NewOrderedMap()
    source.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", source.Ipv4Address})
    source.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", source.Ipv6Address})

    source.EntityData.YListKeys = []string {}

    return &(source.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor_NeighborQueue
// Neighbor Queue
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor_NeighborQueue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Operation Code. The type is interface{} with range: 0..4294967295.
    OperationCode interface{}

    // ACK sequence number. The type is interface{} with range: 0..4294967295.
    AckSequnceNumber interface{}

    // Starting serial number. The type is interface{} with range:
    // 0..18446744073709551615.
    StartSerialNumber interface{}

    // Ending serial number. The type is interface{} with range:
    // 0..18446744073709551615.
    EndSerialNumber interface{}

    // Pregenerated pak. The type is bool.
    Pregenerated interface{}

    // pak len. The type is interface{} with range: 0..4294967295.
    PacketLength interface{}

    // Has a pak been sent. The type is bool.
    TimeSentFlag interface{}

    // Time sent. The type is interface{} with range: 0..4294967295.
    TimeSent interface{}

    // Is the init bit set. The type is bool.
    InitBitSet interface{}

    // Is it sequenced. The type is bool.
    Sequenced interface{}
}

func (neighborQueue *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_Neighbors_Neighbor_NeighborQueue) GetEntityData() *types.CommonEntityData {
    neighborQueue.EntityData.YFilter = neighborQueue.YFilter
    neighborQueue.EntityData.YangName = "neighbor-queue"
    neighborQueue.EntityData.BundleName = "cisco_ios_xr"
    neighborQueue.EntityData.ParentYangName = "neighbor"
    neighborQueue.EntityData.SegmentPath = "neighbor-queue" + types.AddNoKeyToken(neighborQueue)
    neighborQueue.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/neighbors/neighbor/" + neighborQueue.EntityData.SegmentPath
    neighborQueue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborQueue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborQueue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborQueue.EntityData.Children = types.NewOrderedMap()
    neighborQueue.EntityData.Leafs = types.NewOrderedMap()
    neighborQueue.EntityData.Leafs.Append("operation-code", types.YLeaf{"OperationCode", neighborQueue.OperationCode})
    neighborQueue.EntityData.Leafs.Append("ack-sequnce-number", types.YLeaf{"AckSequnceNumber", neighborQueue.AckSequnceNumber})
    neighborQueue.EntityData.Leafs.Append("start-serial-number", types.YLeaf{"StartSerialNumber", neighborQueue.StartSerialNumber})
    neighborQueue.EntityData.Leafs.Append("end-serial-number", types.YLeaf{"EndSerialNumber", neighborQueue.EndSerialNumber})
    neighborQueue.EntityData.Leafs.Append("pregenerated", types.YLeaf{"Pregenerated", neighborQueue.Pregenerated})
    neighborQueue.EntityData.Leafs.Append("packet-length", types.YLeaf{"PacketLength", neighborQueue.PacketLength})
    neighborQueue.EntityData.Leafs.Append("time-sent-flag", types.YLeaf{"TimeSentFlag", neighborQueue.TimeSentFlag})
    neighborQueue.EntityData.Leafs.Append("time-sent", types.YLeaf{"TimeSent", neighborQueue.TimeSent})
    neighborQueue.EntityData.Leafs.Append("init-bit-set", types.YLeaf{"InitBitSet", neighborQueue.InitBitSet})
    neighborQueue.EntityData.Leafs.Append("sequenced", types.YLeaf{"Sequenced", neighborQueue.Sequenced})

    neighborQueue.EntityData.YListKeys = []string {}

    return &(neighborQueue.EntityData)
}

// Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpStatistics
// Statistics for a specific VRF AF
type Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // AS Number. The type is interface{} with range: 0..4294967295.
    AsNumber interface{}

    // Configured interfaces. The type is interface{} with range: 0..4294967295.
    ConfiguredInterfaceCount interface{}

    // Active interfaces. The type is interface{} with range: 0..4294967295.
    ActiveInterfacesCount interface{}

    // Activate address family success count. The type is interface{} with range:
    // 0..18446744073709551615.
    ActivateCount interface{}

    // Activate address family failure count. The type is interface{} with range:
    // 0..18446744073709551615.
    ActivateError interface{}

    // Activate address family last error. The type is interface{} with range:
    // -2147483648..2147483647.
    ActivateLastError interface{}

    // Deactivate address family success count. The type is interface{} with
    // range: 0..18446744073709551615.
    DeactivateCount interface{}

    // Deactivate address family failure count. The type is interface{} with
    // range: 0..18446744073709551615.
    DeactivateError interface{}

    // Deactivate address family last error. The type is interface{} with range:
    // -2147483648..2147483647.
    DeactivateLastError interface{}

    // Socket setup success count. The type is interface{} with range:
    // 0..18446744073709551615.
    SocketSet interface{}

    // Socket setup failure count. The type is interface{} with range:
    // 0..18446744073709551615.
    SocketSetError interface{}

    // Socket setup last error. The type is interface{} with range:
    // -2147483648..2147483647.
    SockSetLastError interface{}

    // Succeeded RAW packets in. The type is interface{} with range:
    // 0..18446744073709551615.
    RawPacketIn interface{}

    // Failed RAW packets ln. The type is interface{} with range:
    // 0..18446744073709551615.
    RawPacketInError interface{}

    // RAW packets in last error. The type is interface{} with range:
    // -2147483648..2147483647.
    RawPacketInLastError interface{}

    // Succeeded RAW packets out. The type is interface{} with range:
    // 0..18446744073709551615.
    RawPacketOut interface{}

    // Failed RAW packets out. The type is interface{} with range:
    // 0..18446744073709551615.
    RawPacketOutError interface{}

    // RAW Packets out last error. The type is interface{} with range:
    // -2147483648..2147483647.
    RawPacketOutLastError interface{}
}

func (eigrpStatistics *Eigrp_Processes_Process_Vrfs_Vrf_Afs_Af_Ases_As_EigrpStatistics) GetEntityData() *types.CommonEntityData {
    eigrpStatistics.EntityData.YFilter = eigrpStatistics.YFilter
    eigrpStatistics.EntityData.YangName = "eigrp-statistics"
    eigrpStatistics.EntityData.BundleName = "cisco_ios_xr"
    eigrpStatistics.EntityData.ParentYangName = "as"
    eigrpStatistics.EntityData.SegmentPath = "eigrp-statistics"
    eigrpStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-eigrp-oper:eigrp/processes/process/vrfs/vrf/afs/af/ases/as/" + eigrpStatistics.EntityData.SegmentPath
    eigrpStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrpStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrpStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrpStatistics.EntityData.Children = types.NewOrderedMap()
    eigrpStatistics.EntityData.Leafs = types.NewOrderedMap()
    eigrpStatistics.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", eigrpStatistics.Afi})
    eigrpStatistics.EntityData.Leafs.Append("as-number", types.YLeaf{"AsNumber", eigrpStatistics.AsNumber})
    eigrpStatistics.EntityData.Leafs.Append("configured-interface-count", types.YLeaf{"ConfiguredInterfaceCount", eigrpStatistics.ConfiguredInterfaceCount})
    eigrpStatistics.EntityData.Leafs.Append("active-interfaces-count", types.YLeaf{"ActiveInterfacesCount", eigrpStatistics.ActiveInterfacesCount})
    eigrpStatistics.EntityData.Leafs.Append("activate-count", types.YLeaf{"ActivateCount", eigrpStatistics.ActivateCount})
    eigrpStatistics.EntityData.Leafs.Append("activate-error", types.YLeaf{"ActivateError", eigrpStatistics.ActivateError})
    eigrpStatistics.EntityData.Leafs.Append("activate-last-error", types.YLeaf{"ActivateLastError", eigrpStatistics.ActivateLastError})
    eigrpStatistics.EntityData.Leafs.Append("deactivate-count", types.YLeaf{"DeactivateCount", eigrpStatistics.DeactivateCount})
    eigrpStatistics.EntityData.Leafs.Append("deactivate-error", types.YLeaf{"DeactivateError", eigrpStatistics.DeactivateError})
    eigrpStatistics.EntityData.Leafs.Append("deactivate-last-error", types.YLeaf{"DeactivateLastError", eigrpStatistics.DeactivateLastError})
    eigrpStatistics.EntityData.Leafs.Append("socket-set", types.YLeaf{"SocketSet", eigrpStatistics.SocketSet})
    eigrpStatistics.EntityData.Leafs.Append("socket-set-error", types.YLeaf{"SocketSetError", eigrpStatistics.SocketSetError})
    eigrpStatistics.EntityData.Leafs.Append("sock-set-last-error", types.YLeaf{"SockSetLastError", eigrpStatistics.SockSetLastError})
    eigrpStatistics.EntityData.Leafs.Append("raw-packet-in", types.YLeaf{"RawPacketIn", eigrpStatistics.RawPacketIn})
    eigrpStatistics.EntityData.Leafs.Append("raw-packet-in-error", types.YLeaf{"RawPacketInError", eigrpStatistics.RawPacketInError})
    eigrpStatistics.EntityData.Leafs.Append("raw-packet-in-last-error", types.YLeaf{"RawPacketInLastError", eigrpStatistics.RawPacketInLastError})
    eigrpStatistics.EntityData.Leafs.Append("raw-packet-out", types.YLeaf{"RawPacketOut", eigrpStatistics.RawPacketOut})
    eigrpStatistics.EntityData.Leafs.Append("raw-packet-out-error", types.YLeaf{"RawPacketOutError", eigrpStatistics.RawPacketOutError})
    eigrpStatistics.EntityData.Leafs.Append("raw-packet-out-last-error", types.YLeaf{"RawPacketOutLastError", eigrpStatistics.RawPacketOutLastError})

    eigrpStatistics.EntityData.YListKeys = []string {}

    return &(eigrpStatistics.EntityData)
}

