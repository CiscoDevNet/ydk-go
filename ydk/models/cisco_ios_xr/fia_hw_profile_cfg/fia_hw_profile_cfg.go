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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure profile.
    Profile HwModuleProfileConfig_Profile

    // Configure Fib for Scale for noTcam LC.
    FibScale HwModuleProfileConfig_FibScale

    // Configure Tcam.
    Tcam HwModuleProfileConfig_Tcam
}

func (hwModuleProfileConfig *HwModuleProfileConfig) GetEntityData() *types.CommonEntityData {
    hwModuleProfileConfig.EntityData.YFilter = hwModuleProfileConfig.YFilter
    hwModuleProfileConfig.EntityData.YangName = "hw-module-profile-config"
    hwModuleProfileConfig.EntityData.BundleName = "cisco_ios_xr"
    hwModuleProfileConfig.EntityData.ParentYangName = "Cisco-IOS-XR-fia-hw-profile-cfg"
    hwModuleProfileConfig.EntityData.SegmentPath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config"
    hwModuleProfileConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwModuleProfileConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwModuleProfileConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwModuleProfileConfig.EntityData.Children = make(map[string]types.YChild)
    hwModuleProfileConfig.EntityData.Children["profile"] = types.YChild{"Profile", &hwModuleProfileConfig.Profile}
    hwModuleProfileConfig.EntityData.Children["fib-scale"] = types.YChild{"FibScale", &hwModuleProfileConfig.FibScale}
    hwModuleProfileConfig.EntityData.Children["tcam"] = types.YChild{"Tcam", &hwModuleProfileConfig.Tcam}
    hwModuleProfileConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hwModuleProfileConfig.EntityData)
}

// HwModuleProfileConfig_Profile
// Configure profile.
type HwModuleProfileConfig_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure profile for TCAM LC cards.
    TcamTable HwModuleProfileConfig_Profile_TcamTable

    // Configure load balance parameters.
    LoadBalance HwModuleProfileConfig_Profile_LoadBalance

    // Configure stats.
    Stats HwModuleProfileConfig_Profile_Stats

    // Configure Netflow profile.
    Netflows HwModuleProfileConfig_Profile_Netflows

    // Configure acl profile.
    ProfileAcl HwModuleProfileConfig_Profile_ProfileAcl

    // Configure Tcam Profile.
    ProfileTcam HwModuleProfileConfig_Profile_ProfileTcam

    // Configure profile.
    Qos HwModuleProfileConfig_Profile_Qos
}

