// This module contains a collection of YANG definitions
// for Cisco IOS-XR vservice package configuration.
// 
// This module contains definitions
// for the following management objects:
//   vservice: configure vservice node
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package vservice_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package vservice_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-vservice-cfg vservice}", reflect.TypeOf(Vservice{}))
    ydk.RegisterEntity("Cisco-IOS-XR-vservice-cfg:vservice", reflect.TypeOf(Vservice{}))
}

// SfcSfTransport represents Sfc sf transport
type SfcSfTransport string

const (
    // vxlan-gpe transport type
    SfcSfTransport_vxlan_gpe SfcSfTransport = "vxlan-gpe"
)

// SfcMetadataDispositionMatch represents Sfc metadata disposition match
type SfcMetadataDispositionMatch string

const (
    // match type 1
    SfcMetadataDispositionMatch_type1_dcalloc_tenant_id SfcMetadataDispositionMatch = "type1-dcalloc-tenant-id"
)

// SfcMetadataAlloc represents Sfc metadata alloc
type SfcMetadataAlloc string

const (
    // type 1 allocation
    SfcMetadataAlloc_type1 SfcMetadataAlloc = "type1"
)

// SfcMetadataType1AllocFormat represents Sfc metadata type1 alloc format
type SfcMetadataType1AllocFormat string

const (
    // data center allocation
    SfcMetadataType1AllocFormat_dc_allocation SfcMetadataType1AllocFormat = "dc-allocation"
)

// SfcMetadataDispositionAction represents Sfc metadata disposition action
type SfcMetadataDispositionAction string

const (
    // redirect nexthop action
    SfcMetadataDispositionAction_redirect_nexthop SfcMetadataDispositionAction = "redirect-nexthop"
)

// Vservice
// configure vservice node
type Vservice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // configure service function locator.
    ServiceFunctionLocator Vservice_ServiceFunctionLocator

    // Configure metadata disposition.
    MetadataDispositions Vservice_MetadataDispositions

    // configure service function forward locator.
    ServiceFunctionForwardLocator Vservice_ServiceFunctionForwardLocator

    // configure metadata imposition.
    MetadataTemplates Vservice_MetadataTemplates

    // service function path.
    ServiceFunctionPath Vservice_ServiceFunctionPath
}

func (vservice *Vservice) GetFilter() yfilter.YFilter { return vservice.YFilter }

func (vservice *Vservice) SetFilter(yf yfilter.YFilter) { vservice.YFilter = yf }

func (vservice *Vservice) GetGoName(yname string) string {
    if yname == "service-function-locator" { return "ServiceFunctionLocator" }
    if yname == "metadata-dispositions" { return "MetadataDispositions" }
    if yname == "service-function-forward-locator" { return "ServiceFunctionForwardLocator" }
    if yname == "metadata-templates" { return "MetadataTemplates" }
    if yname == "service-function-path" { return "ServiceFunctionPath" }
    return ""
}

func (vservice *Vservice) GetSegmentPath() string {
    return "Cisco-IOS-XR-vservice-cfg:vservice"
}

func (vservice *Vservice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-function-locator" {
        return &vservice.ServiceFunctionLocator
    }
    if childYangName == "metadata-dispositions" {
        return &vservice.MetadataDispositions
    }
    if childYangName == "service-function-forward-locator" {
        return &vservice.ServiceFunctionForwardLocator
    }
    if childYangName == "metadata-templates" {
        return &vservice.MetadataTemplates
    }
    if childYangName == "service-function-path" {
        return &vservice.ServiceFunctionPath
    }
    return nil
}

func (vservice *Vservice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-function-locator"] = &vservice.ServiceFunctionLocator
    children["metadata-dispositions"] = &vservice.MetadataDispositions
    children["service-function-forward-locator"] = &vservice.ServiceFunctionForwardLocator
    children["metadata-templates"] = &vservice.MetadataTemplates
    children["service-function-path"] = &vservice.ServiceFunctionPath
    return children
}

func (vservice *Vservice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vservice *Vservice) GetBundleName() string { return "cisco_ios_xr" }

func (vservice *Vservice) GetYangName() string { return "vservice" }

func (vservice *Vservice) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vservice *Vservice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vservice *Vservice) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vservice *Vservice) SetParent(parent types.Entity) { vservice.parent = parent }

func (vservice *Vservice) GetParent() types.Entity { return vservice.parent }

func (vservice *Vservice) GetParentYangName() string { return "Cisco-IOS-XR-vservice-cfg" }

