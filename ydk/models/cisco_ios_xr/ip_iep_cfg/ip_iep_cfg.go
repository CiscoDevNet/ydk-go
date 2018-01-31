// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-iep package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ip-explicit-paths: IP Explicit Path config data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_iep_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_iep_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-iep-cfg ip-explicit-paths}", reflect.TypeOf(IpExplicitPaths{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-iep-cfg:ip-explicit-paths", reflect.TypeOf(IpExplicitPaths{}))
}

// IpIepPath represents Ip iep path
type IpIepPath string

const (
    // Identifier
    IpIepPath_identifier IpIepPath = "identifier"

    // Name
    IpIepPath_name IpIepPath = "name"
)

// IpIepHop represents Ip iep hop
type IpIepHop string

const (
    // NextStrict
    IpIepHop_next_strict IpIepHop = "next-strict"

    // NextLoose
    IpIepHop_next_loose IpIepHop = "next-loose"

    // Exclude
    IpIepHop_exclude IpIepHop = "exclude"

    // Exclude Shared Risk Link Group
    IpIepHop_exclude_srlg IpIepHop = "exclude-srlg"

    // NextLabel
    IpIepHop_next_label IpIepHop = "next-label"
)

// IpIepNum represents Ip iep num
type IpIepNum string

const (
    // Unnumbered
    IpIepNum_unnumbered IpIepNum = "unnumbered"

    // Numbered
    IpIepNum_numbered IpIepNum = "numbered"
)

// IpExplicitPaths
// IP Explicit Path config data
type IpExplicitPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of explicit paths.
    Paths IpExplicitPaths_Paths
}

func (ipExplicitPaths *IpExplicitPaths) GetFilter() yfilter.YFilter { return ipExplicitPaths.YFilter }

func (ipExplicitPaths *IpExplicitPaths) SetFilter(yf yfilter.YFilter) { ipExplicitPaths.YFilter = yf }

func (ipExplicitPaths *IpExplicitPaths) GetGoName(yname string) string {
    if yname == "paths" { return "Paths" }
    return ""
}

func (ipExplicitPaths *IpExplicitPaths) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-iep-cfg:ip-explicit-paths"
}

func (ipExplicitPaths *IpExplicitPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "paths" {
        return &ipExplicitPaths.Paths
    }
    return nil
}

func (ipExplicitPaths *IpExplicitPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["paths"] = &ipExplicitPaths.Paths
    return children
}

func (ipExplicitPaths *IpExplicitPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipExplicitPaths *IpExplicitPaths) GetBundleName() string { return "cisco_ios_xr" }

func (ipExplicitPaths *IpExplicitPaths) GetYangName() string { return "ip-explicit-paths" }

func (ipExplicitPaths *IpExplicitPaths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipExplicitPaths *IpExplicitPaths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipExplicitPaths *IpExplicitPaths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipExplicitPaths *IpExplicitPaths) SetParent(parent types.Entity) { ipExplicitPaths.parent = parent }

func (ipExplicitPaths *IpExplicitPaths) GetParent() types.Entity { return ipExplicitPaths.parent }

func (ipExplicitPaths *IpExplicitPaths) GetParentYangName() string { return "Cisco-IOS-XR-ip-iep-cfg" }

// IpExplicitPaths_Paths
// A list of explicit paths
type IpExplicitPaths_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config data for a specific explicit path. The type is slice of
    // IpExplicitPaths_Paths_Path.
    Path []IpExplicitPaths_Paths_Path
}

func (paths *IpExplicitPaths_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *IpExplicitPaths_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *IpExplicitPaths_Paths) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (paths *IpExplicitPaths_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *IpExplicitPaths_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        for _, c := range paths.Path {
            if paths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpExplicitPaths_Paths_Path{}
        paths.Path = append(paths.Path, child)
        return &paths.Path[len(paths.Path)-1]
    }
    return nil
}

func (paths *IpExplicitPaths_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range paths.Path {
        children[paths.Path[i].GetSegmentPath()] = &paths.Path[i]
    }
    return children
}

