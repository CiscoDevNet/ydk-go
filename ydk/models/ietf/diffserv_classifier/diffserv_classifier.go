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

type DestinationIpAddress struct {
}

func (id DestinationIpAddress) String() string {
	return "ietf-diffserv-classifier:destination-ip-address"
}

type Dscp struct {
}

func (id Dscp) String() string {
	return "ietf-diffserv-classifier:dscp"
}

type MatchAllFilter struct {
}

func (id MatchAllFilter) String() string {
	return "ietf-diffserv-classifier:match-all-filter"
}

type SourceIpAddress struct {
}

func (id SourceIpAddress) String() string {
	return "ietf-diffserv-classifier:source-ip-address"
}

type MatchAnyFilter struct {
}

func (id MatchAnyFilter) String() string {
	return "ietf-diffserv-classifier:match-any-filter"
}

type SourcePort struct {
}

func (id SourcePort) String() string {
	return "ietf-diffserv-classifier:source-port"
}

type FilterType struct {
}

func (id FilterType) String() string {
	return "ietf-diffserv-classifier:filter-type"
}

type ClassifierEntryFilterOperationType struct {
}

func (id ClassifierEntryFilterOperationType) String() string {
	return "ietf-diffserv-classifier:classifier-entry-filter-operation-type"
}

// Classifiers
// list of classifier entry
type Classifiers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // classifier entry template. The type is slice of
    // Classifiers_ClassifierEntry.
    ClassifierEntry []Classifiers_ClassifierEntry
}

func (classifiers *Classifiers) GetFilter() yfilter.YFilter { return classifiers.YFilter }

func (classifiers *Classifiers) SetFilter(yf yfilter.YFilter) { classifiers.YFilter = yf }

func (classifiers *Classifiers) GetGoName(yname string) string {
    if yname == "classifier-entry" { return "ClassifierEntry" }
    return ""
}

func (classifiers *Classifiers) GetSegmentPath() string {
    return "ietf-diffserv-classifier:classifiers"
}

func (classifiers *Classifiers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "classifier-entry" {
        for _, c := range classifiers.ClassifierEntry {
            if classifiers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Classifiers_ClassifierEntry{}
        classifiers.ClassifierEntry = append(classifiers.ClassifierEntry, child)
        return &classifiers.ClassifierEntry[len(classifiers.ClassifierEntry)-1]
    }
    return nil
}

func (classifiers *Classifiers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classifiers.ClassifierEntry {
        children[classifiers.ClassifierEntry[i].GetSegmentPath()] = &classifiers.ClassifierEntry[i]
    }
    return children
}

func (classifiers *Classifiers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classifiers *Classifiers) GetBundleName() string { return "ietf" }

func (classifiers *Classifiers) GetYangName() string { return "classifiers" }

func (classifiers *Classifiers) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (classifiers *Classifiers) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (classifiers *Classifiers) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (classifiers *Classifiers) SetParent(parent types.Entity) { classifiers.parent = parent }

func (classifiers *Classifiers) GetParent() types.Entity { return classifiers.parent }

func (classifiers *Classifiers) GetParentYangName() string { return "ietf-diffserv-classifier" }

// Classifiers_ClassifierEntry
// classifier entry template
type Classifiers_ClassifierEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Diffserv classifier name. The type is string.
    ClassifierEntryName interface{}

    // Description of the class template. The type is string.
    ClassifierEntryDescr interface{}

    // Filters are applicable as any or all filters. The type is one of the
    // following: MatchAllFilterMatchAnyFilter. The default value is
    // match-any-filter.
    ClassifierEntryFilterOperation interface{}

    // Filter configuration. The type is slice of
    // Classifiers_ClassifierEntry_FilterEntry.
    FilterEntry []Classifiers_ClassifierEntry_FilterEntry
}

func (classifierEntry *Classifiers_ClassifierEntry) GetFilter() yfilter.YFilter { return classifierEntry.YFilter }

func (classifierEntry *Classifiers_ClassifierEntry) SetFilter(yf yfilter.YFilter) { classifierEntry.YFilter = yf }

