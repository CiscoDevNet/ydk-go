// This module contains a collection of YANG definitions
// for Cisco IOS-XR traffmon-netflow package configuration.
// 
// This module contains definitions
// for the following management objects:
//   net-flow: NetFlow Configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package traffmon_netflow_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package traffmon_netflow_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-traffmon-netflow-cfg net-flow}", reflect.TypeOf(NetFlow{}))
    ydk.RegisterEntity("Cisco-IOS-XR-traffmon-netflow-cfg:net-flow", reflect.TypeOf(NetFlow{}))
}

// NfSamplingMode represents Nf sampling mode
type NfSamplingMode string

const (
    // Random sampling
    NfSamplingMode_random NfSamplingMode = "random"
)

// NfCacheAgingMode represents Nf cache aging mode
type NfCacheAgingMode string

const (
    // Normal, caches age
    NfCacheAgingMode_normal NfCacheAgingMode = "normal"

    // Permanent, caches never age
    NfCacheAgingMode_permanent NfCacheAgingMode = "permanent"

    // Immediate, caches age immediately
    NfCacheAgingMode_immediate NfCacheAgingMode = "immediate"
)

// NetFlow
// NetFlow Configuration
type NetFlow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure a flow exporter map.
    FlowExporterMaps NetFlow_FlowExporterMaps

    // Flow sampler map configuration.
    FlowSamplerMaps NetFlow_FlowSamplerMaps

    // Flow monitor map configuration.
    FlowMonitorMapTable NetFlow_FlowMonitorMapTable

    // Configure a performance traffic flow monitor map.
    FlowMonitorMapPerformanceTable NetFlow_FlowMonitorMapPerformanceTable
}

func (netFlow *NetFlow) GetFilter() yfilter.YFilter { return netFlow.YFilter }

func (netFlow *NetFlow) SetFilter(yf yfilter.YFilter) { netFlow.YFilter = yf }

func (netFlow *NetFlow) GetGoName(yname string) string {
    if yname == "flow-exporter-maps" { return "FlowExporterMaps" }
    if yname == "flow-sampler-maps" { return "FlowSamplerMaps" }
    if yname == "flow-monitor-map-table" { return "FlowMonitorMapTable" }
    if yname == "flow-monitor-map-performance-table" { return "FlowMonitorMapPerformanceTable" }
    return ""
}

func (netFlow *NetFlow) GetSegmentPath() string {
    return "Cisco-IOS-XR-traffmon-netflow-cfg:net-flow"
}

func (netFlow *NetFlow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-exporter-maps" {
        return &netFlow.FlowExporterMaps
    }
    if childYangName == "flow-sampler-maps" {
        return &netFlow.FlowSamplerMaps
    }
    if childYangName == "flow-monitor-map-table" {
        return &netFlow.FlowMonitorMapTable
    }
    if childYangName == "flow-monitor-map-performance-table" {
        return &netFlow.FlowMonitorMapPerformanceTable
    }
    return nil
}

func (netFlow *NetFlow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flow-exporter-maps"] = &netFlow.FlowExporterMaps
    children["flow-sampler-maps"] = &netFlow.FlowSamplerMaps
    children["flow-monitor-map-table"] = &netFlow.FlowMonitorMapTable
    children["flow-monitor-map-performance-table"] = &netFlow.FlowMonitorMapPerformanceTable
    return children
}

func (netFlow *NetFlow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (netFlow *NetFlow) GetBundleName() string { return "cisco_ios_xr" }

func (netFlow *NetFlow) GetYangName() string { return "net-flow" }

func (netFlow *NetFlow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (netFlow *NetFlow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (netFlow *NetFlow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (netFlow *NetFlow) SetParent(parent types.Entity) { netFlow.parent = parent }

func (netFlow *NetFlow) GetParent() types.Entity { return netFlow.parent }

func (netFlow *NetFlow) GetParentYangName() string { return "Cisco-IOS-XR-traffmon-netflow-cfg" }

// NetFlow_FlowExporterMaps
// Configure a flow exporter map
type NetFlow_FlowExporterMaps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Exporter map name. The type is slice of
    // NetFlow_FlowExporterMaps_FlowExporterMap.
    FlowExporterMap []NetFlow_FlowExporterMaps_FlowExporterMap
}

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetFilter() yfilter.YFilter { return flowExporterMaps.YFilter }

func (flowExporterMaps *NetFlow_FlowExporterMaps) SetFilter(yf yfilter.YFilter) { flowExporterMaps.YFilter = yf }

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetGoName(yname string) string {
    if yname == "flow-exporter-map" { return "FlowExporterMap" }
    return ""
}

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetSegmentPath() string {
    return "flow-exporter-maps"
}

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-exporter-map" {
        for _, c := range flowExporterMaps.FlowExporterMap {
            if flowExporterMaps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetFlow_FlowExporterMaps_FlowExporterMap{}
        flowExporterMaps.FlowExporterMap = append(flowExporterMaps.FlowExporterMap, child)
        return &flowExporterMaps.FlowExporterMap[len(flowExporterMaps.FlowExporterMap)-1]
    }
    return nil
}

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range flowExporterMaps.FlowExporterMap {
        children[flowExporterMaps.FlowExporterMap[i].GetSegmentPath()] = &flowExporterMaps.FlowExporterMap[i]
    }
    return children
}

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetBundleName() string { return "cisco_ios_xr" }

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetYangName() string { return "flow-exporter-maps" }

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowExporterMaps *NetFlow_FlowExporterMaps) SetParent(parent types.Entity) { flowExporterMaps.parent = parent }

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetParent() types.Entity { return flowExporterMaps.parent }

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetParentYangName() string { return "net-flow" }