// Vservice_ServiceFunctionLocator
// configure service function locator
type Vservice_ServiceFunctionLocator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mention the sf/sff name.
    Names Vservice_ServiceFunctionLocator_Names
}

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetFilter() yfilter.YFilter { return serviceFunctionLocator.YFilter }

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) SetFilter(yf yfilter.YFilter) { serviceFunctionLocator.YFilter = yf }

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetGoName(yname string) string {
    if yname == "names" { return "Names" }
    return ""
}

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetSegmentPath() string {
    return "service-function-locator"
}

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "names" {
        return &serviceFunctionLocator.Names
    }
    return nil
}

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["names"] = &serviceFunctionLocator.Names
    return children
}

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetBundleName() string { return "cisco_ios_xr" }

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetYangName() string { return "service-function-locator" }

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) SetParent(parent types.Entity) { serviceFunctionLocator.parent = parent }

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetParent() types.Entity { return serviceFunctionLocator.parent }

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetParentYangName() string { return "vservice" }

// Vservice_ServiceFunctionLocator_Names
// Mention the sf/sff name
type Vservice_ServiceFunctionLocator_Names struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // service function name. The type is slice of
    // Vservice_ServiceFunctionLocator_Names_Name.
    Name []Vservice_ServiceFunctionLocator_Names_Name
}

func (names *Vservice_ServiceFunctionLocator_Names) GetFilter() yfilter.YFilter { return names.YFilter }

func (names *Vservice_ServiceFunctionLocator_Names) SetFilter(yf yfilter.YFilter) { names.YFilter = yf }

func (names *Vservice_ServiceFunctionLocator_Names) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (names *Vservice_ServiceFunctionLocator_Names) GetSegmentPath() string {
    return "names"
}

func (names *Vservice_ServiceFunctionLocator_Names) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "name" {
        for _, c := range names.Name {
            if names.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vservice_ServiceFunctionLocator_Names_Name{}
        names.Name = append(names.Name, child)
        return &names.Name[len(names.Name)-1]
    }
    return nil
}

func (names *Vservice_ServiceFunctionLocator_Names) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range names.Name {
        children[names.Name[i].GetSegmentPath()] = &names.Name[i]
    }
    return children
}

func (names *Vservice_ServiceFunctionLocator_Names) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (names *Vservice_ServiceFunctionLocator_Names) GetBundleName() string { return "cisco_ios_xr" }

func (names *Vservice_ServiceFunctionLocator_Names) GetYangName() string { return "names" }

func (names *Vservice_ServiceFunctionLocator_Names) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (names *Vservice_ServiceFunctionLocator_Names) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (names *Vservice_ServiceFunctionLocator_Names) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (names *Vservice_ServiceFunctionLocator_Names) SetParent(parent types.Entity) { names.parent = parent }

func (names *Vservice_ServiceFunctionLocator_Names) GetParent() types.Entity { return names.parent }

func (names *Vservice_ServiceFunctionLocator_Names) GetParentYangName() string { return "service-function-locator" }

// Vservice_ServiceFunctionLocator_Names_Name
// service function name
type Vservice_ServiceFunctionLocator_Names_Name struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Service function/forwarder name. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    FunctionName interface{}

    // This attribute is a key. Specify locator id. The type is interface{} with
    // range: 1..255.
    LocatorId interface{}

    // configure sff/sffl.
    Node Vservice_ServiceFunctionLocator_Names_Name_Node
}

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetFilter() yfilter.YFilter { return name.YFilter }

func (name *Vservice_ServiceFunctionLocator_Names_Name) SetFilter(yf yfilter.YFilter) { name.YFilter = yf }

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetGoName(yname string) string {
    if yname == "function-name" { return "FunctionName" }
    if yname == "locator-id" { return "LocatorId" }
    if yname == "node" { return "Node" }
    return ""
}

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetSegmentPath() string {
    return "name" + "[function-name='" + fmt.Sprintf("%v", name.FunctionName) + "']" + "[locator-id='" + fmt.Sprintf("%v", name.LocatorId) + "']"
}

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        return &name.Node
    }
    return nil
}

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node"] = &name.Node
    return children
}

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["function-name"] = name.FunctionName
    leafs["locator-id"] = name.LocatorId
    return leafs
}

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetBundleName() string { return "cisco_ios_xr" }

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetYangName() string { return "name" }

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (name *Vservice_ServiceFunctionLocator_Names_Name) SetParent(parent types.Entity) { name.parent = parent }

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetParent() types.Entity { return name.parent }

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetParentYangName() string { return "names" }