func (classifierEntry *Classifiers_ClassifierEntry) GetGoName(yname string) string {
    if yname == "classifier-entry-name" { return "ClassifierEntryName" }
    if yname == "classifier-entry-descr" { return "ClassifierEntryDescr" }
    if yname == "classifier-entry-filter-operation" { return "ClassifierEntryFilterOperation" }
    if yname == "filter-entry" { return "FilterEntry" }
    return ""
}

func (classifierEntry *Classifiers_ClassifierEntry) GetSegmentPath() string {
    return "classifier-entry" + "[classifier-entry-name='" + fmt.Sprintf("%v", classifierEntry.ClassifierEntryName) + "']"
}

func (classifierEntry *Classifiers_ClassifierEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "filter-entry" {
        for _, c := range classifierEntry.FilterEntry {
            if classifierEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Classifiers_ClassifierEntry_FilterEntry{}
        classifierEntry.FilterEntry = append(classifierEntry.FilterEntry, child)
        return &classifierEntry.FilterEntry[len(classifierEntry.FilterEntry)-1]
    }
    return nil
}

func (classifierEntry *Classifiers_ClassifierEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classifierEntry.FilterEntry {
        children[classifierEntry.FilterEntry[i].GetSegmentPath()] = &classifierEntry.FilterEntry[i]
    }
    return children
}

func (classifierEntry *Classifiers_ClassifierEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["classifier-entry-name"] = classifierEntry.ClassifierEntryName
    leafs["classifier-entry-descr"] = classifierEntry.ClassifierEntryDescr
    leafs["classifier-entry-filter-operation"] = classifierEntry.ClassifierEntryFilterOperation
    return leafs
}

func (classifierEntry *Classifiers_ClassifierEntry) GetBundleName() string { return "ietf" }

func (classifierEntry *Classifiers_ClassifierEntry) GetYangName() string { return "classifier-entry" }

func (classifierEntry *Classifiers_ClassifierEntry) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (classifierEntry *Classifiers_ClassifierEntry) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (classifierEntry *Classifiers_ClassifierEntry) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (classifierEntry *Classifiers_ClassifierEntry) SetParent(parent types.Entity) { classifierEntry.parent = parent }

func (classifierEntry *Classifiers_ClassifierEntry) GetParent() types.Entity { return classifierEntry.parent }

func (classifierEntry *Classifiers_ClassifierEntry) GetParentYangName() string { return "classifiers" }

// Classifiers_ClassifierEntry_FilterEntry
// Filter configuration
type Classifiers_ClassifierEntry_FilterEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This leaf defines type of the filter. The type is
    // one of the following:
    // DestinationPortProtocolDestinationIpAddressDscpSourceIpAddressSourcePortInputInterfaceSrcMacApplicationSecurityGroupNameIpv4AclNameFlowDlciDeiPrecPacketLengthIpv4AclFlowDeFlowIpFlowRecordVlanInnerMetadataVlanAtmVciClassMapQosGroupWlanUserPriorityIpRtpIpv6AclAtmClpDstMacCosDeiInnerMplsExpTopCosInnerIpv6AclNameMplsExpImpSecurityGroupTagDiscardClassVpls.
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

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetFilter() yfilter.YFilter { return filterEntry.YFilter }

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) SetFilter(yf yfilter.YFilter) { filterEntry.YFilter = yf }

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetGoName(yname string) string {
    if yname == "filter-type" { return "FilterType" }
    if yname == "filter-logical-not" { return "FilterLogicalNot" }
    if yname == "dscp-cfg" { return "DscpCfg" }
    if yname == "source-ip-address-cfg" { return "SourceIpAddressCfg" }
    if yname == "destination-ip-address-cfg" { return "DestinationIpAddressCfg" }
    if yname == "source-port-cfg" { return "SourcePortCfg" }
    if yname == "destination-port-cfg" { return "DestinationPortCfg" }
    if yname == "protocol-cfg" { return "ProtocolCfg" }
    return ""
}

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetSegmentPath() string {
    return "filter-entry" + "[filter-type='" + fmt.Sprintf("%v", filterEntry.FilterType) + "']" + "[filter-logical-not='" + fmt.Sprintf("%v", filterEntry.FilterLogicalNot) + "']"
}

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dscp-cfg" {
        for _, c := range filterEntry.DscpCfg {
            if filterEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Classifiers_ClassifierEntry_FilterEntry_DscpCfg{}
        filterEntry.DscpCfg = append(filterEntry.DscpCfg, child)
        return &filterEntry.DscpCfg[len(filterEntry.DscpCfg)-1]
    }
    if childYangName == "source-ip-address-cfg" {
        for _, c := range filterEntry.SourceIpAddressCfg {
            if filterEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg{}
        filterEntry.SourceIpAddressCfg = append(filterEntry.SourceIpAddressCfg, child)
        return &filterEntry.SourceIpAddressCfg[len(filterEntry.SourceIpAddressCfg)-1]
    }
    if childYangName == "destination-ip-address-cfg" {
        for _, c := range filterEntry.DestinationIpAddressCfg {
            if filterEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg{}
        filterEntry.DestinationIpAddressCfg = append(filterEntry.DestinationIpAddressCfg, child)
        return &filterEntry.DestinationIpAddressCfg[len(filterEntry.DestinationIpAddressCfg)-1]
    }
    if childYangName == "source-port-cfg" {
        for _, c := range filterEntry.SourcePortCfg {
            if filterEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg{}
        filterEntry.SourcePortCfg = append(filterEntry.SourcePortCfg, child)
        return &filterEntry.SourcePortCfg[len(filterEntry.SourcePortCfg)-1]
    }
    if childYangName == "destination-port-cfg" {
        for _, c := range filterEntry.DestinationPortCfg {
            if filterEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg{}
        filterEntry.DestinationPortCfg = append(filterEntry.DestinationPortCfg, child)
        return &filterEntry.DestinationPortCfg[len(filterEntry.DestinationPortCfg)-1]
    }
    if childYangName == "protocol-cfg" {
        for _, c := range filterEntry.ProtocolCfg {
            if filterEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg{}
        filterEntry.ProtocolCfg = append(filterEntry.ProtocolCfg, child)
        return &filterEntry.ProtocolCfg[len(filterEntry.ProtocolCfg)-1]
    }
    return nil
}

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range filterEntry.DscpCfg {
        children[filterEntry.DscpCfg[i].GetSegmentPath()] = &filterEntry.DscpCfg[i]
    }
    for i := range filterEntry.SourceIpAddressCfg {
        children[filterEntry.SourceIpAddressCfg[i].GetSegmentPath()] = &filterEntry.SourceIpAddressCfg[i]
    }
    for i := range filterEntry.DestinationIpAddressCfg {
        children[filterEntry.DestinationIpAddressCfg[i].GetSegmentPath()] = &filterEntry.DestinationIpAddressCfg[i]
    }
    for i := range filterEntry.SourcePortCfg {
        children[filterEntry.SourcePortCfg[i].GetSegmentPath()] = &filterEntry.SourcePortCfg[i]
    }
    for i := range filterEntry.DestinationPortCfg {
        children[filterEntry.DestinationPortCfg[i].GetSegmentPath()] = &filterEntry.DestinationPortCfg[i]
    }
    for i := range filterEntry.ProtocolCfg {
        children[filterEntry.ProtocolCfg[i].GetSegmentPath()] = &filterEntry.ProtocolCfg[i]
    }
    return children
}

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["filter-type"] = filterEntry.FilterType
    leafs["filter-logical-not"] = filterEntry.FilterLogicalNot
    return leafs
}

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetBundleName() string { return "ietf" }

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetYangName() string { return "filter-entry" }

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) SetParent(parent types.Entity) { filterEntry.parent = parent }

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetParent() types.Entity { return filterEntry.parent }

func (filterEntry *Classifiers_ClassifierEntry_FilterEntry) GetParentYangName() string { return "classifier-entry" }

// Classifiers_ClassifierEntry_FilterEntry_DscpCfg
// list of dscp ranges
type Classifiers_ClassifierEntry_FilterEntry_DscpCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Minimum value of dscp range. The type is
    // interface{} with range: 0..63.
    DscpMin interface{}

    // This attribute is a key. maximum value of dscp range. The type is
    // interface{} with range: 0..63.
    DscpMax interface{}
}

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetFilter() yfilter.YFilter { return dscpCfg.YFilter }

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) SetFilter(yf yfilter.YFilter) { dscpCfg.YFilter = yf }

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetGoName(yname string) string {
    if yname == "dscp-min" { return "DscpMin" }
    if yname == "dscp-max" { return "DscpMax" }
    return ""
}

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetSegmentPath() string {
    return "dscp-cfg" + "[dscp-min='" + fmt.Sprintf("%v", dscpCfg.DscpMin) + "']" + "[dscp-max='" + fmt.Sprintf("%v", dscpCfg.DscpMax) + "']"
}

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp-min"] = dscpCfg.DscpMin
    leafs["dscp-max"] = dscpCfg.DscpMax
    return leafs
}

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetBundleName() string { return "ietf" }

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetYangName() string { return "dscp-cfg" }

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) SetParent(parent types.Entity) { dscpCfg.parent = parent }

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetParent() types.Entity { return dscpCfg.parent }

