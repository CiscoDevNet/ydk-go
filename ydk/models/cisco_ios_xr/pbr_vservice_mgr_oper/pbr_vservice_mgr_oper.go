// This module contains a collection of YANG definitions
// for Cisco IOS-XR pbr-vservice-mgr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   global-service-function-chaining: NSH Service Function
//     Chaining global operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package pbr_vservice_mgr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package pbr_vservice_mgr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-pbr-vservice-mgr-oper global-service-function-chaining}", reflect.TypeOf(GlobalServiceFunctionChaining{}))
    ydk.RegisterEntity("Cisco-IOS-XR-pbr-vservice-mgr-oper:global-service-function-chaining", reflect.TypeOf(GlobalServiceFunctionChaining{}))
}

// VsNshStats represents Vs nsh stats
type VsNshStats string

const (
    // vs nsh stats spi si
    VsNshStats_vs_nsh_stats_spi_si VsNshStats = "vs-nsh-stats-spi-si"

    // vs nsh stats ter min ate
    VsNshStats_vs_nsh_stats_ter_min_ate VsNshStats = "vs-nsh-stats-ter-min-ate"

    // vs nsh stats sf
    VsNshStats_vs_nsh_stats_sf VsNshStats = "vs-nsh-stats-sf"

    // vs nsh stats sff
    VsNshStats_vs_nsh_stats_sff VsNshStats = "vs-nsh-stats-sff"

    // vs nsh stats sff local
    VsNshStats_vs_nsh_stats_sff_local VsNshStats = "vs-nsh-stats-sff-local"

    // vs nsh stats sfp
    VsNshStats_vs_nsh_stats_sfp VsNshStats = "vs-nsh-stats-sfp"

    // vs nsh stats sfp detail
    VsNshStats_vs_nsh_stats_sfp_detail VsNshStats = "vs-nsh-stats-sfp-detail"

    // vs nsh stats max
    VsNshStats_vs_nsh_stats_max VsNshStats = "vs-nsh-stats-max"
)

// GlobalServiceFunctionChaining
// NSH Service Function Chaining global operational
// data
type GlobalServiceFunctionChaining struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service Function Path operational data.
    ServiceFunctionPath GlobalServiceFunctionChaining_ServiceFunctionPath

    // Service Function operational data.
    ServiceFunction GlobalServiceFunctionChaining_ServiceFunction

    // Service Function Forwarder operational data.
    ServiceFunctionForwarder GlobalServiceFunctionChaining_ServiceFunctionForwarder
}