// NetFlow_FlowExporterMaps_FlowExporterMap
// Exporter map name
type NetFlow_FlowExporterMaps_FlowExporterMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Exporter map name. The type is string with length:
    // 1..32.
    ExporterMapName interface{}

    // Configure source interface for collector. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Specify DSCP value for export packets. The type is interface{} with range:
    // 0..63.
    Dscp interface{}

    // Configure Maximum Value for Export Packet size. The type is interface{}
    // with range: 512..1468.
    PacketLength interface{}

    // Use UDP as transport protocol.
    Udp NetFlow_FlowExporterMaps_FlowExporterMap_Udp

    // Specify export version parameters.
    Versions NetFlow_FlowExporterMaps_FlowExporterMap_Versions

    // Configure export destination (collector).
    Destination NetFlow_FlowExporterMaps_FlowExporterMap_Destination
}

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetFilter() yfilter.YFilter { return flowExporterMap.YFilter }

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) SetFilter(yf yfilter.YFilter) { flowExporterMap.YFilter = yf }

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetGoName(yname string) string {
    if yname == "exporter-map-name" { return "ExporterMapName" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "dscp" { return "Dscp" }
    if yname == "packet-length" { return "PacketLength" }
    if yname == "udp" { return "Udp" }
    if yname == "versions" { return "Versions" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetSegmentPath() string {
    return "flow-exporter-map" + "[exporter-map-name='" + fmt.Sprintf("%v", flowExporterMap.ExporterMapName) + "']"
}

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udp" {
        return &flowExporterMap.Udp
    }
    if childYangName == "versions" {
        return &flowExporterMap.Versions
    }
    if childYangName == "destination" {
        return &flowExporterMap.Destination
    }
    return nil
}

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["udp"] = &flowExporterMap.Udp
    children["versions"] = &flowExporterMap.Versions
    children["destination"] = &flowExporterMap.Destination
    return children
}

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exporter-map-name"] = flowExporterMap.ExporterMapName
    leafs["source-interface"] = flowExporterMap.SourceInterface
    leafs["dscp"] = flowExporterMap.Dscp
    leafs["packet-length"] = flowExporterMap.PacketLength
    return leafs
}

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetBundleName() string { return "cisco_ios_xr" }

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetYangName() string { return "flow-exporter-map" }

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) SetParent(parent types.Entity) { flowExporterMap.parent = parent }

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetParent() types.Entity { return flowExporterMap.parent }

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetParentYangName() string { return "flow-exporter-maps" }

// NetFlow_FlowExporterMaps_FlowExporterMap_Udp
// Use UDP as transport protocol
type NetFlow_FlowExporterMaps_FlowExporterMap_Udp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Destination UDP port. The type is interface{} with range:
    // 1024..65535.
    DestinationPort interface{}
}

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetFilter() yfilter.YFilter { return udp.YFilter }

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) SetFilter(yf yfilter.YFilter) { udp.YFilter = yf }

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetGoName(yname string) string {
    if yname == "destination-port" { return "DestinationPort" }
    return ""
}

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetSegmentPath() string {
    return "udp"
}

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-port"] = udp.DestinationPort
    return leafs
}

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetBundleName() string { return "cisco_ios_xr" }

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetYangName() string { return "udp" }

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) SetParent(parent types.Entity) { udp.parent = parent }

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetParent() types.Entity { return udp.parent }

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetParentYangName() string { return "flow-exporter-map" }

// NetFlow_FlowExporterMaps_FlowExporterMap_Versions
// Specify export version parameters
type NetFlow_FlowExporterMaps_FlowExporterMap_Versions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure export version options. The type is slice of
    // NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version.
    Version []NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version
}

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) GetFilter() yfilter.YFilter { return versions.YFilter }

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) SetFilter(yf yfilter.YFilter) { versions.YFilter = yf }

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    return ""
}

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) GetSegmentPath() string {
    return "versions"
}

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "version" {
        for _, c := range versions.Version {
            if versions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version{}
        versions.Version = append(versions.Version, child)
        return &versions.Version[len(versions.Version)-1]
    }
    return nil
}

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range versions.Version {
        children[versions.Version[i].GetSegmentPath()] = &versions.Version[i]
    }
    return children
}

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) GetBundleName() string { return "cisco_ios_xr" }

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) GetYangName() string { return "versions" }

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) SetParent(parent types.Entity) { versions.parent = parent }

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) GetParent() types.Entity { return versions.parent }

func (versions *NetFlow_FlowExporterMaps_FlowExporterMap_Versions) GetParentYangName() string { return "flow-exporter-map" }

// NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version
// Configure export version options
type NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Export version number. The type is interface{}
    // with range: 9..10.
    VersionNumber interface{}

    // Option template configuration options. The type is interface{} with range:
    // 1..604800. Units are second. The default value is 1800.
    OptionsTemplateTimeout interface{}

    // Specify custom timeout for the template. The type is interface{} with
    // range: 1..604800. Units are second. The default value is 1800.
    CommonTemplateTimeout interface{}

    // Data template configuration options. The type is interface{} with range:
    // 1..604800. Units are second. The default value is 1800.
    DataTemplateTimeout interface{}

    // Specify options for exporting templates.
    Options NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options
}

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) GetFilter() yfilter.YFilter { return version.YFilter }

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) SetFilter(yf yfilter.YFilter) { version.YFilter = yf }

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) GetGoName(yname string) string {
    if yname == "version-number" { return "VersionNumber" }
    if yname == "options-template-timeout" { return "OptionsTemplateTimeout" }
    if yname == "common-template-timeout" { return "CommonTemplateTimeout" }
    if yname == "data-template-timeout" { return "DataTemplateTimeout" }
    if yname == "options" { return "Options" }
    return ""
}

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) GetSegmentPath() string {
    return "version" + "[version-number='" + fmt.Sprintf("%v", version.VersionNumber) + "']"
}

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "options" {
        return &version.Options
    }
    return nil
}

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["options"] = &version.Options
    return children
}

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version-number"] = version.VersionNumber
    leafs["options-template-timeout"] = version.OptionsTemplateTimeout
    leafs["common-template-timeout"] = version.CommonTemplateTimeout
    leafs["data-template-timeout"] = version.DataTemplateTimeout
    return leafs
}

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) GetBundleName() string { return "cisco_ios_xr" }

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) GetYangName() string { return "version" }

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) SetParent(parent types.Entity) { version.parent = parent }

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) GetParent() types.Entity { return version.parent }

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version) GetParentYangName() string { return "versions" }

// NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options
// Specify options for exporting templates
type NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify timeout for exporting interface table. The type is interface{} with
    // range: 0..604800. Units are second.
    InterfaceTableExportTimeout interface{}

    // Specify timeout for exporting sampler table. The type is interface{} with
    // range: 0..604800. Units are second.
    SamplerTableExportTimeout interface{}

    // Specify timeout for exporting vrf table. The type is interface{} with
    // range: 0..604800. Units are second.
    VrfTableExportTimeout interface{}
}

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) GetFilter() yfilter.YFilter { return options.YFilter }

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) SetFilter(yf yfilter.YFilter) { options.YFilter = yf }

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) GetGoName(yname string) string {
    if yname == "interface-table-export-timeout" { return "InterfaceTableExportTimeout" }
    if yname == "sampler-table-export-timeout" { return "SamplerTableExportTimeout" }
    if yname == "vrf-table-export-timeout" { return "VrfTableExportTimeout" }
    return ""
}

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) GetSegmentPath() string {
    return "options"
}

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-table-export-timeout"] = options.InterfaceTableExportTimeout
    leafs["sampler-table-export-timeout"] = options.SamplerTableExportTimeout
    leafs["vrf-table-export-timeout"] = options.VrfTableExportTimeout
    return leafs
}

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) GetBundleName() string { return "cisco_ios_xr" }

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) GetYangName() string { return "options" }

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) SetParent(parent types.Entity) { options.parent = parent }

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) GetParent() types.Entity { return options.parent }

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Versions_Version_Options) GetParentYangName() string { return "version" }

// NetFlow_FlowExporterMaps_FlowExporterMap_Destination
// Configure export destination (collector)
type NetFlow_FlowExporterMaps_FlowExporterMap_Destination struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // IPV6 address of the tunnel destination. The type is string.
    Ipv6Address interface{}

    // VRF name. The type is string.
    VrfName interface{}
}

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetFilter() yfilter.YFilter { return destination.YFilter }

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) SetFilter(yf yfilter.YFilter) { destination.YFilter = yf }

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetSegmentPath() string {
    return "destination"
}

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = destination.IpAddress
    leafs["ipv6-address"] = destination.Ipv6Address
    leafs["vrf-name"] = destination.VrfName
    return leafs
}

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetBundleName() string { return "cisco_ios_xr" }

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetYangName() string { return "destination" }

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) SetParent(parent types.Entity) { destination.parent = parent }

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetParent() types.Entity { return destination.parent }

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetParentYangName() string { return "flow-exporter-map" }

