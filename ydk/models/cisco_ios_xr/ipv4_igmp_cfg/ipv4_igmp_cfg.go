// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-igmp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   igmp: IGMPconfiguration
//   amt: amt
//   mld: mld
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_igmp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_igmp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-igmp-cfg igmp}", reflect.TypeOf(Igmp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-igmp-cfg:igmp", reflect.TypeOf(Igmp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-igmp-cfg amt}", reflect.TypeOf(Amt{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-igmp-cfg:amt", reflect.TypeOf(Amt{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-igmp-cfg mld}", reflect.TypeOf(Mld{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-igmp-cfg:mld", reflect.TypeOf(Mld{}))
}

// IgmpFilter represents Igmp filter
type IgmpFilter string

const (
    // Include
    IgmpFilter_include IgmpFilter = "include"

    // Exclude
    IgmpFilter_exclude IgmpFilter = "exclude"

    // StarG
    IgmpFilter_star_g IgmpFilter = "star-g"
)

// Igmp
// IGMPconfiguration
// This type is a presence type.
type Igmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF related configuration.
    Vrfs Igmp_Vrfs

    // Default Context.
    DefaultContext Igmp_DefaultContext
}

func (igmp *Igmp) GetFilter() yfilter.YFilter { return igmp.YFilter }

func (igmp *Igmp) SetFilter(yf yfilter.YFilter) { igmp.YFilter = yf }

func (igmp *Igmp) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    if yname == "default-context" { return "DefaultContext" }
    return ""
}

func (igmp *Igmp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-igmp-cfg:igmp"
}

func (igmp *Igmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &igmp.Vrfs
    }
    if childYangName == "default-context" {
        return &igmp.DefaultContext
    }
    return nil
}

func (igmp *Igmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &igmp.Vrfs
    children["default-context"] = &igmp.DefaultContext
    return children
}

func (igmp *Igmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (igmp *Igmp) GetBundleName() string { return "cisco_ios_xr" }

func (igmp *Igmp) GetYangName() string { return "igmp" }

func (igmp *Igmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igmp *Igmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igmp *Igmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igmp *Igmp) SetParent(parent types.Entity) { igmp.parent = parent }

func (igmp *Igmp) GetParent() types.Entity { return igmp.parent }

func (igmp *Igmp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-igmp-cfg" }

// Igmp_Vrfs
// VRF related configuration
type Igmp_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a particular vrf. The type is slice of Igmp_Vrfs_Vrf.
    Vrf []Igmp_Vrfs_Vrf
}

func (vrfs *Igmp_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Igmp_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Igmp_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Igmp_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Igmp_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Igmp_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Igmp_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Igmp_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Igmp_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Igmp_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Igmp_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Igmp_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Igmp_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Igmp_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Igmp_Vrfs) GetParentYangName() string { return "igmp" }

// Igmp_Vrfs_Vrf
// Configuration for a particular vrf
type Igmp_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name for this vrf. The type is string with length:
    // 1..32.
    VrfName interface{}

    // Enable SSM mapping using DNS Query. The type is interface{}.
    SsmdnsQueryGroup interface{}

    // Configure IGMP Robustness variable. The type is interface{} with range:
    // 2..10. The default value is 2.
    Robustness interface{}

    // Configure IGMP Traffic variables.
    Traffic Igmp_Vrfs_Vrf_Traffic

    // Inheritable Defaults.
    InheritableDefaults Igmp_Vrfs_Vrf_InheritableDefaults

    // IGMP Source specific mode.
    SsmAccessGroups Igmp_Vrfs_Vrf_SsmAccessGroups

    // Configure IGMP State Limits.
    Maximum Igmp_Vrfs_Vrf_Maximum

    // Interface-level configuration.
    Interfaces Igmp_Vrfs_Vrf_Interfaces
}

func (vrf *Igmp_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Igmp_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Igmp_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ssmdns-query-group" { return "SsmdnsQueryGroup" }
    if yname == "robustness" { return "Robustness" }
    if yname == "traffic" { return "Traffic" }
    if yname == "inheritable-defaults" { return "InheritableDefaults" }
    if yname == "ssm-access-groups" { return "SsmAccessGroups" }
    if yname == "maximum" { return "Maximum" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (vrf *Igmp_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Igmp_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traffic" {
        return &vrf.Traffic
    }
    if childYangName == "inheritable-defaults" {
        return &vrf.InheritableDefaults
    }
    if childYangName == "ssm-access-groups" {
        return &vrf.SsmAccessGroups
    }
    if childYangName == "maximum" {
        return &vrf.Maximum
    }
    if childYangName == "interfaces" {
        return &vrf.Interfaces
    }
    return nil
}

func (vrf *Igmp_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["traffic"] = &vrf.Traffic
    children["inheritable-defaults"] = &vrf.InheritableDefaults
    children["ssm-access-groups"] = &vrf.SsmAccessGroups
    children["maximum"] = &vrf.Maximum
    children["interfaces"] = &vrf.Interfaces
    return children
}

func (vrf *Igmp_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["ssmdns-query-group"] = vrf.SsmdnsQueryGroup
    leafs["robustness"] = vrf.Robustness
    return leafs
}

func (vrf *Igmp_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Igmp_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Igmp_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Igmp_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Igmp_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Igmp_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Igmp_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Igmp_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Igmp_Vrfs_Vrf_Traffic
// Configure IGMP Traffic variables
type Igmp_Vrfs_Vrf_Traffic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the route-policy profile. The type is string with length: 1..64.
    Profile interface{}
}

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetFilter() yfilter.YFilter { return traffic.YFilter }

func (traffic *Igmp_Vrfs_Vrf_Traffic) SetFilter(yf yfilter.YFilter) { traffic.YFilter = yf }

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetSegmentPath() string {
    return "traffic"
}

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile"] = traffic.Profile
    return leafs
}

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetYangName() string { return "traffic" }

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traffic *Igmp_Vrfs_Vrf_Traffic) SetParent(parent types.Entity) { traffic.parent = parent }

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetParent() types.Entity { return traffic.parent }

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetParentYangName() string { return "vrf" }

// Igmp_Vrfs_Vrf_InheritableDefaults
// Inheritable Defaults
type Igmp_Vrfs_Vrf_InheritableDefaults struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGMP previous querier timeout. The type is interface{} with range: 60..300.
    // Units are second.
    QueryTimeout interface{}

    // Access list specifying access group range. The type is string with length:
    // 1..64.
    AccessGroup interface{}

    // Query response value in seconds. The type is interface{} with range: 1..12.
    // Units are second. The default value is 10.
    QueryMaxResponseTime interface{}

    // Version number. The type is interface{} with range: 1..3. The default value
    // is 3.
    Version interface{}

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. The default value is true.
    RouterEnable interface{}

    // Query interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second. The default value is 60.
    QueryInterval interface{}

    // Configure maximum number of groups accepted per interface by this router.
    MaximumGroupsPerInterfaceOor Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor

    // IGMPv3 explicit host tracking.
    ExplicitTracking Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking
}

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetFilter() yfilter.YFilter { return inheritableDefaults.YFilter }

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) SetFilter(yf yfilter.YFilter) { inheritableDefaults.YFilter = yf }

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetGoName(yname string) string {
    if yname == "query-timeout" { return "QueryTimeout" }
    if yname == "access-group" { return "AccessGroup" }
    if yname == "query-max-response-time" { return "QueryMaxResponseTime" }
    if yname == "version" { return "Version" }
    if yname == "router-enable" { return "RouterEnable" }
    if yname == "query-interval" { return "QueryInterval" }
    if yname == "maximum-groups-per-interface-oor" { return "MaximumGroupsPerInterfaceOor" }
    if yname == "explicit-tracking" { return "ExplicitTracking" }
    return ""
}

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetSegmentPath() string {
    return "inheritable-defaults"
}

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "maximum-groups-per-interface-oor" {
        return &inheritableDefaults.MaximumGroupsPerInterfaceOor
    }
    if childYangName == "explicit-tracking" {
        return &inheritableDefaults.ExplicitTracking
    }
    return nil
}

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["maximum-groups-per-interface-oor"] = &inheritableDefaults.MaximumGroupsPerInterfaceOor
    children["explicit-tracking"] = &inheritableDefaults.ExplicitTracking
    return children
}

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["query-timeout"] = inheritableDefaults.QueryTimeout
    leafs["access-group"] = inheritableDefaults.AccessGroup
    leafs["query-max-response-time"] = inheritableDefaults.QueryMaxResponseTime
    leafs["version"] = inheritableDefaults.Version
    leafs["router-enable"] = inheritableDefaults.RouterEnable
    leafs["query-interval"] = inheritableDefaults.QueryInterval
    return leafs
}

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetBundleName() string { return "cisco_ios_xr" }

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetYangName() string { return "inheritable-defaults" }

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) SetParent(parent types.Entity) { inheritableDefaults.parent = parent }

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetParent() types.Entity { return inheritableDefaults.parent }

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetParentYangName() string { return "vrf" }

// Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of groups accepted per interface by this router. The type is
    // interface{} with range: 1..40000. This attribute is mandatory.
    MaximumGroups interface{}

    // WarningThreshold for number of groups accepted per interface by this
    // router. The type is interface{} with range: 1..40000. The default value is
    // 25000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetFilter() yfilter.YFilter { return maximumGroupsPerInterfaceOor.YFilter }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) SetFilter(yf yfilter.YFilter) { maximumGroupsPerInterfaceOor.YFilter = yf }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetGoName(yname string) string {
    if yname == "maximum-groups" { return "MaximumGroups" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetSegmentPath() string {
    return "maximum-groups-per-interface-oor"
}

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-groups"] = maximumGroupsPerInterfaceOor.MaximumGroups
    leafs["warning-threshold"] = maximumGroupsPerInterfaceOor.WarningThreshold
    leafs["access-list-name"] = maximumGroupsPerInterfaceOor.AccessListName
    return leafs
}

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetBundleName() string { return "cisco_ios_xr" }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetYangName() string { return "maximum-groups-per-interface-oor" }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) SetParent(parent types.Entity) { maximumGroupsPerInterfaceOor.parent = parent }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetParent() types.Entity { return maximumGroupsPerInterfaceOor.parent }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetParentYangName() string { return "inheritable-defaults" }

// Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetFilter() yfilter.YFilter { return explicitTracking.YFilter }

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) SetFilter(yf yfilter.YFilter) { explicitTracking.YFilter = yf }

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetSegmentPath() string {
    return "explicit-tracking"
}

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = explicitTracking.Enable
    leafs["access-list-name"] = explicitTracking.AccessListName
    return leafs
}

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetBundleName() string { return "cisco_ios_xr" }

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetYangName() string { return "explicit-tracking" }

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) SetParent(parent types.Entity) { explicitTracking.parent = parent }

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetParent() types.Entity { return explicitTracking.parent }

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetParentYangName() string { return "inheritable-defaults" }

// Igmp_Vrfs_Vrf_SsmAccessGroups
// IGMP Source specific mode
type Igmp_Vrfs_Vrf_SsmAccessGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSM static Access Group. The type is slice of
    // Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup.
    SsmAccessGroup []Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup
}

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetFilter() yfilter.YFilter { return ssmAccessGroups.YFilter }

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) SetFilter(yf yfilter.YFilter) { ssmAccessGroups.YFilter = yf }

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetGoName(yname string) string {
    if yname == "ssm-access-group" { return "SsmAccessGroup" }
    return ""
}

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetSegmentPath() string {
    return "ssm-access-groups"
}

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ssm-access-group" {
        for _, c := range ssmAccessGroups.SsmAccessGroup {
            if ssmAccessGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup{}
        ssmAccessGroups.SsmAccessGroup = append(ssmAccessGroups.SsmAccessGroup, child)
        return &ssmAccessGroups.SsmAccessGroup[len(ssmAccessGroups.SsmAccessGroup)-1]
    }
    return nil
}

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ssmAccessGroups.SsmAccessGroup {
        children[ssmAccessGroups.SsmAccessGroup[i].GetSegmentPath()] = &ssmAccessGroups.SsmAccessGroup[i]
    }
    return children
}

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetBundleName() string { return "cisco_ios_xr" }

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetYangName() string { return "ssm-access-groups" }

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) SetParent(parent types.Entity) { ssmAccessGroups.parent = parent }

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetParent() types.Entity { return ssmAccessGroups.parent }

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetParentYangName() string { return "vrf" }

// Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup
// SSM static Access Group
type Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Access list specifying access group. The type is string with length: 1..64.
    // This attribute is mandatory.
    AccessListName interface{}
}

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetFilter() yfilter.YFilter { return ssmAccessGroup.YFilter }

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) SetFilter(yf yfilter.YFilter) { ssmAccessGroup.YFilter = yf }

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetSegmentPath() string {
    return "ssm-access-group" + "[source-address='" + fmt.Sprintf("%v", ssmAccessGroup.SourceAddress) + "']"
}

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = ssmAccessGroup.SourceAddress
    leafs["access-list-name"] = ssmAccessGroup.AccessListName
    return leafs
}

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetBundleName() string { return "cisco_ios_xr" }

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetYangName() string { return "ssm-access-group" }

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) SetParent(parent types.Entity) { ssmAccessGroup.parent = parent }

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetParent() types.Entity { return ssmAccessGroup.parent }

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetParentYangName() string { return "ssm-access-groups" }

// Igmp_Vrfs_Vrf_Maximum
// Configure IGMP State Limits
type Igmp_Vrfs_Vrf_Maximum struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure maximum number of groups accepted by this router. The type is
    // interface{} with range: 1..75000. The default value is 50000.
    MaximumGroups interface{}
}

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetFilter() yfilter.YFilter { return maximum.YFilter }

func (maximum *Igmp_Vrfs_Vrf_Maximum) SetFilter(yf yfilter.YFilter) { maximum.YFilter = yf }

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetGoName(yname string) string {
    if yname == "maximum-groups" { return "MaximumGroups" }
    return ""
}

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetSegmentPath() string {
    return "maximum"
}

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-groups"] = maximum.MaximumGroups
    return leafs
}

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetBundleName() string { return "cisco_ios_xr" }

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetYangName() string { return "maximum" }

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximum *Igmp_Vrfs_Vrf_Maximum) SetParent(parent types.Entity) { maximum.parent = parent }

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetParent() types.Entity { return maximum.parent }

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetParentYangName() string { return "vrf" }

// Igmp_Vrfs_Vrf_Interfaces
// Interface-level configuration
type Igmp_Vrfs_Vrf_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Igmp_Vrfs_Vrf_Interfaces_Interface.
    Interface []Igmp_Vrfs_Vrf_Interfaces_Interface
}

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_Vrfs_Vrf_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetParentYangName() string { return "vrf" }

// Igmp_Vrfs_Vrf_Interfaces_Interface
// The name of the interface
type Igmp_Vrfs_Vrf_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // IGMP previous querier timeout. The type is interface{} with range: 60..300.
    // Units are second.
    QueryTimeout interface{}

    // Access list specifying access group range. The type is string with length:
    // 1..64.
    AccessGroup interface{}

    // Query response value in seconds. The type is interface{} with range: 1..12.
    // Units are second. The default value is 10.
    QueryMaxResponseTime interface{}

    // Version number. The type is interface{} with range: 1..3. The default value
    // is 3.
    Version interface{}

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. The default value is true.
    RouterEnable interface{}

    // Query interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second. The default value is 60.
    QueryInterval interface{}

    // IGMP join multicast group.
    JoinGroups Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups

    // IGMP static multicast group.
    StaticGroupGroupAddresses Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses

    // Configure maximum number of groups accepted per interface by this router.
    MaximumGroupsPerInterfaceOor Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor

    // IGMPv3 explicit host tracking.
    ExplicitTracking Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking
}

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "query-timeout" { return "QueryTimeout" }
    if yname == "access-group" { return "AccessGroup" }
    if yname == "query-max-response-time" { return "QueryMaxResponseTime" }
    if yname == "version" { return "Version" }
    if yname == "router-enable" { return "RouterEnable" }
    if yname == "query-interval" { return "QueryInterval" }
    if yname == "join-groups" { return "JoinGroups" }
    if yname == "static-group-group-addresses" { return "StaticGroupGroupAddresses" }
    if yname == "maximum-groups-per-interface-oor" { return "MaximumGroupsPerInterfaceOor" }
    if yname == "explicit-tracking" { return "ExplicitTracking" }
    return ""
}

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "join-groups" {
        return &self.JoinGroups
    }
    if childYangName == "static-group-group-addresses" {
        return &self.StaticGroupGroupAddresses
    }
    if childYangName == "maximum-groups-per-interface-oor" {
        return &self.MaximumGroupsPerInterfaceOor
    }
    if childYangName == "explicit-tracking" {
        return &self.ExplicitTracking
    }
    return nil
}

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["join-groups"] = &self.JoinGroups
    children["static-group-group-addresses"] = &self.StaticGroupGroupAddresses
    children["maximum-groups-per-interface-oor"] = &self.MaximumGroupsPerInterfaceOor
    children["explicit-tracking"] = &self.ExplicitTracking
    return children
}

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["query-timeout"] = self.QueryTimeout
    leafs["access-group"] = self.AccessGroup
    leafs["query-max-response-time"] = self.QueryMaxResponseTime
    leafs["version"] = self.Version
    leafs["router-enable"] = self.RouterEnable
    leafs["query-interval"] = self.QueryInterval
    return leafs
}

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups
// IGMP join multicast group
// This type is a presence type.
type Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP group address and optional source address to include. The type is slice
    // of Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup.
    JoinGroup []Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup

    // IP group address and optional source address to include. The type is slice
    // of Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress.
    JoinGroupSourceAddress []Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
}

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetFilter() yfilter.YFilter { return joinGroups.YFilter }

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) SetFilter(yf yfilter.YFilter) { joinGroups.YFilter = yf }

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetGoName(yname string) string {
    if yname == "join-group" { return "JoinGroup" }
    if yname == "join-group-source-address" { return "JoinGroupSourceAddress" }
    return ""
}

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetSegmentPath() string {
    return "join-groups"
}

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "join-group" {
        for _, c := range joinGroups.JoinGroup {
            if joinGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup{}
        joinGroups.JoinGroup = append(joinGroups.JoinGroup, child)
        return &joinGroups.JoinGroup[len(joinGroups.JoinGroup)-1]
    }
    if childYangName == "join-group-source-address" {
        for _, c := range joinGroups.JoinGroupSourceAddress {
            if joinGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress{}
        joinGroups.JoinGroupSourceAddress = append(joinGroups.JoinGroupSourceAddress, child)
        return &joinGroups.JoinGroupSourceAddress[len(joinGroups.JoinGroupSourceAddress)-1]
    }
    return nil
}

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range joinGroups.JoinGroup {
        children[joinGroups.JoinGroup[i].GetSegmentPath()] = &joinGroups.JoinGroup[i]
    }
    for i := range joinGroups.JoinGroupSourceAddress {
        children[joinGroups.JoinGroupSourceAddress[i].GetSegmentPath()] = &joinGroups.JoinGroupSourceAddress[i]
    }
    return children
}

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetBundleName() string { return "cisco_ios_xr" }

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetYangName() string { return "join-groups" }

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) SetParent(parent types.Entity) { joinGroups.parent = parent }

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetParent() types.Entity { return joinGroups.parent }

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetParentYangName() string { return "interface" }

// Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // Filter mode. The type is IgmpFilter. This attribute is mandatory.
    Mode interface{}
}

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetFilter() yfilter.YFilter { return joinGroup.YFilter }

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) SetFilter(yf yfilter.YFilter) { joinGroup.YFilter = yf }

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetSegmentPath() string {
    return "join-group" + "[group-address='" + fmt.Sprintf("%v", joinGroup.GroupAddress) + "']"
}

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = joinGroup.GroupAddress
    leafs["mode"] = joinGroup.Mode
    return leafs
}

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetBundleName() string { return "cisco_ios_xr" }

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetYangName() string { return "join-group" }

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) SetParent(parent types.Entity) { joinGroup.parent = parent }

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetParent() types.Entity { return joinGroup.parent }

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetParentYangName() string { return "join-groups" }

// Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Optional IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Filter mode. The type is IgmpFilter. This attribute is mandatory.
    Mode interface{}
}

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetFilter() yfilter.YFilter { return joinGroupSourceAddress.YFilter }

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) SetFilter(yf yfilter.YFilter) { joinGroupSourceAddress.YFilter = yf }

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetSegmentPath() string {
    return "join-group-source-address" + "[group-address='" + fmt.Sprintf("%v", joinGroupSourceAddress.GroupAddress) + "']" + "[source-address='" + fmt.Sprintf("%v", joinGroupSourceAddress.SourceAddress) + "']"
}

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = joinGroupSourceAddress.GroupAddress
    leafs["source-address"] = joinGroupSourceAddress.SourceAddress
    leafs["mode"] = joinGroupSourceAddress.Mode
    return leafs
}

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetBundleName() string { return "cisco_ios_xr" }

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetYangName() string { return "join-group-source-address" }

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) SetParent(parent types.Entity) { joinGroupSourceAddress.parent = parent }

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetParent() types.Entity { return joinGroupSourceAddress.parent }

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetParentYangName() string { return "join-groups" }

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses
// IGMP static multicast group
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress.
    StaticGroupGroupAddress []Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress.
    StaticGroupGroupAddressSourceAddress []Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask.
    StaticGroupGroupAddressSourceAddressSourceAddressMask []Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask.
    StaticGroupGroupAddressGroupAddressMask []Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress.
    StaticGroupGroupAddressGroupAddressMaskSourceAddress []Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.
    StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask []Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
}

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetFilter() yfilter.YFilter { return staticGroupGroupAddresses.YFilter }

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddresses.YFilter = yf }

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetGoName(yname string) string {
    if yname == "static-group-group-address" { return "StaticGroupGroupAddress" }
    if yname == "static-group-group-address-source-address" { return "StaticGroupGroupAddressSourceAddress" }
    if yname == "static-group-group-address-source-address-source-address-mask" { return "StaticGroupGroupAddressSourceAddressSourceAddressMask" }
    if yname == "static-group-group-address-group-address-mask" { return "StaticGroupGroupAddressGroupAddressMask" }
    if yname == "static-group-group-address-group-address-mask-source-address" { return "StaticGroupGroupAddressGroupAddressMaskSourceAddress" }
    if yname == "static-group-group-address-group-address-mask-source-address-source-address-mask" { return "StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask" }
    return ""
}

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetSegmentPath() string {
    return "static-group-group-addresses"
}

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-group-group-address" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddress {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress{}
        staticGroupGroupAddresses.StaticGroupGroupAddress = append(staticGroupGroupAddresses.StaticGroupGroupAddress, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddress[len(staticGroupGroupAddresses.StaticGroupGroupAddress)-1]
    }
    if childYangName == "static-group-group-address-source-address" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress{}
        staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress = append(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[len(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress)-1]
    }
    if childYangName == "static-group-group-address-source-address-source-address-mask" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask{}
        staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask = append(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[len(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask)-1]
    }
    if childYangName == "static-group-group-address-group-address-mask" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask{}
        staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask = append(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[len(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask)-1]
    }
    if childYangName == "static-group-group-address-group-address-mask-source-address" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress{}
        staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress = append(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[len(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress)-1]
    }
    if childYangName == "static-group-group-address-group-address-mask-source-address-source-address-mask" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask{}
        staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask = append(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[len(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask)-1]
    }
    return nil
}

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddress {
        children[staticGroupGroupAddresses.StaticGroupGroupAddress[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddress[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i]
    }
    return children
}

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetYangName() string { return "static-group-group-addresses" }

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) SetParent(parent types.Entity) { staticGroupGroupAddresses.parent = parent }

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetParent() types.Entity { return staticGroupGroupAddresses.parent }

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetParentYangName() string { return "interface" }

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetFilter() yfilter.YFilter { return staticGroupGroupAddress.YFilter }

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddress.YFilter = yf }

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetSegmentPath() string {
    return "static-group-group-address" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddress.GroupAddress) + "']"
}

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddress.GroupAddress
    leafs["group-count"] = staticGroupGroupAddress.GroupCount
    leafs["source-count"] = staticGroupGroupAddress.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddress.SuppressReport
    return leafs
}

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetYangName() string { return "static-group-group-address" }

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) SetParent(parent types.Entity) { staticGroupGroupAddress.parent = parent }

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetParent() types.Entity { return staticGroupGroupAddress.parent }

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetParentYangName() string { return "static-group-group-addresses" }

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetFilter() yfilter.YFilter { return staticGroupGroupAddressSourceAddress.YFilter }

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressSourceAddress.YFilter = yf }

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetSegmentPath() string {
    return "static-group-group-address-source-address" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddress.GroupAddress) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddress.SourceAddress) + "']"
}

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressSourceAddress.GroupAddress
    leafs["source-address"] = staticGroupGroupAddressSourceAddress.SourceAddress
    leafs["group-count"] = staticGroupGroupAddressSourceAddress.GroupCount
    leafs["source-count"] = staticGroupGroupAddressSourceAddress.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressSourceAddress.SuppressReport
    return leafs
}

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetYangName() string { return "static-group-group-address-source-address" }

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) SetParent(parent types.Entity) { staticGroupGroupAddressSourceAddress.parent = parent }

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetParent() types.Entity { return staticGroupGroupAddressSourceAddress.parent }

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetParentYangName() string { return "static-group-group-addresses" }

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Mask for Source Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddressMask interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetFilter() yfilter.YFilter { return staticGroupGroupAddressSourceAddressSourceAddressMask.YFilter }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressSourceAddressSourceAddressMask.YFilter = yf }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "source-address-mask" { return "SourceAddressMask" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetSegmentPath() string {
    return "static-group-group-address-source-address-source-address-mask" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress) + "']" + "[source-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask) + "']"
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress
    leafs["source-address"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress
    leafs["source-address-mask"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask
    leafs["group-count"] = staticGroupGroupAddressSourceAddressSourceAddressMask.GroupCount
    leafs["source-count"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SuppressReport
    return leafs
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetYangName() string { return "static-group-group-address-source-address-source-address-mask" }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) SetParent(parent types.Entity) { staticGroupGroupAddressSourceAddressSourceAddressMask.parent = parent }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetParent() types.Entity { return staticGroupGroupAddressSourceAddressSourceAddressMask.parent }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetParentYangName() string { return "static-group-group-addresses" }

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Mask for Group Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddressMask interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetFilter() yfilter.YFilter { return staticGroupGroupAddressGroupAddressMask.YFilter }

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressGroupAddressMask.YFilter = yf }

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-address-mask" { return "GroupAddressMask" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetSegmentPath() string {
    return "static-group-group-address-group-address-mask" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMask.GroupAddress) + "']" + "[group-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMask.GroupAddressMask) + "']"
}

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressGroupAddressMask.GroupAddress
    leafs["group-address-mask"] = staticGroupGroupAddressGroupAddressMask.GroupAddressMask
    leafs["group-count"] = staticGroupGroupAddressGroupAddressMask.GroupCount
    leafs["source-count"] = staticGroupGroupAddressGroupAddressMask.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressGroupAddressMask.SuppressReport
    return leafs
}

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetYangName() string { return "static-group-group-address-group-address-mask" }

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) SetParent(parent types.Entity) { staticGroupGroupAddressGroupAddressMask.parent = parent }

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetParent() types.Entity { return staticGroupGroupAddressGroupAddressMask.parent }

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetParentYangName() string { return "static-group-group-addresses" }

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Mask for Group Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddressMask interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetFilter() yfilter.YFilter { return staticGroupGroupAddressGroupAddressMaskSourceAddress.YFilter }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressGroupAddressMaskSourceAddress.YFilter = yf }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-address-mask" { return "GroupAddressMask" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetSegmentPath() string {
    return "static-group-group-address-group-address-mask-source-address" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress) + "']" + "[group-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress) + "']"
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress
    leafs["group-address-mask"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask
    leafs["source-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress
    leafs["group-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupCount
    leafs["source-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.SuppressReport
    return leafs
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetYangName() string { return "static-group-group-address-group-address-mask-source-address" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) SetParent(parent types.Entity) { staticGroupGroupAddressGroupAddressMaskSourceAddress.parent = parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetParent() types.Entity { return staticGroupGroupAddressGroupAddressMaskSourceAddress.parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetParentYangName() string { return "static-group-group-addresses" }

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Mask for Group Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddressMask interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Mask for Source Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddressMask interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetFilter() yfilter.YFilter { return staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.YFilter }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.YFilter = yf }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-address-mask" { return "GroupAddressMask" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "source-address-mask" { return "SourceAddressMask" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetSegmentPath() string {
    return "static-group-group-address-group-address-mask-source-address-source-address-mask" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress) + "']" + "[group-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress) + "']" + "[source-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask) + "']"
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress
    leafs["group-address-mask"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask
    leafs["source-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress
    leafs["source-address-mask"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask
    leafs["group-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupCount
    leafs["source-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SuppressReport
    return leafs
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetYangName() string { return "static-group-group-address-group-address-mask-source-address-source-address-mask" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) SetParent(parent types.Entity) { staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.parent = parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetParent() types.Entity { return staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetParentYangName() string { return "static-group-group-addresses" }

// Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of groups accepted per interface by this router. The type is
    // interface{} with range: 1..40000. This attribute is mandatory.
    MaximumGroups interface{}

    // WarningThreshold for number of groups accepted per interface by this
    // router. The type is interface{} with range: 1..40000. The default value is
    // 25000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetFilter() yfilter.YFilter { return maximumGroupsPerInterfaceOor.YFilter }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) SetFilter(yf yfilter.YFilter) { maximumGroupsPerInterfaceOor.YFilter = yf }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetGoName(yname string) string {
    if yname == "maximum-groups" { return "MaximumGroups" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetSegmentPath() string {
    return "maximum-groups-per-interface-oor"
}

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-groups"] = maximumGroupsPerInterfaceOor.MaximumGroups
    leafs["warning-threshold"] = maximumGroupsPerInterfaceOor.WarningThreshold
    leafs["access-list-name"] = maximumGroupsPerInterfaceOor.AccessListName
    return leafs
}

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetBundleName() string { return "cisco_ios_xr" }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetYangName() string { return "maximum-groups-per-interface-oor" }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) SetParent(parent types.Entity) { maximumGroupsPerInterfaceOor.parent = parent }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetParent() types.Entity { return maximumGroupsPerInterfaceOor.parent }

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetParentYangName() string { return "interface" }

// Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetFilter() yfilter.YFilter { return explicitTracking.YFilter }

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) SetFilter(yf yfilter.YFilter) { explicitTracking.YFilter = yf }

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetSegmentPath() string {
    return "explicit-tracking"
}

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = explicitTracking.Enable
    leafs["access-list-name"] = explicitTracking.AccessListName
    return leafs
}

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetBundleName() string { return "cisco_ios_xr" }

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetYangName() string { return "explicit-tracking" }

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) SetParent(parent types.Entity) { explicitTracking.parent = parent }

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetParent() types.Entity { return explicitTracking.parent }

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetParentYangName() string { return "interface" }

// Igmp_DefaultContext
// Default Context
// This type is a presence type.
type Igmp_DefaultContext struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable SSM mapping using DNS Query. The type is interface{}.
    SsmdnsQueryGroup interface{}

    // Configure IGMP Robustness variable. The type is interface{} with range:
    // 2..10. The default value is 2.
    Robustness interface{}

    // Configure NSF specific options.
    Nsf Igmp_DefaultContext_Nsf

    // Configure IGMP QoS shapers for subscriber interfaces.
    UnicastQosAdjust Igmp_DefaultContext_UnicastQosAdjust

    // Configure IGMP accounting for logging join/leave records.
    Accounting Igmp_DefaultContext_Accounting

    // Configure IGMP Traffic variables.
    Traffic Igmp_DefaultContext_Traffic

    // Inheritable Defaults.
    InheritableDefaults Igmp_DefaultContext_InheritableDefaults

    // IGMP Source specific mode.
    SsmAccessGroups Igmp_DefaultContext_SsmAccessGroups

    // Configure IGMP State Limits.
    Maximum Igmp_DefaultContext_Maximum

    // Interface-level configuration.
    Interfaces Igmp_DefaultContext_Interfaces
}

func (defaultContext *Igmp_DefaultContext) GetFilter() yfilter.YFilter { return defaultContext.YFilter }

func (defaultContext *Igmp_DefaultContext) SetFilter(yf yfilter.YFilter) { defaultContext.YFilter = yf }

func (defaultContext *Igmp_DefaultContext) GetGoName(yname string) string {
    if yname == "ssmdns-query-group" { return "SsmdnsQueryGroup" }
    if yname == "robustness" { return "Robustness" }
    if yname == "nsf" { return "Nsf" }
    if yname == "unicast-qos-adjust" { return "UnicastQosAdjust" }
    if yname == "accounting" { return "Accounting" }
    if yname == "traffic" { return "Traffic" }
    if yname == "inheritable-defaults" { return "InheritableDefaults" }
    if yname == "ssm-access-groups" { return "SsmAccessGroups" }
    if yname == "maximum" { return "Maximum" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (defaultContext *Igmp_DefaultContext) GetSegmentPath() string {
    return "default-context"
}

func (defaultContext *Igmp_DefaultContext) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nsf" {
        return &defaultContext.Nsf
    }
    if childYangName == "unicast-qos-adjust" {
        return &defaultContext.UnicastQosAdjust
    }
    if childYangName == "accounting" {
        return &defaultContext.Accounting
    }
    if childYangName == "traffic" {
        return &defaultContext.Traffic
    }
    if childYangName == "inheritable-defaults" {
        return &defaultContext.InheritableDefaults
    }
    if childYangName == "ssm-access-groups" {
        return &defaultContext.SsmAccessGroups
    }
    if childYangName == "maximum" {
        return &defaultContext.Maximum
    }
    if childYangName == "interfaces" {
        return &defaultContext.Interfaces
    }
    return nil
}

func (defaultContext *Igmp_DefaultContext) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nsf"] = &defaultContext.Nsf
    children["unicast-qos-adjust"] = &defaultContext.UnicastQosAdjust
    children["accounting"] = &defaultContext.Accounting
    children["traffic"] = &defaultContext.Traffic
    children["inheritable-defaults"] = &defaultContext.InheritableDefaults
    children["ssm-access-groups"] = &defaultContext.SsmAccessGroups
    children["maximum"] = &defaultContext.Maximum
    children["interfaces"] = &defaultContext.Interfaces
    return children
}

func (defaultContext *Igmp_DefaultContext) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ssmdns-query-group"] = defaultContext.SsmdnsQueryGroup
    leafs["robustness"] = defaultContext.Robustness
    return leafs
}

func (defaultContext *Igmp_DefaultContext) GetBundleName() string { return "cisco_ios_xr" }

func (defaultContext *Igmp_DefaultContext) GetYangName() string { return "default-context" }

func (defaultContext *Igmp_DefaultContext) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultContext *Igmp_DefaultContext) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultContext *Igmp_DefaultContext) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultContext *Igmp_DefaultContext) SetParent(parent types.Entity) { defaultContext.parent = parent }

func (defaultContext *Igmp_DefaultContext) GetParent() types.Entity { return defaultContext.parent }

func (defaultContext *Igmp_DefaultContext) GetParentYangName() string { return "igmp" }

// Igmp_DefaultContext_Nsf
// Configure NSF specific options
type Igmp_DefaultContext_Nsf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum time for IGMP NSF mode in seconds. The type is interface{} with
    // range: 10..3600. Units are second. The default value is 60.
    Lifetime interface{}
}

func (nsf *Igmp_DefaultContext_Nsf) GetFilter() yfilter.YFilter { return nsf.YFilter }

func (nsf *Igmp_DefaultContext_Nsf) SetFilter(yf yfilter.YFilter) { nsf.YFilter = yf }

func (nsf *Igmp_DefaultContext_Nsf) GetGoName(yname string) string {
    if yname == "lifetime" { return "Lifetime" }
    return ""
}

func (nsf *Igmp_DefaultContext_Nsf) GetSegmentPath() string {
    return "nsf"
}

func (nsf *Igmp_DefaultContext_Nsf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nsf *Igmp_DefaultContext_Nsf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nsf *Igmp_DefaultContext_Nsf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lifetime"] = nsf.Lifetime
    return leafs
}

