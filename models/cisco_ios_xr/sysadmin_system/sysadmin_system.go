// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module defines the system users authentication 
// credentials and virtual IP that can be modified in
// runtime. 
// 
// Copyright(c) 2011-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_system

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_system"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-system mgmt}", reflect.TypeOf(Mgmt{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-system:mgmt", reflect.TypeOf(Mgmt{}))
}

// Mgmt
type Mgmt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Ipv4 Mgmt_Ipv4

    
    Ipv6 Mgmt_Ipv6
}

func (mgmt *Mgmt) GetEntityData() *types.CommonEntityData {
    mgmt.EntityData.YFilter = mgmt.YFilter
    mgmt.EntityData.YangName = "mgmt"
    mgmt.EntityData.BundleName = "cisco_ios_xr"
    mgmt.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-system"
    mgmt.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-system:mgmt"
    mgmt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mgmt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mgmt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mgmt.EntityData.Children = types.NewOrderedMap()
    mgmt.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &mgmt.Ipv4})
    mgmt.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &mgmt.Ipv6})
    mgmt.EntityData.Leafs = types.NewOrderedMap()

    mgmt.EntityData.YListKeys = []string {}

    return &(mgmt.EntityData)
}

// Mgmt_Ipv4
type Mgmt_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(/(([0-9])|([1-2][0-9])|(3[0-2])))?.
    Address interface{}

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SubnetMaskIp interface{}
}

func (ipv4 *Mgmt_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "mgmt"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4.Address})
    ipv4.EntityData.Leafs.Append("subnet-mask-ip", types.YLeaf{"SubnetMaskIp", ipv4.SubnetMaskIp})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Mgmt_Ipv6
type Mgmt_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))?.
    Address interface{}

    // The type is interface{} with range: 0..128.
    Prefix interface{}
}

func (ipv6 *Mgmt_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "mgmt"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv6.Address})
    ipv6.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", ipv6.Prefix})

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

