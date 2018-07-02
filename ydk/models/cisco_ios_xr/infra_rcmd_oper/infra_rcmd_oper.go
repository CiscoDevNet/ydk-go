// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-rcmd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   rcmd: Show command for Route Convergence Monitoring &
//     Diagnostics
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_rcmd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_rcmd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rcmd-oper rcmd}", reflect.TypeOf(Rcmd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rcmd-oper:rcmd", reflect.TypeOf(Rcmd{}))
}

// RcmdBagEnblDsbl represents status enum
type RcmdBagEnblDsbl string

const (
    // Disabled
    RcmdBagEnblDsbl_dsbl RcmdBagEnblDsbl = "dsbl"

    // Enabled
    RcmdBagEnblDsbl_enbl RcmdBagEnblDsbl = "enbl"
)

// RcmdBoolYesNo represents Boolean enum
type RcmdBoolYesNo string

const (
    // No
    RcmdBoolYesNo_no RcmdBoolYesNo = "no"

    // Yes
    RcmdBoolYesNo_yes RcmdBoolYesNo = "yes"
)

// RcmdPriorityLevel represents Level of priority
type RcmdPriorityLevel string

const (
    // Critical
    RcmdPriorityLevel_critical RcmdPriorityLevel = "critical"

    // High
    RcmdPriorityLevel_high RcmdPriorityLevel = "high"

    // Medium
    RcmdPriorityLevel_medium RcmdPriorityLevel = "medium"

    // Low
    RcmdPriorityLevel_low RcmdPriorityLevel = "low"
)

// RcmdSpfState represents Type of an ISIS Level
type RcmdSpfState string

const (
    // Complete
    RcmdSpfState_complete RcmdSpfState = "complete"

    // InComplete
    RcmdSpfState_in_complete RcmdSpfState = "in-complete"

    // Collecting data
    RcmdSpfState_collecting RcmdSpfState = "collecting"

    // No Route Change
    RcmdSpfState_no_route_change RcmdSpfState = "no-route-change"
)

// RcmdShowIpfrrLfa represents IP Frr LFA Types
type RcmdShowIpfrrLfa string

const (
    // No IP Frr LFA Type 
    RcmdShowIpfrrLfa_none RcmdShowIpfrrLfa = "none"

    // IP Frr Local LFA
    RcmdShowIpfrrLfa_local RcmdShowIpfrrLfa = "local"

    // IP Frr Remote LFA
    RcmdShowIpfrrLfa_remote RcmdShowIpfrrLfa = "remote"
)

// RcmdShowPrcsState represents Post Processing Info
type RcmdShowPrcsState string

const (
    // Success
    RcmdShowPrcsState_success RcmdShowPrcsState = "success"

    // Cpu overload
    RcmdShowPrcsState_cpu RcmdShowPrcsState = "cpu"

    // Memory overload
    RcmdShowPrcsState_memory RcmdShowPrcsState = "memory"
)

// RcmdShowCompId represents Component Info
type RcmdShowCompId string

const (
    // OSPF component
    RcmdShowCompId_ospf RcmdShowCompId = "ospf"

    // ISIS component
    RcmdShowCompId_isis RcmdShowCompId = "isis"

    // Max NA
    RcmdShowCompId_un_known RcmdShowCompId = "un-known"
)

// RcmdShowLdpSessionState represents LDP Session State
type RcmdShowLdpSessionState string

const (
    // GR Down State
    RcmdShowLdpSessionState_gr_down RcmdShowLdpSessionState = "gr-down"

    // GR Converging State
    RcmdShowLdpSessionState_gr_converging RcmdShowLdpSessionState = "gr-converging"

    // Establishing State
    RcmdShowLdpSessionState_establishing RcmdShowLdpSessionState = "establishing"

    // Converging State
    RcmdShowLdpSessionState_converging RcmdShowLdpSessionState = "converging"

    // Converged State
    RcmdShowLdpSessionState_converged RcmdShowLdpSessionState = "converged"

    // Retrying State
    RcmdShowLdpSessionState_retrying RcmdShowLdpSessionState = "retrying"

    // Cumulative Coverage for all the States
    RcmdShowLdpSessionState_total RcmdShowLdpSessionState = "total"
)

// RcmdProtocolId represents Protocol Info
type RcmdProtocolId string

const (
    // OSPF protocol
    RcmdProtocolId_ospf RcmdProtocolId = "ospf"

    // ISIS Prrotocol
    RcmdProtocolId_isis RcmdProtocolId = "isis"

    // Max NA
    RcmdProtocolId_na RcmdProtocolId = "na"
)

// RcmdShowLdpConvState represents LDP Convergence States
type RcmdShowLdpConvState string

const (
    // Not Fully Covered
    RcmdShowLdpConvState_not_full RcmdShowLdpConvState = "not-full"

    // Fully Covered
    RcmdShowLdpConvState_fully_covered RcmdShowLdpConvState = "fully-covered"

    // Backup Coverage Above Threshold
    RcmdShowLdpConvState_coverage_above_threshold RcmdShowLdpConvState = "coverage-above-threshold"

    // Backup Coverage Below Threshold
    RcmdShowLdpConvState_coverage_below_threshold RcmdShowLdpConvState = "coverage-below-threshold"

    // Backup Coverage is Flappping
    RcmdShowLdpConvState_coverage_flapping RcmdShowLdpConvState = "coverage-flapping"
)

// RcmdLinecardSpeed represents Comparative speed of programming on linecard
type RcmdLinecardSpeed string

const (
    // Other linecard
    RcmdLinecardSpeed_other RcmdLinecardSpeed = "other"

    // Fastest linecard
    RcmdLinecardSpeed_fastest RcmdLinecardSpeed = "fastest"

    // Slowest linecard
    RcmdLinecardSpeed_slowest RcmdLinecardSpeed = "slowest"
)

// RcmdShowNode represents Type of Node
type RcmdShowNode string

const (
    // Unknown Type
    RcmdShowNode_unknown RcmdShowNode = "unknown"

    // LC Type
    RcmdShowNode_lc RcmdShowNode = "lc"

    // RP Type
    RcmdShowNode_rp RcmdShowNode = "rp"
)

// RcmdShowLdpNeighbourStatus represents LDP Adjacency Session Status
type RcmdShowLdpNeighbourStatus string

const (
    // Down State
    RcmdShowLdpNeighbourStatus_down RcmdShowLdpNeighbourStatus = "down"

    // Up State
    RcmdShowLdpNeighbourStatus_up RcmdShowLdpNeighbourStatus = "up"
)

// RcmdIsisSpf represents Type of an ISIS SPF run
type RcmdIsisSpf string

const (
    // Full
    RcmdIsisSpf_full RcmdIsisSpf = "full"

    // Incremental
    RcmdIsisSpf_incremental RcmdIsisSpf = "incremental"

    // Next hop calculation
    RcmdIsisSpf_next_hop RcmdIsisSpf = "next-hop"

    // Partial route calculation
    RcmdIsisSpf_partial_route RcmdIsisSpf = "partial-route"
)

// RcmdShowRoutePathChange represents Type of route change
type RcmdShowRoutePathChange string

const (
    // Primary path is changed
    RcmdShowRoutePathChange_primary RcmdShowRoutePathChange = "primary"

    // Backup path is changed
    RcmdShowRoutePathChange_backup RcmdShowRoutePathChange = "backup"
)

// RcmdIsisLvl represents Type of an ISIS Level
type RcmdIsisLvl string

const (
    // Level 1
    RcmdIsisLvl_l1 RcmdIsisLvl = "l1"

    // Level 2
    RcmdIsisLvl_l2 RcmdIsisLvl = "l2"
)

// RcmdChange represents Type of change
type RcmdChange string

const (
    // Invalid
    RcmdChange_none RcmdChange = "none"

    // Added
    RcmdChange_add RcmdChange = "add"

    // Deleted
    RcmdChange_delete_ RcmdChange = "delete"

    // Modified
    RcmdChange_modify RcmdChange = "modify"

    // No Change
    RcmdChange_no_change RcmdChange = "no-change"
)

// RcmdShowRoute represents Route Types
type RcmdShowRoute string

const (
    // OSPF route init
    RcmdShowRoute_ospf RcmdShowRoute = "ospf"

    // OSPF Intra route
    RcmdShowRoute_intra RcmdShowRoute = "intra"

    // OSPF Inter route
    RcmdShowRoute_inter RcmdShowRoute = "inter"

    // OSPF External Type-1 Route
    RcmdShowRoute_ext_1 RcmdShowRoute = "ext-1"

    // OSPF External Type-2 Route
    RcmdShowRoute_ext_2 RcmdShowRoute = "ext-2"

    // OSPF NSSA Type-1 Route
    RcmdShowRoute_nssa_1 RcmdShowRoute = "nssa-1"

    // OSPF NSSA Type-2 Route
    RcmdShowRoute_nssa_2 RcmdShowRoute = "nssa-2"

    // ISIS route init
    RcmdShowRoute_isis RcmdShowRoute = "isis"

    // ISIS L1 Summary
    RcmdShowRoute_l1_summary RcmdShowRoute = "l1-summary"

    // ISIS L1
    RcmdShowRoute_l1 RcmdShowRoute = "l1"

    // ISIS L2 Summary
    RcmdShowRoute_l2_summary RcmdShowRoute = "l2-summary"

    // ISIS L2
    RcmdShowRoute_l2 RcmdShowRoute = "l2"

    // ISIS Inter Area Summary
    RcmdShowRoute_inter_area_summary RcmdShowRoute = "inter-area-summary"

    // ISIS Inter Area
    RcmdShowRoute_inter_area RcmdShowRoute = "inter-area"

    // ISIS Default Route Attached
    RcmdShowRoute_default_attached RcmdShowRoute = "default-attached"
)

// RcmdLdpEvent represents Type of LDP Event
type RcmdLdpEvent string

const (
    // Neighbor Event
    RcmdLdpEvent_neighbor RcmdLdpEvent = "neighbor"

    // Adjacency Event
    RcmdLdpEvent_adjacency RcmdLdpEvent = "adjacency"
)

// RcmdLsa represents Type of LSA
type RcmdLsa string

const (
    // Invalid LSA
    RcmdLsa_unknown RcmdLsa = "unknown"

    // Router LSA
    RcmdLsa_router RcmdLsa = "router"

    // Network LSA
    RcmdLsa_network RcmdLsa = "network"

    // Summary LSA
    RcmdLsa_summary RcmdLsa = "summary"

    // ASBR LSA
    RcmdLsa_asbr RcmdLsa = "asbr"

    // External LSA
    RcmdLsa_external RcmdLsa = "external"

    // Multicast LSA
    RcmdLsa_multicast RcmdLsa = "multicast"

    // NSSA LSA
    RcmdLsa_nssa RcmdLsa = "nssa"
)

// RcmdShowMem represents RCMD Memory Manager type
type RcmdShowMem string

const (
    // Standard type
    RcmdShowMem_standard RcmdShowMem = "standard"

    // Chunk type
    RcmdShowMem_chunk RcmdShowMem = "chunk"

    // EDM type
    RcmdShowMem_edm RcmdShowMem = "edm"

    // String type
    RcmdShowMem_string_ RcmdShowMem = "string"

    // Static type
    RcmdShowMem_static RcmdShowMem = "static"

    // Unknown type
    RcmdShowMem_unknown RcmdShowMem = "unknown"
)

// RcmdBagEnableDisable represents status enum
type RcmdBagEnableDisable string

const (
    // Disabled
    RcmdBagEnableDisable_disable RcmdBagEnableDisable = "disable"

    // Enabled
    RcmdBagEnableDisable_enable RcmdBagEnableDisable = "enable"
)

// RcmdShowInstState represents instance state
type RcmdShowInstState string

const (
    // Unknown state
    RcmdShowInstState_unknown RcmdShowInstState = "unknown"

    // Active state
    RcmdShowInstState_active RcmdShowInstState = "active"

    // InActive state
    RcmdShowInstState_in_active RcmdShowInstState = "in-active"

    // Max state
    RcmdShowInstState_na RcmdShowInstState = "na"
)

// RcmdLsChange represents Type of change
type RcmdLsChange string

const (
    // Added
    RcmdLsChange_new_ RcmdLsChange = "new"

    // Deleted
    RcmdLsChange_delete_ RcmdLsChange = "delete"

    // Modified
    RcmdLsChange_modify RcmdLsChange = "modify"

    // No operation
    RcmdLsChange_noop RcmdLsChange = "noop"
)

// RcmdShowIntfEvent represents Rcmd show intf event
type RcmdShowIntfEvent string

const (
    // Create
    RcmdShowIntfEvent_create RcmdShowIntfEvent = "create"

    // Delete
    RcmdShowIntfEvent_delete_ RcmdShowIntfEvent = "delete"

    // LinkUp
    RcmdShowIntfEvent_link_up RcmdShowIntfEvent = "link-up"

    // LinkDown
    RcmdShowIntfEvent_link_down RcmdShowIntfEvent = "link-down"

    // PrimaryAddress
    RcmdShowIntfEvent_primary_address RcmdShowIntfEvent = "primary-address"

    // SecondaryAddress
    RcmdShowIntfEvent_secondary_address RcmdShowIntfEvent = "secondary-address"

    // Ipv6LinkLocalAddress
    RcmdShowIntfEvent_ipv6_link_local_address RcmdShowIntfEvent = "ipv6-link-local-address"

    // Ipv6GlobalAddress
    RcmdShowIntfEvent_ipv6_global_address RcmdShowIntfEvent = "ipv6-global-address"

    // MTU
    RcmdShowIntfEvent_mtu RcmdShowIntfEvent = "mtu"

    // BandWidth
    RcmdShowIntfEvent_band_width RcmdShowIntfEvent = "band-width"

    // LDPSync
    RcmdShowIntfEvent_ldp_sync RcmdShowIntfEvent = "ldp-sync"

    // ForwardReference
    RcmdShowIntfEvent_forward_reference RcmdShowIntfEvent = "forward-reference"

    // LDPNoSync
    RcmdShowIntfEvent_ldp_no_sync RcmdShowIntfEvent = "ldp-no-sync"
)

// Rcmd
// Show command for Route Convergence Monitoring &
// Diagnostics
type Rcmd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for OSPF.
    Ospf Rcmd_Ospf

    // Server Info.
    Server Rcmd_Server

    // Node Info.
    Node Rcmd_Node

    // Operational data for ISIS.
    Isis Rcmd_Isis

    // Memory Info.
    Memory Rcmd_Memory

    // LDP data.
    Ldp Rcmd_Ldp

    // Interface data.
    Intf Rcmd_Intf

    // Process information.
    Process Rcmd_Process
}

func (rcmd *Rcmd) GetEntityData() *types.CommonEntityData {
    rcmd.EntityData.YFilter = rcmd.YFilter
    rcmd.EntityData.YangName = "rcmd"
    rcmd.EntityData.BundleName = "cisco_ios_xr"
    rcmd.EntityData.ParentYangName = "Cisco-IOS-XR-infra-rcmd-oper"
    rcmd.EntityData.SegmentPath = "Cisco-IOS-XR-infra-rcmd-oper:rcmd"
    rcmd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rcmd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rcmd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rcmd.EntityData.Children = types.NewOrderedMap()
    rcmd.EntityData.Children.Append("ospf", types.YChild{"Ospf", &rcmd.Ospf})
    rcmd.EntityData.Children.Append("server", types.YChild{"Server", &rcmd.Server})
    rcmd.EntityData.Children.Append("node", types.YChild{"Node", &rcmd.Node})
    rcmd.EntityData.Children.Append("isis", types.YChild{"Isis", &rcmd.Isis})
    rcmd.EntityData.Children.Append("memory", types.YChild{"Memory", &rcmd.Memory})
    rcmd.EntityData.Children.Append("ldp", types.YChild{"Ldp", &rcmd.Ldp})
    rcmd.EntityData.Children.Append("intf", types.YChild{"Intf", &rcmd.Intf})
    rcmd.EntityData.Children.Append("process", types.YChild{"Process", &rcmd.Process})
    rcmd.EntityData.Leafs = types.NewOrderedMap()

    rcmd.EntityData.YListKeys = []string {}

    return &(rcmd.EntityData)
}

// Rcmd_Ospf
// Operational data for OSPF
type Rcmd_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data.
    Instances Rcmd_Ospf_Instances
}

func (ospf *Rcmd_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "rcmd"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = types.NewOrderedMap()
    ospf.EntityData.Children.Append("instances", types.YChild{"Instances", &ospf.Instances})
    ospf.EntityData.Leafs = types.NewOrderedMap()

    ospf.EntityData.YListKeys = []string {}

    return &(ospf.EntityData)
}

// Rcmd_Ospf_Instances
// Operational data
type Rcmd_Ospf_Instances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a particular instance. The type is slice of
    // Rcmd_Ospf_Instances_Instance.
    Instance []*Rcmd_Ospf_Instances_Instance
}

func (instances *Rcmd_Ospf_Instances) GetEntityData() *types.CommonEntityData {
    instances.EntityData.YFilter = instances.YFilter
    instances.EntityData.YangName = "instances"
    instances.EntityData.BundleName = "cisco_ios_xr"
    instances.EntityData.ParentYangName = "ospf"
    instances.EntityData.SegmentPath = "instances"
    instances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instances.EntityData.Children = types.NewOrderedMap()
    instances.EntityData.Children.Append("instance", types.YChild{"Instance", nil})
    for i := range instances.Instance {
        instances.EntityData.Children.Append(types.GetSegmentPath(instances.Instance[i]), types.YChild{"Instance", instances.Instance[i]})
    }
    instances.EntityData.Leafs = types.NewOrderedMap()

    instances.EntityData.YListKeys = []string {}

    return &(instances.EntityData)
}

// Rcmd_Ospf_Instances_Instance
// Operational data for a particular instance
type Rcmd_Ospf_Instances_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Operational data for a particular instance. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // OSPF IP-FRR events summary data.
    IpfrrEventSummaries Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries

    // OSPF Prefix events summary data.
    PrefixEventStatistics Rcmd_Ospf_Instances_Instance_PrefixEventStatistics

    // OSPF SPF run summary data.
    SpfRunSummaries Rcmd_Ospf_Instances_Instance_SpfRunSummaries

    // OSPF IP-FRR Event offline data.
    IpfrrEventOfflines Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines

    // OSPF SPF run offline data.
    SpfRunOfflines Rcmd_Ospf_Instances_Instance_SpfRunOfflines

    // OSPF Summary-External Prefix events summary data.
    SummaryExternalEventSummaries Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries

    // OSPF Prefix events summary data.
    PrefixEventSummaries Rcmd_Ospf_Instances_Instance_PrefixEventSummaries

    // OSPF Summary-External Prefix events offline data.
    SummaryExternalEventOfflines Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines

    // OSPF Prefix events offline data.
    PrefixEventOfflines Rcmd_Ospf_Instances_Instance_PrefixEventOfflines

    // Summary-External prefix monitoring statistics.
    SummaryExternalEventStatistics Rcmd_Ospf_Instances_Instance_SummaryExternalEventStatistics
}

func (instance *Rcmd_Ospf_Instances_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "instances"
    instance.EntityData.SegmentPath = "instance" + types.AddKeyToken(instance.InstanceName, "instance-name")
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Children.Append("ipfrr-event-summaries", types.YChild{"IpfrrEventSummaries", &instance.IpfrrEventSummaries})
    instance.EntityData.Children.Append("prefix-event-statistics", types.YChild{"PrefixEventStatistics", &instance.PrefixEventStatistics})
    instance.EntityData.Children.Append("spf-run-summaries", types.YChild{"SpfRunSummaries", &instance.SpfRunSummaries})
    instance.EntityData.Children.Append("ipfrr-event-offlines", types.YChild{"IpfrrEventOfflines", &instance.IpfrrEventOfflines})
    instance.EntityData.Children.Append("spf-run-offlines", types.YChild{"SpfRunOfflines", &instance.SpfRunOfflines})
    instance.EntityData.Children.Append("summary-external-event-summaries", types.YChild{"SummaryExternalEventSummaries", &instance.SummaryExternalEventSummaries})
    instance.EntityData.Children.Append("prefix-event-summaries", types.YChild{"PrefixEventSummaries", &instance.PrefixEventSummaries})
    instance.EntityData.Children.Append("summary-external-event-offlines", types.YChild{"SummaryExternalEventOfflines", &instance.SummaryExternalEventOfflines})
    instance.EntityData.Children.Append("prefix-event-offlines", types.YChild{"PrefixEventOfflines", &instance.PrefixEventOfflines})
    instance.EntityData.Children.Append("summary-external-event-statistics", types.YChild{"SummaryExternalEventStatistics", &instance.SummaryExternalEventStatistics})
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", instance.InstanceName})

    instance.EntityData.YListKeys = []string {"InstanceName"}

    return &(instance.EntityData)
}

// Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries
// OSPF IP-FRR events summary data
type Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP-FRR Event data. The type is slice of
    // Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary.
    IpfrrEventSummary []*Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary
}

func (ipfrrEventSummaries *Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries) GetEntityData() *types.CommonEntityData {
    ipfrrEventSummaries.EntityData.YFilter = ipfrrEventSummaries.YFilter
    ipfrrEventSummaries.EntityData.YangName = "ipfrr-event-summaries"
    ipfrrEventSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipfrrEventSummaries.EntityData.ParentYangName = "instance"
    ipfrrEventSummaries.EntityData.SegmentPath = "ipfrr-event-summaries"
    ipfrrEventSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipfrrEventSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipfrrEventSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipfrrEventSummaries.EntityData.Children = types.NewOrderedMap()
    ipfrrEventSummaries.EntityData.Children.Append("ipfrr-event-summary", types.YChild{"IpfrrEventSummary", nil})
    for i := range ipfrrEventSummaries.IpfrrEventSummary {
        ipfrrEventSummaries.EntityData.Children.Append(types.GetSegmentPath(ipfrrEventSummaries.IpfrrEventSummary[i]), types.YChild{"IpfrrEventSummary", ipfrrEventSummaries.IpfrrEventSummary[i]})
    }
    ipfrrEventSummaries.EntityData.Leafs = types.NewOrderedMap()

    ipfrrEventSummaries.EntityData.YListKeys = []string {}

    return &(ipfrrEventSummaries.EntityData)
}

// Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary
// IP-FRR Event data
type Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific IP-FRR Event. The type is interface{}
    // with range: 1..4294967295.
    EventId interface{}

    // IP-Frr Event ID. The type is interface{} with range: 0..4294967295.
    EventIdXr interface{}

    // Trigger time  (eg: Apr 24 13:16:04.961). The type is string.
    TriggerTime interface{}

    // IP-Frr Triggered reference SPF Run Number. The type is interface{} with
    // range: 0..4294967295.
    TriggerSpfRun interface{}

    // Waiting Time (in milliseconds). The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    WaitTime interface{}

    // Start Time offset from trigger time (in milliseconds). The type is string.
    // Units are millisecond.
    StartTimeOffset interface{}

    // Duration for the calculation (in milliseconds). The type is string. Units
    // are millisecond.
    Duration interface{}

    // IP-Frr Completed reference SPF Run Number. The type is interface{} with
    // range: 0..4294967295.
    CompletedSpfRun interface{}

    // Cumulative Number of Routes for all priorities. The type is interface{}
    // with range: 0..4294967295.
    TotalRoutes interface{}

    // Cumulative Number of Fully Protected Routes. The type is interface{} with
    // range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Cumulative Number of Partially Protected Routes. The type is interface{}
    // with range: 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage for all priorities. The type is string. Units are
    // percentage.
    Coverage interface{}

    // IP-Frr Statistics categorized by priority. The type is slice of
    // Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_IpfrrStatistic.
    IpfrrStatistic []*Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_IpfrrStatistic

    // Remote Node Information. The type is slice of
    // Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode.
    RemoteNode []*Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode
}

func (ipfrrEventSummary *Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary) GetEntityData() *types.CommonEntityData {
    ipfrrEventSummary.EntityData.YFilter = ipfrrEventSummary.YFilter
    ipfrrEventSummary.EntityData.YangName = "ipfrr-event-summary"
    ipfrrEventSummary.EntityData.BundleName = "cisco_ios_xr"
    ipfrrEventSummary.EntityData.ParentYangName = "ipfrr-event-summaries"
    ipfrrEventSummary.EntityData.SegmentPath = "ipfrr-event-summary" + types.AddKeyToken(ipfrrEventSummary.EventId, "event-id")
    ipfrrEventSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipfrrEventSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipfrrEventSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipfrrEventSummary.EntityData.Children = types.NewOrderedMap()
    ipfrrEventSummary.EntityData.Children.Append("ipfrr-statistic", types.YChild{"IpfrrStatistic", nil})
    for i := range ipfrrEventSummary.IpfrrStatistic {
        ipfrrEventSummary.EntityData.Children.Append(types.GetSegmentPath(ipfrrEventSummary.IpfrrStatistic[i]), types.YChild{"IpfrrStatistic", ipfrrEventSummary.IpfrrStatistic[i]})
    }
    ipfrrEventSummary.EntityData.Children.Append("remote-node", types.YChild{"RemoteNode", nil})
    for i := range ipfrrEventSummary.RemoteNode {
        ipfrrEventSummary.EntityData.Children.Append(types.GetSegmentPath(ipfrrEventSummary.RemoteNode[i]), types.YChild{"RemoteNode", ipfrrEventSummary.RemoteNode[i]})
    }
    ipfrrEventSummary.EntityData.Leafs = types.NewOrderedMap()
    ipfrrEventSummary.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", ipfrrEventSummary.EventId})
    ipfrrEventSummary.EntityData.Leafs.Append("event-id-xr", types.YLeaf{"EventIdXr", ipfrrEventSummary.EventIdXr})
    ipfrrEventSummary.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", ipfrrEventSummary.TriggerTime})
    ipfrrEventSummary.EntityData.Leafs.Append("trigger-spf-run", types.YLeaf{"TriggerSpfRun", ipfrrEventSummary.TriggerSpfRun})
    ipfrrEventSummary.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", ipfrrEventSummary.WaitTime})
    ipfrrEventSummary.EntityData.Leafs.Append("start-time-offset", types.YLeaf{"StartTimeOffset", ipfrrEventSummary.StartTimeOffset})
    ipfrrEventSummary.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ipfrrEventSummary.Duration})
    ipfrrEventSummary.EntityData.Leafs.Append("completed-spf-run", types.YLeaf{"CompletedSpfRun", ipfrrEventSummary.CompletedSpfRun})
    ipfrrEventSummary.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", ipfrrEventSummary.TotalRoutes})
    ipfrrEventSummary.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", ipfrrEventSummary.FullyProtectedRoutes})
    ipfrrEventSummary.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", ipfrrEventSummary.PartiallyProtectedRoutes})
    ipfrrEventSummary.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", ipfrrEventSummary.Coverage})

    ipfrrEventSummary.EntityData.YListKeys = []string {"EventId"}

    return &(ipfrrEventSummary.EntityData)
}

// Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_IpfrrStatistic
// IP-Frr Statistics categorized by priority
type Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_IpfrrStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Priority. The type is RcmdPriorityLevel.
    Priority interface{}

    // Total Number of Routes. The type is interface{} with range: 0..4294967295.
    TotalRoutes interface{}

    // Fully Protected Routes. The type is interface{} with range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Partially Protected Routes. The type is interface{} with range:
    // 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage. The type is string. Units are percentage.
    Coverage interface{}

    // Local LFA Coverage in percentage. The type is string. Units are percentage.
    LocalLfaCoverage interface{}

    // Remote LFA Coverage in percentage. The type is string. Units are
    // percentage.
    RemoteLfaCoverage interface{}

    // Covearge is below Configured Threshold. The type is bool.
    BelowThreshold interface{}
}

func (ipfrrStatistic *Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_IpfrrStatistic) GetEntityData() *types.CommonEntityData {
    ipfrrStatistic.EntityData.YFilter = ipfrrStatistic.YFilter
    ipfrrStatistic.EntityData.YangName = "ipfrr-statistic"
    ipfrrStatistic.EntityData.BundleName = "cisco_ios_xr"
    ipfrrStatistic.EntityData.ParentYangName = "ipfrr-event-summary"
    ipfrrStatistic.EntityData.SegmentPath = "ipfrr-statistic"
    ipfrrStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipfrrStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipfrrStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipfrrStatistic.EntityData.Children = types.NewOrderedMap()
    ipfrrStatistic.EntityData.Leafs = types.NewOrderedMap()
    ipfrrStatistic.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", ipfrrStatistic.Priority})
    ipfrrStatistic.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", ipfrrStatistic.TotalRoutes})
    ipfrrStatistic.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", ipfrrStatistic.FullyProtectedRoutes})
    ipfrrStatistic.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", ipfrrStatistic.PartiallyProtectedRoutes})
    ipfrrStatistic.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", ipfrrStatistic.Coverage})
    ipfrrStatistic.EntityData.Leafs.Append("local-lfa-coverage", types.YLeaf{"LocalLfaCoverage", ipfrrStatistic.LocalLfaCoverage})
    ipfrrStatistic.EntityData.Leafs.Append("remote-lfa-coverage", types.YLeaf{"RemoteLfaCoverage", ipfrrStatistic.RemoteLfaCoverage})
    ipfrrStatistic.EntityData.Leafs.Append("below-threshold", types.YLeaf{"BelowThreshold", ipfrrStatistic.BelowThreshold})

    ipfrrStatistic.EntityData.YListKeys = []string {}

    return &(ipfrrStatistic.EntityData)
}

// Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode
// Remote Node Information
type Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote-LFA Node ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteNodeId interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Number of paths protected by this Remote Node. The type is interface{} with
    // range: 0..4294967295.
    PathCount interface{}

    // Inuse time of the Remote Node (eg: Apr 24 13:16 :04.961). The type is
    // string.
    InUseTime interface{}

    // Protected Primary Paths. The type is slice of
    // Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode_PrimaryPath.
    PrimaryPath []*Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode_PrimaryPath
}

func (remoteNode *Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode) GetEntityData() *types.CommonEntityData {
    remoteNode.EntityData.YFilter = remoteNode.YFilter
    remoteNode.EntityData.YangName = "remote-node"
    remoteNode.EntityData.BundleName = "cisco_ios_xr"
    remoteNode.EntityData.ParentYangName = "ipfrr-event-summary"
    remoteNode.EntityData.SegmentPath = "remote-node"
    remoteNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNode.EntityData.Children = types.NewOrderedMap()
    remoteNode.EntityData.Children.Append("primary-path", types.YChild{"PrimaryPath", nil})
    for i := range remoteNode.PrimaryPath {
        remoteNode.EntityData.Children.Append(types.GetSegmentPath(remoteNode.PrimaryPath[i]), types.YChild{"PrimaryPath", remoteNode.PrimaryPath[i]})
    }
    remoteNode.EntityData.Leafs = types.NewOrderedMap()
    remoteNode.EntityData.Leafs.Append("remote-node-id", types.YLeaf{"RemoteNodeId", remoteNode.RemoteNodeId})
    remoteNode.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", remoteNode.InterfaceName})
    remoteNode.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", remoteNode.NeighbourAddress})
    remoteNode.EntityData.Leafs.Append("path-count", types.YLeaf{"PathCount", remoteNode.PathCount})
    remoteNode.EntityData.Leafs.Append("in-use-time", types.YLeaf{"InUseTime", remoteNode.InUseTime})

    remoteNode.EntityData.YListKeys = []string {}

    return &(remoteNode.EntityData)
}

// Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode_PrimaryPath
// Protected Primary Paths
type Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode_PrimaryPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}
}

func (primaryPath *Rcmd_Ospf_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode_PrimaryPath) GetEntityData() *types.CommonEntityData {
    primaryPath.EntityData.YFilter = primaryPath.YFilter
    primaryPath.EntityData.YangName = "primary-path"
    primaryPath.EntityData.BundleName = "cisco_ios_xr"
    primaryPath.EntityData.ParentYangName = "remote-node"
    primaryPath.EntityData.SegmentPath = "primary-path"
    primaryPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryPath.EntityData.Children = types.NewOrderedMap()
    primaryPath.EntityData.Leafs = types.NewOrderedMap()
    primaryPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", primaryPath.InterfaceName})
    primaryPath.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", primaryPath.NeighbourAddress})

    primaryPath.EntityData.YListKeys = []string {}

    return &(primaryPath.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventStatistics
// OSPF Prefix events summary data
type Rcmd_Ospf_Instances_Instance_PrefixEventStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix Event statistics. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventStatistics_PrefixEventStatistic.
    PrefixEventStatistic []*Rcmd_Ospf_Instances_Instance_PrefixEventStatistics_PrefixEventStatistic
}

func (prefixEventStatistics *Rcmd_Ospf_Instances_Instance_PrefixEventStatistics) GetEntityData() *types.CommonEntityData {
    prefixEventStatistics.EntityData.YFilter = prefixEventStatistics.YFilter
    prefixEventStatistics.EntityData.YangName = "prefix-event-statistics"
    prefixEventStatistics.EntityData.BundleName = "cisco_ios_xr"
    prefixEventStatistics.EntityData.ParentYangName = "instance"
    prefixEventStatistics.EntityData.SegmentPath = "prefix-event-statistics"
    prefixEventStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixEventStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixEventStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixEventStatistics.EntityData.Children = types.NewOrderedMap()
    prefixEventStatistics.EntityData.Children.Append("prefix-event-statistic", types.YChild{"PrefixEventStatistic", nil})
    for i := range prefixEventStatistics.PrefixEventStatistic {
        prefixEventStatistics.EntityData.Children.Append(types.GetSegmentPath(prefixEventStatistics.PrefixEventStatistic[i]), types.YChild{"PrefixEventStatistic", prefixEventStatistics.PrefixEventStatistic[i]})
    }
    prefixEventStatistics.EntityData.Leafs = types.NewOrderedMap()

    prefixEventStatistics.EntityData.YListKeys = []string {}

    return &(prefixEventStatistics.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventStatistics_PrefixEventStatistic
// Prefix Event statistics
type Rcmd_Ospf_Instances_Instance_PrefixEventStatistics_PrefixEventStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Events with Prefix. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    PrefixInfo interface{}

    // Prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLenth interface{}

    // Last event trigger time. The type is string.
    LastEventTime interface{}

    // Last event processed priority. The type is RcmdPriorityLevel.
    LastPriority interface{}

    // Last event Route Type. The type is RcmdShowRoute.
    LastRouteType interface{}

    // Last event Add/Delete. The type is RcmdChange.
    LastChangeType interface{}

    // Last Known Cost. The type is interface{} with range: 0..4294967295.
    LastCost interface{}

    // No. of times processed under Critical Priority. The type is interface{}
    // with range: 0..4294967295.
    CriticalPriority interface{}

    // No. of times processed under High Priority. The type is interface{} with
    // range: 0..4294967295.
    HighPriority interface{}

    // No. of times processed under Medium Priority. The type is interface{} with
    // range: 0..4294967295.
    MediumPriority interface{}

    // No. of times processed under Low Priority. The type is interface{} with
    // range: 0..4294967295.
    LowPriority interface{}

    // No. of times route gets Added. The type is interface{} with range:
    // 0..4294967295.
    AddCount interface{}

    // No. of times route gets Deleted. The type is interface{} with range:
    // 0..4294967295.
    ModifyCount interface{}

    // No. of times route gets Deleted. The type is interface{} with range:
    // 0..4294967295.
    DeleteCount interface{}

    // No. of times threshold got exceeded. The type is interface{} with range:
    // 0..4294967295.
    ThresholdExceedCount interface{}
}

func (prefixEventStatistic *Rcmd_Ospf_Instances_Instance_PrefixEventStatistics_PrefixEventStatistic) GetEntityData() *types.CommonEntityData {
    prefixEventStatistic.EntityData.YFilter = prefixEventStatistic.YFilter
    prefixEventStatistic.EntityData.YangName = "prefix-event-statistic"
    prefixEventStatistic.EntityData.BundleName = "cisco_ios_xr"
    prefixEventStatistic.EntityData.ParentYangName = "prefix-event-statistics"
    prefixEventStatistic.EntityData.SegmentPath = "prefix-event-statistic" + types.AddKeyToken(prefixEventStatistic.PrefixInfo, "prefix-info")
    prefixEventStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixEventStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixEventStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixEventStatistic.EntityData.Children = types.NewOrderedMap()
    prefixEventStatistic.EntityData.Leafs = types.NewOrderedMap()
    prefixEventStatistic.EntityData.Leafs.Append("prefix-info", types.YLeaf{"PrefixInfo", prefixEventStatistic.PrefixInfo})
    prefixEventStatistic.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", prefixEventStatistic.Prefix})
    prefixEventStatistic.EntityData.Leafs.Append("prefix-lenth", types.YLeaf{"PrefixLenth", prefixEventStatistic.PrefixLenth})
    prefixEventStatistic.EntityData.Leafs.Append("last-event-time", types.YLeaf{"LastEventTime", prefixEventStatistic.LastEventTime})
    prefixEventStatistic.EntityData.Leafs.Append("last-priority", types.YLeaf{"LastPriority", prefixEventStatistic.LastPriority})
    prefixEventStatistic.EntityData.Leafs.Append("last-route-type", types.YLeaf{"LastRouteType", prefixEventStatistic.LastRouteType})
    prefixEventStatistic.EntityData.Leafs.Append("last-change-type", types.YLeaf{"LastChangeType", prefixEventStatistic.LastChangeType})
    prefixEventStatistic.EntityData.Leafs.Append("last-cost", types.YLeaf{"LastCost", prefixEventStatistic.LastCost})
    prefixEventStatistic.EntityData.Leafs.Append("critical-priority", types.YLeaf{"CriticalPriority", prefixEventStatistic.CriticalPriority})
    prefixEventStatistic.EntityData.Leafs.Append("high-priority", types.YLeaf{"HighPriority", prefixEventStatistic.HighPriority})
    prefixEventStatistic.EntityData.Leafs.Append("medium-priority", types.YLeaf{"MediumPriority", prefixEventStatistic.MediumPriority})
    prefixEventStatistic.EntityData.Leafs.Append("low-priority", types.YLeaf{"LowPriority", prefixEventStatistic.LowPriority})
    prefixEventStatistic.EntityData.Leafs.Append("add-count", types.YLeaf{"AddCount", prefixEventStatistic.AddCount})
    prefixEventStatistic.EntityData.Leafs.Append("modify-count", types.YLeaf{"ModifyCount", prefixEventStatistic.ModifyCount})
    prefixEventStatistic.EntityData.Leafs.Append("delete-count", types.YLeaf{"DeleteCount", prefixEventStatistic.DeleteCount})
    prefixEventStatistic.EntityData.Leafs.Append("threshold-exceed-count", types.YLeaf{"ThresholdExceedCount", prefixEventStatistic.ThresholdExceedCount})

    prefixEventStatistic.EntityData.YListKeys = []string {"PrefixInfo"}

    return &(prefixEventStatistic.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries
// OSPF SPF run summary data
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SPF Event data. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary.
    SpfRunSummary []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary
}

func (spfRunSummaries *Rcmd_Ospf_Instances_Instance_SpfRunSummaries) GetEntityData() *types.CommonEntityData {
    spfRunSummaries.EntityData.YFilter = spfRunSummaries.YFilter
    spfRunSummaries.EntityData.YangName = "spf-run-summaries"
    spfRunSummaries.EntityData.BundleName = "cisco_ios_xr"
    spfRunSummaries.EntityData.ParentYangName = "instance"
    spfRunSummaries.EntityData.SegmentPath = "spf-run-summaries"
    spfRunSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfRunSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfRunSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfRunSummaries.EntityData.Children = types.NewOrderedMap()
    spfRunSummaries.EntityData.Children.Append("spf-run-summary", types.YChild{"SpfRunSummary", nil})
    for i := range spfRunSummaries.SpfRunSummary {
        spfRunSummaries.EntityData.Children.Append(types.GetSegmentPath(spfRunSummaries.SpfRunSummary[i]), types.YChild{"SpfRunSummary", spfRunSummaries.SpfRunSummary[i]})
    }
    spfRunSummaries.EntityData.Leafs = types.NewOrderedMap()

    spfRunSummaries.EntityData.YListKeys = []string {}

    return &(spfRunSummaries.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary
// SPF Event data
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific SPF run. The type is interface{} with
    // range: 1..4294967295.
    SpfRunNumber interface{}

    // SPF summary information.
    SpfSummary Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary

    // List of Dijkstra runs. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun.
    DijkstraRun []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun

    // Inter-area & external calculation information. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal.
    InterAreaAndExternal []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal
}

func (spfRunSummary *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary) GetEntityData() *types.CommonEntityData {
    spfRunSummary.EntityData.YFilter = spfRunSummary.YFilter
    spfRunSummary.EntityData.YangName = "spf-run-summary"
    spfRunSummary.EntityData.BundleName = "cisco_ios_xr"
    spfRunSummary.EntityData.ParentYangName = "spf-run-summaries"
    spfRunSummary.EntityData.SegmentPath = "spf-run-summary" + types.AddKeyToken(spfRunSummary.SpfRunNumber, "spf-run-number")
    spfRunSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfRunSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfRunSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfRunSummary.EntityData.Children = types.NewOrderedMap()
    spfRunSummary.EntityData.Children.Append("spf-summary", types.YChild{"SpfSummary", &spfRunSummary.SpfSummary})
    spfRunSummary.EntityData.Children.Append("dijkstra-run", types.YChild{"DijkstraRun", nil})
    for i := range spfRunSummary.DijkstraRun {
        spfRunSummary.EntityData.Children.Append(types.GetSegmentPath(spfRunSummary.DijkstraRun[i]), types.YChild{"DijkstraRun", spfRunSummary.DijkstraRun[i]})
    }
    spfRunSummary.EntityData.Children.Append("inter-area-and-external", types.YChild{"InterAreaAndExternal", nil})
    for i := range spfRunSummary.InterAreaAndExternal {
        spfRunSummary.EntityData.Children.Append(types.GetSegmentPath(spfRunSummary.InterAreaAndExternal[i]), types.YChild{"InterAreaAndExternal", spfRunSummary.InterAreaAndExternal[i]})
    }
    spfRunSummary.EntityData.Leafs = types.NewOrderedMap()
    spfRunSummary.EntityData.Leafs.Append("spf-run-number", types.YLeaf{"SpfRunNumber", spfRunSummary.SpfRunNumber})

    spfRunSummary.EntityData.YListKeys = []string {"SpfRunNumber"}

    return &(spfRunSummary.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary
// SPF summary information
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SPF state. The type is RcmdSpfState.
    State interface{}

    // Whether the event has all information. The type is bool.
    IsDataComplete interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Trigger time (in hh:mm:ss.msec). The type is string.
    TriggerTime interface{}

    // Start time (offset from event trigger time in ss .msec). The type is
    // string.
    StartTime interface{}

    // Duration of complete SPF calculation (in ss .msec). The type is string.
    Duration interface{}

    // Total number of Dijkstra runs. The type is interface{} with range:
    // 0..65535.
    TotalDijkstraRuns interface{}

    // Total number of inter-area/external computation batches. The type is
    // interface{} with range: 0..65535.
    TotalInterAreaAndExternalBatches interface{}

    // Total number of Type 1/2 LSA changes processed. The type is interface{}
    // with range: 0..65535.
    TotalType12lsaChanges interface{}

    // Total number of Type 3/5/7 LSA changes processed. The type is interface{}
    // with range: 0..65535.
    TotalType357lsaChanges interface{}

    // Convergence information summary on per-priority basis. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary.
    PrioritySummary []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary
}

func (spfSummary *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary) GetEntityData() *types.CommonEntityData {
    spfSummary.EntityData.YFilter = spfSummary.YFilter
    spfSummary.EntityData.YangName = "spf-summary"
    spfSummary.EntityData.BundleName = "cisco_ios_xr"
    spfSummary.EntityData.ParentYangName = "spf-run-summary"
    spfSummary.EntityData.SegmentPath = "spf-summary"
    spfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfSummary.EntityData.Children = types.NewOrderedMap()
    spfSummary.EntityData.Children.Append("priority-summary", types.YChild{"PrioritySummary", nil})
    for i := range spfSummary.PrioritySummary {
        spfSummary.EntityData.Children.Append(types.GetSegmentPath(spfSummary.PrioritySummary[i]), types.YChild{"PrioritySummary", spfSummary.PrioritySummary[i]})
    }
    spfSummary.EntityData.Leafs = types.NewOrderedMap()
    spfSummary.EntityData.Leafs.Append("state", types.YLeaf{"State", spfSummary.State})
    spfSummary.EntityData.Leafs.Append("is-data-complete", types.YLeaf{"IsDataComplete", spfSummary.IsDataComplete})
    spfSummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", spfSummary.ThresholdExceeded})
    spfSummary.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", spfSummary.TriggerTime})
    spfSummary.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", spfSummary.StartTime})
    spfSummary.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", spfSummary.Duration})
    spfSummary.EntityData.Leafs.Append("total-dijkstra-runs", types.YLeaf{"TotalDijkstraRuns", spfSummary.TotalDijkstraRuns})
    spfSummary.EntityData.Leafs.Append("total-inter-area-and-external-batches", types.YLeaf{"TotalInterAreaAndExternalBatches", spfSummary.TotalInterAreaAndExternalBatches})
    spfSummary.EntityData.Leafs.Append("total-type12lsa-changes", types.YLeaf{"TotalType12lsaChanges", spfSummary.TotalType12lsaChanges})
    spfSummary.EntityData.Leafs.Append("total-type357lsa-changes", types.YLeaf{"TotalType357lsaChanges", spfSummary.TotalType357lsaChanges})

    spfSummary.EntityData.YListKeys = []string {}

    return &(spfSummary.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary
// Convergence information summary on per-priority
// basis
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Critical, High, Medium or Low. The type is RcmdPriorityLevel.
    Level interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Route statistics.
    RouteStatistics Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_RouteStatistics

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_MplsConvergenceTime

    // Fast Re-Route Statistics. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_FrrStatistic.
    FrrStatistic []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_FrrStatistic
}

func (prioritySummary *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary) GetEntityData() *types.CommonEntityData {
    prioritySummary.EntityData.YFilter = prioritySummary.YFilter
    prioritySummary.EntityData.YangName = "priority-summary"
    prioritySummary.EntityData.BundleName = "cisco_ios_xr"
    prioritySummary.EntityData.ParentYangName = "spf-summary"
    prioritySummary.EntityData.SegmentPath = "priority-summary"
    prioritySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prioritySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prioritySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prioritySummary.EntityData.Children = types.NewOrderedMap()
    prioritySummary.EntityData.Children.Append("route-statistics", types.YChild{"RouteStatistics", &prioritySummary.RouteStatistics})
    prioritySummary.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prioritySummary.IpConvergenceTime})
    prioritySummary.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prioritySummary.MplsConvergenceTime})
    prioritySummary.EntityData.Children.Append("frr-statistic", types.YChild{"FrrStatistic", nil})
    for i := range prioritySummary.FrrStatistic {
        prioritySummary.EntityData.Children.Append(types.GetSegmentPath(prioritySummary.FrrStatistic[i]), types.YChild{"FrrStatistic", prioritySummary.FrrStatistic[i]})
    }
    prioritySummary.EntityData.Leafs = types.NewOrderedMap()
    prioritySummary.EntityData.Leafs.Append("level", types.YLeaf{"Level", prioritySummary.Level})
    prioritySummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prioritySummary.ThresholdExceeded})

    prioritySummary.EntityData.YListKeys = []string {}

    return &(prioritySummary.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_RouteStatistics
// Route statistics
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_RouteStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Added. The type is interface{} with range: 0..4294967295.
    Adds interface{}

    // Deleted. The type is interface{} with range: 0..4294967295.
    Deletes interface{}

    // Modified. The type is interface{} with range: 0..4294967295.
    Modifies interface{}

    // Reachable. The type is interface{} with range: 0..4294967295.
    Reachables interface{}

    // Unreachable. The type is interface{} with range: 0..4294967295.
    Unreachables interface{}

    // Touched. The type is interface{} with range: 0..4294967295.
    Touches interface{}
}

func (routeStatistics *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_RouteStatistics) GetEntityData() *types.CommonEntityData {
    routeStatistics.EntityData.YFilter = routeStatistics.YFilter
    routeStatistics.EntityData.YangName = "route-statistics"
    routeStatistics.EntityData.BundleName = "cisco_ios_xr"
    routeStatistics.EntityData.ParentYangName = "priority-summary"
    routeStatistics.EntityData.SegmentPath = "route-statistics"
    routeStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeStatistics.EntityData.Children = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs.Append("adds", types.YLeaf{"Adds", routeStatistics.Adds})
    routeStatistics.EntityData.Leafs.Append("deletes", types.YLeaf{"Deletes", routeStatistics.Deletes})
    routeStatistics.EntityData.Leafs.Append("modifies", types.YLeaf{"Modifies", routeStatistics.Modifies})
    routeStatistics.EntityData.Leafs.Append("reachables", types.YLeaf{"Reachables", routeStatistics.Reachables})
    routeStatistics.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", routeStatistics.Unreachables})
    routeStatistics.EntityData.Leafs.Append("touches", types.YLeaf{"Touches", routeStatistics.Touches})

    routeStatistics.EntityData.YListKeys = []string {}

    return &(routeStatistics.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "priority-summary"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "priority-summary"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_FrrStatistic
// Fast Re-Route Statistics
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_FrrStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total Number of Routes. The type is interface{} with range: 0..4294967295.
    TotalRoutes interface{}

    // Fully Protected Routes. The type is interface{} with range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Partially Protected Routes. The type is interface{} with range:
    // 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage. The type is string. Units are percentage.
    Coverage interface{}
}

func (frrStatistic *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_FrrStatistic) GetEntityData() *types.CommonEntityData {
    frrStatistic.EntityData.YFilter = frrStatistic.YFilter
    frrStatistic.EntityData.YangName = "frr-statistic"
    frrStatistic.EntityData.BundleName = "cisco_ios_xr"
    frrStatistic.EntityData.ParentYangName = "priority-summary"
    frrStatistic.EntityData.SegmentPath = "frr-statistic"
    frrStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrStatistic.EntityData.Children = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", frrStatistic.TotalRoutes})
    frrStatistic.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", frrStatistic.FullyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", frrStatistic.PartiallyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", frrStatistic.Coverage})

    frrStatistic.EntityData.YListKeys = []string {}

    return &(frrStatistic.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun
// List of Dijkstra runs
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Area Dijkstra run number. The type is interface{} with range:
    // 0..4294967295.
    DijkstraRunNumber interface{}

    // Area ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    AreaId interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Trigger time (in hh:mm:ss.msec). The type is string.
    TriggerTime interface{}

    // Start time (offset from event trigger time in ss .msec). The type is
    // string.
    StartTime interface{}

    // Wait time (offset from event trigger time in ss .msec). The type is
    // interface{} with range: 0..4294967295.
    WaitTime interface{}

    // Duration of Dijktra calculation (in ss.msec). The type is string.
    Duration interface{}

    // LSA that triggered the Dijkstra run. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_TriggerLsa.
    TriggerLsa []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_TriggerLsa

    // Convergence information on per-priority basis. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority.
    Priority []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority

    // List of type 1/2 LSA changes processed. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_LsaProcessed.
    LsaProcessed []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_LsaProcessed
}

func (dijkstraRun *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun) GetEntityData() *types.CommonEntityData {
    dijkstraRun.EntityData.YFilter = dijkstraRun.YFilter
    dijkstraRun.EntityData.YangName = "dijkstra-run"
    dijkstraRun.EntityData.BundleName = "cisco_ios_xr"
    dijkstraRun.EntityData.ParentYangName = "spf-run-summary"
    dijkstraRun.EntityData.SegmentPath = "dijkstra-run"
    dijkstraRun.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dijkstraRun.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dijkstraRun.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dijkstraRun.EntityData.Children = types.NewOrderedMap()
    dijkstraRun.EntityData.Children.Append("trigger-lsa", types.YChild{"TriggerLsa", nil})
    for i := range dijkstraRun.TriggerLsa {
        dijkstraRun.EntityData.Children.Append(types.GetSegmentPath(dijkstraRun.TriggerLsa[i]), types.YChild{"TriggerLsa", dijkstraRun.TriggerLsa[i]})
    }
    dijkstraRun.EntityData.Children.Append("priority", types.YChild{"Priority", nil})
    for i := range dijkstraRun.Priority {
        dijkstraRun.EntityData.Children.Append(types.GetSegmentPath(dijkstraRun.Priority[i]), types.YChild{"Priority", dijkstraRun.Priority[i]})
    }
    dijkstraRun.EntityData.Children.Append("lsa-processed", types.YChild{"LsaProcessed", nil})
    for i := range dijkstraRun.LsaProcessed {
        dijkstraRun.EntityData.Children.Append(types.GetSegmentPath(dijkstraRun.LsaProcessed[i]), types.YChild{"LsaProcessed", dijkstraRun.LsaProcessed[i]})
    }
    dijkstraRun.EntityData.Leafs = types.NewOrderedMap()
    dijkstraRun.EntityData.Leafs.Append("dijkstra-run-number", types.YLeaf{"DijkstraRunNumber", dijkstraRun.DijkstraRunNumber})
    dijkstraRun.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", dijkstraRun.AreaId})
    dijkstraRun.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", dijkstraRun.ThresholdExceeded})
    dijkstraRun.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", dijkstraRun.TriggerTime})
    dijkstraRun.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", dijkstraRun.StartTime})
    dijkstraRun.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", dijkstraRun.WaitTime})
    dijkstraRun.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", dijkstraRun.Duration})

    dijkstraRun.EntityData.YListKeys = []string {}

    return &(dijkstraRun.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_TriggerLsa
// LSA that triggered the Dijkstra run
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_TriggerLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (triggerLsa *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_TriggerLsa) GetEntityData() *types.CommonEntityData {
    triggerLsa.EntityData.YFilter = triggerLsa.YFilter
    triggerLsa.EntityData.YangName = "trigger-lsa"
    triggerLsa.EntityData.BundleName = "cisco_ios_xr"
    triggerLsa.EntityData.ParentYangName = "dijkstra-run"
    triggerLsa.EntityData.SegmentPath = "trigger-lsa"
    triggerLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    triggerLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    triggerLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    triggerLsa.EntityData.Children = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", triggerLsa.LsaId})
    triggerLsa.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", triggerLsa.SequenceNumber})
    triggerLsa.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", triggerLsa.LsaType})
    triggerLsa.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", triggerLsa.OriginRouterId})
    triggerLsa.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", triggerLsa.ChangeType})
    triggerLsa.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", triggerLsa.ReceptionTime})

    triggerLsa.EntityData.YListKeys = []string {}

    return &(triggerLsa.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority
// Convergence information on per-priority basis
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of the priority.
    PrioritySummary Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary

    // Convergence timeline details. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline.
    ConvergenceTimeline []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline

    // List of Leaf Networks Added. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_LeafNetworksAdded.
    LeafNetworksAdded []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_LeafNetworksAdded

    // List of Leaf Networks Deleted. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_LeafNetworksDeleted.
    LeafNetworksDeleted []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_LeafNetworksDeleted
}

func (priority *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "dijkstra-run"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Children.Append("priority-summary", types.YChild{"PrioritySummary", &priority.PrioritySummary})
    priority.EntityData.Children.Append("convergence-timeline", types.YChild{"ConvergenceTimeline", nil})
    for i := range priority.ConvergenceTimeline {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.ConvergenceTimeline[i]), types.YChild{"ConvergenceTimeline", priority.ConvergenceTimeline[i]})
    }
    priority.EntityData.Children.Append("leaf-networks-added", types.YChild{"LeafNetworksAdded", nil})
    for i := range priority.LeafNetworksAdded {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.LeafNetworksAdded[i]), types.YChild{"LeafNetworksAdded", priority.LeafNetworksAdded[i]})
    }
    priority.EntityData.Children.Append("leaf-networks-deleted", types.YChild{"LeafNetworksDeleted", nil})
    for i := range priority.LeafNetworksDeleted {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.LeafNetworksDeleted[i]), types.YChild{"LeafNetworksDeleted", priority.LeafNetworksDeleted[i]})
    }
    priority.EntityData.Leafs = types.NewOrderedMap()

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary
// Summary of the priority
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Critical, High, Medium or Low. The type is RcmdPriorityLevel.
    Level interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Route statistics.
    RouteStatistics Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_RouteStatistics

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_MplsConvergenceTime

    // Fast Re-Route Statistics. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_FrrStatistic.
    FrrStatistic []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_FrrStatistic
}

func (prioritySummary *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary) GetEntityData() *types.CommonEntityData {
    prioritySummary.EntityData.YFilter = prioritySummary.YFilter
    prioritySummary.EntityData.YangName = "priority-summary"
    prioritySummary.EntityData.BundleName = "cisco_ios_xr"
    prioritySummary.EntityData.ParentYangName = "priority"
    prioritySummary.EntityData.SegmentPath = "priority-summary"
    prioritySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prioritySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prioritySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prioritySummary.EntityData.Children = types.NewOrderedMap()
    prioritySummary.EntityData.Children.Append("route-statistics", types.YChild{"RouteStatistics", &prioritySummary.RouteStatistics})
    prioritySummary.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prioritySummary.IpConvergenceTime})
    prioritySummary.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prioritySummary.MplsConvergenceTime})
    prioritySummary.EntityData.Children.Append("frr-statistic", types.YChild{"FrrStatistic", nil})
    for i := range prioritySummary.FrrStatistic {
        prioritySummary.EntityData.Children.Append(types.GetSegmentPath(prioritySummary.FrrStatistic[i]), types.YChild{"FrrStatistic", prioritySummary.FrrStatistic[i]})
    }
    prioritySummary.EntityData.Leafs = types.NewOrderedMap()
    prioritySummary.EntityData.Leafs.Append("level", types.YLeaf{"Level", prioritySummary.Level})
    prioritySummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prioritySummary.ThresholdExceeded})

    prioritySummary.EntityData.YListKeys = []string {}

    return &(prioritySummary.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_RouteStatistics
// Route statistics
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_RouteStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Added. The type is interface{} with range: 0..4294967295.
    Adds interface{}

    // Deleted. The type is interface{} with range: 0..4294967295.
    Deletes interface{}

    // Modified. The type is interface{} with range: 0..4294967295.
    Modifies interface{}

    // Reachable. The type is interface{} with range: 0..4294967295.
    Reachables interface{}

    // Unreachable. The type is interface{} with range: 0..4294967295.
    Unreachables interface{}

    // Touched. The type is interface{} with range: 0..4294967295.
    Touches interface{}
}

func (routeStatistics *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_RouteStatistics) GetEntityData() *types.CommonEntityData {
    routeStatistics.EntityData.YFilter = routeStatistics.YFilter
    routeStatistics.EntityData.YangName = "route-statistics"
    routeStatistics.EntityData.BundleName = "cisco_ios_xr"
    routeStatistics.EntityData.ParentYangName = "priority-summary"
    routeStatistics.EntityData.SegmentPath = "route-statistics"
    routeStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeStatistics.EntityData.Children = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs.Append("adds", types.YLeaf{"Adds", routeStatistics.Adds})
    routeStatistics.EntityData.Leafs.Append("deletes", types.YLeaf{"Deletes", routeStatistics.Deletes})
    routeStatistics.EntityData.Leafs.Append("modifies", types.YLeaf{"Modifies", routeStatistics.Modifies})
    routeStatistics.EntityData.Leafs.Append("reachables", types.YLeaf{"Reachables", routeStatistics.Reachables})
    routeStatistics.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", routeStatistics.Unreachables})
    routeStatistics.EntityData.Leafs.Append("touches", types.YLeaf{"Touches", routeStatistics.Touches})

    routeStatistics.EntityData.YListKeys = []string {}

    return &(routeStatistics.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "priority-summary"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "priority-summary"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_FrrStatistic
// Fast Re-Route Statistics
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_FrrStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total Number of Routes. The type is interface{} with range: 0..4294967295.
    TotalRoutes interface{}

    // Fully Protected Routes. The type is interface{} with range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Partially Protected Routes. The type is interface{} with range:
    // 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage. The type is string. Units are percentage.
    Coverage interface{}
}

func (frrStatistic *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_PrioritySummary_FrrStatistic) GetEntityData() *types.CommonEntityData {
    frrStatistic.EntityData.YFilter = frrStatistic.YFilter
    frrStatistic.EntityData.YangName = "frr-statistic"
    frrStatistic.EntityData.BundleName = "cisco_ios_xr"
    frrStatistic.EntityData.ParentYangName = "priority-summary"
    frrStatistic.EntityData.SegmentPath = "frr-statistic"
    frrStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrStatistic.EntityData.Children = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", frrStatistic.TotalRoutes})
    frrStatistic.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", frrStatistic.FullyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", frrStatistic.PartiallyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", frrStatistic.Coverage})

    frrStatistic.EntityData.YListKeys = []string {}

    return &(frrStatistic.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline
// Convergence timeline details
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route origin (routing protocol).
    RouteOrigin Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RouteOrigin

    // Entry point of IPv4 RIB.
    RiBv4Enter Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Enter

    // Exit point from IPv4 RIB to FIBs.
    RiBv4Exit Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Exit

    // Route Redistribute point from IPv4 RIB to LDP.
    RiBv4Redistribute Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Redistribute

    // Entry point of LDP.
    LdpEnter Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LdpEnter

    // Exit point of LDP to LSD.
    LdpExit Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LdpExit

    // Entry point of LSD.
    LsdEnter Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LsdEnter

    // Exit point of LSD to FIBs.
    LsdExit Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LsdExit

    // List of Linecards' completion point for IP routes. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcIp.
    LcIp []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcIp

    // List of Linecards' completion point for MPLS labels. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcMpls.
    LcMpls []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcMpls
}

func (convergenceTimeline *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline) GetEntityData() *types.CommonEntityData {
    convergenceTimeline.EntityData.YFilter = convergenceTimeline.YFilter
    convergenceTimeline.EntityData.YangName = "convergence-timeline"
    convergenceTimeline.EntityData.BundleName = "cisco_ios_xr"
    convergenceTimeline.EntityData.ParentYangName = "priority"
    convergenceTimeline.EntityData.SegmentPath = "convergence-timeline"
    convergenceTimeline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    convergenceTimeline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    convergenceTimeline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    convergenceTimeline.EntityData.Children = types.NewOrderedMap()
    convergenceTimeline.EntityData.Children.Append("route-origin", types.YChild{"RouteOrigin", &convergenceTimeline.RouteOrigin})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-enter", types.YChild{"RiBv4Enter", &convergenceTimeline.RiBv4Enter})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-exit", types.YChild{"RiBv4Exit", &convergenceTimeline.RiBv4Exit})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-redistribute", types.YChild{"RiBv4Redistribute", &convergenceTimeline.RiBv4Redistribute})
    convergenceTimeline.EntityData.Children.Append("ldp-enter", types.YChild{"LdpEnter", &convergenceTimeline.LdpEnter})
    convergenceTimeline.EntityData.Children.Append("ldp-exit", types.YChild{"LdpExit", &convergenceTimeline.LdpExit})
    convergenceTimeline.EntityData.Children.Append("lsd-enter", types.YChild{"LsdEnter", &convergenceTimeline.LsdEnter})
    convergenceTimeline.EntityData.Children.Append("lsd-exit", types.YChild{"LsdExit", &convergenceTimeline.LsdExit})
    convergenceTimeline.EntityData.Children.Append("lc-ip", types.YChild{"LcIp", nil})
    for i := range convergenceTimeline.LcIp {
        convergenceTimeline.EntityData.Children.Append(types.GetSegmentPath(convergenceTimeline.LcIp[i]), types.YChild{"LcIp", convergenceTimeline.LcIp[i]})
    }
    convergenceTimeline.EntityData.Children.Append("lc-mpls", types.YChild{"LcMpls", nil})
    for i := range convergenceTimeline.LcMpls {
        convergenceTimeline.EntityData.Children.Append(types.GetSegmentPath(convergenceTimeline.LcMpls[i]), types.YChild{"LcMpls", convergenceTimeline.LcMpls[i]})
    }
    convergenceTimeline.EntityData.Leafs = types.NewOrderedMap()

    convergenceTimeline.EntityData.YListKeys = []string {}

    return &(convergenceTimeline.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RouteOrigin
// Route origin (routing protocol)
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RouteOrigin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (routeOrigin *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RouteOrigin) GetEntityData() *types.CommonEntityData {
    routeOrigin.EntityData.YFilter = routeOrigin.YFilter
    routeOrigin.EntityData.YangName = "route-origin"
    routeOrigin.EntityData.BundleName = "cisco_ios_xr"
    routeOrigin.EntityData.ParentYangName = "convergence-timeline"
    routeOrigin.EntityData.SegmentPath = "route-origin"
    routeOrigin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeOrigin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeOrigin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeOrigin.EntityData.Children = types.NewOrderedMap()
    routeOrigin.EntityData.Leafs = types.NewOrderedMap()
    routeOrigin.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", routeOrigin.StartTime})
    routeOrigin.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", routeOrigin.EndTime})
    routeOrigin.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", routeOrigin.Duration})

    routeOrigin.EntityData.YListKeys = []string {}

    return &(routeOrigin.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Enter
// Entry point of IPv4 RIB
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Enter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Enter *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Enter) GetEntityData() *types.CommonEntityData {
    riBv4Enter.EntityData.YFilter = riBv4Enter.YFilter
    riBv4Enter.EntityData.YangName = "ri-bv4-enter"
    riBv4Enter.EntityData.BundleName = "cisco_ios_xr"
    riBv4Enter.EntityData.ParentYangName = "convergence-timeline"
    riBv4Enter.EntityData.SegmentPath = "ri-bv4-enter"
    riBv4Enter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Enter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Enter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Enter.EntityData.Children = types.NewOrderedMap()
    riBv4Enter.EntityData.Leafs = types.NewOrderedMap()
    riBv4Enter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Enter.StartTime})
    riBv4Enter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Enter.EndTime})
    riBv4Enter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Enter.Duration})

    riBv4Enter.EntityData.YListKeys = []string {}

    return &(riBv4Enter.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Exit
// Exit point from IPv4 RIB to FIBs
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Exit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Exit *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Exit) GetEntityData() *types.CommonEntityData {
    riBv4Exit.EntityData.YFilter = riBv4Exit.YFilter
    riBv4Exit.EntityData.YangName = "ri-bv4-exit"
    riBv4Exit.EntityData.BundleName = "cisco_ios_xr"
    riBv4Exit.EntityData.ParentYangName = "convergence-timeline"
    riBv4Exit.EntityData.SegmentPath = "ri-bv4-exit"
    riBv4Exit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Exit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Exit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Exit.EntityData.Children = types.NewOrderedMap()
    riBv4Exit.EntityData.Leafs = types.NewOrderedMap()
    riBv4Exit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Exit.StartTime})
    riBv4Exit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Exit.EndTime})
    riBv4Exit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Exit.Duration})

    riBv4Exit.EntityData.YListKeys = []string {}

    return &(riBv4Exit.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Redistribute
// Route Redistribute point from IPv4 RIB to LDP
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Redistribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Redistribute *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Redistribute) GetEntityData() *types.CommonEntityData {
    riBv4Redistribute.EntityData.YFilter = riBv4Redistribute.YFilter
    riBv4Redistribute.EntityData.YangName = "ri-bv4-redistribute"
    riBv4Redistribute.EntityData.BundleName = "cisco_ios_xr"
    riBv4Redistribute.EntityData.ParentYangName = "convergence-timeline"
    riBv4Redistribute.EntityData.SegmentPath = "ri-bv4-redistribute"
    riBv4Redistribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Redistribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Redistribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Redistribute.EntityData.Children = types.NewOrderedMap()
    riBv4Redistribute.EntityData.Leafs = types.NewOrderedMap()
    riBv4Redistribute.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Redistribute.StartTime})
    riBv4Redistribute.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Redistribute.EndTime})
    riBv4Redistribute.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Redistribute.Duration})

    riBv4Redistribute.EntityData.YListKeys = []string {}

    return &(riBv4Redistribute.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LdpEnter
// Entry point of LDP
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LdpEnter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (ldpEnter *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LdpEnter) GetEntityData() *types.CommonEntityData {
    ldpEnter.EntityData.YFilter = ldpEnter.YFilter
    ldpEnter.EntityData.YangName = "ldp-enter"
    ldpEnter.EntityData.BundleName = "cisco_ios_xr"
    ldpEnter.EntityData.ParentYangName = "convergence-timeline"
    ldpEnter.EntityData.SegmentPath = "ldp-enter"
    ldpEnter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpEnter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpEnter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpEnter.EntityData.Children = types.NewOrderedMap()
    ldpEnter.EntityData.Leafs = types.NewOrderedMap()
    ldpEnter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", ldpEnter.StartTime})
    ldpEnter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", ldpEnter.EndTime})
    ldpEnter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ldpEnter.Duration})

    ldpEnter.EntityData.YListKeys = []string {}

    return &(ldpEnter.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LdpExit
// Exit point of LDP to LSD
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LdpExit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (ldpExit *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LdpExit) GetEntityData() *types.CommonEntityData {
    ldpExit.EntityData.YFilter = ldpExit.YFilter
    ldpExit.EntityData.YangName = "ldp-exit"
    ldpExit.EntityData.BundleName = "cisco_ios_xr"
    ldpExit.EntityData.ParentYangName = "convergence-timeline"
    ldpExit.EntityData.SegmentPath = "ldp-exit"
    ldpExit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpExit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpExit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpExit.EntityData.Children = types.NewOrderedMap()
    ldpExit.EntityData.Leafs = types.NewOrderedMap()
    ldpExit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", ldpExit.StartTime})
    ldpExit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", ldpExit.EndTime})
    ldpExit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ldpExit.Duration})

    ldpExit.EntityData.YListKeys = []string {}

    return &(ldpExit.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LsdEnter
// Entry point of LSD
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LsdEnter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (lsdEnter *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LsdEnter) GetEntityData() *types.CommonEntityData {
    lsdEnter.EntityData.YFilter = lsdEnter.YFilter
    lsdEnter.EntityData.YangName = "lsd-enter"
    lsdEnter.EntityData.BundleName = "cisco_ios_xr"
    lsdEnter.EntityData.ParentYangName = "convergence-timeline"
    lsdEnter.EntityData.SegmentPath = "lsd-enter"
    lsdEnter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsdEnter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsdEnter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsdEnter.EntityData.Children = types.NewOrderedMap()
    lsdEnter.EntityData.Leafs = types.NewOrderedMap()
    lsdEnter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", lsdEnter.StartTime})
    lsdEnter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", lsdEnter.EndTime})
    lsdEnter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", lsdEnter.Duration})

    lsdEnter.EntityData.YListKeys = []string {}

    return &(lsdEnter.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LsdExit
// Exit point of LSD to FIBs
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LsdExit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (lsdExit *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LsdExit) GetEntityData() *types.CommonEntityData {
    lsdExit.EntityData.YFilter = lsdExit.YFilter
    lsdExit.EntityData.YangName = "lsd-exit"
    lsdExit.EntityData.BundleName = "cisco_ios_xr"
    lsdExit.EntityData.ParentYangName = "convergence-timeline"
    lsdExit.EntityData.SegmentPath = "lsd-exit"
    lsdExit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsdExit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsdExit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsdExit.EntityData.Children = types.NewOrderedMap()
    lsdExit.EntityData.Leafs = types.NewOrderedMap()
    lsdExit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", lsdExit.StartTime})
    lsdExit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", lsdExit.EndTime})
    lsdExit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", lsdExit.Duration})

    lsdExit.EntityData.YListKeys = []string {}

    return &(lsdExit.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcIp
// List of Linecards' completion point for IP
// routes
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB.
    FibComplete Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcIp_FibComplete
}

func (lcIp *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcIp) GetEntityData() *types.CommonEntityData {
    lcIp.EntityData.YFilter = lcIp.YFilter
    lcIp.EntityData.YangName = "lc-ip"
    lcIp.EntityData.BundleName = "cisco_ios_xr"
    lcIp.EntityData.ParentYangName = "convergence-timeline"
    lcIp.EntityData.SegmentPath = "lc-ip"
    lcIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcIp.EntityData.Children = types.NewOrderedMap()
    lcIp.EntityData.Children.Append("fib-complete", types.YChild{"FibComplete", &lcIp.FibComplete})
    lcIp.EntityData.Leafs = types.NewOrderedMap()
    lcIp.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcIp.NodeName})
    lcIp.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcIp.Speed})

    lcIp.EntityData.YListKeys = []string {}

    return &(lcIp.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcIp_FibComplete
// Completion point of FIB
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcIp_FibComplete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (fibComplete *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcIp_FibComplete) GetEntityData() *types.CommonEntityData {
    fibComplete.EntityData.YFilter = fibComplete.YFilter
    fibComplete.EntityData.YangName = "fib-complete"
    fibComplete.EntityData.BundleName = "cisco_ios_xr"
    fibComplete.EntityData.ParentYangName = "lc-ip"
    fibComplete.EntityData.SegmentPath = "fib-complete"
    fibComplete.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibComplete.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibComplete.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibComplete.EntityData.Children = types.NewOrderedMap()
    fibComplete.EntityData.Leafs = types.NewOrderedMap()
    fibComplete.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", fibComplete.StartTime})
    fibComplete.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", fibComplete.EndTime})
    fibComplete.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", fibComplete.Duration})

    fibComplete.EntityData.YListKeys = []string {}

    return &(fibComplete.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcMpls
// List of Linecards' completion point for MPLS
// labels
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB.
    FibComplete Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcMpls_FibComplete
}

func (lcMpls *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcMpls) GetEntityData() *types.CommonEntityData {
    lcMpls.EntityData.YFilter = lcMpls.YFilter
    lcMpls.EntityData.YangName = "lc-mpls"
    lcMpls.EntityData.BundleName = "cisco_ios_xr"
    lcMpls.EntityData.ParentYangName = "convergence-timeline"
    lcMpls.EntityData.SegmentPath = "lc-mpls"
    lcMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcMpls.EntityData.Children = types.NewOrderedMap()
    lcMpls.EntityData.Children.Append("fib-complete", types.YChild{"FibComplete", &lcMpls.FibComplete})
    lcMpls.EntityData.Leafs = types.NewOrderedMap()
    lcMpls.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcMpls.NodeName})
    lcMpls.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcMpls.Speed})

    lcMpls.EntityData.YListKeys = []string {}

    return &(lcMpls.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcMpls_FibComplete
// Completion point of FIB
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcMpls_FibComplete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (fibComplete *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_ConvergenceTimeline_LcMpls_FibComplete) GetEntityData() *types.CommonEntityData {
    fibComplete.EntityData.YFilter = fibComplete.YFilter
    fibComplete.EntityData.YangName = "fib-complete"
    fibComplete.EntityData.BundleName = "cisco_ios_xr"
    fibComplete.EntityData.ParentYangName = "lc-mpls"
    fibComplete.EntityData.SegmentPath = "fib-complete"
    fibComplete.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibComplete.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibComplete.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibComplete.EntityData.Children = types.NewOrderedMap()
    fibComplete.EntityData.Leafs = types.NewOrderedMap()
    fibComplete.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", fibComplete.StartTime})
    fibComplete.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", fibComplete.EndTime})
    fibComplete.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", fibComplete.Duration})

    fibComplete.EntityData.YListKeys = []string {}

    return &(fibComplete.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_LeafNetworksAdded
// List of Leaf Networks Added
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_LeafNetworksAdded struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is interface{} with range: 0..255.
    NetMask interface{}
}

func (leafNetworksAdded *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_LeafNetworksAdded) GetEntityData() *types.CommonEntityData {
    leafNetworksAdded.EntityData.YFilter = leafNetworksAdded.YFilter
    leafNetworksAdded.EntityData.YangName = "leaf-networks-added"
    leafNetworksAdded.EntityData.BundleName = "cisco_ios_xr"
    leafNetworksAdded.EntityData.ParentYangName = "priority"
    leafNetworksAdded.EntityData.SegmentPath = "leaf-networks-added"
    leafNetworksAdded.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leafNetworksAdded.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leafNetworksAdded.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leafNetworksAdded.EntityData.Children = types.NewOrderedMap()
    leafNetworksAdded.EntityData.Leafs = types.NewOrderedMap()
    leafNetworksAdded.EntityData.Leafs.Append("address", types.YLeaf{"Address", leafNetworksAdded.Address})
    leafNetworksAdded.EntityData.Leafs.Append("net-mask", types.YLeaf{"NetMask", leafNetworksAdded.NetMask})

    leafNetworksAdded.EntityData.YListKeys = []string {}

    return &(leafNetworksAdded.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_LeafNetworksDeleted
// List of Leaf Networks Deleted
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_LeafNetworksDeleted struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is interface{} with range: 0..255.
    NetMask interface{}
}

func (leafNetworksDeleted *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_Priority_LeafNetworksDeleted) GetEntityData() *types.CommonEntityData {
    leafNetworksDeleted.EntityData.YFilter = leafNetworksDeleted.YFilter
    leafNetworksDeleted.EntityData.YangName = "leaf-networks-deleted"
    leafNetworksDeleted.EntityData.BundleName = "cisco_ios_xr"
    leafNetworksDeleted.EntityData.ParentYangName = "priority"
    leafNetworksDeleted.EntityData.SegmentPath = "leaf-networks-deleted"
    leafNetworksDeleted.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leafNetworksDeleted.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leafNetworksDeleted.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leafNetworksDeleted.EntityData.Children = types.NewOrderedMap()
    leafNetworksDeleted.EntityData.Leafs = types.NewOrderedMap()
    leafNetworksDeleted.EntityData.Leafs.Append("address", types.YLeaf{"Address", leafNetworksDeleted.Address})
    leafNetworksDeleted.EntityData.Leafs.Append("net-mask", types.YLeaf{"NetMask", leafNetworksDeleted.NetMask})

    leafNetworksDeleted.EntityData.YListKeys = []string {}

    return &(leafNetworksDeleted.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_LsaProcessed
// List of type 1/2 LSA changes processed
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_LsaProcessed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (lsaProcessed *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_DijkstraRun_LsaProcessed) GetEntityData() *types.CommonEntityData {
    lsaProcessed.EntityData.YFilter = lsaProcessed.YFilter
    lsaProcessed.EntityData.YangName = "lsa-processed"
    lsaProcessed.EntityData.BundleName = "cisco_ios_xr"
    lsaProcessed.EntityData.ParentYangName = "dijkstra-run"
    lsaProcessed.EntityData.SegmentPath = "lsa-processed"
    lsaProcessed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsaProcessed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsaProcessed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsaProcessed.EntityData.Children = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", lsaProcessed.LsaId})
    lsaProcessed.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", lsaProcessed.SequenceNumber})
    lsaProcessed.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", lsaProcessed.LsaType})
    lsaProcessed.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", lsaProcessed.OriginRouterId})
    lsaProcessed.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lsaProcessed.ChangeType})
    lsaProcessed.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", lsaProcessed.ReceptionTime})

    lsaProcessed.EntityData.YListKeys = []string {}

    return &(lsaProcessed.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal
// Inter-area & external calculation information
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Convergence information on a per-priority basis. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority.
    Priority []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority
}

func (interAreaAndExternal *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal) GetEntityData() *types.CommonEntityData {
    interAreaAndExternal.EntityData.YFilter = interAreaAndExternal.YFilter
    interAreaAndExternal.EntityData.YangName = "inter-area-and-external"
    interAreaAndExternal.EntityData.BundleName = "cisco_ios_xr"
    interAreaAndExternal.EntityData.ParentYangName = "spf-run-summary"
    interAreaAndExternal.EntityData.SegmentPath = "inter-area-and-external"
    interAreaAndExternal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interAreaAndExternal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interAreaAndExternal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interAreaAndExternal.EntityData.Children = types.NewOrderedMap()
    interAreaAndExternal.EntityData.Children.Append("priority", types.YChild{"Priority", nil})
    for i := range interAreaAndExternal.Priority {
        interAreaAndExternal.EntityData.Children.Append(types.GetSegmentPath(interAreaAndExternal.Priority[i]), types.YChild{"Priority", interAreaAndExternal.Priority[i]})
    }
    interAreaAndExternal.EntityData.Leafs = types.NewOrderedMap()

    interAreaAndExternal.EntityData.YListKeys = []string {}

    return &(interAreaAndExternal.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority
// Convergence information on a per-priority basis
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of the priority.
    PrioritySummary Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary

    // Convergence timeline details. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline.
    ConvergenceTimeline []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline

    // List of Leaf Networks Added. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_LeafNetworksAdded.
    LeafNetworksAdded []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_LeafNetworksAdded

    // List of Leaf Networks Deleted. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_LeafNetworksDeleted.
    LeafNetworksDeleted []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_LeafNetworksDeleted
}

func (priority *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "inter-area-and-external"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Children.Append("priority-summary", types.YChild{"PrioritySummary", &priority.PrioritySummary})
    priority.EntityData.Children.Append("convergence-timeline", types.YChild{"ConvergenceTimeline", nil})
    for i := range priority.ConvergenceTimeline {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.ConvergenceTimeline[i]), types.YChild{"ConvergenceTimeline", priority.ConvergenceTimeline[i]})
    }
    priority.EntityData.Children.Append("leaf-networks-added", types.YChild{"LeafNetworksAdded", nil})
    for i := range priority.LeafNetworksAdded {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.LeafNetworksAdded[i]), types.YChild{"LeafNetworksAdded", priority.LeafNetworksAdded[i]})
    }
    priority.EntityData.Children.Append("leaf-networks-deleted", types.YChild{"LeafNetworksDeleted", nil})
    for i := range priority.LeafNetworksDeleted {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.LeafNetworksDeleted[i]), types.YChild{"LeafNetworksDeleted", priority.LeafNetworksDeleted[i]})
    }
    priority.EntityData.Leafs = types.NewOrderedMap()

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary
// Summary of the priority
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Critical, High, Medium or Low. The type is RcmdPriorityLevel.
    Level interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Number of Type 3 LSA. The type is interface{} with range: 0..4294967295.
    Type3lsAs interface{}

    // Number of Type 4 LSA. The type is interface{} with range: 0..4294967295.
    Type4lsAs interface{}

    // Number of Type 5/7 LSA. The type is interface{} with range: 0..4294967295.
    Type57lsAs interface{}

    // Route statistics.
    RouteStatistics Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary_RouteStatistics

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary_MplsConvergenceTime
}

func (prioritySummary *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary) GetEntityData() *types.CommonEntityData {
    prioritySummary.EntityData.YFilter = prioritySummary.YFilter
    prioritySummary.EntityData.YangName = "priority-summary"
    prioritySummary.EntityData.BundleName = "cisco_ios_xr"
    prioritySummary.EntityData.ParentYangName = "priority"
    prioritySummary.EntityData.SegmentPath = "priority-summary"
    prioritySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prioritySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prioritySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prioritySummary.EntityData.Children = types.NewOrderedMap()
    prioritySummary.EntityData.Children.Append("route-statistics", types.YChild{"RouteStatistics", &prioritySummary.RouteStatistics})
    prioritySummary.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prioritySummary.IpConvergenceTime})
    prioritySummary.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prioritySummary.MplsConvergenceTime})
    prioritySummary.EntityData.Leafs = types.NewOrderedMap()
    prioritySummary.EntityData.Leafs.Append("level", types.YLeaf{"Level", prioritySummary.Level})
    prioritySummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prioritySummary.ThresholdExceeded})
    prioritySummary.EntityData.Leafs.Append("type3ls-as", types.YLeaf{"Type3lsAs", prioritySummary.Type3lsAs})
    prioritySummary.EntityData.Leafs.Append("type4ls-as", types.YLeaf{"Type4lsAs", prioritySummary.Type4lsAs})
    prioritySummary.EntityData.Leafs.Append("type57ls-as", types.YLeaf{"Type57lsAs", prioritySummary.Type57lsAs})

    prioritySummary.EntityData.YListKeys = []string {}

    return &(prioritySummary.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary_RouteStatistics
// Route statistics
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary_RouteStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Added. The type is interface{} with range: 0..4294967295.
    Adds interface{}

    // Deleted. The type is interface{} with range: 0..4294967295.
    Deletes interface{}

    // Modified. The type is interface{} with range: 0..4294967295.
    Modifies interface{}

    // Reachable. The type is interface{} with range: 0..4294967295.
    Reachables interface{}

    // Unreachable. The type is interface{} with range: 0..4294967295.
    Unreachables interface{}

    // Touched. The type is interface{} with range: 0..4294967295.
    Touches interface{}
}

func (routeStatistics *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary_RouteStatistics) GetEntityData() *types.CommonEntityData {
    routeStatistics.EntityData.YFilter = routeStatistics.YFilter
    routeStatistics.EntityData.YangName = "route-statistics"
    routeStatistics.EntityData.BundleName = "cisco_ios_xr"
    routeStatistics.EntityData.ParentYangName = "priority-summary"
    routeStatistics.EntityData.SegmentPath = "route-statistics"
    routeStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeStatistics.EntityData.Children = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs.Append("adds", types.YLeaf{"Adds", routeStatistics.Adds})
    routeStatistics.EntityData.Leafs.Append("deletes", types.YLeaf{"Deletes", routeStatistics.Deletes})
    routeStatistics.EntityData.Leafs.Append("modifies", types.YLeaf{"Modifies", routeStatistics.Modifies})
    routeStatistics.EntityData.Leafs.Append("reachables", types.YLeaf{"Reachables", routeStatistics.Reachables})
    routeStatistics.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", routeStatistics.Unreachables})
    routeStatistics.EntityData.Leafs.Append("touches", types.YLeaf{"Touches", routeStatistics.Touches})

    routeStatistics.EntityData.YListKeys = []string {}

    return &(routeStatistics.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "priority-summary"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_PrioritySummary_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "priority-summary"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline
// Convergence timeline details
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route origin (routing protocol).
    RouteOrigin Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RouteOrigin

    // Entry point of IPv4 RIB.
    RiBv4Enter Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Enter

    // Exit point from IPv4 RIB to FIBs.
    RiBv4Exit Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Exit

    // Route Redistribute point from IPv4 RIB to LDP.
    RiBv4Redistribute Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Redistribute

    // Entry point of LDP.
    LdpEnter Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpEnter

    // Exit point of LDP to LSD.
    LdpExit Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpExit

    // Entry point of LSD.
    LsdEnter Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdEnter

    // Exit point of LSD to FIBs.
    LsdExit Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdExit

    // List of Linecards' completion point for IP routes. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp.
    LcIp []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp

    // List of Linecards' completion point for MPLS labels. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls.
    LcMpls []*Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls
}

func (convergenceTimeline *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline) GetEntityData() *types.CommonEntityData {
    convergenceTimeline.EntityData.YFilter = convergenceTimeline.YFilter
    convergenceTimeline.EntityData.YangName = "convergence-timeline"
    convergenceTimeline.EntityData.BundleName = "cisco_ios_xr"
    convergenceTimeline.EntityData.ParentYangName = "priority"
    convergenceTimeline.EntityData.SegmentPath = "convergence-timeline"
    convergenceTimeline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    convergenceTimeline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    convergenceTimeline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    convergenceTimeline.EntityData.Children = types.NewOrderedMap()
    convergenceTimeline.EntityData.Children.Append("route-origin", types.YChild{"RouteOrigin", &convergenceTimeline.RouteOrigin})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-enter", types.YChild{"RiBv4Enter", &convergenceTimeline.RiBv4Enter})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-exit", types.YChild{"RiBv4Exit", &convergenceTimeline.RiBv4Exit})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-redistribute", types.YChild{"RiBv4Redistribute", &convergenceTimeline.RiBv4Redistribute})
    convergenceTimeline.EntityData.Children.Append("ldp-enter", types.YChild{"LdpEnter", &convergenceTimeline.LdpEnter})
    convergenceTimeline.EntityData.Children.Append("ldp-exit", types.YChild{"LdpExit", &convergenceTimeline.LdpExit})
    convergenceTimeline.EntityData.Children.Append("lsd-enter", types.YChild{"LsdEnter", &convergenceTimeline.LsdEnter})
    convergenceTimeline.EntityData.Children.Append("lsd-exit", types.YChild{"LsdExit", &convergenceTimeline.LsdExit})
    convergenceTimeline.EntityData.Children.Append("lc-ip", types.YChild{"LcIp", nil})
    for i := range convergenceTimeline.LcIp {
        convergenceTimeline.EntityData.Children.Append(types.GetSegmentPath(convergenceTimeline.LcIp[i]), types.YChild{"LcIp", convergenceTimeline.LcIp[i]})
    }
    convergenceTimeline.EntityData.Children.Append("lc-mpls", types.YChild{"LcMpls", nil})
    for i := range convergenceTimeline.LcMpls {
        convergenceTimeline.EntityData.Children.Append(types.GetSegmentPath(convergenceTimeline.LcMpls[i]), types.YChild{"LcMpls", convergenceTimeline.LcMpls[i]})
    }
    convergenceTimeline.EntityData.Leafs = types.NewOrderedMap()

    convergenceTimeline.EntityData.YListKeys = []string {}

    return &(convergenceTimeline.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RouteOrigin
// Route origin (routing protocol)
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RouteOrigin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (routeOrigin *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RouteOrigin) GetEntityData() *types.CommonEntityData {
    routeOrigin.EntityData.YFilter = routeOrigin.YFilter
    routeOrigin.EntityData.YangName = "route-origin"
    routeOrigin.EntityData.BundleName = "cisco_ios_xr"
    routeOrigin.EntityData.ParentYangName = "convergence-timeline"
    routeOrigin.EntityData.SegmentPath = "route-origin"
    routeOrigin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeOrigin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeOrigin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeOrigin.EntityData.Children = types.NewOrderedMap()
    routeOrigin.EntityData.Leafs = types.NewOrderedMap()
    routeOrigin.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", routeOrigin.StartTime})
    routeOrigin.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", routeOrigin.EndTime})
    routeOrigin.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", routeOrigin.Duration})

    routeOrigin.EntityData.YListKeys = []string {}

    return &(routeOrigin.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Enter
// Entry point of IPv4 RIB
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Enter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Enter *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Enter) GetEntityData() *types.CommonEntityData {
    riBv4Enter.EntityData.YFilter = riBv4Enter.YFilter
    riBv4Enter.EntityData.YangName = "ri-bv4-enter"
    riBv4Enter.EntityData.BundleName = "cisco_ios_xr"
    riBv4Enter.EntityData.ParentYangName = "convergence-timeline"
    riBv4Enter.EntityData.SegmentPath = "ri-bv4-enter"
    riBv4Enter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Enter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Enter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Enter.EntityData.Children = types.NewOrderedMap()
    riBv4Enter.EntityData.Leafs = types.NewOrderedMap()
    riBv4Enter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Enter.StartTime})
    riBv4Enter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Enter.EndTime})
    riBv4Enter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Enter.Duration})

    riBv4Enter.EntityData.YListKeys = []string {}

    return &(riBv4Enter.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Exit
// Exit point from IPv4 RIB to FIBs
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Exit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Exit *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Exit) GetEntityData() *types.CommonEntityData {
    riBv4Exit.EntityData.YFilter = riBv4Exit.YFilter
    riBv4Exit.EntityData.YangName = "ri-bv4-exit"
    riBv4Exit.EntityData.BundleName = "cisco_ios_xr"
    riBv4Exit.EntityData.ParentYangName = "convergence-timeline"
    riBv4Exit.EntityData.SegmentPath = "ri-bv4-exit"
    riBv4Exit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Exit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Exit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Exit.EntityData.Children = types.NewOrderedMap()
    riBv4Exit.EntityData.Leafs = types.NewOrderedMap()
    riBv4Exit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Exit.StartTime})
    riBv4Exit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Exit.EndTime})
    riBv4Exit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Exit.Duration})

    riBv4Exit.EntityData.YListKeys = []string {}

    return &(riBv4Exit.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Redistribute
// Route Redistribute point from IPv4 RIB to LDP
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Redistribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Redistribute *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Redistribute) GetEntityData() *types.CommonEntityData {
    riBv4Redistribute.EntityData.YFilter = riBv4Redistribute.YFilter
    riBv4Redistribute.EntityData.YangName = "ri-bv4-redistribute"
    riBv4Redistribute.EntityData.BundleName = "cisco_ios_xr"
    riBv4Redistribute.EntityData.ParentYangName = "convergence-timeline"
    riBv4Redistribute.EntityData.SegmentPath = "ri-bv4-redistribute"
    riBv4Redistribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Redistribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Redistribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Redistribute.EntityData.Children = types.NewOrderedMap()
    riBv4Redistribute.EntityData.Leafs = types.NewOrderedMap()
    riBv4Redistribute.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Redistribute.StartTime})
    riBv4Redistribute.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Redistribute.EndTime})
    riBv4Redistribute.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Redistribute.Duration})

    riBv4Redistribute.EntityData.YListKeys = []string {}

    return &(riBv4Redistribute.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpEnter
// Entry point of LDP
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpEnter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (ldpEnter *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpEnter) GetEntityData() *types.CommonEntityData {
    ldpEnter.EntityData.YFilter = ldpEnter.YFilter
    ldpEnter.EntityData.YangName = "ldp-enter"
    ldpEnter.EntityData.BundleName = "cisco_ios_xr"
    ldpEnter.EntityData.ParentYangName = "convergence-timeline"
    ldpEnter.EntityData.SegmentPath = "ldp-enter"
    ldpEnter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpEnter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpEnter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpEnter.EntityData.Children = types.NewOrderedMap()
    ldpEnter.EntityData.Leafs = types.NewOrderedMap()
    ldpEnter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", ldpEnter.StartTime})
    ldpEnter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", ldpEnter.EndTime})
    ldpEnter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ldpEnter.Duration})

    ldpEnter.EntityData.YListKeys = []string {}

    return &(ldpEnter.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpExit
// Exit point of LDP to LSD
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpExit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (ldpExit *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpExit) GetEntityData() *types.CommonEntityData {
    ldpExit.EntityData.YFilter = ldpExit.YFilter
    ldpExit.EntityData.YangName = "ldp-exit"
    ldpExit.EntityData.BundleName = "cisco_ios_xr"
    ldpExit.EntityData.ParentYangName = "convergence-timeline"
    ldpExit.EntityData.SegmentPath = "ldp-exit"
    ldpExit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpExit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpExit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpExit.EntityData.Children = types.NewOrderedMap()
    ldpExit.EntityData.Leafs = types.NewOrderedMap()
    ldpExit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", ldpExit.StartTime})
    ldpExit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", ldpExit.EndTime})
    ldpExit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ldpExit.Duration})

    ldpExit.EntityData.YListKeys = []string {}

    return &(ldpExit.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdEnter
// Entry point of LSD
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdEnter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (lsdEnter *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdEnter) GetEntityData() *types.CommonEntityData {
    lsdEnter.EntityData.YFilter = lsdEnter.YFilter
    lsdEnter.EntityData.YangName = "lsd-enter"
    lsdEnter.EntityData.BundleName = "cisco_ios_xr"
    lsdEnter.EntityData.ParentYangName = "convergence-timeline"
    lsdEnter.EntityData.SegmentPath = "lsd-enter"
    lsdEnter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsdEnter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsdEnter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsdEnter.EntityData.Children = types.NewOrderedMap()
    lsdEnter.EntityData.Leafs = types.NewOrderedMap()
    lsdEnter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", lsdEnter.StartTime})
    lsdEnter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", lsdEnter.EndTime})
    lsdEnter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", lsdEnter.Duration})

    lsdEnter.EntityData.YListKeys = []string {}

    return &(lsdEnter.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdExit
// Exit point of LSD to FIBs
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdExit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (lsdExit *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdExit) GetEntityData() *types.CommonEntityData {
    lsdExit.EntityData.YFilter = lsdExit.YFilter
    lsdExit.EntityData.YangName = "lsd-exit"
    lsdExit.EntityData.BundleName = "cisco_ios_xr"
    lsdExit.EntityData.ParentYangName = "convergence-timeline"
    lsdExit.EntityData.SegmentPath = "lsd-exit"
    lsdExit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsdExit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsdExit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsdExit.EntityData.Children = types.NewOrderedMap()
    lsdExit.EntityData.Leafs = types.NewOrderedMap()
    lsdExit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", lsdExit.StartTime})
    lsdExit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", lsdExit.EndTime})
    lsdExit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", lsdExit.Duration})

    lsdExit.EntityData.YListKeys = []string {}

    return &(lsdExit.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp
// List of Linecards' completion point for IP
// routes
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB.
    FibComplete Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp_FibComplete
}

func (lcIp *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp) GetEntityData() *types.CommonEntityData {
    lcIp.EntityData.YFilter = lcIp.YFilter
    lcIp.EntityData.YangName = "lc-ip"
    lcIp.EntityData.BundleName = "cisco_ios_xr"
    lcIp.EntityData.ParentYangName = "convergence-timeline"
    lcIp.EntityData.SegmentPath = "lc-ip"
    lcIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcIp.EntityData.Children = types.NewOrderedMap()
    lcIp.EntityData.Children.Append("fib-complete", types.YChild{"FibComplete", &lcIp.FibComplete})
    lcIp.EntityData.Leafs = types.NewOrderedMap()
    lcIp.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcIp.NodeName})
    lcIp.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcIp.Speed})

    lcIp.EntityData.YListKeys = []string {}

    return &(lcIp.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp_FibComplete
// Completion point of FIB
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp_FibComplete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (fibComplete *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp_FibComplete) GetEntityData() *types.CommonEntityData {
    fibComplete.EntityData.YFilter = fibComplete.YFilter
    fibComplete.EntityData.YangName = "fib-complete"
    fibComplete.EntityData.BundleName = "cisco_ios_xr"
    fibComplete.EntityData.ParentYangName = "lc-ip"
    fibComplete.EntityData.SegmentPath = "fib-complete"
    fibComplete.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibComplete.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibComplete.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibComplete.EntityData.Children = types.NewOrderedMap()
    fibComplete.EntityData.Leafs = types.NewOrderedMap()
    fibComplete.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", fibComplete.StartTime})
    fibComplete.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", fibComplete.EndTime})
    fibComplete.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", fibComplete.Duration})

    fibComplete.EntityData.YListKeys = []string {}

    return &(fibComplete.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls
// List of Linecards' completion point for MPLS
// labels
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB.
    FibComplete Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls_FibComplete
}

func (lcMpls *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls) GetEntityData() *types.CommonEntityData {
    lcMpls.EntityData.YFilter = lcMpls.YFilter
    lcMpls.EntityData.YangName = "lc-mpls"
    lcMpls.EntityData.BundleName = "cisco_ios_xr"
    lcMpls.EntityData.ParentYangName = "convergence-timeline"
    lcMpls.EntityData.SegmentPath = "lc-mpls"
    lcMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcMpls.EntityData.Children = types.NewOrderedMap()
    lcMpls.EntityData.Children.Append("fib-complete", types.YChild{"FibComplete", &lcMpls.FibComplete})
    lcMpls.EntityData.Leafs = types.NewOrderedMap()
    lcMpls.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcMpls.NodeName})
    lcMpls.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcMpls.Speed})

    lcMpls.EntityData.YListKeys = []string {}

    return &(lcMpls.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls_FibComplete
// Completion point of FIB
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls_FibComplete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (fibComplete *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls_FibComplete) GetEntityData() *types.CommonEntityData {
    fibComplete.EntityData.YFilter = fibComplete.YFilter
    fibComplete.EntityData.YangName = "fib-complete"
    fibComplete.EntityData.BundleName = "cisco_ios_xr"
    fibComplete.EntityData.ParentYangName = "lc-mpls"
    fibComplete.EntityData.SegmentPath = "fib-complete"
    fibComplete.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibComplete.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibComplete.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibComplete.EntityData.Children = types.NewOrderedMap()
    fibComplete.EntityData.Leafs = types.NewOrderedMap()
    fibComplete.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", fibComplete.StartTime})
    fibComplete.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", fibComplete.EndTime})
    fibComplete.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", fibComplete.Duration})

    fibComplete.EntityData.YListKeys = []string {}

    return &(fibComplete.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_LeafNetworksAdded
// List of Leaf Networks Added
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_LeafNetworksAdded struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is interface{} with range: 0..255.
    NetMask interface{}
}

func (leafNetworksAdded *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_LeafNetworksAdded) GetEntityData() *types.CommonEntityData {
    leafNetworksAdded.EntityData.YFilter = leafNetworksAdded.YFilter
    leafNetworksAdded.EntityData.YangName = "leaf-networks-added"
    leafNetworksAdded.EntityData.BundleName = "cisco_ios_xr"
    leafNetworksAdded.EntityData.ParentYangName = "priority"
    leafNetworksAdded.EntityData.SegmentPath = "leaf-networks-added"
    leafNetworksAdded.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leafNetworksAdded.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leafNetworksAdded.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leafNetworksAdded.EntityData.Children = types.NewOrderedMap()
    leafNetworksAdded.EntityData.Leafs = types.NewOrderedMap()
    leafNetworksAdded.EntityData.Leafs.Append("address", types.YLeaf{"Address", leafNetworksAdded.Address})
    leafNetworksAdded.EntityData.Leafs.Append("net-mask", types.YLeaf{"NetMask", leafNetworksAdded.NetMask})

    leafNetworksAdded.EntityData.YListKeys = []string {}

    return &(leafNetworksAdded.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_LeafNetworksDeleted
// List of Leaf Networks Deleted
type Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_LeafNetworksDeleted struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is interface{} with range: 0..255.
    NetMask interface{}
}

func (leafNetworksDeleted *Rcmd_Ospf_Instances_Instance_SpfRunSummaries_SpfRunSummary_InterAreaAndExternal_Priority_LeafNetworksDeleted) GetEntityData() *types.CommonEntityData {
    leafNetworksDeleted.EntityData.YFilter = leafNetworksDeleted.YFilter
    leafNetworksDeleted.EntityData.YangName = "leaf-networks-deleted"
    leafNetworksDeleted.EntityData.BundleName = "cisco_ios_xr"
    leafNetworksDeleted.EntityData.ParentYangName = "priority"
    leafNetworksDeleted.EntityData.SegmentPath = "leaf-networks-deleted"
    leafNetworksDeleted.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leafNetworksDeleted.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leafNetworksDeleted.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leafNetworksDeleted.EntityData.Children = types.NewOrderedMap()
    leafNetworksDeleted.EntityData.Leafs = types.NewOrderedMap()
    leafNetworksDeleted.EntityData.Leafs.Append("address", types.YLeaf{"Address", leafNetworksDeleted.Address})
    leafNetworksDeleted.EntityData.Leafs.Append("net-mask", types.YLeaf{"NetMask", leafNetworksDeleted.NetMask})

    leafNetworksDeleted.EntityData.YListKeys = []string {}

    return &(leafNetworksDeleted.EntityData)
}

// Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines
// OSPF IP-FRR Event offline data
type Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Offline operational data for particular OSPF IP-FRR Event. The type is
    // slice of Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline.
    IpfrrEventOffline []*Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline
}

func (ipfrrEventOfflines *Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines) GetEntityData() *types.CommonEntityData {
    ipfrrEventOfflines.EntityData.YFilter = ipfrrEventOfflines.YFilter
    ipfrrEventOfflines.EntityData.YangName = "ipfrr-event-offlines"
    ipfrrEventOfflines.EntityData.BundleName = "cisco_ios_xr"
    ipfrrEventOfflines.EntityData.ParentYangName = "instance"
    ipfrrEventOfflines.EntityData.SegmentPath = "ipfrr-event-offlines"
    ipfrrEventOfflines.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipfrrEventOfflines.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipfrrEventOfflines.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipfrrEventOfflines.EntityData.Children = types.NewOrderedMap()
    ipfrrEventOfflines.EntityData.Children.Append("ipfrr-event-offline", types.YChild{"IpfrrEventOffline", nil})
    for i := range ipfrrEventOfflines.IpfrrEventOffline {
        ipfrrEventOfflines.EntityData.Children.Append(types.GetSegmentPath(ipfrrEventOfflines.IpfrrEventOffline[i]), types.YChild{"IpfrrEventOffline", ipfrrEventOfflines.IpfrrEventOffline[i]})
    }
    ipfrrEventOfflines.EntityData.Leafs = types.NewOrderedMap()

    ipfrrEventOfflines.EntityData.YListKeys = []string {}

    return &(ipfrrEventOfflines.EntityData)
}

// Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline
// Offline operational data for particular OSPF
// IP-FRR Event
type Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific IP-FRR Event. The type is interface{}
    // with range: 1..4294967295.
    EventId interface{}

    // IP-Frr Event ID. The type is interface{} with range: 0..4294967295.
    EventIdXr interface{}

    // Trigger time  (eg: Apr 24 13:16:04.961). The type is string.
    TriggerTime interface{}

    // IP-Frr Triggered reference SPF Run Number. The type is interface{} with
    // range: 0..4294967295.
    TriggerSpfRun interface{}

    // Waiting Time (in milliseconds). The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    WaitTime interface{}

    // Start Time offset from trigger time (in milliseconds). The type is string.
    // Units are millisecond.
    StartTimeOffset interface{}

    // Duration for the calculation (in milliseconds). The type is string. Units
    // are millisecond.
    Duration interface{}

    // IP-Frr Completed reference SPF Run Number. The type is interface{} with
    // range: 0..4294967295.
    CompletedSpfRun interface{}

    // Cumulative Number of Routes for all priorities. The type is interface{}
    // with range: 0..4294967295.
    TotalRoutes interface{}

    // Cumulative Number of Fully Protected Routes. The type is interface{} with
    // range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Cumulative Number of Partially Protected Routes. The type is interface{}
    // with range: 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage for all priorities. The type is string. Units are
    // percentage.
    Coverage interface{}

    // IP-Frr Statistics categorized by priority. The type is slice of
    // Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_IpfrrStatistic.
    IpfrrStatistic []*Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_IpfrrStatistic

    // Remote Node Information. The type is slice of
    // Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode.
    RemoteNode []*Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode
}

func (ipfrrEventOffline *Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline) GetEntityData() *types.CommonEntityData {
    ipfrrEventOffline.EntityData.YFilter = ipfrrEventOffline.YFilter
    ipfrrEventOffline.EntityData.YangName = "ipfrr-event-offline"
    ipfrrEventOffline.EntityData.BundleName = "cisco_ios_xr"
    ipfrrEventOffline.EntityData.ParentYangName = "ipfrr-event-offlines"
    ipfrrEventOffline.EntityData.SegmentPath = "ipfrr-event-offline" + types.AddKeyToken(ipfrrEventOffline.EventId, "event-id")
    ipfrrEventOffline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipfrrEventOffline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipfrrEventOffline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipfrrEventOffline.EntityData.Children = types.NewOrderedMap()
    ipfrrEventOffline.EntityData.Children.Append("ipfrr-statistic", types.YChild{"IpfrrStatistic", nil})
    for i := range ipfrrEventOffline.IpfrrStatistic {
        ipfrrEventOffline.EntityData.Children.Append(types.GetSegmentPath(ipfrrEventOffline.IpfrrStatistic[i]), types.YChild{"IpfrrStatistic", ipfrrEventOffline.IpfrrStatistic[i]})
    }
    ipfrrEventOffline.EntityData.Children.Append("remote-node", types.YChild{"RemoteNode", nil})
    for i := range ipfrrEventOffline.RemoteNode {
        ipfrrEventOffline.EntityData.Children.Append(types.GetSegmentPath(ipfrrEventOffline.RemoteNode[i]), types.YChild{"RemoteNode", ipfrrEventOffline.RemoteNode[i]})
    }
    ipfrrEventOffline.EntityData.Leafs = types.NewOrderedMap()
    ipfrrEventOffline.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", ipfrrEventOffline.EventId})
    ipfrrEventOffline.EntityData.Leafs.Append("event-id-xr", types.YLeaf{"EventIdXr", ipfrrEventOffline.EventIdXr})
    ipfrrEventOffline.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", ipfrrEventOffline.TriggerTime})
    ipfrrEventOffline.EntityData.Leafs.Append("trigger-spf-run", types.YLeaf{"TriggerSpfRun", ipfrrEventOffline.TriggerSpfRun})
    ipfrrEventOffline.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", ipfrrEventOffline.WaitTime})
    ipfrrEventOffline.EntityData.Leafs.Append("start-time-offset", types.YLeaf{"StartTimeOffset", ipfrrEventOffline.StartTimeOffset})
    ipfrrEventOffline.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ipfrrEventOffline.Duration})
    ipfrrEventOffline.EntityData.Leafs.Append("completed-spf-run", types.YLeaf{"CompletedSpfRun", ipfrrEventOffline.CompletedSpfRun})
    ipfrrEventOffline.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", ipfrrEventOffline.TotalRoutes})
    ipfrrEventOffline.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", ipfrrEventOffline.FullyProtectedRoutes})
    ipfrrEventOffline.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", ipfrrEventOffline.PartiallyProtectedRoutes})
    ipfrrEventOffline.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", ipfrrEventOffline.Coverage})

    ipfrrEventOffline.EntityData.YListKeys = []string {"EventId"}

    return &(ipfrrEventOffline.EntityData)
}

// Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_IpfrrStatistic
// IP-Frr Statistics categorized by priority
type Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_IpfrrStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Priority. The type is RcmdPriorityLevel.
    Priority interface{}

    // Total Number of Routes. The type is interface{} with range: 0..4294967295.
    TotalRoutes interface{}

    // Fully Protected Routes. The type is interface{} with range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Partially Protected Routes. The type is interface{} with range:
    // 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage. The type is string. Units are percentage.
    Coverage interface{}

    // Local LFA Coverage in percentage. The type is string. Units are percentage.
    LocalLfaCoverage interface{}

    // Remote LFA Coverage in percentage. The type is string. Units are
    // percentage.
    RemoteLfaCoverage interface{}

    // Covearge is below Configured Threshold. The type is bool.
    BelowThreshold interface{}
}

func (ipfrrStatistic *Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_IpfrrStatistic) GetEntityData() *types.CommonEntityData {
    ipfrrStatistic.EntityData.YFilter = ipfrrStatistic.YFilter
    ipfrrStatistic.EntityData.YangName = "ipfrr-statistic"
    ipfrrStatistic.EntityData.BundleName = "cisco_ios_xr"
    ipfrrStatistic.EntityData.ParentYangName = "ipfrr-event-offline"
    ipfrrStatistic.EntityData.SegmentPath = "ipfrr-statistic"
    ipfrrStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipfrrStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipfrrStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipfrrStatistic.EntityData.Children = types.NewOrderedMap()
    ipfrrStatistic.EntityData.Leafs = types.NewOrderedMap()
    ipfrrStatistic.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", ipfrrStatistic.Priority})
    ipfrrStatistic.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", ipfrrStatistic.TotalRoutes})
    ipfrrStatistic.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", ipfrrStatistic.FullyProtectedRoutes})
    ipfrrStatistic.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", ipfrrStatistic.PartiallyProtectedRoutes})
    ipfrrStatistic.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", ipfrrStatistic.Coverage})
    ipfrrStatistic.EntityData.Leafs.Append("local-lfa-coverage", types.YLeaf{"LocalLfaCoverage", ipfrrStatistic.LocalLfaCoverage})
    ipfrrStatistic.EntityData.Leafs.Append("remote-lfa-coverage", types.YLeaf{"RemoteLfaCoverage", ipfrrStatistic.RemoteLfaCoverage})
    ipfrrStatistic.EntityData.Leafs.Append("below-threshold", types.YLeaf{"BelowThreshold", ipfrrStatistic.BelowThreshold})

    ipfrrStatistic.EntityData.YListKeys = []string {}

    return &(ipfrrStatistic.EntityData)
}

// Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode
// Remote Node Information
type Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote-LFA Node ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteNodeId interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Number of paths protected by this Remote Node. The type is interface{} with
    // range: 0..4294967295.
    PathCount interface{}

    // Inuse time of the Remote Node (eg: Apr 24 13:16 :04.961). The type is
    // string.
    InUseTime interface{}

    // Protected Primary Paths. The type is slice of
    // Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode_PrimaryPath.
    PrimaryPath []*Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode_PrimaryPath
}

func (remoteNode *Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode) GetEntityData() *types.CommonEntityData {
    remoteNode.EntityData.YFilter = remoteNode.YFilter
    remoteNode.EntityData.YangName = "remote-node"
    remoteNode.EntityData.BundleName = "cisco_ios_xr"
    remoteNode.EntityData.ParentYangName = "ipfrr-event-offline"
    remoteNode.EntityData.SegmentPath = "remote-node"
    remoteNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNode.EntityData.Children = types.NewOrderedMap()
    remoteNode.EntityData.Children.Append("primary-path", types.YChild{"PrimaryPath", nil})
    for i := range remoteNode.PrimaryPath {
        remoteNode.EntityData.Children.Append(types.GetSegmentPath(remoteNode.PrimaryPath[i]), types.YChild{"PrimaryPath", remoteNode.PrimaryPath[i]})
    }
    remoteNode.EntityData.Leafs = types.NewOrderedMap()
    remoteNode.EntityData.Leafs.Append("remote-node-id", types.YLeaf{"RemoteNodeId", remoteNode.RemoteNodeId})
    remoteNode.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", remoteNode.InterfaceName})
    remoteNode.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", remoteNode.NeighbourAddress})
    remoteNode.EntityData.Leafs.Append("path-count", types.YLeaf{"PathCount", remoteNode.PathCount})
    remoteNode.EntityData.Leafs.Append("in-use-time", types.YLeaf{"InUseTime", remoteNode.InUseTime})

    remoteNode.EntityData.YListKeys = []string {}

    return &(remoteNode.EntityData)
}

// Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode_PrimaryPath
// Protected Primary Paths
type Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode_PrimaryPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}
}

func (primaryPath *Rcmd_Ospf_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode_PrimaryPath) GetEntityData() *types.CommonEntityData {
    primaryPath.EntityData.YFilter = primaryPath.YFilter
    primaryPath.EntityData.YangName = "primary-path"
    primaryPath.EntityData.BundleName = "cisco_ios_xr"
    primaryPath.EntityData.ParentYangName = "remote-node"
    primaryPath.EntityData.SegmentPath = "primary-path"
    primaryPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryPath.EntityData.Children = types.NewOrderedMap()
    primaryPath.EntityData.Leafs = types.NewOrderedMap()
    primaryPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", primaryPath.InterfaceName})
    primaryPath.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", primaryPath.NeighbourAddress})

    primaryPath.EntityData.YListKeys = []string {}

    return &(primaryPath.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines
// OSPF SPF run offline data
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Offline operational data for particular OSPF SPF run. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline.
    SpfRunOffline []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline
}

func (spfRunOfflines *Rcmd_Ospf_Instances_Instance_SpfRunOfflines) GetEntityData() *types.CommonEntityData {
    spfRunOfflines.EntityData.YFilter = spfRunOfflines.YFilter
    spfRunOfflines.EntityData.YangName = "spf-run-offlines"
    spfRunOfflines.EntityData.BundleName = "cisco_ios_xr"
    spfRunOfflines.EntityData.ParentYangName = "instance"
    spfRunOfflines.EntityData.SegmentPath = "spf-run-offlines"
    spfRunOfflines.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfRunOfflines.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfRunOfflines.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfRunOfflines.EntityData.Children = types.NewOrderedMap()
    spfRunOfflines.EntityData.Children.Append("spf-run-offline", types.YChild{"SpfRunOffline", nil})
    for i := range spfRunOfflines.SpfRunOffline {
        spfRunOfflines.EntityData.Children.Append(types.GetSegmentPath(spfRunOfflines.SpfRunOffline[i]), types.YChild{"SpfRunOffline", spfRunOfflines.SpfRunOffline[i]})
    }
    spfRunOfflines.EntityData.Leafs = types.NewOrderedMap()

    spfRunOfflines.EntityData.YListKeys = []string {}

    return &(spfRunOfflines.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline
// Offline operational data for particular OSPF
// SPF run
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific SPF run. The type is interface{} with
    // range: 1..4294967295.
    SpfRunNumber interface{}

    // SPF summary information.
    SpfSummary Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary

    // List of Dijkstra runs. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun.
    DijkstraRun []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun

    // Inter-area & external calculation information. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal.
    InterAreaAndExternal []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal
}

func (spfRunOffline *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline) GetEntityData() *types.CommonEntityData {
    spfRunOffline.EntityData.YFilter = spfRunOffline.YFilter
    spfRunOffline.EntityData.YangName = "spf-run-offline"
    spfRunOffline.EntityData.BundleName = "cisco_ios_xr"
    spfRunOffline.EntityData.ParentYangName = "spf-run-offlines"
    spfRunOffline.EntityData.SegmentPath = "spf-run-offline" + types.AddKeyToken(spfRunOffline.SpfRunNumber, "spf-run-number")
    spfRunOffline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfRunOffline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfRunOffline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfRunOffline.EntityData.Children = types.NewOrderedMap()
    spfRunOffline.EntityData.Children.Append("spf-summary", types.YChild{"SpfSummary", &spfRunOffline.SpfSummary})
    spfRunOffline.EntityData.Children.Append("dijkstra-run", types.YChild{"DijkstraRun", nil})
    for i := range spfRunOffline.DijkstraRun {
        spfRunOffline.EntityData.Children.Append(types.GetSegmentPath(spfRunOffline.DijkstraRun[i]), types.YChild{"DijkstraRun", spfRunOffline.DijkstraRun[i]})
    }
    spfRunOffline.EntityData.Children.Append("inter-area-and-external", types.YChild{"InterAreaAndExternal", nil})
    for i := range spfRunOffline.InterAreaAndExternal {
        spfRunOffline.EntityData.Children.Append(types.GetSegmentPath(spfRunOffline.InterAreaAndExternal[i]), types.YChild{"InterAreaAndExternal", spfRunOffline.InterAreaAndExternal[i]})
    }
    spfRunOffline.EntityData.Leafs = types.NewOrderedMap()
    spfRunOffline.EntityData.Leafs.Append("spf-run-number", types.YLeaf{"SpfRunNumber", spfRunOffline.SpfRunNumber})

    spfRunOffline.EntityData.YListKeys = []string {"SpfRunNumber"}

    return &(spfRunOffline.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary
// SPF summary information
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SPF state. The type is RcmdSpfState.
    State interface{}

    // Whether the event has all information. The type is bool.
    IsDataComplete interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Trigger time (in hh:mm:ss.msec). The type is string.
    TriggerTime interface{}

    // Start time (offset from event trigger time in ss .msec). The type is
    // string.
    StartTime interface{}

    // Duration of complete SPF calculation (in ss .msec). The type is string.
    Duration interface{}

    // Total number of Dijkstra runs. The type is interface{} with range:
    // 0..65535.
    TotalDijkstraRuns interface{}

    // Total number of inter-area/external computation batches. The type is
    // interface{} with range: 0..65535.
    TotalInterAreaAndExternalBatches interface{}

    // Total number of Type 1/2 LSA changes processed. The type is interface{}
    // with range: 0..65535.
    TotalType12lsaChanges interface{}

    // Total number of Type 3/5/7 LSA changes processed. The type is interface{}
    // with range: 0..65535.
    TotalType357lsaChanges interface{}

    // Convergence information summary on per-priority basis. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary.
    PrioritySummary []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary
}

func (spfSummary *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary) GetEntityData() *types.CommonEntityData {
    spfSummary.EntityData.YFilter = spfSummary.YFilter
    spfSummary.EntityData.YangName = "spf-summary"
    spfSummary.EntityData.BundleName = "cisco_ios_xr"
    spfSummary.EntityData.ParentYangName = "spf-run-offline"
    spfSummary.EntityData.SegmentPath = "spf-summary"
    spfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfSummary.EntityData.Children = types.NewOrderedMap()
    spfSummary.EntityData.Children.Append("priority-summary", types.YChild{"PrioritySummary", nil})
    for i := range spfSummary.PrioritySummary {
        spfSummary.EntityData.Children.Append(types.GetSegmentPath(spfSummary.PrioritySummary[i]), types.YChild{"PrioritySummary", spfSummary.PrioritySummary[i]})
    }
    spfSummary.EntityData.Leafs = types.NewOrderedMap()
    spfSummary.EntityData.Leafs.Append("state", types.YLeaf{"State", spfSummary.State})
    spfSummary.EntityData.Leafs.Append("is-data-complete", types.YLeaf{"IsDataComplete", spfSummary.IsDataComplete})
    spfSummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", spfSummary.ThresholdExceeded})
    spfSummary.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", spfSummary.TriggerTime})
    spfSummary.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", spfSummary.StartTime})
    spfSummary.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", spfSummary.Duration})
    spfSummary.EntityData.Leafs.Append("total-dijkstra-runs", types.YLeaf{"TotalDijkstraRuns", spfSummary.TotalDijkstraRuns})
    spfSummary.EntityData.Leafs.Append("total-inter-area-and-external-batches", types.YLeaf{"TotalInterAreaAndExternalBatches", spfSummary.TotalInterAreaAndExternalBatches})
    spfSummary.EntityData.Leafs.Append("total-type12lsa-changes", types.YLeaf{"TotalType12lsaChanges", spfSummary.TotalType12lsaChanges})
    spfSummary.EntityData.Leafs.Append("total-type357lsa-changes", types.YLeaf{"TotalType357lsaChanges", spfSummary.TotalType357lsaChanges})

    spfSummary.EntityData.YListKeys = []string {}

    return &(spfSummary.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary
// Convergence information summary on per-priority
// basis
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Critical, High, Medium or Low. The type is RcmdPriorityLevel.
    Level interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Route statistics.
    RouteStatistics Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_RouteStatistics

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_MplsConvergenceTime

    // Fast Re-Route Statistics. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_FrrStatistic.
    FrrStatistic []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_FrrStatistic
}

func (prioritySummary *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary) GetEntityData() *types.CommonEntityData {
    prioritySummary.EntityData.YFilter = prioritySummary.YFilter
    prioritySummary.EntityData.YangName = "priority-summary"
    prioritySummary.EntityData.BundleName = "cisco_ios_xr"
    prioritySummary.EntityData.ParentYangName = "spf-summary"
    prioritySummary.EntityData.SegmentPath = "priority-summary"
    prioritySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prioritySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prioritySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prioritySummary.EntityData.Children = types.NewOrderedMap()
    prioritySummary.EntityData.Children.Append("route-statistics", types.YChild{"RouteStatistics", &prioritySummary.RouteStatistics})
    prioritySummary.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prioritySummary.IpConvergenceTime})
    prioritySummary.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prioritySummary.MplsConvergenceTime})
    prioritySummary.EntityData.Children.Append("frr-statistic", types.YChild{"FrrStatistic", nil})
    for i := range prioritySummary.FrrStatistic {
        prioritySummary.EntityData.Children.Append(types.GetSegmentPath(prioritySummary.FrrStatistic[i]), types.YChild{"FrrStatistic", prioritySummary.FrrStatistic[i]})
    }
    prioritySummary.EntityData.Leafs = types.NewOrderedMap()
    prioritySummary.EntityData.Leafs.Append("level", types.YLeaf{"Level", prioritySummary.Level})
    prioritySummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prioritySummary.ThresholdExceeded})

    prioritySummary.EntityData.YListKeys = []string {}

    return &(prioritySummary.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_RouteStatistics
// Route statistics
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_RouteStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Added. The type is interface{} with range: 0..4294967295.
    Adds interface{}

    // Deleted. The type is interface{} with range: 0..4294967295.
    Deletes interface{}

    // Modified. The type is interface{} with range: 0..4294967295.
    Modifies interface{}

    // Reachable. The type is interface{} with range: 0..4294967295.
    Reachables interface{}

    // Unreachable. The type is interface{} with range: 0..4294967295.
    Unreachables interface{}

    // Touched. The type is interface{} with range: 0..4294967295.
    Touches interface{}
}

func (routeStatistics *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_RouteStatistics) GetEntityData() *types.CommonEntityData {
    routeStatistics.EntityData.YFilter = routeStatistics.YFilter
    routeStatistics.EntityData.YangName = "route-statistics"
    routeStatistics.EntityData.BundleName = "cisco_ios_xr"
    routeStatistics.EntityData.ParentYangName = "priority-summary"
    routeStatistics.EntityData.SegmentPath = "route-statistics"
    routeStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeStatistics.EntityData.Children = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs.Append("adds", types.YLeaf{"Adds", routeStatistics.Adds})
    routeStatistics.EntityData.Leafs.Append("deletes", types.YLeaf{"Deletes", routeStatistics.Deletes})
    routeStatistics.EntityData.Leafs.Append("modifies", types.YLeaf{"Modifies", routeStatistics.Modifies})
    routeStatistics.EntityData.Leafs.Append("reachables", types.YLeaf{"Reachables", routeStatistics.Reachables})
    routeStatistics.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", routeStatistics.Unreachables})
    routeStatistics.EntityData.Leafs.Append("touches", types.YLeaf{"Touches", routeStatistics.Touches})

    routeStatistics.EntityData.YListKeys = []string {}

    return &(routeStatistics.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "priority-summary"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "priority-summary"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_FrrStatistic
// Fast Re-Route Statistics
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_FrrStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total Number of Routes. The type is interface{} with range: 0..4294967295.
    TotalRoutes interface{}

    // Fully Protected Routes. The type is interface{} with range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Partially Protected Routes. The type is interface{} with range:
    // 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage. The type is string. Units are percentage.
    Coverage interface{}
}

func (frrStatistic *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_FrrStatistic) GetEntityData() *types.CommonEntityData {
    frrStatistic.EntityData.YFilter = frrStatistic.YFilter
    frrStatistic.EntityData.YangName = "frr-statistic"
    frrStatistic.EntityData.BundleName = "cisco_ios_xr"
    frrStatistic.EntityData.ParentYangName = "priority-summary"
    frrStatistic.EntityData.SegmentPath = "frr-statistic"
    frrStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrStatistic.EntityData.Children = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", frrStatistic.TotalRoutes})
    frrStatistic.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", frrStatistic.FullyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", frrStatistic.PartiallyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", frrStatistic.Coverage})

    frrStatistic.EntityData.YListKeys = []string {}

    return &(frrStatistic.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun
// List of Dijkstra runs
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Area Dijkstra run number. The type is interface{} with range:
    // 0..4294967295.
    DijkstraRunNumber interface{}

    // Area ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    AreaId interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Trigger time (in hh:mm:ss.msec). The type is string.
    TriggerTime interface{}

    // Start time (offset from event trigger time in ss .msec). The type is
    // string.
    StartTime interface{}

    // Wait time (offset from event trigger time in ss .msec). The type is
    // interface{} with range: 0..4294967295.
    WaitTime interface{}

    // Duration of Dijktra calculation (in ss.msec). The type is string.
    Duration interface{}

    // LSA that triggered the Dijkstra run. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_TriggerLsa.
    TriggerLsa []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_TriggerLsa

    // Convergence information on per-priority basis. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority.
    Priority []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority

    // List of type 1/2 LSA changes processed. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_LsaProcessed.
    LsaProcessed []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_LsaProcessed
}

func (dijkstraRun *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun) GetEntityData() *types.CommonEntityData {
    dijkstraRun.EntityData.YFilter = dijkstraRun.YFilter
    dijkstraRun.EntityData.YangName = "dijkstra-run"
    dijkstraRun.EntityData.BundleName = "cisco_ios_xr"
    dijkstraRun.EntityData.ParentYangName = "spf-run-offline"
    dijkstraRun.EntityData.SegmentPath = "dijkstra-run"
    dijkstraRun.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dijkstraRun.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dijkstraRun.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dijkstraRun.EntityData.Children = types.NewOrderedMap()
    dijkstraRun.EntityData.Children.Append("trigger-lsa", types.YChild{"TriggerLsa", nil})
    for i := range dijkstraRun.TriggerLsa {
        dijkstraRun.EntityData.Children.Append(types.GetSegmentPath(dijkstraRun.TriggerLsa[i]), types.YChild{"TriggerLsa", dijkstraRun.TriggerLsa[i]})
    }
    dijkstraRun.EntityData.Children.Append("priority", types.YChild{"Priority", nil})
    for i := range dijkstraRun.Priority {
        dijkstraRun.EntityData.Children.Append(types.GetSegmentPath(dijkstraRun.Priority[i]), types.YChild{"Priority", dijkstraRun.Priority[i]})
    }
    dijkstraRun.EntityData.Children.Append("lsa-processed", types.YChild{"LsaProcessed", nil})
    for i := range dijkstraRun.LsaProcessed {
        dijkstraRun.EntityData.Children.Append(types.GetSegmentPath(dijkstraRun.LsaProcessed[i]), types.YChild{"LsaProcessed", dijkstraRun.LsaProcessed[i]})
    }
    dijkstraRun.EntityData.Leafs = types.NewOrderedMap()
    dijkstraRun.EntityData.Leafs.Append("dijkstra-run-number", types.YLeaf{"DijkstraRunNumber", dijkstraRun.DijkstraRunNumber})
    dijkstraRun.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", dijkstraRun.AreaId})
    dijkstraRun.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", dijkstraRun.ThresholdExceeded})
    dijkstraRun.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", dijkstraRun.TriggerTime})
    dijkstraRun.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", dijkstraRun.StartTime})
    dijkstraRun.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", dijkstraRun.WaitTime})
    dijkstraRun.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", dijkstraRun.Duration})

    dijkstraRun.EntityData.YListKeys = []string {}

    return &(dijkstraRun.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_TriggerLsa
// LSA that triggered the Dijkstra run
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_TriggerLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (triggerLsa *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_TriggerLsa) GetEntityData() *types.CommonEntityData {
    triggerLsa.EntityData.YFilter = triggerLsa.YFilter
    triggerLsa.EntityData.YangName = "trigger-lsa"
    triggerLsa.EntityData.BundleName = "cisco_ios_xr"
    triggerLsa.EntityData.ParentYangName = "dijkstra-run"
    triggerLsa.EntityData.SegmentPath = "trigger-lsa"
    triggerLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    triggerLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    triggerLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    triggerLsa.EntityData.Children = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", triggerLsa.LsaId})
    triggerLsa.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", triggerLsa.SequenceNumber})
    triggerLsa.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", triggerLsa.LsaType})
    triggerLsa.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", triggerLsa.OriginRouterId})
    triggerLsa.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", triggerLsa.ChangeType})
    triggerLsa.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", triggerLsa.ReceptionTime})

    triggerLsa.EntityData.YListKeys = []string {}

    return &(triggerLsa.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority
// Convergence information on per-priority basis
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of the priority.
    PrioritySummary Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary

    // Convergence timeline details. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline.
    ConvergenceTimeline []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline

    // List of Leaf Networks Added. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_LeafNetworksAdded.
    LeafNetworksAdded []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_LeafNetworksAdded

    // List of Leaf Networks Deleted. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_LeafNetworksDeleted.
    LeafNetworksDeleted []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_LeafNetworksDeleted
}

func (priority *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "dijkstra-run"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Children.Append("priority-summary", types.YChild{"PrioritySummary", &priority.PrioritySummary})
    priority.EntityData.Children.Append("convergence-timeline", types.YChild{"ConvergenceTimeline", nil})
    for i := range priority.ConvergenceTimeline {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.ConvergenceTimeline[i]), types.YChild{"ConvergenceTimeline", priority.ConvergenceTimeline[i]})
    }
    priority.EntityData.Children.Append("leaf-networks-added", types.YChild{"LeafNetworksAdded", nil})
    for i := range priority.LeafNetworksAdded {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.LeafNetworksAdded[i]), types.YChild{"LeafNetworksAdded", priority.LeafNetworksAdded[i]})
    }
    priority.EntityData.Children.Append("leaf-networks-deleted", types.YChild{"LeafNetworksDeleted", nil})
    for i := range priority.LeafNetworksDeleted {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.LeafNetworksDeleted[i]), types.YChild{"LeafNetworksDeleted", priority.LeafNetworksDeleted[i]})
    }
    priority.EntityData.Leafs = types.NewOrderedMap()

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary
// Summary of the priority
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Critical, High, Medium or Low. The type is RcmdPriorityLevel.
    Level interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Route statistics.
    RouteStatistics Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_RouteStatistics

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_MplsConvergenceTime

    // Fast Re-Route Statistics. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_FrrStatistic.
    FrrStatistic []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_FrrStatistic
}

func (prioritySummary *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary) GetEntityData() *types.CommonEntityData {
    prioritySummary.EntityData.YFilter = prioritySummary.YFilter
    prioritySummary.EntityData.YangName = "priority-summary"
    prioritySummary.EntityData.BundleName = "cisco_ios_xr"
    prioritySummary.EntityData.ParentYangName = "priority"
    prioritySummary.EntityData.SegmentPath = "priority-summary"
    prioritySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prioritySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prioritySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prioritySummary.EntityData.Children = types.NewOrderedMap()
    prioritySummary.EntityData.Children.Append("route-statistics", types.YChild{"RouteStatistics", &prioritySummary.RouteStatistics})
    prioritySummary.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prioritySummary.IpConvergenceTime})
    prioritySummary.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prioritySummary.MplsConvergenceTime})
    prioritySummary.EntityData.Children.Append("frr-statistic", types.YChild{"FrrStatistic", nil})
    for i := range prioritySummary.FrrStatistic {
        prioritySummary.EntityData.Children.Append(types.GetSegmentPath(prioritySummary.FrrStatistic[i]), types.YChild{"FrrStatistic", prioritySummary.FrrStatistic[i]})
    }
    prioritySummary.EntityData.Leafs = types.NewOrderedMap()
    prioritySummary.EntityData.Leafs.Append("level", types.YLeaf{"Level", prioritySummary.Level})
    prioritySummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prioritySummary.ThresholdExceeded})

    prioritySummary.EntityData.YListKeys = []string {}

    return &(prioritySummary.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_RouteStatistics
// Route statistics
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_RouteStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Added. The type is interface{} with range: 0..4294967295.
    Adds interface{}

    // Deleted. The type is interface{} with range: 0..4294967295.
    Deletes interface{}

    // Modified. The type is interface{} with range: 0..4294967295.
    Modifies interface{}

    // Reachable. The type is interface{} with range: 0..4294967295.
    Reachables interface{}

    // Unreachable. The type is interface{} with range: 0..4294967295.
    Unreachables interface{}

    // Touched. The type is interface{} with range: 0..4294967295.
    Touches interface{}
}

func (routeStatistics *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_RouteStatistics) GetEntityData() *types.CommonEntityData {
    routeStatistics.EntityData.YFilter = routeStatistics.YFilter
    routeStatistics.EntityData.YangName = "route-statistics"
    routeStatistics.EntityData.BundleName = "cisco_ios_xr"
    routeStatistics.EntityData.ParentYangName = "priority-summary"
    routeStatistics.EntityData.SegmentPath = "route-statistics"
    routeStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeStatistics.EntityData.Children = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs.Append("adds", types.YLeaf{"Adds", routeStatistics.Adds})
    routeStatistics.EntityData.Leafs.Append("deletes", types.YLeaf{"Deletes", routeStatistics.Deletes})
    routeStatistics.EntityData.Leafs.Append("modifies", types.YLeaf{"Modifies", routeStatistics.Modifies})
    routeStatistics.EntityData.Leafs.Append("reachables", types.YLeaf{"Reachables", routeStatistics.Reachables})
    routeStatistics.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", routeStatistics.Unreachables})
    routeStatistics.EntityData.Leafs.Append("touches", types.YLeaf{"Touches", routeStatistics.Touches})

    routeStatistics.EntityData.YListKeys = []string {}

    return &(routeStatistics.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "priority-summary"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "priority-summary"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_FrrStatistic
// Fast Re-Route Statistics
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_FrrStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total Number of Routes. The type is interface{} with range: 0..4294967295.
    TotalRoutes interface{}

    // Fully Protected Routes. The type is interface{} with range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Partially Protected Routes. The type is interface{} with range:
    // 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage. The type is string. Units are percentage.
    Coverage interface{}
}

func (frrStatistic *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_PrioritySummary_FrrStatistic) GetEntityData() *types.CommonEntityData {
    frrStatistic.EntityData.YFilter = frrStatistic.YFilter
    frrStatistic.EntityData.YangName = "frr-statistic"
    frrStatistic.EntityData.BundleName = "cisco_ios_xr"
    frrStatistic.EntityData.ParentYangName = "priority-summary"
    frrStatistic.EntityData.SegmentPath = "frr-statistic"
    frrStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrStatistic.EntityData.Children = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", frrStatistic.TotalRoutes})
    frrStatistic.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", frrStatistic.FullyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", frrStatistic.PartiallyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", frrStatistic.Coverage})

    frrStatistic.EntityData.YListKeys = []string {}

    return &(frrStatistic.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline
// Convergence timeline details
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route origin (routing protocol).
    RouteOrigin Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RouteOrigin

    // Entry point of IPv4 RIB.
    RiBv4Enter Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Enter

    // Exit point from IPv4 RIB to FIBs.
    RiBv4Exit Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Exit

    // Route Redistribute point from IPv4 RIB to LDP.
    RiBv4Redistribute Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Redistribute

    // Entry point of LDP.
    LdpEnter Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LdpEnter

    // Exit point of LDP to LSD.
    LdpExit Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LdpExit

    // Entry point of LSD.
    LsdEnter Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LsdEnter

    // Exit point of LSD to FIBs.
    LsdExit Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LsdExit

    // List of Linecards' completion point for IP routes. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcIp.
    LcIp []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcIp

    // List of Linecards' completion point for MPLS labels. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcMpls.
    LcMpls []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcMpls
}

func (convergenceTimeline *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline) GetEntityData() *types.CommonEntityData {
    convergenceTimeline.EntityData.YFilter = convergenceTimeline.YFilter
    convergenceTimeline.EntityData.YangName = "convergence-timeline"
    convergenceTimeline.EntityData.BundleName = "cisco_ios_xr"
    convergenceTimeline.EntityData.ParentYangName = "priority"
    convergenceTimeline.EntityData.SegmentPath = "convergence-timeline"
    convergenceTimeline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    convergenceTimeline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    convergenceTimeline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    convergenceTimeline.EntityData.Children = types.NewOrderedMap()
    convergenceTimeline.EntityData.Children.Append("route-origin", types.YChild{"RouteOrigin", &convergenceTimeline.RouteOrigin})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-enter", types.YChild{"RiBv4Enter", &convergenceTimeline.RiBv4Enter})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-exit", types.YChild{"RiBv4Exit", &convergenceTimeline.RiBv4Exit})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-redistribute", types.YChild{"RiBv4Redistribute", &convergenceTimeline.RiBv4Redistribute})
    convergenceTimeline.EntityData.Children.Append("ldp-enter", types.YChild{"LdpEnter", &convergenceTimeline.LdpEnter})
    convergenceTimeline.EntityData.Children.Append("ldp-exit", types.YChild{"LdpExit", &convergenceTimeline.LdpExit})
    convergenceTimeline.EntityData.Children.Append("lsd-enter", types.YChild{"LsdEnter", &convergenceTimeline.LsdEnter})
    convergenceTimeline.EntityData.Children.Append("lsd-exit", types.YChild{"LsdExit", &convergenceTimeline.LsdExit})
    convergenceTimeline.EntityData.Children.Append("lc-ip", types.YChild{"LcIp", nil})
    for i := range convergenceTimeline.LcIp {
        convergenceTimeline.EntityData.Children.Append(types.GetSegmentPath(convergenceTimeline.LcIp[i]), types.YChild{"LcIp", convergenceTimeline.LcIp[i]})
    }
    convergenceTimeline.EntityData.Children.Append("lc-mpls", types.YChild{"LcMpls", nil})
    for i := range convergenceTimeline.LcMpls {
        convergenceTimeline.EntityData.Children.Append(types.GetSegmentPath(convergenceTimeline.LcMpls[i]), types.YChild{"LcMpls", convergenceTimeline.LcMpls[i]})
    }
    convergenceTimeline.EntityData.Leafs = types.NewOrderedMap()

    convergenceTimeline.EntityData.YListKeys = []string {}

    return &(convergenceTimeline.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RouteOrigin
// Route origin (routing protocol)
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RouteOrigin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (routeOrigin *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RouteOrigin) GetEntityData() *types.CommonEntityData {
    routeOrigin.EntityData.YFilter = routeOrigin.YFilter
    routeOrigin.EntityData.YangName = "route-origin"
    routeOrigin.EntityData.BundleName = "cisco_ios_xr"
    routeOrigin.EntityData.ParentYangName = "convergence-timeline"
    routeOrigin.EntityData.SegmentPath = "route-origin"
    routeOrigin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeOrigin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeOrigin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeOrigin.EntityData.Children = types.NewOrderedMap()
    routeOrigin.EntityData.Leafs = types.NewOrderedMap()
    routeOrigin.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", routeOrigin.StartTime})
    routeOrigin.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", routeOrigin.EndTime})
    routeOrigin.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", routeOrigin.Duration})

    routeOrigin.EntityData.YListKeys = []string {}

    return &(routeOrigin.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Enter
// Entry point of IPv4 RIB
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Enter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Enter *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Enter) GetEntityData() *types.CommonEntityData {
    riBv4Enter.EntityData.YFilter = riBv4Enter.YFilter
    riBv4Enter.EntityData.YangName = "ri-bv4-enter"
    riBv4Enter.EntityData.BundleName = "cisco_ios_xr"
    riBv4Enter.EntityData.ParentYangName = "convergence-timeline"
    riBv4Enter.EntityData.SegmentPath = "ri-bv4-enter"
    riBv4Enter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Enter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Enter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Enter.EntityData.Children = types.NewOrderedMap()
    riBv4Enter.EntityData.Leafs = types.NewOrderedMap()
    riBv4Enter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Enter.StartTime})
    riBv4Enter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Enter.EndTime})
    riBv4Enter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Enter.Duration})

    riBv4Enter.EntityData.YListKeys = []string {}

    return &(riBv4Enter.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Exit
// Exit point from IPv4 RIB to FIBs
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Exit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Exit *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Exit) GetEntityData() *types.CommonEntityData {
    riBv4Exit.EntityData.YFilter = riBv4Exit.YFilter
    riBv4Exit.EntityData.YangName = "ri-bv4-exit"
    riBv4Exit.EntityData.BundleName = "cisco_ios_xr"
    riBv4Exit.EntityData.ParentYangName = "convergence-timeline"
    riBv4Exit.EntityData.SegmentPath = "ri-bv4-exit"
    riBv4Exit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Exit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Exit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Exit.EntityData.Children = types.NewOrderedMap()
    riBv4Exit.EntityData.Leafs = types.NewOrderedMap()
    riBv4Exit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Exit.StartTime})
    riBv4Exit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Exit.EndTime})
    riBv4Exit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Exit.Duration})

    riBv4Exit.EntityData.YListKeys = []string {}

    return &(riBv4Exit.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Redistribute
// Route Redistribute point from IPv4 RIB to LDP
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Redistribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Redistribute *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_RiBv4Redistribute) GetEntityData() *types.CommonEntityData {
    riBv4Redistribute.EntityData.YFilter = riBv4Redistribute.YFilter
    riBv4Redistribute.EntityData.YangName = "ri-bv4-redistribute"
    riBv4Redistribute.EntityData.BundleName = "cisco_ios_xr"
    riBv4Redistribute.EntityData.ParentYangName = "convergence-timeline"
    riBv4Redistribute.EntityData.SegmentPath = "ri-bv4-redistribute"
    riBv4Redistribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Redistribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Redistribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Redistribute.EntityData.Children = types.NewOrderedMap()
    riBv4Redistribute.EntityData.Leafs = types.NewOrderedMap()
    riBv4Redistribute.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Redistribute.StartTime})
    riBv4Redistribute.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Redistribute.EndTime})
    riBv4Redistribute.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Redistribute.Duration})

    riBv4Redistribute.EntityData.YListKeys = []string {}

    return &(riBv4Redistribute.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LdpEnter
// Entry point of LDP
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LdpEnter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (ldpEnter *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LdpEnter) GetEntityData() *types.CommonEntityData {
    ldpEnter.EntityData.YFilter = ldpEnter.YFilter
    ldpEnter.EntityData.YangName = "ldp-enter"
    ldpEnter.EntityData.BundleName = "cisco_ios_xr"
    ldpEnter.EntityData.ParentYangName = "convergence-timeline"
    ldpEnter.EntityData.SegmentPath = "ldp-enter"
    ldpEnter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpEnter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpEnter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpEnter.EntityData.Children = types.NewOrderedMap()
    ldpEnter.EntityData.Leafs = types.NewOrderedMap()
    ldpEnter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", ldpEnter.StartTime})
    ldpEnter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", ldpEnter.EndTime})
    ldpEnter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ldpEnter.Duration})

    ldpEnter.EntityData.YListKeys = []string {}

    return &(ldpEnter.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LdpExit
// Exit point of LDP to LSD
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LdpExit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (ldpExit *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LdpExit) GetEntityData() *types.CommonEntityData {
    ldpExit.EntityData.YFilter = ldpExit.YFilter
    ldpExit.EntityData.YangName = "ldp-exit"
    ldpExit.EntityData.BundleName = "cisco_ios_xr"
    ldpExit.EntityData.ParentYangName = "convergence-timeline"
    ldpExit.EntityData.SegmentPath = "ldp-exit"
    ldpExit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpExit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpExit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpExit.EntityData.Children = types.NewOrderedMap()
    ldpExit.EntityData.Leafs = types.NewOrderedMap()
    ldpExit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", ldpExit.StartTime})
    ldpExit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", ldpExit.EndTime})
    ldpExit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ldpExit.Duration})

    ldpExit.EntityData.YListKeys = []string {}

    return &(ldpExit.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LsdEnter
// Entry point of LSD
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LsdEnter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (lsdEnter *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LsdEnter) GetEntityData() *types.CommonEntityData {
    lsdEnter.EntityData.YFilter = lsdEnter.YFilter
    lsdEnter.EntityData.YangName = "lsd-enter"
    lsdEnter.EntityData.BundleName = "cisco_ios_xr"
    lsdEnter.EntityData.ParentYangName = "convergence-timeline"
    lsdEnter.EntityData.SegmentPath = "lsd-enter"
    lsdEnter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsdEnter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsdEnter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsdEnter.EntityData.Children = types.NewOrderedMap()
    lsdEnter.EntityData.Leafs = types.NewOrderedMap()
    lsdEnter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", lsdEnter.StartTime})
    lsdEnter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", lsdEnter.EndTime})
    lsdEnter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", lsdEnter.Duration})

    lsdEnter.EntityData.YListKeys = []string {}

    return &(lsdEnter.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LsdExit
// Exit point of LSD to FIBs
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LsdExit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (lsdExit *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LsdExit) GetEntityData() *types.CommonEntityData {
    lsdExit.EntityData.YFilter = lsdExit.YFilter
    lsdExit.EntityData.YangName = "lsd-exit"
    lsdExit.EntityData.BundleName = "cisco_ios_xr"
    lsdExit.EntityData.ParentYangName = "convergence-timeline"
    lsdExit.EntityData.SegmentPath = "lsd-exit"
    lsdExit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsdExit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsdExit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsdExit.EntityData.Children = types.NewOrderedMap()
    lsdExit.EntityData.Leafs = types.NewOrderedMap()
    lsdExit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", lsdExit.StartTime})
    lsdExit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", lsdExit.EndTime})
    lsdExit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", lsdExit.Duration})

    lsdExit.EntityData.YListKeys = []string {}

    return &(lsdExit.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcIp
// List of Linecards' completion point for IP
// routes
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB.
    FibComplete Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcIp_FibComplete
}

func (lcIp *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcIp) GetEntityData() *types.CommonEntityData {
    lcIp.EntityData.YFilter = lcIp.YFilter
    lcIp.EntityData.YangName = "lc-ip"
    lcIp.EntityData.BundleName = "cisco_ios_xr"
    lcIp.EntityData.ParentYangName = "convergence-timeline"
    lcIp.EntityData.SegmentPath = "lc-ip"
    lcIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcIp.EntityData.Children = types.NewOrderedMap()
    lcIp.EntityData.Children.Append("fib-complete", types.YChild{"FibComplete", &lcIp.FibComplete})
    lcIp.EntityData.Leafs = types.NewOrderedMap()
    lcIp.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcIp.NodeName})
    lcIp.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcIp.Speed})

    lcIp.EntityData.YListKeys = []string {}

    return &(lcIp.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcIp_FibComplete
// Completion point of FIB
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcIp_FibComplete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (fibComplete *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcIp_FibComplete) GetEntityData() *types.CommonEntityData {
    fibComplete.EntityData.YFilter = fibComplete.YFilter
    fibComplete.EntityData.YangName = "fib-complete"
    fibComplete.EntityData.BundleName = "cisco_ios_xr"
    fibComplete.EntityData.ParentYangName = "lc-ip"
    fibComplete.EntityData.SegmentPath = "fib-complete"
    fibComplete.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibComplete.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibComplete.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibComplete.EntityData.Children = types.NewOrderedMap()
    fibComplete.EntityData.Leafs = types.NewOrderedMap()
    fibComplete.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", fibComplete.StartTime})
    fibComplete.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", fibComplete.EndTime})
    fibComplete.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", fibComplete.Duration})

    fibComplete.EntityData.YListKeys = []string {}

    return &(fibComplete.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcMpls
// List of Linecards' completion point for MPLS
// labels
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB.
    FibComplete Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcMpls_FibComplete
}

func (lcMpls *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcMpls) GetEntityData() *types.CommonEntityData {
    lcMpls.EntityData.YFilter = lcMpls.YFilter
    lcMpls.EntityData.YangName = "lc-mpls"
    lcMpls.EntityData.BundleName = "cisco_ios_xr"
    lcMpls.EntityData.ParentYangName = "convergence-timeline"
    lcMpls.EntityData.SegmentPath = "lc-mpls"
    lcMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcMpls.EntityData.Children = types.NewOrderedMap()
    lcMpls.EntityData.Children.Append("fib-complete", types.YChild{"FibComplete", &lcMpls.FibComplete})
    lcMpls.EntityData.Leafs = types.NewOrderedMap()
    lcMpls.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcMpls.NodeName})
    lcMpls.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcMpls.Speed})

    lcMpls.EntityData.YListKeys = []string {}

    return &(lcMpls.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcMpls_FibComplete
// Completion point of FIB
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcMpls_FibComplete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (fibComplete *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_ConvergenceTimeline_LcMpls_FibComplete) GetEntityData() *types.CommonEntityData {
    fibComplete.EntityData.YFilter = fibComplete.YFilter
    fibComplete.EntityData.YangName = "fib-complete"
    fibComplete.EntityData.BundleName = "cisco_ios_xr"
    fibComplete.EntityData.ParentYangName = "lc-mpls"
    fibComplete.EntityData.SegmentPath = "fib-complete"
    fibComplete.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibComplete.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibComplete.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibComplete.EntityData.Children = types.NewOrderedMap()
    fibComplete.EntityData.Leafs = types.NewOrderedMap()
    fibComplete.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", fibComplete.StartTime})
    fibComplete.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", fibComplete.EndTime})
    fibComplete.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", fibComplete.Duration})

    fibComplete.EntityData.YListKeys = []string {}

    return &(fibComplete.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_LeafNetworksAdded
// List of Leaf Networks Added
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_LeafNetworksAdded struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is interface{} with range: 0..255.
    NetMask interface{}
}

func (leafNetworksAdded *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_LeafNetworksAdded) GetEntityData() *types.CommonEntityData {
    leafNetworksAdded.EntityData.YFilter = leafNetworksAdded.YFilter
    leafNetworksAdded.EntityData.YangName = "leaf-networks-added"
    leafNetworksAdded.EntityData.BundleName = "cisco_ios_xr"
    leafNetworksAdded.EntityData.ParentYangName = "priority"
    leafNetworksAdded.EntityData.SegmentPath = "leaf-networks-added"
    leafNetworksAdded.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leafNetworksAdded.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leafNetworksAdded.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leafNetworksAdded.EntityData.Children = types.NewOrderedMap()
    leafNetworksAdded.EntityData.Leafs = types.NewOrderedMap()
    leafNetworksAdded.EntityData.Leafs.Append("address", types.YLeaf{"Address", leafNetworksAdded.Address})
    leafNetworksAdded.EntityData.Leafs.Append("net-mask", types.YLeaf{"NetMask", leafNetworksAdded.NetMask})

    leafNetworksAdded.EntityData.YListKeys = []string {}

    return &(leafNetworksAdded.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_LeafNetworksDeleted
// List of Leaf Networks Deleted
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_LeafNetworksDeleted struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is interface{} with range: 0..255.
    NetMask interface{}
}

func (leafNetworksDeleted *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_Priority_LeafNetworksDeleted) GetEntityData() *types.CommonEntityData {
    leafNetworksDeleted.EntityData.YFilter = leafNetworksDeleted.YFilter
    leafNetworksDeleted.EntityData.YangName = "leaf-networks-deleted"
    leafNetworksDeleted.EntityData.BundleName = "cisco_ios_xr"
    leafNetworksDeleted.EntityData.ParentYangName = "priority"
    leafNetworksDeleted.EntityData.SegmentPath = "leaf-networks-deleted"
    leafNetworksDeleted.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leafNetworksDeleted.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leafNetworksDeleted.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leafNetworksDeleted.EntityData.Children = types.NewOrderedMap()
    leafNetworksDeleted.EntityData.Leafs = types.NewOrderedMap()
    leafNetworksDeleted.EntityData.Leafs.Append("address", types.YLeaf{"Address", leafNetworksDeleted.Address})
    leafNetworksDeleted.EntityData.Leafs.Append("net-mask", types.YLeaf{"NetMask", leafNetworksDeleted.NetMask})

    leafNetworksDeleted.EntityData.YListKeys = []string {}

    return &(leafNetworksDeleted.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_LsaProcessed
// List of type 1/2 LSA changes processed
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_LsaProcessed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (lsaProcessed *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_DijkstraRun_LsaProcessed) GetEntityData() *types.CommonEntityData {
    lsaProcessed.EntityData.YFilter = lsaProcessed.YFilter
    lsaProcessed.EntityData.YangName = "lsa-processed"
    lsaProcessed.EntityData.BundleName = "cisco_ios_xr"
    lsaProcessed.EntityData.ParentYangName = "dijkstra-run"
    lsaProcessed.EntityData.SegmentPath = "lsa-processed"
    lsaProcessed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsaProcessed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsaProcessed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsaProcessed.EntityData.Children = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", lsaProcessed.LsaId})
    lsaProcessed.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", lsaProcessed.SequenceNumber})
    lsaProcessed.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", lsaProcessed.LsaType})
    lsaProcessed.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", lsaProcessed.OriginRouterId})
    lsaProcessed.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lsaProcessed.ChangeType})
    lsaProcessed.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", lsaProcessed.ReceptionTime})

    lsaProcessed.EntityData.YListKeys = []string {}

    return &(lsaProcessed.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal
// Inter-area & external calculation information
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Convergence information on a per-priority basis. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority.
    Priority []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority
}

func (interAreaAndExternal *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal) GetEntityData() *types.CommonEntityData {
    interAreaAndExternal.EntityData.YFilter = interAreaAndExternal.YFilter
    interAreaAndExternal.EntityData.YangName = "inter-area-and-external"
    interAreaAndExternal.EntityData.BundleName = "cisco_ios_xr"
    interAreaAndExternal.EntityData.ParentYangName = "spf-run-offline"
    interAreaAndExternal.EntityData.SegmentPath = "inter-area-and-external"
    interAreaAndExternal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interAreaAndExternal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interAreaAndExternal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interAreaAndExternal.EntityData.Children = types.NewOrderedMap()
    interAreaAndExternal.EntityData.Children.Append("priority", types.YChild{"Priority", nil})
    for i := range interAreaAndExternal.Priority {
        interAreaAndExternal.EntityData.Children.Append(types.GetSegmentPath(interAreaAndExternal.Priority[i]), types.YChild{"Priority", interAreaAndExternal.Priority[i]})
    }
    interAreaAndExternal.EntityData.Leafs = types.NewOrderedMap()

    interAreaAndExternal.EntityData.YListKeys = []string {}

    return &(interAreaAndExternal.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority
// Convergence information on a per-priority basis
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of the priority.
    PrioritySummary Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary

    // Convergence timeline details. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline.
    ConvergenceTimeline []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline

    // List of Leaf Networks Added. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_LeafNetworksAdded.
    LeafNetworksAdded []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_LeafNetworksAdded

    // List of Leaf Networks Deleted. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_LeafNetworksDeleted.
    LeafNetworksDeleted []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_LeafNetworksDeleted
}

func (priority *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "inter-area-and-external"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Children.Append("priority-summary", types.YChild{"PrioritySummary", &priority.PrioritySummary})
    priority.EntityData.Children.Append("convergence-timeline", types.YChild{"ConvergenceTimeline", nil})
    for i := range priority.ConvergenceTimeline {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.ConvergenceTimeline[i]), types.YChild{"ConvergenceTimeline", priority.ConvergenceTimeline[i]})
    }
    priority.EntityData.Children.Append("leaf-networks-added", types.YChild{"LeafNetworksAdded", nil})
    for i := range priority.LeafNetworksAdded {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.LeafNetworksAdded[i]), types.YChild{"LeafNetworksAdded", priority.LeafNetworksAdded[i]})
    }
    priority.EntityData.Children.Append("leaf-networks-deleted", types.YChild{"LeafNetworksDeleted", nil})
    for i := range priority.LeafNetworksDeleted {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.LeafNetworksDeleted[i]), types.YChild{"LeafNetworksDeleted", priority.LeafNetworksDeleted[i]})
    }
    priority.EntityData.Leafs = types.NewOrderedMap()

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary
// Summary of the priority
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Critical, High, Medium or Low. The type is RcmdPriorityLevel.
    Level interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Number of Type 3 LSA. The type is interface{} with range: 0..4294967295.
    Type3lsAs interface{}

    // Number of Type 4 LSA. The type is interface{} with range: 0..4294967295.
    Type4lsAs interface{}

    // Number of Type 5/7 LSA. The type is interface{} with range: 0..4294967295.
    Type57lsAs interface{}

    // Route statistics.
    RouteStatistics Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary_RouteStatistics

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary_MplsConvergenceTime
}

func (prioritySummary *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary) GetEntityData() *types.CommonEntityData {
    prioritySummary.EntityData.YFilter = prioritySummary.YFilter
    prioritySummary.EntityData.YangName = "priority-summary"
    prioritySummary.EntityData.BundleName = "cisco_ios_xr"
    prioritySummary.EntityData.ParentYangName = "priority"
    prioritySummary.EntityData.SegmentPath = "priority-summary"
    prioritySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prioritySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prioritySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prioritySummary.EntityData.Children = types.NewOrderedMap()
    prioritySummary.EntityData.Children.Append("route-statistics", types.YChild{"RouteStatistics", &prioritySummary.RouteStatistics})
    prioritySummary.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prioritySummary.IpConvergenceTime})
    prioritySummary.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prioritySummary.MplsConvergenceTime})
    prioritySummary.EntityData.Leafs = types.NewOrderedMap()
    prioritySummary.EntityData.Leafs.Append("level", types.YLeaf{"Level", prioritySummary.Level})
    prioritySummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prioritySummary.ThresholdExceeded})
    prioritySummary.EntityData.Leafs.Append("type3ls-as", types.YLeaf{"Type3lsAs", prioritySummary.Type3lsAs})
    prioritySummary.EntityData.Leafs.Append("type4ls-as", types.YLeaf{"Type4lsAs", prioritySummary.Type4lsAs})
    prioritySummary.EntityData.Leafs.Append("type57ls-as", types.YLeaf{"Type57lsAs", prioritySummary.Type57lsAs})

    prioritySummary.EntityData.YListKeys = []string {}

    return &(prioritySummary.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary_RouteStatistics
// Route statistics
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary_RouteStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Added. The type is interface{} with range: 0..4294967295.
    Adds interface{}

    // Deleted. The type is interface{} with range: 0..4294967295.
    Deletes interface{}

    // Modified. The type is interface{} with range: 0..4294967295.
    Modifies interface{}

    // Reachable. The type is interface{} with range: 0..4294967295.
    Reachables interface{}

    // Unreachable. The type is interface{} with range: 0..4294967295.
    Unreachables interface{}

    // Touched. The type is interface{} with range: 0..4294967295.
    Touches interface{}
}

func (routeStatistics *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary_RouteStatistics) GetEntityData() *types.CommonEntityData {
    routeStatistics.EntityData.YFilter = routeStatistics.YFilter
    routeStatistics.EntityData.YangName = "route-statistics"
    routeStatistics.EntityData.BundleName = "cisco_ios_xr"
    routeStatistics.EntityData.ParentYangName = "priority-summary"
    routeStatistics.EntityData.SegmentPath = "route-statistics"
    routeStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeStatistics.EntityData.Children = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs.Append("adds", types.YLeaf{"Adds", routeStatistics.Adds})
    routeStatistics.EntityData.Leafs.Append("deletes", types.YLeaf{"Deletes", routeStatistics.Deletes})
    routeStatistics.EntityData.Leafs.Append("modifies", types.YLeaf{"Modifies", routeStatistics.Modifies})
    routeStatistics.EntityData.Leafs.Append("reachables", types.YLeaf{"Reachables", routeStatistics.Reachables})
    routeStatistics.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", routeStatistics.Unreachables})
    routeStatistics.EntityData.Leafs.Append("touches", types.YLeaf{"Touches", routeStatistics.Touches})

    routeStatistics.EntityData.YListKeys = []string {}

    return &(routeStatistics.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "priority-summary"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_PrioritySummary_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "priority-summary"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline
// Convergence timeline details
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route origin (routing protocol).
    RouteOrigin Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RouteOrigin

    // Entry point of IPv4 RIB.
    RiBv4Enter Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Enter

    // Exit point from IPv4 RIB to FIBs.
    RiBv4Exit Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Exit

    // Route Redistribute point from IPv4 RIB to LDP.
    RiBv4Redistribute Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Redistribute

    // Entry point of LDP.
    LdpEnter Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpEnter

    // Exit point of LDP to LSD.
    LdpExit Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpExit

    // Entry point of LSD.
    LsdEnter Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdEnter

    // Exit point of LSD to FIBs.
    LsdExit Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdExit

    // List of Linecards' completion point for IP routes. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp.
    LcIp []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp

    // List of Linecards' completion point for MPLS labels. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls.
    LcMpls []*Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls
}

func (convergenceTimeline *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline) GetEntityData() *types.CommonEntityData {
    convergenceTimeline.EntityData.YFilter = convergenceTimeline.YFilter
    convergenceTimeline.EntityData.YangName = "convergence-timeline"
    convergenceTimeline.EntityData.BundleName = "cisco_ios_xr"
    convergenceTimeline.EntityData.ParentYangName = "priority"
    convergenceTimeline.EntityData.SegmentPath = "convergence-timeline"
    convergenceTimeline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    convergenceTimeline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    convergenceTimeline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    convergenceTimeline.EntityData.Children = types.NewOrderedMap()
    convergenceTimeline.EntityData.Children.Append("route-origin", types.YChild{"RouteOrigin", &convergenceTimeline.RouteOrigin})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-enter", types.YChild{"RiBv4Enter", &convergenceTimeline.RiBv4Enter})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-exit", types.YChild{"RiBv4Exit", &convergenceTimeline.RiBv4Exit})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-redistribute", types.YChild{"RiBv4Redistribute", &convergenceTimeline.RiBv4Redistribute})
    convergenceTimeline.EntityData.Children.Append("ldp-enter", types.YChild{"LdpEnter", &convergenceTimeline.LdpEnter})
    convergenceTimeline.EntityData.Children.Append("ldp-exit", types.YChild{"LdpExit", &convergenceTimeline.LdpExit})
    convergenceTimeline.EntityData.Children.Append("lsd-enter", types.YChild{"LsdEnter", &convergenceTimeline.LsdEnter})
    convergenceTimeline.EntityData.Children.Append("lsd-exit", types.YChild{"LsdExit", &convergenceTimeline.LsdExit})
    convergenceTimeline.EntityData.Children.Append("lc-ip", types.YChild{"LcIp", nil})
    for i := range convergenceTimeline.LcIp {
        convergenceTimeline.EntityData.Children.Append(types.GetSegmentPath(convergenceTimeline.LcIp[i]), types.YChild{"LcIp", convergenceTimeline.LcIp[i]})
    }
    convergenceTimeline.EntityData.Children.Append("lc-mpls", types.YChild{"LcMpls", nil})
    for i := range convergenceTimeline.LcMpls {
        convergenceTimeline.EntityData.Children.Append(types.GetSegmentPath(convergenceTimeline.LcMpls[i]), types.YChild{"LcMpls", convergenceTimeline.LcMpls[i]})
    }
    convergenceTimeline.EntityData.Leafs = types.NewOrderedMap()

    convergenceTimeline.EntityData.YListKeys = []string {}

    return &(convergenceTimeline.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RouteOrigin
// Route origin (routing protocol)
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RouteOrigin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (routeOrigin *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RouteOrigin) GetEntityData() *types.CommonEntityData {
    routeOrigin.EntityData.YFilter = routeOrigin.YFilter
    routeOrigin.EntityData.YangName = "route-origin"
    routeOrigin.EntityData.BundleName = "cisco_ios_xr"
    routeOrigin.EntityData.ParentYangName = "convergence-timeline"
    routeOrigin.EntityData.SegmentPath = "route-origin"
    routeOrigin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeOrigin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeOrigin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeOrigin.EntityData.Children = types.NewOrderedMap()
    routeOrigin.EntityData.Leafs = types.NewOrderedMap()
    routeOrigin.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", routeOrigin.StartTime})
    routeOrigin.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", routeOrigin.EndTime})
    routeOrigin.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", routeOrigin.Duration})

    routeOrigin.EntityData.YListKeys = []string {}

    return &(routeOrigin.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Enter
// Entry point of IPv4 RIB
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Enter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Enter *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Enter) GetEntityData() *types.CommonEntityData {
    riBv4Enter.EntityData.YFilter = riBv4Enter.YFilter
    riBv4Enter.EntityData.YangName = "ri-bv4-enter"
    riBv4Enter.EntityData.BundleName = "cisco_ios_xr"
    riBv4Enter.EntityData.ParentYangName = "convergence-timeline"
    riBv4Enter.EntityData.SegmentPath = "ri-bv4-enter"
    riBv4Enter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Enter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Enter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Enter.EntityData.Children = types.NewOrderedMap()
    riBv4Enter.EntityData.Leafs = types.NewOrderedMap()
    riBv4Enter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Enter.StartTime})
    riBv4Enter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Enter.EndTime})
    riBv4Enter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Enter.Duration})

    riBv4Enter.EntityData.YListKeys = []string {}

    return &(riBv4Enter.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Exit
// Exit point from IPv4 RIB to FIBs
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Exit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Exit *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Exit) GetEntityData() *types.CommonEntityData {
    riBv4Exit.EntityData.YFilter = riBv4Exit.YFilter
    riBv4Exit.EntityData.YangName = "ri-bv4-exit"
    riBv4Exit.EntityData.BundleName = "cisco_ios_xr"
    riBv4Exit.EntityData.ParentYangName = "convergence-timeline"
    riBv4Exit.EntityData.SegmentPath = "ri-bv4-exit"
    riBv4Exit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Exit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Exit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Exit.EntityData.Children = types.NewOrderedMap()
    riBv4Exit.EntityData.Leafs = types.NewOrderedMap()
    riBv4Exit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Exit.StartTime})
    riBv4Exit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Exit.EndTime})
    riBv4Exit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Exit.Duration})

    riBv4Exit.EntityData.YListKeys = []string {}

    return &(riBv4Exit.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Redistribute
// Route Redistribute point from IPv4 RIB to LDP
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Redistribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Redistribute *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_RiBv4Redistribute) GetEntityData() *types.CommonEntityData {
    riBv4Redistribute.EntityData.YFilter = riBv4Redistribute.YFilter
    riBv4Redistribute.EntityData.YangName = "ri-bv4-redistribute"
    riBv4Redistribute.EntityData.BundleName = "cisco_ios_xr"
    riBv4Redistribute.EntityData.ParentYangName = "convergence-timeline"
    riBv4Redistribute.EntityData.SegmentPath = "ri-bv4-redistribute"
    riBv4Redistribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Redistribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Redistribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Redistribute.EntityData.Children = types.NewOrderedMap()
    riBv4Redistribute.EntityData.Leafs = types.NewOrderedMap()
    riBv4Redistribute.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Redistribute.StartTime})
    riBv4Redistribute.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Redistribute.EndTime})
    riBv4Redistribute.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Redistribute.Duration})

    riBv4Redistribute.EntityData.YListKeys = []string {}

    return &(riBv4Redistribute.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpEnter
// Entry point of LDP
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpEnter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (ldpEnter *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpEnter) GetEntityData() *types.CommonEntityData {
    ldpEnter.EntityData.YFilter = ldpEnter.YFilter
    ldpEnter.EntityData.YangName = "ldp-enter"
    ldpEnter.EntityData.BundleName = "cisco_ios_xr"
    ldpEnter.EntityData.ParentYangName = "convergence-timeline"
    ldpEnter.EntityData.SegmentPath = "ldp-enter"
    ldpEnter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpEnter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpEnter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpEnter.EntityData.Children = types.NewOrderedMap()
    ldpEnter.EntityData.Leafs = types.NewOrderedMap()
    ldpEnter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", ldpEnter.StartTime})
    ldpEnter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", ldpEnter.EndTime})
    ldpEnter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ldpEnter.Duration})

    ldpEnter.EntityData.YListKeys = []string {}

    return &(ldpEnter.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpExit
// Exit point of LDP to LSD
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpExit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (ldpExit *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LdpExit) GetEntityData() *types.CommonEntityData {
    ldpExit.EntityData.YFilter = ldpExit.YFilter
    ldpExit.EntityData.YangName = "ldp-exit"
    ldpExit.EntityData.BundleName = "cisco_ios_xr"
    ldpExit.EntityData.ParentYangName = "convergence-timeline"
    ldpExit.EntityData.SegmentPath = "ldp-exit"
    ldpExit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpExit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpExit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpExit.EntityData.Children = types.NewOrderedMap()
    ldpExit.EntityData.Leafs = types.NewOrderedMap()
    ldpExit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", ldpExit.StartTime})
    ldpExit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", ldpExit.EndTime})
    ldpExit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ldpExit.Duration})

    ldpExit.EntityData.YListKeys = []string {}

    return &(ldpExit.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdEnter
// Entry point of LSD
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdEnter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (lsdEnter *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdEnter) GetEntityData() *types.CommonEntityData {
    lsdEnter.EntityData.YFilter = lsdEnter.YFilter
    lsdEnter.EntityData.YangName = "lsd-enter"
    lsdEnter.EntityData.BundleName = "cisco_ios_xr"
    lsdEnter.EntityData.ParentYangName = "convergence-timeline"
    lsdEnter.EntityData.SegmentPath = "lsd-enter"
    lsdEnter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsdEnter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsdEnter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsdEnter.EntityData.Children = types.NewOrderedMap()
    lsdEnter.EntityData.Leafs = types.NewOrderedMap()
    lsdEnter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", lsdEnter.StartTime})
    lsdEnter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", lsdEnter.EndTime})
    lsdEnter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", lsdEnter.Duration})

    lsdEnter.EntityData.YListKeys = []string {}

    return &(lsdEnter.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdExit
// Exit point of LSD to FIBs
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdExit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (lsdExit *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LsdExit) GetEntityData() *types.CommonEntityData {
    lsdExit.EntityData.YFilter = lsdExit.YFilter
    lsdExit.EntityData.YangName = "lsd-exit"
    lsdExit.EntityData.BundleName = "cisco_ios_xr"
    lsdExit.EntityData.ParentYangName = "convergence-timeline"
    lsdExit.EntityData.SegmentPath = "lsd-exit"
    lsdExit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsdExit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsdExit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsdExit.EntityData.Children = types.NewOrderedMap()
    lsdExit.EntityData.Leafs = types.NewOrderedMap()
    lsdExit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", lsdExit.StartTime})
    lsdExit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", lsdExit.EndTime})
    lsdExit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", lsdExit.Duration})

    lsdExit.EntityData.YListKeys = []string {}

    return &(lsdExit.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp
// List of Linecards' completion point for IP
// routes
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB.
    FibComplete Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp_FibComplete
}

func (lcIp *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp) GetEntityData() *types.CommonEntityData {
    lcIp.EntityData.YFilter = lcIp.YFilter
    lcIp.EntityData.YangName = "lc-ip"
    lcIp.EntityData.BundleName = "cisco_ios_xr"
    lcIp.EntityData.ParentYangName = "convergence-timeline"
    lcIp.EntityData.SegmentPath = "lc-ip"
    lcIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcIp.EntityData.Children = types.NewOrderedMap()
    lcIp.EntityData.Children.Append("fib-complete", types.YChild{"FibComplete", &lcIp.FibComplete})
    lcIp.EntityData.Leafs = types.NewOrderedMap()
    lcIp.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcIp.NodeName})
    lcIp.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcIp.Speed})

    lcIp.EntityData.YListKeys = []string {}

    return &(lcIp.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp_FibComplete
// Completion point of FIB
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp_FibComplete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (fibComplete *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcIp_FibComplete) GetEntityData() *types.CommonEntityData {
    fibComplete.EntityData.YFilter = fibComplete.YFilter
    fibComplete.EntityData.YangName = "fib-complete"
    fibComplete.EntityData.BundleName = "cisco_ios_xr"
    fibComplete.EntityData.ParentYangName = "lc-ip"
    fibComplete.EntityData.SegmentPath = "fib-complete"
    fibComplete.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibComplete.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibComplete.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibComplete.EntityData.Children = types.NewOrderedMap()
    fibComplete.EntityData.Leafs = types.NewOrderedMap()
    fibComplete.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", fibComplete.StartTime})
    fibComplete.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", fibComplete.EndTime})
    fibComplete.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", fibComplete.Duration})

    fibComplete.EntityData.YListKeys = []string {}

    return &(fibComplete.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls
// List of Linecards' completion point for MPLS
// labels
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB.
    FibComplete Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls_FibComplete
}

func (lcMpls *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls) GetEntityData() *types.CommonEntityData {
    lcMpls.EntityData.YFilter = lcMpls.YFilter
    lcMpls.EntityData.YangName = "lc-mpls"
    lcMpls.EntityData.BundleName = "cisco_ios_xr"
    lcMpls.EntityData.ParentYangName = "convergence-timeline"
    lcMpls.EntityData.SegmentPath = "lc-mpls"
    lcMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcMpls.EntityData.Children = types.NewOrderedMap()
    lcMpls.EntityData.Children.Append("fib-complete", types.YChild{"FibComplete", &lcMpls.FibComplete})
    lcMpls.EntityData.Leafs = types.NewOrderedMap()
    lcMpls.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcMpls.NodeName})
    lcMpls.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcMpls.Speed})

    lcMpls.EntityData.YListKeys = []string {}

    return &(lcMpls.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls_FibComplete
// Completion point of FIB
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls_FibComplete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (fibComplete *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_ConvergenceTimeline_LcMpls_FibComplete) GetEntityData() *types.CommonEntityData {
    fibComplete.EntityData.YFilter = fibComplete.YFilter
    fibComplete.EntityData.YangName = "fib-complete"
    fibComplete.EntityData.BundleName = "cisco_ios_xr"
    fibComplete.EntityData.ParentYangName = "lc-mpls"
    fibComplete.EntityData.SegmentPath = "fib-complete"
    fibComplete.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibComplete.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibComplete.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibComplete.EntityData.Children = types.NewOrderedMap()
    fibComplete.EntityData.Leafs = types.NewOrderedMap()
    fibComplete.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", fibComplete.StartTime})
    fibComplete.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", fibComplete.EndTime})
    fibComplete.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", fibComplete.Duration})

    fibComplete.EntityData.YListKeys = []string {}

    return &(fibComplete.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_LeafNetworksAdded
// List of Leaf Networks Added
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_LeafNetworksAdded struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is interface{} with range: 0..255.
    NetMask interface{}
}

func (leafNetworksAdded *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_LeafNetworksAdded) GetEntityData() *types.CommonEntityData {
    leafNetworksAdded.EntityData.YFilter = leafNetworksAdded.YFilter
    leafNetworksAdded.EntityData.YangName = "leaf-networks-added"
    leafNetworksAdded.EntityData.BundleName = "cisco_ios_xr"
    leafNetworksAdded.EntityData.ParentYangName = "priority"
    leafNetworksAdded.EntityData.SegmentPath = "leaf-networks-added"
    leafNetworksAdded.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leafNetworksAdded.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leafNetworksAdded.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leafNetworksAdded.EntityData.Children = types.NewOrderedMap()
    leafNetworksAdded.EntityData.Leafs = types.NewOrderedMap()
    leafNetworksAdded.EntityData.Leafs.Append("address", types.YLeaf{"Address", leafNetworksAdded.Address})
    leafNetworksAdded.EntityData.Leafs.Append("net-mask", types.YLeaf{"NetMask", leafNetworksAdded.NetMask})

    leafNetworksAdded.EntityData.YListKeys = []string {}

    return &(leafNetworksAdded.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_LeafNetworksDeleted
// List of Leaf Networks Deleted
type Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_LeafNetworksDeleted struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is interface{} with range: 0..255.
    NetMask interface{}
}

func (leafNetworksDeleted *Rcmd_Ospf_Instances_Instance_SpfRunOfflines_SpfRunOffline_InterAreaAndExternal_Priority_LeafNetworksDeleted) GetEntityData() *types.CommonEntityData {
    leafNetworksDeleted.EntityData.YFilter = leafNetworksDeleted.YFilter
    leafNetworksDeleted.EntityData.YangName = "leaf-networks-deleted"
    leafNetworksDeleted.EntityData.BundleName = "cisco_ios_xr"
    leafNetworksDeleted.EntityData.ParentYangName = "priority"
    leafNetworksDeleted.EntityData.SegmentPath = "leaf-networks-deleted"
    leafNetworksDeleted.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leafNetworksDeleted.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leafNetworksDeleted.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leafNetworksDeleted.EntityData.Children = types.NewOrderedMap()
    leafNetworksDeleted.EntityData.Leafs = types.NewOrderedMap()
    leafNetworksDeleted.EntityData.Leafs.Append("address", types.YLeaf{"Address", leafNetworksDeleted.Address})
    leafNetworksDeleted.EntityData.Leafs.Append("net-mask", types.YLeaf{"NetMask", leafNetworksDeleted.NetMask})

    leafNetworksDeleted.EntityData.YListKeys = []string {}

    return &(leafNetworksDeleted.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries
// OSPF Summary-External Prefix events summary
// data
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF Summary-External Prefix Event data. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary.
    SummaryExternalEventSummary []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary
}

func (summaryExternalEventSummaries *Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries) GetEntityData() *types.CommonEntityData {
    summaryExternalEventSummaries.EntityData.YFilter = summaryExternalEventSummaries.YFilter
    summaryExternalEventSummaries.EntityData.YangName = "summary-external-event-summaries"
    summaryExternalEventSummaries.EntityData.BundleName = "cisco_ios_xr"
    summaryExternalEventSummaries.EntityData.ParentYangName = "instance"
    summaryExternalEventSummaries.EntityData.SegmentPath = "summary-external-event-summaries"
    summaryExternalEventSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryExternalEventSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryExternalEventSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryExternalEventSummaries.EntityData.Children = types.NewOrderedMap()
    summaryExternalEventSummaries.EntityData.Children.Append("summary-external-event-summary", types.YChild{"SummaryExternalEventSummary", nil})
    for i := range summaryExternalEventSummaries.SummaryExternalEventSummary {
        summaryExternalEventSummaries.EntityData.Children.Append(types.GetSegmentPath(summaryExternalEventSummaries.SummaryExternalEventSummary[i]), types.YChild{"SummaryExternalEventSummary", summaryExternalEventSummaries.SummaryExternalEventSummary[i]})
    }
    summaryExternalEventSummaries.EntityData.Leafs = types.NewOrderedMap()

    summaryExternalEventSummaries.EntityData.YListKeys = []string {}

    return &(summaryExternalEventSummaries.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary
// OSPF Summary-External Prefix Event data
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Event ID. The type is interface{} with
    // range: 1..4294967295.
    EventId interface{}

    // Prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLenth interface{}

    // Referenced SPF Run No (0 - Not Applicable). The type is interface{} with
    // range: 0..4294967295.
    SpfRunNo interface{}

    // Referenced IP-FRR Event ID (0 - Not Applicable). The type is interface{}
    // with range: 0..4294967295.
    IpfrrEventId interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Event processed priority. The type is RcmdPriorityLevel.
    Priority interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Route Type intra/inter/l1/l2. The type is RcmdShowRoute.
    RouteType interface{}

    // Route Path Change Type. The type is RcmdShowRoutePathChange.
    RoutePathChangeType interface{}

    // Protocol route cost. The type is interface{} with range: 0..4294967295.
    Cost interface{}

    // Event trigger time. The type is string.
    TriggerTime interface{}

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_MplsConvergenceTime

    // Path information. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_Path.
    Path []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_Path

    // LSA that triggered this event. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TriggerLsa.
    TriggerLsa []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TriggerLsa

    // Timeline information. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine.
    TimeLine []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine

    // List of LSAs processed. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_LsaProcessed.
    LsaProcessed []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_LsaProcessed
}

func (summaryExternalEventSummary *Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary) GetEntityData() *types.CommonEntityData {
    summaryExternalEventSummary.EntityData.YFilter = summaryExternalEventSummary.YFilter
    summaryExternalEventSummary.EntityData.YangName = "summary-external-event-summary"
    summaryExternalEventSummary.EntityData.BundleName = "cisco_ios_xr"
    summaryExternalEventSummary.EntityData.ParentYangName = "summary-external-event-summaries"
    summaryExternalEventSummary.EntityData.SegmentPath = "summary-external-event-summary" + types.AddKeyToken(summaryExternalEventSummary.EventId, "event-id")
    summaryExternalEventSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryExternalEventSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryExternalEventSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryExternalEventSummary.EntityData.Children = types.NewOrderedMap()
    summaryExternalEventSummary.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &summaryExternalEventSummary.IpConvergenceTime})
    summaryExternalEventSummary.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &summaryExternalEventSummary.MplsConvergenceTime})
    summaryExternalEventSummary.EntityData.Children.Append("path", types.YChild{"Path", nil})
    for i := range summaryExternalEventSummary.Path {
        summaryExternalEventSummary.EntityData.Children.Append(types.GetSegmentPath(summaryExternalEventSummary.Path[i]), types.YChild{"Path", summaryExternalEventSummary.Path[i]})
    }
    summaryExternalEventSummary.EntityData.Children.Append("trigger-lsa", types.YChild{"TriggerLsa", nil})
    for i := range summaryExternalEventSummary.TriggerLsa {
        summaryExternalEventSummary.EntityData.Children.Append(types.GetSegmentPath(summaryExternalEventSummary.TriggerLsa[i]), types.YChild{"TriggerLsa", summaryExternalEventSummary.TriggerLsa[i]})
    }
    summaryExternalEventSummary.EntityData.Children.Append("time-line", types.YChild{"TimeLine", nil})
    for i := range summaryExternalEventSummary.TimeLine {
        summaryExternalEventSummary.EntityData.Children.Append(types.GetSegmentPath(summaryExternalEventSummary.TimeLine[i]), types.YChild{"TimeLine", summaryExternalEventSummary.TimeLine[i]})
    }
    summaryExternalEventSummary.EntityData.Children.Append("lsa-processed", types.YChild{"LsaProcessed", nil})
    for i := range summaryExternalEventSummary.LsaProcessed {
        summaryExternalEventSummary.EntityData.Children.Append(types.GetSegmentPath(summaryExternalEventSummary.LsaProcessed[i]), types.YChild{"LsaProcessed", summaryExternalEventSummary.LsaProcessed[i]})
    }
    summaryExternalEventSummary.EntityData.Leafs = types.NewOrderedMap()
    summaryExternalEventSummary.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", summaryExternalEventSummary.EventId})
    summaryExternalEventSummary.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", summaryExternalEventSummary.Prefix})
    summaryExternalEventSummary.EntityData.Leafs.Append("prefix-lenth", types.YLeaf{"PrefixLenth", summaryExternalEventSummary.PrefixLenth})
    summaryExternalEventSummary.EntityData.Leafs.Append("spf-run-no", types.YLeaf{"SpfRunNo", summaryExternalEventSummary.SpfRunNo})
    summaryExternalEventSummary.EntityData.Leafs.Append("ipfrr-event-id", types.YLeaf{"IpfrrEventId", summaryExternalEventSummary.IpfrrEventId})
    summaryExternalEventSummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", summaryExternalEventSummary.ThresholdExceeded})
    summaryExternalEventSummary.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", summaryExternalEventSummary.Priority})
    summaryExternalEventSummary.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", summaryExternalEventSummary.ChangeType})
    summaryExternalEventSummary.EntityData.Leafs.Append("route-type", types.YLeaf{"RouteType", summaryExternalEventSummary.RouteType})
    summaryExternalEventSummary.EntityData.Leafs.Append("route-path-change-type", types.YLeaf{"RoutePathChangeType", summaryExternalEventSummary.RoutePathChangeType})
    summaryExternalEventSummary.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", summaryExternalEventSummary.Cost})
    summaryExternalEventSummary.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", summaryExternalEventSummary.TriggerTime})

    summaryExternalEventSummary.EntityData.YListKeys = []string {"EventId"}

    return &(summaryExternalEventSummary.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "summary-external-event-summary"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "summary-external-event-summary"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_Path
// Path information
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Path Metric. The type is interface{} with range: 0..4294967295.
    PathMetric interface{}

    // Backup Path Informatoin. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_Path_LfaPath.
    LfaPath []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_Path_LfaPath
}

func (path *Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "summary-external-event-summary"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Children.Append("lfa-path", types.YChild{"LfaPath", nil})
    for i := range path.LfaPath {
        path.EntityData.Children.Append(types.GetSegmentPath(path.LfaPath[i]), types.YChild{"LfaPath", path.LfaPath[i]})
    }
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", path.InterfaceName})
    path.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", path.NeighbourAddress})
    path.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", path.ChangeType})
    path.EntityData.Leafs.Append("path-metric", types.YLeaf{"PathMetric", path.PathMetric})

    path.EntityData.YListKeys = []string {}

    return &(path.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_Path_LfaPath
// Backup Path Informatoin
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_Path_LfaPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of LFA. The type is RcmdShowIpfrrLfa.
    LfaType interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Path Metric. The type is interface{} with range: 0..4294967295.
    PathMetric interface{}

    // Remote Node ID, in case of Remote LFA. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteNodeId interface{}
}

func (lfaPath *Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_Path_LfaPath) GetEntityData() *types.CommonEntityData {
    lfaPath.EntityData.YFilter = lfaPath.YFilter
    lfaPath.EntityData.YangName = "lfa-path"
    lfaPath.EntityData.BundleName = "cisco_ios_xr"
    lfaPath.EntityData.ParentYangName = "path"
    lfaPath.EntityData.SegmentPath = "lfa-path"
    lfaPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lfaPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lfaPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lfaPath.EntityData.Children = types.NewOrderedMap()
    lfaPath.EntityData.Leafs = types.NewOrderedMap()
    lfaPath.EntityData.Leafs.Append("lfa-type", types.YLeaf{"LfaType", lfaPath.LfaType})
    lfaPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", lfaPath.InterfaceName})
    lfaPath.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", lfaPath.NeighbourAddress})
    lfaPath.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lfaPath.ChangeType})
    lfaPath.EntityData.Leafs.Append("path-metric", types.YLeaf{"PathMetric", lfaPath.PathMetric})
    lfaPath.EntityData.Leafs.Append("remote-node-id", types.YLeaf{"RemoteNodeId", lfaPath.RemoteNodeId})

    lfaPath.EntityData.YListKeys = []string {}

    return &(lfaPath.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TriggerLsa
// LSA that triggered this event
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TriggerLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (triggerLsa *Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TriggerLsa) GetEntityData() *types.CommonEntityData {
    triggerLsa.EntityData.YFilter = triggerLsa.YFilter
    triggerLsa.EntityData.YangName = "trigger-lsa"
    triggerLsa.EntityData.BundleName = "cisco_ios_xr"
    triggerLsa.EntityData.ParentYangName = "summary-external-event-summary"
    triggerLsa.EntityData.SegmentPath = "trigger-lsa"
    triggerLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    triggerLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    triggerLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    triggerLsa.EntityData.Children = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", triggerLsa.LsaId})
    triggerLsa.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", triggerLsa.SequenceNumber})
    triggerLsa.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", triggerLsa.LsaType})
    triggerLsa.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", triggerLsa.OriginRouterId})
    triggerLsa.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", triggerLsa.ChangeType})
    triggerLsa.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", triggerLsa.ReceptionTime})

    triggerLsa.EntityData.YListKeys = []string {}

    return &(triggerLsa.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine
// Timeline information
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route origin (routing protocol). The type is string.
    RouteOrigin interface{}

    // Entry point of IPv4 RIB. The type is string.
    RiBv4Enter interface{}

    // Exit point from IPv4 RIB to FIBs. The type is string.
    RiBv4Exit interface{}

    // Route Redistribute point from IPv4 RIB to LDP. The type is string.
    RiBv4Redistribute interface{}

    // Entry point of LDP. The type is string.
    LdpEnter interface{}

    // Exit point of LDP to LSD. The type is string.
    LdpExit interface{}

    // Entry point of LSD. The type is string.
    LsdEnter interface{}

    // Exit point of LSD to FIBs. The type is string.
    LsdExit interface{}

    // List of Linecards' completion point for IP routes. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine_LcIp.
    LcIp []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine_LcIp

    // List of Linecards' completion point for MPLS labels. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine_LcMpls.
    LcMpls []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine_LcMpls
}

func (timeLine *Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine) GetEntityData() *types.CommonEntityData {
    timeLine.EntityData.YFilter = timeLine.YFilter
    timeLine.EntityData.YangName = "time-line"
    timeLine.EntityData.BundleName = "cisco_ios_xr"
    timeLine.EntityData.ParentYangName = "summary-external-event-summary"
    timeLine.EntityData.SegmentPath = "time-line"
    timeLine.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeLine.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeLine.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeLine.EntityData.Children = types.NewOrderedMap()
    timeLine.EntityData.Children.Append("lc-ip", types.YChild{"LcIp", nil})
    for i := range timeLine.LcIp {
        timeLine.EntityData.Children.Append(types.GetSegmentPath(timeLine.LcIp[i]), types.YChild{"LcIp", timeLine.LcIp[i]})
    }
    timeLine.EntityData.Children.Append("lc-mpls", types.YChild{"LcMpls", nil})
    for i := range timeLine.LcMpls {
        timeLine.EntityData.Children.Append(types.GetSegmentPath(timeLine.LcMpls[i]), types.YChild{"LcMpls", timeLine.LcMpls[i]})
    }
    timeLine.EntityData.Leafs = types.NewOrderedMap()
    timeLine.EntityData.Leafs.Append("route-origin", types.YLeaf{"RouteOrigin", timeLine.RouteOrigin})
    timeLine.EntityData.Leafs.Append("ri-bv4-enter", types.YLeaf{"RiBv4Enter", timeLine.RiBv4Enter})
    timeLine.EntityData.Leafs.Append("ri-bv4-exit", types.YLeaf{"RiBv4Exit", timeLine.RiBv4Exit})
    timeLine.EntityData.Leafs.Append("ri-bv4-redistribute", types.YLeaf{"RiBv4Redistribute", timeLine.RiBv4Redistribute})
    timeLine.EntityData.Leafs.Append("ldp-enter", types.YLeaf{"LdpEnter", timeLine.LdpEnter})
    timeLine.EntityData.Leafs.Append("ldp-exit", types.YLeaf{"LdpExit", timeLine.LdpExit})
    timeLine.EntityData.Leafs.Append("lsd-enter", types.YLeaf{"LsdEnter", timeLine.LsdEnter})
    timeLine.EntityData.Leafs.Append("lsd-exit", types.YLeaf{"LsdExit", timeLine.LsdExit})

    timeLine.EntityData.YListKeys = []string {}

    return &(timeLine.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine_LcIp
// List of Linecards' completion point for IP
// routes
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine_LcIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB. The type is string.
    FibComplete interface{}
}

func (lcIp *Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine_LcIp) GetEntityData() *types.CommonEntityData {
    lcIp.EntityData.YFilter = lcIp.YFilter
    lcIp.EntityData.YangName = "lc-ip"
    lcIp.EntityData.BundleName = "cisco_ios_xr"
    lcIp.EntityData.ParentYangName = "time-line"
    lcIp.EntityData.SegmentPath = "lc-ip"
    lcIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcIp.EntityData.Children = types.NewOrderedMap()
    lcIp.EntityData.Leafs = types.NewOrderedMap()
    lcIp.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcIp.NodeName})
    lcIp.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcIp.Speed})
    lcIp.EntityData.Leafs.Append("fib-complete", types.YLeaf{"FibComplete", lcIp.FibComplete})

    lcIp.EntityData.YListKeys = []string {}

    return &(lcIp.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine_LcMpls
// List of Linecards' completion point for MPLS
// labels
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine_LcMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB. The type is string.
    FibComplete interface{}
}

func (lcMpls *Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_TimeLine_LcMpls) GetEntityData() *types.CommonEntityData {
    lcMpls.EntityData.YFilter = lcMpls.YFilter
    lcMpls.EntityData.YangName = "lc-mpls"
    lcMpls.EntityData.BundleName = "cisco_ios_xr"
    lcMpls.EntityData.ParentYangName = "time-line"
    lcMpls.EntityData.SegmentPath = "lc-mpls"
    lcMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcMpls.EntityData.Children = types.NewOrderedMap()
    lcMpls.EntityData.Leafs = types.NewOrderedMap()
    lcMpls.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcMpls.NodeName})
    lcMpls.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcMpls.Speed})
    lcMpls.EntityData.Leafs.Append("fib-complete", types.YLeaf{"FibComplete", lcMpls.FibComplete})

    lcMpls.EntityData.YListKeys = []string {}

    return &(lcMpls.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_LsaProcessed
// List of LSAs processed
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_LsaProcessed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (lsaProcessed *Rcmd_Ospf_Instances_Instance_SummaryExternalEventSummaries_SummaryExternalEventSummary_LsaProcessed) GetEntityData() *types.CommonEntityData {
    lsaProcessed.EntityData.YFilter = lsaProcessed.YFilter
    lsaProcessed.EntityData.YangName = "lsa-processed"
    lsaProcessed.EntityData.BundleName = "cisco_ios_xr"
    lsaProcessed.EntityData.ParentYangName = "summary-external-event-summary"
    lsaProcessed.EntityData.SegmentPath = "lsa-processed"
    lsaProcessed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsaProcessed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsaProcessed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsaProcessed.EntityData.Children = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", lsaProcessed.LsaId})
    lsaProcessed.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", lsaProcessed.SequenceNumber})
    lsaProcessed.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", lsaProcessed.LsaType})
    lsaProcessed.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", lsaProcessed.OriginRouterId})
    lsaProcessed.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lsaProcessed.ChangeType})
    lsaProcessed.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", lsaProcessed.ReceptionTime})

    lsaProcessed.EntityData.YListKeys = []string {}

    return &(lsaProcessed.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventSummaries
// OSPF Prefix events summary data
type Rcmd_Ospf_Instances_Instance_PrefixEventSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF Prefix Event data. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary.
    PrefixEventSummary []*Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary
}

func (prefixEventSummaries *Rcmd_Ospf_Instances_Instance_PrefixEventSummaries) GetEntityData() *types.CommonEntityData {
    prefixEventSummaries.EntityData.YFilter = prefixEventSummaries.YFilter
    prefixEventSummaries.EntityData.YangName = "prefix-event-summaries"
    prefixEventSummaries.EntityData.BundleName = "cisco_ios_xr"
    prefixEventSummaries.EntityData.ParentYangName = "instance"
    prefixEventSummaries.EntityData.SegmentPath = "prefix-event-summaries"
    prefixEventSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixEventSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixEventSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixEventSummaries.EntityData.Children = types.NewOrderedMap()
    prefixEventSummaries.EntityData.Children.Append("prefix-event-summary", types.YChild{"PrefixEventSummary", nil})
    for i := range prefixEventSummaries.PrefixEventSummary {
        prefixEventSummaries.EntityData.Children.Append(types.GetSegmentPath(prefixEventSummaries.PrefixEventSummary[i]), types.YChild{"PrefixEventSummary", prefixEventSummaries.PrefixEventSummary[i]})
    }
    prefixEventSummaries.EntityData.Leafs = types.NewOrderedMap()

    prefixEventSummaries.EntityData.YListKeys = []string {}

    return &(prefixEventSummaries.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary
// OSPF Prefix Event data
type Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Event ID. The type is interface{} with
    // range: 1..4294967295.
    EventId interface{}

    // Prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLenth interface{}

    // Referenced SPF Run No (0 - Not Applicable). The type is interface{} with
    // range: 0..4294967295.
    SpfRunNo interface{}

    // Referenced IP-FRR Event ID (0 - Not Applicable). The type is interface{}
    // with range: 0..4294967295.
    IpfrrEventId interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Event processed priority. The type is RcmdPriorityLevel.
    Priority interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Route Type intra/inter/l1/l2. The type is RcmdShowRoute.
    RouteType interface{}

    // Route Path Change Type. The type is RcmdShowRoutePathChange.
    RoutePathChangeType interface{}

    // Protocol route cost. The type is interface{} with range: 0..4294967295.
    Cost interface{}

    // Event trigger time. The type is string.
    TriggerTime interface{}

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_MplsConvergenceTime

    // Path information. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path.
    Path []*Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path

    // LSA that triggered this event. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TriggerLsa.
    TriggerLsa []*Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TriggerLsa

    // Timeline information. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine.
    TimeLine []*Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine

    // List of LSAs processed. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_LsaProcessed.
    LsaProcessed []*Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_LsaProcessed
}

func (prefixEventSummary *Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary) GetEntityData() *types.CommonEntityData {
    prefixEventSummary.EntityData.YFilter = prefixEventSummary.YFilter
    prefixEventSummary.EntityData.YangName = "prefix-event-summary"
    prefixEventSummary.EntityData.BundleName = "cisco_ios_xr"
    prefixEventSummary.EntityData.ParentYangName = "prefix-event-summaries"
    prefixEventSummary.EntityData.SegmentPath = "prefix-event-summary" + types.AddKeyToken(prefixEventSummary.EventId, "event-id")
    prefixEventSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixEventSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixEventSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixEventSummary.EntityData.Children = types.NewOrderedMap()
    prefixEventSummary.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prefixEventSummary.IpConvergenceTime})
    prefixEventSummary.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prefixEventSummary.MplsConvergenceTime})
    prefixEventSummary.EntityData.Children.Append("path", types.YChild{"Path", nil})
    for i := range prefixEventSummary.Path {
        prefixEventSummary.EntityData.Children.Append(types.GetSegmentPath(prefixEventSummary.Path[i]), types.YChild{"Path", prefixEventSummary.Path[i]})
    }
    prefixEventSummary.EntityData.Children.Append("trigger-lsa", types.YChild{"TriggerLsa", nil})
    for i := range prefixEventSummary.TriggerLsa {
        prefixEventSummary.EntityData.Children.Append(types.GetSegmentPath(prefixEventSummary.TriggerLsa[i]), types.YChild{"TriggerLsa", prefixEventSummary.TriggerLsa[i]})
    }
    prefixEventSummary.EntityData.Children.Append("time-line", types.YChild{"TimeLine", nil})
    for i := range prefixEventSummary.TimeLine {
        prefixEventSummary.EntityData.Children.Append(types.GetSegmentPath(prefixEventSummary.TimeLine[i]), types.YChild{"TimeLine", prefixEventSummary.TimeLine[i]})
    }
    prefixEventSummary.EntityData.Children.Append("lsa-processed", types.YChild{"LsaProcessed", nil})
    for i := range prefixEventSummary.LsaProcessed {
        prefixEventSummary.EntityData.Children.Append(types.GetSegmentPath(prefixEventSummary.LsaProcessed[i]), types.YChild{"LsaProcessed", prefixEventSummary.LsaProcessed[i]})
    }
    prefixEventSummary.EntityData.Leafs = types.NewOrderedMap()
    prefixEventSummary.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", prefixEventSummary.EventId})
    prefixEventSummary.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", prefixEventSummary.Prefix})
    prefixEventSummary.EntityData.Leafs.Append("prefix-lenth", types.YLeaf{"PrefixLenth", prefixEventSummary.PrefixLenth})
    prefixEventSummary.EntityData.Leafs.Append("spf-run-no", types.YLeaf{"SpfRunNo", prefixEventSummary.SpfRunNo})
    prefixEventSummary.EntityData.Leafs.Append("ipfrr-event-id", types.YLeaf{"IpfrrEventId", prefixEventSummary.IpfrrEventId})
    prefixEventSummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prefixEventSummary.ThresholdExceeded})
    prefixEventSummary.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", prefixEventSummary.Priority})
    prefixEventSummary.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", prefixEventSummary.ChangeType})
    prefixEventSummary.EntityData.Leafs.Append("route-type", types.YLeaf{"RouteType", prefixEventSummary.RouteType})
    prefixEventSummary.EntityData.Leafs.Append("route-path-change-type", types.YLeaf{"RoutePathChangeType", prefixEventSummary.RoutePathChangeType})
    prefixEventSummary.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", prefixEventSummary.Cost})
    prefixEventSummary.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", prefixEventSummary.TriggerTime})

    prefixEventSummary.EntityData.YListKeys = []string {"EventId"}

    return &(prefixEventSummary.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "prefix-event-summary"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "prefix-event-summary"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path
// Path information
type Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Path Metric. The type is interface{} with range: 0..4294967295.
    PathMetric interface{}

    // Backup Path Informatoin. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path_LfaPath.
    LfaPath []*Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path_LfaPath
}

func (path *Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "prefix-event-summary"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Children.Append("lfa-path", types.YChild{"LfaPath", nil})
    for i := range path.LfaPath {
        path.EntityData.Children.Append(types.GetSegmentPath(path.LfaPath[i]), types.YChild{"LfaPath", path.LfaPath[i]})
    }
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", path.InterfaceName})
    path.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", path.NeighbourAddress})
    path.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", path.ChangeType})
    path.EntityData.Leafs.Append("path-metric", types.YLeaf{"PathMetric", path.PathMetric})

    path.EntityData.YListKeys = []string {}

    return &(path.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path_LfaPath
// Backup Path Informatoin
type Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path_LfaPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of LFA. The type is RcmdShowIpfrrLfa.
    LfaType interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Path Metric. The type is interface{} with range: 0..4294967295.
    PathMetric interface{}

    // Remote Node ID, in case of Remote LFA. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteNodeId interface{}
}

func (lfaPath *Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path_LfaPath) GetEntityData() *types.CommonEntityData {
    lfaPath.EntityData.YFilter = lfaPath.YFilter
    lfaPath.EntityData.YangName = "lfa-path"
    lfaPath.EntityData.BundleName = "cisco_ios_xr"
    lfaPath.EntityData.ParentYangName = "path"
    lfaPath.EntityData.SegmentPath = "lfa-path"
    lfaPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lfaPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lfaPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lfaPath.EntityData.Children = types.NewOrderedMap()
    lfaPath.EntityData.Leafs = types.NewOrderedMap()
    lfaPath.EntityData.Leafs.Append("lfa-type", types.YLeaf{"LfaType", lfaPath.LfaType})
    lfaPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", lfaPath.InterfaceName})
    lfaPath.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", lfaPath.NeighbourAddress})
    lfaPath.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lfaPath.ChangeType})
    lfaPath.EntityData.Leafs.Append("path-metric", types.YLeaf{"PathMetric", lfaPath.PathMetric})
    lfaPath.EntityData.Leafs.Append("remote-node-id", types.YLeaf{"RemoteNodeId", lfaPath.RemoteNodeId})

    lfaPath.EntityData.YListKeys = []string {}

    return &(lfaPath.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TriggerLsa
// LSA that triggered this event
type Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TriggerLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (triggerLsa *Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TriggerLsa) GetEntityData() *types.CommonEntityData {
    triggerLsa.EntityData.YFilter = triggerLsa.YFilter
    triggerLsa.EntityData.YangName = "trigger-lsa"
    triggerLsa.EntityData.BundleName = "cisco_ios_xr"
    triggerLsa.EntityData.ParentYangName = "prefix-event-summary"
    triggerLsa.EntityData.SegmentPath = "trigger-lsa"
    triggerLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    triggerLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    triggerLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    triggerLsa.EntityData.Children = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", triggerLsa.LsaId})
    triggerLsa.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", triggerLsa.SequenceNumber})
    triggerLsa.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", triggerLsa.LsaType})
    triggerLsa.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", triggerLsa.OriginRouterId})
    triggerLsa.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", triggerLsa.ChangeType})
    triggerLsa.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", triggerLsa.ReceptionTime})

    triggerLsa.EntityData.YListKeys = []string {}

    return &(triggerLsa.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine
// Timeline information
type Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route origin (routing protocol). The type is string.
    RouteOrigin interface{}

    // Entry point of IPv4 RIB. The type is string.
    RiBv4Enter interface{}

    // Exit point from IPv4 RIB to FIBs. The type is string.
    RiBv4Exit interface{}

    // Route Redistribute point from IPv4 RIB to LDP. The type is string.
    RiBv4Redistribute interface{}

    // Entry point of LDP. The type is string.
    LdpEnter interface{}

    // Exit point of LDP to LSD. The type is string.
    LdpExit interface{}

    // Entry point of LSD. The type is string.
    LsdEnter interface{}

    // Exit point of LSD to FIBs. The type is string.
    LsdExit interface{}

    // List of Linecards' completion point for IP routes. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcIp.
    LcIp []*Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcIp

    // List of Linecards' completion point for MPLS labels. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcMpls.
    LcMpls []*Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcMpls
}

func (timeLine *Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine) GetEntityData() *types.CommonEntityData {
    timeLine.EntityData.YFilter = timeLine.YFilter
    timeLine.EntityData.YangName = "time-line"
    timeLine.EntityData.BundleName = "cisco_ios_xr"
    timeLine.EntityData.ParentYangName = "prefix-event-summary"
    timeLine.EntityData.SegmentPath = "time-line"
    timeLine.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeLine.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeLine.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeLine.EntityData.Children = types.NewOrderedMap()
    timeLine.EntityData.Children.Append("lc-ip", types.YChild{"LcIp", nil})
    for i := range timeLine.LcIp {
        timeLine.EntityData.Children.Append(types.GetSegmentPath(timeLine.LcIp[i]), types.YChild{"LcIp", timeLine.LcIp[i]})
    }
    timeLine.EntityData.Children.Append("lc-mpls", types.YChild{"LcMpls", nil})
    for i := range timeLine.LcMpls {
        timeLine.EntityData.Children.Append(types.GetSegmentPath(timeLine.LcMpls[i]), types.YChild{"LcMpls", timeLine.LcMpls[i]})
    }
    timeLine.EntityData.Leafs = types.NewOrderedMap()
    timeLine.EntityData.Leafs.Append("route-origin", types.YLeaf{"RouteOrigin", timeLine.RouteOrigin})
    timeLine.EntityData.Leafs.Append("ri-bv4-enter", types.YLeaf{"RiBv4Enter", timeLine.RiBv4Enter})
    timeLine.EntityData.Leafs.Append("ri-bv4-exit", types.YLeaf{"RiBv4Exit", timeLine.RiBv4Exit})
    timeLine.EntityData.Leafs.Append("ri-bv4-redistribute", types.YLeaf{"RiBv4Redistribute", timeLine.RiBv4Redistribute})
    timeLine.EntityData.Leafs.Append("ldp-enter", types.YLeaf{"LdpEnter", timeLine.LdpEnter})
    timeLine.EntityData.Leafs.Append("ldp-exit", types.YLeaf{"LdpExit", timeLine.LdpExit})
    timeLine.EntityData.Leafs.Append("lsd-enter", types.YLeaf{"LsdEnter", timeLine.LsdEnter})
    timeLine.EntityData.Leafs.Append("lsd-exit", types.YLeaf{"LsdExit", timeLine.LsdExit})

    timeLine.EntityData.YListKeys = []string {}

    return &(timeLine.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcIp
// List of Linecards' completion point for IP
// routes
type Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB. The type is string.
    FibComplete interface{}
}

func (lcIp *Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcIp) GetEntityData() *types.CommonEntityData {
    lcIp.EntityData.YFilter = lcIp.YFilter
    lcIp.EntityData.YangName = "lc-ip"
    lcIp.EntityData.BundleName = "cisco_ios_xr"
    lcIp.EntityData.ParentYangName = "time-line"
    lcIp.EntityData.SegmentPath = "lc-ip"
    lcIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcIp.EntityData.Children = types.NewOrderedMap()
    lcIp.EntityData.Leafs = types.NewOrderedMap()
    lcIp.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcIp.NodeName})
    lcIp.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcIp.Speed})
    lcIp.EntityData.Leafs.Append("fib-complete", types.YLeaf{"FibComplete", lcIp.FibComplete})

    lcIp.EntityData.YListKeys = []string {}

    return &(lcIp.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcMpls
// List of Linecards' completion point for MPLS
// labels
type Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB. The type is string.
    FibComplete interface{}
}

func (lcMpls *Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcMpls) GetEntityData() *types.CommonEntityData {
    lcMpls.EntityData.YFilter = lcMpls.YFilter
    lcMpls.EntityData.YangName = "lc-mpls"
    lcMpls.EntityData.BundleName = "cisco_ios_xr"
    lcMpls.EntityData.ParentYangName = "time-line"
    lcMpls.EntityData.SegmentPath = "lc-mpls"
    lcMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcMpls.EntityData.Children = types.NewOrderedMap()
    lcMpls.EntityData.Leafs = types.NewOrderedMap()
    lcMpls.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcMpls.NodeName})
    lcMpls.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcMpls.Speed})
    lcMpls.EntityData.Leafs.Append("fib-complete", types.YLeaf{"FibComplete", lcMpls.FibComplete})

    lcMpls.EntityData.YListKeys = []string {}

    return &(lcMpls.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_LsaProcessed
// List of LSAs processed
type Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_LsaProcessed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (lsaProcessed *Rcmd_Ospf_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_LsaProcessed) GetEntityData() *types.CommonEntityData {
    lsaProcessed.EntityData.YFilter = lsaProcessed.YFilter
    lsaProcessed.EntityData.YangName = "lsa-processed"
    lsaProcessed.EntityData.BundleName = "cisco_ios_xr"
    lsaProcessed.EntityData.ParentYangName = "prefix-event-summary"
    lsaProcessed.EntityData.SegmentPath = "lsa-processed"
    lsaProcessed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsaProcessed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsaProcessed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsaProcessed.EntityData.Children = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", lsaProcessed.LsaId})
    lsaProcessed.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", lsaProcessed.SequenceNumber})
    lsaProcessed.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", lsaProcessed.LsaType})
    lsaProcessed.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", lsaProcessed.OriginRouterId})
    lsaProcessed.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lsaProcessed.ChangeType})
    lsaProcessed.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", lsaProcessed.ReceptionTime})

    lsaProcessed.EntityData.YListKeys = []string {}

    return &(lsaProcessed.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines
// OSPF Summary-External Prefix events offline
// data
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Offline operational data for particular OSPF Prefix Event. The type is
    // slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline.
    SummaryExternalEventOffline []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline
}

func (summaryExternalEventOfflines *Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines) GetEntityData() *types.CommonEntityData {
    summaryExternalEventOfflines.EntityData.YFilter = summaryExternalEventOfflines.YFilter
    summaryExternalEventOfflines.EntityData.YangName = "summary-external-event-offlines"
    summaryExternalEventOfflines.EntityData.BundleName = "cisco_ios_xr"
    summaryExternalEventOfflines.EntityData.ParentYangName = "instance"
    summaryExternalEventOfflines.EntityData.SegmentPath = "summary-external-event-offlines"
    summaryExternalEventOfflines.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryExternalEventOfflines.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryExternalEventOfflines.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryExternalEventOfflines.EntityData.Children = types.NewOrderedMap()
    summaryExternalEventOfflines.EntityData.Children.Append("summary-external-event-offline", types.YChild{"SummaryExternalEventOffline", nil})
    for i := range summaryExternalEventOfflines.SummaryExternalEventOffline {
        summaryExternalEventOfflines.EntityData.Children.Append(types.GetSegmentPath(summaryExternalEventOfflines.SummaryExternalEventOffline[i]), types.YChild{"SummaryExternalEventOffline", summaryExternalEventOfflines.SummaryExternalEventOffline[i]})
    }
    summaryExternalEventOfflines.EntityData.Leafs = types.NewOrderedMap()

    summaryExternalEventOfflines.EntityData.YListKeys = []string {}

    return &(summaryExternalEventOfflines.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline
// Offline operational data for particular OSPF
// Prefix Event
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Event ID. The type is interface{} with
    // range: 1..4294967295.
    EventId interface{}

    // Prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLenth interface{}

    // Referenced SPF Run No (0 - Not Applicable). The type is interface{} with
    // range: 0..4294967295.
    SpfRunNo interface{}

    // Referenced IP-FRR Event ID (0 - Not Applicable). The type is interface{}
    // with range: 0..4294967295.
    IpfrrEventId interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Event processed priority. The type is RcmdPriorityLevel.
    Priority interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Route Type intra/inter/l1/l2. The type is RcmdShowRoute.
    RouteType interface{}

    // Route Path Change Type. The type is RcmdShowRoutePathChange.
    RoutePathChangeType interface{}

    // Protocol route cost. The type is interface{} with range: 0..4294967295.
    Cost interface{}

    // Event trigger time. The type is string.
    TriggerTime interface{}

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_MplsConvergenceTime

    // Path information. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_Path.
    Path []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_Path

    // LSA that triggered this event. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TriggerLsa.
    TriggerLsa []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TriggerLsa

    // Timeline information. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine.
    TimeLine []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine

    // List of LSAs processed. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_LsaProcessed.
    LsaProcessed []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_LsaProcessed
}

func (summaryExternalEventOffline *Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline) GetEntityData() *types.CommonEntityData {
    summaryExternalEventOffline.EntityData.YFilter = summaryExternalEventOffline.YFilter
    summaryExternalEventOffline.EntityData.YangName = "summary-external-event-offline"
    summaryExternalEventOffline.EntityData.BundleName = "cisco_ios_xr"
    summaryExternalEventOffline.EntityData.ParentYangName = "summary-external-event-offlines"
    summaryExternalEventOffline.EntityData.SegmentPath = "summary-external-event-offline" + types.AddKeyToken(summaryExternalEventOffline.EventId, "event-id")
    summaryExternalEventOffline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryExternalEventOffline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryExternalEventOffline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryExternalEventOffline.EntityData.Children = types.NewOrderedMap()
    summaryExternalEventOffline.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &summaryExternalEventOffline.IpConvergenceTime})
    summaryExternalEventOffline.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &summaryExternalEventOffline.MplsConvergenceTime})
    summaryExternalEventOffline.EntityData.Children.Append("path", types.YChild{"Path", nil})
    for i := range summaryExternalEventOffline.Path {
        summaryExternalEventOffline.EntityData.Children.Append(types.GetSegmentPath(summaryExternalEventOffline.Path[i]), types.YChild{"Path", summaryExternalEventOffline.Path[i]})
    }
    summaryExternalEventOffline.EntityData.Children.Append("trigger-lsa", types.YChild{"TriggerLsa", nil})
    for i := range summaryExternalEventOffline.TriggerLsa {
        summaryExternalEventOffline.EntityData.Children.Append(types.GetSegmentPath(summaryExternalEventOffline.TriggerLsa[i]), types.YChild{"TriggerLsa", summaryExternalEventOffline.TriggerLsa[i]})
    }
    summaryExternalEventOffline.EntityData.Children.Append("time-line", types.YChild{"TimeLine", nil})
    for i := range summaryExternalEventOffline.TimeLine {
        summaryExternalEventOffline.EntityData.Children.Append(types.GetSegmentPath(summaryExternalEventOffline.TimeLine[i]), types.YChild{"TimeLine", summaryExternalEventOffline.TimeLine[i]})
    }
    summaryExternalEventOffline.EntityData.Children.Append("lsa-processed", types.YChild{"LsaProcessed", nil})
    for i := range summaryExternalEventOffline.LsaProcessed {
        summaryExternalEventOffline.EntityData.Children.Append(types.GetSegmentPath(summaryExternalEventOffline.LsaProcessed[i]), types.YChild{"LsaProcessed", summaryExternalEventOffline.LsaProcessed[i]})
    }
    summaryExternalEventOffline.EntityData.Leafs = types.NewOrderedMap()
    summaryExternalEventOffline.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", summaryExternalEventOffline.EventId})
    summaryExternalEventOffline.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", summaryExternalEventOffline.Prefix})
    summaryExternalEventOffline.EntityData.Leafs.Append("prefix-lenth", types.YLeaf{"PrefixLenth", summaryExternalEventOffline.PrefixLenth})
    summaryExternalEventOffline.EntityData.Leafs.Append("spf-run-no", types.YLeaf{"SpfRunNo", summaryExternalEventOffline.SpfRunNo})
    summaryExternalEventOffline.EntityData.Leafs.Append("ipfrr-event-id", types.YLeaf{"IpfrrEventId", summaryExternalEventOffline.IpfrrEventId})
    summaryExternalEventOffline.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", summaryExternalEventOffline.ThresholdExceeded})
    summaryExternalEventOffline.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", summaryExternalEventOffline.Priority})
    summaryExternalEventOffline.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", summaryExternalEventOffline.ChangeType})
    summaryExternalEventOffline.EntityData.Leafs.Append("route-type", types.YLeaf{"RouteType", summaryExternalEventOffline.RouteType})
    summaryExternalEventOffline.EntityData.Leafs.Append("route-path-change-type", types.YLeaf{"RoutePathChangeType", summaryExternalEventOffline.RoutePathChangeType})
    summaryExternalEventOffline.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", summaryExternalEventOffline.Cost})
    summaryExternalEventOffline.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", summaryExternalEventOffline.TriggerTime})

    summaryExternalEventOffline.EntityData.YListKeys = []string {"EventId"}

    return &(summaryExternalEventOffline.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "summary-external-event-offline"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "summary-external-event-offline"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_Path
// Path information
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Path Metric. The type is interface{} with range: 0..4294967295.
    PathMetric interface{}

    // Backup Path Informatoin. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_Path_LfaPath.
    LfaPath []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_Path_LfaPath
}

func (path *Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "summary-external-event-offline"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Children.Append("lfa-path", types.YChild{"LfaPath", nil})
    for i := range path.LfaPath {
        path.EntityData.Children.Append(types.GetSegmentPath(path.LfaPath[i]), types.YChild{"LfaPath", path.LfaPath[i]})
    }
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", path.InterfaceName})
    path.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", path.NeighbourAddress})
    path.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", path.ChangeType})
    path.EntityData.Leafs.Append("path-metric", types.YLeaf{"PathMetric", path.PathMetric})

    path.EntityData.YListKeys = []string {}

    return &(path.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_Path_LfaPath
// Backup Path Informatoin
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_Path_LfaPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of LFA. The type is RcmdShowIpfrrLfa.
    LfaType interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Path Metric. The type is interface{} with range: 0..4294967295.
    PathMetric interface{}

    // Remote Node ID, in case of Remote LFA. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteNodeId interface{}
}

func (lfaPath *Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_Path_LfaPath) GetEntityData() *types.CommonEntityData {
    lfaPath.EntityData.YFilter = lfaPath.YFilter
    lfaPath.EntityData.YangName = "lfa-path"
    lfaPath.EntityData.BundleName = "cisco_ios_xr"
    lfaPath.EntityData.ParentYangName = "path"
    lfaPath.EntityData.SegmentPath = "lfa-path"
    lfaPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lfaPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lfaPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lfaPath.EntityData.Children = types.NewOrderedMap()
    lfaPath.EntityData.Leafs = types.NewOrderedMap()
    lfaPath.EntityData.Leafs.Append("lfa-type", types.YLeaf{"LfaType", lfaPath.LfaType})
    lfaPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", lfaPath.InterfaceName})
    lfaPath.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", lfaPath.NeighbourAddress})
    lfaPath.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lfaPath.ChangeType})
    lfaPath.EntityData.Leafs.Append("path-metric", types.YLeaf{"PathMetric", lfaPath.PathMetric})
    lfaPath.EntityData.Leafs.Append("remote-node-id", types.YLeaf{"RemoteNodeId", lfaPath.RemoteNodeId})

    lfaPath.EntityData.YListKeys = []string {}

    return &(lfaPath.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TriggerLsa
// LSA that triggered this event
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TriggerLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (triggerLsa *Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TriggerLsa) GetEntityData() *types.CommonEntityData {
    triggerLsa.EntityData.YFilter = triggerLsa.YFilter
    triggerLsa.EntityData.YangName = "trigger-lsa"
    triggerLsa.EntityData.BundleName = "cisco_ios_xr"
    triggerLsa.EntityData.ParentYangName = "summary-external-event-offline"
    triggerLsa.EntityData.SegmentPath = "trigger-lsa"
    triggerLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    triggerLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    triggerLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    triggerLsa.EntityData.Children = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", triggerLsa.LsaId})
    triggerLsa.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", triggerLsa.SequenceNumber})
    triggerLsa.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", triggerLsa.LsaType})
    triggerLsa.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", triggerLsa.OriginRouterId})
    triggerLsa.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", triggerLsa.ChangeType})
    triggerLsa.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", triggerLsa.ReceptionTime})

    triggerLsa.EntityData.YListKeys = []string {}

    return &(triggerLsa.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine
// Timeline information
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route origin (routing protocol). The type is string.
    RouteOrigin interface{}

    // Entry point of IPv4 RIB. The type is string.
    RiBv4Enter interface{}

    // Exit point from IPv4 RIB to FIBs. The type is string.
    RiBv4Exit interface{}

    // Route Redistribute point from IPv4 RIB to LDP. The type is string.
    RiBv4Redistribute interface{}

    // Entry point of LDP. The type is string.
    LdpEnter interface{}

    // Exit point of LDP to LSD. The type is string.
    LdpExit interface{}

    // Entry point of LSD. The type is string.
    LsdEnter interface{}

    // Exit point of LSD to FIBs. The type is string.
    LsdExit interface{}

    // List of Linecards' completion point for IP routes. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine_LcIp.
    LcIp []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine_LcIp

    // List of Linecards' completion point for MPLS labels. The type is slice of
    // Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine_LcMpls.
    LcMpls []*Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine_LcMpls
}

func (timeLine *Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine) GetEntityData() *types.CommonEntityData {
    timeLine.EntityData.YFilter = timeLine.YFilter
    timeLine.EntityData.YangName = "time-line"
    timeLine.EntityData.BundleName = "cisco_ios_xr"
    timeLine.EntityData.ParentYangName = "summary-external-event-offline"
    timeLine.EntityData.SegmentPath = "time-line"
    timeLine.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeLine.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeLine.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeLine.EntityData.Children = types.NewOrderedMap()
    timeLine.EntityData.Children.Append("lc-ip", types.YChild{"LcIp", nil})
    for i := range timeLine.LcIp {
        timeLine.EntityData.Children.Append(types.GetSegmentPath(timeLine.LcIp[i]), types.YChild{"LcIp", timeLine.LcIp[i]})
    }
    timeLine.EntityData.Children.Append("lc-mpls", types.YChild{"LcMpls", nil})
    for i := range timeLine.LcMpls {
        timeLine.EntityData.Children.Append(types.GetSegmentPath(timeLine.LcMpls[i]), types.YChild{"LcMpls", timeLine.LcMpls[i]})
    }
    timeLine.EntityData.Leafs = types.NewOrderedMap()
    timeLine.EntityData.Leafs.Append("route-origin", types.YLeaf{"RouteOrigin", timeLine.RouteOrigin})
    timeLine.EntityData.Leafs.Append("ri-bv4-enter", types.YLeaf{"RiBv4Enter", timeLine.RiBv4Enter})
    timeLine.EntityData.Leafs.Append("ri-bv4-exit", types.YLeaf{"RiBv4Exit", timeLine.RiBv4Exit})
    timeLine.EntityData.Leafs.Append("ri-bv4-redistribute", types.YLeaf{"RiBv4Redistribute", timeLine.RiBv4Redistribute})
    timeLine.EntityData.Leafs.Append("ldp-enter", types.YLeaf{"LdpEnter", timeLine.LdpEnter})
    timeLine.EntityData.Leafs.Append("ldp-exit", types.YLeaf{"LdpExit", timeLine.LdpExit})
    timeLine.EntityData.Leafs.Append("lsd-enter", types.YLeaf{"LsdEnter", timeLine.LsdEnter})
    timeLine.EntityData.Leafs.Append("lsd-exit", types.YLeaf{"LsdExit", timeLine.LsdExit})

    timeLine.EntityData.YListKeys = []string {}

    return &(timeLine.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine_LcIp
// List of Linecards' completion point for IP
// routes
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine_LcIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB. The type is string.
    FibComplete interface{}
}

func (lcIp *Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine_LcIp) GetEntityData() *types.CommonEntityData {
    lcIp.EntityData.YFilter = lcIp.YFilter
    lcIp.EntityData.YangName = "lc-ip"
    lcIp.EntityData.BundleName = "cisco_ios_xr"
    lcIp.EntityData.ParentYangName = "time-line"
    lcIp.EntityData.SegmentPath = "lc-ip"
    lcIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcIp.EntityData.Children = types.NewOrderedMap()
    lcIp.EntityData.Leafs = types.NewOrderedMap()
    lcIp.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcIp.NodeName})
    lcIp.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcIp.Speed})
    lcIp.EntityData.Leafs.Append("fib-complete", types.YLeaf{"FibComplete", lcIp.FibComplete})

    lcIp.EntityData.YListKeys = []string {}

    return &(lcIp.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine_LcMpls
// List of Linecards' completion point for MPLS
// labels
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine_LcMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB. The type is string.
    FibComplete interface{}
}

func (lcMpls *Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_TimeLine_LcMpls) GetEntityData() *types.CommonEntityData {
    lcMpls.EntityData.YFilter = lcMpls.YFilter
    lcMpls.EntityData.YangName = "lc-mpls"
    lcMpls.EntityData.BundleName = "cisco_ios_xr"
    lcMpls.EntityData.ParentYangName = "time-line"
    lcMpls.EntityData.SegmentPath = "lc-mpls"
    lcMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcMpls.EntityData.Children = types.NewOrderedMap()
    lcMpls.EntityData.Leafs = types.NewOrderedMap()
    lcMpls.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcMpls.NodeName})
    lcMpls.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcMpls.Speed})
    lcMpls.EntityData.Leafs.Append("fib-complete", types.YLeaf{"FibComplete", lcMpls.FibComplete})

    lcMpls.EntityData.YListKeys = []string {}

    return &(lcMpls.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_LsaProcessed
// List of LSAs processed
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_LsaProcessed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (lsaProcessed *Rcmd_Ospf_Instances_Instance_SummaryExternalEventOfflines_SummaryExternalEventOffline_LsaProcessed) GetEntityData() *types.CommonEntityData {
    lsaProcessed.EntityData.YFilter = lsaProcessed.YFilter
    lsaProcessed.EntityData.YangName = "lsa-processed"
    lsaProcessed.EntityData.BundleName = "cisco_ios_xr"
    lsaProcessed.EntityData.ParentYangName = "summary-external-event-offline"
    lsaProcessed.EntityData.SegmentPath = "lsa-processed"
    lsaProcessed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsaProcessed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsaProcessed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsaProcessed.EntityData.Children = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", lsaProcessed.LsaId})
    lsaProcessed.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", lsaProcessed.SequenceNumber})
    lsaProcessed.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", lsaProcessed.LsaType})
    lsaProcessed.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", lsaProcessed.OriginRouterId})
    lsaProcessed.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lsaProcessed.ChangeType})
    lsaProcessed.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", lsaProcessed.ReceptionTime})

    lsaProcessed.EntityData.YListKeys = []string {}

    return &(lsaProcessed.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventOfflines
// OSPF Prefix events offline data
type Rcmd_Ospf_Instances_Instance_PrefixEventOfflines struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Offline operational data for particular OSPF Prefix Event. The type is
    // slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline.
    PrefixEventOffline []*Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline
}

func (prefixEventOfflines *Rcmd_Ospf_Instances_Instance_PrefixEventOfflines) GetEntityData() *types.CommonEntityData {
    prefixEventOfflines.EntityData.YFilter = prefixEventOfflines.YFilter
    prefixEventOfflines.EntityData.YangName = "prefix-event-offlines"
    prefixEventOfflines.EntityData.BundleName = "cisco_ios_xr"
    prefixEventOfflines.EntityData.ParentYangName = "instance"
    prefixEventOfflines.EntityData.SegmentPath = "prefix-event-offlines"
    prefixEventOfflines.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixEventOfflines.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixEventOfflines.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixEventOfflines.EntityData.Children = types.NewOrderedMap()
    prefixEventOfflines.EntityData.Children.Append("prefix-event-offline", types.YChild{"PrefixEventOffline", nil})
    for i := range prefixEventOfflines.PrefixEventOffline {
        prefixEventOfflines.EntityData.Children.Append(types.GetSegmentPath(prefixEventOfflines.PrefixEventOffline[i]), types.YChild{"PrefixEventOffline", prefixEventOfflines.PrefixEventOffline[i]})
    }
    prefixEventOfflines.EntityData.Leafs = types.NewOrderedMap()

    prefixEventOfflines.EntityData.YListKeys = []string {}

    return &(prefixEventOfflines.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline
// Offline operational data for particular OSPF
// Prefix Event
type Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Event ID. The type is interface{} with
    // range: 1..4294967295.
    EventId interface{}

    // Prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLenth interface{}

    // Referenced SPF Run No (0 - Not Applicable). The type is interface{} with
    // range: 0..4294967295.
    SpfRunNo interface{}

    // Referenced IP-FRR Event ID (0 - Not Applicable). The type is interface{}
    // with range: 0..4294967295.
    IpfrrEventId interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Event processed priority. The type is RcmdPriorityLevel.
    Priority interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Route Type intra/inter/l1/l2. The type is RcmdShowRoute.
    RouteType interface{}

    // Route Path Change Type. The type is RcmdShowRoutePathChange.
    RoutePathChangeType interface{}

    // Protocol route cost. The type is interface{} with range: 0..4294967295.
    Cost interface{}

    // Event trigger time. The type is string.
    TriggerTime interface{}

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_MplsConvergenceTime

    // Path information. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path.
    Path []*Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path

    // LSA that triggered this event. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TriggerLsa.
    TriggerLsa []*Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TriggerLsa

    // Timeline information. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine.
    TimeLine []*Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine

    // List of LSAs processed. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_LsaProcessed.
    LsaProcessed []*Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_LsaProcessed
}

func (prefixEventOffline *Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline) GetEntityData() *types.CommonEntityData {
    prefixEventOffline.EntityData.YFilter = prefixEventOffline.YFilter
    prefixEventOffline.EntityData.YangName = "prefix-event-offline"
    prefixEventOffline.EntityData.BundleName = "cisco_ios_xr"
    prefixEventOffline.EntityData.ParentYangName = "prefix-event-offlines"
    prefixEventOffline.EntityData.SegmentPath = "prefix-event-offline" + types.AddKeyToken(prefixEventOffline.EventId, "event-id")
    prefixEventOffline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixEventOffline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixEventOffline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixEventOffline.EntityData.Children = types.NewOrderedMap()
    prefixEventOffline.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prefixEventOffline.IpConvergenceTime})
    prefixEventOffline.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prefixEventOffline.MplsConvergenceTime})
    prefixEventOffline.EntityData.Children.Append("path", types.YChild{"Path", nil})
    for i := range prefixEventOffline.Path {
        prefixEventOffline.EntityData.Children.Append(types.GetSegmentPath(prefixEventOffline.Path[i]), types.YChild{"Path", prefixEventOffline.Path[i]})
    }
    prefixEventOffline.EntityData.Children.Append("trigger-lsa", types.YChild{"TriggerLsa", nil})
    for i := range prefixEventOffline.TriggerLsa {
        prefixEventOffline.EntityData.Children.Append(types.GetSegmentPath(prefixEventOffline.TriggerLsa[i]), types.YChild{"TriggerLsa", prefixEventOffline.TriggerLsa[i]})
    }
    prefixEventOffline.EntityData.Children.Append("time-line", types.YChild{"TimeLine", nil})
    for i := range prefixEventOffline.TimeLine {
        prefixEventOffline.EntityData.Children.Append(types.GetSegmentPath(prefixEventOffline.TimeLine[i]), types.YChild{"TimeLine", prefixEventOffline.TimeLine[i]})
    }
    prefixEventOffline.EntityData.Children.Append("lsa-processed", types.YChild{"LsaProcessed", nil})
    for i := range prefixEventOffline.LsaProcessed {
        prefixEventOffline.EntityData.Children.Append(types.GetSegmentPath(prefixEventOffline.LsaProcessed[i]), types.YChild{"LsaProcessed", prefixEventOffline.LsaProcessed[i]})
    }
    prefixEventOffline.EntityData.Leafs = types.NewOrderedMap()
    prefixEventOffline.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", prefixEventOffline.EventId})
    prefixEventOffline.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", prefixEventOffline.Prefix})
    prefixEventOffline.EntityData.Leafs.Append("prefix-lenth", types.YLeaf{"PrefixLenth", prefixEventOffline.PrefixLenth})
    prefixEventOffline.EntityData.Leafs.Append("spf-run-no", types.YLeaf{"SpfRunNo", prefixEventOffline.SpfRunNo})
    prefixEventOffline.EntityData.Leafs.Append("ipfrr-event-id", types.YLeaf{"IpfrrEventId", prefixEventOffline.IpfrrEventId})
    prefixEventOffline.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prefixEventOffline.ThresholdExceeded})
    prefixEventOffline.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", prefixEventOffline.Priority})
    prefixEventOffline.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", prefixEventOffline.ChangeType})
    prefixEventOffline.EntityData.Leafs.Append("route-type", types.YLeaf{"RouteType", prefixEventOffline.RouteType})
    prefixEventOffline.EntityData.Leafs.Append("route-path-change-type", types.YLeaf{"RoutePathChangeType", prefixEventOffline.RoutePathChangeType})
    prefixEventOffline.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", prefixEventOffline.Cost})
    prefixEventOffline.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", prefixEventOffline.TriggerTime})

    prefixEventOffline.EntityData.YListKeys = []string {"EventId"}

    return &(prefixEventOffline.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "prefix-event-offline"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "prefix-event-offline"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path
// Path information
type Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Path Metric. The type is interface{} with range: 0..4294967295.
    PathMetric interface{}

    // Backup Path Informatoin. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path_LfaPath.
    LfaPath []*Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path_LfaPath
}

func (path *Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "prefix-event-offline"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Children.Append("lfa-path", types.YChild{"LfaPath", nil})
    for i := range path.LfaPath {
        path.EntityData.Children.Append(types.GetSegmentPath(path.LfaPath[i]), types.YChild{"LfaPath", path.LfaPath[i]})
    }
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", path.InterfaceName})
    path.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", path.NeighbourAddress})
    path.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", path.ChangeType})
    path.EntityData.Leafs.Append("path-metric", types.YLeaf{"PathMetric", path.PathMetric})

    path.EntityData.YListKeys = []string {}

    return &(path.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path_LfaPath
// Backup Path Informatoin
type Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path_LfaPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of LFA. The type is RcmdShowIpfrrLfa.
    LfaType interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Path Metric. The type is interface{} with range: 0..4294967295.
    PathMetric interface{}

    // Remote Node ID, in case of Remote LFA. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteNodeId interface{}
}

func (lfaPath *Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path_LfaPath) GetEntityData() *types.CommonEntityData {
    lfaPath.EntityData.YFilter = lfaPath.YFilter
    lfaPath.EntityData.YangName = "lfa-path"
    lfaPath.EntityData.BundleName = "cisco_ios_xr"
    lfaPath.EntityData.ParentYangName = "path"
    lfaPath.EntityData.SegmentPath = "lfa-path"
    lfaPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lfaPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lfaPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lfaPath.EntityData.Children = types.NewOrderedMap()
    lfaPath.EntityData.Leafs = types.NewOrderedMap()
    lfaPath.EntityData.Leafs.Append("lfa-type", types.YLeaf{"LfaType", lfaPath.LfaType})
    lfaPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", lfaPath.InterfaceName})
    lfaPath.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", lfaPath.NeighbourAddress})
    lfaPath.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lfaPath.ChangeType})
    lfaPath.EntityData.Leafs.Append("path-metric", types.YLeaf{"PathMetric", lfaPath.PathMetric})
    lfaPath.EntityData.Leafs.Append("remote-node-id", types.YLeaf{"RemoteNodeId", lfaPath.RemoteNodeId})

    lfaPath.EntityData.YListKeys = []string {}

    return &(lfaPath.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TriggerLsa
// LSA that triggered this event
type Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TriggerLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (triggerLsa *Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TriggerLsa) GetEntityData() *types.CommonEntityData {
    triggerLsa.EntityData.YFilter = triggerLsa.YFilter
    triggerLsa.EntityData.YangName = "trigger-lsa"
    triggerLsa.EntityData.BundleName = "cisco_ios_xr"
    triggerLsa.EntityData.ParentYangName = "prefix-event-offline"
    triggerLsa.EntityData.SegmentPath = "trigger-lsa"
    triggerLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    triggerLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    triggerLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    triggerLsa.EntityData.Children = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", triggerLsa.LsaId})
    triggerLsa.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", triggerLsa.SequenceNumber})
    triggerLsa.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", triggerLsa.LsaType})
    triggerLsa.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", triggerLsa.OriginRouterId})
    triggerLsa.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", triggerLsa.ChangeType})
    triggerLsa.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", triggerLsa.ReceptionTime})

    triggerLsa.EntityData.YListKeys = []string {}

    return &(triggerLsa.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine
// Timeline information
type Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route origin (routing protocol). The type is string.
    RouteOrigin interface{}

    // Entry point of IPv4 RIB. The type is string.
    RiBv4Enter interface{}

    // Exit point from IPv4 RIB to FIBs. The type is string.
    RiBv4Exit interface{}

    // Route Redistribute point from IPv4 RIB to LDP. The type is string.
    RiBv4Redistribute interface{}

    // Entry point of LDP. The type is string.
    LdpEnter interface{}

    // Exit point of LDP to LSD. The type is string.
    LdpExit interface{}

    // Entry point of LSD. The type is string.
    LsdEnter interface{}

    // Exit point of LSD to FIBs. The type is string.
    LsdExit interface{}

    // List of Linecards' completion point for IP routes. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcIp.
    LcIp []*Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcIp

    // List of Linecards' completion point for MPLS labels. The type is slice of
    // Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcMpls.
    LcMpls []*Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcMpls
}

func (timeLine *Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine) GetEntityData() *types.CommonEntityData {
    timeLine.EntityData.YFilter = timeLine.YFilter
    timeLine.EntityData.YangName = "time-line"
    timeLine.EntityData.BundleName = "cisco_ios_xr"
    timeLine.EntityData.ParentYangName = "prefix-event-offline"
    timeLine.EntityData.SegmentPath = "time-line"
    timeLine.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeLine.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeLine.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeLine.EntityData.Children = types.NewOrderedMap()
    timeLine.EntityData.Children.Append("lc-ip", types.YChild{"LcIp", nil})
    for i := range timeLine.LcIp {
        timeLine.EntityData.Children.Append(types.GetSegmentPath(timeLine.LcIp[i]), types.YChild{"LcIp", timeLine.LcIp[i]})
    }
    timeLine.EntityData.Children.Append("lc-mpls", types.YChild{"LcMpls", nil})
    for i := range timeLine.LcMpls {
        timeLine.EntityData.Children.Append(types.GetSegmentPath(timeLine.LcMpls[i]), types.YChild{"LcMpls", timeLine.LcMpls[i]})
    }
    timeLine.EntityData.Leafs = types.NewOrderedMap()
    timeLine.EntityData.Leafs.Append("route-origin", types.YLeaf{"RouteOrigin", timeLine.RouteOrigin})
    timeLine.EntityData.Leafs.Append("ri-bv4-enter", types.YLeaf{"RiBv4Enter", timeLine.RiBv4Enter})
    timeLine.EntityData.Leafs.Append("ri-bv4-exit", types.YLeaf{"RiBv4Exit", timeLine.RiBv4Exit})
    timeLine.EntityData.Leafs.Append("ri-bv4-redistribute", types.YLeaf{"RiBv4Redistribute", timeLine.RiBv4Redistribute})
    timeLine.EntityData.Leafs.Append("ldp-enter", types.YLeaf{"LdpEnter", timeLine.LdpEnter})
    timeLine.EntityData.Leafs.Append("ldp-exit", types.YLeaf{"LdpExit", timeLine.LdpExit})
    timeLine.EntityData.Leafs.Append("lsd-enter", types.YLeaf{"LsdEnter", timeLine.LsdEnter})
    timeLine.EntityData.Leafs.Append("lsd-exit", types.YLeaf{"LsdExit", timeLine.LsdExit})

    timeLine.EntityData.YListKeys = []string {}

    return &(timeLine.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcIp
// List of Linecards' completion point for IP
// routes
type Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB. The type is string.
    FibComplete interface{}
}

func (lcIp *Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcIp) GetEntityData() *types.CommonEntityData {
    lcIp.EntityData.YFilter = lcIp.YFilter
    lcIp.EntityData.YangName = "lc-ip"
    lcIp.EntityData.BundleName = "cisco_ios_xr"
    lcIp.EntityData.ParentYangName = "time-line"
    lcIp.EntityData.SegmentPath = "lc-ip"
    lcIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcIp.EntityData.Children = types.NewOrderedMap()
    lcIp.EntityData.Leafs = types.NewOrderedMap()
    lcIp.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcIp.NodeName})
    lcIp.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcIp.Speed})
    lcIp.EntityData.Leafs.Append("fib-complete", types.YLeaf{"FibComplete", lcIp.FibComplete})

    lcIp.EntityData.YListKeys = []string {}

    return &(lcIp.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcMpls
// List of Linecards' completion point for MPLS
// labels
type Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB. The type is string.
    FibComplete interface{}
}

func (lcMpls *Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcMpls) GetEntityData() *types.CommonEntityData {
    lcMpls.EntityData.YFilter = lcMpls.YFilter
    lcMpls.EntityData.YangName = "lc-mpls"
    lcMpls.EntityData.BundleName = "cisco_ios_xr"
    lcMpls.EntityData.ParentYangName = "time-line"
    lcMpls.EntityData.SegmentPath = "lc-mpls"
    lcMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcMpls.EntityData.Children = types.NewOrderedMap()
    lcMpls.EntityData.Leafs = types.NewOrderedMap()
    lcMpls.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcMpls.NodeName})
    lcMpls.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcMpls.Speed})
    lcMpls.EntityData.Leafs.Append("fib-complete", types.YLeaf{"FibComplete", lcMpls.FibComplete})

    lcMpls.EntityData.YListKeys = []string {}

    return &(lcMpls.EntityData)
}

// Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_LsaProcessed
// List of LSAs processed
type Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_LsaProcessed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (lsaProcessed *Rcmd_Ospf_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_LsaProcessed) GetEntityData() *types.CommonEntityData {
    lsaProcessed.EntityData.YFilter = lsaProcessed.YFilter
    lsaProcessed.EntityData.YangName = "lsa-processed"
    lsaProcessed.EntityData.BundleName = "cisco_ios_xr"
    lsaProcessed.EntityData.ParentYangName = "prefix-event-offline"
    lsaProcessed.EntityData.SegmentPath = "lsa-processed"
    lsaProcessed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsaProcessed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsaProcessed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsaProcessed.EntityData.Children = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", lsaProcessed.LsaId})
    lsaProcessed.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", lsaProcessed.SequenceNumber})
    lsaProcessed.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", lsaProcessed.LsaType})
    lsaProcessed.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", lsaProcessed.OriginRouterId})
    lsaProcessed.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lsaProcessed.ChangeType})
    lsaProcessed.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", lsaProcessed.ReceptionTime})

    lsaProcessed.EntityData.YListKeys = []string {}

    return &(lsaProcessed.EntityData)
}

// Rcmd_Ospf_Instances_Instance_SummaryExternalEventStatistics
// Summary-External prefix monitoring statistics
type Rcmd_Ospf_Instances_Instance_SummaryExternalEventStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total Inter-Area Routes. The type is interface{} with range: 0..4294967295.
    InterAreaRoutes interface{}

    // Total IA Routes Added. The type is interface{} with range: 0..4294967295.
    InterAreaAdded interface{}

    // Total IA Routes Modified. The type is interface{} with range:
    // 0..4294967295.
    InterAreaModified interface{}

    // Total IA Routes Deleted. The type is interface{} with range: 0..4294967295.
    InterAreaDeleted interface{}

    // Total IA Routes Critical. The type is interface{} with range:
    // 0..4294967295.
    InterAreaCritical interface{}

    // Total IA Routes High. The type is interface{} with range: 0..4294967295.
    InterAreaHigh interface{}

    // Total IA Routes Medium. The type is interface{} with range: 0..4294967295.
    InterAreaMedium interface{}

    // Total IA Routes Low. The type is interface{} with range: 0..4294967295.
    InterAreaLow interface{}

    // Total External Routes. The type is interface{} with range: 0..4294967295.
    ExternalRoutes interface{}

    // Total Ext Routes Added. The type is interface{} with range: 0..4294967295.
    ExternalAdded interface{}

    // Total Ext Routes Modified. The type is interface{} with range:
    // 0..4294967295.
    ExternalModified interface{}

    // Total Ext Routes Deleted. The type is interface{} with range:
    // 0..4294967295.
    ExternalDeleted interface{}

    // Total Ext Routes Critical. The type is interface{} with range:
    // 0..4294967295.
    ExternalCritical interface{}

    // Total Ext Routes High. The type is interface{} with range: 0..4294967295.
    ExternalHigh interface{}

    // Total Ext Routes Medium. The type is interface{} with range: 0..4294967295.
    ExternalMedium interface{}

    // Total Ext Routes Low. The type is interface{} with range: 0..4294967295.
    ExternalLow interface{}
}

func (summaryExternalEventStatistics *Rcmd_Ospf_Instances_Instance_SummaryExternalEventStatistics) GetEntityData() *types.CommonEntityData {
    summaryExternalEventStatistics.EntityData.YFilter = summaryExternalEventStatistics.YFilter
    summaryExternalEventStatistics.EntityData.YangName = "summary-external-event-statistics"
    summaryExternalEventStatistics.EntityData.BundleName = "cisco_ios_xr"
    summaryExternalEventStatistics.EntityData.ParentYangName = "instance"
    summaryExternalEventStatistics.EntityData.SegmentPath = "summary-external-event-statistics"
    summaryExternalEventStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryExternalEventStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryExternalEventStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryExternalEventStatistics.EntityData.Children = types.NewOrderedMap()
    summaryExternalEventStatistics.EntityData.Leafs = types.NewOrderedMap()
    summaryExternalEventStatistics.EntityData.Leafs.Append("inter-area-routes", types.YLeaf{"InterAreaRoutes", summaryExternalEventStatistics.InterAreaRoutes})
    summaryExternalEventStatistics.EntityData.Leafs.Append("inter-area-added", types.YLeaf{"InterAreaAdded", summaryExternalEventStatistics.InterAreaAdded})
    summaryExternalEventStatistics.EntityData.Leafs.Append("inter-area-modified", types.YLeaf{"InterAreaModified", summaryExternalEventStatistics.InterAreaModified})
    summaryExternalEventStatistics.EntityData.Leafs.Append("inter-area-deleted", types.YLeaf{"InterAreaDeleted", summaryExternalEventStatistics.InterAreaDeleted})
    summaryExternalEventStatistics.EntityData.Leafs.Append("inter-area-critical", types.YLeaf{"InterAreaCritical", summaryExternalEventStatistics.InterAreaCritical})
    summaryExternalEventStatistics.EntityData.Leafs.Append("inter-area-high", types.YLeaf{"InterAreaHigh", summaryExternalEventStatistics.InterAreaHigh})
    summaryExternalEventStatistics.EntityData.Leafs.Append("inter-area-medium", types.YLeaf{"InterAreaMedium", summaryExternalEventStatistics.InterAreaMedium})
    summaryExternalEventStatistics.EntityData.Leafs.Append("inter-area-low", types.YLeaf{"InterAreaLow", summaryExternalEventStatistics.InterAreaLow})
    summaryExternalEventStatistics.EntityData.Leafs.Append("external-routes", types.YLeaf{"ExternalRoutes", summaryExternalEventStatistics.ExternalRoutes})
    summaryExternalEventStatistics.EntityData.Leafs.Append("external-added", types.YLeaf{"ExternalAdded", summaryExternalEventStatistics.ExternalAdded})
    summaryExternalEventStatistics.EntityData.Leafs.Append("external-modified", types.YLeaf{"ExternalModified", summaryExternalEventStatistics.ExternalModified})
    summaryExternalEventStatistics.EntityData.Leafs.Append("external-deleted", types.YLeaf{"ExternalDeleted", summaryExternalEventStatistics.ExternalDeleted})
    summaryExternalEventStatistics.EntityData.Leafs.Append("external-critical", types.YLeaf{"ExternalCritical", summaryExternalEventStatistics.ExternalCritical})
    summaryExternalEventStatistics.EntityData.Leafs.Append("external-high", types.YLeaf{"ExternalHigh", summaryExternalEventStatistics.ExternalHigh})
    summaryExternalEventStatistics.EntityData.Leafs.Append("external-medium", types.YLeaf{"ExternalMedium", summaryExternalEventStatistics.ExternalMedium})
    summaryExternalEventStatistics.EntityData.Leafs.Append("external-low", types.YLeaf{"ExternalLow", summaryExternalEventStatistics.ExternalLow})

    summaryExternalEventStatistics.EntityData.YListKeys = []string {}

    return &(summaryExternalEventStatistics.EntityData)
}

// Rcmd_Server
// Server Info
type Rcmd_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Server Info.
    Normal Rcmd_Server_Normal

    // Server Info.
    Detail Rcmd_Server_Detail
}

func (server *Rcmd_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "rcmd"
    server.EntityData.SegmentPath = "server"
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = types.NewOrderedMap()
    server.EntityData.Children.Append("normal", types.YChild{"Normal", &server.Normal})
    server.EntityData.Children.Append("detail", types.YChild{"Detail", &server.Detail})
    server.EntityData.Leafs = types.NewOrderedMap()

    server.EntityData.YListKeys = []string {}

    return &(server.EntityData)
}

// Rcmd_Server_Normal
// Server Info
type Rcmd_Server_Normal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured Hostname. The type is string.
    HostName interface{}

    // Server Status. The type is RcmdBagEnableDisable.
    Status interface{}

    // Maximum Events. The type is interface{} with range: 0..4294967295.
    MaxEvents interface{}

    // Event Buffer Size. The type is interface{} with range: 0..4294967295.
    EventBufferSize interface{}

    // Configured Monitor Interval. The type is interface{} with range:
    // 0..4294967295.
    MonitoringInterval interface{}

    // Time for next processing. The type is interface{} with range:
    // 0..4294967295.
    NextInterval interface{}

    // Max Interface events count. The type is interface{} with range:
    // 0..4294967295.
    MaxInterfaceCount interface{}

    // Interface events count. The type is interface{} with range: 0..4294967295.
    InterfaceCount interface{}

    // RP count. The type is interface{} with range: 0..4294967295.
    NodeRpCount interface{}

    // LC count. The type is interface{} with range: 0..4294967295.
    NodeLcCount interface{}

    // Diag Node count. The type is interface{} with range: 0..4294967295.
    DiagNodeCount interface{}

    // Disabled Node count. The type is interface{} with range: 0..4294967295.
    DisabledNodeCount interface{}

    // Disabled Node count. The type is interface{} with range: 0..4294967295.
    InActiveNodeCount interface{}

    // Last Processing Start Time. The type is string.
    LastProcessStartTime interface{}

    // Last Processing Duration. The type is string.
    LastProcessDuration interface{}

    // Process state. The type is RcmdShowPrcsState.
    LastProcessState interface{}

    // Post Processing count. The type is interface{} with range: 0..4294967295.
    ProcessCount interface{}

    // SPF Processing count. The type is interface{} with range: 0..4294967295.
    SpfProcessCount interface{}

    // Reports Archival Path. The type is string.
    ReportsArchivePath interface{}

    // Reports Archival Node (Applicable for local location). The type is string.
    ReportsArchiveNode interface{}

    // Last Archival Status. The type is string.
    LastArchivalStatus interface{}

    // Last Archival Error. The type is string.
    LastArchivalError interface{}

    // Last Archival Status. The type is string.
    LastArchivalErrorTime interface{}

    // Archive Count. The type is interface{} with range: 0..4294967295.
    ArchiveCount interface{}

    // Diagnostics Archival Path. The type is string.
    DiagnosticsArchivePath interface{}

    // Diagnostics Archival Node (Applicable for local location). The type is
    // string.
    DiagnosticsArchiveNode interface{}

    // Protocol level configuration. The type is slice of
    // Rcmd_Server_Normal_ProtocolConfig.
    ProtocolConfig []*Rcmd_Server_Normal_ProtocolConfig

    // Detailed Information. The type is slice of Rcmd_Server_Normal_ServerDetail.
    ServerDetail []*Rcmd_Server_Normal_ServerDetail
}

func (normal *Rcmd_Server_Normal) GetEntityData() *types.CommonEntityData {
    normal.EntityData.YFilter = normal.YFilter
    normal.EntityData.YangName = "normal"
    normal.EntityData.BundleName = "cisco_ios_xr"
    normal.EntityData.ParentYangName = "server"
    normal.EntityData.SegmentPath = "normal"
    normal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    normal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    normal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    normal.EntityData.Children = types.NewOrderedMap()
    normal.EntityData.Children.Append("protocol-config", types.YChild{"ProtocolConfig", nil})
    for i := range normal.ProtocolConfig {
        normal.EntityData.Children.Append(types.GetSegmentPath(normal.ProtocolConfig[i]), types.YChild{"ProtocolConfig", normal.ProtocolConfig[i]})
    }
    normal.EntityData.Children.Append("server-detail", types.YChild{"ServerDetail", nil})
    for i := range normal.ServerDetail {
        normal.EntityData.Children.Append(types.GetSegmentPath(normal.ServerDetail[i]), types.YChild{"ServerDetail", normal.ServerDetail[i]})
    }
    normal.EntityData.Leafs = types.NewOrderedMap()
    normal.EntityData.Leafs.Append("host-name", types.YLeaf{"HostName", normal.HostName})
    normal.EntityData.Leafs.Append("status", types.YLeaf{"Status", normal.Status})
    normal.EntityData.Leafs.Append("max-events", types.YLeaf{"MaxEvents", normal.MaxEvents})
    normal.EntityData.Leafs.Append("event-buffer-size", types.YLeaf{"EventBufferSize", normal.EventBufferSize})
    normal.EntityData.Leafs.Append("monitoring-interval", types.YLeaf{"MonitoringInterval", normal.MonitoringInterval})
    normal.EntityData.Leafs.Append("next-interval", types.YLeaf{"NextInterval", normal.NextInterval})
    normal.EntityData.Leafs.Append("max-interface-count", types.YLeaf{"MaxInterfaceCount", normal.MaxInterfaceCount})
    normal.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", normal.InterfaceCount})
    normal.EntityData.Leafs.Append("node-rp-count", types.YLeaf{"NodeRpCount", normal.NodeRpCount})
    normal.EntityData.Leafs.Append("node-lc-count", types.YLeaf{"NodeLcCount", normal.NodeLcCount})
    normal.EntityData.Leafs.Append("diag-node-count", types.YLeaf{"DiagNodeCount", normal.DiagNodeCount})
    normal.EntityData.Leafs.Append("disabled-node-count", types.YLeaf{"DisabledNodeCount", normal.DisabledNodeCount})
    normal.EntityData.Leafs.Append("in-active-node-count", types.YLeaf{"InActiveNodeCount", normal.InActiveNodeCount})
    normal.EntityData.Leafs.Append("last-process-start-time", types.YLeaf{"LastProcessStartTime", normal.LastProcessStartTime})
    normal.EntityData.Leafs.Append("last-process-duration", types.YLeaf{"LastProcessDuration", normal.LastProcessDuration})
    normal.EntityData.Leafs.Append("last-process-state", types.YLeaf{"LastProcessState", normal.LastProcessState})
    normal.EntityData.Leafs.Append("process-count", types.YLeaf{"ProcessCount", normal.ProcessCount})
    normal.EntityData.Leafs.Append("spf-process-count", types.YLeaf{"SpfProcessCount", normal.SpfProcessCount})
    normal.EntityData.Leafs.Append("reports-archive-path", types.YLeaf{"ReportsArchivePath", normal.ReportsArchivePath})
    normal.EntityData.Leafs.Append("reports-archive-node", types.YLeaf{"ReportsArchiveNode", normal.ReportsArchiveNode})
    normal.EntityData.Leafs.Append("last-archival-status", types.YLeaf{"LastArchivalStatus", normal.LastArchivalStatus})
    normal.EntityData.Leafs.Append("last-archival-error", types.YLeaf{"LastArchivalError", normal.LastArchivalError})
    normal.EntityData.Leafs.Append("last-archival-error-time", types.YLeaf{"LastArchivalErrorTime", normal.LastArchivalErrorTime})
    normal.EntityData.Leafs.Append("archive-count", types.YLeaf{"ArchiveCount", normal.ArchiveCount})
    normal.EntityData.Leafs.Append("diagnostics-archive-path", types.YLeaf{"DiagnosticsArchivePath", normal.DiagnosticsArchivePath})
    normal.EntityData.Leafs.Append("diagnostics-archive-node", types.YLeaf{"DiagnosticsArchiveNode", normal.DiagnosticsArchiveNode})

    normal.EntityData.YListKeys = []string {}

    return &(normal.EntityData)
}

// Rcmd_Server_Normal_ProtocolConfig
// Protocol level configuration
type Rcmd_Server_Normal_ProtocolConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol Name. The type is string.
    ProtocolName interface{}

    // Priority level configuration. The type is slice of
    // Rcmd_Server_Normal_ProtocolConfig_Priority.
    Priority []*Rcmd_Server_Normal_ProtocolConfig_Priority
}

func (protocolConfig *Rcmd_Server_Normal_ProtocolConfig) GetEntityData() *types.CommonEntityData {
    protocolConfig.EntityData.YFilter = protocolConfig.YFilter
    protocolConfig.EntityData.YangName = "protocol-config"
    protocolConfig.EntityData.BundleName = "cisco_ios_xr"
    protocolConfig.EntityData.ParentYangName = "normal"
    protocolConfig.EntityData.SegmentPath = "protocol-config"
    protocolConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocolConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocolConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocolConfig.EntityData.Children = types.NewOrderedMap()
    protocolConfig.EntityData.Children.Append("priority", types.YChild{"Priority", nil})
    for i := range protocolConfig.Priority {
        protocolConfig.EntityData.Children.Append(types.GetSegmentPath(protocolConfig.Priority[i]), types.YChild{"Priority", protocolConfig.Priority[i]})
    }
    protocolConfig.EntityData.Leafs = types.NewOrderedMap()
    protocolConfig.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", protocolConfig.ProtocolName})

    protocolConfig.EntityData.YListKeys = []string {}

    return &(protocolConfig.EntityData)
}

// Rcmd_Server_Normal_ProtocolConfig_Priority
// Priority level configuration
type Rcmd_Server_Normal_ProtocolConfig_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Priority Level. The type is RcmdPriorityLevel.
    PriorityName interface{}

    // threshold value. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Enable/Disable cfg. The type is RcmdBoolYesNo.
    Disable interface{}
}

func (priority *Rcmd_Server_Normal_ProtocolConfig_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "protocol-config"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Leafs = types.NewOrderedMap()
    priority.EntityData.Leafs.Append("priority-name", types.YLeaf{"PriorityName", priority.PriorityName})
    priority.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", priority.Threshold})
    priority.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", priority.Disable})

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// Rcmd_Server_Normal_ServerDetail
// Detailed Information
type Rcmd_Server_Normal_ServerDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Overload suspend. The type is interface{} with range: 0..4294967295.
    OverloadSuspend interface{}

    // Memory Suspend. The type is interface{} with range: 0..4294967295.
    MemorySuspend interface{}

    // Trace Information. The type is slice of
    // Rcmd_Server_Normal_ServerDetail_TraceInformation.
    TraceInformation []*Rcmd_Server_Normal_ServerDetail_TraceInformation
}

func (serverDetail *Rcmd_Server_Normal_ServerDetail) GetEntityData() *types.CommonEntityData {
    serverDetail.EntityData.YFilter = serverDetail.YFilter
    serverDetail.EntityData.YangName = "server-detail"
    serverDetail.EntityData.BundleName = "cisco_ios_xr"
    serverDetail.EntityData.ParentYangName = "normal"
    serverDetail.EntityData.SegmentPath = "server-detail"
    serverDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverDetail.EntityData.Children = types.NewOrderedMap()
    serverDetail.EntityData.Children.Append("trace-information", types.YChild{"TraceInformation", nil})
    for i := range serverDetail.TraceInformation {
        serverDetail.EntityData.Children.Append(types.GetSegmentPath(serverDetail.TraceInformation[i]), types.YChild{"TraceInformation", serverDetail.TraceInformation[i]})
    }
    serverDetail.EntityData.Leafs = types.NewOrderedMap()
    serverDetail.EntityData.Leafs.Append("overload-suspend", types.YLeaf{"OverloadSuspend", serverDetail.OverloadSuspend})
    serverDetail.EntityData.Leafs.Append("memory-suspend", types.YLeaf{"MemorySuspend", serverDetail.MemorySuspend})

    serverDetail.EntityData.YListKeys = []string {}

    return &(serverDetail.EntityData)
}

// Rcmd_Server_Normal_ServerDetail_TraceInformation
// Trace Information
type Rcmd_Server_Normal_ServerDetail_TraceInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured Hostname. The type is string.
    TraceName interface{}

    // Server Total Status. The type is interface{} with range: 0..4294967295.
    TotalStats interface{}

    // Server Last Run Status. The type is interface{} with range: 0..4294967295.
    LastRunStats interface{}

    // Server Error Status. The type is interface{} with range: 0..4294967295.
    ErrorStats interface{}
}

func (traceInformation *Rcmd_Server_Normal_ServerDetail_TraceInformation) GetEntityData() *types.CommonEntityData {
    traceInformation.EntityData.YFilter = traceInformation.YFilter
    traceInformation.EntityData.YangName = "trace-information"
    traceInformation.EntityData.BundleName = "cisco_ios_xr"
    traceInformation.EntityData.ParentYangName = "server-detail"
    traceInformation.EntityData.SegmentPath = "trace-information"
    traceInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceInformation.EntityData.Children = types.NewOrderedMap()
    traceInformation.EntityData.Leafs = types.NewOrderedMap()
    traceInformation.EntityData.Leafs.Append("trace-name", types.YLeaf{"TraceName", traceInformation.TraceName})
    traceInformation.EntityData.Leafs.Append("total-stats", types.YLeaf{"TotalStats", traceInformation.TotalStats})
    traceInformation.EntityData.Leafs.Append("last-run-stats", types.YLeaf{"LastRunStats", traceInformation.LastRunStats})
    traceInformation.EntityData.Leafs.Append("error-stats", types.YLeaf{"ErrorStats", traceInformation.ErrorStats})

    traceInformation.EntityData.YListKeys = []string {}

    return &(traceInformation.EntityData)
}

// Rcmd_Server_Detail
// Server Info
type Rcmd_Server_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured Hostname. The type is string.
    HostName interface{}

    // Server Status. The type is RcmdBagEnableDisable.
    Status interface{}

    // Maximum Events. The type is interface{} with range: 0..4294967295.
    MaxEvents interface{}

    // Event Buffer Size. The type is interface{} with range: 0..4294967295.
    EventBufferSize interface{}

    // Configured Monitor Interval. The type is interface{} with range:
    // 0..4294967295.
    MonitoringInterval interface{}

    // Time for next processing. The type is interface{} with range:
    // 0..4294967295.
    NextInterval interface{}

    // Max Interface events count. The type is interface{} with range:
    // 0..4294967295.
    MaxInterfaceCount interface{}

    // Interface events count. The type is interface{} with range: 0..4294967295.
    InterfaceCount interface{}

    // RP count. The type is interface{} with range: 0..4294967295.
    NodeRpCount interface{}

    // LC count. The type is interface{} with range: 0..4294967295.
    NodeLcCount interface{}

    // Diag Node count. The type is interface{} with range: 0..4294967295.
    DiagNodeCount interface{}

    // Disabled Node count. The type is interface{} with range: 0..4294967295.
    DisabledNodeCount interface{}

    // Disabled Node count. The type is interface{} with range: 0..4294967295.
    InActiveNodeCount interface{}

    // Last Processing Start Time. The type is string.
    LastProcessStartTime interface{}

    // Last Processing Duration. The type is string.
    LastProcessDuration interface{}

    // Process state. The type is RcmdShowPrcsState.
    LastProcessState interface{}

    // Post Processing count. The type is interface{} with range: 0..4294967295.
    ProcessCount interface{}

    // SPF Processing count. The type is interface{} with range: 0..4294967295.
    SpfProcessCount interface{}

    // Reports Archival Path. The type is string.
    ReportsArchivePath interface{}

    // Reports Archival Node (Applicable for local location). The type is string.
    ReportsArchiveNode interface{}

    // Last Archival Status. The type is string.
    LastArchivalStatus interface{}

    // Last Archival Error. The type is string.
    LastArchivalError interface{}

    // Last Archival Status. The type is string.
    LastArchivalErrorTime interface{}

    // Archive Count. The type is interface{} with range: 0..4294967295.
    ArchiveCount interface{}

    // Diagnostics Archival Path. The type is string.
    DiagnosticsArchivePath interface{}

    // Diagnostics Archival Node (Applicable for local location). The type is
    // string.
    DiagnosticsArchiveNode interface{}

    // Protocol level configuration. The type is slice of
    // Rcmd_Server_Detail_ProtocolConfig.
    ProtocolConfig []*Rcmd_Server_Detail_ProtocolConfig

    // Detailed Information. The type is slice of Rcmd_Server_Detail_ServerDetail.
    ServerDetail []*Rcmd_Server_Detail_ServerDetail
}

func (detail *Rcmd_Server_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "server"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("protocol-config", types.YChild{"ProtocolConfig", nil})
    for i := range detail.ProtocolConfig {
        detail.EntityData.Children.Append(types.GetSegmentPath(detail.ProtocolConfig[i]), types.YChild{"ProtocolConfig", detail.ProtocolConfig[i]})
    }
    detail.EntityData.Children.Append("server-detail", types.YChild{"ServerDetail", nil})
    for i := range detail.ServerDetail {
        detail.EntityData.Children.Append(types.GetSegmentPath(detail.ServerDetail[i]), types.YChild{"ServerDetail", detail.ServerDetail[i]})
    }
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("host-name", types.YLeaf{"HostName", detail.HostName})
    detail.EntityData.Leafs.Append("status", types.YLeaf{"Status", detail.Status})
    detail.EntityData.Leafs.Append("max-events", types.YLeaf{"MaxEvents", detail.MaxEvents})
    detail.EntityData.Leafs.Append("event-buffer-size", types.YLeaf{"EventBufferSize", detail.EventBufferSize})
    detail.EntityData.Leafs.Append("monitoring-interval", types.YLeaf{"MonitoringInterval", detail.MonitoringInterval})
    detail.EntityData.Leafs.Append("next-interval", types.YLeaf{"NextInterval", detail.NextInterval})
    detail.EntityData.Leafs.Append("max-interface-count", types.YLeaf{"MaxInterfaceCount", detail.MaxInterfaceCount})
    detail.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", detail.InterfaceCount})
    detail.EntityData.Leafs.Append("node-rp-count", types.YLeaf{"NodeRpCount", detail.NodeRpCount})
    detail.EntityData.Leafs.Append("node-lc-count", types.YLeaf{"NodeLcCount", detail.NodeLcCount})
    detail.EntityData.Leafs.Append("diag-node-count", types.YLeaf{"DiagNodeCount", detail.DiagNodeCount})
    detail.EntityData.Leafs.Append("disabled-node-count", types.YLeaf{"DisabledNodeCount", detail.DisabledNodeCount})
    detail.EntityData.Leafs.Append("in-active-node-count", types.YLeaf{"InActiveNodeCount", detail.InActiveNodeCount})
    detail.EntityData.Leafs.Append("last-process-start-time", types.YLeaf{"LastProcessStartTime", detail.LastProcessStartTime})
    detail.EntityData.Leafs.Append("last-process-duration", types.YLeaf{"LastProcessDuration", detail.LastProcessDuration})
    detail.EntityData.Leafs.Append("last-process-state", types.YLeaf{"LastProcessState", detail.LastProcessState})
    detail.EntityData.Leafs.Append("process-count", types.YLeaf{"ProcessCount", detail.ProcessCount})
    detail.EntityData.Leafs.Append("spf-process-count", types.YLeaf{"SpfProcessCount", detail.SpfProcessCount})
    detail.EntityData.Leafs.Append("reports-archive-path", types.YLeaf{"ReportsArchivePath", detail.ReportsArchivePath})
    detail.EntityData.Leafs.Append("reports-archive-node", types.YLeaf{"ReportsArchiveNode", detail.ReportsArchiveNode})
    detail.EntityData.Leafs.Append("last-archival-status", types.YLeaf{"LastArchivalStatus", detail.LastArchivalStatus})
    detail.EntityData.Leafs.Append("last-archival-error", types.YLeaf{"LastArchivalError", detail.LastArchivalError})
    detail.EntityData.Leafs.Append("last-archival-error-time", types.YLeaf{"LastArchivalErrorTime", detail.LastArchivalErrorTime})
    detail.EntityData.Leafs.Append("archive-count", types.YLeaf{"ArchiveCount", detail.ArchiveCount})
    detail.EntityData.Leafs.Append("diagnostics-archive-path", types.YLeaf{"DiagnosticsArchivePath", detail.DiagnosticsArchivePath})
    detail.EntityData.Leafs.Append("diagnostics-archive-node", types.YLeaf{"DiagnosticsArchiveNode", detail.DiagnosticsArchiveNode})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Rcmd_Server_Detail_ProtocolConfig
// Protocol level configuration
type Rcmd_Server_Detail_ProtocolConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol Name. The type is string.
    ProtocolName interface{}

    // Priority level configuration. The type is slice of
    // Rcmd_Server_Detail_ProtocolConfig_Priority.
    Priority []*Rcmd_Server_Detail_ProtocolConfig_Priority
}

func (protocolConfig *Rcmd_Server_Detail_ProtocolConfig) GetEntityData() *types.CommonEntityData {
    protocolConfig.EntityData.YFilter = protocolConfig.YFilter
    protocolConfig.EntityData.YangName = "protocol-config"
    protocolConfig.EntityData.BundleName = "cisco_ios_xr"
    protocolConfig.EntityData.ParentYangName = "detail"
    protocolConfig.EntityData.SegmentPath = "protocol-config"
    protocolConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocolConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocolConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocolConfig.EntityData.Children = types.NewOrderedMap()
    protocolConfig.EntityData.Children.Append("priority", types.YChild{"Priority", nil})
    for i := range protocolConfig.Priority {
        protocolConfig.EntityData.Children.Append(types.GetSegmentPath(protocolConfig.Priority[i]), types.YChild{"Priority", protocolConfig.Priority[i]})
    }
    protocolConfig.EntityData.Leafs = types.NewOrderedMap()
    protocolConfig.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", protocolConfig.ProtocolName})

    protocolConfig.EntityData.YListKeys = []string {}

    return &(protocolConfig.EntityData)
}

// Rcmd_Server_Detail_ProtocolConfig_Priority
// Priority level configuration
type Rcmd_Server_Detail_ProtocolConfig_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Priority Level. The type is RcmdPriorityLevel.
    PriorityName interface{}

    // threshold value. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Enable/Disable cfg. The type is RcmdBoolYesNo.
    Disable interface{}
}

func (priority *Rcmd_Server_Detail_ProtocolConfig_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "protocol-config"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Leafs = types.NewOrderedMap()
    priority.EntityData.Leafs.Append("priority-name", types.YLeaf{"PriorityName", priority.PriorityName})
    priority.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", priority.Threshold})
    priority.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", priority.Disable})

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// Rcmd_Server_Detail_ServerDetail
// Detailed Information
type Rcmd_Server_Detail_ServerDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Overload suspend. The type is interface{} with range: 0..4294967295.
    OverloadSuspend interface{}

    // Memory Suspend. The type is interface{} with range: 0..4294967295.
    MemorySuspend interface{}

    // Trace Information. The type is slice of
    // Rcmd_Server_Detail_ServerDetail_TraceInformation.
    TraceInformation []*Rcmd_Server_Detail_ServerDetail_TraceInformation
}

func (serverDetail *Rcmd_Server_Detail_ServerDetail) GetEntityData() *types.CommonEntityData {
    serverDetail.EntityData.YFilter = serverDetail.YFilter
    serverDetail.EntityData.YangName = "server-detail"
    serverDetail.EntityData.BundleName = "cisco_ios_xr"
    serverDetail.EntityData.ParentYangName = "detail"
    serverDetail.EntityData.SegmentPath = "server-detail"
    serverDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverDetail.EntityData.Children = types.NewOrderedMap()
    serverDetail.EntityData.Children.Append("trace-information", types.YChild{"TraceInformation", nil})
    for i := range serverDetail.TraceInformation {
        serverDetail.EntityData.Children.Append(types.GetSegmentPath(serverDetail.TraceInformation[i]), types.YChild{"TraceInformation", serverDetail.TraceInformation[i]})
    }
    serverDetail.EntityData.Leafs = types.NewOrderedMap()
    serverDetail.EntityData.Leafs.Append("overload-suspend", types.YLeaf{"OverloadSuspend", serverDetail.OverloadSuspend})
    serverDetail.EntityData.Leafs.Append("memory-suspend", types.YLeaf{"MemorySuspend", serverDetail.MemorySuspend})

    serverDetail.EntityData.YListKeys = []string {}

    return &(serverDetail.EntityData)
}

// Rcmd_Server_Detail_ServerDetail_TraceInformation
// Trace Information
type Rcmd_Server_Detail_ServerDetail_TraceInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured Hostname. The type is string.
    TraceName interface{}

    // Server Total Status. The type is interface{} with range: 0..4294967295.
    TotalStats interface{}

    // Server Last Run Status. The type is interface{} with range: 0..4294967295.
    LastRunStats interface{}

    // Server Error Status. The type is interface{} with range: 0..4294967295.
    ErrorStats interface{}
}

func (traceInformation *Rcmd_Server_Detail_ServerDetail_TraceInformation) GetEntityData() *types.CommonEntityData {
    traceInformation.EntityData.YFilter = traceInformation.YFilter
    traceInformation.EntityData.YangName = "trace-information"
    traceInformation.EntityData.BundleName = "cisco_ios_xr"
    traceInformation.EntityData.ParentYangName = "server-detail"
    traceInformation.EntityData.SegmentPath = "trace-information"
    traceInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceInformation.EntityData.Children = types.NewOrderedMap()
    traceInformation.EntityData.Leafs = types.NewOrderedMap()
    traceInformation.EntityData.Leafs.Append("trace-name", types.YLeaf{"TraceName", traceInformation.TraceName})
    traceInformation.EntityData.Leafs.Append("total-stats", types.YLeaf{"TotalStats", traceInformation.TotalStats})
    traceInformation.EntityData.Leafs.Append("last-run-stats", types.YLeaf{"LastRunStats", traceInformation.LastRunStats})
    traceInformation.EntityData.Leafs.Append("error-stats", types.YLeaf{"ErrorStats", traceInformation.ErrorStats})

    traceInformation.EntityData.YListKeys = []string {}

    return &(traceInformation.EntityData)
}

// Rcmd_Node
// Node Info
type Rcmd_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Info. The type is slice of Rcmd_Node_NodeInformation.
    NodeInformation []*Rcmd_Node_NodeInformation
}

func (node *Rcmd_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "rcmd"
    node.EntityData.SegmentPath = "node"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("node-information", types.YChild{"NodeInformation", nil})
    for i := range node.NodeInformation {
        node.EntityData.Children.Append(types.GetSegmentPath(node.NodeInformation[i]), types.YChild{"NodeInformation", node.NodeInformation[i]})
    }
    node.EntityData.Leafs = types.NewOrderedMap()

    node.EntityData.YListKeys = []string {}

    return &(node.EntityData)
}

// Rcmd_Node_NodeInformation
// Node Info
type Rcmd_Node_NodeInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Id. The type is interface{} with range: 0..4294967295.
    NodeId interface{}

    // Node Name. The type is string.
    NodeName interface{}

    // Rack Id. The type is interface{} with range: 0..4294967295.
    RackId interface{}

    // Last Updated Time. The type is string.
    LastUpdateTime interface{}

    // Forward Referenced. The type is RcmdBoolYesNo.
    FwdReferenced interface{}

    // Node Type. The type is RcmdShowNode.
    NodeType interface{}

    // Software State. The type is interface{} with range: 0..4294967295.
    SoftwareState interface{}

    // Card State. The type is interface{} with range: 0..4294967295.
    CardState interface{}

    // Node State. The type is RcmdBoolYesNo.
    NodeState interface{}

    // Status. The type is RcmdBagEnblDsbl.
    Status interface{}

    // Diag Mode. The type is interface{} with range: 0..4294967295.
    DiagMode interface{}

    // Redundancy State. The type is interface{} with range: 0..4294967295.
    RedundancyState interface{}
}

func (nodeInformation *Rcmd_Node_NodeInformation) GetEntityData() *types.CommonEntityData {
    nodeInformation.EntityData.YFilter = nodeInformation.YFilter
    nodeInformation.EntityData.YangName = "node-information"
    nodeInformation.EntityData.BundleName = "cisco_ios_xr"
    nodeInformation.EntityData.ParentYangName = "node"
    nodeInformation.EntityData.SegmentPath = "node-information"
    nodeInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInformation.EntityData.Children = types.NewOrderedMap()
    nodeInformation.EntityData.Leafs = types.NewOrderedMap()
    nodeInformation.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", nodeInformation.NodeId})
    nodeInformation.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", nodeInformation.NodeName})
    nodeInformation.EntityData.Leafs.Append("rack-id", types.YLeaf{"RackId", nodeInformation.RackId})
    nodeInformation.EntityData.Leafs.Append("last-update-time", types.YLeaf{"LastUpdateTime", nodeInformation.LastUpdateTime})
    nodeInformation.EntityData.Leafs.Append("fwd-referenced", types.YLeaf{"FwdReferenced", nodeInformation.FwdReferenced})
    nodeInformation.EntityData.Leafs.Append("node-type", types.YLeaf{"NodeType", nodeInformation.NodeType})
    nodeInformation.EntityData.Leafs.Append("software-state", types.YLeaf{"SoftwareState", nodeInformation.SoftwareState})
    nodeInformation.EntityData.Leafs.Append("card-state", types.YLeaf{"CardState", nodeInformation.CardState})
    nodeInformation.EntityData.Leafs.Append("node-state", types.YLeaf{"NodeState", nodeInformation.NodeState})
    nodeInformation.EntityData.Leafs.Append("status", types.YLeaf{"Status", nodeInformation.Status})
    nodeInformation.EntityData.Leafs.Append("diag-mode", types.YLeaf{"DiagMode", nodeInformation.DiagMode})
    nodeInformation.EntityData.Leafs.Append("redundancy-state", types.YLeaf{"RedundancyState", nodeInformation.RedundancyState})

    nodeInformation.EntityData.YListKeys = []string {}

    return &(nodeInformation.EntityData)
}

// Rcmd_Isis
// Operational data for ISIS
type Rcmd_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data.
    Instances Rcmd_Isis_Instances
}

func (isis *Rcmd_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "rcmd"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = types.NewOrderedMap()
    isis.EntityData.Children.Append("instances", types.YChild{"Instances", &isis.Instances})
    isis.EntityData.Leafs = types.NewOrderedMap()

    isis.EntityData.YListKeys = []string {}

    return &(isis.EntityData)
}

// Rcmd_Isis_Instances
// Operational data
type Rcmd_Isis_Instances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a particular instance. The type is slice of
    // Rcmd_Isis_Instances_Instance.
    Instance []*Rcmd_Isis_Instances_Instance
}

func (instances *Rcmd_Isis_Instances) GetEntityData() *types.CommonEntityData {
    instances.EntityData.YFilter = instances.YFilter
    instances.EntityData.YangName = "instances"
    instances.EntityData.BundleName = "cisco_ios_xr"
    instances.EntityData.ParentYangName = "isis"
    instances.EntityData.SegmentPath = "instances"
    instances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instances.EntityData.Children = types.NewOrderedMap()
    instances.EntityData.Children.Append("instance", types.YChild{"Instance", nil})
    for i := range instances.Instance {
        instances.EntityData.Children.Append(types.GetSegmentPath(instances.Instance[i]), types.YChild{"Instance", instances.Instance[i]})
    }
    instances.EntityData.Leafs = types.NewOrderedMap()

    instances.EntityData.YListKeys = []string {}

    return &(instances.EntityData)
}

// Rcmd_Isis_Instances_Instance
// Operational data for a particular instance
type Rcmd_Isis_Instances_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Operational data for a particular instance. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // ISIS IP-FRR events summary data.
    IpfrrEventSummaries Rcmd_Isis_Instances_Instance_IpfrrEventSummaries

    // ISIS Prefix events statistics data.
    PrefixEventStatistics Rcmd_Isis_Instances_Instance_PrefixEventStatistics

    // ISIS SPF run summary data.
    SpfRunSummaries Rcmd_Isis_Instances_Instance_SpfRunSummaries

    // ISIS IP-FRR Event offline data.
    IpfrrEventOfflines Rcmd_Isis_Instances_Instance_IpfrrEventOfflines

    // ISIS SPF run offline data.
    SpfRunOfflines Rcmd_Isis_Instances_Instance_SpfRunOfflines

    // ISIS Prefix events summary data.
    PrefixEventSummaries Rcmd_Isis_Instances_Instance_PrefixEventSummaries

    // ISIS Prefix events offline data.
    PrefixEventOfflines Rcmd_Isis_Instances_Instance_PrefixEventOfflines

    // Regenerated LSP data.
    LspRegenerateds Rcmd_Isis_Instances_Instance_LspRegenerateds
}

func (instance *Rcmd_Isis_Instances_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "instances"
    instance.EntityData.SegmentPath = "instance" + types.AddKeyToken(instance.InstanceName, "instance-name")
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Children.Append("ipfrr-event-summaries", types.YChild{"IpfrrEventSummaries", &instance.IpfrrEventSummaries})
    instance.EntityData.Children.Append("prefix-event-statistics", types.YChild{"PrefixEventStatistics", &instance.PrefixEventStatistics})
    instance.EntityData.Children.Append("spf-run-summaries", types.YChild{"SpfRunSummaries", &instance.SpfRunSummaries})
    instance.EntityData.Children.Append("ipfrr-event-offlines", types.YChild{"IpfrrEventOfflines", &instance.IpfrrEventOfflines})
    instance.EntityData.Children.Append("spf-run-offlines", types.YChild{"SpfRunOfflines", &instance.SpfRunOfflines})
    instance.EntityData.Children.Append("prefix-event-summaries", types.YChild{"PrefixEventSummaries", &instance.PrefixEventSummaries})
    instance.EntityData.Children.Append("prefix-event-offlines", types.YChild{"PrefixEventOfflines", &instance.PrefixEventOfflines})
    instance.EntityData.Children.Append("lsp-regenerateds", types.YChild{"LspRegenerateds", &instance.LspRegenerateds})
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", instance.InstanceName})

    instance.EntityData.YListKeys = []string {"InstanceName"}

    return &(instance.EntityData)
}

// Rcmd_Isis_Instances_Instance_IpfrrEventSummaries
// ISIS IP-FRR events summary data
type Rcmd_Isis_Instances_Instance_IpfrrEventSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP-FRR Event data. The type is slice of
    // Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary.
    IpfrrEventSummary []*Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary
}

func (ipfrrEventSummaries *Rcmd_Isis_Instances_Instance_IpfrrEventSummaries) GetEntityData() *types.CommonEntityData {
    ipfrrEventSummaries.EntityData.YFilter = ipfrrEventSummaries.YFilter
    ipfrrEventSummaries.EntityData.YangName = "ipfrr-event-summaries"
    ipfrrEventSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipfrrEventSummaries.EntityData.ParentYangName = "instance"
    ipfrrEventSummaries.EntityData.SegmentPath = "ipfrr-event-summaries"
    ipfrrEventSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipfrrEventSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipfrrEventSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipfrrEventSummaries.EntityData.Children = types.NewOrderedMap()
    ipfrrEventSummaries.EntityData.Children.Append("ipfrr-event-summary", types.YChild{"IpfrrEventSummary", nil})
    for i := range ipfrrEventSummaries.IpfrrEventSummary {
        ipfrrEventSummaries.EntityData.Children.Append(types.GetSegmentPath(ipfrrEventSummaries.IpfrrEventSummary[i]), types.YChild{"IpfrrEventSummary", ipfrrEventSummaries.IpfrrEventSummary[i]})
    }
    ipfrrEventSummaries.EntityData.Leafs = types.NewOrderedMap()

    ipfrrEventSummaries.EntityData.YListKeys = []string {}

    return &(ipfrrEventSummaries.EntityData)
}

// Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary
// IP-FRR Event data
type Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific IP-FRR Event. The type is interface{}
    // with range: 1..4294967295.
    EventId interface{}

    // IP-Frr Event ID. The type is interface{} with range: 0..4294967295.
    EventIdXr interface{}

    // Trigger time  (eg: Apr 24 13:16:04.961). The type is string.
    TriggerTime interface{}

    // IP-Frr Triggered reference SPF Run Number. The type is interface{} with
    // range: 0..4294967295.
    TriggerSpfRun interface{}

    // Waiting Time (in milliseconds). The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    WaitTime interface{}

    // Start Time offset from trigger time (in milliseconds). The type is string.
    // Units are millisecond.
    StartTimeOffset interface{}

    // Duration for the calculation (in milliseconds). The type is string. Units
    // are millisecond.
    Duration interface{}

    // IP-Frr Completed reference SPF Run Number. The type is interface{} with
    // range: 0..4294967295.
    CompletedSpfRun interface{}

    // Cumulative Number of Routes for all priorities. The type is interface{}
    // with range: 0..4294967295.
    TotalRoutes interface{}

    // Cumulative Number of Fully Protected Routes. The type is interface{} with
    // range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Cumulative Number of Partially Protected Routes. The type is interface{}
    // with range: 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage for all priorities. The type is string. Units are
    // percentage.
    Coverage interface{}

    // IP-Frr Statistics categorized by priority. The type is slice of
    // Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_IpfrrStatistic.
    IpfrrStatistic []*Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_IpfrrStatistic

    // Remote Node Information. The type is slice of
    // Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode.
    RemoteNode []*Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode
}

func (ipfrrEventSummary *Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary) GetEntityData() *types.CommonEntityData {
    ipfrrEventSummary.EntityData.YFilter = ipfrrEventSummary.YFilter
    ipfrrEventSummary.EntityData.YangName = "ipfrr-event-summary"
    ipfrrEventSummary.EntityData.BundleName = "cisco_ios_xr"
    ipfrrEventSummary.EntityData.ParentYangName = "ipfrr-event-summaries"
    ipfrrEventSummary.EntityData.SegmentPath = "ipfrr-event-summary" + types.AddKeyToken(ipfrrEventSummary.EventId, "event-id")
    ipfrrEventSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipfrrEventSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipfrrEventSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipfrrEventSummary.EntityData.Children = types.NewOrderedMap()
    ipfrrEventSummary.EntityData.Children.Append("ipfrr-statistic", types.YChild{"IpfrrStatistic", nil})
    for i := range ipfrrEventSummary.IpfrrStatistic {
        ipfrrEventSummary.EntityData.Children.Append(types.GetSegmentPath(ipfrrEventSummary.IpfrrStatistic[i]), types.YChild{"IpfrrStatistic", ipfrrEventSummary.IpfrrStatistic[i]})
    }
    ipfrrEventSummary.EntityData.Children.Append("remote-node", types.YChild{"RemoteNode", nil})
    for i := range ipfrrEventSummary.RemoteNode {
        ipfrrEventSummary.EntityData.Children.Append(types.GetSegmentPath(ipfrrEventSummary.RemoteNode[i]), types.YChild{"RemoteNode", ipfrrEventSummary.RemoteNode[i]})
    }
    ipfrrEventSummary.EntityData.Leafs = types.NewOrderedMap()
    ipfrrEventSummary.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", ipfrrEventSummary.EventId})
    ipfrrEventSummary.EntityData.Leafs.Append("event-id-xr", types.YLeaf{"EventIdXr", ipfrrEventSummary.EventIdXr})
    ipfrrEventSummary.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", ipfrrEventSummary.TriggerTime})
    ipfrrEventSummary.EntityData.Leafs.Append("trigger-spf-run", types.YLeaf{"TriggerSpfRun", ipfrrEventSummary.TriggerSpfRun})
    ipfrrEventSummary.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", ipfrrEventSummary.WaitTime})
    ipfrrEventSummary.EntityData.Leafs.Append("start-time-offset", types.YLeaf{"StartTimeOffset", ipfrrEventSummary.StartTimeOffset})
    ipfrrEventSummary.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ipfrrEventSummary.Duration})
    ipfrrEventSummary.EntityData.Leafs.Append("completed-spf-run", types.YLeaf{"CompletedSpfRun", ipfrrEventSummary.CompletedSpfRun})
    ipfrrEventSummary.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", ipfrrEventSummary.TotalRoutes})
    ipfrrEventSummary.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", ipfrrEventSummary.FullyProtectedRoutes})
    ipfrrEventSummary.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", ipfrrEventSummary.PartiallyProtectedRoutes})
    ipfrrEventSummary.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", ipfrrEventSummary.Coverage})

    ipfrrEventSummary.EntityData.YListKeys = []string {"EventId"}

    return &(ipfrrEventSummary.EntityData)
}

// Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_IpfrrStatistic
// IP-Frr Statistics categorized by priority
type Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_IpfrrStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Priority. The type is RcmdPriorityLevel.
    Priority interface{}

    // Total Number of Routes. The type is interface{} with range: 0..4294967295.
    TotalRoutes interface{}

    // Fully Protected Routes. The type is interface{} with range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Partially Protected Routes. The type is interface{} with range:
    // 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage. The type is string. Units are percentage.
    Coverage interface{}

    // Local LFA Coverage in percentage. The type is string. Units are percentage.
    LocalLfaCoverage interface{}

    // Remote LFA Coverage in percentage. The type is string. Units are
    // percentage.
    RemoteLfaCoverage interface{}

    // Covearge is below Configured Threshold. The type is bool.
    BelowThreshold interface{}
}

func (ipfrrStatistic *Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_IpfrrStatistic) GetEntityData() *types.CommonEntityData {
    ipfrrStatistic.EntityData.YFilter = ipfrrStatistic.YFilter
    ipfrrStatistic.EntityData.YangName = "ipfrr-statistic"
    ipfrrStatistic.EntityData.BundleName = "cisco_ios_xr"
    ipfrrStatistic.EntityData.ParentYangName = "ipfrr-event-summary"
    ipfrrStatistic.EntityData.SegmentPath = "ipfrr-statistic"
    ipfrrStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipfrrStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipfrrStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipfrrStatistic.EntityData.Children = types.NewOrderedMap()
    ipfrrStatistic.EntityData.Leafs = types.NewOrderedMap()
    ipfrrStatistic.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", ipfrrStatistic.Priority})
    ipfrrStatistic.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", ipfrrStatistic.TotalRoutes})
    ipfrrStatistic.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", ipfrrStatistic.FullyProtectedRoutes})
    ipfrrStatistic.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", ipfrrStatistic.PartiallyProtectedRoutes})
    ipfrrStatistic.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", ipfrrStatistic.Coverage})
    ipfrrStatistic.EntityData.Leafs.Append("local-lfa-coverage", types.YLeaf{"LocalLfaCoverage", ipfrrStatistic.LocalLfaCoverage})
    ipfrrStatistic.EntityData.Leafs.Append("remote-lfa-coverage", types.YLeaf{"RemoteLfaCoverage", ipfrrStatistic.RemoteLfaCoverage})
    ipfrrStatistic.EntityData.Leafs.Append("below-threshold", types.YLeaf{"BelowThreshold", ipfrrStatistic.BelowThreshold})

    ipfrrStatistic.EntityData.YListKeys = []string {}

    return &(ipfrrStatistic.EntityData)
}

// Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode
// Remote Node Information
type Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote-LFA Node ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteNodeId interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Number of paths protected by this Remote Node. The type is interface{} with
    // range: 0..4294967295.
    PathCount interface{}

    // Inuse time of the Remote Node (eg: Apr 24 13:16 :04.961). The type is
    // string.
    InUseTime interface{}

    // Protected Primary Paths. The type is slice of
    // Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode_PrimaryPath.
    PrimaryPath []*Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode_PrimaryPath
}

func (remoteNode *Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode) GetEntityData() *types.CommonEntityData {
    remoteNode.EntityData.YFilter = remoteNode.YFilter
    remoteNode.EntityData.YangName = "remote-node"
    remoteNode.EntityData.BundleName = "cisco_ios_xr"
    remoteNode.EntityData.ParentYangName = "ipfrr-event-summary"
    remoteNode.EntityData.SegmentPath = "remote-node"
    remoteNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNode.EntityData.Children = types.NewOrderedMap()
    remoteNode.EntityData.Children.Append("primary-path", types.YChild{"PrimaryPath", nil})
    for i := range remoteNode.PrimaryPath {
        remoteNode.EntityData.Children.Append(types.GetSegmentPath(remoteNode.PrimaryPath[i]), types.YChild{"PrimaryPath", remoteNode.PrimaryPath[i]})
    }
    remoteNode.EntityData.Leafs = types.NewOrderedMap()
    remoteNode.EntityData.Leafs.Append("remote-node-id", types.YLeaf{"RemoteNodeId", remoteNode.RemoteNodeId})
    remoteNode.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", remoteNode.InterfaceName})
    remoteNode.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", remoteNode.NeighbourAddress})
    remoteNode.EntityData.Leafs.Append("path-count", types.YLeaf{"PathCount", remoteNode.PathCount})
    remoteNode.EntityData.Leafs.Append("in-use-time", types.YLeaf{"InUseTime", remoteNode.InUseTime})

    remoteNode.EntityData.YListKeys = []string {}

    return &(remoteNode.EntityData)
}

// Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode_PrimaryPath
// Protected Primary Paths
type Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode_PrimaryPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}
}

func (primaryPath *Rcmd_Isis_Instances_Instance_IpfrrEventSummaries_IpfrrEventSummary_RemoteNode_PrimaryPath) GetEntityData() *types.CommonEntityData {
    primaryPath.EntityData.YFilter = primaryPath.YFilter
    primaryPath.EntityData.YangName = "primary-path"
    primaryPath.EntityData.BundleName = "cisco_ios_xr"
    primaryPath.EntityData.ParentYangName = "remote-node"
    primaryPath.EntityData.SegmentPath = "primary-path"
    primaryPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryPath.EntityData.Children = types.NewOrderedMap()
    primaryPath.EntityData.Leafs = types.NewOrderedMap()
    primaryPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", primaryPath.InterfaceName})
    primaryPath.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", primaryPath.NeighbourAddress})

    primaryPath.EntityData.YListKeys = []string {}

    return &(primaryPath.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventStatistics
// ISIS Prefix events statistics data
type Rcmd_Isis_Instances_Instance_PrefixEventStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Monitoring Statistics. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventStatistics_PrefixEventStatistic.
    PrefixEventStatistic []*Rcmd_Isis_Instances_Instance_PrefixEventStatistics_PrefixEventStatistic
}

func (prefixEventStatistics *Rcmd_Isis_Instances_Instance_PrefixEventStatistics) GetEntityData() *types.CommonEntityData {
    prefixEventStatistics.EntityData.YFilter = prefixEventStatistics.YFilter
    prefixEventStatistics.EntityData.YangName = "prefix-event-statistics"
    prefixEventStatistics.EntityData.BundleName = "cisco_ios_xr"
    prefixEventStatistics.EntityData.ParentYangName = "instance"
    prefixEventStatistics.EntityData.SegmentPath = "prefix-event-statistics"
    prefixEventStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixEventStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixEventStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixEventStatistics.EntityData.Children = types.NewOrderedMap()
    prefixEventStatistics.EntityData.Children.Append("prefix-event-statistic", types.YChild{"PrefixEventStatistic", nil})
    for i := range prefixEventStatistics.PrefixEventStatistic {
        prefixEventStatistics.EntityData.Children.Append(types.GetSegmentPath(prefixEventStatistics.PrefixEventStatistic[i]), types.YChild{"PrefixEventStatistic", prefixEventStatistics.PrefixEventStatistic[i]})
    }
    prefixEventStatistics.EntityData.Leafs = types.NewOrderedMap()

    prefixEventStatistics.EntityData.YListKeys = []string {}

    return &(prefixEventStatistics.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventStatistics_PrefixEventStatistic
// Monitoring Statistics
type Rcmd_Isis_Instances_Instance_PrefixEventStatistics_PrefixEventStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Events with Prefix. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    PrefixInfo interface{}

    // Prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLenth interface{}

    // Last event trigger time. The type is string.
    LastEventTime interface{}

    // Last event processed priority. The type is RcmdPriorityLevel.
    LastPriority interface{}

    // Last event Route Type. The type is RcmdShowRoute.
    LastRouteType interface{}

    // Last event Add/Delete. The type is RcmdChange.
    LastChangeType interface{}

    // Last Known Cost. The type is interface{} with range: 0..4294967295.
    LastCost interface{}

    // No. of times processed under Critical Priority. The type is interface{}
    // with range: 0..4294967295.
    CriticalPriority interface{}

    // No. of times processed under High Priority. The type is interface{} with
    // range: 0..4294967295.
    HighPriority interface{}

    // No. of times processed under Medium Priority. The type is interface{} with
    // range: 0..4294967295.
    MediumPriority interface{}

    // No. of times processed under Low Priority. The type is interface{} with
    // range: 0..4294967295.
    LowPriority interface{}

    // No. of times route gets Added. The type is interface{} with range:
    // 0..4294967295.
    AddCount interface{}

    // No. of times route gets Deleted. The type is interface{} with range:
    // 0..4294967295.
    ModifyCount interface{}

    // No. of times route gets Deleted. The type is interface{} with range:
    // 0..4294967295.
    DeleteCount interface{}

    // No. of times threshold got exceeded. The type is interface{} with range:
    // 0..4294967295.
    ThresholdExceedCount interface{}
}

func (prefixEventStatistic *Rcmd_Isis_Instances_Instance_PrefixEventStatistics_PrefixEventStatistic) GetEntityData() *types.CommonEntityData {
    prefixEventStatistic.EntityData.YFilter = prefixEventStatistic.YFilter
    prefixEventStatistic.EntityData.YangName = "prefix-event-statistic"
    prefixEventStatistic.EntityData.BundleName = "cisco_ios_xr"
    prefixEventStatistic.EntityData.ParentYangName = "prefix-event-statistics"
    prefixEventStatistic.EntityData.SegmentPath = "prefix-event-statistic" + types.AddKeyToken(prefixEventStatistic.PrefixInfo, "prefix-info")
    prefixEventStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixEventStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixEventStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixEventStatistic.EntityData.Children = types.NewOrderedMap()
    prefixEventStatistic.EntityData.Leafs = types.NewOrderedMap()
    prefixEventStatistic.EntityData.Leafs.Append("prefix-info", types.YLeaf{"PrefixInfo", prefixEventStatistic.PrefixInfo})
    prefixEventStatistic.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", prefixEventStatistic.Prefix})
    prefixEventStatistic.EntityData.Leafs.Append("prefix-lenth", types.YLeaf{"PrefixLenth", prefixEventStatistic.PrefixLenth})
    prefixEventStatistic.EntityData.Leafs.Append("last-event-time", types.YLeaf{"LastEventTime", prefixEventStatistic.LastEventTime})
    prefixEventStatistic.EntityData.Leafs.Append("last-priority", types.YLeaf{"LastPriority", prefixEventStatistic.LastPriority})
    prefixEventStatistic.EntityData.Leafs.Append("last-route-type", types.YLeaf{"LastRouteType", prefixEventStatistic.LastRouteType})
    prefixEventStatistic.EntityData.Leafs.Append("last-change-type", types.YLeaf{"LastChangeType", prefixEventStatistic.LastChangeType})
    prefixEventStatistic.EntityData.Leafs.Append("last-cost", types.YLeaf{"LastCost", prefixEventStatistic.LastCost})
    prefixEventStatistic.EntityData.Leafs.Append("critical-priority", types.YLeaf{"CriticalPriority", prefixEventStatistic.CriticalPriority})
    prefixEventStatistic.EntityData.Leafs.Append("high-priority", types.YLeaf{"HighPriority", prefixEventStatistic.HighPriority})
    prefixEventStatistic.EntityData.Leafs.Append("medium-priority", types.YLeaf{"MediumPriority", prefixEventStatistic.MediumPriority})
    prefixEventStatistic.EntityData.Leafs.Append("low-priority", types.YLeaf{"LowPriority", prefixEventStatistic.LowPriority})
    prefixEventStatistic.EntityData.Leafs.Append("add-count", types.YLeaf{"AddCount", prefixEventStatistic.AddCount})
    prefixEventStatistic.EntityData.Leafs.Append("modify-count", types.YLeaf{"ModifyCount", prefixEventStatistic.ModifyCount})
    prefixEventStatistic.EntityData.Leafs.Append("delete-count", types.YLeaf{"DeleteCount", prefixEventStatistic.DeleteCount})
    prefixEventStatistic.EntityData.Leafs.Append("threshold-exceed-count", types.YLeaf{"ThresholdExceedCount", prefixEventStatistic.ThresholdExceedCount})

    prefixEventStatistic.EntityData.YListKeys = []string {"PrefixInfo"}

    return &(prefixEventStatistic.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries
// ISIS SPF run summary data
type Rcmd_Isis_Instances_Instance_SpfRunSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SPF Event data. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary.
    SpfRunSummary []*Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary
}

func (spfRunSummaries *Rcmd_Isis_Instances_Instance_SpfRunSummaries) GetEntityData() *types.CommonEntityData {
    spfRunSummaries.EntityData.YFilter = spfRunSummaries.YFilter
    spfRunSummaries.EntityData.YangName = "spf-run-summaries"
    spfRunSummaries.EntityData.BundleName = "cisco_ios_xr"
    spfRunSummaries.EntityData.ParentYangName = "instance"
    spfRunSummaries.EntityData.SegmentPath = "spf-run-summaries"
    spfRunSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfRunSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfRunSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfRunSummaries.EntityData.Children = types.NewOrderedMap()
    spfRunSummaries.EntityData.Children.Append("spf-run-summary", types.YChild{"SpfRunSummary", nil})
    for i := range spfRunSummaries.SpfRunSummary {
        spfRunSummaries.EntityData.Children.Append(types.GetSegmentPath(spfRunSummaries.SpfRunSummary[i]), types.YChild{"SpfRunSummary", spfRunSummaries.SpfRunSummary[i]})
    }
    spfRunSummaries.EntityData.Leafs = types.NewOrderedMap()

    spfRunSummaries.EntityData.YListKeys = []string {}

    return &(spfRunSummaries.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary
// SPF Event data
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific SPF run. The type is interface{} with
    // range: 1..4294967295.
    SpfRunNumber interface{}

    // Start time (offset from event trigger time in ss .msec). The type is
    // string.
    StartTime interface{}

    // Wait time applied at SPF schedule (in msec). The type is interface{} with
    // range: 0..4294967295.
    WaitTime interface{}

    // Trigger reasons for SPF run. Example: pr^ - periodic, cr^ - clear (Check
    // the documentation for the entire list). The type is string.
    Reason interface{}

    // SPF summary information.
    SpfSummary Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary

    // SPF Node statistics.
    NodeStatistics Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_NodeStatistics

    // Trigger LSP. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_TriggerLsp.
    TriggerLsp []*Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_TriggerLsp

    // Convergence information on per-priority basis. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority.
    Priority []*Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority

    // List of LSP changes processed. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_LspProcessed.
    LspProcessed []*Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_LspProcessed

    // List of LSP regenerated. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_LspRegenerated.
    LspRegenerated []*Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_LspRegenerated
}

func (spfRunSummary *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary) GetEntityData() *types.CommonEntityData {
    spfRunSummary.EntityData.YFilter = spfRunSummary.YFilter
    spfRunSummary.EntityData.YangName = "spf-run-summary"
    spfRunSummary.EntityData.BundleName = "cisco_ios_xr"
    spfRunSummary.EntityData.ParentYangName = "spf-run-summaries"
    spfRunSummary.EntityData.SegmentPath = "spf-run-summary" + types.AddKeyToken(spfRunSummary.SpfRunNumber, "spf-run-number")
    spfRunSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfRunSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfRunSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfRunSummary.EntityData.Children = types.NewOrderedMap()
    spfRunSummary.EntityData.Children.Append("spf-summary", types.YChild{"SpfSummary", &spfRunSummary.SpfSummary})
    spfRunSummary.EntityData.Children.Append("node-statistics", types.YChild{"NodeStatistics", &spfRunSummary.NodeStatistics})
    spfRunSummary.EntityData.Children.Append("trigger-lsp", types.YChild{"TriggerLsp", nil})
    for i := range spfRunSummary.TriggerLsp {
        spfRunSummary.EntityData.Children.Append(types.GetSegmentPath(spfRunSummary.TriggerLsp[i]), types.YChild{"TriggerLsp", spfRunSummary.TriggerLsp[i]})
    }
    spfRunSummary.EntityData.Children.Append("priority", types.YChild{"Priority", nil})
    for i := range spfRunSummary.Priority {
        spfRunSummary.EntityData.Children.Append(types.GetSegmentPath(spfRunSummary.Priority[i]), types.YChild{"Priority", spfRunSummary.Priority[i]})
    }
    spfRunSummary.EntityData.Children.Append("lsp-processed", types.YChild{"LspProcessed", nil})
    for i := range spfRunSummary.LspProcessed {
        spfRunSummary.EntityData.Children.Append(types.GetSegmentPath(spfRunSummary.LspProcessed[i]), types.YChild{"LspProcessed", spfRunSummary.LspProcessed[i]})
    }
    spfRunSummary.EntityData.Children.Append("lsp-regenerated", types.YChild{"LspRegenerated", nil})
    for i := range spfRunSummary.LspRegenerated {
        spfRunSummary.EntityData.Children.Append(types.GetSegmentPath(spfRunSummary.LspRegenerated[i]), types.YChild{"LspRegenerated", spfRunSummary.LspRegenerated[i]})
    }
    spfRunSummary.EntityData.Leafs = types.NewOrderedMap()
    spfRunSummary.EntityData.Leafs.Append("spf-run-number", types.YLeaf{"SpfRunNumber", spfRunSummary.SpfRunNumber})
    spfRunSummary.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", spfRunSummary.StartTime})
    spfRunSummary.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", spfRunSummary.WaitTime})
    spfRunSummary.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", spfRunSummary.Reason})

    spfRunSummary.EntityData.YListKeys = []string {"SpfRunNumber"}

    return &(spfRunSummary.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary
// SPF summary information
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Topology index (multi-topology). The type is interface{} with range:
    // 0..4294967295.
    Topology interface{}

    // ISIS Level. The type is RcmdIsisLvl.
    IsisLevel interface{}

    // Type of SPF. The type is RcmdIsisSpf.
    Type interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // SPF state. The type is RcmdSpfState.
    State interface{}

    // Whether the event has all information. The type is bool.
    IsDataComplete interface{}

    // Trigger time (in hh:mm:ss.msec). The type is string.
    TriggerTime interface{}

    // Duration of SPF calculation (in ss.msec). The type is string.
    Duration interface{}

    // Total number of LSP changes processed. The type is interface{} with range:
    // 0..65535.
    TotalLspChanges interface{}

    // Convergence information summary on per-priority basis. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary.
    PrioritySummary []*Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary
}

func (spfSummary *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary) GetEntityData() *types.CommonEntityData {
    spfSummary.EntityData.YFilter = spfSummary.YFilter
    spfSummary.EntityData.YangName = "spf-summary"
    spfSummary.EntityData.BundleName = "cisco_ios_xr"
    spfSummary.EntityData.ParentYangName = "spf-run-summary"
    spfSummary.EntityData.SegmentPath = "spf-summary"
    spfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfSummary.EntityData.Children = types.NewOrderedMap()
    spfSummary.EntityData.Children.Append("priority-summary", types.YChild{"PrioritySummary", nil})
    for i := range spfSummary.PrioritySummary {
        spfSummary.EntityData.Children.Append(types.GetSegmentPath(spfSummary.PrioritySummary[i]), types.YChild{"PrioritySummary", spfSummary.PrioritySummary[i]})
    }
    spfSummary.EntityData.Leafs = types.NewOrderedMap()
    spfSummary.EntityData.Leafs.Append("topology", types.YLeaf{"Topology", spfSummary.Topology})
    spfSummary.EntityData.Leafs.Append("isis-level", types.YLeaf{"IsisLevel", spfSummary.IsisLevel})
    spfSummary.EntityData.Leafs.Append("type", types.YLeaf{"Type", spfSummary.Type})
    spfSummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", spfSummary.ThresholdExceeded})
    spfSummary.EntityData.Leafs.Append("state", types.YLeaf{"State", spfSummary.State})
    spfSummary.EntityData.Leafs.Append("is-data-complete", types.YLeaf{"IsDataComplete", spfSummary.IsDataComplete})
    spfSummary.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", spfSummary.TriggerTime})
    spfSummary.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", spfSummary.Duration})
    spfSummary.EntityData.Leafs.Append("total-lsp-changes", types.YLeaf{"TotalLspChanges", spfSummary.TotalLspChanges})

    spfSummary.EntityData.YListKeys = []string {}

    return &(spfSummary.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary
// Convergence information summary on per-priority
// basis
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Critical, High, Medium or Low. The type is RcmdPriorityLevel.
    Level interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Route statistics.
    RouteStatistics Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_RouteStatistics

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_MplsConvergenceTime

    // Fast Re-Route Statistics. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_FrrStatistic.
    FrrStatistic []*Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_FrrStatistic
}

func (prioritySummary *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary) GetEntityData() *types.CommonEntityData {
    prioritySummary.EntityData.YFilter = prioritySummary.YFilter
    prioritySummary.EntityData.YangName = "priority-summary"
    prioritySummary.EntityData.BundleName = "cisco_ios_xr"
    prioritySummary.EntityData.ParentYangName = "spf-summary"
    prioritySummary.EntityData.SegmentPath = "priority-summary"
    prioritySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prioritySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prioritySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prioritySummary.EntityData.Children = types.NewOrderedMap()
    prioritySummary.EntityData.Children.Append("route-statistics", types.YChild{"RouteStatistics", &prioritySummary.RouteStatistics})
    prioritySummary.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prioritySummary.IpConvergenceTime})
    prioritySummary.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prioritySummary.MplsConvergenceTime})
    prioritySummary.EntityData.Children.Append("frr-statistic", types.YChild{"FrrStatistic", nil})
    for i := range prioritySummary.FrrStatistic {
        prioritySummary.EntityData.Children.Append(types.GetSegmentPath(prioritySummary.FrrStatistic[i]), types.YChild{"FrrStatistic", prioritySummary.FrrStatistic[i]})
    }
    prioritySummary.EntityData.Leafs = types.NewOrderedMap()
    prioritySummary.EntityData.Leafs.Append("level", types.YLeaf{"Level", prioritySummary.Level})
    prioritySummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prioritySummary.ThresholdExceeded})

    prioritySummary.EntityData.YListKeys = []string {}

    return &(prioritySummary.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_RouteStatistics
// Route statistics
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_RouteStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Added. The type is interface{} with range: 0..4294967295.
    Adds interface{}

    // Deleted. The type is interface{} with range: 0..4294967295.
    Deletes interface{}

    // Modified. The type is interface{} with range: 0..4294967295.
    Modifies interface{}

    // Reachable. The type is interface{} with range: 0..4294967295.
    Reachables interface{}

    // Unreachable. The type is interface{} with range: 0..4294967295.
    Unreachables interface{}

    // Touched. The type is interface{} with range: 0..4294967295.
    Touches interface{}
}

func (routeStatistics *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_RouteStatistics) GetEntityData() *types.CommonEntityData {
    routeStatistics.EntityData.YFilter = routeStatistics.YFilter
    routeStatistics.EntityData.YangName = "route-statistics"
    routeStatistics.EntityData.BundleName = "cisco_ios_xr"
    routeStatistics.EntityData.ParentYangName = "priority-summary"
    routeStatistics.EntityData.SegmentPath = "route-statistics"
    routeStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeStatistics.EntityData.Children = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs.Append("adds", types.YLeaf{"Adds", routeStatistics.Adds})
    routeStatistics.EntityData.Leafs.Append("deletes", types.YLeaf{"Deletes", routeStatistics.Deletes})
    routeStatistics.EntityData.Leafs.Append("modifies", types.YLeaf{"Modifies", routeStatistics.Modifies})
    routeStatistics.EntityData.Leafs.Append("reachables", types.YLeaf{"Reachables", routeStatistics.Reachables})
    routeStatistics.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", routeStatistics.Unreachables})
    routeStatistics.EntityData.Leafs.Append("touches", types.YLeaf{"Touches", routeStatistics.Touches})

    routeStatistics.EntityData.YListKeys = []string {}

    return &(routeStatistics.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "priority-summary"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "priority-summary"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_FrrStatistic
// Fast Re-Route Statistics
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_FrrStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total Number of Routes. The type is interface{} with range: 0..4294967295.
    TotalRoutes interface{}

    // Fully Protected Routes. The type is interface{} with range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Partially Protected Routes. The type is interface{} with range:
    // 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage. The type is string. Units are percentage.
    Coverage interface{}
}

func (frrStatistic *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_SpfSummary_PrioritySummary_FrrStatistic) GetEntityData() *types.CommonEntityData {
    frrStatistic.EntityData.YFilter = frrStatistic.YFilter
    frrStatistic.EntityData.YangName = "frr-statistic"
    frrStatistic.EntityData.BundleName = "cisco_ios_xr"
    frrStatistic.EntityData.ParentYangName = "priority-summary"
    frrStatistic.EntityData.SegmentPath = "frr-statistic"
    frrStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrStatistic.EntityData.Children = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", frrStatistic.TotalRoutes})
    frrStatistic.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", frrStatistic.FullyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", frrStatistic.PartiallyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", frrStatistic.Coverage})

    frrStatistic.EntityData.YListKeys = []string {}

    return &(frrStatistic.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_NodeStatistics
// SPF Node statistics
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_NodeStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Added. The type is interface{} with range: 0..4294967295.
    Adds interface{}

    // Deleted. The type is interface{} with range: 0..4294967295.
    Deletes interface{}

    // Modified. The type is interface{} with range: 0..4294967295.
    Modifies interface{}

    // Reachable. The type is interface{} with range: 0..4294967295.
    Reachables interface{}

    // Unreachable. The type is interface{} with range: 0..4294967295.
    Unreachables interface{}

    // Touched. The type is interface{} with range: 0..4294967295.
    Touches interface{}
}

func (nodeStatistics *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_NodeStatistics) GetEntityData() *types.CommonEntityData {
    nodeStatistics.EntityData.YFilter = nodeStatistics.YFilter
    nodeStatistics.EntityData.YangName = "node-statistics"
    nodeStatistics.EntityData.BundleName = "cisco_ios_xr"
    nodeStatistics.EntityData.ParentYangName = "spf-run-summary"
    nodeStatistics.EntityData.SegmentPath = "node-statistics"
    nodeStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeStatistics.EntityData.Children = types.NewOrderedMap()
    nodeStatistics.EntityData.Leafs = types.NewOrderedMap()
    nodeStatistics.EntityData.Leafs.Append("adds", types.YLeaf{"Adds", nodeStatistics.Adds})
    nodeStatistics.EntityData.Leafs.Append("deletes", types.YLeaf{"Deletes", nodeStatistics.Deletes})
    nodeStatistics.EntityData.Leafs.Append("modifies", types.YLeaf{"Modifies", nodeStatistics.Modifies})
    nodeStatistics.EntityData.Leafs.Append("reachables", types.YLeaf{"Reachables", nodeStatistics.Reachables})
    nodeStatistics.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", nodeStatistics.Unreachables})
    nodeStatistics.EntityData.Leafs.Append("touches", types.YLeaf{"Touches", nodeStatistics.Touches})

    nodeStatistics.EntityData.YListKeys = []string {}

    return &(nodeStatistics.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_TriggerLsp
// Trigger LSP
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_TriggerLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP ID. The type is string.
    LspId interface{}

    // Sequence Number. The type is interface{} with range: 0..4294967295.
    SequenceNumber interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (triggerLsp *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_TriggerLsp) GetEntityData() *types.CommonEntityData {
    triggerLsp.EntityData.YFilter = triggerLsp.YFilter
    triggerLsp.EntityData.YangName = "trigger-lsp"
    triggerLsp.EntityData.BundleName = "cisco_ios_xr"
    triggerLsp.EntityData.ParentYangName = "spf-run-summary"
    triggerLsp.EntityData.SegmentPath = "trigger-lsp"
    triggerLsp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    triggerLsp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    triggerLsp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    triggerLsp.EntityData.Children = types.NewOrderedMap()
    triggerLsp.EntityData.Leafs = types.NewOrderedMap()
    triggerLsp.EntityData.Leafs.Append("lsp-id", types.YLeaf{"LspId", triggerLsp.LspId})
    triggerLsp.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", triggerLsp.SequenceNumber})
    triggerLsp.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", triggerLsp.ChangeType})
    triggerLsp.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", triggerLsp.ReceptionTime})

    triggerLsp.EntityData.YListKeys = []string {}

    return &(triggerLsp.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority
// Convergence information on per-priority basis
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of the priority.
    PrioritySummary Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary

    // Convergence timeline details. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline.
    ConvergenceTimeline []*Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline

    // List of Leaf Networks Added. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_LeafNetworksAdded.
    LeafNetworksAdded []*Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_LeafNetworksAdded

    // List of Leaf Networks Deleted. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_LeafNetworksDeleted.
    LeafNetworksDeleted []*Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_LeafNetworksDeleted
}

func (priority *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "spf-run-summary"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Children.Append("priority-summary", types.YChild{"PrioritySummary", &priority.PrioritySummary})
    priority.EntityData.Children.Append("convergence-timeline", types.YChild{"ConvergenceTimeline", nil})
    for i := range priority.ConvergenceTimeline {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.ConvergenceTimeline[i]), types.YChild{"ConvergenceTimeline", priority.ConvergenceTimeline[i]})
    }
    priority.EntityData.Children.Append("leaf-networks-added", types.YChild{"LeafNetworksAdded", nil})
    for i := range priority.LeafNetworksAdded {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.LeafNetworksAdded[i]), types.YChild{"LeafNetworksAdded", priority.LeafNetworksAdded[i]})
    }
    priority.EntityData.Children.Append("leaf-networks-deleted", types.YChild{"LeafNetworksDeleted", nil})
    for i := range priority.LeafNetworksDeleted {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.LeafNetworksDeleted[i]), types.YChild{"LeafNetworksDeleted", priority.LeafNetworksDeleted[i]})
    }
    priority.EntityData.Leafs = types.NewOrderedMap()

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary
// Summary of the priority
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Critical, High, Medium or Low. The type is RcmdPriorityLevel.
    Level interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Route statistics.
    RouteStatistics Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_RouteStatistics

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_MplsConvergenceTime

    // Fast Re-Route Statistics. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_FrrStatistic.
    FrrStatistic []*Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_FrrStatistic
}

func (prioritySummary *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary) GetEntityData() *types.CommonEntityData {
    prioritySummary.EntityData.YFilter = prioritySummary.YFilter
    prioritySummary.EntityData.YangName = "priority-summary"
    prioritySummary.EntityData.BundleName = "cisco_ios_xr"
    prioritySummary.EntityData.ParentYangName = "priority"
    prioritySummary.EntityData.SegmentPath = "priority-summary"
    prioritySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prioritySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prioritySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prioritySummary.EntityData.Children = types.NewOrderedMap()
    prioritySummary.EntityData.Children.Append("route-statistics", types.YChild{"RouteStatistics", &prioritySummary.RouteStatistics})
    prioritySummary.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prioritySummary.IpConvergenceTime})
    prioritySummary.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prioritySummary.MplsConvergenceTime})
    prioritySummary.EntityData.Children.Append("frr-statistic", types.YChild{"FrrStatistic", nil})
    for i := range prioritySummary.FrrStatistic {
        prioritySummary.EntityData.Children.Append(types.GetSegmentPath(prioritySummary.FrrStatistic[i]), types.YChild{"FrrStatistic", prioritySummary.FrrStatistic[i]})
    }
    prioritySummary.EntityData.Leafs = types.NewOrderedMap()
    prioritySummary.EntityData.Leafs.Append("level", types.YLeaf{"Level", prioritySummary.Level})
    prioritySummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prioritySummary.ThresholdExceeded})

    prioritySummary.EntityData.YListKeys = []string {}

    return &(prioritySummary.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_RouteStatistics
// Route statistics
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_RouteStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Added. The type is interface{} with range: 0..4294967295.
    Adds interface{}

    // Deleted. The type is interface{} with range: 0..4294967295.
    Deletes interface{}

    // Modified. The type is interface{} with range: 0..4294967295.
    Modifies interface{}

    // Reachable. The type is interface{} with range: 0..4294967295.
    Reachables interface{}

    // Unreachable. The type is interface{} with range: 0..4294967295.
    Unreachables interface{}

    // Touched. The type is interface{} with range: 0..4294967295.
    Touches interface{}
}

func (routeStatistics *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_RouteStatistics) GetEntityData() *types.CommonEntityData {
    routeStatistics.EntityData.YFilter = routeStatistics.YFilter
    routeStatistics.EntityData.YangName = "route-statistics"
    routeStatistics.EntityData.BundleName = "cisco_ios_xr"
    routeStatistics.EntityData.ParentYangName = "priority-summary"
    routeStatistics.EntityData.SegmentPath = "route-statistics"
    routeStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeStatistics.EntityData.Children = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs.Append("adds", types.YLeaf{"Adds", routeStatistics.Adds})
    routeStatistics.EntityData.Leafs.Append("deletes", types.YLeaf{"Deletes", routeStatistics.Deletes})
    routeStatistics.EntityData.Leafs.Append("modifies", types.YLeaf{"Modifies", routeStatistics.Modifies})
    routeStatistics.EntityData.Leafs.Append("reachables", types.YLeaf{"Reachables", routeStatistics.Reachables})
    routeStatistics.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", routeStatistics.Unreachables})
    routeStatistics.EntityData.Leafs.Append("touches", types.YLeaf{"Touches", routeStatistics.Touches})

    routeStatistics.EntityData.YListKeys = []string {}

    return &(routeStatistics.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "priority-summary"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "priority-summary"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_FrrStatistic
// Fast Re-Route Statistics
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_FrrStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total Number of Routes. The type is interface{} with range: 0..4294967295.
    TotalRoutes interface{}

    // Fully Protected Routes. The type is interface{} with range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Partially Protected Routes. The type is interface{} with range:
    // 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage. The type is string. Units are percentage.
    Coverage interface{}
}

func (frrStatistic *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_PrioritySummary_FrrStatistic) GetEntityData() *types.CommonEntityData {
    frrStatistic.EntityData.YFilter = frrStatistic.YFilter
    frrStatistic.EntityData.YangName = "frr-statistic"
    frrStatistic.EntityData.BundleName = "cisco_ios_xr"
    frrStatistic.EntityData.ParentYangName = "priority-summary"
    frrStatistic.EntityData.SegmentPath = "frr-statistic"
    frrStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrStatistic.EntityData.Children = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", frrStatistic.TotalRoutes})
    frrStatistic.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", frrStatistic.FullyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", frrStatistic.PartiallyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", frrStatistic.Coverage})

    frrStatistic.EntityData.YListKeys = []string {}

    return &(frrStatistic.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline
// Convergence timeline details
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route origin (routing protocol).
    RouteOrigin Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RouteOrigin

    // Entry point of IPv4 RIB.
    RiBv4Enter Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RiBv4Enter

    // Exit point from IPv4 RIB to FIBs.
    RiBv4Exit Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RiBv4Exit

    // Route Redistribute point from IPv4 RIB to LDP.
    RiBv4Redistribute Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RiBv4Redistribute

    // Entry point of LDP.
    LdpEnter Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LdpEnter

    // Exit point of LDP to LSD.
    LdpExit Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LdpExit

    // Entry point of LSD.
    LsdEnter Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LsdEnter

    // Exit point of LSD to FIBs.
    LsdExit Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LsdExit

    // List of Linecards' completion point for IP routes. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcIp.
    LcIp []*Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcIp

    // List of Linecards' completion point for MPLS labels. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcMpls.
    LcMpls []*Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcMpls
}

func (convergenceTimeline *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline) GetEntityData() *types.CommonEntityData {
    convergenceTimeline.EntityData.YFilter = convergenceTimeline.YFilter
    convergenceTimeline.EntityData.YangName = "convergence-timeline"
    convergenceTimeline.EntityData.BundleName = "cisco_ios_xr"
    convergenceTimeline.EntityData.ParentYangName = "priority"
    convergenceTimeline.EntityData.SegmentPath = "convergence-timeline"
    convergenceTimeline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    convergenceTimeline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    convergenceTimeline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    convergenceTimeline.EntityData.Children = types.NewOrderedMap()
    convergenceTimeline.EntityData.Children.Append("route-origin", types.YChild{"RouteOrigin", &convergenceTimeline.RouteOrigin})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-enter", types.YChild{"RiBv4Enter", &convergenceTimeline.RiBv4Enter})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-exit", types.YChild{"RiBv4Exit", &convergenceTimeline.RiBv4Exit})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-redistribute", types.YChild{"RiBv4Redistribute", &convergenceTimeline.RiBv4Redistribute})
    convergenceTimeline.EntityData.Children.Append("ldp-enter", types.YChild{"LdpEnter", &convergenceTimeline.LdpEnter})
    convergenceTimeline.EntityData.Children.Append("ldp-exit", types.YChild{"LdpExit", &convergenceTimeline.LdpExit})
    convergenceTimeline.EntityData.Children.Append("lsd-enter", types.YChild{"LsdEnter", &convergenceTimeline.LsdEnter})
    convergenceTimeline.EntityData.Children.Append("lsd-exit", types.YChild{"LsdExit", &convergenceTimeline.LsdExit})
    convergenceTimeline.EntityData.Children.Append("lc-ip", types.YChild{"LcIp", nil})
    for i := range convergenceTimeline.LcIp {
        convergenceTimeline.EntityData.Children.Append(types.GetSegmentPath(convergenceTimeline.LcIp[i]), types.YChild{"LcIp", convergenceTimeline.LcIp[i]})
    }
    convergenceTimeline.EntityData.Children.Append("lc-mpls", types.YChild{"LcMpls", nil})
    for i := range convergenceTimeline.LcMpls {
        convergenceTimeline.EntityData.Children.Append(types.GetSegmentPath(convergenceTimeline.LcMpls[i]), types.YChild{"LcMpls", convergenceTimeline.LcMpls[i]})
    }
    convergenceTimeline.EntityData.Leafs = types.NewOrderedMap()

    convergenceTimeline.EntityData.YListKeys = []string {}

    return &(convergenceTimeline.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RouteOrigin
// Route origin (routing protocol)
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RouteOrigin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (routeOrigin *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RouteOrigin) GetEntityData() *types.CommonEntityData {
    routeOrigin.EntityData.YFilter = routeOrigin.YFilter
    routeOrigin.EntityData.YangName = "route-origin"
    routeOrigin.EntityData.BundleName = "cisco_ios_xr"
    routeOrigin.EntityData.ParentYangName = "convergence-timeline"
    routeOrigin.EntityData.SegmentPath = "route-origin"
    routeOrigin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeOrigin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeOrigin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeOrigin.EntityData.Children = types.NewOrderedMap()
    routeOrigin.EntityData.Leafs = types.NewOrderedMap()
    routeOrigin.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", routeOrigin.StartTime})
    routeOrigin.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", routeOrigin.EndTime})
    routeOrigin.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", routeOrigin.Duration})

    routeOrigin.EntityData.YListKeys = []string {}

    return &(routeOrigin.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RiBv4Enter
// Entry point of IPv4 RIB
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RiBv4Enter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Enter *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RiBv4Enter) GetEntityData() *types.CommonEntityData {
    riBv4Enter.EntityData.YFilter = riBv4Enter.YFilter
    riBv4Enter.EntityData.YangName = "ri-bv4-enter"
    riBv4Enter.EntityData.BundleName = "cisco_ios_xr"
    riBv4Enter.EntityData.ParentYangName = "convergence-timeline"
    riBv4Enter.EntityData.SegmentPath = "ri-bv4-enter"
    riBv4Enter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Enter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Enter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Enter.EntityData.Children = types.NewOrderedMap()
    riBv4Enter.EntityData.Leafs = types.NewOrderedMap()
    riBv4Enter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Enter.StartTime})
    riBv4Enter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Enter.EndTime})
    riBv4Enter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Enter.Duration})

    riBv4Enter.EntityData.YListKeys = []string {}

    return &(riBv4Enter.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RiBv4Exit
// Exit point from IPv4 RIB to FIBs
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RiBv4Exit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Exit *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RiBv4Exit) GetEntityData() *types.CommonEntityData {
    riBv4Exit.EntityData.YFilter = riBv4Exit.YFilter
    riBv4Exit.EntityData.YangName = "ri-bv4-exit"
    riBv4Exit.EntityData.BundleName = "cisco_ios_xr"
    riBv4Exit.EntityData.ParentYangName = "convergence-timeline"
    riBv4Exit.EntityData.SegmentPath = "ri-bv4-exit"
    riBv4Exit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Exit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Exit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Exit.EntityData.Children = types.NewOrderedMap()
    riBv4Exit.EntityData.Leafs = types.NewOrderedMap()
    riBv4Exit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Exit.StartTime})
    riBv4Exit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Exit.EndTime})
    riBv4Exit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Exit.Duration})

    riBv4Exit.EntityData.YListKeys = []string {}

    return &(riBv4Exit.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RiBv4Redistribute
// Route Redistribute point from IPv4 RIB to LDP
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RiBv4Redistribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Redistribute *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_RiBv4Redistribute) GetEntityData() *types.CommonEntityData {
    riBv4Redistribute.EntityData.YFilter = riBv4Redistribute.YFilter
    riBv4Redistribute.EntityData.YangName = "ri-bv4-redistribute"
    riBv4Redistribute.EntityData.BundleName = "cisco_ios_xr"
    riBv4Redistribute.EntityData.ParentYangName = "convergence-timeline"
    riBv4Redistribute.EntityData.SegmentPath = "ri-bv4-redistribute"
    riBv4Redistribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Redistribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Redistribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Redistribute.EntityData.Children = types.NewOrderedMap()
    riBv4Redistribute.EntityData.Leafs = types.NewOrderedMap()
    riBv4Redistribute.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Redistribute.StartTime})
    riBv4Redistribute.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Redistribute.EndTime})
    riBv4Redistribute.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Redistribute.Duration})

    riBv4Redistribute.EntityData.YListKeys = []string {}

    return &(riBv4Redistribute.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LdpEnter
// Entry point of LDP
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LdpEnter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (ldpEnter *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LdpEnter) GetEntityData() *types.CommonEntityData {
    ldpEnter.EntityData.YFilter = ldpEnter.YFilter
    ldpEnter.EntityData.YangName = "ldp-enter"
    ldpEnter.EntityData.BundleName = "cisco_ios_xr"
    ldpEnter.EntityData.ParentYangName = "convergence-timeline"
    ldpEnter.EntityData.SegmentPath = "ldp-enter"
    ldpEnter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpEnter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpEnter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpEnter.EntityData.Children = types.NewOrderedMap()
    ldpEnter.EntityData.Leafs = types.NewOrderedMap()
    ldpEnter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", ldpEnter.StartTime})
    ldpEnter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", ldpEnter.EndTime})
    ldpEnter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ldpEnter.Duration})

    ldpEnter.EntityData.YListKeys = []string {}

    return &(ldpEnter.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LdpExit
// Exit point of LDP to LSD
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LdpExit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (ldpExit *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LdpExit) GetEntityData() *types.CommonEntityData {
    ldpExit.EntityData.YFilter = ldpExit.YFilter
    ldpExit.EntityData.YangName = "ldp-exit"
    ldpExit.EntityData.BundleName = "cisco_ios_xr"
    ldpExit.EntityData.ParentYangName = "convergence-timeline"
    ldpExit.EntityData.SegmentPath = "ldp-exit"
    ldpExit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpExit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpExit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpExit.EntityData.Children = types.NewOrderedMap()
    ldpExit.EntityData.Leafs = types.NewOrderedMap()
    ldpExit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", ldpExit.StartTime})
    ldpExit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", ldpExit.EndTime})
    ldpExit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ldpExit.Duration})

    ldpExit.EntityData.YListKeys = []string {}

    return &(ldpExit.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LsdEnter
// Entry point of LSD
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LsdEnter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (lsdEnter *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LsdEnter) GetEntityData() *types.CommonEntityData {
    lsdEnter.EntityData.YFilter = lsdEnter.YFilter
    lsdEnter.EntityData.YangName = "lsd-enter"
    lsdEnter.EntityData.BundleName = "cisco_ios_xr"
    lsdEnter.EntityData.ParentYangName = "convergence-timeline"
    lsdEnter.EntityData.SegmentPath = "lsd-enter"
    lsdEnter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsdEnter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsdEnter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsdEnter.EntityData.Children = types.NewOrderedMap()
    lsdEnter.EntityData.Leafs = types.NewOrderedMap()
    lsdEnter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", lsdEnter.StartTime})
    lsdEnter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", lsdEnter.EndTime})
    lsdEnter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", lsdEnter.Duration})

    lsdEnter.EntityData.YListKeys = []string {}

    return &(lsdEnter.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LsdExit
// Exit point of LSD to FIBs
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LsdExit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (lsdExit *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LsdExit) GetEntityData() *types.CommonEntityData {
    lsdExit.EntityData.YFilter = lsdExit.YFilter
    lsdExit.EntityData.YangName = "lsd-exit"
    lsdExit.EntityData.BundleName = "cisco_ios_xr"
    lsdExit.EntityData.ParentYangName = "convergence-timeline"
    lsdExit.EntityData.SegmentPath = "lsd-exit"
    lsdExit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsdExit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsdExit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsdExit.EntityData.Children = types.NewOrderedMap()
    lsdExit.EntityData.Leafs = types.NewOrderedMap()
    lsdExit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", lsdExit.StartTime})
    lsdExit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", lsdExit.EndTime})
    lsdExit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", lsdExit.Duration})

    lsdExit.EntityData.YListKeys = []string {}

    return &(lsdExit.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcIp
// List of Linecards' completion point for IP
// routes
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB.
    FibComplete Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcIp_FibComplete
}

func (lcIp *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcIp) GetEntityData() *types.CommonEntityData {
    lcIp.EntityData.YFilter = lcIp.YFilter
    lcIp.EntityData.YangName = "lc-ip"
    lcIp.EntityData.BundleName = "cisco_ios_xr"
    lcIp.EntityData.ParentYangName = "convergence-timeline"
    lcIp.EntityData.SegmentPath = "lc-ip"
    lcIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcIp.EntityData.Children = types.NewOrderedMap()
    lcIp.EntityData.Children.Append("fib-complete", types.YChild{"FibComplete", &lcIp.FibComplete})
    lcIp.EntityData.Leafs = types.NewOrderedMap()
    lcIp.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcIp.NodeName})
    lcIp.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcIp.Speed})

    lcIp.EntityData.YListKeys = []string {}

    return &(lcIp.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcIp_FibComplete
// Completion point of FIB
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcIp_FibComplete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (fibComplete *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcIp_FibComplete) GetEntityData() *types.CommonEntityData {
    fibComplete.EntityData.YFilter = fibComplete.YFilter
    fibComplete.EntityData.YangName = "fib-complete"
    fibComplete.EntityData.BundleName = "cisco_ios_xr"
    fibComplete.EntityData.ParentYangName = "lc-ip"
    fibComplete.EntityData.SegmentPath = "fib-complete"
    fibComplete.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibComplete.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibComplete.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibComplete.EntityData.Children = types.NewOrderedMap()
    fibComplete.EntityData.Leafs = types.NewOrderedMap()
    fibComplete.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", fibComplete.StartTime})
    fibComplete.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", fibComplete.EndTime})
    fibComplete.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", fibComplete.Duration})

    fibComplete.EntityData.YListKeys = []string {}

    return &(fibComplete.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcMpls
// List of Linecards' completion point for MPLS
// labels
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB.
    FibComplete Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcMpls_FibComplete
}

func (lcMpls *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcMpls) GetEntityData() *types.CommonEntityData {
    lcMpls.EntityData.YFilter = lcMpls.YFilter
    lcMpls.EntityData.YangName = "lc-mpls"
    lcMpls.EntityData.BundleName = "cisco_ios_xr"
    lcMpls.EntityData.ParentYangName = "convergence-timeline"
    lcMpls.EntityData.SegmentPath = "lc-mpls"
    lcMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcMpls.EntityData.Children = types.NewOrderedMap()
    lcMpls.EntityData.Children.Append("fib-complete", types.YChild{"FibComplete", &lcMpls.FibComplete})
    lcMpls.EntityData.Leafs = types.NewOrderedMap()
    lcMpls.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcMpls.NodeName})
    lcMpls.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcMpls.Speed})

    lcMpls.EntityData.YListKeys = []string {}

    return &(lcMpls.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcMpls_FibComplete
// Completion point of FIB
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcMpls_FibComplete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (fibComplete *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_ConvergenceTimeline_LcMpls_FibComplete) GetEntityData() *types.CommonEntityData {
    fibComplete.EntityData.YFilter = fibComplete.YFilter
    fibComplete.EntityData.YangName = "fib-complete"
    fibComplete.EntityData.BundleName = "cisco_ios_xr"
    fibComplete.EntityData.ParentYangName = "lc-mpls"
    fibComplete.EntityData.SegmentPath = "fib-complete"
    fibComplete.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibComplete.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibComplete.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibComplete.EntityData.Children = types.NewOrderedMap()
    fibComplete.EntityData.Leafs = types.NewOrderedMap()
    fibComplete.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", fibComplete.StartTime})
    fibComplete.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", fibComplete.EndTime})
    fibComplete.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", fibComplete.Duration})

    fibComplete.EntityData.YListKeys = []string {}

    return &(fibComplete.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_LeafNetworksAdded
// List of Leaf Networks Added
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_LeafNetworksAdded struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is interface{} with range: 0..255.
    NetMask interface{}
}

func (leafNetworksAdded *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_LeafNetworksAdded) GetEntityData() *types.CommonEntityData {
    leafNetworksAdded.EntityData.YFilter = leafNetworksAdded.YFilter
    leafNetworksAdded.EntityData.YangName = "leaf-networks-added"
    leafNetworksAdded.EntityData.BundleName = "cisco_ios_xr"
    leafNetworksAdded.EntityData.ParentYangName = "priority"
    leafNetworksAdded.EntityData.SegmentPath = "leaf-networks-added"
    leafNetworksAdded.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leafNetworksAdded.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leafNetworksAdded.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leafNetworksAdded.EntityData.Children = types.NewOrderedMap()
    leafNetworksAdded.EntityData.Leafs = types.NewOrderedMap()
    leafNetworksAdded.EntityData.Leafs.Append("address", types.YLeaf{"Address", leafNetworksAdded.Address})
    leafNetworksAdded.EntityData.Leafs.Append("net-mask", types.YLeaf{"NetMask", leafNetworksAdded.NetMask})

    leafNetworksAdded.EntityData.YListKeys = []string {}

    return &(leafNetworksAdded.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_LeafNetworksDeleted
// List of Leaf Networks Deleted
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_LeafNetworksDeleted struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is interface{} with range: 0..255.
    NetMask interface{}
}

func (leafNetworksDeleted *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_Priority_LeafNetworksDeleted) GetEntityData() *types.CommonEntityData {
    leafNetworksDeleted.EntityData.YFilter = leafNetworksDeleted.YFilter
    leafNetworksDeleted.EntityData.YangName = "leaf-networks-deleted"
    leafNetworksDeleted.EntityData.BundleName = "cisco_ios_xr"
    leafNetworksDeleted.EntityData.ParentYangName = "priority"
    leafNetworksDeleted.EntityData.SegmentPath = "leaf-networks-deleted"
    leafNetworksDeleted.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leafNetworksDeleted.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leafNetworksDeleted.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leafNetworksDeleted.EntityData.Children = types.NewOrderedMap()
    leafNetworksDeleted.EntityData.Leafs = types.NewOrderedMap()
    leafNetworksDeleted.EntityData.Leafs.Append("address", types.YLeaf{"Address", leafNetworksDeleted.Address})
    leafNetworksDeleted.EntityData.Leafs.Append("net-mask", types.YLeaf{"NetMask", leafNetworksDeleted.NetMask})

    leafNetworksDeleted.EntityData.YListKeys = []string {}

    return &(leafNetworksDeleted.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_LspProcessed
// List of LSP changes processed
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_LspProcessed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP ID. The type is string.
    LspId interface{}

    // Sequence Number. The type is interface{} with range: 0..4294967295.
    SequenceNumber interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (lspProcessed *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_LspProcessed) GetEntityData() *types.CommonEntityData {
    lspProcessed.EntityData.YFilter = lspProcessed.YFilter
    lspProcessed.EntityData.YangName = "lsp-processed"
    lspProcessed.EntityData.BundleName = "cisco_ios_xr"
    lspProcessed.EntityData.ParentYangName = "spf-run-summary"
    lspProcessed.EntityData.SegmentPath = "lsp-processed"
    lspProcessed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspProcessed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspProcessed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspProcessed.EntityData.Children = types.NewOrderedMap()
    lspProcessed.EntityData.Leafs = types.NewOrderedMap()
    lspProcessed.EntityData.Leafs.Append("lsp-id", types.YLeaf{"LspId", lspProcessed.LspId})
    lspProcessed.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", lspProcessed.SequenceNumber})
    lspProcessed.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lspProcessed.ChangeType})
    lspProcessed.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", lspProcessed.ReceptionTime})

    lspProcessed.EntityData.YListKeys = []string {}

    return &(lspProcessed.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_LspRegenerated
// List of LSP regenerated
type Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_LspRegenerated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Serial Number of the session event. The type is interface{} with range:
    // 0..4294967295.
    SerialNumberXr interface{}

    // LSP ID. The type is string.
    LspId interface{}

    // Sequence Number. The type is interface{} with range: 0..4294967295.
    SequenceNumber interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}

    // ISIS Level. The type is RcmdIsisLvl.
    IsisLevel interface{}

    // SPF Run Number. The type is interface{} with range: 0..4294967295.
    SpfRunNumber interface{}

    // Trigger reasons for LSP regeneration. Example: pr^ - periodic, cr^ - clear
    // (Check the documentation for the entire list). The type is string.
    Reason interface{}
}

func (lspRegenerated *Rcmd_Isis_Instances_Instance_SpfRunSummaries_SpfRunSummary_LspRegenerated) GetEntityData() *types.CommonEntityData {
    lspRegenerated.EntityData.YFilter = lspRegenerated.YFilter
    lspRegenerated.EntityData.YangName = "lsp-regenerated"
    lspRegenerated.EntityData.BundleName = "cisco_ios_xr"
    lspRegenerated.EntityData.ParentYangName = "spf-run-summary"
    lspRegenerated.EntityData.SegmentPath = "lsp-regenerated"
    lspRegenerated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspRegenerated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspRegenerated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspRegenerated.EntityData.Children = types.NewOrderedMap()
    lspRegenerated.EntityData.Leafs = types.NewOrderedMap()
    lspRegenerated.EntityData.Leafs.Append("serial-number-xr", types.YLeaf{"SerialNumberXr", lspRegenerated.SerialNumberXr})
    lspRegenerated.EntityData.Leafs.Append("lsp-id", types.YLeaf{"LspId", lspRegenerated.LspId})
    lspRegenerated.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", lspRegenerated.SequenceNumber})
    lspRegenerated.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", lspRegenerated.ReceptionTime})
    lspRegenerated.EntityData.Leafs.Append("isis-level", types.YLeaf{"IsisLevel", lspRegenerated.IsisLevel})
    lspRegenerated.EntityData.Leafs.Append("spf-run-number", types.YLeaf{"SpfRunNumber", lspRegenerated.SpfRunNumber})
    lspRegenerated.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", lspRegenerated.Reason})

    lspRegenerated.EntityData.YListKeys = []string {}

    return &(lspRegenerated.EntityData)
}

// Rcmd_Isis_Instances_Instance_IpfrrEventOfflines
// ISIS IP-FRR Event offline data
type Rcmd_Isis_Instances_Instance_IpfrrEventOfflines struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Offline operational data for particular ISIS IP-FRR Event. The type is
    // slice of Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline.
    IpfrrEventOffline []*Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline
}

func (ipfrrEventOfflines *Rcmd_Isis_Instances_Instance_IpfrrEventOfflines) GetEntityData() *types.CommonEntityData {
    ipfrrEventOfflines.EntityData.YFilter = ipfrrEventOfflines.YFilter
    ipfrrEventOfflines.EntityData.YangName = "ipfrr-event-offlines"
    ipfrrEventOfflines.EntityData.BundleName = "cisco_ios_xr"
    ipfrrEventOfflines.EntityData.ParentYangName = "instance"
    ipfrrEventOfflines.EntityData.SegmentPath = "ipfrr-event-offlines"
    ipfrrEventOfflines.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipfrrEventOfflines.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipfrrEventOfflines.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipfrrEventOfflines.EntityData.Children = types.NewOrderedMap()
    ipfrrEventOfflines.EntityData.Children.Append("ipfrr-event-offline", types.YChild{"IpfrrEventOffline", nil})
    for i := range ipfrrEventOfflines.IpfrrEventOffline {
        ipfrrEventOfflines.EntityData.Children.Append(types.GetSegmentPath(ipfrrEventOfflines.IpfrrEventOffline[i]), types.YChild{"IpfrrEventOffline", ipfrrEventOfflines.IpfrrEventOffline[i]})
    }
    ipfrrEventOfflines.EntityData.Leafs = types.NewOrderedMap()

    ipfrrEventOfflines.EntityData.YListKeys = []string {}

    return &(ipfrrEventOfflines.EntityData)
}

// Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline
// Offline operational data for particular ISIS
// IP-FRR Event
type Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific IP-FRR Event. The type is interface{}
    // with range: 1..4294967295.
    EventId interface{}

    // IP-Frr Event ID. The type is interface{} with range: 0..4294967295.
    EventIdXr interface{}

    // Trigger time  (eg: Apr 24 13:16:04.961). The type is string.
    TriggerTime interface{}

    // IP-Frr Triggered reference SPF Run Number. The type is interface{} with
    // range: 0..4294967295.
    TriggerSpfRun interface{}

    // Waiting Time (in milliseconds). The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    WaitTime interface{}

    // Start Time offset from trigger time (in milliseconds). The type is string.
    // Units are millisecond.
    StartTimeOffset interface{}

    // Duration for the calculation (in milliseconds). The type is string. Units
    // are millisecond.
    Duration interface{}

    // IP-Frr Completed reference SPF Run Number. The type is interface{} with
    // range: 0..4294967295.
    CompletedSpfRun interface{}

    // Cumulative Number of Routes for all priorities. The type is interface{}
    // with range: 0..4294967295.
    TotalRoutes interface{}

    // Cumulative Number of Fully Protected Routes. The type is interface{} with
    // range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Cumulative Number of Partially Protected Routes. The type is interface{}
    // with range: 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage for all priorities. The type is string. Units are
    // percentage.
    Coverage interface{}

    // IP-Frr Statistics categorized by priority. The type is slice of
    // Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_IpfrrStatistic.
    IpfrrStatistic []*Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_IpfrrStatistic

    // Remote Node Information. The type is slice of
    // Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode.
    RemoteNode []*Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode
}

func (ipfrrEventOffline *Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline) GetEntityData() *types.CommonEntityData {
    ipfrrEventOffline.EntityData.YFilter = ipfrrEventOffline.YFilter
    ipfrrEventOffline.EntityData.YangName = "ipfrr-event-offline"
    ipfrrEventOffline.EntityData.BundleName = "cisco_ios_xr"
    ipfrrEventOffline.EntityData.ParentYangName = "ipfrr-event-offlines"
    ipfrrEventOffline.EntityData.SegmentPath = "ipfrr-event-offline" + types.AddKeyToken(ipfrrEventOffline.EventId, "event-id")
    ipfrrEventOffline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipfrrEventOffline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipfrrEventOffline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipfrrEventOffline.EntityData.Children = types.NewOrderedMap()
    ipfrrEventOffline.EntityData.Children.Append("ipfrr-statistic", types.YChild{"IpfrrStatistic", nil})
    for i := range ipfrrEventOffline.IpfrrStatistic {
        ipfrrEventOffline.EntityData.Children.Append(types.GetSegmentPath(ipfrrEventOffline.IpfrrStatistic[i]), types.YChild{"IpfrrStatistic", ipfrrEventOffline.IpfrrStatistic[i]})
    }
    ipfrrEventOffline.EntityData.Children.Append("remote-node", types.YChild{"RemoteNode", nil})
    for i := range ipfrrEventOffline.RemoteNode {
        ipfrrEventOffline.EntityData.Children.Append(types.GetSegmentPath(ipfrrEventOffline.RemoteNode[i]), types.YChild{"RemoteNode", ipfrrEventOffline.RemoteNode[i]})
    }
    ipfrrEventOffline.EntityData.Leafs = types.NewOrderedMap()
    ipfrrEventOffline.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", ipfrrEventOffline.EventId})
    ipfrrEventOffline.EntityData.Leafs.Append("event-id-xr", types.YLeaf{"EventIdXr", ipfrrEventOffline.EventIdXr})
    ipfrrEventOffline.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", ipfrrEventOffline.TriggerTime})
    ipfrrEventOffline.EntityData.Leafs.Append("trigger-spf-run", types.YLeaf{"TriggerSpfRun", ipfrrEventOffline.TriggerSpfRun})
    ipfrrEventOffline.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", ipfrrEventOffline.WaitTime})
    ipfrrEventOffline.EntityData.Leafs.Append("start-time-offset", types.YLeaf{"StartTimeOffset", ipfrrEventOffline.StartTimeOffset})
    ipfrrEventOffline.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ipfrrEventOffline.Duration})
    ipfrrEventOffline.EntityData.Leafs.Append("completed-spf-run", types.YLeaf{"CompletedSpfRun", ipfrrEventOffline.CompletedSpfRun})
    ipfrrEventOffline.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", ipfrrEventOffline.TotalRoutes})
    ipfrrEventOffline.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", ipfrrEventOffline.FullyProtectedRoutes})
    ipfrrEventOffline.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", ipfrrEventOffline.PartiallyProtectedRoutes})
    ipfrrEventOffline.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", ipfrrEventOffline.Coverage})

    ipfrrEventOffline.EntityData.YListKeys = []string {"EventId"}

    return &(ipfrrEventOffline.EntityData)
}

// Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_IpfrrStatistic
// IP-Frr Statistics categorized by priority
type Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_IpfrrStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Priority. The type is RcmdPriorityLevel.
    Priority interface{}

    // Total Number of Routes. The type is interface{} with range: 0..4294967295.
    TotalRoutes interface{}

    // Fully Protected Routes. The type is interface{} with range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Partially Protected Routes. The type is interface{} with range:
    // 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage. The type is string. Units are percentage.
    Coverage interface{}

    // Local LFA Coverage in percentage. The type is string. Units are percentage.
    LocalLfaCoverage interface{}

    // Remote LFA Coverage in percentage. The type is string. Units are
    // percentage.
    RemoteLfaCoverage interface{}

    // Covearge is below Configured Threshold. The type is bool.
    BelowThreshold interface{}
}

func (ipfrrStatistic *Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_IpfrrStatistic) GetEntityData() *types.CommonEntityData {
    ipfrrStatistic.EntityData.YFilter = ipfrrStatistic.YFilter
    ipfrrStatistic.EntityData.YangName = "ipfrr-statistic"
    ipfrrStatistic.EntityData.BundleName = "cisco_ios_xr"
    ipfrrStatistic.EntityData.ParentYangName = "ipfrr-event-offline"
    ipfrrStatistic.EntityData.SegmentPath = "ipfrr-statistic"
    ipfrrStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipfrrStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipfrrStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipfrrStatistic.EntityData.Children = types.NewOrderedMap()
    ipfrrStatistic.EntityData.Leafs = types.NewOrderedMap()
    ipfrrStatistic.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", ipfrrStatistic.Priority})
    ipfrrStatistic.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", ipfrrStatistic.TotalRoutes})
    ipfrrStatistic.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", ipfrrStatistic.FullyProtectedRoutes})
    ipfrrStatistic.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", ipfrrStatistic.PartiallyProtectedRoutes})
    ipfrrStatistic.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", ipfrrStatistic.Coverage})
    ipfrrStatistic.EntityData.Leafs.Append("local-lfa-coverage", types.YLeaf{"LocalLfaCoverage", ipfrrStatistic.LocalLfaCoverage})
    ipfrrStatistic.EntityData.Leafs.Append("remote-lfa-coverage", types.YLeaf{"RemoteLfaCoverage", ipfrrStatistic.RemoteLfaCoverage})
    ipfrrStatistic.EntityData.Leafs.Append("below-threshold", types.YLeaf{"BelowThreshold", ipfrrStatistic.BelowThreshold})

    ipfrrStatistic.EntityData.YListKeys = []string {}

    return &(ipfrrStatistic.EntityData)
}

// Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode
// Remote Node Information
type Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote-LFA Node ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteNodeId interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Number of paths protected by this Remote Node. The type is interface{} with
    // range: 0..4294967295.
    PathCount interface{}

    // Inuse time of the Remote Node (eg: Apr 24 13:16 :04.961). The type is
    // string.
    InUseTime interface{}

    // Protected Primary Paths. The type is slice of
    // Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode_PrimaryPath.
    PrimaryPath []*Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode_PrimaryPath
}

func (remoteNode *Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode) GetEntityData() *types.CommonEntityData {
    remoteNode.EntityData.YFilter = remoteNode.YFilter
    remoteNode.EntityData.YangName = "remote-node"
    remoteNode.EntityData.BundleName = "cisco_ios_xr"
    remoteNode.EntityData.ParentYangName = "ipfrr-event-offline"
    remoteNode.EntityData.SegmentPath = "remote-node"
    remoteNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNode.EntityData.Children = types.NewOrderedMap()
    remoteNode.EntityData.Children.Append("primary-path", types.YChild{"PrimaryPath", nil})
    for i := range remoteNode.PrimaryPath {
        remoteNode.EntityData.Children.Append(types.GetSegmentPath(remoteNode.PrimaryPath[i]), types.YChild{"PrimaryPath", remoteNode.PrimaryPath[i]})
    }
    remoteNode.EntityData.Leafs = types.NewOrderedMap()
    remoteNode.EntityData.Leafs.Append("remote-node-id", types.YLeaf{"RemoteNodeId", remoteNode.RemoteNodeId})
    remoteNode.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", remoteNode.InterfaceName})
    remoteNode.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", remoteNode.NeighbourAddress})
    remoteNode.EntityData.Leafs.Append("path-count", types.YLeaf{"PathCount", remoteNode.PathCount})
    remoteNode.EntityData.Leafs.Append("in-use-time", types.YLeaf{"InUseTime", remoteNode.InUseTime})

    remoteNode.EntityData.YListKeys = []string {}

    return &(remoteNode.EntityData)
}

// Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode_PrimaryPath
// Protected Primary Paths
type Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode_PrimaryPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}
}

func (primaryPath *Rcmd_Isis_Instances_Instance_IpfrrEventOfflines_IpfrrEventOffline_RemoteNode_PrimaryPath) GetEntityData() *types.CommonEntityData {
    primaryPath.EntityData.YFilter = primaryPath.YFilter
    primaryPath.EntityData.YangName = "primary-path"
    primaryPath.EntityData.BundleName = "cisco_ios_xr"
    primaryPath.EntityData.ParentYangName = "remote-node"
    primaryPath.EntityData.SegmentPath = "primary-path"
    primaryPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryPath.EntityData.Children = types.NewOrderedMap()
    primaryPath.EntityData.Leafs = types.NewOrderedMap()
    primaryPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", primaryPath.InterfaceName})
    primaryPath.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", primaryPath.NeighbourAddress})

    primaryPath.EntityData.YListKeys = []string {}

    return &(primaryPath.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines
// ISIS SPF run offline data
type Rcmd_Isis_Instances_Instance_SpfRunOfflines struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Offline operational data for particular ISIS SPF run. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline.
    SpfRunOffline []*Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline
}

func (spfRunOfflines *Rcmd_Isis_Instances_Instance_SpfRunOfflines) GetEntityData() *types.CommonEntityData {
    spfRunOfflines.EntityData.YFilter = spfRunOfflines.YFilter
    spfRunOfflines.EntityData.YangName = "spf-run-offlines"
    spfRunOfflines.EntityData.BundleName = "cisco_ios_xr"
    spfRunOfflines.EntityData.ParentYangName = "instance"
    spfRunOfflines.EntityData.SegmentPath = "spf-run-offlines"
    spfRunOfflines.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfRunOfflines.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfRunOfflines.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfRunOfflines.EntityData.Children = types.NewOrderedMap()
    spfRunOfflines.EntityData.Children.Append("spf-run-offline", types.YChild{"SpfRunOffline", nil})
    for i := range spfRunOfflines.SpfRunOffline {
        spfRunOfflines.EntityData.Children.Append(types.GetSegmentPath(spfRunOfflines.SpfRunOffline[i]), types.YChild{"SpfRunOffline", spfRunOfflines.SpfRunOffline[i]})
    }
    spfRunOfflines.EntityData.Leafs = types.NewOrderedMap()

    spfRunOfflines.EntityData.YListKeys = []string {}

    return &(spfRunOfflines.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline
// Offline operational data for particular ISIS
// SPF run
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific SPF run. The type is interface{} with
    // range: 1..4294967295.
    SpfRunNumber interface{}

    // Start time (offset from event trigger time in ss .msec). The type is
    // string.
    StartTime interface{}

    // Wait time applied at SPF schedule (in msec). The type is interface{} with
    // range: 0..4294967295.
    WaitTime interface{}

    // Trigger reasons for SPF run. Example: pr^ - periodic, cr^ - clear (Check
    // the documentation for the entire list). The type is string.
    Reason interface{}

    // SPF summary information.
    SpfSummary Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary

    // SPF Node statistics.
    NodeStatistics Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_NodeStatistics

    // Trigger LSP. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_TriggerLsp.
    TriggerLsp []*Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_TriggerLsp

    // Convergence information on per-priority basis. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority.
    Priority []*Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority

    // List of LSP changes processed. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_LspProcessed.
    LspProcessed []*Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_LspProcessed

    // List of LSP regenerated. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_LspRegenerated.
    LspRegenerated []*Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_LspRegenerated
}

func (spfRunOffline *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline) GetEntityData() *types.CommonEntityData {
    spfRunOffline.EntityData.YFilter = spfRunOffline.YFilter
    spfRunOffline.EntityData.YangName = "spf-run-offline"
    spfRunOffline.EntityData.BundleName = "cisco_ios_xr"
    spfRunOffline.EntityData.ParentYangName = "spf-run-offlines"
    spfRunOffline.EntityData.SegmentPath = "spf-run-offline" + types.AddKeyToken(spfRunOffline.SpfRunNumber, "spf-run-number")
    spfRunOffline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfRunOffline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfRunOffline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfRunOffline.EntityData.Children = types.NewOrderedMap()
    spfRunOffline.EntityData.Children.Append("spf-summary", types.YChild{"SpfSummary", &spfRunOffline.SpfSummary})
    spfRunOffline.EntityData.Children.Append("node-statistics", types.YChild{"NodeStatistics", &spfRunOffline.NodeStatistics})
    spfRunOffline.EntityData.Children.Append("trigger-lsp", types.YChild{"TriggerLsp", nil})
    for i := range spfRunOffline.TriggerLsp {
        spfRunOffline.EntityData.Children.Append(types.GetSegmentPath(spfRunOffline.TriggerLsp[i]), types.YChild{"TriggerLsp", spfRunOffline.TriggerLsp[i]})
    }
    spfRunOffline.EntityData.Children.Append("priority", types.YChild{"Priority", nil})
    for i := range spfRunOffline.Priority {
        spfRunOffline.EntityData.Children.Append(types.GetSegmentPath(spfRunOffline.Priority[i]), types.YChild{"Priority", spfRunOffline.Priority[i]})
    }
    spfRunOffline.EntityData.Children.Append("lsp-processed", types.YChild{"LspProcessed", nil})
    for i := range spfRunOffline.LspProcessed {
        spfRunOffline.EntityData.Children.Append(types.GetSegmentPath(spfRunOffline.LspProcessed[i]), types.YChild{"LspProcessed", spfRunOffline.LspProcessed[i]})
    }
    spfRunOffline.EntityData.Children.Append("lsp-regenerated", types.YChild{"LspRegenerated", nil})
    for i := range spfRunOffline.LspRegenerated {
        spfRunOffline.EntityData.Children.Append(types.GetSegmentPath(spfRunOffline.LspRegenerated[i]), types.YChild{"LspRegenerated", spfRunOffline.LspRegenerated[i]})
    }
    spfRunOffline.EntityData.Leafs = types.NewOrderedMap()
    spfRunOffline.EntityData.Leafs.Append("spf-run-number", types.YLeaf{"SpfRunNumber", spfRunOffline.SpfRunNumber})
    spfRunOffline.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", spfRunOffline.StartTime})
    spfRunOffline.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", spfRunOffline.WaitTime})
    spfRunOffline.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", spfRunOffline.Reason})

    spfRunOffline.EntityData.YListKeys = []string {"SpfRunNumber"}

    return &(spfRunOffline.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary
// SPF summary information
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Topology index (multi-topology). The type is interface{} with range:
    // 0..4294967295.
    Topology interface{}

    // ISIS Level. The type is RcmdIsisLvl.
    IsisLevel interface{}

    // Type of SPF. The type is RcmdIsisSpf.
    Type interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // SPF state. The type is RcmdSpfState.
    State interface{}

    // Whether the event has all information. The type is bool.
    IsDataComplete interface{}

    // Trigger time (in hh:mm:ss.msec). The type is string.
    TriggerTime interface{}

    // Duration of SPF calculation (in ss.msec). The type is string.
    Duration interface{}

    // Total number of LSP changes processed. The type is interface{} with range:
    // 0..65535.
    TotalLspChanges interface{}

    // Convergence information summary on per-priority basis. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary.
    PrioritySummary []*Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary
}

func (spfSummary *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary) GetEntityData() *types.CommonEntityData {
    spfSummary.EntityData.YFilter = spfSummary.YFilter
    spfSummary.EntityData.YangName = "spf-summary"
    spfSummary.EntityData.BundleName = "cisco_ios_xr"
    spfSummary.EntityData.ParentYangName = "spf-run-offline"
    spfSummary.EntityData.SegmentPath = "spf-summary"
    spfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfSummary.EntityData.Children = types.NewOrderedMap()
    spfSummary.EntityData.Children.Append("priority-summary", types.YChild{"PrioritySummary", nil})
    for i := range spfSummary.PrioritySummary {
        spfSummary.EntityData.Children.Append(types.GetSegmentPath(spfSummary.PrioritySummary[i]), types.YChild{"PrioritySummary", spfSummary.PrioritySummary[i]})
    }
    spfSummary.EntityData.Leafs = types.NewOrderedMap()
    spfSummary.EntityData.Leafs.Append("topology", types.YLeaf{"Topology", spfSummary.Topology})
    spfSummary.EntityData.Leafs.Append("isis-level", types.YLeaf{"IsisLevel", spfSummary.IsisLevel})
    spfSummary.EntityData.Leafs.Append("type", types.YLeaf{"Type", spfSummary.Type})
    spfSummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", spfSummary.ThresholdExceeded})
    spfSummary.EntityData.Leafs.Append("state", types.YLeaf{"State", spfSummary.State})
    spfSummary.EntityData.Leafs.Append("is-data-complete", types.YLeaf{"IsDataComplete", spfSummary.IsDataComplete})
    spfSummary.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", spfSummary.TriggerTime})
    spfSummary.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", spfSummary.Duration})
    spfSummary.EntityData.Leafs.Append("total-lsp-changes", types.YLeaf{"TotalLspChanges", spfSummary.TotalLspChanges})

    spfSummary.EntityData.YListKeys = []string {}

    return &(spfSummary.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary
// Convergence information summary on per-priority
// basis
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Critical, High, Medium or Low. The type is RcmdPriorityLevel.
    Level interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Route statistics.
    RouteStatistics Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_RouteStatistics

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_MplsConvergenceTime

    // Fast Re-Route Statistics. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_FrrStatistic.
    FrrStatistic []*Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_FrrStatistic
}

func (prioritySummary *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary) GetEntityData() *types.CommonEntityData {
    prioritySummary.EntityData.YFilter = prioritySummary.YFilter
    prioritySummary.EntityData.YangName = "priority-summary"
    prioritySummary.EntityData.BundleName = "cisco_ios_xr"
    prioritySummary.EntityData.ParentYangName = "spf-summary"
    prioritySummary.EntityData.SegmentPath = "priority-summary"
    prioritySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prioritySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prioritySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prioritySummary.EntityData.Children = types.NewOrderedMap()
    prioritySummary.EntityData.Children.Append("route-statistics", types.YChild{"RouteStatistics", &prioritySummary.RouteStatistics})
    prioritySummary.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prioritySummary.IpConvergenceTime})
    prioritySummary.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prioritySummary.MplsConvergenceTime})
    prioritySummary.EntityData.Children.Append("frr-statistic", types.YChild{"FrrStatistic", nil})
    for i := range prioritySummary.FrrStatistic {
        prioritySummary.EntityData.Children.Append(types.GetSegmentPath(prioritySummary.FrrStatistic[i]), types.YChild{"FrrStatistic", prioritySummary.FrrStatistic[i]})
    }
    prioritySummary.EntityData.Leafs = types.NewOrderedMap()
    prioritySummary.EntityData.Leafs.Append("level", types.YLeaf{"Level", prioritySummary.Level})
    prioritySummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prioritySummary.ThresholdExceeded})

    prioritySummary.EntityData.YListKeys = []string {}

    return &(prioritySummary.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_RouteStatistics
// Route statistics
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_RouteStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Added. The type is interface{} with range: 0..4294967295.
    Adds interface{}

    // Deleted. The type is interface{} with range: 0..4294967295.
    Deletes interface{}

    // Modified. The type is interface{} with range: 0..4294967295.
    Modifies interface{}

    // Reachable. The type is interface{} with range: 0..4294967295.
    Reachables interface{}

    // Unreachable. The type is interface{} with range: 0..4294967295.
    Unreachables interface{}

    // Touched. The type is interface{} with range: 0..4294967295.
    Touches interface{}
}

func (routeStatistics *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_RouteStatistics) GetEntityData() *types.CommonEntityData {
    routeStatistics.EntityData.YFilter = routeStatistics.YFilter
    routeStatistics.EntityData.YangName = "route-statistics"
    routeStatistics.EntityData.BundleName = "cisco_ios_xr"
    routeStatistics.EntityData.ParentYangName = "priority-summary"
    routeStatistics.EntityData.SegmentPath = "route-statistics"
    routeStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeStatistics.EntityData.Children = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs.Append("adds", types.YLeaf{"Adds", routeStatistics.Adds})
    routeStatistics.EntityData.Leafs.Append("deletes", types.YLeaf{"Deletes", routeStatistics.Deletes})
    routeStatistics.EntityData.Leafs.Append("modifies", types.YLeaf{"Modifies", routeStatistics.Modifies})
    routeStatistics.EntityData.Leafs.Append("reachables", types.YLeaf{"Reachables", routeStatistics.Reachables})
    routeStatistics.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", routeStatistics.Unreachables})
    routeStatistics.EntityData.Leafs.Append("touches", types.YLeaf{"Touches", routeStatistics.Touches})

    routeStatistics.EntityData.YListKeys = []string {}

    return &(routeStatistics.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "priority-summary"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "priority-summary"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_FrrStatistic
// Fast Re-Route Statistics
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_FrrStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total Number of Routes. The type is interface{} with range: 0..4294967295.
    TotalRoutes interface{}

    // Fully Protected Routes. The type is interface{} with range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Partially Protected Routes. The type is interface{} with range:
    // 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage. The type is string. Units are percentage.
    Coverage interface{}
}

func (frrStatistic *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_SpfSummary_PrioritySummary_FrrStatistic) GetEntityData() *types.CommonEntityData {
    frrStatistic.EntityData.YFilter = frrStatistic.YFilter
    frrStatistic.EntityData.YangName = "frr-statistic"
    frrStatistic.EntityData.BundleName = "cisco_ios_xr"
    frrStatistic.EntityData.ParentYangName = "priority-summary"
    frrStatistic.EntityData.SegmentPath = "frr-statistic"
    frrStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrStatistic.EntityData.Children = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", frrStatistic.TotalRoutes})
    frrStatistic.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", frrStatistic.FullyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", frrStatistic.PartiallyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", frrStatistic.Coverage})

    frrStatistic.EntityData.YListKeys = []string {}

    return &(frrStatistic.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_NodeStatistics
// SPF Node statistics
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_NodeStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Added. The type is interface{} with range: 0..4294967295.
    Adds interface{}

    // Deleted. The type is interface{} with range: 0..4294967295.
    Deletes interface{}

    // Modified. The type is interface{} with range: 0..4294967295.
    Modifies interface{}

    // Reachable. The type is interface{} with range: 0..4294967295.
    Reachables interface{}

    // Unreachable. The type is interface{} with range: 0..4294967295.
    Unreachables interface{}

    // Touched. The type is interface{} with range: 0..4294967295.
    Touches interface{}
}

func (nodeStatistics *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_NodeStatistics) GetEntityData() *types.CommonEntityData {
    nodeStatistics.EntityData.YFilter = nodeStatistics.YFilter
    nodeStatistics.EntityData.YangName = "node-statistics"
    nodeStatistics.EntityData.BundleName = "cisco_ios_xr"
    nodeStatistics.EntityData.ParentYangName = "spf-run-offline"
    nodeStatistics.EntityData.SegmentPath = "node-statistics"
    nodeStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeStatistics.EntityData.Children = types.NewOrderedMap()
    nodeStatistics.EntityData.Leafs = types.NewOrderedMap()
    nodeStatistics.EntityData.Leafs.Append("adds", types.YLeaf{"Adds", nodeStatistics.Adds})
    nodeStatistics.EntityData.Leafs.Append("deletes", types.YLeaf{"Deletes", nodeStatistics.Deletes})
    nodeStatistics.EntityData.Leafs.Append("modifies", types.YLeaf{"Modifies", nodeStatistics.Modifies})
    nodeStatistics.EntityData.Leafs.Append("reachables", types.YLeaf{"Reachables", nodeStatistics.Reachables})
    nodeStatistics.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", nodeStatistics.Unreachables})
    nodeStatistics.EntityData.Leafs.Append("touches", types.YLeaf{"Touches", nodeStatistics.Touches})

    nodeStatistics.EntityData.YListKeys = []string {}

    return &(nodeStatistics.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_TriggerLsp
// Trigger LSP
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_TriggerLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP ID. The type is string.
    LspId interface{}

    // Sequence Number. The type is interface{} with range: 0..4294967295.
    SequenceNumber interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (triggerLsp *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_TriggerLsp) GetEntityData() *types.CommonEntityData {
    triggerLsp.EntityData.YFilter = triggerLsp.YFilter
    triggerLsp.EntityData.YangName = "trigger-lsp"
    triggerLsp.EntityData.BundleName = "cisco_ios_xr"
    triggerLsp.EntityData.ParentYangName = "spf-run-offline"
    triggerLsp.EntityData.SegmentPath = "trigger-lsp"
    triggerLsp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    triggerLsp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    triggerLsp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    triggerLsp.EntityData.Children = types.NewOrderedMap()
    triggerLsp.EntityData.Leafs = types.NewOrderedMap()
    triggerLsp.EntityData.Leafs.Append("lsp-id", types.YLeaf{"LspId", triggerLsp.LspId})
    triggerLsp.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", triggerLsp.SequenceNumber})
    triggerLsp.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", triggerLsp.ChangeType})
    triggerLsp.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", triggerLsp.ReceptionTime})

    triggerLsp.EntityData.YListKeys = []string {}

    return &(triggerLsp.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority
// Convergence information on per-priority basis
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of the priority.
    PrioritySummary Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary

    // Convergence timeline details. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline.
    ConvergenceTimeline []*Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline

    // List of Leaf Networks Added. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_LeafNetworksAdded.
    LeafNetworksAdded []*Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_LeafNetworksAdded

    // List of Leaf Networks Deleted. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_LeafNetworksDeleted.
    LeafNetworksDeleted []*Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_LeafNetworksDeleted
}

func (priority *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "spf-run-offline"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Children.Append("priority-summary", types.YChild{"PrioritySummary", &priority.PrioritySummary})
    priority.EntityData.Children.Append("convergence-timeline", types.YChild{"ConvergenceTimeline", nil})
    for i := range priority.ConvergenceTimeline {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.ConvergenceTimeline[i]), types.YChild{"ConvergenceTimeline", priority.ConvergenceTimeline[i]})
    }
    priority.EntityData.Children.Append("leaf-networks-added", types.YChild{"LeafNetworksAdded", nil})
    for i := range priority.LeafNetworksAdded {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.LeafNetworksAdded[i]), types.YChild{"LeafNetworksAdded", priority.LeafNetworksAdded[i]})
    }
    priority.EntityData.Children.Append("leaf-networks-deleted", types.YChild{"LeafNetworksDeleted", nil})
    for i := range priority.LeafNetworksDeleted {
        priority.EntityData.Children.Append(types.GetSegmentPath(priority.LeafNetworksDeleted[i]), types.YChild{"LeafNetworksDeleted", priority.LeafNetworksDeleted[i]})
    }
    priority.EntityData.Leafs = types.NewOrderedMap()

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary
// Summary of the priority
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Critical, High, Medium or Low. The type is RcmdPriorityLevel.
    Level interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Route statistics.
    RouteStatistics Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_RouteStatistics

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_MplsConvergenceTime

    // Fast Re-Route Statistics. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_FrrStatistic.
    FrrStatistic []*Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_FrrStatistic
}

func (prioritySummary *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary) GetEntityData() *types.CommonEntityData {
    prioritySummary.EntityData.YFilter = prioritySummary.YFilter
    prioritySummary.EntityData.YangName = "priority-summary"
    prioritySummary.EntityData.BundleName = "cisco_ios_xr"
    prioritySummary.EntityData.ParentYangName = "priority"
    prioritySummary.EntityData.SegmentPath = "priority-summary"
    prioritySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prioritySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prioritySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prioritySummary.EntityData.Children = types.NewOrderedMap()
    prioritySummary.EntityData.Children.Append("route-statistics", types.YChild{"RouteStatistics", &prioritySummary.RouteStatistics})
    prioritySummary.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prioritySummary.IpConvergenceTime})
    prioritySummary.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prioritySummary.MplsConvergenceTime})
    prioritySummary.EntityData.Children.Append("frr-statistic", types.YChild{"FrrStatistic", nil})
    for i := range prioritySummary.FrrStatistic {
        prioritySummary.EntityData.Children.Append(types.GetSegmentPath(prioritySummary.FrrStatistic[i]), types.YChild{"FrrStatistic", prioritySummary.FrrStatistic[i]})
    }
    prioritySummary.EntityData.Leafs = types.NewOrderedMap()
    prioritySummary.EntityData.Leafs.Append("level", types.YLeaf{"Level", prioritySummary.Level})
    prioritySummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prioritySummary.ThresholdExceeded})

    prioritySummary.EntityData.YListKeys = []string {}

    return &(prioritySummary.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_RouteStatistics
// Route statistics
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_RouteStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Added. The type is interface{} with range: 0..4294967295.
    Adds interface{}

    // Deleted. The type is interface{} with range: 0..4294967295.
    Deletes interface{}

    // Modified. The type is interface{} with range: 0..4294967295.
    Modifies interface{}

    // Reachable. The type is interface{} with range: 0..4294967295.
    Reachables interface{}

    // Unreachable. The type is interface{} with range: 0..4294967295.
    Unreachables interface{}

    // Touched. The type is interface{} with range: 0..4294967295.
    Touches interface{}
}

func (routeStatistics *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_RouteStatistics) GetEntityData() *types.CommonEntityData {
    routeStatistics.EntityData.YFilter = routeStatistics.YFilter
    routeStatistics.EntityData.YangName = "route-statistics"
    routeStatistics.EntityData.BundleName = "cisco_ios_xr"
    routeStatistics.EntityData.ParentYangName = "priority-summary"
    routeStatistics.EntityData.SegmentPath = "route-statistics"
    routeStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeStatistics.EntityData.Children = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs = types.NewOrderedMap()
    routeStatistics.EntityData.Leafs.Append("adds", types.YLeaf{"Adds", routeStatistics.Adds})
    routeStatistics.EntityData.Leafs.Append("deletes", types.YLeaf{"Deletes", routeStatistics.Deletes})
    routeStatistics.EntityData.Leafs.Append("modifies", types.YLeaf{"Modifies", routeStatistics.Modifies})
    routeStatistics.EntityData.Leafs.Append("reachables", types.YLeaf{"Reachables", routeStatistics.Reachables})
    routeStatistics.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", routeStatistics.Unreachables})
    routeStatistics.EntityData.Leafs.Append("touches", types.YLeaf{"Touches", routeStatistics.Touches})

    routeStatistics.EntityData.YListKeys = []string {}

    return &(routeStatistics.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "priority-summary"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "priority-summary"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_FrrStatistic
// Fast Re-Route Statistics
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_FrrStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total Number of Routes. The type is interface{} with range: 0..4294967295.
    TotalRoutes interface{}

    // Fully Protected Routes. The type is interface{} with range: 0..4294967295.
    FullyProtectedRoutes interface{}

    // Partially Protected Routes. The type is interface{} with range:
    // 0..4294967295.
    PartiallyProtectedRoutes interface{}

    // Coverage in percentage. The type is string. Units are percentage.
    Coverage interface{}
}

func (frrStatistic *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_PrioritySummary_FrrStatistic) GetEntityData() *types.CommonEntityData {
    frrStatistic.EntityData.YFilter = frrStatistic.YFilter
    frrStatistic.EntityData.YangName = "frr-statistic"
    frrStatistic.EntityData.BundleName = "cisco_ios_xr"
    frrStatistic.EntityData.ParentYangName = "priority-summary"
    frrStatistic.EntityData.SegmentPath = "frr-statistic"
    frrStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrStatistic.EntityData.Children = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs = types.NewOrderedMap()
    frrStatistic.EntityData.Leafs.Append("total-routes", types.YLeaf{"TotalRoutes", frrStatistic.TotalRoutes})
    frrStatistic.EntityData.Leafs.Append("fully-protected-routes", types.YLeaf{"FullyProtectedRoutes", frrStatistic.FullyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("partially-protected-routes", types.YLeaf{"PartiallyProtectedRoutes", frrStatistic.PartiallyProtectedRoutes})
    frrStatistic.EntityData.Leafs.Append("coverage", types.YLeaf{"Coverage", frrStatistic.Coverage})

    frrStatistic.EntityData.YListKeys = []string {}

    return &(frrStatistic.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline
// Convergence timeline details
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route origin (routing protocol).
    RouteOrigin Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RouteOrigin

    // Entry point of IPv4 RIB.
    RiBv4Enter Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RiBv4Enter

    // Exit point from IPv4 RIB to FIBs.
    RiBv4Exit Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RiBv4Exit

    // Route Redistribute point from IPv4 RIB to LDP.
    RiBv4Redistribute Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RiBv4Redistribute

    // Entry point of LDP.
    LdpEnter Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LdpEnter

    // Exit point of LDP to LSD.
    LdpExit Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LdpExit

    // Entry point of LSD.
    LsdEnter Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LsdEnter

    // Exit point of LSD to FIBs.
    LsdExit Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LsdExit

    // List of Linecards' completion point for IP routes. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcIp.
    LcIp []*Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcIp

    // List of Linecards' completion point for MPLS labels. The type is slice of
    // Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcMpls.
    LcMpls []*Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcMpls
}

func (convergenceTimeline *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline) GetEntityData() *types.CommonEntityData {
    convergenceTimeline.EntityData.YFilter = convergenceTimeline.YFilter
    convergenceTimeline.EntityData.YangName = "convergence-timeline"
    convergenceTimeline.EntityData.BundleName = "cisco_ios_xr"
    convergenceTimeline.EntityData.ParentYangName = "priority"
    convergenceTimeline.EntityData.SegmentPath = "convergence-timeline"
    convergenceTimeline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    convergenceTimeline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    convergenceTimeline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    convergenceTimeline.EntityData.Children = types.NewOrderedMap()
    convergenceTimeline.EntityData.Children.Append("route-origin", types.YChild{"RouteOrigin", &convergenceTimeline.RouteOrigin})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-enter", types.YChild{"RiBv4Enter", &convergenceTimeline.RiBv4Enter})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-exit", types.YChild{"RiBv4Exit", &convergenceTimeline.RiBv4Exit})
    convergenceTimeline.EntityData.Children.Append("ri-bv4-redistribute", types.YChild{"RiBv4Redistribute", &convergenceTimeline.RiBv4Redistribute})
    convergenceTimeline.EntityData.Children.Append("ldp-enter", types.YChild{"LdpEnter", &convergenceTimeline.LdpEnter})
    convergenceTimeline.EntityData.Children.Append("ldp-exit", types.YChild{"LdpExit", &convergenceTimeline.LdpExit})
    convergenceTimeline.EntityData.Children.Append("lsd-enter", types.YChild{"LsdEnter", &convergenceTimeline.LsdEnter})
    convergenceTimeline.EntityData.Children.Append("lsd-exit", types.YChild{"LsdExit", &convergenceTimeline.LsdExit})
    convergenceTimeline.EntityData.Children.Append("lc-ip", types.YChild{"LcIp", nil})
    for i := range convergenceTimeline.LcIp {
        convergenceTimeline.EntityData.Children.Append(types.GetSegmentPath(convergenceTimeline.LcIp[i]), types.YChild{"LcIp", convergenceTimeline.LcIp[i]})
    }
    convergenceTimeline.EntityData.Children.Append("lc-mpls", types.YChild{"LcMpls", nil})
    for i := range convergenceTimeline.LcMpls {
        convergenceTimeline.EntityData.Children.Append(types.GetSegmentPath(convergenceTimeline.LcMpls[i]), types.YChild{"LcMpls", convergenceTimeline.LcMpls[i]})
    }
    convergenceTimeline.EntityData.Leafs = types.NewOrderedMap()

    convergenceTimeline.EntityData.YListKeys = []string {}

    return &(convergenceTimeline.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RouteOrigin
// Route origin (routing protocol)
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RouteOrigin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (routeOrigin *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RouteOrigin) GetEntityData() *types.CommonEntityData {
    routeOrigin.EntityData.YFilter = routeOrigin.YFilter
    routeOrigin.EntityData.YangName = "route-origin"
    routeOrigin.EntityData.BundleName = "cisco_ios_xr"
    routeOrigin.EntityData.ParentYangName = "convergence-timeline"
    routeOrigin.EntityData.SegmentPath = "route-origin"
    routeOrigin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeOrigin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeOrigin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeOrigin.EntityData.Children = types.NewOrderedMap()
    routeOrigin.EntityData.Leafs = types.NewOrderedMap()
    routeOrigin.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", routeOrigin.StartTime})
    routeOrigin.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", routeOrigin.EndTime})
    routeOrigin.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", routeOrigin.Duration})

    routeOrigin.EntityData.YListKeys = []string {}

    return &(routeOrigin.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RiBv4Enter
// Entry point of IPv4 RIB
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RiBv4Enter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Enter *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RiBv4Enter) GetEntityData() *types.CommonEntityData {
    riBv4Enter.EntityData.YFilter = riBv4Enter.YFilter
    riBv4Enter.EntityData.YangName = "ri-bv4-enter"
    riBv4Enter.EntityData.BundleName = "cisco_ios_xr"
    riBv4Enter.EntityData.ParentYangName = "convergence-timeline"
    riBv4Enter.EntityData.SegmentPath = "ri-bv4-enter"
    riBv4Enter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Enter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Enter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Enter.EntityData.Children = types.NewOrderedMap()
    riBv4Enter.EntityData.Leafs = types.NewOrderedMap()
    riBv4Enter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Enter.StartTime})
    riBv4Enter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Enter.EndTime})
    riBv4Enter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Enter.Duration})

    riBv4Enter.EntityData.YListKeys = []string {}

    return &(riBv4Enter.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RiBv4Exit
// Exit point from IPv4 RIB to FIBs
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RiBv4Exit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Exit *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RiBv4Exit) GetEntityData() *types.CommonEntityData {
    riBv4Exit.EntityData.YFilter = riBv4Exit.YFilter
    riBv4Exit.EntityData.YangName = "ri-bv4-exit"
    riBv4Exit.EntityData.BundleName = "cisco_ios_xr"
    riBv4Exit.EntityData.ParentYangName = "convergence-timeline"
    riBv4Exit.EntityData.SegmentPath = "ri-bv4-exit"
    riBv4Exit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Exit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Exit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Exit.EntityData.Children = types.NewOrderedMap()
    riBv4Exit.EntityData.Leafs = types.NewOrderedMap()
    riBv4Exit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Exit.StartTime})
    riBv4Exit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Exit.EndTime})
    riBv4Exit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Exit.Duration})

    riBv4Exit.EntityData.YListKeys = []string {}

    return &(riBv4Exit.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RiBv4Redistribute
// Route Redistribute point from IPv4 RIB to LDP
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RiBv4Redistribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (riBv4Redistribute *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_RiBv4Redistribute) GetEntityData() *types.CommonEntityData {
    riBv4Redistribute.EntityData.YFilter = riBv4Redistribute.YFilter
    riBv4Redistribute.EntityData.YangName = "ri-bv4-redistribute"
    riBv4Redistribute.EntityData.BundleName = "cisco_ios_xr"
    riBv4Redistribute.EntityData.ParentYangName = "convergence-timeline"
    riBv4Redistribute.EntityData.SegmentPath = "ri-bv4-redistribute"
    riBv4Redistribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    riBv4Redistribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    riBv4Redistribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    riBv4Redistribute.EntityData.Children = types.NewOrderedMap()
    riBv4Redistribute.EntityData.Leafs = types.NewOrderedMap()
    riBv4Redistribute.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", riBv4Redistribute.StartTime})
    riBv4Redistribute.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", riBv4Redistribute.EndTime})
    riBv4Redistribute.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", riBv4Redistribute.Duration})

    riBv4Redistribute.EntityData.YListKeys = []string {}

    return &(riBv4Redistribute.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LdpEnter
// Entry point of LDP
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LdpEnter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (ldpEnter *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LdpEnter) GetEntityData() *types.CommonEntityData {
    ldpEnter.EntityData.YFilter = ldpEnter.YFilter
    ldpEnter.EntityData.YangName = "ldp-enter"
    ldpEnter.EntityData.BundleName = "cisco_ios_xr"
    ldpEnter.EntityData.ParentYangName = "convergence-timeline"
    ldpEnter.EntityData.SegmentPath = "ldp-enter"
    ldpEnter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpEnter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpEnter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpEnter.EntityData.Children = types.NewOrderedMap()
    ldpEnter.EntityData.Leafs = types.NewOrderedMap()
    ldpEnter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", ldpEnter.StartTime})
    ldpEnter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", ldpEnter.EndTime})
    ldpEnter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ldpEnter.Duration})

    ldpEnter.EntityData.YListKeys = []string {}

    return &(ldpEnter.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LdpExit
// Exit point of LDP to LSD
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LdpExit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (ldpExit *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LdpExit) GetEntityData() *types.CommonEntityData {
    ldpExit.EntityData.YFilter = ldpExit.YFilter
    ldpExit.EntityData.YangName = "ldp-exit"
    ldpExit.EntityData.BundleName = "cisco_ios_xr"
    ldpExit.EntityData.ParentYangName = "convergence-timeline"
    ldpExit.EntityData.SegmentPath = "ldp-exit"
    ldpExit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpExit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpExit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpExit.EntityData.Children = types.NewOrderedMap()
    ldpExit.EntityData.Leafs = types.NewOrderedMap()
    ldpExit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", ldpExit.StartTime})
    ldpExit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", ldpExit.EndTime})
    ldpExit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", ldpExit.Duration})

    ldpExit.EntityData.YListKeys = []string {}

    return &(ldpExit.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LsdEnter
// Entry point of LSD
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LsdEnter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (lsdEnter *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LsdEnter) GetEntityData() *types.CommonEntityData {
    lsdEnter.EntityData.YFilter = lsdEnter.YFilter
    lsdEnter.EntityData.YangName = "lsd-enter"
    lsdEnter.EntityData.BundleName = "cisco_ios_xr"
    lsdEnter.EntityData.ParentYangName = "convergence-timeline"
    lsdEnter.EntityData.SegmentPath = "lsd-enter"
    lsdEnter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsdEnter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsdEnter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsdEnter.EntityData.Children = types.NewOrderedMap()
    lsdEnter.EntityData.Leafs = types.NewOrderedMap()
    lsdEnter.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", lsdEnter.StartTime})
    lsdEnter.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", lsdEnter.EndTime})
    lsdEnter.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", lsdEnter.Duration})

    lsdEnter.EntityData.YListKeys = []string {}

    return &(lsdEnter.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LsdExit
// Exit point of LSD to FIBs
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LsdExit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (lsdExit *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LsdExit) GetEntityData() *types.CommonEntityData {
    lsdExit.EntityData.YFilter = lsdExit.YFilter
    lsdExit.EntityData.YangName = "lsd-exit"
    lsdExit.EntityData.BundleName = "cisco_ios_xr"
    lsdExit.EntityData.ParentYangName = "convergence-timeline"
    lsdExit.EntityData.SegmentPath = "lsd-exit"
    lsdExit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsdExit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsdExit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsdExit.EntityData.Children = types.NewOrderedMap()
    lsdExit.EntityData.Leafs = types.NewOrderedMap()
    lsdExit.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", lsdExit.StartTime})
    lsdExit.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", lsdExit.EndTime})
    lsdExit.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", lsdExit.Duration})

    lsdExit.EntityData.YListKeys = []string {}

    return &(lsdExit.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcIp
// List of Linecards' completion point for IP
// routes
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB.
    FibComplete Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcIp_FibComplete
}

func (lcIp *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcIp) GetEntityData() *types.CommonEntityData {
    lcIp.EntityData.YFilter = lcIp.YFilter
    lcIp.EntityData.YangName = "lc-ip"
    lcIp.EntityData.BundleName = "cisco_ios_xr"
    lcIp.EntityData.ParentYangName = "convergence-timeline"
    lcIp.EntityData.SegmentPath = "lc-ip"
    lcIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcIp.EntityData.Children = types.NewOrderedMap()
    lcIp.EntityData.Children.Append("fib-complete", types.YChild{"FibComplete", &lcIp.FibComplete})
    lcIp.EntityData.Leafs = types.NewOrderedMap()
    lcIp.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcIp.NodeName})
    lcIp.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcIp.Speed})

    lcIp.EntityData.YListKeys = []string {}

    return &(lcIp.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcIp_FibComplete
// Completion point of FIB
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcIp_FibComplete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (fibComplete *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcIp_FibComplete) GetEntityData() *types.CommonEntityData {
    fibComplete.EntityData.YFilter = fibComplete.YFilter
    fibComplete.EntityData.YangName = "fib-complete"
    fibComplete.EntityData.BundleName = "cisco_ios_xr"
    fibComplete.EntityData.ParentYangName = "lc-ip"
    fibComplete.EntityData.SegmentPath = "fib-complete"
    fibComplete.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibComplete.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibComplete.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibComplete.EntityData.Children = types.NewOrderedMap()
    fibComplete.EntityData.Leafs = types.NewOrderedMap()
    fibComplete.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", fibComplete.StartTime})
    fibComplete.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", fibComplete.EndTime})
    fibComplete.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", fibComplete.Duration})

    fibComplete.EntityData.YListKeys = []string {}

    return &(fibComplete.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcMpls
// List of Linecards' completion point for MPLS
// labels
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB.
    FibComplete Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcMpls_FibComplete
}

func (lcMpls *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcMpls) GetEntityData() *types.CommonEntityData {
    lcMpls.EntityData.YFilter = lcMpls.YFilter
    lcMpls.EntityData.YangName = "lc-mpls"
    lcMpls.EntityData.BundleName = "cisco_ios_xr"
    lcMpls.EntityData.ParentYangName = "convergence-timeline"
    lcMpls.EntityData.SegmentPath = "lc-mpls"
    lcMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcMpls.EntityData.Children = types.NewOrderedMap()
    lcMpls.EntityData.Children.Append("fib-complete", types.YChild{"FibComplete", &lcMpls.FibComplete})
    lcMpls.EntityData.Leafs = types.NewOrderedMap()
    lcMpls.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcMpls.NodeName})
    lcMpls.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcMpls.Speed})

    lcMpls.EntityData.YListKeys = []string {}

    return &(lcMpls.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcMpls_FibComplete
// Completion point of FIB
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcMpls_FibComplete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First route process time relative to event trigger time (in ss.msec). The
    // type is string.
    StartTime interface{}

    // Last route process time relative to event trigger time (in ss.msec). The
    // type is string.
    EndTime interface{}

    // Duration of processing (in ss.msec). The type is string.
    Duration interface{}
}

func (fibComplete *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_ConvergenceTimeline_LcMpls_FibComplete) GetEntityData() *types.CommonEntityData {
    fibComplete.EntityData.YFilter = fibComplete.YFilter
    fibComplete.EntityData.YangName = "fib-complete"
    fibComplete.EntityData.BundleName = "cisco_ios_xr"
    fibComplete.EntityData.ParentYangName = "lc-mpls"
    fibComplete.EntityData.SegmentPath = "fib-complete"
    fibComplete.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibComplete.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibComplete.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibComplete.EntityData.Children = types.NewOrderedMap()
    fibComplete.EntityData.Leafs = types.NewOrderedMap()
    fibComplete.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", fibComplete.StartTime})
    fibComplete.EntityData.Leafs.Append("end-time", types.YLeaf{"EndTime", fibComplete.EndTime})
    fibComplete.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", fibComplete.Duration})

    fibComplete.EntityData.YListKeys = []string {}

    return &(fibComplete.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_LeafNetworksAdded
// List of Leaf Networks Added
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_LeafNetworksAdded struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is interface{} with range: 0..255.
    NetMask interface{}
}

func (leafNetworksAdded *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_LeafNetworksAdded) GetEntityData() *types.CommonEntityData {
    leafNetworksAdded.EntityData.YFilter = leafNetworksAdded.YFilter
    leafNetworksAdded.EntityData.YangName = "leaf-networks-added"
    leafNetworksAdded.EntityData.BundleName = "cisco_ios_xr"
    leafNetworksAdded.EntityData.ParentYangName = "priority"
    leafNetworksAdded.EntityData.SegmentPath = "leaf-networks-added"
    leafNetworksAdded.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leafNetworksAdded.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leafNetworksAdded.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leafNetworksAdded.EntityData.Children = types.NewOrderedMap()
    leafNetworksAdded.EntityData.Leafs = types.NewOrderedMap()
    leafNetworksAdded.EntityData.Leafs.Append("address", types.YLeaf{"Address", leafNetworksAdded.Address})
    leafNetworksAdded.EntityData.Leafs.Append("net-mask", types.YLeaf{"NetMask", leafNetworksAdded.NetMask})

    leafNetworksAdded.EntityData.YListKeys = []string {}

    return &(leafNetworksAdded.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_LeafNetworksDeleted
// List of Leaf Networks Deleted
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_LeafNetworksDeleted struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is interface{} with range: 0..255.
    NetMask interface{}
}

func (leafNetworksDeleted *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_Priority_LeafNetworksDeleted) GetEntityData() *types.CommonEntityData {
    leafNetworksDeleted.EntityData.YFilter = leafNetworksDeleted.YFilter
    leafNetworksDeleted.EntityData.YangName = "leaf-networks-deleted"
    leafNetworksDeleted.EntityData.BundleName = "cisco_ios_xr"
    leafNetworksDeleted.EntityData.ParentYangName = "priority"
    leafNetworksDeleted.EntityData.SegmentPath = "leaf-networks-deleted"
    leafNetworksDeleted.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leafNetworksDeleted.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leafNetworksDeleted.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leafNetworksDeleted.EntityData.Children = types.NewOrderedMap()
    leafNetworksDeleted.EntityData.Leafs = types.NewOrderedMap()
    leafNetworksDeleted.EntityData.Leafs.Append("address", types.YLeaf{"Address", leafNetworksDeleted.Address})
    leafNetworksDeleted.EntityData.Leafs.Append("net-mask", types.YLeaf{"NetMask", leafNetworksDeleted.NetMask})

    leafNetworksDeleted.EntityData.YListKeys = []string {}

    return &(leafNetworksDeleted.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_LspProcessed
// List of LSP changes processed
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_LspProcessed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP ID. The type is string.
    LspId interface{}

    // Sequence Number. The type is interface{} with range: 0..4294967295.
    SequenceNumber interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (lspProcessed *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_LspProcessed) GetEntityData() *types.CommonEntityData {
    lspProcessed.EntityData.YFilter = lspProcessed.YFilter
    lspProcessed.EntityData.YangName = "lsp-processed"
    lspProcessed.EntityData.BundleName = "cisco_ios_xr"
    lspProcessed.EntityData.ParentYangName = "spf-run-offline"
    lspProcessed.EntityData.SegmentPath = "lsp-processed"
    lspProcessed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspProcessed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspProcessed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspProcessed.EntityData.Children = types.NewOrderedMap()
    lspProcessed.EntityData.Leafs = types.NewOrderedMap()
    lspProcessed.EntityData.Leafs.Append("lsp-id", types.YLeaf{"LspId", lspProcessed.LspId})
    lspProcessed.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", lspProcessed.SequenceNumber})
    lspProcessed.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lspProcessed.ChangeType})
    lspProcessed.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", lspProcessed.ReceptionTime})

    lspProcessed.EntityData.YListKeys = []string {}

    return &(lspProcessed.EntityData)
}

// Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_LspRegenerated
// List of LSP regenerated
type Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_LspRegenerated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Serial Number of the session event. The type is interface{} with range:
    // 0..4294967295.
    SerialNumberXr interface{}

    // LSP ID. The type is string.
    LspId interface{}

    // Sequence Number. The type is interface{} with range: 0..4294967295.
    SequenceNumber interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}

    // ISIS Level. The type is RcmdIsisLvl.
    IsisLevel interface{}

    // SPF Run Number. The type is interface{} with range: 0..4294967295.
    SpfRunNumber interface{}

    // Trigger reasons for LSP regeneration. Example: pr^ - periodic, cr^ - clear
    // (Check the documentation for the entire list). The type is string.
    Reason interface{}
}

func (lspRegenerated *Rcmd_Isis_Instances_Instance_SpfRunOfflines_SpfRunOffline_LspRegenerated) GetEntityData() *types.CommonEntityData {
    lspRegenerated.EntityData.YFilter = lspRegenerated.YFilter
    lspRegenerated.EntityData.YangName = "lsp-regenerated"
    lspRegenerated.EntityData.BundleName = "cisco_ios_xr"
    lspRegenerated.EntityData.ParentYangName = "spf-run-offline"
    lspRegenerated.EntityData.SegmentPath = "lsp-regenerated"
    lspRegenerated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspRegenerated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspRegenerated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspRegenerated.EntityData.Children = types.NewOrderedMap()
    lspRegenerated.EntityData.Leafs = types.NewOrderedMap()
    lspRegenerated.EntityData.Leafs.Append("serial-number-xr", types.YLeaf{"SerialNumberXr", lspRegenerated.SerialNumberXr})
    lspRegenerated.EntityData.Leafs.Append("lsp-id", types.YLeaf{"LspId", lspRegenerated.LspId})
    lspRegenerated.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", lspRegenerated.SequenceNumber})
    lspRegenerated.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", lspRegenerated.ReceptionTime})
    lspRegenerated.EntityData.Leafs.Append("isis-level", types.YLeaf{"IsisLevel", lspRegenerated.IsisLevel})
    lspRegenerated.EntityData.Leafs.Append("spf-run-number", types.YLeaf{"SpfRunNumber", lspRegenerated.SpfRunNumber})
    lspRegenerated.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", lspRegenerated.Reason})

    lspRegenerated.EntityData.YListKeys = []string {}

    return &(lspRegenerated.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventSummaries
// ISIS Prefix events summary data
type Rcmd_Isis_Instances_Instance_PrefixEventSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix Event data. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary.
    PrefixEventSummary []*Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary
}

func (prefixEventSummaries *Rcmd_Isis_Instances_Instance_PrefixEventSummaries) GetEntityData() *types.CommonEntityData {
    prefixEventSummaries.EntityData.YFilter = prefixEventSummaries.YFilter
    prefixEventSummaries.EntityData.YangName = "prefix-event-summaries"
    prefixEventSummaries.EntityData.BundleName = "cisco_ios_xr"
    prefixEventSummaries.EntityData.ParentYangName = "instance"
    prefixEventSummaries.EntityData.SegmentPath = "prefix-event-summaries"
    prefixEventSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixEventSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixEventSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixEventSummaries.EntityData.Children = types.NewOrderedMap()
    prefixEventSummaries.EntityData.Children.Append("prefix-event-summary", types.YChild{"PrefixEventSummary", nil})
    for i := range prefixEventSummaries.PrefixEventSummary {
        prefixEventSummaries.EntityData.Children.Append(types.GetSegmentPath(prefixEventSummaries.PrefixEventSummary[i]), types.YChild{"PrefixEventSummary", prefixEventSummaries.PrefixEventSummary[i]})
    }
    prefixEventSummaries.EntityData.Leafs = types.NewOrderedMap()

    prefixEventSummaries.EntityData.YListKeys = []string {}

    return &(prefixEventSummaries.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary
// Prefix Event data
type Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Event ID. The type is interface{} with
    // range: 1..4294967295.
    EventId interface{}

    // Prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLenth interface{}

    // Referenced SPF Run No (0 - Not Applicable). The type is interface{} with
    // range: 0..4294967295.
    SpfRunNo interface{}

    // Referenced IP-FRR Event ID (0 - Not Applicable). The type is interface{}
    // with range: 0..4294967295.
    IpfrrEventId interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Event processed priority. The type is RcmdPriorityLevel.
    Priority interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Route Type intra/inter/l1/l2. The type is RcmdShowRoute.
    RouteType interface{}

    // Route Path Change Type. The type is RcmdShowRoutePathChange.
    RoutePathChangeType interface{}

    // Protocol route cost. The type is interface{} with range: 0..4294967295.
    Cost interface{}

    // Event trigger time. The type is string.
    TriggerTime interface{}

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_MplsConvergenceTime

    // Path information. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path.
    Path []*Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path

    // LSA that triggered this event. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TriggerLsa.
    TriggerLsa []*Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TriggerLsa

    // Timeline information. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine.
    TimeLine []*Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine

    // List of LSAs processed. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_LsaProcessed.
    LsaProcessed []*Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_LsaProcessed
}

func (prefixEventSummary *Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary) GetEntityData() *types.CommonEntityData {
    prefixEventSummary.EntityData.YFilter = prefixEventSummary.YFilter
    prefixEventSummary.EntityData.YangName = "prefix-event-summary"
    prefixEventSummary.EntityData.BundleName = "cisco_ios_xr"
    prefixEventSummary.EntityData.ParentYangName = "prefix-event-summaries"
    prefixEventSummary.EntityData.SegmentPath = "prefix-event-summary" + types.AddKeyToken(prefixEventSummary.EventId, "event-id")
    prefixEventSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixEventSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixEventSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixEventSummary.EntityData.Children = types.NewOrderedMap()
    prefixEventSummary.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prefixEventSummary.IpConvergenceTime})
    prefixEventSummary.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prefixEventSummary.MplsConvergenceTime})
    prefixEventSummary.EntityData.Children.Append("path", types.YChild{"Path", nil})
    for i := range prefixEventSummary.Path {
        prefixEventSummary.EntityData.Children.Append(types.GetSegmentPath(prefixEventSummary.Path[i]), types.YChild{"Path", prefixEventSummary.Path[i]})
    }
    prefixEventSummary.EntityData.Children.Append("trigger-lsa", types.YChild{"TriggerLsa", nil})
    for i := range prefixEventSummary.TriggerLsa {
        prefixEventSummary.EntityData.Children.Append(types.GetSegmentPath(prefixEventSummary.TriggerLsa[i]), types.YChild{"TriggerLsa", prefixEventSummary.TriggerLsa[i]})
    }
    prefixEventSummary.EntityData.Children.Append("time-line", types.YChild{"TimeLine", nil})
    for i := range prefixEventSummary.TimeLine {
        prefixEventSummary.EntityData.Children.Append(types.GetSegmentPath(prefixEventSummary.TimeLine[i]), types.YChild{"TimeLine", prefixEventSummary.TimeLine[i]})
    }
    prefixEventSummary.EntityData.Children.Append("lsa-processed", types.YChild{"LsaProcessed", nil})
    for i := range prefixEventSummary.LsaProcessed {
        prefixEventSummary.EntityData.Children.Append(types.GetSegmentPath(prefixEventSummary.LsaProcessed[i]), types.YChild{"LsaProcessed", prefixEventSummary.LsaProcessed[i]})
    }
    prefixEventSummary.EntityData.Leafs = types.NewOrderedMap()
    prefixEventSummary.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", prefixEventSummary.EventId})
    prefixEventSummary.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", prefixEventSummary.Prefix})
    prefixEventSummary.EntityData.Leafs.Append("prefix-lenth", types.YLeaf{"PrefixLenth", prefixEventSummary.PrefixLenth})
    prefixEventSummary.EntityData.Leafs.Append("spf-run-no", types.YLeaf{"SpfRunNo", prefixEventSummary.SpfRunNo})
    prefixEventSummary.EntityData.Leafs.Append("ipfrr-event-id", types.YLeaf{"IpfrrEventId", prefixEventSummary.IpfrrEventId})
    prefixEventSummary.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prefixEventSummary.ThresholdExceeded})
    prefixEventSummary.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", prefixEventSummary.Priority})
    prefixEventSummary.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", prefixEventSummary.ChangeType})
    prefixEventSummary.EntityData.Leafs.Append("route-type", types.YLeaf{"RouteType", prefixEventSummary.RouteType})
    prefixEventSummary.EntityData.Leafs.Append("route-path-change-type", types.YLeaf{"RoutePathChangeType", prefixEventSummary.RoutePathChangeType})
    prefixEventSummary.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", prefixEventSummary.Cost})
    prefixEventSummary.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", prefixEventSummary.TriggerTime})

    prefixEventSummary.EntityData.YListKeys = []string {"EventId"}

    return &(prefixEventSummary.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "prefix-event-summary"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "prefix-event-summary"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path
// Path information
type Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Path Metric. The type is interface{} with range: 0..4294967295.
    PathMetric interface{}

    // Backup Path Informatoin. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path_LfaPath.
    LfaPath []*Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path_LfaPath
}

func (path *Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "prefix-event-summary"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Children.Append("lfa-path", types.YChild{"LfaPath", nil})
    for i := range path.LfaPath {
        path.EntityData.Children.Append(types.GetSegmentPath(path.LfaPath[i]), types.YChild{"LfaPath", path.LfaPath[i]})
    }
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", path.InterfaceName})
    path.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", path.NeighbourAddress})
    path.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", path.ChangeType})
    path.EntityData.Leafs.Append("path-metric", types.YLeaf{"PathMetric", path.PathMetric})

    path.EntityData.YListKeys = []string {}

    return &(path.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path_LfaPath
// Backup Path Informatoin
type Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path_LfaPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of LFA. The type is RcmdShowIpfrrLfa.
    LfaType interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Path Metric. The type is interface{} with range: 0..4294967295.
    PathMetric interface{}

    // Remote Node ID, in case of Remote LFA. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteNodeId interface{}
}

func (lfaPath *Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_Path_LfaPath) GetEntityData() *types.CommonEntityData {
    lfaPath.EntityData.YFilter = lfaPath.YFilter
    lfaPath.EntityData.YangName = "lfa-path"
    lfaPath.EntityData.BundleName = "cisco_ios_xr"
    lfaPath.EntityData.ParentYangName = "path"
    lfaPath.EntityData.SegmentPath = "lfa-path"
    lfaPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lfaPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lfaPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lfaPath.EntityData.Children = types.NewOrderedMap()
    lfaPath.EntityData.Leafs = types.NewOrderedMap()
    lfaPath.EntityData.Leafs.Append("lfa-type", types.YLeaf{"LfaType", lfaPath.LfaType})
    lfaPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", lfaPath.InterfaceName})
    lfaPath.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", lfaPath.NeighbourAddress})
    lfaPath.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lfaPath.ChangeType})
    lfaPath.EntityData.Leafs.Append("path-metric", types.YLeaf{"PathMetric", lfaPath.PathMetric})
    lfaPath.EntityData.Leafs.Append("remote-node-id", types.YLeaf{"RemoteNodeId", lfaPath.RemoteNodeId})

    lfaPath.EntityData.YListKeys = []string {}

    return &(lfaPath.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TriggerLsa
// LSA that triggered this event
type Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TriggerLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (triggerLsa *Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TriggerLsa) GetEntityData() *types.CommonEntityData {
    triggerLsa.EntityData.YFilter = triggerLsa.YFilter
    triggerLsa.EntityData.YangName = "trigger-lsa"
    triggerLsa.EntityData.BundleName = "cisco_ios_xr"
    triggerLsa.EntityData.ParentYangName = "prefix-event-summary"
    triggerLsa.EntityData.SegmentPath = "trigger-lsa"
    triggerLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    triggerLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    triggerLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    triggerLsa.EntityData.Children = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", triggerLsa.LsaId})
    triggerLsa.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", triggerLsa.SequenceNumber})
    triggerLsa.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", triggerLsa.LsaType})
    triggerLsa.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", triggerLsa.OriginRouterId})
    triggerLsa.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", triggerLsa.ChangeType})
    triggerLsa.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", triggerLsa.ReceptionTime})

    triggerLsa.EntityData.YListKeys = []string {}

    return &(triggerLsa.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine
// Timeline information
type Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route origin (routing protocol). The type is string.
    RouteOrigin interface{}

    // Entry point of IPv4 RIB. The type is string.
    RiBv4Enter interface{}

    // Exit point from IPv4 RIB to FIBs. The type is string.
    RiBv4Exit interface{}

    // Route Redistribute point from IPv4 RIB to LDP. The type is string.
    RiBv4Redistribute interface{}

    // Entry point of LDP. The type is string.
    LdpEnter interface{}

    // Exit point of LDP to LSD. The type is string.
    LdpExit interface{}

    // Entry point of LSD. The type is string.
    LsdEnter interface{}

    // Exit point of LSD to FIBs. The type is string.
    LsdExit interface{}

    // List of Linecards' completion point for IP routes. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcIp.
    LcIp []*Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcIp

    // List of Linecards' completion point for MPLS labels. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcMpls.
    LcMpls []*Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcMpls
}

func (timeLine *Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine) GetEntityData() *types.CommonEntityData {
    timeLine.EntityData.YFilter = timeLine.YFilter
    timeLine.EntityData.YangName = "time-line"
    timeLine.EntityData.BundleName = "cisco_ios_xr"
    timeLine.EntityData.ParentYangName = "prefix-event-summary"
    timeLine.EntityData.SegmentPath = "time-line"
    timeLine.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeLine.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeLine.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeLine.EntityData.Children = types.NewOrderedMap()
    timeLine.EntityData.Children.Append("lc-ip", types.YChild{"LcIp", nil})
    for i := range timeLine.LcIp {
        timeLine.EntityData.Children.Append(types.GetSegmentPath(timeLine.LcIp[i]), types.YChild{"LcIp", timeLine.LcIp[i]})
    }
    timeLine.EntityData.Children.Append("lc-mpls", types.YChild{"LcMpls", nil})
    for i := range timeLine.LcMpls {
        timeLine.EntityData.Children.Append(types.GetSegmentPath(timeLine.LcMpls[i]), types.YChild{"LcMpls", timeLine.LcMpls[i]})
    }
    timeLine.EntityData.Leafs = types.NewOrderedMap()
    timeLine.EntityData.Leafs.Append("route-origin", types.YLeaf{"RouteOrigin", timeLine.RouteOrigin})
    timeLine.EntityData.Leafs.Append("ri-bv4-enter", types.YLeaf{"RiBv4Enter", timeLine.RiBv4Enter})
    timeLine.EntityData.Leafs.Append("ri-bv4-exit", types.YLeaf{"RiBv4Exit", timeLine.RiBv4Exit})
    timeLine.EntityData.Leafs.Append("ri-bv4-redistribute", types.YLeaf{"RiBv4Redistribute", timeLine.RiBv4Redistribute})
    timeLine.EntityData.Leafs.Append("ldp-enter", types.YLeaf{"LdpEnter", timeLine.LdpEnter})
    timeLine.EntityData.Leafs.Append("ldp-exit", types.YLeaf{"LdpExit", timeLine.LdpExit})
    timeLine.EntityData.Leafs.Append("lsd-enter", types.YLeaf{"LsdEnter", timeLine.LsdEnter})
    timeLine.EntityData.Leafs.Append("lsd-exit", types.YLeaf{"LsdExit", timeLine.LsdExit})

    timeLine.EntityData.YListKeys = []string {}

    return &(timeLine.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcIp
// List of Linecards' completion point for IP
// routes
type Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB. The type is string.
    FibComplete interface{}
}

func (lcIp *Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcIp) GetEntityData() *types.CommonEntityData {
    lcIp.EntityData.YFilter = lcIp.YFilter
    lcIp.EntityData.YangName = "lc-ip"
    lcIp.EntityData.BundleName = "cisco_ios_xr"
    lcIp.EntityData.ParentYangName = "time-line"
    lcIp.EntityData.SegmentPath = "lc-ip"
    lcIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcIp.EntityData.Children = types.NewOrderedMap()
    lcIp.EntityData.Leafs = types.NewOrderedMap()
    lcIp.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcIp.NodeName})
    lcIp.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcIp.Speed})
    lcIp.EntityData.Leafs.Append("fib-complete", types.YLeaf{"FibComplete", lcIp.FibComplete})

    lcIp.EntityData.YListKeys = []string {}

    return &(lcIp.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcMpls
// List of Linecards' completion point for MPLS
// labels
type Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB. The type is string.
    FibComplete interface{}
}

func (lcMpls *Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_TimeLine_LcMpls) GetEntityData() *types.CommonEntityData {
    lcMpls.EntityData.YFilter = lcMpls.YFilter
    lcMpls.EntityData.YangName = "lc-mpls"
    lcMpls.EntityData.BundleName = "cisco_ios_xr"
    lcMpls.EntityData.ParentYangName = "time-line"
    lcMpls.EntityData.SegmentPath = "lc-mpls"
    lcMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcMpls.EntityData.Children = types.NewOrderedMap()
    lcMpls.EntityData.Leafs = types.NewOrderedMap()
    lcMpls.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcMpls.NodeName})
    lcMpls.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcMpls.Speed})
    lcMpls.EntityData.Leafs.Append("fib-complete", types.YLeaf{"FibComplete", lcMpls.FibComplete})

    lcMpls.EntityData.YListKeys = []string {}

    return &(lcMpls.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_LsaProcessed
// List of LSAs processed
type Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_LsaProcessed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (lsaProcessed *Rcmd_Isis_Instances_Instance_PrefixEventSummaries_PrefixEventSummary_LsaProcessed) GetEntityData() *types.CommonEntityData {
    lsaProcessed.EntityData.YFilter = lsaProcessed.YFilter
    lsaProcessed.EntityData.YangName = "lsa-processed"
    lsaProcessed.EntityData.BundleName = "cisco_ios_xr"
    lsaProcessed.EntityData.ParentYangName = "prefix-event-summary"
    lsaProcessed.EntityData.SegmentPath = "lsa-processed"
    lsaProcessed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsaProcessed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsaProcessed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsaProcessed.EntityData.Children = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", lsaProcessed.LsaId})
    lsaProcessed.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", lsaProcessed.SequenceNumber})
    lsaProcessed.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", lsaProcessed.LsaType})
    lsaProcessed.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", lsaProcessed.OriginRouterId})
    lsaProcessed.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lsaProcessed.ChangeType})
    lsaProcessed.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", lsaProcessed.ReceptionTime})

    lsaProcessed.EntityData.YListKeys = []string {}

    return &(lsaProcessed.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventOfflines
// ISIS Prefix events offline data
type Rcmd_Isis_Instances_Instance_PrefixEventOfflines struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Offline operational data for particular ISIS Prefix Event. The type is
    // slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline.
    PrefixEventOffline []*Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline
}

func (prefixEventOfflines *Rcmd_Isis_Instances_Instance_PrefixEventOfflines) GetEntityData() *types.CommonEntityData {
    prefixEventOfflines.EntityData.YFilter = prefixEventOfflines.YFilter
    prefixEventOfflines.EntityData.YangName = "prefix-event-offlines"
    prefixEventOfflines.EntityData.BundleName = "cisco_ios_xr"
    prefixEventOfflines.EntityData.ParentYangName = "instance"
    prefixEventOfflines.EntityData.SegmentPath = "prefix-event-offlines"
    prefixEventOfflines.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixEventOfflines.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixEventOfflines.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixEventOfflines.EntityData.Children = types.NewOrderedMap()
    prefixEventOfflines.EntityData.Children.Append("prefix-event-offline", types.YChild{"PrefixEventOffline", nil})
    for i := range prefixEventOfflines.PrefixEventOffline {
        prefixEventOfflines.EntityData.Children.Append(types.GetSegmentPath(prefixEventOfflines.PrefixEventOffline[i]), types.YChild{"PrefixEventOffline", prefixEventOfflines.PrefixEventOffline[i]})
    }
    prefixEventOfflines.EntityData.Leafs = types.NewOrderedMap()

    prefixEventOfflines.EntityData.YListKeys = []string {}

    return &(prefixEventOfflines.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline
// Offline operational data for particular ISIS
// Prefix Event
type Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Event ID. The type is interface{} with
    // range: 1..4294967295.
    EventId interface{}

    // Prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLenth interface{}

    // Referenced SPF Run No (0 - Not Applicable). The type is interface{} with
    // range: 0..4294967295.
    SpfRunNo interface{}

    // Referenced IP-FRR Event ID (0 - Not Applicable). The type is interface{}
    // with range: 0..4294967295.
    IpfrrEventId interface{}

    // Threshold exceeded. The type is bool.
    ThresholdExceeded interface{}

    // Event processed priority. The type is RcmdPriorityLevel.
    Priority interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Route Type intra/inter/l1/l2. The type is RcmdShowRoute.
    RouteType interface{}

    // Route Path Change Type. The type is RcmdShowRoutePathChange.
    RoutePathChangeType interface{}

    // Protocol route cost. The type is interface{} with range: 0..4294967295.
    Cost interface{}

    // Event trigger time. The type is string.
    TriggerTime interface{}

    // Convergence time for IP route programming.
    IpConvergenceTime Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_IpConvergenceTime

    // Convergence time for MPLS label programming.
    MplsConvergenceTime Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_MplsConvergenceTime

    // Path information. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path.
    Path []*Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path

    // LSA that triggered this event. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TriggerLsa.
    TriggerLsa []*Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TriggerLsa

    // Timeline information. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine.
    TimeLine []*Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine

    // List of LSAs processed. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_LsaProcessed.
    LsaProcessed []*Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_LsaProcessed
}

func (prefixEventOffline *Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline) GetEntityData() *types.CommonEntityData {
    prefixEventOffline.EntityData.YFilter = prefixEventOffline.YFilter
    prefixEventOffline.EntityData.YangName = "prefix-event-offline"
    prefixEventOffline.EntityData.BundleName = "cisco_ios_xr"
    prefixEventOffline.EntityData.ParentYangName = "prefix-event-offlines"
    prefixEventOffline.EntityData.SegmentPath = "prefix-event-offline" + types.AddKeyToken(prefixEventOffline.EventId, "event-id")
    prefixEventOffline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixEventOffline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixEventOffline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixEventOffline.EntityData.Children = types.NewOrderedMap()
    prefixEventOffline.EntityData.Children.Append("ip-convergence-time", types.YChild{"IpConvergenceTime", &prefixEventOffline.IpConvergenceTime})
    prefixEventOffline.EntityData.Children.Append("mpls-convergence-time", types.YChild{"MplsConvergenceTime", &prefixEventOffline.MplsConvergenceTime})
    prefixEventOffline.EntityData.Children.Append("path", types.YChild{"Path", nil})
    for i := range prefixEventOffline.Path {
        prefixEventOffline.EntityData.Children.Append(types.GetSegmentPath(prefixEventOffline.Path[i]), types.YChild{"Path", prefixEventOffline.Path[i]})
    }
    prefixEventOffline.EntityData.Children.Append("trigger-lsa", types.YChild{"TriggerLsa", nil})
    for i := range prefixEventOffline.TriggerLsa {
        prefixEventOffline.EntityData.Children.Append(types.GetSegmentPath(prefixEventOffline.TriggerLsa[i]), types.YChild{"TriggerLsa", prefixEventOffline.TriggerLsa[i]})
    }
    prefixEventOffline.EntityData.Children.Append("time-line", types.YChild{"TimeLine", nil})
    for i := range prefixEventOffline.TimeLine {
        prefixEventOffline.EntityData.Children.Append(types.GetSegmentPath(prefixEventOffline.TimeLine[i]), types.YChild{"TimeLine", prefixEventOffline.TimeLine[i]})
    }
    prefixEventOffline.EntityData.Children.Append("lsa-processed", types.YChild{"LsaProcessed", nil})
    for i := range prefixEventOffline.LsaProcessed {
        prefixEventOffline.EntityData.Children.Append(types.GetSegmentPath(prefixEventOffline.LsaProcessed[i]), types.YChild{"LsaProcessed", prefixEventOffline.LsaProcessed[i]})
    }
    prefixEventOffline.EntityData.Leafs = types.NewOrderedMap()
    prefixEventOffline.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", prefixEventOffline.EventId})
    prefixEventOffline.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", prefixEventOffline.Prefix})
    prefixEventOffline.EntityData.Leafs.Append("prefix-lenth", types.YLeaf{"PrefixLenth", prefixEventOffline.PrefixLenth})
    prefixEventOffline.EntityData.Leafs.Append("spf-run-no", types.YLeaf{"SpfRunNo", prefixEventOffline.SpfRunNo})
    prefixEventOffline.EntityData.Leafs.Append("ipfrr-event-id", types.YLeaf{"IpfrrEventId", prefixEventOffline.IpfrrEventId})
    prefixEventOffline.EntityData.Leafs.Append("threshold-exceeded", types.YLeaf{"ThresholdExceeded", prefixEventOffline.ThresholdExceeded})
    prefixEventOffline.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", prefixEventOffline.Priority})
    prefixEventOffline.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", prefixEventOffline.ChangeType})
    prefixEventOffline.EntityData.Leafs.Append("route-type", types.YLeaf{"RouteType", prefixEventOffline.RouteType})
    prefixEventOffline.EntityData.Leafs.Append("route-path-change-type", types.YLeaf{"RoutePathChangeType", prefixEventOffline.RoutePathChangeType})
    prefixEventOffline.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", prefixEventOffline.Cost})
    prefixEventOffline.EntityData.Leafs.Append("trigger-time", types.YLeaf{"TriggerTime", prefixEventOffline.TriggerTime})

    prefixEventOffline.EntityData.YListKeys = []string {"EventId"}

    return &(prefixEventOffline.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_IpConvergenceTime
// Convergence time for IP route programming
type Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_IpConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (ipConvergenceTime *Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_IpConvergenceTime) GetEntityData() *types.CommonEntityData {
    ipConvergenceTime.EntityData.YFilter = ipConvergenceTime.YFilter
    ipConvergenceTime.EntityData.YangName = "ip-convergence-time"
    ipConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    ipConvergenceTime.EntityData.ParentYangName = "prefix-event-offline"
    ipConvergenceTime.EntityData.SegmentPath = "ip-convergence-time"
    ipConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipConvergenceTime.EntityData.Children = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    ipConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", ipConvergenceTime.MinimumTime})
    ipConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", ipConvergenceTime.MaximumTime})
    ipConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", ipConvergenceTime.SlowestNodeName})
    ipConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", ipConvergenceTime.FastestNodeName})

    ipConvergenceTime.EntityData.YListKeys = []string {}

    return &(ipConvergenceTime.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_MplsConvergenceTime
// Convergence time for MPLS label programming
type Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_MplsConvergenceTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MinimumTime interface{}

    // Maximum time(in seconds.milliseconds). The type is string. Units are
    // millisecond.
    MaximumTime interface{}

    // Linecard node name which took the maximum time. The type is string.
    SlowestNodeName interface{}

    // Linecard node name which took the minimum time. The type is string.
    FastestNodeName interface{}
}

func (mplsConvergenceTime *Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_MplsConvergenceTime) GetEntityData() *types.CommonEntityData {
    mplsConvergenceTime.EntityData.YFilter = mplsConvergenceTime.YFilter
    mplsConvergenceTime.EntityData.YangName = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.BundleName = "cisco_ios_xr"
    mplsConvergenceTime.EntityData.ParentYangName = "prefix-event-offline"
    mplsConvergenceTime.EntityData.SegmentPath = "mpls-convergence-time"
    mplsConvergenceTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsConvergenceTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsConvergenceTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsConvergenceTime.EntityData.Children = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs = types.NewOrderedMap()
    mplsConvergenceTime.EntityData.Leafs.Append("minimum-time", types.YLeaf{"MinimumTime", mplsConvergenceTime.MinimumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("maximum-time", types.YLeaf{"MaximumTime", mplsConvergenceTime.MaximumTime})
    mplsConvergenceTime.EntityData.Leafs.Append("slowest-node-name", types.YLeaf{"SlowestNodeName", mplsConvergenceTime.SlowestNodeName})
    mplsConvergenceTime.EntityData.Leafs.Append("fastest-node-name", types.YLeaf{"FastestNodeName", mplsConvergenceTime.FastestNodeName})

    mplsConvergenceTime.EntityData.YListKeys = []string {}

    return &(mplsConvergenceTime.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path
// Path information
type Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Path Metric. The type is interface{} with range: 0..4294967295.
    PathMetric interface{}

    // Backup Path Informatoin. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path_LfaPath.
    LfaPath []*Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path_LfaPath
}

func (path *Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "prefix-event-offline"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Children.Append("lfa-path", types.YChild{"LfaPath", nil})
    for i := range path.LfaPath {
        path.EntityData.Children.Append(types.GetSegmentPath(path.LfaPath[i]), types.YChild{"LfaPath", path.LfaPath[i]})
    }
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", path.InterfaceName})
    path.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", path.NeighbourAddress})
    path.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", path.ChangeType})
    path.EntityData.Leafs.Append("path-metric", types.YLeaf{"PathMetric", path.PathMetric})

    path.EntityData.YListKeys = []string {}

    return &(path.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path_LfaPath
// Backup Path Informatoin
type Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path_LfaPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of LFA. The type is RcmdShowIpfrrLfa.
    LfaType interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Nexthop Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighbourAddress interface{}

    // Event Add/Delete. The type is RcmdChange.
    ChangeType interface{}

    // Path Metric. The type is interface{} with range: 0..4294967295.
    PathMetric interface{}

    // Remote Node ID, in case of Remote LFA. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteNodeId interface{}
}

func (lfaPath *Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_Path_LfaPath) GetEntityData() *types.CommonEntityData {
    lfaPath.EntityData.YFilter = lfaPath.YFilter
    lfaPath.EntityData.YangName = "lfa-path"
    lfaPath.EntityData.BundleName = "cisco_ios_xr"
    lfaPath.EntityData.ParentYangName = "path"
    lfaPath.EntityData.SegmentPath = "lfa-path"
    lfaPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lfaPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lfaPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lfaPath.EntityData.Children = types.NewOrderedMap()
    lfaPath.EntityData.Leafs = types.NewOrderedMap()
    lfaPath.EntityData.Leafs.Append("lfa-type", types.YLeaf{"LfaType", lfaPath.LfaType})
    lfaPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", lfaPath.InterfaceName})
    lfaPath.EntityData.Leafs.Append("neighbour-address", types.YLeaf{"NeighbourAddress", lfaPath.NeighbourAddress})
    lfaPath.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lfaPath.ChangeType})
    lfaPath.EntityData.Leafs.Append("path-metric", types.YLeaf{"PathMetric", lfaPath.PathMetric})
    lfaPath.EntityData.Leafs.Append("remote-node-id", types.YLeaf{"RemoteNodeId", lfaPath.RemoteNodeId})

    lfaPath.EntityData.YListKeys = []string {}

    return &(lfaPath.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TriggerLsa
// LSA that triggered this event
type Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TriggerLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (triggerLsa *Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TriggerLsa) GetEntityData() *types.CommonEntityData {
    triggerLsa.EntityData.YFilter = triggerLsa.YFilter
    triggerLsa.EntityData.YangName = "trigger-lsa"
    triggerLsa.EntityData.BundleName = "cisco_ios_xr"
    triggerLsa.EntityData.ParentYangName = "prefix-event-offline"
    triggerLsa.EntityData.SegmentPath = "trigger-lsa"
    triggerLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    triggerLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    triggerLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    triggerLsa.EntityData.Children = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs = types.NewOrderedMap()
    triggerLsa.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", triggerLsa.LsaId})
    triggerLsa.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", triggerLsa.SequenceNumber})
    triggerLsa.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", triggerLsa.LsaType})
    triggerLsa.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", triggerLsa.OriginRouterId})
    triggerLsa.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", triggerLsa.ChangeType})
    triggerLsa.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", triggerLsa.ReceptionTime})

    triggerLsa.EntityData.YListKeys = []string {}

    return &(triggerLsa.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine
// Timeline information
type Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route origin (routing protocol). The type is string.
    RouteOrigin interface{}

    // Entry point of IPv4 RIB. The type is string.
    RiBv4Enter interface{}

    // Exit point from IPv4 RIB to FIBs. The type is string.
    RiBv4Exit interface{}

    // Route Redistribute point from IPv4 RIB to LDP. The type is string.
    RiBv4Redistribute interface{}

    // Entry point of LDP. The type is string.
    LdpEnter interface{}

    // Exit point of LDP to LSD. The type is string.
    LdpExit interface{}

    // Entry point of LSD. The type is string.
    LsdEnter interface{}

    // Exit point of LSD to FIBs. The type is string.
    LsdExit interface{}

    // List of Linecards' completion point for IP routes. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcIp.
    LcIp []*Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcIp

    // List of Linecards' completion point for MPLS labels. The type is slice of
    // Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcMpls.
    LcMpls []*Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcMpls
}

func (timeLine *Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine) GetEntityData() *types.CommonEntityData {
    timeLine.EntityData.YFilter = timeLine.YFilter
    timeLine.EntityData.YangName = "time-line"
    timeLine.EntityData.BundleName = "cisco_ios_xr"
    timeLine.EntityData.ParentYangName = "prefix-event-offline"
    timeLine.EntityData.SegmentPath = "time-line"
    timeLine.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeLine.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeLine.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeLine.EntityData.Children = types.NewOrderedMap()
    timeLine.EntityData.Children.Append("lc-ip", types.YChild{"LcIp", nil})
    for i := range timeLine.LcIp {
        timeLine.EntityData.Children.Append(types.GetSegmentPath(timeLine.LcIp[i]), types.YChild{"LcIp", timeLine.LcIp[i]})
    }
    timeLine.EntityData.Children.Append("lc-mpls", types.YChild{"LcMpls", nil})
    for i := range timeLine.LcMpls {
        timeLine.EntityData.Children.Append(types.GetSegmentPath(timeLine.LcMpls[i]), types.YChild{"LcMpls", timeLine.LcMpls[i]})
    }
    timeLine.EntityData.Leafs = types.NewOrderedMap()
    timeLine.EntityData.Leafs.Append("route-origin", types.YLeaf{"RouteOrigin", timeLine.RouteOrigin})
    timeLine.EntityData.Leafs.Append("ri-bv4-enter", types.YLeaf{"RiBv4Enter", timeLine.RiBv4Enter})
    timeLine.EntityData.Leafs.Append("ri-bv4-exit", types.YLeaf{"RiBv4Exit", timeLine.RiBv4Exit})
    timeLine.EntityData.Leafs.Append("ri-bv4-redistribute", types.YLeaf{"RiBv4Redistribute", timeLine.RiBv4Redistribute})
    timeLine.EntityData.Leafs.Append("ldp-enter", types.YLeaf{"LdpEnter", timeLine.LdpEnter})
    timeLine.EntityData.Leafs.Append("ldp-exit", types.YLeaf{"LdpExit", timeLine.LdpExit})
    timeLine.EntityData.Leafs.Append("lsd-enter", types.YLeaf{"LsdEnter", timeLine.LsdEnter})
    timeLine.EntityData.Leafs.Append("lsd-exit", types.YLeaf{"LsdExit", timeLine.LsdExit})

    timeLine.EntityData.YListKeys = []string {}

    return &(timeLine.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcIp
// List of Linecards' completion point for IP
// routes
type Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB. The type is string.
    FibComplete interface{}
}

func (lcIp *Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcIp) GetEntityData() *types.CommonEntityData {
    lcIp.EntityData.YFilter = lcIp.YFilter
    lcIp.EntityData.YangName = "lc-ip"
    lcIp.EntityData.BundleName = "cisco_ios_xr"
    lcIp.EntityData.ParentYangName = "time-line"
    lcIp.EntityData.SegmentPath = "lc-ip"
    lcIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcIp.EntityData.Children = types.NewOrderedMap()
    lcIp.EntityData.Leafs = types.NewOrderedMap()
    lcIp.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcIp.NodeName})
    lcIp.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcIp.Speed})
    lcIp.EntityData.Leafs.Append("fib-complete", types.YLeaf{"FibComplete", lcIp.FibComplete})

    lcIp.EntityData.YListKeys = []string {}

    return &(lcIp.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcMpls
// List of Linecards' completion point for MPLS
// labels
type Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard node name. The type is string.
    NodeName interface{}

    // Relative convergence speed. The type is RcmdLinecardSpeed.
    Speed interface{}

    // Completion point of FIB. The type is string.
    FibComplete interface{}
}

func (lcMpls *Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_TimeLine_LcMpls) GetEntityData() *types.CommonEntityData {
    lcMpls.EntityData.YFilter = lcMpls.YFilter
    lcMpls.EntityData.YangName = "lc-mpls"
    lcMpls.EntityData.BundleName = "cisco_ios_xr"
    lcMpls.EntityData.ParentYangName = "time-line"
    lcMpls.EntityData.SegmentPath = "lc-mpls"
    lcMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcMpls.EntityData.Children = types.NewOrderedMap()
    lcMpls.EntityData.Leafs = types.NewOrderedMap()
    lcMpls.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", lcMpls.NodeName})
    lcMpls.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", lcMpls.Speed})
    lcMpls.EntityData.Leafs.Append("fib-complete", types.YLeaf{"FibComplete", lcMpls.FibComplete})

    lcMpls.EntityData.YListKeys = []string {}

    return &(lcMpls.EntityData)
}

// Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_LsaProcessed
// List of LSAs processed
type Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_LsaProcessed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsaId interface{}

    // Sequence Number. The type is string.
    SequenceNumber interface{}

    // LSA type. The type is RcmdLsa.
    LsaType interface{}

    // Originating Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginRouterId interface{}

    // Add, Delete, Modify. The type is RcmdLsChange.
    ChangeType interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}
}

func (lsaProcessed *Rcmd_Isis_Instances_Instance_PrefixEventOfflines_PrefixEventOffline_LsaProcessed) GetEntityData() *types.CommonEntityData {
    lsaProcessed.EntityData.YFilter = lsaProcessed.YFilter
    lsaProcessed.EntityData.YangName = "lsa-processed"
    lsaProcessed.EntityData.BundleName = "cisco_ios_xr"
    lsaProcessed.EntityData.ParentYangName = "prefix-event-offline"
    lsaProcessed.EntityData.SegmentPath = "lsa-processed"
    lsaProcessed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsaProcessed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsaProcessed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsaProcessed.EntityData.Children = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs = types.NewOrderedMap()
    lsaProcessed.EntityData.Leafs.Append("lsa-id", types.YLeaf{"LsaId", lsaProcessed.LsaId})
    lsaProcessed.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", lsaProcessed.SequenceNumber})
    lsaProcessed.EntityData.Leafs.Append("lsa-type", types.YLeaf{"LsaType", lsaProcessed.LsaType})
    lsaProcessed.EntityData.Leafs.Append("origin-router-id", types.YLeaf{"OriginRouterId", lsaProcessed.OriginRouterId})
    lsaProcessed.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", lsaProcessed.ChangeType})
    lsaProcessed.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", lsaProcessed.ReceptionTime})

    lsaProcessed.EntityData.YListKeys = []string {}

    return &(lsaProcessed.EntityData)
}

// Rcmd_Isis_Instances_Instance_LspRegenerateds
// Regenerated LSP data
type Rcmd_Isis_Instances_Instance_LspRegenerateds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Regenerated LSP data. The type is slice of
    // Rcmd_Isis_Instances_Instance_LspRegenerateds_LspRegenerated.
    LspRegenerated []*Rcmd_Isis_Instances_Instance_LspRegenerateds_LspRegenerated
}

func (lspRegenerateds *Rcmd_Isis_Instances_Instance_LspRegenerateds) GetEntityData() *types.CommonEntityData {
    lspRegenerateds.EntityData.YFilter = lspRegenerateds.YFilter
    lspRegenerateds.EntityData.YangName = "lsp-regenerateds"
    lspRegenerateds.EntityData.BundleName = "cisco_ios_xr"
    lspRegenerateds.EntityData.ParentYangName = "instance"
    lspRegenerateds.EntityData.SegmentPath = "lsp-regenerateds"
    lspRegenerateds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspRegenerateds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspRegenerateds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspRegenerateds.EntityData.Children = types.NewOrderedMap()
    lspRegenerateds.EntityData.Children.Append("lsp-regenerated", types.YChild{"LspRegenerated", nil})
    for i := range lspRegenerateds.LspRegenerated {
        lspRegenerateds.EntityData.Children.Append(types.GetSegmentPath(lspRegenerateds.LspRegenerated[i]), types.YChild{"LspRegenerated", lspRegenerateds.LspRegenerated[i]})
    }
    lspRegenerateds.EntityData.Leafs = types.NewOrderedMap()

    lspRegenerateds.EntityData.YListKeys = []string {}

    return &(lspRegenerateds.EntityData)
}

// Rcmd_Isis_Instances_Instance_LspRegenerateds_LspRegenerated
// Regenerated LSP data
type Rcmd_Isis_Instances_Instance_LspRegenerateds_LspRegenerated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Data for a particular regenerated LSP. The type is
    // interface{} with range: 1..4294967295.
    SerialNumber interface{}

    // Serial Number of the session event. The type is interface{} with range:
    // 0..4294967295.
    SerialNumberXr interface{}

    // LSP ID. The type is string.
    LspId interface{}

    // Sequence Number. The type is interface{} with range: 0..4294967295.
    SequenceNumber interface{}

    // Reception Time on router (in hh:mm:ss.msec). The type is string.
    ReceptionTime interface{}

    // ISIS Level. The type is RcmdIsisLvl.
    IsisLevel interface{}

    // SPF Run Number. The type is interface{} with range: 0..4294967295.
    SpfRunNumber interface{}

    // Trigger reasons for LSP regeneration. Example: pr^ - periodic, cr^ - clear
    // (Check the documentation for the entire list). The type is string.
    Reason interface{}
}

func (lspRegenerated *Rcmd_Isis_Instances_Instance_LspRegenerateds_LspRegenerated) GetEntityData() *types.CommonEntityData {
    lspRegenerated.EntityData.YFilter = lspRegenerated.YFilter
    lspRegenerated.EntityData.YangName = "lsp-regenerated"
    lspRegenerated.EntityData.BundleName = "cisco_ios_xr"
    lspRegenerated.EntityData.ParentYangName = "lsp-regenerateds"
    lspRegenerated.EntityData.SegmentPath = "lsp-regenerated" + types.AddKeyToken(lspRegenerated.SerialNumber, "serial-number")
    lspRegenerated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspRegenerated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspRegenerated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspRegenerated.EntityData.Children = types.NewOrderedMap()
    lspRegenerated.EntityData.Leafs = types.NewOrderedMap()
    lspRegenerated.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", lspRegenerated.SerialNumber})
    lspRegenerated.EntityData.Leafs.Append("serial-number-xr", types.YLeaf{"SerialNumberXr", lspRegenerated.SerialNumberXr})
    lspRegenerated.EntityData.Leafs.Append("lsp-id", types.YLeaf{"LspId", lspRegenerated.LspId})
    lspRegenerated.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", lspRegenerated.SequenceNumber})
    lspRegenerated.EntityData.Leafs.Append("reception-time", types.YLeaf{"ReceptionTime", lspRegenerated.ReceptionTime})
    lspRegenerated.EntityData.Leafs.Append("isis-level", types.YLeaf{"IsisLevel", lspRegenerated.IsisLevel})
    lspRegenerated.EntityData.Leafs.Append("spf-run-number", types.YLeaf{"SpfRunNumber", lspRegenerated.SpfRunNumber})
    lspRegenerated.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", lspRegenerated.Reason})

    lspRegenerated.EntityData.YListKeys = []string {"SerialNumber"}

    return &(lspRegenerated.EntityData)
}

// Rcmd_Memory
// Memory Info
type Rcmd_Memory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Memory Info. The type is slice of Rcmd_Memory_MemoryInfo.
    MemoryInfo []*Rcmd_Memory_MemoryInfo

    // Memory Info. The type is slice of Rcmd_Memory_EdmMemoryInfo.
    EdmMemoryInfo []*Rcmd_Memory_EdmMemoryInfo

    // Memory Info. The type is slice of Rcmd_Memory_StringMemoryInfo.
    StringMemoryInfo []*Rcmd_Memory_StringMemoryInfo
}

func (memory *Rcmd_Memory) GetEntityData() *types.CommonEntityData {
    memory.EntityData.YFilter = memory.YFilter
    memory.EntityData.YangName = "memory"
    memory.EntityData.BundleName = "cisco_ios_xr"
    memory.EntityData.ParentYangName = "rcmd"
    memory.EntityData.SegmentPath = "memory"
    memory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memory.EntityData.Children = types.NewOrderedMap()
    memory.EntityData.Children.Append("memory-info", types.YChild{"MemoryInfo", nil})
    for i := range memory.MemoryInfo {
        memory.EntityData.Children.Append(types.GetSegmentPath(memory.MemoryInfo[i]), types.YChild{"MemoryInfo", memory.MemoryInfo[i]})
    }
    memory.EntityData.Children.Append("edm-memory-info", types.YChild{"EdmMemoryInfo", nil})
    for i := range memory.EdmMemoryInfo {
        memory.EntityData.Children.Append(types.GetSegmentPath(memory.EdmMemoryInfo[i]), types.YChild{"EdmMemoryInfo", memory.EdmMemoryInfo[i]})
    }
    memory.EntityData.Children.Append("string-memory-info", types.YChild{"StringMemoryInfo", nil})
    for i := range memory.StringMemoryInfo {
        memory.EntityData.Children.Append(types.GetSegmentPath(memory.StringMemoryInfo[i]), types.YChild{"StringMemoryInfo", memory.StringMemoryInfo[i]})
    }
    memory.EntityData.Leafs = types.NewOrderedMap()

    memory.EntityData.YListKeys = []string {}

    return &(memory.EntityData)
}

// Rcmd_Memory_MemoryInfo
// Memory Info
type Rcmd_Memory_MemoryInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Structure Name. The type is string.
    StructureName interface{}

    // Size of the datastructure. The type is interface{} with range:
    // 0..4294967295.
    Size interface{}

    // Current Count. The type is interface{} with range: 0..4294967295.
    CurrentCount interface{}

    // Allocation Fails. The type is interface{} with range: 0..4294967295.
    AllocFails interface{}

    // Allocated count. The type is interface{} with range: 0..4294967295.
    AllocCount interface{}

    // Freed Count. The type is interface{} with range: 0..4294967295.
    FreedCount interface{}

    // Memory Type. The type is RcmdShowMem.
    MemoryType interface{}
}

func (memoryInfo *Rcmd_Memory_MemoryInfo) GetEntityData() *types.CommonEntityData {
    memoryInfo.EntityData.YFilter = memoryInfo.YFilter
    memoryInfo.EntityData.YangName = "memory-info"
    memoryInfo.EntityData.BundleName = "cisco_ios_xr"
    memoryInfo.EntityData.ParentYangName = "memory"
    memoryInfo.EntityData.SegmentPath = "memory-info"
    memoryInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryInfo.EntityData.Children = types.NewOrderedMap()
    memoryInfo.EntityData.Leafs = types.NewOrderedMap()
    memoryInfo.EntityData.Leafs.Append("structure-name", types.YLeaf{"StructureName", memoryInfo.StructureName})
    memoryInfo.EntityData.Leafs.Append("size", types.YLeaf{"Size", memoryInfo.Size})
    memoryInfo.EntityData.Leafs.Append("current-count", types.YLeaf{"CurrentCount", memoryInfo.CurrentCount})
    memoryInfo.EntityData.Leafs.Append("alloc-fails", types.YLeaf{"AllocFails", memoryInfo.AllocFails})
    memoryInfo.EntityData.Leafs.Append("alloc-count", types.YLeaf{"AllocCount", memoryInfo.AllocCount})
    memoryInfo.EntityData.Leafs.Append("freed-count", types.YLeaf{"FreedCount", memoryInfo.FreedCount})
    memoryInfo.EntityData.Leafs.Append("memory-type", types.YLeaf{"MemoryType", memoryInfo.MemoryType})

    memoryInfo.EntityData.YListKeys = []string {}

    return &(memoryInfo.EntityData)
}

// Rcmd_Memory_EdmMemoryInfo
// Memory Info
type Rcmd_Memory_EdmMemoryInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Size of the block. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // Total request. The type is interface{} with range: 0..4294967295.
    Total interface{}

    // Cache-hit success. The type is interface{} with range: 0..4294967295.
    Success interface{}

    // Cache-hit failure. The type is interface{} with range: 0..4294967295.
    Failure interface{}
}

func (edmMemoryInfo *Rcmd_Memory_EdmMemoryInfo) GetEntityData() *types.CommonEntityData {
    edmMemoryInfo.EntityData.YFilter = edmMemoryInfo.YFilter
    edmMemoryInfo.EntityData.YangName = "edm-memory-info"
    edmMemoryInfo.EntityData.BundleName = "cisco_ios_xr"
    edmMemoryInfo.EntityData.ParentYangName = "memory"
    edmMemoryInfo.EntityData.SegmentPath = "edm-memory-info"
    edmMemoryInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    edmMemoryInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    edmMemoryInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    edmMemoryInfo.EntityData.Children = types.NewOrderedMap()
    edmMemoryInfo.EntityData.Leafs = types.NewOrderedMap()
    edmMemoryInfo.EntityData.Leafs.Append("size", types.YLeaf{"Size", edmMemoryInfo.Size})
    edmMemoryInfo.EntityData.Leafs.Append("total", types.YLeaf{"Total", edmMemoryInfo.Total})
    edmMemoryInfo.EntityData.Leafs.Append("success", types.YLeaf{"Success", edmMemoryInfo.Success})
    edmMemoryInfo.EntityData.Leafs.Append("failure", types.YLeaf{"Failure", edmMemoryInfo.Failure})

    edmMemoryInfo.EntityData.YListKeys = []string {}

    return &(edmMemoryInfo.EntityData)
}

// Rcmd_Memory_StringMemoryInfo
// Memory Info
type Rcmd_Memory_StringMemoryInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Size of the block. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // Total request. The type is interface{} with range: 0..4294967295.
    Total interface{}

    // Cache-hit success. The type is interface{} with range: 0..4294967295.
    Success interface{}

    // Cache-hit failure. The type is interface{} with range: 0..4294967295.
    Failure interface{}
}

func (stringMemoryInfo *Rcmd_Memory_StringMemoryInfo) GetEntityData() *types.CommonEntityData {
    stringMemoryInfo.EntityData.YFilter = stringMemoryInfo.YFilter
    stringMemoryInfo.EntityData.YangName = "string-memory-info"
    stringMemoryInfo.EntityData.BundleName = "cisco_ios_xr"
    stringMemoryInfo.EntityData.ParentYangName = "memory"
    stringMemoryInfo.EntityData.SegmentPath = "string-memory-info"
    stringMemoryInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stringMemoryInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stringMemoryInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stringMemoryInfo.EntityData.Children = types.NewOrderedMap()
    stringMemoryInfo.EntityData.Leafs = types.NewOrderedMap()
    stringMemoryInfo.EntityData.Leafs.Append("size", types.YLeaf{"Size", stringMemoryInfo.Size})
    stringMemoryInfo.EntityData.Leafs.Append("total", types.YLeaf{"Total", stringMemoryInfo.Total})
    stringMemoryInfo.EntityData.Leafs.Append("success", types.YLeaf{"Success", stringMemoryInfo.Success})
    stringMemoryInfo.EntityData.Leafs.Append("failure", types.YLeaf{"Failure", stringMemoryInfo.Failure})

    stringMemoryInfo.EntityData.YListKeys = []string {}

    return &(stringMemoryInfo.EntityData)
}

// Rcmd_Ldp
// LDP data
type Rcmd_Ldp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session Events.
    Sessions Rcmd_Ldp_Sessions

    // Remote LFA Coverage Events.
    RemoteLfaS Rcmd_Ldp_RemoteLfaS

    // Remote LFA Coverage Events.
    RemoteLfaSummaries Rcmd_Ldp_RemoteLfaSummaries
}

func (ldp *Rcmd_Ldp) GetEntityData() *types.CommonEntityData {
    ldp.EntityData.YFilter = ldp.YFilter
    ldp.EntityData.YangName = "ldp"
    ldp.EntityData.BundleName = "cisco_ios_xr"
    ldp.EntityData.ParentYangName = "rcmd"
    ldp.EntityData.SegmentPath = "ldp"
    ldp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldp.EntityData.Children = types.NewOrderedMap()
    ldp.EntityData.Children.Append("sessions", types.YChild{"Sessions", &ldp.Sessions})
    ldp.EntityData.Children.Append("remote-lfa-s", types.YChild{"RemoteLfaS", &ldp.RemoteLfaS})
    ldp.EntityData.Children.Append("remote-lfa-summaries", types.YChild{"RemoteLfaSummaries", &ldp.RemoteLfaSummaries})
    ldp.EntityData.Leafs = types.NewOrderedMap()

    ldp.EntityData.YListKeys = []string {}

    return &(ldp.EntityData)
}

// Rcmd_Ldp_Sessions
// Session Events
type Rcmd_Ldp_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session. The type is slice of Rcmd_Ldp_Sessions_Session.
    Session []*Rcmd_Ldp_Sessions_Session
}

func (sessions *Rcmd_Ldp_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "ldp"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// Rcmd_Ldp_Sessions_Session
// Session
type Rcmd_Ldp_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Event ID. The type is interface{} with
    // range: 1..4294967295.
    EventId interface{}

    // Event ID. The type is interface{} with range: 0..4294967295.
    EventIdXr interface{}

    // Type of event. The type is RcmdLdpEvent.
    EventType interface{}

    // Event Time. The type is string.
    EventTime interface{}

    // Label Space Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // transport address or adjacency address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Adjacency Session Status. The type is RcmdShowLdpNeighbourStatus.
    State interface{}
}

func (session *Rcmd_Ldp_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.EventId, "event-id")
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", session.EventId})
    session.EntityData.Leafs.Append("event-id-xr", types.YLeaf{"EventIdXr", session.EventIdXr})
    session.EntityData.Leafs.Append("event-type", types.YLeaf{"EventType", session.EventType})
    session.EntityData.Leafs.Append("event-time", types.YLeaf{"EventTime", session.EventTime})
    session.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", session.LsrId})
    session.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", session.InterfaceName})
    session.EntityData.Leafs.Append("address", types.YLeaf{"Address", session.Address})
    session.EntityData.Leafs.Append("state", types.YLeaf{"State", session.State})

    session.EntityData.YListKeys = []string {"EventId"}

    return &(session.EntityData)
}

// Rcmd_Ldp_RemoteLfaS
// Remote LFA Coverage Events
type Rcmd_Ldp_RemoteLfaS struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RemoteLFA. The type is slice of Rcmd_Ldp_RemoteLfaS_RemoteLfa.
    RemoteLfa []*Rcmd_Ldp_RemoteLfaS_RemoteLfa
}

func (remoteLfaS *Rcmd_Ldp_RemoteLfaS) GetEntityData() *types.CommonEntityData {
    remoteLfaS.EntityData.YFilter = remoteLfaS.YFilter
    remoteLfaS.EntityData.YangName = "remote-lfa-s"
    remoteLfaS.EntityData.BundleName = "cisco_ios_xr"
    remoteLfaS.EntityData.ParentYangName = "ldp"
    remoteLfaS.EntityData.SegmentPath = "remote-lfa-s"
    remoteLfaS.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteLfaS.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteLfaS.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteLfaS.EntityData.Children = types.NewOrderedMap()
    remoteLfaS.EntityData.Children.Append("remote-lfa", types.YChild{"RemoteLfa", nil})
    for i := range remoteLfaS.RemoteLfa {
        remoteLfaS.EntityData.Children.Append(types.GetSegmentPath(remoteLfaS.RemoteLfa[i]), types.YChild{"RemoteLfa", remoteLfaS.RemoteLfa[i]})
    }
    remoteLfaS.EntityData.Leafs = types.NewOrderedMap()

    remoteLfaS.EntityData.YListKeys = []string {}

    return &(remoteLfaS.EntityData)
}

// Rcmd_Ldp_RemoteLfaS_RemoteLfa
// RemoteLFA
type Rcmd_Ldp_RemoteLfaS_RemoteLfa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Event ID. The type is interface{} with
    // range: 1..4294967295.
    EventId interface{}

    // LDP-rLFA Event ID. The type is interface{} with range: 0..4294967295.
    EventIdXr interface{}

    // End of IGP LFA Calculation Time (eg: Apr 24 13 :16:04.961). The type is
    // string.
    EndOfCalculationTime interface{}

    // IGP Protocol. The type is RcmdProtocolId.
    IgpProtocol interface{}

    // Process Name. The type is string.
    ProcessName interface{}

    // IGP IP-FRR Event ID (ref: rcmd_show_ipfrr_event_info(EventID)). The type is
    // interface{} with range: 0..4294967295.
    IpfrrEventId interface{}

    // Coverage Below Threshold. The type is bool.
    BelowThreshold interface{}

    // RLFA Statistics categorized by session state. The type is slice of
    // Rcmd_Ldp_RemoteLfaS_RemoteLfa_SessionStatistic.
    SessionStatistic []*Rcmd_Ldp_RemoteLfaS_RemoteLfa_SessionStatistic

    // Remote Node Information. The type is slice of
    // Rcmd_Ldp_RemoteLfaS_RemoteLfa_RemoteNode.
    RemoteNode []*Rcmd_Ldp_RemoteLfaS_RemoteLfa_RemoteNode

    // Logs Information. The type is slice of Rcmd_Ldp_RemoteLfaS_RemoteLfa_Logs.
    Logs []*Rcmd_Ldp_RemoteLfaS_RemoteLfa_Logs
}

func (remoteLfa *Rcmd_Ldp_RemoteLfaS_RemoteLfa) GetEntityData() *types.CommonEntityData {
    remoteLfa.EntityData.YFilter = remoteLfa.YFilter
    remoteLfa.EntityData.YangName = "remote-lfa"
    remoteLfa.EntityData.BundleName = "cisco_ios_xr"
    remoteLfa.EntityData.ParentYangName = "remote-lfa-s"
    remoteLfa.EntityData.SegmentPath = "remote-lfa" + types.AddKeyToken(remoteLfa.EventId, "event-id")
    remoteLfa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteLfa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteLfa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteLfa.EntityData.Children = types.NewOrderedMap()
    remoteLfa.EntityData.Children.Append("session-statistic", types.YChild{"SessionStatistic", nil})
    for i := range remoteLfa.SessionStatistic {
        remoteLfa.EntityData.Children.Append(types.GetSegmentPath(remoteLfa.SessionStatistic[i]), types.YChild{"SessionStatistic", remoteLfa.SessionStatistic[i]})
    }
    remoteLfa.EntityData.Children.Append("remote-node", types.YChild{"RemoteNode", nil})
    for i := range remoteLfa.RemoteNode {
        remoteLfa.EntityData.Children.Append(types.GetSegmentPath(remoteLfa.RemoteNode[i]), types.YChild{"RemoteNode", remoteLfa.RemoteNode[i]})
    }
    remoteLfa.EntityData.Children.Append("logs", types.YChild{"Logs", nil})
    for i := range remoteLfa.Logs {
        remoteLfa.EntityData.Children.Append(types.GetSegmentPath(remoteLfa.Logs[i]), types.YChild{"Logs", remoteLfa.Logs[i]})
    }
    remoteLfa.EntityData.Leafs = types.NewOrderedMap()
    remoteLfa.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", remoteLfa.EventId})
    remoteLfa.EntityData.Leafs.Append("event-id-xr", types.YLeaf{"EventIdXr", remoteLfa.EventIdXr})
    remoteLfa.EntityData.Leafs.Append("end-of-calculation-time", types.YLeaf{"EndOfCalculationTime", remoteLfa.EndOfCalculationTime})
    remoteLfa.EntityData.Leafs.Append("igp-protocol", types.YLeaf{"IgpProtocol", remoteLfa.IgpProtocol})
    remoteLfa.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", remoteLfa.ProcessName})
    remoteLfa.EntityData.Leafs.Append("ipfrr-event-id", types.YLeaf{"IpfrrEventId", remoteLfa.IpfrrEventId})
    remoteLfa.EntityData.Leafs.Append("below-threshold", types.YLeaf{"BelowThreshold", remoteLfa.BelowThreshold})

    remoteLfa.EntityData.YListKeys = []string {"EventId"}

    return &(remoteLfa.EntityData)
}

// Rcmd_Ldp_RemoteLfaS_RemoteLfa_SessionStatistic
// RLFA Statistics categorized by session state
type Rcmd_Ldp_RemoteLfaS_RemoteLfa_SessionStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session State. The type is RcmdShowLdpSessionState.
    SessionState interface{}

    // LDP Session Count. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}

    // Total Route Count. The type is interface{} with range: 0..4294967295.
    RouteCount interface{}

    // Total Path Count. The type is interface{} with range: 0..4294967295.
    PathCount interface{}

    // Remote Label Count. The type is interface{} with range: 0..4294967295.
    RemoteLabelCount interface{}

    // Protected Route Count. The type is interface{} with range: 0..4294967295.
    ProtectedRouteCount interface{}

    // Protected Path Count. The type is interface{} with range: 0..4294967295.
    ProtectedPathCount interface{}
}

func (sessionStatistic *Rcmd_Ldp_RemoteLfaS_RemoteLfa_SessionStatistic) GetEntityData() *types.CommonEntityData {
    sessionStatistic.EntityData.YFilter = sessionStatistic.YFilter
    sessionStatistic.EntityData.YangName = "session-statistic"
    sessionStatistic.EntityData.BundleName = "cisco_ios_xr"
    sessionStatistic.EntityData.ParentYangName = "remote-lfa"
    sessionStatistic.EntityData.SegmentPath = "session-statistic"
    sessionStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionStatistic.EntityData.Children = types.NewOrderedMap()
    sessionStatistic.EntityData.Leafs = types.NewOrderedMap()
    sessionStatistic.EntityData.Leafs.Append("session-state", types.YLeaf{"SessionState", sessionStatistic.SessionState})
    sessionStatistic.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", sessionStatistic.SessionCount})
    sessionStatistic.EntityData.Leafs.Append("route-count", types.YLeaf{"RouteCount", sessionStatistic.RouteCount})
    sessionStatistic.EntityData.Leafs.Append("path-count", types.YLeaf{"PathCount", sessionStatistic.PathCount})
    sessionStatistic.EntityData.Leafs.Append("remote-label-count", types.YLeaf{"RemoteLabelCount", sessionStatistic.RemoteLabelCount})
    sessionStatistic.EntityData.Leafs.Append("protected-route-count", types.YLeaf{"ProtectedRouteCount", sessionStatistic.ProtectedRouteCount})
    sessionStatistic.EntityData.Leafs.Append("protected-path-count", types.YLeaf{"ProtectedPathCount", sessionStatistic.ProtectedPathCount})

    sessionStatistic.EntityData.YListKeys = []string {}

    return &(sessionStatistic.EntityData)
}

// Rcmd_Ldp_RemoteLfaS_RemoteLfa_RemoteNode
// Remote Node Information
type Rcmd_Ldp_RemoteLfaS_RemoteLfa_RemoteNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote Node ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteNodeId interface{}

    // Label Space Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // Transport Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    TransportAddress interface{}

    // Inuse time of the Session. The type is string.
    InUseTime interface{}

    // Session State. The type is RcmdShowLdpSessionState.
    SessionState interface{}

    // Total Route Count. The type is interface{} with range: 0..4294967295.
    RouteCount interface{}

    // Total Path Count. The type is interface{} with range: 0..4294967295.
    PathCount interface{}

    // Remote Label Count. The type is interface{} with range: 0..4294967295.
    RemoteLabelCount interface{}

    // Protected Route Count. The type is interface{} with range: 0..4294967295.
    ProtectedRouteCount interface{}

    // Protected Path Count. The type is interface{} with range: 0..4294967295.
    ProtectedPathCount interface{}
}

func (remoteNode *Rcmd_Ldp_RemoteLfaS_RemoteLfa_RemoteNode) GetEntityData() *types.CommonEntityData {
    remoteNode.EntityData.YFilter = remoteNode.YFilter
    remoteNode.EntityData.YangName = "remote-node"
    remoteNode.EntityData.BundleName = "cisco_ios_xr"
    remoteNode.EntityData.ParentYangName = "remote-lfa"
    remoteNode.EntityData.SegmentPath = "remote-node"
    remoteNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNode.EntityData.Children = types.NewOrderedMap()
    remoteNode.EntityData.Leafs = types.NewOrderedMap()
    remoteNode.EntityData.Leafs.Append("remote-node-id", types.YLeaf{"RemoteNodeId", remoteNode.RemoteNodeId})
    remoteNode.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", remoteNode.LsrId})
    remoteNode.EntityData.Leafs.Append("transport-address", types.YLeaf{"TransportAddress", remoteNode.TransportAddress})
    remoteNode.EntityData.Leafs.Append("in-use-time", types.YLeaf{"InUseTime", remoteNode.InUseTime})
    remoteNode.EntityData.Leafs.Append("session-state", types.YLeaf{"SessionState", remoteNode.SessionState})
    remoteNode.EntityData.Leafs.Append("route-count", types.YLeaf{"RouteCount", remoteNode.RouteCount})
    remoteNode.EntityData.Leafs.Append("path-count", types.YLeaf{"PathCount", remoteNode.PathCount})
    remoteNode.EntityData.Leafs.Append("remote-label-count", types.YLeaf{"RemoteLabelCount", remoteNode.RemoteLabelCount})
    remoteNode.EntityData.Leafs.Append("protected-route-count", types.YLeaf{"ProtectedRouteCount", remoteNode.ProtectedRouteCount})
    remoteNode.EntityData.Leafs.Append("protected-path-count", types.YLeaf{"ProtectedPathCount", remoteNode.ProtectedPathCount})

    remoteNode.EntityData.YListKeys = []string {}

    return &(remoteNode.EntityData)
}

// Rcmd_Ldp_RemoteLfaS_RemoteLfa_Logs
// Logs Information
type Rcmd_Ldp_RemoteLfaS_RemoteLfa_Logs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event Time (eg: Apr 24 13:16:04.961). The type is string.
    LogTime interface{}

    // Label Coverage State. The type is RcmdShowLdpConvState.
    LabelCoverageState interface{}

    // Total Route Count. The type is interface{} with range: 0..4294967295.
    RouteCount interface{}

    // Remote Label Count. The type is interface{} with range: 0..4294967295.
    RemoteLabelCount interface{}
}

func (logs *Rcmd_Ldp_RemoteLfaS_RemoteLfa_Logs) GetEntityData() *types.CommonEntityData {
    logs.EntityData.YFilter = logs.YFilter
    logs.EntityData.YangName = "logs"
    logs.EntityData.BundleName = "cisco_ios_xr"
    logs.EntityData.ParentYangName = "remote-lfa"
    logs.EntityData.SegmentPath = "logs"
    logs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logs.EntityData.Children = types.NewOrderedMap()
    logs.EntityData.Leafs = types.NewOrderedMap()
    logs.EntityData.Leafs.Append("log-time", types.YLeaf{"LogTime", logs.LogTime})
    logs.EntityData.Leafs.Append("label-coverage-state", types.YLeaf{"LabelCoverageState", logs.LabelCoverageState})
    logs.EntityData.Leafs.Append("route-count", types.YLeaf{"RouteCount", logs.RouteCount})
    logs.EntityData.Leafs.Append("remote-label-count", types.YLeaf{"RemoteLabelCount", logs.RemoteLabelCount})

    logs.EntityData.YListKeys = []string {}

    return &(logs.EntityData)
}

// Rcmd_Ldp_RemoteLfaSummaries
// Remote LFA Coverage Events
type Rcmd_Ldp_RemoteLfaSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary operational data for Remote LFA. The type is slice of
    // Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary.
    RemoteLfaSummary []*Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary
}

func (remoteLfaSummaries *Rcmd_Ldp_RemoteLfaSummaries) GetEntityData() *types.CommonEntityData {
    remoteLfaSummaries.EntityData.YFilter = remoteLfaSummaries.YFilter
    remoteLfaSummaries.EntityData.YangName = "remote-lfa-summaries"
    remoteLfaSummaries.EntityData.BundleName = "cisco_ios_xr"
    remoteLfaSummaries.EntityData.ParentYangName = "ldp"
    remoteLfaSummaries.EntityData.SegmentPath = "remote-lfa-summaries"
    remoteLfaSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteLfaSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteLfaSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteLfaSummaries.EntityData.Children = types.NewOrderedMap()
    remoteLfaSummaries.EntityData.Children.Append("remote-lfa-summary", types.YChild{"RemoteLfaSummary", nil})
    for i := range remoteLfaSummaries.RemoteLfaSummary {
        remoteLfaSummaries.EntityData.Children.Append(types.GetSegmentPath(remoteLfaSummaries.RemoteLfaSummary[i]), types.YChild{"RemoteLfaSummary", remoteLfaSummaries.RemoteLfaSummary[i]})
    }
    remoteLfaSummaries.EntityData.Leafs = types.NewOrderedMap()

    remoteLfaSummaries.EntityData.YListKeys = []string {}

    return &(remoteLfaSummaries.EntityData)
}

// Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary
// Summary operational data for Remote LFA
type Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Event ID. The type is interface{} with
    // range: 1..4294967295.
    EventId interface{}

    // LDP-rLFA Event ID. The type is interface{} with range: 0..4294967295.
    EventIdXr interface{}

    // End of IGP LFA Calculation Time (eg: Apr 24 13 :16:04.961). The type is
    // string.
    EndOfCalculationTime interface{}

    // IGP Protocol. The type is RcmdProtocolId.
    IgpProtocol interface{}

    // Process Name. The type is string.
    ProcessName interface{}

    // IGP IP-FRR Event ID (ref: rcmd_show_ipfrr_event_info(EventID)). The type is
    // interface{} with range: 0..4294967295.
    IpfrrEventId interface{}

    // Coverage Below Threshold. The type is bool.
    BelowThreshold interface{}

    // RLFA Statistics categorized by session state. The type is slice of
    // Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_SessionStatistic.
    SessionStatistic []*Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_SessionStatistic

    // Remote Node Information. The type is slice of
    // Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_RemoteNode.
    RemoteNode []*Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_RemoteNode

    // Logs Information. The type is slice of
    // Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_Logs.
    Logs []*Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_Logs
}

func (remoteLfaSummary *Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary) GetEntityData() *types.CommonEntityData {
    remoteLfaSummary.EntityData.YFilter = remoteLfaSummary.YFilter
    remoteLfaSummary.EntityData.YangName = "remote-lfa-summary"
    remoteLfaSummary.EntityData.BundleName = "cisco_ios_xr"
    remoteLfaSummary.EntityData.ParentYangName = "remote-lfa-summaries"
    remoteLfaSummary.EntityData.SegmentPath = "remote-lfa-summary" + types.AddKeyToken(remoteLfaSummary.EventId, "event-id")
    remoteLfaSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteLfaSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteLfaSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteLfaSummary.EntityData.Children = types.NewOrderedMap()
    remoteLfaSummary.EntityData.Children.Append("session-statistic", types.YChild{"SessionStatistic", nil})
    for i := range remoteLfaSummary.SessionStatistic {
        remoteLfaSummary.EntityData.Children.Append(types.GetSegmentPath(remoteLfaSummary.SessionStatistic[i]), types.YChild{"SessionStatistic", remoteLfaSummary.SessionStatistic[i]})
    }
    remoteLfaSummary.EntityData.Children.Append("remote-node", types.YChild{"RemoteNode", nil})
    for i := range remoteLfaSummary.RemoteNode {
        remoteLfaSummary.EntityData.Children.Append(types.GetSegmentPath(remoteLfaSummary.RemoteNode[i]), types.YChild{"RemoteNode", remoteLfaSummary.RemoteNode[i]})
    }
    remoteLfaSummary.EntityData.Children.Append("logs", types.YChild{"Logs", nil})
    for i := range remoteLfaSummary.Logs {
        remoteLfaSummary.EntityData.Children.Append(types.GetSegmentPath(remoteLfaSummary.Logs[i]), types.YChild{"Logs", remoteLfaSummary.Logs[i]})
    }
    remoteLfaSummary.EntityData.Leafs = types.NewOrderedMap()
    remoteLfaSummary.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", remoteLfaSummary.EventId})
    remoteLfaSummary.EntityData.Leafs.Append("event-id-xr", types.YLeaf{"EventIdXr", remoteLfaSummary.EventIdXr})
    remoteLfaSummary.EntityData.Leafs.Append("end-of-calculation-time", types.YLeaf{"EndOfCalculationTime", remoteLfaSummary.EndOfCalculationTime})
    remoteLfaSummary.EntityData.Leafs.Append("igp-protocol", types.YLeaf{"IgpProtocol", remoteLfaSummary.IgpProtocol})
    remoteLfaSummary.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", remoteLfaSummary.ProcessName})
    remoteLfaSummary.EntityData.Leafs.Append("ipfrr-event-id", types.YLeaf{"IpfrrEventId", remoteLfaSummary.IpfrrEventId})
    remoteLfaSummary.EntityData.Leafs.Append("below-threshold", types.YLeaf{"BelowThreshold", remoteLfaSummary.BelowThreshold})

    remoteLfaSummary.EntityData.YListKeys = []string {"EventId"}

    return &(remoteLfaSummary.EntityData)
}

// Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_SessionStatistic
// RLFA Statistics categorized by session state
type Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_SessionStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session State. The type is RcmdShowLdpSessionState.
    SessionState interface{}

    // LDP Session Count. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}

    // Total Route Count. The type is interface{} with range: 0..4294967295.
    RouteCount interface{}

    // Total Path Count. The type is interface{} with range: 0..4294967295.
    PathCount interface{}

    // Remote Label Count. The type is interface{} with range: 0..4294967295.
    RemoteLabelCount interface{}

    // Protected Route Count. The type is interface{} with range: 0..4294967295.
    ProtectedRouteCount interface{}

    // Protected Path Count. The type is interface{} with range: 0..4294967295.
    ProtectedPathCount interface{}
}

func (sessionStatistic *Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_SessionStatistic) GetEntityData() *types.CommonEntityData {
    sessionStatistic.EntityData.YFilter = sessionStatistic.YFilter
    sessionStatistic.EntityData.YangName = "session-statistic"
    sessionStatistic.EntityData.BundleName = "cisco_ios_xr"
    sessionStatistic.EntityData.ParentYangName = "remote-lfa-summary"
    sessionStatistic.EntityData.SegmentPath = "session-statistic"
    sessionStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionStatistic.EntityData.Children = types.NewOrderedMap()
    sessionStatistic.EntityData.Leafs = types.NewOrderedMap()
    sessionStatistic.EntityData.Leafs.Append("session-state", types.YLeaf{"SessionState", sessionStatistic.SessionState})
    sessionStatistic.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", sessionStatistic.SessionCount})
    sessionStatistic.EntityData.Leafs.Append("route-count", types.YLeaf{"RouteCount", sessionStatistic.RouteCount})
    sessionStatistic.EntityData.Leafs.Append("path-count", types.YLeaf{"PathCount", sessionStatistic.PathCount})
    sessionStatistic.EntityData.Leafs.Append("remote-label-count", types.YLeaf{"RemoteLabelCount", sessionStatistic.RemoteLabelCount})
    sessionStatistic.EntityData.Leafs.Append("protected-route-count", types.YLeaf{"ProtectedRouteCount", sessionStatistic.ProtectedRouteCount})
    sessionStatistic.EntityData.Leafs.Append("protected-path-count", types.YLeaf{"ProtectedPathCount", sessionStatistic.ProtectedPathCount})

    sessionStatistic.EntityData.YListKeys = []string {}

    return &(sessionStatistic.EntityData)
}

// Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_RemoteNode
// Remote Node Information
type Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_RemoteNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote Node ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteNodeId interface{}

    // Label Space Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // Transport Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    TransportAddress interface{}

    // Inuse time of the Session. The type is string.
    InUseTime interface{}

    // Session State. The type is RcmdShowLdpSessionState.
    SessionState interface{}

    // Total Route Count. The type is interface{} with range: 0..4294967295.
    RouteCount interface{}

    // Total Path Count. The type is interface{} with range: 0..4294967295.
    PathCount interface{}

    // Remote Label Count. The type is interface{} with range: 0..4294967295.
    RemoteLabelCount interface{}

    // Protected Route Count. The type is interface{} with range: 0..4294967295.
    ProtectedRouteCount interface{}

    // Protected Path Count. The type is interface{} with range: 0..4294967295.
    ProtectedPathCount interface{}
}

func (remoteNode *Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_RemoteNode) GetEntityData() *types.CommonEntityData {
    remoteNode.EntityData.YFilter = remoteNode.YFilter
    remoteNode.EntityData.YangName = "remote-node"
    remoteNode.EntityData.BundleName = "cisco_ios_xr"
    remoteNode.EntityData.ParentYangName = "remote-lfa-summary"
    remoteNode.EntityData.SegmentPath = "remote-node"
    remoteNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNode.EntityData.Children = types.NewOrderedMap()
    remoteNode.EntityData.Leafs = types.NewOrderedMap()
    remoteNode.EntityData.Leafs.Append("remote-node-id", types.YLeaf{"RemoteNodeId", remoteNode.RemoteNodeId})
    remoteNode.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", remoteNode.LsrId})
    remoteNode.EntityData.Leafs.Append("transport-address", types.YLeaf{"TransportAddress", remoteNode.TransportAddress})
    remoteNode.EntityData.Leafs.Append("in-use-time", types.YLeaf{"InUseTime", remoteNode.InUseTime})
    remoteNode.EntityData.Leafs.Append("session-state", types.YLeaf{"SessionState", remoteNode.SessionState})
    remoteNode.EntityData.Leafs.Append("route-count", types.YLeaf{"RouteCount", remoteNode.RouteCount})
    remoteNode.EntityData.Leafs.Append("path-count", types.YLeaf{"PathCount", remoteNode.PathCount})
    remoteNode.EntityData.Leafs.Append("remote-label-count", types.YLeaf{"RemoteLabelCount", remoteNode.RemoteLabelCount})
    remoteNode.EntityData.Leafs.Append("protected-route-count", types.YLeaf{"ProtectedRouteCount", remoteNode.ProtectedRouteCount})
    remoteNode.EntityData.Leafs.Append("protected-path-count", types.YLeaf{"ProtectedPathCount", remoteNode.ProtectedPathCount})

    remoteNode.EntityData.YListKeys = []string {}

    return &(remoteNode.EntityData)
}

// Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_Logs
// Logs Information
type Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_Logs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event Time (eg: Apr 24 13:16:04.961). The type is string.
    LogTime interface{}

    // Label Coverage State. The type is RcmdShowLdpConvState.
    LabelCoverageState interface{}

    // Total Route Count. The type is interface{} with range: 0..4294967295.
    RouteCount interface{}

    // Remote Label Count. The type is interface{} with range: 0..4294967295.
    RemoteLabelCount interface{}
}

func (logs *Rcmd_Ldp_RemoteLfaSummaries_RemoteLfaSummary_Logs) GetEntityData() *types.CommonEntityData {
    logs.EntityData.YFilter = logs.YFilter
    logs.EntityData.YangName = "logs"
    logs.EntityData.BundleName = "cisco_ios_xr"
    logs.EntityData.ParentYangName = "remote-lfa-summary"
    logs.EntityData.SegmentPath = "logs"
    logs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logs.EntityData.Children = types.NewOrderedMap()
    logs.EntityData.Leafs = types.NewOrderedMap()
    logs.EntityData.Leafs.Append("log-time", types.YLeaf{"LogTime", logs.LogTime})
    logs.EntityData.Leafs.Append("label-coverage-state", types.YLeaf{"LabelCoverageState", logs.LabelCoverageState})
    logs.EntityData.Leafs.Append("route-count", types.YLeaf{"RouteCount", logs.RouteCount})
    logs.EntityData.Leafs.Append("remote-label-count", types.YLeaf{"RemoteLabelCount", logs.RemoteLabelCount})

    logs.EntityData.YListKeys = []string {}

    return &(logs.EntityData)
}

// Rcmd_Intf
// Interface data
type Rcmd_Intf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Events.
    Events Rcmd_Intf_Events
}

func (intf *Rcmd_Intf) GetEntityData() *types.CommonEntityData {
    intf.EntityData.YFilter = intf.YFilter
    intf.EntityData.YangName = "intf"
    intf.EntityData.BundleName = "cisco_ios_xr"
    intf.EntityData.ParentYangName = "rcmd"
    intf.EntityData.SegmentPath = "intf"
    intf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    intf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    intf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    intf.EntityData.Children = types.NewOrderedMap()
    intf.EntityData.Children.Append("events", types.YChild{"Events", &intf.Events})
    intf.EntityData.Leafs = types.NewOrderedMap()

    intf.EntityData.YListKeys = []string {}

    return &(intf.EntityData)
}

// Rcmd_Intf_Events
// Events
type Rcmd_Intf_Events struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Events. The type is slice of Rcmd_Intf_Events_Event.
    Event []*Rcmd_Intf_Events_Event
}

func (events *Rcmd_Intf_Events) GetEntityData() *types.CommonEntityData {
    events.EntityData.YFilter = events.YFilter
    events.EntityData.YangName = "events"
    events.EntityData.BundleName = "cisco_ios_xr"
    events.EntityData.ParentYangName = "intf"
    events.EntityData.SegmentPath = "events"
    events.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    events.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    events.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    events.EntityData.Children = types.NewOrderedMap()
    events.EntityData.Children.Append("event", types.YChild{"Event", nil})
    for i := range events.Event {
        events.EntityData.Children.Append(types.GetSegmentPath(events.Event[i]), types.YChild{"Event", events.Event[i]})
    }
    events.EntityData.Leafs = types.NewOrderedMap()

    events.EntityData.YListKeys = []string {}

    return &(events.EntityData)
}

// Rcmd_Intf_Events_Event
// Events
type Rcmd_Intf_Events_Event struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Event No. The type is interface{} with
    // range: 1..4294967295.
    EventNo interface{}

    // Sequence No. The type is interface{} with range: 0..4294967295.
    SequenceNo interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Component info. The type is RcmdShowCompId.
    Component interface{}

    // Event Info. The type is RcmdShowIntfEvent.
    EventType interface{}

    // Event Time. The type is string.
    EventTime interface{}

    // Primary Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrimaryAddress interface{}
}

func (event *Rcmd_Intf_Events_Event) GetEntityData() *types.CommonEntityData {
    event.EntityData.YFilter = event.YFilter
    event.EntityData.YangName = "event"
    event.EntityData.BundleName = "cisco_ios_xr"
    event.EntityData.ParentYangName = "events"
    event.EntityData.SegmentPath = "event" + types.AddKeyToken(event.EventNo, "event-no")
    event.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    event.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    event.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    event.EntityData.Children = types.NewOrderedMap()
    event.EntityData.Leafs = types.NewOrderedMap()
    event.EntityData.Leafs.Append("event-no", types.YLeaf{"EventNo", event.EventNo})
    event.EntityData.Leafs.Append("sequence-no", types.YLeaf{"SequenceNo", event.SequenceNo})
    event.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", event.InterfaceName})
    event.EntityData.Leafs.Append("component", types.YLeaf{"Component", event.Component})
    event.EntityData.Leafs.Append("event-type", types.YLeaf{"EventType", event.EventType})
    event.EntityData.Leafs.Append("event-time", types.YLeaf{"EventTime", event.EventTime})
    event.EntityData.Leafs.Append("primary-address", types.YLeaf{"PrimaryAddress", event.PrimaryAddress})

    event.EntityData.YListKeys = []string {"EventNo"}

    return &(event.EntityData)
}

// Rcmd_Process
// Process information
type Rcmd_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS Process Information.
    Isis Rcmd_Process_Isis

    // OSPF Process Information.
    Ospf Rcmd_Process_Ospf

    // LDP Process Information.
    Ldp Rcmd_Process_Ldp
}

func (process *Rcmd_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "rcmd"
    process.EntityData.SegmentPath = "process"
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Children.Append("isis", types.YChild{"Isis", &process.Isis})
    process.EntityData.Children.Append("ospf", types.YChild{"Ospf", &process.Ospf})
    process.EntityData.Children.Append("ldp", types.YChild{"Ldp", &process.Ldp})
    process.EntityData.Leafs = types.NewOrderedMap()

    process.EntityData.YListKeys = []string {}

    return &(process.EntityData)
}

// Rcmd_Process_Isis
// ISIS Process Information
type Rcmd_Process_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process Information. The type is slice of Rcmd_Process_Isis_Process.
    Process []*Rcmd_Process_Isis_Process
}

func (isis *Rcmd_Process_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "process"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = types.NewOrderedMap()
    isis.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range isis.Process {
        isis.EntityData.Children.Append(types.GetSegmentPath(isis.Process[i]), types.YChild{"Process", isis.Process[i]})
    }
    isis.EntityData.Leafs = types.NewOrderedMap()

    isis.EntityData.YListKeys = []string {}

    return &(isis.EntityData)
}

// Rcmd_Process_Isis_Process
// Process Information
type Rcmd_Process_Isis_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol id. The type is RcmdProtocolId.
    ProtocolId interface{}

    // Process Name. The type is string.
    ProcessName interface{}

    // Instance/VRF Name. The type is slice of
    // Rcmd_Process_Isis_Process_InstanceName.
    InstanceName []*Rcmd_Process_Isis_Process_InstanceName
}

func (process *Rcmd_Process_Isis_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "isis"
    process.EntityData.SegmentPath = "process"
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Children.Append("instance-name", types.YChild{"InstanceName", nil})
    for i := range process.InstanceName {
        process.EntityData.Children.Append(types.GetSegmentPath(process.InstanceName[i]), types.YChild{"InstanceName", process.InstanceName[i]})
    }
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("protocol-id", types.YLeaf{"ProtocolId", process.ProtocolId})
    process.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", process.ProcessName})

    process.EntityData.YListKeys = []string {}

    return &(process.EntityData)
}

// Rcmd_Process_Isis_Process_InstanceName
// Instance/VRF Name
type Rcmd_Process_Isis_Process_InstanceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance Name. The type is string.
    Name interface{}

    // Last Updated Time. The type is string.
    LastUpdateTime interface{}

    // Total spf nos. The type is interface{} with range: 0..4294967295.
    TotalSpfNos interface{}

    // Route change spf nos. The type is interface{} with range: 0..4294967295.
    RouteChangeSpfNos interface{}

    // No Route change spf nos. The type is interface{} with range: 0..4294967295.
    NoRouteChangeSpfNos interface{}

    // Not Interested SPF nos. The type is interface{} with range: 0..4294967295.
    NotInterestedSpfNos interface{}

    // LSP Regen Count. The type is interface{} with range: 0..4294967295.
    LspRegenerationCount interface{}

    // Last Serial. The type is interface{} with range: 0..4294967295.
    LspRegenerationSerial interface{}

    // Archive SPF event. The type is interface{} with range: 0..4294967295.
    ArchSpfEvent interface{}

    // Archive Lsp regen. The type is interface{} with range: 0..4294967295.
    ArchLspRegeneration interface{}

    // Instance Information. The type is slice of
    // Rcmd_Process_Isis_Process_InstanceName_Instance.
    Instance []*Rcmd_Process_Isis_Process_InstanceName_Instance
}

func (instanceName *Rcmd_Process_Isis_Process_InstanceName) GetEntityData() *types.CommonEntityData {
    instanceName.EntityData.YFilter = instanceName.YFilter
    instanceName.EntityData.YangName = "instance-name"
    instanceName.EntityData.BundleName = "cisco_ios_xr"
    instanceName.EntityData.ParentYangName = "process"
    instanceName.EntityData.SegmentPath = "instance-name"
    instanceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instanceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instanceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instanceName.EntityData.Children = types.NewOrderedMap()
    instanceName.EntityData.Children.Append("instance", types.YChild{"Instance", nil})
    for i := range instanceName.Instance {
        instanceName.EntityData.Children.Append(types.GetSegmentPath(instanceName.Instance[i]), types.YChild{"Instance", instanceName.Instance[i]})
    }
    instanceName.EntityData.Leafs = types.NewOrderedMap()
    instanceName.EntityData.Leafs.Append("name", types.YLeaf{"Name", instanceName.Name})
    instanceName.EntityData.Leafs.Append("last-update-time", types.YLeaf{"LastUpdateTime", instanceName.LastUpdateTime})
    instanceName.EntityData.Leafs.Append("total-spf-nos", types.YLeaf{"TotalSpfNos", instanceName.TotalSpfNos})
    instanceName.EntityData.Leafs.Append("route-change-spf-nos", types.YLeaf{"RouteChangeSpfNos", instanceName.RouteChangeSpfNos})
    instanceName.EntityData.Leafs.Append("no-route-change-spf-nos", types.YLeaf{"NoRouteChangeSpfNos", instanceName.NoRouteChangeSpfNos})
    instanceName.EntityData.Leafs.Append("not-interested-spf-nos", types.YLeaf{"NotInterestedSpfNos", instanceName.NotInterestedSpfNos})
    instanceName.EntityData.Leafs.Append("lsp-regeneration-count", types.YLeaf{"LspRegenerationCount", instanceName.LspRegenerationCount})
    instanceName.EntityData.Leafs.Append("lsp-regeneration-serial", types.YLeaf{"LspRegenerationSerial", instanceName.LspRegenerationSerial})
    instanceName.EntityData.Leafs.Append("arch-spf-event", types.YLeaf{"ArchSpfEvent", instanceName.ArchSpfEvent})
    instanceName.EntityData.Leafs.Append("arch-lsp-regeneration", types.YLeaf{"ArchLspRegeneration", instanceName.ArchLspRegeneration})

    instanceName.EntityData.YListKeys = []string {}

    return &(instanceName.EntityData)
}

// Rcmd_Process_Isis_Process_InstanceName_Instance
// Instance Information
type Rcmd_Process_Isis_Process_InstanceName_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance Id. The type is interface{} with range: 0..4294967295.
    InstanceId interface{}

    // Instance State. The type is RcmdShowInstState.
    InstanceState interface{}

    // Instance Deleted. The type is RcmdBoolYesNo.
    InstanceDeleted interface{}

    // Forward Referenced. The type is RcmdBoolYesNo.
    FwdReferenced interface{}

    // Last Updated Time. The type is string.
    LastUpdateTime interface{}

    // Node Id. The type is interface{} with range: 0..4294967295.
    NodeId interface{}

    // SPF Offset. The type is interface{} with range: 0..4294967295.
    SpfOffset interface{}

    // Total spf nos. The type is interface{} with range: 0..4294967295.
    TotalSpfNos interface{}

    // spf run can be archived. The type is interface{} with range: 0..4294967295.
    ArchSpfRun interface{}

    // Route change spf nos. The type is interface{} with range: 0..4294967295.
    RouteChangeSpfNos interface{}

    // No Route change spf nos. The type is interface{} with range: 0..4294967295.
    NoRouteChangeSpfNos interface{}

    // Not Interested SPF nos. The type is interface{} with range: 0..4294967295.
    NotInterestedSpfNos interface{}

    // Total spt nos. The type is interface{} with range: 0..4294967295.
    TotalSptNos interface{}
}

func (instance *Rcmd_Process_Isis_Process_InstanceName_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "instance-name"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", instance.InstanceId})
    instance.EntityData.Leafs.Append("instance-state", types.YLeaf{"InstanceState", instance.InstanceState})
    instance.EntityData.Leafs.Append("instance-deleted", types.YLeaf{"InstanceDeleted", instance.InstanceDeleted})
    instance.EntityData.Leafs.Append("fwd-referenced", types.YLeaf{"FwdReferenced", instance.FwdReferenced})
    instance.EntityData.Leafs.Append("last-update-time", types.YLeaf{"LastUpdateTime", instance.LastUpdateTime})
    instance.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", instance.NodeId})
    instance.EntityData.Leafs.Append("spf-offset", types.YLeaf{"SpfOffset", instance.SpfOffset})
    instance.EntityData.Leafs.Append("total-spf-nos", types.YLeaf{"TotalSpfNos", instance.TotalSpfNos})
    instance.EntityData.Leafs.Append("arch-spf-run", types.YLeaf{"ArchSpfRun", instance.ArchSpfRun})
    instance.EntityData.Leafs.Append("route-change-spf-nos", types.YLeaf{"RouteChangeSpfNos", instance.RouteChangeSpfNos})
    instance.EntityData.Leafs.Append("no-route-change-spf-nos", types.YLeaf{"NoRouteChangeSpfNos", instance.NoRouteChangeSpfNos})
    instance.EntityData.Leafs.Append("not-interested-spf-nos", types.YLeaf{"NotInterestedSpfNos", instance.NotInterestedSpfNos})
    instance.EntityData.Leafs.Append("total-spt-nos", types.YLeaf{"TotalSptNos", instance.TotalSptNos})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// Rcmd_Process_Ospf
// OSPF Process Information
type Rcmd_Process_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process Information. The type is slice of Rcmd_Process_Ospf_Process.
    Process []*Rcmd_Process_Ospf_Process
}

func (ospf *Rcmd_Process_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "process"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = types.NewOrderedMap()
    ospf.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range ospf.Process {
        ospf.EntityData.Children.Append(types.GetSegmentPath(ospf.Process[i]), types.YChild{"Process", ospf.Process[i]})
    }
    ospf.EntityData.Leafs = types.NewOrderedMap()

    ospf.EntityData.YListKeys = []string {}

    return &(ospf.EntityData)
}

// Rcmd_Process_Ospf_Process
// Process Information
type Rcmd_Process_Ospf_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol id. The type is RcmdProtocolId.
    ProtocolId interface{}

    // Process Name. The type is string.
    ProcessName interface{}

    // Instance/VRF Name. The type is slice of
    // Rcmd_Process_Ospf_Process_InstanceName.
    InstanceName []*Rcmd_Process_Ospf_Process_InstanceName
}

func (process *Rcmd_Process_Ospf_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "ospf"
    process.EntityData.SegmentPath = "process"
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Children.Append("instance-name", types.YChild{"InstanceName", nil})
    for i := range process.InstanceName {
        process.EntityData.Children.Append(types.GetSegmentPath(process.InstanceName[i]), types.YChild{"InstanceName", process.InstanceName[i]})
    }
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("protocol-id", types.YLeaf{"ProtocolId", process.ProtocolId})
    process.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", process.ProcessName})

    process.EntityData.YListKeys = []string {}

    return &(process.EntityData)
}

// Rcmd_Process_Ospf_Process_InstanceName
// Instance/VRF Name
type Rcmd_Process_Ospf_Process_InstanceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance Name. The type is string.
    Name interface{}

    // Last Updated Time. The type is string.
    LastUpdateTime interface{}

    // Total spf nos. The type is interface{} with range: 0..4294967295.
    TotalSpfNos interface{}

    // Route change spf nos. The type is interface{} with range: 0..4294967295.
    RouteChangeSpfNos interface{}

    // No Route change spf nos. The type is interface{} with range: 0..4294967295.
    NoRouteChangeSpfNos interface{}

    // Not Interested SPF nos. The type is interface{} with range: 0..4294967295.
    NotInterestedSpfNos interface{}

    // LSP Regen Count. The type is interface{} with range: 0..4294967295.
    LspRegenerationCount interface{}

    // Last Serial. The type is interface{} with range: 0..4294967295.
    LspRegenerationSerial interface{}

    // Archive SPF event. The type is interface{} with range: 0..4294967295.
    ArchSpfEvent interface{}

    // Archive Lsp regen. The type is interface{} with range: 0..4294967295.
    ArchLspRegeneration interface{}

    // Instance Information. The type is slice of
    // Rcmd_Process_Ospf_Process_InstanceName_Instance.
    Instance []*Rcmd_Process_Ospf_Process_InstanceName_Instance
}

func (instanceName *Rcmd_Process_Ospf_Process_InstanceName) GetEntityData() *types.CommonEntityData {
    instanceName.EntityData.YFilter = instanceName.YFilter
    instanceName.EntityData.YangName = "instance-name"
    instanceName.EntityData.BundleName = "cisco_ios_xr"
    instanceName.EntityData.ParentYangName = "process"
    instanceName.EntityData.SegmentPath = "instance-name"
    instanceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instanceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instanceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instanceName.EntityData.Children = types.NewOrderedMap()
    instanceName.EntityData.Children.Append("instance", types.YChild{"Instance", nil})
    for i := range instanceName.Instance {
        instanceName.EntityData.Children.Append(types.GetSegmentPath(instanceName.Instance[i]), types.YChild{"Instance", instanceName.Instance[i]})
    }
    instanceName.EntityData.Leafs = types.NewOrderedMap()
    instanceName.EntityData.Leafs.Append("name", types.YLeaf{"Name", instanceName.Name})
    instanceName.EntityData.Leafs.Append("last-update-time", types.YLeaf{"LastUpdateTime", instanceName.LastUpdateTime})
    instanceName.EntityData.Leafs.Append("total-spf-nos", types.YLeaf{"TotalSpfNos", instanceName.TotalSpfNos})
    instanceName.EntityData.Leafs.Append("route-change-spf-nos", types.YLeaf{"RouteChangeSpfNos", instanceName.RouteChangeSpfNos})
    instanceName.EntityData.Leafs.Append("no-route-change-spf-nos", types.YLeaf{"NoRouteChangeSpfNos", instanceName.NoRouteChangeSpfNos})
    instanceName.EntityData.Leafs.Append("not-interested-spf-nos", types.YLeaf{"NotInterestedSpfNos", instanceName.NotInterestedSpfNos})
    instanceName.EntityData.Leafs.Append("lsp-regeneration-count", types.YLeaf{"LspRegenerationCount", instanceName.LspRegenerationCount})
    instanceName.EntityData.Leafs.Append("lsp-regeneration-serial", types.YLeaf{"LspRegenerationSerial", instanceName.LspRegenerationSerial})
    instanceName.EntityData.Leafs.Append("arch-spf-event", types.YLeaf{"ArchSpfEvent", instanceName.ArchSpfEvent})
    instanceName.EntityData.Leafs.Append("arch-lsp-regeneration", types.YLeaf{"ArchLspRegeneration", instanceName.ArchLspRegeneration})

    instanceName.EntityData.YListKeys = []string {}

    return &(instanceName.EntityData)
}

// Rcmd_Process_Ospf_Process_InstanceName_Instance
// Instance Information
type Rcmd_Process_Ospf_Process_InstanceName_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance Id. The type is interface{} with range: 0..4294967295.
    InstanceId interface{}

    // Instance State. The type is RcmdShowInstState.
    InstanceState interface{}

    // Instance Deleted. The type is RcmdBoolYesNo.
    InstanceDeleted interface{}

    // Forward Referenced. The type is RcmdBoolYesNo.
    FwdReferenced interface{}

    // Last Updated Time. The type is string.
    LastUpdateTime interface{}

    // Node Id. The type is interface{} with range: 0..4294967295.
    NodeId interface{}

    // SPF Offset. The type is interface{} with range: 0..4294967295.
    SpfOffset interface{}

    // Total spf nos. The type is interface{} with range: 0..4294967295.
    TotalSpfNos interface{}

    // spf run can be archived. The type is interface{} with range: 0..4294967295.
    ArchSpfRun interface{}

    // Route change spf nos. The type is interface{} with range: 0..4294967295.
    RouteChangeSpfNos interface{}

    // No Route change spf nos. The type is interface{} with range: 0..4294967295.
    NoRouteChangeSpfNos interface{}

    // Not Interested SPF nos. The type is interface{} with range: 0..4294967295.
    NotInterestedSpfNos interface{}

    // Total spt nos. The type is interface{} with range: 0..4294967295.
    TotalSptNos interface{}
}

func (instance *Rcmd_Process_Ospf_Process_InstanceName_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "instance-name"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", instance.InstanceId})
    instance.EntityData.Leafs.Append("instance-state", types.YLeaf{"InstanceState", instance.InstanceState})
    instance.EntityData.Leafs.Append("instance-deleted", types.YLeaf{"InstanceDeleted", instance.InstanceDeleted})
    instance.EntityData.Leafs.Append("fwd-referenced", types.YLeaf{"FwdReferenced", instance.FwdReferenced})
    instance.EntityData.Leafs.Append("last-update-time", types.YLeaf{"LastUpdateTime", instance.LastUpdateTime})
    instance.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", instance.NodeId})
    instance.EntityData.Leafs.Append("spf-offset", types.YLeaf{"SpfOffset", instance.SpfOffset})
    instance.EntityData.Leafs.Append("total-spf-nos", types.YLeaf{"TotalSpfNos", instance.TotalSpfNos})
    instance.EntityData.Leafs.Append("arch-spf-run", types.YLeaf{"ArchSpfRun", instance.ArchSpfRun})
    instance.EntityData.Leafs.Append("route-change-spf-nos", types.YLeaf{"RouteChangeSpfNos", instance.RouteChangeSpfNos})
    instance.EntityData.Leafs.Append("no-route-change-spf-nos", types.YLeaf{"NoRouteChangeSpfNos", instance.NoRouteChangeSpfNos})
    instance.EntityData.Leafs.Append("not-interested-spf-nos", types.YLeaf{"NotInterestedSpfNos", instance.NotInterestedSpfNos})
    instance.EntityData.Leafs.Append("total-spt-nos", types.YLeaf{"TotalSptNos", instance.TotalSptNos})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// Rcmd_Process_Ldp
// LDP Process Information
type Rcmd_Process_Ldp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process Information. The type is slice of Rcmd_Process_Ldp_Process.
    Process []*Rcmd_Process_Ldp_Process
}

func (ldp *Rcmd_Process_Ldp) GetEntityData() *types.CommonEntityData {
    ldp.EntityData.YFilter = ldp.YFilter
    ldp.EntityData.YangName = "ldp"
    ldp.EntityData.BundleName = "cisco_ios_xr"
    ldp.EntityData.ParentYangName = "process"
    ldp.EntityData.SegmentPath = "ldp"
    ldp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldp.EntityData.Children = types.NewOrderedMap()
    ldp.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range ldp.Process {
        ldp.EntityData.Children.Append(types.GetSegmentPath(ldp.Process[i]), types.YChild{"Process", ldp.Process[i]})
    }
    ldp.EntityData.Leafs = types.NewOrderedMap()

    ldp.EntityData.YListKeys = []string {}

    return &(ldp.EntityData)
}

// Rcmd_Process_Ldp_Process
// Process Information
type Rcmd_Process_Ldp_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol id. The type is RcmdProtocolId.
    ProtocolId interface{}

    // Process Name. The type is string.
    ProcessName interface{}

    // Instance/VRF Name. The type is slice of
    // Rcmd_Process_Ldp_Process_InstanceName.
    InstanceName []*Rcmd_Process_Ldp_Process_InstanceName
}

func (process *Rcmd_Process_Ldp_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "ldp"
    process.EntityData.SegmentPath = "process"
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Children.Append("instance-name", types.YChild{"InstanceName", nil})
    for i := range process.InstanceName {
        process.EntityData.Children.Append(types.GetSegmentPath(process.InstanceName[i]), types.YChild{"InstanceName", process.InstanceName[i]})
    }
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("protocol-id", types.YLeaf{"ProtocolId", process.ProtocolId})
    process.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", process.ProcessName})

    process.EntityData.YListKeys = []string {}

    return &(process.EntityData)
}

// Rcmd_Process_Ldp_Process_InstanceName
// Instance/VRF Name
type Rcmd_Process_Ldp_Process_InstanceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance Name. The type is string.
    Name interface{}

    // Last Updated Time. The type is string.
    LastUpdateTime interface{}

    // Total spf nos. The type is interface{} with range: 0..4294967295.
    TotalSpfNos interface{}

    // Route change spf nos. The type is interface{} with range: 0..4294967295.
    RouteChangeSpfNos interface{}

    // No Route change spf nos. The type is interface{} with range: 0..4294967295.
    NoRouteChangeSpfNos interface{}

    // Not Interested SPF nos. The type is interface{} with range: 0..4294967295.
    NotInterestedSpfNos interface{}

    // LSP Regen Count. The type is interface{} with range: 0..4294967295.
    LspRegenerationCount interface{}

    // Last Serial. The type is interface{} with range: 0..4294967295.
    LspRegenerationSerial interface{}

    // Archive SPF event. The type is interface{} with range: 0..4294967295.
    ArchSpfEvent interface{}

    // Archive Lsp regen. The type is interface{} with range: 0..4294967295.
    ArchLspRegeneration interface{}

    // Instance Information. The type is slice of
    // Rcmd_Process_Ldp_Process_InstanceName_Instance.
    Instance []*Rcmd_Process_Ldp_Process_InstanceName_Instance
}

func (instanceName *Rcmd_Process_Ldp_Process_InstanceName) GetEntityData() *types.CommonEntityData {
    instanceName.EntityData.YFilter = instanceName.YFilter
    instanceName.EntityData.YangName = "instance-name"
    instanceName.EntityData.BundleName = "cisco_ios_xr"
    instanceName.EntityData.ParentYangName = "process"
    instanceName.EntityData.SegmentPath = "instance-name"
    instanceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instanceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instanceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instanceName.EntityData.Children = types.NewOrderedMap()
    instanceName.EntityData.Children.Append("instance", types.YChild{"Instance", nil})
    for i := range instanceName.Instance {
        instanceName.EntityData.Children.Append(types.GetSegmentPath(instanceName.Instance[i]), types.YChild{"Instance", instanceName.Instance[i]})
    }
    instanceName.EntityData.Leafs = types.NewOrderedMap()
    instanceName.EntityData.Leafs.Append("name", types.YLeaf{"Name", instanceName.Name})
    instanceName.EntityData.Leafs.Append("last-update-time", types.YLeaf{"LastUpdateTime", instanceName.LastUpdateTime})
    instanceName.EntityData.Leafs.Append("total-spf-nos", types.YLeaf{"TotalSpfNos", instanceName.TotalSpfNos})
    instanceName.EntityData.Leafs.Append("route-change-spf-nos", types.YLeaf{"RouteChangeSpfNos", instanceName.RouteChangeSpfNos})
    instanceName.EntityData.Leafs.Append("no-route-change-spf-nos", types.YLeaf{"NoRouteChangeSpfNos", instanceName.NoRouteChangeSpfNos})
    instanceName.EntityData.Leafs.Append("not-interested-spf-nos", types.YLeaf{"NotInterestedSpfNos", instanceName.NotInterestedSpfNos})
    instanceName.EntityData.Leafs.Append("lsp-regeneration-count", types.YLeaf{"LspRegenerationCount", instanceName.LspRegenerationCount})
    instanceName.EntityData.Leafs.Append("lsp-regeneration-serial", types.YLeaf{"LspRegenerationSerial", instanceName.LspRegenerationSerial})
    instanceName.EntityData.Leafs.Append("arch-spf-event", types.YLeaf{"ArchSpfEvent", instanceName.ArchSpfEvent})
    instanceName.EntityData.Leafs.Append("arch-lsp-regeneration", types.YLeaf{"ArchLspRegeneration", instanceName.ArchLspRegeneration})

    instanceName.EntityData.YListKeys = []string {}

    return &(instanceName.EntityData)
}

// Rcmd_Process_Ldp_Process_InstanceName_Instance
// Instance Information
type Rcmd_Process_Ldp_Process_InstanceName_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance Id. The type is interface{} with range: 0..4294967295.
    InstanceId interface{}

    // Instance State. The type is RcmdShowInstState.
    InstanceState interface{}

    // Instance Deleted. The type is RcmdBoolYesNo.
    InstanceDeleted interface{}

    // Forward Referenced. The type is RcmdBoolYesNo.
    FwdReferenced interface{}

    // Last Updated Time. The type is string.
    LastUpdateTime interface{}

    // Node Id. The type is interface{} with range: 0..4294967295.
    NodeId interface{}

    // SPF Offset. The type is interface{} with range: 0..4294967295.
    SpfOffset interface{}

    // Total spf nos. The type is interface{} with range: 0..4294967295.
    TotalSpfNos interface{}

    // spf run can be archived. The type is interface{} with range: 0..4294967295.
    ArchSpfRun interface{}

    // Route change spf nos. The type is interface{} with range: 0..4294967295.
    RouteChangeSpfNos interface{}

    // No Route change spf nos. The type is interface{} with range: 0..4294967295.
    NoRouteChangeSpfNos interface{}

    // Not Interested SPF nos. The type is interface{} with range: 0..4294967295.
    NotInterestedSpfNos interface{}

    // Total spt nos. The type is interface{} with range: 0..4294967295.
    TotalSptNos interface{}
}

func (instance *Rcmd_Process_Ldp_Process_InstanceName_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "instance-name"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", instance.InstanceId})
    instance.EntityData.Leafs.Append("instance-state", types.YLeaf{"InstanceState", instance.InstanceState})
    instance.EntityData.Leafs.Append("instance-deleted", types.YLeaf{"InstanceDeleted", instance.InstanceDeleted})
    instance.EntityData.Leafs.Append("fwd-referenced", types.YLeaf{"FwdReferenced", instance.FwdReferenced})
    instance.EntityData.Leafs.Append("last-update-time", types.YLeaf{"LastUpdateTime", instance.LastUpdateTime})
    instance.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", instance.NodeId})
    instance.EntityData.Leafs.Append("spf-offset", types.YLeaf{"SpfOffset", instance.SpfOffset})
    instance.EntityData.Leafs.Append("total-spf-nos", types.YLeaf{"TotalSpfNos", instance.TotalSpfNos})
    instance.EntityData.Leafs.Append("arch-spf-run", types.YLeaf{"ArchSpfRun", instance.ArchSpfRun})
    instance.EntityData.Leafs.Append("route-change-spf-nos", types.YLeaf{"RouteChangeSpfNos", instance.RouteChangeSpfNos})
    instance.EntityData.Leafs.Append("no-route-change-spf-nos", types.YLeaf{"NoRouteChangeSpfNos", instance.NoRouteChangeSpfNos})
    instance.EntityData.Leafs.Append("not-interested-spf-nos", types.YLeaf{"NotInterestedSpfNos", instance.NotInterestedSpfNos})
    instance.EntityData.Leafs.Append("total-spt-nos", types.YLeaf{"TotalSptNos", instance.TotalSptNos})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

