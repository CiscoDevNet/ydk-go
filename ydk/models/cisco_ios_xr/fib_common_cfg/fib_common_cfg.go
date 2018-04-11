// This module contains a collection of YANG definitions
// for Cisco IOS-XR fib-common package configuration.
// 
// This module contains definitions
// for the following management objects:
//   fib: CEF configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package fib_common_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package fib_common_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-fib-common-cfg fib}", reflect.TypeOf(Fib{}))
    ydk.RegisterEntity("Cisco-IOS-XR-fib-common-cfg:fib", reflect.TypeOf(Fib{}))
}

// FibPbtsForwardClass represents Fib pbts forward class
type FibPbtsForwardClass string

const (
    // Any class
    FibPbtsForwardClass_any FibPbtsForwardClass = "any"
)

// FibPbtsFallback represents Fib pbts fallback
type FibPbtsFallback string

const (
    // Fallback to class number list
    FibPbtsFallback_list FibPbtsFallback = "list"

    // Fallback to any class
    FibPbtsFallback_any FibPbtsFallback = "any"

    // Fallback to drop
    FibPbtsFallback_drop FibPbtsFallback = "drop"
)

// Fib
// CEF configuration
type Fib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set options for adjacency routes overriding RIB routes. The type is bool.
    PreferAibRoutes interface{}

    // Set option for fast-reroute to follow BGP PIC update, not to wait for
    // timeout. The type is bool.
    FrrFollowBgpPic interface{}

    // PBTS class configuration.
    PbtsForwardClassFallbacks Fib_PbtsForwardClassFallbacks

    // FIB platform parameters.
    Platform Fib_Platform
}

func (fib *Fib) GetEntityData() *types.CommonEntityData {
    fib.EntityData.YFilter = fib.YFilter
    fib.EntityData.YangName = "fib"
    fib.EntityData.BundleName = "cisco_ios_xr"
    fib.EntityData.ParentYangName = "Cisco-IOS-XR-fib-common-cfg"
    fib.EntityData.SegmentPath = "Cisco-IOS-XR-fib-common-cfg:fib"
    fib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fib.EntityData.Children = make(map[string]types.YChild)
    fib.EntityData.Children["pbts-forward-class-fallbacks"] = types.YChild{"PbtsForwardClassFallbacks", &fib.PbtsForwardClassFallbacks}
    fib.EntityData.Children["platform"] = types.YChild{"Platform", &fib.Platform}
    fib.EntityData.Leafs = make(map[string]types.YLeaf)
    fib.EntityData.Leafs["prefer-aib-routes"] = types.YLeaf{"PreferAibRoutes", fib.PreferAibRoutes}
    fib.EntityData.Leafs["frr-follow-bgp-pic"] = types.YLeaf{"FrrFollowBgpPic", fib.FrrFollowBgpPic}
    return &(fib.EntityData)
}

// Fib_PbtsForwardClassFallbacks
// PBTS class configuration
type Fib_PbtsForwardClassFallbacks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set PBTS class for fallback. The type is slice of
    // Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback.
    PbtsForwardClassFallback []Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback
}

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetEntityData() *types.CommonEntityData {
    pbtsForwardClassFallbacks.EntityData.YFilter = pbtsForwardClassFallbacks.YFilter
    pbtsForwardClassFallbacks.EntityData.YangName = "pbts-forward-class-fallbacks"
    pbtsForwardClassFallbacks.EntityData.BundleName = "cisco_ios_xr"
    pbtsForwardClassFallbacks.EntityData.ParentYangName = "fib"
    pbtsForwardClassFallbacks.EntityData.SegmentPath = "pbts-forward-class-fallbacks"
    pbtsForwardClassFallbacks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbtsForwardClassFallbacks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbtsForwardClassFallbacks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbtsForwardClassFallbacks.EntityData.Children = make(map[string]types.YChild)
    pbtsForwardClassFallbacks.EntityData.Children["pbts-forward-class-fallback"] = types.YChild{"PbtsForwardClassFallback", nil}
    for i := range pbtsForwardClassFallbacks.PbtsForwardClassFallback {
        pbtsForwardClassFallbacks.EntityData.Children[types.GetSegmentPath(&pbtsForwardClassFallbacks.PbtsForwardClassFallback[i])] = types.YChild{"PbtsForwardClassFallback", &pbtsForwardClassFallbacks.PbtsForwardClassFallback[i]}
    }
    pbtsForwardClassFallbacks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pbtsForwardClassFallbacks.EntityData)
}

// Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback
// Set PBTS class for fallback
type Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. PBTS forward class number. The type is one of the
    // following types: enumeration FibPbtsForwardClass, or int with range: 0..8.
    ForwardClassNumber interface{}

    // Set PBTS fallback type. The type is FibPbtsFallback. This attribute is
    // mandatory.
    FallbackType interface{}

    // Set PBTS fallback class number array. The type is slice of interface{} with
    // range: 0..7.
    FallbackClassNumberArray []interface{}
}

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetEntityData() *types.CommonEntityData {
    pbtsForwardClassFallback.EntityData.YFilter = pbtsForwardClassFallback.YFilter
    pbtsForwardClassFallback.EntityData.YangName = "pbts-forward-class-fallback"
    pbtsForwardClassFallback.EntityData.BundleName = "cisco_ios_xr"
    pbtsForwardClassFallback.EntityData.ParentYangName = "pbts-forward-class-fallbacks"
    pbtsForwardClassFallback.EntityData.SegmentPath = "pbts-forward-class-fallback" + "[forward-class-number='" + fmt.Sprintf("%v", pbtsForwardClassFallback.ForwardClassNumber) + "']"
    pbtsForwardClassFallback.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbtsForwardClassFallback.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbtsForwardClassFallback.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbtsForwardClassFallback.EntityData.Children = make(map[string]types.YChild)
    pbtsForwardClassFallback.EntityData.Leafs = make(map[string]types.YLeaf)
    pbtsForwardClassFallback.EntityData.Leafs["forward-class-number"] = types.YLeaf{"ForwardClassNumber", pbtsForwardClassFallback.ForwardClassNumber}
    pbtsForwardClassFallback.EntityData.Leafs["fallback-type"] = types.YLeaf{"FallbackType", pbtsForwardClassFallback.FallbackType}
    pbtsForwardClassFallback.EntityData.Leafs["fallback-class-number-array"] = types.YLeaf{"FallbackClassNumberArray", pbtsForwardClassFallback.FallbackClassNumberArray}
    return &(pbtsForwardClassFallback.EntityData)
}

// Fib_Platform
// FIB platform parameters
type Fib_Platform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Options for label-switched-multicast parameters.
    LabelSwitchedMulticast Fib_Platform_LabelSwitchedMulticast
}

func (platform *Fib_Platform) GetEntityData() *types.CommonEntityData {
    platform.EntityData.YFilter = platform.YFilter
    platform.EntityData.YangName = "platform"
    platform.EntityData.BundleName = "cisco_ios_xr"
    platform.EntityData.ParentYangName = "fib"
    platform.EntityData.SegmentPath = "platform"
    platform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platform.EntityData.Children = make(map[string]types.YChild)
    platform.EntityData.Children["label-switched-multicast"] = types.YChild{"LabelSwitchedMulticast", &platform.LabelSwitchedMulticast}
    platform.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(platform.EntityData)
}

// Fib_Platform_LabelSwitchedMulticast
// Options for label-switched-multicast parameters
type Fib_Platform_LabelSwitchedMulticast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set time to keep FRR slots programmed post FRR. The type is interface{}
    // with range: 3..180. Units are second.
    FrrHoldtime interface{}
}

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetEntityData() *types.CommonEntityData {
    labelSwitchedMulticast.EntityData.YFilter = labelSwitchedMulticast.YFilter
    labelSwitchedMulticast.EntityData.YangName = "label-switched-multicast"
    labelSwitchedMulticast.EntityData.BundleName = "cisco_ios_xr"
    labelSwitchedMulticast.EntityData.ParentYangName = "platform"
    labelSwitchedMulticast.EntityData.SegmentPath = "label-switched-multicast"
    labelSwitchedMulticast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelSwitchedMulticast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelSwitchedMulticast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelSwitchedMulticast.EntityData.Children = make(map[string]types.YChild)
    labelSwitchedMulticast.EntityData.Leafs = make(map[string]types.YLeaf)
    labelSwitchedMulticast.EntityData.Leafs["frr-holdtime"] = types.YLeaf{"FrrHoldtime", labelSwitchedMulticast.FrrHoldtime}
    return &(labelSwitchedMulticast.EntityData)
}

