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
type Igmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF related configuration.
    Vrfs Igmp_Vrfs

    // Default Context.
    DefaultContext Igmp_DefaultContext
}

func (igmp *Igmp) GetEntityData() *types.CommonEntityData {
    igmp.EntityData.YFilter = igmp.YFilter
    igmp.EntityData.YangName = "igmp"
    igmp.EntityData.BundleName = "cisco_ios_xr"
    igmp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-igmp-cfg"
    igmp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-igmp-cfg:igmp"
    igmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igmp.EntityData.Children = types.NewOrderedMap()
    igmp.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &igmp.Vrfs})
    igmp.EntityData.Children.Append("default-context", types.YChild{"DefaultContext", &igmp.DefaultContext})
    igmp.EntityData.Leafs = types.NewOrderedMap()

    igmp.EntityData.YListKeys = []string {}

    return &(igmp.EntityData)
}

// Igmp_Vrfs
// VRF related configuration
type Igmp_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a particular vrf. The type is slice of Igmp_Vrfs_Vrf.
    Vrf []*Igmp_Vrfs_Vrf
}

func (vrfs *Igmp_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "igmp"
    vrfs.EntityData.SegmentPath = "vrfs"
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

// Igmp_Vrfs_Vrf
// Configuration for a particular vrf
type Igmp_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
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

func (vrf *Igmp_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("traffic", types.YChild{"Traffic", &vrf.Traffic})
    vrf.EntityData.Children.Append("inheritable-defaults", types.YChild{"InheritableDefaults", &vrf.InheritableDefaults})
    vrf.EntityData.Children.Append("ssm-access-groups", types.YChild{"SsmAccessGroups", &vrf.SsmAccessGroups})
    vrf.EntityData.Children.Append("maximum", types.YChild{"Maximum", &vrf.Maximum})
    vrf.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &vrf.Interfaces})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("ssmdns-query-group", types.YLeaf{"SsmdnsQueryGroup", vrf.SsmdnsQueryGroup})
    vrf.EntityData.Leafs.Append("robustness", types.YLeaf{"Robustness", vrf.Robustness})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Igmp_Vrfs_Vrf_Traffic
// Configure IGMP Traffic variables
type Igmp_Vrfs_Vrf_Traffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the route-policy profile. The type is string with length: 1..64.
    Profile interface{}
}

func (traffic *Igmp_Vrfs_Vrf_Traffic) GetEntityData() *types.CommonEntityData {
    traffic.EntityData.YFilter = traffic.YFilter
    traffic.EntityData.YangName = "traffic"
    traffic.EntityData.BundleName = "cisco_ios_xr"
    traffic.EntityData.ParentYangName = "vrf"
    traffic.EntityData.SegmentPath = "traffic"
    traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traffic.EntityData.Children = types.NewOrderedMap()
    traffic.EntityData.Leafs = types.NewOrderedMap()
    traffic.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", traffic.Profile})

    traffic.EntityData.YListKeys = []string {}

    return &(traffic.EntityData)
}

// Igmp_Vrfs_Vrf_InheritableDefaults
// Inheritable Defaults
type Igmp_Vrfs_Vrf_InheritableDefaults struct {
    EntityData types.CommonEntityData
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

func (inheritableDefaults *Igmp_Vrfs_Vrf_InheritableDefaults) GetEntityData() *types.CommonEntityData {
    inheritableDefaults.EntityData.YFilter = inheritableDefaults.YFilter
    inheritableDefaults.EntityData.YangName = "inheritable-defaults"
    inheritableDefaults.EntityData.BundleName = "cisco_ios_xr"
    inheritableDefaults.EntityData.ParentYangName = "vrf"
    inheritableDefaults.EntityData.SegmentPath = "inheritable-defaults"
    inheritableDefaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritableDefaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritableDefaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritableDefaults.EntityData.Children = types.NewOrderedMap()
    inheritableDefaults.EntityData.Children.Append("maximum-groups-per-interface-oor", types.YChild{"MaximumGroupsPerInterfaceOor", &inheritableDefaults.MaximumGroupsPerInterfaceOor})
    inheritableDefaults.EntityData.Children.Append("explicit-tracking", types.YChild{"ExplicitTracking", &inheritableDefaults.ExplicitTracking})
    inheritableDefaults.EntityData.Leafs = types.NewOrderedMap()
    inheritableDefaults.EntityData.Leafs.Append("query-timeout", types.YLeaf{"QueryTimeout", inheritableDefaults.QueryTimeout})
    inheritableDefaults.EntityData.Leafs.Append("access-group", types.YLeaf{"AccessGroup", inheritableDefaults.AccessGroup})
    inheritableDefaults.EntityData.Leafs.Append("query-max-response-time", types.YLeaf{"QueryMaxResponseTime", inheritableDefaults.QueryMaxResponseTime})
    inheritableDefaults.EntityData.Leafs.Append("version", types.YLeaf{"Version", inheritableDefaults.Version})
    inheritableDefaults.EntityData.Leafs.Append("router-enable", types.YLeaf{"RouterEnable", inheritableDefaults.RouterEnable})
    inheritableDefaults.EntityData.Leafs.Append("query-interval", types.YLeaf{"QueryInterval", inheritableDefaults.QueryInterval})

    inheritableDefaults.EntityData.YListKeys = []string {}

    return &(inheritableDefaults.EntityData)
}

// Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetEntityData() *types.CommonEntityData {
    maximumGroupsPerInterfaceOor.EntityData.YFilter = maximumGroupsPerInterfaceOor.YFilter
    maximumGroupsPerInterfaceOor.EntityData.YangName = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.BundleName = "cisco_ios_xr"
    maximumGroupsPerInterfaceOor.EntityData.ParentYangName = "inheritable-defaults"
    maximumGroupsPerInterfaceOor.EntityData.SegmentPath = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumGroupsPerInterfaceOor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumGroupsPerInterfaceOor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumGroupsPerInterfaceOor.EntityData.Children = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("maximum-groups", types.YLeaf{"MaximumGroups", maximumGroupsPerInterfaceOor.MaximumGroups})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", maximumGroupsPerInterfaceOor.WarningThreshold})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", maximumGroupsPerInterfaceOor.AccessListName})

    maximumGroupsPerInterfaceOor.EntityData.YListKeys = []string {}

    return &(maximumGroupsPerInterfaceOor.EntityData)
}

// Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Igmp_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetEntityData() *types.CommonEntityData {
    explicitTracking.EntityData.YFilter = explicitTracking.YFilter
    explicitTracking.EntityData.YangName = "explicit-tracking"
    explicitTracking.EntityData.BundleName = "cisco_ios_xr"
    explicitTracking.EntityData.ParentYangName = "inheritable-defaults"
    explicitTracking.EntityData.SegmentPath = "explicit-tracking"
    explicitTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitTracking.EntityData.Children = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", explicitTracking.Enable})
    explicitTracking.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", explicitTracking.AccessListName})

    explicitTracking.EntityData.YListKeys = []string {}

    return &(explicitTracking.EntityData)
}

// Igmp_Vrfs_Vrf_SsmAccessGroups
// IGMP Source specific mode
type Igmp_Vrfs_Vrf_SsmAccessGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SSM static Access Group. The type is slice of
    // Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup.
    SsmAccessGroup []*Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup
}

func (ssmAccessGroups *Igmp_Vrfs_Vrf_SsmAccessGroups) GetEntityData() *types.CommonEntityData {
    ssmAccessGroups.EntityData.YFilter = ssmAccessGroups.YFilter
    ssmAccessGroups.EntityData.YangName = "ssm-access-groups"
    ssmAccessGroups.EntityData.BundleName = "cisco_ios_xr"
    ssmAccessGroups.EntityData.ParentYangName = "vrf"
    ssmAccessGroups.EntityData.SegmentPath = "ssm-access-groups"
    ssmAccessGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssmAccessGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssmAccessGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssmAccessGroups.EntityData.Children = types.NewOrderedMap()
    ssmAccessGroups.EntityData.Children.Append("ssm-access-group", types.YChild{"SsmAccessGroup", nil})
    for i := range ssmAccessGroups.SsmAccessGroup {
        ssmAccessGroups.EntityData.Children.Append(types.GetSegmentPath(ssmAccessGroups.SsmAccessGroup[i]), types.YChild{"SsmAccessGroup", ssmAccessGroups.SsmAccessGroup[i]})
    }
    ssmAccessGroups.EntityData.Leafs = types.NewOrderedMap()

    ssmAccessGroups.EntityData.YListKeys = []string {}

    return &(ssmAccessGroups.EntityData)
}

// Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup
// SSM static Access Group
type Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup struct {
    EntityData types.CommonEntityData
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

func (ssmAccessGroup *Igmp_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetEntityData() *types.CommonEntityData {
    ssmAccessGroup.EntityData.YFilter = ssmAccessGroup.YFilter
    ssmAccessGroup.EntityData.YangName = "ssm-access-group"
    ssmAccessGroup.EntityData.BundleName = "cisco_ios_xr"
    ssmAccessGroup.EntityData.ParentYangName = "ssm-access-groups"
    ssmAccessGroup.EntityData.SegmentPath = "ssm-access-group" + types.AddKeyToken(ssmAccessGroup.SourceAddress, "source-address")
    ssmAccessGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssmAccessGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssmAccessGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssmAccessGroup.EntityData.Children = types.NewOrderedMap()
    ssmAccessGroup.EntityData.Leafs = types.NewOrderedMap()
    ssmAccessGroup.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", ssmAccessGroup.SourceAddress})
    ssmAccessGroup.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", ssmAccessGroup.AccessListName})

    ssmAccessGroup.EntityData.YListKeys = []string {"SourceAddress"}

    return &(ssmAccessGroup.EntityData)
}

// Igmp_Vrfs_Vrf_Maximum
// Configure IGMP State Limits
type Igmp_Vrfs_Vrf_Maximum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure maximum number of groups accepted by this router. The type is
    // interface{} with range: 1..75000. The default value is 50000.
    MaximumGroups interface{}
}

func (maximum *Igmp_Vrfs_Vrf_Maximum) GetEntityData() *types.CommonEntityData {
    maximum.EntityData.YFilter = maximum.YFilter
    maximum.EntityData.YangName = "maximum"
    maximum.EntityData.BundleName = "cisco_ios_xr"
    maximum.EntityData.ParentYangName = "vrf"
    maximum.EntityData.SegmentPath = "maximum"
    maximum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximum.EntityData.Children = types.NewOrderedMap()
    maximum.EntityData.Leafs = types.NewOrderedMap()
    maximum.EntityData.Leafs.Append("maximum-groups", types.YLeaf{"MaximumGroups", maximum.MaximumGroups})

    maximum.EntityData.YListKeys = []string {}

    return &(maximum.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces
// Interface-level configuration
type Igmp_Vrfs_Vrf_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Igmp_Vrfs_Vrf_Interfaces_Interface.
    Interface []*Igmp_Vrfs_Vrf_Interfaces_Interface
}

func (interfaces *Igmp_Vrfs_Vrf_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "vrf"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces_Interface
// The name of the interface
type Igmp_Vrfs_Vrf_Interfaces_Interface struct {
    EntityData types.CommonEntityData
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

func (self *Igmp_Vrfs_Vrf_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("join-groups", types.YChild{"JoinGroups", &self.JoinGroups})
    self.EntityData.Children.Append("static-group-group-addresses", types.YChild{"StaticGroupGroupAddresses", &self.StaticGroupGroupAddresses})
    self.EntityData.Children.Append("maximum-groups-per-interface-oor", types.YChild{"MaximumGroupsPerInterfaceOor", &self.MaximumGroupsPerInterfaceOor})
    self.EntityData.Children.Append("explicit-tracking", types.YChild{"ExplicitTracking", &self.ExplicitTracking})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("query-timeout", types.YLeaf{"QueryTimeout", self.QueryTimeout})
    self.EntityData.Leafs.Append("access-group", types.YLeaf{"AccessGroup", self.AccessGroup})
    self.EntityData.Leafs.Append("query-max-response-time", types.YLeaf{"QueryMaxResponseTime", self.QueryMaxResponseTime})
    self.EntityData.Leafs.Append("version", types.YLeaf{"Version", self.Version})
    self.EntityData.Leafs.Append("router-enable", types.YLeaf{"RouterEnable", self.RouterEnable})
    self.EntityData.Leafs.Append("query-interval", types.YLeaf{"QueryInterval", self.QueryInterval})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups
// IGMP join multicast group
// This type is a presence type.
type Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IP group address and optional source address to include. The type is slice
    // of Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup.
    JoinGroup []*Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup

    // IP group address and optional source address to include. The type is slice
    // of Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress.
    JoinGroupSourceAddress []*Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
}

func (joinGroups *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetEntityData() *types.CommonEntityData {
    joinGroups.EntityData.YFilter = joinGroups.YFilter
    joinGroups.EntityData.YangName = "join-groups"
    joinGroups.EntityData.BundleName = "cisco_ios_xr"
    joinGroups.EntityData.ParentYangName = "interface"
    joinGroups.EntityData.SegmentPath = "join-groups"
    joinGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    joinGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    joinGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    joinGroups.EntityData.Children = types.NewOrderedMap()
    joinGroups.EntityData.Children.Append("join-group", types.YChild{"JoinGroup", nil})
    for i := range joinGroups.JoinGroup {
        joinGroups.EntityData.Children.Append(types.GetSegmentPath(joinGroups.JoinGroup[i]), types.YChild{"JoinGroup", joinGroups.JoinGroup[i]})
    }
    joinGroups.EntityData.Children.Append("join-group-source-address", types.YChild{"JoinGroupSourceAddress", nil})
    for i := range joinGroups.JoinGroupSourceAddress {
        joinGroups.EntityData.Children.Append(types.GetSegmentPath(joinGroups.JoinGroupSourceAddress[i]), types.YChild{"JoinGroupSourceAddress", joinGroups.JoinGroupSourceAddress[i]})
    }
    joinGroups.EntityData.Leafs = types.NewOrderedMap()

    joinGroups.EntityData.YListKeys = []string {}

    return &(joinGroups.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup struct {
    EntityData types.CommonEntityData
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

func (joinGroup *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetEntityData() *types.CommonEntityData {
    joinGroup.EntityData.YFilter = joinGroup.YFilter
    joinGroup.EntityData.YangName = "join-group"
    joinGroup.EntityData.BundleName = "cisco_ios_xr"
    joinGroup.EntityData.ParentYangName = "join-groups"
    joinGroup.EntityData.SegmentPath = "join-group" + types.AddKeyToken(joinGroup.GroupAddress, "group-address")
    joinGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    joinGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    joinGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    joinGroup.EntityData.Children = types.NewOrderedMap()
    joinGroup.EntityData.Leafs = types.NewOrderedMap()
    joinGroup.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", joinGroup.GroupAddress})
    joinGroup.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", joinGroup.Mode})

    joinGroup.EntityData.YListKeys = []string {"GroupAddress"}

    return &(joinGroup.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress struct {
    EntityData types.CommonEntityData
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

func (joinGroupSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetEntityData() *types.CommonEntityData {
    joinGroupSourceAddress.EntityData.YFilter = joinGroupSourceAddress.YFilter
    joinGroupSourceAddress.EntityData.YangName = "join-group-source-address"
    joinGroupSourceAddress.EntityData.BundleName = "cisco_ios_xr"
    joinGroupSourceAddress.EntityData.ParentYangName = "join-groups"
    joinGroupSourceAddress.EntityData.SegmentPath = "join-group-source-address" + types.AddKeyToken(joinGroupSourceAddress.GroupAddress, "group-address") + types.AddKeyToken(joinGroupSourceAddress.SourceAddress, "source-address")
    joinGroupSourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    joinGroupSourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    joinGroupSourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    joinGroupSourceAddress.EntityData.Children = types.NewOrderedMap()
    joinGroupSourceAddress.EntityData.Leafs = types.NewOrderedMap()
    joinGroupSourceAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", joinGroupSourceAddress.GroupAddress})
    joinGroupSourceAddress.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", joinGroupSourceAddress.SourceAddress})
    joinGroupSourceAddress.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", joinGroupSourceAddress.Mode})

    joinGroupSourceAddress.EntityData.YListKeys = []string {"GroupAddress", "SourceAddress"}

    return &(joinGroupSourceAddress.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses
// IGMP static multicast group
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress.
    StaticGroupGroupAddress []*Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress.
    StaticGroupGroupAddressSourceAddress []*Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask.
    StaticGroupGroupAddressSourceAddressSourceAddressMask []*Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask.
    StaticGroupGroupAddressGroupAddressMask []*Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress.
    StaticGroupGroupAddressGroupAddressMaskSourceAddress []*Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.
    StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask []*Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
}

func (staticGroupGroupAddresses *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddresses.EntityData.YFilter = staticGroupGroupAddresses.YFilter
    staticGroupGroupAddresses.EntityData.YangName = "static-group-group-addresses"
    staticGroupGroupAddresses.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddresses.EntityData.ParentYangName = "interface"
    staticGroupGroupAddresses.EntityData.SegmentPath = "static-group-group-addresses"
    staticGroupGroupAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddresses.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address", types.YChild{"StaticGroupGroupAddress", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddress {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddress[i]), types.YChild{"StaticGroupGroupAddress", staticGroupGroupAddresses.StaticGroupGroupAddress[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-source-address", types.YChild{"StaticGroupGroupAddressSourceAddress", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i]), types.YChild{"StaticGroupGroupAddressSourceAddress", staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-source-address-source-address-mask", types.YChild{"StaticGroupGroupAddressSourceAddressSourceAddressMask", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i]), types.YChild{"StaticGroupGroupAddressSourceAddressSourceAddressMask", staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-group-address-mask", types.YChild{"StaticGroupGroupAddressGroupAddressMask", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i]), types.YChild{"StaticGroupGroupAddressGroupAddressMask", staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-group-address-mask-source-address", types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddress", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i]), types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddress", staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-group-address-mask-source-address-source-address-mask", types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i]), types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask", staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i]})
    }
    staticGroupGroupAddresses.EntityData.Leafs = types.NewOrderedMap()

    staticGroupGroupAddresses.EntityData.YListKeys = []string {}

    return &(staticGroupGroupAddresses.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddress.EntityData.YFilter = staticGroupGroupAddress.YFilter
    staticGroupGroupAddress.EntityData.YangName = "static-group-group-address"
    staticGroupGroupAddress.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddress.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddress.EntityData.SegmentPath = "static-group-group-address" + types.AddKeyToken(staticGroupGroupAddress.GroupAddress, "group-address")
    staticGroupGroupAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddress.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddress.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddress.GroupAddress})
    staticGroupGroupAddress.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddress.GroupCount})
    staticGroupGroupAddress.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddress.SourceCount})
    staticGroupGroupAddress.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddress.SuppressReport})

    staticGroupGroupAddress.EntityData.YListKeys = []string {"GroupAddress"}

    return &(staticGroupGroupAddress.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressSourceAddress.EntityData.YFilter = staticGroupGroupAddressSourceAddress.YFilter
    staticGroupGroupAddressSourceAddress.EntityData.YangName = "static-group-group-address-source-address"
    staticGroupGroupAddressSourceAddress.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressSourceAddress.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressSourceAddress.EntityData.SegmentPath = "static-group-group-address-source-address" + types.AddKeyToken(staticGroupGroupAddressSourceAddress.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressSourceAddress.SourceAddress, "source-address")
    staticGroupGroupAddressSourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressSourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressSourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressSourceAddress.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddress.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressSourceAddress.GroupAddress})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressSourceAddress.SourceAddress})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressSourceAddress.GroupCount})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressSourceAddress.SourceCount})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressSourceAddress.SuppressReport})

    staticGroupGroupAddressSourceAddress.EntityData.YListKeys = []string {"GroupAddress", "SourceAddress"}

    return &(staticGroupGroupAddressSourceAddress.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.YFilter = staticGroupGroupAddressSourceAddressSourceAddressMask.YFilter
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.YangName = "static-group-group-address-source-address-source-address-mask"
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.SegmentPath = "static-group-group-address-source-address-source-address-mask" + types.AddKeyToken(staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress, "source-address") + types.AddKeyToken(staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask, "source-address-mask")
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address-mask", types.YLeaf{"SourceAddressMask", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressSourceAddressSourceAddressMask.GroupCount})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceCount})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressSourceAddressSourceAddressMask.SuppressReport})

    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.YListKeys = []string {"GroupAddress", "SourceAddress", "SourceAddressMask"}

    return &(staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressGroupAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressGroupAddressMask.EntityData.YFilter = staticGroupGroupAddressGroupAddressMask.YFilter
    staticGroupGroupAddressGroupAddressMask.EntityData.YangName = "static-group-group-address-group-address-mask"
    staticGroupGroupAddressGroupAddressMask.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressGroupAddressMask.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressGroupAddressMask.EntityData.SegmentPath = "static-group-group-address-group-address-mask" + types.AddKeyToken(staticGroupGroupAddressGroupAddressMask.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMask.GroupAddressMask, "group-address-mask")
    staticGroupGroupAddressGroupAddressMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressGroupAddressMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressGroupAddressMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressGroupAddressMask.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressGroupAddressMask.GroupAddress})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("group-address-mask", types.YLeaf{"GroupAddressMask", staticGroupGroupAddressGroupAddressMask.GroupAddressMask})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressGroupAddressMask.GroupCount})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressGroupAddressMask.SourceCount})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressGroupAddressMask.SuppressReport})

    staticGroupGroupAddressGroupAddressMask.EntityData.YListKeys = []string {"GroupAddress", "GroupAddressMask"}

    return &(staticGroupGroupAddressGroupAddressMask.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.YFilter = staticGroupGroupAddressGroupAddressMaskSourceAddress.YFilter
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.YangName = "static-group-group-address-group-address-mask-source-address"
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.SegmentPath = "static-group-group-address-group-address-mask-source-address" + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask, "group-address-mask") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress, "source-address")
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("group-address-mask", types.YLeaf{"GroupAddressMask", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressGroupAddressMaskSourceAddress.SuppressReport})

    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.YListKeys = []string {"GroupAddress", "GroupAddressMask", "SourceAddress"}

    return &(staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.YFilter = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.YFilter
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.YangName = "static-group-group-address-group-address-mask-source-address-source-address-mask"
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.SegmentPath = "static-group-group-address-group-address-mask-source-address-source-address-mask" + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask, "group-address-mask") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress, "source-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask, "source-address-mask")
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-address-mask", types.YLeaf{"GroupAddressMask", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address-mask", types.YLeaf{"SourceAddressMask", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SuppressReport})

    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.YListKeys = []string {"GroupAddress", "GroupAddressMask", "SourceAddress", "SourceAddressMask"}

    return &(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (maximumGroupsPerInterfaceOor *Igmp_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetEntityData() *types.CommonEntityData {
    maximumGroupsPerInterfaceOor.EntityData.YFilter = maximumGroupsPerInterfaceOor.YFilter
    maximumGroupsPerInterfaceOor.EntityData.YangName = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.BundleName = "cisco_ios_xr"
    maximumGroupsPerInterfaceOor.EntityData.ParentYangName = "interface"
    maximumGroupsPerInterfaceOor.EntityData.SegmentPath = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumGroupsPerInterfaceOor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumGroupsPerInterfaceOor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumGroupsPerInterfaceOor.EntityData.Children = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("maximum-groups", types.YLeaf{"MaximumGroups", maximumGroupsPerInterfaceOor.MaximumGroups})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", maximumGroupsPerInterfaceOor.WarningThreshold})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", maximumGroupsPerInterfaceOor.AccessListName})

    maximumGroupsPerInterfaceOor.EntityData.YListKeys = []string {}

    return &(maximumGroupsPerInterfaceOor.EntityData)
}

// Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Igmp_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetEntityData() *types.CommonEntityData {
    explicitTracking.EntityData.YFilter = explicitTracking.YFilter
    explicitTracking.EntityData.YangName = "explicit-tracking"
    explicitTracking.EntityData.BundleName = "cisco_ios_xr"
    explicitTracking.EntityData.ParentYangName = "interface"
    explicitTracking.EntityData.SegmentPath = "explicit-tracking"
    explicitTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitTracking.EntityData.Children = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", explicitTracking.Enable})
    explicitTracking.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", explicitTracking.AccessListName})

    explicitTracking.EntityData.YListKeys = []string {}

    return &(explicitTracking.EntityData)
}

// Igmp_DefaultContext
// Default Context
// This type is a presence type.
type Igmp_DefaultContext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (defaultContext *Igmp_DefaultContext) GetEntityData() *types.CommonEntityData {
    defaultContext.EntityData.YFilter = defaultContext.YFilter
    defaultContext.EntityData.YangName = "default-context"
    defaultContext.EntityData.BundleName = "cisco_ios_xr"
    defaultContext.EntityData.ParentYangName = "igmp"
    defaultContext.EntityData.SegmentPath = "default-context"
    defaultContext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultContext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultContext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultContext.EntityData.Children = types.NewOrderedMap()
    defaultContext.EntityData.Children.Append("nsf", types.YChild{"Nsf", &defaultContext.Nsf})
    defaultContext.EntityData.Children.Append("unicast-qos-adjust", types.YChild{"UnicastQosAdjust", &defaultContext.UnicastQosAdjust})
    defaultContext.EntityData.Children.Append("accounting", types.YChild{"Accounting", &defaultContext.Accounting})
    defaultContext.EntityData.Children.Append("traffic", types.YChild{"Traffic", &defaultContext.Traffic})
    defaultContext.EntityData.Children.Append("inheritable-defaults", types.YChild{"InheritableDefaults", &defaultContext.InheritableDefaults})
    defaultContext.EntityData.Children.Append("ssm-access-groups", types.YChild{"SsmAccessGroups", &defaultContext.SsmAccessGroups})
    defaultContext.EntityData.Children.Append("maximum", types.YChild{"Maximum", &defaultContext.Maximum})
    defaultContext.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &defaultContext.Interfaces})
    defaultContext.EntityData.Leafs = types.NewOrderedMap()
    defaultContext.EntityData.Leafs.Append("ssmdns-query-group", types.YLeaf{"SsmdnsQueryGroup", defaultContext.SsmdnsQueryGroup})
    defaultContext.EntityData.Leafs.Append("robustness", types.YLeaf{"Robustness", defaultContext.Robustness})

    defaultContext.EntityData.YListKeys = []string {}

    return &(defaultContext.EntityData)
}

// Igmp_DefaultContext_Nsf
// Configure NSF specific options
type Igmp_DefaultContext_Nsf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum time for IGMP NSF mode in seconds. The type is interface{} with
    // range: 10..3600. Units are second. The default value is 60.
    Lifetime interface{}
}

func (nsf *Igmp_DefaultContext_Nsf) GetEntityData() *types.CommonEntityData {
    nsf.EntityData.YFilter = nsf.YFilter
    nsf.EntityData.YangName = "nsf"
    nsf.EntityData.BundleName = "cisco_ios_xr"
    nsf.EntityData.ParentYangName = "default-context"
    nsf.EntityData.SegmentPath = "nsf"
    nsf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nsf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nsf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nsf.EntityData.Children = types.NewOrderedMap()
    nsf.EntityData.Leafs = types.NewOrderedMap()
    nsf.EntityData.Leafs.Append("lifetime", types.YLeaf{"Lifetime", nsf.Lifetime})

    nsf.EntityData.YListKeys = []string {}

    return &(nsf.EntityData)
}