// NetFlow_FlowSamplerMaps
// Flow sampler map configuration
type NetFlow_FlowSamplerMaps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sampler map name. The type is slice of
    // NetFlow_FlowSamplerMaps_FlowSamplerMap.
    FlowSamplerMap []NetFlow_FlowSamplerMaps_FlowSamplerMap
}

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetFilter() yfilter.YFilter { return flowSamplerMaps.YFilter }

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) SetFilter(yf yfilter.YFilter) { flowSamplerMaps.YFilter = yf }

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetGoName(yname string) string {
    if yname == "flow-sampler-map" { return "FlowSamplerMap" }
    return ""
}

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetSegmentPath() string {
    return "flow-sampler-maps"
}

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-sampler-map" {
        for _, c := range flowSamplerMaps.FlowSamplerMap {
            if flowSamplerMaps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetFlow_FlowSamplerMaps_FlowSamplerMap{}
        flowSamplerMaps.FlowSamplerMap = append(flowSamplerMaps.FlowSamplerMap, child)
        return &flowSamplerMaps.FlowSamplerMap[len(flowSamplerMaps.FlowSamplerMap)-1]
    }
    return nil
}

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range flowSamplerMaps.FlowSamplerMap {
        children[flowSamplerMaps.FlowSamplerMap[i].GetSegmentPath()] = &flowSamplerMaps.FlowSamplerMap[i]
    }
    return children
}

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetBundleName() string { return "cisco_ios_xr" }

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetYangName() string { return "flow-sampler-maps" }

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) SetParent(parent types.Entity) { flowSamplerMaps.parent = parent }

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetParent() types.Entity { return flowSamplerMaps.parent }

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetParentYangName() string { return "net-flow" }

// NetFlow_FlowSamplerMaps_FlowSamplerMap
// Sampler map name
type NetFlow_FlowSamplerMaps_FlowSamplerMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sampler map name. The type is string with length:
    // 1..32.
    SamplerMapName interface{}

    // Configure packet sampling mode.
    SamplingModes NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes
}

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetFilter() yfilter.YFilter { return flowSamplerMap.YFilter }

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) SetFilter(yf yfilter.YFilter) { flowSamplerMap.YFilter = yf }

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetGoName(yname string) string {
    if yname == "sampler-map-name" { return "SamplerMapName" }
    if yname == "sampling-modes" { return "SamplingModes" }
    return ""
}

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetSegmentPath() string {
    return "flow-sampler-map" + "[sampler-map-name='" + fmt.Sprintf("%v", flowSamplerMap.SamplerMapName) + "']"
}

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sampling-modes" {
        return &flowSamplerMap.SamplingModes
    }
    return nil
}

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sampling-modes"] = &flowSamplerMap.SamplingModes
    return children
}

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sampler-map-name"] = flowSamplerMap.SamplerMapName
    return leafs
}

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetBundleName() string { return "cisco_ios_xr" }

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetYangName() string { return "flow-sampler-map" }

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) SetParent(parent types.Entity) { flowSamplerMap.parent = parent }

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetParent() types.Entity { return flowSamplerMap.parent }

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetParentYangName() string { return "flow-sampler-maps" }

// NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes
// Configure packet sampling mode
type NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure sampling mode. The type is slice of
    // NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode.
    SamplingMode []NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode
}

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetFilter() yfilter.YFilter { return samplingModes.YFilter }

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) SetFilter(yf yfilter.YFilter) { samplingModes.YFilter = yf }

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetGoName(yname string) string {
    if yname == "sampling-mode" { return "SamplingMode" }
    return ""
}

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetSegmentPath() string {
    return "sampling-modes"
}

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sampling-mode" {
        for _, c := range samplingModes.SamplingMode {
            if samplingModes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode{}
        samplingModes.SamplingMode = append(samplingModes.SamplingMode, child)
        return &samplingModes.SamplingMode[len(samplingModes.SamplingMode)-1]
    }
    return nil
}

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range samplingModes.SamplingMode {
        children[samplingModes.SamplingMode[i].GetSegmentPath()] = &samplingModes.SamplingMode[i]
    }
    return children
}

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetBundleName() string { return "cisco_ios_xr" }

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetYangName() string { return "sampling-modes" }

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) SetParent(parent types.Entity) { samplingModes.parent = parent }

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetParent() types.Entity { return samplingModes.parent }

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetParentYangName() string { return "flow-sampler-map" }

// NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode
// Configure sampling mode
type NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sampling mode. The type is NfSamplingMode.
    Mode interface{}

    // Number of packets to be sampled in the sampling interval. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    SampleNumber interface{}

    // Sampling interval in units of packets. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    Interval interface{}
}

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetFilter() yfilter.YFilter { return samplingMode.YFilter }

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) SetFilter(yf yfilter.YFilter) { samplingMode.YFilter = yf }

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetGoName(yname string) string {
    if yname == "mode" { return "Mode" }
    if yname == "sample-number" { return "SampleNumber" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetSegmentPath() string {
    return "sampling-mode" + "[mode='" + fmt.Sprintf("%v", samplingMode.Mode) + "']"
}

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mode"] = samplingMode.Mode
    leafs["sample-number"] = samplingMode.SampleNumber
    leafs["interval"] = samplingMode.Interval
    return leafs
}

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetBundleName() string { return "cisco_ios_xr" }

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetYangName() string { return "sampling-mode" }

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) SetParent(parent types.Entity) { samplingMode.parent = parent }

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetParent() types.Entity { return samplingMode.parent }

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetParentYangName() string { return "sampling-modes" }