func (paths *IpExplicitPaths_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (paths *IpExplicitPaths_Paths) GetBundleName() string { return "cisco_ios_xr" }

func (paths *IpExplicitPaths_Paths) GetYangName() string { return "paths" }

func (paths *IpExplicitPaths_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paths *IpExplicitPaths_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paths *IpExplicitPaths_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paths *IpExplicitPaths_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *IpExplicitPaths_Paths) GetParent() types.Entity { return paths.parent }

func (paths *IpExplicitPaths_Paths) GetParentYangName() string { return "ip-explicit-paths" }

// IpExplicitPaths_Paths_Path
// Config data for a specific explicit path
type IpExplicitPaths_Paths_Path struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Path type. The type is IpIepPath.
    Type interface{}

    // name. The type is slice of IpExplicitPaths_Paths_Path_Name.
    Name []IpExplicitPaths_Paths_Path_Name

    // identifier. The type is slice of IpExplicitPaths_Paths_Path_Identifier.
    Identifier []IpExplicitPaths_Paths_Path_Identifier
}

func (path *IpExplicitPaths_Paths_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *IpExplicitPaths_Paths_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *IpExplicitPaths_Paths_Path) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "name" { return "Name" }
    if yname == "identifier" { return "Identifier" }
    return ""
}

func (path *IpExplicitPaths_Paths_Path) GetSegmentPath() string {
    return "path" + "[type='" + fmt.Sprintf("%v", path.Type) + "']"
}

func (path *IpExplicitPaths_Paths_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "name" {
        for _, c := range path.Name {
            if path.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpExplicitPaths_Paths_Path_Name{}
        path.Name = append(path.Name, child)
        return &path.Name[len(path.Name)-1]
    }
    if childYangName == "identifier" {
        for _, c := range path.Identifier {
            if path.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpExplicitPaths_Paths_Path_Identifier{}
        path.Identifier = append(path.Identifier, child)
        return &path.Identifier[len(path.Identifier)-1]
    }
    return nil
}

func (path *IpExplicitPaths_Paths_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range path.Name {
        children[path.Name[i].GetSegmentPath()] = &path.Name[i]
    }
    for i := range path.Identifier {
        children[path.Identifier[i].GetSegmentPath()] = &path.Identifier[i]
    }
    return children
}

func (path *IpExplicitPaths_Paths_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = path.Type
    return leafs
}

func (path *IpExplicitPaths_Paths_Path) GetBundleName() string { return "cisco_ios_xr" }

func (path *IpExplicitPaths_Paths_Path) GetYangName() string { return "path" }

func (path *IpExplicitPaths_Paths_Path) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (path *IpExplicitPaths_Paths_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (path *IpExplicitPaths_Paths_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (path *IpExplicitPaths_Paths_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *IpExplicitPaths_Paths_Path) GetParent() types.Entity { return path.parent }

func (path *IpExplicitPaths_Paths_Path) GetParentYangName() string { return "paths" }

// IpExplicitPaths_Paths_Path_Name
// name
type IpExplicitPaths_Paths_Path_Name struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Path name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Disable the explicit path. The type is interface{}.
    Disable interface{}

    // List of Hops.
    Hops IpExplicitPaths_Paths_Path_Name_Hops
}

func (name *IpExplicitPaths_Paths_Path_Name) GetFilter() yfilter.YFilter { return name.YFilter }

func (name *IpExplicitPaths_Paths_Path_Name) SetFilter(yf yfilter.YFilter) { name.YFilter = yf }

func (name *IpExplicitPaths_Paths_Path_Name) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "disable" { return "Disable" }
    if yname == "hops" { return "Hops" }
    return ""
}

func (name *IpExplicitPaths_Paths_Path_Name) GetSegmentPath() string {
    return "name" + "[name='" + fmt.Sprintf("%v", name.Name) + "']"
}

func (name *IpExplicitPaths_Paths_Path_Name) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hops" {
        return &name.Hops
    }
    return nil
}

func (name *IpExplicitPaths_Paths_Path_Name) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hops"] = &name.Hops
    return children
}

func (name *IpExplicitPaths_Paths_Path_Name) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = name.Name
    leafs["disable"] = name.Disable
    return leafs
}

func (name *IpExplicitPaths_Paths_Path_Name) GetBundleName() string { return "cisco_ios_xr" }

func (name *IpExplicitPaths_Paths_Path_Name) GetYangName() string { return "name" }

