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
    EntityData types.CommonEntityData
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

func (vservice *Vservice) GetEntityData() *types.CommonEntityData {
    vservice.EntityData.YFilter = vservice.YFilter
    vservice.EntityData.YangName = "vservice"
    vservice.EntityData.BundleName = "cisco_ios_xr"
    vservice.EntityData.ParentYangName = "Cisco-IOS-XR-vservice-cfg"
    vservice.EntityData.SegmentPath = "Cisco-IOS-XR-vservice-cfg:vservice"
    vservice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vservice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vservice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vservice.EntityData.Children = types.NewOrderedMap()
    vservice.EntityData.Children.Append("service-function-locator", types.YChild{"ServiceFunctionLocator", &vservice.ServiceFunctionLocator})
    vservice.EntityData.Children.Append("metadata-dispositions", types.YChild{"MetadataDispositions", &vservice.MetadataDispositions})
    vservice.EntityData.Children.Append("service-function-forward-locator", types.YChild{"ServiceFunctionForwardLocator", &vservice.ServiceFunctionForwardLocator})
    vservice.EntityData.Children.Append("metadata-templates", types.YChild{"MetadataTemplates", &vservice.MetadataTemplates})
    vservice.EntityData.Children.Append("service-function-path", types.YChild{"ServiceFunctionPath", &vservice.ServiceFunctionPath})
    vservice.EntityData.Leafs = types.NewOrderedMap()

    vservice.EntityData.YListKeys = []string {}

    return &(vservice.EntityData)
}

// Vservice_ServiceFunctionLocator
// configure service function locator
type Vservice_ServiceFunctionLocator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mention the sf/sff name.
    Names Vservice_ServiceFunctionLocator_Names
}

func (serviceFunctionLocator *Vservice_ServiceFunctionLocator) GetEntityData() *types.CommonEntityData {
    serviceFunctionLocator.EntityData.YFilter = serviceFunctionLocator.YFilter
    serviceFunctionLocator.EntityData.YangName = "service-function-locator"
    serviceFunctionLocator.EntityData.BundleName = "cisco_ios_xr"
    serviceFunctionLocator.EntityData.ParentYangName = "vservice"
    serviceFunctionLocator.EntityData.SegmentPath = "service-function-locator"
    serviceFunctionLocator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceFunctionLocator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceFunctionLocator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceFunctionLocator.EntityData.Children = types.NewOrderedMap()
    serviceFunctionLocator.EntityData.Children.Append("names", types.YChild{"Names", &serviceFunctionLocator.Names})
    serviceFunctionLocator.EntityData.Leafs = types.NewOrderedMap()

    serviceFunctionLocator.EntityData.YListKeys = []string {}

    return &(serviceFunctionLocator.EntityData)
}

// Vservice_ServiceFunctionLocator_Names
// Mention the sf/sff name
type Vservice_ServiceFunctionLocator_Names struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // service function name. The type is slice of
    // Vservice_ServiceFunctionLocator_Names_Name.
    Name []*Vservice_ServiceFunctionLocator_Names_Name
}

func (names *Vservice_ServiceFunctionLocator_Names) GetEntityData() *types.CommonEntityData {
    names.EntityData.YFilter = names.YFilter
    names.EntityData.YangName = "names"
    names.EntityData.BundleName = "cisco_ios_xr"
    names.EntityData.ParentYangName = "service-function-locator"
    names.EntityData.SegmentPath = "names"
    names.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    names.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    names.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    names.EntityData.Children = types.NewOrderedMap()
    names.EntityData.Children.Append("name", types.YChild{"Name", nil})
    for i := range names.Name {
        names.EntityData.Children.Append(types.GetSegmentPath(names.Name[i]), types.YChild{"Name", names.Name[i]})
    }
    names.EntityData.Leafs = types.NewOrderedMap()

    names.EntityData.YListKeys = []string {}

    return &(names.EntityData)
}

