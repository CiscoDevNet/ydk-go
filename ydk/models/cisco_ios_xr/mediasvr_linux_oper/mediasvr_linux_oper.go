// This module contains a collection of YANG definitions
// for Cisco IOS-XR mediasvr-linux package operational data.
// 
// This module contains definitions
// for the following management objects:
//   media-svr: Media server CLI operations
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package mediasvr_linux_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mediasvr_linux_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mediasvr-linux-oper media-svr}", reflect.TypeOf(MediaSvr{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mediasvr-linux-oper:media-svr", reflect.TypeOf(MediaSvr{}))
}

// MediaSvr
// Media server CLI operations
type MediaSvr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show Media bag.
    All MediaSvr_All

    // Show Media.
    LocationDescriptions MediaSvr_LocationDescriptions
}

func (mediaSvr *MediaSvr) GetEntityData() *types.CommonEntityData {
    mediaSvr.EntityData.YFilter = mediaSvr.YFilter
    mediaSvr.EntityData.YangName = "media-svr"
    mediaSvr.EntityData.BundleName = "cisco_ios_xr"
    mediaSvr.EntityData.ParentYangName = "Cisco-IOS-XR-mediasvr-linux-oper"
    mediaSvr.EntityData.SegmentPath = "Cisco-IOS-XR-mediasvr-linux-oper:media-svr"
    mediaSvr.EntityData.AbsolutePath = mediaSvr.EntityData.SegmentPath
    mediaSvr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mediaSvr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mediaSvr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mediaSvr.EntityData.Children = types.NewOrderedMap()
    mediaSvr.EntityData.Children.Append("all", types.YChild{"All", &mediaSvr.All})
    mediaSvr.EntityData.Children.Append("location-descriptions", types.YChild{"LocationDescriptions", &mediaSvr.LocationDescriptions})
    mediaSvr.EntityData.Leafs = types.NewOrderedMap()

    mediaSvr.EntityData.YListKeys = []string {}

    return &(mediaSvr.EntityData)
}

// MediaSvr_All
// Show Media bag
type MediaSvr_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // string output. The type is string.
    ShowOutput interface{}
}

func (all *MediaSvr_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "media-svr"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-mediasvr-linux-oper:media-svr/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("show-output", types.YLeaf{"ShowOutput", all.ShowOutput})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// MediaSvr_LocationDescriptions
// Show Media
type MediaSvr_LocationDescriptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Location specified in location. The type is slice of
    // MediaSvr_LocationDescriptions_LocationDescription.
    LocationDescription []*MediaSvr_LocationDescriptions_LocationDescription
}

func (locationDescriptions *MediaSvr_LocationDescriptions) GetEntityData() *types.CommonEntityData {
    locationDescriptions.EntityData.YFilter = locationDescriptions.YFilter
    locationDescriptions.EntityData.YangName = "location-descriptions"
    locationDescriptions.EntityData.BundleName = "cisco_ios_xr"
    locationDescriptions.EntityData.ParentYangName = "media-svr"
    locationDescriptions.EntityData.SegmentPath = "location-descriptions"
    locationDescriptions.EntityData.AbsolutePath = "Cisco-IOS-XR-mediasvr-linux-oper:media-svr/" + locationDescriptions.EntityData.SegmentPath
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

// MediaSvr_LocationDescriptions_LocationDescription
// Location specified in location
type MediaSvr_LocationDescriptions_LocationDescription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // string output. The type is string.
    ShowOutput interface{}
}

func (locationDescription *MediaSvr_LocationDescriptions_LocationDescription) GetEntityData() *types.CommonEntityData {
    locationDescription.EntityData.YFilter = locationDescription.YFilter
    locationDescription.EntityData.YangName = "location-description"
    locationDescription.EntityData.BundleName = "cisco_ios_xr"
    locationDescription.EntityData.ParentYangName = "location-descriptions"
    locationDescription.EntityData.SegmentPath = "location-description" + types.AddKeyToken(locationDescription.Node, "node")
    locationDescription.EntityData.AbsolutePath = "Cisco-IOS-XR-mediasvr-linux-oper:media-svr/location-descriptions/" + locationDescription.EntityData.SegmentPath
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

