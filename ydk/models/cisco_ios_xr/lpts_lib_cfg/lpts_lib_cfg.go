// This module contains a collection of YANG definitions
// for Cisco IOS-XR lpts-lib package configuration.
// 
// This module contains definitions
// for the following management objects:
//   lpts: lpts configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package lpts_lib_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lpts_lib_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lpts-lib-cfg lpts}", reflect.TypeOf(Lpts{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lpts-lib-cfg:lpts", reflect.TypeOf(Lpts{}))
}

// Lpts
// lpts configuration commands
type Lpts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pre IFiB Configuration .
    Ipolicer Lpts_Ipolicer

    // Configure penalty timeout value.
    Punt Lpts_Punt
}

func (lpts *Lpts) GetEntityData() *types.CommonEntityData {
    lpts.EntityData.YFilter = lpts.YFilter
    lpts.EntityData.YangName = "lpts"
    lpts.EntityData.BundleName = "cisco_ios_xr"
    lpts.EntityData.ParentYangName = "Cisco-IOS-XR-lpts-lib-cfg"
    lpts.EntityData.SegmentPath = "Cisco-IOS-XR-lpts-lib-cfg:lpts"
    lpts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lpts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lpts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lpts.EntityData.Children = types.NewOrderedMap()
    lpts.EntityData.Children.Append("Cisco-IOS-XR-lpts-pre-ifib-cfg:ipolicer", types.YChild{"Ipolicer", &lpts.Ipolicer})
    lpts.EntityData.Children.Append("Cisco-IOS-XR-lpts-punt-flowtrap-cfg:punt", types.YChild{"Punt", &lpts.Punt})
    lpts.EntityData.Leafs = types.NewOrderedMap()

    lpts.EntityData.YListKeys = []string {}

    return &(lpts.EntityData)
}

// Lpts_Ipolicer
// Pre IFiB Configuration 
// This type is a presence type.
type Lpts_Ipolicer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enabled. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Table for ACLs.
    Acls Lpts_Ipolicer_Acls

    // Table for Flows.
    Flows Lpts_Ipolicer_Flows
}

func (ipolicer *Lpts_Ipolicer) GetEntityData() *types.CommonEntityData {
    ipolicer.EntityData.YFilter = ipolicer.YFilter
    ipolicer.EntityData.YangName = "ipolicer"
    ipolicer.EntityData.BundleName = "cisco_ios_xr"
    ipolicer.EntityData.ParentYangName = "lpts"
    ipolicer.EntityData.SegmentPath = "Cisco-IOS-XR-lpts-pre-ifib-cfg:ipolicer"
    ipolicer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipolicer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipolicer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipolicer.EntityData.Children = types.NewOrderedMap()
    ipolicer.EntityData.Children.Append("acls", types.YChild{"Acls", &ipolicer.Acls})
    ipolicer.EntityData.Children.Append("flows", types.YChild{"Flows", &ipolicer.Flows})
    ipolicer.EntityData.Leafs = types.NewOrderedMap()
    ipolicer.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ipolicer.Enable})

    ipolicer.EntityData.YListKeys = []string {}

    return &(ipolicer.EntityData)
}

// Lpts_Ipolicer_Acls
// Table for ACLs
type Lpts_Ipolicer_Acls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ACL name. The type is slice of Lpts_Ipolicer_Acls_Acl.
    Acl []*Lpts_Ipolicer_Acls_Acl
}

func (acls *Lpts_Ipolicer_Acls) GetEntityData() *types.CommonEntityData {
    acls.EntityData.YFilter = acls.YFilter
    acls.EntityData.YangName = "acls"
    acls.EntityData.BundleName = "cisco_ios_xr"
    acls.EntityData.ParentYangName = "ipolicer"
    acls.EntityData.SegmentPath = "acls"
    acls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acls.EntityData.Children = types.NewOrderedMap()
    acls.EntityData.Children.Append("acl", types.YChild{"Acl", nil})
    for i := range acls.Acl {
        acls.EntityData.Children.Append(types.GetSegmentPath(acls.Acl[i]), types.YChild{"Acl", acls.Acl[i]})
    }
    acls.EntityData.Leafs = types.NewOrderedMap()

    acls.EntityData.YListKeys = []string {}

    return &(acls.EntityData)
}