func (nsf *Igmp_DefaultContext_Nsf) GetBundleName() string { return "cisco_ios_xr" }

func (nsf *Igmp_DefaultContext_Nsf) GetYangName() string { return "nsf" }

func (nsf *Igmp_DefaultContext_Nsf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nsf *Igmp_DefaultContext_Nsf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nsf *Igmp_DefaultContext_Nsf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nsf *Igmp_DefaultContext_Nsf) SetParent(parent types.Entity) { nsf.parent = parent }

func (nsf *Igmp_DefaultContext_Nsf) GetParent() types.Entity { return nsf.parent }

func (nsf *Igmp_DefaultContext_Nsf) GetParentYangName() string { return "default-context" }

// Igmp_DefaultContext_UnicastQosAdjust
// Configure IGMP QoS shapers for subscriber
// interfaces
type Igmp_DefaultContext_UnicastQosAdjust struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the QoS download interval (in milliseconds). The type is
    // interface{} with range: 10..500. Units are millisecond. The default value
    // is 100.
    DownloadInterval interface{}

    // Configure the QoS delay before programming (in seconds). The type is
    // interface{} with range: 0..10. Units are second. The default value is 1.
    AdjustmentDelay interface{}

    // Configure the QoS hold off time (in seconds). The type is interface{} with
    // range: 5..1800. Units are second. The default value is 180.
    HoldOff interface{}
}

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetFilter() yfilter.YFilter { return unicastQosAdjust.YFilter }

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) SetFilter(yf yfilter.YFilter) { unicastQosAdjust.YFilter = yf }

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetGoName(yname string) string {
    if yname == "download-interval" { return "DownloadInterval" }
    if yname == "adjustment-delay" { return "AdjustmentDelay" }
    if yname == "hold-off" { return "HoldOff" }
    return ""
}

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetSegmentPath() string {
    return "unicast-qos-adjust"
}

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["download-interval"] = unicastQosAdjust.DownloadInterval
    leafs["adjustment-delay"] = unicastQosAdjust.AdjustmentDelay
    leafs["hold-off"] = unicastQosAdjust.HoldOff
    return leafs
}

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetBundleName() string { return "cisco_ios_xr" }

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetYangName() string { return "unicast-qos-adjust" }

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) SetParent(parent types.Entity) { unicastQosAdjust.parent = parent }

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetParent() types.Entity { return unicastQosAdjust.parent }

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetParentYangName() string { return "default-context" }

// Igmp_DefaultContext_Accounting
// Configure IGMP accounting for logging
// join/leave records
type Igmp_DefaultContext_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure IGMP accounting Maximum History setting. The type is interface{}
    // with range: 0..365. Units are day. The default value is 0.
    MaxHistory interface{}
}

func (accounting *Igmp_DefaultContext_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *Igmp_DefaultContext_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *Igmp_DefaultContext_Accounting) GetGoName(yname string) string {
    if yname == "max-history" { return "MaxHistory" }
    return ""
}

func (accounting *Igmp_DefaultContext_Accounting) GetSegmentPath() string {
    return "accounting"
}

func (accounting *Igmp_DefaultContext_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accounting *Igmp_DefaultContext_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accounting *Igmp_DefaultContext_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-history"] = accounting.MaxHistory
    return leafs
}

func (accounting *Igmp_DefaultContext_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *Igmp_DefaultContext_Accounting) GetYangName() string { return "accounting" }

func (accounting *Igmp_DefaultContext_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *Igmp_DefaultContext_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *Igmp_DefaultContext_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *Igmp_DefaultContext_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *Igmp_DefaultContext_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *Igmp_DefaultContext_Accounting) GetParentYangName() string { return "default-context" }

// Igmp_DefaultContext_Traffic
// Configure IGMP Traffic variables
type Igmp_DefaultContext_Traffic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the route-policy profile. The type is string with length: 1..64.
    Profile interface{}
}

func (traffic *Igmp_DefaultContext_Traffic) GetFilter() yfilter.YFilter { return traffic.YFilter }

func (traffic *Igmp_DefaultContext_Traffic) SetFilter(yf yfilter.YFilter) { traffic.YFilter = yf }

func (traffic *Igmp_DefaultContext_Traffic) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (traffic *Igmp_DefaultContext_Traffic) GetSegmentPath() string {
    return "traffic"
}

func (traffic *Igmp_DefaultContext_Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (traffic *Igmp_DefaultContext_Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (traffic *Igmp_DefaultContext_Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile"] = traffic.Profile
    return leafs
}

func (traffic *Igmp_DefaultContext_Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (traffic *Igmp_DefaultContext_Traffic) GetYangName() string { return "traffic" }

func (traffic *Igmp_DefaultContext_Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traffic *Igmp_DefaultContext_Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traffic *Igmp_DefaultContext_Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traffic *Igmp_DefaultContext_Traffic) SetParent(parent types.Entity) { traffic.parent = parent }

func (traffic *Igmp_DefaultContext_Traffic) GetParent() types.Entity { return traffic.parent }

func (traffic *Igmp_DefaultContext_Traffic) GetParentYangName() string { return "default-context" }

// Igmp_DefaultContext_InheritableDefaults
// Inheritable Defaults
type Igmp_DefaultContext_InheritableDefaults struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGMP previous querier timeout. The type is interface{} with range: 60..300.
    // Units are second.
    QueryTimeout interface{}

    // Access list specifying access group range. The type is string with length:
    // 1..64.
    AccessGroup interface{}

    // Query response value in seconds. The type is interface{} with range: 1..12.
    // Units are second. The default value is 10.
    QueryMaxResponseTime interface{}

    // Version number. The type is interface{} with range: 1..3. The default value
    // is 3.
    Version interface{}

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. The default value is true.
    RouterEnable interface{}

    // Query interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second. The default value is 60.
    QueryInterval interface{}

    // Configure maximum number of groups accepted per interface by this router.
    MaximumGroupsPerInterfaceOor Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor

    // IGMPv3 explicit host tracking.
    ExplicitTracking Igmp_DefaultContext_InheritableDefaults_ExplicitTracking
}

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetFilter() yfilter.YFilter { return inheritableDefaults.YFilter }

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) SetFilter(yf yfilter.YFilter) { inheritableDefaults.YFilter = yf }

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetGoName(yname string) string {
    if yname == "query-timeout" { return "QueryTimeout" }
    if yname == "access-group" { return "AccessGroup" }
    if yname == "query-max-response-time" { return "QueryMaxResponseTime" }
    if yname == "version" { return "Version" }
    if yname == "router-enable" { return "RouterEnable" }
    if yname == "query-interval" { return "QueryInterval" }
    if yname == "maximum-groups-per-interface-oor" { return "MaximumGroupsPerInterfaceOor" }
    if yname == "explicit-tracking" { return "ExplicitTracking" }
    return ""
}

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetSegmentPath() string {
    return "inheritable-defaults"
}

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "maximum-groups-per-interface-oor" {
        return &inheritableDefaults.MaximumGroupsPerInterfaceOor
    }
    if childYangName == "explicit-tracking" {
        return &inheritableDefaults.ExplicitTracking
    }
    return nil
}

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["maximum-groups-per-interface-oor"] = &inheritableDefaults.MaximumGroupsPerInterfaceOor
    children["explicit-tracking"] = &inheritableDefaults.ExplicitTracking
    return children
}

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["query-timeout"] = inheritableDefaults.QueryTimeout
    leafs["access-group"] = inheritableDefaults.AccessGroup
    leafs["query-max-response-time"] = inheritableDefaults.QueryMaxResponseTime
    leafs["version"] = inheritableDefaults.Version
    leafs["router-enable"] = inheritableDefaults.RouterEnable
    leafs["query-interval"] = inheritableDefaults.QueryInterval
    return leafs
}

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetBundleName() string { return "cisco_ios_xr" }

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetYangName() string { return "inheritable-defaults" }

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) SetParent(parent types.Entity) { inheritableDefaults.parent = parent }

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetParent() types.Entity { return inheritableDefaults.parent }

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetParentYangName() string { return "default-context" }

// Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of groups accepted per interface by this router. The type is
    // interface{} with range: 1..40000. This attribute is mandatory.
    MaximumGroups interface{}

    // WarningThreshold for number of groups accepted per interface by this
    // router. The type is interface{} with range: 1..40000. The default value is
    // 25000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetFilter() yfilter.YFilter { return maximumGroupsPerInterfaceOor.YFilter }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) SetFilter(yf yfilter.YFilter) { maximumGroupsPerInterfaceOor.YFilter = yf }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetGoName(yname string) string {
    if yname == "maximum-groups" { return "MaximumGroups" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetSegmentPath() string {
    return "maximum-groups-per-interface-oor"
}

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-groups"] = maximumGroupsPerInterfaceOor.MaximumGroups
    leafs["warning-threshold"] = maximumGroupsPerInterfaceOor.WarningThreshold
    leafs["access-list-name"] = maximumGroupsPerInterfaceOor.AccessListName
    return leafs
}

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetBundleName() string { return "cisco_ios_xr" }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetYangName() string { return "maximum-groups-per-interface-oor" }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) SetParent(parent types.Entity) { maximumGroupsPerInterfaceOor.parent = parent }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetParent() types.Entity { return maximumGroupsPerInterfaceOor.parent }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetParentYangName() string { return "inheritable-defaults" }

// Igmp_DefaultContext_InheritableDefaults_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Igmp_DefaultContext_InheritableDefaults_ExplicitTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetFilter() yfilter.YFilter { return explicitTracking.YFilter }

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) SetFilter(yf yfilter.YFilter) { explicitTracking.YFilter = yf }

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetSegmentPath() string {
    return "explicit-tracking"
}

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = explicitTracking.Enable
    leafs["access-list-name"] = explicitTracking.AccessListName
    return leafs
}

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetBundleName() string { return "cisco_ios_xr" }

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetYangName() string { return "explicit-tracking" }

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) SetParent(parent types.Entity) { explicitTracking.parent = parent }

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetParent() types.Entity { return explicitTracking.parent }

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetParentYangName() string { return "inheritable-defaults" }

// Igmp_DefaultContext_SsmAccessGroups
// IGMP Source specific mode
type Igmp_DefaultContext_SsmAccessGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSM static Access Group. The type is slice of
    // Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup.
    SsmAccessGroup []Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup
}

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetFilter() yfilter.YFilter { return ssmAccessGroups.YFilter }

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) SetFilter(yf yfilter.YFilter) { ssmAccessGroups.YFilter = yf }

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetGoName(yname string) string {
    if yname == "ssm-access-group" { return "SsmAccessGroup" }
    return ""
}

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetSegmentPath() string {
    return "ssm-access-groups"
}

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ssm-access-group" {
        for _, c := range ssmAccessGroups.SsmAccessGroup {
            if ssmAccessGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup{}
        ssmAccessGroups.SsmAccessGroup = append(ssmAccessGroups.SsmAccessGroup, child)
        return &ssmAccessGroups.SsmAccessGroup[len(ssmAccessGroups.SsmAccessGroup)-1]
    }
    return nil
}

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ssmAccessGroups.SsmAccessGroup {
        children[ssmAccessGroups.SsmAccessGroup[i].GetSegmentPath()] = &ssmAccessGroups.SsmAccessGroup[i]
    }
    return children
}

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetBundleName() string { return "cisco_ios_xr" }

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetYangName() string { return "ssm-access-groups" }

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) SetParent(parent types.Entity) { ssmAccessGroups.parent = parent }

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetParent() types.Entity { return ssmAccessGroups.parent }

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetParentYangName() string { return "default-context" }

// Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup
// SSM static Access Group
type Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Access list specifying access group. The type is string with length: 1..64.
    // This attribute is mandatory.
    AccessListName interface{}
}

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetFilter() yfilter.YFilter { return ssmAccessGroup.YFilter }

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) SetFilter(yf yfilter.YFilter) { ssmAccessGroup.YFilter = yf }

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetSegmentPath() string {
    return "ssm-access-group" + "[source-address='" + fmt.Sprintf("%v", ssmAccessGroup.SourceAddress) + "']"
}

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = ssmAccessGroup.SourceAddress
    leafs["access-list-name"] = ssmAccessGroup.AccessListName
    return leafs
}

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetBundleName() string { return "cisco_ios_xr" }

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetYangName() string { return "ssm-access-group" }

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) SetParent(parent types.Entity) { ssmAccessGroup.parent = parent }

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetParent() types.Entity { return ssmAccessGroup.parent }

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetParentYangName() string { return "ssm-access-groups" }

// Igmp_DefaultContext_Maximum
// Configure IGMP State Limits
type Igmp_DefaultContext_Maximum struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure maximum number of groups accepted by this router. The type is
    // interface{} with range: 1..75000. The default value is 50000.
    MaximumGroups interface{}
}

func (maximum *Igmp_DefaultContext_Maximum) GetFilter() yfilter.YFilter { return maximum.YFilter }

func (maximum *Igmp_DefaultContext_Maximum) SetFilter(yf yfilter.YFilter) { maximum.YFilter = yf }

func (maximum *Igmp_DefaultContext_Maximum) GetGoName(yname string) string {
    if yname == "maximum-groups" { return "MaximumGroups" }
    return ""
}

func (maximum *Igmp_DefaultContext_Maximum) GetSegmentPath() string {
    return "maximum"
}

func (maximum *Igmp_DefaultContext_Maximum) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximum *Igmp_DefaultContext_Maximum) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximum *Igmp_DefaultContext_Maximum) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-groups"] = maximum.MaximumGroups
    return leafs
}

func (maximum *Igmp_DefaultContext_Maximum) GetBundleName() string { return "cisco_ios_xr" }

func (maximum *Igmp_DefaultContext_Maximum) GetYangName() string { return "maximum" }

func (maximum *Igmp_DefaultContext_Maximum) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximum *Igmp_DefaultContext_Maximum) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximum *Igmp_DefaultContext_Maximum) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximum *Igmp_DefaultContext_Maximum) SetParent(parent types.Entity) { maximum.parent = parent }

func (maximum *Igmp_DefaultContext_Maximum) GetParent() types.Entity { return maximum.parent }

func (maximum *Igmp_DefaultContext_Maximum) GetParentYangName() string { return "default-context" }

// Igmp_DefaultContext_Interfaces
// Interface-level configuration
type Igmp_DefaultContext_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Igmp_DefaultContext_Interfaces_Interface.
    Interface []Igmp_DefaultContext_Interfaces_Interface
}

func (interfaces *Igmp_DefaultContext_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Igmp_DefaultContext_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Igmp_DefaultContext_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Igmp_DefaultContext_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Igmp_DefaultContext_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_DefaultContext_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Igmp_DefaultContext_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Igmp_DefaultContext_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Igmp_DefaultContext_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Igmp_DefaultContext_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Igmp_DefaultContext_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Igmp_DefaultContext_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Igmp_DefaultContext_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Igmp_DefaultContext_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Igmp_DefaultContext_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Igmp_DefaultContext_Interfaces) GetParentYangName() string { return "default-context" }

// Igmp_DefaultContext_Interfaces_Interface
// The name of the interface
type Igmp_DefaultContext_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // IGMP previous querier timeout. The type is interface{} with range: 60..300.
    // Units are second.
    QueryTimeout interface{}

    // Access list specifying access group range. The type is string with length:
    // 1..64.
    AccessGroup interface{}

    // Query response value in seconds. The type is interface{} with range: 1..12.
    // Units are second. The default value is 10.
    QueryMaxResponseTime interface{}

    // Version number. The type is interface{} with range: 1..3. The default value
    // is 3.
    Version interface{}

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. The default value is true.
    RouterEnable interface{}

    // Query interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second. The default value is 60.
    QueryInterval interface{}

    // IGMP join multicast group.
    JoinGroups Igmp_DefaultContext_Interfaces_Interface_JoinGroups

    // IGMP static multicast group.
    StaticGroupGroupAddresses Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses

    // Configure maximum number of groups accepted per interface by this router.
    MaximumGroupsPerInterfaceOor Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor

    // IGMPv3 explicit host tracking.
    ExplicitTracking Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking
}

func (self *Igmp_DefaultContext_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Igmp_DefaultContext_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Igmp_DefaultContext_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "query-timeout" { return "QueryTimeout" }
    if yname == "access-group" { return "AccessGroup" }
    if yname == "query-max-response-time" { return "QueryMaxResponseTime" }
    if yname == "version" { return "Version" }
    if yname == "router-enable" { return "RouterEnable" }
    if yname == "query-interval" { return "QueryInterval" }
    if yname == "join-groups" { return "JoinGroups" }
    if yname == "static-group-group-addresses" { return "StaticGroupGroupAddresses" }
    if yname == "maximum-groups-per-interface-oor" { return "MaximumGroupsPerInterfaceOor" }
    if yname == "explicit-tracking" { return "ExplicitTracking" }
    return ""
}

func (self *Igmp_DefaultContext_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Igmp_DefaultContext_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "join-groups" {
        return &self.JoinGroups
    }
    if childYangName == "static-group-group-addresses" {
        return &self.StaticGroupGroupAddresses
    }
    if childYangName == "maximum-groups-per-interface-oor" {
        return &self.MaximumGroupsPerInterfaceOor
    }
    if childYangName == "explicit-tracking" {
        return &self.ExplicitTracking
    }
    return nil
}

func (self *Igmp_DefaultContext_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["join-groups"] = &self.JoinGroups
    children["static-group-group-addresses"] = &self.StaticGroupGroupAddresses
    children["maximum-groups-per-interface-oor"] = &self.MaximumGroupsPerInterfaceOor
    children["explicit-tracking"] = &self.ExplicitTracking
    return children
}

func (self *Igmp_DefaultContext_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["query-timeout"] = self.QueryTimeout
    leafs["access-group"] = self.AccessGroup
    leafs["query-max-response-time"] = self.QueryMaxResponseTime
    leafs["version"] = self.Version
    leafs["router-enable"] = self.RouterEnable
    leafs["query-interval"] = self.QueryInterval
    return leafs
}

