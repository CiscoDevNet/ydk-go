// This module contains a collection of YANG definitions for
// configuring diffserv specification implementations.
// 
// Copyright (c) 2014 IETF Trust and the persons identified as
// authors of the code.  All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD License
// set forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// This version of this YANG module is part of RFC XXXX; see
// the RFC itself for full legal notices.
package diffserv_classifier

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/ietf"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package diffserv_classifier"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-diffserv-classifier classifiers}", reflect.TypeOf(Classifiers{}))
    ydk.RegisterEntity("ietf-diffserv-classifier:classifiers", reflect.TypeOf(Classifiers{}))
}

type FilterType struct {
}

func (id FilterType) String() string {
	return "ietf-diffserv-classifier:filter-type"
}

type Dscp struct {
}

func (id Dscp) String() string {
	return "ietf-diffserv-classifier:dscp"
}

type SourceIpAddress struct {
}

func (id SourceIpAddress) String() string {
	return "ietf-diffserv-classifier:source-ip-address"
}

type DestinationIpAddress struct {
}

func (id DestinationIpAddress) String() string {
	return "ietf-diffserv-classifier:destination-ip-address"
}

type SourcePort struct {
}

func (id SourcePort) String() string {
	return "ietf-diffserv-classifier:source-port"
}

type DestinationPort struct {
}

func (id DestinationPort) String() string {
	return "ietf-diffserv-classifier:destination-port"
}

type Protocol struct {
}

func (id Protocol) String() string {
	return "ietf-diffserv-classifier:protocol"
}

type ClassifierEntryFilterOperationType struct {
}

func (id ClassifierEntryFilterOperationType) String() string {
	return "ietf-diffserv-classifier:classifier-entry-filter-operation-type"
}

type MatchAnyFilter struct {
}

func (id MatchAnyFilter) String() string {
	return "ietf-diffserv-classifier:match-any-filter"
}

type MatchAllFilter struct {
}

func (id MatchAllFilter) String() string {
	return "ietf-diffserv-classifier:match-all-filter"
}

// Classifiers
// list of classifier entry
type Classifiers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // classifier entry template. The type is slice of
    // Classifiers_ClassifierEntry.
    ClassifierEntry []Classifiers_ClassifierEntry
}

func (classifiers *Classifiers) GetEntityData() *types.CommonEntityData {
    classifiers.EntityData.YFilter = classifiers.YFilter
    classifiers.EntityData.YangName = "classifiers"
    classifiers.EntityData.BundleName = "ietf"
    classifiers.EntityData.ParentYangName = "ietf-diffserv-classifier"
    classifiers.EntityData.SegmentPath = "ietf-diffserv-classifier:classifiers"
    classifiers.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    classifiers.EntityData.NamespaceTable = ietf.GetNamespaces()
    classifiers.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    classifiers.EntityData.Children = make(map[string]types.YChild)
    classifiers.EntityData.Children["classifier-entry"] = types.YChild{"ClassifierEntry", nil}
    for i := range classifiers.ClassifierEntry {
        classifiers.EntityData.Children[types.GetSegmentPath(&classifiers.ClassifierEntry[i])] = types.YChild{"ClassifierEntry", &classifiers.ClassifierEntry[i]}
    }
    classifiers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(classifiers.EntityData)
}

// Classifiers_ClassifierEntry
// classifier entry template
type Classifiers_ClassifierEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Diffserv classifier name. The type is string.
    ClassifierEntryName interface{}

    // Description of the class template. The type is string.
    ClassifierEntryDescr interface{}

    // Filters are applicable as any or all filters. The type is one of the
    // following: MatchAnyFilterMatchAllFilter. The default value is
    // match-any-filter.
    ClassifierEntryFilterOperation interface{}

    // Filter configuration. The type is slice of
    // Classifiers_ClassifierEntry_FilterEntry.
    FilterEntry []Classifiers_ClassifierEntry_FilterEntry
}