// Vservice_ServiceFunctionLocator_Names_Name_Node
// configure sff/sffl
type Vservice_ServiceFunctionLocator_Names_Name_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transport type. The type is SfcSfTransport.
    Transport interface{}

    // IPv4 source address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4SourceAddress interface{}

    // IPv4 destination address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4DestinationAddress interface{}

    // VNI. The type is interface{} with range: -2147483648..2147483647.
    Vni interface{}
}

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetGoName(yname string) string {
    if yname == "transport" { return "Transport" }
    if yname == "ipv4-source-address" { return "Ipv4SourceAddress" }
    if yname == "ipv4-destination-address" { return "Ipv4DestinationAddress" }
    if yname == "vni" { return "Vni" }
    return ""
}

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetSegmentPath() string {
    return "node"
}

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transport"] = node.Transport
    leafs["ipv4-source-address"] = node.Ipv4SourceAddress
    leafs["ipv4-destination-address"] = node.Ipv4DestinationAddress
    leafs["vni"] = node.Vni
    return leafs
}

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetYangName() string { return "node" }

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetParent() types.Entity { return node.parent }

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetParentYangName() string { return "name" }

// Vservice_MetadataDispositions
// Configure metadata disposition
type Vservice_MetadataDispositions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // metadata disposition name. The type is slice of
    // Vservice_MetadataDispositions_MetadataDisposition.
    MetadataDisposition []Vservice_MetadataDispositions_MetadataDisposition
}

func (metadataDispositions *Vservice_MetadataDispositions) GetFilter() yfilter.YFilter { return metadataDispositions.YFilter }

func (metadataDispositions *Vservice_MetadataDispositions) SetFilter(yf yfilter.YFilter) { metadataDispositions.YFilter = yf }

func (metadataDispositions *Vservice_MetadataDispositions) GetGoName(yname string) string {
    if yname == "metadata-disposition" { return "MetadataDisposition" }
    return ""
}

func (metadataDispositions *Vservice_MetadataDispositions) GetSegmentPath() string {
    return "metadata-dispositions"
}

func (metadataDispositions *Vservice_MetadataDispositions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "metadata-disposition" {
        for _, c := range metadataDispositions.MetadataDisposition {
            if metadataDispositions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vservice_MetadataDispositions_MetadataDisposition{}
        metadataDispositions.MetadataDisposition = append(metadataDispositions.MetadataDisposition, child)
        return &metadataDispositions.MetadataDisposition[len(metadataDispositions.MetadataDisposition)-1]
    }
    return nil
}

func (metadataDispositions *Vservice_MetadataDispositions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range metadataDispositions.MetadataDisposition {
        children[metadataDispositions.MetadataDisposition[i].GetSegmentPath()] = &metadataDispositions.MetadataDisposition[i]
    }
    return children
}

func (metadataDispositions *Vservice_MetadataDispositions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (metadataDispositions *Vservice_MetadataDispositions) GetBundleName() string { return "cisco_ios_xr" }

func (metadataDispositions *Vservice_MetadataDispositions) GetYangName() string { return "metadata-dispositions" }

func (metadataDispositions *Vservice_MetadataDispositions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metadataDispositions *Vservice_MetadataDispositions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metadataDispositions *Vservice_MetadataDispositions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metadataDispositions *Vservice_MetadataDispositions) SetParent(parent types.Entity) { metadataDispositions.parent = parent }

func (metadataDispositions *Vservice_MetadataDispositions) GetParent() types.Entity { return metadataDispositions.parent }

func (metadataDispositions *Vservice_MetadataDispositions) GetParentYangName() string { return "vservice" }

// Vservice_MetadataDispositions_MetadataDisposition
// metadata disposition name
type Vservice_MetadataDispositions_MetadataDisposition struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. disposition name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    DispositionName interface{}

    // This attribute is a key. Specify Format. The type is
    // SfcMetadataType1AllocFormat.
    Format interface{}

    // match entry name. The type is slice of
    // Vservice_MetadataDispositions_MetadataDisposition_MatchEntry.
    MatchEntry []Vservice_MetadataDispositions_MetadataDisposition_MatchEntry
}

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetFilter() yfilter.YFilter { return metadataDisposition.YFilter }

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) SetFilter(yf yfilter.YFilter) { metadataDisposition.YFilter = yf }

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetGoName(yname string) string {
    if yname == "disposition-name" { return "DispositionName" }
    if yname == "format" { return "Format" }
    if yname == "match-entry" { return "MatchEntry" }
    return ""
}

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetSegmentPath() string {
    return "metadata-disposition" + "[disposition-name='" + fmt.Sprintf("%v", metadataDisposition.DispositionName) + "']" + "[format='" + fmt.Sprintf("%v", metadataDisposition.Format) + "']"
}

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "match-entry" {
        for _, c := range metadataDisposition.MatchEntry {
            if metadataDisposition.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vservice_MetadataDispositions_MetadataDisposition_MatchEntry{}
        metadataDisposition.MatchEntry = append(metadataDisposition.MatchEntry, child)
        return &metadataDisposition.MatchEntry[len(metadataDisposition.MatchEntry)-1]
    }
    return nil
}

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range metadataDisposition.MatchEntry {
        children[metadataDisposition.MatchEntry[i].GetSegmentPath()] = &metadataDisposition.MatchEntry[i]
    }
    return children
}

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disposition-name"] = metadataDisposition.DispositionName
    leafs["format"] = metadataDisposition.Format
    return leafs
}

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetBundleName() string { return "cisco_ios_xr" }

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetYangName() string { return "metadata-disposition" }

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) SetParent(parent types.Entity) { metadataDisposition.parent = parent }

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetParent() types.Entity { return metadataDisposition.parent }

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetParentYangName() string { return "metadata-dispositions" }