func (profile *HwModuleProfileConfig_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "hw-module-profile-config"
    profile.EntityData.SegmentPath = "profile"
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Children["tcam-table"] = types.YChild{"TcamTable", &profile.TcamTable}
    profile.EntityData.Children["load-balance"] = types.YChild{"LoadBalance", &profile.LoadBalance}
    profile.EntityData.Children["stats"] = types.YChild{"Stats", &profile.Stats}
    profile.EntityData.Children["netflows"] = types.YChild{"Netflows", &profile.Netflows}
    profile.EntityData.Children["profile-acl"] = types.YChild{"ProfileAcl", &profile.ProfileAcl}
    profile.EntityData.Children["profile-tcam"] = types.YChild{"ProfileTcam", &profile.ProfileTcam}
    profile.EntityData.Children["qos"] = types.YChild{"Qos", &profile.Qos}
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(profile.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable
// Configure profile for TCAM LC cards
type HwModuleProfileConfig_Profile_TcamTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FIB table for TCAM LC cards.
    FibTable HwModuleProfileConfig_Profile_TcamTable_FibTable
}

func (tcamTable *HwModuleProfileConfig_Profile_TcamTable) GetEntityData() *types.CommonEntityData {
    tcamTable.EntityData.YFilter = tcamTable.YFilter
    tcamTable.EntityData.YangName = "tcam-table"
    tcamTable.EntityData.BundleName = "cisco_ios_xr"
    tcamTable.EntityData.ParentYangName = "profile"
    tcamTable.EntityData.SegmentPath = "tcam-table"
    tcamTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcamTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcamTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcamTable.EntityData.Children = make(map[string]types.YChild)
    tcamTable.EntityData.Children["fib-table"] = types.YChild{"FibTable", &tcamTable.FibTable}
    tcamTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tcamTable.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable_FibTable
// FIB table for TCAM LC cards
type HwModuleProfileConfig_Profile_TcamTable_FibTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 table for TCAM LC cards.
    Ipv4Address HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address

    // IPv6 table for TCAM LC cards.
    Ipv6Address HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address
}

func (fibTable *HwModuleProfileConfig_Profile_TcamTable_FibTable) GetEntityData() *types.CommonEntityData {
    fibTable.EntityData.YFilter = fibTable.YFilter
    fibTable.EntityData.YangName = "fib-table"
    fibTable.EntityData.BundleName = "cisco_ios_xr"
    fibTable.EntityData.ParentYangName = "tcam-table"
    fibTable.EntityData.SegmentPath = "fib-table"
    fibTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibTable.EntityData.Children = make(map[string]types.YChild)
    fibTable.EntityData.Children["ipv4-address"] = types.YChild{"Ipv4Address", &fibTable.Ipv4Address}
    fibTable.EntityData.Children["ipv6-address"] = types.YChild{"Ipv6Address", &fibTable.Ipv6Address}
    fibTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fibTable.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address
// IPv4 table for TCAM LC cards
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unicast table for TCAM LC cards.
    Ipv4Unicast HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast
}

func (ipv4Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address) GetEntityData() *types.CommonEntityData {
    ipv4Address.EntityData.YFilter = ipv4Address.YFilter
    ipv4Address.EntityData.YangName = "ipv4-address"
    ipv4Address.EntityData.BundleName = "cisco_ios_xr"
    ipv4Address.EntityData.ParentYangName = "fib-table"
    ipv4Address.EntityData.SegmentPath = "ipv4-address"
    ipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Address.EntityData.Children = make(map[string]types.YChild)
    ipv4Address.EntityData.Children["ipv4-unicast"] = types.YChild{"Ipv4Unicast", &ipv4Address.Ipv4Unicast}
    ipv4Address.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4Address.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast
// Unicast table for TCAM LC cards
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // curve out percentage of TCAM table entries. The type is interface{} with
    // range: 0..100. Units are percentage.
    Ipv4UnicastPercent interface{}

    // IPv4 Unicast prefix .
    Ipv4UnicastPrefixLengths HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths
}

func (ipv4Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast) GetEntityData() *types.CommonEntityData {
    ipv4Unicast.EntityData.YFilter = ipv4Unicast.YFilter
    ipv4Unicast.EntityData.YangName = "ipv4-unicast"
    ipv4Unicast.EntityData.BundleName = "cisco_ios_xr"
    ipv4Unicast.EntityData.ParentYangName = "ipv4-address"
    ipv4Unicast.EntityData.SegmentPath = "ipv4-unicast"
    ipv4Unicast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Unicast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Unicast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Unicast.EntityData.Children = make(map[string]types.YChild)
    ipv4Unicast.EntityData.Children["ipv4-unicast-prefix-lengths"] = types.YChild{"Ipv4UnicastPrefixLengths", &ipv4Unicast.Ipv4UnicastPrefixLengths}
    ipv4Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4Unicast.EntityData.Leafs["ipv4-unicast-percent"] = types.YLeaf{"Ipv4UnicastPercent", ipv4Unicast.Ipv4UnicastPercent}
    return &(ipv4Unicast.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths
// IPv4 Unicast prefix 
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Unicast prefix length. The type is slice of
    // HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength.
    Ipv4UnicastPrefixLength []HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength
}

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetEntityData() *types.CommonEntityData {
    ipv4UnicastPrefixLengths.EntityData.YFilter = ipv4UnicastPrefixLengths.YFilter
    ipv4UnicastPrefixLengths.EntityData.YangName = "ipv4-unicast-prefix-lengths"
    ipv4UnicastPrefixLengths.EntityData.BundleName = "cisco_ios_xr"
    ipv4UnicastPrefixLengths.EntityData.ParentYangName = "ipv4-unicast"
    ipv4UnicastPrefixLengths.EntityData.SegmentPath = "ipv4-unicast-prefix-lengths"
    ipv4UnicastPrefixLengths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4UnicastPrefixLengths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4UnicastPrefixLengths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4UnicastPrefixLengths.EntityData.Children = make(map[string]types.YChild)
    ipv4UnicastPrefixLengths.EntityData.Children["ipv4-unicast-prefix-length"] = types.YChild{"Ipv4UnicastPrefixLength", nil}
    for i := range ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength {
        ipv4UnicastPrefixLengths.EntityData.Children[types.GetSegmentPath(&ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength[i])] = types.YChild{"Ipv4UnicastPrefixLength", &ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength[i]}
    }
    ipv4UnicastPrefixLengths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4UnicastPrefixLengths.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength
// IPv4 Unicast prefix length
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix length. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}

    // curve out percentage of TCAM table entries. The type is string. Units are
    // percentage.
    Ipv4UnicastPrefixPercent interface{}
}

func (ipv4UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength) GetEntityData() *types.CommonEntityData {
    ipv4UnicastPrefixLength.EntityData.YFilter = ipv4UnicastPrefixLength.YFilter
    ipv4UnicastPrefixLength.EntityData.YangName = "ipv4-unicast-prefix-length"
    ipv4UnicastPrefixLength.EntityData.BundleName = "cisco_ios_xr"
    ipv4UnicastPrefixLength.EntityData.ParentYangName = "ipv4-unicast-prefix-lengths"
    ipv4UnicastPrefixLength.EntityData.SegmentPath = "ipv4-unicast-prefix-length" + "[prefix-length='" + fmt.Sprintf("%v", ipv4UnicastPrefixLength.PrefixLength) + "']"
    ipv4UnicastPrefixLength.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4UnicastPrefixLength.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4UnicastPrefixLength.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4UnicastPrefixLength.EntityData.Children = make(map[string]types.YChild)
    ipv4UnicastPrefixLength.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4UnicastPrefixLength.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", ipv4UnicastPrefixLength.PrefixLength}
    ipv4UnicastPrefixLength.EntityData.Leafs["ipv4-unicast-prefix-percent"] = types.YLeaf{"Ipv4UnicastPrefixPercent", ipv4UnicastPrefixLength.Ipv4UnicastPrefixPercent}
    return &(ipv4UnicastPrefixLength.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address
// IPv6 table for TCAM LC cards
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unicast table for TCAM LC cards.
    Ipv6Unicast HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast
}

func (ipv6Address *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address) GetEntityData() *types.CommonEntityData {
    ipv6Address.EntityData.YFilter = ipv6Address.YFilter
    ipv6Address.EntityData.YangName = "ipv6-address"
    ipv6Address.EntityData.BundleName = "cisco_ios_xr"
    ipv6Address.EntityData.ParentYangName = "fib-table"
    ipv6Address.EntityData.SegmentPath = "ipv6-address"
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = make(map[string]types.YChild)
    ipv6Address.EntityData.Children["ipv6-unicast"] = types.YChild{"Ipv6Unicast", &ipv6Address.Ipv6Unicast}
    ipv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6Address.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast
// Unicast table for TCAM LC cards
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // curve out percentage of TCAM table entries. The type is interface{} with
    // range: 0..100. Units are percentage.
    Ipv6UnicastPercent interface{}

    // IPv6 Unicast prefix .
    Ipv6UnicastPrefixLengths HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths
}

func (ipv6Unicast *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast) GetEntityData() *types.CommonEntityData {
    ipv6Unicast.EntityData.YFilter = ipv6Unicast.YFilter
    ipv6Unicast.EntityData.YangName = "ipv6-unicast"
    ipv6Unicast.EntityData.BundleName = "cisco_ios_xr"
    ipv6Unicast.EntityData.ParentYangName = "ipv6-address"
    ipv6Unicast.EntityData.SegmentPath = "ipv6-unicast"
    ipv6Unicast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Unicast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Unicast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Unicast.EntityData.Children = make(map[string]types.YChild)
    ipv6Unicast.EntityData.Children["ipv6-unicast-prefix-lengths"] = types.YChild{"Ipv6UnicastPrefixLengths", &ipv6Unicast.Ipv6UnicastPrefixLengths}
    ipv6Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Unicast.EntityData.Leafs["ipv6-unicast-percent"] = types.YLeaf{"Ipv6UnicastPercent", ipv6Unicast.Ipv6UnicastPercent}
    return &(ipv6Unicast.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths
// IPv6 Unicast prefix 
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Unicast prefix length. The type is slice of
    // HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength.
    Ipv6UnicastPrefixLength []HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength
}

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetEntityData() *types.CommonEntityData {
    ipv6UnicastPrefixLengths.EntityData.YFilter = ipv6UnicastPrefixLengths.YFilter
    ipv6UnicastPrefixLengths.EntityData.YangName = "ipv6-unicast-prefix-lengths"
    ipv6UnicastPrefixLengths.EntityData.BundleName = "cisco_ios_xr"
    ipv6UnicastPrefixLengths.EntityData.ParentYangName = "ipv6-unicast"
    ipv6UnicastPrefixLengths.EntityData.SegmentPath = "ipv6-unicast-prefix-lengths"
    ipv6UnicastPrefixLengths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6UnicastPrefixLengths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6UnicastPrefixLengths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6UnicastPrefixLengths.EntityData.Children = make(map[string]types.YChild)
    ipv6UnicastPrefixLengths.EntityData.Children["ipv6-unicast-prefix-length"] = types.YChild{"Ipv6UnicastPrefixLength", nil}
    for i := range ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength {
        ipv6UnicastPrefixLengths.EntityData.Children[types.GetSegmentPath(&ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength[i])] = types.YChild{"Ipv6UnicastPrefixLength", &ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength[i]}
    }
    ipv6UnicastPrefixLengths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6UnicastPrefixLengths.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength
// IPv6 Unicast prefix length
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. prefix length. The type is interface{} with range:
    // 0..128.
    PrefixLength interface{}

    // curve out percentage of TCAM table entries. The type is string. Units are
    // percentage.
    Ipv6UnicastPrefixPercent interface{}
}

func (ipv6UnicastPrefixLength *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength) GetEntityData() *types.CommonEntityData {
    ipv6UnicastPrefixLength.EntityData.YFilter = ipv6UnicastPrefixLength.YFilter
    ipv6UnicastPrefixLength.EntityData.YangName = "ipv6-unicast-prefix-length"
    ipv6UnicastPrefixLength.EntityData.BundleName = "cisco_ios_xr"
    ipv6UnicastPrefixLength.EntityData.ParentYangName = "ipv6-unicast-prefix-lengths"
    ipv6UnicastPrefixLength.EntityData.SegmentPath = "ipv6-unicast-prefix-length" + "[prefix-length='" + fmt.Sprintf("%v", ipv6UnicastPrefixLength.PrefixLength) + "']"
    ipv6UnicastPrefixLength.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6UnicastPrefixLength.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6UnicastPrefixLength.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6UnicastPrefixLength.EntityData.Children = make(map[string]types.YChild)
    ipv6UnicastPrefixLength.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6UnicastPrefixLength.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", ipv6UnicastPrefixLength.PrefixLength}
    ipv6UnicastPrefixLength.EntityData.Leafs["ipv6-unicast-prefix-percent"] = types.YLeaf{"Ipv6UnicastPrefixPercent", ipv6UnicastPrefixLength.Ipv6UnicastPrefixPercent}
    return &(ipv6UnicastPrefixLength.EntityData)
}

// HwModuleProfileConfig_Profile_LoadBalance
// Configure load balance parameters
type HwModuleProfileConfig_Profile_LoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure load balance parameters. The type is interface{} with range:
    // -2147483648..2147483647.
    LoadBalanceProfile interface{}
}

func (loadBalance *HwModuleProfileConfig_Profile_LoadBalance) GetEntityData() *types.CommonEntityData {
    loadBalance.EntityData.YFilter = loadBalance.YFilter
    loadBalance.EntityData.YangName = "load-balance"
    loadBalance.EntityData.BundleName = "cisco_ios_xr"
    loadBalance.EntityData.ParentYangName = "profile"
    loadBalance.EntityData.SegmentPath = "load-balance"
    loadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadBalance.EntityData.Children = make(map[string]types.YChild)
    loadBalance.EntityData.Leafs = make(map[string]types.YLeaf)
    loadBalance.EntityData.Leafs["load-balance-profile"] = types.YLeaf{"LoadBalanceProfile", loadBalance.LoadBalanceProfile}
    return &(loadBalance.EntityData)
}

// HwModuleProfileConfig_Profile_Stats
// Configure stats
type HwModuleProfileConfig_Profile_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure stats for qos-enhanced and acl-permit. The type is interface{}
    // with range: -2147483648..2147483647.
    CounterProfile interface{}
}

func (stats *HwModuleProfileConfig_Profile_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "profile"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["counter-profile"] = types.YLeaf{"CounterProfile", stats.CounterProfile}
    return &(stats.EntityData)
}

// HwModuleProfileConfig_Profile_Netflows
// Configure Netflow profile.
type HwModuleProfileConfig_Profile_Netflows struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of HwModuleProfileConfig_Profile_Netflows_Netflow.
    Netflow []HwModuleProfileConfig_Profile_Netflows_Netflow
}