// Igmp_DefaultContext_UnicastQosAdjust
// Configure IGMP QoS shapers for subscriber
// interfaces
type Igmp_DefaultContext_UnicastQosAdjust struct {
    EntityData types.CommonEntityData
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

func (unicastQosAdjust *Igmp_DefaultContext_UnicastQosAdjust) GetEntityData() *types.CommonEntityData {
    unicastQosAdjust.EntityData.YFilter = unicastQosAdjust.YFilter
    unicastQosAdjust.EntityData.YangName = "unicast-qos-adjust"
    unicastQosAdjust.EntityData.BundleName = "cisco_ios_xr"
    unicastQosAdjust.EntityData.ParentYangName = "default-context"
    unicastQosAdjust.EntityData.SegmentPath = "unicast-qos-adjust"
    unicastQosAdjust.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unicastQosAdjust.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unicastQosAdjust.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unicastQosAdjust.EntityData.Children = types.NewOrderedMap()
    unicastQosAdjust.EntityData.Leafs = types.NewOrderedMap()
    unicastQosAdjust.EntityData.Leafs.Append("download-interval", types.YLeaf{"DownloadInterval", unicastQosAdjust.DownloadInterval})
    unicastQosAdjust.EntityData.Leafs.Append("adjustment-delay", types.YLeaf{"AdjustmentDelay", unicastQosAdjust.AdjustmentDelay})
    unicastQosAdjust.EntityData.Leafs.Append("hold-off", types.YLeaf{"HoldOff", unicastQosAdjust.HoldOff})

    unicastQosAdjust.EntityData.YListKeys = []string {}

    return &(unicastQosAdjust.EntityData)
}

// Igmp_DefaultContext_Accounting
// Configure IGMP accounting for logging
// join/leave records
type Igmp_DefaultContext_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure IGMP accounting Maximum History setting. The type is interface{}
    // with range: 1..365. Units are day. The default value is 1.
    MaxHistory interface{}
}

func (accounting *Igmp_DefaultContext_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "default-context"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Leafs = types.NewOrderedMap()
    accounting.EntityData.Leafs.Append("max-history", types.YLeaf{"MaxHistory", accounting.MaxHistory})

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// Igmp_DefaultContext_Traffic
// Configure IGMP Traffic variables
type Igmp_DefaultContext_Traffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the route-policy profile. The type is string with length: 1..64.
    Profile interface{}
}

func (traffic *Igmp_DefaultContext_Traffic) GetEntityData() *types.CommonEntityData {
    traffic.EntityData.YFilter = traffic.YFilter
    traffic.EntityData.YangName = "traffic"
    traffic.EntityData.BundleName = "cisco_ios_xr"
    traffic.EntityData.ParentYangName = "default-context"
    traffic.EntityData.SegmentPath = "traffic"
    traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traffic.EntityData.Children = types.NewOrderedMap()
    traffic.EntityData.Leafs = types.NewOrderedMap()
    traffic.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", traffic.Profile})

    traffic.EntityData.YListKeys = []string {}

    return &(traffic.EntityData)
}

// Igmp_DefaultContext_InheritableDefaults
// Inheritable Defaults
type Igmp_DefaultContext_InheritableDefaults struct {
    EntityData types.CommonEntityData
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

func (inheritableDefaults *Igmp_DefaultContext_InheritableDefaults) GetEntityData() *types.CommonEntityData {
    inheritableDefaults.EntityData.YFilter = inheritableDefaults.YFilter
    inheritableDefaults.EntityData.YangName = "inheritable-defaults"
    inheritableDefaults.EntityData.BundleName = "cisco_ios_xr"
    inheritableDefaults.EntityData.ParentYangName = "default-context"
    inheritableDefaults.EntityData.SegmentPath = "inheritable-defaults"
    inheritableDefaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritableDefaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritableDefaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritableDefaults.EntityData.Children = types.NewOrderedMap()
    inheritableDefaults.EntityData.Children.Append("maximum-groups-per-interface-oor", types.YChild{"MaximumGroupsPerInterfaceOor", &inheritableDefaults.MaximumGroupsPerInterfaceOor})
    inheritableDefaults.EntityData.Children.Append("explicit-tracking", types.YChild{"ExplicitTracking", &inheritableDefaults.ExplicitTracking})
    inheritableDefaults.EntityData.Leafs = types.NewOrderedMap()
    inheritableDefaults.EntityData.Leafs.Append("query-timeout", types.YLeaf{"QueryTimeout", inheritableDefaults.QueryTimeout})
    inheritableDefaults.EntityData.Leafs.Append("access-group", types.YLeaf{"AccessGroup", inheritableDefaults.AccessGroup})
    inheritableDefaults.EntityData.Leafs.Append("query-max-response-time", types.YLeaf{"QueryMaxResponseTime", inheritableDefaults.QueryMaxResponseTime})
    inheritableDefaults.EntityData.Leafs.Append("version", types.YLeaf{"Version", inheritableDefaults.Version})
    inheritableDefaults.EntityData.Leafs.Append("router-enable", types.YLeaf{"RouterEnable", inheritableDefaults.RouterEnable})
    inheritableDefaults.EntityData.Leafs.Append("query-interval", types.YLeaf{"QueryInterval", inheritableDefaults.QueryInterval})

    inheritableDefaults.EntityData.YListKeys = []string {}

    return &(inheritableDefaults.EntityData)
}

// Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetEntityData() *types.CommonEntityData {
    maximumGroupsPerInterfaceOor.EntityData.YFilter = maximumGroupsPerInterfaceOor.YFilter
    maximumGroupsPerInterfaceOor.EntityData.YangName = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.BundleName = "cisco_ios_xr"
    maximumGroupsPerInterfaceOor.EntityData.ParentYangName = "inheritable-defaults"
    maximumGroupsPerInterfaceOor.EntityData.SegmentPath = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumGroupsPerInterfaceOor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumGroupsPerInterfaceOor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumGroupsPerInterfaceOor.EntityData.Children = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("maximum-groups", types.YLeaf{"MaximumGroups", maximumGroupsPerInterfaceOor.MaximumGroups})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", maximumGroupsPerInterfaceOor.WarningThreshold})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", maximumGroupsPerInterfaceOor.AccessListName})

    maximumGroupsPerInterfaceOor.EntityData.YListKeys = []string {}

    return &(maximumGroupsPerInterfaceOor.EntityData)
}

// Igmp_DefaultContext_InheritableDefaults_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Igmp_DefaultContext_InheritableDefaults_ExplicitTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Igmp_DefaultContext_InheritableDefaults_ExplicitTracking) GetEntityData() *types.CommonEntityData {
    explicitTracking.EntityData.YFilter = explicitTracking.YFilter
    explicitTracking.EntityData.YangName = "explicit-tracking"
    explicitTracking.EntityData.BundleName = "cisco_ios_xr"
    explicitTracking.EntityData.ParentYangName = "inheritable-defaults"
    explicitTracking.EntityData.SegmentPath = "explicit-tracking"
    explicitTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitTracking.EntityData.Children = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", explicitTracking.Enable})
    explicitTracking.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", explicitTracking.AccessListName})

    explicitTracking.EntityData.YListKeys = []string {}

    return &(explicitTracking.EntityData)
}

// Igmp_DefaultContext_SsmAccessGroups
// IGMP Source specific mode
type Igmp_DefaultContext_SsmAccessGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SSM static Access Group. The type is slice of
    // Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup.
    SsmAccessGroup []*Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup
}

func (ssmAccessGroups *Igmp_DefaultContext_SsmAccessGroups) GetEntityData() *types.CommonEntityData {
    ssmAccessGroups.EntityData.YFilter = ssmAccessGroups.YFilter
    ssmAccessGroups.EntityData.YangName = "ssm-access-groups"
    ssmAccessGroups.EntityData.BundleName = "cisco_ios_xr"
    ssmAccessGroups.EntityData.ParentYangName = "default-context"
    ssmAccessGroups.EntityData.SegmentPath = "ssm-access-groups"
    ssmAccessGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssmAccessGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssmAccessGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssmAccessGroups.EntityData.Children = types.NewOrderedMap()
    ssmAccessGroups.EntityData.Children.Append("ssm-access-group", types.YChild{"SsmAccessGroup", nil})
    for i := range ssmAccessGroups.SsmAccessGroup {
        ssmAccessGroups.EntityData.Children.Append(types.GetSegmentPath(ssmAccessGroups.SsmAccessGroup[i]), types.YChild{"SsmAccessGroup", ssmAccessGroups.SsmAccessGroup[i]})
    }
    ssmAccessGroups.EntityData.Leafs = types.NewOrderedMap()

    ssmAccessGroups.EntityData.YListKeys = []string {}

    return &(ssmAccessGroups.EntityData)
}

// Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup
// SSM static Access Group
type Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup struct {
    EntityData types.CommonEntityData
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

func (ssmAccessGroup *Igmp_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetEntityData() *types.CommonEntityData {
    ssmAccessGroup.EntityData.YFilter = ssmAccessGroup.YFilter
    ssmAccessGroup.EntityData.YangName = "ssm-access-group"
    ssmAccessGroup.EntityData.BundleName = "cisco_ios_xr"
    ssmAccessGroup.EntityData.ParentYangName = "ssm-access-groups"
    ssmAccessGroup.EntityData.SegmentPath = "ssm-access-group" + types.AddKeyToken(ssmAccessGroup.SourceAddress, "source-address")
    ssmAccessGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssmAccessGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssmAccessGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssmAccessGroup.EntityData.Children = types.NewOrderedMap()
    ssmAccessGroup.EntityData.Leafs = types.NewOrderedMap()
    ssmAccessGroup.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", ssmAccessGroup.SourceAddress})
    ssmAccessGroup.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", ssmAccessGroup.AccessListName})

    ssmAccessGroup.EntityData.YListKeys = []string {"SourceAddress"}

    return &(ssmAccessGroup.EntityData)
}

// Igmp_DefaultContext_Maximum
// Configure IGMP State Limits
type Igmp_DefaultContext_Maximum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure maximum number of groups accepted by this router. The type is
    // interface{} with range: 1..75000. The default value is 50000.
    MaximumGroups interface{}
}

func (maximum *Igmp_DefaultContext_Maximum) GetEntityData() *types.CommonEntityData {
    maximum.EntityData.YFilter = maximum.YFilter
    maximum.EntityData.YangName = "maximum"
    maximum.EntityData.BundleName = "cisco_ios_xr"
    maximum.EntityData.ParentYangName = "default-context"
    maximum.EntityData.SegmentPath = "maximum"
    maximum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximum.EntityData.Children = types.NewOrderedMap()
    maximum.EntityData.Leafs = types.NewOrderedMap()
    maximum.EntityData.Leafs.Append("maximum-groups", types.YLeaf{"MaximumGroups", maximum.MaximumGroups})

    maximum.EntityData.YListKeys = []string {}

    return &(maximum.EntityData)
}

// Igmp_DefaultContext_Interfaces
// Interface-level configuration
type Igmp_DefaultContext_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Igmp_DefaultContext_Interfaces_Interface.
    Interface []*Igmp_DefaultContext_Interfaces_Interface
}

