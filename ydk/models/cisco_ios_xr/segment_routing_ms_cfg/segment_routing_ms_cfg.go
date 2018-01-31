// This module contains a collection of YANG definitions
// for Cisco IOS-XR segment-routing-ms package configuration.
// 
// This module contains definitions
// for the following management objects:
//   sr: Segment Routing
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package segment_routing_ms_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package segment_routing_ms_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-segment-routing-ms-cfg sr}", reflect.TypeOf(Sr{}))
    ydk.RegisterEntity("Cisco-IOS-XR-segment-routing-ms-cfg:sr", reflect.TypeOf(Sr{}))
}

// SrmsMiFlag represents Srms mi flag
type SrmsMiFlag string

const (
    // Disable flag
    SrmsMiFlag_disable SrmsMiFlag = "disable"

    // Enable flag
    SrmsMiFlag_enable SrmsMiFlag = "enable"
)

// Sr
// Segment Routing
type Sr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // enable SR. The type is interface{}.
    Enable interface{}

    // Global Block Segment Routing.
    GlobalBlock Sr_GlobalBlock

    // Segment Routing Local Block of Labels.
    LocalBlock Sr_LocalBlock

    // Mapping Server.
    Mappings Sr_Mappings
}

func (sr *Sr) GetFilter() yfilter.YFilter { return sr.YFilter }

func (sr *Sr) SetFilter(yf yfilter.YFilter) { sr.YFilter = yf }

func (sr *Sr) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "global-block" { return "GlobalBlock" }
    if yname == "local-block" { return "LocalBlock" }
    if yname == "mappings" { return "Mappings" }
    return ""
}

func (sr *Sr) GetSegmentPath() string {
    return "Cisco-IOS-XR-segment-routing-ms-cfg:sr"
}

func (sr *Sr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-block" {
        return &sr.GlobalBlock
    }
    if childYangName == "local-block" {
        return &sr.LocalBlock
    }
    if childYangName == "mappings" {
        return &sr.Mappings
    }
    return nil
}

func (sr *Sr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["global-block"] = &sr.GlobalBlock
    children["local-block"] = &sr.LocalBlock
    children["mappings"] = &sr.Mappings
    return children
}

func (sr *Sr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = sr.Enable
    return leafs
}

func (sr *Sr) GetBundleName() string { return "cisco_ios_xr" }

func (sr *Sr) GetYangName() string { return "sr" }

func (sr *Sr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sr *Sr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sr *Sr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sr *Sr) SetParent(parent types.Entity) { sr.parent = parent }

func (sr *Sr) GetParent() types.Entity { return sr.parent }

func (sr *Sr) GetParentYangName() string { return "Cisco-IOS-XR-segment-routing-ms-cfg" }

// Sr_GlobalBlock
// Global Block Segment Routing
// This type is a presence type.
type Sr_GlobalBlock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRGB Lower Bound. The type is interface{} with range: 16000..1048574. This
    // attribute is mandatory.
    LowerBound interface{}

    // SRGB Upper Bound. The type is interface{} with range: 16001..1048575. This
    // attribute is mandatory.
    UpperBound interface{}
}

func (globalBlock *Sr_GlobalBlock) GetFilter() yfilter.YFilter { return globalBlock.YFilter }

func (globalBlock *Sr_GlobalBlock) SetFilter(yf yfilter.YFilter) { globalBlock.YFilter = yf }

func (globalBlock *Sr_GlobalBlock) GetGoName(yname string) string {
    if yname == "lower-bound" { return "LowerBound" }
    if yname == "upper-bound" { return "UpperBound" }
    return ""
}

func (globalBlock *Sr_GlobalBlock) GetSegmentPath() string {
    return "global-block"
}

func (globalBlock *Sr_GlobalBlock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalBlock *Sr_GlobalBlock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalBlock *Sr_GlobalBlock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower-bound"] = globalBlock.LowerBound
    leafs["upper-bound"] = globalBlock.UpperBound
    return leafs
}

func (globalBlock *Sr_GlobalBlock) GetBundleName() string { return "cisco_ios_xr" }

func (globalBlock *Sr_GlobalBlock) GetYangName() string { return "global-block" }

func (globalBlock *Sr_GlobalBlock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalBlock *Sr_GlobalBlock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalBlock *Sr_GlobalBlock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalBlock *Sr_GlobalBlock) SetParent(parent types.Entity) { globalBlock.parent = parent }

func (globalBlock *Sr_GlobalBlock) GetParent() types.Entity { return globalBlock.parent }

func (globalBlock *Sr_GlobalBlock) GetParentYangName() string { return "sr" }

// Sr_LocalBlock
// Segment Routing Local Block of Labels
// This type is a presence type.
type Sr_LocalBlock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRLB Lower Bound. The type is interface{} with range: 15000..1048574. This
    // attribute is mandatory.
    LowerBound interface{}

    // SRLB Upper Bound. The type is interface{} with range: 15001..1048575. This
    // attribute is mandatory.
    UpperBound interface{}
}

func (localBlock *Sr_LocalBlock) GetFilter() yfilter.YFilter { return localBlock.YFilter }

func (localBlock *Sr_LocalBlock) SetFilter(yf yfilter.YFilter) { localBlock.YFilter = yf }

func (localBlock *Sr_LocalBlock) GetGoName(yname string) string {
    if yname == "lower-bound" { return "LowerBound" }
    if yname == "upper-bound" { return "UpperBound" }
    return ""
}

