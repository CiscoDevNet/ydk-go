// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-iep package operational data.
// 
// This module contains definitions
// for the following management objects:
//   explicit-paths: Configured IP explicit paths
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ip_iep_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_iep_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-iep-oper explicit-paths}", reflect.TypeOf(ExplicitPaths{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-iep-oper:explicit-paths", reflect.TypeOf(ExplicitPaths{}))
}

// IepStatus represents Status of the path
type IepStatus string

const (
    // Status is enabled
    IepStatus_enabled IepStatus = "enabled"

    // Status is disabled
    IepStatus_disabled IepStatus = "disabled"
)

// IepAddress represents Address types
type IepAddress string

const (
    // Address type is next
    IepAddress_next IepAddress = "next"

    // Address type is exclude
    IepAddress_exclude IepAddress = "exclude"

    // Address type is exclude SRLG
    IepAddress_exclude_srlg IepAddress = "exclude-srlg"
)

// IepHop represents Hop types of the next-address configured
type IepHop string

const (
    // Hop type is strict
    IepHop_strict IepHop = "strict"

    // Hop type is loose
    IepHop_loose IepHop = "loose"
)

// ExplicitPaths
// Configured IP explicit paths
type ExplicitPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of configured IP explicit path identifiers, this corresponds to
    // mplsTunnelHopTable in TE MIB.
    Identifiers ExplicitPaths_Identifiers

    // List of configured IP explicit path names, this corresponds to
    // mplsTunnelHopTable in TE MIB.
    Names ExplicitPaths_Names
}

func (explicitPaths *ExplicitPaths) GetEntityData() *types.CommonEntityData {
    explicitPaths.EntityData.YFilter = explicitPaths.YFilter
    explicitPaths.EntityData.YangName = "explicit-paths"
    explicitPaths.EntityData.BundleName = "cisco_ios_xr"
    explicitPaths.EntityData.ParentYangName = "Cisco-IOS-XR-ip-iep-oper"
    explicitPaths.EntityData.SegmentPath = "Cisco-IOS-XR-ip-iep-oper:explicit-paths"
    explicitPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitPaths.EntityData.Children = types.NewOrderedMap()
    explicitPaths.EntityData.Children.Append("identifiers", types.YChild{"Identifiers", &explicitPaths.Identifiers})
    explicitPaths.EntityData.Children.Append("names", types.YChild{"Names", &explicitPaths.Names})
    explicitPaths.EntityData.Leafs = types.NewOrderedMap()

    explicitPaths.EntityData.YListKeys = []string {}

    return &(explicitPaths.EntityData)
}

// ExplicitPaths_Identifiers
// List of configured IP explicit path identifiers,
// this corresponds to mplsTunnelHopTable in TE MIB
type ExplicitPaths_Identifiers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP explicit path configured for a particular identifier. The type is slice
    // of ExplicitPaths_Identifiers_Identifier.
    Identifier []*ExplicitPaths_Identifiers_Identifier
}

func (identifiers *ExplicitPaths_Identifiers) GetEntityData() *types.CommonEntityData {
    identifiers.EntityData.YFilter = identifiers.YFilter
    identifiers.EntityData.YangName = "identifiers"
    identifiers.EntityData.BundleName = "cisco_ios_xr"
    identifiers.EntityData.ParentYangName = "explicit-paths"
    identifiers.EntityData.SegmentPath = "identifiers"
    identifiers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    identifiers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    identifiers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    identifiers.EntityData.Children = types.NewOrderedMap()
    identifiers.EntityData.Children.Append("identifier", types.YChild{"Identifier", nil})
    for i := range identifiers.Identifier {
        identifiers.EntityData.Children.Append(types.GetSegmentPath(identifiers.Identifier[i]), types.YChild{"Identifier", identifiers.Identifier[i]})
    }
    identifiers.EntityData.Leafs = types.NewOrderedMap()

    identifiers.EntityData.YListKeys = []string {}

    return &(identifiers.EntityData)
}

// ExplicitPaths_Identifiers_Identifier
// IP explicit path configured for a particular
// identifier
type ExplicitPaths_Identifiers_Identifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Identifier ID. The type is interface{} with range:
    // 0..4294967295.
    IdentifierId interface{}

    // Status of the path. The type is IepStatus.
    Status interface{}

    // List of IP addresses configured in the explicit path. The type is slice of
    // ExplicitPaths_Identifiers_Identifier_Address.
    Address []*ExplicitPaths_Identifiers_Identifier_Address
}

func (identifier *ExplicitPaths_Identifiers_Identifier) GetEntityData() *types.CommonEntityData {
    identifier.EntityData.YFilter = identifier.YFilter
    identifier.EntityData.YangName = "identifier"
    identifier.EntityData.BundleName = "cisco_ios_xr"
    identifier.EntityData.ParentYangName = "identifiers"
    identifier.EntityData.SegmentPath = "identifier" + types.AddKeyToken(identifier.IdentifierId, "identifier-id")
    identifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    identifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    identifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    identifier.EntityData.Children = types.NewOrderedMap()
    identifier.EntityData.Children.Append("address", types.YChild{"Address", nil})
    for i := range identifier.Address {
        identifier.EntityData.Children.Append(types.GetSegmentPath(identifier.Address[i]), types.YChild{"Address", identifier.Address[i]})
    }
    identifier.EntityData.Leafs = types.NewOrderedMap()
    identifier.EntityData.Leafs.Append("identifier-id", types.YLeaf{"IdentifierId", identifier.IdentifierId})
    identifier.EntityData.Leafs.Append("status", types.YLeaf{"Status", identifier.Status})

    identifier.EntityData.YListKeys = []string {"IdentifierId"}

    return &(identifier.EntityData)
}