func (interfaces *Igmp_DefaultContext_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "default-context"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Igmp_DefaultContext_Interfaces_Interface
// The name of the interface
type Igmp_DefaultContext_Interfaces_Interface struct {
    EntityData types.CommonEntityData
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

func (self *Igmp_DefaultContext_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("join-groups", types.YChild{"JoinGroups", &self.JoinGroups})
    self.EntityData.Children.Append("static-group-group-addresses", types.YChild{"StaticGroupGroupAddresses", &self.StaticGroupGroupAddresses})
    self.EntityData.Children.Append("maximum-groups-per-interface-oor", types.YChild{"MaximumGroupsPerInterfaceOor", &self.MaximumGroupsPerInterfaceOor})
    self.EntityData.Children.Append("explicit-tracking", types.YChild{"ExplicitTracking", &self.ExplicitTracking})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("query-timeout", types.YLeaf{"QueryTimeout", self.QueryTimeout})
    self.EntityData.Leafs.Append("access-group", types.YLeaf{"AccessGroup", self.AccessGroup})
    self.EntityData.Leafs.Append("query-max-response-time", types.YLeaf{"QueryMaxResponseTime", self.QueryMaxResponseTime})
    self.EntityData.Leafs.Append("version", types.YLeaf{"Version", self.Version})
    self.EntityData.Leafs.Append("router-enable", types.YLeaf{"RouterEnable", self.RouterEnable})
    self.EntityData.Leafs.Append("query-interval", types.YLeaf{"QueryInterval", self.QueryInterval})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Igmp_DefaultContext_Interfaces_Interface_JoinGroups
// IGMP join multicast group
// This type is a presence type.
type Igmp_DefaultContext_Interfaces_Interface_JoinGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IP group address and optional source address to include. The type is slice
    // of Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup.
    JoinGroup []*Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress.
    JoinGroupSourceAddress []*Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
}

func (joinGroups *Igmp_DefaultContext_Interfaces_Interface_JoinGroups) GetEntityData() *types.CommonEntityData {
    joinGroups.EntityData.YFilter = joinGroups.YFilter
    joinGroups.EntityData.YangName = "join-groups"
    joinGroups.EntityData.BundleName = "cisco_ios_xr"
    joinGroups.EntityData.ParentYangName = "interface"
    joinGroups.EntityData.SegmentPath = "join-groups"
    joinGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    joinGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    joinGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    joinGroups.EntityData.Children = types.NewOrderedMap()
    joinGroups.EntityData.Children.Append("join-group", types.YChild{"JoinGroup", nil})
    for i := range joinGroups.JoinGroup {
        joinGroups.EntityData.Children.Append(types.GetSegmentPath(joinGroups.JoinGroup[i]), types.YChild{"JoinGroup", joinGroups.JoinGroup[i]})
    }
    joinGroups.EntityData.Children.Append("join-group-source-address", types.YChild{"JoinGroupSourceAddress", nil})
    for i := range joinGroups.JoinGroupSourceAddress {
        joinGroups.EntityData.Children.Append(types.GetSegmentPath(joinGroups.JoinGroupSourceAddress[i]), types.YChild{"JoinGroupSourceAddress", joinGroups.JoinGroupSourceAddress[i]})
    }
    joinGroups.EntityData.Leafs = types.NewOrderedMap()

    joinGroups.EntityData.YListKeys = []string {}

    return &(joinGroups.EntityData)
}

// Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup struct {
    EntityData types.CommonEntityData
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

func (joinGroup *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetEntityData() *types.CommonEntityData {
    joinGroup.EntityData.YFilter = joinGroup.YFilter
    joinGroup.EntityData.YangName = "join-group"
    joinGroup.EntityData.BundleName = "cisco_ios_xr"
    joinGroup.EntityData.ParentYangName = "join-groups"
    joinGroup.EntityData.SegmentPath = "join-group" + types.AddKeyToken(joinGroup.GroupAddress, "group-address")
    joinGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    joinGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    joinGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    joinGroup.EntityData.Children = types.NewOrderedMap()
    joinGroup.EntityData.Leafs = types.NewOrderedMap()
    joinGroup.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", joinGroup.GroupAddress})
    joinGroup.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", joinGroup.Mode})

    joinGroup.EntityData.YListKeys = []string {"GroupAddress"}

    return &(joinGroup.EntityData)
}

// Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress struct {
    EntityData types.CommonEntityData
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

func (joinGroupSourceAddress *Igmp_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetEntityData() *types.CommonEntityData {
    joinGroupSourceAddress.EntityData.YFilter = joinGroupSourceAddress.YFilter
    joinGroupSourceAddress.EntityData.YangName = "join-group-source-address"
    joinGroupSourceAddress.EntityData.BundleName = "cisco_ios_xr"
    joinGroupSourceAddress.EntityData.ParentYangName = "join-groups"
    joinGroupSourceAddress.EntityData.SegmentPath = "join-group-source-address" + types.AddKeyToken(joinGroupSourceAddress.GroupAddress, "group-address") + types.AddKeyToken(joinGroupSourceAddress.SourceAddress, "source-address")
    joinGroupSourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    joinGroupSourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    joinGroupSourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    joinGroupSourceAddress.EntityData.Children = types.NewOrderedMap()
    joinGroupSourceAddress.EntityData.Leafs = types.NewOrderedMap()
    joinGroupSourceAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", joinGroupSourceAddress.GroupAddress})
    joinGroupSourceAddress.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", joinGroupSourceAddress.SourceAddress})
    joinGroupSourceAddress.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", joinGroupSourceAddress.Mode})

    joinGroupSourceAddress.EntityData.YListKeys = []string {"GroupAddress", "SourceAddress"}

    return &(joinGroupSourceAddress.EntityData)
}

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses
// IGMP static multicast group
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress.
    StaticGroupGroupAddress []*Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress.
    StaticGroupGroupAddressSourceAddress []*Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask.
    StaticGroupGroupAddressSourceAddressSourceAddressMask []*Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask.
    StaticGroupGroupAddressGroupAddressMask []*Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress.
    StaticGroupGroupAddressGroupAddressMaskSourceAddress []*Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.
    StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask []*Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
}

func (staticGroupGroupAddresses *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddresses.EntityData.YFilter = staticGroupGroupAddresses.YFilter
    staticGroupGroupAddresses.EntityData.YangName = "static-group-group-addresses"
    staticGroupGroupAddresses.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddresses.EntityData.ParentYangName = "interface"
    staticGroupGroupAddresses.EntityData.SegmentPath = "static-group-group-addresses"
    staticGroupGroupAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddresses.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address", types.YChild{"StaticGroupGroupAddress", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddress {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddress[i]), types.YChild{"StaticGroupGroupAddress", staticGroupGroupAddresses.StaticGroupGroupAddress[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-source-address", types.YChild{"StaticGroupGroupAddressSourceAddress", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i]), types.YChild{"StaticGroupGroupAddressSourceAddress", staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-source-address-source-address-mask", types.YChild{"StaticGroupGroupAddressSourceAddressSourceAddressMask", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i]), types.YChild{"StaticGroupGroupAddressSourceAddressSourceAddressMask", staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-group-address-mask", types.YChild{"StaticGroupGroupAddressGroupAddressMask", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i]), types.YChild{"StaticGroupGroupAddressGroupAddressMask", staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-group-address-mask-source-address", types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddress", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i]), types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddress", staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-group-address-mask-source-address-source-address-mask", types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i]), types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask", staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i]})
    }
    staticGroupGroupAddresses.EntityData.Leafs = types.NewOrderedMap()

    staticGroupGroupAddresses.EntityData.YListKeys = []string {}

    return &(staticGroupGroupAddresses.EntityData)
}

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddress.EntityData.YFilter = staticGroupGroupAddress.YFilter
    staticGroupGroupAddress.EntityData.YangName = "static-group-group-address"
    staticGroupGroupAddress.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddress.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddress.EntityData.SegmentPath = "static-group-group-address" + types.AddKeyToken(staticGroupGroupAddress.GroupAddress, "group-address")
    staticGroupGroupAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddress.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddress.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddress.GroupAddress})
    staticGroupGroupAddress.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddress.GroupCount})
    staticGroupGroupAddress.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddress.SourceCount})
    staticGroupGroupAddress.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddress.SuppressReport})

    staticGroupGroupAddress.EntityData.YListKeys = []string {"GroupAddress"}

    return &(staticGroupGroupAddress.EntityData)
}

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressSourceAddress.EntityData.YFilter = staticGroupGroupAddressSourceAddress.YFilter
    staticGroupGroupAddressSourceAddress.EntityData.YangName = "static-group-group-address-source-address"
    staticGroupGroupAddressSourceAddress.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressSourceAddress.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressSourceAddress.EntityData.SegmentPath = "static-group-group-address-source-address" + types.AddKeyToken(staticGroupGroupAddressSourceAddress.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressSourceAddress.SourceAddress, "source-address")
    staticGroupGroupAddressSourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressSourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressSourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressSourceAddress.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddress.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressSourceAddress.GroupAddress})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressSourceAddress.SourceAddress})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressSourceAddress.GroupCount})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressSourceAddress.SourceCount})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressSourceAddress.SuppressReport})

    staticGroupGroupAddressSourceAddress.EntityData.YListKeys = []string {"GroupAddress", "SourceAddress"}

    return &(staticGroupGroupAddressSourceAddress.EntityData)
}

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.YFilter = staticGroupGroupAddressSourceAddressSourceAddressMask.YFilter
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.YangName = "static-group-group-address-source-address-source-address-mask"
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.SegmentPath = "static-group-group-address-source-address-source-address-mask" + types.AddKeyToken(staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress, "source-address") + types.AddKeyToken(staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask, "source-address-mask")
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address-mask", types.YLeaf{"SourceAddressMask", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressSourceAddressSourceAddressMask.GroupCount})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceCount})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressSourceAddressSourceAddressMask.SuppressReport})

    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.YListKeys = []string {"GroupAddress", "SourceAddress", "SourceAddressMask"}

    return &(staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData)
}

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressGroupAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressGroupAddressMask.EntityData.YFilter = staticGroupGroupAddressGroupAddressMask.YFilter
    staticGroupGroupAddressGroupAddressMask.EntityData.YangName = "static-group-group-address-group-address-mask"
    staticGroupGroupAddressGroupAddressMask.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressGroupAddressMask.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressGroupAddressMask.EntityData.SegmentPath = "static-group-group-address-group-address-mask" + types.AddKeyToken(staticGroupGroupAddressGroupAddressMask.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMask.GroupAddressMask, "group-address-mask")
    staticGroupGroupAddressGroupAddressMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressGroupAddressMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressGroupAddressMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressGroupAddressMask.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressGroupAddressMask.GroupAddress})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("group-address-mask", types.YLeaf{"GroupAddressMask", staticGroupGroupAddressGroupAddressMask.GroupAddressMask})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressGroupAddressMask.GroupCount})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressGroupAddressMask.SourceCount})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressGroupAddressMask.SuppressReport})

    staticGroupGroupAddressGroupAddressMask.EntityData.YListKeys = []string {"GroupAddress", "GroupAddressMask"}

    return &(staticGroupGroupAddressGroupAddressMask.EntityData)
}

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.YFilter = staticGroupGroupAddressGroupAddressMaskSourceAddress.YFilter
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.YangName = "static-group-group-address-group-address-mask-source-address"
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.SegmentPath = "static-group-group-address-group-address-mask-source-address" + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask, "group-address-mask") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress, "source-address")
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("group-address-mask", types.YLeaf{"GroupAddressMask", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressGroupAddressMaskSourceAddress.SuppressReport})

    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.YListKeys = []string {"GroupAddress", "GroupAddressMask", "SourceAddress"}

    return &(staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData)
}

// Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Igmp_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.YFilter = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.YFilter
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.YangName = "static-group-group-address-group-address-mask-source-address-source-address-mask"
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.SegmentPath = "static-group-group-address-group-address-mask-source-address-source-address-mask" + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask, "group-address-mask") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress, "source-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask, "source-address-mask")
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-address-mask", types.YLeaf{"GroupAddressMask", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address-mask", types.YLeaf{"SourceAddressMask", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SuppressReport})

    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.YListKeys = []string {"GroupAddress", "GroupAddressMask", "SourceAddress", "SourceAddressMask"}

    return &(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData)
}

// Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (maximumGroupsPerInterfaceOor *Igmp_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetEntityData() *types.CommonEntityData {
    maximumGroupsPerInterfaceOor.EntityData.YFilter = maximumGroupsPerInterfaceOor.YFilter
    maximumGroupsPerInterfaceOor.EntityData.YangName = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.BundleName = "cisco_ios_xr"
    maximumGroupsPerInterfaceOor.EntityData.ParentYangName = "interface"
    maximumGroupsPerInterfaceOor.EntityData.SegmentPath = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumGroupsPerInterfaceOor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumGroupsPerInterfaceOor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumGroupsPerInterfaceOor.EntityData.Children = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("maximum-groups", types.YLeaf{"MaximumGroups", maximumGroupsPerInterfaceOor.MaximumGroups})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", maximumGroupsPerInterfaceOor.WarningThreshold})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", maximumGroupsPerInterfaceOor.AccessListName})

    maximumGroupsPerInterfaceOor.EntityData.YListKeys = []string {}

    return &(maximumGroupsPerInterfaceOor.EntityData)
}

// Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Igmp_DefaultContext_Interfaces_Interface_ExplicitTracking) GetEntityData() *types.CommonEntityData {
    explicitTracking.EntityData.YFilter = explicitTracking.YFilter
    explicitTracking.EntityData.YangName = "explicit-tracking"
    explicitTracking.EntityData.BundleName = "cisco_ios_xr"
    explicitTracking.EntityData.ParentYangName = "interface"
    explicitTracking.EntityData.SegmentPath = "explicit-tracking"
    explicitTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitTracking.EntityData.Children = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", explicitTracking.Enable})
    explicitTracking.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", explicitTracking.AccessListName})

    explicitTracking.EntityData.YListKeys = []string {}

    return &(explicitTracking.EntityData)
}

// Amt
// amt
type Amt struct {
    EntityData types.CommonEntityData
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

func (amt *Amt) GetEntityData() *types.CommonEntityData {
    amt.EntityData.YFilter = amt.YFilter
    amt.EntityData.YangName = "amt"
    amt.EntityData.BundleName = "cisco_ios_xr"
    amt.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-igmp-cfg"
    amt.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-igmp-cfg:amt"
    amt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    amt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    amt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    amt.EntityData.Children = types.NewOrderedMap()
    amt.EntityData.Children.Append("relay-adv-add", types.YChild{"RelayAdvAdd", &amt.RelayAdvAdd})
    amt.EntityData.Children.Append("relay-anycast-prefix", types.YChild{"RelayAnycastPrefix", &amt.RelayAnycastPrefix})
    amt.EntityData.Leafs = types.NewOrderedMap()
    amt.EntityData.Leafs.Append("maximum-v4-route-gateway", types.YLeaf{"MaximumV4RouteGateway", amt.MaximumV4RouteGateway})
    amt.EntityData.Leafs.Append("gateway-filter", types.YLeaf{"GatewayFilter", amt.GatewayFilter})
    amt.EntityData.Leafs.Append("maximum-v4-routes", types.YLeaf{"MaximumV4Routes", amt.MaximumV4Routes})
    amt.EntityData.Leafs.Append("amttos", types.YLeaf{"Amttos", amt.Amttos})
    amt.EntityData.Leafs.Append("amtttl", types.YLeaf{"Amtttl", amt.Amtttl})
    amt.EntityData.Leafs.Append("maximum-v6-route-gateway", types.YLeaf{"MaximumV6RouteGateway", amt.MaximumV6RouteGateway})
    amt.EntityData.Leafs.Append("maximum-gateway", types.YLeaf{"MaximumGateway", amt.MaximumGateway})
    amt.EntityData.Leafs.Append("maximum-v6-routes", types.YLeaf{"MaximumV6Routes", amt.MaximumV6Routes})
    amt.EntityData.Leafs.Append("amtqqic", types.YLeaf{"Amtqqic", amt.Amtqqic})
    amt.EntityData.Leafs.Append("amtmtu", types.YLeaf{"Amtmtu", amt.Amtmtu})

    amt.EntityData.YListKeys = []string {}

    return &(amt.EntityData)
}

// Amt_RelayAdvAdd
// Configure AMT Relay IPv4 Advertisement Address
// This type is a presence type.
type Amt_RelayAdvAdd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // AMT Relay IPv4 Advertisement Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}

    // Relay Advertisement Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}
}

func (relayAdvAdd *Amt_RelayAdvAdd) GetEntityData() *types.CommonEntityData {
    relayAdvAdd.EntityData.YFilter = relayAdvAdd.YFilter
    relayAdvAdd.EntityData.YangName = "relay-adv-add"
    relayAdvAdd.EntityData.BundleName = "cisco_ios_xr"
    relayAdvAdd.EntityData.ParentYangName = "amt"
    relayAdvAdd.EntityData.SegmentPath = "relay-adv-add"
    relayAdvAdd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayAdvAdd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayAdvAdd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayAdvAdd.EntityData.Children = types.NewOrderedMap()
    relayAdvAdd.EntityData.Leafs = types.NewOrderedMap()
    relayAdvAdd.EntityData.Leafs.Append("address", types.YLeaf{"Address", relayAdvAdd.Address})
    relayAdvAdd.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", relayAdvAdd.Interface})

    relayAdvAdd.EntityData.YListKeys = []string {}

    return &(relayAdvAdd.EntityData)
}

// Amt_RelayAnycastPrefix
// Configure AMT Relay IPv4 Anycast-Prefix
// This type is a presence type.
type Amt_RelayAnycastPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Anycast-Prefix Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}

    // Mask Length for Anycast Prefix. The type is interface{} with range: 1..32.
    PrefixLength interface{}
}

