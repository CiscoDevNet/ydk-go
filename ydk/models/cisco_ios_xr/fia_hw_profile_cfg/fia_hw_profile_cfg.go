// This module contains a collection of YANG definitions
// for Cisco IOS-XR fia-hw-profile package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hw-module-profile-config: none
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

    // Configure OLR.
    OrchestratedLinecardReload HwModuleProfileConfig_OrchestratedLinecardReload

    // Configure Tcam.
    Tcam HwModuleProfileConfig_Tcam

    // Configure profile.
    Qosqppb HwModuleProfileConfig_Qosqppb
}

func (hwModuleProfileConfig *HwModuleProfileConfig) GetEntityData() *types.CommonEntityData {
    hwModuleProfileConfig.EntityData.YFilter = hwModuleProfileConfig.YFilter
    hwModuleProfileConfig.EntityData.YangName = "hw-module-profile-config"
    hwModuleProfileConfig.EntityData.BundleName = "cisco_ios_xr"
    hwModuleProfileConfig.EntityData.ParentYangName = "Cisco-IOS-XR-fia-hw-profile-cfg"
    hwModuleProfileConfig.EntityData.SegmentPath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config"
    hwModuleProfileConfig.EntityData.AbsolutePath = hwModuleProfileConfig.EntityData.SegmentPath
    hwModuleProfileConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwModuleProfileConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwModuleProfileConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwModuleProfileConfig.EntityData.Children = types.NewOrderedMap()
    hwModuleProfileConfig.EntityData.Children.Append("profile", types.YChild{"Profile", &hwModuleProfileConfig.Profile})
    hwModuleProfileConfig.EntityData.Children.Append("fib-scale", types.YChild{"FibScale", &hwModuleProfileConfig.FibScale})
    hwModuleProfileConfig.EntityData.Children.Append("orchestrated-linecard-reload", types.YChild{"OrchestratedLinecardReload", &hwModuleProfileConfig.OrchestratedLinecardReload})
    hwModuleProfileConfig.EntityData.Children.Append("tcam", types.YChild{"Tcam", &hwModuleProfileConfig.Tcam})
    hwModuleProfileConfig.EntityData.Children.Append("qosqppb", types.YChild{"Qosqppb", &hwModuleProfileConfig.Qosqppb})
    hwModuleProfileConfig.EntityData.Leafs = types.NewOrderedMap()

    hwModuleProfileConfig.EntityData.YListKeys = []string {}

    return &(hwModuleProfileConfig.EntityData)
}

// HwModuleProfileConfig_Profile
// Configure profile.
type HwModuleProfileConfig_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure profile for TCAM LC cards.
    TcamTable HwModuleProfileConfig_Profile_TcamTable

    // Configure Netflow profile.
    Netflow HwModuleProfileConfig_Profile_Netflow

    // Configure Flowspec profile.
    Flowspecs HwModuleProfileConfig_Profile_Flowspecs

    // Configure Segment Routing profile.
    SegmentRoutings HwModuleProfileConfig_Profile_SegmentRoutings

    // Configure load balance parameters.
    LoadBalance HwModuleProfileConfig_Profile_LoadBalance

    // Configure stats.
    Stats HwModuleProfileConfig_Profile_Stats

    // Configure acl profile.
    ProfileAcl HwModuleProfileConfig_Profile_ProfileAcl

    // Configure Bundle profile.
    BundleScale HwModuleProfileConfig_Profile_BundleScale

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
    profile.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/" + profile.EntityData.SegmentPath
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Children.Append("tcam-table", types.YChild{"TcamTable", &profile.TcamTable})
    profile.EntityData.Children.Append("netflow", types.YChild{"Netflow", &profile.Netflow})
    profile.EntityData.Children.Append("flowspecs", types.YChild{"Flowspecs", &profile.Flowspecs})
    profile.EntityData.Children.Append("segment-routings", types.YChild{"SegmentRoutings", &profile.SegmentRoutings})
    profile.EntityData.Children.Append("load-balance", types.YChild{"LoadBalance", &profile.LoadBalance})
    profile.EntityData.Children.Append("stats", types.YChild{"Stats", &profile.Stats})
    profile.EntityData.Children.Append("profile-acl", types.YChild{"ProfileAcl", &profile.ProfileAcl})
    profile.EntityData.Children.Append("bundle-scale", types.YChild{"BundleScale", &profile.BundleScale})
    profile.EntityData.Children.Append("profile-tcam", types.YChild{"ProfileTcam", &profile.ProfileTcam})
    profile.EntityData.Children.Append("qos", types.YChild{"Qos", &profile.Qos})
    profile.EntityData.Leafs = types.NewOrderedMap()

    profile.EntityData.YListKeys = []string {}

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
    tcamTable.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/" + tcamTable.EntityData.SegmentPath
    tcamTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcamTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcamTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcamTable.EntityData.Children = types.NewOrderedMap()
    tcamTable.EntityData.Children.Append("fib-table", types.YChild{"FibTable", &tcamTable.FibTable})
    tcamTable.EntityData.Leafs = types.NewOrderedMap()

    tcamTable.EntityData.YListKeys = []string {}

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
    fibTable.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/tcam-table/" + fibTable.EntityData.SegmentPath
    fibTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibTable.EntityData.Children = types.NewOrderedMap()
    fibTable.EntityData.Children.Append("ipv4-address", types.YChild{"Ipv4Address", &fibTable.Ipv4Address})
    fibTable.EntityData.Children.Append("ipv6-address", types.YChild{"Ipv6Address", &fibTable.Ipv6Address})
    fibTable.EntityData.Leafs = types.NewOrderedMap()

    fibTable.EntityData.YListKeys = []string {}

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
    ipv4Address.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/tcam-table/fib-table/" + ipv4Address.EntityData.SegmentPath
    ipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Address.EntityData.Children = types.NewOrderedMap()
    ipv4Address.EntityData.Children.Append("ipv4-unicast", types.YChild{"Ipv4Unicast", &ipv4Address.Ipv4Unicast})
    ipv4Address.EntityData.Leafs = types.NewOrderedMap()

    ipv4Address.EntityData.YListKeys = []string {}

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
    ipv4Unicast.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/tcam-table/fib-table/ipv4-address/" + ipv4Unicast.EntityData.SegmentPath
    ipv4Unicast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Unicast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Unicast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Unicast.EntityData.Children = types.NewOrderedMap()
    ipv4Unicast.EntityData.Children.Append("ipv4-unicast-prefix-lengths", types.YChild{"Ipv4UnicastPrefixLengths", &ipv4Unicast.Ipv4UnicastPrefixLengths})
    ipv4Unicast.EntityData.Leafs = types.NewOrderedMap()
    ipv4Unicast.EntityData.Leafs.Append("ipv4-unicast-percent", types.YLeaf{"Ipv4UnicastPercent", ipv4Unicast.Ipv4UnicastPercent})

    ipv4Unicast.EntityData.YListKeys = []string {}

    return &(ipv4Unicast.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths
// IPv4 Unicast prefix 
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Unicast prefix length. The type is slice of
    // HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength.
    Ipv4UnicastPrefixLength []*HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength
}