// NetFlow_FlowMonitorMapTable
// Flow monitor map configuration
type NetFlow_FlowMonitorMapTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Monitor map name. The type is slice of
    // NetFlow_FlowMonitorMapTable_FlowMonitorMap.
    FlowMonitorMap []NetFlow_FlowMonitorMapTable_FlowMonitorMap
}

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetFilter() yfilter.YFilter { return flowMonitorMapTable.YFilter }

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) SetFilter(yf yfilter.YFilter) { flowMonitorMapTable.YFilter = yf }

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetGoName(yname string) string {
    if yname == "flow-monitor-map" { return "FlowMonitorMap" }
    return ""
}

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetSegmentPath() string {
    return "flow-monitor-map-table"
}

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-monitor-map" {
        for _, c := range flowMonitorMapTable.FlowMonitorMap {
            if flowMonitorMapTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetFlow_FlowMonitorMapTable_FlowMonitorMap{}
        flowMonitorMapTable.FlowMonitorMap = append(flowMonitorMapTable.FlowMonitorMap, child)
        return &flowMonitorMapTable.FlowMonitorMap[len(flowMonitorMapTable.FlowMonitorMap)-1]
    }
    return nil
}

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range flowMonitorMapTable.FlowMonitorMap {
        children[flowMonitorMapTable.FlowMonitorMap[i].GetSegmentPath()] = &flowMonitorMapTable.FlowMonitorMap[i]
    }
    return children
}

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetBundleName() string { return "cisco_ios_xr" }

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetYangName() string { return "flow-monitor-map-table" }

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) SetParent(parent types.Entity) { flowMonitorMapTable.parent = parent }

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetParent() types.Entity { return flowMonitorMapTable.parent }

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetParentYangName() string { return "net-flow" }

// NetFlow_FlowMonitorMapTable_FlowMonitorMap
// Monitor map name
type NetFlow_FlowMonitorMapTable_FlowMonitorMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Monitor map name. The type is string with length:
    // 1..32.
    MonitorMapName interface{}

    // Specify the update flow cache aging timeout. The type is interface{} with
    // range: 1..604800. Units are second. The default value is 1800.
    CacheUpdateAgingTimeout interface{}

    // Specify the number of entries in the flow cache. The type is interface{}
    // with range: 4096..1000000. The default value is 65535.
    CacheEntries interface{}

    // Specify the inactive flow cache aging timeout. The type is interface{} with
    // range: 0..604800. Units are second. The default value is 15.
    CacheInactiveAgingTimeout interface{}

    // Specify the active flow cache aging timeout. The type is interface{} with
    // range: 1..604800. Units are second. The default value is 1800.
    CacheActiveAgingTimeout interface{}

    // Specify the maximum number of entries to age each second. The type is
    // interface{} with range: 1..1000000. The default value is 2000.
    CacheTimeoutRateLimit interface{}

    // Specify the flow cache aging mode. The type is NfCacheAgingMode. The
    // default value is normal.
    CacheAgingMode interface{}

    // Specify an option for the flow cache.
    Option NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option

    // Configure exporters to be used by the monitor-map.
    Exporters NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters

    // Specify a flow record format.
    Record NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record
}

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetFilter() yfilter.YFilter { return flowMonitorMap.YFilter }

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) SetFilter(yf yfilter.YFilter) { flowMonitorMap.YFilter = yf }

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetGoName(yname string) string {
    if yname == "monitor-map-name" { return "MonitorMapName" }
    if yname == "cache-update-aging-timeout" { return "CacheUpdateAgingTimeout" }
    if yname == "cache-entries" { return "CacheEntries" }
    if yname == "cache-inactive-aging-timeout" { return "CacheInactiveAgingTimeout" }
    if yname == "cache-active-aging-timeout" { return "CacheActiveAgingTimeout" }
    if yname == "cache-timeout-rate-limit" { return "CacheTimeoutRateLimit" }
    if yname == "cache-aging-mode" { return "CacheAgingMode" }
    if yname == "option" { return "Option" }
    if yname == "exporters" { return "Exporters" }
    if yname == "record" { return "Record" }
    return ""
}

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetSegmentPath() string {
    return "flow-monitor-map" + "[monitor-map-name='" + fmt.Sprintf("%v", flowMonitorMap.MonitorMapName) + "']"
}

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option" {
        return &flowMonitorMap.Option
    }
    if childYangName == "exporters" {
        return &flowMonitorMap.Exporters
    }
    if childYangName == "record" {
        return &flowMonitorMap.Record
    }
    return nil
}

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["option"] = &flowMonitorMap.Option
    children["exporters"] = &flowMonitorMap.Exporters
    children["record"] = &flowMonitorMap.Record
    return children
}

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["monitor-map-name"] = flowMonitorMap.MonitorMapName
    leafs["cache-update-aging-timeout"] = flowMonitorMap.CacheUpdateAgingTimeout
    leafs["cache-entries"] = flowMonitorMap.CacheEntries
    leafs["cache-inactive-aging-timeout"] = flowMonitorMap.CacheInactiveAgingTimeout
    leafs["cache-active-aging-timeout"] = flowMonitorMap.CacheActiveAgingTimeout
    leafs["cache-timeout-rate-limit"] = flowMonitorMap.CacheTimeoutRateLimit
    leafs["cache-aging-mode"] = flowMonitorMap.CacheAgingMode
    return leafs
}

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetBundleName() string { return "cisco_ios_xr" }

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetYangName() string { return "flow-monitor-map" }

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) SetParent(parent types.Entity) { flowMonitorMap.parent = parent }

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetParent() types.Entity { return flowMonitorMap.parent }

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetParentYangName() string { return "flow-monitor-map-table" }

// NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option
// Specify an option for the flow cache
type NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify whether data should be filtered. The type is interface{}.
    Filtered interface{}

    // Specify whether to export physical ifh for bundle interface. The type is
    // interface{}.
    OutBundleMember interface{}

    // Specify whether it exports the physical output interface. The type is
    // interface{}.
    OutPhysInt interface{}

    // Specify if BGP Attributes AS_PATH STD_COMM should be exported. The type is
    // interface{}.
    BgpAttr interface{}
}

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetFilter() yfilter.YFilter { return option.YFilter }

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) SetFilter(yf yfilter.YFilter) { option.YFilter = yf }

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetGoName(yname string) string {
    if yname == "filtered" { return "Filtered" }
    if yname == "out-bundle-member" { return "OutBundleMember" }
    if yname == "out-phys-int" { return "OutPhysInt" }
    if yname == "bgp-attr" { return "BgpAttr" }
    return ""
}

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetSegmentPath() string {
    return "option"
}

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["filtered"] = option.Filtered
    leafs["out-bundle-member"] = option.OutBundleMember
    leafs["out-phys-int"] = option.OutPhysInt
    leafs["bgp-attr"] = option.BgpAttr
    return leafs
}

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetBundleName() string { return "cisco_ios_xr" }

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetYangName() string { return "option" }

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) SetParent(parent types.Entity) { option.parent = parent }

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetParent() types.Entity { return option.parent }

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetParentYangName() string { return "flow-monitor-map" }

// NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters
// Configure exporters to be used by the
// monitor-map
type NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure exporter to be used by the monitor-map. The type is slice of
    // NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter.
    Exporter []NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter
}

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetFilter() yfilter.YFilter { return exporters.YFilter }

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) SetFilter(yf yfilter.YFilter) { exporters.YFilter = yf }

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetGoName(yname string) string {
    if yname == "exporter" { return "Exporter" }
    return ""
}

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetSegmentPath() string {
    return "exporters"
}

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "exporter" {
        for _, c := range exporters.Exporter {
            if exporters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter{}
        exporters.Exporter = append(exporters.Exporter, child)
        return &exporters.Exporter[len(exporters.Exporter)-1]
    }
    return nil
}

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range exporters.Exporter {
        children[exporters.Exporter[i].GetSegmentPath()] = &exporters.Exporter[i]
    }
    return children
}

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetBundleName() string { return "cisco_ios_xr" }

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetYangName() string { return "exporters" }

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) SetParent(parent types.Entity) { exporters.parent = parent }

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetParent() types.Entity { return exporters.parent }

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetParentYangName() string { return "flow-monitor-map" }

// NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter
// Configure exporter to be used by the
// monitor-map
type NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Exporter name. The type is string with length:
    // 1..32.
    ExporterName interface{}
}

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetFilter() yfilter.YFilter { return exporter.YFilter }

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) SetFilter(yf yfilter.YFilter) { exporter.YFilter = yf }

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetGoName(yname string) string {
    if yname == "exporter-name" { return "ExporterName" }
    return ""
}

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetSegmentPath() string {
    return "exporter" + "[exporter-name='" + fmt.Sprintf("%v", exporter.ExporterName) + "']"
}

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exporter-name"] = exporter.ExporterName
    return leafs
}

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetBundleName() string { return "cisco_ios_xr" }

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetYangName() string { return "exporter" }

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) SetParent(parent types.Entity) { exporter.parent = parent }

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetParent() types.Entity { return exporter.parent }

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetParentYangName() string { return "exporters" }

// NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record
// Specify a flow record format
// This type is a presence type.
type NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flow record format (Either 'ipv4-raw' ,'ipv4-peer-as', 'ipv6', 'mpls',
    // 'mpls-ipv4', 'mpls-ipv6', 'mpls-ipv4-ipv6', 'ipv6-peer-as'). The type is
    // string with length: 1..32. This attribute is mandatory.
    RecordName interface{}

    // Enter label value for MPLS record type. The type is interface{} with range:
    // 1..6.
    Label interface{}
}

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetFilter() yfilter.YFilter { return record.YFilter }

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) SetFilter(yf yfilter.YFilter) { record.YFilter = yf }

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetGoName(yname string) string {
    if yname == "record-name" { return "RecordName" }
    if yname == "label" { return "Label" }
    return ""
}

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetSegmentPath() string {
    return "record"
}

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["record-name"] = record.RecordName
    leafs["label"] = record.Label
    return leafs
}

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetBundleName() string { return "cisco_ios_xr" }

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetYangName() string { return "record" }

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) SetParent(parent types.Entity) { record.parent = parent }

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetParent() types.Entity { return record.parent }

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetParentYangName() string { return "flow-monitor-map" }