// Lpts_Ipolicer_Acls_Acl
// ACL name
type Lpts_Ipolicer_Acls_Acl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. ACL name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    AclName interface{}

    // AFI Family.
    AfiTypes Lpts_Ipolicer_Acls_Acl_AfiTypes
}

func (acl *Lpts_Ipolicer_Acls_Acl) GetEntityData() *types.CommonEntityData {
    acl.EntityData.YFilter = acl.YFilter
    acl.EntityData.YangName = "acl"
    acl.EntityData.BundleName = "cisco_ios_xr"
    acl.EntityData.ParentYangName = "acls"
    acl.EntityData.SegmentPath = "acl" + types.AddKeyToken(acl.AclName, "acl-name")
    acl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acl.EntityData.Children = types.NewOrderedMap()
    acl.EntityData.Children.Append("afi-types", types.YChild{"AfiTypes", &acl.AfiTypes})
    acl.EntityData.Leafs = types.NewOrderedMap()
    acl.EntityData.Leafs.Append("acl-name", types.YLeaf{"AclName", acl.AclName})

    acl.EntityData.YListKeys = []string {"AclName"}

    return &(acl.EntityData)
}

// Lpts_Ipolicer_Acls_Acl_AfiTypes
// AFI Family
type Lpts_Ipolicer_Acls_Acl_AfiTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI Family type. The type is slice of
    // Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType.
    AfiType []*Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType
}

func (afiTypes *Lpts_Ipolicer_Acls_Acl_AfiTypes) GetEntityData() *types.CommonEntityData {
    afiTypes.EntityData.YFilter = afiTypes.YFilter
    afiTypes.EntityData.YangName = "afi-types"
    afiTypes.EntityData.BundleName = "cisco_ios_xr"
    afiTypes.EntityData.ParentYangName = "acl"
    afiTypes.EntityData.SegmentPath = "afi-types"
    afiTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afiTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afiTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afiTypes.EntityData.Children = types.NewOrderedMap()
    afiTypes.EntityData.Children.Append("afi-type", types.YChild{"AfiType", nil})
    for i := range afiTypes.AfiType {
        afiTypes.EntityData.Children.Append(types.GetSegmentPath(afiTypes.AfiType[i]), types.YChild{"AfiType", afiTypes.AfiType[i]})
    }
    afiTypes.EntityData.Leafs = types.NewOrderedMap()

    afiTypes.EntityData.YListKeys = []string {}

    return &(afiTypes.EntityData)
}

// Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType
// AFI Family type
type Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. AFI Family Type. The type is Lptsafi.
    AfiFamilyType interface{}

    // VRF list.
    VrfNames Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType_VrfNames
}

func (afiType *Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType) GetEntityData() *types.CommonEntityData {
    afiType.EntityData.YFilter = afiType.YFilter
    afiType.EntityData.YangName = "afi-type"
    afiType.EntityData.BundleName = "cisco_ios_xr"
    afiType.EntityData.ParentYangName = "afi-types"
    afiType.EntityData.SegmentPath = "afi-type" + types.AddKeyToken(afiType.AfiFamilyType, "afi-family-type")
    afiType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afiType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afiType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afiType.EntityData.Children = types.NewOrderedMap()
    afiType.EntityData.Children.Append("vrf-names", types.YChild{"VrfNames", &afiType.VrfNames})
    afiType.EntityData.Leafs = types.NewOrderedMap()
    afiType.EntityData.Leafs.Append("afi-family-type", types.YLeaf{"AfiFamilyType", afiType.AfiFamilyType})

    afiType.EntityData.YListKeys = []string {"AfiFamilyType"}

    return &(afiType.EntityData)
}

// Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType_VrfNames
// VRF list
type Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType_VrfNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is slice of
    // Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType_VrfNames_VrfName.
    VrfName []*Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType_VrfNames_VrfName
}

func (vrfNames *Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType_VrfNames) GetEntityData() *types.CommonEntityData {
    vrfNames.EntityData.YFilter = vrfNames.YFilter
    vrfNames.EntityData.YangName = "vrf-names"
    vrfNames.EntityData.BundleName = "cisco_ios_xr"
    vrfNames.EntityData.ParentYangName = "afi-type"
    vrfNames.EntityData.SegmentPath = "vrf-names"
    vrfNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfNames.EntityData.Children = types.NewOrderedMap()
    vrfNames.EntityData.Children.Append("vrf-name", types.YChild{"VrfName", nil})
    for i := range vrfNames.VrfName {
        vrfNames.EntityData.Children.Append(types.GetSegmentPath(vrfNames.VrfName[i]), types.YChild{"VrfName", vrfNames.VrfName[i]})
    }
    vrfNames.EntityData.Leafs = types.NewOrderedMap()

    vrfNames.EntityData.YListKeys = []string {}

    return &(vrfNames.EntityData)
}

// Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType_VrfNames_VrfName
// VRF name
type Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType_VrfNames_VrfName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // pre-ifib policer rate config commands. The type is interface{} with range:
    // 0..100000.
    AclRate interface{}
}

func (vrfName *Lpts_Ipolicer_Acls_Acl_AfiTypes_AfiType_VrfNames_VrfName) GetEntityData() *types.CommonEntityData {
    vrfName.EntityData.YFilter = vrfName.YFilter
    vrfName.EntityData.YangName = "vrf-name"
    vrfName.EntityData.BundleName = "cisco_ios_xr"
    vrfName.EntityData.ParentYangName = "vrf-names"
    vrfName.EntityData.SegmentPath = "vrf-name" + types.AddKeyToken(vrfName.VrfName, "vrf-name")
    vrfName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfName.EntityData.Children = types.NewOrderedMap()
    vrfName.EntityData.Leafs = types.NewOrderedMap()
    vrfName.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrfName.VrfName})
    vrfName.EntityData.Leafs.Append("acl-rate", types.YLeaf{"AclRate", vrfName.AclRate})

    vrfName.EntityData.YListKeys = []string {"VrfName"}

    return &(vrfName.EntityData)
}

// Lpts_Ipolicer_Flows
// Table for Flows
type Lpts_Ipolicer_Flows struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // selected flow type. The type is slice of Lpts_Ipolicer_Flows_Flow.
    Flow []*Lpts_Ipolicer_Flows_Flow
}

func (flows *Lpts_Ipolicer_Flows) GetEntityData() *types.CommonEntityData {
    flows.EntityData.YFilter = flows.YFilter
    flows.EntityData.YangName = "flows"
    flows.EntityData.BundleName = "cisco_ios_xr"
    flows.EntityData.ParentYangName = "ipolicer"
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

// Lpts_Ipolicer_Flows_Flow
// selected flow type
type Lpts_Ipolicer_Flows_Flow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Flow Type. The type is LptsFlow.
    FlowType interface{}

    // Configured rate value. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // TOS Precedence value(s).
    Precedences Lpts_Ipolicer_Flows_Flow_Precedences
}

func (flow *Lpts_Ipolicer_Flows_Flow) GetEntityData() *types.CommonEntityData {
    flow.EntityData.YFilter = flow.YFilter
    flow.EntityData.YangName = "flow"
    flow.EntityData.BundleName = "cisco_ios_xr"
    flow.EntityData.ParentYangName = "flows"
    flow.EntityData.SegmentPath = "flow" + types.AddKeyToken(flow.FlowType, "flow-type")
    flow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flow.EntityData.Children = types.NewOrderedMap()
    flow.EntityData.Children.Append("precedences", types.YChild{"Precedences", &flow.Precedences})
    flow.EntityData.Leafs = types.NewOrderedMap()
    flow.EntityData.Leafs.Append("flow-type", types.YLeaf{"FlowType", flow.FlowType})
    flow.EntityData.Leafs.Append("rate", types.YLeaf{"Rate", flow.Rate})

    flow.EntityData.YListKeys = []string {"FlowType"}

    return &(flow.EntityData)
}