func (relayAnycastPrefix *Amt_RelayAnycastPrefix) GetEntityData() *types.CommonEntityData {
    relayAnycastPrefix.EntityData.YFilter = relayAnycastPrefix.YFilter
    relayAnycastPrefix.EntityData.YangName = "relay-anycast-prefix"
    relayAnycastPrefix.EntityData.BundleName = "cisco_ios_xr"
    relayAnycastPrefix.EntityData.ParentYangName = "amt"
    relayAnycastPrefix.EntityData.SegmentPath = "relay-anycast-prefix"
    relayAnycastPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayAnycastPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayAnycastPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayAnycastPrefix.EntityData.Children = types.NewOrderedMap()
    relayAnycastPrefix.EntityData.Leafs = types.NewOrderedMap()
    relayAnycastPrefix.EntityData.Leafs.Append("address", types.YLeaf{"Address", relayAnycastPrefix.Address})
    relayAnycastPrefix.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", relayAnycastPrefix.PrefixLength})

    relayAnycastPrefix.EntityData.YListKeys = []string {}

    return &(relayAnycastPrefix.EntityData)
}

// Mld
// mld
type Mld struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF related configuration.
    Vrfs Mld_Vrfs

    // Default Context.
    DefaultContext Mld_DefaultContext
}

func (mld *Mld) GetEntityData() *types.CommonEntityData {
    mld.EntityData.YFilter = mld.YFilter
    mld.EntityData.YangName = "mld"
    mld.EntityData.BundleName = "cisco_ios_xr"
    mld.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-igmp-cfg"
    mld.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-igmp-cfg:mld"
    mld.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mld.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mld.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mld.EntityData.Children = types.NewOrderedMap()
    mld.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &mld.Vrfs})
    mld.EntityData.Children.Append("default-context", types.YChild{"DefaultContext", &mld.DefaultContext})
    mld.EntityData.Leafs = types.NewOrderedMap()

    mld.EntityData.YListKeys = []string {}

    return &(mld.EntityData)
}

// Mld_Vrfs
// VRF related configuration
type Mld_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a particular vrf. The type is slice of Mld_Vrfs_Vrf.
    Vrf []*Mld_Vrfs_Vrf
}

func (vrfs *Mld_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "mld"
    vrfs.EntityData.SegmentPath = "vrfs"
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

// Mld_Vrfs_Vrf
// Configuration for a particular vrf
type Mld_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
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

func (vrf *Mld_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("traffic", types.YChild{"Traffic", &vrf.Traffic})
    vrf.EntityData.Children.Append("inheritable-defaults", types.YChild{"InheritableDefaults", &vrf.InheritableDefaults})
    vrf.EntityData.Children.Append("ssm-access-groups", types.YChild{"SsmAccessGroups", &vrf.SsmAccessGroups})
    vrf.EntityData.Children.Append("maximum", types.YChild{"Maximum", &vrf.Maximum})
    vrf.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &vrf.Interfaces})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("ssmdns-query-group", types.YLeaf{"SsmdnsQueryGroup", vrf.SsmdnsQueryGroup})
    vrf.EntityData.Leafs.Append("robustness", types.YLeaf{"Robustness", vrf.Robustness})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Mld_Vrfs_Vrf_Traffic
// Configure IGMP Traffic variables
type Mld_Vrfs_Vrf_Traffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the route-policy profile. The type is string with length: 1..64.
    Profile interface{}
}

func (traffic *Mld_Vrfs_Vrf_Traffic) GetEntityData() *types.CommonEntityData {
    traffic.EntityData.YFilter = traffic.YFilter
    traffic.EntityData.YangName = "traffic"
    traffic.EntityData.BundleName = "cisco_ios_xr"
    traffic.EntityData.ParentYangName = "vrf"
    traffic.EntityData.SegmentPath = "traffic"
    traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traffic.EntityData.Children = types.NewOrderedMap()
    traffic.EntityData.Leafs = types.NewOrderedMap()
    traffic.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", traffic.Profile})

    traffic.EntityData.YListKeys = []string {}

    return &(traffic.EntityData)
}

// Mld_Vrfs_Vrf_InheritableDefaults
// Inheritable Defaults
type Mld_Vrfs_Vrf_InheritableDefaults struct {
    EntityData types.CommonEntityData
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

func (inheritableDefaults *Mld_Vrfs_Vrf_InheritableDefaults) GetEntityData() *types.CommonEntityData {
    inheritableDefaults.EntityData.YFilter = inheritableDefaults.YFilter
    inheritableDefaults.EntityData.YangName = "inheritable-defaults"
    inheritableDefaults.EntityData.BundleName = "cisco_ios_xr"
    inheritableDefaults.EntityData.ParentYangName = "vrf"
    inheritableDefaults.EntityData.SegmentPath = "inheritable-defaults"
    inheritableDefaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritableDefaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritableDefaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritableDefaults.EntityData.Children = types.NewOrderedMap()
    inheritableDefaults.EntityData.Children.Append("maximum-groups-per-interface-oor", types.YChild{"MaximumGroupsPerInterfaceOor", &inheritableDefaults.MaximumGroupsPerInterfaceOor})
    inheritableDefaults.EntityData.Children.Append("explicit-tracking", types.YChild{"ExplicitTracking", &inheritableDefaults.ExplicitTracking})
    inheritableDefaults.EntityData.Leafs = types.NewOrderedMap()
    inheritableDefaults.EntityData.Leafs.Append("query-timeout", types.YLeaf{"QueryTimeout", inheritableDefaults.QueryTimeout})
    inheritableDefaults.EntityData.Leafs.Append("access-group", types.YLeaf{"AccessGroup", inheritableDefaults.AccessGroup})
    inheritableDefaults.EntityData.Leafs.Append("query-max-response-time", types.YLeaf{"QueryMaxResponseTime", inheritableDefaults.QueryMaxResponseTime})
    inheritableDefaults.EntityData.Leafs.Append("version", types.YLeaf{"Version", inheritableDefaults.Version})
    inheritableDefaults.EntityData.Leafs.Append("router-enable", types.YLeaf{"RouterEnable", inheritableDefaults.RouterEnable})
    inheritableDefaults.EntityData.Leafs.Append("query-interval", types.YLeaf{"QueryInterval", inheritableDefaults.QueryInterval})

    inheritableDefaults.EntityData.YListKeys = []string {}

    return &(inheritableDefaults.EntityData)
}

// Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetEntityData() *types.CommonEntityData {
    maximumGroupsPerInterfaceOor.EntityData.YFilter = maximumGroupsPerInterfaceOor.YFilter
    maximumGroupsPerInterfaceOor.EntityData.YangName = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.BundleName = "cisco_ios_xr"
    maximumGroupsPerInterfaceOor.EntityData.ParentYangName = "inheritable-defaults"
    maximumGroupsPerInterfaceOor.EntityData.SegmentPath = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumGroupsPerInterfaceOor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumGroupsPerInterfaceOor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumGroupsPerInterfaceOor.EntityData.Children = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("maximum-groups", types.YLeaf{"MaximumGroups", maximumGroupsPerInterfaceOor.MaximumGroups})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", maximumGroupsPerInterfaceOor.WarningThreshold})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", maximumGroupsPerInterfaceOor.AccessListName})

    maximumGroupsPerInterfaceOor.EntityData.YListKeys = []string {}

    return &(maximumGroupsPerInterfaceOor.EntityData)
}

// Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Mld_Vrfs_Vrf_InheritableDefaults_ExplicitTracking) GetEntityData() *types.CommonEntityData {
    explicitTracking.EntityData.YFilter = explicitTracking.YFilter
    explicitTracking.EntityData.YangName = "explicit-tracking"
    explicitTracking.EntityData.BundleName = "cisco_ios_xr"
    explicitTracking.EntityData.ParentYangName = "inheritable-defaults"
    explicitTracking.EntityData.SegmentPath = "explicit-tracking"
    explicitTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitTracking.EntityData.Children = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", explicitTracking.Enable})
    explicitTracking.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", explicitTracking.AccessListName})

    explicitTracking.EntityData.YListKeys = []string {}

    return &(explicitTracking.EntityData)
}

// Mld_Vrfs_Vrf_SsmAccessGroups
// IGMP Source specific mode
type Mld_Vrfs_Vrf_SsmAccessGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SSM static Access Group. The type is slice of
    // Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup.
    SsmAccessGroup []*Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup
}

func (ssmAccessGroups *Mld_Vrfs_Vrf_SsmAccessGroups) GetEntityData() *types.CommonEntityData {
    ssmAccessGroups.EntityData.YFilter = ssmAccessGroups.YFilter
    ssmAccessGroups.EntityData.YangName = "ssm-access-groups"
    ssmAccessGroups.EntityData.BundleName = "cisco_ios_xr"
    ssmAccessGroups.EntityData.ParentYangName = "vrf"
    ssmAccessGroups.EntityData.SegmentPath = "ssm-access-groups"
    ssmAccessGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssmAccessGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssmAccessGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssmAccessGroups.EntityData.Children = types.NewOrderedMap()
    ssmAccessGroups.EntityData.Children.Append("ssm-access-group", types.YChild{"SsmAccessGroup", nil})
    for i := range ssmAccessGroups.SsmAccessGroup {
        ssmAccessGroups.EntityData.Children.Append(types.GetSegmentPath(ssmAccessGroups.SsmAccessGroup[i]), types.YChild{"SsmAccessGroup", ssmAccessGroups.SsmAccessGroup[i]})
    }
    ssmAccessGroups.EntityData.Leafs = types.NewOrderedMap()

    ssmAccessGroups.EntityData.YListKeys = []string {}

    return &(ssmAccessGroups.EntityData)
}

// Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup
// SSM static Access Group
type Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup struct {
    EntityData types.CommonEntityData
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

func (ssmAccessGroup *Mld_Vrfs_Vrf_SsmAccessGroups_SsmAccessGroup) GetEntityData() *types.CommonEntityData {
    ssmAccessGroup.EntityData.YFilter = ssmAccessGroup.YFilter
    ssmAccessGroup.EntityData.YangName = "ssm-access-group"
    ssmAccessGroup.EntityData.BundleName = "cisco_ios_xr"
    ssmAccessGroup.EntityData.ParentYangName = "ssm-access-groups"
    ssmAccessGroup.EntityData.SegmentPath = "ssm-access-group" + types.AddKeyToken(ssmAccessGroup.SourceAddress, "source-address")
    ssmAccessGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssmAccessGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssmAccessGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssmAccessGroup.EntityData.Children = types.NewOrderedMap()
    ssmAccessGroup.EntityData.Leafs = types.NewOrderedMap()
    ssmAccessGroup.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", ssmAccessGroup.SourceAddress})
    ssmAccessGroup.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", ssmAccessGroup.AccessListName})

    ssmAccessGroup.EntityData.YListKeys = []string {"SourceAddress"}

    return &(ssmAccessGroup.EntityData)
}

// Mld_Vrfs_Vrf_Maximum
// Configure IGMP State Limits
type Mld_Vrfs_Vrf_Maximum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure maximum number of groups accepted by this router. The type is
    // interface{} with range: 1..75000. The default value is 50000.
    MaximumGroups interface{}
}