// Vservice_ServiceFunctionLocator_Names_Name
// service function name
type Vservice_ServiceFunctionLocator_Names_Name struct {
    EntityData types.CommonEntityData
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

func (name *Vservice_ServiceFunctionLocator_Names_Name) GetEntityData() *types.CommonEntityData {
    name.EntityData.YFilter = name.YFilter
    name.EntityData.YangName = "name"
    name.EntityData.BundleName = "cisco_ios_xr"
    name.EntityData.ParentYangName = "names"
    name.EntityData.SegmentPath = "name" + types.AddKeyToken(name.FunctionName, "function-name") + types.AddKeyToken(name.LocatorId, "locator-id")
    name.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    name.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    name.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    name.EntityData.Children = types.NewOrderedMap()
    name.EntityData.Children.Append("node", types.YChild{"Node", &name.Node})
    name.EntityData.Leafs = types.NewOrderedMap()
    name.EntityData.Leafs.Append("function-name", types.YLeaf{"FunctionName", name.FunctionName})
    name.EntityData.Leafs.Append("locator-id", types.YLeaf{"LocatorId", name.LocatorId})

    name.EntityData.YListKeys = []string {"FunctionName", "LocatorId"}

    return &(name.EntityData)
}

// Vservice_ServiceFunctionLocator_Names_Name_Node
// configure sff/sffl
type Vservice_ServiceFunctionLocator_Names_Name_Node struct {
    EntityData types.CommonEntityData
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

func (node *Vservice_ServiceFunctionLocator_Names_Name_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "name"
    node.EntityData.SegmentPath = "node"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("transport", types.YLeaf{"Transport", node.Transport})
    node.EntityData.Leafs.Append("ipv4-source-address", types.YLeaf{"Ipv4SourceAddress", node.Ipv4SourceAddress})
    node.EntityData.Leafs.Append("ipv4-destination-address", types.YLeaf{"Ipv4DestinationAddress", node.Ipv4DestinationAddress})
    node.EntityData.Leafs.Append("vni", types.YLeaf{"Vni", node.Vni})

    node.EntityData.YListKeys = []string {}

    return &(node.EntityData)
}

// Vservice_MetadataDispositions
// Configure metadata disposition
type Vservice_MetadataDispositions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // metadata disposition name. The type is slice of
    // Vservice_MetadataDispositions_MetadataDisposition.
    MetadataDisposition []*Vservice_MetadataDispositions_MetadataDisposition
}

func (metadataDispositions *Vservice_MetadataDispositions) GetEntityData() *types.CommonEntityData {
    metadataDispositions.EntityData.YFilter = metadataDispositions.YFilter
    metadataDispositions.EntityData.YangName = "metadata-dispositions"
    metadataDispositions.EntityData.BundleName = "cisco_ios_xr"
    metadataDispositions.EntityData.ParentYangName = "vservice"
    metadataDispositions.EntityData.SegmentPath = "metadata-dispositions"
    metadataDispositions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metadataDispositions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metadataDispositions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metadataDispositions.EntityData.Children = types.NewOrderedMap()
    metadataDispositions.EntityData.Children.Append("metadata-disposition", types.YChild{"MetadataDisposition", nil})
    for i := range metadataDispositions.MetadataDisposition {
        metadataDispositions.EntityData.Children.Append(types.GetSegmentPath(metadataDispositions.MetadataDisposition[i]), types.YChild{"MetadataDisposition", metadataDispositions.MetadataDisposition[i]})
    }
    metadataDispositions.EntityData.Leafs = types.NewOrderedMap()

    metadataDispositions.EntityData.YListKeys = []string {}

    return &(metadataDispositions.EntityData)
}

// Vservice_MetadataDispositions_MetadataDisposition
// metadata disposition name
type Vservice_MetadataDispositions_MetadataDisposition struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. disposition name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    DispositionName interface{}

    // This attribute is a key. Specify Format. The type is
    // SfcMetadataType1AllocFormat.
    Format interface{}

    // match entry name. The type is slice of
    // Vservice_MetadataDispositions_MetadataDisposition_MatchEntry.
    MatchEntry []*Vservice_MetadataDispositions_MetadataDisposition_MatchEntry
}

