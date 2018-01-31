// This module contains a collection of YANG definitions
// for Cisco IOS-XR fia-hw-profile package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hw-module-profile-config: none
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package fia_hw_profile_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package fia_hw_profile_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-fia-hw-profile-cfg hw-module-profile-config}", reflect.TypeOf(HwModuleProfileConfig{}))
    ydk.RegisterEntity("Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config", reflect.TypeOf(HwModuleProfileConfig{}))
}

// HwModuleProfileConfig
// none
type HwModuleProfileConfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure profile.
    Profile HwModuleProfileConfig_Profile

    // Configure Fib for Scale for noTcam LC.
    FibScale HwModuleProfileConfig_FibScale

    // Configure Tcam.
    Tcam HwModuleProfileConfig_Tcam
}

func (hwModuleProfileConfig *HwModuleProfileConfig) GetFilter() yfilter.YFilter { return hwModuleProfileConfig.YFilter }

func (hwModuleProfileConfig *HwModuleProfileConfig) SetFilter(yf yfilter.YFilter) { hwModuleProfileConfig.YFilter = yf }

func (hwModuleProfileConfig *HwModuleProfileConfig) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    if yname == "fib-scale" { return "FibScale" }
    if yname == "tcam" { return "Tcam" }
    return ""
}

func (hwModuleProfileConfig *HwModuleProfileConfig) GetSegmentPath() string {
    return "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config"
}

func (hwModuleProfileConfig *HwModuleProfileConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        return &hwModuleProfileConfig.Profile
    }
    if childYangName == "fib-scale" {
        return &hwModuleProfileConfig.FibScale
    }
    if childYangName == "tcam" {
        return &hwModuleProfileConfig.Tcam
    }
    return nil
}

func (hwModuleProfileConfig *HwModuleProfileConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["profile"] = &hwModuleProfileConfig.Profile
    children["fib-scale"] = &hwModuleProfileConfig.FibScale
    children["tcam"] = &hwModuleProfileConfig.Tcam
    return children
}

func (hwModuleProfileConfig *HwModuleProfileConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hwModuleProfileConfig *HwModuleProfileConfig) GetBundleName() string { return "cisco_ios_xr" }

func (hwModuleProfileConfig *HwModuleProfileConfig) GetYangName() string { return "hw-module-profile-config" }

func (hwModuleProfileConfig *HwModuleProfileConfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwModuleProfileConfig *HwModuleProfileConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwModuleProfileConfig *HwModuleProfileConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwModuleProfileConfig *HwModuleProfileConfig) SetParent(parent types.Entity) { hwModuleProfileConfig.parent = parent }

func (hwModuleProfileConfig *HwModuleProfileConfig) GetParent() types.Entity { return hwModuleProfileConfig.parent }

func (hwModuleProfileConfig *HwModuleProfileConfig) GetParentYangName() string { return "Cisco-IOS-XR-fia-hw-profile-cfg" }

// HwModuleProfileConfig_Profile
// Configure profile.
type HwModuleProfileConfig_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure profile for TCAM LC cards.
    TcamTable HwModuleProfileConfig_Profile_TcamTable

    // Configure stats.
    Stats HwModuleProfileConfig_Profile_Stats

    // Configure profile.
    Qos HwModuleProfileConfig_Profile_Qos
}

func (profile *HwModuleProfileConfig_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *HwModuleProfileConfig_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *HwModuleProfileConfig_Profile) GetGoName(yname string) string {
    if yname == "tcam-table" { return "TcamTable" }
    if yname == "stats" { return "Stats" }
    if yname == "qos" { return "Qos" }
    return ""
}

func (profile *HwModuleProfileConfig_Profile) GetSegmentPath() string {
    return "profile"
}

func (profile *HwModuleProfileConfig_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcam-table" {
        return &profile.TcamTable
    }
    if childYangName == "stats" {
        return &profile.Stats
    }
    if childYangName == "qos" {
        return &profile.Qos
    }
    return nil
}

func (profile *HwModuleProfileConfig_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tcam-table"] = &profile.TcamTable
    children["stats"] = &profile.Stats
    children["qos"] = &profile.Qos
    return children
}

func (profile *HwModuleProfileConfig_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profile *HwModuleProfileConfig_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *HwModuleProfileConfig_Profile) GetYangName() string { return "profile" }

func (profile *HwModuleProfileConfig_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *HwModuleProfileConfig_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *HwModuleProfileConfig_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *HwModuleProfileConfig_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *HwModuleProfileConfig_Profile) GetParent() types.Entity { return profile.parent }

func (profile *HwModuleProfileConfig_Profile) GetParentYangName() string { return "hw-module-profile-config" }

// HwModuleProfileConfig_Profile_TcamTable
// Configure profile for TCAM LC cards
type HwModuleProfileConfig_Profile_TcamTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // FIB table for TCAM LC cards.
    FibTable HwModuleProfileConfig_Profile_TcamTable_FibTable
}

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetFilter() yfilter.YFilter { return tcamTable.YFilter }

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) SetFilter(yf yfilter.YFilter) { tcamTable.YFilter = yf }

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetGoName(yname string) string {
    if yname == "fib-table" { return "FibTable" }
    return ""
}

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetSegmentPath() string {
    return "tcam-table"
}

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fib-table" {
        return &tcamTable.FibTable
    }
    return nil
}

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fib-table"] = &tcamTable.FibTable
    return children
}

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetBundleName() string { return "cisco_ios_xr" }

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetYangName() string { return "tcam-table" }

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) SetParent(parent types.Entity) { tcamTable.parent = parent }

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetParent() types.Entity { return tcamTable.parent }

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetParentYangName() string { return "profile" }