func (name *IpExplicitPaths_Paths_Path_Name) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (name *IpExplicitPaths_Paths_Path_Name) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (name *IpExplicitPaths_Paths_Path_Name) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (name *IpExplicitPaths_Paths_Path_Name) SetParent(parent types.Entity) { name.parent = parent }

func (name *IpExplicitPaths_Paths_Path_Name) GetParent() types.Entity { return name.parent }

func (name *IpExplicitPaths_Paths_Path_Name) GetParentYangName() string { return "path" }

// IpExplicitPaths_Paths_Path_Name_Hops
// List of Hops
type IpExplicitPaths_Paths_Path_Name_Hops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Hop Information. The type is slice of
    // IpExplicitPaths_Paths_Path_Name_Hops_Hop.
    Hop []IpExplicitPaths_Paths_Path_Name_Hops_Hop
}

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetFilter() yfilter.YFilter { return hops.YFilter }

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) SetFilter(yf yfilter.YFilter) { hops.YFilter = yf }

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetGoName(yname string) string {
    if yname == "hop" { return "Hop" }
    return ""
}

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetSegmentPath() string {
    return "hops"
}

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hop" {
        for _, c := range hops.Hop {
            if hops.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpExplicitPaths_Paths_Path_Name_Hops_Hop{}
        hops.Hop = append(hops.Hop, child)
        return &hops.Hop[len(hops.Hop)-1]
    }
    return nil
}

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hops.Hop {
        children[hops.Hop[i].GetSegmentPath()] = &hops.Hop[i]
    }
    return children
}

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetBundleName() string { return "cisco_ios_xr" }

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetYangName() string { return "hops" }

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) SetParent(parent types.Entity) { hops.parent = parent }

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetParent() types.Entity { return hops.parent }

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetParentYangName() string { return "name" }

// IpExplicitPaths_Paths_Path_Name_Hops_Hop
// Hop Information
type IpExplicitPaths_Paths_Path_Name_Hops_Hop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index number. The type is interface{} with range:
    // 1..65535.
    IndexNumber interface{}

    // IP address of the hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // The default value is 0.0.0.0.
    IpAddress interface{}

    // Include or exclude this hop in the path. The type is IpIepHop. The default
    // value is next-strict.
    HopType interface{}

    // Ifindex value. The type is interface{} with range: -2147483648..2147483647.
    // The default value is 0.
    IfIndex interface{}

    // Number type Numbered or Unnumbered. The type is IpIepNum. The default value
    // is numbered.
    NumType interface{}

    // MPLS Label. The type is interface{} with range: 0..1048575.
    MplsLabel interface{}
}

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetFilter() yfilter.YFilter { return hop.YFilter }

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) SetFilter(yf yfilter.YFilter) { hop.YFilter = yf }

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetGoName(yname string) string {
    if yname == "index-number" { return "IndexNumber" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "hop-type" { return "HopType" }
    if yname == "if-index" { return "IfIndex" }
    if yname == "num-type" { return "NumType" }
    if yname == "mpls-label" { return "MplsLabel" }
    return ""
}

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetSegmentPath() string {
    return "hop" + "[index-number='" + fmt.Sprintf("%v", hop.IndexNumber) + "']"
}

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index-number"] = hop.IndexNumber
    leafs["ip-address"] = hop.IpAddress
    leafs["hop-type"] = hop.HopType
    leafs["if-index"] = hop.IfIndex
    leafs["num-type"] = hop.NumType
    leafs["mpls-label"] = hop.MplsLabel
    return leafs
}

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetBundleName() string { return "cisco_ios_xr" }

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetYangName() string { return "hop" }

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) SetParent(parent types.Entity) { hop.parent = parent }

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetParent() types.Entity { return hop.parent }

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetParentYangName() string { return "hops" }

// IpExplicitPaths_Paths_Path_Identifier
// identifier
type IpExplicitPaths_Paths_Path_Identifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Path identifier. The type is interface{} with
    // range: 1..65535.
    Id interface{}

    // Disable the explicit path. The type is interface{}.
    Disable interface{}

    // List of Hops.
    Hops IpExplicitPaths_Paths_Path_Identifier_Hops
}

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetFilter() yfilter.YFilter { return identifier.YFilter }