func (dscpCfg *Classifiers_ClassifierEntry_FilterEntry_DscpCfg) GetParentYangName() string { return "filter-entry" }

// Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg
// list of source ip address
type Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. source ip prefix. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    SourceIpAddr interface{}
}

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetFilter() yfilter.YFilter { return sourceIpAddressCfg.YFilter }

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) SetFilter(yf yfilter.YFilter) { sourceIpAddressCfg.YFilter = yf }

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetGoName(yname string) string {
    if yname == "source-ip-addr" { return "SourceIpAddr" }
    return ""
}

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetSegmentPath() string {
    return "source-ip-address-cfg" + "[source-ip-addr='" + fmt.Sprintf("%v", sourceIpAddressCfg.SourceIpAddr) + "']"
}

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-ip-addr"] = sourceIpAddressCfg.SourceIpAddr
    return leafs
}

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetBundleName() string { return "ietf" }

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetYangName() string { return "source-ip-address-cfg" }

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) SetParent(parent types.Entity) { sourceIpAddressCfg.parent = parent }

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetParent() types.Entity { return sourceIpAddressCfg.parent }

func (sourceIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetParentYangName() string { return "filter-entry" }

// Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg
// list of destination ip address
type Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. destination ip prefix. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    DestinationIpAddr interface{}
}

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetFilter() yfilter.YFilter { return destinationIpAddressCfg.YFilter }

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) SetFilter(yf yfilter.YFilter) { destinationIpAddressCfg.YFilter = yf }

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetGoName(yname string) string {
    if yname == "destination-ip-addr" { return "DestinationIpAddr" }
    return ""
}

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetSegmentPath() string {
    return "destination-ip-address-cfg" + "[destination-ip-addr='" + fmt.Sprintf("%v", destinationIpAddressCfg.DestinationIpAddr) + "']"
}

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-ip-addr"] = destinationIpAddressCfg.DestinationIpAddr
    return leafs
}

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetBundleName() string { return "ietf" }

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetYangName() string { return "destination-ip-address-cfg" }

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) SetParent(parent types.Entity) { destinationIpAddressCfg.parent = parent }

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetParent() types.Entity { return destinationIpAddressCfg.parent }

