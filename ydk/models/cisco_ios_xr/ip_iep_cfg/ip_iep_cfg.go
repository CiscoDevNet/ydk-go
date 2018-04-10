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

// IpIepNum represents Ip iep num
type IpIepNum string

const (
    // Unnumbered
    IpIepNum_unnumbered IpIepNum = "unnumbered"

    // Numbered
    IpIepNum_numbered IpIepNum = "numbered"
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

// IpIepPath represents Ip iep path
type IpIepPath string

const (
    // Identifier
    IpIepPath_identifier IpIepPath = "identifier"

    // Name
    IpIepPath_name IpIepPath = "name"
)

// IpExplicitPaths
// IP Explicit Path config data
type IpExplicitPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of explicit paths.
    Paths IpExplicitPaths_Paths
}

func (ipExplicitPaths *IpExplicitPaths) GetEntityData() *types.CommonEntityData {
    ipExplicitPaths.EntityData.YFilter = ipExplicitPaths.YFilter
    ipExplicitPaths.EntityData.YangName = "ip-explicit-paths"
    ipExplicitPaths.EntityData.BundleName = "cisco_ios_xr"
    ipExplicitPaths.EntityData.ParentYangName = "Cisco-IOS-XR-ip-iep-cfg"
    ipExplicitPaths.EntityData.SegmentPath = "Cisco-IOS-XR-ip-iep-cfg:ip-explicit-paths"
    ipExplicitPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipExplicitPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipExplicitPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipExplicitPaths.EntityData.Children = make(map[string]types.YChild)
    ipExplicitPaths.EntityData.Children["paths"] = types.YChild{"Paths", &ipExplicitPaths.Paths}
    ipExplicitPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipExplicitPaths.EntityData)
}

// IpExplicitPaths_Paths
// A list of explicit paths
type IpExplicitPaths_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config data for a specific explicit path. The type is slice of
    // IpExplicitPaths_Paths_Path.
    Path []IpExplicitPaths_Paths_Path
}

func (paths *IpExplicitPaths_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "ip-explicit-paths"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = make(map[string]types.YChild)
    paths.EntityData.Children["path"] = types.YChild{"Path", nil}
    for i := range paths.Path {
        paths.EntityData.Children[types.GetSegmentPath(&paths.Path[i])] = types.YChild{"Path", &paths.Path[i]}
    }
    paths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(paths.EntityData)
}

// IpExplicitPaths_Paths_Path
// Config data for a specific explicit path
type IpExplicitPaths_Paths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path type. The type is IpIepPath.
    Type_ interface{}

    // name. The type is slice of IpExplicitPaths_Paths_Path_Name.
    Name []IpExplicitPaths_Paths_Path_Name

    // identifier. The type is slice of IpExplicitPaths_Paths_Path_Identifier.
    Identifier []IpExplicitPaths_Paths_Path_Identifier
}

func (path *IpExplicitPaths_Paths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "paths"
    path.EntityData.SegmentPath = "path" + "[type='" + fmt.Sprintf("%v", path.Type_) + "']"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Children["name"] = types.YChild{"Name", nil}
    for i := range path.Name {
        path.EntityData.Children[types.GetSegmentPath(&path.Name[i])] = types.YChild{"Name", &path.Name[i]}
    }
    path.EntityData.Children["identifier"] = types.YChild{"Identifier", nil}
    for i := range path.Identifier {
        path.EntityData.Children[types.GetSegmentPath(&path.Identifier[i])] = types.YChild{"Identifier", &path.Identifier[i]}
    }
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["type"] = types.YLeaf{"Type_", path.Type_}
    return &(path.EntityData)
}

// IpExplicitPaths_Paths_Path_Name
// name
type IpExplicitPaths_Paths_Path_Name struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // Disable the explicit path. The type is interface{}.
    Disable interface{}

    // List of Hops.
    Hops IpExplicitPaths_Paths_Path_Name_Hops
}

func (name *IpExplicitPaths_Paths_Path_Name) GetEntityData() *types.CommonEntityData {
    name.EntityData.YFilter = name.YFilter
    name.EntityData.YangName = "name"
    name.EntityData.BundleName = "cisco_ios_xr"
    name.EntityData.ParentYangName = "path"
    name.EntityData.SegmentPath = "name" + "[name='" + fmt.Sprintf("%v", name.Name) + "']"
    name.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    name.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    name.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    name.EntityData.Children = make(map[string]types.YChild)
    name.EntityData.Children["hops"] = types.YChild{"Hops", &name.Hops}
    name.EntityData.Leafs = make(map[string]types.YLeaf)
    name.EntityData.Leafs["name"] = types.YLeaf{"Name", name.Name}
    name.EntityData.Leafs["disable"] = types.YLeaf{"Disable", name.Disable}
    return &(name.EntityData)
}

// IpExplicitPaths_Paths_Path_Name_Hops
// List of Hops
type IpExplicitPaths_Paths_Path_Name_Hops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hop Information. The type is slice of
    // IpExplicitPaths_Paths_Path_Name_Hops_Hop.
    Hop []IpExplicitPaths_Paths_Path_Name_Hops_Hop
}

func (hops *IpExplicitPaths_Paths_Path_Name_Hops) GetEntityData() *types.CommonEntityData {
    hops.EntityData.YFilter = hops.YFilter
    hops.EntityData.YangName = "hops"
    hops.EntityData.BundleName = "cisco_ios_xr"
    hops.EntityData.ParentYangName = "name"
    hops.EntityData.SegmentPath = "hops"
    hops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hops.EntityData.Children = make(map[string]types.YChild)
    hops.EntityData.Children["hop"] = types.YChild{"Hop", nil}
    for i := range hops.Hop {
        hops.EntityData.Children[types.GetSegmentPath(&hops.Hop[i])] = types.YChild{"Hop", &hops.Hop[i]}
    }
    hops.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hops.EntityData)
}

