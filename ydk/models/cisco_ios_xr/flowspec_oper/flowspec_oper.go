// This module contains a collection of YANG definitions
// for Cisco IOS-XR flowspec package operational data.
// 
// This module contains definitions
// for the following management objects:
//   flow-spec: FlowSpec information
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package flowspec_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package flowspec_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-flowspec-oper flow-spec}", reflect.TypeOf(FlowSpec{}))
    ydk.RegisterEntity("Cisco-IOS-XR-flowspec-oper:flow-spec", reflect.TypeOf(FlowSpec{}))
}

// FsClient represents Fs client
type FsClient string

const (
    // FlowSpec BGP Client
    FsClient_bgp FsClient = "bgp"

    // FlowSpec OneP Client
    FsClient_one_pk FsClient = "one-pk"

    // FlowSpec Policy Client
    FsClient_policy FsClient = "policy"

    // FlowSpec HA Client
    FsClient_ha FsClient = "ha"

    // FlowSpec Test Client
    FsClient_test FsClient = "test"
)

// FsMgrClientState represents FlowSpec manager client state
type FsMgrClientState string

const (
    // Dormant
    FsMgrClientState_dormant FsMgrClientState = "dormant"

    // Connected
    FsMgrClientState_connected FsMgrClientState = "connected"

    // Replay
    FsMgrClientState_replay FsMgrClientState = "replay"

    // Unconfigured
    FsMgrClientState_unconfigured FsMgrClientState = "unconfigured"
)

// FlowSpec
// FlowSpec information
type FlowSpec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Client.
    Clients FlowSpec_Clients

    // FlowSpec summary.
    Summary FlowSpec_Summary

    // Table of VRF.
    Vrfs FlowSpec_Vrfs
}

func (flowSpec *FlowSpec) GetEntityData() *types.CommonEntityData {
    flowSpec.EntityData.YFilter = flowSpec.YFilter
    flowSpec.EntityData.YangName = "flow-spec"
    flowSpec.EntityData.BundleName = "cisco_ios_xr"
    flowSpec.EntityData.ParentYangName = "Cisco-IOS-XR-flowspec-oper"
    flowSpec.EntityData.SegmentPath = "Cisco-IOS-XR-flowspec-oper:flow-spec"
    flowSpec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowSpec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowSpec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowSpec.EntityData.Children = types.NewOrderedMap()
    flowSpec.EntityData.Children.Append("clients", types.YChild{"Clients", &flowSpec.Clients})
    flowSpec.EntityData.Children.Append("summary", types.YChild{"Summary", &flowSpec.Summary})
    flowSpec.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &flowSpec.Vrfs})
    flowSpec.EntityData.Leafs = types.NewOrderedMap()

    flowSpec.EntityData.YListKeys = []string {}

    return &(flowSpec.EntityData)
}

// FlowSpec_Clients
// Table of Client
type FlowSpec_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FlowSpec client information. The type is slice of FlowSpec_Clients_Client.
    Client []*FlowSpec_Clients_Client
}

func (clients *FlowSpec_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "flow-spec"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Children.Append("client", types.YChild{"Client", nil})
    for i := range clients.Client {
        clients.EntityData.Children.Append(types.GetSegmentPath(clients.Client[i]), types.YChild{"Client", clients.Client[i]})
    }
    clients.EntityData.Leafs = types.NewOrderedMap()

    clients.EntityData.YListKeys = []string {}

    return &(clients.EntityData)
}

// FlowSpec_Clients_Client
// FlowSpec client information
type FlowSpec_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FlowSpec Client Type. The type is FsClient.
    ClientType interface{}

    // FlowSpec client ID. The type is interface{} with range: 0..4294967295.
    ClientId interface{}

    // State of FS client. The type is FsMgrClientState.
    ClientState interface{}

    // Number of client flows. The type is interface{} with range: 0..4294967295.
    ClientFlows interface{}
}

func (client *FlowSpec_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("client-type", types.YLeaf{"ClientType", client.ClientType})
    client.EntityData.Leafs.Append("client-id", types.YLeaf{"ClientId", client.ClientId})
    client.EntityData.Leafs.Append("client-state", types.YLeaf{"ClientState", client.ClientState})
    client.EntityData.Leafs.Append("client-flows", types.YLeaf{"ClientFlows", client.ClientFlows})

    client.EntityData.YListKeys = []string {}

    return &(client.EntityData)
}

// FlowSpec_Summary
// FlowSpec summary
type FlowSpec_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total VRF+AFI tables. The type is interface{} with range: 0..4294967295.
    VrfafiTables interface{}

    // Total flows. The type is interface{} with range: 0..4294967295.
    Flows interface{}
}