func (destinationIpAddressCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetParentYangName() string { return "filter-entry" }

// Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg
// list of ranges of source port
type Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. minimum value of source port range. The type is
    // interface{} with range: 0..65535.
    SourcePortMin interface{}

    // This attribute is a key. maximum value of source port range. The type is
    // interface{} with range: 0..65535.
    SourcePortMax interface{}
}

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetFilter() yfilter.YFilter { return sourcePortCfg.YFilter }

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) SetFilter(yf yfilter.YFilter) { sourcePortCfg.YFilter = yf }

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetGoName(yname string) string {
    if yname == "source-port-min" { return "SourcePortMin" }
    if yname == "source-port-max" { return "SourcePortMax" }
    return ""
}

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetSegmentPath() string {
    return "source-port-cfg" + "[source-port-min='" + fmt.Sprintf("%v", sourcePortCfg.SourcePortMin) + "']" + "[source-port-max='" + fmt.Sprintf("%v", sourcePortCfg.SourcePortMax) + "']"
}

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-port-min"] = sourcePortCfg.SourcePortMin
    leafs["source-port-max"] = sourcePortCfg.SourcePortMax
    return leafs
}

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetBundleName() string { return "ietf" }

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetYangName() string { return "source-port-cfg" }

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) SetParent(parent types.Entity) { sourcePortCfg.parent = parent }

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetParent() types.Entity { return sourcePortCfg.parent }

