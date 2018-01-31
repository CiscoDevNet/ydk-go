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
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF Name. The type is slice of VrfPolicy_Vrf.
    Vrf []VrfPolicy_Vrf
}

func (vrfPolicy *VrfPolicy) GetFilter() yfilter.YFilter { return vrfPolicy.YFilter }

func (vrfPolicy *VrfPolicy) SetFilter(yf yfilter.YFilter) { vrfPolicy.YFilter = yf }

func (vrfPolicy *VrfPolicy) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfPolicy *VrfPolicy) GetSegmentPath() string {
    return "Cisco-IOS-XR-pbr-vrf-policy-cfg:vrf-policy"
}

func (vrfPolicy *VrfPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfPolicy.Vrf {
            if vrfPolicy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := VrfPolicy_Vrf{}
        vrfPolicy.Vrf = append(vrfPolicy.Vrf, child)
        return &vrfPolicy.Vrf[len(vrfPolicy.Vrf)-1]
    }
    return nil
}

func (vrfPolicy *VrfPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfPolicy.Vrf {
        children[vrfPolicy.Vrf[i].GetSegmentPath()] = &vrfPolicy.Vrf[i]
    }
    return children
}

func (vrfPolicy *VrfPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfPolicy *VrfPolicy) GetBundleName() string { return "cisco_ios_xr" }

func (vrfPolicy *VrfPolicy) GetYangName() string { return "vrf-policy" }

func (vrfPolicy *VrfPolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfPolicy *VrfPolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfPolicy *VrfPolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfPolicy *VrfPolicy) SetParent(parent types.Entity) { vrfPolicy.parent = parent }

func (vrfPolicy *VrfPolicy) GetParent() types.Entity { return vrfPolicy.parent }

func (vrfPolicy *VrfPolicy) GetParentYangName() string { return "Cisco-IOS-XR-pbr-vrf-policy-cfg" }

// VrfPolicy_Vrf
// VRF Name
type VrfPolicy_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // address family. The type is slice of VrfPolicy_Vrf_Afi.
    Afi []VrfPolicy_Vrf_Afi
}

func (vrf *VrfPolicy_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *VrfPolicy_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *VrfPolicy_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "afi" { return "Afi" }
    return ""
}

func (vrf *VrfPolicy_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *VrfPolicy_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afi" {
        for _, c := range vrf.Afi {
            if vrf.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := VrfPolicy_Vrf_Afi{}
        vrf.Afi = append(vrf.Afi, child)
        return &vrf.Afi[len(vrf.Afi)-1]
    }
    return nil
}

func (vrf *VrfPolicy_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrf.Afi {
        children[vrf.Afi[i].GetSegmentPath()] = &vrf.Afi[i]
    }
    return children
}

func (vrf *VrfPolicy_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *VrfPolicy_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *VrfPolicy_Vrf) GetYangName() string { return "vrf" }

func (vrf *VrfPolicy_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *VrfPolicy_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *VrfPolicy_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *VrfPolicy_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *VrfPolicy_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *VrfPolicy_Vrf) GetParentYangName() string { return "vrf-policy" }

// VrfPolicy_Vrf_Afi
// address family
type VrfPolicy_Vrf_Afi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. AFI name. The type is string with pattern: (ipv4).
    AfiType interface{}

    // Policy map name. The type is string.
    ServicePolicyIn interface{}
}

func (afi *VrfPolicy_Vrf_Afi) GetFilter() yfilter.YFilter { return afi.YFilter }

func (afi *VrfPolicy_Vrf_Afi) SetFilter(yf yfilter.YFilter) { afi.YFilter = yf }

func (afi *VrfPolicy_Vrf_Afi) GetGoName(yname string) string {
    if yname == "afi-type" { return "AfiType" }
    if yname == "service-policy-in" { return "ServicePolicyIn" }
    return ""
}

func (afi *VrfPolicy_Vrf_Afi) GetSegmentPath() string {
    return "afi" + "[afi-type='" + fmt.Sprintf("%v", afi.AfiType) + "']"
}

func (afi *VrfPolicy_Vrf_Afi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (afi *VrfPolicy_Vrf_Afi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (afi *VrfPolicy_Vrf_Afi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-type"] = afi.AfiType
    leafs["service-policy-in"] = afi.ServicePolicyIn
    return leafs
}

func (afi *VrfPolicy_Vrf_Afi) GetBundleName() string { return "cisco_ios_xr" }

func (afi *VrfPolicy_Vrf_Afi) GetYangName() string { return "afi" }

func (afi *VrfPolicy_Vrf_Afi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afi *VrfPolicy_Vrf_Afi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afi *VrfPolicy_Vrf_Afi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afi *VrfPolicy_Vrf_Afi) SetParent(parent types.Entity) { afi.parent = parent }

func (afi *VrfPolicy_Vrf_Afi) GetParent() types.Entity { return afi.parent }

func (afi *VrfPolicy_Vrf_Afi) GetParentYangName() string { return "vrf" }