func (globalServiceFunctionChaining *GlobalServiceFunctionChaining) GetEntityData() *types.CommonEntityData {
    globalServiceFunctionChaining.EntityData.YFilter = globalServiceFunctionChaining.YFilter
    globalServiceFunctionChaining.EntityData.YangName = "global-service-function-chaining"
    globalServiceFunctionChaining.EntityData.BundleName = "cisco_ios_xr"
    globalServiceFunctionChaining.EntityData.ParentYangName = "Cisco-IOS-XR-pbr-vservice-mgr-oper"
    globalServiceFunctionChaining.EntityData.SegmentPath = "Cisco-IOS-XR-pbr-vservice-mgr-oper:global-service-function-chaining"
    globalServiceFunctionChaining.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalServiceFunctionChaining.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalServiceFunctionChaining.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalServiceFunctionChaining.EntityData.Children = make(map[string]types.YChild)
    globalServiceFunctionChaining.EntityData.Children["service-function-path"] = types.YChild{"ServiceFunctionPath", &globalServiceFunctionChaining.ServiceFunctionPath}
    globalServiceFunctionChaining.EntityData.Children["service-function"] = types.YChild{"ServiceFunction", &globalServiceFunctionChaining.ServiceFunction}
    globalServiceFunctionChaining.EntityData.Children["service-function-forwarder"] = types.YChild{"ServiceFunctionForwarder", &globalServiceFunctionChaining.ServiceFunctionForwarder}
    globalServiceFunctionChaining.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(globalServiceFunctionChaining.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath
// Service Function Path operational data
type GlobalServiceFunctionChaining_ServiceFunctionPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service Function Path Id .
    PathIds GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds
}

func (serviceFunctionPath *GlobalServiceFunctionChaining_ServiceFunctionPath) GetEntityData() *types.CommonEntityData {
    serviceFunctionPath.EntityData.YFilter = serviceFunctionPath.YFilter
    serviceFunctionPath.EntityData.YangName = "service-function-path"
    serviceFunctionPath.EntityData.BundleName = "cisco_ios_xr"
    serviceFunctionPath.EntityData.ParentYangName = "global-service-function-chaining"
    serviceFunctionPath.EntityData.SegmentPath = "service-function-path"
    serviceFunctionPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceFunctionPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceFunctionPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceFunctionPath.EntityData.Children = make(map[string]types.YChild)
    serviceFunctionPath.EntityData.Children["path-ids"] = types.YChild{"PathIds", &serviceFunctionPath.PathIds}
    serviceFunctionPath.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(serviceFunctionPath.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds
// Service Function Path Id 
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specific Service-Function-Path identifier . The type is slice of
    // GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId.
    PathId []GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId
}

func (pathIds *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds) GetEntityData() *types.CommonEntityData {
    pathIds.EntityData.YFilter = pathIds.YFilter
    pathIds.EntityData.YangName = "path-ids"
    pathIds.EntityData.BundleName = "cisco_ios_xr"
    pathIds.EntityData.ParentYangName = "service-function-path"
    pathIds.EntityData.SegmentPath = "path-ids"
    pathIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathIds.EntityData.Children = make(map[string]types.YChild)
    pathIds.EntityData.Children["path-id"] = types.YChild{"PathId", nil}
    for i := range pathIds.PathId {
        pathIds.EntityData.Children[types.GetSegmentPath(&pathIds.PathId[i])] = types.YChild{"PathId", &pathIds.PathId[i]}
    }
    pathIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pathIds.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId
// Specific Service-Function-Path identifier 
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Service-Function-Path identifier. The
    // type is interface{} with range: 1..16777215.
    Id interface{}

    // SFP Statistics.
    Stats GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats

    // Service Index Belonging to Path.
    ServiceIndexes GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes
}

func (pathId *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId) GetEntityData() *types.CommonEntityData {
    pathId.EntityData.YFilter = pathId.YFilter
    pathId.EntityData.YangName = "path-id"
    pathId.EntityData.BundleName = "cisco_ios_xr"
    pathId.EntityData.ParentYangName = "path-ids"
    pathId.EntityData.SegmentPath = "path-id" + "[id='" + fmt.Sprintf("%v", pathId.Id) + "']"
    pathId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathId.EntityData.Children = make(map[string]types.YChild)
    pathId.EntityData.Children["stats"] = types.YChild{"Stats", &pathId.Stats}
    pathId.EntityData.Children["service-indexes"] = types.YChild{"ServiceIndexes", &pathId.ServiceIndexes}
    pathId.EntityData.Leafs = make(map[string]types.YLeaf)
    pathId.EntityData.Leafs["id"] = types.YLeaf{"Id", pathId.Id}
    return &(pathId.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats
// SFP Statistics
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detail statistics per service index .
    Detail GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail

    // Combined statistics of all service index in service functionpath.
    Summarized GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized
}

func (stats *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "path-id"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Children["detail"] = types.YChild{"Detail", &stats.Detail}
    stats.EntityData.Children["summarized"] = types.YChild{"Summarized", &stats.Summarized}
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stats.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail
// Detail statistics per service index 
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics data.
    Data GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data

    // SI array in case of detail stats. The type is slice of
    // GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr.
    SiArr []GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr
}

func (detail *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "stats"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["data"] = types.YChild{"Data", &detail.Data}
    detail.EntityData.Children["si-arr"] = types.YChild{"SiArr", nil}
    for i := range detail.SiArr {
        detail.EntityData.Children[types.GetSegmentPath(&detail.SiArr[i])] = types.YChild{"SiArr", &detail.SiArr[i]}
    }
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(detail.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data
// Statistics data
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type_ interface{}

    // SFP stats.
    Sfp GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp

    // SPI SI stats.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi

    // Terminate stats.
    Term GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term

    // Service function stats.
    Sf GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf

    // Service function forwarder stats.
    Sff GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff

    // Local service function forwarder stats.
    SffLocal GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal
}

func (data *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "detail"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = make(map[string]types.YChild)
    data.EntityData.Children["sfp"] = types.YChild{"Sfp", &data.Sfp}
    data.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &data.SpiSi}
    data.EntityData.Children["term"] = types.YChild{"Term", &data.Term}
    data.EntityData.Children["sf"] = types.YChild{"Sf", &data.Sf}
    data.EntityData.Children["sff"] = types.YChild{"Sff", &data.Sff}
    data.EntityData.Children["sff-local"] = types.YChild{"SffLocal", &data.SffLocal}
    data.EntityData.Leafs = make(map[string]types.YLeaf)
    data.EntityData.Leafs["type"] = types.YLeaf{"Type_", data.Type_}
    return &(data.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp
// SFP stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service index counters.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi

    // Terminate counters.
    Term GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term
}

func (sfp *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetEntityData() *types.CommonEntityData {
    sfp.EntityData.YFilter = sfp.YFilter
    sfp.EntityData.YangName = "sfp"
    sfp.EntityData.BundleName = "cisco_ios_xr"
    sfp.EntityData.ParentYangName = "data"
    sfp.EntityData.SegmentPath = "sfp"
    sfp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfp.EntityData.Children = make(map[string]types.YChild)
    sfp.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &sfp.SpiSi}
    sfp.EntityData.Children["term"] = types.YChild{"Term", &sfp.Term}
    sfp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sfp.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi
// Service index counters
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "sfp"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term
// Terminate counters
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "sfp"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi
// SPI SI stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "data"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term
// Terminate stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "data"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf
// Service function stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sf *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetEntityData() *types.CommonEntityData {
    sf.EntityData.YFilter = sf.YFilter
    sf.EntityData.YangName = "sf"
    sf.EntityData.BundleName = "cisco_ios_xr"
    sf.EntityData.ParentYangName = "data"
    sf.EntityData.SegmentPath = "sf"
    sf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sf.EntityData.Children = make(map[string]types.YChild)
    sf.EntityData.Leafs = make(map[string]types.YLeaf)
    sf.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", sf.ProcessedPkts}
    sf.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", sf.ProcessedBytes}
    return &(sf.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff
// Service function forwarder stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sff *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetEntityData() *types.CommonEntityData {
    sff.EntityData.YFilter = sff.YFilter
    sff.EntityData.YangName = "sff"
    sff.EntityData.BundleName = "cisco_ios_xr"
    sff.EntityData.ParentYangName = "data"
    sff.EntityData.SegmentPath = "sff"
    sff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sff.EntityData.Children = make(map[string]types.YChild)
    sff.EntityData.Leafs = make(map[string]types.YLeaf)
    sff.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", sff.ProcessedPkts}
    sff.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", sff.ProcessedBytes}
    return &(sff.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal
// Local service function forwarder stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets with invalid NSH header. The type is interface{} with
    // range: 0..18446744073709551615.
    MalformedErrPkts interface{}

    // Number of packets with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615.
    LookupErrPkts interface{}

    // Total bytes with invalid NSH header. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    MalformedErrBytes interface{}

    // Total bytes with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    LookupErrBytes interface{}
}

func (sffLocal *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetEntityData() *types.CommonEntityData {
    sffLocal.EntityData.YFilter = sffLocal.YFilter
    sffLocal.EntityData.YangName = "sff-local"
    sffLocal.EntityData.BundleName = "cisco_ios_xr"
    sffLocal.EntityData.ParentYangName = "data"
    sffLocal.EntityData.SegmentPath = "sff-local"
    sffLocal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sffLocal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sffLocal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sffLocal.EntityData.Children = make(map[string]types.YChild)
    sffLocal.EntityData.Leafs = make(map[string]types.YLeaf)
    sffLocal.EntityData.Leafs["malformed-err-pkts"] = types.YLeaf{"MalformedErrPkts", sffLocal.MalformedErrPkts}
    sffLocal.EntityData.Leafs["lookup-err-pkts"] = types.YLeaf{"LookupErrPkts", sffLocal.LookupErrPkts}
    sffLocal.EntityData.Leafs["malformed-err-bytes"] = types.YLeaf{"MalformedErrBytes", sffLocal.MalformedErrBytes}
    sffLocal.EntityData.Leafs["lookup-err-bytes"] = types.YLeaf{"LookupErrBytes", sffLocal.LookupErrBytes}
    return &(sffLocal.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr
// SI array in case of detail stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service index. The type is interface{} with range: 0..255.
    Si interface{}

    // Stats counter for this index.
    Data GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data
}

func (siArr *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetEntityData() *types.CommonEntityData {
    siArr.EntityData.YFilter = siArr.YFilter
    siArr.EntityData.YangName = "si-arr"
    siArr.EntityData.BundleName = "cisco_ios_xr"
    siArr.EntityData.ParentYangName = "detail"
    siArr.EntityData.SegmentPath = "si-arr"
    siArr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    siArr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    siArr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    siArr.EntityData.Children = make(map[string]types.YChild)
    siArr.EntityData.Children["data"] = types.YChild{"Data", &siArr.Data}
    siArr.EntityData.Leafs = make(map[string]types.YLeaf)
    siArr.EntityData.Leafs["si"] = types.YLeaf{"Si", siArr.Si}
    return &(siArr.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data
// Stats counter for this index
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type_ interface{}

    // SF/SFF stats.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi

    // Terminate stats.
    Term GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term
}

func (data *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "si-arr"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = make(map[string]types.YChild)
    data.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &data.SpiSi}
    data.EntityData.Children["term"] = types.YChild{"Term", &data.Term}
    data.EntityData.Leafs = make(map[string]types.YLeaf)
    data.EntityData.Leafs["type"] = types.YLeaf{"Type_", data.Type_}
    return &(data.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi
// SF/SFF stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "data"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term
// Terminate stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "data"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized
// Combined statistics of all service index in
// service functionpath
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics data.
    Data GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data

    // SI array in case of detail stats. The type is slice of
    // GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr.
    SiArr []GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr
}

func (summarized *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetEntityData() *types.CommonEntityData {
    summarized.EntityData.YFilter = summarized.YFilter
    summarized.EntityData.YangName = "summarized"
    summarized.EntityData.BundleName = "cisco_ios_xr"
    summarized.EntityData.ParentYangName = "stats"
    summarized.EntityData.SegmentPath = "summarized"
    summarized.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summarized.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summarized.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summarized.EntityData.Children = make(map[string]types.YChild)
    summarized.EntityData.Children["data"] = types.YChild{"Data", &summarized.Data}
    summarized.EntityData.Children["si-arr"] = types.YChild{"SiArr", nil}
    for i := range summarized.SiArr {
        summarized.EntityData.Children[types.GetSegmentPath(&summarized.SiArr[i])] = types.YChild{"SiArr", &summarized.SiArr[i]}
    }
    summarized.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(summarized.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data
// Statistics data
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type_ interface{}

    // SFP stats.
    Sfp GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp

    // SPI SI stats.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi

    // Terminate stats.
    Term GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term

    // Service function stats.
    Sf GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf

    // Service function forwarder stats.
    Sff GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff

    // Local service function forwarder stats.
    SffLocal GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal
}

func (data *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "summarized"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = make(map[string]types.YChild)
    data.EntityData.Children["sfp"] = types.YChild{"Sfp", &data.Sfp}
    data.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &data.SpiSi}
    data.EntityData.Children["term"] = types.YChild{"Term", &data.Term}
    data.EntityData.Children["sf"] = types.YChild{"Sf", &data.Sf}
    data.EntityData.Children["sff"] = types.YChild{"Sff", &data.Sff}
    data.EntityData.Children["sff-local"] = types.YChild{"SffLocal", &data.SffLocal}
    data.EntityData.Leafs = make(map[string]types.YLeaf)
    data.EntityData.Leafs["type"] = types.YLeaf{"Type_", data.Type_}
    return &(data.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp
// SFP stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service index counters.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi

    // Terminate counters.
    Term GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term
}

func (sfp *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetEntityData() *types.CommonEntityData {
    sfp.EntityData.YFilter = sfp.YFilter
    sfp.EntityData.YangName = "sfp"
    sfp.EntityData.BundleName = "cisco_ios_xr"
    sfp.EntityData.ParentYangName = "data"
    sfp.EntityData.SegmentPath = "sfp"
    sfp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfp.EntityData.Children = make(map[string]types.YChild)
    sfp.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &sfp.SpiSi}
    sfp.EntityData.Children["term"] = types.YChild{"Term", &sfp.Term}
    sfp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sfp.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi
// Service index counters
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "sfp"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term
// Terminate counters
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "sfp"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi
// SPI SI stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "data"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term
// Terminate stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "data"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf
// Service function stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sf *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetEntityData() *types.CommonEntityData {
    sf.EntityData.YFilter = sf.YFilter
    sf.EntityData.YangName = "sf"
    sf.EntityData.BundleName = "cisco_ios_xr"
    sf.EntityData.ParentYangName = "data"
    sf.EntityData.SegmentPath = "sf"
    sf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sf.EntityData.Children = make(map[string]types.YChild)
    sf.EntityData.Leafs = make(map[string]types.YLeaf)
    sf.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", sf.ProcessedPkts}
    sf.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", sf.ProcessedBytes}
    return &(sf.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff
// Service function forwarder stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sff *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetEntityData() *types.CommonEntityData {
    sff.EntityData.YFilter = sff.YFilter
    sff.EntityData.YangName = "sff"
    sff.EntityData.BundleName = "cisco_ios_xr"
    sff.EntityData.ParentYangName = "data"
    sff.EntityData.SegmentPath = "sff"
    sff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sff.EntityData.Children = make(map[string]types.YChild)
    sff.EntityData.Leafs = make(map[string]types.YLeaf)
    sff.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", sff.ProcessedPkts}
    sff.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", sff.ProcessedBytes}
    return &(sff.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal
// Local service function forwarder stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets with invalid NSH header. The type is interface{} with
    // range: 0..18446744073709551615.
    MalformedErrPkts interface{}

    // Number of packets with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615.
    LookupErrPkts interface{}

    // Total bytes with invalid NSH header. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    MalformedErrBytes interface{}

    // Total bytes with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    LookupErrBytes interface{}
}

func (sffLocal *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetEntityData() *types.CommonEntityData {
    sffLocal.EntityData.YFilter = sffLocal.YFilter
    sffLocal.EntityData.YangName = "sff-local"
    sffLocal.EntityData.BundleName = "cisco_ios_xr"
    sffLocal.EntityData.ParentYangName = "data"
    sffLocal.EntityData.SegmentPath = "sff-local"
    sffLocal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sffLocal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sffLocal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sffLocal.EntityData.Children = make(map[string]types.YChild)
    sffLocal.EntityData.Leafs = make(map[string]types.YLeaf)
    sffLocal.EntityData.Leafs["malformed-err-pkts"] = types.YLeaf{"MalformedErrPkts", sffLocal.MalformedErrPkts}
    sffLocal.EntityData.Leafs["lookup-err-pkts"] = types.YLeaf{"LookupErrPkts", sffLocal.LookupErrPkts}
    sffLocal.EntityData.Leafs["malformed-err-bytes"] = types.YLeaf{"MalformedErrBytes", sffLocal.MalformedErrBytes}
    sffLocal.EntityData.Leafs["lookup-err-bytes"] = types.YLeaf{"LookupErrBytes", sffLocal.LookupErrBytes}
    return &(sffLocal.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr
// SI array in case of detail stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service index. The type is interface{} with range: 0..255.
    Si interface{}

    // Stats counter for this index.
    Data GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data
}

func (siArr *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetEntityData() *types.CommonEntityData {
    siArr.EntityData.YFilter = siArr.YFilter
    siArr.EntityData.YangName = "si-arr"
    siArr.EntityData.BundleName = "cisco_ios_xr"
    siArr.EntityData.ParentYangName = "summarized"
    siArr.EntityData.SegmentPath = "si-arr"
    siArr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    siArr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    siArr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    siArr.EntityData.Children = make(map[string]types.YChild)
    siArr.EntityData.Children["data"] = types.YChild{"Data", &siArr.Data}
    siArr.EntityData.Leafs = make(map[string]types.YLeaf)
    siArr.EntityData.Leafs["si"] = types.YLeaf{"Si", siArr.Si}
    return &(siArr.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data
// Stats counter for this index
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type_ interface{}

    // SF/SFF stats.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi

    // Terminate stats.
    Term GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term
}

func (data *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "si-arr"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = make(map[string]types.YChild)
    data.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &data.SpiSi}
    data.EntityData.Children["term"] = types.YChild{"Term", &data.Term}
    data.EntityData.Leafs = make(map[string]types.YLeaf)
    data.EntityData.Leafs["type"] = types.YLeaf{"Type_", data.Type_}
    return &(data.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi
// SF/SFF stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "data"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term
// Terminate stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "data"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes
// Service Index Belonging to Path
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service index operational data belonging to this path. The type is slice of
    // GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex.
    ServiceIndex []GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex
}

func (serviceIndexes *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetEntityData() *types.CommonEntityData {
    serviceIndexes.EntityData.YFilter = serviceIndexes.YFilter
    serviceIndexes.EntityData.YangName = "service-indexes"
    serviceIndexes.EntityData.BundleName = "cisco_ios_xr"
    serviceIndexes.EntityData.ParentYangName = "path-id"
    serviceIndexes.EntityData.SegmentPath = "service-indexes"
    serviceIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceIndexes.EntityData.Children = make(map[string]types.YChild)
    serviceIndexes.EntityData.Children["service-index"] = types.YChild{"ServiceIndex", nil}
    for i := range serviceIndexes.ServiceIndex {
        serviceIndexes.EntityData.Children[types.GetSegmentPath(&serviceIndexes.ServiceIndex[i])] = types.YChild{"ServiceIndex", &serviceIndexes.ServiceIndex[i]}
    }
    serviceIndexes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(serviceIndexes.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex
// Service index operational data belonging to
// this path
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Service Index. The type is interface{} with range:
    // 1..255.
    Index interface{}

    // Statistics data.
    Data GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data

    // SI array in case of detail stats. The type is slice of
    // GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr.
    SiArr []GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr
}

func (serviceIndex *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetEntityData() *types.CommonEntityData {
    serviceIndex.EntityData.YFilter = serviceIndex.YFilter
    serviceIndex.EntityData.YangName = "service-index"
    serviceIndex.EntityData.BundleName = "cisco_ios_xr"
    serviceIndex.EntityData.ParentYangName = "service-indexes"
    serviceIndex.EntityData.SegmentPath = "service-index" + "[index='" + fmt.Sprintf("%v", serviceIndex.Index) + "']"
    serviceIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceIndex.EntityData.Children = make(map[string]types.YChild)
    serviceIndex.EntityData.Children["data"] = types.YChild{"Data", &serviceIndex.Data}
    serviceIndex.EntityData.Children["si-arr"] = types.YChild{"SiArr", nil}
    for i := range serviceIndex.SiArr {
        serviceIndex.EntityData.Children[types.GetSegmentPath(&serviceIndex.SiArr[i])] = types.YChild{"SiArr", &serviceIndex.SiArr[i]}
    }
    serviceIndex.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceIndex.EntityData.Leafs["index"] = types.YLeaf{"Index", serviceIndex.Index}
    return &(serviceIndex.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data
// Statistics data
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type_ interface{}

    // SFP stats.
    Sfp GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp

    // SPI SI stats.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi

    // Terminate stats.
    Term GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term

    // Service function stats.
    Sf GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf

    // Service function forwarder stats.
    Sff GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff

    // Local service function forwarder stats.
    SffLocal GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal
}

func (data *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "service-index"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = make(map[string]types.YChild)
    data.EntityData.Children["sfp"] = types.YChild{"Sfp", &data.Sfp}
    data.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &data.SpiSi}
    data.EntityData.Children["term"] = types.YChild{"Term", &data.Term}
    data.EntityData.Children["sf"] = types.YChild{"Sf", &data.Sf}
    data.EntityData.Children["sff"] = types.YChild{"Sff", &data.Sff}
    data.EntityData.Children["sff-local"] = types.YChild{"SffLocal", &data.SffLocal}
    data.EntityData.Leafs = make(map[string]types.YLeaf)
    data.EntityData.Leafs["type"] = types.YLeaf{"Type_", data.Type_}
    return &(data.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp
// SFP stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service index counters.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi

    // Terminate counters.
    Term GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term
}

func (sfp *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetEntityData() *types.CommonEntityData {
    sfp.EntityData.YFilter = sfp.YFilter
    sfp.EntityData.YangName = "sfp"
    sfp.EntityData.BundleName = "cisco_ios_xr"
    sfp.EntityData.ParentYangName = "data"
    sfp.EntityData.SegmentPath = "sfp"
    sfp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfp.EntityData.Children = make(map[string]types.YChild)
    sfp.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &sfp.SpiSi}
    sfp.EntityData.Children["term"] = types.YChild{"Term", &sfp.Term}
    sfp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sfp.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi
// Service index counters
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "sfp"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term
// Terminate counters
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "sfp"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi
// SPI SI stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "data"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term
// Terminate stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "data"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf
// Service function stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sf *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetEntityData() *types.CommonEntityData {
    sf.EntityData.YFilter = sf.YFilter
    sf.EntityData.YangName = "sf"
    sf.EntityData.BundleName = "cisco_ios_xr"
    sf.EntityData.ParentYangName = "data"
    sf.EntityData.SegmentPath = "sf"
    sf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sf.EntityData.Children = make(map[string]types.YChild)
    sf.EntityData.Leafs = make(map[string]types.YLeaf)
    sf.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", sf.ProcessedPkts}
    sf.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", sf.ProcessedBytes}
    return &(sf.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff
// Service function forwarder stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sff *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetEntityData() *types.CommonEntityData {
    sff.EntityData.YFilter = sff.YFilter
    sff.EntityData.YangName = "sff"
    sff.EntityData.BundleName = "cisco_ios_xr"
    sff.EntityData.ParentYangName = "data"
    sff.EntityData.SegmentPath = "sff"
    sff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sff.EntityData.Children = make(map[string]types.YChild)
    sff.EntityData.Leafs = make(map[string]types.YLeaf)
    sff.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", sff.ProcessedPkts}
    sff.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", sff.ProcessedBytes}
    return &(sff.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal
// Local service function forwarder stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets with invalid NSH header. The type is interface{} with
    // range: 0..18446744073709551615.
    MalformedErrPkts interface{}

    // Number of packets with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615.
    LookupErrPkts interface{}

    // Total bytes with invalid NSH header. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    MalformedErrBytes interface{}

    // Total bytes with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    LookupErrBytes interface{}
}

func (sffLocal *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetEntityData() *types.CommonEntityData {
    sffLocal.EntityData.YFilter = sffLocal.YFilter
    sffLocal.EntityData.YangName = "sff-local"
    sffLocal.EntityData.BundleName = "cisco_ios_xr"
    sffLocal.EntityData.ParentYangName = "data"
    sffLocal.EntityData.SegmentPath = "sff-local"
    sffLocal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sffLocal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sffLocal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sffLocal.EntityData.Children = make(map[string]types.YChild)
    sffLocal.EntityData.Leafs = make(map[string]types.YLeaf)
    sffLocal.EntityData.Leafs["malformed-err-pkts"] = types.YLeaf{"MalformedErrPkts", sffLocal.MalformedErrPkts}
    sffLocal.EntityData.Leafs["lookup-err-pkts"] = types.YLeaf{"LookupErrPkts", sffLocal.LookupErrPkts}
    sffLocal.EntityData.Leafs["malformed-err-bytes"] = types.YLeaf{"MalformedErrBytes", sffLocal.MalformedErrBytes}
    sffLocal.EntityData.Leafs["lookup-err-bytes"] = types.YLeaf{"LookupErrBytes", sffLocal.LookupErrBytes}
    return &(sffLocal.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr
// SI array in case of detail stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service index. The type is interface{} with range: 0..255.
    Si interface{}

    // Stats counter for this index.
    Data GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data
}

func (siArr *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetEntityData() *types.CommonEntityData {
    siArr.EntityData.YFilter = siArr.YFilter
    siArr.EntityData.YangName = "si-arr"
    siArr.EntityData.BundleName = "cisco_ios_xr"
    siArr.EntityData.ParentYangName = "service-index"
    siArr.EntityData.SegmentPath = "si-arr"
    siArr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    siArr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    siArr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    siArr.EntityData.Children = make(map[string]types.YChild)
    siArr.EntityData.Children["data"] = types.YChild{"Data", &siArr.Data}
    siArr.EntityData.Leafs = make(map[string]types.YLeaf)
    siArr.EntityData.Leafs["si"] = types.YLeaf{"Si", siArr.Si}
    return &(siArr.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data
// Stats counter for this index
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type_ interface{}

    // SF/SFF stats.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi

    // Terminate stats.
    Term GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term
}

func (data *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "si-arr"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = make(map[string]types.YChild)
    data.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &data.SpiSi}
    data.EntityData.Children["term"] = types.YChild{"Term", &data.Term}
    data.EntityData.Leafs = make(map[string]types.YLeaf)
    data.EntityData.Leafs["type"] = types.YLeaf{"Type_", data.Type_}
    return &(data.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi
// SF/SFF stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "data"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term
// Terminate stats
type GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "data"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction
// Service Function operational data
type GlobalServiceFunctionChaining_ServiceFunction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of Service Function Names.
    SfNames GlobalServiceFunctionChaining_ServiceFunction_SfNames
}

func (serviceFunction *GlobalServiceFunctionChaining_ServiceFunction) GetEntityData() *types.CommonEntityData {
    serviceFunction.EntityData.YFilter = serviceFunction.YFilter
    serviceFunction.EntityData.YangName = "service-function"
    serviceFunction.EntityData.BundleName = "cisco_ios_xr"
    serviceFunction.EntityData.ParentYangName = "global-service-function-chaining"
    serviceFunction.EntityData.SegmentPath = "service-function"
    serviceFunction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceFunction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceFunction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceFunction.EntityData.Children = make(map[string]types.YChild)
    serviceFunction.EntityData.Children["sf-names"] = types.YChild{"SfNames", &serviceFunction.SfNames}
    serviceFunction.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(serviceFunction.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames
// List of Service Function Names
type GlobalServiceFunctionChaining_ServiceFunction_SfNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of Service Function. The type is slice of
    // GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName.
    SfName []GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName
}

func (sfNames *GlobalServiceFunctionChaining_ServiceFunction_SfNames) GetEntityData() *types.CommonEntityData {
    sfNames.EntityData.YFilter = sfNames.YFilter
    sfNames.EntityData.YangName = "sf-names"
    sfNames.EntityData.BundleName = "cisco_ios_xr"
    sfNames.EntityData.ParentYangName = "service-function"
    sfNames.EntityData.SegmentPath = "sf-names"
    sfNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfNames.EntityData.Children = make(map[string]types.YChild)
    sfNames.EntityData.Children["sf-name"] = types.YChild{"SfName", nil}
    for i := range sfNames.SfName {
        sfNames.EntityData.Children[types.GetSegmentPath(&sfNames.SfName[i])] = types.YChild{"SfName", &sfNames.SfName[i]}
    }
    sfNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sfNames.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName
// Name of Service Function
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name. The type is string with length: 1..32.
    Name interface{}

    // Statistics data.
    Data GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data

    // SI array in case of detail stats. The type is slice of
    // GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr.
    SiArr []GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr
}

func (sfName *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName) GetEntityData() *types.CommonEntityData {
    sfName.EntityData.YFilter = sfName.YFilter
    sfName.EntityData.YangName = "sf-name"
    sfName.EntityData.BundleName = "cisco_ios_xr"
    sfName.EntityData.ParentYangName = "sf-names"
    sfName.EntityData.SegmentPath = "sf-name" + "[name='" + fmt.Sprintf("%v", sfName.Name) + "']"
    sfName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfName.EntityData.Children = make(map[string]types.YChild)
    sfName.EntityData.Children["data"] = types.YChild{"Data", &sfName.Data}
    sfName.EntityData.Children["si-arr"] = types.YChild{"SiArr", nil}
    for i := range sfName.SiArr {
        sfName.EntityData.Children[types.GetSegmentPath(&sfName.SiArr[i])] = types.YChild{"SiArr", &sfName.SiArr[i]}
    }
    sfName.EntityData.Leafs = make(map[string]types.YLeaf)
    sfName.EntityData.Leafs["name"] = types.YLeaf{"Name", sfName.Name}
    return &(sfName.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data
// Statistics data
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type_ interface{}

    // SFP stats.
    Sfp GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sfp

    // SPI SI stats.
    SpiSi GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_SpiSi

    // Terminate stats.
    Term GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Term

    // Service function stats.
    Sf GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sf

    // Service function forwarder stats.
    Sff GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sff

    // Local service function forwarder stats.
    SffLocal GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_SffLocal
}

func (data *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "sf-name"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = make(map[string]types.YChild)
    data.EntityData.Children["sfp"] = types.YChild{"Sfp", &data.Sfp}
    data.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &data.SpiSi}
    data.EntityData.Children["term"] = types.YChild{"Term", &data.Term}
    data.EntityData.Children["sf"] = types.YChild{"Sf", &data.Sf}
    data.EntityData.Children["sff"] = types.YChild{"Sff", &data.Sff}
    data.EntityData.Children["sff-local"] = types.YChild{"SffLocal", &data.SffLocal}
    data.EntityData.Leafs = make(map[string]types.YLeaf)
    data.EntityData.Leafs["type"] = types.YLeaf{"Type_", data.Type_}
    return &(data.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sfp
// SFP stats
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sfp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service index counters.
    SpiSi GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi

    // Terminate counters.
    Term GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sfp_Term
}

func (sfp *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sfp) GetEntityData() *types.CommonEntityData {
    sfp.EntityData.YFilter = sfp.YFilter
    sfp.EntityData.YangName = "sfp"
    sfp.EntityData.BundleName = "cisco_ios_xr"
    sfp.EntityData.ParentYangName = "data"
    sfp.EntityData.SegmentPath = "sfp"
    sfp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfp.EntityData.Children = make(map[string]types.YChild)
    sfp.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &sfp.SpiSi}
    sfp.EntityData.Children["term"] = types.YChild{"Term", &sfp.Term}
    sfp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sfp.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi
// Service index counters
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "sfp"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sfp_Term
// Terminate counters
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sfp_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "sfp"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_SpiSi
// SPI SI stats
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "data"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Term
// Terminate stats
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "data"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sf
// Service function stats
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sf *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sf) GetEntityData() *types.CommonEntityData {
    sf.EntityData.YFilter = sf.YFilter
    sf.EntityData.YangName = "sf"
    sf.EntityData.BundleName = "cisco_ios_xr"
    sf.EntityData.ParentYangName = "data"
    sf.EntityData.SegmentPath = "sf"
    sf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sf.EntityData.Children = make(map[string]types.YChild)
    sf.EntityData.Leafs = make(map[string]types.YLeaf)
    sf.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", sf.ProcessedPkts}
    sf.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", sf.ProcessedBytes}
    return &(sf.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sff
// Service function forwarder stats
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sff *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_Sff) GetEntityData() *types.CommonEntityData {
    sff.EntityData.YFilter = sff.YFilter
    sff.EntityData.YangName = "sff"
    sff.EntityData.BundleName = "cisco_ios_xr"
    sff.EntityData.ParentYangName = "data"
    sff.EntityData.SegmentPath = "sff"
    sff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sff.EntityData.Children = make(map[string]types.YChild)
    sff.EntityData.Leafs = make(map[string]types.YLeaf)
    sff.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", sff.ProcessedPkts}
    sff.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", sff.ProcessedBytes}
    return &(sff.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_SffLocal
// Local service function forwarder stats
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_SffLocal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets with invalid NSH header. The type is interface{} with
    // range: 0..18446744073709551615.
    MalformedErrPkts interface{}

    // Number of packets with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615.
    LookupErrPkts interface{}

    // Total bytes with invalid NSH header. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    MalformedErrBytes interface{}

    // Total bytes with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    LookupErrBytes interface{}
}

func (sffLocal *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_Data_SffLocal) GetEntityData() *types.CommonEntityData {
    sffLocal.EntityData.YFilter = sffLocal.YFilter
    sffLocal.EntityData.YangName = "sff-local"
    sffLocal.EntityData.BundleName = "cisco_ios_xr"
    sffLocal.EntityData.ParentYangName = "data"
    sffLocal.EntityData.SegmentPath = "sff-local"
    sffLocal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sffLocal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sffLocal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sffLocal.EntityData.Children = make(map[string]types.YChild)
    sffLocal.EntityData.Leafs = make(map[string]types.YLeaf)
    sffLocal.EntityData.Leafs["malformed-err-pkts"] = types.YLeaf{"MalformedErrPkts", sffLocal.MalformedErrPkts}
    sffLocal.EntityData.Leafs["lookup-err-pkts"] = types.YLeaf{"LookupErrPkts", sffLocal.LookupErrPkts}
    sffLocal.EntityData.Leafs["malformed-err-bytes"] = types.YLeaf{"MalformedErrBytes", sffLocal.MalformedErrBytes}
    sffLocal.EntityData.Leafs["lookup-err-bytes"] = types.YLeaf{"LookupErrBytes", sffLocal.LookupErrBytes}
    return &(sffLocal.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr
// SI array in case of detail stats
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service index. The type is interface{} with range: 0..255.
    Si interface{}

    // Stats counter for this index.
    Data GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr_Data
}

func (siArr *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr) GetEntityData() *types.CommonEntityData {
    siArr.EntityData.YFilter = siArr.YFilter
    siArr.EntityData.YangName = "si-arr"
    siArr.EntityData.BundleName = "cisco_ios_xr"
    siArr.EntityData.ParentYangName = "sf-name"
    siArr.EntityData.SegmentPath = "si-arr"
    siArr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    siArr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    siArr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    siArr.EntityData.Children = make(map[string]types.YChild)
    siArr.EntityData.Children["data"] = types.YChild{"Data", &siArr.Data}
    siArr.EntityData.Leafs = make(map[string]types.YLeaf)
    siArr.EntityData.Leafs["si"] = types.YLeaf{"Si", siArr.Si}
    return &(siArr.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr_Data
// Stats counter for this index
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type_ interface{}

    // SF/SFF stats.
    SpiSi GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi

    // Terminate stats.
    Term GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr_Data_Term
}

func (data *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "si-arr"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = make(map[string]types.YChild)
    data.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &data.SpiSi}
    data.EntityData.Children["term"] = types.YChild{"Term", &data.Term}
    data.EntityData.Leafs = make(map[string]types.YLeaf)
    data.EntityData.Leafs["type"] = types.YLeaf{"Type_", data.Type_}
    return &(data.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi
// SF/SFF stats
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "data"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr_Data_Term
// Terminate stats
type GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr_Data_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "data"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder
// Service Function Forwarder operational data
type GlobalServiceFunctionChaining_ServiceFunctionForwarder struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of Service Function Forwarder Names.
    SffNames GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames

    // Local Service Function Forwarder operational data.
    Local GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local
}

func (serviceFunctionForwarder *GlobalServiceFunctionChaining_ServiceFunctionForwarder) GetEntityData() *types.CommonEntityData {
    serviceFunctionForwarder.EntityData.YFilter = serviceFunctionForwarder.YFilter
    serviceFunctionForwarder.EntityData.YangName = "service-function-forwarder"
    serviceFunctionForwarder.EntityData.BundleName = "cisco_ios_xr"
    serviceFunctionForwarder.EntityData.ParentYangName = "global-service-function-chaining"
    serviceFunctionForwarder.EntityData.SegmentPath = "service-function-forwarder"
    serviceFunctionForwarder.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceFunctionForwarder.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceFunctionForwarder.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceFunctionForwarder.EntityData.Children = make(map[string]types.YChild)
    serviceFunctionForwarder.EntityData.Children["sff-names"] = types.YChild{"SffNames", &serviceFunctionForwarder.SffNames}
    serviceFunctionForwarder.EntityData.Children["local"] = types.YChild{"Local", &serviceFunctionForwarder.Local}
    serviceFunctionForwarder.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(serviceFunctionForwarder.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames
// List of Service Function Forwarder Names
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of Service Function Forwarder. The type is slice of
    // GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName.
    SffName []GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName
}

func (sffNames *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames) GetEntityData() *types.CommonEntityData {
    sffNames.EntityData.YFilter = sffNames.YFilter
    sffNames.EntityData.YangName = "sff-names"
    sffNames.EntityData.BundleName = "cisco_ios_xr"
    sffNames.EntityData.ParentYangName = "service-function-forwarder"
    sffNames.EntityData.SegmentPath = "sff-names"
    sffNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sffNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sffNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sffNames.EntityData.Children = make(map[string]types.YChild)
    sffNames.EntityData.Children["sff-name"] = types.YChild{"SffName", nil}
    for i := range sffNames.SffName {
        sffNames.EntityData.Children[types.GetSegmentPath(&sffNames.SffName[i])] = types.YChild{"SffName", &sffNames.SffName[i]}
    }
    sffNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sffNames.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName
// Name of Service Function Forwarder
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name. The type is string with length: 1..32.
    Name interface{}

    // Statistics data.
    Data GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data

    // SI array in case of detail stats. The type is slice of
    // GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr.
    SiArr []GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr
}

func (sffName *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName) GetEntityData() *types.CommonEntityData {
    sffName.EntityData.YFilter = sffName.YFilter
    sffName.EntityData.YangName = "sff-name"
    sffName.EntityData.BundleName = "cisco_ios_xr"
    sffName.EntityData.ParentYangName = "sff-names"
    sffName.EntityData.SegmentPath = "sff-name" + "[name='" + fmt.Sprintf("%v", sffName.Name) + "']"
    sffName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sffName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sffName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sffName.EntityData.Children = make(map[string]types.YChild)
    sffName.EntityData.Children["data"] = types.YChild{"Data", &sffName.Data}
    sffName.EntityData.Children["si-arr"] = types.YChild{"SiArr", nil}
    for i := range sffName.SiArr {
        sffName.EntityData.Children[types.GetSegmentPath(&sffName.SiArr[i])] = types.YChild{"SiArr", &sffName.SiArr[i]}
    }
    sffName.EntityData.Leafs = make(map[string]types.YLeaf)
    sffName.EntityData.Leafs["name"] = types.YLeaf{"Name", sffName.Name}
    return &(sffName.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data
// Statistics data
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type_ interface{}

    // SFP stats.
    Sfp GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp

    // SPI SI stats.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi

    // Terminate stats.
    Term GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Term

    // Service function stats.
    Sf GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sf

    // Service function forwarder stats.
    Sff GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sff

    // Local service function forwarder stats.
    SffLocal GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal
}

func (data *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "sff-name"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = make(map[string]types.YChild)
    data.EntityData.Children["sfp"] = types.YChild{"Sfp", &data.Sfp}
    data.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &data.SpiSi}
    data.EntityData.Children["term"] = types.YChild{"Term", &data.Term}
    data.EntityData.Children["sf"] = types.YChild{"Sf", &data.Sf}
    data.EntityData.Children["sff"] = types.YChild{"Sff", &data.Sff}
    data.EntityData.Children["sff-local"] = types.YChild{"SffLocal", &data.SffLocal}
    data.EntityData.Leafs = make(map[string]types.YLeaf)
    data.EntityData.Leafs["type"] = types.YLeaf{"Type_", data.Type_}
    return &(data.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp
// SFP stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service index counters.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi

    // Terminate counters.
    Term GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term
}

func (sfp *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetEntityData() *types.CommonEntityData {
    sfp.EntityData.YFilter = sfp.YFilter
    sfp.EntityData.YangName = "sfp"
    sfp.EntityData.BundleName = "cisco_ios_xr"
    sfp.EntityData.ParentYangName = "data"
    sfp.EntityData.SegmentPath = "sfp"
    sfp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfp.EntityData.Children = make(map[string]types.YChild)
    sfp.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &sfp.SpiSi}
    sfp.EntityData.Children["term"] = types.YChild{"Term", &sfp.Term}
    sfp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sfp.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi
// Service index counters
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "sfp"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term
// Terminate counters
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "sfp"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi
// SPI SI stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "data"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Term
// Terminate stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "data"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sf
// Service function stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sf *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetEntityData() *types.CommonEntityData {
    sf.EntityData.YFilter = sf.YFilter
    sf.EntityData.YangName = "sf"
    sf.EntityData.BundleName = "cisco_ios_xr"
    sf.EntityData.ParentYangName = "data"
    sf.EntityData.SegmentPath = "sf"
    sf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sf.EntityData.Children = make(map[string]types.YChild)
    sf.EntityData.Leafs = make(map[string]types.YLeaf)
    sf.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", sf.ProcessedPkts}
    sf.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", sf.ProcessedBytes}
    return &(sf.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sff
// Service function forwarder stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sff *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetEntityData() *types.CommonEntityData {
    sff.EntityData.YFilter = sff.YFilter
    sff.EntityData.YangName = "sff"
    sff.EntityData.BundleName = "cisco_ios_xr"
    sff.EntityData.ParentYangName = "data"
    sff.EntityData.SegmentPath = "sff"
    sff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sff.EntityData.Children = make(map[string]types.YChild)
    sff.EntityData.Leafs = make(map[string]types.YLeaf)
    sff.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", sff.ProcessedPkts}
    sff.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", sff.ProcessedBytes}
    return &(sff.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal
// Local service function forwarder stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets with invalid NSH header. The type is interface{} with
    // range: 0..18446744073709551615.
    MalformedErrPkts interface{}

    // Number of packets with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615.
    LookupErrPkts interface{}

    // Total bytes with invalid NSH header. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    MalformedErrBytes interface{}

    // Total bytes with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    LookupErrBytes interface{}
}

func (sffLocal *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetEntityData() *types.CommonEntityData {
    sffLocal.EntityData.YFilter = sffLocal.YFilter
    sffLocal.EntityData.YangName = "sff-local"
    sffLocal.EntityData.BundleName = "cisco_ios_xr"
    sffLocal.EntityData.ParentYangName = "data"
    sffLocal.EntityData.SegmentPath = "sff-local"
    sffLocal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sffLocal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sffLocal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sffLocal.EntityData.Children = make(map[string]types.YChild)
    sffLocal.EntityData.Leafs = make(map[string]types.YLeaf)
    sffLocal.EntityData.Leafs["malformed-err-pkts"] = types.YLeaf{"MalformedErrPkts", sffLocal.MalformedErrPkts}
    sffLocal.EntityData.Leafs["lookup-err-pkts"] = types.YLeaf{"LookupErrPkts", sffLocal.LookupErrPkts}
    sffLocal.EntityData.Leafs["malformed-err-bytes"] = types.YLeaf{"MalformedErrBytes", sffLocal.MalformedErrBytes}
    sffLocal.EntityData.Leafs["lookup-err-bytes"] = types.YLeaf{"LookupErrBytes", sffLocal.LookupErrBytes}
    return &(sffLocal.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr
// SI array in case of detail stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service index. The type is interface{} with range: 0..255.
    Si interface{}

    // Stats counter for this index.
    Data GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data
}

func (siArr *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetEntityData() *types.CommonEntityData {
    siArr.EntityData.YFilter = siArr.YFilter
    siArr.EntityData.YangName = "si-arr"
    siArr.EntityData.BundleName = "cisco_ios_xr"
    siArr.EntityData.ParentYangName = "sff-name"
    siArr.EntityData.SegmentPath = "si-arr"
    siArr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    siArr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    siArr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    siArr.EntityData.Children = make(map[string]types.YChild)
    siArr.EntityData.Children["data"] = types.YChild{"Data", &siArr.Data}
    siArr.EntityData.Leafs = make(map[string]types.YLeaf)
    siArr.EntityData.Leafs["si"] = types.YLeaf{"Si", siArr.Si}
    return &(siArr.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data
// Stats counter for this index
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type_ interface{}

    // SF/SFF stats.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi

    // Terminate stats.
    Term GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term
}

func (data *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "si-arr"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = make(map[string]types.YChild)
    data.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &data.SpiSi}
    data.EntityData.Children["term"] = types.YChild{"Term", &data.Term}
    data.EntityData.Leafs = make(map[string]types.YLeaf)
    data.EntityData.Leafs["type"] = types.YLeaf{"Type_", data.Type_}
    return &(data.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi
// SF/SFF stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "data"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term
// Terminate stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "data"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local
// Local Service Function Forwarder operational
// data
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Error Statistics for local service function forwarder.
    Error GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error
}

func (local *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local) GetEntityData() *types.CommonEntityData {
    local.EntityData.YFilter = local.YFilter
    local.EntityData.YangName = "local"
    local.EntityData.BundleName = "cisco_ios_xr"
    local.EntityData.ParentYangName = "service-function-forwarder"
    local.EntityData.SegmentPath = "local"
    local.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    local.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    local.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    local.EntityData.Children = make(map[string]types.YChild)
    local.EntityData.Children["error"] = types.YChild{"Error", &local.Error}
    local.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(local.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error
// Error Statistics for local service function
// forwarder
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics data.
    Data GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data

    // SI array in case of detail stats. The type is slice of
    // GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr.
    SiArr []GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr
}

func (error *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "local"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["data"] = types.YChild{"Data", &error.Data}
    error.EntityData.Children["si-arr"] = types.YChild{"SiArr", nil}
    for i := range error.SiArr {
        error.EntityData.Children[types.GetSegmentPath(&error.SiArr[i])] = types.YChild{"SiArr", &error.SiArr[i]}
    }
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(error.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data
// Statistics data
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type_ interface{}

    // SFP stats.
    Sfp GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sfp

    // SPI SI stats.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_SpiSi

    // Terminate stats.
    Term GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Term

    // Service function stats.
    Sf GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sf

    // Service function forwarder stats.
    Sff GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sff

    // Local service function forwarder stats.
    SffLocal GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_SffLocal
}

func (data *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "error"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = make(map[string]types.YChild)
    data.EntityData.Children["sfp"] = types.YChild{"Sfp", &data.Sfp}
    data.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &data.SpiSi}
    data.EntityData.Children["term"] = types.YChild{"Term", &data.Term}
    data.EntityData.Children["sf"] = types.YChild{"Sf", &data.Sf}
    data.EntityData.Children["sff"] = types.YChild{"Sff", &data.Sff}
    data.EntityData.Children["sff-local"] = types.YChild{"SffLocal", &data.SffLocal}
    data.EntityData.Leafs = make(map[string]types.YLeaf)
    data.EntityData.Leafs["type"] = types.YLeaf{"Type_", data.Type_}
    return &(data.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sfp
// SFP stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sfp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service index counters.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi

    // Terminate counters.
    Term GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term
}

func (sfp *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetEntityData() *types.CommonEntityData {
    sfp.EntityData.YFilter = sfp.YFilter
    sfp.EntityData.YangName = "sfp"
    sfp.EntityData.BundleName = "cisco_ios_xr"
    sfp.EntityData.ParentYangName = "data"
    sfp.EntityData.SegmentPath = "sfp"
    sfp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfp.EntityData.Children = make(map[string]types.YChild)
    sfp.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &sfp.SpiSi}
    sfp.EntityData.Children["term"] = types.YChild{"Term", &sfp.Term}
    sfp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sfp.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi
// Service index counters
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "sfp"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term
// Terminate counters
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "sfp"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_SpiSi
// SPI SI stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "data"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Term
// Terminate stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "data"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sf
// Service function stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sf *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sf) GetEntityData() *types.CommonEntityData {
    sf.EntityData.YFilter = sf.YFilter
    sf.EntityData.YangName = "sf"
    sf.EntityData.BundleName = "cisco_ios_xr"
    sf.EntityData.ParentYangName = "data"
    sf.EntityData.SegmentPath = "sf"
    sf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sf.EntityData.Children = make(map[string]types.YChild)
    sf.EntityData.Leafs = make(map[string]types.YLeaf)
    sf.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", sf.ProcessedPkts}
    sf.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", sf.ProcessedBytes}
    return &(sf.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sff
// Service function forwarder stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sff *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_Sff) GetEntityData() *types.CommonEntityData {
    sff.EntityData.YFilter = sff.YFilter
    sff.EntityData.YangName = "sff"
    sff.EntityData.BundleName = "cisco_ios_xr"
    sff.EntityData.ParentYangName = "data"
    sff.EntityData.SegmentPath = "sff"
    sff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sff.EntityData.Children = make(map[string]types.YChild)
    sff.EntityData.Leafs = make(map[string]types.YLeaf)
    sff.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", sff.ProcessedPkts}
    sff.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", sff.ProcessedBytes}
    return &(sff.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_SffLocal
// Local service function forwarder stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_SffLocal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets with invalid NSH header. The type is interface{} with
    // range: 0..18446744073709551615.
    MalformedErrPkts interface{}

    // Number of packets with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615.
    LookupErrPkts interface{}

    // Total bytes with invalid NSH header. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    MalformedErrBytes interface{}

    // Total bytes with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    LookupErrBytes interface{}
}

func (sffLocal *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetEntityData() *types.CommonEntityData {
    sffLocal.EntityData.YFilter = sffLocal.YFilter
    sffLocal.EntityData.YangName = "sff-local"
    sffLocal.EntityData.BundleName = "cisco_ios_xr"
    sffLocal.EntityData.ParentYangName = "data"
    sffLocal.EntityData.SegmentPath = "sff-local"
    sffLocal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sffLocal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sffLocal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sffLocal.EntityData.Children = make(map[string]types.YChild)
    sffLocal.EntityData.Leafs = make(map[string]types.YLeaf)
    sffLocal.EntityData.Leafs["malformed-err-pkts"] = types.YLeaf{"MalformedErrPkts", sffLocal.MalformedErrPkts}
    sffLocal.EntityData.Leafs["lookup-err-pkts"] = types.YLeaf{"LookupErrPkts", sffLocal.LookupErrPkts}
    sffLocal.EntityData.Leafs["malformed-err-bytes"] = types.YLeaf{"MalformedErrBytes", sffLocal.MalformedErrBytes}
    sffLocal.EntityData.Leafs["lookup-err-bytes"] = types.YLeaf{"LookupErrBytes", sffLocal.LookupErrBytes}
    return &(sffLocal.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr
// SI array in case of detail stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service index. The type is interface{} with range: 0..255.
    Si interface{}

    // Stats counter for this index.
    Data GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr_Data
}

func (siArr *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr) GetEntityData() *types.CommonEntityData {
    siArr.EntityData.YFilter = siArr.YFilter
    siArr.EntityData.YangName = "si-arr"
    siArr.EntityData.BundleName = "cisco_ios_xr"
    siArr.EntityData.ParentYangName = "error"
    siArr.EntityData.SegmentPath = "si-arr"
    siArr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    siArr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    siArr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    siArr.EntityData.Children = make(map[string]types.YChild)
    siArr.EntityData.Children["data"] = types.YChild{"Data", &siArr.Data}
    siArr.EntityData.Leafs = make(map[string]types.YLeaf)
    siArr.EntityData.Leafs["si"] = types.YLeaf{"Si", siArr.Si}
    return &(siArr.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr_Data
// Stats counter for this index
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type_ interface{}

    // SF/SFF stats.
    SpiSi GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi

    // Terminate stats.
    Term GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term
}

func (data *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "si-arr"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = make(map[string]types.YChild)
    data.EntityData.Children["spi-si"] = types.YChild{"SpiSi", &data.SpiSi}
    data.EntityData.Children["term"] = types.YChild{"Term", &data.Term}
    data.EntityData.Leafs = make(map[string]types.YLeaf)
    data.EntityData.Leafs["type"] = types.YLeaf{"Type_", data.Type_}
    return &(data.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi
// SF/SFF stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetEntityData() *types.CommonEntityData {
    spiSi.EntityData.YFilter = spiSi.YFilter
    spiSi.EntityData.YangName = "spi-si"
    spiSi.EntityData.BundleName = "cisco_ios_xr"
    spiSi.EntityData.ParentYangName = "data"
    spiSi.EntityData.SegmentPath = "spi-si"
    spiSi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spiSi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spiSi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spiSi.EntityData.Children = make(map[string]types.YChild)
    spiSi.EntityData.Leafs = make(map[string]types.YLeaf)
    spiSi.EntityData.Leafs["processed-pkts"] = types.YLeaf{"ProcessedPkts", spiSi.ProcessedPkts}
    spiSi.EntityData.Leafs["processed-bytes"] = types.YLeaf{"ProcessedBytes", spiSi.ProcessedBytes}
    return &(spiSi.EntityData)
}

// GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term
// Terminate stats
type GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *GlobalServiceFunctionChaining_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetEntityData() *types.CommonEntityData {
    term.EntityData.YFilter = term.YFilter
    term.EntityData.YangName = "term"
    term.EntityData.BundleName = "cisco_ios_xr"
    term.EntityData.ParentYangName = "data"
    term.EntityData.SegmentPath = "term"
    term.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    term.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    term.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    term.EntityData.Children = make(map[string]types.YChild)
    term.EntityData.Leafs = make(map[string]types.YLeaf)
    term.EntityData.Leafs["terminated-pkts"] = types.YLeaf{"TerminatedPkts", term.TerminatedPkts}
    term.EntityData.Leafs["terminated-bytes"] = types.YLeaf{"TerminatedBytes", term.TerminatedBytes}
    return &(term.EntityData)
}

