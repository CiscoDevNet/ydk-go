// This module contains a collection of YANG definitions
// for Cisco IOS-XR fib-common package configuration.
// 
// This module contains definitions
// for the following management objects:
//   fib: CEF configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

    // Set option for automatcially recovering consistent-hashing state on
    // interface up. The type is bool.
    AutoHashRecover interface{}

    // Set options for adjacency routes overriding RIB routes. The type is bool.
    PreferAibRoutes interface{}

    // Set true to disable encapsulation sharing. The type is bool.
    EncapSharingDisable interface{}

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
    fib.EntityData.AbsolutePath = fib.EntityData.SegmentPath
    fib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fib.EntityData.Children = types.NewOrderedMap()
    fib.EntityData.Children.Append("pbts-forward-class-fallbacks", types.YChild{"PbtsForwardClassFallbacks", &fib.PbtsForwardClassFallbacks})
    fib.EntityData.Children.Append("platform", types.YChild{"Platform", &fib.Platform})
    fib.EntityData.Leafs = types.NewOrderedMap()
    fib.EntityData.Leafs.Append("auto-hash-recover", types.YLeaf{"AutoHashRecover", fib.AutoHashRecover})
    fib.EntityData.Leafs.Append("prefer-aib-routes", types.YLeaf{"PreferAibRoutes", fib.PreferAibRoutes})
    fib.EntityData.Leafs.Append("encap-sharing-disable", types.YLeaf{"EncapSharingDisable", fib.EncapSharingDisable})
    fib.EntityData.Leafs.Append("frr-follow-bgp-pic", types.YLeaf{"FrrFollowBgpPic", fib.FrrFollowBgpPic})

    fib.EntityData.YListKeys = []string {}

    return &(fib.EntityData)
}

// Fib_PbtsForwardClassFallbacks
// PBTS class configuration
type Fib_PbtsForwardClassFallbacks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set PBTS class for fallback. The type is slice of
    // Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback.
    PbtsForwardClassFallback []*Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback
}

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetEntityData() *types.CommonEntityData {
    pbtsForwardClassFallbacks.EntityData.YFilter = pbtsForwardClassFallbacks.YFilter
    pbtsForwardClassFallbacks.EntityData.YangName = "pbts-forward-class-fallbacks"
    pbtsForwardClassFallbacks.EntityData.BundleName = "cisco_ios_xr"
    pbtsForwardClassFallbacks.EntityData.ParentYangName = "fib"
    pbtsForwardClassFallbacks.EntityData.SegmentPath = "pbts-forward-class-fallbacks"
    pbtsForwardClassFallbacks.EntityData.AbsolutePath = "Cisco-IOS-XR-fib-common-cfg:fib/" + pbtsForwardClassFallbacks.EntityData.SegmentPath
    pbtsForwardClassFallbacks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbtsForwardClassFallbacks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbtsForwardClassFallbacks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbtsForwardClassFallbacks.EntityData.Children = types.NewOrderedMap()
    pbtsForwardClassFallbacks.EntityData.Children.Append("pbts-forward-class-fallback", types.YChild{"PbtsForwardClassFallback", nil})
    for i := range pbtsForwardClassFallbacks.PbtsForwardClassFallback {
        pbtsForwardClassFallbacks.EntityData.Children.Append(types.GetSegmentPath(pbtsForwardClassFallbacks.PbtsForwardClassFallback[i]), types.YChild{"PbtsForwardClassFallback", pbtsForwardClassFallbacks.PbtsForwardClassFallback[i]})
    }
    pbtsForwardClassFallbacks.EntityData.Leafs = types.NewOrderedMap()

    pbtsForwardClassFallbacks.EntityData.YListKeys = []string {}

    return &(pbtsForwardClassFallbacks.EntityData)
}

// Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback
// Set PBTS class for fallback
type Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    pbtsForwardClassFallback.EntityData.SegmentPath = "pbts-forward-class-fallback" + types.AddKeyToken(pbtsForwardClassFallback.ForwardClassNumber, "forward-class-number")
    pbtsForwardClassFallback.EntityData.AbsolutePath = "Cisco-IOS-XR-fib-common-cfg:fib/pbts-forward-class-fallbacks/" + pbtsForwardClassFallback.EntityData.SegmentPath
    pbtsForwardClassFallback.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbtsForwardClassFallback.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbtsForwardClassFallback.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbtsForwardClassFallback.EntityData.Children = types.NewOrderedMap()
    pbtsForwardClassFallback.EntityData.Leafs = types.NewOrderedMap()
    pbtsForwardClassFallback.EntityData.Leafs.Append("forward-class-number", types.YLeaf{"ForwardClassNumber", pbtsForwardClassFallback.ForwardClassNumber})
    pbtsForwardClassFallback.EntityData.Leafs.Append("fallback-type", types.YLeaf{"FallbackType", pbtsForwardClassFallback.FallbackType})
    pbtsForwardClassFallback.EntityData.Leafs.Append("fallback-class-number-array", types.YLeaf{"FallbackClassNumberArray", pbtsForwardClassFallback.FallbackClassNumberArray})

    pbtsForwardClassFallback.EntityData.YListKeys = []string {"ForwardClassNumber"}

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
    platform.EntityData.AbsolutePath = "Cisco-IOS-XR-fib-common-cfg:fib/" + platform.EntityData.SegmentPath
    platform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platform.EntityData.Children = types.NewOrderedMap()
    platform.EntityData.Children.Append("label-switched-multicast", types.YChild{"LabelSwitchedMulticast", &platform.LabelSwitchedMulticast})
    platform.EntityData.Leafs = types.NewOrderedMap()

    platform.EntityData.YListKeys = []string {}

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
    labelSwitchedMulticast.EntityData.AbsolutePath = "Cisco-IOS-XR-fib-common-cfg:fib/platform/" + labelSwitchedMulticast.EntityData.SegmentPath
    labelSwitchedMulticast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelSwitchedMulticast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelSwitchedMulticast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelSwitchedMulticast.EntityData.Children = types.NewOrderedMap()
    labelSwitchedMulticast.EntityData.Leafs = types.NewOrderedMap()
    labelSwitchedMulticast.EntityData.Leafs.Append("frr-holdtime", types.YLeaf{"FrrHoldtime", labelSwitchedMulticast.FrrHoldtime})

    labelSwitchedMulticast.EntityData.YListKeys = []string {}

    return &(labelSwitchedMulticast.EntityData)
}