// HwModuleProfileConfig_Profile_TcamTable_FibTable
// FIB table for TCAM LC cards
type HwModuleProfileConfig_Profile_TcamTable_FibTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 table for TCAM LC cards.
    Ipv4Address HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address

    // IPv6 table for TCAM LC cards.
    Ipv6Address HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address
}

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetFilter() yfilter.YFilter { return fibTable.YFilter }

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) SetFilter(yf yfilter.YFilter) { fibTable.YFilter = yf }

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetSegmentPath() string {
    return "fib-table"
}

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-address" {
        return &fibTable.Ipv4Address
    }
    if childYangName == "ipv6-address" {
        return &fibTable.Ipv6Address
    }
    return nil
}

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-address"] = &fibTable.Ipv4Address
    children["ipv6-address"] = &fibTable.Ipv6Address
    return children
}

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetBundleName() string { return "cisco_ios_xr" }

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetYangName() string { return "fib-table" }

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) SetParent(parent types.Entity) { fibTable.parent = parent }

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetParent() types.Entity { return fibTable.parent }

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetParentYangName() string { return "tcam-table" }

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address
// IPv4 table for TCAM LC cards
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unicast table for TCAM LC cards.
    Ipv4Unicast HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast
}

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetFilter() yfilter.YFilter { return ipv4Address.YFilter }

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) SetFilter(yf yfilter.YFilter) { ipv4Address.YFilter = yf }

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetGoName(yname string) string {
    if yname == "ipv4-unicast" { return "Ipv4Unicast" }
    return ""
}

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetSegmentPath() string {
    return "ipv4-address"
}

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-unicast" {
        return &ipv4Address.Ipv4Unicast
    }
    return nil
}

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-unicast"] = &ipv4Address.Ipv4Unicast
    return children
}

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetYangName() string { return "ipv4-address" }

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) SetParent(parent types.Entity) { ipv4Address.parent = parent }

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetParent() types.Entity { return ipv4Address.parent }

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetParentYangName() string { return "fib-table" }

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast
// Unicast table for TCAM LC cards
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // curve out percentage of TCAM table entries. The type is interface{} with
    // range: 0..100. Units are percentage.
    Ipv4UnicastPercent interface{}

    // IPv4 Unicast prefix .
    Ipv4UnicastPrefixLengths HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths
}

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetFilter() yfilter.YFilter { return ipv4Unicast.YFilter }

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) SetFilter(yf yfilter.YFilter) { ipv4Unicast.YFilter = yf }

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetGoName(yname string) string {
    if yname == "ipv4-unicast-percent" { return "Ipv4UnicastPercent" }
    if yname == "ipv4-unicast-prefix-lengths" { return "Ipv4UnicastPrefixLengths" }
    return ""
}

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetSegmentPath() string {
    return "ipv4-unicast"
}

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-unicast-prefix-lengths" {
        return &ipv4Unicast.Ipv4UnicastPrefixLengths
    }
    return nil
}

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-unicast-prefix-lengths"] = &ipv4Unicast.Ipv4UnicastPrefixLengths
    return children
}

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-unicast-percent"] = ipv4Unicast.Ipv4UnicastPercent
    return leafs
}

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetYangName() string { return "ipv4-unicast" }

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) SetParent(parent types.Entity) { ipv4Unicast.parent = parent }

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetParent() types.Entity { return ipv4Unicast.parent }

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetParentYangName() string { return "ipv4-address" }

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths
// IPv4 Unicast prefix 
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Unicast prefix length. The type is slice of
    // HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength.
    Ipv4UnicastPrefixLength []HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength
}

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetFilter() yfilter.YFilter { return ipv4UnicastPrefixLengths.YFilter }

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) SetFilter(yf yfilter.YFilter) { ipv4UnicastPrefixLengths.YFilter = yf }

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetGoName(yname string) string {
    if yname == "ipv4-unicast-prefix-length" { return "Ipv4UnicastPrefixLength" }
    return ""
}

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetSegmentPath() string {
    return "ipv4-unicast-prefix-lengths"
}

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-unicast-prefix-length" {
        for _, c := range ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength {
            if ipv4UnicastPrefixLengths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength{}
        ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength = append(ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength, child)
        return &ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength[len(ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength)-1]
    }
    return nil
}

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength {
        children[ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength[i].GetSegmentPath()] = &ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength[i]
    }
    return children
}

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetYangName() string { return "ipv4-unicast-prefix-lengths" }

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) SetParent(parent types.Entity) { ipv4UnicastPrefixLengths.parent = parent }

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetParent() types.Entity { return ipv4UnicastPrefixLengths.parent }

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetParentYangName() string { return "ipv4-unicast" }

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength
// IPv4 Unicast prefix length
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. prefix length. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}

    // curve out percentage of TCAM table entries. The type is string. Units are
    // percentage.
    Ipv4UnicastPrefixPercent interface{}
}

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetFilter() yfilter.YFilter { return ipv4UnicastPrefixLength.YFilter }

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) SetFilter(yf yfilter.YFilter) { ipv4UnicastPrefixLength.YFilter = yf }

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "ipv4-unicast-prefix-percent" { return "Ipv4UnicastPrefixPercent" }
    return ""
}

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetSegmentPath() string {
    return "ipv4-unicast-prefix-length" + "[prefix-length='" + fmt.Sprintf("%v", ipv4UnicastPrefixLength.PrefixLength) + "']"
}

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = ipv4UnicastPrefixLength.PrefixLength
    leafs["ipv4-unicast-prefix-percent"] = ipv4UnicastPrefixLength.Ipv4UnicastPrefixPercent
    return leafs
}

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetYangName() string { return "ipv4-unicast-prefix-length" }

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) SetParent(parent types.Entity) { ipv4UnicastPrefixLength.parent = parent }

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetParent() types.Entity { return ipv4UnicastPrefixLength.parent }

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetParentYangName() string { return "ipv4-unicast-prefix-lengths" }

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address
// IPv6 table for TCAM LC cards
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unicast table for TCAM LC cards.
    Ipv6Unicast HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast
}

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetFilter() yfilter.YFilter { return ipv6Address.YFilter }

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) SetFilter(yf yfilter.YFilter) { ipv6Address.YFilter = yf }

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetGoName(yname string) string {
    if yname == "ipv6-unicast" { return "Ipv6Unicast" }
    return ""
}

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetSegmentPath() string {
    return "ipv6-address"
}

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-unicast" {
        return &ipv6Address.Ipv6Unicast
    }
    return nil
}

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6-unicast"] = &ipv6Address.Ipv6Unicast
    return children
}

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetYangName() string { return "ipv6-address" }

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) SetParent(parent types.Entity) { ipv6Address.parent = parent }

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetParent() types.Entity { return ipv6Address.parent }

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetParentYangName() string { return "fib-table" }

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast
// Unicast table for TCAM LC cards
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // curve out percentage of TCAM table entries. The type is interface{} with
    // range: 0..100. Units are percentage.
    Ipv6UnicastPercent interface{}

    // IPv6 Unicast prefix .
    Ipv6UnicastPrefixLengths HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths
}

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetFilter() yfilter.YFilter { return ipv6Unicast.YFilter }

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) SetFilter(yf yfilter.YFilter) { ipv6Unicast.YFilter = yf }

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetGoName(yname string) string {
    if yname == "ipv6-unicast-percent" { return "Ipv6UnicastPercent" }
    if yname == "ipv6-unicast-prefix-lengths" { return "Ipv6UnicastPrefixLengths" }
    return ""
}

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetSegmentPath() string {
    return "ipv6-unicast"
}

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-unicast-prefix-lengths" {
        return &ipv6Unicast.Ipv6UnicastPrefixLengths
    }
    return nil
}

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6-unicast-prefix-lengths"] = &ipv6Unicast.Ipv6UnicastPrefixLengths
    return children
}

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-unicast-percent"] = ipv6Unicast.Ipv6UnicastPercent
    return leafs
}

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetYangName() string { return "ipv6-unicast" }

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) SetParent(parent types.Entity) { ipv6Unicast.parent = parent }

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetParent() types.Entity { return ipv6Unicast.parent }

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetParentYangName() string { return "ipv6-address" }

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths
// IPv6 Unicast prefix 
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Unicast prefix length. The type is slice of
    // HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength.
    Ipv6UnicastPrefixLength []HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength
}

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetFilter() yfilter.YFilter { return ipv6UnicastPrefixLengths.YFilter }

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) SetFilter(yf yfilter.YFilter) { ipv6UnicastPrefixLengths.YFilter = yf }

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetGoName(yname string) string {
    if yname == "ipv6-unicast-prefix-length" { return "Ipv6UnicastPrefixLength" }
    return ""
}

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetSegmentPath() string {
    return "ipv6-unicast-prefix-lengths"
}

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-unicast-prefix-length" {
        for _, c := range ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength {
            if ipv6UnicastPrefixLengths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength{}
        ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength = append(ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength, child)
        return &ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength[len(ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength)-1]
    }
    return nil
}

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength {
        children[ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength[i].GetSegmentPath()] = &ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength[i]
    }
    return children
}

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetYangName() string { return "ipv6-unicast-prefix-lengths" }

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) SetParent(parent types.Entity) { ipv6UnicastPrefixLengths.parent = parent }

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetParent() types.Entity { return ipv6UnicastPrefixLengths.parent }

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetParentYangName() string { return "ipv6-unicast" }

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength
// IPv6 Unicast prefix length
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. prefix length. The type is interface{} with range:
    // 0..128.
    PrefixLength interface{}

    // curve out percentage of TCAM table entries. The type is string. Units are
    // percentage.
    Ipv6UnicastPrefixPercent interface{}
}

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetFilter() yfilter.YFilter { return ipv6UnicastPrefixLength.YFilter }

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) SetFilter(yf yfilter.YFilter) { ipv6UnicastPrefixLength.YFilter = yf }

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "ipv6-unicast-prefix-percent" { return "Ipv6UnicastPrefixPercent" }
    return ""
}

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetSegmentPath() string {
    return "ipv6-unicast-prefix-length" + "[prefix-length='" + fmt.Sprintf("%v", ipv6UnicastPrefixLength.PrefixLength) + "']"
}

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = ipv6UnicastPrefixLength.PrefixLength
    leafs["ipv6-unicast-prefix-percent"] = ipv6UnicastPrefixLength.Ipv6UnicastPrefixPercent
    return leafs
}

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetYangName() string { return "ipv6-unicast-prefix-length" }

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) SetParent(parent types.Entity) { ipv6UnicastPrefixLength.parent = parent }

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetParent() types.Entity { return ipv6UnicastPrefixLength.parent }

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetParentYangName() string { return "ipv6-unicast-prefix-lengths" }