// IpExplicitPaths_Paths_Path_Name_Hops_Hop
// Hop Information
type IpExplicitPaths_Paths_Path_Name_Hops_Hop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index number. The type is interface{} with range:
    // 1..65535.
    IndexNumber interface{}

    // IP address of the hop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (hop *IpExplicitPaths_Paths_Path_Name_Hops_Hop) GetEntityData() *types.CommonEntityData {
    hop.EntityData.YFilter = hop.YFilter
    hop.EntityData.YangName = "hop"
    hop.EntityData.BundleName = "cisco_ios_xr"
    hop.EntityData.ParentYangName = "hops"
    hop.EntityData.SegmentPath = "hop" + "[index-number='" + fmt.Sprintf("%v", hop.IndexNumber) + "']"
    hop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hop.EntityData.Children = make(map[string]types.YChild)
    hop.EntityData.Leafs = make(map[string]types.YLeaf)
    hop.EntityData.Leafs["index-number"] = types.YLeaf{"IndexNumber", hop.IndexNumber}
    hop.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", hop.IpAddress}
    hop.EntityData.Leafs["hop-type"] = types.YLeaf{"HopType", hop.HopType}
    hop.EntityData.Leafs["if-index"] = types.YLeaf{"IfIndex", hop.IfIndex}
    hop.EntityData.Leafs["num-type"] = types.YLeaf{"NumType", hop.NumType}
    hop.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", hop.MplsLabel}
    return &(hop.EntityData)
}

// IpExplicitPaths_Paths_Path_Identifier
// identifier
type IpExplicitPaths_Paths_Path_Identifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path identifier. The type is interface{} with
    // range: 1..65535.
    Id interface{}

    // Disable the explicit path. The type is interface{}.
    Disable interface{}

    // List of Hops.
    Hops IpExplicitPaths_Paths_Path_Identifier_Hops
}

func (identifier *IpExplicitPaths_Paths_Path_Identifier) GetEntityData() *types.CommonEntityData {
    identifier.EntityData.YFilter = identifier.YFilter
    identifier.EntityData.YangName = "identifier"
    identifier.EntityData.BundleName = "cisco_ios_xr"
    identifier.EntityData.ParentYangName = "path"
    identifier.EntityData.SegmentPath = "identifier" + "[id='" + fmt.Sprintf("%v", identifier.Id) + "']"
    identifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    identifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    identifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    identifier.EntityData.Children = make(map[string]types.YChild)
    identifier.EntityData.Children["hops"] = types.YChild{"Hops", &identifier.Hops}
    identifier.EntityData.Leafs = make(map[string]types.YLeaf)
    identifier.EntityData.Leafs["id"] = types.YLeaf{"Id", identifier.Id}
    identifier.EntityData.Leafs["disable"] = types.YLeaf{"Disable", identifier.Disable}
    return &(identifier.EntityData)
}

// IpExplicitPaths_Paths_Path_Identifier_Hops
// List of Hops
type IpExplicitPaths_Paths_Path_Identifier_Hops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hop Information. The type is slice of
    // IpExplicitPaths_Paths_Path_Identifier_Hops_Hop.
    Hop []IpExplicitPaths_Paths_Path_Identifier_Hops_Hop
}

func (hops *IpExplicitPaths_Paths_Path_Identifier_Hops) GetEntityData() *types.CommonEntityData {
    hops.EntityData.YFilter = hops.YFilter
    hops.EntityData.YangName = "hops"
    hops.EntityData.BundleName = "cisco_ios_xr"
    hops.EntityData.ParentYangName = "identifier"
    hops.EntityData.SegmentPath = "hops"
    hops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hops.EntityData.Children = make(map[string]types.YChild)
    hops.EntityData.Children["hop"] = types.YChild{"Hop", nil}
    for i := range hops.Hop {
        hops.EntityData.Children[types.GetSegmentPath(&hops.Hop[i])] = types.YChild{"Hop", &hops.Hop[i]}
    }
    hops.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hops.EntityData)
}

// IpExplicitPaths_Paths_Path_Identifier_Hops_Hop
// Hop Information
type IpExplicitPaths_Paths_Path_Identifier_Hops_Hop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index number. The type is interface{} with range:
    // 1..65535.
    IndexNumber interface{}

    // IP address of the hop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (hop *IpExplicitPaths_Paths_Path_Identifier_Hops_Hop) GetEntityData() *types.CommonEntityData {
    hop.EntityData.YFilter = hop.YFilter
    hop.EntityData.YangName = "hop"
    hop.EntityData.BundleName = "cisco_ios_xr"
    hop.EntityData.ParentYangName = "hops"
    hop.EntityData.SegmentPath = "hop" + "[index-number='" + fmt.Sprintf("%v", hop.IndexNumber) + "']"
    hop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hop.EntityData.Children = make(map[string]types.YChild)
    hop.EntityData.Leafs = make(map[string]types.YLeaf)
    hop.EntityData.Leafs["index-number"] = types.YLeaf{"IndexNumber", hop.IndexNumber}
    hop.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", hop.IpAddress}
    hop.EntityData.Leafs["hop-type"] = types.YLeaf{"HopType", hop.HopType}
    hop.EntityData.Leafs["if-index"] = types.YLeaf{"IfIndex", hop.IfIndex}
    hop.EntityData.Leafs["num-type"] = types.YLeaf{"NumType", hop.NumType}
    hop.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", hop.MplsLabel}
    return &(hop.EntityData)
}