func (ipv4UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths) GetEntityData() *types.CommonEntityData {
    ipv4UnicastPrefixLengths.EntityData.YFilter = ipv4UnicastPrefixLengths.YFilter
    ipv4UnicastPrefixLengths.EntityData.YangName = "ipv4-unicast-prefix-lengths"
    ipv4UnicastPrefixLengths.EntityData.BundleName = "cisco_ios_xr"
    ipv4UnicastPrefixLengths.EntityData.ParentYangName = "ipv4-unicast"
    ipv4UnicastPrefixLengths.EntityData.SegmentPath = "ipv4-unicast-prefix-lengths"
    ipv4UnicastPrefixLengths.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/tcam-table/fib-table/ipv4-address/ipv4-unicast/" + ipv4UnicastPrefixLengths.EntityData.SegmentPath
    ipv4UnicastPrefixLengths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4UnicastPrefixLengths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4UnicastPrefixLengths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4UnicastPrefixLengths.EntityData.Children = types.NewOrderedMap()
    ipv4UnicastPrefixLengths.EntityData.Children.Append("ipv4-unicast-prefix-length", types.YChild{"Ipv4UnicastPrefixLength", nil})
    for i := range ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength {
        ipv4UnicastPrefixLengths.EntityData.Children.Append(types.GetSegmentPath(ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength[i]), types.YChild{"Ipv4UnicastPrefixLength", ipv4UnicastPrefixLengths.Ipv4UnicastPrefixLength[i]})
    }
    ipv4UnicastPrefixLengths.EntityData.Leafs = types.NewOrderedMap()

    ipv4UnicastPrefixLengths.EntityData.YListKeys = []string {}

    return &(ipv4UnicastPrefixLengths.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength
// IPv4 Unicast prefix length
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv4Address_Ipv4Unicast_Ipv4UnicastPrefixLengths_Ipv4UnicastPrefixLength struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    ipv4UnicastPrefixLength.EntityData.SegmentPath = "ipv4-unicast-prefix-length" + types.AddKeyToken(ipv4UnicastPrefixLength.PrefixLength, "prefix-length")
    ipv4UnicastPrefixLength.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/tcam-table/fib-table/ipv4-address/ipv4-unicast/ipv4-unicast-prefix-lengths/" + ipv4UnicastPrefixLength.EntityData.SegmentPath
    ipv4UnicastPrefixLength.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4UnicastPrefixLength.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4UnicastPrefixLength.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4UnicastPrefixLength.EntityData.Children = types.NewOrderedMap()
    ipv4UnicastPrefixLength.EntityData.Leafs = types.NewOrderedMap()
    ipv4UnicastPrefixLength.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", ipv4UnicastPrefixLength.PrefixLength})
    ipv4UnicastPrefixLength.EntityData.Leafs.Append("ipv4-unicast-prefix-percent", types.YLeaf{"Ipv4UnicastPrefixPercent", ipv4UnicastPrefixLength.Ipv4UnicastPrefixPercent})

    ipv4UnicastPrefixLength.EntityData.YListKeys = []string {"PrefixLength"}

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
    ipv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/tcam-table/fib-table/" + ipv6Address.EntityData.SegmentPath
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = types.NewOrderedMap()
    ipv6Address.EntityData.Children.Append("ipv6-unicast", types.YChild{"Ipv6Unicast", &ipv6Address.Ipv6Unicast})
    ipv6Address.EntityData.Leafs = types.NewOrderedMap()

    ipv6Address.EntityData.YListKeys = []string {}

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
    ipv6Unicast.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/tcam-table/fib-table/ipv6-address/" + ipv6Unicast.EntityData.SegmentPath
    ipv6Unicast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Unicast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Unicast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Unicast.EntityData.Children = types.NewOrderedMap()
    ipv6Unicast.EntityData.Children.Append("ipv6-unicast-prefix-lengths", types.YChild{"Ipv6UnicastPrefixLengths", &ipv6Unicast.Ipv6UnicastPrefixLengths})
    ipv6Unicast.EntityData.Leafs = types.NewOrderedMap()
    ipv6Unicast.EntityData.Leafs.Append("ipv6-unicast-percent", types.YLeaf{"Ipv6UnicastPercent", ipv6Unicast.Ipv6UnicastPercent})

    ipv6Unicast.EntityData.YListKeys = []string {}

    return &(ipv6Unicast.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths
// IPv6 Unicast prefix 
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Unicast prefix length. The type is slice of
    // HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength.
    Ipv6UnicastPrefixLength []*HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength
}

func (ipv6UnicastPrefixLengths *HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths) GetEntityData() *types.CommonEntityData {
    ipv6UnicastPrefixLengths.EntityData.YFilter = ipv6UnicastPrefixLengths.YFilter
    ipv6UnicastPrefixLengths.EntityData.YangName = "ipv6-unicast-prefix-lengths"
    ipv6UnicastPrefixLengths.EntityData.BundleName = "cisco_ios_xr"
    ipv6UnicastPrefixLengths.EntityData.ParentYangName = "ipv6-unicast"
    ipv6UnicastPrefixLengths.EntityData.SegmentPath = "ipv6-unicast-prefix-lengths"
    ipv6UnicastPrefixLengths.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/tcam-table/fib-table/ipv6-address/ipv6-unicast/" + ipv6UnicastPrefixLengths.EntityData.SegmentPath
    ipv6UnicastPrefixLengths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6UnicastPrefixLengths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6UnicastPrefixLengths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6UnicastPrefixLengths.EntityData.Children = types.NewOrderedMap()
    ipv6UnicastPrefixLengths.EntityData.Children.Append("ipv6-unicast-prefix-length", types.YChild{"Ipv6UnicastPrefixLength", nil})
    for i := range ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength {
        ipv6UnicastPrefixLengths.EntityData.Children.Append(types.GetSegmentPath(ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength[i]), types.YChild{"Ipv6UnicastPrefixLength", ipv6UnicastPrefixLengths.Ipv6UnicastPrefixLength[i]})
    }
    ipv6UnicastPrefixLengths.EntityData.Leafs = types.NewOrderedMap()

    ipv6UnicastPrefixLengths.EntityData.YListKeys = []string {}

    return &(ipv6UnicastPrefixLengths.EntityData)
}

// HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength
// IPv6 Unicast prefix length
type HwModuleProfileConfig_Profile_TcamTable_FibTable_Ipv6Address_Ipv6Unicast_Ipv6UnicastPrefixLengths_Ipv6UnicastPrefixLength struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    ipv6UnicastPrefixLength.EntityData.SegmentPath = "ipv6-unicast-prefix-length" + types.AddKeyToken(ipv6UnicastPrefixLength.PrefixLength, "prefix-length")
    ipv6UnicastPrefixLength.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/tcam-table/fib-table/ipv6-address/ipv6-unicast/ipv6-unicast-prefix-lengths/" + ipv6UnicastPrefixLength.EntityData.SegmentPath
    ipv6UnicastPrefixLength.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6UnicastPrefixLength.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6UnicastPrefixLength.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6UnicastPrefixLength.EntityData.Children = types.NewOrderedMap()
    ipv6UnicastPrefixLength.EntityData.Leafs = types.NewOrderedMap()
    ipv6UnicastPrefixLength.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", ipv6UnicastPrefixLength.PrefixLength})
    ipv6UnicastPrefixLength.EntityData.Leafs.Append("ipv6-unicast-prefix-percent", types.YLeaf{"Ipv6UnicastPrefixPercent", ipv6UnicastPrefixLength.Ipv6UnicastPrefixPercent})

    ipv6UnicastPrefixLength.EntityData.YListKeys = []string {"PrefixLength"}

    return &(ipv6UnicastPrefixLength.EntityData)
}

// HwModuleProfileConfig_Profile_Netflow
// Configure Netflow profile.
type HwModuleProfileConfig_Profile_Netflow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none.
    NetflowLocations HwModuleProfileConfig_Profile_Netflow_NetflowLocations

    // IPFIX315 Location All.
    LocationAll HwModuleProfileConfig_Profile_Netflow_LocationAll
}

func (netflow *HwModuleProfileConfig_Profile_Netflow) GetEntityData() *types.CommonEntityData {
    netflow.EntityData.YFilter = netflow.YFilter
    netflow.EntityData.YangName = "netflow"
    netflow.EntityData.BundleName = "cisco_ios_xr"
    netflow.EntityData.ParentYangName = "profile"
    netflow.EntityData.SegmentPath = "netflow"
    netflow.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/" + netflow.EntityData.SegmentPath
    netflow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netflow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netflow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netflow.EntityData.Children = types.NewOrderedMap()
    netflow.EntityData.Children.Append("netflow-locations", types.YChild{"NetflowLocations", &netflow.NetflowLocations})
    netflow.EntityData.Children.Append("location-all", types.YChild{"LocationAll", &netflow.LocationAll})
    netflow.EntityData.Leafs = types.NewOrderedMap()

    netflow.EntityData.YListKeys = []string {}

    return &(netflow.EntityData)
}

// HwModuleProfileConfig_Profile_Netflow_NetflowLocations
// none
type HwModuleProfileConfig_Profile_Netflow_NetflowLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of
    // HwModuleProfileConfig_Profile_Netflow_NetflowLocations_NetflowLocation.
    NetflowLocation []*HwModuleProfileConfig_Profile_Netflow_NetflowLocations_NetflowLocation
}

func (netflowLocations *HwModuleProfileConfig_Profile_Netflow_NetflowLocations) GetEntityData() *types.CommonEntityData {
    netflowLocations.EntityData.YFilter = netflowLocations.YFilter
    netflowLocations.EntityData.YangName = "netflow-locations"
    netflowLocations.EntityData.BundleName = "cisco_ios_xr"
    netflowLocations.EntityData.ParentYangName = "netflow"
    netflowLocations.EntityData.SegmentPath = "netflow-locations"
    netflowLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/netflow/" + netflowLocations.EntityData.SegmentPath
    netflowLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netflowLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netflowLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netflowLocations.EntityData.Children = types.NewOrderedMap()
    netflowLocations.EntityData.Children.Append("netflow-location", types.YChild{"NetflowLocation", nil})
    for i := range netflowLocations.NetflowLocation {
        netflowLocations.EntityData.Children.Append(types.GetSegmentPath(netflowLocations.NetflowLocation[i]), types.YChild{"NetflowLocation", netflowLocations.NetflowLocation[i]})
    }
    netflowLocations.EntityData.Leafs = types.NewOrderedMap()

    netflowLocations.EntityData.YListKeys = []string {}

    return &(netflowLocations.EntityData)
}

// HwModuleProfileConfig_Profile_Netflow_NetflowLocations_NetflowLocation
// none
type HwModuleProfileConfig_Profile_Netflow_NetflowLocations_NetflowLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Location of NETFLOW config. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    LocationString interface{}

    // none. The type is slice of
    // HwModuleProfileConfig_Profile_Netflow_NetflowLocations_NetflowLocation_NetflowLocationLeaf.
    NetflowLocationLeaf []*HwModuleProfileConfig_Profile_Netflow_NetflowLocations_NetflowLocation_NetflowLocationLeaf
}

func (netflowLocation *HwModuleProfileConfig_Profile_Netflow_NetflowLocations_NetflowLocation) GetEntityData() *types.CommonEntityData {
    netflowLocation.EntityData.YFilter = netflowLocation.YFilter
    netflowLocation.EntityData.YangName = "netflow-location"
    netflowLocation.EntityData.BundleName = "cisco_ios_xr"
    netflowLocation.EntityData.ParentYangName = "netflow-locations"
    netflowLocation.EntityData.SegmentPath = "netflow-location" + types.AddKeyToken(netflowLocation.LocationString, "location-string")
    netflowLocation.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/netflow/netflow-locations/" + netflowLocation.EntityData.SegmentPath
    netflowLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netflowLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netflowLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netflowLocation.EntityData.Children = types.NewOrderedMap()
    netflowLocation.EntityData.Children.Append("netflow-location-leaf", types.YChild{"NetflowLocationLeaf", nil})
    for i := range netflowLocation.NetflowLocationLeaf {
        netflowLocation.EntityData.Children.Append(types.GetSegmentPath(netflowLocation.NetflowLocationLeaf[i]), types.YChild{"NetflowLocationLeaf", netflowLocation.NetflowLocationLeaf[i]})
    }
    netflowLocation.EntityData.Leafs = types.NewOrderedMap()
    netflowLocation.EntityData.Leafs.Append("location-string", types.YLeaf{"LocationString", netflowLocation.LocationString})

    netflowLocation.EntityData.YListKeys = []string {"LocationString"}

    return &(netflowLocation.EntityData)
}

// HwModuleProfileConfig_Profile_Netflow_NetflowLocations_NetflowLocation_NetflowLocationLeaf
// none
type HwModuleProfileConfig_Profile_Netflow_NetflowLocations_NetflowLocation_NetflowLocationLeaf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Location ID. The type is interface{} with range:
    // 0..4294967295.
    LocationId interface{}

    // If Enabled set value to 65535. The type is interface{} with range:
    // 0..4294967295. This attribute is mandatory.
    EnableVal interface{}
}

func (netflowLocationLeaf *HwModuleProfileConfig_Profile_Netflow_NetflowLocations_NetflowLocation_NetflowLocationLeaf) GetEntityData() *types.CommonEntityData {
    netflowLocationLeaf.EntityData.YFilter = netflowLocationLeaf.YFilter
    netflowLocationLeaf.EntityData.YangName = "netflow-location-leaf"
    netflowLocationLeaf.EntityData.BundleName = "cisco_ios_xr"
    netflowLocationLeaf.EntityData.ParentYangName = "netflow-location"
    netflowLocationLeaf.EntityData.SegmentPath = "netflow-location-leaf" + types.AddKeyToken(netflowLocationLeaf.LocationId, "location-id")
    netflowLocationLeaf.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/netflow/netflow-locations/netflow-location/" + netflowLocationLeaf.EntityData.SegmentPath
    netflowLocationLeaf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netflowLocationLeaf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netflowLocationLeaf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netflowLocationLeaf.EntityData.Children = types.NewOrderedMap()
    netflowLocationLeaf.EntityData.Leafs = types.NewOrderedMap()
    netflowLocationLeaf.EntityData.Leafs.Append("location-id", types.YLeaf{"LocationId", netflowLocationLeaf.LocationId})
    netflowLocationLeaf.EntityData.Leafs.Append("enable-val", types.YLeaf{"EnableVal", netflowLocationLeaf.EnableVal})

    netflowLocationLeaf.EntityData.YListKeys = []string {"LocationId"}

    return &(netflowLocationLeaf.EntityData)
}

