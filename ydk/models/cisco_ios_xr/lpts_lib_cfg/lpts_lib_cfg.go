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

    lpts.EntityData.Children = make(map[string]types.YChild)
    lpts.EntityData.Children["Cisco-IOS-XR-lpts-pre-ifib-cfg:ipolicer"] = types.YChild{"Ipolicer", &lpts.Ipolicer}
    lpts.EntityData.Children["Cisco-IOS-XR-lpts-punt-flowtrap-cfg:punt"] = types.YChild{"Punt", &lpts.Punt}
    lpts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lpts.EntityData)
}

// Lpts_Ipolicer
// Pre IFiB Configuration 
// This type is a presence type.
type Lpts_Ipolicer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enabled. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Table for ACLs.
    Ipv4Acls Lpts_Ipolicer_Ipv4Acls

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

    ipolicer.EntityData.Children = make(map[string]types.YChild)
    ipolicer.EntityData.Children["ipv4acls"] = types.YChild{"Ipv4Acls", &ipolicer.Ipv4Acls}
    ipolicer.EntityData.Children["flows"] = types.YChild{"Flows", &ipolicer.Flows}
    ipolicer.EntityData.Leafs = make(map[string]types.YLeaf)
    ipolicer.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ipolicer.Enable}
    return &(ipolicer.EntityData)
}

// Lpts_Ipolicer_Ipv4Acls
// Table for ACLs
type Lpts_Ipolicer_Ipv4Acls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ACL name. The type is slice of Lpts_Ipolicer_Ipv4Acls_Ipv4Acl.
    Ipv4Acl []Lpts_Ipolicer_Ipv4Acls_Ipv4Acl
}

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetEntityData() *types.CommonEntityData {
    ipv4Acls.EntityData.YFilter = ipv4Acls.YFilter
    ipv4Acls.EntityData.YangName = "ipv4acls"
    ipv4Acls.EntityData.BundleName = "cisco_ios_xr"
    ipv4Acls.EntityData.ParentYangName = "ipolicer"
    ipv4Acls.EntityData.SegmentPath = "ipv4acls"
    ipv4Acls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Acls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Acls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Acls.EntityData.Children = make(map[string]types.YChild)
    ipv4Acls.EntityData.Children["ipv4acl"] = types.YChild{"Ipv4Acl", nil}
    for i := range ipv4Acls.Ipv4Acl {
        ipv4Acls.EntityData.Children[types.GetSegmentPath(&ipv4Acls.Ipv4Acl[i])] = types.YChild{"Ipv4Acl", &ipv4Acls.Ipv4Acl[i]}
    }
    ipv4Acls.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4Acls.EntityData)
}

// Lpts_Ipolicer_Ipv4Acls_Ipv4Acl
// ACL name
type Lpts_Ipolicer_Ipv4Acls_Ipv4Acl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. ACL name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    AclName interface{}

    // VRF list.
    Ipv4VrfNames Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames
}

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetEntityData() *types.CommonEntityData {
    ipv4Acl.EntityData.YFilter = ipv4Acl.YFilter
    ipv4Acl.EntityData.YangName = "ipv4acl"
    ipv4Acl.EntityData.BundleName = "cisco_ios_xr"
    ipv4Acl.EntityData.ParentYangName = "ipv4acls"
    ipv4Acl.EntityData.SegmentPath = "ipv4acl" + "[acl-name='" + fmt.Sprintf("%v", ipv4Acl.AclName) + "']"
    ipv4Acl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Acl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Acl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Acl.EntityData.Children = make(map[string]types.YChild)
    ipv4Acl.EntityData.Children["ipv4vrf-names"] = types.YChild{"Ipv4VrfNames", &ipv4Acl.Ipv4VrfNames}
    ipv4Acl.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4Acl.EntityData.Leafs["acl-name"] = types.YLeaf{"AclName", ipv4Acl.AclName}
    return &(ipv4Acl.EntityData)
}

// Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames
// VRF list
type Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is slice of
    // Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName.
    Ipv4VrfName []Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName
}

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetEntityData() *types.CommonEntityData {
    ipv4VrfNames.EntityData.YFilter = ipv4VrfNames.YFilter
    ipv4VrfNames.EntityData.YangName = "ipv4vrf-names"
    ipv4VrfNames.EntityData.BundleName = "cisco_ios_xr"
    ipv4VrfNames.EntityData.ParentYangName = "ipv4acl"
    ipv4VrfNames.EntityData.SegmentPath = "ipv4vrf-names"
    ipv4VrfNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4VrfNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4VrfNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4VrfNames.EntityData.Children = make(map[string]types.YChild)
    ipv4VrfNames.EntityData.Children["ipv4vrf-name"] = types.YChild{"Ipv4VrfName", nil}
    for i := range ipv4VrfNames.Ipv4VrfName {
        ipv4VrfNames.EntityData.Children[types.GetSegmentPath(&ipv4VrfNames.Ipv4VrfName[i])] = types.YChild{"Ipv4VrfName", &ipv4VrfNames.Ipv4VrfName[i]}
    }
    ipv4VrfNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4VrfNames.EntityData)
}

// Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName
// VRF name
type Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // pre-ifib policer rate config commands. The type is interface{} with range:
    // 0..100000.
    AclRate interface{}
}

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetEntityData() *types.CommonEntityData {
    ipv4VrfName.EntityData.YFilter = ipv4VrfName.YFilter
    ipv4VrfName.EntityData.YangName = "ipv4vrf-name"
    ipv4VrfName.EntityData.BundleName = "cisco_ios_xr"
    ipv4VrfName.EntityData.ParentYangName = "ipv4vrf-names"
    ipv4VrfName.EntityData.SegmentPath = "ipv4vrf-name" + "[vrf-name='" + fmt.Sprintf("%v", ipv4VrfName.VrfName) + "']"
    ipv4VrfName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4VrfName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4VrfName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4VrfName.EntityData.Children = make(map[string]types.YChild)
    ipv4VrfName.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4VrfName.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv4VrfName.VrfName}
    ipv4VrfName.EntityData.Leafs["acl-rate"] = types.YLeaf{"AclRate", ipv4VrfName.AclRate}
    return &(ipv4VrfName.EntityData)
}

// Lpts_Ipolicer_Flows
// Table for Flows
type Lpts_Ipolicer_Flows struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // selected flow type. The type is slice of Lpts_Ipolicer_Flows_Flow.
    Flow []Lpts_Ipolicer_Flows_Flow
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

    flows.EntityData.Children = make(map[string]types.YChild)
    flows.EntityData.Children["flow"] = types.YChild{"Flow", nil}
    for i := range flows.Flow {
        flows.EntityData.Children[types.GetSegmentPath(&flows.Flow[i])] = types.YChild{"Flow", &flows.Flow[i]}
    }
    flows.EntityData.Leafs = make(map[string]types.YLeaf)
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
    flow.EntityData.SegmentPath = "flow" + "[flow-type='" + fmt.Sprintf("%v", flow.FlowType) + "']"
    flow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flow.EntityData.Children = make(map[string]types.YChild)
    flow.EntityData.Children["precedences"] = types.YChild{"Precedences", &flow.Precedences}
    flow.EntityData.Leafs = make(map[string]types.YLeaf)
    flow.EntityData.Leafs["flow-type"] = types.YLeaf{"FlowType", flow.FlowType}
    flow.EntityData.Leafs["rate"] = types.YLeaf{"Rate", flow.Rate}
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

    precedences.EntityData.Children = make(map[string]types.YChild)
    precedences.EntityData.Leafs = make(map[string]types.YLeaf)
    precedences.EntityData.Leafs["precedence"] = types.YLeaf{"Precedence", precedences.Precedence}
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

    punt.EntityData.Children = make(map[string]types.YChild)
    punt.EntityData.Children["flowtrap"] = types.YChild{"Flowtrap", &punt.Flowtrap}
    punt.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // interface{} with range: -2147483648..2147483647.
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

    flowtrap.EntityData.Children = make(map[string]types.YChild)
    flowtrap.EntityData.Children["penalty-rates"] = types.YChild{"PenaltyRates", &flowtrap.PenaltyRates}
    flowtrap.EntityData.Children["penalty-timeouts"] = types.YChild{"PenaltyTimeouts", &flowtrap.PenaltyTimeouts}
    flowtrap.EntityData.Children["exclude"] = types.YChild{"Exclude", &flowtrap.Exclude}
    flowtrap.EntityData.Leafs = make(map[string]types.YLeaf)
    flowtrap.EntityData.Leafs["max-flow-gap"] = types.YLeaf{"MaxFlowGap", flowtrap.MaxFlowGap}
    flowtrap.EntityData.Leafs["et-size"] = types.YLeaf{"EtSize", flowtrap.EtSize}
    flowtrap.EntityData.Leafs["eviction-threshold"] = types.YLeaf{"EvictionThreshold", flowtrap.EvictionThreshold}
    flowtrap.EntityData.Leafs["report-threshold"] = types.YLeaf{"ReportThreshold", flowtrap.ReportThreshold}
    flowtrap.EntityData.Leafs["non-subscriber-interfaces"] = types.YLeaf{"NonSubscriberInterfaces", flowtrap.NonSubscriberInterfaces}
    flowtrap.EntityData.Leafs["sample-prob"] = types.YLeaf{"SampleProb", flowtrap.SampleProb}
    flowtrap.EntityData.Leafs["eviction-search-limit"] = types.YLeaf{"EvictionSearchLimit", flowtrap.EvictionSearchLimit}
    flowtrap.EntityData.Leafs["routing-protocols-enable"] = types.YLeaf{"RoutingProtocolsEnable", flowtrap.RoutingProtocolsEnable}
    flowtrap.EntityData.Leafs["subscriber-interfaces"] = types.YLeaf{"SubscriberInterfaces", flowtrap.SubscriberInterfaces}
    flowtrap.EntityData.Leafs["interface-based-flow"] = types.YLeaf{"InterfaceBasedFlow", flowtrap.InterfaceBasedFlow}
    flowtrap.EntityData.Leafs["dampening"] = types.YLeaf{"Dampening", flowtrap.Dampening}
    return &(flowtrap.EntityData)
}

