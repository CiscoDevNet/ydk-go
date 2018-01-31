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
package diffserv_policy

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/ietf"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package diffserv_policy"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-diffserv-policy policies}", reflect.TypeOf(Policies{}))
    ydk.RegisterEntity("ietf-diffserv-policy:policies", reflect.TypeOf(Policies{}))
}

type ActionType struct {
}

func (id ActionType) String() string {
	return "ietf-diffserv-policy:action-type"
}

// Policies
// list of policy templates
type Policies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // policy template. The type is slice of Policies_PolicyEntry.
    PolicyEntry []Policies_PolicyEntry
}

func (policies *Policies) GetFilter() yfilter.YFilter { return policies.YFilter }

func (policies *Policies) SetFilter(yf yfilter.YFilter) { policies.YFilter = yf }

func (policies *Policies) GetGoName(yname string) string {
    if yname == "policy-entry" { return "PolicyEntry" }
    return ""
}

func (policies *Policies) GetSegmentPath() string {
    return "ietf-diffserv-policy:policies"
}

func (policies *Policies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-entry" {
        for _, c := range policies.PolicyEntry {
            if policies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Policies_PolicyEntry{}
        policies.PolicyEntry = append(policies.PolicyEntry, child)
        return &policies.PolicyEntry[len(policies.PolicyEntry)-1]
    }
    return nil
}

func (policies *Policies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policies.PolicyEntry {
        children[policies.PolicyEntry[i].GetSegmentPath()] = &policies.PolicyEntry[i]
    }
    return children
}

func (policies *Policies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policies *Policies) GetBundleName() string { return "ietf" }

func (policies *Policies) GetYangName() string { return "policies" }

func (policies *Policies) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (policies *Policies) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (policies *Policies) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (policies *Policies) SetParent(parent types.Entity) { policies.parent = parent }

func (policies *Policies) GetParent() types.Entity { return policies.parent }

func (policies *Policies) GetParentYangName() string { return "ietf-diffserv-policy" }

// Policies_PolicyEntry
// policy template
type Policies_PolicyEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Diffserv policy name. The type is string.
    PolicyName interface{}

    // Diffserv policy description. The type is string.
    PolicyDescr interface{}

    // Classifier entry configuration in a policy. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry.
    ClassifierEntry []Policies_PolicyEntry_ClassifierEntry
}

func (policyEntry *Policies_PolicyEntry) GetFilter() yfilter.YFilter { return policyEntry.YFilter }

func (policyEntry *Policies_PolicyEntry) SetFilter(yf yfilter.YFilter) { policyEntry.YFilter = yf }

func (policyEntry *Policies_PolicyEntry) GetGoName(yname string) string {
    if yname == "policy-name" { return "PolicyName" }
    if yname == "policy-descr" { return "PolicyDescr" }
    if yname == "classifier-entry" { return "ClassifierEntry" }
    return ""
}

func (policyEntry *Policies_PolicyEntry) GetSegmentPath() string {
    return "policy-entry" + "[policy-name='" + fmt.Sprintf("%v", policyEntry.PolicyName) + "']"
}

func (policyEntry *Policies_PolicyEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "classifier-entry" {
        for _, c := range policyEntry.ClassifierEntry {
            if policyEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Policies_PolicyEntry_ClassifierEntry{}
        policyEntry.ClassifierEntry = append(policyEntry.ClassifierEntry, child)
        return &policyEntry.ClassifierEntry[len(policyEntry.ClassifierEntry)-1]
    }
    return nil
}

func (policyEntry *Policies_PolicyEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policyEntry.ClassifierEntry {
        children[policyEntry.ClassifierEntry[i].GetSegmentPath()] = &policyEntry.ClassifierEntry[i]
    }
    return children
}

func (policyEntry *Policies_PolicyEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-name"] = policyEntry.PolicyName
    leafs["policy-descr"] = policyEntry.PolicyDescr
    return leafs
}

func (policyEntry *Policies_PolicyEntry) GetBundleName() string { return "ietf" }

func (policyEntry *Policies_PolicyEntry) GetYangName() string { return "policy-entry" }

func (policyEntry *Policies_PolicyEntry) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (policyEntry *Policies_PolicyEntry) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (policyEntry *Policies_PolicyEntry) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (policyEntry *Policies_PolicyEntry) SetParent(parent types.Entity) { policyEntry.parent = parent }

func (policyEntry *Policies_PolicyEntry) GetParent() types.Entity { return policyEntry.parent }

func (policyEntry *Policies_PolicyEntry) GetParentYangName() string { return "policies" }

// Policies_PolicyEntry_ClassifierEntry
// Classifier entry configuration in a policy
type Policies_PolicyEntry_ClassifierEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Diffserv classifier entry name. The type is
    // string.
    ClassifierEntryName interface{}

    // Indication of inline classifier entry. The type is bool. The default value
    // is false.
    ClassifierEntryInline interface{}

    // Filters are applicable as any or all filters. The type is one of the
    // following: MatchAllFilterMatchAnyFilter. The default value is
    // match-any-filter.
    ClassifierEntryFilterOper interface{}

    // Filters configured inline in a policy. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_FilterEntry.
    FilterEntry []Policies_PolicyEntry_ClassifierEntry_FilterEntry

    // Configuration of classifier & associated actions. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg.
    ClassifierActionEntryCfg []Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg
}

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetFilter() yfilter.YFilter { return classifierEntry.YFilter }

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) SetFilter(yf yfilter.YFilter) { classifierEntry.YFilter = yf }

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetGoName(yname string) string {
    if yname == "classifier-entry-name" { return "ClassifierEntryName" }
    if yname == "classifier-entry-inline" { return "ClassifierEntryInline" }
    if yname == "classifier-entry-filter-oper" { return "ClassifierEntryFilterOper" }
    if yname == "filter-entry" { return "FilterEntry" }
    if yname == "classifier-action-entry-cfg" { return "ClassifierActionEntryCfg" }
    return ""
}

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetSegmentPath() string {
    return "classifier-entry" + "[classifier-entry-name='" + fmt.Sprintf("%v", classifierEntry.ClassifierEntryName) + "']"
}

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "filter-entry" {
        for _, c := range classifierEntry.FilterEntry {
            if classifierEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Policies_PolicyEntry_ClassifierEntry_FilterEntry{}
        classifierEntry.FilterEntry = append(classifierEntry.FilterEntry, child)
        return &classifierEntry.FilterEntry[len(classifierEntry.FilterEntry)-1]
    }
    if childYangName == "classifier-action-entry-cfg" {
        for _, c := range classifierEntry.ClassifierActionEntryCfg {
            if classifierEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg{}
        classifierEntry.ClassifierActionEntryCfg = append(classifierEntry.ClassifierActionEntryCfg, child)
        return &classifierEntry.ClassifierActionEntryCfg[len(classifierEntry.ClassifierActionEntryCfg)-1]
    }
    return nil
}

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classifierEntry.FilterEntry {
        children[classifierEntry.FilterEntry[i].GetSegmentPath()] = &classifierEntry.FilterEntry[i]
    }
    for i := range classifierEntry.ClassifierActionEntryCfg {
        children[classifierEntry.ClassifierActionEntryCfg[i].GetSegmentPath()] = &classifierEntry.ClassifierActionEntryCfg[i]
    }
    return children
}

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["classifier-entry-name"] = classifierEntry.ClassifierEntryName
    leafs["classifier-entry-inline"] = classifierEntry.ClassifierEntryInline
    leafs["classifier-entry-filter-oper"] = classifierEntry.ClassifierEntryFilterOper
    return leafs
}

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetBundleName() string { return "ietf" }

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetYangName() string { return "classifier-entry" }

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) SetParent(parent types.Entity) { classifierEntry.parent = parent }

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetParent() types.Entity { return classifierEntry.parent }

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetParentYangName() string { return "policy-entry" }