func (maximum *Mld_Vrfs_Vrf_Maximum) GetEntityData() *types.CommonEntityData {
    maximum.EntityData.YFilter = maximum.YFilter
    maximum.EntityData.YangName = "maximum"
    maximum.EntityData.BundleName = "cisco_ios_xr"
    maximum.EntityData.ParentYangName = "vrf"
    maximum.EntityData.SegmentPath = "maximum"
    maximum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximum.EntityData.Children = types.NewOrderedMap()
    maximum.EntityData.Leafs = types.NewOrderedMap()
    maximum.EntityData.Leafs.Append("maximum-groups", types.YLeaf{"MaximumGroups", maximum.MaximumGroups})

    maximum.EntityData.YListKeys = []string {}

    return &(maximum.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces
// Interface-level configuration
type Mld_Vrfs_Vrf_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Mld_Vrfs_Vrf_Interfaces_Interface.
    Interface []*Mld_Vrfs_Vrf_Interfaces_Interface
}

func (interfaces *Mld_Vrfs_Vrf_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "vrf"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces_Interface
// The name of the interface
type Mld_Vrfs_Vrf_Interfaces_Interface struct {
    EntityData types.CommonEntityData
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

func (self *Mld_Vrfs_Vrf_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("join-groups", types.YChild{"JoinGroups", &self.JoinGroups})
    self.EntityData.Children.Append("static-group-group-addresses", types.YChild{"StaticGroupGroupAddresses", &self.StaticGroupGroupAddresses})
    self.EntityData.Children.Append("maximum-groups-per-interface-oor", types.YChild{"MaximumGroupsPerInterfaceOor", &self.MaximumGroupsPerInterfaceOor})
    self.EntityData.Children.Append("explicit-tracking", types.YChild{"ExplicitTracking", &self.ExplicitTracking})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("query-timeout", types.YLeaf{"QueryTimeout", self.QueryTimeout})
    self.EntityData.Leafs.Append("access-group", types.YLeaf{"AccessGroup", self.AccessGroup})
    self.EntityData.Leafs.Append("query-max-response-time", types.YLeaf{"QueryMaxResponseTime", self.QueryMaxResponseTime})
    self.EntityData.Leafs.Append("version", types.YLeaf{"Version", self.Version})
    self.EntityData.Leafs.Append("router-enable", types.YLeaf{"RouterEnable", self.RouterEnable})
    self.EntityData.Leafs.Append("query-interval", types.YLeaf{"QueryInterval", self.QueryInterval})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups
// IGMP join multicast group
// This type is a presence type.
type Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IP group address and optional source address to include. The type is slice
    // of Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup.
    JoinGroup []*Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup

    // IP group address and optional source address to include. The type is slice
    // of Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress.
    JoinGroupSourceAddress []*Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
}

func (joinGroups *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups) GetEntityData() *types.CommonEntityData {
    joinGroups.EntityData.YFilter = joinGroups.YFilter
    joinGroups.EntityData.YangName = "join-groups"
    joinGroups.EntityData.BundleName = "cisco_ios_xr"
    joinGroups.EntityData.ParentYangName = "interface"
    joinGroups.EntityData.SegmentPath = "join-groups"
    joinGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    joinGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    joinGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    joinGroups.EntityData.Children = types.NewOrderedMap()
    joinGroups.EntityData.Children.Append("join-group", types.YChild{"JoinGroup", nil})
    for i := range joinGroups.JoinGroup {
        joinGroups.EntityData.Children.Append(types.GetSegmentPath(joinGroups.JoinGroup[i]), types.YChild{"JoinGroup", joinGroups.JoinGroup[i]})
    }
    joinGroups.EntityData.Children.Append("join-group-source-address", types.YChild{"JoinGroupSourceAddress", nil})
    for i := range joinGroups.JoinGroupSourceAddress {
        joinGroups.EntityData.Children.Append(types.GetSegmentPath(joinGroups.JoinGroupSourceAddress[i]), types.YChild{"JoinGroupSourceAddress", joinGroups.JoinGroupSourceAddress[i]})
    }
    joinGroups.EntityData.Leafs = types.NewOrderedMap()

    joinGroups.EntityData.YListKeys = []string {}

    return &(joinGroups.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup struct {
    EntityData types.CommonEntityData
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

func (joinGroup *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroup) GetEntityData() *types.CommonEntityData {
    joinGroup.EntityData.YFilter = joinGroup.YFilter
    joinGroup.EntityData.YangName = "join-group"
    joinGroup.EntityData.BundleName = "cisco_ios_xr"
    joinGroup.EntityData.ParentYangName = "join-groups"
    joinGroup.EntityData.SegmentPath = "join-group" + types.AddKeyToken(joinGroup.GroupAddress, "group-address")
    joinGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    joinGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    joinGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    joinGroup.EntityData.Children = types.NewOrderedMap()
    joinGroup.EntityData.Leafs = types.NewOrderedMap()
    joinGroup.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", joinGroup.GroupAddress})
    joinGroup.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", joinGroup.Mode})

    joinGroup.EntityData.YListKeys = []string {"GroupAddress"}

    return &(joinGroup.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress struct {
    EntityData types.CommonEntityData
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

func (joinGroupSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetEntityData() *types.CommonEntityData {
    joinGroupSourceAddress.EntityData.YFilter = joinGroupSourceAddress.YFilter
    joinGroupSourceAddress.EntityData.YangName = "join-group-source-address"
    joinGroupSourceAddress.EntityData.BundleName = "cisco_ios_xr"
    joinGroupSourceAddress.EntityData.ParentYangName = "join-groups"
    joinGroupSourceAddress.EntityData.SegmentPath = "join-group-source-address" + types.AddKeyToken(joinGroupSourceAddress.GroupAddress, "group-address") + types.AddKeyToken(joinGroupSourceAddress.SourceAddress, "source-address")
    joinGroupSourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    joinGroupSourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    joinGroupSourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    joinGroupSourceAddress.EntityData.Children = types.NewOrderedMap()
    joinGroupSourceAddress.EntityData.Leafs = types.NewOrderedMap()
    joinGroupSourceAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", joinGroupSourceAddress.GroupAddress})
    joinGroupSourceAddress.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", joinGroupSourceAddress.SourceAddress})
    joinGroupSourceAddress.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", joinGroupSourceAddress.Mode})

    joinGroupSourceAddress.EntityData.YListKeys = []string {"GroupAddress", "SourceAddress"}

    return &(joinGroupSourceAddress.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses
// IGMP static multicast group
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress.
    StaticGroupGroupAddress []*Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress.
    StaticGroupGroupAddressSourceAddress []*Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask.
    StaticGroupGroupAddressSourceAddressSourceAddressMask []*Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask.
    StaticGroupGroupAddressGroupAddressMask []*Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress.
    StaticGroupGroupAddressGroupAddressMaskSourceAddress []*Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.
    StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask []*Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
}

func (staticGroupGroupAddresses *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddresses.EntityData.YFilter = staticGroupGroupAddresses.YFilter
    staticGroupGroupAddresses.EntityData.YangName = "static-group-group-addresses"
    staticGroupGroupAddresses.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddresses.EntityData.ParentYangName = "interface"
    staticGroupGroupAddresses.EntityData.SegmentPath = "static-group-group-addresses"
    staticGroupGroupAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddresses.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address", types.YChild{"StaticGroupGroupAddress", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddress {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddress[i]), types.YChild{"StaticGroupGroupAddress", staticGroupGroupAddresses.StaticGroupGroupAddress[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-source-address", types.YChild{"StaticGroupGroupAddressSourceAddress", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i]), types.YChild{"StaticGroupGroupAddressSourceAddress", staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-source-address-source-address-mask", types.YChild{"StaticGroupGroupAddressSourceAddressSourceAddressMask", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i]), types.YChild{"StaticGroupGroupAddressSourceAddressSourceAddressMask", staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-group-address-mask", types.YChild{"StaticGroupGroupAddressGroupAddressMask", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i]), types.YChild{"StaticGroupGroupAddressGroupAddressMask", staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-group-address-mask-source-address", types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddress", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i]), types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddress", staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-group-address-mask-source-address-source-address-mask", types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i]), types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask", staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i]})
    }
    staticGroupGroupAddresses.EntityData.Leafs = types.NewOrderedMap()

    staticGroupGroupAddresses.EntityData.YListKeys = []string {}

    return &(staticGroupGroupAddresses.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddress.EntityData.YFilter = staticGroupGroupAddress.YFilter
    staticGroupGroupAddress.EntityData.YangName = "static-group-group-address"
    staticGroupGroupAddress.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddress.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddress.EntityData.SegmentPath = "static-group-group-address" + types.AddKeyToken(staticGroupGroupAddress.GroupAddress, "group-address")
    staticGroupGroupAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddress.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddress.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddress.GroupAddress})
    staticGroupGroupAddress.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddress.GroupCount})
    staticGroupGroupAddress.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddress.SourceCount})
    staticGroupGroupAddress.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddress.SuppressReport})

    staticGroupGroupAddress.EntityData.YListKeys = []string {"GroupAddress"}

    return &(staticGroupGroupAddress.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressSourceAddress.EntityData.YFilter = staticGroupGroupAddressSourceAddress.YFilter
    staticGroupGroupAddressSourceAddress.EntityData.YangName = "static-group-group-address-source-address"
    staticGroupGroupAddressSourceAddress.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressSourceAddress.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressSourceAddress.EntityData.SegmentPath = "static-group-group-address-source-address" + types.AddKeyToken(staticGroupGroupAddressSourceAddress.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressSourceAddress.SourceAddress, "source-address")
    staticGroupGroupAddressSourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressSourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressSourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressSourceAddress.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddress.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressSourceAddress.GroupAddress})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressSourceAddress.SourceAddress})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressSourceAddress.GroupCount})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressSourceAddress.SourceCount})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressSourceAddress.SuppressReport})

    staticGroupGroupAddressSourceAddress.EntityData.YListKeys = []string {"GroupAddress", "SourceAddress"}

    return &(staticGroupGroupAddressSourceAddress.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.YFilter = staticGroupGroupAddressSourceAddressSourceAddressMask.YFilter
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.YangName = "static-group-group-address-source-address-source-address-mask"
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.SegmentPath = "static-group-group-address-source-address-source-address-mask" + types.AddKeyToken(staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress, "source-address") + types.AddKeyToken(staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask, "source-address-mask")
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address-mask", types.YLeaf{"SourceAddressMask", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressSourceAddressSourceAddressMask.GroupCount})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceCount})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressSourceAddressSourceAddressMask.SuppressReport})

    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.YListKeys = []string {"GroupAddress", "SourceAddress", "SourceAddressMask"}

    return &(staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressGroupAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressGroupAddressMask.EntityData.YFilter = staticGroupGroupAddressGroupAddressMask.YFilter
    staticGroupGroupAddressGroupAddressMask.EntityData.YangName = "static-group-group-address-group-address-mask"
    staticGroupGroupAddressGroupAddressMask.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressGroupAddressMask.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressGroupAddressMask.EntityData.SegmentPath = "static-group-group-address-group-address-mask" + types.AddKeyToken(staticGroupGroupAddressGroupAddressMask.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMask.GroupAddressMask, "group-address-mask")
    staticGroupGroupAddressGroupAddressMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressGroupAddressMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressGroupAddressMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressGroupAddressMask.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressGroupAddressMask.GroupAddress})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("group-address-mask", types.YLeaf{"GroupAddressMask", staticGroupGroupAddressGroupAddressMask.GroupAddressMask})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressGroupAddressMask.GroupCount})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressGroupAddressMask.SourceCount})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressGroupAddressMask.SuppressReport})

    staticGroupGroupAddressGroupAddressMask.EntityData.YListKeys = []string {"GroupAddress", "GroupAddressMask"}

    return &(staticGroupGroupAddressGroupAddressMask.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.YFilter = staticGroupGroupAddressGroupAddressMaskSourceAddress.YFilter
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.YangName = "static-group-group-address-group-address-mask-source-address"
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.SegmentPath = "static-group-group-address-group-address-mask-source-address" + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask, "group-address-mask") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress, "source-address")
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("group-address-mask", types.YLeaf{"GroupAddressMask", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressGroupAddressMaskSourceAddress.SuppressReport})

    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.YListKeys = []string {"GroupAddress", "GroupAddressMask", "SourceAddress"}

    return &(staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_Vrfs_Vrf_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.YFilter = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.YFilter
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.YangName = "static-group-group-address-group-address-mask-source-address-source-address-mask"
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.SegmentPath = "static-group-group-address-group-address-mask-source-address-source-address-mask" + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask, "group-address-mask") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress, "source-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask, "source-address-mask")
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-address-mask", types.YLeaf{"GroupAddressMask", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address-mask", types.YLeaf{"SourceAddressMask", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SuppressReport})

    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.YListKeys = []string {"GroupAddress", "GroupAddressMask", "SourceAddress", "SourceAddressMask"}

    return &(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (maximumGroupsPerInterfaceOor *Mld_Vrfs_Vrf_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetEntityData() *types.CommonEntityData {
    maximumGroupsPerInterfaceOor.EntityData.YFilter = maximumGroupsPerInterfaceOor.YFilter
    maximumGroupsPerInterfaceOor.EntityData.YangName = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.BundleName = "cisco_ios_xr"
    maximumGroupsPerInterfaceOor.EntityData.ParentYangName = "interface"
    maximumGroupsPerInterfaceOor.EntityData.SegmentPath = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumGroupsPerInterfaceOor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumGroupsPerInterfaceOor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumGroupsPerInterfaceOor.EntityData.Children = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("maximum-groups", types.YLeaf{"MaximumGroups", maximumGroupsPerInterfaceOor.MaximumGroups})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", maximumGroupsPerInterfaceOor.WarningThreshold})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", maximumGroupsPerInterfaceOor.AccessListName})

    maximumGroupsPerInterfaceOor.EntityData.YListKeys = []string {}

    return &(maximumGroupsPerInterfaceOor.EntityData)
}

// Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Mld_Vrfs_Vrf_Interfaces_Interface_ExplicitTracking) GetEntityData() *types.CommonEntityData {
    explicitTracking.EntityData.YFilter = explicitTracking.YFilter
    explicitTracking.EntityData.YangName = "explicit-tracking"
    explicitTracking.EntityData.BundleName = "cisco_ios_xr"
    explicitTracking.EntityData.ParentYangName = "interface"
    explicitTracking.EntityData.SegmentPath = "explicit-tracking"
    explicitTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitTracking.EntityData.Children = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", explicitTracking.Enable})
    explicitTracking.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", explicitTracking.AccessListName})

    explicitTracking.EntityData.YListKeys = []string {}

    return &(explicitTracking.EntityData)
}

// Mld_DefaultContext
// Default Context
// This type is a presence type.
type Mld_DefaultContext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (defaultContext *Mld_DefaultContext) GetEntityData() *types.CommonEntityData {
    defaultContext.EntityData.YFilter = defaultContext.YFilter
    defaultContext.EntityData.YangName = "default-context"
    defaultContext.EntityData.BundleName = "cisco_ios_xr"
    defaultContext.EntityData.ParentYangName = "mld"
    defaultContext.EntityData.SegmentPath = "default-context"
    defaultContext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultContext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultContext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultContext.EntityData.Children = types.NewOrderedMap()
    defaultContext.EntityData.Children.Append("nsf", types.YChild{"Nsf", &defaultContext.Nsf})
    defaultContext.EntityData.Children.Append("unicast-qos-adjust", types.YChild{"UnicastQosAdjust", &defaultContext.UnicastQosAdjust})
    defaultContext.EntityData.Children.Append("accounting", types.YChild{"Accounting", &defaultContext.Accounting})
    defaultContext.EntityData.Children.Append("traffic", types.YChild{"Traffic", &defaultContext.Traffic})
    defaultContext.EntityData.Children.Append("inheritable-defaults", types.YChild{"InheritableDefaults", &defaultContext.InheritableDefaults})
    defaultContext.EntityData.Children.Append("ssm-access-groups", types.YChild{"SsmAccessGroups", &defaultContext.SsmAccessGroups})
    defaultContext.EntityData.Children.Append("maximum", types.YChild{"Maximum", &defaultContext.Maximum})
    defaultContext.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &defaultContext.Interfaces})
    defaultContext.EntityData.Leafs = types.NewOrderedMap()
    defaultContext.EntityData.Leafs.Append("ssmdns-query-group", types.YLeaf{"SsmdnsQueryGroup", defaultContext.SsmdnsQueryGroup})
    defaultContext.EntityData.Leafs.Append("robustness", types.YLeaf{"Robustness", defaultContext.Robustness})

    defaultContext.EntityData.YListKeys = []string {}

    return &(defaultContext.EntityData)
}

// Mld_DefaultContext_Nsf
// Configure NSF specific options
type Mld_DefaultContext_Nsf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum time for IGMP NSF mode in seconds. The type is interface{} with
    // range: 10..3600. Units are second. The default value is 60.
    Lifetime interface{}
}

