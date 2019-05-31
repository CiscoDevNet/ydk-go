// This module contains a collection of YANG definitions
// for Cisco IOS-XR linux-os-heap-summary package operational data.
// 
// This module contains definitions
// for the following management objects:
//   heap-summary: Heap Summary
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package linux_os_heap_summary_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package linux_os_heap_summary_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-linux-os-heap-summary-oper heap-summary}", reflect.TypeOf(HeapSummary{}))
    ydk.RegisterEntity("Cisco-IOS-XR-linux-os-heap-summary-oper:heap-summary", reflect.TypeOf(HeapSummary{}))
}

// HeapSummary
// Heap Summary
type HeapSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Location.
    LocationDescriptions HeapSummary_LocationDescriptions

    // All locations.
    All HeapSummary_All
}

func (heapSummary *HeapSummary) GetEntityData() *types.CommonEntityData {
    heapSummary.EntityData.YFilter = heapSummary.YFilter
    heapSummary.EntityData.YangName = "heap-summary"
    heapSummary.EntityData.BundleName = "cisco_ios_xr"
    heapSummary.EntityData.ParentYangName = "Cisco-IOS-XR-linux-os-heap-summary-oper"
    heapSummary.EntityData.SegmentPath = "Cisco-IOS-XR-linux-os-heap-summary-oper:heap-summary"
    heapSummary.EntityData.AbsolutePath = heapSummary.EntityData.SegmentPath
    heapSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    heapSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    heapSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    heapSummary.EntityData.Children = types.NewOrderedMap()
    heapSummary.EntityData.Children.Append("location-descriptions", types.YChild{"LocationDescriptions", &heapSummary.LocationDescriptions})
    heapSummary.EntityData.Children.Append("all", types.YChild{"All", &heapSummary.All})
    heapSummary.EntityData.Leafs = types.NewOrderedMap()

    heapSummary.EntityData.YListKeys = []string {}

    return &(heapSummary.EntityData)
}

// HeapSummary_LocationDescriptions
// Location
type HeapSummary_LocationDescriptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Location specified in location. The type is slice of
    // HeapSummary_LocationDescriptions_LocationDescription.
    LocationDescription []*HeapSummary_LocationDescriptions_LocationDescription
}

func (locationDescriptions *HeapSummary_LocationDescriptions) GetEntityData() *types.CommonEntityData {
    locationDescriptions.EntityData.YFilter = locationDescriptions.YFilter
    locationDescriptions.EntityData.YangName = "location-descriptions"
    locationDescriptions.EntityData.BundleName = "cisco_ios_xr"
    locationDescriptions.EntityData.ParentYangName = "heap-summary"
    locationDescriptions.EntityData.SegmentPath = "location-descriptions"
    locationDescriptions.EntityData.AbsolutePath = "Cisco-IOS-XR-linux-os-heap-summary-oper:heap-summary/" + locationDescriptions.EntityData.SegmentPath
    locationDescriptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locationDescriptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locationDescriptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locationDescriptions.EntityData.Children = types.NewOrderedMap()
    locationDescriptions.EntityData.Children.Append("location-description", types.YChild{"LocationDescription", nil})
    for i := range locationDescriptions.LocationDescription {
        locationDescriptions.EntityData.Children.Append(types.GetSegmentPath(locationDescriptions.LocationDescription[i]), types.YChild{"LocationDescription", locationDescriptions.LocationDescription[i]})
    }
    locationDescriptions.EntityData.Leafs = types.NewOrderedMap()

    locationDescriptions.EntityData.YListKeys = []string {}

    return &(locationDescriptions.EntityData)
}

// HeapSummary_LocationDescriptions_LocationDescription
// Location specified in location
type HeapSummary_LocationDescriptions_LocationDescription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node location. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // string output. The type is string.
    ShowOutput interface{}
}

func (locationDescription *HeapSummary_LocationDescriptions_LocationDescription) GetEntityData() *types.CommonEntityData {
    locationDescription.EntityData.YFilter = locationDescription.YFilter
    locationDescription.EntityData.YangName = "location-description"
    locationDescription.EntityData.BundleName = "cisco_ios_xr"
    locationDescription.EntityData.ParentYangName = "location-descriptions"
    locationDescription.EntityData.SegmentPath = "location-description" + types.AddKeyToken(locationDescription.Node, "node")
    locationDescription.EntityData.AbsolutePath = "Cisco-IOS-XR-linux-os-heap-summary-oper:heap-summary/location-descriptions/" + locationDescription.EntityData.SegmentPath
    locationDescription.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locationDescription.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locationDescription.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locationDescription.EntityData.Children = types.NewOrderedMap()
    locationDescription.EntityData.Leafs = types.NewOrderedMap()
    locationDescription.EntityData.Leafs.Append("node", types.YLeaf{"Node", locationDescription.Node})
    locationDescription.EntityData.Leafs.Append("show-output", types.YLeaf{"ShowOutput", locationDescription.ShowOutput})

    locationDescription.EntityData.YListKeys = []string {"Node"}

    return &(locationDescription.EntityData)
}

// HeapSummary_All
// All locations
type HeapSummary_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // string output. The type is string.
    ShowOutput interface{}
}

func (all *HeapSummary_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "heap-summary"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-linux-os-heap-summary-oper:heap-summary/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("show-output", types.YLeaf{"ShowOutput", all.ShowOutput})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

