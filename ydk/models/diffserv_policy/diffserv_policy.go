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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // policy template. The type is slice of Policies_PolicyEntry.
    PolicyEntry []*Policies_PolicyEntry
}

func (policies *Policies) GetEntityData() *types.CommonEntityData {
    policies.EntityData.YFilter = policies.YFilter
    policies.EntityData.YangName = "policies"
    policies.EntityData.BundleName = "ietf"
    policies.EntityData.ParentYangName = "ietf-diffserv-policy"
    policies.EntityData.SegmentPath = "ietf-diffserv-policy:policies"
    policies.EntityData.AbsolutePath = policies.EntityData.SegmentPath
    policies.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    policies.EntityData.NamespaceTable = ietf.GetNamespaces()
    policies.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    policies.EntityData.Children = types.NewOrderedMap()
    policies.EntityData.Children.Append("policy-entry", types.YChild{"PolicyEntry", nil})
    for i := range policies.PolicyEntry {
        policies.EntityData.Children.Append(types.GetSegmentPath(policies.PolicyEntry[i]), types.YChild{"PolicyEntry", policies.PolicyEntry[i]})
    }
    policies.EntityData.Leafs = types.NewOrderedMap()

    policies.EntityData.YListKeys = []string {}

    return &(policies.EntityData)
}

// Policies_PolicyEntry
// policy template
type Policies_PolicyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Diffserv policy name. The type is string.
    PolicyName interface{}

    // Diffserv policy description. The type is string.
    PolicyDescr interface{}

    // Classifier entry configuration in a policy. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry.
    ClassifierEntry []*Policies_PolicyEntry_ClassifierEntry
}