func (summary *FlowSpec_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "flow-spec"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("vrfafi-tables", types.YLeaf{"VrfafiTables", summary.VrfafiTables})
    summary.EntityData.Leafs.Append("flows", types.YLeaf{"Flows", summary.Flows})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// FlowSpec_Vrfs
// Table of VRF
type FlowSpec_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF information. The type is slice of FlowSpec_Vrfs_Vrf.
    Vrf []*FlowSpec_Vrfs_Vrf
}

func (vrfs *FlowSpec_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "flow-spec"
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

// FlowSpec_Vrfs_Vrf
// VRF information
type FlowSpec_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with length: 1..32.
    VrfName interface{}

    // Table of AF.
    Afs FlowSpec_Vrfs_Vrf_Afs
}

func (vrf *FlowSpec_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
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

// FlowSpec_Vrfs_Vrf_Afs
// Table of AF
type FlowSpec_Vrfs_Vrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI Type (IPv4/IPv6). The type is slice of FlowSpec_Vrfs_Vrf_Afs_Af.
    Af []*FlowSpec_Vrfs_Vrf_Afs_Af
}

func (afs *FlowSpec_Vrfs_Vrf_Afs) GetEntityData() *types.CommonEntityData {
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

// FlowSpec_Vrfs_Vrf_Afs_Af
// AFI Type (IPv4/IPv6)
type FlowSpec_Vrfs_Vrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set string. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    AfName interface{}

    // FlowSpec summary for VRF+AFI tables.
    TableSummary FlowSpec_Vrfs_Vrf_Afs_Af_TableSummary

    // Table of NLRI.
    Nlris FlowSpec_Vrfs_Vrf_Afs_Af_Nlris

    // Table of Flow.
    Flows FlowSpec_Vrfs_Vrf_Afs_Af_Flows
}

func (af *FlowSpec_Vrfs_Vrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("table-summary", types.YChild{"TableSummary", &af.TableSummary})
    af.EntityData.Children.Append("nlris", types.YChild{"Nlris", &af.Nlris})
    af.EntityData.Children.Append("flows", types.YChild{"Flows", &af.Flows})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_TableSummary
// FlowSpec summary for VRF+AFI tables
type FlowSpec_Vrfs_Vrf_Afs_Af_TableSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of flows. The type is interface{} with range: 0..4294967295.
    TotalFlows interface{}

    // Total number of attached service policies. The type is interface{} with
    // range: 0..4294967295.
    ServicePolicies interface{}

    // Local installation of flowspec rules. The type is bool.
    LocalInstallEnabled interface{}
}