func (netflows *HwModuleProfileConfig_Profile_Netflows) GetEntityData() *types.CommonEntityData {
    netflows.EntityData.YFilter = netflows.YFilter
    netflows.EntityData.YangName = "netflows"
    netflows.EntityData.BundleName = "cisco_ios_xr"
    netflows.EntityData.ParentYangName = "profile"
    netflows.EntityData.SegmentPath = "netflows"
    netflows.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netflows.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netflows.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netflows.EntityData.Children = make(map[string]types.YChild)
    netflows.EntityData.Children["netflow"] = types.YChild{"Netflow", nil}
    for i := range netflows.Netflow {
        netflows.EntityData.Children[types.GetSegmentPath(&netflows.Netflow[i])] = types.YChild{"Netflow", &netflows.Netflow[i]}
    }
    netflows.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(netflows.EntityData)
}

// HwModuleProfileConfig_Profile_Netflows_Netflow
// none
type HwModuleProfileConfig_Profile_Netflows_Netflow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Ipfix315Enable interface{}

    // This attribute is a key. Location of NETFLOW config. The type is string
    // with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    LocationString interface{}

    // This attribute is a key. Location ID hex to Decimal 0xffff for all. The
    // type is interface{} with range: -2147483648..2147483647.
    LocationId interface{}

    // If Enabled set value to 65535. The type is interface{} with range:
    // -2147483648..2147483647. This attribute is mandatory.
    EnableVal interface{}
}

func (netflow *HwModuleProfileConfig_Profile_Netflows_Netflow) GetEntityData() *types.CommonEntityData {
    netflow.EntityData.YFilter = netflow.YFilter
    netflow.EntityData.YangName = "netflow"
    netflow.EntityData.BundleName = "cisco_ios_xr"
    netflow.EntityData.ParentYangName = "netflows"
    netflow.EntityData.SegmentPath = "netflow" + "[ipfix315-enable='" + fmt.Sprintf("%v", netflow.Ipfix315Enable) + "']" + "[location-string='" + fmt.Sprintf("%v", netflow.LocationString) + "']" + "[location-id='" + fmt.Sprintf("%v", netflow.LocationId) + "']"
    netflow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netflow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netflow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netflow.EntityData.Children = make(map[string]types.YChild)
    netflow.EntityData.Leafs = make(map[string]types.YLeaf)
    netflow.EntityData.Leafs["ipfix315-enable"] = types.YLeaf{"Ipfix315Enable", netflow.Ipfix315Enable}
    netflow.EntityData.Leafs["location-string"] = types.YLeaf{"LocationString", netflow.LocationString}
    netflow.EntityData.Leafs["location-id"] = types.YLeaf{"LocationId", netflow.LocationId}
    netflow.EntityData.Leafs["enable-val"] = types.YLeaf{"EnableVal", netflow.EnableVal}
    return &(netflow.EntityData)
}