func (nsf *Mld_DefaultContext_Nsf) GetEntityData() *types.CommonEntityData {
    nsf.EntityData.YFilter = nsf.YFilter
    nsf.EntityData.YangName = "nsf"
    nsf.EntityData.BundleName = "cisco_ios_xr"
    nsf.EntityData.ParentYangName = "default-context"
    nsf.EntityData.SegmentPath = "nsf"
    nsf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nsf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nsf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nsf.EntityData.Children = types.NewOrderedMap()
    nsf.EntityData.Leafs = types.NewOrderedMap()
    nsf.EntityData.Leafs.Append("lifetime", types.YLeaf{"Lifetime", nsf.Lifetime})

    nsf.EntityData.YListKeys = []string {}

    return &(nsf.EntityData)
}

// Mld_DefaultContext_UnicastQosAdjust
// Configure IGMP QoS shapers for subscriber
// interfaces
type Mld_DefaultContext_UnicastQosAdjust struct {
    EntityData types.CommonEntityData
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

func (unicastQosAdjust *Mld_DefaultContext_UnicastQosAdjust) GetEntityData() *types.CommonEntityData {
    unicastQosAdjust.EntityData.YFilter = unicastQosAdjust.YFilter
    unicastQosAdjust.EntityData.YangName = "unicast-qos-adjust"
    unicastQosAdjust.EntityData.BundleName = "cisco_ios_xr"
    unicastQosAdjust.EntityData.ParentYangName = "default-context"
    unicastQosAdjust.EntityData.SegmentPath = "unicast-qos-adjust"
    unicastQosAdjust.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unicastQosAdjust.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unicastQosAdjust.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unicastQosAdjust.EntityData.Children = types.NewOrderedMap()
    unicastQosAdjust.EntityData.Leafs = types.NewOrderedMap()
    unicastQosAdjust.EntityData.Leafs.Append("download-interval", types.YLeaf{"DownloadInterval", unicastQosAdjust.DownloadInterval})
    unicastQosAdjust.EntityData.Leafs.Append("adjustment-delay", types.YLeaf{"AdjustmentDelay", unicastQosAdjust.AdjustmentDelay})
    unicastQosAdjust.EntityData.Leafs.Append("hold-off", types.YLeaf{"HoldOff", unicastQosAdjust.HoldOff})

    unicastQosAdjust.EntityData.YListKeys = []string {}

    return &(unicastQosAdjust.EntityData)
}

// Mld_DefaultContext_Accounting
// Configure IGMP accounting for logging
// join/leave records
type Mld_DefaultContext_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure IGMP accounting Maximum History setting. The type is interface{}
    // with range: 1..365. Units are day. The default value is 1.
    MaxHistory interface{}
}

func (accounting *Mld_DefaultContext_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "default-context"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Leafs = types.NewOrderedMap()
    accounting.EntityData.Leafs.Append("max-history", types.YLeaf{"MaxHistory", accounting.MaxHistory})

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// Mld_DefaultContext_Traffic
// Configure IGMP Traffic variables
type Mld_DefaultContext_Traffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the route-policy profile. The type is string with length: 1..64.
    Profile interface{}
}

func (traffic *Mld_DefaultContext_Traffic) GetEntityData() *types.CommonEntityData {
    traffic.EntityData.YFilter = traffic.YFilter
    traffic.EntityData.YangName = "traffic"
    traffic.EntityData.BundleName = "cisco_ios_xr"
    traffic.EntityData.ParentYangName = "default-context"
    traffic.EntityData.SegmentPath = "traffic"
    traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traffic.EntityData.Children = types.NewOrderedMap()
    traffic.EntityData.Leafs = types.NewOrderedMap()
    traffic.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", traffic.Profile})

    traffic.EntityData.YListKeys = []string {}

    return &(traffic.EntityData)
}

// Mld_DefaultContext_InheritableDefaults
// Inheritable Defaults
type Mld_DefaultContext_InheritableDefaults struct {
    EntityData types.CommonEntityData
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

func (inheritableDefaults *Mld_DefaultContext_InheritableDefaults) GetEntityData() *types.CommonEntityData {
    inheritableDefaults.EntityData.YFilter = inheritableDefaults.YFilter
    inheritableDefaults.EntityData.YangName = "inheritable-defaults"
    inheritableDefaults.EntityData.BundleName = "cisco_ios_xr"
    inheritableDefaults.EntityData.ParentYangName = "default-context"
    inheritableDefaults.EntityData.SegmentPath = "inheritable-defaults"
    inheritableDefaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inheritableDefaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inheritableDefaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inheritableDefaults.EntityData.Children = types.NewOrderedMap()
    inheritableDefaults.EntityData.Children.Append("maximum-groups-per-interface-oor", types.YChild{"MaximumGroupsPerInterfaceOor", &inheritableDefaults.MaximumGroupsPerInterfaceOor})
    inheritableDefaults.EntityData.Children.Append("explicit-tracking", types.YChild{"ExplicitTracking", &inheritableDefaults.ExplicitTracking})
    inheritableDefaults.EntityData.Leafs = types.NewOrderedMap()
    inheritableDefaults.EntityData.Leafs.Append("query-timeout", types.YLeaf{"QueryTimeout", inheritableDefaults.QueryTimeout})
    inheritableDefaults.EntityData.Leafs.Append("access-group", types.YLeaf{"AccessGroup", inheritableDefaults.AccessGroup})
    inheritableDefaults.EntityData.Leafs.Append("query-max-response-time", types.YLeaf{"QueryMaxResponseTime", inheritableDefaults.QueryMaxResponseTime})
    inheritableDefaults.EntityData.Leafs.Append("version", types.YLeaf{"Version", inheritableDefaults.Version})
    inheritableDefaults.EntityData.Leafs.Append("router-enable", types.YLeaf{"RouterEnable", inheritableDefaults.RouterEnable})
    inheritableDefaults.EntityData.Leafs.Append("query-interval", types.YLeaf{"QueryInterval", inheritableDefaults.QueryInterval})

    inheritableDefaults.EntityData.YListKeys = []string {}

    return &(inheritableDefaults.EntityData)
}

// Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_InheritableDefaults_MaximumGroupsPerInterfaceOor) GetEntityData() *types.CommonEntityData {
    maximumGroupsPerInterfaceOor.EntityData.YFilter = maximumGroupsPerInterfaceOor.YFilter
    maximumGroupsPerInterfaceOor.EntityData.YangName = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.BundleName = "cisco_ios_xr"
    maximumGroupsPerInterfaceOor.EntityData.ParentYangName = "inheritable-defaults"
    maximumGroupsPerInterfaceOor.EntityData.SegmentPath = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumGroupsPerInterfaceOor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumGroupsPerInterfaceOor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumGroupsPerInterfaceOor.EntityData.Children = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("maximum-groups", types.YLeaf{"MaximumGroups", maximumGroupsPerInterfaceOor.MaximumGroups})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", maximumGroupsPerInterfaceOor.WarningThreshold})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", maximumGroupsPerInterfaceOor.AccessListName})

    maximumGroupsPerInterfaceOor.EntityData.YListKeys = []string {}

    return &(maximumGroupsPerInterfaceOor.EntityData)
}

// Mld_DefaultContext_InheritableDefaults_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Mld_DefaultContext_InheritableDefaults_ExplicitTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Mld_DefaultContext_InheritableDefaults_ExplicitTracking) GetEntityData() *types.CommonEntityData {
    explicitTracking.EntityData.YFilter = explicitTracking.YFilter
    explicitTracking.EntityData.YangName = "explicit-tracking"
    explicitTracking.EntityData.BundleName = "cisco_ios_xr"
    explicitTracking.EntityData.ParentYangName = "inheritable-defaults"
    explicitTracking.EntityData.SegmentPath = "explicit-tracking"
    explicitTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitTracking.EntityData.Children = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", explicitTracking.Enable})
    explicitTracking.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", explicitTracking.AccessListName})

    explicitTracking.EntityData.YListKeys = []string {}

    return &(explicitTracking.EntityData)
}

// Mld_DefaultContext_SsmAccessGroups
// IGMP Source specific mode
type Mld_DefaultContext_SsmAccessGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SSM static Access Group. The type is slice of
    // Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup.
    SsmAccessGroup []*Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup
}

func (ssmAccessGroups *Mld_DefaultContext_SsmAccessGroups) GetEntityData() *types.CommonEntityData {
    ssmAccessGroups.EntityData.YFilter = ssmAccessGroups.YFilter
    ssmAccessGroups.EntityData.YangName = "ssm-access-groups"
    ssmAccessGroups.EntityData.BundleName = "cisco_ios_xr"
    ssmAccessGroups.EntityData.ParentYangName = "default-context"
    ssmAccessGroups.EntityData.SegmentPath = "ssm-access-groups"
    ssmAccessGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssmAccessGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssmAccessGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssmAccessGroups.EntityData.Children = types.NewOrderedMap()
    ssmAccessGroups.EntityData.Children.Append("ssm-access-group", types.YChild{"SsmAccessGroup", nil})
    for i := range ssmAccessGroups.SsmAccessGroup {
        ssmAccessGroups.EntityData.Children.Append(types.GetSegmentPath(ssmAccessGroups.SsmAccessGroup[i]), types.YChild{"SsmAccessGroup", ssmAccessGroups.SsmAccessGroup[i]})
    }
    ssmAccessGroups.EntityData.Leafs = types.NewOrderedMap()

    ssmAccessGroups.EntityData.YListKeys = []string {}

    return &(ssmAccessGroups.EntityData)
}

// Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup
// SSM static Access Group
type Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup struct {
    EntityData types.CommonEntityData
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

func (ssmAccessGroup *Mld_DefaultContext_SsmAccessGroups_SsmAccessGroup) GetEntityData() *types.CommonEntityData {
    ssmAccessGroup.EntityData.YFilter = ssmAccessGroup.YFilter
    ssmAccessGroup.EntityData.YangName = "ssm-access-group"
    ssmAccessGroup.EntityData.BundleName = "cisco_ios_xr"
    ssmAccessGroup.EntityData.ParentYangName = "ssm-access-groups"
    ssmAccessGroup.EntityData.SegmentPath = "ssm-access-group" + types.AddKeyToken(ssmAccessGroup.SourceAddress, "source-address")
    ssmAccessGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssmAccessGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssmAccessGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssmAccessGroup.EntityData.Children = types.NewOrderedMap()
    ssmAccessGroup.EntityData.Leafs = types.NewOrderedMap()
    ssmAccessGroup.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", ssmAccessGroup.SourceAddress})
    ssmAccessGroup.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", ssmAccessGroup.AccessListName})

    ssmAccessGroup.EntityData.YListKeys = []string {"SourceAddress"}

    return &(ssmAccessGroup.EntityData)
}

// Mld_DefaultContext_Maximum
// Configure IGMP State Limits
type Mld_DefaultContext_Maximum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure maximum number of groups accepted by this router. The type is
    // interface{} with range: 1..75000. The default value is 50000.
    MaximumGroups interface{}
}

func (maximum *Mld_DefaultContext_Maximum) GetEntityData() *types.CommonEntityData {
    maximum.EntityData.YFilter = maximum.YFilter
    maximum.EntityData.YangName = "maximum"
    maximum.EntityData.BundleName = "cisco_ios_xr"
    maximum.EntityData.ParentYangName = "default-context"
    maximum.EntityData.SegmentPath = "maximum"
    maximum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximum.EntityData.Children = types.NewOrderedMap()
    maximum.EntityData.Leafs = types.NewOrderedMap()
    maximum.EntityData.Leafs.Append("maximum-groups", types.YLeaf{"MaximumGroups", maximum.MaximumGroups})

    maximum.EntityData.YListKeys = []string {}

    return &(maximum.EntityData)
}

// Mld_DefaultContext_Interfaces
// Interface-level configuration
type Mld_DefaultContext_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Mld_DefaultContext_Interfaces_Interface.
    Interface []*Mld_DefaultContext_Interfaces_Interface
}

