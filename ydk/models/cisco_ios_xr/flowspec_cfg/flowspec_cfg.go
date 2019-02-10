// This module contains a collection of YANG definitions
// for Cisco IOS-XR flowspec package configuration.
// 
// This module contains definitions
// for the following management objects:
//   flow-spec: FlowSpec configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package flowspec_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package flowspec_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-flowspec-cfg flow-spec}", reflect.TypeOf(FlowSpec{}))
    ydk.RegisterEntity("Cisco-IOS-XR-flowspec-cfg:flow-spec", reflect.TypeOf(FlowSpec{}))
}

// FsAddf represents Fs addf
type FsAddf string

const (
    // IPv4 FlowSpec
    FsAddf_ipv4 FsAddf = "ipv4"

    // IPv6 FlowSpec
    FsAddf_ipv6 FsAddf = "ipv6"
)

// FsVrfAf represents Fs vrf af
type FsVrfAf string

const (
    // IPv4 FlowSpec
    FsVrfAf_ipv4 FsVrfAf = "ipv4"

    // IPv6 FlowSpec
    FsVrfAf_ipv6 FsVrfAf = "ipv6"
)

// FsVrfAfP represents Fs vrf af p
type FsVrfAfP string

const (
    // PBR policy type
    FsVrfAfP_pbr FsVrfAfP = "pbr"
)

// FsAfP represents Fs af p
type FsAfP string

const (
    // PBR policy type
    FsAfP_pbr FsAfP = "pbr"
)

// FlowSpec
// FlowSpec configuration
type FlowSpec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable FlowSpec configuration. Deletion of this object also causes deletion
    // of all associated objects under FlowSpec. The type is interface{}.
    Enable interface{}

    // Install FlowSpec policy on all interfaces. The type is interface{}.
    InterfaceAll interface{}

    // Table of AF.
    Afs FlowSpec_Afs

    // Table of VRF.
    Vrfs FlowSpec_Vrfs
}

func (flowSpec *FlowSpec) GetEntityData() *types.CommonEntityData {
    flowSpec.EntityData.YFilter = flowSpec.YFilter
    flowSpec.EntityData.YangName = "flow-spec"
    flowSpec.EntityData.BundleName = "cisco_ios_xr"
    flowSpec.EntityData.ParentYangName = "Cisco-IOS-XR-flowspec-cfg"
    flowSpec.EntityData.SegmentPath = "Cisco-IOS-XR-flowspec-cfg:flow-spec"
    flowSpec.EntityData.AbsolutePath = flowSpec.EntityData.SegmentPath
    flowSpec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowSpec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowSpec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowSpec.EntityData.Children = types.NewOrderedMap()
    flowSpec.EntityData.Children.Append("afs", types.YChild{"Afs", &flowSpec.Afs})
    flowSpec.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &flowSpec.Vrfs})
    flowSpec.EntityData.Leafs = types.NewOrderedMap()
    flowSpec.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", flowSpec.Enable})
    flowSpec.EntityData.Leafs.Append("interface-all", types.YLeaf{"InterfaceAll", flowSpec.InterfaceAll})

    flowSpec.EntityData.YListKeys = []string {}

    return &(flowSpec.EntityData)
}

// FlowSpec_Afs
// Table of AF
type FlowSpec_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Family Identifier Type (IPv4/IPv6). The type is slice of
    // FlowSpec_Afs_Af.
    Af []*FlowSpec_Afs_Af
}

func (afs *FlowSpec_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "flow-spec"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.AbsolutePath = "Cisco-IOS-XR-flowspec-cfg:flow-spec/" + afs.EntityData.SegmentPath
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// FlowSpec_Afs_Af
// Address Family Identifier Type (IPv4/IPv6)
type FlowSpec_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. AFI type. The type is FsAddf.
    AfName interface{}

    // Install FlowSpec policy on all interfaces. The type is interface{}.
    InterfaceAll interface{}

    // Table of ServicePolicy.
    ServicePolicies FlowSpec_Afs_Af_ServicePolicies
}