// HwModuleProfileConfig_Profile_ProfileAcl
// Configure acl profile
type HwModuleProfileConfig_Profile_ProfileAcl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enabled or disabled. The type is bool.
    Egress interface{}
}

func (profileAcl *HwModuleProfileConfig_Profile_ProfileAcl) GetEntityData() *types.CommonEntityData {
    profileAcl.EntityData.YFilter = profileAcl.YFilter
    profileAcl.EntityData.YangName = "profile-acl"
    profileAcl.EntityData.BundleName = "cisco_ios_xr"
    profileAcl.EntityData.ParentYangName = "profile"
    profileAcl.EntityData.SegmentPath = "profile-acl"
    profileAcl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profileAcl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profileAcl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profileAcl.EntityData.Children = make(map[string]types.YChild)
    profileAcl.EntityData.Leafs = make(map[string]types.YLeaf)
    profileAcl.EntityData.Leafs["egress"] = types.YLeaf{"Egress", profileAcl.Egress}
    return &(profileAcl.EntityData)
}

// HwModuleProfileConfig_Profile_ProfileTcam
// Configure Tcam Profile
type HwModuleProfileConfig_Profile_ProfileTcam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none.
    KeyFormat HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat
}

func (profileTcam *HwModuleProfileConfig_Profile_ProfileTcam) GetEntityData() *types.CommonEntityData {
    profileTcam.EntityData.YFilter = profileTcam.YFilter
    profileTcam.EntityData.YangName = "profile-tcam"
    profileTcam.EntityData.BundleName = "cisco_ios_xr"
    profileTcam.EntityData.ParentYangName = "profile"
    profileTcam.EntityData.SegmentPath = "profile-tcam"
    profileTcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profileTcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profileTcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profileTcam.EntityData.Children = make(map[string]types.YChild)
    profileTcam.EntityData.Children["key-format"] = types.YChild{"KeyFormat", &profileTcam.KeyFormat}
    profileTcam.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(profileTcam.EntityData)
}

// HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat
// none
type HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure acl profile.
    AclTables HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_AclTables
}

func (keyFormat *HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat) GetEntityData() *types.CommonEntityData {
    keyFormat.EntityData.YFilter = keyFormat.YFilter
    keyFormat.EntityData.YangName = "key-format"
    keyFormat.EntityData.BundleName = "cisco_ios_xr"
    keyFormat.EntityData.ParentYangName = "profile-tcam"
    keyFormat.EntityData.SegmentPath = "key-format"
    keyFormat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyFormat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyFormat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyFormat.EntityData.Children = make(map[string]types.YChild)
    keyFormat.EntityData.Children["acl-tables"] = types.YChild{"AclTables", &keyFormat.AclTables}
    keyFormat.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(keyFormat.EntityData)
}

// HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_AclTables
// Configure acl profile
type HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_AclTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure format for acl profile. The type is slice of
    // HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_AclTables_AclTable.
    AclTable []HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_AclTables_AclTable
}

func (aclTables *HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_AclTables) GetEntityData() *types.CommonEntityData {
    aclTables.EntityData.YFilter = aclTables.YFilter
    aclTables.EntityData.YangName = "acl-tables"
    aclTables.EntityData.BundleName = "cisco_ios_xr"
    aclTables.EntityData.ParentYangName = "key-format"
    aclTables.EntityData.SegmentPath = "acl-tables"
    aclTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aclTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aclTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aclTables.EntityData.Children = make(map[string]types.YChild)
    aclTables.EntityData.Children["acl-table"] = types.YChild{"AclTable", nil}
    for i := range aclTables.AclTable {
        aclTables.EntityData.Children[types.GetSegmentPath(&aclTables.AclTable[i])] = types.YChild{"AclTable", &aclTables.AclTable[i]}
    }
    aclTables.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aclTables.EntityData)
}

// HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_AclTables_AclTable
// Configure format for acl profile
type HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_AclTables_AclTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. ipv4/ipv6. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    AddressFamily interface{}

    // This attribute is a key. Location string (all) if for all LCs. The type is
    // string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    LocationString interface{}

    // This attribute is a key. Location ID hex to Decimal 0xffff for all. The
    // type is interface{} with range: -2147483648..2147483647.
    LocationId interface{}

    // Source Address 32 bit qual. The type is interface{} with range:
    // -2147483648..2147483647.
    SourceAddr interface{}

    // Destination Address 32 bit qual. The type is interface{} with range:
    // -2147483648..2147483647.
    DestinationAddr interface{}

    // Source Port. The type is interface{} with range: -2147483648..2147483647.
    SourcePort interface{}

    // Destination Port. The type is interface{} with range:
    // -2147483648..2147483647.
    DestPort interface{}

    // Protocol Type. The type is interface{} with range: -2147483648..2147483647.
    ProtType interface{}

    // TCP Flags. The type is interface{} with range: -2147483648..2147483647.
    TcpFlag interface{}

    // Packet Length. The type is interface{} with range: -2147483648..2147483647.
    PackLen interface{}

    // Fragment Bit. The type is interface{} with range: -2147483648..2147483647.
    FragBit interface{}

    // Precedence. The type is interface{} with range: -2147483648..2147483647.
    Precedence interface{}

    // PortRange. The type is interface{} with range: -2147483648..2147483647.
    PortRange interface{}

    // UDF name. The type is string.
    Udf1 interface{}

    // UDF name. The type is string.
    Udf2 interface{}

    // UDF name. The type is string.
    Udf3 interface{}

    // UDF name. The type is string.
    Udf4 interface{}

    // UDF name. The type is string.
    Udf5 interface{}

    // UDF name. The type is string.
    Udf6 interface{}

    // UDF name. The type is string.
    Udf7 interface{}

    // UDF name. The type is string.
    Udf8 interface{}

    // Enable Capture. The type is interface{} with range:
    // -2147483648..2147483647.
    EnCapture interface{}

    // Enable Setting TTL. The type is interface{} with range:
    // -2147483648..2147483647.
    EnTtl interface{}

    // Enable Matching TTL. The type is interface{} with range:
    // -2147483648..2147483647.
    EnMatch interface{}

    // Enable Non Shared Interface ACL. The type is interface{} with range:
    // -2147483648..2147483647.
    EnShareAcl interface{}
}