// HwModuleProfileConfig_Profile_Stats
// Configure stats
type HwModuleProfileConfig_Profile_Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure stats for qos-enhanced and acl-permit. The type is interface{}
    // with range: -2147483648..2147483647.
    CounterProfile interface{}
}

func (stats *HwModuleProfileConfig_Profile_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *HwModuleProfileConfig_Profile_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *HwModuleProfileConfig_Profile_Stats) GetGoName(yname string) string {
    if yname == "counter-profile" { return "CounterProfile" }
    return ""
}

func (stats *HwModuleProfileConfig_Profile_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *HwModuleProfileConfig_Profile_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stats *HwModuleProfileConfig_Profile_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stats *HwModuleProfileConfig_Profile_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter-profile"] = stats.CounterProfile
    return leafs
}

func (stats *HwModuleProfileConfig_Profile_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *HwModuleProfileConfig_Profile_Stats) GetYangName() string { return "stats" }

func (stats *HwModuleProfileConfig_Profile_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *HwModuleProfileConfig_Profile_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *HwModuleProfileConfig_Profile_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *HwModuleProfileConfig_Profile_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *HwModuleProfileConfig_Profile_Stats) GetParent() types.Entity { return stats.parent }

func (stats *HwModuleProfileConfig_Profile_Stats) GetParentYangName() string { return "profile" }