func (metadataDisposition *Vservice_MetadataDispositions_MetadataDisposition) GetEntityData() *types.CommonEntityData {
    metadataDisposition.EntityData.YFilter = metadataDisposition.YFilter
    metadataDisposition.EntityData.YangName = "metadata-disposition"
    metadataDisposition.EntityData.BundleName = "cisco_ios_xr"
    metadataDisposition.EntityData.ParentYangName = "metadata-dispositions"
    metadataDisposition.EntityData.SegmentPath = "metadata-disposition" + types.AddKeyToken(metadataDisposition.DispositionName, "disposition-name") + types.AddKeyToken(metadataDisposition.Format, "format")
    metadataDisposition.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metadataDisposition.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metadataDisposition.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metadataDisposition.EntityData.Children = types.NewOrderedMap()
    metadataDisposition.EntityData.Children.Append("match-entry", types.YChild{"MatchEntry", nil})
    for i := range metadataDisposition.MatchEntry {
        metadataDisposition.EntityData.Children.Append(types.GetSegmentPath(metadataDisposition.MatchEntry[i]), types.YChild{"MatchEntry", metadataDisposition.MatchEntry[i]})
    }
    metadataDisposition.EntityData.Leafs = types.NewOrderedMap()
    metadataDisposition.EntityData.Leafs.Append("disposition-name", types.YLeaf{"DispositionName", metadataDisposition.DispositionName})
    metadataDisposition.EntityData.Leafs.Append("format", types.YLeaf{"Format", metadataDisposition.Format})

    metadataDisposition.EntityData.YListKeys = []string {"DispositionName", "Format"}

    return &(metadataDisposition.EntityData)
}

// Vservice_MetadataDispositions_MetadataDisposition_MatchEntry
// match entry name
type Vservice_MetadataDispositions_MetadataDisposition_MatchEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. match entry name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    MatchEntryName interface{}

    // configure disposition data.
    Node Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node
}

func (matchEntry *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry) GetEntityData() *types.CommonEntityData {
    matchEntry.EntityData.YFilter = matchEntry.YFilter
    matchEntry.EntityData.YangName = "match-entry"
    matchEntry.EntityData.BundleName = "cisco_ios_xr"
    matchEntry.EntityData.ParentYangName = "metadata-disposition"
    matchEntry.EntityData.SegmentPath = "match-entry" + types.AddKeyToken(matchEntry.MatchEntryName, "match-entry-name")
    matchEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    matchEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    matchEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    matchEntry.EntityData.Children = types.NewOrderedMap()
    matchEntry.EntityData.Children.Append("node", types.YChild{"Node", &matchEntry.Node})
    matchEntry.EntityData.Leafs = types.NewOrderedMap()
    matchEntry.EntityData.Leafs.Append("match-entry-name", types.YLeaf{"MatchEntryName", matchEntry.MatchEntryName})

    matchEntry.EntityData.YListKeys = []string {"MatchEntryName"}

    return &(matchEntry.EntityData)
}

// Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node
// configure disposition data
type Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node struct {
    EntityData types.CommonEntityData
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

func (node *Vservice_MetadataDispositions_MetadataDisposition_MatchEntry_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "match-entry"
    node.EntityData.SegmentPath = "node"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("match-type", types.YLeaf{"MatchType", node.MatchType})
    node.EntityData.Leafs.Append("action-type", types.YLeaf{"ActionType", node.ActionType})
    node.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", node.Vrf})
    node.EntityData.Leafs.Append("nexthop-ipv4-address", types.YLeaf{"NexthopIpv4Address", node.NexthopIpv4Address})
    node.EntityData.Leafs.Append("tenant-id", types.YLeaf{"TenantId", node.TenantId})

    node.EntityData.YListKeys = []string {}

    return &(node.EntityData)
}

// Vservice_ServiceFunctionForwardLocator
// configure service function forward locator
type Vservice_ServiceFunctionForwardLocator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mention the sf/sff name.
    Names Vservice_ServiceFunctionForwardLocator_Names
}

func (serviceFunctionForwardLocator *Vservice_ServiceFunctionForwardLocator) GetEntityData() *types.CommonEntityData {
    serviceFunctionForwardLocator.EntityData.YFilter = serviceFunctionForwardLocator.YFilter
    serviceFunctionForwardLocator.EntityData.YangName = "service-function-forward-locator"
    serviceFunctionForwardLocator.EntityData.BundleName = "cisco_ios_xr"
    serviceFunctionForwardLocator.EntityData.ParentYangName = "vservice"
    serviceFunctionForwardLocator.EntityData.SegmentPath = "service-function-forward-locator"
    serviceFunctionForwardLocator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceFunctionForwardLocator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceFunctionForwardLocator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceFunctionForwardLocator.EntityData.Children = types.NewOrderedMap()
    serviceFunctionForwardLocator.EntityData.Children.Append("names", types.YChild{"Names", &serviceFunctionForwardLocator.Names})
    serviceFunctionForwardLocator.EntityData.Leafs = types.NewOrderedMap()

    serviceFunctionForwardLocator.EntityData.YListKeys = []string {}

    return &(serviceFunctionForwardLocator.EntityData)
}