// Lpts_Ipolicer_Flows_Flow_Precedences
// TOS Precedence value(s)
type Lpts_Ipolicer_Flows_Flow_Precedences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Precedence values. The type is one of the following types: slice of  
    // :go:struct:`LptsPreIFibPrecedenceNumber
    // <ydk/models/lpts_pre_ifib_cfg/LptsPreIFibPrecedenceNumber>`, or slice of
    // int with range: 0..7.
    Precedence []interface{}
}

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetEntityData() *types.CommonEntityData {
    precedences.EntityData.YFilter = precedences.YFilter
    precedences.EntityData.YangName = "precedences"
    precedences.EntityData.BundleName = "cisco_ios_xr"
    precedences.EntityData.ParentYangName = "flow"
    precedences.EntityData.SegmentPath = "precedences"
    precedences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    precedences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    precedences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    precedences.EntityData.Children = types.NewOrderedMap()
    precedences.EntityData.Leafs = types.NewOrderedMap()
    precedences.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", precedences.Precedence})

    precedences.EntityData.YListKeys = []string {}

    return &(precedences.EntityData)
}

// Lpts_Punt
// Configure penalty timeout value
type Lpts_Punt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // excessive punt flow trap configuration commands.
    Flowtrap Lpts_Punt_Flowtrap
}

func (punt *Lpts_Punt) GetEntityData() *types.CommonEntityData {
    punt.EntityData.YFilter = punt.YFilter
    punt.EntityData.YangName = "punt"
    punt.EntityData.BundleName = "cisco_ios_xr"
    punt.EntityData.ParentYangName = "lpts"
    punt.EntityData.SegmentPath = "Cisco-IOS-XR-lpts-punt-flowtrap-cfg:punt"
    punt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    punt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    punt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    punt.EntityData.Children = types.NewOrderedMap()
    punt.EntityData.Children.Append("flowtrap", types.YChild{"Flowtrap", &punt.Flowtrap})
    punt.EntityData.Leafs = types.NewOrderedMap()

    punt.EntityData.YListKeys = []string {}

    return &(punt.EntityData)
}

// Lpts_Punt_Flowtrap
// excessive punt flow trap configuration commands
type Lpts_Punt_Flowtrap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum flow gap in milliseconds. The type is interface{} with range:
    // 1..60000.
    MaxFlowGap interface{}

    // Should be power of 2. Any one of 1,2,4,8,16,32 ,64,128. The type is
    // interface{} with range: 1..128.
    EtSize interface{}

    // Eviction threshold, should be less than report-threshold. The type is
    // interface{} with range: 1..65535.
    EvictionThreshold interface{}

    // Threshold to cross for a flow to be considered as bad actor flow. The type
    // is interface{} with range: 1..65535.
    ReportThreshold interface{}

    // Enable trap based on source mac on non-subscriber interface. The type is
    // interface{} with range: 0..4294967295.
    NonSubscriberInterfaces interface{}

    // Probability of packets to be sampled. The type is string with length:
    // 1..32.
    SampleProb interface{}

    // Eviction search limit, should be less than trap-size. The type is
    // interface{} with range: 1..128.
    EvictionSearchLimit interface{}

    // Allow routing protocols to pass through copp sampler. The type is bool.
    RoutingProtocolsEnable interface{}

    // Enable the trap on subscriber interfaces. The type is bool.
    SubscriberInterfaces interface{}

    // Identify flow based on interface and flowtype. The type is bool.
    InterfaceBasedFlow interface{}

    // Dampening period for a bad actor flow in milliseconds. The type is
    // interface{} with range: 5000..60000.
    Dampening interface{}

    // Configure penalty policing rate.
    PenaltyRates Lpts_Punt_Flowtrap_PenaltyRates

    // Configure penalty timeout value.
    PenaltyTimeouts Lpts_Punt_Flowtrap_PenaltyTimeouts

    // Exclude an item from all traps.
    Exclude Lpts_Punt_Flowtrap_Exclude
}

