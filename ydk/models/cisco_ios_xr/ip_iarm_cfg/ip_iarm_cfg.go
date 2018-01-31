// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-iarm package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ip-arm: IP Address Repository Manager (IPv4/IPv6 ARM)
//     configuration data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_iarm_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_iarm_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-iarm-cfg ip-arm}", reflect.TypeOf(IpArm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-iarm-cfg:ip-arm", reflect.TypeOf(IpArm{}))
}

// IpArmConflictPolicy represents Ip arm conflict policy
type IpArmConflictPolicy string

const (
    // Lowest Rack/Slot
    IpArmConflictPolicy_lowest_rack_slot IpArmConflictPolicy = "lowest-rack-slot"

    // UP interfaces unaffected
    IpArmConflictPolicy_static IpArmConflictPolicy = "static"

    // Longest Prefix
    IpArmConflictPolicy_longest_prefix IpArmConflictPolicy = "longest-prefix"

    // Highest IP
    IpArmConflictPolicy_highest_ip IpArmConflictPolicy = "highest-ip"
)

// IpArm
// IP Address Repository Manager (IPv4/IPv6 ARM)
// configuration data
type IpArm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 ARM configuration.
    Ipv4 IpArm_Ipv4

    // IPv6 ARM configuration.
    Ipv6 IpArm_Ipv6
}

func (ipArm *IpArm) GetFilter() yfilter.YFilter { return ipArm.YFilter }

func (ipArm *IpArm) SetFilter(yf yfilter.YFilter) { ipArm.YFilter = yf }

func (ipArm *IpArm) GetGoName(yname string) string {
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (ipArm *IpArm) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-iarm-cfg:ip-arm"
}

func (ipArm *IpArm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4" {
        return &ipArm.Ipv4
    }
    if childYangName == "ipv6" {
        return &ipArm.Ipv6
    }
    return nil
}

func (ipArm *IpArm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4"] = &ipArm.Ipv4
    children["ipv6"] = &ipArm.Ipv6
    return children
}

func (ipArm *IpArm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipArm *IpArm) GetBundleName() string { return "cisco_ios_xr" }

func (ipArm *IpArm) GetYangName() string { return "ip-arm" }

func (ipArm *IpArm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipArm *IpArm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipArm *IpArm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipArm *IpArm) SetParent(parent types.Entity) { ipArm.parent = parent }

func (ipArm *IpArm) GetParent() types.Entity { return ipArm.parent }

func (ipArm *IpArm) GetParentYangName() string { return "Cisco-IOS-XR-ip-iarm-cfg" }

// IpArm_Ipv4
// IPv4 ARM configuration
type IpArm_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP ARM conflict policy configuration.
    ConflictPolicyTable IpArm_Ipv4_ConflictPolicyTable

    // IP ARM Multicast Host configuration.
    MulticastHost IpArm_Ipv4_MulticastHost
}

func (ipv4 *IpArm_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *IpArm_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *IpArm_Ipv4) GetGoName(yname string) string {
    if yname == "conflict-policy-table" { return "ConflictPolicyTable" }
    if yname == "multicast-host" { return "MulticastHost" }
    return ""
}

func (ipv4 *IpArm_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *IpArm_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "conflict-policy-table" {
        return &ipv4.ConflictPolicyTable
    }
    if childYangName == "multicast-host" {
        return &ipv4.MulticastHost
    }
    return nil
}

func (ipv4 *IpArm_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["conflict-policy-table"] = &ipv4.ConflictPolicyTable
    children["multicast-host"] = &ipv4.MulticastHost
    return children
}

func (ipv4 *IpArm_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *IpArm_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *IpArm_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *IpArm_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *IpArm_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *IpArm_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *IpArm_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *IpArm_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *IpArm_Ipv4) GetParentYangName() string { return "ip-arm" }

// IpArm_Ipv4_ConflictPolicyTable
// IP ARM conflict policy configuration
type IpArm_Ipv4_ConflictPolicyTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP ARM conflict policy value definitions. The type is IpArmConflictPolicy.
    ConflictPolicy interface{}
}

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetFilter() yfilter.YFilter { return conflictPolicyTable.YFilter }

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) SetFilter(yf yfilter.YFilter) { conflictPolicyTable.YFilter = yf }

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetGoName(yname string) string {
    if yname == "conflict-policy" { return "ConflictPolicy" }
    return ""
}

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetSegmentPath() string {
    return "conflict-policy-table"
}

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["conflict-policy"] = conflictPolicyTable.ConflictPolicy
    return leafs
}

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetBundleName() string { return "cisco_ios_xr" }

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetYangName() string { return "conflict-policy-table" }

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) SetParent(parent types.Entity) { conflictPolicyTable.parent = parent }

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetParent() types.Entity { return conflictPolicyTable.parent }

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetParentYangName() string { return "ipv4" }

// IpArm_Ipv4_MulticastHost
// IP ARM Multicast Host configuration
type IpArm_Ipv4_MulticastHost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default multicast host interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    MulticastHostInterface interface{}
}

func (multicastHost *IpArm_Ipv4_MulticastHost) GetFilter() yfilter.YFilter { return multicastHost.YFilter }

func (multicastHost *IpArm_Ipv4_MulticastHost) SetFilter(yf yfilter.YFilter) { multicastHost.YFilter = yf }