// Vservice_ServiceFunctionForwardLocator_Names
// Mention the sf/sff name
type Vservice_ServiceFunctionForwardLocator_Names struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // service function name. The type is slice of
    // Vservice_ServiceFunctionForwardLocator_Names_Name.
    Name []*Vservice_ServiceFunctionForwardLocator_Names_Name
}

func (names *Vservice_ServiceFunctionForwardLocator_Names) GetEntityData() *types.CommonEntityData {
    names.EntityData.YFilter = names.YFilter
    names.EntityData.YangName = "names"
    names.EntityData.BundleName = "cisco_ios_xr"
    names.EntityData.ParentYangName = "service-function-forward-locator"
    names.EntityData.SegmentPath = "names"
    names.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    names.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    names.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    names.EntityData.Children = types.NewOrderedMap()
    names.EntityData.Children.Append("name", types.YChild{"Name", nil})
    for i := range names.Name {
        names.EntityData.Children.Append(types.GetSegmentPath(names.Name[i]), types.YChild{"Name", names.Name[i]})
    }
    names.EntityData.Leafs = types.NewOrderedMap()

    names.EntityData.YListKeys = []string {}

    return &(names.EntityData)
}

// Vservice_ServiceFunctionForwardLocator_Names_Name
// service function name
type Vservice_ServiceFunctionForwardLocator_Names_Name struct {
    EntityData types.CommonEntityData
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

func (name *Vservice_ServiceFunctionForwardLocator_Names_Name) GetEntityData() *types.CommonEntityData {
    name.EntityData.YFilter = name.YFilter
    name.EntityData.YangName = "name"
    name.EntityData.BundleName = "cisco_ios_xr"
    name.EntityData.ParentYangName = "names"
    name.EntityData.SegmentPath = "name" + types.AddKeyToken(name.FunctionName, "function-name") + types.AddKeyToken(name.LocatorId, "locator-id")
    name.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    name.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    name.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    name.EntityData.Children = types.NewOrderedMap()
    name.EntityData.Children.Append("node", types.YChild{"Node", &name.Node})
    name.EntityData.Leafs = types.NewOrderedMap()
    name.EntityData.Leafs.Append("function-name", types.YLeaf{"FunctionName", name.FunctionName})
    name.EntityData.Leafs.Append("locator-id", types.YLeaf{"LocatorId", name.LocatorId})

    name.EntityData.YListKeys = []string {"FunctionName", "LocatorId"}

    return &(name.EntityData)
}

// Vservice_ServiceFunctionForwardLocator_Names_Name_Node
// configure sff/sffl
type Vservice_ServiceFunctionForwardLocator_Names_Name_Node struct {
    EntityData types.CommonEntityData
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

func (node *Vservice_ServiceFunctionForwardLocator_Names_Name_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "name"
    node.EntityData.SegmentPath = "node"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("transport", types.YLeaf{"Transport", node.Transport})
    node.EntityData.Leafs.Append("ipv4-source-address", types.YLeaf{"Ipv4SourceAddress", node.Ipv4SourceAddress})
    node.EntityData.Leafs.Append("ipv4-destination-address", types.YLeaf{"Ipv4DestinationAddress", node.Ipv4DestinationAddress})
    node.EntityData.Leafs.Append("vni", types.YLeaf{"Vni", node.Vni})

    node.EntityData.YListKeys = []string {}

    return &(node.EntityData)
}

// Vservice_MetadataTemplates
// configure metadata imposition
type Vservice_MetadataTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // metadata name, type and format. The type is slice of
    // Vservice_MetadataTemplates_MetadataTemplate.
    MetadataTemplate []*Vservice_MetadataTemplates_MetadataTemplate
}

func (metadataTemplates *Vservice_MetadataTemplates) GetEntityData() *types.CommonEntityData {
    metadataTemplates.EntityData.YFilter = metadataTemplates.YFilter
    metadataTemplates.EntityData.YangName = "metadata-templates"
    metadataTemplates.EntityData.BundleName = "cisco_ios_xr"
    metadataTemplates.EntityData.ParentYangName = "vservice"
    metadataTemplates.EntityData.SegmentPath = "metadata-templates"
    metadataTemplates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metadataTemplates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metadataTemplates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metadataTemplates.EntityData.Children = types.NewOrderedMap()
    metadataTemplates.EntityData.Children.Append("metadata-template", types.YChild{"MetadataTemplate", nil})
    for i := range metadataTemplates.MetadataTemplate {
        metadataTemplates.EntityData.Children.Append(types.GetSegmentPath(metadataTemplates.MetadataTemplate[i]), types.YChild{"MetadataTemplate", metadataTemplates.MetadataTemplate[i]})
    }
    metadataTemplates.EntityData.Leafs = types.NewOrderedMap()

    metadataTemplates.EntityData.YListKeys = []string {}

    return &(metadataTemplates.EntityData)
}

// Vservice_MetadataTemplates_MetadataTemplate
// metadata name, type and format
type Vservice_MetadataTemplates_MetadataTemplate struct {
    EntityData types.CommonEntityData
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

func (metadataTemplate *Vservice_MetadataTemplates_MetadataTemplate) GetEntityData() *types.CommonEntityData {
    metadataTemplate.EntityData.YFilter = metadataTemplate.YFilter
    metadataTemplate.EntityData.YangName = "metadata-template"
    metadataTemplate.EntityData.BundleName = "cisco_ios_xr"
    metadataTemplate.EntityData.ParentYangName = "metadata-templates"
    metadataTemplate.EntityData.SegmentPath = "metadata-template" + types.AddKeyToken(metadataTemplate.MetadataName, "metadata-name") + types.AddKeyToken(metadataTemplate.Type, "type") + types.AddKeyToken(metadataTemplate.Format, "format")
    metadataTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metadataTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metadataTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metadataTemplate.EntityData.Children = types.NewOrderedMap()
    metadataTemplate.EntityData.Leafs = types.NewOrderedMap()
    metadataTemplate.EntityData.Leafs.Append("metadata-name", types.YLeaf{"MetadataName", metadataTemplate.MetadataName})
    metadataTemplate.EntityData.Leafs.Append("type", types.YLeaf{"Type", metadataTemplate.Type})
    metadataTemplate.EntityData.Leafs.Append("format", types.YLeaf{"Format", metadataTemplate.Format})
    metadataTemplate.EntityData.Leafs.Append("tenant-id", types.YLeaf{"TenantId", metadataTemplate.TenantId})

    metadataTemplate.EntityData.YListKeys = []string {"MetadataName", "Type", "Format"}

    return &(metadataTemplate.EntityData)
}

// Vservice_ServiceFunctionPath
// service function path
type Vservice_ServiceFunctionPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // service function path id.
    Paths Vservice_ServiceFunctionPath_Paths
}

func (serviceFunctionPath *Vservice_ServiceFunctionPath) GetEntityData() *types.CommonEntityData {
    serviceFunctionPath.EntityData.YFilter = serviceFunctionPath.YFilter
    serviceFunctionPath.EntityData.YangName = "service-function-path"
    serviceFunctionPath.EntityData.BundleName = "cisco_ios_xr"
    serviceFunctionPath.EntityData.ParentYangName = "vservice"
    serviceFunctionPath.EntityData.SegmentPath = "service-function-path"
    serviceFunctionPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceFunctionPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceFunctionPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceFunctionPath.EntityData.Children = types.NewOrderedMap()
    serviceFunctionPath.EntityData.Children.Append("paths", types.YChild{"Paths", &serviceFunctionPath.Paths})
    serviceFunctionPath.EntityData.Leafs = types.NewOrderedMap()

    serviceFunctionPath.EntityData.YListKeys = []string {}

    return &(serviceFunctionPath.EntityData)
}

// Vservice_ServiceFunctionPath_Paths
// service function path id
type Vservice_ServiceFunctionPath_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // specify the service function path id. The type is slice of
    // Vservice_ServiceFunctionPath_Paths_Path.
    Path []*Vservice_ServiceFunctionPath_Paths_Path
}