// HwModuleProfileConfig_Profile_Netflow_LocationAll
// IPFIX315 Location All
type HwModuleProfileConfig_Profile_Netflow_LocationAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If Enabled set value to 65535. The type is interface{} with range:
    // 0..4294967295.
    LocationAllLeaf interface{}
}

func (locationAll *HwModuleProfileConfig_Profile_Netflow_LocationAll) GetEntityData() *types.CommonEntityData {
    locationAll.EntityData.YFilter = locationAll.YFilter
    locationAll.EntityData.YangName = "location-all"
    locationAll.EntityData.BundleName = "cisco_ios_xr"
    locationAll.EntityData.ParentYangName = "netflow"
    locationAll.EntityData.SegmentPath = "location-all"
    locationAll.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/netflow/" + locationAll.EntityData.SegmentPath
    locationAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locationAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locationAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locationAll.EntityData.Children = types.NewOrderedMap()
    locationAll.EntityData.Leafs = types.NewOrderedMap()
    locationAll.EntityData.Leafs.Append("location-all-leaf", types.YLeaf{"LocationAllLeaf", locationAll.LocationAllLeaf})

    locationAll.EntityData.YListKeys = []string {}

    return &(locationAll.EntityData)
}

// HwModuleProfileConfig_Profile_Flowspecs
// Configure Flowspec profile.
type HwModuleProfileConfig_Profile_Flowspecs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of
    // HwModuleProfileConfig_Profile_Flowspecs_Flowspec.
    Flowspec []*HwModuleProfileConfig_Profile_Flowspecs_Flowspec
}

func (flowspecs *HwModuleProfileConfig_Profile_Flowspecs) GetEntityData() *types.CommonEntityData {
    flowspecs.EntityData.YFilter = flowspecs.YFilter
    flowspecs.EntityData.YangName = "flowspecs"
    flowspecs.EntityData.BundleName = "cisco_ios_xr"
    flowspecs.EntityData.ParentYangName = "profile"
    flowspecs.EntityData.SegmentPath = "flowspecs"
    flowspecs.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/" + flowspecs.EntityData.SegmentPath
    flowspecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowspecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowspecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowspecs.EntityData.Children = types.NewOrderedMap()
    flowspecs.EntityData.Children.Append("flowspec", types.YChild{"Flowspec", nil})
    for i := range flowspecs.Flowspec {
        flowspecs.EntityData.Children.Append(types.GetSegmentPath(flowspecs.Flowspec[i]), types.YChild{"Flowspec", flowspecs.Flowspec[i]})
    }
    flowspecs.EntityData.Leafs = types.NewOrderedMap()

    flowspecs.EntityData.YListKeys = []string {}

    return &(flowspecs.EntityData)
}

// HwModuleProfileConfig_Profile_Flowspecs_Flowspec
// none
type HwModuleProfileConfig_Profile_Flowspecs_Flowspec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. none. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    V6Enable interface{}

    // This attribute is a key. Location of FLOWSPEC config. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    LocationString interface{}

    // This attribute is a key. Location ID hex to Decimal 0xffff for all. The
    // type is interface{} with range: 0..4294967295.
    LocationId interface{}

    // If Enabled set value to 65535. The type is interface{} with range:
    // 0..4294967295. This attribute is mandatory.
    EnableVal interface{}
}

func (flowspec *HwModuleProfileConfig_Profile_Flowspecs_Flowspec) GetEntityData() *types.CommonEntityData {
    flowspec.EntityData.YFilter = flowspec.YFilter
    flowspec.EntityData.YangName = "flowspec"
    flowspec.EntityData.BundleName = "cisco_ios_xr"
    flowspec.EntityData.ParentYangName = "flowspecs"
    flowspec.EntityData.SegmentPath = "flowspec" + types.AddKeyToken(flowspec.V6Enable, "v6-enable") + types.AddKeyToken(flowspec.LocationString, "location-string") + types.AddKeyToken(flowspec.LocationId, "location-id")
    flowspec.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/flowspecs/" + flowspec.EntityData.SegmentPath
    flowspec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowspec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowspec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowspec.EntityData.Children = types.NewOrderedMap()
    flowspec.EntityData.Leafs = types.NewOrderedMap()
    flowspec.EntityData.Leafs.Append("v6-enable", types.YLeaf{"V6Enable", flowspec.V6Enable})
    flowspec.EntityData.Leafs.Append("location-string", types.YLeaf{"LocationString", flowspec.LocationString})
    flowspec.EntityData.Leafs.Append("location-id", types.YLeaf{"LocationId", flowspec.LocationId})
    flowspec.EntityData.Leafs.Append("enable-val", types.YLeaf{"EnableVal", flowspec.EnableVal})

    flowspec.EntityData.YListKeys = []string {"V6Enable", "LocationString", "LocationId"}

    return &(flowspec.EntityData)
}

// HwModuleProfileConfig_Profile_SegmentRoutings
// Configure Segment Routing profile.
type HwModuleProfileConfig_Profile_SegmentRoutings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of
    // HwModuleProfileConfig_Profile_SegmentRoutings_SegmentRouting.
    SegmentRouting []*HwModuleProfileConfig_Profile_SegmentRoutings_SegmentRouting
}

func (segmentRoutings *HwModuleProfileConfig_Profile_SegmentRoutings) GetEntityData() *types.CommonEntityData {
    segmentRoutings.EntityData.YFilter = segmentRoutings.YFilter
    segmentRoutings.EntityData.YangName = "segment-routings"
    segmentRoutings.EntityData.BundleName = "cisco_ios_xr"
    segmentRoutings.EntityData.ParentYangName = "profile"
    segmentRoutings.EntityData.SegmentPath = "segment-routings"
    segmentRoutings.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/" + segmentRoutings.EntityData.SegmentPath
    segmentRoutings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segmentRoutings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segmentRoutings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segmentRoutings.EntityData.Children = types.NewOrderedMap()
    segmentRoutings.EntityData.Children.Append("segment-routing", types.YChild{"SegmentRouting", nil})
    for i := range segmentRoutings.SegmentRouting {
        segmentRoutings.EntityData.Children.Append(types.GetSegmentPath(segmentRoutings.SegmentRouting[i]), types.YChild{"SegmentRouting", segmentRoutings.SegmentRouting[i]})
    }
    segmentRoutings.EntityData.Leafs = types.NewOrderedMap()

    segmentRoutings.EntityData.YListKeys = []string {}

    return &(segmentRoutings.EntityData)
}

// HwModuleProfileConfig_Profile_SegmentRoutings_SegmentRouting
// none
type HwModuleProfileConfig_Profile_SegmentRoutings_SegmentRouting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. none. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Srv6 interface{}

    // If Enabled set value to 1. The type is interface{} with range:
    // 0..4294967295. This attribute is mandatory.
    EnableVal interface{}
}

func (segmentRouting *HwModuleProfileConfig_Profile_SegmentRoutings_SegmentRouting) GetEntityData() *types.CommonEntityData {
    segmentRouting.EntityData.YFilter = segmentRouting.YFilter
    segmentRouting.EntityData.YangName = "segment-routing"
    segmentRouting.EntityData.BundleName = "cisco_ios_xr"
    segmentRouting.EntityData.ParentYangName = "segment-routings"
    segmentRouting.EntityData.SegmentPath = "segment-routing" + types.AddKeyToken(segmentRouting.Srv6, "srv6")
    segmentRouting.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/segment-routings/" + segmentRouting.EntityData.SegmentPath
    segmentRouting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segmentRouting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segmentRouting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segmentRouting.EntityData.Children = types.NewOrderedMap()
    segmentRouting.EntityData.Leafs = types.NewOrderedMap()
    segmentRouting.EntityData.Leafs.Append("srv6", types.YLeaf{"Srv6", segmentRouting.Srv6})
    segmentRouting.EntityData.Leafs.Append("enable-val", types.YLeaf{"EnableVal", segmentRouting.EnableVal})

    segmentRouting.EntityData.YListKeys = []string {"Srv6"}

    return &(segmentRouting.EntityData)
}

// HwModuleProfileConfig_Profile_LoadBalance
// Configure load balance parameters
type HwModuleProfileConfig_Profile_LoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure load balance parameters. The type is interface{} with range:
    // 0..4294967295.
    LoadBalanceProfile interface{}
}

func (loadBalance *HwModuleProfileConfig_Profile_LoadBalance) GetEntityData() *types.CommonEntityData {
    loadBalance.EntityData.YFilter = loadBalance.YFilter
    loadBalance.EntityData.YangName = "load-balance"
    loadBalance.EntityData.BundleName = "cisco_ios_xr"
    loadBalance.EntityData.ParentYangName = "profile"
    loadBalance.EntityData.SegmentPath = "load-balance"
    loadBalance.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/" + loadBalance.EntityData.SegmentPath
    loadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadBalance.EntityData.Children = types.NewOrderedMap()
    loadBalance.EntityData.Leafs = types.NewOrderedMap()
    loadBalance.EntityData.Leafs.Append("load-balance-profile", types.YLeaf{"LoadBalanceProfile", loadBalance.LoadBalanceProfile})

    loadBalance.EntityData.YListKeys = []string {}

    return &(loadBalance.EntityData)
}

// HwModuleProfileConfig_Profile_Stats
// Configure stats
type HwModuleProfileConfig_Profile_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure stats for qos-enhanced and acl-permit.
    StatsProfileModes HwModuleProfileConfig_Profile_Stats_StatsProfileModes
}

func (stats *HwModuleProfileConfig_Profile_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "profile"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/" + stats.EntityData.SegmentPath
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Children.Append("stats-profile-modes", types.YChild{"StatsProfileModes", &stats.StatsProfileModes})
    stats.EntityData.Leafs = types.NewOrderedMap()

    stats.EntityData.YListKeys = []string {}

    return &(stats.EntityData)
}

// HwModuleProfileConfig_Profile_Stats_StatsProfileModes
// Configure stats for qos-enhanced and
// acl-permit
type HwModuleProfileConfig_Profile_Stats_StatsProfileModes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure stats for qos-enhanced and acl-permit. The type is slice of
    // HwModuleProfileConfig_Profile_Stats_StatsProfileModes_StatsProfileMode.
    StatsProfileMode []*HwModuleProfileConfig_Profile_Stats_StatsProfileModes_StatsProfileMode
}