func (identifier *IpExplicitPaths_Paths_Path_Identifier) SetFilter(yf yfilter.YFilter) { identifier.YFilter = yf }

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "disable" { return "Disable" }
    if yname == "hops" { return "Hops" }
    return ""
}

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetSegmentPath() string {
    return "identifier" + "[id='" + fmt.Sprintf("%v", identifier.Id) + "']"
}

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hops" {
        return &identifier.Hops
    }
    return nil
}

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hops"] = &identifier.Hops
    return children
}

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = identifier.Id
    leafs["disable"] = identifier.Disable
    return leafs
}

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetBundleName() string { return "cisco_ios_xr" }

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetYangName() string { return "identifier" }

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (identifier *IpExplicitPaths_Paths_Path_Identifier) SetParent(parent types.Entity) { identifier.parent = parent }

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetParent() types.Entity { return identifier.parent }

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetParentYangName() string { return "path" }

// IpExplicitPaths_Paths_Path_Identifier_Hops
// List of Hops
type IpExplicitPaths_Paths_Path_Identifier_Hops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Hop Information. The type is slice of
    // IpExplicitPaths_Paths_Path_Identifier_Hops_Hop.
    Hop []IpExplicitPaths_Paths_Path_Identifier_Hops_Hop
}

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetFilter() yfilter.YFilter { return hops.YFilter }

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) SetFilter(yf yfilter.YFilter) { hops.YFilter = yf }

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetGoName(yname string) string {
    if yname == "hop" { return "Hop" }
    return ""
}

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetSegmentPath() string {
    return "hops"
}

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hop" {
        for _, c := range hops.Hop {
            if hops.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpExplicitPaths_Paths_Path_Identifier_Hops_Hop{}
        hops.Hop = append(hops.Hop, child)
        return &hops.Hop[len(hops.Hop)-1]
    }
    return nil
}

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hops.Hop {
        children[hops.Hop[i].GetSegmentPath()] = &hops.Hop[i]
    }
    return children
}

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetBundleName() string { return "cisco_ios_xr" }

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetYangName() string { return "hops" }

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) SetParent(parent types.Entity) { hops.parent = parent }

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetParent() types.Entity { return hops.parent }

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetParentYangName() string { return "identifier" }

// IpExplicitPaths_Paths_Path_Identifier_Hops_Hop
// Hop Information
type IpExplicitPaths_Paths_Path_Identifier_Hops_Hop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index number. The type is interface{} with range:
    // 1..65535.
    IndexNumber interface{}

    // IP address of the hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // The default value is 0.0.0.0.
    IpAddress interface{}

    // Include or exclude this hop in the path. The type is IpIepHop. The default
    // value is next-strict.
    HopType interface{}

    // Ifindex value. The type is interface{} with range: -2147483648..2147483647.
    // The default value is 0.
    IfIndex interface{}

    // Number type Numbered or Unnumbered. The type is IpIepNum. The default value
    // is numbered.
    NumType interface{}

    // MPLS Label. The type is interface{} with range: 0..1048575.
    MplsLabel interface{}
}

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetFilter() yfilter.YFilter { return hop.YFilter }

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) SetFilter(yf yfilter.YFilter) { hop.YFilter = yf }

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetGoName(yname string) string {
    if yname == "index-number" { return "IndexNumber" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "hop-type" { return "HopType" }
    if yname == "if-index" { return "IfIndex" }
    if yname == "num-type" { return "NumType" }
    if yname == "mpls-label" { return "MplsLabel" }
    return ""
}

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetSegmentPath() string {
    return "hop" + "[index-number='" + fmt.Sprintf("%v", hop.IndexNumber) + "']"
}

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index-number"] = hop.IndexNumber
    leafs["ip-address"] = hop.IpAddress
    leafs["hop-type"] = hop.HopType
    leafs["if-index"] = hop.IfIndex
    leafs["num-type"] = hop.NumType
    leafs["mpls-label"] = hop.MplsLabel
    return leafs
}

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetBundleName() string { return "cisco_ios_xr" }

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetYangName() string { return "hop" }

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) SetParent(parent types.Entity) { hop.parent = parent }

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetParent() types.Entity { return hop.parent }

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetParentYangName() string { return "hops" }