// HwModuleProfileConfig_Profile_Qos
// Configure profile.
type HwModuleProfileConfig_Profile_Qos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Hqos profile.
    HqosEnableAll HwModuleProfileConfig_Profile_Qos_HqosEnableAll

    // Configure Ingress Model Default.
    IngressModelRootDef HwModuleProfileConfig_Profile_Qos_IngressModelRootDef

    // Configure Ingress Model Root.
    IngressModels HwModuleProfileConfig_Profile_Qos_IngressModels

    // Configure Max Trunk Size.
    Trunks HwModuleProfileConfig_Profile_Qos_Trunks

    // Configure Class Maps Default.
    ClassMapsRootDef HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef

    // Configure Class Map Root.
    ClassMaps HwModuleProfileConfig_Profile_Qos_ClassMaps
}

func (qos *HwModuleProfileConfig_Profile_Qos) GetFilter() yfilter.YFilter { return qos.YFilter }

func (qos *HwModuleProfileConfig_Profile_Qos) SetFilter(yf yfilter.YFilter) { qos.YFilter = yf }

func (qos *HwModuleProfileConfig_Profile_Qos) GetGoName(yname string) string {
    if yname == "hqos-enable-all" { return "HqosEnableAll" }
    if yname == "ingress-model-root-def" { return "IngressModelRootDef" }
    if yname == "ingress-models" { return "IngressModels" }
    if yname == "trunks" { return "Trunks" }
    if yname == "class-maps-root-def" { return "ClassMapsRootDef" }
    if yname == "class-maps" { return "ClassMaps" }
    return ""
}

func (qos *HwModuleProfileConfig_Profile_Qos) GetSegmentPath() string {
    return "qos"
}

func (qos *HwModuleProfileConfig_Profile_Qos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hqos-enable-all" {
        return &qos.HqosEnableAll
    }
    if childYangName == "ingress-model-root-def" {
        return &qos.IngressModelRootDef
    }
    if childYangName == "ingress-models" {
        return &qos.IngressModels
    }
    if childYangName == "trunks" {
        return &qos.Trunks
    }
    if childYangName == "class-maps-root-def" {
        return &qos.ClassMapsRootDef
    }
    if childYangName == "class-maps" {
        return &qos.ClassMaps
    }
    return nil
}

func (qos *HwModuleProfileConfig_Profile_Qos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hqos-enable-all"] = &qos.HqosEnableAll
    children["ingress-model-root-def"] = &qos.IngressModelRootDef
    children["ingress-models"] = &qos.IngressModels
    children["trunks"] = &qos.Trunks
    children["class-maps-root-def"] = &qos.ClassMapsRootDef
    children["class-maps"] = &qos.ClassMaps
    return children
}

func (qos *HwModuleProfileConfig_Profile_Qos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qos *HwModuleProfileConfig_Profile_Qos) GetBundleName() string { return "cisco_ios_xr" }

func (qos *HwModuleProfileConfig_Profile_Qos) GetYangName() string { return "qos" }

func (qos *HwModuleProfileConfig_Profile_Qos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qos *HwModuleProfileConfig_Profile_Qos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qos *HwModuleProfileConfig_Profile_Qos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qos *HwModuleProfileConfig_Profile_Qos) SetParent(parent types.Entity) { qos.parent = parent }

func (qos *HwModuleProfileConfig_Profile_Qos) GetParent() types.Entity { return qos.parent }

func (qos *HwModuleProfileConfig_Profile_Qos) GetParentYangName() string { return "profile" }

// HwModuleProfileConfig_Profile_Qos_HqosEnableAll
// Configure Hqos profile
type HwModuleProfileConfig_Profile_Qos_HqosEnableAll struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Hqos profile value. The type is interface{} with range:
    // -2147483648..2147483647.
    HqosEnable interface{}
}

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetFilter() yfilter.YFilter { return hqosEnableAll.YFilter }

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) SetFilter(yf yfilter.YFilter) { hqosEnableAll.YFilter = yf }

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetGoName(yname string) string {
    if yname == "hqos-enable" { return "HqosEnable" }
    return ""
}

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetSegmentPath() string {
    return "hqos-enable-all"
}

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hqos-enable"] = hqosEnableAll.HqosEnable
    return leafs
}

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetBundleName() string { return "cisco_ios_xr" }

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetYangName() string { return "hqos-enable-all" }

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) SetParent(parent types.Entity) { hqosEnableAll.parent = parent }

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetParent() types.Entity { return hqosEnableAll.parent }

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetParentYangName() string { return "qos" }