func (statsProfileModes *HwModuleProfileConfig_Profile_Stats_StatsProfileModes) GetEntityData() *types.CommonEntityData {
    statsProfileModes.EntityData.YFilter = statsProfileModes.YFilter
    statsProfileModes.EntityData.YangName = "stats-profile-modes"
    statsProfileModes.EntityData.BundleName = "cisco_ios_xr"
    statsProfileModes.EntityData.ParentYangName = "stats"
    statsProfileModes.EntityData.SegmentPath = "stats-profile-modes"
    statsProfileModes.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/stats/" + statsProfileModes.EntityData.SegmentPath
    statsProfileModes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statsProfileModes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statsProfileModes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statsProfileModes.EntityData.Children = types.NewOrderedMap()
    statsProfileModes.EntityData.Children.Append("stats-profile-mode", types.YChild{"StatsProfileMode", nil})
    for i := range statsProfileModes.StatsProfileMode {
        statsProfileModes.EntityData.Children.Append(types.GetSegmentPath(statsProfileModes.StatsProfileMode[i]), types.YChild{"StatsProfileMode", statsProfileModes.StatsProfileMode[i]})
    }
    statsProfileModes.EntityData.Leafs = types.NewOrderedMap()

    statsProfileModes.EntityData.YListKeys = []string {}

    return &(statsProfileModes.EntityData)
}

// HwModuleProfileConfig_Profile_Stats_StatsProfileModes_StatsProfileMode
// Configure stats for qos-enhanced and
// acl-permit
type HwModuleProfileConfig_Profile_Stats_StatsProfileModes_StatsProfileMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Give Default Stats Mode. The type is interface{}
    // with range: 0..4294967295.
    StatsModeDefault interface{}

    // Stats Mode. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    Mode interface{}
}

func (statsProfileMode *HwModuleProfileConfig_Profile_Stats_StatsProfileModes_StatsProfileMode) GetEntityData() *types.CommonEntityData {
    statsProfileMode.EntityData.YFilter = statsProfileMode.YFilter
    statsProfileMode.EntityData.YangName = "stats-profile-mode"
    statsProfileMode.EntityData.BundleName = "cisco_ios_xr"
    statsProfileMode.EntityData.ParentYangName = "stats-profile-modes"
    statsProfileMode.EntityData.SegmentPath = "stats-profile-mode" + types.AddKeyToken(statsProfileMode.StatsModeDefault, "stats-mode-default")
    statsProfileMode.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/stats/stats-profile-modes/" + statsProfileMode.EntityData.SegmentPath
    statsProfileMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statsProfileMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statsProfileMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statsProfileMode.EntityData.Children = types.NewOrderedMap()
    statsProfileMode.EntityData.Leafs = types.NewOrderedMap()
    statsProfileMode.EntityData.Leafs.Append("stats-mode-default", types.YLeaf{"StatsModeDefault", statsProfileMode.StatsModeDefault})
    statsProfileMode.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", statsProfileMode.Mode})

    statsProfileMode.EntityData.YListKeys = []string {"StatsModeDefault"}

    return &(statsProfileMode.EntityData)
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
    profileAcl.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/" + profileAcl.EntityData.SegmentPath
    profileAcl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profileAcl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profileAcl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profileAcl.EntityData.Children = types.NewOrderedMap()
    profileAcl.EntityData.Leafs = types.NewOrderedMap()
    profileAcl.EntityData.Leafs.Append("egress", types.YLeaf{"Egress", profileAcl.Egress})

    profileAcl.EntityData.YListKeys = []string {}

    return &(profileAcl.EntityData)
}

// HwModuleProfileConfig_Profile_BundleScale
// Configure Bundle profile.
type HwModuleProfileConfig_Profile_BundleScale struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Max Trunk Size. The type is interface{} with range:
    // 0..4294967295.
    TrunkSize interface{}
}

func (bundleScale *HwModuleProfileConfig_Profile_BundleScale) GetEntityData() *types.CommonEntityData {
    bundleScale.EntityData.YFilter = bundleScale.YFilter
    bundleScale.EntityData.YangName = "bundle-scale"
    bundleScale.EntityData.BundleName = "cisco_ios_xr"
    bundleScale.EntityData.ParentYangName = "profile"
    bundleScale.EntityData.SegmentPath = "bundle-scale"
    bundleScale.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/" + bundleScale.EntityData.SegmentPath
    bundleScale.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleScale.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleScale.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleScale.EntityData.Children = types.NewOrderedMap()
    bundleScale.EntityData.Leafs = types.NewOrderedMap()
    bundleScale.EntityData.Leafs.Append("trunk-size", types.YLeaf{"TrunkSize", bundleScale.TrunkSize})

    bundleScale.EntityData.YListKeys = []string {}

    return &(bundleScale.EntityData)
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
    profileTcam.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/" + profileTcam.EntityData.SegmentPath
    profileTcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profileTcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profileTcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profileTcam.EntityData.Children = types.NewOrderedMap()
    profileTcam.EntityData.Children.Append("key-format", types.YChild{"KeyFormat", &profileTcam.KeyFormat})
    profileTcam.EntityData.Leafs = types.NewOrderedMap()

    profileTcam.EntityData.YListKeys = []string {}

    return &(profileTcam.EntityData)
}

// HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat
// none
type HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure acl profile.
    KeyFormatAclTable HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable
}

func (keyFormat *HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat) GetEntityData() *types.CommonEntityData {
    keyFormat.EntityData.YFilter = keyFormat.YFilter
    keyFormat.EntityData.YangName = "key-format"
    keyFormat.EntityData.BundleName = "cisco_ios_xr"
    keyFormat.EntityData.ParentYangName = "profile-tcam"
    keyFormat.EntityData.SegmentPath = "key-format"
    keyFormat.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/profile-tcam/" + keyFormat.EntityData.SegmentPath
    keyFormat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyFormat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyFormat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyFormat.EntityData.Children = types.NewOrderedMap()
    keyFormat.EntityData.Children.Append("key-format-acl-table", types.YChild{"KeyFormatAclTable", &keyFormat.KeyFormatAclTable})
    keyFormat.EntityData.Leafs = types.NewOrderedMap()

    keyFormat.EntityData.YListKeys = []string {}

    return &(keyFormat.EntityData)
}

// HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable
// Configure acl profile
type HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure ipv6 acl profile.
    Ipv6AclTables HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv6AclTables

    // Configure ipv4 acl profile.
    Ipv4AclTables HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv4AclTables
}

func (keyFormatAclTable *HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable) GetEntityData() *types.CommonEntityData {
    keyFormatAclTable.EntityData.YFilter = keyFormatAclTable.YFilter
    keyFormatAclTable.EntityData.YangName = "key-format-acl-table"
    keyFormatAclTable.EntityData.BundleName = "cisco_ios_xr"
    keyFormatAclTable.EntityData.ParentYangName = "key-format"
    keyFormatAclTable.EntityData.SegmentPath = "key-format-acl-table"
    keyFormatAclTable.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/profile-tcam/key-format/" + keyFormatAclTable.EntityData.SegmentPath
    keyFormatAclTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyFormatAclTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyFormatAclTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyFormatAclTable.EntityData.Children = types.NewOrderedMap()
    keyFormatAclTable.EntityData.Children.Append("ipv6-acl-tables", types.YChild{"Ipv6AclTables", &keyFormatAclTable.Ipv6AclTables})
    keyFormatAclTable.EntityData.Children.Append("ipv4-acl-tables", types.YChild{"Ipv4AclTables", &keyFormatAclTable.Ipv4AclTables})
    keyFormatAclTable.EntityData.Leafs = types.NewOrderedMap()

    keyFormatAclTable.EntityData.YListKeys = []string {}

    return &(keyFormatAclTable.EntityData)
}

// HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv6AclTables
// Configure ipv6 acl profile
type HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv6AclTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure format for ipv6 acl profile. The type is slice of
    // HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv6AclTables_Ipv6AclTable.
    Ipv6AclTable []*HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv6AclTables_Ipv6AclTable
}

func (ipv6AclTables *HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv6AclTables) GetEntityData() *types.CommonEntityData {
    ipv6AclTables.EntityData.YFilter = ipv6AclTables.YFilter
    ipv6AclTables.EntityData.YangName = "ipv6-acl-tables"
    ipv6AclTables.EntityData.BundleName = "cisco_ios_xr"
    ipv6AclTables.EntityData.ParentYangName = "key-format-acl-table"
    ipv6AclTables.EntityData.SegmentPath = "ipv6-acl-tables"
    ipv6AclTables.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/profile-tcam/key-format/key-format-acl-table/" + ipv6AclTables.EntityData.SegmentPath
    ipv6AclTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6AclTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6AclTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6AclTables.EntityData.Children = types.NewOrderedMap()
    ipv6AclTables.EntityData.Children.Append("ipv6-acl-table", types.YChild{"Ipv6AclTable", nil})
    for i := range ipv6AclTables.Ipv6AclTable {
        ipv6AclTables.EntityData.Children.Append(types.GetSegmentPath(ipv6AclTables.Ipv6AclTable[i]), types.YChild{"Ipv6AclTable", ipv6AclTables.Ipv6AclTable[i]})
    }
    ipv6AclTables.EntityData.Leafs = types.NewOrderedMap()

    ipv6AclTables.EntityData.YListKeys = []string {}

    return &(ipv6AclTables.EntityData)
}

// HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv6AclTables_Ipv6AclTable
// Configure format for ipv6 acl profile
type HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv6AclTables_Ipv6AclTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Location string (all) if for all LCs. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    LocationString interface{}

    // This attribute is a key. Location ID hex to Decimal 0xffff for all. The
    // type is interface{} with range: 0..4294967295.
    LocationId interface{}

    // Source Address. The type is interface{} with range: 0..4294967295.
    SourceAddr interface{}

    // Source Port. The type is interface{} with range: 0..4294967295.
    SourcePort interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddr interface{}

    // Destination Port. The type is interface{} with range: 0..4294967295.
    DestPort interface{}

    // Next Header Type. The type is interface{} with range: 0..4294967295.
    ProtType interface{}

    // TCP Flags. The type is interface{} with range: 0..4294967295.
    TcpFlag interface{}

    // Packet Length. The type is interface{} with range: 0..4294967295.
    PackLen interface{}

    // Precedence. The type is interface{} with range: 0..4294967295.
    Precedence interface{}

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

    // Enable Capture. The type is interface{} with range: 0..4294967295.
    EnCapture interface{}

    // Enable Setting TTL. The type is interface{} with range: 0..4294967295.
    EnTtl interface{}

    // Enable Matching TTL. The type is interface{} with range: 0..4294967295.
    EnMatch interface{}

    // Enable Non Shared Interface ACL. The type is interface{} with range:
    // 0..4294967295.
    EnShareAcl interface{}
}

