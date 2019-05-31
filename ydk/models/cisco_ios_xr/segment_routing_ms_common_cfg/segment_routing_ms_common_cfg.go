// This module contains a collection of YANG definitions
// for Cisco IOS-XR segment-routing-ms-common package configuration.
// 
// This module contains definitions
// for the following management objects:
//   sr: Segment Routing
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package segment_routing_ms_common_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package segment_routing_ms_common_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-segment-routing-ms-common-cfg sr}", reflect.TypeOf(Sr{}))
    ydk.RegisterEntity("Cisco-IOS-XR-segment-routing-ms-common-cfg:sr", reflect.TypeOf(Sr{}))
}

// Sr
// Segment Routing
type Sr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (sr *Sr) GetEntityData() *types.CommonEntityData {
    sr.EntityData.YFilter = sr.YFilter
    sr.EntityData.YangName = "sr"
    sr.EntityData.BundleName = "cisco_ios_xr"
    sr.EntityData.ParentYangName = "Cisco-IOS-XR-segment-routing-ms-common-cfg"
    sr.EntityData.SegmentPath = "Cisco-IOS-XR-segment-routing-ms-common-cfg:sr"
    sr.EntityData.AbsolutePath = sr.EntityData.SegmentPath
    sr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sr.EntityData.Children = types.NewOrderedMap()
    sr.EntityData.Leafs = types.NewOrderedMap()

    sr.EntityData.YListKeys = []string {}

    return &(sr.EntityData)
}