func (aclTable *HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_AclTables_AclTable) GetEntityData() *types.CommonEntityData {
    aclTable.EntityData.YFilter = aclTable.YFilter
    aclTable.EntityData.YangName = "acl-table"
    aclTable.EntityData.BundleName = "cisco_ios_xr"
    aclTable.EntityData.ParentYangName = "acl-tables"
    aclTable.EntityData.SegmentPath = "acl-table" + "[address-family='" + fmt.Sprintf("%v", aclTable.AddressFamily) + "']" + "[location-string='" + fmt.Sprintf("%v", aclTable.LocationString) + "']" + "[location-id='" + fmt.Sprintf("%v", aclTable.LocationId) + "']"
    aclTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aclTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aclTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aclTable.EntityData.Children = make(map[string]types.YChild)
    aclTable.EntityData.Leafs = make(map[string]types.YLeaf)
    aclTable.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", aclTable.AddressFamily}
    aclTable.EntityData.Leafs["location-string"] = types.YLeaf{"LocationString", aclTable.LocationString}
    aclTable.EntityData.Leafs["location-id"] = types.YLeaf{"LocationId", aclTable.LocationId}
    aclTable.EntityData.Leafs["source-addr"] = types.YLeaf{"SourceAddr", aclTable.SourceAddr}
    aclTable.EntityData.Leafs["destination-addr"] = types.YLeaf{"DestinationAddr", aclTable.DestinationAddr}
    aclTable.EntityData.Leafs["source-port"] = types.YLeaf{"SourcePort", aclTable.SourcePort}
    aclTable.EntityData.Leafs["dest-port"] = types.YLeaf{"DestPort", aclTable.DestPort}
    aclTable.EntityData.Leafs["prot-type"] = types.YLeaf{"ProtType", aclTable.ProtType}
    aclTable.EntityData.Leafs["tcp-flag"] = types.YLeaf{"TcpFlag", aclTable.TcpFlag}
    aclTable.EntityData.Leafs["pack-len"] = types.YLeaf{"PackLen", aclTable.PackLen}
    aclTable.EntityData.Leafs["frag-bit"] = types.YLeaf{"FragBit", aclTable.FragBit}
    aclTable.EntityData.Leafs["precedence"] = types.YLeaf{"Precedence", aclTable.Precedence}
    aclTable.EntityData.Leafs["port-range"] = types.YLeaf{"PortRange", aclTable.PortRange}
    aclTable.EntityData.Leafs["udf1"] = types.YLeaf{"Udf1", aclTable.Udf1}
    aclTable.EntityData.Leafs["udf2"] = types.YLeaf{"Udf2", aclTable.Udf2}
    aclTable.EntityData.Leafs["udf3"] = types.YLeaf{"Udf3", aclTable.Udf3}
    aclTable.EntityData.Leafs["udf4"] = types.YLeaf{"Udf4", aclTable.Udf4}
    aclTable.EntityData.Leafs["udf5"] = types.YLeaf{"Udf5", aclTable.Udf5}
    aclTable.EntityData.Leafs["udf6"] = types.YLeaf{"Udf6", aclTable.Udf6}
    aclTable.EntityData.Leafs["udf7"] = types.YLeaf{"Udf7", aclTable.Udf7}
    aclTable.EntityData.Leafs["udf8"] = types.YLeaf{"Udf8", aclTable.Udf8}
    aclTable.EntityData.Leafs["en-capture"] = types.YLeaf{"EnCapture", aclTable.EnCapture}
    aclTable.EntityData.Leafs["en-ttl"] = types.YLeaf{"EnTtl", aclTable.EnTtl}
    aclTable.EntityData.Leafs["en-match"] = types.YLeaf{"EnMatch", aclTable.EnMatch}
    aclTable.EntityData.Leafs["en-share-acl"] = types.YLeaf{"EnShareAcl", aclTable.EnShareAcl}
    return &(aclTable.EntityData)
}