// Vservice_MetadataDispositions_MetadataDisposition_MatchEntry
// match entry name
type Vservice_MetadataDispositions_MetadataDisposition_MatchEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. match entry name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    MatchEntryName interface{}

    // configure disposition data.
    Node Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node
}

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetFilter() yfilter.YFilter { return matchEntry.YFilter }

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) SetFilter(yf yfilter.YFilter) { matchEntry.YFilter = yf }

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetGoName(yname string) string {
    if yname == "match-entry-name" { return "MatchEntryName" }
    if yname == "node" { return "Node" }
    return ""
}

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetSegmentPath() string {
    return "match-entry" + "[match-entry-name='" + fmt.Sprintf("%v", matchEntry.MatchEntryName) + "']"
}

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        return &matchEntry.Node
    }
    return nil
}

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node"] = &matchEntry.Node
    return children
}

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["match-entry-name"] = matchEntry.MatchEntryName
    return leafs
}

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetBundleName() string { return "cisco_ios_xr" }

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetYangName() string { return "match-entry" }

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) SetParent(parent types.Entity) { matchEntry.parent = parent }

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetParent() types.Entity { return matchEntry.parent }

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetParentYangName() string { return "metadata-disposition" }

// Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node
// configure disposition data
type Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // match type. The type is SfcMetadataDispositionMatch.
    MatchType interface{}

    // action type. The type is SfcMetadataDispositionAction.
    ActionType interface{}

    // VRF name. The type is string.
    Vrf interface{}

    // IPv4 nexthop address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NexthopIpv4Address interface{}

    // 24-bit tenant id. The type is slice of interface{} with range:
    // -2147483648..2147483647.
    TenantId []interface{}
}

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetGoName(yname string) string {
    if yname == "match-type" { return "MatchType" }
    if yname == "action-type" { return "ActionType" }
    if yname == "vrf" { return "Vrf" }
    if yname == "nexthop-ipv4-address" { return "NexthopIpv4Address" }
    if yname == "tenant-id" { return "TenantId" }
    return ""
}

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetSegmentPath() string {
    return "node"
}

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["match-type"] = node.MatchType
    leafs["action-type"] = node.ActionType
    leafs["vrf"] = node.Vrf
    leafs["nexthop-ipv4-address"] = node.NexthopIpv4Address
    leafs["tenant-id"] = node.TenantId
    return leafs
}

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetYangName() string { return "node" }

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetParent() types.Entity { return node.parent }

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetParentYangName() string { return "match-entry" }

// Vservice_ServiceFunctionForwardLocator
// configure service function forward locator
type Vservice_ServiceFunctionForwardLocator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mention the sf/sff name.
    Names Vservice_ServiceFunctionForwardLocator_Names
}

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetFilter() yfilter.YFilter { return serviceFunctionForwardLocator.YFilter }

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) SetFilter(yf yfilter.YFilter) { serviceFunctionForwardLocator.YFilter = yf }

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetGoName(yname string) string {
    if yname == "names" { return "Names" }
    return ""
}

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetSegmentPath() string {
    return "service-function-forward-locator"
}

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "names" {
        return &serviceFunctionForwardLocator.Names
    }
    return nil
}

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["names"] = &serviceFunctionForwardLocator.Names
    return children
}

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetBundleName() string { return "cisco_ios_xr" }

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetYangName() string { return "service-function-forward-locator" }

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) SetParent(parent types.Entity) { serviceFunctionForwardLocator.parent = parent }

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetParent() types.Entity { return serviceFunctionForwardLocator.parent }

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetParentYangName() string { return "vservice" }

// Vservice_ServiceFunctionForwardLocator_Names
// Mention the sf/sff name
type Vservice_ServiceFunctionForwardLocator_Names struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // service function name. The type is slice of
    // Vservice_ServiceFunctionForwardLocator_Names_Name.
    Name []Vservice_ServiceFunctionForwardLocator_Names_Name
}

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetFilter() yfilter.YFilter { return names.YFilter }