func (tableSummary *FlowSpec_Vrfs_Vrf_Afs_Af_TableSummary) GetEntityData() *types.CommonEntityData {
    tableSummary.EntityData.YFilter = tableSummary.YFilter
    tableSummary.EntityData.YangName = "table-summary"
    tableSummary.EntityData.BundleName = "cisco_ios_xr"
    tableSummary.EntityData.ParentYangName = "af"
    tableSummary.EntityData.SegmentPath = "table-summary"
    tableSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tableSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tableSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tableSummary.EntityData.Children = types.NewOrderedMap()
    tableSummary.EntityData.Leafs = types.NewOrderedMap()
    tableSummary.EntityData.Leafs.Append("total-flows", types.YLeaf{"TotalFlows", tableSummary.TotalFlows})
    tableSummary.EntityData.Leafs.Append("service-policies", types.YLeaf{"ServicePolicies", tableSummary.ServicePolicies})
    tableSummary.EntityData.Leafs.Append("local-install-enabled", types.YLeaf{"LocalInstallEnabled", tableSummary.LocalInstallEnabled})

    tableSummary.EntityData.YListKeys = []string {}

    return &(tableSummary.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_Nlris
// Table of NLRI
type FlowSpec_Vrfs_Vrf_Afs_Af_Nlris struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NLRI information (hexdump). The type is slice of
    // FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri.
    Nlri []*FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri
}

func (nlris *FlowSpec_Vrfs_Vrf_Afs_Af_Nlris) GetEntityData() *types.CommonEntityData {
    nlris.EntityData.YFilter = nlris.YFilter
    nlris.EntityData.YangName = "nlris"
    nlris.EntityData.BundleName = "cisco_ios_xr"
    nlris.EntityData.ParentYangName = "af"
    nlris.EntityData.SegmentPath = "nlris"
    nlris.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nlris.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nlris.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nlris.EntityData.Children = types.NewOrderedMap()
    nlris.EntityData.Children.Append("nlri", types.YChild{"Nlri", nil})
    for i := range nlris.Nlri {
        nlris.EntityData.Children.Append(types.GetSegmentPath(nlris.Nlri[i]), types.YChild{"Nlri", nlris.Nlri[i]})
    }
    nlris.EntityData.Leafs = types.NewOrderedMap()

    nlris.EntityData.YListKeys = []string {}

    return &(nlris.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri
// NLRI information (hexdump)
type FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Enter NLRI hex string. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    NlriBytes interface{}

    // Flow statistics.
    FlowStatistics FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri_FlowStatistics
}

func (nlri *FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri) GetEntityData() *types.CommonEntityData {
    nlri.EntityData.YFilter = nlri.YFilter
    nlri.EntityData.YangName = "nlri"
    nlri.EntityData.BundleName = "cisco_ios_xr"
    nlri.EntityData.ParentYangName = "nlris"
    nlri.EntityData.SegmentPath = "nlri" + types.AddKeyToken(nlri.NlriBytes, "nlri-bytes")
    nlri.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nlri.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nlri.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nlri.EntityData.Children = types.NewOrderedMap()
    nlri.EntityData.Children.Append("flow-statistics", types.YChild{"FlowStatistics", &nlri.FlowStatistics})
    nlri.EntityData.Leafs = types.NewOrderedMap()
    nlri.EntityData.Leafs.Append("nlri-bytes", types.YLeaf{"NlriBytes", nlri.NlriBytes})

    nlri.EntityData.YListKeys = []string {"NlriBytes"}

    return &(nlri.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri_FlowStatistics
// Flow statistics
type FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri_FlowStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Classified statistics.
    Classified FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri_FlowStatistics_Classified

    // Drop statistics.
    Dropped FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri_FlowStatistics_Dropped
}

func (flowStatistics *FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri_FlowStatistics) GetEntityData() *types.CommonEntityData {
    flowStatistics.EntityData.YFilter = flowStatistics.YFilter
    flowStatistics.EntityData.YangName = "flow-statistics"
    flowStatistics.EntityData.BundleName = "cisco_ios_xr"
    flowStatistics.EntityData.ParentYangName = "nlri"
    flowStatistics.EntityData.SegmentPath = "flow-statistics"
    flowStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowStatistics.EntityData.Children = types.NewOrderedMap()
    flowStatistics.EntityData.Children.Append("classified", types.YChild{"Classified", &flowStatistics.Classified})
    flowStatistics.EntityData.Children.Append("dropped", types.YChild{"Dropped", &flowStatistics.Dropped})
    flowStatistics.EntityData.Leafs = types.NewOrderedMap()

    flowStatistics.EntityData.YListKeys = []string {}

    return &(flowStatistics.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri_FlowStatistics_Classified
// Classified statistics
type FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri_FlowStatistics_Classified struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Number of bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    Bytes interface{}
}

func (classified *FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri_FlowStatistics_Classified) GetEntityData() *types.CommonEntityData {
    classified.EntityData.YFilter = classified.YFilter
    classified.EntityData.YangName = "classified"
    classified.EntityData.BundleName = "cisco_ios_xr"
    classified.EntityData.ParentYangName = "flow-statistics"
    classified.EntityData.SegmentPath = "classified"
    classified.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classified.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classified.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classified.EntityData.Children = types.NewOrderedMap()
    classified.EntityData.Leafs = types.NewOrderedMap()
    classified.EntityData.Leafs.Append("packets", types.YLeaf{"Packets", classified.Packets})
    classified.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", classified.Bytes})

    classified.EntityData.YListKeys = []string {}

    return &(classified.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri_FlowStatistics_Dropped
// Drop statistics
type FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri_FlowStatistics_Dropped struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Number of bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    Bytes interface{}
}

func (dropped *FlowSpec_Vrfs_Vrf_Afs_Af_Nlris_Nlri_FlowStatistics_Dropped) GetEntityData() *types.CommonEntityData {
    dropped.EntityData.YFilter = dropped.YFilter
    dropped.EntityData.YangName = "dropped"
    dropped.EntityData.BundleName = "cisco_ios_xr"
    dropped.EntityData.ParentYangName = "flow-statistics"
    dropped.EntityData.SegmentPath = "dropped"
    dropped.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dropped.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dropped.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dropped.EntityData.Children = types.NewOrderedMap()
    dropped.EntityData.Leafs = types.NewOrderedMap()
    dropped.EntityData.Leafs.Append("packets", types.YLeaf{"Packets", dropped.Packets})
    dropped.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", dropped.Bytes})

    dropped.EntityData.YListKeys = []string {}

    return &(dropped.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_Flows
// Table of Flow
type FlowSpec_Vrfs_Vrf_Afs_Af_Flows struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow notation string. The type is slice of
    // FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow.
    Flow []*FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow
}

func (flows *FlowSpec_Vrfs_Vrf_Afs_Af_Flows) GetEntityData() *types.CommonEntityData {
    flows.EntityData.YFilter = flows.YFilter
    flows.EntityData.YangName = "flows"
    flows.EntityData.BundleName = "cisco_ios_xr"
    flows.EntityData.ParentYangName = "af"
    flows.EntityData.SegmentPath = "flows"
    flows.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flows.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flows.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flows.EntityData.Children = types.NewOrderedMap()
    flows.EntityData.Children.Append("flow", types.YChild{"Flow", nil})
    for i := range flows.Flow {
        flows.EntityData.Children.Append(types.GetSegmentPath(flows.Flow[i]), types.YChild{"Flow", flows.Flow[i]})
    }
    flows.EntityData.Leafs = types.NewOrderedMap()

    flows.EntityData.YListKeys = []string {}

    return &(flows.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow
// Flow notation string
type FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Enter the Flow notation string. The type is
    // string.
    FlowNotation interface{}

    // Flow statistics.
    FlowStatistics FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow_FlowStatistics
}

func (flow *FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow) GetEntityData() *types.CommonEntityData {
    flow.EntityData.YFilter = flow.YFilter
    flow.EntityData.YangName = "flow"
    flow.EntityData.BundleName = "cisco_ios_xr"
    flow.EntityData.ParentYangName = "flows"
    flow.EntityData.SegmentPath = "flow" + types.AddKeyToken(flow.FlowNotation, "flow-notation")
    flow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flow.EntityData.Children = types.NewOrderedMap()
    flow.EntityData.Children.Append("flow-statistics", types.YChild{"FlowStatistics", &flow.FlowStatistics})
    flow.EntityData.Leafs = types.NewOrderedMap()
    flow.EntityData.Leafs.Append("flow-notation", types.YLeaf{"FlowNotation", flow.FlowNotation})

    flow.EntityData.YListKeys = []string {"FlowNotation"}

    return &(flow.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow_FlowStatistics
// Flow statistics
type FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow_FlowStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Classified statistics.
    Classified FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow_FlowStatistics_Classified

    // Drop statistics.
    Dropped FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow_FlowStatistics_Dropped
}

func (flowStatistics *FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow_FlowStatistics) GetEntityData() *types.CommonEntityData {
    flowStatistics.EntityData.YFilter = flowStatistics.YFilter
    flowStatistics.EntityData.YangName = "flow-statistics"
    flowStatistics.EntityData.BundleName = "cisco_ios_xr"
    flowStatistics.EntityData.ParentYangName = "flow"
    flowStatistics.EntityData.SegmentPath = "flow-statistics"
    flowStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowStatistics.EntityData.Children = types.NewOrderedMap()
    flowStatistics.EntityData.Children.Append("classified", types.YChild{"Classified", &flowStatistics.Classified})
    flowStatistics.EntityData.Children.Append("dropped", types.YChild{"Dropped", &flowStatistics.Dropped})
    flowStatistics.EntityData.Leafs = types.NewOrderedMap()

    flowStatistics.EntityData.YListKeys = []string {}

    return &(flowStatistics.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow_FlowStatistics_Classified
// Classified statistics
type FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow_FlowStatistics_Classified struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Number of bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    Bytes interface{}
}

func (classified *FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow_FlowStatistics_Classified) GetEntityData() *types.CommonEntityData {
    classified.EntityData.YFilter = classified.YFilter
    classified.EntityData.YangName = "classified"
    classified.EntityData.BundleName = "cisco_ios_xr"
    classified.EntityData.ParentYangName = "flow-statistics"
    classified.EntityData.SegmentPath = "classified"
    classified.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classified.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classified.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classified.EntityData.Children = types.NewOrderedMap()
    classified.EntityData.Leafs = types.NewOrderedMap()
    classified.EntityData.Leafs.Append("packets", types.YLeaf{"Packets", classified.Packets})
    classified.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", classified.Bytes})

    classified.EntityData.YListKeys = []string {}

    return &(classified.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow_FlowStatistics_Dropped
// Drop statistics
type FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow_FlowStatistics_Dropped struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Number of bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    Bytes interface{}
}

func (dropped *FlowSpec_Vrfs_Vrf_Afs_Af_Flows_Flow_FlowStatistics_Dropped) GetEntityData() *types.CommonEntityData {
    dropped.EntityData.YFilter = dropped.YFilter
    dropped.EntityData.YangName = "dropped"
    dropped.EntityData.BundleName = "cisco_ios_xr"
    dropped.EntityData.ParentYangName = "flow-statistics"
    dropped.EntityData.SegmentPath = "dropped"
    dropped.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dropped.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dropped.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dropped.EntityData.Children = types.NewOrderedMap()
    dropped.EntityData.Leafs = types.NewOrderedMap()
    dropped.EntityData.Leafs.Append("packets", types.YLeaf{"Packets", dropped.Packets})
    dropped.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", dropped.Bytes})

    dropped.EntityData.YListKeys = []string {}

    return &(dropped.EntityData)
}