// HwModuleProfileConfig_Profile_Qos
// Configure profile.
type HwModuleProfileConfig_Profile_Qos struct {
    EntityData types.CommonEntityData
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

func (qos *HwModuleProfileConfig_Profile_Qos) GetEntityData() *types.CommonEntityData {
    qos.EntityData.YFilter = qos.YFilter
    qos.EntityData.YangName = "qos"
    qos.EntityData.BundleName = "cisco_ios_xr"
    qos.EntityData.ParentYangName = "profile"
    qos.EntityData.SegmentPath = "qos"
    qos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qos.EntityData.Children = make(map[string]types.YChild)
    qos.EntityData.Children["hqos-enable-all"] = types.YChild{"HqosEnableAll", &qos.HqosEnableAll}
    qos.EntityData.Children["ingress-model-root-def"] = types.YChild{"IngressModelRootDef", &qos.IngressModelRootDef}
    qos.EntityData.Children["ingress-models"] = types.YChild{"IngressModels", &qos.IngressModels}
    qos.EntityData.Children["trunks"] = types.YChild{"Trunks", &qos.Trunks}
    qos.EntityData.Children["class-maps-root-def"] = types.YChild{"ClassMapsRootDef", &qos.ClassMapsRootDef}
    qos.EntityData.Children["class-maps"] = types.YChild{"ClassMaps", &qos.ClassMaps}
    qos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qos.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_HqosEnableAll
// Configure Hqos profile
type HwModuleProfileConfig_Profile_Qos_HqosEnableAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hqos profile value. The type is interface{} with range:
    // -2147483648..2147483647.
    HqosEnable interface{}
}

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetEntityData() *types.CommonEntityData {
    hqosEnableAll.EntityData.YFilter = hqosEnableAll.YFilter
    hqosEnableAll.EntityData.YangName = "hqos-enable-all"
    hqosEnableAll.EntityData.BundleName = "cisco_ios_xr"
    hqosEnableAll.EntityData.ParentYangName = "qos"
    hqosEnableAll.EntityData.SegmentPath = "hqos-enable-all"
    hqosEnableAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hqosEnableAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hqosEnableAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hqosEnableAll.EntityData.Children = make(map[string]types.YChild)
    hqosEnableAll.EntityData.Leafs = make(map[string]types.YLeaf)
    hqosEnableAll.EntityData.Leafs["hqos-enable"] = types.YLeaf{"HqosEnable", hqosEnableAll.HqosEnable}
    return &(hqosEnableAll.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_IngressModelRootDef
// Configure Ingress Model Default
type HwModuleProfileConfig_Profile_Qos_IngressModelRootDef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ingress Model Default. The type is interface{} with range:
    // -2147483648..2147483647.
    IngressModelLeafDef interface{}
}

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetEntityData() *types.CommonEntityData {
    ingressModelRootDef.EntityData.YFilter = ingressModelRootDef.YFilter
    ingressModelRootDef.EntityData.YangName = "ingress-model-root-def"
    ingressModelRootDef.EntityData.BundleName = "cisco_ios_xr"
    ingressModelRootDef.EntityData.ParentYangName = "qos"
    ingressModelRootDef.EntityData.SegmentPath = "ingress-model-root-def"
    ingressModelRootDef.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressModelRootDef.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressModelRootDef.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressModelRootDef.EntityData.Children = make(map[string]types.YChild)
    ingressModelRootDef.EntityData.Leafs = make(map[string]types.YLeaf)
    ingressModelRootDef.EntityData.Leafs["ingress-model-leaf-def"] = types.YLeaf{"IngressModelLeafDef", ingressModelRootDef.IngressModelLeafDef}
    return &(ingressModelRootDef.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_IngressModels
// Configure Ingress Model Root
type HwModuleProfileConfig_Profile_Qos_IngressModels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Ingress Model. The type is slice of
    // HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel.
    IngressModel []HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel
}

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetEntityData() *types.CommonEntityData {
    ingressModels.EntityData.YFilter = ingressModels.YFilter
    ingressModels.EntityData.YangName = "ingress-models"
    ingressModels.EntityData.BundleName = "cisco_ios_xr"
    ingressModels.EntityData.ParentYangName = "qos"
    ingressModels.EntityData.SegmentPath = "ingress-models"
    ingressModels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressModels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressModels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressModels.EntityData.Children = make(map[string]types.YChild)
    ingressModels.EntityData.Children["ingress-model"] = types.YChild{"IngressModel", nil}
    for i := range ingressModels.IngressModel {
        ingressModels.EntityData.Children[types.GetSegmentPath(&ingressModels.IngressModel[i])] = types.YChild{"IngressModel", &ingressModels.IngressModel[i]}
    }
    ingressModels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ingressModels.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel
// Configure Ingress Model
type HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. NodeName. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    NodeName interface{}

    // Configure Ingress Model Leaf. The type is slice of
    // HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf.
    IngressModelLeaf []HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf
}

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetEntityData() *types.CommonEntityData {
    ingressModel.EntityData.YFilter = ingressModel.YFilter
    ingressModel.EntityData.YangName = "ingress-model"
    ingressModel.EntityData.BundleName = "cisco_ios_xr"
    ingressModel.EntityData.ParentYangName = "ingress-models"
    ingressModel.EntityData.SegmentPath = "ingress-model" + "[node-name='" + fmt.Sprintf("%v", ingressModel.NodeName) + "']"
    ingressModel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressModel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressModel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressModel.EntityData.Children = make(map[string]types.YChild)
    ingressModel.EntityData.Children["ingress-model-leaf"] = types.YChild{"IngressModelLeaf", nil}
    for i := range ingressModel.IngressModelLeaf {
        ingressModel.EntityData.Children[types.GetSegmentPath(&ingressModel.IngressModelLeaf[i])] = types.YChild{"IngressModelLeaf", &ingressModel.IngressModelLeaf[i]}
    }
    ingressModel.EntityData.Leafs = make(map[string]types.YLeaf)
    ingressModel.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", ingressModel.NodeName}
    return &(ingressModel.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf
// Configure Ingress Model Leaf
type HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is interface{} with range:
    // -2147483648..2147483647.
    Location interface{}

    // IngressModelLeaf. The type is interface{} with range:
    // -2147483648..2147483647. This attribute is mandatory.
    IngressModelLeaf interface{}
}

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetEntityData() *types.CommonEntityData {
    ingressModelLeaf.EntityData.YFilter = ingressModelLeaf.YFilter
    ingressModelLeaf.EntityData.YangName = "ingress-model-leaf"
    ingressModelLeaf.EntityData.BundleName = "cisco_ios_xr"
    ingressModelLeaf.EntityData.ParentYangName = "ingress-model"
    ingressModelLeaf.EntityData.SegmentPath = "ingress-model-leaf" + "[location='" + fmt.Sprintf("%v", ingressModelLeaf.Location) + "']"
    ingressModelLeaf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressModelLeaf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressModelLeaf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressModelLeaf.EntityData.Children = make(map[string]types.YChild)
    ingressModelLeaf.EntityData.Leafs = make(map[string]types.YLeaf)
    ingressModelLeaf.EntityData.Leafs["location"] = types.YLeaf{"Location", ingressModelLeaf.Location}
    ingressModelLeaf.EntityData.Leafs["ingress-model-leaf"] = types.YLeaf{"IngressModelLeaf", ingressModelLeaf.IngressModelLeaf}
    return &(ingressModelLeaf.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_Trunks
// Configure Max Trunk Size
type HwModuleProfileConfig_Profile_Qos_Trunks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Trunk Size. The type is interface{} with range:
    // -2147483648..2147483647.
    TrunkSize interface{}
}

func (trunks *HwModuleProfileConfig_Profile_Qos_Trunks) GetEntityData() *types.CommonEntityData {
    trunks.EntityData.YFilter = trunks.YFilter
    trunks.EntityData.YangName = "trunks"
    trunks.EntityData.BundleName = "cisco_ios_xr"
    trunks.EntityData.ParentYangName = "qos"
    trunks.EntityData.SegmentPath = "trunks"
    trunks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trunks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trunks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trunks.EntityData.Children = make(map[string]types.YChild)
    trunks.EntityData.Leafs = make(map[string]types.YLeaf)
    trunks.EntityData.Leafs["trunk-size"] = types.YLeaf{"TrunkSize", trunks.TrunkSize}
    return &(trunks.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef
// Configure Class Maps Default
type HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class Map Size Default. The type is interface{} with range:
    // -2147483648..2147483647.
    ClassMapSizeDef interface{}
}

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetEntityData() *types.CommonEntityData {
    classMapsRootDef.EntityData.YFilter = classMapsRootDef.YFilter
    classMapsRootDef.EntityData.YangName = "class-maps-root-def"
    classMapsRootDef.EntityData.BundleName = "cisco_ios_xr"
    classMapsRootDef.EntityData.ParentYangName = "qos"
    classMapsRootDef.EntityData.SegmentPath = "class-maps-root-def"
    classMapsRootDef.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMapsRootDef.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMapsRootDef.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMapsRootDef.EntityData.Children = make(map[string]types.YChild)
    classMapsRootDef.EntityData.Leafs = make(map[string]types.YLeaf)
    classMapsRootDef.EntityData.Leafs["class-map-size-def"] = types.YLeaf{"ClassMapSizeDef", classMapsRootDef.ClassMapSizeDef}
    return &(classMapsRootDef.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_ClassMaps
// Configure Class Map Root
type HwModuleProfileConfig_Profile_Qos_ClassMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Class Maps. The type is slice of
    // HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap.
    ClassMap []HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap
}

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetEntityData() *types.CommonEntityData {
    classMaps.EntityData.YFilter = classMaps.YFilter
    classMaps.EntityData.YangName = "class-maps"
    classMaps.EntityData.BundleName = "cisco_ios_xr"
    classMaps.EntityData.ParentYangName = "qos"
    classMaps.EntityData.SegmentPath = "class-maps"
    classMaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMaps.EntityData.Children = make(map[string]types.YChild)
    classMaps.EntityData.Children["class-map"] = types.YChild{"ClassMap", nil}
    for i := range classMaps.ClassMap {
        classMaps.EntityData.Children[types.GetSegmentPath(&classMaps.ClassMap[i])] = types.YChild{"ClassMap", &classMaps.ClassMap[i]}
    }
    classMaps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(classMaps.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap
// Configure Class Maps
type HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. NodeName. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    NodeName interface{}

    // Class Map Size. The type is slice of
    // HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize.
    ClassMapSize []HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize
}

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetEntityData() *types.CommonEntityData {
    classMap.EntityData.YFilter = classMap.YFilter
    classMap.EntityData.YangName = "class-map"
    classMap.EntityData.BundleName = "cisco_ios_xr"
    classMap.EntityData.ParentYangName = "class-maps"
    classMap.EntityData.SegmentPath = "class-map" + "[node-name='" + fmt.Sprintf("%v", classMap.NodeName) + "']"
    classMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMap.EntityData.Children = make(map[string]types.YChild)
    classMap.EntityData.Children["class-map-size"] = types.YChild{"ClassMapSize", nil}
    for i := range classMap.ClassMapSize {
        classMap.EntityData.Children[types.GetSegmentPath(&classMap.ClassMapSize[i])] = types.YChild{"ClassMapSize", &classMap.ClassMapSize[i]}
    }
    classMap.EntityData.Leafs = make(map[string]types.YLeaf)
    classMap.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", classMap.NodeName}
    return &(classMap.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize
// Class Map Size
type HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is interface{} with range:
    // -2147483648..2147483647.
    Location interface{}

    // ClassMapSize. The type is interface{} with range: -2147483648..2147483647.
    // This attribute is mandatory.
    ClassMapSize interface{}
}

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetEntityData() *types.CommonEntityData {
    classMapSize.EntityData.YFilter = classMapSize.YFilter
    classMapSize.EntityData.YangName = "class-map-size"
    classMapSize.EntityData.BundleName = "cisco_ios_xr"
    classMapSize.EntityData.ParentYangName = "class-map"
    classMapSize.EntityData.SegmentPath = "class-map-size" + "[location='" + fmt.Sprintf("%v", classMapSize.Location) + "']"
    classMapSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMapSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMapSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMapSize.EntityData.Children = make(map[string]types.YChild)
    classMapSize.EntityData.Leafs = make(map[string]types.YLeaf)
    classMapSize.EntityData.Leafs["location"] = types.YLeaf{"Location", classMapSize.Location}
    classMapSize.EntityData.Leafs["class-map-size"] = types.YLeaf{"ClassMapSize", classMapSize.ClassMapSize}
    return &(classMapSize.EntityData)
}

// HwModuleProfileConfig_FibScale
// Configure Fib for Scale for noTcam LC.
type HwModuleProfileConfig_FibScale struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 table for NOTCAM LC Scale.
    Ipv6UnicastScaleNoTcam HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam

    // IPv4 table for NOTCAM LC Scale.
    Ipv4UnicastScaleNoTcam HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam
}

func (fibScale *HwModuleProfileConfig_FibScale) GetEntityData() *types.CommonEntityData {
    fibScale.EntityData.YFilter = fibScale.YFilter
    fibScale.EntityData.YangName = "fib-scale"
    fibScale.EntityData.BundleName = "cisco_ios_xr"
    fibScale.EntityData.ParentYangName = "hw-module-profile-config"
    fibScale.EntityData.SegmentPath = "fib-scale"
    fibScale.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibScale.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibScale.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibScale.EntityData.Children = make(map[string]types.YChild)
    fibScale.EntityData.Children["ipv6-unicast-scale-no-tcam"] = types.YChild{"Ipv6UnicastScaleNoTcam", &fibScale.Ipv6UnicastScaleNoTcam}
    fibScale.EntityData.Children["ipv4-unicast-scale-no-tcam"] = types.YChild{"Ipv4UnicastScaleNoTcam", &fibScale.Ipv4UnicastScaleNoTcam}
    fibScale.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fibScale.EntityData)
}

// HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam
// IPv6 table for NOTCAM LC Scale.
type HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Scale for IPv6 table for NoTCAM LC.
    ScaleIpv6NoTcam HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam
}

func (ipv6UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam) GetEntityData() *types.CommonEntityData {
    ipv6UnicastScaleNoTcam.EntityData.YFilter = ipv6UnicastScaleNoTcam.YFilter
    ipv6UnicastScaleNoTcam.EntityData.YangName = "ipv6-unicast-scale-no-tcam"
    ipv6UnicastScaleNoTcam.EntityData.BundleName = "cisco_ios_xr"
    ipv6UnicastScaleNoTcam.EntityData.ParentYangName = "fib-scale"
    ipv6UnicastScaleNoTcam.EntityData.SegmentPath = "ipv6-unicast-scale-no-tcam"
    ipv6UnicastScaleNoTcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6UnicastScaleNoTcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6UnicastScaleNoTcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6UnicastScaleNoTcam.EntityData.Children = make(map[string]types.YChild)
    ipv6UnicastScaleNoTcam.EntityData.Children["scale-ipv6-no-tcam"] = types.YChild{"ScaleIpv6NoTcam", &ipv6UnicastScaleNoTcam.ScaleIpv6NoTcam}
    ipv6UnicastScaleNoTcam.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6UnicastScaleNoTcam.EntityData)
}

// HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam
// Scale for IPv6 table for NoTCAM LC.
type HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Internet-optimized Scale for IPv6 table for NoTCAM LC. The type is string.
    InternetOptimizedIpv6NoTcam interface{}
}

func (scaleIpv6NoTcam *HwModuleProfileConfig_FibScale_Ipv6UnicastScaleNoTcam_ScaleIpv6NoTcam) GetEntityData() *types.CommonEntityData {
    scaleIpv6NoTcam.EntityData.YFilter = scaleIpv6NoTcam.YFilter
    scaleIpv6NoTcam.EntityData.YangName = "scale-ipv6-no-tcam"
    scaleIpv6NoTcam.EntityData.BundleName = "cisco_ios_xr"
    scaleIpv6NoTcam.EntityData.ParentYangName = "ipv6-unicast-scale-no-tcam"
    scaleIpv6NoTcam.EntityData.SegmentPath = "scale-ipv6-no-tcam"
    scaleIpv6NoTcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scaleIpv6NoTcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scaleIpv6NoTcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scaleIpv6NoTcam.EntityData.Children = make(map[string]types.YChild)
    scaleIpv6NoTcam.EntityData.Leafs = make(map[string]types.YLeaf)
    scaleIpv6NoTcam.EntityData.Leafs["internet-optimized-ipv6-no-tcam"] = types.YLeaf{"InternetOptimizedIpv6NoTcam", scaleIpv6NoTcam.InternetOptimizedIpv6NoTcam}
    return &(scaleIpv6NoTcam.EntityData)
}

// HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam
// IPv4 table for NOTCAM LC Scale.
type HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Scale for IPv4 table for NoTCAM LC.
    ScaleIpv4NoTcam HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam
}

func (ipv4UnicastScaleNoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam) GetEntityData() *types.CommonEntityData {
    ipv4UnicastScaleNoTcam.EntityData.YFilter = ipv4UnicastScaleNoTcam.YFilter
    ipv4UnicastScaleNoTcam.EntityData.YangName = "ipv4-unicast-scale-no-tcam"
    ipv4UnicastScaleNoTcam.EntityData.BundleName = "cisco_ios_xr"
    ipv4UnicastScaleNoTcam.EntityData.ParentYangName = "fib-scale"
    ipv4UnicastScaleNoTcam.EntityData.SegmentPath = "ipv4-unicast-scale-no-tcam"
    ipv4UnicastScaleNoTcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4UnicastScaleNoTcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4UnicastScaleNoTcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4UnicastScaleNoTcam.EntityData.Children = make(map[string]types.YChild)
    ipv4UnicastScaleNoTcam.EntityData.Children["scale-ipv4-no-tcam"] = types.YChild{"ScaleIpv4NoTcam", &ipv4UnicastScaleNoTcam.ScaleIpv4NoTcam}
    ipv4UnicastScaleNoTcam.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4UnicastScaleNoTcam.EntityData)
}

// HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam
// Scale for IPv4 table for NoTCAM LC.
type HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Optimized Scale for IPv4 table for NoTCAM LC. The type is string.
    OptimizedIpv4NoTcam interface{}
}

func (scaleIpv4NoTcam *HwModuleProfileConfig_FibScale_Ipv4UnicastScaleNoTcam_ScaleIpv4NoTcam) GetEntityData() *types.CommonEntityData {
    scaleIpv4NoTcam.EntityData.YFilter = scaleIpv4NoTcam.YFilter
    scaleIpv4NoTcam.EntityData.YangName = "scale-ipv4-no-tcam"
    scaleIpv4NoTcam.EntityData.BundleName = "cisco_ios_xr"
    scaleIpv4NoTcam.EntityData.ParentYangName = "ipv4-unicast-scale-no-tcam"
    scaleIpv4NoTcam.EntityData.SegmentPath = "scale-ipv4-no-tcam"
    scaleIpv4NoTcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scaleIpv4NoTcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scaleIpv4NoTcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scaleIpv4NoTcam.EntityData.Children = make(map[string]types.YChild)
    scaleIpv4NoTcam.EntityData.Leafs = make(map[string]types.YLeaf)
    scaleIpv4NoTcam.EntityData.Leafs["optimized-ipv4-no-tcam"] = types.YLeaf{"OptimizedIpv4NoTcam", scaleIpv4NoTcam.OptimizedIpv4NoTcam}
    return &(scaleIpv4NoTcam.EntityData)
}

// HwModuleProfileConfig_Tcam
// Configure Tcam.
type HwModuleProfileConfig_Tcam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Fib iscale for Tcam.
    FibTcamScale HwModuleProfileConfig_Tcam_FibTcamScale
}