func (af *FlowSpec_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-flowspec-cfg:flow-spec/afs/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("service-policies", types.YChild{"ServicePolicies", &af.ServicePolicies})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("interface-all", types.YLeaf{"InterfaceAll", af.InterfaceAll})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// FlowSpec_Afs_Af_ServicePolicies
// Table of ServicePolicy
type FlowSpec_Afs_Af_ServicePolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service Policy configuration. The type is slice of
    // FlowSpec_Afs_Af_ServicePolicies_ServicePolicy.
    ServicePolicy []*FlowSpec_Afs_Af_ServicePolicies_ServicePolicy
}

func (servicePolicies *FlowSpec_Afs_Af_ServicePolicies) GetEntityData() *types.CommonEntityData {
    servicePolicies.EntityData.YFilter = servicePolicies.YFilter
    servicePolicies.EntityData.YangName = "service-policies"
    servicePolicies.EntityData.BundleName = "cisco_ios_xr"
    servicePolicies.EntityData.ParentYangName = "af"
    servicePolicies.EntityData.SegmentPath = "service-policies"
    servicePolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-flowspec-cfg:flow-spec/afs/af/" + servicePolicies.EntityData.SegmentPath
    servicePolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicies.EntityData.Children = types.NewOrderedMap()
    servicePolicies.EntityData.Children.Append("service-policy", types.YChild{"ServicePolicy", nil})
    for i := range servicePolicies.ServicePolicy {
        servicePolicies.EntityData.Children.Append(types.GetSegmentPath(servicePolicies.ServicePolicy[i]), types.YChild{"ServicePolicy", servicePolicies.ServicePolicy[i]})
    }
    servicePolicies.EntityData.Leafs = types.NewOrderedMap()

    servicePolicies.EntityData.YListKeys = []string {}

    return &(servicePolicies.EntityData)
}

// FlowSpec_Afs_Af_ServicePolicies_ServicePolicy
// Service Policy configuration
type FlowSpec_Afs_Af_ServicePolicies_ServicePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Policy map name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    PolicyName interface{}

    // keys: policy-type. The type is slice of
    // FlowSpec_Afs_Af_ServicePolicies_ServicePolicy_PolicyType.
    PolicyType []*FlowSpec_Afs_Af_ServicePolicies_ServicePolicy_PolicyType
}

func (servicePolicy *FlowSpec_Afs_Af_ServicePolicies_ServicePolicy) GetEntityData() *types.CommonEntityData {
    servicePolicy.EntityData.YFilter = servicePolicy.YFilter
    servicePolicy.EntityData.YangName = "service-policy"
    servicePolicy.EntityData.BundleName = "cisco_ios_xr"
    servicePolicy.EntityData.ParentYangName = "service-policies"
    servicePolicy.EntityData.SegmentPath = "service-policy" + types.AddKeyToken(servicePolicy.PolicyName, "policy-name")
    servicePolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-flowspec-cfg:flow-spec/afs/af/service-policies/" + servicePolicy.EntityData.SegmentPath
    servicePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicy.EntityData.Children = types.NewOrderedMap()
    servicePolicy.EntityData.Children.Append("policy-type", types.YChild{"PolicyType", nil})
    for i := range servicePolicy.PolicyType {
        servicePolicy.EntityData.Children.Append(types.GetSegmentPath(servicePolicy.PolicyType[i]), types.YChild{"PolicyType", servicePolicy.PolicyType[i]})
    }
    servicePolicy.EntityData.Leafs = types.NewOrderedMap()
    servicePolicy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", servicePolicy.PolicyName})

    servicePolicy.EntityData.YListKeys = []string {"PolicyName"}

    return &(servicePolicy.EntityData)
}