func (flowtrap *Lpts_Punt_Flowtrap) GetEntityData() *types.CommonEntityData {
    flowtrap.EntityData.YFilter = flowtrap.YFilter
    flowtrap.EntityData.YangName = "flowtrap"
    flowtrap.EntityData.BundleName = "cisco_ios_xr"
    flowtrap.EntityData.ParentYangName = "punt"
    flowtrap.EntityData.SegmentPath = "flowtrap"
    flowtrap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowtrap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowtrap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowtrap.EntityData.Children = types.NewOrderedMap()
    flowtrap.EntityData.Children.Append("penalty-rates", types.YChild{"PenaltyRates", &flowtrap.PenaltyRates})
    flowtrap.EntityData.Children.Append("penalty-timeouts", types.YChild{"PenaltyTimeouts", &flowtrap.PenaltyTimeouts})
    flowtrap.EntityData.Children.Append("exclude", types.YChild{"Exclude", &flowtrap.Exclude})
    flowtrap.EntityData.Leafs = types.NewOrderedMap()
    flowtrap.EntityData.Leafs.Append("max-flow-gap", types.YLeaf{"MaxFlowGap", flowtrap.MaxFlowGap})
    flowtrap.EntityData.Leafs.Append("et-size", types.YLeaf{"EtSize", flowtrap.EtSize})
    flowtrap.EntityData.Leafs.Append("eviction-threshold", types.YLeaf{"EvictionThreshold", flowtrap.EvictionThreshold})
    flowtrap.EntityData.Leafs.Append("report-threshold", types.YLeaf{"ReportThreshold", flowtrap.ReportThreshold})
    flowtrap.EntityData.Leafs.Append("non-subscriber-interfaces", types.YLeaf{"NonSubscriberInterfaces", flowtrap.NonSubscriberInterfaces})
    flowtrap.EntityData.Leafs.Append("sample-prob", types.YLeaf{"SampleProb", flowtrap.SampleProb})
    flowtrap.EntityData.Leafs.Append("eviction-search-limit", types.YLeaf{"EvictionSearchLimit", flowtrap.EvictionSearchLimit})
    flowtrap.EntityData.Leafs.Append("routing-protocols-enable", types.YLeaf{"RoutingProtocolsEnable", flowtrap.RoutingProtocolsEnable})
    flowtrap.EntityData.Leafs.Append("subscriber-interfaces", types.YLeaf{"SubscriberInterfaces", flowtrap.SubscriberInterfaces})
    flowtrap.EntityData.Leafs.Append("interface-based-flow", types.YLeaf{"InterfaceBasedFlow", flowtrap.InterfaceBasedFlow})
    flowtrap.EntityData.Leafs.Append("dampening", types.YLeaf{"Dampening", flowtrap.Dampening})

    flowtrap.EntityData.YListKeys = []string {}

    return &(flowtrap.EntityData)
}

// Lpts_Punt_Flowtrap_PenaltyRates
// Configure penalty policing rate
type Lpts_Punt_Flowtrap_PenaltyRates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate.
    PenaltyRate []*Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate
}

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetEntityData() *types.CommonEntityData {
    penaltyRates.EntityData.YFilter = penaltyRates.YFilter
    penaltyRates.EntityData.YangName = "penalty-rates"
    penaltyRates.EntityData.BundleName = "cisco_ios_xr"
    penaltyRates.EntityData.ParentYangName = "flowtrap"
    penaltyRates.EntityData.SegmentPath = "penalty-rates"
    penaltyRates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    penaltyRates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    penaltyRates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    penaltyRates.EntityData.Children = types.NewOrderedMap()
    penaltyRates.EntityData.Children.Append("penalty-rate", types.YChild{"PenaltyRate", nil})
    for i := range penaltyRates.PenaltyRate {
        penaltyRates.EntityData.Children.Append(types.GetSegmentPath(penaltyRates.PenaltyRate[i]), types.YChild{"PenaltyRate", penaltyRates.PenaltyRate[i]})
    }
    penaltyRates.EntityData.Leafs = types.NewOrderedMap()

    penaltyRates.EntityData.YListKeys = []string {}

    return &(penaltyRates.EntityData)
}

// Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate
// none
type Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is LptsPuntFlowtrapProtoId.
    ProtocolName interface{}

    // Penalty policer rate in packets-per-second. The type is interface{} with
    // range: 2..100. This attribute is mandatory.
    Rate interface{}
}

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetEntityData() *types.CommonEntityData {
    penaltyRate.EntityData.YFilter = penaltyRate.YFilter
    penaltyRate.EntityData.YangName = "penalty-rate"
    penaltyRate.EntityData.BundleName = "cisco_ios_xr"
    penaltyRate.EntityData.ParentYangName = "penalty-rates"
    penaltyRate.EntityData.SegmentPath = "penalty-rate" + types.AddKeyToken(penaltyRate.ProtocolName, "protocol-name")
    penaltyRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    penaltyRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    penaltyRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    penaltyRate.EntityData.Children = types.NewOrderedMap()
    penaltyRate.EntityData.Leafs = types.NewOrderedMap()
    penaltyRate.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", penaltyRate.ProtocolName})
    penaltyRate.EntityData.Leafs.Append("rate", types.YLeaf{"Rate", penaltyRate.Rate})

    penaltyRate.EntityData.YListKeys = []string {"ProtocolName"}

    return &(penaltyRate.EntityData)
}

// Lpts_Punt_Flowtrap_PenaltyTimeouts
// Configure penalty timeout value
type Lpts_Punt_Flowtrap_PenaltyTimeouts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of
    // Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout.
    PenaltyTimeout []*Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout
}

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetEntityData() *types.CommonEntityData {
    penaltyTimeouts.EntityData.YFilter = penaltyTimeouts.YFilter
    penaltyTimeouts.EntityData.YangName = "penalty-timeouts"
    penaltyTimeouts.EntityData.BundleName = "cisco_ios_xr"
    penaltyTimeouts.EntityData.ParentYangName = "flowtrap"
    penaltyTimeouts.EntityData.SegmentPath = "penalty-timeouts"
    penaltyTimeouts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    penaltyTimeouts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    penaltyTimeouts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    penaltyTimeouts.EntityData.Children = types.NewOrderedMap()
    penaltyTimeouts.EntityData.Children.Append("penalty-timeout", types.YChild{"PenaltyTimeout", nil})
    for i := range penaltyTimeouts.PenaltyTimeout {
        penaltyTimeouts.EntityData.Children.Append(types.GetSegmentPath(penaltyTimeouts.PenaltyTimeout[i]), types.YChild{"PenaltyTimeout", penaltyTimeouts.PenaltyTimeout[i]})
    }
    penaltyTimeouts.EntityData.Leafs = types.NewOrderedMap()

    penaltyTimeouts.EntityData.YListKeys = []string {}

    return &(penaltyTimeouts.EntityData)
}

// Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout
// none
type Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is LptsPuntFlowtrapProtoId.
    ProtocolName interface{}

    // Timeout value in minutes. The type is interface{} with range: 1..1000. This
    // attribute is mandatory.
    Timeout interface{}
}

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetEntityData() *types.CommonEntityData {
    penaltyTimeout.EntityData.YFilter = penaltyTimeout.YFilter
    penaltyTimeout.EntityData.YangName = "penalty-timeout"
    penaltyTimeout.EntityData.BundleName = "cisco_ios_xr"
    penaltyTimeout.EntityData.ParentYangName = "penalty-timeouts"
    penaltyTimeout.EntityData.SegmentPath = "penalty-timeout" + types.AddKeyToken(penaltyTimeout.ProtocolName, "protocol-name")
    penaltyTimeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    penaltyTimeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    penaltyTimeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    penaltyTimeout.EntityData.Children = types.NewOrderedMap()
    penaltyTimeout.EntityData.Leafs = types.NewOrderedMap()
    penaltyTimeout.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", penaltyTimeout.ProtocolName})
    penaltyTimeout.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", penaltyTimeout.Timeout})

    penaltyTimeout.EntityData.YListKeys = []string {"ProtocolName"}

    return &(penaltyTimeout.EntityData)
}

