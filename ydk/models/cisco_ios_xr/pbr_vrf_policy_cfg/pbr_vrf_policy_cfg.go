// This module contains a collection of YANG definitions
// for Cisco IOS-XR pbr-vrf-policy package configuration.
// 
// This module contains definitions
// for the following management objects:
//   vrf-policy: VRF Policy PBR configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package pbr_vrf_policy_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package pbr_vrf_policy_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-pbr-vrf-policy-cfg vrf-policy}", reflect.TypeOf(VrfPolicy{}))
    ydk.RegisterEntity("Cisco-IOS-XR-pbr-vrf-policy-cfg:vrf-policy", reflect.TypeOf(VrfPolicy{}))
}

// VrfPolicy
// VRF Policy PBR configuration
type VrfPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF Name. The type is slice of VrfPolicy_Vrf.
    Vrf []*VrfPolicy_Vrf
}

func (vrfPolicy *VrfPolicy) GetEntityData() *types.CommonEntityData {
    vrfPolicy.EntityData.YFilter = vrfPolicy.YFilter
    vrfPolicy.EntityData.YangName = "vrf-policy"
    vrfPolicy.EntityData.BundleName = "cisco_ios_xr"
    vrfPolicy.EntityData.ParentYangName = "Cisco-IOS-XR-pbr-vrf-policy-cfg"
    vrfPolicy.EntityData.SegmentPath = "Cisco-IOS-XR-pbr-vrf-policy-cfg:vrf-policy"
    vrfPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfPolicy.EntityData.Children = types.NewOrderedMap()
    vrfPolicy.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfPolicy.Vrf {
        vrfPolicy.EntityData.Children.Append(types.GetSegmentPath(vrfPolicy.Vrf[i]), types.YChild{"Vrf", vrfPolicy.Vrf[i]})
    }
    vrfPolicy.EntityData.Leafs = types.NewOrderedMap()

    vrfPolicy.EntityData.YListKeys = []string {}

    return &(vrfPolicy.EntityData)
}

// VrfPolicy_Vrf
// VRF Name
type VrfPolicy_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // address family. The type is slice of VrfPolicy_Vrf_Afi.
    Afi []*VrfPolicy_Vrf_Afi
}

func (vrf *VrfPolicy_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrf-policy"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("afi", types.YChild{"Afi", nil})
    for i := range vrf.Afi {
        vrf.EntityData.Children.Append(types.GetSegmentPath(vrf.Afi[i]), types.YChild{"Afi", vrf.Afi[i]})
    }
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// VrfPolicy_Vrf_Afi
// address family
type VrfPolicy_Vrf_Afi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. AFI name. The type is string with pattern: (ipv4).
    AfiType interface{}

    // Policy map name. The type is string.
    ServicePolicyIn interface{}
}

func (afi *VrfPolicy_Vrf_Afi) GetEntityData() *types.CommonEntityData {
    afi.EntityData.YFilter = afi.YFilter
    afi.EntityData.YangName = "afi"
    afi.EntityData.BundleName = "cisco_ios_xr"
    afi.EntityData.ParentYangName = "vrf"
    afi.EntityData.SegmentPath = "afi" + types.AddKeyToken(afi.AfiType, "afi-type")
    afi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afi.EntityData.Children = types.NewOrderedMap()
    afi.EntityData.Leafs = types.NewOrderedMap()
    afi.EntityData.Leafs.Append("afi-type", types.YLeaf{"AfiType", afi.AfiType})
    afi.EntityData.Leafs.Append("service-policy-in", types.YLeaf{"ServicePolicyIn", afi.ServicePolicyIn})

    afi.EntityData.YListKeys = []string {"AfiType"}

    return &(afi.EntityData)
}