// FlowSpec_Afs_Af_ServicePolicies_ServicePolicy_PolicyType
// keys: policy-type
type FlowSpec_Afs_Af_ServicePolicies_ServicePolicy_PolicyType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Choose the Policy type. The type is FsAfP.
    PolicyType interface{}

    // Set constant integer. The type is bool. This attribute is mandatory.
    Local interface{}
}

func (policyType *FlowSpec_Afs_Af_ServicePolicies_ServicePolicy_PolicyType) GetEntityData() *types.CommonEntityData {
    policyType.EntityData.YFilter = policyType.YFilter
    policyType.EntityData.YangName = "policy-type"
    policyType.EntityData.BundleName = "cisco_ios_xr"
    policyType.EntityData.ParentYangName = "service-policy"
    policyType.EntityData.SegmentPath = "policy-type" + types.AddKeyToken(policyType.PolicyType, "policy-type")
    policyType.EntityData.AbsolutePath = "Cisco-IOS-XR-flowspec-cfg:flow-spec/afs/af/service-policies/service-policy/" + policyType.EntityData.SegmentPath
    policyType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyType.EntityData.Children = types.NewOrderedMap()
    policyType.EntityData.Leafs = types.NewOrderedMap()
    policyType.EntityData.Leafs.Append("policy-type", types.YLeaf{"PolicyType", policyType.PolicyType})
    policyType.EntityData.Leafs.Append("local", types.YLeaf{"Local", policyType.Local})

    policyType.EntityData.YListKeys = []string {"PolicyType"}

    return &(policyType.EntityData)
}

// FlowSpec_Vrfs
// Table of VRF
type FlowSpec_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF configuration. The type is slice of FlowSpec_Vrfs_Vrf.
    Vrf []*FlowSpec_Vrfs_Vrf
}

func (vrfs *FlowSpec_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "flow-spec"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-flowspec-cfg:flow-spec/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// FlowSpec_Vrfs_Vrf
// VRF configuration
type FlowSpec_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF Name. The type is string with length: 1..32.
    VrfName interface{}

    // Table of AF.
    Afs FlowSpec_Vrfs_Vrf_Afs
}

func (vrf *FlowSpec_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-flowspec-cfg:flow-spec/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("afs", types.YChild{"Afs", &vrf.Afs})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs
// Table of AF
type FlowSpec_Vrfs_Vrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Family Identifier Type (IPv4/IPv6). The type is slice of
    // FlowSpec_Vrfs_Vrf_Afs_Af.
    Af []*FlowSpec_Vrfs_Vrf_Afs_Af
}

func (afs *FlowSpec_Vrfs_Vrf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "vrf"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.AbsolutePath = "Cisco-IOS-XR-flowspec-cfg:flow-spec/vrfs/vrf/" + afs.EntityData.SegmentPath
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af
// Address Family Identifier Type (IPv4/IPv6)
type FlowSpec_Vrfs_Vrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. AFI type. The type is FsVrfAf.
    AfName interface{}

    // Install FlowSpec policy on all interfaces. The type is interface{}.
    InterfaceAll interface{}

    // Table of ServicePolicy.
    ServicePolicies FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies
}

func (af *FlowSpec_Vrfs_Vrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-flowspec-cfg:flow-spec/vrfs/vrf/afs/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("service-policies", types.YChild{"ServicePolicies", &af.ServicePolicies})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("interface-all", types.YLeaf{"InterfaceAll", af.InterfaceAll})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies
// Table of ServicePolicy
type FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service Policy configuration. The type is slice of
    // FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies_ServicePolicy.
    ServicePolicy []*FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies_ServicePolicy
}