func (self *Igmp_DefaultContext_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Igmp_DefaultContext_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Igmp_DefaultContext_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Igmp_DefaultContext_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Igmp_DefaultContext_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Igmp_DefaultContext_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Igmp_DefaultContext_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Igmp_DefaultContext_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Igmp_DefaultContext_Interfaces_Interface_JoinGroups
// IGMP join multicast group
// This type is a presence type.
type Igmp_DefaultContext_Interfaces_Interface_JoinGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP group address and optional source address to include. The type is slice
    // of Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup.
    JoinGroup []Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress.
    JoinGroupSourceAddress []Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
}

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetFilter() yfilter.YFilter { return joinGroups.YFilter }

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) SetFilter(yf yfilter.YFilter) { joinGroups.YFilter = yf }

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetGoName(yname string) string {
    if yname == "join-group" { return "JoinGroup" }
    if yname == "join-group-source-address" { return "JoinGroupSourceAddress" }
    return ""
}

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetSegmentPath() string {
    return "join-groups"
}

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "join-group" {
        for _, c := range joinGroups.JoinGroup {
            if joinGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup{}
        joinGroups.JoinGroup = append(joinGroups.JoinGroup, child)
        return &joinGroups.JoinGroup[len(joinGroups.JoinGroup)-1]
    }
    if childYangName == "join-group-source-address" {
        for _, c := range joinGroups.JoinGroupSourceAddress {
            if joinGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress{}
        joinGroups.JoinGroupSourceAddress = append(joinGroups.JoinGroupSourceAddress, child)
        return &joinGroups.JoinGroupSourceAddress[len(joinGroups.JoinGroupSourceAddress)-1]
    }
    return nil
}

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range joinGroups.JoinGroup {
        children[joinGroups.JoinGroup[i].GetSegmentPath()] = &joinGroups.JoinGroup[i]
    }
    for i := range joinGroups.JoinGroupSourceAddress {
        children[joinGroups.JoinGroupSourceAddress[i].GetSegmentPath()] = &joinGroups.JoinGroupSourceAddress[i]
    }
    return children
}

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetBundleName() string { return "cisco_ios_xr" }

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetYangName() string { return "join-groups" }

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) SetParent(parent types.Entity) { joinGroups.parent = parent }

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetParent() types.Entity { return joinGroups.parent }

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetParentYangName() string { return "interface" }

// Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // Filter mode. The type is IgmpFilter. This attribute is mandatory.
    Mode interface{}
}

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetFilter() yfilter.YFilter { return joinGroup.YFilter }

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) SetFilter(yf yfilter.YFilter) { joinGroup.YFilter = yf }

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetSegmentPath() string {
    return "join-group" + "[group-address='" + fmt.Sprintf("%v", joinGroup.GroupAddress) + "']"
}

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = joinGroup.GroupAddress
    leafs["mode"] = joinGroup.Mode
    return leafs
}

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetBundleName() string { return "cisco_ios_xr" }

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetYangName() string { return "join-group" }

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) SetParent(parent types.Entity) { joinGroup.parent = parent }

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetParent() types.Entity { return joinGroup.parent }

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetParentYangName() string { return "join-groups" }

// Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Optional IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Filter mode. The type is IgmpFilter. This attribute is mandatory.
    Mode interface{}
}

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetFilter() yfilter.YFilter { return joinGroupSourceAddress.YFilter }

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) SetFilter(yf yfilter.YFilter) { joinGroupSourceAddress.YFilter = yf }

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetSegmentPath() string {
    return "join-group-source-address" + "[group-address='" + fmt.Sprintf("%v", joinGroupSourceAddress.GroupAddress) + "']" + "[source-address='" + fmt.Sprintf("%v", joinGroupSourceAddress.SourceAddress) + "']"
}

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = joinGroupSourceAddress.GroupAddress
    leafs["source-address"] = joinGroupSourceAddress.SourceAddress
    leafs["mode"] = joinGroupSourceAddress.Mode
    return leafs
}

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetBundleName() string { return "cisco_ios_xr" }

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetYangName() string { return "join-group-source-address" }

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) SetParent(parent types.Entity) { joinGroupSourceAddress.parent = parent }

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetParent() types.Entity { return joinGroupSourceAddress.parent }

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetParentYangName() string { return "join-groups" }

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses
// IGMP static multicast group
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress.
    StaticGroupGroupAddress []Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress.
    StaticGroupGroupAddressSourceAddress []Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask.
    StaticGroupGroupAddressSourceAddressSourceAddressMask []Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask.
    StaticGroupGroupAddressGroupAddressMask []Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress.
    StaticGroupGroupAddressGroupAddressMaskSourceAddress []Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.
    StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask []Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
}

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetFilter() yfilter.YFilter { return staticGroupGroupAddresses.YFilter }

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddresses.YFilter = yf }

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetGoName(yname string) string {
    if yname == "static-group-group-address" { return "StaticGroupGroupAddress" }
    if yname == "static-group-group-address-source-address" { return "StaticGroupGroupAddressSourceAddress" }
    if yname == "static-group-group-address-source-address-source-address-mask" { return "StaticGroupGroupAddressSourceAddressSourceAddressMask" }
    if yname == "static-group-group-address-group-address-mask" { return "StaticGroupGroupAddressGroupAddressMask" }
    if yname == "static-group-group-address-group-address-mask-source-address" { return "StaticGroupGroupAddressGroupAddressMaskSourceAddress" }
    if yname == "static-group-group-address-group-address-mask-source-address-source-address-mask" { return "StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask" }
    return ""
}

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetSegmentPath() string {
    return "static-group-group-addresses"
}

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-group-group-address" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddress {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress{}
        staticGroupGroupAddresses.StaticGroupGroupAddress = append(staticGroupGroupAddresses.StaticGroupGroupAddress, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddress[len(staticGroupGroupAddresses.StaticGroupGroupAddress)-1]
    }
    if childYangName == "static-group-group-address-source-address" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress{}
        staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress = append(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[len(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress)-1]
    }
    if childYangName == "static-group-group-address-source-address-source-address-mask" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask{}
        staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask = append(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[len(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask)-1]
    }
    if childYangName == "static-group-group-address-group-address-mask" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask{}
        staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask = append(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[len(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask)-1]
    }
    if childYangName == "static-group-group-address-group-address-mask-source-address" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress{}
        staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress = append(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[len(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress)-1]
    }
    if childYangName == "static-group-group-address-group-address-mask-source-address-source-address-mask" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask{}
        staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask = append(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[len(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask)-1]
    }
    return nil
}

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddress {
        children[staticGroupGroupAddresses.StaticGroupGroupAddress[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddress[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i]
    }
    return children
}

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetYangName() string { return "static-group-group-addresses" }

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) SetParent(parent types.Entity) { staticGroupGroupAddresses.parent = parent }

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetParent() types.Entity { return staticGroupGroupAddresses.parent }

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetParentYangName() string { return "interface" }

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetFilter() yfilter.YFilter { return staticGroupGroupAddress.YFilter }

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddress.YFilter = yf }

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetSegmentPath() string {
    return "static-group-group-address" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddress.GroupAddress) + "']"
}

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddress.GroupAddress
    leafs["group-count"] = staticGroupGroupAddress.GroupCount
    leafs["source-count"] = staticGroupGroupAddress.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddress.SuppressReport
    return leafs
}

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetYangName() string { return "static-group-group-address" }

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) SetParent(parent types.Entity) { staticGroupGroupAddress.parent = parent }

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetParent() types.Entity { return staticGroupGroupAddress.parent }

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetParentYangName() string { return "static-group-group-addresses" }

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetFilter() yfilter.YFilter { return staticGroupGroupAddressSourceAddress.YFilter }

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressSourceAddress.YFilter = yf }

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetSegmentPath() string {
    return "static-group-group-address-source-address" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddress.GroupAddress) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddress.SourceAddress) + "']"
}

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressSourceAddress.GroupAddress
    leafs["source-address"] = staticGroupGroupAddressSourceAddress.SourceAddress
    leafs["group-count"] = staticGroupGroupAddressSourceAddress.GroupCount
    leafs["source-count"] = staticGroupGroupAddressSourceAddress.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressSourceAddress.SuppressReport
    return leafs
}

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetYangName() string { return "static-group-group-address-source-address" }

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) SetParent(parent types.Entity) { staticGroupGroupAddressSourceAddress.parent = parent }

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetParent() types.Entity { return staticGroupGroupAddressSourceAddress.parent }

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetParentYangName() string { return "static-group-group-addresses" }

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Mask for Source Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddressMask interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetFilter() yfilter.YFilter { return staticGroupGroupAddressSourceAddressSourceAddressMask.YFilter }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressSourceAddressSourceAddressMask.YFilter = yf }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "source-address-mask" { return "SourceAddressMask" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetSegmentPath() string {
    return "static-group-group-address-source-address-source-address-mask" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress) + "']" + "[source-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask) + "']"
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress
    leafs["source-address"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress
    leafs["source-address-mask"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask
    leafs["group-count"] = staticGroupGroupAddressSourceAddressSourceAddressMask.GroupCount
    leafs["source-count"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SuppressReport
    return leafs
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetYangName() string { return "static-group-group-address-source-address-source-address-mask" }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) SetParent(parent types.Entity) { staticGroupGroupAddressSourceAddressSourceAddressMask.parent = parent }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetParent() types.Entity { return staticGroupGroupAddressSourceAddressSourceAddressMask.parent }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetParentYangName() string { return "static-group-group-addresses" }

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Mask for Group Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddressMask interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetFilter() yfilter.YFilter { return staticGroupGroupAddressGroupAddressMask.YFilter }

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressGroupAddressMask.YFilter = yf }

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-address-mask" { return "GroupAddressMask" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetSegmentPath() string {
    return "static-group-group-address-group-address-mask" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMask.GroupAddress) + "']" + "[group-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMask.GroupAddressMask) + "']"
}

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressGroupAddressMask.GroupAddress
    leafs["group-address-mask"] = staticGroupGroupAddressGroupAddressMask.GroupAddressMask
    leafs["group-count"] = staticGroupGroupAddressGroupAddressMask.GroupCount
    leafs["source-count"] = staticGroupGroupAddressGroupAddressMask.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressGroupAddressMask.SuppressReport
    return leafs
}

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetYangName() string { return "static-group-group-address-group-address-mask" }

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) SetParent(parent types.Entity) { staticGroupGroupAddressGroupAddressMask.parent = parent }

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetParent() types.Entity { return staticGroupGroupAddressGroupAddressMask.parent }

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetParentYangName() string { return "static-group-group-addresses" }

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Mask for Group Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddressMask interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetFilter() yfilter.YFilter { return staticGroupGroupAddressGroupAddressMaskSourceAddress.YFilter }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressGroupAddressMaskSourceAddress.YFilter = yf }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-address-mask" { return "GroupAddressMask" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetSegmentPath() string {
    return "static-group-group-address-group-address-mask-source-address" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress) + "']" + "[group-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress) + "']"
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress
    leafs["group-address-mask"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask
    leafs["source-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress
    leafs["group-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupCount
    leafs["source-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.SuppressReport
    return leafs
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetYangName() string { return "static-group-group-address-group-address-mask-source-address" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) SetParent(parent types.Entity) { staticGroupGroupAddressGroupAddressMaskSourceAddress.parent = parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetParent() types.Entity { return staticGroupGroupAddressGroupAddressMaskSourceAddress.parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetParentYangName() string { return "static-group-group-addresses" }

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Mask for Group Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddressMask interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Mask for Source Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddressMask interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetFilter() yfilter.YFilter { return staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.YFilter }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.YFilter = yf }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-address-mask" { return "GroupAddressMask" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "source-address-mask" { return "SourceAddressMask" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetSegmentPath() string {
    return "static-group-group-address-group-address-mask-source-address-source-address-mask" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress) + "']" + "[group-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress) + "']" + "[source-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask) + "']"
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress
    leafs["group-address-mask"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask
    leafs["source-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress
    leafs["source-address-mask"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask
    leafs["group-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupCount
    leafs["source-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SuppressReport
    return leafs
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetYangName() string { return "static-group-group-address-group-address-mask-source-address-source-address-mask" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) SetParent(parent types.Entity) { staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.parent = parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetParent() types.Entity { return staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetParentYangName() string { return "static-group-group-addresses" }

// Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of groups accepted per interface by this router. The type is
    // interface{} with range: 1..40000. This attribute is mandatory.
    MaximumGroups interface{}

    // WarningThreshold for number of groups accepted per interface by this
    // router. The type is interface{} with range: 1..40000. The default value is
    // 25000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetFilter() yfilter.YFilter { return maximumGroupsPerInterfaceOor.YFilter }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) SetFilter(yf yfilter.YFilter) { maximumGroupsPerInterfaceOor.YFilter = yf }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetGoName(yname string) string {
    if yname == "maximum-groups" { return "MaximumGroups" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetSegmentPath() string {
    return "maximum-groups-per-interface-oor"
}

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-groups"] = maximumGroupsPerInterfaceOor.MaximumGroups
    leafs["warning-threshold"] = maximumGroupsPerInterfaceOor.WarningThreshold
    leafs["access-list-name"] = maximumGroupsPerInterfaceOor.AccessListName
    return leafs
}

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetBundleName() string { return "cisco_ios_xr" }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetYangName() string { return "maximum-groups-per-interface-oor" }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) SetParent(parent types.Entity) { maximumGroupsPerInterfaceOor.parent = parent }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetParent() types.Entity { return maximumGroupsPerInterfaceOor.parent }

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetParentYangName() string { return "interface" }

// Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetFilter() yfilter.YFilter { return explicitTracking.YFilter }

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) SetFilter(yf yfilter.YFilter) { explicitTracking.YFilter = yf }

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetSegmentPath() string {
    return "explicit-tracking"
}

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = explicitTracking.Enable
    leafs["access-list-name"] = explicitTracking.AccessListName
    return leafs
}

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetBundleName() string { return "cisco_ios_xr" }

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetYangName() string { return "explicit-tracking" }

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) SetParent(parent types.Entity) { explicitTracking.parent = parent }

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetParent() types.Entity { return explicitTracking.parent }

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetParentYangName() string { return "interface" }

// Amt
// amt
type Amt struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Maximum number of IPv4 route-gateways (Tunnels). The type is
    // interface{} with range: 1..4294967295.
    MaximumV4RouteGateway interface{}

    // Access list for Gateway Filter. The type is string with length: 1..64.
    GatewayFilter interface{}

    // Configure Maximum number of IPv4 Routes. The type is interface{} with
    // range: 1..4294967295.
    MaximumV4Routes interface{}

    // Configure AMT Relay TOS. The type is interface{} with range: 1..255.
    Amttos interface{}

    // Configure AMT Relay TTL. The type is interface{} with range: 1..255.
    Amtttl interface{}

    // Configure Maximum number of IPv6 route-gateways (Tunnels). The type is
    // interface{} with range: 1..4294967295.
    MaximumV6RouteGateway interface{}

    // Configure AMT maximum number of Gateways. The type is interface{} with
    // range: 1..4294967295.
    MaximumGateway interface{}

    // Configure Maximum number of IPv6 Routes. The type is interface{} with
    // range: 1..4294967295.
    MaximumV6Routes interface{}

    // Configure AMT QQIC value. The type is interface{} with range:
    // 1..4294967295.
    Amtqqic interface{}

    // Configure AMT Relay MTU. The type is interface{} with range: 100..65535.
    Amtmtu interface{}

    // Configure AMT Relay IPv4 Advertisement Address.
    RelayAdvAdd Amt_RelayAdvAdd

    // Configure AMT Relay IPv4 Anycast-Prefix.
    RelayAnycastPrefix Amt_RelayAnycastPrefix
}

func (amt *Amt) GetFilter() yfilter.YFilter { return amt.YFilter }

func (amt *Amt) SetFilter(yf yfilter.YFilter) { amt.YFilter = yf }

func (amt *Amt) GetGoName(yname string) string {
    if yname == "maximum-v4-route-gateway" { return "MaximumV4RouteGateway" }
    if yname == "gateway-filter" { return "GatewayFilter" }
    if yname == "maximum-v4-routes" { return "MaximumV4Routes" }
    if yname == "amttos" { return "Amttos" }
    if yname == "amtttl" { return "Amtttl" }
    if yname == "maximum-v6-route-gateway" { return "MaximumV6RouteGateway" }
    if yname == "maximum-gateway" { return "MaximumGateway" }
    if yname == "maximum-v6-routes" { return "MaximumV6Routes" }
    if yname == "amtqqic" { return "Amtqqic" }
    if yname == "amtmtu" { return "Amtmtu" }
    if yname == "relay-adv-add" { return "RelayAdvAdd" }
    if yname == "relay-anycast-prefix" { return "RelayAnycastPrefix" }
    return ""
}

func (amt *Amt) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-igmp-cfg:amt"
}

func (amt *Amt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "relay-adv-add" {
        return &amt.RelayAdvAdd
    }
    if childYangName == "relay-anycast-prefix" {
        return &amt.RelayAnycastPrefix
    }
    return nil
}

func (amt *Amt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["relay-adv-add"] = &amt.RelayAdvAdd
    children["relay-anycast-prefix"] = &amt.RelayAnycastPrefix
    return children
}

func (amt *Amt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-v4-route-gateway"] = amt.MaximumV4RouteGateway
    leafs["gateway-filter"] = amt.GatewayFilter
    leafs["maximum-v4-routes"] = amt.MaximumV4Routes
    leafs["amttos"] = amt.Amttos
    leafs["amtttl"] = amt.Amtttl
    leafs["maximum-v6-route-gateway"] = amt.MaximumV6RouteGateway
    leafs["maximum-gateway"] = amt.MaximumGateway
    leafs["maximum-v6-routes"] = amt.MaximumV6Routes
    leafs["amtqqic"] = amt.Amtqqic
    leafs["amtmtu"] = amt.Amtmtu
    return leafs
}

func (amt *Amt) GetBundleName() string { return "cisco_ios_xr" }

func (amt *Amt) GetYangName() string { return "amt" }

func (amt *Amt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (amt *Amt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (amt *Amt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (amt *Amt) SetParent(parent types.Entity) { amt.parent = parent }

func (amt *Amt) GetParent() types.Entity { return amt.parent }

func (amt *Amt) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-igmp-cfg" }

// Amt_RelayAdvAdd
// Configure AMT Relay IPv4 Advertisement Address
// This type is a presence type.
type Amt_RelayAdvAdd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AMT Relay IPv4 Advertisement Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}

    // Relay Advertisement Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}
}

func (relayAdvAdd *Amt_RelayAdvAdd) GetFilter() yfilter.YFilter { return relayAdvAdd.YFilter }

func (relayAdvAdd *Amt_RelayAdvAdd) SetFilter(yf yfilter.YFilter) { relayAdvAdd.YFilter = yf }

func (relayAdvAdd *Amt_RelayAdvAdd) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (relayAdvAdd *Amt_RelayAdvAdd) GetSegmentPath() string {
    return "relay-adv-add"
}

func (relayAdvAdd *Amt_RelayAdvAdd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (relayAdvAdd *Amt_RelayAdvAdd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (relayAdvAdd *Amt_RelayAdvAdd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = relayAdvAdd.Address
    leafs["interface"] = relayAdvAdd.Interface
    return leafs
}

func (relayAdvAdd *Amt_RelayAdvAdd) GetBundleName() string { return "cisco_ios_xr" }

func (relayAdvAdd *Amt_RelayAdvAdd) GetYangName() string { return "relay-adv-add" }

func (relayAdvAdd *Amt_RelayAdvAdd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relayAdvAdd *Amt_RelayAdvAdd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relayAdvAdd *Amt_RelayAdvAdd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relayAdvAdd *Amt_RelayAdvAdd) SetParent(parent types.Entity) { relayAdvAdd.parent = parent }

func (relayAdvAdd *Amt_RelayAdvAdd) GetParent() types.Entity { return relayAdvAdd.parent }

func (relayAdvAdd *Amt_RelayAdvAdd) GetParentYangName() string { return "amt" }

// Amt_RelayAnycastPrefix
// Configure AMT Relay IPv4 Anycast-Prefix
// This type is a presence type.
type Amt_RelayAnycastPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Anycast-Prefix Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}

    // Mask Length for Anycast Prefix. The type is interface{} with range: 1..32.
    PrefixLength interface{}
}

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetFilter() yfilter.YFilter { return relayAnycastPrefix.YFilter }

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) SetFilter(yf yfilter.YFilter) { relayAnycastPrefix.YFilter = yf }

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetSegmentPath() string {
    return "relay-anycast-prefix"
}

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = relayAnycastPrefix.Address
    leafs["prefix-length"] = relayAnycastPrefix.PrefixLength
    return leafs
}

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetYangName() string { return "relay-anycast-prefix" }

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) SetParent(parent types.Entity) { relayAnycastPrefix.parent = parent }

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetParent() types.Entity { return relayAnycastPrefix.parent }

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetParentYangName() string { return "amt" }

// Mld
// mld
// This type is a presence type.
type Mld struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF related configuration.
    Vrfs Mld_Vrfs

    // Default Context.
    DefaultContext Mld_DefaultContext
}

func (mld *Mld) GetFilter() yfilter.YFilter { return mld.YFilter }

func (mld *Mld) SetFilter(yf yfilter.YFilter) { mld.YFilter = yf }

func (mld *Mld) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    if yname == "default-context" { return "DefaultContext" }
    return ""
}