// HwModuleProfileConfig_Profile_Qos_IngressModelRootDef
// Configure Ingress Model Default
type HwModuleProfileConfig_Profile_Qos_IngressModelRootDef struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ingress Model Default. The type is interface{} with range:
    // -2147483648..2147483647.
    IngressModelLeafDef interface{}
}

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetFilter() yfilter.YFilter { return ingressModelRootDef.YFilter }

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) SetFilter(yf yfilter.YFilter) { ingressModelRootDef.YFilter = yf }

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetGoName(yname string) string {
    if yname == "ingress-model-leaf-def" { return "IngressModelLeafDef" }
    return ""
}

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetSegmentPath() string {
    return "ingress-model-root-def"
}

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ingress-model-leaf-def"] = ingressModelRootDef.IngressModelLeafDef
    return leafs
}

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetBundleName() string { return "cisco_ios_xr" }

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetYangName() string { return "ingress-model-root-def" }

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) SetParent(parent types.Entity) { ingressModelRootDef.parent = parent }

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetParent() types.Entity { return ingressModelRootDef.parent }

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetParentYangName() string { return "qos" }

// HwModuleProfileConfig_Profile_Qos_IngressModels
// Configure Ingress Model Root
type HwModuleProfileConfig_Profile_Qos_IngressModels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Ingress Model. The type is slice of
    // HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel.
    IngressModel []HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel
}

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetFilter() yfilter.YFilter { return ingressModels.YFilter }

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) SetFilter(yf yfilter.YFilter) { ingressModels.YFilter = yf }

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetGoName(yname string) string {
    if yname == "ingress-model" { return "IngressModel" }
    return ""
}

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetSegmentPath() string {
    return "ingress-models"
}

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ingress-model" {
        for _, c := range ingressModels.IngressModel {
            if ingressModels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel{}
        ingressModels.IngressModel = append(ingressModels.IngressModel, child)
        return &ingressModels.IngressModel[len(ingressModels.IngressModel)-1]
    }
    return nil
}

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ingressModels.IngressModel {
        children[ingressModels.IngressModel[i].GetSegmentPath()] = &ingressModels.IngressModel[i]
    }
    return children
}

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetBundleName() string { return "cisco_ios_xr" }

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetYangName() string { return "ingress-models" }

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) SetParent(parent types.Entity) { ingressModels.parent = parent }

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetParent() types.Entity { return ingressModels.parent }

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetParentYangName() string { return "qos" }

// HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel
// Configure Ingress Model
type HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. NodeName. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    NodeName interface{}

    // Configure Ingress Model Leaf. The type is slice of
    // HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf.
    IngressModelLeaf []HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf
}

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetFilter() yfilter.YFilter { return ingressModel.YFilter }

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) SetFilter(yf yfilter.YFilter) { ingressModel.YFilter = yf }

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "ingress-model-leaf" { return "IngressModelLeaf" }
    return ""
}

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetSegmentPath() string {
    return "ingress-model" + "[node-name='" + fmt.Sprintf("%v", ingressModel.NodeName) + "']"
}

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ingress-model-leaf" {
        for _, c := range ingressModel.IngressModelLeaf {
            if ingressModel.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf{}
        ingressModel.IngressModelLeaf = append(ingressModel.IngressModelLeaf, child)
        return &ingressModel.IngressModelLeaf[len(ingressModel.IngressModelLeaf)-1]
    }
    return nil
}

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ingressModel.IngressModelLeaf {
        children[ingressModel.IngressModelLeaf[i].GetSegmentPath()] = &ingressModel.IngressModelLeaf[i]
    }
    return children
}

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = ingressModel.NodeName
    return leafs
}

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetBundleName() string { return "cisco_ios_xr" }

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetYangName() string { return "ingress-model" }

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) SetParent(parent types.Entity) { ingressModel.parent = parent }

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetParent() types.Entity { return ingressModel.parent }

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetParentYangName() string { return "ingress-models" }

// HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf
// Configure Ingress Model Leaf
type HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is interface{} with range:
    // -2147483648..2147483647.
    Location interface{}

    // IngressModelLeaf. The type is interface{} with range:
    // -2147483648..2147483647. This attribute is mandatory.
    IngressModelLeaf interface{}
}

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetFilter() yfilter.YFilter { return ingressModelLeaf.YFilter }

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) SetFilter(yf yfilter.YFilter) { ingressModelLeaf.YFilter = yf }

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "ingress-model-leaf" { return "IngressModelLeaf" }
    return ""
}

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetSegmentPath() string {
    return "ingress-model-leaf" + "[location='" + fmt.Sprintf("%v", ingressModelLeaf.Location) + "']"
}

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = ingressModelLeaf.Location
    leafs["ingress-model-leaf"] = ingressModelLeaf.IngressModelLeaf
    return leafs
}

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetBundleName() string { return "cisco_ios_xr" }

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetYangName() string { return "ingress-model-leaf" }

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) SetParent(parent types.Entity) { ingressModelLeaf.parent = parent }

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetParent() types.Entity { return ingressModelLeaf.parent }

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetParentYangName() string { return "ingress-model" }