func (paths *Vservice_ServiceFunctionPath_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "service-function-path"
    paths.EntityData.SegmentPath = "paths"
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

// Vservice_ServiceFunctionPath_Paths_Path
// specify the service function path id
type Vservice_ServiceFunctionPath_Paths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specify the service function path id. The type is
    // interface{} with range: 1..16777215.
    PathId interface{}

    // specify the service index. The type is slice of
    // Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex.
    ServiceIndex []*Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex
}

func (path *Vservice_ServiceFunctionPath_Paths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "paths"
    path.EntityData.SegmentPath = "path" + types.AddKeyToken(path.PathId, "path-id")
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Children.Append("service-index", types.YChild{"ServiceIndex", nil})
    for i := range path.ServiceIndex {
        path.EntityData.Children.Append(types.GetSegmentPath(path.ServiceIndex[i]), types.YChild{"ServiceIndex", path.ServiceIndex[i]})
    }
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", path.PathId})

    path.EntityData.YListKeys = []string {"PathId"}

    return &(path.EntityData)
}

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex
// specify the service index
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex struct {
    EntityData types.CommonEntityData
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

func (serviceIndex *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex) GetEntityData() *types.CommonEntityData {
    serviceIndex.EntityData.YFilter = serviceIndex.YFilter
    serviceIndex.EntityData.YangName = "service-index"
    serviceIndex.EntityData.BundleName = "cisco_ios_xr"
    serviceIndex.EntityData.ParentYangName = "path"
    serviceIndex.EntityData.SegmentPath = "service-index" + types.AddKeyToken(serviceIndex.Index, "index")
    serviceIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceIndex.EntityData.Children = types.NewOrderedMap()
    serviceIndex.EntityData.Children.Append("terminate", types.YChild{"Terminate", &serviceIndex.Terminate})
    serviceIndex.EntityData.Children.Append("sff-names", types.YChild{"SffNames", &serviceIndex.SffNames})
    serviceIndex.EntityData.Children.Append("sf-names", types.YChild{"SfNames", &serviceIndex.SfNames})
    serviceIndex.EntityData.Leafs = types.NewOrderedMap()
    serviceIndex.EntityData.Leafs.Append("index", types.YLeaf{"Index", serviceIndex.Index})

    serviceIndex.EntityData.YListKeys = []string {"Index"}

    return &(serviceIndex.EntityData)
}

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate
// configure terminate
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // configure default terminate action.
    Node Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node
}