func (sourcePortCfg *Classifiers_ClassifierEntry_FilterEntry_SourcePortCfg) GetParentYangName() string { return "filter-entry" }

// Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg
// list of ranges of destination port
type Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. minimum value of destination port range. The type
    // is interface{} with range: 0..65535.
    DestinationPortMin interface{}

    // This attribute is a key. maximum value of destination port range. The type
    // is interface{} with range: 0..65535.
    DestinationPortMax interface{}
}

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetFilter() yfilter.YFilter { return destinationPortCfg.YFilter }

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) SetFilter(yf yfilter.YFilter) { destinationPortCfg.YFilter = yf }

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetGoName(yname string) string {
    if yname == "destination-port-min" { return "DestinationPortMin" }
    if yname == "destination-port-max" { return "DestinationPortMax" }
    return ""
}

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetSegmentPath() string {
    return "destination-port-cfg" + "[destination-port-min='" + fmt.Sprintf("%v", destinationPortCfg.DestinationPortMin) + "']" + "[destination-port-max='" + fmt.Sprintf("%v", destinationPortCfg.DestinationPortMax) + "']"
}

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-port-min"] = destinationPortCfg.DestinationPortMin
    leafs["destination-port-max"] = destinationPortCfg.DestinationPortMax
    return leafs
}

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetBundleName() string { return "ietf" }

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetYangName() string { return "destination-port-cfg" }

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) SetParent(parent types.Entity) { destinationPortCfg.parent = parent }

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetParent() types.Entity { return destinationPortCfg.parent }

func (destinationPortCfg *Classifiers_ClassifierEntry_FilterEntry_DestinationPortCfg) GetParentYangName() string { return "filter-entry" }

// Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg
// list of ranges of protocol values
type Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. minimum value of protocol range. The type is
    // interface{} with range: 0..255.
    ProtocolMin interface{}

    // This attribute is a key. maximum value of protocol range. The type is
    // interface{} with range: 0..255.
    ProtocolMax interface{}
}

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetFilter() yfilter.YFilter { return protocolCfg.YFilter }

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) SetFilter(yf yfilter.YFilter) { protocolCfg.YFilter = yf }

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetGoName(yname string) string {
    if yname == "protocol-min" { return "ProtocolMin" }
    if yname == "protocol-max" { return "ProtocolMax" }
    return ""
}

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetSegmentPath() string {
    return "protocol-cfg" + "[protocol-min='" + fmt.Sprintf("%v", protocolCfg.ProtocolMin) + "']" + "[protocol-max='" + fmt.Sprintf("%v", protocolCfg.ProtocolMax) + "']"
}

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol-min"] = protocolCfg.ProtocolMin
    leafs["protocol-max"] = protocolCfg.ProtocolMax
    return leafs
}

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetBundleName() string { return "ietf" }

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetYangName() string { return "protocol-cfg" }

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) SetParent(parent types.Entity) { protocolCfg.parent = parent }

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetParent() types.Entity { return protocolCfg.parent }

func (protocolCfg *Classifiers_ClassifierEntry_FilterEntry_ProtocolCfg) GetParentYangName() string { return "filter-entry" }