func (policyEntry *Policies_PolicyEntry) GetEntityData() *types.CommonEntityData {
    policyEntry.EntityData.YFilter = policyEntry.YFilter
    policyEntry.EntityData.YangName = "policy-entry"
    policyEntry.EntityData.BundleName = "ietf"
    policyEntry.EntityData.ParentYangName = "policies"
    policyEntry.EntityData.SegmentPath = "policy-entry" + types.AddKeyToken(policyEntry.PolicyName, "policy-name")
    policyEntry.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/" + policyEntry.EntityData.SegmentPath
    policyEntry.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    policyEntry.EntityData.NamespaceTable = ietf.GetNamespaces()
    policyEntry.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    policyEntry.EntityData.Children = types.NewOrderedMap()
    policyEntry.EntityData.Children.Append("classifier-entry", types.YChild{"ClassifierEntry", nil})
    for i := range policyEntry.ClassifierEntry {
        policyEntry.EntityData.Children.Append(types.GetSegmentPath(policyEntry.ClassifierEntry[i]), types.YChild{"ClassifierEntry", policyEntry.ClassifierEntry[i]})
    }
    policyEntry.EntityData.Leafs = types.NewOrderedMap()
    policyEntry.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", policyEntry.PolicyName})
    policyEntry.EntityData.Leafs.Append("policy-descr", types.YLeaf{"PolicyDescr", policyEntry.PolicyDescr})

    policyEntry.EntityData.YListKeys = []string {"PolicyName"}

    return &(policyEntry.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry
// Classifier entry configuration in a policy
type Policies_PolicyEntry_ClassifierEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    FilterEntry []*Policies_PolicyEntry_ClassifierEntry_FilterEntry

    // Configuration of classifier & associated actions. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg.
    ClassifierActionEntryCfg []*Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg
}

func (classifierEntry *Policies_PolicyEntry_ClassifierEntry) GetEntityData() *types.CommonEntityData {
    classifierEntry.EntityData.YFilter = classifierEntry.YFilter
    classifierEntry.EntityData.YangName = "classifier-entry"
    classifierEntry.EntityData.BundleName = "ietf"
    classifierEntry.EntityData.ParentYangName = "policy-entry"
    classifierEntry.EntityData.SegmentPath = "classifier-entry" + types.AddKeyToken(classifierEntry.ClassifierEntryName, "classifier-entry-name")
    classifierEntry.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/" + classifierEntry.EntityData.SegmentPath
    classifierEntry.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    classifierEntry.EntityData.NamespaceTable = ietf.GetNamespaces()
    classifierEntry.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    classifierEntry.EntityData.Children = types.NewOrderedMap()
    classifierEntry.EntityData.Children.Append("filter-entry", types.YChild{"FilterEntry", nil})
    for i := range classifierEntry.FilterEntry {
        classifierEntry.EntityData.Children.Append(types.GetSegmentPath(classifierEntry.FilterEntry[i]), types.YChild{"FilterEntry", classifierEntry.FilterEntry[i]})
    }
    classifierEntry.EntityData.Children.Append("classifier-action-entry-cfg", types.YChild{"ClassifierActionEntryCfg", nil})
    for i := range classifierEntry.ClassifierActionEntryCfg {
        classifierEntry.EntityData.Children.Append(types.GetSegmentPath(classifierEntry.ClassifierActionEntryCfg[i]), types.YChild{"ClassifierActionEntryCfg", classifierEntry.ClassifierActionEntryCfg[i]})
    }
    classifierEntry.EntityData.Leafs = types.NewOrderedMap()
    classifierEntry.EntityData.Leafs.Append("classifier-entry-name", types.YLeaf{"ClassifierEntryName", classifierEntry.ClassifierEntryName})
    classifierEntry.EntityData.Leafs.Append("classifier-entry-inline", types.YLeaf{"ClassifierEntryInline", classifierEntry.ClassifierEntryInline})
    classifierEntry.EntityData.Leafs.Append("classifier-entry-filter-oper", types.YLeaf{"ClassifierEntryFilterOper", classifierEntry.ClassifierEntryFilterOper})

    classifierEntry.EntityData.YListKeys = []string {"ClassifierEntryName"}

    return &(classifierEntry.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_FilterEntry
// Filters configured inline in a policy
type Policies_PolicyEntry_ClassifierEntry_FilterEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    DscpCfg []*Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg

    // list of source ip address. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg.
    SourceIpAddressCfg []*Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg

    // list of destination ip address. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg.
    DestinationIpAddressCfg []*Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg

    // list of ranges of source port. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg.
    SourcePortCfg []*Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg

    // list of ranges of destination port. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg.
    DestinationPortCfg []*Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg

    // list of ranges of protocol values. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg.
    ProtocolCfg []*Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg
}

func (filterEntry *Policies_PolicyEntry_ClassifierEntry_FilterEntry) GetEntityData() *types.CommonEntityData {
    filterEntry.EntityData.YFilter = filterEntry.YFilter
    filterEntry.EntityData.YangName = "filter-entry"
    filterEntry.EntityData.BundleName = "ietf"
    filterEntry.EntityData.ParentYangName = "classifier-entry"
    filterEntry.EntityData.SegmentPath = "filter-entry" + types.AddKeyToken(filterEntry.FilterType, "filter-type") + types.AddKeyToken(filterEntry.FilterLogicalNot, "filter-logical-not")
    filterEntry.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/" + filterEntry.EntityData.SegmentPath
    filterEntry.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    filterEntry.EntityData.NamespaceTable = ietf.GetNamespaces()
    filterEntry.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    filterEntry.EntityData.Children = types.NewOrderedMap()
    filterEntry.EntityData.Children.Append("dscp-cfg", types.YChild{"DscpCfg", nil})
    for i := range filterEntry.DscpCfg {
        filterEntry.EntityData.Children.Append(types.GetSegmentPath(filterEntry.DscpCfg[i]), types.YChild{"DscpCfg", filterEntry.DscpCfg[i]})
    }
    filterEntry.EntityData.Children.Append("source-ip-address-cfg", types.YChild{"SourceIpAddressCfg", nil})
    for i := range filterEntry.SourceIpAddressCfg {
        filterEntry.EntityData.Children.Append(types.GetSegmentPath(filterEntry.SourceIpAddressCfg[i]), types.YChild{"SourceIpAddressCfg", filterEntry.SourceIpAddressCfg[i]})
    }
    filterEntry.EntityData.Children.Append("destination-ip-address-cfg", types.YChild{"DestinationIpAddressCfg", nil})
    for i := range filterEntry.DestinationIpAddressCfg {
        filterEntry.EntityData.Children.Append(types.GetSegmentPath(filterEntry.DestinationIpAddressCfg[i]), types.YChild{"DestinationIpAddressCfg", filterEntry.DestinationIpAddressCfg[i]})
    }
    filterEntry.EntityData.Children.Append("source-port-cfg", types.YChild{"SourcePortCfg", nil})
    for i := range filterEntry.SourcePortCfg {
        filterEntry.EntityData.Children.Append(types.GetSegmentPath(filterEntry.SourcePortCfg[i]), types.YChild{"SourcePortCfg", filterEntry.SourcePortCfg[i]})
    }
    filterEntry.EntityData.Children.Append("destination-port-cfg", types.YChild{"DestinationPortCfg", nil})
    for i := range filterEntry.DestinationPortCfg {
        filterEntry.EntityData.Children.Append(types.GetSegmentPath(filterEntry.DestinationPortCfg[i]), types.YChild{"DestinationPortCfg", filterEntry.DestinationPortCfg[i]})
    }
    filterEntry.EntityData.Children.Append("protocol-cfg", types.YChild{"ProtocolCfg", nil})
    for i := range filterEntry.ProtocolCfg {
        filterEntry.EntityData.Children.Append(types.GetSegmentPath(filterEntry.ProtocolCfg[i]), types.YChild{"ProtocolCfg", filterEntry.ProtocolCfg[i]})
    }
    filterEntry.EntityData.Leafs = types.NewOrderedMap()
    filterEntry.EntityData.Leafs.Append("filter-type", types.YLeaf{"FilterType", filterEntry.FilterType})
    filterEntry.EntityData.Leafs.Append("filter-logical-not", types.YLeaf{"FilterLogicalNot", filterEntry.FilterLogicalNot})

    filterEntry.EntityData.YListKeys = []string {"FilterType", "FilterLogicalNot"}

    return &(filterEntry.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg
// list of dscp ranges
type Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Minimum value of dscp range. The type is
    // interface{} with range: 0..63.
    DscpMin interface{}

    // This attribute is a key. maximum value of dscp range. The type is
    // interface{} with range: 0..63.
    DscpMax interface{}
}

func (dscpCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DscpCfg) GetEntityData() *types.CommonEntityData {
    dscpCfg.EntityData.YFilter = dscpCfg.YFilter
    dscpCfg.EntityData.YangName = "dscp-cfg"
    dscpCfg.EntityData.BundleName = "ietf"
    dscpCfg.EntityData.ParentYangName = "filter-entry"
    dscpCfg.EntityData.SegmentPath = "dscp-cfg" + types.AddKeyToken(dscpCfg.DscpMin, "dscp-min") + types.AddKeyToken(dscpCfg.DscpMax, "dscp-max")
    dscpCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/filter-entry/" + dscpCfg.EntityData.SegmentPath
    dscpCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    dscpCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    dscpCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    dscpCfg.EntityData.Children = types.NewOrderedMap()
    dscpCfg.EntityData.Leafs = types.NewOrderedMap()
    dscpCfg.EntityData.Leafs.Append("dscp-min", types.YLeaf{"DscpMin", dscpCfg.DscpMin})
    dscpCfg.EntityData.Leafs.Append("dscp-max", types.YLeaf{"DscpMax", dscpCfg.DscpMax})

    dscpCfg.EntityData.YListKeys = []string {"DscpMin", "DscpMax"}

    return &(dscpCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg
// list of source ip address
type Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. source ip prefix. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    SourceIpAddr interface{}
}

func (sourceIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourceIpAddressCfg) GetEntityData() *types.CommonEntityData {
    sourceIpAddressCfg.EntityData.YFilter = sourceIpAddressCfg.YFilter
    sourceIpAddressCfg.EntityData.YangName = "source-ip-address-cfg"
    sourceIpAddressCfg.EntityData.BundleName = "ietf"
    sourceIpAddressCfg.EntityData.ParentYangName = "filter-entry"
    sourceIpAddressCfg.EntityData.SegmentPath = "source-ip-address-cfg" + types.AddKeyToken(sourceIpAddressCfg.SourceIpAddr, "source-ip-addr")
    sourceIpAddressCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/filter-entry/" + sourceIpAddressCfg.EntityData.SegmentPath
    sourceIpAddressCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    sourceIpAddressCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    sourceIpAddressCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    sourceIpAddressCfg.EntityData.Children = types.NewOrderedMap()
    sourceIpAddressCfg.EntityData.Leafs = types.NewOrderedMap()
    sourceIpAddressCfg.EntityData.Leafs.Append("source-ip-addr", types.YLeaf{"SourceIpAddr", sourceIpAddressCfg.SourceIpAddr})

    sourceIpAddressCfg.EntityData.YListKeys = []string {"SourceIpAddr"}

    return &(sourceIpAddressCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg
// list of destination ip address
type Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. destination ip prefix. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    DestinationIpAddr interface{}
}

func (destinationIpAddressCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationIpAddressCfg) GetEntityData() *types.CommonEntityData {
    destinationIpAddressCfg.EntityData.YFilter = destinationIpAddressCfg.YFilter
    destinationIpAddressCfg.EntityData.YangName = "destination-ip-address-cfg"
    destinationIpAddressCfg.EntityData.BundleName = "ietf"
    destinationIpAddressCfg.EntityData.ParentYangName = "filter-entry"
    destinationIpAddressCfg.EntityData.SegmentPath = "destination-ip-address-cfg" + types.AddKeyToken(destinationIpAddressCfg.DestinationIpAddr, "destination-ip-addr")
    destinationIpAddressCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/filter-entry/" + destinationIpAddressCfg.EntityData.SegmentPath
    destinationIpAddressCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    destinationIpAddressCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    destinationIpAddressCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    destinationIpAddressCfg.EntityData.Children = types.NewOrderedMap()
    destinationIpAddressCfg.EntityData.Leafs = types.NewOrderedMap()
    destinationIpAddressCfg.EntityData.Leafs.Append("destination-ip-addr", types.YLeaf{"DestinationIpAddr", destinationIpAddressCfg.DestinationIpAddr})

    destinationIpAddressCfg.EntityData.YListKeys = []string {"DestinationIpAddr"}

    return &(destinationIpAddressCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg
// list of ranges of source port
type Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. minimum value of source port range. The type is
    // interface{} with range: 0..65535.
    SourcePortMin interface{}

    // This attribute is a key. maximum value of source port range. The type is
    // interface{} with range: 0..65535.
    SourcePortMax interface{}
}

func (sourcePortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_SourcePortCfg) GetEntityData() *types.CommonEntityData {
    sourcePortCfg.EntityData.YFilter = sourcePortCfg.YFilter
    sourcePortCfg.EntityData.YangName = "source-port-cfg"
    sourcePortCfg.EntityData.BundleName = "ietf"
    sourcePortCfg.EntityData.ParentYangName = "filter-entry"
    sourcePortCfg.EntityData.SegmentPath = "source-port-cfg" + types.AddKeyToken(sourcePortCfg.SourcePortMin, "source-port-min") + types.AddKeyToken(sourcePortCfg.SourcePortMax, "source-port-max")
    sourcePortCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/filter-entry/" + sourcePortCfg.EntityData.SegmentPath
    sourcePortCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    sourcePortCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    sourcePortCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    sourcePortCfg.EntityData.Children = types.NewOrderedMap()
    sourcePortCfg.EntityData.Leafs = types.NewOrderedMap()
    sourcePortCfg.EntityData.Leafs.Append("source-port-min", types.YLeaf{"SourcePortMin", sourcePortCfg.SourcePortMin})
    sourcePortCfg.EntityData.Leafs.Append("source-port-max", types.YLeaf{"SourcePortMax", sourcePortCfg.SourcePortMax})

    sourcePortCfg.EntityData.YListKeys = []string {"SourcePortMin", "SourcePortMax"}

    return &(sourcePortCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg
// list of ranges of destination port
type Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. minimum value of destination port range. The type
    // is interface{} with range: 0..65535.
    DestinationPortMin interface{}

    // This attribute is a key. maximum value of destination port range. The type
    // is interface{} with range: 0..65535.
    DestinationPortMax interface{}
}

func (destinationPortCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_DestinationPortCfg) GetEntityData() *types.CommonEntityData {
    destinationPortCfg.EntityData.YFilter = destinationPortCfg.YFilter
    destinationPortCfg.EntityData.YangName = "destination-port-cfg"
    destinationPortCfg.EntityData.BundleName = "ietf"
    destinationPortCfg.EntityData.ParentYangName = "filter-entry"
    destinationPortCfg.EntityData.SegmentPath = "destination-port-cfg" + types.AddKeyToken(destinationPortCfg.DestinationPortMin, "destination-port-min") + types.AddKeyToken(destinationPortCfg.DestinationPortMax, "destination-port-max")
    destinationPortCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/filter-entry/" + destinationPortCfg.EntityData.SegmentPath
    destinationPortCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    destinationPortCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    destinationPortCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    destinationPortCfg.EntityData.Children = types.NewOrderedMap()
    destinationPortCfg.EntityData.Leafs = types.NewOrderedMap()
    destinationPortCfg.EntityData.Leafs.Append("destination-port-min", types.YLeaf{"DestinationPortMin", destinationPortCfg.DestinationPortMin})
    destinationPortCfg.EntityData.Leafs.Append("destination-port-max", types.YLeaf{"DestinationPortMax", destinationPortCfg.DestinationPortMax})

    destinationPortCfg.EntityData.YListKeys = []string {"DestinationPortMin", "DestinationPortMax"}

    return &(destinationPortCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg
// list of ranges of protocol values
type Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. minimum value of protocol range. The type is
    // interface{} with range: 0..255.
    ProtocolMin interface{}

    // This attribute is a key. maximum value of protocol range. The type is
    // interface{} with range: 0..255.
    ProtocolMax interface{}
}

func (protocolCfg *Policies_PolicyEntry_ClassifierEntry_FilterEntry_ProtocolCfg) GetEntityData() *types.CommonEntityData {
    protocolCfg.EntityData.YFilter = protocolCfg.YFilter
    protocolCfg.EntityData.YangName = "protocol-cfg"
    protocolCfg.EntityData.BundleName = "ietf"
    protocolCfg.EntityData.ParentYangName = "filter-entry"
    protocolCfg.EntityData.SegmentPath = "protocol-cfg" + types.AddKeyToken(protocolCfg.ProtocolMin, "protocol-min") + types.AddKeyToken(protocolCfg.ProtocolMax, "protocol-max")
    protocolCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/filter-entry/" + protocolCfg.EntityData.SegmentPath
    protocolCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    protocolCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    protocolCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    protocolCfg.EntityData.Children = types.NewOrderedMap()
    protocolCfg.EntityData.Leafs = types.NewOrderedMap()
    protocolCfg.EntityData.Leafs.Append("protocol-min", types.YLeaf{"ProtocolMin", protocolCfg.ProtocolMin})
    protocolCfg.EntityData.Leafs.Append("protocol-max", types.YLeaf{"ProtocolMax", protocolCfg.ProtocolMax})

    protocolCfg.EntityData.YListKeys = []string {"ProtocolMin", "ProtocolMax"}

    return &(protocolCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg
// Configuration of classifier & associated actions
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (classifierActionEntryCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg) GetEntityData() *types.CommonEntityData {
    classifierActionEntryCfg.EntityData.YFilter = classifierActionEntryCfg.YFilter
    classifierActionEntryCfg.EntityData.YangName = "classifier-action-entry-cfg"
    classifierActionEntryCfg.EntityData.BundleName = "ietf"
    classifierActionEntryCfg.EntityData.ParentYangName = "classifier-entry"
    classifierActionEntryCfg.EntityData.SegmentPath = "classifier-action-entry-cfg" + types.AddKeyToken(classifierActionEntryCfg.ActionType, "action-type")
    classifierActionEntryCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/" + classifierActionEntryCfg.EntityData.SegmentPath
    classifierActionEntryCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    classifierActionEntryCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    classifierActionEntryCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    classifierActionEntryCfg.EntityData.Children = types.NewOrderedMap()
    classifierActionEntryCfg.EntityData.Children.Append("ietf-diffserv-action:marking-cfg", types.YChild{"MarkingCfg", &classifierActionEntryCfg.MarkingCfg})
    classifierActionEntryCfg.EntityData.Children.Append("ietf-diffserv-action:priority-cfg", types.YChild{"PriorityCfg", &classifierActionEntryCfg.PriorityCfg})
    classifierActionEntryCfg.EntityData.Children.Append("ietf-diffserv-action:meter-cfg", types.YChild{"MeterCfg", &classifierActionEntryCfg.MeterCfg})
    classifierActionEntryCfg.EntityData.Children.Append("ietf-diffserv-action:min-rate-cfg", types.YChild{"MinRateCfg", &classifierActionEntryCfg.MinRateCfg})
    classifierActionEntryCfg.EntityData.Children.Append("ietf-diffserv-action:max-rate-cfg", types.YChild{"MaxRateCfg", &classifierActionEntryCfg.MaxRateCfg})
    classifierActionEntryCfg.EntityData.Children.Append("ietf-diffserv-action:drop-cfg", types.YChild{"DropCfg", &classifierActionEntryCfg.DropCfg})
    classifierActionEntryCfg.EntityData.Children.Append("ietf-diffserv-action:tail-drop-cfg", types.YChild{"TailDropCfg", &classifierActionEntryCfg.TailDropCfg})
    classifierActionEntryCfg.EntityData.Children.Append("ietf-diffserv-action:random-detect-cfg", types.YChild{"RandomDetectCfg", &classifierActionEntryCfg.RandomDetectCfg})
    classifierActionEntryCfg.EntityData.Leafs = types.NewOrderedMap()
    classifierActionEntryCfg.EntityData.Leafs.Append("action-type", types.YLeaf{"ActionType", classifierActionEntryCfg.ActionType})

    classifierActionEntryCfg.EntityData.YListKeys = []string {"ActionType"}

    return &(classifierActionEntryCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg
// Marking configuration container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dscp marking. The type is interface{} with range: 0..63.
    Dscp interface{}
}

func (markingCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MarkingCfg) GetEntityData() *types.CommonEntityData {
    markingCfg.EntityData.YFilter = markingCfg.YFilter
    markingCfg.EntityData.YangName = "marking-cfg"
    markingCfg.EntityData.BundleName = "ietf"
    markingCfg.EntityData.ParentYangName = "classifier-action-entry-cfg"
    markingCfg.EntityData.SegmentPath = "ietf-diffserv-action:marking-cfg"
    markingCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/" + markingCfg.EntityData.SegmentPath
    markingCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    markingCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    markingCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    markingCfg.EntityData.Children = types.NewOrderedMap()
    markingCfg.EntityData.Leafs = types.NewOrderedMap()
    markingCfg.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", markingCfg.Dscp})

    markingCfg.EntityData.YListKeys = []string {}

    return &(markingCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg
// priority attributes container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // priority level. The type is interface{} with range: 0..255.
    PriorityLevel interface{}

    // absolute priority rate with/without burst rateand absolute percent.
    RateBurst Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst
}

func (priorityCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg) GetEntityData() *types.CommonEntityData {
    priorityCfg.EntityData.YFilter = priorityCfg.YFilter
    priorityCfg.EntityData.YangName = "priority-cfg"
    priorityCfg.EntityData.BundleName = "ietf"
    priorityCfg.EntityData.ParentYangName = "classifier-action-entry-cfg"
    priorityCfg.EntityData.SegmentPath = "ietf-diffserv-action:priority-cfg"
    priorityCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/" + priorityCfg.EntityData.SegmentPath
    priorityCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    priorityCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    priorityCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    priorityCfg.EntityData.Children = types.NewOrderedMap()
    priorityCfg.EntityData.Children.Append("rate-burst", types.YChild{"RateBurst", &priorityCfg.RateBurst})
    priorityCfg.EntityData.Leafs = types.NewOrderedMap()
    priorityCfg.EntityData.Leafs.Append("priority-level", types.YLeaf{"PriorityLevel", priorityCfg.PriorityLevel})

    priorityCfg.EntityData.YListKeys = []string {}

    return &(priorityCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst
// absolute priority rate with/without burst rateand absolute percent
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst struct {
    EntityData types.CommonEntityData
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

func (rateBurst *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_PriorityCfg_RateBurst) GetEntityData() *types.CommonEntityData {
    rateBurst.EntityData.YFilter = rateBurst.YFilter
    rateBurst.EntityData.YangName = "rate-burst"
    rateBurst.EntityData.BundleName = "ietf"
    rateBurst.EntityData.ParentYangName = "priority-cfg"
    rateBurst.EntityData.SegmentPath = "rate-burst"
    rateBurst.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/ietf-diffserv-action:priority-cfg/" + rateBurst.EntityData.SegmentPath
    rateBurst.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    rateBurst.EntityData.NamespaceTable = ietf.GetNamespaces()
    rateBurst.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    rateBurst.EntityData.Children = types.NewOrderedMap()
    rateBurst.EntityData.Leafs = types.NewOrderedMap()
    rateBurst.EntityData.Leafs.Append("rate", types.YLeaf{"Rate", rateBurst.Rate})
    rateBurst.EntityData.Leafs.Append("absolute-rate-metric", types.YLeaf{"AbsoluteRateMetric", rateBurst.AbsoluteRateMetric})
    rateBurst.EntityData.Leafs.Append("absolute-rate-units", types.YLeaf{"AbsoluteRateUnits", rateBurst.AbsoluteRateUnits})
    rateBurst.EntityData.Leafs.Append("rate-percent", types.YLeaf{"RatePercent", rateBurst.RatePercent})
    rateBurst.EntityData.Leafs.Append("rate-ratio", types.YLeaf{"RateRatio", rateBurst.RateRatio})
    rateBurst.EntityData.Leafs.Append("burst-size", types.YLeaf{"BurstSize", rateBurst.BurstSize})
    rateBurst.EntityData.Leafs.Append("burst-interval", types.YLeaf{"BurstInterval", rateBurst.BurstInterval})

    rateBurst.EntityData.YListKeys = []string {}

    return &(rateBurst.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg
// Meter list configuration container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Meter configuration. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList.
    MeterList []*Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList
}

func (meterCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg) GetEntityData() *types.CommonEntityData {
    meterCfg.EntityData.YFilter = meterCfg.YFilter
    meterCfg.EntityData.YangName = "meter-cfg"
    meterCfg.EntityData.BundleName = "ietf"
    meterCfg.EntityData.ParentYangName = "classifier-action-entry-cfg"
    meterCfg.EntityData.SegmentPath = "ietf-diffserv-action:meter-cfg"
    meterCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/" + meterCfg.EntityData.SegmentPath
    meterCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    meterCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    meterCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    meterCfg.EntityData.Children = types.NewOrderedMap()
    meterCfg.EntityData.Children.Append("meter-list", types.YChild{"MeterList", nil})
    for i := range meterCfg.MeterList {
        meterCfg.EntityData.Children.Append(types.GetSegmentPath(meterCfg.MeterList[i]), types.YChild{"MeterList", meterCfg.MeterList[i]})
    }
    meterCfg.EntityData.Leafs = types.NewOrderedMap()

    meterCfg.EntityData.YListKeys = []string {}

    return &(meterCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList
// Meter configuration
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (meterList *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList) GetEntityData() *types.CommonEntityData {
    meterList.EntityData.YFilter = meterList.YFilter
    meterList.EntityData.YangName = "meter-list"
    meterList.EntityData.BundleName = "ietf"
    meterList.EntityData.ParentYangName = "meter-cfg"
    meterList.EntityData.SegmentPath = "meter-list" + types.AddKeyToken(meterList.MeterId, "meter-id")
    meterList.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/ietf-diffserv-action:meter-cfg/" + meterList.EntityData.SegmentPath
    meterList.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    meterList.EntityData.NamespaceTable = ietf.GetNamespaces()
    meterList.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    meterList.EntityData.Children = types.NewOrderedMap()
    meterList.EntityData.Children.Append("color", types.YChild{"Color", &meterList.Color})
    meterList.EntityData.Children.Append("succeed-action", types.YChild{"SucceedAction", &meterList.SucceedAction})
    meterList.EntityData.Children.Append("fail-action", types.YChild{"FailAction", &meterList.FailAction})
    meterList.EntityData.Leafs = types.NewOrderedMap()
    meterList.EntityData.Leafs.Append("meter-id", types.YLeaf{"MeterId", meterList.MeterId})
    meterList.EntityData.Leafs.Append("meter-rate", types.YLeaf{"MeterRate", meterList.MeterRate})
    meterList.EntityData.Leafs.Append("burst-size", types.YLeaf{"BurstSize", meterList.BurstSize})
    meterList.EntityData.Leafs.Append("burst-interval", types.YLeaf{"BurstInterval", meterList.BurstInterval})

    meterList.EntityData.YListKeys = []string {"MeterId"}

    return &(meterList.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color
// color aware & color blind attributes container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color struct {
    EntityData types.CommonEntityData
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

func (color *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_Color) GetEntityData() *types.CommonEntityData {
    color.EntityData.YFilter = color.YFilter
    color.EntityData.YangName = "color"
    color.EntityData.BundleName = "ietf"
    color.EntityData.ParentYangName = "meter-list"
    color.EntityData.SegmentPath = "color"
    color.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/ietf-diffserv-action:meter-cfg/meter-list/" + color.EntityData.SegmentPath
    color.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    color.EntityData.NamespaceTable = ietf.GetNamespaces()
    color.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    color.EntityData.Children = types.NewOrderedMap()
    color.EntityData.Leafs = types.NewOrderedMap()
    color.EntityData.Leafs.Append("classifier-entry-name", types.YLeaf{"ClassifierEntryName", color.ClassifierEntryName})
    color.EntityData.Leafs.Append("classifier-entry-descr", types.YLeaf{"ClassifierEntryDescr", color.ClassifierEntryDescr})
    color.EntityData.Leafs.Append("classifier-entry-filter-operation", types.YLeaf{"ClassifierEntryFilterOperation", color.ClassifierEntryFilterOperation})

    color.EntityData.YListKeys = []string {}

    return &(color.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction
// confirm action
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction struct {
    EntityData types.CommonEntityData
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

func (succeedAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_SucceedAction) GetEntityData() *types.CommonEntityData {
    succeedAction.EntityData.YFilter = succeedAction.YFilter
    succeedAction.EntityData.YangName = "succeed-action"
    succeedAction.EntityData.BundleName = "ietf"
    succeedAction.EntityData.ParentYangName = "meter-list"
    succeedAction.EntityData.SegmentPath = "succeed-action"
    succeedAction.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/ietf-diffserv-action:meter-cfg/meter-list/" + succeedAction.EntityData.SegmentPath
    succeedAction.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    succeedAction.EntityData.NamespaceTable = ietf.GetNamespaces()
    succeedAction.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    succeedAction.EntityData.Children = types.NewOrderedMap()
    succeedAction.EntityData.Leafs = types.NewOrderedMap()
    succeedAction.EntityData.Leafs.Append("meter-action-type", types.YLeaf{"MeterActionType", succeedAction.MeterActionType})
    succeedAction.EntityData.Leafs.Append("next-meter-id", types.YLeaf{"NextMeterId", succeedAction.NextMeterId})
    succeedAction.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", succeedAction.Dscp})
    succeedAction.EntityData.Leafs.Append("drop-action", types.YLeaf{"DropAction", succeedAction.DropAction})

    succeedAction.EntityData.YListKeys = []string {}

    return &(succeedAction.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction
// exceed action
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction struct {
    EntityData types.CommonEntityData
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

func (failAction *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MeterCfg_MeterList_FailAction) GetEntityData() *types.CommonEntityData {
    failAction.EntityData.YFilter = failAction.YFilter
    failAction.EntityData.YangName = "fail-action"
    failAction.EntityData.BundleName = "ietf"
    failAction.EntityData.ParentYangName = "meter-list"
    failAction.EntityData.SegmentPath = "fail-action"
    failAction.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/ietf-diffserv-action:meter-cfg/meter-list/" + failAction.EntityData.SegmentPath
    failAction.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    failAction.EntityData.NamespaceTable = ietf.GetNamespaces()
    failAction.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    failAction.EntityData.Children = types.NewOrderedMap()
    failAction.EntityData.Leafs = types.NewOrderedMap()
    failAction.EntityData.Leafs.Append("meter-action-type", types.YLeaf{"MeterActionType", failAction.MeterActionType})
    failAction.EntityData.Leafs.Append("next-meter-id", types.YLeaf{"NextMeterId", failAction.NextMeterId})
    failAction.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", failAction.Dscp})
    failAction.EntityData.Leafs.Append("drop-action", types.YLeaf{"DropAction", failAction.DropAction})

    failAction.EntityData.YListKeys = []string {}

    return &(failAction.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg
// min guaranteed bandwidth
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg struct {
    EntityData types.CommonEntityData
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

func (minRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg) GetEntityData() *types.CommonEntityData {
    minRateCfg.EntityData.YFilter = minRateCfg.YFilter
    minRateCfg.EntityData.YangName = "min-rate-cfg"
    minRateCfg.EntityData.BundleName = "ietf"
    minRateCfg.EntityData.ParentYangName = "classifier-action-entry-cfg"
    minRateCfg.EntityData.SegmentPath = "ietf-diffserv-action:min-rate-cfg"
    minRateCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/" + minRateCfg.EntityData.SegmentPath
    minRateCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    minRateCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    minRateCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    minRateCfg.EntityData.Children = types.NewOrderedMap()
    minRateCfg.EntityData.Children.Append("bw-excess-share-cfg", types.YChild{"BwExcessShareCfg", &minRateCfg.BwExcessShareCfg})
    minRateCfg.EntityData.Leafs = types.NewOrderedMap()
    minRateCfg.EntityData.Leafs.Append("min-rate", types.YLeaf{"MinRate", minRateCfg.MinRate})
    minRateCfg.EntityData.Leafs.Append("absolute-rate-metric", types.YLeaf{"AbsoluteRateMetric", minRateCfg.AbsoluteRateMetric})
    minRateCfg.EntityData.Leafs.Append("absolute-rate-units", types.YLeaf{"AbsoluteRateUnits", minRateCfg.AbsoluteRateUnits})
    minRateCfg.EntityData.Leafs.Append("rate-percent", types.YLeaf{"RatePercent", minRateCfg.RatePercent})
    minRateCfg.EntityData.Leafs.Append("rate-ratio", types.YLeaf{"RateRatio", minRateCfg.RateRatio})

    minRateCfg.EntityData.YListKeys = []string {}

    return &(minRateCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg
// share the bandwidth remaming
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg struct {
    EntityData types.CommonEntityData
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

func (bwExcessShareCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MinRateCfg_BwExcessShareCfg) GetEntityData() *types.CommonEntityData {
    bwExcessShareCfg.EntityData.YFilter = bwExcessShareCfg.YFilter
    bwExcessShareCfg.EntityData.YangName = "bw-excess-share-cfg"
    bwExcessShareCfg.EntityData.BundleName = "ietf"
    bwExcessShareCfg.EntityData.ParentYangName = "min-rate-cfg"
    bwExcessShareCfg.EntityData.SegmentPath = "bw-excess-share-cfg"
    bwExcessShareCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/ietf-diffserv-action:min-rate-cfg/" + bwExcessShareCfg.EntityData.SegmentPath
    bwExcessShareCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    bwExcessShareCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    bwExcessShareCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    bwExcessShareCfg.EntityData.Children = types.NewOrderedMap()
    bwExcessShareCfg.EntityData.Leafs = types.NewOrderedMap()
    bwExcessShareCfg.EntityData.Leafs.Append("value", types.YLeaf{"Value", bwExcessShareCfg.Value})
    bwExcessShareCfg.EntityData.Leafs.Append("absolute-rate-metric", types.YLeaf{"AbsoluteRateMetric", bwExcessShareCfg.AbsoluteRateMetric})
    bwExcessShareCfg.EntityData.Leafs.Append("absolute-rate-units", types.YLeaf{"AbsoluteRateUnits", bwExcessShareCfg.AbsoluteRateUnits})
    bwExcessShareCfg.EntityData.Leafs.Append("rate-percent", types.YLeaf{"RatePercent", bwExcessShareCfg.RatePercent})
    bwExcessShareCfg.EntityData.Leafs.Append("rate-ratio", types.YLeaf{"RateRatio", bwExcessShareCfg.RateRatio})

    bwExcessShareCfg.EntityData.YListKeys = []string {}

    return &(bwExcessShareCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg
// maximum rate attributes
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg struct {
    EntityData types.CommonEntityData
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

func (maxRateCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_MaxRateCfg) GetEntityData() *types.CommonEntityData {
    maxRateCfg.EntityData.YFilter = maxRateCfg.YFilter
    maxRateCfg.EntityData.YangName = "max-rate-cfg"
    maxRateCfg.EntityData.BundleName = "ietf"
    maxRateCfg.EntityData.ParentYangName = "classifier-action-entry-cfg"
    maxRateCfg.EntityData.SegmentPath = "ietf-diffserv-action:max-rate-cfg"
    maxRateCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/" + maxRateCfg.EntityData.SegmentPath
    maxRateCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    maxRateCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    maxRateCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    maxRateCfg.EntityData.Children = types.NewOrderedMap()
    maxRateCfg.EntityData.Leafs = types.NewOrderedMap()
    maxRateCfg.EntityData.Leafs.Append("absolute-rate", types.YLeaf{"AbsoluteRate", maxRateCfg.AbsoluteRate})
    maxRateCfg.EntityData.Leafs.Append("burst-size", types.YLeaf{"BurstSize", maxRateCfg.BurstSize})
    maxRateCfg.EntityData.Leafs.Append("burst-interval", types.YLeaf{"BurstInterval", maxRateCfg.BurstInterval})
    maxRateCfg.EntityData.Leafs.Append("absolute-rate-metric", types.YLeaf{"AbsoluteRateMetric", maxRateCfg.AbsoluteRateMetric})
    maxRateCfg.EntityData.Leafs.Append("absolute-rate-units", types.YLeaf{"AbsoluteRateUnits", maxRateCfg.AbsoluteRateUnits})
    maxRateCfg.EntityData.Leafs.Append("rate-percent", types.YLeaf{"RatePercent", maxRateCfg.RatePercent})
    maxRateCfg.EntityData.Leafs.Append("rate-ratio", types.YLeaf{"RateRatio", maxRateCfg.RateRatio})

    maxRateCfg.EntityData.YListKeys = []string {}

    return &(maxRateCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg
// Always Drop configuration container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // always drop algorithm. The type is interface{}.
    DropAction interface{}
}

func (dropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_DropCfg) GetEntityData() *types.CommonEntityData {
    dropCfg.EntityData.YFilter = dropCfg.YFilter
    dropCfg.EntityData.YangName = "drop-cfg"
    dropCfg.EntityData.BundleName = "ietf"
    dropCfg.EntityData.ParentYangName = "classifier-action-entry-cfg"
    dropCfg.EntityData.SegmentPath = "ietf-diffserv-action:drop-cfg"
    dropCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/" + dropCfg.EntityData.SegmentPath
    dropCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    dropCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    dropCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    dropCfg.EntityData.Children = types.NewOrderedMap()
    dropCfg.EntityData.Leafs = types.NewOrderedMap()
    dropCfg.EntityData.Leafs.Append("drop-action", types.YLeaf{"DropAction", dropCfg.DropAction})

    dropCfg.EntityData.YListKeys = []string {}

    return &(dropCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg
// Tail Drop configuration container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // the queue limit per dscp range. The type is slice of
    // Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh.
    QlimitDscpThresh []*Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh
}

func (tailDropCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg) GetEntityData() *types.CommonEntityData {
    tailDropCfg.EntityData.YFilter = tailDropCfg.YFilter
    tailDropCfg.EntityData.YangName = "tail-drop-cfg"
    tailDropCfg.EntityData.BundleName = "ietf"
    tailDropCfg.EntityData.ParentYangName = "classifier-action-entry-cfg"
    tailDropCfg.EntityData.SegmentPath = "ietf-diffserv-action:tail-drop-cfg"
    tailDropCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/" + tailDropCfg.EntityData.SegmentPath
    tailDropCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    tailDropCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    tailDropCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    tailDropCfg.EntityData.Children = types.NewOrderedMap()
    tailDropCfg.EntityData.Children.Append("qlimit-dscp-thresh", types.YChild{"QlimitDscpThresh", nil})
    for i := range tailDropCfg.QlimitDscpThresh {
        tailDropCfg.EntityData.Children.Append(types.GetSegmentPath(tailDropCfg.QlimitDscpThresh[i]), types.YChild{"QlimitDscpThresh", tailDropCfg.QlimitDscpThresh[i]})
    }
    tailDropCfg.EntityData.Leafs = types.NewOrderedMap()

    tailDropCfg.EntityData.YListKeys = []string {}

    return &(tailDropCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh
// the queue limit per dscp range
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Minimum of dscp range. The type is interface{}
    // with range: 0..63.
    DscpMin interface{}

    // This attribute is a key. Maximum of dscp range. The type is interface{}
    // with range: 0..63.
    DscpMax interface{}

    // threshold.
    Threshold Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold
}

func (qlimitDscpThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh) GetEntityData() *types.CommonEntityData {
    qlimitDscpThresh.EntityData.YFilter = qlimitDscpThresh.YFilter
    qlimitDscpThresh.EntityData.YangName = "qlimit-dscp-thresh"
    qlimitDscpThresh.EntityData.BundleName = "ietf"
    qlimitDscpThresh.EntityData.ParentYangName = "tail-drop-cfg"
    qlimitDscpThresh.EntityData.SegmentPath = "qlimit-dscp-thresh" + types.AddKeyToken(qlimitDscpThresh.DscpMin, "dscp-min") + types.AddKeyToken(qlimitDscpThresh.DscpMax, "dscp-max")
    qlimitDscpThresh.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/ietf-diffserv-action:tail-drop-cfg/" + qlimitDscpThresh.EntityData.SegmentPath
    qlimitDscpThresh.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    qlimitDscpThresh.EntityData.NamespaceTable = ietf.GetNamespaces()
    qlimitDscpThresh.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    qlimitDscpThresh.EntityData.Children = types.NewOrderedMap()
    qlimitDscpThresh.EntityData.Children.Append("threshold", types.YChild{"Threshold", &qlimitDscpThresh.Threshold})
    qlimitDscpThresh.EntityData.Leafs = types.NewOrderedMap()
    qlimitDscpThresh.EntityData.Leafs.Append("dscp-min", types.YLeaf{"DscpMin", qlimitDscpThresh.DscpMin})
    qlimitDscpThresh.EntityData.Leafs.Append("dscp-max", types.YLeaf{"DscpMax", qlimitDscpThresh.DscpMax})

    qlimitDscpThresh.EntityData.YListKeys = []string {"DscpMin", "DscpMax"}

    return &(qlimitDscpThresh.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold
// threshold
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold size. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    ThresholdSize interface{}

    // Threshold interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are microsecond.
    ThresholdInterval interface{}
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_TailDropCfg_QlimitDscpThresh_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "ietf"
    threshold.EntityData.ParentYangName = "qlimit-dscp-thresh"
    threshold.EntityData.SegmentPath = "threshold"
    threshold.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/ietf-diffserv-action:tail-drop-cfg/qlimit-dscp-thresh/" + threshold.EntityData.SegmentPath
    threshold.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    threshold.EntityData.NamespaceTable = ietf.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    threshold.EntityData.Children = types.NewOrderedMap()
    threshold.EntityData.Leafs = types.NewOrderedMap()
    threshold.EntityData.Leafs.Append("threshold-size", types.YLeaf{"ThresholdSize", threshold.ThresholdSize})
    threshold.EntityData.Leafs.Append("threshold-interval", types.YLeaf{"ThresholdInterval", threshold.ThresholdInterval})

    threshold.EntityData.YListKeys = []string {}

    return &(threshold.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg
// Random Detect configuration container
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg struct {
    EntityData types.CommonEntityData
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

func (randomDetectCfg *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg) GetEntityData() *types.CommonEntityData {
    randomDetectCfg.EntityData.YFilter = randomDetectCfg.YFilter
    randomDetectCfg.EntityData.YangName = "random-detect-cfg"
    randomDetectCfg.EntityData.BundleName = "ietf"
    randomDetectCfg.EntityData.ParentYangName = "classifier-action-entry-cfg"
    randomDetectCfg.EntityData.SegmentPath = "ietf-diffserv-action:random-detect-cfg"
    randomDetectCfg.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/" + randomDetectCfg.EntityData.SegmentPath
    randomDetectCfg.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    randomDetectCfg.EntityData.NamespaceTable = ietf.GetNamespaces()
    randomDetectCfg.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    randomDetectCfg.EntityData.Children = types.NewOrderedMap()
    randomDetectCfg.EntityData.Children.Append("red-min-thresh", types.YChild{"RedMinThresh", &randomDetectCfg.RedMinThresh})
    randomDetectCfg.EntityData.Children.Append("red-max-thresh", types.YChild{"RedMaxThresh", &randomDetectCfg.RedMaxThresh})
    randomDetectCfg.EntityData.Leafs = types.NewOrderedMap()
    randomDetectCfg.EntityData.Leafs.Append("exp-weighting-const", types.YLeaf{"ExpWeightingConst", randomDetectCfg.ExpWeightingConst})
    randomDetectCfg.EntityData.Leafs.Append("mark-probability", types.YLeaf{"MarkProbability", randomDetectCfg.MarkProbability})

    randomDetectCfg.EntityData.YListKeys = []string {}

    return &(randomDetectCfg.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh
// Minimum threshold
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // threshold.
    Threshold Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold
}

func (redMinThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh) GetEntityData() *types.CommonEntityData {
    redMinThresh.EntityData.YFilter = redMinThresh.YFilter
    redMinThresh.EntityData.YangName = "red-min-thresh"
    redMinThresh.EntityData.BundleName = "ietf"
    redMinThresh.EntityData.ParentYangName = "random-detect-cfg"
    redMinThresh.EntityData.SegmentPath = "red-min-thresh"
    redMinThresh.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/ietf-diffserv-action:random-detect-cfg/" + redMinThresh.EntityData.SegmentPath
    redMinThresh.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    redMinThresh.EntityData.NamespaceTable = ietf.GetNamespaces()
    redMinThresh.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    redMinThresh.EntityData.Children = types.NewOrderedMap()
    redMinThresh.EntityData.Children.Append("threshold", types.YChild{"Threshold", &redMinThresh.Threshold})
    redMinThresh.EntityData.Leafs = types.NewOrderedMap()

    redMinThresh.EntityData.YListKeys = []string {}

    return &(redMinThresh.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold
// threshold
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold size. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    ThresholdSize interface{}

    // Threshold interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are microsecond.
    ThresholdInterval interface{}
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMinThresh_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "ietf"
    threshold.EntityData.ParentYangName = "red-min-thresh"
    threshold.EntityData.SegmentPath = "threshold"
    threshold.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/ietf-diffserv-action:random-detect-cfg/red-min-thresh/" + threshold.EntityData.SegmentPath
    threshold.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    threshold.EntityData.NamespaceTable = ietf.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    threshold.EntityData.Children = types.NewOrderedMap()
    threshold.EntityData.Leafs = types.NewOrderedMap()
    threshold.EntityData.Leafs.Append("threshold-size", types.YLeaf{"ThresholdSize", threshold.ThresholdSize})
    threshold.EntityData.Leafs.Append("threshold-interval", types.YLeaf{"ThresholdInterval", threshold.ThresholdInterval})

    threshold.EntityData.YListKeys = []string {}

    return &(threshold.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh
// Maximum threshold
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // threshold.
    Threshold Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold
}

func (redMaxThresh *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh) GetEntityData() *types.CommonEntityData {
    redMaxThresh.EntityData.YFilter = redMaxThresh.YFilter
    redMaxThresh.EntityData.YangName = "red-max-thresh"
    redMaxThresh.EntityData.BundleName = "ietf"
    redMaxThresh.EntityData.ParentYangName = "random-detect-cfg"
    redMaxThresh.EntityData.SegmentPath = "red-max-thresh"
    redMaxThresh.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/ietf-diffserv-action:random-detect-cfg/" + redMaxThresh.EntityData.SegmentPath
    redMaxThresh.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    redMaxThresh.EntityData.NamespaceTable = ietf.GetNamespaces()
    redMaxThresh.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    redMaxThresh.EntityData.Children = types.NewOrderedMap()
    redMaxThresh.EntityData.Children.Append("threshold", types.YChild{"Threshold", &redMaxThresh.Threshold})
    redMaxThresh.EntityData.Leafs = types.NewOrderedMap()

    redMaxThresh.EntityData.YListKeys = []string {}

    return &(redMaxThresh.EntityData)
}

// Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold
// threshold
type Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold size. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    ThresholdSize interface{}

    // Threshold interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are microsecond.
    ThresholdInterval interface{}
}

func (threshold *Policies_PolicyEntry_ClassifierEntry_ClassifierActionEntryCfg_RandomDetectCfg_RedMaxThresh_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "ietf"
    threshold.EntityData.ParentYangName = "red-max-thresh"
    threshold.EntityData.SegmentPath = "threshold"
    threshold.EntityData.AbsolutePath = "ietf-diffserv-policy:policies/policy-entry/classifier-entry/classifier-action-entry-cfg/ietf-diffserv-action:random-detect-cfg/red-max-thresh/" + threshold.EntityData.SegmentPath
    threshold.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    threshold.EntityData.NamespaceTable = ietf.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    threshold.EntityData.Children = types.NewOrderedMap()
    threshold.EntityData.Leafs = types.NewOrderedMap()
    threshold.EntityData.Leafs.Append("threshold-size", types.YLeaf{"ThresholdSize", threshold.ThresholdSize})
    threshold.EntityData.Leafs.Append("threshold-interval", types.YLeaf{"ThresholdInterval", threshold.ThresholdInterval})

    threshold.EntityData.YListKeys = []string {}

    return &(threshold.EntityData)
}