func (names *Vservice_ServiceFunctionForwardLocator_Names) SetFilter(yf yfilter.YFilter) { names.YFilter = yf }

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetSegmentPath() string {
    return "names"
}

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "name" {
        for _, c := range names.Name {
            if names.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vservice_ServiceFunctionForwardLocator_Names_Name{}
        names.Name = append(names.Name, child)
        return &names.Name[len(names.Name)-1]
    }
    return nil
}

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range names.Name {
        children[names.Name[i].GetSegmentPath()] = &names.Name[i]
    }
    return children
}

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetBundleName() string { return "cisco_ios_xr" }

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetYangName() string { return "names" }

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (names *Vservice_ServiceFunctionForwardLocator_Names) SetParent(parent types.Entity) { names.parent = parent }

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetParent() types.Entity { return names.parent }

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetParentYangName() string { return "service-function-forward-locator" }

// Vservice_ServiceFunctionForwardLocator_Names_Name
// service function name
type Vservice_ServiceFunctionForwardLocator_Names_Name struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Service function/forwarder name. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    FunctionName interface{}

    // This attribute is a key. Specify locator id. The type is interface{} with
    // range: 1..255.
    LocatorId interface{}

    // configure sff/sffl.
    Node Vservice_ServiceFunctionForwardLocator_Names_Name_Node
}

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetFilter() yfilter.YFilter { return name.YFilter }

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) SetFilter(yf yfilter.YFilter) { name.YFilter = yf }

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetGoName(yname string) string {
    if yname == "function-name" { return "FunctionName" }
    if yname == "locator-id" { return "LocatorId" }
    if yname == "node" { return "Node" }
    return ""
}

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetSegmentPath() string {
    return "name" + "[function-name='" + fmt.Sprintf("%v", name.FunctionName) + "']" + "[locator-id='" + fmt.Sprintf("%v", name.LocatorId) + "']"
}

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        return &name.Node
    }
    return nil
}

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node"] = &name.Node
    return children
}

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["function-name"] = name.FunctionName
    leafs["locator-id"] = name.LocatorId
    return leafs
}

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetBundleName() string { return "cisco_ios_xr" }

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetYangName() string { return "name" }

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) SetParent(parent types.Entity) { name.parent = parent }

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetParent() types.Entity { return name.parent }

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetParentYangName() string { return "names" }

// Vservice_ServiceFunctionForwardLocator_Names_Name_Node
// configure sff/sffl
type Vservice_ServiceFunctionForwardLocator_Names_Name_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transport type. The type is SfcSfTransport.
    Transport interface{}

    // IPv4 source address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4SourceAddress interface{}

    // IPv4 destination address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4DestinationAddress interface{}

    // VNI. The type is interface{} with range: -2147483648..2147483647.
    Vni interface{}
}

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetGoName(yname string) string {
    if yname == "transport" { return "Transport" }
    if yname == "ipv4-source-address" { return "Ipv4SourceAddress" }
    if yname == "ipv4-destination-address" { return "Ipv4DestinationAddress" }
    if yname == "vni" { return "Vni" }
    return ""
}

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetSegmentPath() string {
    return "node"
}

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transport"] = node.Transport
    leafs["ipv4-source-address"] = node.Ipv4SourceAddress
    leafs["ipv4-destination-address"] = node.Ipv4DestinationAddress
    leafs["vni"] = node.Vni
    return leafs
}

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetYangName() string { return "node" }

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetParent() types.Entity { return node.parent }

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetParentYangName() string { return "name" }

// Vservice_MetadataTemplates
// configure metadata imposition
type Vservice_MetadataTemplates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // metadata name, type and format. The type is slice of
    // Vservice_MetadataTemplates_MetadataTemplate.
    MetadataTemplate []Vservice_MetadataTemplates_MetadataTemplate
}

func (metadataTemplates *Vservice_MetadataTemplates) GetFilter() yfilter.YFilter { return metadataTemplates.YFilter }

func (metadataTemplates *Vservice_MetadataTemplates) SetFilter(yf yfilter.YFilter) { metadataTemplates.YFilter = yf }

func (metadataTemplates *Vservice_MetadataTemplates) GetGoName(yname string) string {
    if yname == "metadata-template" { return "MetadataTemplate" }
    return ""
}

func (metadataTemplates *Vservice_MetadataTemplates) GetSegmentPath() string {
    return "metadata-templates"
}