func (terminate *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate) GetEntityData() *types.CommonEntityData {
    terminate.EntityData.YFilter = terminate.YFilter
    terminate.EntityData.YangName = "terminate"
    terminate.EntityData.BundleName = "cisco_ios_xr"
    terminate.EntityData.ParentYangName = "service-index"
    terminate.EntityData.SegmentPath = "terminate"
    terminate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    terminate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    terminate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    terminate.EntityData.Children = types.NewOrderedMap()
    terminate.EntityData.Children.Append("node", types.YChild{"Node", &terminate.Node})
    terminate.EntityData.Leafs = types.NewOrderedMap()

    terminate.EntityData.YListKeys = []string {}

    return &(terminate.EntityData)
}

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node
// configure default terminate action
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node struct {
    EntityData types.CommonEntityData
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

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_Terminate_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "terminate"
    node.EntityData.SegmentPath = "node"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("action", types.YLeaf{"Action", node.Action})
    node.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", node.Vrf})
    node.EntityData.Leafs.Append("nexthop-ipv4-address", types.YLeaf{"NexthopIpv4Address", node.NexthopIpv4Address})
    node.EntityData.Leafs.Append("metatdata-disposition", types.YLeaf{"MetatdataDisposition", node.MetatdataDisposition})

    node.EntityData.YListKeys = []string {}

    return &(node.EntityData)
}

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames
// service function forwarder 
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // service function forwarder name. The type is slice of
    // Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName.
    SffName []*Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName
}

func (sffNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames) GetEntityData() *types.CommonEntityData {
    sffNames.EntityData.YFilter = sffNames.YFilter
    sffNames.EntityData.YangName = "sff-names"
    sffNames.EntityData.BundleName = "cisco_ios_xr"
    sffNames.EntityData.ParentYangName = "service-index"
    sffNames.EntityData.SegmentPath = "sff-names"
    sffNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sffNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sffNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sffNames.EntityData.Children = types.NewOrderedMap()
    sffNames.EntityData.Children.Append("sff-name", types.YChild{"SffName", nil})
    for i := range sffNames.SffName {
        sffNames.EntityData.Children.Append(types.GetSegmentPath(sffNames.SffName[i]), types.YChild{"SffName", sffNames.SffName[i]})
    }
    sffNames.EntityData.Leafs = types.NewOrderedMap()

    sffNames.EntityData.YListKeys = []string {}

    return &(sffNames.EntityData)
}

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName
// service function forwarder name
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SFF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // configure SFP.
    Node Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node
}