// Lpts_Punt_Flowtrap_Exclude
// Exclude an item from all traps
type Lpts_Punt_Flowtrap_Exclude struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none.
    InterfaceNames Lpts_Punt_Flowtrap_Exclude_InterfaceNames
}

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetEntityData() *types.CommonEntityData {
    exclude.EntityData.YFilter = exclude.YFilter
    exclude.EntityData.YangName = "exclude"
    exclude.EntityData.BundleName = "cisco_ios_xr"
    exclude.EntityData.ParentYangName = "flowtrap"
    exclude.EntityData.SegmentPath = "exclude"
    exclude.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exclude.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exclude.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exclude.EntityData.Children = types.NewOrderedMap()
    exclude.EntityData.Children.Append("interface-names", types.YChild{"InterfaceNames", &exclude.InterfaceNames})
    exclude.EntityData.Leafs = types.NewOrderedMap()

    exclude.EntityData.YListKeys = []string {}

    return &(exclude.EntityData)
}

// Lpts_Punt_Flowtrap_Exclude_InterfaceNames
// none
type Lpts_Punt_Flowtrap_Exclude_InterfaceNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of interface to exclude from all traps. The type is slice of
    // Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName.
    InterfaceName []*Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName
}

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetEntityData() *types.CommonEntityData {
    interfaceNames.EntityData.YFilter = interfaceNames.YFilter
    interfaceNames.EntityData.YangName = "interface-names"
    interfaceNames.EntityData.BundleName = "cisco_ios_xr"
    interfaceNames.EntityData.ParentYangName = "exclude"
    interfaceNames.EntityData.SegmentPath = "interface-names"
    interfaceNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceNames.EntityData.Children = types.NewOrderedMap()
    interfaceNames.EntityData.Children.Append("interface-name", types.YChild{"InterfaceName", nil})
    for i := range interfaceNames.InterfaceName {
        interfaceNames.EntityData.Children.Append(types.GetSegmentPath(interfaceNames.InterfaceName[i]), types.YChild{"InterfaceName", interfaceNames.InterfaceName[i]})
    }
    interfaceNames.EntityData.Leafs = types.NewOrderedMap()

    interfaceNames.EntityData.YListKeys = []string {}

    return &(interfaceNames.EntityData)
}

// Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName
// Name of interface to exclude from all traps
type Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of interface to exclude from all traps. The
    // type is string with pattern: [a-zA-Z0-9./-]+.
    Ifname interface{}

    // Enabled or disabled. The type is bool. This attribute is mandatory.
    Id1 interface{}
}

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetEntityData() *types.CommonEntityData {
    interfaceName.EntityData.YFilter = interfaceName.YFilter
    interfaceName.EntityData.YangName = "interface-name"
    interfaceName.EntityData.BundleName = "cisco_ios_xr"
    interfaceName.EntityData.ParentYangName = "interface-names"
    interfaceName.EntityData.SegmentPath = "interface-name" + types.AddKeyToken(interfaceName.Ifname, "ifname")
    interfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceName.EntityData.Children = types.NewOrderedMap()
    interfaceName.EntityData.Leafs = types.NewOrderedMap()
    interfaceName.EntityData.Leafs.Append("ifname", types.YLeaf{"Ifname", interfaceName.Ifname})
    interfaceName.EntityData.Leafs.Append("id1", types.YLeaf{"Id1", interfaceName.Id1})

    interfaceName.EntityData.YListKeys = []string {"Ifname"}

    return &(interfaceName.EntityData)
}