// ExplicitPaths_Identifiers_Identifier_Address
// List of IP addresses configured in the explicit
// path
type ExplicitPaths_Identifiers_Identifier_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Index number at which the path entry is inserted or modified. The type is
    // interface{} with range: 0..4294967295.
    Index interface{}

    // Interface Index of the path. The type is interface{} with range:
    // 0..4294967295.
    IfIndex interface{}

    // Specifies the address type. The type is IepAddress.
    AddressType interface{}

    // Specifies the next unicast address in the path as a strict or loose hop.
    // The type is IepHop.
    HopType interface{}

    // IPv4 unicast address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // MPLS label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}
}

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "identifier"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("index", types.YLeaf{"Index", address.Index})
    address.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", address.IfIndex})
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("hop-type", types.YLeaf{"HopType", address.HopType})
    address.EntityData.Leafs.Append("address", types.YLeaf{"Address", address.Address})
    address.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", address.MplsLabel})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// ExplicitPaths_Names
// List of configured IP explicit path names, this
// corresponds to mplsTunnelHopTable in TE MIB
type ExplicitPaths_Names struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP explicit path configured for a particular path name. The type is slice
    // of ExplicitPaths_Names_Name.
    Name []*ExplicitPaths_Names_Name
}

func (names *ExplicitPaths_Names) GetEntityData() *types.CommonEntityData {
    names.EntityData.YFilter = names.YFilter
    names.EntityData.YangName = "names"
    names.EntityData.BundleName = "cisco_ios_xr"
    names.EntityData.ParentYangName = "explicit-paths"
    names.EntityData.SegmentPath = "names"
    names.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    names.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    names.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    names.EntityData.Children = types.NewOrderedMap()
    names.EntityData.Children.Append("name", types.YChild{"Name", nil})
    for i := range names.Name {
        names.EntityData.Children.Append(types.GetSegmentPath(names.Name[i]), types.YChild{"Name", names.Name[i]})
    }
    names.EntityData.Leafs = types.NewOrderedMap()

    names.EntityData.YListKeys = []string {}

    return &(names.EntityData)
}

// ExplicitPaths_Names_Name
// IP explicit path configured for a particular
// path name
type ExplicitPaths_Names_Name struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    PathName interface{}

    // Status of the path. The type is IepStatus.
    Status interface{}

    // List of IP addresses configured in the explicit path. The type is slice of
    // ExplicitPaths_Names_Name_Address.
    Address []*ExplicitPaths_Names_Name_Address
}

func (name *ExplicitPaths_Names_Name) GetEntityData() *types.CommonEntityData {
    name.EntityData.YFilter = name.YFilter
    name.EntityData.YangName = "name"
    name.EntityData.BundleName = "cisco_ios_xr"
    name.EntityData.ParentYangName = "names"
    name.EntityData.SegmentPath = "name" + types.AddKeyToken(name.PathName, "path-name")
    name.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    name.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    name.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    name.EntityData.Children = types.NewOrderedMap()
    name.EntityData.Children.Append("address", types.YChild{"Address", nil})
    for i := range name.Address {
        name.EntityData.Children.Append(types.GetSegmentPath(name.Address[i]), types.YChild{"Address", name.Address[i]})
    }
    name.EntityData.Leafs = types.NewOrderedMap()
    name.EntityData.Leafs.Append("path-name", types.YLeaf{"PathName", name.PathName})
    name.EntityData.Leafs.Append("status", types.YLeaf{"Status", name.Status})

    name.EntityData.YListKeys = []string {"PathName"}

    return &(name.EntityData)
}

// ExplicitPaths_Names_Name_Address
// List of IP addresses configured in the explicit
// path
type ExplicitPaths_Names_Name_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Index number at which the path entry is inserted or modified. The type is
    // interface{} with range: 0..4294967295.
    Index interface{}

    // Interface Index of the path. The type is interface{} with range:
    // 0..4294967295.
    IfIndex interface{}

    // Specifies the address type. The type is IepAddress.
    AddressType interface{}

    // Specifies the next unicast address in the path as a strict or loose hop.
    // The type is IepHop.
    HopType interface{}

    // IPv4 unicast address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // MPLS label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}
}

func (address *ExplicitPaths_Names_Name_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "name"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("index", types.YLeaf{"Index", address.Index})
    address.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", address.IfIndex})
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("hop-type", types.YLeaf{"HopType", address.HopType})
    address.EntityData.Leafs.Append("address", types.YLeaf{"Address", address.Address})
    address.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", address.MplsLabel})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