func (ipv6AclTable *HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv6AclTables_Ipv6AclTable) GetEntityData() *types.CommonEntityData {
    ipv6AclTable.EntityData.YFilter = ipv6AclTable.YFilter
    ipv6AclTable.EntityData.YangName = "ipv6-acl-table"
    ipv6AclTable.EntityData.BundleName = "cisco_ios_xr"
    ipv6AclTable.EntityData.ParentYangName = "ipv6-acl-tables"
    ipv6AclTable.EntityData.SegmentPath = "ipv6-acl-table" + types.AddKeyToken(ipv6AclTable.LocationString, "location-string") + types.AddKeyToken(ipv6AclTable.LocationId, "location-id")
    ipv6AclTable.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/profile-tcam/key-format/key-format-acl-table/ipv6-acl-tables/" + ipv6AclTable.EntityData.SegmentPath
    ipv6AclTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6AclTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6AclTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6AclTable.EntityData.Children = types.NewOrderedMap()
    ipv6AclTable.EntityData.Leafs = types.NewOrderedMap()
    ipv6AclTable.EntityData.Leafs.Append("location-string", types.YLeaf{"LocationString", ipv6AclTable.LocationString})
    ipv6AclTable.EntityData.Leafs.Append("location-id", types.YLeaf{"LocationId", ipv6AclTable.LocationId})
    ipv6AclTable.EntityData.Leafs.Append("source-addr", types.YLeaf{"SourceAddr", ipv6AclTable.SourceAddr})
    ipv6AclTable.EntityData.Leafs.Append("source-port", types.YLeaf{"SourcePort", ipv6AclTable.SourcePort})
    ipv6AclTable.EntityData.Leafs.Append("destination-addr", types.YLeaf{"DestinationAddr", ipv6AclTable.DestinationAddr})
    ipv6AclTable.EntityData.Leafs.Append("dest-port", types.YLeaf{"DestPort", ipv6AclTable.DestPort})
    ipv6AclTable.EntityData.Leafs.Append("prot-type", types.YLeaf{"ProtType", ipv6AclTable.ProtType})
    ipv6AclTable.EntityData.Leafs.Append("tcp-flag", types.YLeaf{"TcpFlag", ipv6AclTable.TcpFlag})
    ipv6AclTable.EntityData.Leafs.Append("pack-len", types.YLeaf{"PackLen", ipv6AclTable.PackLen})
    ipv6AclTable.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", ipv6AclTable.Precedence})
    ipv6AclTable.EntityData.Leafs.Append("udf1", types.YLeaf{"Udf1", ipv6AclTable.Udf1})
    ipv6AclTable.EntityData.Leafs.Append("udf2", types.YLeaf{"Udf2", ipv6AclTable.Udf2})
    ipv6AclTable.EntityData.Leafs.Append("udf3", types.YLeaf{"Udf3", ipv6AclTable.Udf3})
    ipv6AclTable.EntityData.Leafs.Append("udf4", types.YLeaf{"Udf4", ipv6AclTable.Udf4})
    ipv6AclTable.EntityData.Leafs.Append("udf5", types.YLeaf{"Udf5", ipv6AclTable.Udf5})
    ipv6AclTable.EntityData.Leafs.Append("udf6", types.YLeaf{"Udf6", ipv6AclTable.Udf6})
    ipv6AclTable.EntityData.Leafs.Append("udf7", types.YLeaf{"Udf7", ipv6AclTable.Udf7})
    ipv6AclTable.EntityData.Leafs.Append("udf8", types.YLeaf{"Udf8", ipv6AclTable.Udf8})
    ipv6AclTable.EntityData.Leafs.Append("en-capture", types.YLeaf{"EnCapture", ipv6AclTable.EnCapture})
    ipv6AclTable.EntityData.Leafs.Append("en-ttl", types.YLeaf{"EnTtl", ipv6AclTable.EnTtl})
    ipv6AclTable.EntityData.Leafs.Append("en-match", types.YLeaf{"EnMatch", ipv6AclTable.EnMatch})
    ipv6AclTable.EntityData.Leafs.Append("en-share-acl", types.YLeaf{"EnShareAcl", ipv6AclTable.EnShareAcl})

    ipv6AclTable.EntityData.YListKeys = []string {"LocationString", "LocationId"}

    return &(ipv6AclTable.EntityData)
}

// HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv4AclTables
// Configure ipv4 acl profile
type HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv4AclTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure format for ipv4 acl profile. The type is slice of
    // HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv4AclTables_Ipv4AclTable.
    Ipv4AclTable []*HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv4AclTables_Ipv4AclTable
}

func (ipv4AclTables *HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv4AclTables) GetEntityData() *types.CommonEntityData {
    ipv4AclTables.EntityData.YFilter = ipv4AclTables.YFilter
    ipv4AclTables.EntityData.YangName = "ipv4-acl-tables"
    ipv4AclTables.EntityData.BundleName = "cisco_ios_xr"
    ipv4AclTables.EntityData.ParentYangName = "key-format-acl-table"
    ipv4AclTables.EntityData.SegmentPath = "ipv4-acl-tables"
    ipv4AclTables.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/profile-tcam/key-format/key-format-acl-table/" + ipv4AclTables.EntityData.SegmentPath
    ipv4AclTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AclTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AclTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AclTables.EntityData.Children = types.NewOrderedMap()
    ipv4AclTables.EntityData.Children.Append("ipv4-acl-table", types.YChild{"Ipv4AclTable", nil})
    for i := range ipv4AclTables.Ipv4AclTable {
        ipv4AclTables.EntityData.Children.Append(types.GetSegmentPath(ipv4AclTables.Ipv4AclTable[i]), types.YChild{"Ipv4AclTable", ipv4AclTables.Ipv4AclTable[i]})
    }
    ipv4AclTables.EntityData.Leafs = types.NewOrderedMap()

    ipv4AclTables.EntityData.YListKeys = []string {}

    return &(ipv4AclTables.EntityData)
}

// HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv4AclTables_Ipv4AclTable
// Configure format for ipv4 acl profile
type HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv4AclTables_Ipv4AclTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Location string (all) if for all LCs. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    LocationString interface{}

    // This attribute is a key. Location ID hex to Decimal 0xffff for all. The
    // type is interface{} with range: 0..4294967295.
    LocationId interface{}

    // Source Address. The type is interface{} with range: 0..4294967295.
    SourceAddr interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddr interface{}

    // Source Port. The type is interface{} with range: 0..4294967295.
    SourcePort interface{}

    // Destination Port. The type is interface{} with range: 0..4294967295.
    DestPort interface{}

    // Protocol Type. The type is interface{} with range: 0..4294967295.
    ProtType interface{}

    // TCP Flags. The type is interface{} with range: 0..4294967295.
    TcpFlag interface{}

    // Packet Length. The type is interface{} with range: 0..4294967295.
    PackLen interface{}

    // Fragment Bit. The type is interface{} with range: 0..4294967295.
    FragBit interface{}

    // Precedence. The type is interface{} with range: 0..4294967295.
    Precedence interface{}

    // PortRange. The type is interface{} with range: 0..4294967295.
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

    // Enable Capture. The type is interface{} with range: 0..4294967295.
    EnCapture interface{}

    // Enable Setting TTL. The type is interface{} with range: 0..4294967295.
    EnTtl interface{}

    // Enable Matching TTL. The type is interface{} with range: 0..4294967295.
    EnMatch interface{}

    // Enable Non Shared Interface ACL. The type is interface{} with range:
    // 0..4294967295.
    EnShareAcl interface{}
}