func (sffName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName) GetEntityData() *types.CommonEntityData {
    sffName.EntityData.YFilter = sffName.YFilter
    sffName.EntityData.YangName = "sff-name"
    sffName.EntityData.BundleName = "cisco_ios_xr"
    sffName.EntityData.ParentYangName = "sff-names"
    sffName.EntityData.SegmentPath = "sff-name" + types.AddKeyToken(sffName.Name, "name")
    sffName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sffName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sffName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sffName.EntityData.Children = types.NewOrderedMap()
    sffName.EntityData.Children.Append("node", types.YChild{"Node", &sffName.Node})
    sffName.EntityData.Leafs = types.NewOrderedMap()
    sffName.EntityData.Leafs.Append("name", types.YLeaf{"Name", sffName.Name})

    sffName.EntityData.YListKeys = []string {"Name"}

    return &(sffName.EntityData)
}

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node
// configure SFP
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Service function path. The type is interface{}.
    Enable interface{}

    // Dummy. The type is interface{}.
    Reserved interface{}
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SffNames_SffName_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "sff-name"
    node.EntityData.SegmentPath = "node"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", node.Enable})
    node.EntityData.Leafs.Append("reserved", types.YLeaf{"Reserved", node.Reserved})

    node.EntityData.YListKeys = []string {}

    return &(node.EntityData)
}

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames
// service function 
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // service function name. The type is slice of
    // Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName.
    SfName []*Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName
}

func (sfNames *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames) GetEntityData() *types.CommonEntityData {
    sfNames.EntityData.YFilter = sfNames.YFilter
    sfNames.EntityData.YangName = "sf-names"
    sfNames.EntityData.BundleName = "cisco_ios_xr"
    sfNames.EntityData.ParentYangName = "service-index"
    sfNames.EntityData.SegmentPath = "sf-names"
    sfNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfNames.EntityData.Children = types.NewOrderedMap()
    sfNames.EntityData.Children.Append("sf-name", types.YChild{"SfName", nil})
    for i := range sfNames.SfName {
        sfNames.EntityData.Children.Append(types.GetSegmentPath(sfNames.SfName[i]), types.YChild{"SfName", sfNames.SfName[i]})
    }
    sfNames.EntityData.Leafs = types.NewOrderedMap()

    sfNames.EntityData.YListKeys = []string {}

    return &(sfNames.EntityData)
}

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName
// service function name
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // configure SFP.
    Node Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node
}

func (sfName *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName) GetEntityData() *types.CommonEntityData {
    sfName.EntityData.YFilter = sfName.YFilter
    sfName.EntityData.YangName = "sf-name"
    sfName.EntityData.BundleName = "cisco_ios_xr"
    sfName.EntityData.ParentYangName = "sf-names"
    sfName.EntityData.SegmentPath = "sf-name" + types.AddKeyToken(sfName.Name, "name")
    sfName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfName.EntityData.Children = types.NewOrderedMap()
    sfName.EntityData.Children.Append("node", types.YChild{"Node", &sfName.Node})
    sfName.EntityData.Leafs = types.NewOrderedMap()
    sfName.EntityData.Leafs.Append("name", types.YLeaf{"Name", sfName.Name})

    sfName.EntityData.YListKeys = []string {"Name"}

    return &(sfName.EntityData)
}

// Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node
// configure SFP
type Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Service function path. The type is interface{}.
    Enable interface{}

    // Dummy. The type is interface{}.
    Reserved interface{}
}

func (node *Vservice_ServiceFunctionPath_Paths_Path_ServiceIndex_SfNames_SfName_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "sf-name"
    node.EntityData.SegmentPath = "node"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", node.Enable})
    node.EntityData.Leafs.Append("reserved", types.YLeaf{"Reserved", node.Reserved})

    node.EntityData.YListKeys = []string {}

    return &(node.EntityData)
}