func (servicePolicies *FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies) GetEntityData() *types.CommonEntityData {
    servicePolicies.EntityData.YFilter = servicePolicies.YFilter
    servicePolicies.EntityData.YangName = "service-policies"
    servicePolicies.EntityData.BundleName = "cisco_ios_xr"
    servicePolicies.EntityData.ParentYangName = "af"
    servicePolicies.EntityData.SegmentPath = "service-policies"
    servicePolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-flowspec-cfg:flow-spec/vrfs/vrf/afs/af/" + servicePolicies.EntityData.SegmentPath
    servicePolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicies.EntityData.Children = types.NewOrderedMap()
    servicePolicies.EntityData.Children.Append("service-policy", types.YChild{"ServicePolicy", nil})
    for i := range servicePolicies.ServicePolicy {
        servicePolicies.EntityData.Children.Append(types.GetSegmentPath(servicePolicies.ServicePolicy[i]), types.YChild{"ServicePolicy", servicePolicies.ServicePolicy[i]})
    }
    servicePolicies.EntityData.Leafs = types.NewOrderedMap()

    servicePolicies.EntityData.YListKeys = []string {}

    return &(servicePolicies.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies_ServicePolicy
// Service Policy configuration
type FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies_ServicePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Policy map name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    PolicyName interface{}

    // keys: policy-type. The type is slice of
    // FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies_ServicePolicy_PolicyType.
    PolicyType []*FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies_ServicePolicy_PolicyType
}

func (servicePolicy *FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies_ServicePolicy) GetEntityData() *types.CommonEntityData {
    servicePolicy.EntityData.YFilter = servicePolicy.YFilter
    servicePolicy.EntityData.YangName = "service-policy"
    servicePolicy.EntityData.BundleName = "cisco_ios_xr"
    servicePolicy.EntityData.ParentYangName = "service-policies"
    servicePolicy.EntityData.SegmentPath = "service-policy" + types.AddKeyToken(servicePolicy.PolicyName, "policy-name")
    servicePolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-flowspec-cfg:flow-spec/vrfs/vrf/afs/af/service-policies/" + servicePolicy.EntityData.SegmentPath
    servicePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicy.EntityData.Children = types.NewOrderedMap()
    servicePolicy.EntityData.Children.Append("policy-type", types.YChild{"PolicyType", nil})
    for i := range servicePolicy.PolicyType {
        servicePolicy.EntityData.Children.Append(types.GetSegmentPath(servicePolicy.PolicyType[i]), types.YChild{"PolicyType", servicePolicy.PolicyType[i]})
    }
    servicePolicy.EntityData.Leafs = types.NewOrderedMap()
    servicePolicy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", servicePolicy.PolicyName})

    servicePolicy.EntityData.YListKeys = []string {"PolicyName"}

    return &(servicePolicy.EntityData)
}

// FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies_ServicePolicy_PolicyType
// keys: policy-type
type FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies_ServicePolicy_PolicyType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Choose the Policy type. The type is FsVrfAfP.
    PolicyType interface{}

    // Set constant integer. The type is bool. This attribute is mandatory.
    Local interface{}
}

func (policyType *FlowSpec_Vrfs_Vrf_Afs_Af_ServicePolicies_ServicePolicy_PolicyType) GetEntityData() *types.CommonEntityData {
    policyType.EntityData.YFilter = policyType.YFilter
    policyType.EntityData.YangName = "policy-type"
    policyType.EntityData.BundleName = "cisco_ios_xr"
    policyType.EntityData.ParentYangName = "service-policy"
    policyType.EntityData.SegmentPath = "policy-type" + types.AddKeyToken(policyType.PolicyType, "policy-type")
    policyType.EntityData.AbsolutePath = "Cisco-IOS-XR-flowspec-cfg:flow-spec/vrfs/vrf/afs/af/service-policies/service-policy/" + policyType.EntityData.SegmentPath
    policyType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyType.EntityData.Children = types.NewOrderedMap()
    policyType.EntityData.Leafs = types.NewOrderedMap()
    policyType.EntityData.Leafs.Append("policy-type", types.YLeaf{"PolicyType", policyType.PolicyType})
    policyType.EntityData.Leafs.Append("local", types.YLeaf{"Local", policyType.Local})

    policyType.EntityData.YListKeys = []string {"PolicyType"}

    return &(policyType.EntityData)
}