func (ipv4AclTable *HwModuleProfileConfig_Profile_ProfileTcam_KeyFormat_KeyFormatAclTable_Ipv4AclTables_Ipv4AclTable) GetEntityData() *types.CommonEntityData {
    ipv4AclTable.EntityData.YFilter = ipv4AclTable.YFilter
    ipv4AclTable.EntityData.YangName = "ipv4-acl-table"
    ipv4AclTable.EntityData.BundleName = "cisco_ios_xr"
    ipv4AclTable.EntityData.ParentYangName = "ipv4-acl-tables"
    ipv4AclTable.EntityData.SegmentPath = "ipv4-acl-table" + types.AddKeyToken(ipv4AclTable.LocationString, "location-string") + types.AddKeyToken(ipv4AclTable.LocationId, "location-id")
    ipv4AclTable.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/profile-tcam/key-format/key-format-acl-table/ipv4-acl-tables/" + ipv4AclTable.EntityData.SegmentPath
    ipv4AclTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AclTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AclTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AclTable.EntityData.Children = types.NewOrderedMap()
    ipv4AclTable.EntityData.Leafs = types.NewOrderedMap()
    ipv4AclTable.EntityData.Leafs.Append("location-string", types.YLeaf{"LocationString", ipv4AclTable.LocationString})
    ipv4AclTable.EntityData.Leafs.Append("location-id", types.YLeaf{"LocationId", ipv4AclTable.LocationId})
    ipv4AclTable.EntityData.Leafs.Append("source-addr", types.YLeaf{"SourceAddr", ipv4AclTable.SourceAddr})
    ipv4AclTable.EntityData.Leafs.Append("destination-addr", types.YLeaf{"DestinationAddr", ipv4AclTable.DestinationAddr})
    ipv4AclTable.EntityData.Leafs.Append("source-port", types.YLeaf{"SourcePort", ipv4AclTable.SourcePort})
    ipv4AclTable.EntityData.Leafs.Append("dest-port", types.YLeaf{"DestPort", ipv4AclTable.DestPort})
    ipv4AclTable.EntityData.Leafs.Append("prot-type", types.YLeaf{"ProtType", ipv4AclTable.ProtType})
    ipv4AclTable.EntityData.Leafs.Append("tcp-flag", types.YLeaf{"TcpFlag", ipv4AclTable.TcpFlag})
    ipv4AclTable.EntityData.Leafs.Append("pack-len", types.YLeaf{"PackLen", ipv4AclTable.PackLen})
    ipv4AclTable.EntityData.Leafs.Append("frag-bit", types.YLeaf{"FragBit", ipv4AclTable.FragBit})
    ipv4AclTable.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", ipv4AclTable.Precedence})
    ipv4AclTable.EntityData.Leafs.Append("port-range", types.YLeaf{"PortRange", ipv4AclTable.PortRange})
    ipv4AclTable.EntityData.Leafs.Append("udf1", types.YLeaf{"Udf1", ipv4AclTable.Udf1})
    ipv4AclTable.EntityData.Leafs.Append("udf2", types.YLeaf{"Udf2", ipv4AclTable.Udf2})
    ipv4AclTable.EntityData.Leafs.Append("udf3", types.YLeaf{"Udf3", ipv4AclTable.Udf3})
    ipv4AclTable.EntityData.Leafs.Append("udf4", types.YLeaf{"Udf4", ipv4AclTable.Udf4})
    ipv4AclTable.EntityData.Leafs.Append("udf5", types.YLeaf{"Udf5", ipv4AclTable.Udf5})
    ipv4AclTable.EntityData.Leafs.Append("udf6", types.YLeaf{"Udf6", ipv4AclTable.Udf6})
    ipv4AclTable.EntityData.Leafs.Append("udf7", types.YLeaf{"Udf7", ipv4AclTable.Udf7})
    ipv4AclTable.EntityData.Leafs.Append("udf8", types.YLeaf{"Udf8", ipv4AclTable.Udf8})
    ipv4AclTable.EntityData.Leafs.Append("en-capture", types.YLeaf{"EnCapture", ipv4AclTable.EnCapture})
    ipv4AclTable.EntityData.Leafs.Append("en-ttl", types.YLeaf{"EnTtl", ipv4AclTable.EnTtl})
    ipv4AclTable.EntityData.Leafs.Append("en-match", types.YLeaf{"EnMatch", ipv4AclTable.EnMatch})
    ipv4AclTable.EntityData.Leafs.Append("en-share-acl", types.YLeaf{"EnShareAcl", ipv4AclTable.EnShareAcl})

    ipv4AclTable.EntityData.YListKeys = []string {"LocationString", "LocationId"}

    return &(ipv4AclTable.EntityData)
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
    qos.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/" + qos.EntityData.SegmentPath
    qos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qos.EntityData.Children = types.NewOrderedMap()
    qos.EntityData.Children.Append("hqos-enable-all", types.YChild{"HqosEnableAll", &qos.HqosEnableAll})
    qos.EntityData.Children.Append("ingress-model-root-def", types.YChild{"IngressModelRootDef", &qos.IngressModelRootDef})
    qos.EntityData.Children.Append("ingress-models", types.YChild{"IngressModels", &qos.IngressModels})
    qos.EntityData.Children.Append("class-maps-root-def", types.YChild{"ClassMapsRootDef", &qos.ClassMapsRootDef})
    qos.EntityData.Children.Append("class-maps", types.YChild{"ClassMaps", &qos.ClassMaps})
    qos.EntityData.Leafs = types.NewOrderedMap()

    qos.EntityData.YListKeys = []string {}

    return &(qos.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_HqosEnableAll
// Configure Hqos profile
type HwModuleProfileConfig_Profile_Qos_HqosEnableAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hqos profile value. The type is interface{} with range: 0..4294967295.
    HqosEnable interface{}
}

func (hqosEnableAll *HwModuleProfileConfig_Profile_Qos_HqosEnableAll) GetEntityData() *types.CommonEntityData {
    hqosEnableAll.EntityData.YFilter = hqosEnableAll.YFilter
    hqosEnableAll.EntityData.YangName = "hqos-enable-all"
    hqosEnableAll.EntityData.BundleName = "cisco_ios_xr"
    hqosEnableAll.EntityData.ParentYangName = "qos"
    hqosEnableAll.EntityData.SegmentPath = "hqos-enable-all"
    hqosEnableAll.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/qos/" + hqosEnableAll.EntityData.SegmentPath
    hqosEnableAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hqosEnableAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hqosEnableAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hqosEnableAll.EntityData.Children = types.NewOrderedMap()
    hqosEnableAll.EntityData.Leafs = types.NewOrderedMap()
    hqosEnableAll.EntityData.Leafs.Append("hqos-enable", types.YLeaf{"HqosEnable", hqosEnableAll.HqosEnable})

    hqosEnableAll.EntityData.YListKeys = []string {}

    return &(hqosEnableAll.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_IngressModelRootDef
// Configure Ingress Model Default
type HwModuleProfileConfig_Profile_Qos_IngressModelRootDef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ingress Model Default. The type is interface{} with range: 0..4294967295.
    IngressModelLeafDef interface{}
}

func (ingressModelRootDef *HwModuleProfileConfig_Profile_Qos_IngressModelRootDef) GetEntityData() *types.CommonEntityData {
    ingressModelRootDef.EntityData.YFilter = ingressModelRootDef.YFilter
    ingressModelRootDef.EntityData.YangName = "ingress-model-root-def"
    ingressModelRootDef.EntityData.BundleName = "cisco_ios_xr"
    ingressModelRootDef.EntityData.ParentYangName = "qos"
    ingressModelRootDef.EntityData.SegmentPath = "ingress-model-root-def"
    ingressModelRootDef.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/qos/" + ingressModelRootDef.EntityData.SegmentPath
    ingressModelRootDef.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressModelRootDef.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressModelRootDef.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressModelRootDef.EntityData.Children = types.NewOrderedMap()
    ingressModelRootDef.EntityData.Leafs = types.NewOrderedMap()
    ingressModelRootDef.EntityData.Leafs.Append("ingress-model-leaf-def", types.YLeaf{"IngressModelLeafDef", ingressModelRootDef.IngressModelLeafDef})

    ingressModelRootDef.EntityData.YListKeys = []string {}

    return &(ingressModelRootDef.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_IngressModels
// Configure Ingress Model Root
type HwModuleProfileConfig_Profile_Qos_IngressModels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Ingress Model. The type is slice of
    // HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel.
    IngressModel []*HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel
}

func (ingressModels *HwModuleProfileConfig_Profile_Qos_IngressModels) GetEntityData() *types.CommonEntityData {
    ingressModels.EntityData.YFilter = ingressModels.YFilter
    ingressModels.EntityData.YangName = "ingress-models"
    ingressModels.EntityData.BundleName = "cisco_ios_xr"
    ingressModels.EntityData.ParentYangName = "qos"
    ingressModels.EntityData.SegmentPath = "ingress-models"
    ingressModels.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/qos/" + ingressModels.EntityData.SegmentPath
    ingressModels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressModels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressModels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressModels.EntityData.Children = types.NewOrderedMap()
    ingressModels.EntityData.Children.Append("ingress-model", types.YChild{"IngressModel", nil})
    for i := range ingressModels.IngressModel {
        ingressModels.EntityData.Children.Append(types.GetSegmentPath(ingressModels.IngressModel[i]), types.YChild{"IngressModel", ingressModels.IngressModel[i]})
    }
    ingressModels.EntityData.Leafs = types.NewOrderedMap()

    ingressModels.EntityData.YListKeys = []string {}

    return &(ingressModels.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel
// Configure Ingress Model
type HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. NodeName. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    NodeName interface{}

    // Configure Ingress Model Leaf. The type is slice of
    // HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf.
    IngressModelLeaf []*HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf
}

func (ingressModel *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel) GetEntityData() *types.CommonEntityData {
    ingressModel.EntityData.YFilter = ingressModel.YFilter
    ingressModel.EntityData.YangName = "ingress-model"
    ingressModel.EntityData.BundleName = "cisco_ios_xr"
    ingressModel.EntityData.ParentYangName = "ingress-models"
    ingressModel.EntityData.SegmentPath = "ingress-model" + types.AddKeyToken(ingressModel.NodeName, "node-name")
    ingressModel.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/qos/ingress-models/" + ingressModel.EntityData.SegmentPath
    ingressModel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressModel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressModel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressModel.EntityData.Children = types.NewOrderedMap()
    ingressModel.EntityData.Children.Append("ingress-model-leaf", types.YChild{"IngressModelLeaf", nil})
    for i := range ingressModel.IngressModelLeaf {
        ingressModel.EntityData.Children.Append(types.GetSegmentPath(ingressModel.IngressModelLeaf[i]), types.YChild{"IngressModelLeaf", ingressModel.IngressModelLeaf[i]})
    }
    ingressModel.EntityData.Leafs = types.NewOrderedMap()
    ingressModel.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", ingressModel.NodeName})

    ingressModel.EntityData.YListKeys = []string {"NodeName"}

    return &(ingressModel.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf
// Configure Ingress Model Leaf
type HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Location. The type is interface{} with range:
    // 0..4294967295.
    Location interface{}

    // IngressModelLeaf. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    IngressModelLeaf interface{}
}

func (ingressModelLeaf *HwModuleProfileConfig_Profile_Qos_IngressModels_IngressModel_IngressModelLeaf) GetEntityData() *types.CommonEntityData {
    ingressModelLeaf.EntityData.YFilter = ingressModelLeaf.YFilter
    ingressModelLeaf.EntityData.YangName = "ingress-model-leaf"
    ingressModelLeaf.EntityData.BundleName = "cisco_ios_xr"
    ingressModelLeaf.EntityData.ParentYangName = "ingress-model"
    ingressModelLeaf.EntityData.SegmentPath = "ingress-model-leaf" + types.AddKeyToken(ingressModelLeaf.Location, "location")
    ingressModelLeaf.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/qos/ingress-models/ingress-model/" + ingressModelLeaf.EntityData.SegmentPath
    ingressModelLeaf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressModelLeaf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressModelLeaf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressModelLeaf.EntityData.Children = types.NewOrderedMap()
    ingressModelLeaf.EntityData.Leafs = types.NewOrderedMap()
    ingressModelLeaf.EntityData.Leafs.Append("location", types.YLeaf{"Location", ingressModelLeaf.Location})
    ingressModelLeaf.EntityData.Leafs.Append("ingress-model-leaf", types.YLeaf{"IngressModelLeaf", ingressModelLeaf.IngressModelLeaf})

    ingressModelLeaf.EntityData.YListKeys = []string {"Location"}

    return &(ingressModelLeaf.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef
// Configure Class Maps Default
type HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class Map Size Default. The type is interface{} with range: 0..4294967295.
    ClassMapSizeDef interface{}
}

func (classMapsRootDef *HwModuleProfileConfig_Profile_Qos_ClassMapsRootDef) GetEntityData() *types.CommonEntityData {
    classMapsRootDef.EntityData.YFilter = classMapsRootDef.YFilter
    classMapsRootDef.EntityData.YangName = "class-maps-root-def"
    classMapsRootDef.EntityData.BundleName = "cisco_ios_xr"
    classMapsRootDef.EntityData.ParentYangName = "qos"
    classMapsRootDef.EntityData.SegmentPath = "class-maps-root-def"
    classMapsRootDef.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/qos/" + classMapsRootDef.EntityData.SegmentPath
    classMapsRootDef.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMapsRootDef.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMapsRootDef.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMapsRootDef.EntityData.Children = types.NewOrderedMap()
    classMapsRootDef.EntityData.Leafs = types.NewOrderedMap()
    classMapsRootDef.EntityData.Leafs.Append("class-map-size-def", types.YLeaf{"ClassMapSizeDef", classMapsRootDef.ClassMapSizeDef})

    classMapsRootDef.EntityData.YListKeys = []string {}

    return &(classMapsRootDef.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_ClassMaps
// Configure Class Map Root
type HwModuleProfileConfig_Profile_Qos_ClassMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Class Maps. The type is slice of
    // HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap.
    ClassMap []*HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap
}

func (classMaps *HwModuleProfileConfig_Profile_Qos_ClassMaps) GetEntityData() *types.CommonEntityData {
    classMaps.EntityData.YFilter = classMaps.YFilter
    classMaps.EntityData.YangName = "class-maps"
    classMaps.EntityData.BundleName = "cisco_ios_xr"
    classMaps.EntityData.ParentYangName = "qos"
    classMaps.EntityData.SegmentPath = "class-maps"
    classMaps.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/qos/" + classMaps.EntityData.SegmentPath
    classMaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMaps.EntityData.Children = types.NewOrderedMap()
    classMaps.EntityData.Children.Append("class-map", types.YChild{"ClassMap", nil})
    for i := range classMaps.ClassMap {
        classMaps.EntityData.Children.Append(types.GetSegmentPath(classMaps.ClassMap[i]), types.YChild{"ClassMap", classMaps.ClassMap[i]})
    }
    classMaps.EntityData.Leafs = types.NewOrderedMap()

    classMaps.EntityData.YListKeys = []string {}

    return &(classMaps.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap
// Configure Class Maps
type HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. NodeName. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    NodeName interface{}

    // Class Map Size. The type is slice of
    // HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize.
    ClassMapSize []*HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize
}

func (classMap *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap) GetEntityData() *types.CommonEntityData {
    classMap.EntityData.YFilter = classMap.YFilter
    classMap.EntityData.YangName = "class-map"
    classMap.EntityData.BundleName = "cisco_ios_xr"
    classMap.EntityData.ParentYangName = "class-maps"
    classMap.EntityData.SegmentPath = "class-map" + types.AddKeyToken(classMap.NodeName, "node-name")
    classMap.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/qos/class-maps/" + classMap.EntityData.SegmentPath
    classMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMap.EntityData.Children = types.NewOrderedMap()
    classMap.EntityData.Children.Append("class-map-size", types.YChild{"ClassMapSize", nil})
    for i := range classMap.ClassMapSize {
        classMap.EntityData.Children.Append(types.GetSegmentPath(classMap.ClassMapSize[i]), types.YChild{"ClassMapSize", classMap.ClassMapSize[i]})
    }
    classMap.EntityData.Leafs = types.NewOrderedMap()
    classMap.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", classMap.NodeName})

    classMap.EntityData.YListKeys = []string {"NodeName"}

    return &(classMap.EntityData)
}

// HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize
// Class Map Size
type HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Location. The type is interface{} with range:
    // 0..4294967295.
    Location interface{}

    // ClassMapSize. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    ClassMapSize interface{}
}

func (classMapSize *HwModuleProfileConfig_Profile_Qos_ClassMaps_ClassMap_ClassMapSize) GetEntityData() *types.CommonEntityData {
    classMapSize.EntityData.YFilter = classMapSize.YFilter
    classMapSize.EntityData.YangName = "class-map-size"
    classMapSize.EntityData.BundleName = "cisco_ios_xr"
    classMapSize.EntityData.ParentYangName = "class-map"
    classMapSize.EntityData.SegmentPath = "class-map-size" + types.AddKeyToken(classMapSize.Location, "location")
    classMapSize.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/profile/qos/class-maps/class-map/" + classMapSize.EntityData.SegmentPath
    classMapSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMapSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMapSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMapSize.EntityData.Children = types.NewOrderedMap()
    classMapSize.EntityData.Leafs = types.NewOrderedMap()
    classMapSize.EntityData.Leafs.Append("location", types.YLeaf{"Location", classMapSize.Location})
    classMapSize.EntityData.Leafs.Append("class-map-size", types.YLeaf{"ClassMapSize", classMapSize.ClassMapSize})

    classMapSize.EntityData.YListKeys = []string {"Location"}

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
    fibScale.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/" + fibScale.EntityData.SegmentPath
    fibScale.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibScale.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibScale.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibScale.EntityData.Children = types.NewOrderedMap()
    fibScale.EntityData.Children.Append("ipv6-unicast-scale-no-tcam", types.YChild{"Ipv6UnicastScaleNoTcam", &fibScale.Ipv6UnicastScaleNoTcam})
    fibScale.EntityData.Children.Append("ipv4-unicast-scale-no-tcam", types.YChild{"Ipv4UnicastScaleNoTcam", &fibScale.Ipv4UnicastScaleNoTcam})
    fibScale.EntityData.Leafs = types.NewOrderedMap()

    fibScale.EntityData.YListKeys = []string {}

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
    ipv6UnicastScaleNoTcam.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/fib-scale/" + ipv6UnicastScaleNoTcam.EntityData.SegmentPath
    ipv6UnicastScaleNoTcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6UnicastScaleNoTcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6UnicastScaleNoTcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6UnicastScaleNoTcam.EntityData.Children = types.NewOrderedMap()
    ipv6UnicastScaleNoTcam.EntityData.Children.Append("scale-ipv6-no-tcam", types.YChild{"ScaleIpv6NoTcam", &ipv6UnicastScaleNoTcam.ScaleIpv6NoTcam})
    ipv6UnicastScaleNoTcam.EntityData.Leafs = types.NewOrderedMap()

    ipv6UnicastScaleNoTcam.EntityData.YListKeys = []string {}

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
    scaleIpv6NoTcam.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/fib-scale/ipv6-unicast-scale-no-tcam/" + scaleIpv6NoTcam.EntityData.SegmentPath
    scaleIpv6NoTcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scaleIpv6NoTcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scaleIpv6NoTcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scaleIpv6NoTcam.EntityData.Children = types.NewOrderedMap()
    scaleIpv6NoTcam.EntityData.Leafs = types.NewOrderedMap()
    scaleIpv6NoTcam.EntityData.Leafs.Append("internet-optimized-ipv6-no-tcam", types.YLeaf{"InternetOptimizedIpv6NoTcam", scaleIpv6NoTcam.InternetOptimizedIpv6NoTcam})

    scaleIpv6NoTcam.EntityData.YListKeys = []string {}

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
    ipv4UnicastScaleNoTcam.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/fib-scale/" + ipv4UnicastScaleNoTcam.EntityData.SegmentPath
    ipv4UnicastScaleNoTcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4UnicastScaleNoTcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4UnicastScaleNoTcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4UnicastScaleNoTcam.EntityData.Children = types.NewOrderedMap()
    ipv4UnicastScaleNoTcam.EntityData.Children.Append("scale-ipv4-no-tcam", types.YChild{"ScaleIpv4NoTcam", &ipv4UnicastScaleNoTcam.ScaleIpv4NoTcam})
    ipv4UnicastScaleNoTcam.EntityData.Leafs = types.NewOrderedMap()

    ipv4UnicastScaleNoTcam.EntityData.YListKeys = []string {}

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
    scaleIpv4NoTcam.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/fib-scale/ipv4-unicast-scale-no-tcam/" + scaleIpv4NoTcam.EntityData.SegmentPath
    scaleIpv4NoTcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scaleIpv4NoTcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scaleIpv4NoTcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scaleIpv4NoTcam.EntityData.Children = types.NewOrderedMap()
    scaleIpv4NoTcam.EntityData.Leafs = types.NewOrderedMap()
    scaleIpv4NoTcam.EntityData.Leafs.Append("optimized-ipv4-no-tcam", types.YLeaf{"OptimizedIpv4NoTcam", scaleIpv4NoTcam.OptimizedIpv4NoTcam})

    scaleIpv4NoTcam.EntityData.YListKeys = []string {}

    return &(scaleIpv4NoTcam.EntityData)
}

// HwModuleProfileConfig_OrchestratedLinecardReload
// Configure OLR.
type HwModuleProfileConfig_OrchestratedLinecardReload struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Plane type.
    PlaneTableEntries HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries
}

func (orchestratedLinecardReload *HwModuleProfileConfig_OrchestratedLinecardReload) GetEntityData() *types.CommonEntityData {
    orchestratedLinecardReload.EntityData.YFilter = orchestratedLinecardReload.YFilter
    orchestratedLinecardReload.EntityData.YangName = "orchestrated-linecard-reload"
    orchestratedLinecardReload.EntityData.BundleName = "cisco_ios_xr"
    orchestratedLinecardReload.EntityData.ParentYangName = "hw-module-profile-config"
    orchestratedLinecardReload.EntityData.SegmentPath = "orchestrated-linecard-reload"
    orchestratedLinecardReload.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/" + orchestratedLinecardReload.EntityData.SegmentPath
    orchestratedLinecardReload.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    orchestratedLinecardReload.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    orchestratedLinecardReload.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    orchestratedLinecardReload.EntityData.Children = types.NewOrderedMap()
    orchestratedLinecardReload.EntityData.Children.Append("plane-table-entries", types.YChild{"PlaneTableEntries", &orchestratedLinecardReload.PlaneTableEntries})
    orchestratedLinecardReload.EntityData.Leafs = types.NewOrderedMap()

    orchestratedLinecardReload.EntityData.YListKeys = []string {}

    return &(orchestratedLinecardReload.EntityData)
}

// HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries
// Plane type
type HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Plane : A or B. The type is slice of
    // HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry.
    PlaneTableEntry []*HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry
}

func (planeTableEntries *HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries) GetEntityData() *types.CommonEntityData {
    planeTableEntries.EntityData.YFilter = planeTableEntries.YFilter
    planeTableEntries.EntityData.YangName = "plane-table-entries"
    planeTableEntries.EntityData.BundleName = "cisco_ios_xr"
    planeTableEntries.EntityData.ParentYangName = "orchestrated-linecard-reload"
    planeTableEntries.EntityData.SegmentPath = "plane-table-entries"
    planeTableEntries.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/orchestrated-linecard-reload/" + planeTableEntries.EntityData.SegmentPath
    planeTableEntries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    planeTableEntries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    planeTableEntries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    planeTableEntries.EntityData.Children = types.NewOrderedMap()
    planeTableEntries.EntityData.Children.Append("plane-table-entry", types.YChild{"PlaneTableEntry", nil})
    for i := range planeTableEntries.PlaneTableEntry {
        planeTableEntries.EntityData.Children.Append(types.GetSegmentPath(planeTableEntries.PlaneTableEntry[i]), types.YChild{"PlaneTableEntry", planeTableEntries.PlaneTableEntry[i]})
    }
    planeTableEntries.EntityData.Leafs = types.NewOrderedMap()

    planeTableEntries.EntityData.YListKeys = []string {}

    return &(planeTableEntries.EntityData)
}

// HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry
// Plane : A or B
type HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Plane Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    PlaneName interface{}

    // Rack Information.
    RackTableEntries HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry_RackTableEntries
}

func (planeTableEntry *HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry) GetEntityData() *types.CommonEntityData {
    planeTableEntry.EntityData.YFilter = planeTableEntry.YFilter
    planeTableEntry.EntityData.YangName = "plane-table-entry"
    planeTableEntry.EntityData.BundleName = "cisco_ios_xr"
    planeTableEntry.EntityData.ParentYangName = "plane-table-entries"
    planeTableEntry.EntityData.SegmentPath = "plane-table-entry" + types.AddKeyToken(planeTableEntry.PlaneName, "plane-name")
    planeTableEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/orchestrated-linecard-reload/plane-table-entries/" + planeTableEntry.EntityData.SegmentPath
    planeTableEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    planeTableEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    planeTableEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    planeTableEntry.EntityData.Children = types.NewOrderedMap()
    planeTableEntry.EntityData.Children.Append("rack-table-entries", types.YChild{"RackTableEntries", &planeTableEntry.RackTableEntries})
    planeTableEntry.EntityData.Leafs = types.NewOrderedMap()
    planeTableEntry.EntityData.Leafs.Append("plane-name", types.YLeaf{"PlaneName", planeTableEntry.PlaneName})

    planeTableEntry.EntityData.YListKeys = []string {"PlaneName"}

    return &(planeTableEntry.EntityData)
}

// HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry_RackTableEntries
// Rack Information
type HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry_RackTableEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack Number. The type is slice of
    // HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry_RackTableEntries_RackTableEntry.
    RackTableEntry []*HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry_RackTableEntries_RackTableEntry
}

func (rackTableEntries *HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry_RackTableEntries) GetEntityData() *types.CommonEntityData {
    rackTableEntries.EntityData.YFilter = rackTableEntries.YFilter
    rackTableEntries.EntityData.YangName = "rack-table-entries"
    rackTableEntries.EntityData.BundleName = "cisco_ios_xr"
    rackTableEntries.EntityData.ParentYangName = "plane-table-entry"
    rackTableEntries.EntityData.SegmentPath = "rack-table-entries"
    rackTableEntries.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/orchestrated-linecard-reload/plane-table-entries/plane-table-entry/" + rackTableEntries.EntityData.SegmentPath
    rackTableEntries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackTableEntries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackTableEntries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackTableEntries.EntityData.Children = types.NewOrderedMap()
    rackTableEntries.EntityData.Children.Append("rack-table-entry", types.YChild{"RackTableEntry", nil})
    for i := range rackTableEntries.RackTableEntry {
        rackTableEntries.EntityData.Children.Append(types.GetSegmentPath(rackTableEntries.RackTableEntry[i]), types.YChild{"RackTableEntry", rackTableEntries.RackTableEntry[i]})
    }
    rackTableEntries.EntityData.Leafs = types.NewOrderedMap()

    rackTableEntries.EntityData.YListKeys = []string {}

    return &(rackTableEntries.EntityData)
}

// HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry_RackTableEntries_RackTableEntry
// Rack Number
type HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry_RackTableEntries_RackTableEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Rack Number. The type is interface{} with range:
    // 0..15.
    RackNum interface{}

    // Please enter linecards separated via comma eg 0,3,6,11,13. The type is
    // string.
    NodesList interface{}
}

func (rackTableEntry *HwModuleProfileConfig_OrchestratedLinecardReload_PlaneTableEntries_PlaneTableEntry_RackTableEntries_RackTableEntry) GetEntityData() *types.CommonEntityData {
    rackTableEntry.EntityData.YFilter = rackTableEntry.YFilter
    rackTableEntry.EntityData.YangName = "rack-table-entry"
    rackTableEntry.EntityData.BundleName = "cisco_ios_xr"
    rackTableEntry.EntityData.ParentYangName = "rack-table-entries"
    rackTableEntry.EntityData.SegmentPath = "rack-table-entry" + types.AddKeyToken(rackTableEntry.RackNum, "rack-num")
    rackTableEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/orchestrated-linecard-reload/plane-table-entries/plane-table-entry/rack-table-entries/" + rackTableEntry.EntityData.SegmentPath
    rackTableEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackTableEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackTableEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackTableEntry.EntityData.Children = types.NewOrderedMap()
    rackTableEntry.EntityData.Leafs = types.NewOrderedMap()
    rackTableEntry.EntityData.Leafs.Append("rack-num", types.YLeaf{"RackNum", rackTableEntry.RackNum})
    rackTableEntry.EntityData.Leafs.Append("nodes-list", types.YLeaf{"NodesList", rackTableEntry.NodesList})

    rackTableEntry.EntityData.YListKeys = []string {"RackNum"}

    return &(rackTableEntry.EntityData)
}

// HwModuleProfileConfig_Tcam
// Configure Tcam.
type HwModuleProfileConfig_Tcam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Fib scale for Tcam.
    FibTcamScale HwModuleProfileConfig_Tcam_FibTcamScale
}

func (tcam *HwModuleProfileConfig_Tcam) GetEntityData() *types.CommonEntityData {
    tcam.EntityData.YFilter = tcam.YFilter
    tcam.EntityData.YangName = "tcam"
    tcam.EntityData.BundleName = "cisco_ios_xr"
    tcam.EntityData.ParentYangName = "hw-module-profile-config"
    tcam.EntityData.SegmentPath = "tcam"
    tcam.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/" + tcam.EntityData.SegmentPath
    tcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcam.EntityData.Children = types.NewOrderedMap()
    tcam.EntityData.Children.Append("fib-tcam-scale", types.YChild{"FibTcamScale", &tcam.FibTcamScale})
    tcam.EntityData.Leafs = types.NewOrderedMap()

    tcam.EntityData.YListKeys = []string {}

    return &(tcam.EntityData)
}

// HwModuleProfileConfig_Tcam_FibTcamScale
// Configure Fib scale for Tcam.
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
    fibTcamScale.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/tcam/" + fibTcamScale.EntityData.SegmentPath
    fibTcamScale.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibTcamScale.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibTcamScale.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibTcamScale.EntityData.Children = types.NewOrderedMap()
    fibTcamScale.EntityData.Children.Append("ipv4-unicast-scale", types.YChild{"Ipv4UnicastScale", &fibTcamScale.Ipv4UnicastScale})
    fibTcamScale.EntityData.Leafs = types.NewOrderedMap()

    fibTcamScale.EntityData.YListKeys = []string {}

    return &(fibTcamScale.EntityData)
}

// HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale
// IPv4 table for TCAM LC Scale.
type HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Scale for IPv4 table for TCAM LC. The type is bool.
    Ipv4Scale interface{}
}

func (ipv4UnicastScale *HwModuleProfileConfig_Tcam_FibTcamScale_Ipv4UnicastScale) GetEntityData() *types.CommonEntityData {
    ipv4UnicastScale.EntityData.YFilter = ipv4UnicastScale.YFilter
    ipv4UnicastScale.EntityData.YangName = "ipv4-unicast-scale"
    ipv4UnicastScale.EntityData.BundleName = "cisco_ios_xr"
    ipv4UnicastScale.EntityData.ParentYangName = "fib-tcam-scale"
    ipv4UnicastScale.EntityData.SegmentPath = "ipv4-unicast-scale"
    ipv4UnicastScale.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/tcam/fib-tcam-scale/" + ipv4UnicastScale.EntityData.SegmentPath
    ipv4UnicastScale.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4UnicastScale.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4UnicastScale.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4UnicastScale.EntityData.Children = types.NewOrderedMap()
    ipv4UnicastScale.EntityData.Leafs = types.NewOrderedMap()
    ipv4UnicastScale.EntityData.Leafs.Append("ipv4-scale", types.YLeaf{"Ipv4Scale", ipv4UnicastScale.Ipv4Scale})

    ipv4UnicastScale.EntityData.YListKeys = []string {}

    return &(ipv4UnicastScale.EntityData)
}