func (interfaces *Mld_DefaultContext_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "default-context"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Mld_DefaultContext_Interfaces_Interface
// The name of the interface
type Mld_DefaultContext_Interfaces_Interface struct {
    EntityData types.CommonEntityData
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

func (self *Mld_DefaultContext_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("join-groups", types.YChild{"JoinGroups", &self.JoinGroups})
    self.EntityData.Children.Append("static-group-group-addresses", types.YChild{"StaticGroupGroupAddresses", &self.StaticGroupGroupAddresses})
    self.EntityData.Children.Append("maximum-groups-per-interface-oor", types.YChild{"MaximumGroupsPerInterfaceOor", &self.MaximumGroupsPerInterfaceOor})
    self.EntityData.Children.Append("explicit-tracking", types.YChild{"ExplicitTracking", &self.ExplicitTracking})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("query-timeout", types.YLeaf{"QueryTimeout", self.QueryTimeout})
    self.EntityData.Leafs.Append("access-group", types.YLeaf{"AccessGroup", self.AccessGroup})
    self.EntityData.Leafs.Append("query-max-response-time", types.YLeaf{"QueryMaxResponseTime", self.QueryMaxResponseTime})
    self.EntityData.Leafs.Append("version", types.YLeaf{"Version", self.Version})
    self.EntityData.Leafs.Append("router-enable", types.YLeaf{"RouterEnable", self.RouterEnable})
    self.EntityData.Leafs.Append("query-interval", types.YLeaf{"QueryInterval", self.QueryInterval})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Mld_DefaultContext_Interfaces_Interface_JoinGroups
// IGMP join multicast group
// This type is a presence type.
type Mld_DefaultContext_Interfaces_Interface_JoinGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IP group address and optional source address to include. The type is slice
    // of Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup.
    JoinGroup []*Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress.
    JoinGroupSourceAddress []*Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
}

func (joinGroups *Mld_DefaultContext_Interfaces_Interface_JoinGroups) GetEntityData() *types.CommonEntityData {
    joinGroups.EntityData.YFilter = joinGroups.YFilter
    joinGroups.EntityData.YangName = "join-groups"
    joinGroups.EntityData.BundleName = "cisco_ios_xr"
    joinGroups.EntityData.ParentYangName = "interface"
    joinGroups.EntityData.SegmentPath = "join-groups"
    joinGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    joinGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    joinGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    joinGroups.EntityData.Children = types.NewOrderedMap()
    joinGroups.EntityData.Children.Append("join-group", types.YChild{"JoinGroup", nil})
    for i := range joinGroups.JoinGroup {
        joinGroups.EntityData.Children.Append(types.GetSegmentPath(joinGroups.JoinGroup[i]), types.YChild{"JoinGroup", joinGroups.JoinGroup[i]})
    }
    joinGroups.EntityData.Children.Append("join-group-source-address", types.YChild{"JoinGroupSourceAddress", nil})
    for i := range joinGroups.JoinGroupSourceAddress {
        joinGroups.EntityData.Children.Append(types.GetSegmentPath(joinGroups.JoinGroupSourceAddress[i]), types.YChild{"JoinGroupSourceAddress", joinGroups.JoinGroupSourceAddress[i]})
    }
    joinGroups.EntityData.Leafs = types.NewOrderedMap()

    joinGroups.EntityData.YListKeys = []string {}

    return &(joinGroups.EntityData)
}

// Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup struct {
    EntityData types.CommonEntityData
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

func (joinGroup *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroup) GetEntityData() *types.CommonEntityData {
    joinGroup.EntityData.YFilter = joinGroup.YFilter
    joinGroup.EntityData.YangName = "join-group"
    joinGroup.EntityData.BundleName = "cisco_ios_xr"
    joinGroup.EntityData.ParentYangName = "join-groups"
    joinGroup.EntityData.SegmentPath = "join-group" + types.AddKeyToken(joinGroup.GroupAddress, "group-address")
    joinGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    joinGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    joinGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    joinGroup.EntityData.Children = types.NewOrderedMap()
    joinGroup.EntityData.Leafs = types.NewOrderedMap()
    joinGroup.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", joinGroup.GroupAddress})
    joinGroup.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", joinGroup.Mode})

    joinGroup.EntityData.YListKeys = []string {"GroupAddress"}

    return &(joinGroup.EntityData)
}

// Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress struct {
    EntityData types.CommonEntityData
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

func (joinGroupSourceAddress *Mld_DefaultContext_Interfaces_Interface_JoinGroups_JoinGroupSourceAddress) GetEntityData() *types.CommonEntityData {
    joinGroupSourceAddress.EntityData.YFilter = joinGroupSourceAddress.YFilter
    joinGroupSourceAddress.EntityData.YangName = "join-group-source-address"
    joinGroupSourceAddress.EntityData.BundleName = "cisco_ios_xr"
    joinGroupSourceAddress.EntityData.ParentYangName = "join-groups"
    joinGroupSourceAddress.EntityData.SegmentPath = "join-group-source-address" + types.AddKeyToken(joinGroupSourceAddress.GroupAddress, "group-address") + types.AddKeyToken(joinGroupSourceAddress.SourceAddress, "source-address")
    joinGroupSourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    joinGroupSourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    joinGroupSourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    joinGroupSourceAddress.EntityData.Children = types.NewOrderedMap()
    joinGroupSourceAddress.EntityData.Leafs = types.NewOrderedMap()
    joinGroupSourceAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", joinGroupSourceAddress.GroupAddress})
    joinGroupSourceAddress.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", joinGroupSourceAddress.SourceAddress})
    joinGroupSourceAddress.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", joinGroupSourceAddress.Mode})

    joinGroupSourceAddress.EntityData.YListKeys = []string {"GroupAddress", "SourceAddress"}

    return &(joinGroupSourceAddress.EntityData)
}

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses
// IGMP static multicast group
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress.
    StaticGroupGroupAddress []*Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress.
    StaticGroupGroupAddressSourceAddress []*Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask.
    StaticGroupGroupAddressSourceAddressSourceAddressMask []*Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask.
    StaticGroupGroupAddressGroupAddressMask []*Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress.
    StaticGroupGroupAddressGroupAddressMaskSourceAddress []*Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress

    // IP group address and optional source address to include. The type is slice
    // of
    // Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.
    StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask []*Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
}

func (staticGroupGroupAddresses *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddresses.EntityData.YFilter = staticGroupGroupAddresses.YFilter
    staticGroupGroupAddresses.EntityData.YangName = "static-group-group-addresses"
    staticGroupGroupAddresses.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddresses.EntityData.ParentYangName = "interface"
    staticGroupGroupAddresses.EntityData.SegmentPath = "static-group-group-addresses"
    staticGroupGroupAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddresses.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address", types.YChild{"StaticGroupGroupAddress", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddress {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddress[i]), types.YChild{"StaticGroupGroupAddress", staticGroupGroupAddresses.StaticGroupGroupAddress[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-source-address", types.YChild{"StaticGroupGroupAddressSourceAddress", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i]), types.YChild{"StaticGroupGroupAddressSourceAddress", staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddress[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-source-address-source-address-mask", types.YChild{"StaticGroupGroupAddressSourceAddressSourceAddressMask", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i]), types.YChild{"StaticGroupGroupAddressSourceAddressSourceAddressMask", staticGroupGroupAddresses.StaticGroupGroupAddressSourceAddressSourceAddressMask[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-group-address-mask", types.YChild{"StaticGroupGroupAddressGroupAddressMask", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i]), types.YChild{"StaticGroupGroupAddressGroupAddressMask", staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMask[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-group-address-mask-source-address", types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddress", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i]), types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddress", staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddress[i]})
    }
    staticGroupGroupAddresses.EntityData.Children.Append("static-group-group-address-group-address-mask-source-address-source-address-mask", types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask", nil})
    for i := range staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask {
        staticGroupGroupAddresses.EntityData.Children.Append(types.GetSegmentPath(staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i]), types.YChild{"StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask", staticGroupGroupAddresses.StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask[i]})
    }
    staticGroupGroupAddresses.EntityData.Leafs = types.NewOrderedMap()

    staticGroupGroupAddresses.EntityData.YListKeys = []string {}

    return &(staticGroupGroupAddresses.EntityData)
}

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddress) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddress.EntityData.YFilter = staticGroupGroupAddress.YFilter
    staticGroupGroupAddress.EntityData.YangName = "static-group-group-address"
    staticGroupGroupAddress.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddress.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddress.EntityData.SegmentPath = "static-group-group-address" + types.AddKeyToken(staticGroupGroupAddress.GroupAddress, "group-address")
    staticGroupGroupAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddress.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddress.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddress.GroupAddress})
    staticGroupGroupAddress.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddress.GroupCount})
    staticGroupGroupAddress.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddress.SourceCount})
    staticGroupGroupAddress.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddress.SuppressReport})

    staticGroupGroupAddress.EntityData.YListKeys = []string {"GroupAddress"}

    return &(staticGroupGroupAddress.EntityData)
}

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddress) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressSourceAddress.EntityData.YFilter = staticGroupGroupAddressSourceAddress.YFilter
    staticGroupGroupAddressSourceAddress.EntityData.YangName = "static-group-group-address-source-address"
    staticGroupGroupAddressSourceAddress.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressSourceAddress.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressSourceAddress.EntityData.SegmentPath = "static-group-group-address-source-address" + types.AddKeyToken(staticGroupGroupAddressSourceAddress.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressSourceAddress.SourceAddress, "source-address")
    staticGroupGroupAddressSourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressSourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressSourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressSourceAddress.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddress.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressSourceAddress.GroupAddress})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressSourceAddress.SourceAddress})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressSourceAddress.GroupCount})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressSourceAddress.SourceCount})
    staticGroupGroupAddressSourceAddress.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressSourceAddress.SuppressReport})

    staticGroupGroupAddressSourceAddress.EntityData.YListKeys = []string {"GroupAddress", "SourceAddress"}

    return &(staticGroupGroupAddressSourceAddress.EntityData)
}

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressSourceAddressSourceAddressMask) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.YFilter = staticGroupGroupAddressSourceAddressSourceAddressMask.YFilter
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.YangName = "static-group-group-address-source-address-source-address-mask"
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.SegmentPath = "static-group-group-address-source-address-source-address-mask" + types.AddKeyToken(staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress, "source-address") + types.AddKeyToken(staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask, "source-address-mask")
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressSourceAddressSourceAddressMask.GroupAddress})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddress})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address-mask", types.YLeaf{"SourceAddressMask", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceAddressMask})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressSourceAddressSourceAddressMask.GroupCount})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressSourceAddressSourceAddressMask.SourceCount})
    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressSourceAddressSourceAddressMask.SuppressReport})

    staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData.YListKeys = []string {"GroupAddress", "SourceAddress", "SourceAddressMask"}

    return &(staticGroupGroupAddressSourceAddressSourceAddressMask.EntityData)
}

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressGroupAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMask) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressGroupAddressMask.EntityData.YFilter = staticGroupGroupAddressGroupAddressMask.YFilter
    staticGroupGroupAddressGroupAddressMask.EntityData.YangName = "static-group-group-address-group-address-mask"
    staticGroupGroupAddressGroupAddressMask.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressGroupAddressMask.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressGroupAddressMask.EntityData.SegmentPath = "static-group-group-address-group-address-mask" + types.AddKeyToken(staticGroupGroupAddressGroupAddressMask.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMask.GroupAddressMask, "group-address-mask")
    staticGroupGroupAddressGroupAddressMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressGroupAddressMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressGroupAddressMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressGroupAddressMask.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressGroupAddressMask.GroupAddress})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("group-address-mask", types.YLeaf{"GroupAddressMask", staticGroupGroupAddressGroupAddressMask.GroupAddressMask})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressGroupAddressMask.GroupCount})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressGroupAddressMask.SourceCount})
    staticGroupGroupAddressGroupAddressMask.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressGroupAddressMask.SuppressReport})

    staticGroupGroupAddressGroupAddressMask.EntityData.YListKeys = []string {"GroupAddress", "GroupAddressMask"}

    return &(staticGroupGroupAddressGroupAddressMask.EntityData)
}

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressGroupAddressMaskSourceAddress *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddress) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.YFilter = staticGroupGroupAddressGroupAddressMaskSourceAddress.YFilter
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.YangName = "static-group-group-address-group-address-mask-source-address"
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.SegmentPath = "static-group-group-address-group-address-mask-source-address" + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask, "group-address-mask") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress, "source-address")
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("group-address-mask", types.YLeaf{"GroupAddressMask", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupAddressMask})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressGroupAddressMaskSourceAddress.GroupCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressGroupAddressMaskSourceAddress.SourceCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressGroupAddressMaskSourceAddress.SuppressReport})

    staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData.YListKeys = []string {"GroupAddress", "GroupAddressMask", "SourceAddress"}

    return &(staticGroupGroupAddressGroupAddressMaskSourceAddress.EntityData)
}

// Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask
// IP group address and optional source address
// to include
type Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask struct {
    EntityData types.CommonEntityData
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

func (staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask *Mld_DefaultContext_Interfaces_Interface_StaticGroupGroupAddresses_StaticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask) GetEntityData() *types.CommonEntityData {
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.YFilter = staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.YFilter
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.YangName = "static-group-group-address-group-address-mask-source-address-source-address-mask"
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.BundleName = "cisco_ios_xr"
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.ParentYangName = "static-group-group-addresses"
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.SegmentPath = "static-group-group-address-group-address-mask-source-address-source-address-mask" + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress, "group-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask, "group-address-mask") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress, "source-address") + types.AddKeyToken(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask, "source-address-mask")
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Children = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs = types.NewOrderedMap()
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-address", types.YLeaf{"GroupAddress", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-address-mask", types.YLeaf{"GroupAddressMask", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupAddressMask})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddress})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-address-mask", types.YLeaf{"SourceAddressMask", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceAddressMask})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.GroupCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("source-count", types.YLeaf{"SourceCount", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SourceCount})
    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.Leafs.Append("suppress-report", types.YLeaf{"SuppressReport", staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.SuppressReport})

    staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData.YListKeys = []string {"GroupAddress", "GroupAddressMask", "SourceAddress", "SourceAddressMask"}

    return &(staticGroupGroupAddressGroupAddressMaskSourceAddressSourceAddressMask.EntityData)
}

// Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor
// Configure maximum number of groups accepted per
// interface by this router
// This type is a presence type.
type Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (maximumGroupsPerInterfaceOor *Mld_DefaultContext_Interfaces_Interface_MaximumGroupsPerInterfaceOor) GetEntityData() *types.CommonEntityData {
    maximumGroupsPerInterfaceOor.EntityData.YFilter = maximumGroupsPerInterfaceOor.YFilter
    maximumGroupsPerInterfaceOor.EntityData.YangName = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.BundleName = "cisco_ios_xr"
    maximumGroupsPerInterfaceOor.EntityData.ParentYangName = "interface"
    maximumGroupsPerInterfaceOor.EntityData.SegmentPath = "maximum-groups-per-interface-oor"
    maximumGroupsPerInterfaceOor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximumGroupsPerInterfaceOor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximumGroupsPerInterfaceOor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximumGroupsPerInterfaceOor.EntityData.Children = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs = types.NewOrderedMap()
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("maximum-groups", types.YLeaf{"MaximumGroups", maximumGroupsPerInterfaceOor.MaximumGroups})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("warning-threshold", types.YLeaf{"WarningThreshold", maximumGroupsPerInterfaceOor.WarningThreshold})
    maximumGroupsPerInterfaceOor.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", maximumGroupsPerInterfaceOor.AccessListName})

    maximumGroupsPerInterfaceOor.EntityData.YListKeys = []string {}

    return &(maximumGroupsPerInterfaceOor.EntityData)
}

// Mld_DefaultContext_Interfaces_Interface_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type Mld_DefaultContext_Interfaces_Interface_ExplicitTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enabled or disabled, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *Mld_DefaultContext_Interfaces_Interface_ExplicitTracking) GetEntityData() *types.CommonEntityData {
    explicitTracking.EntityData.YFilter = explicitTracking.YFilter
    explicitTracking.EntityData.YangName = "explicit-tracking"
    explicitTracking.EntityData.BundleName = "cisco_ios_xr"
    explicitTracking.EntityData.ParentYangName = "interface"
    explicitTracking.EntityData.SegmentPath = "explicit-tracking"
    explicitTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitTracking.EntityData.Children = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", explicitTracking.Enable})
    explicitTracking.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", explicitTracking.AccessListName})

    explicitTracking.EntityData.YListKeys = []string {}

    return &(explicitTracking.EntityData)
}