func (mld *Mld) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-igmp-cfg:mld"
}

func (mld *Mld) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &mld.Vrfs
    }
    if childYangName == "default-context" {
        return &mld.DefaultContext
    }
    return nil
}

func (mld *Mld) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &mld.Vrfs
    children["default-context"] = &mld.DefaultContext
    return children
}

func (mld *Mld) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mld *Mld) GetBundleName() string { return "cisco_ios_xr" }

func (mld *Mld) GetYangName() string { return "mld" }

func (mld *Mld) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mld *Mld) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mld *Mld) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mld *Mld) SetParent(parent types.Entity) { mld.parent = parent }

func (mld *Mld) GetParent() types.Entity { return mld.parent }

func (mld *Mld) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-igmp-cfg" }

// Mld_Vrfs
// VRF related configuration
type Mld_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a particular vrf. The type is slice of Mld_Vrfs_Vrf.
    Vrf []Mld_Vrfs_Vrf
}

func (vrfs *Mld_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Mld_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Mld_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Mld_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Mld_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Mld_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Mld_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Mld_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Mld_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Mld_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Mld_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Mld_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Mld_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Mld_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Mld_Vrfs) GetParentYangName() string { return "mld" }

// Mld_Vrfs_Vrf
// Configuration for a particular vrf
type Mld_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name for this vrf. The type is string with length:
    // 1..32.
    VrfName interface{}

    // Enable SSM mapping using DNS Query. The type is interface{}.
    SsmdnsQueryGroup interface{}

    // Configure IGMP Robustness variable. The type is interface{} with range:
    // 2..10. The default value is 2.
    Robustness interface{}

    // Configure IGMP Traffic variables.
    Traffic Mld_Vrfs_Vrf_Traffic

    // Inheritable Defaults.
    InheritableDefaults Mld_Vrfs_Vrf_InheritableDefaults

    // IGMP Source specific mode.
    SsmAccessGroups Mld_Vrfs_Vrf_SsmAccessGroups

    // Configure IGMP State Limits.
    Maximum Mld_Vrfs_Vrf_Maximum

    // Interface-level configuration.
    Interfaces Mld_Vrfs_Vrf_Interfaces
}

func (vrf *Mld_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Mld_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Mld_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ssmdns-query-group" { return "SsmdnsQueryGroup" }
    if yname == "robustness" { return "Robustness" }
    if yname == "traffic" { return "Traffic" }
    if yname == "inheritable-defaults" { return "InheritableDefaults" }
    if yname == "ssm-access-groups" { return "SsmAccessGroups" }
    if yname == "maximum" { return "Maximum" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (vrf *Mld_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Mld_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traffic" {
        return &vrf.Traffic
    }
    if childYangName == "inheritable-defaults" {
        return &vrf.InheritableDefaults
    }
    if childYangName == "ssm-access-groups" {
        return &vrf.SsmAccessGroups
    }
    if childYangName == "maximum" {
        return &vrf.Maximum
    }
    if childYangName == "interfaces" {
        return &vrf.Interfaces
    }
    return nil
}

func (vrf *Mld_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["traffic"] = &vrf.Traffic
    children["inheritable-defaults"] = &vrf.InheritableDefaults
    children["ssm-access-groups"] = &vrf.SsmAccessGroups
    children["maximum"] = &vrf.Maximum
    children["interfaces"] = &vrf.Interfaces
    return children
}

func (vrf *Mld_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["ssmdns-query-group"] = vrf.SsmdnsQueryGroup
    leafs["robustness"] = vrf.Robustness
    return leafs
}

func (vrf *Mld_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Mld_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Mld_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Mld_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Mld_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Mld_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Mld_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Mld_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Mld_Vrfs_Vrf_Traffic
// Configure IGMP Traffic variables
type Mld_Vrfs_Vrf_Traffic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the route-policy profile. The type is string with length: 1..64.
    Profile interface{}
}

func (traffic *Mld_Vrfs_Vrf_Traffic) GetFilter() yfilter.YFilter { return traffic.YFilter }

func (traffic *Mld_Vrfs_Vrf_Traffic) SetFilter(yf yfilter.YFilter) { traffic.YFilter = yf }

func (traffic *Mld_Vrfs_Vrf_Traffic) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (traffic *Mld_Vrfs_Vrf_Traffic) GetSegmentPath() string {
    return "traffic"
}

func (traffic *Mld_Vrfs_Vrf_Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (traffic *Mld_Vrfs_Vrf_Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (traffic *Mld_Vrfs_Vrf_Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile"] = traffic.Profile
    return leafs
}

func (traffic *Mld_Vrfs_Vrf_Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (traffic *Mld_Vrfs_Vrf_Traffic) GetYangName() string { return "traffic" }

func (traffic *Mld_Vrfs_Vrf_Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traffic *Mld_Vrfs_Vrf_Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traffic *Mld_Vrfs_Vrf_Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traffic *Mld_Vrfs_Vrf_Traffic) SetParent(parent types.Entity) { traffic.parent = parent }

func (traffic *Mld_Vrfs_Vrf_Traffic) GetParent() types.Entity { return traffic.parent }

func (traffic *Mld_Vrfs_Vrf_Traffic) GetParentYangName() string { return "vrf" }

// Mld_Vrfs_Vrf_InheritableDefaults
// Inheritable Defaults
type Mld_Vrfs_Vrf_InheritableDefaults struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGMP previous querier timeout. The type is interface{} with range: 60..300.
    // Units are second.
    QueryTimeout interface{}

    // Access list specifying access group range. The type is string with length:
    // 1..64.
    AccessGroup interface{}

    // Query response value in seconds. The type is interface{} with range: 1..12.
    // Units are second. The default value is 10.
    QueryMaxResponseTime interface{}

    // Version number. The type is interface{} with range: 1..3. The default value
    // is 3.
    Version interface{}

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. The default value is true.
    RouterEnable interface{}

    // Query interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second. The default value is 60.
    QueryInterval interface{}

    // Configure maximum number of groups accepted per interface by this router.
    MaximumGroupsPerInterfaceOor Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor

    // IGMPv3 explicit host tracking.
    ExplicitTracking Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking
}

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetFilter() yfilter.YFilter { return inheritableDefaults.YFilter }

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) SetFilter(yf yfilter.YFilter) { inheritableDefaults.YFilter = yf }

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetGoName(yname string) string {
    if yname == "query-timeout" { return "QueryTimeout" }
    if yname == "access-group" { return "AccessGroup" }
    if yname == "query-max-response-time" { return "QueryMaxResponseTime" }
    if yname == "version" { return "Version" }
    if yname == "router-enable" { return "RouterEnable" }
    if yname == "query-interval" { return "QueryInterval" }
    if yname == "maximum-groups-per-interface-oor" { return "MaximumGroupsPerInterfaceOor" }
    if yname == "explicit-tracking" { return "ExplicitTracking" }
    return ""
}

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetSegmentPath() string {
    return "inheritable-defaults"
}

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "maximum-groups-per-interface-oor" {
        return &inheritableDefaults.MaximumGroupsPerInterfaceOor
    }
    if childYangName == "explicit-tracking" {
        return &inheritableDefaults.ExplicitTracking
    }
    return nil
}

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["maximum-groups-per-interface-oor"] = &inheritableDefaults.MaximumGroupsPerInterfaceOor
    children["explicit-tracking"] = &inheritableDefaults.ExplicitTracking
    return children
}

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["query-timeout"] = inheritableDefaults.QueryTimeout
    leafs["access-group"] = inheritableDefaults.AccessGroup
    leafs["query-max-response-time"] = inheritableDefaults.QueryMaxResponseTime
    leafs["version"] = inheritableDefaults.Version
    leafs["router-enable"] = inheritableDefaults.RouterEnable
    leafs["query-interval"] = inheritableDefaults.QueryInterval
    return leafs
}

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetBundleName() string { return "cisco_ios_xr" }

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetYangName() string { return "inheritable-defaults" }

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) SetParent(parent types.Entity) { inheritableDefaults.parent = parent }

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetParent() types.Entity { return inheritableDefaults.parent }

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetParentYangName() string { return "vrf" }

// Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of groups accepted per interface by this router. The type is
    // interface{} with range: 1..40000. This attribute is mandatory.
    MaximumGroups interface{}

    // WarningThreshold for number of groups accepted per interface by this
    // router. The type is interface{} with range: 1..40000. The default value is
    // 25000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetFilter() yfilter.YFilter { return maximumGroupsPerInterfaceOor.YFilter }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) SetFilter(yf yfilter.YFilter) { maximumGroupsPerInterfaceOor.YFilter = yf }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetGoName(yname string) string {
    if yname == "maximum-groups" { return "MaximumGroups" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetSegmentPath() string {
    return "maximum-groups-per-interface-oor"
}

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-groups"] = maximumGroupsPerInterfaceOor.MaximumGroups
    leafs["warning-threshold"] = maximumGroupsPerInterfaceOor.WarningThreshold
    leafs["access-list-name"] = maximumGroupsPerInterfaceOor.AccessListName
    return leafs
}

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetBundleName() string { return "cisco_ios_xr" }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetYangName() string { return "maximum-groups-per-interface-oor" }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) SetParent(parent types.Entity) { maximumGroupsPerInterfaceOor.parent = parent }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetParent() types.Entity { return maximumGroupsPerInterfaceOor.parent }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetParentYangName() string { return "inheritable-defaults" }

// Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetFilter() yfilter.YFilter { return explicitTracking.YFilter }

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) SetFilter(yf yfilter.YFilter) { explicitTracking.YFilter = yf }

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetSegmentPath() string {
    return "explicit-tracking"
}

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = explicitTracking.Enable
    leafs["access-list-name"] = explicitTracking.AccessListName
    return leafs
}

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetBundleName() string { return "cisco_ios_xr" }

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetYangName() string { return "explicit-tracking" }

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) SetParent(parent types.Entity) { explicitTracking.parent = parent }

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetParent() types.Entity { return explicitTracking.parent }

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetParentYangName() string { return "inheritable-defaults" }

// Mld_Vrfs_Vrf_SsmAccessGroups
// IGMP Source specific mode
type Mld_Vrfs_Vrf_SsmAccessGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSM static Access Group. The type is slice of
    // Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup.
    SsmAccessGroup []Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup
}

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetFilter() yfilter.YFilter { return ssmAccessGroups.YFilter }

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) SetFilter(yf yfilter.YFilter) { ssmAccessGroups.YFilter = yf }

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetGoName(yname string) string {
    if yname == "ssm-access-group" { return "SsmAccessGroup" }
    return ""
}

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetSegmentPath() string {
    return "ssm-access-groups"
}

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ssm-access-group" {
        for _, c := range ssmAccessGroups.SsmAccessGroup {
            if ssmAccessGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup{}
        ssmAccessGroups.SsmAccessGroup = append(ssmAccessGroups.SsmAccessGroup, child)
        return &ssmAccessGroups.SsmAccessGroup[len(ssmAccessGroups.SsmAccessGroup)-1]
    }
    return nil
}

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ssmAccessGroups.SsmAccessGroup {
        children[ssmAccessGroups.SsmAccessGroup[i].GetSegmentPath()] = &ssmAccessGroups.SsmAccessGroup[i]
    }
    return children
}

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetBundleName() string { return "cisco_ios_xr" }

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetYangName() string { return "ssm-access-groups" }

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) SetParent(parent types.Entity) { ssmAccessGroups.parent = parent }

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetParent() types.Entity { return ssmAccessGroups.parent }

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetParentYangName() string { return "vrf" }

// Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup
// SSM static Access Group
type Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Access list specifying access group. The type is string with length: 1..64.
    // This attribute is mandatory.
    AccessListName interface{}
}

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetFilter() yfilter.YFilter { return ssmAccessGroup.YFilter }

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) SetFilter(yf yfilter.YFilter) { ssmAccessGroup.YFilter = yf }

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetSegmentPath() string {
    return "ssm-access-group" + "[source-address='" + fmt.Sprintf("%v", ssmAccessGroup.SourceAddress) + "']"
}

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = ssmAccessGroup.SourceAddress
    leafs["access-list-name"] = ssmAccessGroup.AccessListName
    return leafs
}

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetBundleName() string { return "cisco_ios_xr" }

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetYangName() string { return "ssm-access-group" }

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) SetParent(parent types.Entity) { ssmAccessGroup.parent = parent }

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetParent() types.Entity { return ssmAccessGroup.parent }

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetParentYangName() string { return "ssm-access-groups" }

// Mld_Vrfs_Vrf_Maximum
// Configure IGMP State Limits
type Mld_Vrfs_Vrf_Maximum struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure maximum number of groups accepted by this router. The type is
    // interface{} with range: 1..75000. The default value is 50000.
    MaximumGroups interface{}
}

func (maximum *Mld_Vrfs_Vrf_Maximum) GetFilter() yfilter.YFilter { return maximum.YFilter }

func (maximum *Mld_Vrfs_Vrf_Maximum) SetFilter(yf yfilter.YFilter) { maximum.YFilter = yf }

func (maximum *Mld_Vrfs_Vrf_Maximum) GetGoName(yname string) string {
    if yname == "maximum-groups" { return "MaximumGroups" }
    return ""
}

func (maximum *Mld_Vrfs_Vrf_Maximum) GetSegmentPath() string {
    return "maximum"
}

func (maximum *Mld_Vrfs_Vrf_Maximum) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximum *Mld_Vrfs_Vrf_Maximum) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximum *Mld_Vrfs_Vrf_Maximum) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-groups"] = maximum.MaximumGroups
    return leafs
}

func (maximum *Mld_Vrfs_Vrf_Maximum) GetBundleName() string { return "cisco_ios_xr" }

func (maximum *Mld_Vrfs_Vrf_Maximum) GetYangName() string { return "maximum" }

func (maximum *Mld_Vrfs_Vrf_Maximum) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximum *Mld_Vrfs_Vrf_Maximum) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximum *Mld_Vrfs_Vrf_Maximum) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximum *Mld_Vrfs_Vrf_Maximum) SetParent(parent types.Entity) { maximum.parent = parent }

func (maximum *Mld_Vrfs_Vrf_Maximum) GetParent() types.Entity { return maximum.parent }

func (maximum *Mld_Vrfs_Vrf_Maximum) GetParentYangName() string { return "vrf" }

// Mld_Vrfs_Vrf_Interfaces
// Interface-level configuration
type Mld_Vrfs_Vrf_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Mld_Vrfs_Vrf_Interfaces_Interface.
    Interface []Mld_Vrfs_Vrf_Interfaces_Interface
}

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Mld_Vrfs_Vrf_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_Vrfs_Vrf_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Mld_Vrfs_Vrf_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetParentYangName() string { return "vrf" }