// HwModuleProfileConfig_Qosqppb
// Configure profile.
type HwModuleProfileConfig_Qosqppb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure IPv6 destination short profile.
    Ipv6Scale HwModuleProfileConfig_Qosqppb_Ipv6Scale
}

func (qosqppb *HwModuleProfileConfig_Qosqppb) GetEntityData() *types.CommonEntityData {
    qosqppb.EntityData.YFilter = qosqppb.YFilter
    qosqppb.EntityData.YangName = "qosqppb"
    qosqppb.EntityData.BundleName = "cisco_ios_xr"
    qosqppb.EntityData.ParentYangName = "hw-module-profile-config"
    qosqppb.EntityData.SegmentPath = "qosqppb"
    qosqppb.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/" + qosqppb.EntityData.SegmentPath
    qosqppb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosqppb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosqppb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosqppb.EntityData.Children = types.NewOrderedMap()
    qosqppb.EntityData.Children.Append("ipv6-scale", types.YChild{"Ipv6Scale", &qosqppb.Ipv6Scale})
    qosqppb.EntityData.Leafs = types.NewOrderedMap()

    qosqppb.EntityData.YListKeys = []string {}

    return &(qosqppb.EntityData)
}

// HwModuleProfileConfig_Qosqppb_Ipv6Scale
// Configure IPv6 destination short profile
type HwModuleProfileConfig_Qosqppb_Ipv6Scale struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 short profile value. The type is bool.
    Ipv6Short interface{}
}

func (ipv6Scale *HwModuleProfileConfig_Qosqppb_Ipv6Scale) GetEntityData() *types.CommonEntityData {
    ipv6Scale.EntityData.YFilter = ipv6Scale.YFilter
    ipv6Scale.EntityData.YangName = "ipv6-scale"
    ipv6Scale.EntityData.BundleName = "cisco_ios_xr"
    ipv6Scale.EntityData.ParentYangName = "qosqppb"
    ipv6Scale.EntityData.SegmentPath = "ipv6-scale"
    ipv6Scale.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-hw-profile-cfg:hw-module-profile-config/qosqppb/" + ipv6Scale.EntityData.SegmentPath
    ipv6Scale.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Scale.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Scale.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Scale.EntityData.Children = types.NewOrderedMap()
    ipv6Scale.EntityData.Leafs = types.NewOrderedMap()
    ipv6Scale.EntityData.Leafs.Append("ipv6-short", types.YLeaf{"Ipv6Short", ipv6Scale.Ipv6Short})

    ipv6Scale.EntityData.YListKeys = []string {}

    return &(ipv6Scale.EntityData)
}

