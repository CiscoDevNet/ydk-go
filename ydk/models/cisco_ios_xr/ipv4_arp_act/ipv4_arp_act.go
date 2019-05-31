// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_arp_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_arp_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-arp-act clear-arp-api-stats-api}", reflect.TypeOf(ClearArpApiStatsApi{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-arp-act:clear-arp-api-stats-api", reflect.TypeOf(ClearArpApiStatsApi{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-arp-act clear-arp-api-stats-location}", reflect.TypeOf(ClearArpApiStatsLocation{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-arp-act:clear-arp-api-stats-location", reflect.TypeOf(ClearArpApiStatsLocation{}))
}

// ClearArpApiStatsApi
type ClearArpApiStatsApi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearArpApiStatsApi_Input
}

func (clearArpApiStatsApi *ClearArpApiStatsApi) GetEntityData() *types.CommonEntityData {
    clearArpApiStatsApi.EntityData.YFilter = clearArpApiStatsApi.YFilter
    clearArpApiStatsApi.EntityData.YangName = "clear-arp-api-stats-api"
    clearArpApiStatsApi.EntityData.BundleName = "cisco_ios_xr"
    clearArpApiStatsApi.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-arp-act"
    clearArpApiStatsApi.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-arp-act:clear-arp-api-stats-api"
    clearArpApiStatsApi.EntityData.AbsolutePath = clearArpApiStatsApi.EntityData.SegmentPath
    clearArpApiStatsApi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearArpApiStatsApi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearArpApiStatsApi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearArpApiStatsApi.EntityData.Children = types.NewOrderedMap()
    clearArpApiStatsApi.EntityData.Children.Append("input", types.YChild{"Input", &clearArpApiStatsApi.Input})
    clearArpApiStatsApi.EntityData.Leafs = types.NewOrderedMap()

    clearArpApiStatsApi.EntityData.YListKeys = []string {}

    return &(clearArpApiStatsApi.EntityData)
}

// ClearArpApiStatsApi_Input
type ClearArpApiStatsApi_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string. This attribute is mandatory.
    Name interface{}
}

func (input *ClearArpApiStatsApi_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-arp-api-stats-api"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-arp-act:clear-arp-api-stats-api/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("name", types.YLeaf{"Name", input.Name})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearArpApiStatsLocation
type ClearArpApiStatsLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearArpApiStatsLocation_Input
}

func (clearArpApiStatsLocation *ClearArpApiStatsLocation) GetEntityData() *types.CommonEntityData {
    clearArpApiStatsLocation.EntityData.YFilter = clearArpApiStatsLocation.YFilter
    clearArpApiStatsLocation.EntityData.YangName = "clear-arp-api-stats-location"
    clearArpApiStatsLocation.EntityData.BundleName = "cisco_ios_xr"
    clearArpApiStatsLocation.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-arp-act"
    clearArpApiStatsLocation.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-arp-act:clear-arp-api-stats-location"
    clearArpApiStatsLocation.EntityData.AbsolutePath = clearArpApiStatsLocation.EntityData.SegmentPath
    clearArpApiStatsLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearArpApiStatsLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearArpApiStatsLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearArpApiStatsLocation.EntityData.Children = types.NewOrderedMap()
    clearArpApiStatsLocation.EntityData.Children.Append("input", types.YChild{"Input", &clearArpApiStatsLocation.Input})
    clearArpApiStatsLocation.EntityData.Leafs = types.NewOrderedMap()

    clearArpApiStatsLocation.EntityData.YListKeys = []string {}

    return &(clearArpApiStatsLocation.EntityData)
}

// ClearArpApiStatsLocation_Input
type ClearArpApiStatsLocation_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string. This attribute is mandatory.
    NodeLocation interface{}
}

func (input *ClearArpApiStatsLocation_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-arp-api-stats-location"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-arp-act:clear-arp-api-stats-location/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("node-location", types.YLeaf{"NodeLocation", input.NodeLocation})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