// Mld_Vrfs_Vrf_Interfaces_Interface
// The name of the interface
type Mld_Vrfs_Vrf_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // IGMP previous querier timeout. The type is interface{} with range: 60..300.
    // Units are second.
    QueryTimeout interface{}

    // Access list specifying access group range. The type is string with length:
    // 1..64.
    AccessGroup interface{}

    // Query response value in seconds. The type is interface{} with range: 1..12.
    // Units are second. The default value is 10.
    QueryMaxResponseTime interface{}

    // Version number. The type is interface{} with range: 1..3. The default value
    // is 3.
    Version interface{}

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. The default value is true.
    RouterEnable interface{}

    // Query interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second. The default value is 60.
    QueryInterval interface{}

    // IGMP join multicast group.
    JoinGroups Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups

    // IGMP static multicast group.
    StaticGroupGroupAddresses Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses

    // Configure maximum number of groups accepted per interface by this router.
    MaximumGroupsPerInterfaceOor Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor

    // IGMPv3 explicit host tracking.
    ExplicitTracking Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking
}

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "query-timeout" { return "QueryTimeout" }
    if yname == "access-group" { return "AccessGroup" }
    if yname == "query-max-response-time" { return "QueryMaxResponseTime" }
    if yname == "version" { return "Version" }
    if yname == "router-enable" { return "RouterEnable" }
    if yname == "query-interval" { return "QueryInterval" }
    if yname == "join-groups" { return "JoinGroups" }
    if yname == "static-group-group-addresses" { return "StaticGroupGroupAddresses" }
    if yname == "maximum-groups-per-interface-oor" { return "MaximumGroupsPerInterfaceOor" }
    if yname == "explicit-tracking" { return "ExplicitTracking" }
    return ""
}

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "join-groups" {
        return &self.JoinGroups
    }
    if childYangName == "static-group-group-addresses" {
        return &self.StaticGroupGroupAddresses
    }
    if childYangName == "maximum-groups-per-interface-oor" {
        return &self.MaximumGroupsPerInterfaceOor
    }
    if childYangName == "explicit-tracking" {
        return &self.ExplicitTracking
    }
    return nil
}

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["join-groups"] = &self.JoinGroups
    children["static-group-group-addresses"] = &self.StaticGroupGroupAddresses
    children["maximum-groups-per-interface-oor"] = &self.MaximumGroupsPerInterfaceOor
    children["explicit-tracking"] = &self.ExplicitTracking
    return children
}

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["query-timeout"] = self.QueryTimeout
    leafs["access-group"] = self.AccessGroup
    leafs["query-max-response-time"] = self.QueryMaxResponseTime
    leafs["version"] = self.Version
    leafs["router-enable"] = self.RouterEnable
    leafs["query-interval"] = self.QueryInterval
    return leafs
}

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups
// IGMP join multicast group
// This type is a presence type.
type Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP group address and optional source address to include. The type is slice
    // of Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup.
    JoinGroup []Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup

    // IP group address and optional source address to include. The type is slice
    // of Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress.
    JoinGroupSourceAddress []Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
}

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetFilter() yfilter.YFilter { return joinGroups.YFilter }

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) SetFilter(yf yfilter.YFilter) { joinGroups.YFilter = yf }

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetGoName(yname string) string {
    if yname == "join-group" { return "JoinGroup" }
    if yname == "join-group-source-address" { return "JoinGroupSourceAddress" }
    return ""
}

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetSegmentPath() string {
    return "join-groups"
}

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "join-group" {
        for _, c := range joinGroups.JoinGroup {
            if joinGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup{}
        joinGroups.JoinGroup = append(joinGroups.JoinGroup, child)
        return &joinGroups.JoinGroup[len(joinGroups.JoinGroup)-1]
    }
    if childYangName == "join-group-source-address" {
        for _, c := range joinGroups.JoinGroupSourceAddress {
            if joinGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress{}
        joinGroups.JoinGroupSourceAddress = append(joinGroups.JoinGroupSourceAddress, child)
        return &joinGroups.JoinGroupSourceAddress[len(joinGroups.JoinGroupSourceAddress)-1]
    }
    return nil
}

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range joinGroups.JoinGroup {
        children[joinGroups.JoinGroup[i].GetSegmentPath()] = &joinGroups.JoinGroup[i]
    }
    for i := range joinGroups.JoinGroupSourceAddress {
        children[joinGroups.JoinGroupSourceAddress[i].GetSegmentPath()] = &joinGroups.JoinGroupSourceAddress[i]
    }
    return children
}

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetBundleName() string { return "cisco_ios_xr" }

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetYangName() string { return "join-groups" }

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) SetParent(parent types.Entity) { joinGroups.parent = parent }

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetParent() types.Entity { return joinGroups.parent }

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetParentYangName() string { return "interface" }

// Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // Filter mode. The type is IgmpFilter. This attribute is mandatory.
    Mode interface{}
}

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetFilter() yfilter.YFilter { return joinGroup.YFilter }

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) SetFilter(yf yfilter.YFilter) { joinGroup.YFilter = yf }

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetSegmentPath() string {
    return "join-group" + "[group-address='" + fmt.Sprintf("%v", joinGroup.GroupAddress) + "']"
}

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = joinGroup.GroupAddress
    leafs["mode"] = joinGroup.Mode
    return leafs
}

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetBundleName() string { return "cisco_ios_xr" }

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetYangName() string { return "join-group" }

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) SetParent(parent types.Entity) { joinGroup.parent = parent }

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetParent() types.Entity { return joinGroup.parent }

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetParentYangName() string { return "join-groups" }

// Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Optional IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Filter mode. The type is IgmpFilter. This attribute is mandatory.
    Mode interface{}
}

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetFilter() yfilter.YFilter { return joinGroupSourceAddress.YFilter }

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) SetFilter(yf yfilter.YFilter) { joinGroupSourceAddress.YFilter = yf }

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetSegmentPath() string {
    return "join-group-source-address" + "[group-address='" + fmt.Sprintf("%v", joinGroupSourceAddress.GroupAddress) + "']" + "[source-address='" + fmt.Sprintf("%v", joinGroupSourceAddress.SourceAddress) + "']"
}

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = joinGroupSourceAddress.GroupAddress
    leafs["source-address"] = joinGroupSourceAddress.SourceAddress
    leafs["mode"] = joinGroupSourceAddress.Mode
    return leafs
}

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetBundleName() string { return "cisco_ios_xr" }

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetYangName() string { return "join-group-source-address" }

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) SetParent(parent types.Entity) { joinGroupSourceAddress.parent = parent }

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetParent() types.Entity { return joinGroupSourceAddress.parent }

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetParentYangName() string { return "join-groups" }

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses
// IGMP static multicast group
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress.
    StaticGroupGroupAddress []Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress.
    StaticGroupGroupAddressSourceAddress []Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask.
    StaticGroupGroupAddressSourceAddressSourceAddressMask []Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask.
    StaticGroupGroupAddressGroupAddressMask []Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress.
    StaticGroupGroupAddressGroupAddressMaskSourceAddress []Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.
    StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask []Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
}

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetFilter() yfilter.YFilter { return staticGroupGroupAddresses.YFilter }

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddresses.YFilter = yf }

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetGoName(yname string) string {
    if yname == "static-group-group-address" { return "StaticGroupGroupAddress" }
    if yname == "static-group-group-address-source-address" { return "StaticGroupGroupAddressSourceAddress" }
    if yname == "static-group-group-address-source-address-source-address-mask" { return "StaticGroupGroupAddressSourceAddressSourceAddressMask" }
    if yname == "static-group-group-address-group-address-mask" { return "StaticGroupGroupAddressGroupAddressMask" }
    if yname == "static-group-group-address-group-address-mask-source-address" { return "StaticGroupGroupAddressGroupAddressMaskSourceAddress" }
    if yname == "static-group-group-address-group-address-mask-source-address-source-address-mask" { return "StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask" }
    return ""
}

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetSegmentPath() string {
    return "static-group-group-addresses"
}

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-group-group-address" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddress {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress{}
        staticGroupGroupAddresses.StaticGroupGroupAddress = append(staticGroupGroupAddresses.StaticGroupGroupAddress, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddress[len(staticGroupGroupAddresses.StaticGroupGroupAddress)-1]
    }
    if childYangName == "static-group-group-address-source-address" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress{}
        staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress = append(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[len(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress)-1]
    }
    if childYangName == "static-group-group-address-source-address-source-address-mask" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask{}
        staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask = append(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[len(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask)-1]
    }
    if childYangName == "static-group-group-address-group-address-mask" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask{}
        staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask = append(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[len(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask)-1]
    }
    if childYangName == "static-group-group-address-group-address-mask-source-address" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress{}
        staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress = append(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[len(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress)-1]
    }
    if childYangName == "static-group-group-address-group-address-mask-source-address-source-address-mask" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask{}
        staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask = append(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[len(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask)-1]
    }
    return nil
}

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddress {
        children[staticGroupGroupAddresses.StaticGroupGroupAddress[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddress[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i]
    }
    return children
}

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetYangName() string { return "static-group-group-addresses" }

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) SetParent(parent types.Entity) { staticGroupGroupAddresses.parent = parent }

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetParent() types.Entity { return staticGroupGroupAddresses.parent }

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetParentYangName() string { return "interface" }

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetFilter() yfilter.YFilter { return staticGroupGroupAddress.YFilter }

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddress.YFilter = yf }

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetSegmentPath() string {
    return "static-group-group-address" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddress.GroupAddress) + "']"
}

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddress.GroupAddress
    leafs["group-count"] = staticGroupGroupAddress.GroupCount
    leafs["source-count"] = staticGroupGroupAddress.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddress.SuppressReport
    return leafs
}

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetYangName() string { return "static-group-group-address" }

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) SetParent(parent types.Entity) { staticGroupGroupAddress.parent = parent }

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetParent() types.Entity { return staticGroupGroupAddress.parent }

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetParentYangName() string { return "static-group-group-addresses" }

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetFilter() yfilter.YFilter { return staticGroupGroupAddressSourceAddress.YFilter }

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressSourceAddress.YFilter = yf }

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetSegmentPath() string {
    return "static-group-group-address-source-address" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddress.GroupAddress) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddress.SourceAddress) + "']"
}

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressSourceAddress.GroupAddress
    leafs["source-address"] = staticGroupGroupAddressSourceAddress.SourceAddress
    leafs["group-count"] = staticGroupGroupAddressSourceAddress.GroupCount
    leafs["source-count"] = staticGroupGroupAddressSourceAddress.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressSourceAddress.SuppressReport
    return leafs
}

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetYangName() string { return "static-group-group-address-source-address" }

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) SetParent(parent types.Entity) { staticGroupGroupAddressSourceAddress.parent = parent }

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetParent() types.Entity { return staticGroupGroupAddressSourceAddress.parent }

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetParentYangName() string { return "static-group-group-addresses" }

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Mask for Source Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddressMask interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetFilter() yfilter.YFilter { return staticGroupGroupAddressSourceAddressSourceAddressMask.YFilter }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressSourceAddressSourceAddressMask.YFilter = yf }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "source-address-mask" { return "SourceAddressMask" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetSegmentPath() string {
    return "static-group-group-address-source-address-source-address-mask" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress) + "']" + "[source-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask) + "']"
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress
    leafs["source-address"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress
    leafs["source-address-mask"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask
    leafs["group-count"] = staticGroupGroupAddressSourceAddressSourceAddressMask.GroupCount
    leafs["source-count"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SuppressReport
    return leafs
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetYangName() string { return "static-group-group-address-source-address-source-address-mask" }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) SetParent(parent types.Entity) { staticGroupGroupAddressSourceAddressSourceAddressMask.parent = parent }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetParent() types.Entity { return staticGroupGroupAddressSourceAddressSourceAddressMask.parent }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetParentYangName() string { return "static-group-group-addresses" }

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Mask for Group Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddressMask interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetFilter() yfilter.YFilter { return staticGroupGroupAddressGroupAddressMask.YFilter }

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressGroupAddressMask.YFilter = yf }

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-address-mask" { return "GroupAddressMask" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetSegmentPath() string {
    return "static-group-group-address-group-address-mask" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMask.GroupAddress) + "']" + "[group-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMask.GroupAddressMask) + "']"
}

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressGroupAddressMask.GroupAddress
    leafs["group-address-mask"] = staticGroupGroupAddressGroupAddressMask.GroupAddressMask
    leafs["group-count"] = staticGroupGroupAddressGroupAddressMask.GroupCount
    leafs["source-count"] = staticGroupGroupAddressGroupAddressMask.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressGroupAddressMask.SuppressReport
    return leafs
}

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetYangName() string { return "static-group-group-address-group-address-mask" }

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) SetParent(parent types.Entity) { staticGroupGroupAddressGroupAddressMask.parent = parent }

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetParent() types.Entity { return staticGroupGroupAddressGroupAddressMask.parent }

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetParentYangName() string { return "static-group-group-addresses" }

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Mask for Group Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddressMask interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetFilter() yfilter.YFilter { return staticGroupGroupAddressGroupAddressMaskSourceAddress.YFilter }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressGroupAddressMaskSourceAddress.YFilter = yf }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-address-mask" { return "GroupAddressMask" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetSegmentPath() string {
    return "static-group-group-address-group-address-mask-source-address" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress) + "']" + "[group-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress) + "']"
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress
    leafs["group-address-mask"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask
    leafs["source-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress
    leafs["group-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupCount
    leafs["source-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.SuppressReport
    return leafs
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetYangName() string { return "static-group-group-address-group-address-mask-source-address" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) SetParent(parent types.Entity) { staticGroupGroupAddressGroupAddressMaskSourceAddress.parent = parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetParent() types.Entity { return staticGroupGroupAddressGroupAddressMaskSourceAddress.parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetParentYangName() string { return "static-group-group-addresses" }

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Mask for Group Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddressMask interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Mask for Source Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddressMask interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetFilter() yfilter.YFilter { return staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.YFilter }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.YFilter = yf }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-address-mask" { return "GroupAddressMask" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "source-address-mask" { return "SourceAddressMask" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetSegmentPath() string {
    return "static-group-group-address-group-address-mask-source-address-source-address-mask" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress) + "']" + "[group-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress) + "']" + "[source-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask) + "']"
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress
    leafs["group-address-mask"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask
    leafs["source-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress
    leafs["source-address-mask"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask
    leafs["group-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupCount
    leafs["source-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SuppressReport
    return leafs
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetYangName() string { return "static-group-group-address-group-address-mask-source-address-source-address-mask" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) SetParent(parent types.Entity) { staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.parent = parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetParent() types.Entity { return staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetParentYangName() string { return "static-group-group-addresses" }

// Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of groups accepted per interface by this router. The type is
    // interface{} with range: 1..40000. This attribute is mandatory.
    MaximumGroups interface{}

    // WarningThreshold for number of groups accepted per interface by this
    // router. The type is interface{} with range: 1..40000. The default value is
    // 25000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetFilter() yfilter.YFilter { return maximumGroupsPerInterfaceOor.YFilter }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) SetFilter(yf yfilter.YFilter) { maximumGroupsPerInterfaceOor.YFilter = yf }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetGoName(yname string) string {
    if yname == "maximum-groups" { return "MaximumGroups" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetSegmentPath() string {
    return "maximum-groups-per-interface-oor"
}

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-groups"] = maximumGroupsPerInterfaceOor.MaximumGroups
    leafs["warning-threshold"] = maximumGroupsPerInterfaceOor.WarningThreshold
    leafs["access-list-name"] = maximumGroupsPerInterfaceOor.AccessListName
    return leafs
}

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetBundleName() string { return "cisco_ios_xr" }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetYangName() string { return "maximum-groups-per-interface-oor" }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) SetParent(parent types.Entity) { maximumGroupsPerInterfaceOor.parent = parent }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetParent() types.Entity { return maximumGroupsPerInterfaceOor.parent }

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetParentYangName() string { return "interface" }

// Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetFilter() yfilter.YFilter { return explicitTracking.YFilter }

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) SetFilter(yf yfilter.YFilter) { explicitTracking.YFilter = yf }

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetSegmentPath() string {
    return "explicit-tracking"
}

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = explicitTracking.Enable
    leafs["access-list-name"] = explicitTracking.AccessListName
    return leafs
}

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetBundleName() string { return "cisco_ios_xr" }

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetYangName() string { return "explicit-tracking" }

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) SetParent(parent types.Entity) { explicitTracking.parent = parent }

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetParent() types.Entity { return explicitTracking.parent }

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetParentYangName() string { return "interface" }

// Mld_DefaultContext
// Default Context
// This type is a presence type.
type Mld_DefaultContext struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable SSM mapping using DNS Query. The type is interface{}.
    SsmdnsQueryGroup interface{}

    // Configure IGMP Robustness variable. The type is interface{} with range:
    // 2..10. The default value is 2.
    Robustness interface{}

    // Configure NSF specific options.
    Nsf Mld_DefaultContext_Nsf

    // Configure IGMP QoS shapers for subscriber interfaces.
    UnicastQosAdjust Mld_DefaultContext_UnicastQosAdjust

    // Configure IGMP accounting for logging join/leave records.
    Accounting Mld_DefaultContext_Accounting

    // Configure IGMP Traffic variables.
    Traffic Mld_DefaultContext_Traffic

    // Inheritable Defaults.
    InheritableDefaults Mld_DefaultContext_InheritableDefaults

    // IGMP Source specific mode.
    SsmAccessGroups Mld_DefaultContext_SsmAccessGroups

    // Configure IGMP State Limits.
    Maximum Mld_DefaultContext_Maximum

    // Interface-level configuration.
    Interfaces Mld_DefaultContext_Interfaces
}

func (defaultContext *Mld_DefaultContext) GetFilter() yfilter.YFilter { return defaultContext.YFilter }

func (defaultContext *Mld_DefaultContext) SetFilter(yf yfilter.YFilter) { defaultContext.YFilter = yf }

func (defaultContext *Mld_DefaultContext) GetGoName(yname string) string {
    if yname == "ssmdns-query-group" { return "SsmdnsQueryGroup" }
    if yname == "robustness" { return "Robustness" }
    if yname == "nsf" { return "Nsf" }
    if yname == "unicast-qos-adjust" { return "UnicastQosAdjust" }
    if yname == "accounting" { return "Accounting" }
    if yname == "traffic" { return "Traffic" }
    if yname == "inheritable-defaults" { return "InheritableDefaults" }
    if yname == "ssm-access-groups" { return "SsmAccessGroups" }
    if yname == "maximum" { return "Maximum" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (defaultContext *Mld_DefaultContext) GetSegmentPath() string {
    return "default-context"
}

func (defaultContext *Mld_DefaultContext) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nsf" {
        return &defaultContext.Nsf
    }
    if childYangName == "unicast-qos-adjust" {
        return &defaultContext.UnicastQosAdjust
    }
    if childYangName == "accounting" {
        return &defaultContext.Accounting
    }
    if childYangName == "traffic" {
        return &defaultContext.Traffic
    }
    if childYangName == "inheritable-defaults" {
        return &defaultContext.InheritableDefaults
    }
    if childYangName == "ssm-access-groups" {
        return &defaultContext.SsmAccessGroups
    }
    if childYangName == "maximum" {
        return &defaultContext.Maximum
    }
    if childYangName == "interfaces" {
        return &defaultContext.Interfaces
    }
    return nil
}

func (defaultContext *Mld_DefaultContext) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nsf"] = &defaultContext.Nsf
    children["unicast-qos-adjust"] = &defaultContext.UnicastQosAdjust
    children["accounting"] = &defaultContext.Accounting
    children["traffic"] = &defaultContext.Traffic
    children["inheritable-defaults"] = &defaultContext.InheritableDefaults
    children["ssm-access-groups"] = &defaultContext.SsmAccessGroups
    children["maximum"] = &defaultContext.Maximum
    children["interfaces"] = &defaultContext.Interfaces
    return children
}

func (defaultContext *Mld_DefaultContext) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ssmdns-query-group"] = defaultContext.SsmdnsQueryGroup
    leafs["robustness"] = defaultContext.Robustness
    return leafs
}

func (defaultContext *Mld_DefaultContext) GetBundleName() string { return "cisco_ios_xr" }

func (defaultContext *Mld_DefaultContext) GetYangName() string { return "default-context" }

func (defaultContext *Mld_DefaultContext) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultContext *Mld_DefaultContext) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultContext *Mld_DefaultContext) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultContext *Mld_DefaultContext) SetParent(parent types.Entity) { defaultContext.parent = parent }