func (classifierEntry *Classifiers_ClassifierEntry) GetEntityData() *types.CommonEntityData {
    classifierEntry.EntityData.YFilter = classifierEntry.YFilter
    classifierEntry.EntityData.YangName = "classifier-entry"
    classifierEntry.EntityData.BundleName = "ietf"
    classifierEntry.EntityData.ParentYangName = "classifiers"
    classifierEntry.EntityData.SegmentPath = "classifier-entry" + "[classifier-entry-name='" + fmt.Sprintf("%v", classifierEntry.ClassifierEntryName) + "']"
    classifierEntry.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    classifierEntry.EntityData.NamespaceTable = ietf.GetNamespaces()
    classifierEntry.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    classifierEntry.EntityData.Children = make(map[string]types.YChild)
    classifierEntry.EntityData.Children["filter-entry"] = types.YChild{"FilterEntry", nil}
    for i := range classifierEntry.FilterEntry {
        classifierEntry.EntityData.Children[types.GetSegmentPath(&classifierEntry.FilterEntry[i])] = types.YChild{"FilterEntry", &classifierEntry.FilterEntry[i]}
    }
    classifierEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    classifierEntry.EntityData.Leafs["classifier-entry-name"] = types.YLeaf{"ClassifierEntryName", classifierEntry.ClassifierEntryName}
    classifierEntry.EntityData.Leafs["classifier-entry-descr"] = types.YLeaf{"ClassifierEntryDescr", classifierEntry.ClassifierEntryDescr}
    classifierEntry.EntityData.Leafs["classifier-entry-filter-operation"] = types.YLeaf{"ClassifierEntryFilterOperation", classifierEntry.ClassifierEntryFilterOperation}
    return &(classifierEntry.EntityData)
}