func (metadataTemplates *Vservice_MetadataTemplates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "metadata-template" {
        for _, c := range metadataTemplates.MetadataTemplate {
            if metadataTemplates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vservice_MetadataTemplates_MetadataTemplate{}
        metadataTemplates.MetadataTemplate = append(metadataTemplates.MetadataTemplate, child)
        return &metadataTemplates.MetadataTemplate[len(metadataTemplates.MetadataTemplate)-1]
    }
    return nil
}

func (metadataTemplates *Vservice_MetadataTemplates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range metadataTemplates.MetadataTemplate {
        children[metadataTemplates.MetadataTemplate[i].GetSegmentPath()] = &metadataTemplates.MetadataTemplate[i]
    }
    return children
}

func (metadataTemplates *Vservice_MetadataTemplates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (metadataTemplates *Vservice_MetadataTemplates) GetBundleName() string { return "cisco_ios_xr" }

func (metadataTemplates *Vservice_MetadataTemplates) GetYangName() string { return "metadata-templates" }

func (metadataTemplates *Vservice_MetadataTemplates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metadataTemplates *Vservice_MetadataTemplates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metadataTemplates *Vservice_MetadataTemplates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metadataTemplates *Vservice_MetadataTemplates) SetParent(parent types.Entity) { metadataTemplates.parent = parent }

func (metadataTemplates *Vservice_MetadataTemplates) GetParent() types.Entity { return metadataTemplates.parent }

func (metadataTemplates *Vservice_MetadataTemplates) GetParentYangName() string { return "vservice" }

// Vservice_MetadataTemplates_MetadataTemplate
// metadata name, type and format
type Vservice_MetadataTemplates_MetadataTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. metadata name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    MetadataName interface{}

    // This attribute is a key. Specify Type . The type is SfcMetadataAlloc.
    Type interface{}

    // This attribute is a key. Specify Format. The type is
    // SfcMetadataType1AllocFormat.
    Format interface{}

    // Enter 24-bit tenant id. The type is interface{} with range: 1..16777215.
    TenantId interface{}
}

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetFilter() yfilter.YFilter { return metadataTemplate.YFilter }

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) SetFilter(yf yfilter.YFilter) { metadataTemplate.YFilter = yf }

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetGoName(yname string) string {
    if yname == "metadata-name" { return "MetadataName" }
    if yname == "type" { return "Type" }
    if yname == "format" { return "Format" }
    if yname == "tenant-id" { return "TenantId" }
    return ""
}

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetSegmentPath() string {
    return "metadata-template" + "[metadata-name='" + fmt.Sprintf("%v", metadataTemplate.MetadataName) + "']" + "[type='" + fmt.Sprintf("%v", metadataTemplate.Type) + "']" + "[format='" + fmt.Sprintf("%v", metadataTemplate.Format) + "']"
}

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metadata-name"] = metadataTemplate.MetadataName
    leafs["type"] = metadataTemplate.Type
    leafs["format"] = metadataTemplate.Format
    leafs["tenant-id"] = metadataTemplate.TenantId
    return leafs
}

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetYangName() string { return "metadata-template" }

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) SetParent(parent types.Entity) { metadataTemplate.parent = parent }

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetParent() types.Entity { return metadataTemplate.parent }

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetParentYangName() string { return "metadata-templates" }

// Vservice_ServiceFunctionPath
// service function path
type Vservice_ServiceFunctionPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // service function path id.
    Paths Vservice_ServiceFunctionPath_Paths
}

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetFilter() yfilter.YFilter { return serviceFunctionPath.YFilter }

func (serviceFunctionPath *Vservice_ServiceFunctionPath) SetFilter(yf yfilter.YFilter) { serviceFunctionPath.YFilter = yf }

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetGoName(yname string) string {
    if yname == "paths" { return "Paths" }
    return ""
}

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetSegmentPath() string {
    return "service-function-path"
}

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "paths" {
        return &serviceFunctionPath.Paths
    }
    return nil
}

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["paths"] = &serviceFunctionPath.Paths
    return children
}

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetBundleName() string { return "cisco_ios_xr" }

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetYangName() string { return "service-function-path" }

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceFunctionPath *Vservice_ServiceFunctionPath) SetParent(parent types.Entity) { serviceFunctionPath.parent = parent }

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetParent() types.Entity { return serviceFunctionPath.parent }

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetParentYangName() string { return "vservice" }

// Vservice_ServiceFunctionPath_Paths
// service function path id
type Vservice_ServiceFunctionPath_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // specify the service function path id. The type is slice of
    // Vservice_ServiceFunctionPath_Paths_Path.
    Path []Vservice_ServiceFunctionPath_Paths_Path
}

func (paths *Vservice_ServiceFunctionPath_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *Vservice_ServiceFunctionPath_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *Vservice_ServiceFunctionPath_Paths) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (paths *Vservice_ServiceFunctionPath_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *Vservice_ServiceFunctionPath_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        for _, c := range paths.Path {
            if paths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vservice_ServiceFunctionPath_Paths_Path{}
        paths.Path = append(paths.Path, child)
        return &paths.Path[len(paths.Path)-1]
    }
    return nil
}

func (paths *Vservice_ServiceFunctionPath_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range paths.Path {
        children[paths.Path[i].GetSegmentPath()] = &paths.Path[i]
    }
    return children
}