// HwModuleProfileConfig_Profile_Qos_Trunks
// Configure Max Trunk Size
type HwModuleProfileConfig_Profile_Qos_Trunks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Trunk Size. The type is interface{} with range:
    // -2147483648..2147483647.
    TrunkSize interface{}
}

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetFilter() yfilter.YFilter { return trunks.YFilter }

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) SetFilter(yf yfilter.YFilter) { trunks.YFilter = yf }

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetGoName(yname string) string {
    if yname == "trunk-size" { return "TrunkSize" }
    return ""
}

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetSegmentPath() string {
    return "trunks"
}

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trunk-size"] = trunks.TrunkSize
    return leafs
}

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetBundleName() string { return "cisco_ios_xr" }

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetYangName() string { return "trunks" }

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) SetParent(parent types.Entity) { trunks.parent = parent }

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetParent() types.Entity { return trunks.parent }

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetParentYangName() string { return "qos" }

// HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef
// Configure Class Maps Default
type HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Class Map Size Default. The type is interface{} with range:
    // -2147483648..2147483647.
    ClassMapSizeDef interface{}
}

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetFilter() yfilter.YFilter { return classMapsRootDef.YFilter }

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) SetFilter(yf yfilter.YFilter) { classMapsRootDef.YFilter = yf }

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetGoName(yname string) string {
    if yname == "class-map-size-def" { return "ClassMapSizeDef" }
    return ""
}

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetSegmentPath() string {
    return "class-maps-root-def"
}

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-map-size-def"] = classMapsRootDef.ClassMapSizeDef
    return leafs
}

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetBundleName() string { return "cisco_ios_xr" }

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetYangName() string { return "class-maps-root-def" }

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) SetParent(parent types.Entity) { classMapsRootDef.parent = parent }

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetParent() types.Entity { return classMapsRootDef.parent }

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetParentYangName() string { return "qos" }

// HwModuleProfileConfig_Profile_Qos_ClassMaps
// Configure Class Map Root
type HwModuleProfileConfig_Profile_Qos_ClassMaps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Class Maps. The type is slice of
    // HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap.
    ClassMap []HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap
}

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetFilter() yfilter.YFilter { return classMaps.YFilter }

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) SetFilter(yf yfilter.YFilter) { classMaps.YFilter = yf }

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetGoName(yname string) string {
    if yname == "class-map" { return "ClassMap" }
    return ""
}

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetSegmentPath() string {
    return "class-maps"
}

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class-map" {
        for _, c := range classMaps.ClassMap {
            if classMaps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap{}
        classMaps.ClassMap = append(classMaps.ClassMap, child)
        return &classMaps.ClassMap[len(classMaps.ClassMap)-1]
    }
    return nil
}

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classMaps.ClassMap {
        children[classMaps.ClassMap[i].GetSegmentPath()] = &classMaps.ClassMap[i]
    }
    return children
}

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetBundleName() string { return "cisco_ios_xr" }

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetYangName() string { return "class-maps" }

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) SetParent(parent types.Entity) { classMaps.parent = parent }

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetParent() types.Entity { return classMaps.parent }

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetParentYangName() string { return "qos" }

// HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap
// Configure Class Maps
type HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. NodeName. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    NodeName interface{}

    // Class Map Size. The type is slice of
    // HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize.
    ClassMapSize []HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize
}

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetFilter() yfilter.YFilter { return classMap.YFilter }

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) SetFilter(yf yfilter.YFilter) { classMap.YFilter = yf }

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "class-map-size" { return "ClassMapSize" }
    return ""
}

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetSegmentPath() string {
    return "class-map" + "[node-name='" + fmt.Sprintf("%v", classMap.NodeName) + "']"
}

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class-map-size" {
        for _, c := range classMap.ClassMapSize {
            if classMap.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize{}
        classMap.ClassMapSize = append(classMap.ClassMapSize, child)
        return &classMap.ClassMapSize[len(classMap.ClassMapSize)-1]
    }
    return nil
}

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classMap.ClassMapSize {
        children[classMap.ClassMapSize[i].GetSegmentPath()] = &classMap.ClassMapSize[i]
    }
    return children
}

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = classMap.NodeName
    return leafs
}

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetBundleName() string { return "cisco_ios_xr" }

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetYangName() string { return "class-map" }

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) SetParent(parent types.Entity) { classMap.parent = parent }

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetParent() types.Entity { return classMap.parent }

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetParentYangName() string { return "class-maps" }

// HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize
// Class Map Size
type HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is interface{} with range:
    // -2147483648..2147483647.
    Location interface{}

    // ClassMapSize. The type is interface{} with range: -2147483648..2147483647.
    // This attribute is mandatory.
    ClassMapSize interface{}
}

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetFilter() yfilter.YFilter { return classMapSize.YFilter }

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) SetFilter(yf yfilter.YFilter) { classMapSize.YFilter = yf }

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "class-map-size" { return "ClassMapSize" }
    return ""
}

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetSegmentPath() string {
    return "class-map-size" + "[location='" + fmt.Sprintf("%v", classMapSize.Location) + "']"
}

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = classMapSize.Location
    leafs["class-map-size"] = classMapSize.ClassMapSize
    return leafs
}

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetBundleName() string { return "cisco_ios_xr" }

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetYangName() string { return "class-map-size" }

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) SetParent(parent types.Entity) { classMapSize.parent = parent }

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetParent() types.Entity { return classMapSize.parent }

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetParentYangName() string { return "class-map" }

// HwModuleProfileConfig_FibScale
// Configure Fib for Scale for noTcam LC.
type HwModuleProfileConfig_FibScale struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 table for NOTCAM LC Scale.
    Ipv6UnicastScaleNoTcam HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam

    // IPv4 table for NOTCAM LC Scale.
    Ipv4UnicastScaleNoTcam HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam
}

func (fibScale *HwModuleProfileConfig_FibScale) GetFilter() yfilter.YFilter { return fibScale.YFilter }