func (localBlock *Sr_LocalBlock) GetSegmentPath() string {
    return "local-block"
}

func (localBlock *Sr_LocalBlock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localBlock *Sr_LocalBlock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localBlock *Sr_LocalBlock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower-bound"] = localBlock.LowerBound
    leafs["upper-bound"] = localBlock.UpperBound
    return leafs
}

func (localBlock *Sr_LocalBlock) GetBundleName() string { return "cisco_ios_xr" }

func (localBlock *Sr_LocalBlock) GetYangName() string { return "local-block" }

func (localBlock *Sr_LocalBlock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localBlock *Sr_LocalBlock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localBlock *Sr_LocalBlock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localBlock *Sr_LocalBlock) SetParent(parent types.Entity) { localBlock.parent = parent }

func (localBlock *Sr_LocalBlock) GetParent() types.Entity { return localBlock.parent }

func (localBlock *Sr_LocalBlock) GetParentYangName() string { return "sr" }

// Sr_Mappings
// Mapping Server
type Sr_Mappings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP prefix to SID mapping. The type is slice of Sr_Mappings_Mapping.
    Mapping []Sr_Mappings_Mapping
}

func (mappings *Sr_Mappings) GetFilter() yfilter.YFilter { return mappings.YFilter }

func (mappings *Sr_Mappings) SetFilter(yf yfilter.YFilter) { mappings.YFilter = yf }

func (mappings *Sr_Mappings) GetGoName(yname string) string {
    if yname == "mapping" { return "Mapping" }
    return ""
}

func (mappings *Sr_Mappings) GetSegmentPath() string {
    return "mappings"
}

func (mappings *Sr_Mappings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mapping" {
        for _, c := range mappings.Mapping {
            if mappings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sr_Mappings_Mapping{}
        mappings.Mapping = append(mappings.Mapping, child)
        return &mappings.Mapping[len(mappings.Mapping)-1]
    }
    return nil
}

func (mappings *Sr_Mappings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mappings.Mapping {
        children[mappings.Mapping[i].GetSegmentPath()] = &mappings.Mapping[i]
    }
    return children
}

func (mappings *Sr_Mappings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mappings *Sr_Mappings) GetBundleName() string { return "cisco_ios_xr" }

func (mappings *Sr_Mappings) GetYangName() string { return "mappings" }

func (mappings *Sr_Mappings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mappings *Sr_Mappings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mappings *Sr_Mappings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mappings *Sr_Mappings) SetParent(parent types.Entity) { mappings.parent = parent }

func (mappings *Sr_Mappings) GetParent() types.Entity { return mappings.parent }

func (mappings *Sr_Mappings) GetParentYangName() string { return "sr" }

// Sr_Mappings_Mapping
// IP prefix to SID mapping
type Sr_Mappings_Mapping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Af interface{}

    // This attribute is a key. IP prefix. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Ip interface{}

    // This attribute is a key. Mask. The type is interface{} with range:
    // -2147483648..2147483647.
    Mask interface{}

    // Start of SID index range. The type is interface{} with range: 0..1048575.
    SidStart interface{}

    // Range (number of SIDs). The type is interface{} with range:
    // -2147483648..2147483647.
    SidRange interface{}

    // Enable/Disable Attached flag. The type is SrmsMiFlag.
    FlagAttached interface{}
}

func (mapping *Sr_Mappings_Mapping) GetFilter() yfilter.YFilter { return mapping.YFilter }

func (mapping *Sr_Mappings_Mapping) SetFilter(yf yfilter.YFilter) { mapping.YFilter = yf }

func (mapping *Sr_Mappings_Mapping) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    if yname == "ip" { return "Ip" }
    if yname == "mask" { return "Mask" }
    if yname == "sid-start" { return "SidStart" }
    if yname == "sid-range" { return "SidRange" }
    if yname == "flag-attached" { return "FlagAttached" }
    return ""
}

func (mapping *Sr_Mappings_Mapping) GetSegmentPath() string {
    return "mapping" + "[af='" + fmt.Sprintf("%v", mapping.Af) + "']" + "[ip='" + fmt.Sprintf("%v", mapping.Ip) + "']" + "[mask='" + fmt.Sprintf("%v", mapping.Mask) + "']"
}

func (mapping *Sr_Mappings_Mapping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mapping *Sr_Mappings_Mapping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mapping *Sr_Mappings_Mapping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af"] = mapping.Af
    leafs["ip"] = mapping.Ip
    leafs["mask"] = mapping.Mask
    leafs["sid-start"] = mapping.SidStart
    leafs["sid-range"] = mapping.SidRange
    leafs["flag-attached"] = mapping.FlagAttached
    return leafs
}

func (mapping *Sr_Mappings_Mapping) GetBundleName() string { return "cisco_ios_xr" }

func (mapping *Sr_Mappings_Mapping) GetYangName() string { return "mapping" }

func (mapping *Sr_Mappings_Mapping) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mapping *Sr_Mappings_Mapping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mapping *Sr_Mappings_Mapping) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mapping *Sr_Mappings_Mapping) SetParent(parent types.Entity) { mapping.parent = parent }

func (mapping *Sr_Mappings_Mapping) GetParent() types.Entity { return mapping.parent }

func (mapping *Sr_Mappings_Mapping) GetParentYangName() string { return "mappings" }