func (defaultContext *Mld_DefaultContext) GetParent() types.Entity { return defaultContext.parent }

func (defaultContext *Mld_DefaultContext) GetParentYangName() string { return "mld" }

// Mld_DefaultContext_Nsf
// Configure NSF specific options
type Mld_DefaultContext_Nsf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum time for IGMP NSF mode in seconds. The type is interface{} with
    // range: 10..3600. Units are second. The default value is 60.
    Lifetime interface{}
}

func (nsf *Mld_DefaultContext_Nsf) GetFilter() yfilter.YFilter { return nsf.YFilter }

func (nsf *Mld_DefaultContext_Nsf) SetFilter(yf yfilter.YFilter) { nsf.YFilter = yf }

func (nsf *Mld_DefaultContext_Nsf) GetGoName(yname string) string {
    if yname == "lifetime" { return "Lifetime" }
    return ""
}

func (nsf *Mld_DefaultContext_Nsf) GetSegmentPath() string {
    return "nsf"
}

func (nsf *Mld_DefaultContext_Nsf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nsf *Mld_DefaultContext_Nsf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nsf *Mld_DefaultContext_Nsf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lifetime"] = nsf.Lifetime
    return leafs
}

func (nsf *Mld_DefaultContext_Nsf) GetBundleName() string { return "cisco_ios_xr" }

func (nsf *Mld_DefaultContext_Nsf) GetYangName() string { return "nsf" }

func (nsf *Mld_DefaultContext_Nsf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nsf *Mld_DefaultContext_Nsf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nsf *Mld_DefaultContext_Nsf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nsf *Mld_DefaultContext_Nsf) SetParent(parent types.Entity) { nsf.parent = parent }

func (nsf *Mld_DefaultContext_Nsf) GetParent() types.Entity { return nsf.parent }

func (nsf *Mld_DefaultContext_Nsf) GetParentYangName() string { return "default-context" }

// Mld_DefaultContext_UnicastQosAdjust
// Configure IGMP QoS shapers for subscriber
// interfaces
type Mld_DefaultContext_UnicastQosAdjust struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the QoS download interval (in milliseconds). The type is
    // interface{} with range: 10..500. Units are millisecond. The default value
    // is 100.
    DownloadInterval interface{}

    // Configure the QoS delay before programming (in seconds). The type is
    // interface{} with range: 0..10. Units are second. The default value is 1.
    AdjustmentDelay interface{}

    // Configure the QoS hold off time (in seconds). The type is interface{} with
    // range: 5..1800. Units are second. The default value is 180.
    HoldOff interface{}
}

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetFilter() yfilter.YFilter { return unicastQosAdjust.YFilter }

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) SetFilter(yf yfilter.YFilter) { unicastQosAdjust.YFilter = yf }

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetGoName(yname string) string {
    if yname == "download-interval" { return "DownloadInterval" }
    if yname == "adjustment-delay" { return "AdjustmentDelay" }
    if yname == "hold-off" { return "HoldOff" }
    return ""
}

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetSegmentPath() string {
    return "unicast-qos-adjust"
}

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["download-interval"] = unicastQosAdjust.DownloadInterval
    leafs["adjustment-delay"] = unicastQosAdjust.AdjustmentDelay
    leafs["hold-off"] = unicastQosAdjust.HoldOff
    return leafs
}

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetBundleName() string { return "cisco_ios_xr" }

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetYangName() string { return "unicast-qos-adjust" }

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) SetParent(parent types.Entity) { unicastQosAdjust.parent = parent }

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetParent() types.Entity { return unicastQosAdjust.parent }

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetParentYangName() string { return "default-context" }

// Mld_DefaultContext_Accounting
// Configure IGMP accounting for logging
// join/leave records
type Mld_DefaultContext_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure IGMP accounting Maximum History setting. The type is interface{}
    // with range: 0..365. Units are day. The default value is 0.
    MaxHistory interface{}
}

func (accounting *Mld_DefaultContext_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *Mld_DefaultContext_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *Mld_DefaultContext_Accounting) GetGoName(yname string) string {
    if yname == "max-history" { return "MaxHistory" }
    return ""
}

func (accounting *Mld_DefaultContext_Accounting) GetSegmentPath() string {
    return "accounting"
}

func (accounting *Mld_DefaultContext_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accounting *Mld_DefaultContext_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accounting *Mld_DefaultContext_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-history"] = accounting.MaxHistory
    return leafs
}

func (accounting *Mld_DefaultContext_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *Mld_DefaultContext_Accounting) GetYangName() string { return "accounting" }

func (accounting *Mld_DefaultContext_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *Mld_DefaultContext_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *Mld_DefaultContext_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *Mld_DefaultContext_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *Mld_DefaultContext_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *Mld_DefaultContext_Accounting) GetParentYangName() string { return "default-context" }

// Mld_DefaultContext_Traffic
// Configure IGMP Traffic variables
type Mld_DefaultContext_Traffic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the route-policy profile. The type is string with length: 1..64.
    Profile interface{}
}

func (traffic *Mld_DefaultContext_Traffic) GetFilter() yfilter.YFilter { return traffic.YFilter }

func (traffic *Mld_DefaultContext_Traffic) SetFilter(yf yfilter.YFilter) { traffic.YFilter = yf }

func (traffic *Mld_DefaultContext_Traffic) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (traffic *Mld_DefaultContext_Traffic) GetSegmentPath() string {
    return "traffic"
}

func (traffic *Mld_DefaultContext_Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (traffic *Mld_DefaultContext_Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (traffic *Mld_DefaultContext_Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile"] = traffic.Profile
    return leafs
}

func (traffic *Mld_DefaultContext_Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (traffic *Mld_DefaultContext_Traffic) GetYangName() string { return "traffic" }

func (traffic *Mld_DefaultContext_Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traffic *Mld_DefaultContext_Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traffic *Mld_DefaultContext_Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traffic *Mld_DefaultContext_Traffic) SetParent(parent types.Entity) { traffic.parent = parent }

func (traffic *Mld_DefaultContext_Traffic) GetParent() types.Entity { return traffic.parent }

func (traffic *Mld_DefaultContext_Traffic) GetParentYangName() string { return "default-context" }

// Mld_DefaultContext_InheritableDefaults
// Inheritable Defaults
type Mld_DefaultContext_InheritableDefaults struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGMP previous querier timeout. The type is interface{} with range: 60..300.
    // Units are second.
    QueryTimeout interface{}

    // Access list specifying access group range. The type is string with length:
    // 1..64.
    AccessGroup interface{}

    // Query response value in seconds. The type is interface{} with range: 1..12.
    // Units are second. The default value is 10.
    QueryMaxResponseTime interface{}

    // Version number. The type is interface{} with range: 1..3. The default value
    // is 3.
    Version interface{}

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. The default value is true.
    RouterEnable interface{}

    // Query interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second. The default value is 60.
    QueryInterval interface{}

    // Configure maximum number of groups accepted per interface by this router.
    MaximumGroupsPerInterfaceOor Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor

    // IGMPv3 explicit host tracking.
    ExplicitTracking Mld_DefaultContext_InheritableDefaults_ExplicitTracking
}

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetFilter() yfilter.YFilter { return inheritableDefaults.YFilter }

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) SetFilter(yf yfilter.YFilter) { inheritableDefaults.YFilter = yf }

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetGoName(yname string) string {
    if yname == "query-timeout" { return "QueryTimeout" }
    if yname == "access-group" { return "AccessGroup" }
    if yname == "query-max-response-time" { return "QueryMaxResponseTime" }
    if yname == "version" { return "Version" }
    if yname == "router-enable" { return "RouterEnable" }
    if yname == "query-interval" { return "QueryInterval" }
    if yname == "maximum-groups-per-interface-oor" { return "MaximumGroupsPerInterfaceOor" }
    if yname == "explicit-tracking" { return "ExplicitTracking" }
    return ""
}

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetSegmentPath() string {
    return "inheritable-defaults"
}

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "maximum-groups-per-interface-oor" {
        return &inheritableDefaults.MaximumGroupsPerInterfaceOor
    }
    if childYangName == "explicit-tracking" {
        return &inheritableDefaults.ExplicitTracking
    }
    return nil
}

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["maximum-groups-per-interface-oor"] = &inheritableDefaults.MaximumGroupsPerInterfaceOor
    children["explicit-tracking"] = &inheritableDefaults.ExplicitTracking
    return children
}

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["query-timeout"] = inheritableDefaults.QueryTimeout
    leafs["access-group"] = inheritableDefaults.AccessGroup
    leafs["query-max-response-time"] = inheritableDefaults.QueryMaxResponseTime
    leafs["version"] = inheritableDefaults.Version
    leafs["router-enable"] = inheritableDefaults.RouterEnable
    leafs["query-interval"] = inheritableDefaults.QueryInterval
    return leafs
}

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetBundleName() string { return "cisco_ios_xr" }

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetYangName() string { return "inheritable-defaults" }

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) SetParent(parent types.Entity) { inheritableDefaults.parent = parent }

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetParent() types.Entity { return inheritableDefaults.parent }

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetParentYangName() string { return "default-context" }

// Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of groups accepted per interface by this router. The type is
    // interface{} with range: 1..40000. This attribute is mandatory.
    MaximumGroups interface{}

    // WarningThreshold for number of groups accepted per interface by this
    // router. The type is interface{} with range: 1..40000. The default value is
    // 25000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetFilter() yfilter.YFilter { return maximumGroupsPerInterfaceOor.YFilter }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) SetFilter(yf yfilter.YFilter) { maximumGroupsPerInterfaceOor.YFilter = yf }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetGoName(yname string) string {
    if yname == "maximum-groups" { return "MaximumGroups" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetSegmentPath() string {
    return "maximum-groups-per-interface-oor"
}

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-groups"] = maximumGroupsPerInterfaceOor.MaximumGroups
    leafs["warning-threshold"] = maximumGroupsPerInterfaceOor.WarningThreshold
    leafs["access-list-name"] = maximumGroupsPerInterfaceOor.AccessListName
    return leafs
}

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetBundleName() string { return "cisco_ios_xr" }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetYangName() string { return "maximum-groups-per-interface-oor" }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) SetParent(parent types.Entity) { maximumGroupsPerInterfaceOor.parent = parent }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetParent() types.Entity { return maximumGroupsPerInterfaceOor.parent }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetParentYangName() string { return "inheritable-defaults" }

// Mld_DefaultContext_InheritableDefaults_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Mld_DefaultContext_InheritableDefaults_ExplicitTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetFilter() yfilter.YFilter { return explicitTracking.YFilter }

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) SetFilter(yf yfilter.YFilter) { explicitTracking.YFilter = yf }

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetSegmentPath() string {
    return "explicit-tracking"
}

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = explicitTracking.Enable
    leafs["access-list-name"] = explicitTracking.AccessListName
    return leafs
}

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetBundleName() string { return "cisco_ios_xr" }

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetYangName() string { return "explicit-tracking" }

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) SetParent(parent types.Entity) { explicitTracking.parent = parent }

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetParent() types.Entity { return explicitTracking.parent }

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetParentYangName() string { return "inheritable-defaults" }

// Mld_DefaultContext_SsmAccessGroups
// IGMP Source specific mode
type Mld_DefaultContext_SsmAccessGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSM static Access Group. The type is slice of
    // Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup.
    SsmAccessGroup []Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup
}

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetFilter() yfilter.YFilter { return ssmAccessGroups.YFilter }

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) SetFilter(yf yfilter.YFilter) { ssmAccessGroups.YFilter = yf }

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetGoName(yname string) string {
    if yname == "ssm-access-group" { return "SsmAccessGroup" }
    return ""
}

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetSegmentPath() string {
    return "ssm-access-groups"
}

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ssm-access-group" {
        for _, c := range ssmAccessGroups.SsmAccessGroup {
            if ssmAccessGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup{}
        ssmAccessGroups.SsmAccessGroup = append(ssmAccessGroups.SsmAccessGroup, child)
        return &ssmAccessGroups.SsmAccessGroup[len(ssmAccessGroups.SsmAccessGroup)-1]
    }
    return nil
}

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ssmAccessGroups.SsmAccessGroup {
        children[ssmAccessGroups.SsmAccessGroup[i].GetSegmentPath()] = &ssmAccessGroups.SsmAccessGroup[i]
    }
    return children
}

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetBundleName() string { return "cisco_ios_xr" }

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetYangName() string { return "ssm-access-groups" }

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) SetParent(parent types.Entity) { ssmAccessGroups.parent = parent }

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetParent() types.Entity { return ssmAccessGroups.parent }

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetParentYangName() string { return "default-context" }

// Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup
// SSM static Access Group
type Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Access list specifying access group. The type is string with length: 1..64.
    // This attribute is mandatory.
    AccessListName interface{}
}

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetFilter() yfilter.YFilter { return ssmAccessGroup.YFilter }

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) SetFilter(yf yfilter.YFilter) { ssmAccessGroup.YFilter = yf }

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetSegmentPath() string {
    return "ssm-access-group" + "[source-address='" + fmt.Sprintf("%v", ssmAccessGroup.SourceAddress) + "']"
}

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = ssmAccessGroup.SourceAddress
    leafs["access-list-name"] = ssmAccessGroup.AccessListName
    return leafs
}

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetBundleName() string { return "cisco_ios_xr" }

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetYangName() string { return "ssm-access-group" }

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) SetParent(parent types.Entity) { ssmAccessGroup.parent = parent }

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetParent() types.Entity { return ssmAccessGroup.parent }

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetParentYangName() string { return "ssm-access-groups" }

// Mld_DefaultContext_Maximum
// Configure IGMP State Limits
type Mld_DefaultContext_Maximum struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure maximum number of groups accepted by this router. The type is
    // interface{} with range: 1..75000. The default value is 50000.
    MaximumGroups interface{}
}

func (maximum *Mld_DefaultContext_Maximum) GetFilter() yfilter.YFilter { return maximum.YFilter }

func (maximum *Mld_DefaultContext_Maximum) SetFilter(yf yfilter.YFilter) { maximum.YFilter = yf }

func (maximum *Mld_DefaultContext_Maximum) GetGoName(yname string) string {
    if yname == "maximum-groups" { return "MaximumGroups" }
    return ""
}

func (maximum *Mld_DefaultContext_Maximum) GetSegmentPath() string {
    return "maximum"
}

func (maximum *Mld_DefaultContext_Maximum) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximum *Mld_DefaultContext_Maximum) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximum *Mld_DefaultContext_Maximum) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-groups"] = maximum.MaximumGroups
    return leafs
}

func (maximum *Mld_DefaultContext_Maximum) GetBundleName() string { return "cisco_ios_xr" }

func (maximum *Mld_DefaultContext_Maximum) GetYangName() string { return "maximum" }

func (maximum *Mld_DefaultContext_Maximum) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximum *Mld_DefaultContext_Maximum) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximum *Mld_DefaultContext_Maximum) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximum *Mld_DefaultContext_Maximum) SetParent(parent types.Entity) { maximum.parent = parent }

func (maximum *Mld_DefaultContext_Maximum) GetParent() types.Entity { return maximum.parent }

func (maximum *Mld_DefaultContext_Maximum) GetParentYangName() string { return "default-context" }

// Mld_DefaultContext_Interfaces
// Interface-level configuration
type Mld_DefaultContext_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Mld_DefaultContext_Interfaces_Interface.
    Interface []Mld_DefaultContext_Interfaces_Interface
}

func (interfaces *Mld_DefaultContext_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Mld_DefaultContext_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Mld_DefaultContext_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Mld_DefaultContext_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Mld_DefaultContext_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_DefaultContext_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Mld_DefaultContext_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Mld_DefaultContext_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Mld_DefaultContext_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Mld_DefaultContext_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Mld_DefaultContext_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Mld_DefaultContext_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Mld_DefaultContext_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Mld_DefaultContext_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Mld_DefaultContext_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Mld_DefaultContext_Interfaces) GetParentYangName() string { return "default-context" }