func (tcam *HwModuleProfileConfig_Tcam) GetEntityData() *types.CommonEntityData {
    tcam.EntityData.YFilter = tcam.YFilter
    tcam.EntityData.YangName = "tcam"
    tcam.EntityData.BundleName = "cisco_ios_xr"
    tcam.EntityData.ParentYangName = "hw-module-profile-config"
    tcam.EntityData.SegmentPath = "tcam"
    tcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcam.EntityData.Children = make(map[string]types.YChild)
    tcam.EntityData.Children["fib-tcam-scale"] = types.YChild{"FibTcamScale", &tcam.FibTcamScale}
    tcam.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tcam.EntityData)
}

// HwModuleProfileConfig_Tcam_FibTcamScale
// Configure Fib iscale for Tcam.
type HwModuleProfileConfig_Tcam_FibTcamScale struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 table for TCAM LC Scale.
    Ipv4UnicastScale HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale
}

func (fibTcamScale *HwModuleProfileConfig_Tcam_FibTcamScale) GetEntityData() *types.CommonEntityData {
    fibTcamScale.EntityData.YFilter = fibTcamScale.YFilter
    fibTcamScale.EntityData.YangName = "fib-tcam-scale"
    fibTcamScale.EntityData.BundleName = "cisco_ios_xr"
    fibTcamScale.EntityData.ParentYangName = "tcam"
    fibTcamScale.EntityData.SegmentPath = "fib-tcam-scale"
    fibTcamScale.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibTcamScale.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibTcamScale.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibTcamScale.EntityData.Children = make(map[string]types.YChild)
    fibTcamScale.EntityData.Children["ipv4-unicast-scale"] = types.YChild{"Ipv4UnicastScale", &fibTcamScale.Ipv4UnicastScale}
    fibTcamScale.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fibTcamScale.EntityData)
}

// HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale
// IPv4 table for TCAM LC Scale.
type HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Scale for IPv4 table for TCAM LC. The type is interface{}.
    Ipv4Scale interface{}
}

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetEntityData() *types.CommonEntityData {
    ipv4UnicastScale.EntityData.YFilter = ipv4UnicastScale.YFilter
    ipv4UnicastScale.EntityData.YangName = "ipv4-unicast-scale"
    ipv4UnicastScale.EntityData.BundleName = "cisco_ios_xr"
    ipv4UnicastScale.EntityData.ParentYangName = "fib-tcam-scale"
    ipv4UnicastScale.EntityData.SegmentPath = "ipv4-unicast-scale"
    ipv4UnicastScale.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4UnicastScale.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4UnicastScale.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4UnicastScale.EntityData.Children = make(map[string]types.YChild)
    ipv4UnicastScale.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4UnicastScale.EntityData.Leafs["ipv4-scale"] = types.YLeaf{"Ipv4Scale", ipv4UnicastScale.Ipv4Scale}
    return &(ipv4UnicastScale.EntityData)
}

