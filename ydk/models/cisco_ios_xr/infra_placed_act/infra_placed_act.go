// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_placed_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_placed_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-placed-act placement-reoptimize}", reflect.TypeOf(PlacementReoptimize{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-placed-act:placement-reoptimize", reflect.TypeOf(PlacementReoptimize{}))
}

// PlacementReoptimize
// Migrate the processes, if placement is not optimal.
// 
type PlacementReoptimize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (placementReoptimize *PlacementReoptimize) GetEntityData() *types.CommonEntityData {
    placementReoptimize.EntityData.YFilter = placementReoptimize.YFilter
    placementReoptimize.EntityData.YangName = "placement-reoptimize"
    placementReoptimize.EntityData.BundleName = "cisco_ios_xr"
    placementReoptimize.EntityData.ParentYangName = "Cisco-IOS-XR-infra-placed-act"
    placementReoptimize.EntityData.SegmentPath = "Cisco-IOS-XR-infra-placed-act:placement-reoptimize"
    placementReoptimize.EntityData.AbsolutePath = placementReoptimize.EntityData.SegmentPath
    placementReoptimize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    placementReoptimize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    placementReoptimize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    placementReoptimize.EntityData.Children = types.NewOrderedMap()
    placementReoptimize.EntityData.Leafs = types.NewOrderedMap()

    placementReoptimize.EntityData.YListKeys = []string {}

    return &(placementReoptimize.EntityData)
}