func (multicastHost *IpArm_Ipv4_MulticastHost) GetGoName(yname string) string {
    if yname == "multicast-host-interface" { return "MulticastHostInterface" }
    return ""
}

func (multicastHost *IpArm_Ipv4_MulticastHost) GetSegmentPath() string {
    return "multicast-host"
}

func (multicastHost *IpArm_Ipv4_MulticastHost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multicastHost *IpArm_Ipv4_MulticastHost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multicastHost *IpArm_Ipv4_MulticastHost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["multicast-host-interface"] = multicastHost.MulticastHostInterface
    return leafs
}

func (multicastHost *IpArm_Ipv4_MulticastHost) GetBundleName() string { return "cisco_ios_xr" }

func (multicastHost *IpArm_Ipv4_MulticastHost) GetYangName() string { return "multicast-host" }

func (multicastHost *IpArm_Ipv4_MulticastHost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicastHost *IpArm_Ipv4_MulticastHost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicastHost *IpArm_Ipv4_MulticastHost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicastHost *IpArm_Ipv4_MulticastHost) SetParent(parent types.Entity) { multicastHost.parent = parent }

func (multicastHost *IpArm_Ipv4_MulticastHost) GetParent() types.Entity { return multicastHost.parent }

func (multicastHost *IpArm_Ipv4_MulticastHost) GetParentYangName() string { return "ipv4" }

// IpArm_Ipv6
// IPv6 ARM configuration
type IpArm_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP ARM conflict policy configuration.
    ConflictPolicyTable IpArm_Ipv6_ConflictPolicyTable

    // IP ARM Multicast Host configuration.
    MulticastHost IpArm_Ipv6_MulticastHost
}

func (ipv6 *IpArm_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *IpArm_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *IpArm_Ipv6) GetGoName(yname string) string {
    if yname == "conflict-policy-table" { return "ConflictPolicyTable" }
    if yname == "multicast-host" { return "MulticastHost" }
    return ""
}

func (ipv6 *IpArm_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *IpArm_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "conflict-policy-table" {
        return &ipv6.ConflictPolicyTable
    }
    if childYangName == "multicast-host" {
        return &ipv6.MulticastHost
    }
    return nil
}

func (ipv6 *IpArm_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["conflict-policy-table"] = &ipv6.ConflictPolicyTable
    children["multicast-host"] = &ipv6.MulticastHost
    return children
}

func (ipv6 *IpArm_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *IpArm_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *IpArm_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *IpArm_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *IpArm_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *IpArm_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *IpArm_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *IpArm_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *IpArm_Ipv6) GetParentYangName() string { return "ip-arm" }

// IpArm_Ipv6_ConflictPolicyTable
// IP ARM conflict policy configuration
type IpArm_Ipv6_ConflictPolicyTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP ARM conflict policy value definitions. The type is IpArmConflictPolicy.
    ConflictPolicy interface{}
}

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetFilter() yfilter.YFilter { return conflictPolicyTable.YFilter }

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) SetFilter(yf yfilter.YFilter) { conflictPolicyTable.YFilter = yf }

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetGoName(yname string) string {
    if yname == "conflict-policy" { return "ConflictPolicy" }
    return ""
}

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetSegmentPath() string {
    return "conflict-policy-table"
}

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["conflict-policy"] = conflictPolicyTable.ConflictPolicy
    return leafs
}

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetBundleName() string { return "cisco_ios_xr" }

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetYangName() string { return "conflict-policy-table" }

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) SetParent(parent types.Entity) { conflictPolicyTable.parent = parent }

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetParent() types.Entity { return conflictPolicyTable.parent }

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetParentYangName() string { return "ipv6" }

// IpArm_Ipv6_MulticastHost
// IP ARM Multicast Host configuration
type IpArm_Ipv6_MulticastHost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default multicast host interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    MulticastHostInterface interface{}
}

func (multicastHost *IpArm_Ipv6_MulticastHost) GetFilter() yfilter.YFilter { return multicastHost.YFilter }

func (multicastHost *IpArm_Ipv6_MulticastHost) SetFilter(yf yfilter.YFilter) { multicastHost.YFilter = yf }

func (multicastHost *IpArm_Ipv6_MulticastHost) GetGoName(yname string) string {
    if yname == "multicast-host-interface" { return "MulticastHostInterface" }
    return ""
}

func (multicastHost *IpArm_Ipv6_MulticastHost) GetSegmentPath() string {
    return "multicast-host"
}

func (multicastHost *IpArm_Ipv6_MulticastHost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multicastHost *IpArm_Ipv6_MulticastHost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multicastHost *IpArm_Ipv6_MulticastHost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["multicast-host-interface"] = multicastHost.MulticastHostInterface
    return leafs
}

func (multicastHost *IpArm_Ipv6_MulticastHost) GetBundleName() string { return "cisco_ios_xr" }

func (multicastHost *IpArm_Ipv6_MulticastHost) GetYangName() string { return "multicast-host" }

func (multicastHost *IpArm_Ipv6_MulticastHost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicastHost *IpArm_Ipv6_MulticastHost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicastHost *IpArm_Ipv6_MulticastHost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicastHost *IpArm_Ipv6_MulticastHost) SetParent(parent types.Entity) { multicastHost.parent = parent }

func (multicastHost *IpArm_Ipv6_MulticastHost) GetParent() types.Entity { return multicastHost.parent }

func (multicastHost *IpArm_Ipv6_MulticastHost) GetParentYangName() string { return "ipv6" }