// NetFlow_FlowMonitorMapPerformanceTable
// Configure a performance traffic flow monitor map
type NetFlow_FlowMonitorMapPerformanceTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Monitor map name. The type is slice of
    // NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap.
    FlowMonitorMap []NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap
}

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetFilter() yfilter.YFilter { return flowMonitorMapPerformanceTable.YFilter }

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) SetFilter(yf yfilter.YFilter) { flowMonitorMapPerformanceTable.YFilter = yf }

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetGoName(yname string) string {
    if yname == "flow-monitor-map" { return "FlowMonitorMap" }
    return ""
}

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetSegmentPath() string {
    return "flow-monitor-map-performance-table"
}

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-monitor-map" {
        for _, c := range flowMonitorMapPerformanceTable.FlowMonitorMap {
            if flowMonitorMapPerformanceTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap{}
        flowMonitorMapPerformanceTable.FlowMonitorMap = append(flowMonitorMapPerformanceTable.FlowMonitorMap, child)
        return &flowMonitorMapPerformanceTable.FlowMonitorMap[len(flowMonitorMapPerformanceTable.FlowMonitorMap)-1]
    }
    return nil
}

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range flowMonitorMapPerformanceTable.FlowMonitorMap {
        children[flowMonitorMapPerformanceTable.FlowMonitorMap[i].GetSegmentPath()] = &flowMonitorMapPerformanceTable.FlowMonitorMap[i]
    }
    return children
}

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetBundleName() string { return "cisco_ios_xr" }

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetYangName() string { return "flow-monitor-map-performance-table" }

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) SetParent(parent types.Entity) { flowMonitorMapPerformanceTable.parent = parent }

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetParent() types.Entity { return flowMonitorMapPerformanceTable.parent }

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetParentYangName() string { return "net-flow" }

// NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap
// Monitor map name
type NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Monitor map name. The type is string with length:
    // 1..32.
    MonitorMapName interface{}

    // Specify the update flow cache aging timeout. The type is interface{} with
    // range: 1..604800. Units are second. The default value is 1800.
    CacheUpdateAgingTimeout interface{}

    // Specify the number of entries in the flow cache. The type is interface{}
    // with range: 4096..1000000. The default value is 65535.
    CacheEntries interface{}

    // Specify the inactive flow cache aging timeout. The type is interface{} with
    // range: 0..604800. Units are second. The default value is 15.
    CacheInactiveAgingTimeout interface{}

    // Specify the active flow cache aging timeout. The type is interface{} with
    // range: 1..604800. Units are second. The default value is 1800.
    CacheActiveAgingTimeout interface{}

    // Specify the maximum number of entries to age each second. The type is
    // interface{} with range: 1..1000000. The default value is 2000.
    CacheTimeoutRateLimit interface{}

    // Specify the flow cache aging mode. The type is NfCacheAgingMode. The
    // default value is normal.
    CacheAgingMode interface{}

    // Specify an option for the flow cache.
    Option NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option

    // Configure exporters to be used by the monitor-map.
    Exporters NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters

    // Specify a flow record format.
    Record NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record
}

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetFilter() yfilter.YFilter { return flowMonitorMap.YFilter }

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) SetFilter(yf yfilter.YFilter) { flowMonitorMap.YFilter = yf }

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetGoName(yname string) string {
    if yname == "monitor-map-name" { return "MonitorMapName" }
    if yname == "cache-update-aging-timeout" { return "CacheUpdateAgingTimeout" }
    if yname == "cache-entries" { return "CacheEntries" }
    if yname == "cache-inactive-aging-timeout" { return "CacheInactiveAgingTimeout" }
    if yname == "cache-active-aging-timeout" { return "CacheActiveAgingTimeout" }
    if yname == "cache-timeout-rate-limit" { return "CacheTimeoutRateLimit" }
    if yname == "cache-aging-mode" { return "CacheAgingMode" }
    if yname == "option" { return "Option" }
    if yname == "exporters" { return "Exporters" }
    if yname == "record" { return "Record" }
    return ""
}

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetSegmentPath() string {
    return "flow-monitor-map" + "[monitor-map-name='" + fmt.Sprintf("%v", flowMonitorMap.MonitorMapName) + "']"
}

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option" {
        return &flowMonitorMap.Option
    }
    if childYangName == "exporters" {
        return &flowMonitorMap.Exporters
    }
    if childYangName == "record" {
        return &flowMonitorMap.Record
    }
    return nil
}

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["option"] = &flowMonitorMap.Option
    children["exporters"] = &flowMonitorMap.Exporters
    children["record"] = &flowMonitorMap.Record
    return children
}

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["monitor-map-name"] = flowMonitorMap.MonitorMapName
    leafs["cache-update-aging-timeout"] = flowMonitorMap.CacheUpdateAgingTimeout
    leafs["cache-entries"] = flowMonitorMap.CacheEntries
    leafs["cache-inactive-aging-timeout"] = flowMonitorMap.CacheInactiveAgingTimeout
    leafs["cache-active-aging-timeout"] = flowMonitorMap.CacheActiveAgingTimeout
    leafs["cache-timeout-rate-limit"] = flowMonitorMap.CacheTimeoutRateLimit
    leafs["cache-aging-mode"] = flowMonitorMap.CacheAgingMode
    return leafs
}

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetBundleName() string { return "cisco_ios_xr" }

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetYangName() string { return "flow-monitor-map" }

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) SetParent(parent types.Entity) { flowMonitorMap.parent = parent }

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetParent() types.Entity { return flowMonitorMap.parent }

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetParentYangName() string { return "flow-monitor-map-performance-table" }

// NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option
// Specify an option for the flow cache
type NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify whether data should be filtered. The type is interface{}.
    Filtered interface{}

    // Specify whether to export physical ifh for bundle interface. The type is
    // interface{}.
    OutBundleMember interface{}

    // Specify whether it exports the physical output interface. The type is
    // interface{}.
    OutPhysInt interface{}

    // Specify if BGP Attributes AS_PATH STD_COMM should be exported. The type is
    // interface{}.
    BgpAttr interface{}
}

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetFilter() yfilter.YFilter { return option.YFilter }

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) SetFilter(yf yfilter.YFilter) { option.YFilter = yf }

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetGoName(yname string) string {
    if yname == "filtered" { return "Filtered" }
    if yname == "out-bundle-member" { return "OutBundleMember" }
    if yname == "out-phys-int" { return "OutPhysInt" }
    if yname == "bgp-attr" { return "BgpAttr" }
    return ""
}

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetSegmentPath() string {
    return "option"
}

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["filtered"] = option.Filtered
    leafs["out-bundle-member"] = option.OutBundleMember
    leafs["out-phys-int"] = option.OutPhysInt
    leafs["bgp-attr"] = option.BgpAttr
    return leafs
}

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetBundleName() string { return "cisco_ios_xr" }

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetYangName() string { return "option" }

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) SetParent(parent types.Entity) { option.parent = parent }

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetParent() types.Entity { return option.parent }

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetParentYangName() string { return "flow-monitor-map" }

// NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters
// Configure exporters to be used by the
// monitor-map
type NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure exporter to be used by the monitor-map. The type is slice of
    // NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter.
    Exporter []NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter
}

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetFilter() yfilter.YFilter { return exporters.YFilter }

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) SetFilter(yf yfilter.YFilter) { exporters.YFilter = yf }

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetGoName(yname string) string {
    if yname == "exporter" { return "Exporter" }
    return ""
}

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetSegmentPath() string {
    return "exporters"
}

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "exporter" {
        for _, c := range exporters.Exporter {
            if exporters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter{}
        exporters.Exporter = append(exporters.Exporter, child)
        return &exporters.Exporter[len(exporters.Exporter)-1]
    }
    return nil
}

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range exporters.Exporter {
        children[exporters.Exporter[i].GetSegmentPath()] = &exporters.Exporter[i]
    }
    return children
}

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetBundleName() string { return "cisco_ios_xr" }

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetYangName() string { return "exporters" }

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) SetParent(parent types.Entity) { exporters.parent = parent }

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetParent() types.Entity { return exporters.parent }

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetParentYangName() string { return "flow-monitor-map" }

// NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter
// Configure exporter to be used by the
// monitor-map
type NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Exporter name. The type is string with length:
    // 1..32.
    ExporterName interface{}
}

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetFilter() yfilter.YFilter { return exporter.YFilter }

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) SetFilter(yf yfilter.YFilter) { exporter.YFilter = yf }

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetGoName(yname string) string {
    if yname == "exporter-name" { return "ExporterName" }
    return ""
}

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetSegmentPath() string {
    return "exporter" + "[exporter-name='" + fmt.Sprintf("%v", exporter.ExporterName) + "']"
}

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exporter-name"] = exporter.ExporterName
    return leafs
}

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetBundleName() string { return "cisco_ios_xr" }

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetYangName() string { return "exporter" }

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) SetParent(parent types.Entity) { exporter.parent = parent }

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetParent() types.Entity { return exporter.parent }

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetParentYangName() string { return "exporters" }

// NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record
// Specify a flow record format
// This type is a presence type.
type NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flow record format (Either 'ipv4-raw' ,'ipv4-peer-as', 'ipv6', 'mpls',
    // 'mpls-ipv4', 'mpls-ipv6', 'mpls-ipv4-ipv6', 'ipv6-peer-as'). The type is
    // string with length: 1..32. This attribute is mandatory.
    RecordName interface{}

    // Enter label value for MPLS record type. The type is interface{} with range:
    // 1..6.
    Label interface{}
}

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetFilter() yfilter.YFilter { return record.YFilter }

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) SetFilter(yf yfilter.YFilter) { record.YFilter = yf }

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetGoName(yname string) string {
    if yname == "record-name" { return "RecordName" }
    if yname == "label" { return "Label" }
    return ""
}

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetSegmentPath() string {
    return "record"
}

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["record-name"] = record.RecordName
    leafs["label"] = record.Label
    return leafs
}

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetBundleName() string { return "cisco_ios_xr" }

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetYangName() string { return "record" }

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) SetParent(parent types.Entity) { record.parent = parent }

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetParent() types.Entity { return record.parent }

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetParentYangName() string { return "flow-monitor-map" }