// Mld_DefaultContext_Interfaces_Interface
// The name of the interface
type Mld_DefaultContext_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // IGMP previous querier timeout. The type is interface{} with range: 60..300.
    // Units are second.
    QueryTimeout interface{}

    // Access list specifying access group range. The type is string with length:
    // 1..64.
    AccessGroup interface{}

    // Query response value in seconds. The type is interface{} with range: 1..12.
    // Units are second. The default value is 10.
    QueryMaxResponseTime interface{}

    // Version number. The type is interface{} with range: 1..3. The default value
    // is 3.
    Version interface{}

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. The default value is true.
    RouterEnable interface{}

    // Query interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second. The default value is 60.
    QueryInterval interface{}

    // IGMP join multicast group.
    JoinGroups Mld_DefaultContext_Interfaces_Interface_JoinGroups

    // IGMP static multicast group.
    StaticGroupGroupAddresses Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses

    // Configure maximum number of groups accepted per interface by this router.
    MaximumGroupsPerInterfaceOor Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor

    // IGMPv3 explicit host tracking.
    ExplicitTracking Mld_DefaultContext_Interfaces_Interface_ExplicitTracking
}

func (self *Mld_DefaultContext_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Mld_DefaultContext_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Mld_DefaultContext_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "query-timeout" { return "QueryTimeout" }
    if yname == "access-group" { return "AccessGroup" }
    if yname == "query-max-response-time" { return "QueryMaxResponseTime" }
    if yname == "version" { return "Version" }
    if yname == "router-enable" { return "RouterEnable" }
    if yname == "query-interval" { return "QueryInterval" }
    if yname == "join-groups" { return "JoinGroups" }
    if yname == "static-group-group-addresses" { return "StaticGroupGroupAddresses" }
    if yname == "maximum-groups-per-interface-oor" { return "MaximumGroupsPerInterfaceOor" }
    if yname == "explicit-tracking" { return "ExplicitTracking" }
    return ""
}

func (self *Mld_DefaultContext_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Mld_DefaultContext_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "join-groups" {
        return &self.JoinGroups
    }
    if childYangName == "static-group-group-addresses" {
        return &self.StaticGroupGroupAddresses
    }
    if childYangName == "maximum-groups-per-interface-oor" {
        return &self.MaximumGroupsPerInterfaceOor
    }
    if childYangName == "explicit-tracking" {
        return &self.ExplicitTracking
    }
    return nil
}

func (self *Mld_DefaultContext_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["join-groups"] = &self.JoinGroups
    children["static-group-group-addresses"] = &self.StaticGroupGroupAddresses
    children["maximum-groups-per-interface-oor"] = &self.MaximumGroupsPerInterfaceOor
    children["explicit-tracking"] = &self.ExplicitTracking
    return children
}

func (self *Mld_DefaultContext_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["query-timeout"] = self.QueryTimeout
    leafs["access-group"] = self.AccessGroup
    leafs["query-max-response-time"] = self.QueryMaxResponseTime
    leafs["version"] = self.Version
    leafs["router-enable"] = self.RouterEnable
    leafs["query-interval"] = self.QueryInterval
    return leafs
}

func (self *Mld_DefaultContext_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Mld_DefaultContext_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Mld_DefaultContext_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Mld_DefaultContext_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Mld_DefaultContext_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Mld_DefaultContext_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Mld_DefaultContext_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Mld_DefaultContext_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Mld_DefaultContext_Interfaces_Interface_JoinGroups
// IGMP join multicast group
// This type is a presence type.
type Mld_DefaultContext_Interfaces_Interface_JoinGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP group address and optional source address to include. The type is slice
    // of Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup.
    JoinGroup []Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress.
    JoinGroupSourceAddress []Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
}

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetFilter() yfilter.YFilter { return joinGroups.YFilter }

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) SetFilter(yf yfilter.YFilter) { joinGroups.YFilter = yf }

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetGoName(yname string) string {
    if yname == "join-group" { return "JoinGroup" }
    if yname == "join-group-source-address" { return "JoinGroupSourceAddress" }
    return ""
}

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetSegmentPath() string {
    return "join-groups"
}

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "join-group" {
        for _, c := range joinGroups.JoinGroup {
            if joinGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup{}
        joinGroups.JoinGroup = append(joinGroups.JoinGroup, child)
        return &joinGroups.JoinGroup[len(joinGroups.JoinGroup)-1]
    }
    if childYangName == "join-group-source-address" {
        for _, c := range joinGroups.JoinGroupSourceAddress {
            if joinGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress{}
        joinGroups.JoinGroupSourceAddress = append(joinGroups.JoinGroupSourceAddress, child)
        return &joinGroups.JoinGroupSourceAddress[len(joinGroups.JoinGroupSourceAddress)-1]
    }
    return nil
}

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range joinGroups.JoinGroup {
        children[joinGroups.JoinGroup[i].GetSegmentPath()] = &joinGroups.JoinGroup[i]
    }
    for i := range joinGroups.JoinGroupSourceAddress {
        children[joinGroups.JoinGroupSourceAddress[i].GetSegmentPath()] = &joinGroups.JoinGroupSourceAddress[i]
    }
    return children
}

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetBundleName() string { return "cisco_ios_xr" }

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetYangName() string { return "join-groups" }

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) SetParent(parent types.Entity) { joinGroups.parent = parent }

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetParent() types.Entity { return joinGroups.parent }

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetParentYangName() string { return "interface" }

// Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // Filter mode. The type is IgmpFilter. This attribute is mandatory.
    Mode interface{}
}

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetFilter() yfilter.YFilter { return joinGroup.YFilter }

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) SetFilter(yf yfilter.YFilter) { joinGroup.YFilter = yf }

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetSegmentPath() string {
    return "join-group" + "[group-address='" + fmt.Sprintf("%v", joinGroup.GroupAddress) + "']"
}

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = joinGroup.GroupAddress
    leafs["mode"] = joinGroup.Mode
    return leafs
}

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetBundleName() string { return "cisco_ios_xr" }

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetYangName() string { return "join-group" }

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) SetParent(parent types.Entity) { joinGroup.parent = parent }

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetParent() types.Entity { return joinGroup.parent }

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetParentYangName() string { return "join-groups" }

// Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Optional IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Filter mode. The type is IgmpFilter. This attribute is mandatory.
    Mode interface{}
}

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetFilter() yfilter.YFilter { return joinGroupSourceAddress.YFilter }

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) SetFilter(yf yfilter.YFilter) { joinGroupSourceAddress.YFilter = yf }

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetSegmentPath() string {
    return "join-group-source-address" + "[group-address='" + fmt.Sprintf("%v", joinGroupSourceAddress.GroupAddress) + "']" + "[source-address='" + fmt.Sprintf("%v", joinGroupSourceAddress.SourceAddress) + "']"
}

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = joinGroupSourceAddress.GroupAddress
    leafs["source-address"] = joinGroupSourceAddress.SourceAddress
    leafs["mode"] = joinGroupSourceAddress.Mode
    return leafs
}

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetBundleName() string { return "cisco_ios_xr" }

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetYangName() string { return "join-group-source-address" }

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) SetParent(parent types.Entity) { joinGroupSourceAddress.parent = parent }

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetParent() types.Entity { return joinGroupSourceAddress.parent }

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetParentYangName() string { return "join-groups" }

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses
// IGMP static multicast group
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress.
    StaticGroupGroupAddress []Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress.
    StaticGroupGroupAddressSourceAddress []Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask.
    StaticGroupGroupAddressSourceAddressSourceAddressMask []Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask.
    StaticGroupGroupAddressGroupAddressMask []Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress.
    StaticGroupGroupAddressGroupAddressMaskSourceAddress []Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.
    StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask []Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
}

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetFilter() yfilter.YFilter { return staticGroupGroupAddresses.YFilter }

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddresses.YFilter = yf }

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetGoName(yname string) string {
    if yname == "static-group-group-address" { return "StaticGroupGroupAddress" }
    if yname == "static-group-group-address-source-address" { return "StaticGroupGroupAddressSourceAddress" }
    if yname == "static-group-group-address-source-address-source-address-mask" { return "StaticGroupGroupAddressSourceAddressSourceAddressMask" }
    if yname == "static-group-group-address-group-address-mask" { return "StaticGroupGroupAddressGroupAddressMask" }
    if yname == "static-group-group-address-group-address-mask-source-address" { return "StaticGroupGroupAddressGroupAddressMaskSourceAddress" }
    if yname == "static-group-group-address-group-address-mask-source-address-source-address-mask" { return "StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask" }
    return ""
}

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetSegmentPath() string {
    return "static-group-group-addresses"
}

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-group-group-address" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddress {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress{}
        staticGroupGroupAddresses.StaticGroupGroupAddress = append(staticGroupGroupAddresses.StaticGroupGroupAddress, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddress[len(staticGroupGroupAddresses.StaticGroupGroupAddress)-1]
    }
    if childYangName == "static-group-group-address-source-address" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress{}
        staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress = append(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[len(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress)-1]
    }
    if childYangName == "static-group-group-address-source-address-source-address-mask" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask{}
        staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask = append(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[len(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask)-1]
    }
    if childYangName == "static-group-group-address-group-address-mask" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask{}
        staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask = append(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[len(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask)-1]
    }
    if childYangName == "static-group-group-address-group-address-mask-source-address" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress{}
        staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress = append(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[len(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress)-1]
    }
    if childYangName == "static-group-group-address-group-address-mask-source-address-source-address-mask" {
        for _, c := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask {
            if staticGroupGroupAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask{}
        staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask = append(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask, child)
        return &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[len(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask)-1]
    }
    return nil
}

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddress {
        children[staticGroupGroupAddresses.StaticGroupGroupAddress[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddress[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i]
    }
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask {
        children[staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i].GetSegmentPath()] = &staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i]
    }
    return children
}

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetYangName() string { return "static-group-group-addresses" }

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) SetParent(parent types.Entity) { staticGroupGroupAddresses.parent = parent }

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetParent() types.Entity { return staticGroupGroupAddresses.parent }

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetParentYangName() string { return "interface" }

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetFilter() yfilter.YFilter { return staticGroupGroupAddress.YFilter }

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddress.YFilter = yf }

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetSegmentPath() string {
    return "static-group-group-address" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddress.GroupAddress) + "']"
}

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddress.GroupAddress
    leafs["group-count"] = staticGroupGroupAddress.GroupCount
    leafs["source-count"] = staticGroupGroupAddress.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddress.SuppressReport
    return leafs
}

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetYangName() string { return "static-group-group-address" }

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) SetParent(parent types.Entity) { staticGroupGroupAddress.parent = parent }

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetParent() types.Entity { return staticGroupGroupAddress.parent }

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetParentYangName() string { return "static-group-group-addresses" }

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetFilter() yfilter.YFilter { return staticGroupGroupAddressSourceAddress.YFilter }

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressSourceAddress.YFilter = yf }

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetSegmentPath() string {
    return "static-group-group-address-source-address" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddress.GroupAddress) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddress.SourceAddress) + "']"
}

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressSourceAddress.GroupAddress
    leafs["source-address"] = staticGroupGroupAddressSourceAddress.SourceAddress
    leafs["group-count"] = staticGroupGroupAddressSourceAddress.GroupCount
    leafs["source-count"] = staticGroupGroupAddressSourceAddress.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressSourceAddress.SuppressReport
    return leafs
}

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetYangName() string { return "static-group-group-address-source-address" }

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) SetParent(parent types.Entity) { staticGroupGroupAddressSourceAddress.parent = parent }

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetParent() types.Entity { return staticGroupGroupAddressSourceAddress.parent }

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetParentYangName() string { return "static-group-group-addresses" }

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Mask for Source Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddressMask interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetFilter() yfilter.YFilter { return staticGroupGroupAddressSourceAddressSourceAddressMask.YFilter }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressSourceAddressSourceAddressMask.YFilter = yf }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "source-address-mask" { return "SourceAddressMask" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetSegmentPath() string {
    return "static-group-group-address-source-address-source-address-mask" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress) + "']" + "[source-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask) + "']"
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress
    leafs["source-address"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress
    leafs["source-address-mask"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask
    leafs["group-count"] = staticGroupGroupAddressSourceAddressSourceAddressMask.GroupCount
    leafs["source-count"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressSourceAddressSourceAddressMask.SuppressReport
    return leafs
}

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetYangName() string { return "static-group-group-address-source-address-source-address-mask" }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) SetParent(parent types.Entity) { staticGroupGroupAddressSourceAddressSourceAddressMask.parent = parent }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetParent() types.Entity { return staticGroupGroupAddressSourceAddressSourceAddressMask.parent }

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetParentYangName() string { return "static-group-group-addresses" }

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Mask for Group Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddressMask interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetFilter() yfilter.YFilter { return staticGroupGroupAddressGroupAddressMask.YFilter }

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressGroupAddressMask.YFilter = yf }

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-address-mask" { return "GroupAddressMask" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetSegmentPath() string {
    return "static-group-group-address-group-address-mask" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMask.GroupAddress) + "']" + "[group-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMask.GroupAddressMask) + "']"
}

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressGroupAddressMask.GroupAddress
    leafs["group-address-mask"] = staticGroupGroupAddressGroupAddressMask.GroupAddressMask
    leafs["group-count"] = staticGroupGroupAddressGroupAddressMask.GroupCount
    leafs["source-count"] = staticGroupGroupAddressGroupAddressMask.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressGroupAddressMask.SuppressReport
    return leafs
}

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetYangName() string { return "static-group-group-address-group-address-mask" }

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) SetParent(parent types.Entity) { staticGroupGroupAddressGroupAddressMask.parent = parent }

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetParent() types.Entity { return staticGroupGroupAddressGroupAddressMask.parent }

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetParentYangName() string { return "static-group-group-addresses" }

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Mask for Group Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddressMask interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetFilter() yfilter.YFilter { return staticGroupGroupAddressGroupAddressMaskSourceAddress.YFilter }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressGroupAddressMaskSourceAddress.YFilter = yf }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-address-mask" { return "GroupAddressMask" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetSegmentPath() string {
    return "static-group-group-address-group-address-mask-source-address" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress) + "']" + "[group-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress) + "']"
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress
    leafs["group-address-mask"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask
    leafs["source-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress
    leafs["group-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupCount
    leafs["source-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressGroupAddressMaskSourceAddress.SuppressReport
    return leafs
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetYangName() string { return "static-group-group-address-group-address-mask-source-address" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) SetParent(parent types.Entity) { staticGroupGroupAddressGroupAddressMaskSourceAddress.parent = parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetParent() types.Entity { return staticGroupGroupAddressGroupAddressMaskSourceAddress.parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetParentYangName() string { return "static-group-group-addresses" }

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP group address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddress interface{}

    // This attribute is a key. Mask for Group Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddressMask interface{}

    // This attribute is a key. IP source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Mask for Source Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddressMask interface{}

    // Number of groups to join (do not set without GroupAddressMask). The type is
    // interface{} with range: 1..512. The default value is 1.
    GroupCount interface{}

    // Number of sources to join (do not set without SourceAddressMask). The type
    // is interface{} with range: 1..512. The default value is 1.
    SourceCount interface{}

    // Suppress reports. The type is bool. The default value is false.
    SuppressReport interface{}
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetFilter() yfilter.YFilter { return staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.YFilter }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) SetFilter(yf yfilter.YFilter) { staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.YFilter = yf }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetGoName(yname string) string {
    if yname == "group-address" { return "GroupAddress" }
    if yname == "group-address-mask" { return "GroupAddressMask" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "source-address-mask" { return "SourceAddressMask" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "source-count" { return "SourceCount" }
    if yname == "suppress-report" { return "SuppressReport" }
    return ""
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetSegmentPath() string {
    return "static-group-group-address-group-address-mask-source-address-source-address-mask" + "[group-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress) + "']" + "[group-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask) + "']" + "[source-address='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress) + "']" + "[source-address-mask='" + fmt.Sprintf("%v", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask) + "']"
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress
    leafs["group-address-mask"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask
    leafs["source-address"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress
    leafs["source-address-mask"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask
    leafs["group-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupCount
    leafs["source-count"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceCount
    leafs["suppress-report"] = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SuppressReport
    return leafs
}

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetBundleName() string { return "cisco_ios_xr" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetYangName() string { return "static-group-group-address-group-address-mask-source-address-source-address-mask" }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) SetParent(parent types.Entity) { staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.parent = parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetParent() types.Entity { return staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.parent }

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetParentYangName() string { return "static-group-group-addresses" }

// Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of groups accepted per interface by this router. The type is
    // interface{} with range: 1..40000. This attribute is mandatory.
    MaximumGroups interface{}

    // WarningThreshold for number of groups accepted per interface by this
    // router. The type is interface{} with range: 1..40000. The default value is
    // 25000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetFilter() yfilter.YFilter { return maximumGroupsPerInterfaceOor.YFilter }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) SetFilter(yf yfilter.YFilter) { maximumGroupsPerInterfaceOor.YFilter = yf }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetGoName(yname string) string {
    if yname == "maximum-groups" { return "MaximumGroups" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetSegmentPath() string {
    return "maximum-groups-per-interface-oor"
}

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-groups"] = maximumGroupsPerInterfaceOor.MaximumGroups
    leafs["warning-threshold"] = maximumGroupsPerInterfaceOor.WarningThreshold
    leafs["access-list-name"] = maximumGroupsPerInterfaceOor.AccessListName
    return leafs
}

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetBundleName() string { return "cisco_ios_xr" }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetYangName() string { return "maximum-groups-per-interface-oor" }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) SetParent(parent types.Entity) { maximumGroupsPerInterfaceOor.parent = parent }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetParent() types.Entity { return maximumGroupsPerInterfaceOor.parent }

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetParentYangName() string { return "interface" }

// Mld_DefaultContext_Interfaces_Interface_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Mld_DefaultContext_Interfaces_Interface_ExplicitTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetFilter() yfilter.YFilter { return explicitTracking.YFilter }

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) SetFilter(yf yfilter.YFilter) { explicitTracking.YFilter = yf }

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetSegmentPath() string {
    return "explicit-tracking"
}

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = explicitTracking.Enable
    leafs["access-list-name"] = explicitTracking.AccessListName
    return leafs
}

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetBundleName() string { return "cisco_ios_xr" }

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetYangName() string { return "explicit-tracking" }

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) SetParent(parent types.Entity) { explicitTracking.parent = parent }

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetParent() types.Entity { return explicitTracking.parent }

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetParentYangName() string { return "interface" }

