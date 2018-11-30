// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-iarm package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ip-arm: IP Address Repository Manager (IPv4/IPv6 ARM)
//     configuration data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 ARM configuration.
    Ipv4 IpArm_Ipv4

    // IPv6 ARM configuration.
    Ipv6 IpArm_Ipv6
}

func (ipArm *IpArm) GetEntityData() *types.CommonEntityData {
    ipArm.EntityData.YFilter = ipArm.YFilter
    ipArm.EntityData.YangName = "ip-arm"
    ipArm.EntityData.BundleName = "cisco_ios_xr"
    ipArm.EntityData.ParentYangName = "Cisco-IOS-XR-ip-iarm-cfg"
    ipArm.EntityData.SegmentPath = "Cisco-IOS-XR-ip-iarm-cfg:ip-arm"
    ipArm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipArm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipArm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipArm.EntityData.Children = types.NewOrderedMap()
    ipArm.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &ipArm.Ipv4})
    ipArm.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &ipArm.Ipv6})
    ipArm.EntityData.Leafs = types.NewOrderedMap()

    ipArm.EntityData.YListKeys = []string {}

    return &(ipArm.EntityData)
}

// IpArm_Ipv4
// IPv4 ARM configuration
type IpArm_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP ARM conflict policy configuration.
    ConflictPolicyTable IpArm_Ipv4_ConflictPolicyTable

    // IP ARM Multicast Host configuration.
    MulticastHost IpArm_Ipv4_MulticastHost
}

func (ipv4 *IpArm_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "ip-arm"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("conflict-policy-table", types.YChild{"ConflictPolicyTable", &ipv4.ConflictPolicyTable})
    ipv4.EntityData.Children.Append("multicast-host", types.YChild{"MulticastHost", &ipv4.MulticastHost})
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// IpArm_Ipv4_ConflictPolicyTable
// IP ARM conflict policy configuration
type IpArm_Ipv4_ConflictPolicyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP ARM conflict policy value definitions. The type is IpArmConflictPolicy.
    ConflictPolicy interface{}
}

func (conflictPolicyTable *IpArm_Ipv4_ConflictPolicyTable) GetEntityData() *types.CommonEntityData {
    conflictPolicyTable.EntityData.YFilter = conflictPolicyTable.YFilter
    conflictPolicyTable.EntityData.YangName = "conflict-policy-table"
    conflictPolicyTable.EntityData.BundleName = "cisco_ios_xr"
    conflictPolicyTable.EntityData.ParentYangName = "ipv4"
    conflictPolicyTable.EntityData.SegmentPath = "conflict-policy-table"
    conflictPolicyTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conflictPolicyTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conflictPolicyTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conflictPolicyTable.EntityData.Children = types.NewOrderedMap()
    conflictPolicyTable.EntityData.Leafs = types.NewOrderedMap()
    conflictPolicyTable.EntityData.Leafs.Append("conflict-policy", types.YLeaf{"ConflictPolicy", conflictPolicyTable.ConflictPolicy})

    conflictPolicyTable.EntityData.YListKeys = []string {}

    return &(conflictPolicyTable.EntityData)
}

// IpArm_Ipv4_MulticastHost
// IP ARM Multicast Host configuration
type IpArm_Ipv4_MulticastHost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default multicast host interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    MulticastHostInterface interface{}
}

func (multicastHost *IpArm_Ipv4_MulticastHost) GetEntityData() *types.CommonEntityData {
    multicastHost.EntityData.YFilter = multicastHost.YFilter
    multicastHost.EntityData.YangName = "multicast-host"
    multicastHost.EntityData.BundleName = "cisco_ios_xr"
    multicastHost.EntityData.ParentYangName = "ipv4"
    multicastHost.EntityData.SegmentPath = "multicast-host"
    multicastHost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastHost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastHost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastHost.EntityData.Children = types.NewOrderedMap()
    multicastHost.EntityData.Leafs = types.NewOrderedMap()
    multicastHost.EntityData.Leafs.Append("multicast-host-interface", types.YLeaf{"MulticastHostInterface", multicastHost.MulticastHostInterface})

    multicastHost.EntityData.YListKeys = []string {}

    return &(multicastHost.EntityData)
}

// IpArm_Ipv6
// IPv6 ARM configuration
type IpArm_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP ARM conflict policy configuration.
    ConflictPolicyTable IpArm_Ipv6_ConflictPolicyTable

    // IP ARM Multicast Host configuration.
    MulticastHost IpArm_Ipv6_MulticastHost
}

func (ipv6 *IpArm_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "ip-arm"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("conflict-policy-table", types.YChild{"ConflictPolicyTable", &ipv6.ConflictPolicyTable})
    ipv6.EntityData.Children.Append("multicast-host", types.YChild{"MulticastHost", &ipv6.MulticastHost})
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// IpArm_Ipv6_ConflictPolicyTable
// IP ARM conflict policy configuration
type IpArm_Ipv6_ConflictPolicyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP ARM conflict policy value definitions. The type is IpArmConflictPolicy.
    ConflictPolicy interface{}
}

func (conflictPolicyTable *IpArm_Ipv6_ConflictPolicyTable) GetEntityData() *types.CommonEntityData {
    conflictPolicyTable.EntityData.YFilter = conflictPolicyTable.YFilter
    conflictPolicyTable.EntityData.YangName = "conflict-policy-table"
    conflictPolicyTable.EntityData.BundleName = "cisco_ios_xr"
    conflictPolicyTable.EntityData.ParentYangName = "ipv6"
    conflictPolicyTable.EntityData.SegmentPath = "conflict-policy-table"
    conflictPolicyTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conflictPolicyTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conflictPolicyTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conflictPolicyTable.EntityData.Children = types.NewOrderedMap()
    conflictPolicyTable.EntityData.Leafs = types.NewOrderedMap()
    conflictPolicyTable.EntityData.Leafs.Append("conflict-policy", types.YLeaf{"ConflictPolicy", conflictPolicyTable.ConflictPolicy})

    conflictPolicyTable.EntityData.YListKeys = []string {}

    return &(conflictPolicyTable.EntityData)
}

// IpArm_Ipv6_MulticastHost
// IP ARM Multicast Host configuration
type IpArm_Ipv6_MulticastHost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default multicast host interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    MulticastHostInterface interface{}
}

func (multicastHost *IpArm_Ipv6_MulticastHost) GetEntityData() *types.CommonEntityData {
    multicastHost.EntityData.YFilter = multicastHost.YFilter
    multicastHost.EntityData.YangName = "multicast-host"
    multicastHost.EntityData.BundleName = "cisco_ios_xr"
    multicastHost.EntityData.ParentYangName = "ipv6"
    multicastHost.EntityData.SegmentPath = "multicast-host"
    multicastHost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastHost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastHost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastHost.EntityData.Children = types.NewOrderedMap()
    multicastHost.EntityData.Leafs = types.NewOrderedMap()
    multicastHost.EntityData.Leafs.Append("multicast-host-interface", types.YLeaf{"MulticastHostInterface", multicastHost.MulticastHostInterface})

    multicastHost.EntityData.YListKeys = []string {}

    return &(multicastHost.EntityData)
}