func (fibScale *HwModuleProfileConfig_FibScale) SetFilter(yf yfilter.YFilter) { fibScale.YFilter = yf }

func (fibScale *HwModuleProfileConfig_FibScale) GetGoName(yname string) string {
    if yname == "ipv6-unicast-scale-no-tcam" { return "Ipv6UnicastScaleNoTcam" }
    if yname == "ipv4-unicast-scale-no-tcam" { return "Ipv4UnicastScaleNoTcam" }
    return ""
}

func (fibScale *HwModuleProfileConfig_FibScale) GetSegmentPath() string {
    return "fib-scale"
}

func (fibScale *HwModuleProfileConfig_FibScale) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-unicast-scale-no-tcam" {
        return &fibScale.Ipv6UnicastScaleNoTcam
    }
    if childYangName == "ipv4-unicast-scale-no-tcam" {
        return &fibScale.Ipv4UnicastScaleNoTcam
    }
    return nil
}

func (fibScale *HwModuleProfileConfig_FibScale) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6-unicast-scale-no-tcam"] = &fibScale.Ipv6UnicastScaleNoTcam
    children["ipv4-unicast-scale-no-tcam"] = &fibScale.Ipv4UnicastScaleNoTcam
    return children
}

func (fibScale *HwModuleProfileConfig_FibScale) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fibScale *HwModuleProfileConfig_FibScale) GetBundleName() string { return "cisco_ios_xr" }

func (fibScale *HwModuleProfileConfig_FibScale) GetYangName() string { return "fib-scale" }

func (fibScale *HwModuleProfileConfig_FibScale) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fibScale *HwModuleProfileConfig_FibScale) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fibScale *HwModuleProfileConfig_FibScale) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fibScale *HwModuleProfileConfig_FibScale) SetParent(parent types.Entity) { fibScale.parent = parent }

func (fibScale *HwModuleProfileConfig_FibScale) GetParent() types.Entity { return fibScale.parent }

func (fibScale *HwModuleProfileConfig_FibScale) GetParentYangName() string { return "hw-module-profile-config" }

// HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam
// IPv6 table for NOTCAM LC Scale.
type HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Scale for IPv6 table for NoTCAM LC.
    ScaleIpv6NoTcam HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam
}

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetFilter() yfilter.YFilter { return ipv6UnicastScaleNoTcam.YFilter }

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) SetFilter(yf yfilter.YFilter) { ipv6UnicastScaleNoTcam.YFilter = yf }

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetGoName(yname string) string {
    if yname == "scale-ipv6-no-tcam" { return "ScaleIpv6NoTcam" }
    return ""
}

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetSegmentPath() string {
    return "ipv6-unicast-scale-no-tcam"
}

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "scale-ipv6-no-tcam" {
        return &ipv6UnicastScaleNoTcam.ScaleIpv6NoTcam
    }
    return nil
}

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["scale-ipv6-no-tcam"] = &ipv6UnicastScaleNoTcam.ScaleIpv6NoTcam
    return children
}

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetYangName() string { return "ipv6-unicast-scale-no-tcam" }

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) SetParent(parent types.Entity) { ipv6UnicastScaleNoTcam.parent = parent }

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetParent() types.Entity { return ipv6UnicastScaleNoTcam.parent }

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetParentYangName() string { return "fib-scale" }

// HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam
// Scale for IPv6 table for NoTCAM LC.
type HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Internet-optimized Scale for IPv6 table for NoTCAM LC. The type is string.
    InternetOptimizedIpv6NoTcam interface{}
}

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetFilter() yfilter.YFilter { return scaleIpv6NoTcam.YFilter }

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) SetFilter(yf yfilter.YFilter) { scaleIpv6NoTcam.YFilter = yf }

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetGoName(yname string) string {
    if yname == "internet-optimized-ipv6-no-tcam" { return "InternetOptimizedIpv6NoTcam" }
    return ""
}

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetSegmentPath() string {
    return "scale-ipv6-no-tcam"
}

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["internet-optimized-ipv6-no-tcam"] = scaleIpv6NoTcam.InternetOptimizedIpv6NoTcam
    return leafs
}

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetBundleName() string { return "cisco_ios_xr" }

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetYangName() string { return "scale-ipv6-no-tcam" }

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) SetParent(parent types.Entity) { scaleIpv6NoTcam.parent = parent }

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetParent() types.Entity { return scaleIpv6NoTcam.parent }

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetParentYangName() string { return "ipv6-unicast-scale-no-tcam" }

// HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam
// IPv4 table for NOTCAM LC Scale.
type HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Scale for IPv4 table for NoTCAM LC.
    ScaleIpv4NoTcam HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam
}

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetFilter() yfilter.YFilter { return ipv4UnicastScaleNoTcam.YFilter }

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) SetFilter(yf yfilter.YFilter) { ipv4UnicastScaleNoTcam.YFilter = yf }

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetGoName(yname string) string {
    if yname == "scale-ipv4-no-tcam" { return "ScaleIpv4NoTcam" }
    return ""
}

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetSegmentPath() string {
    return "ipv4-unicast-scale-no-tcam"
}

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "scale-ipv4-no-tcam" {
        return &ipv4UnicastScaleNoTcam.ScaleIpv4NoTcam
    }
    return nil
}

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["scale-ipv4-no-tcam"] = &ipv4UnicastScaleNoTcam.ScaleIpv4NoTcam
    return children
}

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetYangName() string { return "ipv4-unicast-scale-no-tcam" }

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) SetParent(parent types.Entity) { ipv4UnicastScaleNoTcam.parent = parent }

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetParent() types.Entity { return ipv4UnicastScaleNoTcam.parent }

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetParentYangName() string { return "fib-scale" }

// HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam
// Scale for IPv4 table for NoTCAM LC.
type HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Host-optimized Scale for IPv4 table for NoTCAM LC. The type is string.
    HostOptimizedIpv4NoTcam interface{}

    // Internet-optimized Scale for IPv4 table for NoTCAM LC. The type is string.
    InternetOptimizedIpv4NoTcam interface{}
}

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetFilter() yfilter.YFilter { return scaleIpv4NoTcam.YFilter }

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) SetFilter(yf yfilter.YFilter) { scaleIpv4NoTcam.YFilter = yf }

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetGoName(yname string) string {
    if yname == "host-optimized-ipv4-no-tcam" { return "HostOptimizedIpv4NoTcam" }
    if yname == "internet-optimized-ipv4-no-tcam" { return "InternetOptimizedIpv4NoTcam" }
    return ""
}

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetSegmentPath() string {
    return "scale-ipv4-no-tcam"
}

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-optimized-ipv4-no-tcam"] = scaleIpv4NoTcam.HostOptimizedIpv4NoTcam
    leafs["internet-optimized-ipv4-no-tcam"] = scaleIpv4NoTcam.InternetOptimizedIpv4NoTcam
    return leafs
}

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetBundleName() string { return "cisco_ios_xr" }

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetYangName() string { return "scale-ipv4-no-tcam" }

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) SetParent(parent types.Entity) { scaleIpv4NoTcam.parent = parent }

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetParent() types.Entity { return scaleIpv4NoTcam.parent }

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetParentYangName() string { return "ipv4-unicast-scale-no-tcam" }

// HwModuleProfileConfig_Tcam
// Configure Tcam.
type HwModuleProfileConfig_Tcam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Fib iscale for Tcam.
    FibTcamScale HwModuleProfileConfig_Tcam_FibTcamScale
}

func (tcam *HwModuleProfileConfig_Tcam) GetFilter() yfilter.YFilter { return tcam.YFilter }

func (tcam *HwModuleProfileConfig_Tcam) SetFilter(yf yfilter.YFilter) { tcam.YFilter = yf }

func (tcam *HwModuleProfileConfig_Tcam) GetGoName(yname string) string {
    if yname == "fib-tcam-scale" { return "FibTcamScale" }
    return ""
}

func (tcam *HwModuleProfileConfig_Tcam) GetSegmentPath() string {
    return "tcam"
}

func (tcam *HwModuleProfileConfig_Tcam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fib-tcam-scale" {
        return &tcam.FibTcamScale
    }
    return nil
}

func (tcam *HwModuleProfileConfig_Tcam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fib-tcam-scale"] = &tcam.FibTcamScale
    return children
}

func (tcam *HwModuleProfileConfig_Tcam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tcam *HwModuleProfileConfig_Tcam) GetBundleName() string { return "cisco_ios_xr" }

func (tcam *HwModuleProfileConfig_Tcam) GetYangName() string { return "tcam" }

func (tcam *HwModuleProfileConfig_Tcam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcam *HwModuleProfileConfig_Tcam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcam *HwModuleProfileConfig_Tcam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcam *HwModuleProfileConfig_Tcam) SetParent(parent types.Entity) { tcam.parent = parent }

func (tcam *HwModuleProfileConfig_Tcam) GetParent() types.Entity { return tcam.parent }

func (tcam *HwModuleProfileConfig_Tcam) GetParentYangName() string { return "hw-module-profile-config" }

// HwModuleProfileConfig_Tcam_FibTcamScale
// Configure Fib iscale for Tcam.
type HwModuleProfileConfig_Tcam_FibTcamScale struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 table for TCAM LC Scale.
    Ipv4UnicastScale HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale
}

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetFilter() yfilter.YFilter { return fibTcamScale.YFilter }

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) SetFilter(yf yfilter.YFilter) { fibTcamScale.YFilter = yf }

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetGoName(yname string) string {
    if yname == "ipv4-unicast-scale" { return "Ipv4UnicastScale" }
    return ""
}

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetSegmentPath() string {
    return "fib-tcam-scale"
}

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-unicast-scale" {
        return &fibTcamScale.Ipv4UnicastScale
    }
    return nil
}

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-unicast-scale"] = &fibTcamScale.Ipv4UnicastScale
    return children
}

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetBundleName() string { return "cisco_ios_xr" }

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetYangName() string { return "fib-tcam-scale" }

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) SetParent(parent types.Entity) { fibTcamScale.parent = parent }

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetParent() types.Entity { return fibTcamScale.parent }

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetParentYangName() string { return "tcam" }

// HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale
// IPv4 table for TCAM LC Scale.
type HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Scale for IPv4 table for TCAM LC. The type is interface{}.
    Ipv4Scale interface{}
}

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetFilter() yfilter.YFilter { return ipv4UnicastScale.YFilter }

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) SetFilter(yf yfilter.YFilter) { ipv4UnicastScale.YFilter = yf }

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetGoName(yname string) string {
    if yname == "ipv4-scale" { return "Ipv4Scale" }
    return ""
}

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetSegmentPath() string {
    return "ipv4-unicast-scale"
}

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-scale"] = ipv4UnicastScale.Ipv4Scale
    return leafs
}

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetYangName() string { return "ipv4-unicast-scale" }

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) SetParent(parent types.Entity) { ipv4UnicastScale.parent = parent }

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetParent() types.Entity { return ipv4UnicastScale.parent }

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetParentYangName() string { return "fib-tcam-scale" }