// Policies_PolicyEntry_ClassifierEntry_FilterEntry
// Filters configured inline in a policy
type Policies_PolicyEntry_ClassifierEntry_FilterEntry struct {
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
    // Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg.
    DscpCfg []Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg

    // list of source ip address. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg.
    SourceIpAddressCfg []Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg

    // list of destination ip address. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg.
    DestinationIpAddressCfg []Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg

    // list of ranges of source port. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg.
    SourcePortCfg []Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg

    // list of ranges of destination port. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg.
    DestinationPortCfg []Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg

    // list of ranges of protocol values. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg.
    ProtocolCfg []Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg
}

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetFilter() yfilter.YFilter { return filterEntry.YFilter }

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) SetFilter(yf yfilter.YFilter) { filterEntry.YFilter = yf }

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetGoName(yname string) string {
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

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetSegmentPath() string {
    return "filter-entry" + "[filter-type='" + fmt.Sprintf("%v", filterEntry.FilterType) + "']" + "[filter-logical-not='" + fmt.Sprintf("%v", filterEntry.FilterLogicalNot) + "']"
}

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dscp-cfg" {
        for _, c := range filterEntry.DscpCfg {
            if filterEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg{}
        filterEntry.DscpCfg = append(filterEntry.DscpCfg, child)
        return &filterEntry.DscpCfg[len(filterEntry.DscpCfg)-1]
    }
    if childYangName == "source-ip-address-cfg" {
        for _, c := range filterEntry.SourceIpAddressCfg {
            if filterEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg{}
        filterEntry.SourceIpAddressCfg = append(filterEntry.SourceIpAddressCfg, child)
        return &filterEntry.SourceIpAddressCfg[len(filterEntry.SourceIpAddressCfg)-1]
    }
    if childYangName == "destination-ip-address-cfg" {
        for _, c := range filterEntry.DestinationIpAddressCfg {
            if filterEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg{}
        filterEntry.DestinationIpAddressCfg = append(filterEntry.DestinationIpAddressCfg, child)
        return &filterEntry.DestinationIpAddressCfg[len(filterEntry.DestinationIpAddressCfg)-1]
    }
    if childYangName == "source-port-cfg" {
        for _, c := range filterEntry.SourcePortCfg {
            if filterEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg{}
        filterEntry.SourcePortCfg = append(filterEntry.SourcePortCfg, child)
        return &filterEntry.SourcePortCfg[len(filterEntry.SourcePortCfg)-1]
    }
    if childYangName == "destination-port-cfg" {
        for _, c := range filterEntry.DestinationPortCfg {
            if filterEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg{}
        filterEntry.DestinationPortCfg = append(filterEntry.DestinationPortCfg, child)
        return &filterEntry.DestinationPortCfg[len(filterEntry.DestinationPortCfg)-1]
    }
    if childYangName == "protocol-cfg" {
        for _, c := range filterEntry.ProtocolCfg {
            if filterEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg{}
        filterEntry.ProtocolCfg = append(filterEntry.ProtocolCfg, child)
        return &filterEntry.ProtocolCfg[len(filterEntry.ProtocolCfg)-1]
    }
    return nil
}

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetChildren() map[string]types.Entity {
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

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["filter-type"] = filterEntry.FilterType
    leafs["filter-logical-not"] = filterEntry.FilterLogicalNot
    return leafs
}

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetBundleName() string { return "ietf" }

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetYangName() string { return "filter-entry" }

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) SetParent(parent types.Entity) { filterEntry.parent = parent }

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetParent() types.Entity { return filterEntry.parent }

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetParentYangName() string { return "classifier-entry" }

// Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg
// list of dscp ranges
type Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Minimum value of dscp range. The type is
    // interface{} with range: 0..63.
    DscpMin interface{}

    // This attribute is a key. maximum value of dscp range. The type is
    // interface{} with range: 0..63.
    DscpMax interface{}
}

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetFilter() yfilter.YFilter { return dscpCfg.YFilter }

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) SetFilter(yf yfilter.YFilter) { dscpCfg.YFilter = yf }

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetGoName(yname string) string {
    if yname == "dscp-min" { return "DscpMin" }
    if yname == "dscp-max" { return "DscpMax" }
    return ""
}

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetSegmentPath() string {
    return "dscp-cfg" + "[dscp-min='" + fmt.Sprintf("%v", dscpCfg.DscpMin) + "']" + "[dscp-max='" + fmt.Sprintf("%v", dscpCfg.DscpMax) + "']"
}

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp-min"] = dscpCfg.DscpMin
    leafs["dscp-max"] = dscpCfg.DscpMax
    return leafs
}

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetBundleName() string { return "ietf" }

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetYangName() string { return "dscp-cfg" }

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) SetParent(parent types.Entity) { dscpCfg.parent = parent }

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetParent() types.Entity { return dscpCfg.parent }

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetParentYangName() string { return "filter-entry" }

// Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg
// list of source ip address
type Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. source ip prefix. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    SourceIpAddr interface{}
}

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetFilter() yfilter.YFilter { return sourceIpAddressCfg.YFilter }

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) SetFilter(yf yfilter.YFilter) { sourceIpAddressCfg.YFilter = yf }

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetGoName(yname string) string {
    if yname == "source-ip-addr" { return "SourceIpAddr" }
    return ""
}

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetSegmentPath() string {
    return "source-ip-address-cfg" + "[source-ip-addr='" + fmt.Sprintf("%v", sourceIpAddressCfg.SourceIpAddr) + "']"
}

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-ip-addr"] = sourceIpAddressCfg.SourceIpAddr
    return leafs
}

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetBundleName() string { return "ietf" }

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetYangName() string { return "source-ip-address-cfg" }

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) SetParent(parent types.Entity) { sourceIpAddressCfg.parent = parent }

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetParent() types.Entity { return sourceIpAddressCfg.parent }

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetParentYangName() string { return "filter-entry" }

// Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg
// list of destination ip address
type Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. destination ip prefix. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    DestinationIpAddr interface{}
}

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetFilter() yfilter.YFilter { return destinationIpAddressCfg.YFilter }

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) SetFilter(yf yfilter.YFilter) { destinationIpAddressCfg.YFilter = yf }

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetGoName(yname string) string {
    if yname == "destination-ip-addr" { return "DestinationIpAddr" }
    return ""
}

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetSegmentPath() string {
    return "destination-ip-address-cfg" + "[destination-ip-addr='" + fmt.Sprintf("%v", destinationIpAddressCfg.DestinationIpAddr) + "']"
}

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-ip-addr"] = destinationIpAddressCfg.DestinationIpAddr
    return leafs
}

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetBundleName() string { return "ietf" }

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetYangName() string { return "destination-ip-address-cfg" }

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) SetParent(parent types.Entity) { destinationIpAddressCfg.parent = parent }

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetParent() types.Entity { return destinationIpAddressCfg.parent }

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetParentYangName() string { return "filter-entry" }

// Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg
// list of ranges of source port
type Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. minimum value of source port range. The type is
    // interface{} with range: 0..65535.
    SourcePortMin interface{}

    // This attribute is a key. maximum value of source port range. The type is
    // interface{} with range: 0..65535.
    SourcePortMax interface{}
}

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetFilter() yfilter.YFilter { return sourcePortCfg.YFilter }

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) SetFilter(yf yfilter.YFilter) { sourcePortCfg.YFilter = yf }

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetGoName(yname string) string {
    if yname == "source-port-min" { return "SourcePortMin" }
    if yname == "source-port-max" { return "SourcePortMax" }
    return ""
}

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetSegmentPath() string {
    return "source-port-cfg" + "[source-port-min='" + fmt.Sprintf("%v", sourcePortCfg.SourcePortMin) + "']" + "[source-port-max='" + fmt.Sprintf("%v", sourcePortCfg.SourcePortMax) + "']"
}

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-port-min"] = sourcePortCfg.SourcePortMin
    leafs["source-port-max"] = sourcePortCfg.SourcePortMax
    return leafs
}

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetBundleName() string { return "ietf" }

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetYangName() string { return "source-port-cfg" }

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) SetParent(parent types.Entity) { sourcePortCfg.parent = parent }

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetParent() types.Entity { return sourcePortCfg.parent }

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetParentYangName() string { return "filter-entry" }

// Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg
// list of ranges of destination port
type Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. minimum value of destination port range. The type
    // is interface{} with range: 0..65535.
    DestinationPortMin interface{}

    // This attribute is a key. maximum value of destination port range. The type
    // is interface{} with range: 0..65535.
    DestinationPortMax interface{}
}

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetFilter() yfilter.YFilter { return destinationPortCfg.YFilter }

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) SetFilter(yf yfilter.YFilter) { destinationPortCfg.YFilter = yf }

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetGoName(yname string) string {
    if yname == "destination-port-min" { return "DestinationPortMin" }
    if yname == "destination-port-max" { return "DestinationPortMax" }
    return ""
}

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetSegmentPath() string {
    return "destination-port-cfg" + "[destination-port-min='" + fmt.Sprintf("%v", destinationPortCfg.DestinationPortMin) + "']" + "[destination-port-max='" + fmt.Sprintf("%v", destinationPortCfg.DestinationPortMax) + "']"
}

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-port-min"] = destinationPortCfg.DestinationPortMin
    leafs["destination-port-max"] = destinationPortCfg.DestinationPortMax
    return leafs
}

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetBundleName() string { return "ietf" }

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetYangName() string { return "destination-port-cfg" }

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) SetParent(parent types.Entity) { destinationPortCfg.parent = parent }

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetParent() types.Entity { return destinationPortCfg.parent }

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetParentYangName() string { return "filter-entry" }

// Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg
// list of ranges of protocol values
type Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. minimum value of protocol range. The type is
    // interface{} with range: 0..255.
    ProtocolMin interface{}

    // This attribute is a key. maximum value of protocol range. The type is
    // interface{} with range: 0..255.
    ProtocolMax interface{}
}

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetFilter() yfilter.YFilter { return protocolCfg.YFilter }

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) SetFilter(yf yfilter.YFilter) { protocolCfg.YFilter = yf }

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetGoName(yname string) string {
    if yname == "protocol-min" { return "ProtocolMin" }
    if yname == "protocol-max" { return "ProtocolMax" }
    return ""
}

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetSegmentPath() string {
    return "protocol-cfg" + "[protocol-min='" + fmt.Sprintf("%v", protocolCfg.ProtocolMin) + "']" + "[protocol-max='" + fmt.Sprintf("%v", protocolCfg.ProtocolMax) + "']"
}

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol-min"] = protocolCfg.ProtocolMin
    leafs["protocol-max"] = protocolCfg.ProtocolMax
    return leafs
}

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetBundleName() string { return "ietf" }

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetYangName() string { return "protocol-cfg" }

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) SetParent(parent types.Entity) { protocolCfg.parent = parent }

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetParent() types.Entity { return protocolCfg.parent }

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetParentYangName() string { return "filter-entry" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg
// Configuration of classifier & associated actions
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This defines action type . The type is one of the
    // following: MarkingMinRateMeterPriorityMaxRateAlgorithmicDrop.
    ActionType interface{}

    // Marking configuration container.
    MarkingCfg Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg

    // priority attributes container.
    PriorityCfg Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg

    // Meter list configuration container.
    MeterCfg Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg

    // min guaranteed bandwidth.
    MinRateCfg Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg

    // maximum rate attributes.
    MaxRateCfg Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg

    // Always Drop configuration container.
    DropCfg Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg

    // Tail Drop configuration container.
    TailDropCfg Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg

    // Random Detect configuration container.
    RandomDetectCfg Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg
}

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetFilter() yfilter.YFilter { return classifierActionEntryCfg.YFilter }

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) SetFilter(yf yfilter.YFilter) { classifierActionEntryCfg.YFilter = yf }

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "ietf-diffserv-action:marking-cfg" { return "MarkingCfg" }
    if yname == "ietf-diffserv-action:priority-cfg" { return "PriorityCfg" }
    if yname == "ietf-diffserv-action:meter-cfg" { return "MeterCfg" }
    if yname == "ietf-diffserv-action:min-rate-cfg" { return "MinRateCfg" }
    if yname == "ietf-diffserv-action:max-rate-cfg" { return "MaxRateCfg" }
    if yname == "ietf-diffserv-action:drop-cfg" { return "DropCfg" }
    if yname == "ietf-diffserv-action:tail-drop-cfg" { return "TailDropCfg" }
    if yname == "ietf-diffserv-action:random-detect-cfg" { return "RandomDetectCfg" }
    return ""
}

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetSegmentPath() string {
    return "classifier-action-entry-cfg" + "[action-type='" + fmt.Sprintf("%v", classifierActionEntryCfg.ActionType) + "']"
}

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ietf-diffserv-action:marking-cfg" {
        return &classifierActionEntryCfg.MarkingCfg
    }
    if childYangName == "ietf-diffserv-action:priority-cfg" {
        return &classifierActionEntryCfg.PriorityCfg
    }
    if childYangName == "ietf-diffserv-action:meter-cfg" {
        return &classifierActionEntryCfg.MeterCfg
    }
    if childYangName == "ietf-diffserv-action:min-rate-cfg" {
        return &classifierActionEntryCfg.MinRateCfg
    }
    if childYangName == "ietf-diffserv-action:max-rate-cfg" {
        return &classifierActionEntryCfg.MaxRateCfg
    }
    if childYangName == "ietf-diffserv-action:drop-cfg" {
        return &classifierActionEntryCfg.DropCfg
    }
    if childYangName == "ietf-diffserv-action:tail-drop-cfg" {
        return &classifierActionEntryCfg.TailDropCfg
    }
    if childYangName == "ietf-diffserv-action:random-detect-cfg" {
        return &classifierActionEntryCfg.RandomDetectCfg
    }
    return nil
}

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ietf-diffserv-action:marking-cfg"] = &classifierActionEntryCfg.MarkingCfg
    children["ietf-diffserv-action:priority-cfg"] = &classifierActionEntryCfg.PriorityCfg
    children["ietf-diffserv-action:meter-cfg"] = &classifierActionEntryCfg.MeterCfg
    children["ietf-diffserv-action:min-rate-cfg"] = &classifierActionEntryCfg.MinRateCfg
    children["ietf-diffserv-action:max-rate-cfg"] = &classifierActionEntryCfg.MaxRateCfg
    children["ietf-diffserv-action:drop-cfg"] = &classifierActionEntryCfg.DropCfg
    children["ietf-diffserv-action:tail-drop-cfg"] = &classifierActionEntryCfg.TailDropCfg
    children["ietf-diffserv-action:random-detect-cfg"] = &classifierActionEntryCfg.RandomDetectCfg
    return children
}

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = classifierActionEntryCfg.ActionType
    return leafs
}

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetBundleName() string { return "ietf" }

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetYangName() string { return "classifier-action-entry-cfg" }

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) SetParent(parent types.Entity) { classifierActionEntryCfg.parent = parent }

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetParent() types.Entity { return classifierActionEntryCfg.parent }

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetParentYangName() string { return "classifier-entry" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg
// Marking configuration container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // dscp marking. The type is interface{} with range: 0..63.
    Dscp interface{}
}

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetFilter() yfilter.YFilter { return markingCfg.YFilter }

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) SetFilter(yf yfilter.YFilter) { markingCfg.YFilter = yf }

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    return ""
}

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetSegmentPath() string {
    return "ietf-diffserv-action:marking-cfg"
}

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp"] = markingCfg.Dscp
    return leafs
}

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetBundleName() string { return "ietf" }

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetYangName() string { return "marking-cfg" }

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) SetParent(parent types.Entity) { markingCfg.parent = parent }

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetParent() types.Entity { return markingCfg.parent }

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetParentYangName() string { return "classifier-action-entry-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg
// priority attributes container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // priority level. The type is interface{} with range: 0..255.
    PriorityLevel interface{}

    // absolute priority rate with/without burst rateand absolute percent.
    RateBurst Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst
}

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetFilter() yfilter.YFilter { return priorityCfg.YFilter }

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) SetFilter(yf yfilter.YFilter) { priorityCfg.YFilter = yf }

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetGoName(yname string) string {
    if yname == "priority-level" { return "PriorityLevel" }
    if yname == "rate-burst" { return "RateBurst" }
    return ""
}

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetSegmentPath() string {
    return "ietf-diffserv-action:priority-cfg"
}

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rate-burst" {
        return &priorityCfg.RateBurst
    }
    return nil
}

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rate-burst"] = &priorityCfg.RateBurst
    return children
}

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["priority-level"] = priorityCfg.PriorityLevel
    return leafs
}

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetBundleName() string { return "ietf" }

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetYangName() string { return "priority-cfg" }

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) SetParent(parent types.Entity) { priorityCfg.parent = parent }

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetParent() types.Entity { return priorityCfg.parent }

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetParentYangName() string { return "classifier-action-entry-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst
// absolute priority rate with/without burst rateand absolute percent
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rate value. The type is interface{} with range: 0..18446744073709551615.
    // Units are bits-per-second.
    Rate interface{}

    // Metric. The type is Metric. The default value is none.
    AbsoluteRateMetric interface{}

    // Rate basic units. The type is RateUnit.
    AbsoluteRateUnits interface{}

    // percent. The type is interface{} with range: 1..100.
    RatePercent interface{}

    // The type is interface{} with range: 1..65532.
    RateRatio interface{}

    // burst size. The type is interface{} with range: 0..18446744073709551615.
    // Units are bytes.
    BurstSize interface{}

    // burst interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are microsecond.
    BurstInterval interface{}
}

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetFilter() yfilter.YFilter { return rateBurst.YFilter }

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) SetFilter(yf yfilter.YFilter) { rateBurst.YFilter = yf }

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetGoName(yname string) string {
    if yname == "rate" { return "Rate" }
    if yname == "absolute-rate-metric" { return "AbsoluteRateMetric" }
    if yname == "absolute-rate-units" { return "AbsoluteRateUnits" }
    if yname == "rate-percent" { return "RatePercent" }
    if yname == "rate-ratio" { return "RateRatio" }
    if yname == "burst-size" { return "BurstSize" }
    if yname == "burst-interval" { return "BurstInterval" }
    return ""
}

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetSegmentPath() string {
    return "rate-burst"
}

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rate"] = rateBurst.Rate
    leafs["absolute-rate-metric"] = rateBurst.AbsoluteRateMetric
    leafs["absolute-rate-units"] = rateBurst.AbsoluteRateUnits
    leafs["rate-percent"] = rateBurst.RatePercent
    leafs["rate-ratio"] = rateBurst.RateRatio
    leafs["burst-size"] = rateBurst.BurstSize
    leafs["burst-interval"] = rateBurst.BurstInterval
    return leafs
}

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetBundleName() string { return "ietf" }

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetYangName() string { return "rate-burst" }

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) SetParent(parent types.Entity) { rateBurst.parent = parent }

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetParent() types.Entity { return rateBurst.parent }

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetParentYangName() string { return "priority-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg
// Meter list configuration container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Meter configuration. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList.
    MeterList []Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList
}

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetFilter() yfilter.YFilter { return meterCfg.YFilter }

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) SetFilter(yf yfilter.YFilter) { meterCfg.YFilter = yf }

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetGoName(yname string) string {
    if yname == "meter-list" { return "MeterList" }
    return ""
}

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetSegmentPath() string {
    return "ietf-diffserv-action:meter-cfg"
}

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "meter-list" {
        for _, c := range meterCfg.MeterList {
            if meterCfg.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList{}
        meterCfg.MeterList = append(meterCfg.MeterList, child)
        return &meterCfg.MeterList[len(meterCfg.MeterList)-1]
    }
    return nil
}

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range meterCfg.MeterList {
        children[meterCfg.MeterList[i].GetSegmentPath()] = &meterCfg.MeterList[i]
    }
    return children
}

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetBundleName() string { return "ietf" }

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetYangName() string { return "meter-cfg" }

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) SetParent(parent types.Entity) { meterCfg.parent = parent }

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetParent() types.Entity { return meterCfg.parent }

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetParentYangName() string { return "classifier-action-entry-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList
// Meter configuration
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. meter identifier. The type is interface{} with
    // range: 0..65535.
    MeterId interface{}

    // meter rate. The type is interface{} with range: 0..18446744073709551615.
    // Units are bits-per-second.
    MeterRate interface{}

    // burst size. The type is interface{} with range: 0..18446744073709551615.
    // Units are bytes.
    BurstSize interface{}

    // burst interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are microsecond.
    BurstInterval interface{}

    // color aware & color blind attributes container.
    Color Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color

    // confirm action.
    SucceedAction Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction

    // exceed action.
    FailAction Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction
}

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetFilter() yfilter.YFilter { return meterList.YFilter }

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) SetFilter(yf yfilter.YFilter) { meterList.YFilter = yf }

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetGoName(yname string) string {
    if yname == "meter-id" { return "MeterId" }
    if yname == "meter-rate" { return "MeterRate" }
    if yname == "burst-size" { return "BurstSize" }
    if yname == "burst-interval" { return "BurstInterval" }
    if yname == "color" { return "Color" }
    if yname == "succeed-action" { return "SucceedAction" }
    if yname == "fail-action" { return "FailAction" }
    return ""
}

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetSegmentPath() string {
    return "meter-list" + "[meter-id='" + fmt.Sprintf("%v", meterList.MeterId) + "']"
}

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "color" {
        return &meterList.Color
    }
    if childYangName == "succeed-action" {
        return &meterList.SucceedAction
    }
    if childYangName == "fail-action" {
        return &meterList.FailAction
    }
    return nil
}

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["color"] = &meterList.Color
    children["succeed-action"] = &meterList.SucceedAction
    children["fail-action"] = &meterList.FailAction
    return children
}

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["meter-id"] = meterList.MeterId
    leafs["meter-rate"] = meterList.MeterRate
    leafs["burst-size"] = meterList.BurstSize
    leafs["burst-interval"] = meterList.BurstInterval
    return leafs
}

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetBundleName() string { return "ietf" }

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetYangName() string { return "meter-list" }

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) SetParent(parent types.Entity) { meterList.parent = parent }

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetParent() types.Entity { return meterList.parent }

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetParentYangName() string { return "meter-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color
// color aware & color blind attributes container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Diffserv classifier name. The type is string.
    ClassifierEntryName interface{}

    // Description of the class template. The type is string.
    ClassifierEntryDescr interface{}

    // Filters are applicable as any or all filters. The type is one of the
    // following: MatchAllFilterMatchAnyFilter. The default value is
    // match-any-filter.
    ClassifierEntryFilterOperation interface{}
}

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetFilter() yfilter.YFilter { return color.YFilter }

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) SetFilter(yf yfilter.YFilter) { color.YFilter = yf }

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetGoName(yname string) string {
    if yname == "classifier-entry-name" { return "ClassifierEntryName" }
    if yname == "classifier-entry-descr" { return "ClassifierEntryDescr" }
    if yname == "classifier-entry-filter-operation" { return "ClassifierEntryFilterOperation" }
    return ""
}

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetSegmentPath() string {
    return "color"
}

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["classifier-entry-name"] = color.ClassifierEntryName
    leafs["classifier-entry-descr"] = color.ClassifierEntryDescr
    leafs["classifier-entry-filter-operation"] = color.ClassifierEntryFilterOperation
    return leafs
}

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetBundleName() string { return "ietf" }

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetYangName() string { return "color" }

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) SetParent(parent types.Entity) { color.parent = parent }

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetParent() types.Entity { return color.parent }

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetParentYangName() string { return "meter-list" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction
// confirm action
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // meter action type. The type is one of the following:
    // MeterActionDropMeterActionSet.
    MeterActionType interface{}

    // next meter identifier. The type is interface{} with range: 0..65535.
    NextMeterId interface{}

    // dscp marking. The type is interface{} with range: 0..63.
    Dscp interface{}

    // always drop algorithm. The type is interface{}.
    DropAction interface{}
}

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetFilter() yfilter.YFilter { return succeedAction.YFilter }

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) SetFilter(yf yfilter.YFilter) { succeedAction.YFilter = yf }

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetGoName(yname string) string {
    if yname == "meter-action-type" { return "MeterActionType" }
    if yname == "next-meter-id" { return "NextMeterId" }
    if yname == "dscp" { return "Dscp" }
    if yname == "drop-action" { return "DropAction" }
    return ""
}

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetSegmentPath() string {
    return "succeed-action"
}

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["meter-action-type"] = succeedAction.MeterActionType
    leafs["next-meter-id"] = succeedAction.NextMeterId
    leafs["dscp"] = succeedAction.Dscp
    leafs["drop-action"] = succeedAction.DropAction
    return leafs
}

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetBundleName() string { return "ietf" }

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetYangName() string { return "succeed-action" }

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) SetParent(parent types.Entity) { succeedAction.parent = parent }

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetParent() types.Entity { return succeedAction.parent }

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetParentYangName() string { return "meter-list" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction
// exceed action
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // meter action type. The type is one of the following:
    // MeterActionDropMeterActionSet.
    MeterActionType interface{}

    // next meter identifier. The type is interface{} with range: 0..65535.
    NextMeterId interface{}

    // dscp marking. The type is interface{} with range: 0..63.
    Dscp interface{}

    // always drop algorithm. The type is interface{}.
    DropAction interface{}
}

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetFilter() yfilter.YFilter { return failAction.YFilter }

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) SetFilter(yf yfilter.YFilter) { failAction.YFilter = yf }

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetGoName(yname string) string {
    if yname == "meter-action-type" { return "MeterActionType" }
    if yname == "next-meter-id" { return "NextMeterId" }
    if yname == "dscp" { return "Dscp" }
    if yname == "drop-action" { return "DropAction" }
    return ""
}

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetSegmentPath() string {
    return "fail-action"
}

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["meter-action-type"] = failAction.MeterActionType
    leafs["next-meter-id"] = failAction.NextMeterId
    leafs["dscp"] = failAction.Dscp
    leafs["drop-action"] = failAction.DropAction
    return leafs
}

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetBundleName() string { return "ietf" }

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetYangName() string { return "fail-action" }

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) SetParent(parent types.Entity) { failAction.parent = parent }

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetParent() types.Entity { return failAction.parent }

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetParentYangName() string { return "meter-list" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg
// min guaranteed bandwidth
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // minimum rate. The type is interface{} with range: 0..18446744073709551615.
    // Units are bits-per-second.
    MinRate interface{}

    // Metric. The type is Metric. The default value is none.
    AbsoluteRateMetric interface{}

    // Rate basic units. The type is RateUnit.
    AbsoluteRateUnits interface{}

    // percent. The type is interface{} with range: 1..100.
    RatePercent interface{}

    // The type is interface{} with range: 1..65532.
    RateRatio interface{}

    // share the bandwidth remaming.
    BwExcessShareCfg Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg
}

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetFilter() yfilter.YFilter { return minRateCfg.YFilter }

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) SetFilter(yf yfilter.YFilter) { minRateCfg.YFilter = yf }

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetGoName(yname string) string {
    if yname == "min-rate" { return "MinRate" }
    if yname == "absolute-rate-metric" { return "AbsoluteRateMetric" }
    if yname == "absolute-rate-units" { return "AbsoluteRateUnits" }
    if yname == "rate-percent" { return "RatePercent" }
    if yname == "rate-ratio" { return "RateRatio" }
    if yname == "bw-excess-share-cfg" { return "BwExcessShareCfg" }
    return ""
}

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetSegmentPath() string {
    return "ietf-diffserv-action:min-rate-cfg"
}

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bw-excess-share-cfg" {
        return &minRateCfg.BwExcessShareCfg
    }
    return nil
}

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bw-excess-share-cfg"] = &minRateCfg.BwExcessShareCfg
    return children
}

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min-rate"] = minRateCfg.MinRate
    leafs["absolute-rate-metric"] = minRateCfg.AbsoluteRateMetric
    leafs["absolute-rate-units"] = minRateCfg.AbsoluteRateUnits
    leafs["rate-percent"] = minRateCfg.RatePercent
    leafs["rate-ratio"] = minRateCfg.RateRatio
    return leafs
}

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetBundleName() string { return "ietf" }

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetYangName() string { return "min-rate-cfg" }

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) SetParent(parent types.Entity) { minRateCfg.parent = parent }

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetParent() types.Entity { return minRateCfg.parent }

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetParentYangName() string { return "classifier-action-entry-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg
// share the bandwidth remaming
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // percentage or ratio value. The type is interface{} with range:
    // 0..4294967295.
    Value interface{}

    // Metric. The type is Metric. The default value is none.
    AbsoluteRateMetric interface{}

    // Rate basic units. The type is RateUnit.
    AbsoluteRateUnits interface{}

    // percent. The type is interface{} with range: 1..100.
    RatePercent interface{}

    // The type is interface{} with range: 1..65532.
    RateRatio interface{}
}

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetFilter() yfilter.YFilter { return bwExcessShareCfg.YFilter }

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) SetFilter(yf yfilter.YFilter) { bwExcessShareCfg.YFilter = yf }

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "absolute-rate-metric" { return "AbsoluteRateMetric" }
    if yname == "absolute-rate-units" { return "AbsoluteRateUnits" }
    if yname == "rate-percent" { return "RatePercent" }
    if yname == "rate-ratio" { return "RateRatio" }
    return ""
}

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetSegmentPath() string {
    return "bw-excess-share-cfg"
}

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = bwExcessShareCfg.Value
    leafs["absolute-rate-metric"] = bwExcessShareCfg.AbsoluteRateMetric
    leafs["absolute-rate-units"] = bwExcessShareCfg.AbsoluteRateUnits
    leafs["rate-percent"] = bwExcessShareCfg.RatePercent
    leafs["rate-ratio"] = bwExcessShareCfg.RateRatio
    return leafs
}

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetBundleName() string { return "ietf" }

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetYangName() string { return "bw-excess-share-cfg" }

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) SetParent(parent types.Entity) { bwExcessShareCfg.parent = parent }

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetParent() types.Entity { return bwExcessShareCfg.parent }

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetParentYangName() string { return "min-rate-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg
// maximum rate attributes
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // rate in bits per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are bits-per-second.
    AbsoluteRate interface{}

    // burst size. The type is interface{} with range: 0..18446744073709551615.
    // Units are bytes.
    BurstSize interface{}

    // burst interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are microsecond.
    BurstInterval interface{}

    // Metric. The type is Metric. The default value is none.
    AbsoluteRateMetric interface{}

    // Rate basic units. The type is RateUnit.
    AbsoluteRateUnits interface{}

    // percent. The type is interface{} with range: 1..100.
    RatePercent interface{}

    // The type is interface{} with range: 1..65532.
    RateRatio interface{}
}

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetFilter() yfilter.YFilter { return maxRateCfg.YFilter }

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) SetFilter(yf yfilter.YFilter) { maxRateCfg.YFilter = yf }

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetGoName(yname string) string {
    if yname == "absolute-rate" { return "AbsoluteRate" }
    if yname == "burst-size" { return "BurstSize" }
    if yname == "burst-interval" { return "BurstInterval" }
    if yname == "absolute-rate-metric" { return "AbsoluteRateMetric" }
    if yname == "absolute-rate-units" { return "AbsoluteRateUnits" }
    if yname == "rate-percent" { return "RatePercent" }
    if yname == "rate-ratio" { return "RateRatio" }
    return ""
}

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetSegmentPath() string {
    return "ietf-diffserv-action:max-rate-cfg"
}

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["absolute-rate"] = maxRateCfg.AbsoluteRate
    leafs["burst-size"] = maxRateCfg.BurstSize
    leafs["burst-interval"] = maxRateCfg.BurstInterval
    leafs["absolute-rate-metric"] = maxRateCfg.AbsoluteRateMetric
    leafs["absolute-rate-units"] = maxRateCfg.AbsoluteRateUnits
    leafs["rate-percent"] = maxRateCfg.RatePercent
    leafs["rate-ratio"] = maxRateCfg.RateRatio
    return leafs
}

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetBundleName() string { return "ietf" }

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetYangName() string { return "max-rate-cfg" }

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) SetParent(parent types.Entity) { maxRateCfg.parent = parent }

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetParent() types.Entity { return maxRateCfg.parent }

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetParentYangName() string { return "classifier-action-entry-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg
// Always Drop configuration container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // always drop algorithm. The type is interface{}.
    DropAction interface{}
}

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetFilter() yfilter.YFilter { return dropCfg.YFilter }

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) SetFilter(yf yfilter.YFilter) { dropCfg.YFilter = yf }

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetGoName(yname string) string {
    if yname == "drop-action" { return "DropAction" }
    return ""
}

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetSegmentPath() string {
    return "ietf-diffserv-action:drop-cfg"
}

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["drop-action"] = dropCfg.DropAction
    return leafs
}

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetBundleName() string { return "ietf" }

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetYangName() string { return "drop-cfg" }

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) SetParent(parent types.Entity) { dropCfg.parent = parent }

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetParent() types.Entity { return dropCfg.parent }

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetParentYangName() string { return "classifier-action-entry-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg
// Tail Drop configuration container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // the queue limit per dscp range. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh.
    QlimitDscpThresh []Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh
}

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetFilter() yfilter.YFilter { return tailDropCfg.YFilter }

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) SetFilter(yf yfilter.YFilter) { tailDropCfg.YFilter = yf }

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetGoName(yname string) string {
    if yname == "qlimit-dscp-thresh" { return "QlimitDscpThresh" }
    return ""
}

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetSegmentPath() string {
    return "ietf-diffserv-action:tail-drop-cfg"
}

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qlimit-dscp-thresh" {
        for _, c := range tailDropCfg.QlimitDscpThresh {
            if tailDropCfg.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh{}
        tailDropCfg.QlimitDscpThresh = append(tailDropCfg.QlimitDscpThresh, child)
        return &tailDropCfg.QlimitDscpThresh[len(tailDropCfg.QlimitDscpThresh)-1]
    }
    return nil
}

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tailDropCfg.QlimitDscpThresh {
        children[tailDropCfg.QlimitDscpThresh[i].GetSegmentPath()] = &tailDropCfg.QlimitDscpThresh[i]
    }
    return children
}

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetBundleName() string { return "ietf" }

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetYangName() string { return "tail-drop-cfg" }

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) SetParent(parent types.Entity) { tailDropCfg.parent = parent }

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetParent() types.Entity { return tailDropCfg.parent }

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetParentYangName() string { return "classifier-action-entry-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh
// the queue limit per dscp range
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Minimum of dscp range. The type is interface{}
    // with range: 0..63.
    DscpMin interface{}

    // This attribute is a key. Maximum of dscp range. The type is interface{}
    // with range: 0..63.
    DscpMax interface{}

    // threshold.
    Threshold Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold
}

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetFilter() yfilter.YFilter { return qlimitDscpThresh.YFilter }

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) SetFilter(yf yfilter.YFilter) { qlimitDscpThresh.YFilter = yf }

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetGoName(yname string) string {
    if yname == "dscp-min" { return "DscpMin" }
    if yname == "dscp-max" { return "DscpMax" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetSegmentPath() string {
    return "qlimit-dscp-thresh" + "[dscp-min='" + fmt.Sprintf("%v", qlimitDscpThresh.DscpMin) + "']" + "[dscp-max='" + fmt.Sprintf("%v", qlimitDscpThresh.DscpMax) + "']"
}

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold" {
        return &qlimitDscpThresh.Threshold
    }
    return nil
}

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold"] = &qlimitDscpThresh.Threshold
    return children
}

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp-min"] = qlimitDscpThresh.DscpMin
    leafs["dscp-max"] = qlimitDscpThresh.DscpMax
    return leafs
}

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetBundleName() string { return "ietf" }

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetYangName() string { return "qlimit-dscp-thresh" }

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) SetParent(parent types.Entity) { qlimitDscpThresh.parent = parent }

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetParent() types.Entity { return qlimitDscpThresh.parent }

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetParentYangName() string { return "tail-drop-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold
// threshold
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold size. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    ThresholdSize interface{}

    // Threshold interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are microsecond.
    ThresholdInterval interface{}
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetFilter() yfilter.YFilter { return threshold.YFilter }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) SetFilter(yf yfilter.YFilter) { threshold.YFilter = yf }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetGoName(yname string) string {
    if yname == "threshold-size" { return "ThresholdSize" }
    if yname == "threshold-interval" { return "ThresholdInterval" }
    return ""
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetSegmentPath() string {
    return "threshold"
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-size"] = threshold.ThresholdSize
    leafs["threshold-interval"] = threshold.ThresholdInterval
    return leafs
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetBundleName() string { return "ietf" }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetYangName() string { return "threshold" }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) SetParent(parent types.Entity) { threshold.parent = parent }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetParent() types.Entity { return threshold.parent }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetParentYangName() string { return "qlimit-dscp-thresh" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg
// Random Detect configuration container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Exponential weighting constant factor for red profile . The type is
    // interface{} with range: 0..4294967295.
    ExpWeightingConst interface{}

    // Mark probability. The type is interface{} with range: 1..1000.
    MarkProbability interface{}

    // Minimum threshold.
    RedMinThresh Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh

    // Maximum threshold.
    RedMaxThresh Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh
}

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetFilter() yfilter.YFilter { return randomDetectCfg.YFilter }

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) SetFilter(yf yfilter.YFilter) { randomDetectCfg.YFilter = yf }

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetGoName(yname string) string {
    if yname == "exp-weighting-const" { return "ExpWeightingConst" }
    if yname == "mark-probability" { return "MarkProbability" }
    if yname == "red-min-thresh" { return "RedMinThresh" }
    if yname == "red-max-thresh" { return "RedMaxThresh" }
    return ""
}

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetSegmentPath() string {
    return "ietf-diffserv-action:random-detect-cfg"
}

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "red-min-thresh" {
        return &randomDetectCfg.RedMinThresh
    }
    if childYangName == "red-max-thresh" {
        return &randomDetectCfg.RedMaxThresh
    }
    return nil
}

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["red-min-thresh"] = &randomDetectCfg.RedMinThresh
    children["red-max-thresh"] = &randomDetectCfg.RedMaxThresh
    return children
}

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exp-weighting-const"] = randomDetectCfg.ExpWeightingConst
    leafs["mark-probability"] = randomDetectCfg.MarkProbability
    return leafs
}

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetBundleName() string { return "ietf" }

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetYangName() string { return "random-detect-cfg" }

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) SetParent(parent types.Entity) { randomDetectCfg.parent = parent }

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetParent() types.Entity { return randomDetectCfg.parent }

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetParentYangName() string { return "classifier-action-entry-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh
// Minimum threshold
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // threshold.
    Threshold Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold
}

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetFilter() yfilter.YFilter { return redMinThresh.YFilter }

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) SetFilter(yf yfilter.YFilter) { redMinThresh.YFilter = yf }

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetGoName(yname string) string {
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetSegmentPath() string {
    return "red-min-thresh"
}

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold" {
        return &redMinThresh.Threshold
    }
    return nil
}

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold"] = &redMinThresh.Threshold
    return children
}

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetBundleName() string { return "ietf" }

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetYangName() string { return "red-min-thresh" }

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) SetParent(parent types.Entity) { redMinThresh.parent = parent }

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetParent() types.Entity { return redMinThresh.parent }

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetParentYangName() string { return "random-detect-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold
// threshold
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold size. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    ThresholdSize interface{}

    // Threshold interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are microsecond.
    ThresholdInterval interface{}
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetFilter() yfilter.YFilter { return threshold.YFilter }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) SetFilter(yf yfilter.YFilter) { threshold.YFilter = yf }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetGoName(yname string) string {
    if yname == "threshold-size" { return "ThresholdSize" }
    if yname == "threshold-interval" { return "ThresholdInterval" }
    return ""
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetSegmentPath() string {
    return "threshold"
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-size"] = threshold.ThresholdSize
    leafs["threshold-interval"] = threshold.ThresholdInterval
    return leafs
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetBundleName() string { return "ietf" }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetYangName() string { return "threshold" }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) SetParent(parent types.Entity) { threshold.parent = parent }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetParent() types.Entity { return threshold.parent }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetParentYangName() string { return "red-min-thresh" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh
// Maximum threshold
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // threshold.
    Threshold Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold
}

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetFilter() yfilter.YFilter { return redMaxThresh.YFilter }

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) SetFilter(yf yfilter.YFilter) { redMaxThresh.YFilter = yf }

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetGoName(yname string) string {
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetSegmentPath() string {
    return "red-max-thresh"
}

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold" {
        return &redMaxThresh.Threshold
    }
    return nil
}

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold"] = &redMaxThresh.Threshold
    return children
}

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetBundleName() string { return "ietf" }

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetYangName() string { return "red-max-thresh" }

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) SetParent(parent types.Entity) { redMaxThresh.parent = parent }

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetParent() types.Entity { return redMaxThresh.parent }

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetParentYangName() string { return "random-detect-cfg" }

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold
// threshold
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold size. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    ThresholdSize interface{}

    // Threshold interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are microsecond.
    ThresholdInterval interface{}
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetFilter() yfilter.YFilter { return threshold.YFilter }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) SetFilter(yf yfilter.YFilter) { threshold.YFilter = yf }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetGoName(yname string) string {
    if yname == "threshold-size" { return "ThresholdSize" }
    if yname == "threshold-interval" { return "ThresholdInterval" }
    return ""
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetSegmentPath() string {
    return "threshold"
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-size"] = threshold.ThresholdSize
    leafs["threshold-interval"] = threshold.ThresholdInterval
    return leafs
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetBundleName() string { return "ietf" }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetYangName() string { return "threshold" }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) SetParent(parent types.Entity) { threshold.parent = parent }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetParent() types.Entity { return threshold.parent }

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetParentYangName() string { return "red-max-thresh" }