// Classifiers_ClassifierEntry_FilterEntry
// Filter configuration
type Classifiers_ClassifierEntry_FilterEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This leaf defines type of the filter. The type is
    // one of the following:
    // DscpSourceIpAddressDestinationIpAddressSourcePortDestinationPortProtocolCosCosInnerIpv4AclNameIpv6AclNameIpv4AclIpv6AclInputInterfaceSrcMacDstMacMplsExpTopMplsExpImpPacketLengthPrecQosGroupVlanVlanInnerAtmClpAtmVciDeiDeiInnerFlowIpFlowRecordFlowDeFlowDlciWlanUserPriorityDiscardClassClassMapMetadataApplicationSecurityGroupNameSecurityGroupTagIpRtpVpls.
    FilterType interface{}

    // This attribute is a key.  This is logical-not operator for a filter. When
    // true, it  indicates filter looks for absence of a pattern defined  by the
    // filter . The type is bool.
    FilterLogicalNot interface{}

    // list of dscp ranges. The type is slice of
    // Classifiers_ClassifierEntry_FilterEntry_DscpCfg.
    DscpCfg []Classifiers_ClassifierEntry_FilterEntry_DscpCfg

    // list of source ip address. The type is slice of
    // Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg.
    SourceIpAddressCfg []Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg

    // list of destination ip address. The type is slice of
    // Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg.
    DestinationIpAddressCfg []Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg

    // list of ranges of source port. The type is slice of
    // Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg.
    SourcePortCfg []Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg

    // list of ranges of destination port. The type is slice of
    // Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg.
    DestinationPortCfg []Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg

    // list of ranges of protocol values. The type is slice of
    // Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg.
    ProtocolCfg []Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg
}

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetEntityData() *types.CommonEntityData {
    filterEntry.EntityData.YFilter = filterEntry.YFilter
    filterEntry.EntityData.YangName = "filter-entry"
    filterEntry.EntityData.BundleName = "ietf"
    filterEntry.EntityData.ParentYangName = "classifier-entry"
    filterEntry.EntityData.SegmentPath = "filter-entry" + "[filter-type='" + fmt.Sprintf("%v", filterEntry.FilterType) + "']" + "[filter-logical-not='" + fmt.Sprintf("%v", filterEntry.FilterLogicalNot) + "']"
    filterEntry.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    filterEntry.EntityData.NamespaceTable = ietf.GetNamespaces()
    filterEntry.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    filterEntry.EntityData.Children = make(map[string]types.YChild)
    filterEntry.EntityData.Children["dscp-cfg"] = types.YChild{"DscpCfg", nil}
    for i := range filterEntry.DscpCfg {
        filterEntry.EntityData.Children[types.GetSegmentPath(&filterEntry.DscpCfg[i])] = types.YChild{"DscpCfg", &filterEntry.DscpCfg[i]}
    }
    filterEntry.EntityData.Children["source-ip-address-cfg"] = types.YChild{"SourceIpAddressCfg", nil}
    for i := range filterEntry.SourceIpAddressCfg {
        filterEntry.EntityData.Children[types.GetSegmentPath(&filterEntry.SourceIpAddressCfg[i])] = types.YChild{"SourceIpAddressCfg", &filterEntry.SourceIpAddressCfg[i]}
    }
    filterEntry.EntityData.Children["destination-ip-address-cfg"] = types.YChild{"DestinationIpAddressCfg", nil}
    for i := range filterEntry.DestinationIpAddressCfg {
        filterEntry.EntityData.Children[types.GetSegmentPath(&filterEntry.DestinationIpAddressCfg[i])] = types.YChild{"DestinationIpAddressCfg", &filterEntry.DestinationIpAddressCfg[i]}
    }
    filterEntry.EntityData.Children["source-port-cfg"] = types.YChild{"SourcePortCfg", nil}
    for i := range filterEntry.SourcePortCfg {
        filterEntry.EntityData.Children[types.GetSegmentPath(&filterEntry.SourcePortCfg[i])] = types.YChild{"SourcePortCfg", &filterEntry.SourcePortCfg[i]}
    }
    filterEntry.EntityData.Children["destination-port-cfg"] = types.YChild{"DestinationPortCfg", nil}
    for i := range filterEntry.DestinationPortCfg {
        filterEntry.EntityData.Children[types.GetSegmentPath(&filterEntry.DestinationPortCfg[i])] = types.YChild{"DestinationPortCfg", &filterEntry.DestinationPortCfg[i]}
    }
    filterEntry.EntityData.Children["protocol-cfg"] = types.YChild{"ProtocolCfg", nil}
    for i := range filterEntry.ProtocolCfg {
        filterEntry.EntityData.Children[types.GetSegmentPath(&filterEntry.ProtocolCfg[i])] = types.YChild{"ProtocolCfg", &filterEntry.ProtocolCfg[i]}
    }
    filterEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    filterEntry.EntityData.Leafs["filter-type"] = types.YLeaf{"FilterType", filterEntry.FilterType}
    filterEntry.EntityData.Leafs["filter-logical-not"] = types.YLeaf{"FilterLogicalNot", filterEntry.FilterLogicalNot}
    return &(filterEntry.EntityData)
}

// Classifiers_ClassifierEntry_FilterEntry_DscpCfg
// list of dscp ranges
type Classifiers_ClassifierEntry_FilterEntry_DscpCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Minimum value of dscp range. The type is
    // interface{} with range: 0..63.
    DscpMin interface{}

    // This attribute is a key. maximum value of dscp range. The type is
    // interface{} with range: 0..63.
    DscpMax interface{}
}

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetEntityData() *types.CommonEntityData {
    dscpCfg.EntityData.YFilter = dscpCfg.YFilter
    dscpCfg.EntityData.YangName = "dscp-cfg"
    dscpCfg.EntityData.BundleName = "ietf"
    dscpCfg.EntityData.ParentYangName = "filter-entry"
    dscpCfg.EntityData.SegmentPath = "dscp-cfg" + "[dscp-min='" + fmt.Sprintf("%v", dscpCfg.DscpMin) + "']" + "[dscp-max='" + fmt.Sprintf("%v", dscpCfg.DscpMax) + "']"
    dscpCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    dscpCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    dscpCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    dscpCfg.EntityData.Children = make(map[string]types.YChild)
    dscpCfg.EntityData.Leafs = make(map[string]types.YLeaf)
    dscpCfg.EntityData.Leafs["dscp-min"] = types.YLeaf{"DscpMin", dscpCfg.DscpMin}
    dscpCfg.EntityData.Leafs["dscp-max"] = types.YLeaf{"DscpMax", dscpCfg.DscpMax}
    return &(dscpCfg.EntityData)
}

// Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg
// list of source ip address
type Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. source ip prefix. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    SourceIpAddr interface{}
}

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetEntityData() *types.CommonEntityData {
    sourceIpAddressCfg.EntityData.YFilter = sourceIpAddressCfg.YFilter
    sourceIpAddressCfg.EntityData.YangName = "source-ip-address-cfg"
    sourceIpAddressCfg.EntityData.BundleName = "ietf"
    sourceIpAddressCfg.EntityData.ParentYangName = "filter-entry"
    sourceIpAddressCfg.EntityData.SegmentPath = "source-ip-address-cfg" + "[source-ip-addr='" + fmt.Sprintf("%v", sourceIpAddressCfg.SourceIpAddr) + "']"
    sourceIpAddressCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    sourceIpAddressCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    sourceIpAddressCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    sourceIpAddressCfg.EntityData.Children = make(map[string]types.YChild)
    sourceIpAddressCfg.EntityData.Leafs = make(map[string]types.YLeaf)
    sourceIpAddressCfg.EntityData.Leafs["source-ip-addr"] = types.YLeaf{"SourceIpAddr", sourceIpAddressCfg.SourceIpAddr}
    return &(sourceIpAddressCfg.EntityData)
}

// Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg
// list of destination ip address
type Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. destination ip prefix. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    DestinationIpAddr interface{}
}

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetEntityData() *types.CommonEntityData {
    destinationIpAddressCfg.EntityData.YFilter = destinationIpAddressCfg.YFilter
    destinationIpAddressCfg.EntityData.YangName = "destination-ip-address-cfg"
    destinationIpAddressCfg.EntityData.BundleName = "ietf"
    destinationIpAddressCfg.EntityData.ParentYangName = "filter-entry"
    destinationIpAddressCfg.EntityData.SegmentPath = "destination-ip-address-cfg" + "[destination-ip-addr='" + fmt.Sprintf("%v", destinationIpAddressCfg.DestinationIpAddr) + "']"
    destinationIpAddressCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    destinationIpAddressCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    destinationIpAddressCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    destinationIpAddressCfg.EntityData.Children = make(map[string]types.YChild)
    destinationIpAddressCfg.EntityData.Leafs = make(map[string]types.YLeaf)
    destinationIpAddressCfg.EntityData.Leafs["destination-ip-addr"] = types.YLeaf{"DestinationIpAddr", destinationIpAddressCfg.DestinationIpAddr}
    return &(destinationIpAddressCfg.EntityData)
}

// Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg
// list of ranges of source port
type Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. minimum value of source port range. The type is
    // interface{} with range: 0..65535.
    SourcePortMin interface{}

    // This attribute is a key. maximum value of source port range. The type is
    // interface{} with range: 0..65535.
    SourcePortMax interface{}
}

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetEntityData() *types.CommonEntityData {
    sourcePortCfg.EntityData.YFilter = sourcePortCfg.YFilter
    sourcePortCfg.EntityData.YangName = "source-port-cfg"
    sourcePortCfg.EntityData.BundleName = "ietf"
    sourcePortCfg.EntityData.ParentYangName = "filter-entry"
    sourcePortCfg.EntityData.SegmentPath = "source-port-cfg" + "[source-port-min='" + fmt.Sprintf("%v", sourcePortCfg.SourcePortMin) + "']" + "[source-port-max='" + fmt.Sprintf("%v", sourcePortCfg.SourcePortMax) + "']"
    sourcePortCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    sourcePortCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    sourcePortCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    sourcePortCfg.EntityData.Children = make(map[string]types.YChild)
    sourcePortCfg.EntityData.Leafs = make(map[string]types.YLeaf)
    sourcePortCfg.EntityData.Leafs["source-port-min"] = types.YLeaf{"SourcePortMin", sourcePortCfg.SourcePortMin}
    sourcePortCfg.EntityData.Leafs["source-port-max"] = types.YLeaf{"SourcePortMax", sourcePortCfg.SourcePortMax}
    return &(sourcePortCfg.EntityData)
}

// Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg
// list of ranges of destination port
type Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. minimum value of destination port range. The type
    // is interface{} with range: 0..65535.
    DestinationPortMin interface{}

    // This attribute is a key. maximum value of destination port range. The type
    // is interface{} with range: 0..65535.
    DestinationPortMax interface{}
}

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetEntityData() *types.CommonEntityData {
    destinationPortCfg.EntityData.YFilter = destinationPortCfg.YFilter
    destinationPortCfg.EntityData.YangName = "destination-port-cfg"
    destinationPortCfg.EntityData.BundleName = "ietf"
    destinationPortCfg.EntityData.ParentYangName = "filter-entry"
    destinationPortCfg.EntityData.SegmentPath = "destination-port-cfg" + "[destination-port-min='" + fmt.Sprintf("%v", destinationPortCfg.DestinationPortMin) + "']" + "[destination-port-max='" + fmt.Sprintf("%v", destinationPortCfg.DestinationPortMax) + "']"
    destinationPortCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    destinationPortCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    destinationPortCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    destinationPortCfg.EntityData.Children = make(map[string]types.YChild)
    destinationPortCfg.EntityData.Leafs = make(map[string]types.YLeaf)
    destinationPortCfg.EntityData.Leafs["destination-port-min"] = types.YLeaf{"DestinationPortMin", destinationPortCfg.DestinationPortMin}
    destinationPortCfg.EntityData.Leafs["destination-port-max"] = types.YLeaf{"DestinationPortMax", destinationPortCfg.DestinationPortMax}
    return &(destinationPortCfg.EntityData)
}

// Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg
// list of ranges of protocol values
type Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. minimum value of protocol range. The type is
    // interface{} with range: 0..255.
    ProtocolMin interface{}

    // This attribute is a key. maximum value of protocol range. The type is
    // interface{} with range: 0..255.
    ProtocolMax interface{}
}

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetEntityData() *types.CommonEntityData {
    protocolCfg.EntityData.YFilter = protocolCfg.YFilter
    protocolCfg.EntityData.YangName = "protocol-cfg"
    protocolCfg.EntityData.BundleName = "ietf"
    protocolCfg.EntityData.ParentYangName = "filter-entry"
    protocolCfg.EntityData.SegmentPath = "protocol-cfg" + "[protocol-min='" + fmt.Sprintf("%v", protocolCfg.ProtocolMin) + "']" + "[protocol-max='" + fmt.Sprintf("%v", protocolCfg.ProtocolMax) + "']"
    protocolCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    protocolCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    protocolCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    protocolCfg.EntityData.Children = make(map[string]types.YChild)
    protocolCfg.EntityData.Leafs = make(map[string]types.YLeaf)
    protocolCfg.EntityData.Leafs["protocol-min"] = types.YLeaf{"ProtocolMin", protocolCfg.ProtocolMin}
    protocolCfg.EntityData.Leafs["protocol-max"] = types.YLeaf{"ProtocolMax", protocolCfg.ProtocolMax}
    return &(protocolCfg.EntityData)
}