func (paths *Vservice_ServiceFunctionPath_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (paths *Vservice_ServiceFunctionPath_Paths) GetBundleName() string { return "cisco_ios_xr" }

func (paths *Vservice_ServiceFunctionPath_Paths) GetYangName() string { return "paths" }

func (paths *Vservice_ServiceFunctionPath_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paths *Vservice_ServiceFunctionPath_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paths *Vservice_ServiceFunctionPath_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paths *Vservice_ServiceFunctionPath_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *Vservice_ServiceFunctionPath_Paths) GetParent() types.Entity { return paths.parent }

func (paths *Vservice_ServiceFunctionPath_Paths) GetParentYangName() string { return "service-function-path" }

// Vservice_ServiceFunctionPath_Paths_Path
// specify the service function path id
type Vservice_ServiceFunctionPath_Paths_Path struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specify the service function path id. The type is
    // interface{} with range: 1..16777215.
    PathId interface{}

    // specify the service index. The type is slice of
    // Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex.
    ServiceIndex []Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex
}

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *Vservice_ServiceFunctionPath_Paths_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetGoName(yname string) string {
    if yname == "path-id" { return "PathId" }
    if yname == "service-index" { return "ServiceIndex" }
    return ""
}

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetSegmentPath() string {
    return "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
}

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-index" {
        for _, c := range path.ServiceIndex {
            if path.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex{}
        path.ServiceIndex = append(path.ServiceIndex, child)
        return &path.ServiceIndex[len(path.ServiceIndex)-1]
    }
    return nil
}

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range path.ServiceIndex {
        children[path.ServiceIndex[i].GetSegmentPath()] = &path.ServiceIndex[i]
    }
    return children
}

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-id"] = path.PathId
    return leafs
}

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetBundleName() string { return "cisco_ios_xr" }

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetYangName() string { return "path" }

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (path *Vservice_ServiceFunctionPath_Paths_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetParent() types.Entity { return path.parent }

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetParentYangName() string { return "paths" }

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex
// specify the service index
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specify the id of service function. The type is
    // interface{} with range: 1..255.
    Index interface{}

    // configure terminate.
    Terminate Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate

    // service function forwarder .
    SffNames Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames

    // service function .
    SfNames Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames
}

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetFilter() yfilter.YFilter { return serviceIndex.YFilter }

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) SetFilter(yf yfilter.YFilter) { serviceIndex.YFilter = yf }

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "terminate" { return "Terminate" }
    if yname == "sff-names" { return "SffNames" }
    if yname == "sf-names" { return "SfNames" }
    return ""
}

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetSegmentPath() string {
    return "service-index" + "[index='" + fmt.Sprintf("%v", serviceIndex.Index) + "']"
}

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "terminate" {
        return &serviceIndex.Terminate
    }
    if childYangName == "sff-names" {
        return &serviceIndex.SffNames
    }
    if childYangName == "sf-names" {
        return &serviceIndex.SfNames
    }
    return nil
}

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["terminate"] = &serviceIndex.Terminate
    children["sff-names"] = &serviceIndex.SffNames
    children["sf-names"] = &serviceIndex.SfNames
    return children
}

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = serviceIndex.Index
    return leafs
}

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetBundleName() string { return "cisco_ios_xr" }

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetYangName() string { return "service-index" }

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) SetParent(parent types.Entity) { serviceIndex.parent = parent }

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetParent() types.Entity { return serviceIndex.parent }

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetParentYangName() string { return "path" }

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate
// configure terminate
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // configure default terminate action.
    Node Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node
}

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetFilter() yfilter.YFilter { return terminate.YFilter }

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) SetFilter(yf yfilter.YFilter) { terminate.YFilter = yf }

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetSegmentPath() string {
    return "terminate"
}

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        return &terminate.Node
    }
    return nil
}

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node"] = &terminate.Node
    return children
}

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetBundleName() string { return "cisco_ios_xr" }

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetYangName() string { return "terminate" }

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) SetParent(parent types.Entity) { terminate.parent = parent }

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetParent() types.Entity { return terminate.parent }

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetParentYangName() string { return "service-index" }

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node
// configure default terminate action
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // default action enum. The type is SfcMetadataDispositionAction.
    Action interface{}

    // nexthop vrf name. The type is string.
    Vrf interface{}

    // IPv4 nexthop address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NexthopIpv4Address interface{}

    // metadata-disposition name. The type is string.
    MetatdataDisposition interface{}
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "vrf" { return "Vrf" }
    if yname == "nexthop-ipv4-address" { return "NexthopIpv4Address" }
    if yname == "metatdata-disposition" { return "MetatdataDisposition" }
    return ""
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetSegmentPath() string {
    return "node"
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = node.Action
    leafs["vrf"] = node.Vrf
    leafs["nexthop-ipv4-address"] = node.NexthopIpv4Address
    leafs["metatdata-disposition"] = node.MetatdataDisposition
    return leafs
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetYangName() string { return "node" }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetParent() types.Entity { return node.parent }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetParentYangName() string { return "terminate" }

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames
// service function forwarder 
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // service function forwarder name. The type is slice of
    // Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName.
    SffName []Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName
}

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetFilter() yfilter.YFilter { return sffNames.YFilter }

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) SetFilter(yf yfilter.YFilter) { sffNames.YFilter = yf }

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetGoName(yname string) string {
    if yname == "sff-name" { return "SffName" }
    return ""
}

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetSegmentPath() string {
    return "sff-names"
}

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sff-name" {
        for _, c := range sffNames.SffName {
            if sffNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName{}
        sffNames.SffName = append(sffNames.SffName, child)
        return &sffNames.SffName[len(sffNames.SffName)-1]
    }
    return nil
}

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sffNames.SffName {
        children[sffNames.SffName[i].GetSegmentPath()] = &sffNames.SffName[i]
    }
    return children
}

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetBundleName() string { return "cisco_ios_xr" }

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetYangName() string { return "sff-names" }

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) SetParent(parent types.Entity) { sffNames.parent = parent }

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetParent() types.Entity { return sffNames.parent }

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetParentYangName() string { return "service-index" }

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName
// service function forwarder name
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SFF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // configure SFP.
    Node Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node
}

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetFilter() yfilter.YFilter { return sffName.YFilter }

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) SetFilter(yf yfilter.YFilter) { sffName.YFilter = yf }

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "node" { return "Node" }
    return ""
}

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetSegmentPath() string {
    return "sff-name" + "[name='" + fmt.Sprintf("%v", sffName.Name) + "']"
}

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        return &sffName.Node
    }
    return nil
}

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node"] = &sffName.Node
    return children
}

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sffName.Name
    return leafs
}

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetBundleName() string { return "cisco_ios_xr" }

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetYangName() string { return "sff-name" }

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) SetParent(parent types.Entity) { sffName.parent = parent }

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetParent() types.Entity { return sffName.parent }

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetParentYangName() string { return "sff-names" }

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node
// configure SFP
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Service function path. The type is interface{}.
    Enable interface{}

    // Dummy. The type is interface{}.
    Reserved interface{}
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "reserved" { return "Reserved" }
    return ""
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetSegmentPath() string {
    return "node"
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = node.Enable
    leafs["reserved"] = node.Reserved
    return leafs
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetYangName() string { return "node" }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetParent() types.Entity { return node.parent }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetParentYangName() string { return "sff-name" }

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames
// service function 
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // service function name. The type is slice of
    // Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName.
    SfName []Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName
}

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetFilter() yfilter.YFilter { return sfNames.YFilter }

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) SetFilter(yf yfilter.YFilter) { sfNames.YFilter = yf }

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetGoName(yname string) string {
    if yname == "sf-name" { return "SfName" }
    return ""
}

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetSegmentPath() string {
    return "sf-names"
}

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sf-name" {
        for _, c := range sfNames.SfName {
            if sfNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName{}
        sfNames.SfName = append(sfNames.SfName, child)
        return &sfNames.SfName[len(sfNames.SfName)-1]
    }
    return nil
}

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sfNames.SfName {
        children[sfNames.SfName[i].GetSegmentPath()] = &sfNames.SfName[i]
    }
    return children
}

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetBundleName() string { return "cisco_ios_xr" }

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetYangName() string { return "sf-names" }

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) SetParent(parent types.Entity) { sfNames.parent = parent }

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetParent() types.Entity { return sfNames.parent }

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetParentYangName() string { return "service-index" }

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName
// service function name
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // configure SFP.
    Node Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node
}

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetFilter() yfilter.YFilter { return sfName.YFilter }

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) SetFilter(yf yfilter.YFilter) { sfName.YFilter = yf }

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "node" { return "Node" }
    return ""
}

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetSegmentPath() string {
    return "sf-name" + "[name='" + fmt.Sprintf("%v", sfName.Name) + "']"
}

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        return &sfName.Node
    }
    return nil
}

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node"] = &sfName.Node
    return children
}

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sfName.Name
    return leafs
}

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetBundleName() string { return "cisco_ios_xr" }

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetYangName() string { return "sf-name" }

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) SetParent(parent types.Entity) { sfName.parent = parent }

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetParent() types.Entity { return sfName.parent }

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetParentYangName() string { return "sf-names" }

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node
// configure SFP
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Service function path. The type is interface{}.
    Enable interface{}

    // Dummy. The type is interface{}.
    Reserved interface{}
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "reserved" { return "Reserved" }
    return ""
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetSegmentPath() string {
    return "node"
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = node.Enable
    leafs["reserved"] = node.Reserved
    return leafs
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetYangName() string { return "node" }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetParent() types.Entity { return node.parent }

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetParentYangName() string { return "sf-name" }

