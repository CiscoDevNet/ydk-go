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

// FibPbtsForwardClass represents Fib pbts forward class
type FibPbtsForwardClass string

const (
    // Any class
    FibPbtsForwardClass_any FibPbtsForwardClass = "any"
)

// Fib
// CEF configuration
type Fib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set options for adjacency routes overriding RIB routes. The type is bool.
    PreferAibRoutes interface{}

    // PBTS class configuration.
    PbtsForwardClassFallbacks Fib_PbtsForwardClassFallbacks

    // FIB platform parameters.
    Platform Fib_Platform
}

func (fib *Fib) GetFilter() yfilter.YFilter { return fib.YFilter }

func (fib *Fib) SetFilter(yf yfilter.YFilter) { fib.YFilter = yf }

func (fib *Fib) GetGoName(yname string) string {
    if yname == "prefer-aib-routes" { return "PreferAibRoutes" }
    if yname == "pbts-forward-class-fallbacks" { return "PbtsForwardClassFallbacks" }
    if yname == "platform" { return "Platform" }
    return ""
}

func (fib *Fib) GetSegmentPath() string {
    return "Cisco-IOS-XR-fib-common-cfg:fib"
}

func (fib *Fib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pbts-forward-class-fallbacks" {
        return &fib.PbtsForwardClassFallbacks
    }
    if childYangName == "platform" {
        return &fib.Platform
    }
    return nil
}

func (fib *Fib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pbts-forward-class-fallbacks"] = &fib.PbtsForwardClassFallbacks
    children["platform"] = &fib.Platform
    return children
}

func (fib *Fib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefer-aib-routes"] = fib.PreferAibRoutes
    return leafs
}

func (fib *Fib) GetBundleName() string { return "cisco_ios_xr" }

func (fib *Fib) GetYangName() string { return "fib" }

func (fib *Fib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fib *Fib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fib *Fib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fib *Fib) SetParent(parent types.Entity) { fib.parent = parent }

func (fib *Fib) GetParent() types.Entity { return fib.parent }

func (fib *Fib) GetParentYangName() string { return "Cisco-IOS-XR-fib-common-cfg" }

// Fib_PbtsForwardClassFallbacks
// PBTS class configuration
type Fib_PbtsForwardClassFallbacks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set PBTS class for fallback. The type is slice of
    // Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback.
    PbtsForwardClassFallback []Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback
}

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetFilter() yfilter.YFilter { return pbtsForwardClassFallbacks.YFilter }

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) SetFilter(yf yfilter.YFilter) { pbtsForwardClassFallbacks.YFilter = yf }

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetGoName(yname string) string {
    if yname == "pbts-forward-class-fallback" { return "PbtsForwardClassFallback" }
    return ""
}

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetSegmentPath() string {
    return "pbts-forward-class-fallbacks"
}

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pbts-forward-class-fallback" {
        for _, c := range pbtsForwardClassFallbacks.PbtsForwardClassFallback {
            if pbtsForwardClassFallbacks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback{}
        pbtsForwardClassFallbacks.PbtsForwardClassFallback = append(pbtsForwardClassFallbacks.PbtsForwardClassFallback, child)
        return &pbtsForwardClassFallbacks.PbtsForwardClassFallback[len(pbtsForwardClassFallbacks.PbtsForwardClassFallback)-1]
    }
    return nil
}

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pbtsForwardClassFallbacks.PbtsForwardClassFallback {
        children[pbtsForwardClassFallbacks.PbtsForwardClassFallback[i].GetSegmentPath()] = &pbtsForwardClassFallbacks.PbtsForwardClassFallback[i]
    }
    return children
}

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetBundleName() string { return "cisco_ios_xr" }

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetYangName() string { return "pbts-forward-class-fallbacks" }

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) SetParent(parent types.Entity) { pbtsForwardClassFallbacks.parent = parent }

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetParent() types.Entity { return pbtsForwardClassFallbacks.parent }

func (pbtsForwardClassFallbacks *Fib_PbtsForwardClassFallbacks) GetParentYangName() string { return "fib" }

// Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback
// Set PBTS class for fallback
type Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback struct {
    parent types.Entity
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

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetFilter() yfilter.YFilter { return pbtsForwardClassFallback.YFilter }

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) SetFilter(yf yfilter.YFilter) { pbtsForwardClassFallback.YFilter = yf }

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetGoName(yname string) string {
    if yname == "forward-class-number" { return "ForwardClassNumber" }
    if yname == "fallback-type" { return "FallbackType" }
    if yname == "fallback-class-number-array" { return "FallbackClassNumberArray" }
    return ""
}

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetSegmentPath() string {
    return "pbts-forward-class-fallback" + "[forward-class-number='" + fmt.Sprintf("%v", pbtsForwardClassFallback.ForwardClassNumber) + "']"
}

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["forward-class-number"] = pbtsForwardClassFallback.ForwardClassNumber
    leafs["fallback-type"] = pbtsForwardClassFallback.FallbackType
    leafs["fallback-class-number-array"] = pbtsForwardClassFallback.FallbackClassNumberArray
    return leafs
}

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetBundleName() string { return "cisco_ios_xr" }

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetYangName() string { return "pbts-forward-class-fallback" }

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) SetParent(parent types.Entity) { pbtsForwardClassFallback.parent = parent }

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetParent() types.Entity { return pbtsForwardClassFallback.parent }

func (pbtsForwardClassFallback *Fib_PbtsForwardClassFallbacks_PbtsForwardClassFallback) GetParentYangName() string { return "pbts-forward-class-fallbacks" }

// Fib_Platform
// FIB platform parameters
type Fib_Platform struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Options for label-switched-multicast parameters.
    LabelSwitchedMulticast Fib_Platform_LabelSwitchedMulticast
}

func (platform *Fib_Platform) GetFilter() yfilter.YFilter { return platform.YFilter }

func (platform *Fib_Platform) SetFilter(yf yfilter.YFilter) { platform.YFilter = yf }

func (platform *Fib_Platform) GetGoName(yname string) string {
    if yname == "label-switched-multicast" { return "LabelSwitchedMulticast" }
    return ""
}

func (platform *Fib_Platform) GetSegmentPath() string {
    return "platform"
}

func (platform *Fib_Platform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-switched-multicast" {
        return &platform.LabelSwitchedMulticast
    }
    return nil
}

func (platform *Fib_Platform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["label-switched-multicast"] = &platform.LabelSwitchedMulticast
    return children
}

func (platform *Fib_Platform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platform *Fib_Platform) GetBundleName() string { return "cisco_ios_xr" }

func (platform *Fib_Platform) GetYangName() string { return "platform" }

func (platform *Fib_Platform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platform *Fib_Platform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platform *Fib_Platform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platform *Fib_Platform) SetParent(parent types.Entity) { platform.parent = parent }

func (platform *Fib_Platform) GetParent() types.Entity { return platform.parent }

func (platform *Fib_Platform) GetParentYangName() string { return "fib" }

// Fib_Platform_LabelSwitchedMulticast
// Options for label-switched-multicast parameters
type Fib_Platform_LabelSwitchedMulticast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set time to keep FRR slots programmed post FRR. The type is interface{}
    // with range: 3..180. Units are second.
    FrrHoldtime interface{}
}

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetFilter() yfilter.YFilter { return labelSwitchedMulticast.YFilter }

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) SetFilter(yf yfilter.YFilter) { labelSwitchedMulticast.YFilter = yf }

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetGoName(yname string) string {
    if yname == "frr-holdtime" { return "FrrHoldtime" }
    return ""
}

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetSegmentPath() string {
    return "label-switched-multicast"
}

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frr-holdtime"] = labelSwitchedMulticast.FrrHoldtime
    return leafs
}

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetBundleName() string { return "cisco_ios_xr" }

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetYangName() string { return "label-switched-multicast" }

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) SetParent(parent types.Entity) { labelSwitchedMulticast.parent = parent }

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetParent() types.Entity { return labelSwitchedMulticast.parent }

func (labelSwitchedMulticast *Fib_Platform_LabelSwitchedMulticast) GetParentYangName() string { return "platform" }