// Lpts_Punt_Flowtrap_PenaltyRates
// Configure penalty policing rate
type Lpts_Punt_Flowtrap_PenaltyRates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate.
    PenaltyRate []Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate
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

    penaltyRates.EntityData.Children = make(map[string]types.YChild)
    penaltyRates.EntityData.Children["penalty-rate"] = types.YChild{"PenaltyRate", nil}
    for i := range penaltyRates.PenaltyRate {
        penaltyRates.EntityData.Children[types.GetSegmentPath(&penaltyRates.PenaltyRate[i])] = types.YChild{"PenaltyRate", &penaltyRates.PenaltyRate[i]}
    }
    penaltyRates.EntityData.Leafs = make(map[string]types.YLeaf)
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
    penaltyRate.EntityData.SegmentPath = "penalty-rate" + "[protocol-name='" + fmt.Sprintf("%v", penaltyRate.ProtocolName) + "']"
    penaltyRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    penaltyRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    penaltyRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    penaltyRate.EntityData.Children = make(map[string]types.YChild)
    penaltyRate.EntityData.Leafs = make(map[string]types.YLeaf)
    penaltyRate.EntityData.Leafs["protocol-name"] = types.YLeaf{"ProtocolName", penaltyRate.ProtocolName}
    penaltyRate.EntityData.Leafs["rate"] = types.YLeaf{"Rate", penaltyRate.Rate}
    return &(penaltyRate.EntityData)
}

// Lpts_Punt_Flowtrap_PenaltyTimeouts
// Configure penalty timeout value
type Lpts_Punt_Flowtrap_PenaltyTimeouts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of
    // Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout.
    PenaltyTimeout []Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout
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

    penaltyTimeouts.EntityData.Children = make(map[string]types.YChild)
    penaltyTimeouts.EntityData.Children["penalty-timeout"] = types.YChild{"PenaltyTimeout", nil}
    for i := range penaltyTimeouts.PenaltyTimeout {
        penaltyTimeouts.EntityData.Children[types.GetSegmentPath(&penaltyTimeouts.PenaltyTimeout[i])] = types.YChild{"PenaltyTimeout", &penaltyTimeouts.PenaltyTimeout[i]}
    }
    penaltyTimeouts.EntityData.Leafs = make(map[string]types.YLeaf)
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
    penaltyTimeout.EntityData.SegmentPath = "penalty-timeout" + "[protocol-name='" + fmt.Sprintf("%v", penaltyTimeout.ProtocolName) + "']"
    penaltyTimeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    penaltyTimeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    penaltyTimeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    penaltyTimeout.EntityData.Children = make(map[string]types.YChild)
    penaltyTimeout.EntityData.Leafs = make(map[string]types.YLeaf)
    penaltyTimeout.EntityData.Leafs["protocol-name"] = types.YLeaf{"ProtocolName", penaltyTimeout.ProtocolName}
    penaltyTimeout.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", penaltyTimeout.Timeout}
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

    exclude.EntityData.Children = make(map[string]types.YChild)
    exclude.EntityData.Children["interface-names"] = types.YChild{"InterfaceNames", &exclude.InterfaceNames}
    exclude.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(exclude.EntityData)
}

// Lpts_Punt_Flowtrap_Exclude_InterfaceNames
// none
type Lpts_Punt_Flowtrap_Exclude_InterfaceNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of interface to exclude from all traps. The type is slice of
    // Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName.
    InterfaceName []Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName
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

    interfaceNames.EntityData.Children = make(map[string]types.YChild)
    interfaceNames.EntityData.Children["interface-name"] = types.YChild{"InterfaceName", nil}
    for i := range interfaceNames.InterfaceName {
        interfaceNames.EntityData.Children[types.GetSegmentPath(&interfaceNames.InterfaceName[i])] = types.YChild{"InterfaceName", &interfaceNames.InterfaceName[i]}
    }
    interfaceNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceNames.EntityData)
}

// Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName
// Name of interface to exclude from all traps
type Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of interface to exclude from all traps. The
    // type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Ifname interface{}

    // Enabled or disabled. The type is bool. This attribute is mandatory.
    Id1 interface{}
}

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetEntityData() *types.CommonEntityData {
    interfaceName.EntityData.YFilter = interfaceName.YFilter
    interfaceName.EntityData.YangName = "interface-name"
    interfaceName.EntityData.BundleName = "cisco_ios_xr"
    interfaceName.EntityData.ParentYangName = "interface-names"
    interfaceName.EntityData.SegmentPath = "interface-name" + "[ifname='" + fmt.Sprintf("%v", interfaceName.Ifname) + "']"
    interfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceName.EntityData.Children = make(map[string]types.YChild)
    interfaceName.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceName.EntityData.Leafs["ifname"] = types.YLeaf{"Ifname", interfaceName.Ifname}
    interfaceName.EntityData.Leafs["id1"] = types.YLeaf{"Id1", interfaceName.Id1}
    return &(interfaceName.EntityData)
}

