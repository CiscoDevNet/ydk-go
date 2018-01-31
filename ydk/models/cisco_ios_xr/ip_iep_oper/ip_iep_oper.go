// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-iep package operational data.
// 
// This module contains definitions
// for the following management objects:
//   explicit-paths: Configured IP explicit paths
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // List of configured IP explicit path identifiers, this corresponds to
    // mplsTunnelHopTable in TE MIB.
    Identifiers ExplicitPaths_Identifiers

    // List of configured IP explicit path names, this corresponds to
    // mplsTunnelHopTable in TE MIB.
    Names ExplicitPaths_Names
}

func (explicitPaths *ExplicitPaths) GetFilter() yfilter.YFilter { return explicitPaths.YFilter }

func (explicitPaths *ExplicitPaths) SetFilter(yf yfilter.YFilter) { explicitPaths.YFilter = yf }

func (explicitPaths *ExplicitPaths) GetGoName(yname string) string {
    if yname == "identifiers" { return "Identifiers" }
    if yname == "names" { return "Names" }
    return ""
}

func (explicitPaths *ExplicitPaths) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-iep-oper:explicit-paths"
}

func (explicitPaths *ExplicitPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "identifiers" {
        return &explicitPaths.Identifiers
    }
    if childYangName == "names" {
        return &explicitPaths.Names
    }
    return nil
}

func (explicitPaths *ExplicitPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["identifiers"] = &explicitPaths.Identifiers
    children["names"] = &explicitPaths.Names
    return children
}

func (explicitPaths *ExplicitPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (explicitPaths *ExplicitPaths) GetBundleName() string { return "cisco_ios_xr" }

func (explicitPaths *ExplicitPaths) GetYangName() string { return "explicit-paths" }

func (explicitPaths *ExplicitPaths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitPaths *ExplicitPaths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitPaths *ExplicitPaths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitPaths *ExplicitPaths) SetParent(parent types.Entity) { explicitPaths.parent = parent }

func (explicitPaths *ExplicitPaths) GetParent() types.Entity { return explicitPaths.parent }

func (explicitPaths *ExplicitPaths) GetParentYangName() string { return "Cisco-IOS-XR-ip-iep-oper" }

// ExplicitPaths_Identifiers
// List of configured IP explicit path identifiers,
// this corresponds to mplsTunnelHopTable in TE MIB
type ExplicitPaths_Identifiers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP explicit path configured for a particular identifier. The type is slice
    // of ExplicitPaths_Identifiers_Identifier.
    Identifier []ExplicitPaths_Identifiers_Identifier
}

func (identifiers *ExplicitPaths_Identifiers) GetFilter() yfilter.YFilter { return identifiers.YFilter }

func (identifiers *ExplicitPaths_Identifiers) SetFilter(yf yfilter.YFilter) { identifiers.YFilter = yf }

func (identifiers *ExplicitPaths_Identifiers) GetGoName(yname string) string {
    if yname == "identifier" { return "Identifier" }
    return ""
}

func (identifiers *ExplicitPaths_Identifiers) GetSegmentPath() string {
    return "identifiers"
}

func (identifiers *ExplicitPaths_Identifiers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "identifier" {
        for _, c := range identifiers.Identifier {
            if identifiers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ExplicitPaths_Identifiers_Identifier{}
        identifiers.Identifier = append(identifiers.Identifier, child)
        return &identifiers.Identifier[len(identifiers.Identifier)-1]
    }
    return nil
}

func (identifiers *ExplicitPaths_Identifiers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range identifiers.Identifier {
        children[identifiers.Identifier[i].GetSegmentPath()] = &identifiers.Identifier[i]
    }
    return children
}

func (identifiers *ExplicitPaths_Identifiers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (identifiers *ExplicitPaths_Identifiers) GetBundleName() string { return "cisco_ios_xr" }

func (identifiers *ExplicitPaths_Identifiers) GetYangName() string { return "identifiers" }

func (identifiers *ExplicitPaths_Identifiers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (identifiers *ExplicitPaths_Identifiers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (identifiers *ExplicitPaths_Identifiers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (identifiers *ExplicitPaths_Identifiers) SetParent(parent types.Entity) { identifiers.parent = parent }

func (identifiers *ExplicitPaths_Identifiers) GetParent() types.Entity { return identifiers.parent }

func (identifiers *ExplicitPaths_Identifiers) GetParentYangName() string { return "explicit-paths" }

// ExplicitPaths_Identifiers_Identifier
// IP explicit path configured for a particular
// identifier
type ExplicitPaths_Identifiers_Identifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Identifier ID. The type is interface{} with range:
    // -2147483648..2147483647.
    IdentifierId interface{}

    // Status of the path. The type is IepStatus.
    Status interface{}

    // List of IP addresses configured in the explicit path. The type is slice of
    // ExplicitPaths_Identifiers_Identifier_Address.
    Address []ExplicitPaths_Identifiers_Identifier_Address
}

func (identifier *ExplicitPaths_Identifiers_Identifier) GetFilter() yfilter.YFilter { return identifier.YFilter }

func (identifier *ExplicitPaths_Identifiers_Identifier) SetFilter(yf yfilter.YFilter) { identifier.YFilter = yf }

func (identifier *ExplicitPaths_Identifiers_Identifier) GetGoName(yname string) string {
    if yname == "identifier-id" { return "IdentifierId" }
    if yname == "status" { return "Status" }
    if yname == "address" { return "Address" }
    return ""
}

func (identifier *ExplicitPaths_Identifiers_Identifier) GetSegmentPath() string {
    return "identifier" + "[identifier-id='" + fmt.Sprintf("%v", identifier.IdentifierId) + "']"
}

func (identifier *ExplicitPaths_Identifiers_Identifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range identifier.Address {
            if identifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ExplicitPaths_Identifiers_Identifier_Address{}
        identifier.Address = append(identifier.Address, child)
        return &identifier.Address[len(identifier.Address)-1]
    }
    return nil
}

func (identifier *ExplicitPaths_Identifiers_Identifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range identifier.Address {
        children[identifier.Address[i].GetSegmentPath()] = &identifier.Address[i]
    }
    return children
}

func (identifier *ExplicitPaths_Identifiers_Identifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["identifier-id"] = identifier.IdentifierId
    leafs["status"] = identifier.Status
    return leafs
}

func (identifier *ExplicitPaths_Identifiers_Identifier) GetBundleName() string { return "cisco_ios_xr" }

func (identifier *ExplicitPaths_Identifiers_Identifier) GetYangName() string { return "identifier" }

func (identifier *ExplicitPaths_Identifiers_Identifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (identifier *ExplicitPaths_Identifiers_Identifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (identifier *ExplicitPaths_Identifiers_Identifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (identifier *ExplicitPaths_Identifiers_Identifier) SetParent(parent types.Entity) { identifier.parent = parent }

func (identifier *ExplicitPaths_Identifiers_Identifier) GetParent() types.Entity { return identifier.parent }

func (identifier *ExplicitPaths_Identifiers_Identifier) GetParentYangName() string { return "identifiers" }

// ExplicitPaths_Identifiers_Identifier_Address
// List of IP addresses configured in the explicit
// path
type ExplicitPaths_Identifiers_Identifier_Address struct {
    parent types.Entity
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

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *ExplicitPaths_Identifiers_Identifier_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "if-index" { return "IfIndex" }
    if yname == "address-type" { return "AddressType" }
    if yname == "hop-type" { return "HopType" }
    if yname == "address" { return "Address" }
    if yname == "mpls-label" { return "MplsLabel" }
    return ""
}

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetSegmentPath() string {
    return "address"
}

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = address.Index
    leafs["if-index"] = address.IfIndex
    leafs["address-type"] = address.AddressType
    leafs["hop-type"] = address.HopType
    leafs["address"] = address.Address
    leafs["mpls-label"] = address.MplsLabel
    return leafs
}

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetYangName() string { return "address" }

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *ExplicitPaths_Identifiers_Identifier_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetParent() types.Entity { return address.parent }

func (address *ExplicitPaths_Identifiers_Identifier_Address) GetParentYangName() string { return "identifier" }

// ExplicitPaths_Names
// List of configured IP explicit path names, this
// corresponds to mplsTunnelHopTable in TE MIB
type ExplicitPaths_Names struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP explicit path configured for a particular path name. The type is slice
    // of ExplicitPaths_Names_Name.
    Name []ExplicitPaths_Names_Name
}

func (names *ExplicitPaths_Names) GetFilter() yfilter.YFilter { return names.YFilter }

func (names *ExplicitPaths_Names) SetFilter(yf yfilter.YFilter) { names.YFilter = yf }

func (names *ExplicitPaths_Names) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (names *ExplicitPaths_Names) GetSegmentPath() string {
    return "names"
}

func (names *ExplicitPaths_Names) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "name" {
        for _, c := range names.Name {
            if names.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ExplicitPaths_Names_Name{}
        names.Name = append(names.Name, child)
        return &names.Name[len(names.Name)-1]
    }
    return nil
}

func (names *ExplicitPaths_Names) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range names.Name {
        children[names.Name[i].GetSegmentPath()] = &names.Name[i]
    }
    return children
}

func (names *ExplicitPaths_Names) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (names *ExplicitPaths_Names) GetBundleName() string { return "cisco_ios_xr" }

func (names *ExplicitPaths_Names) GetYangName() string { return "names" }

func (names *ExplicitPaths_Names) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (names *ExplicitPaths_Names) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (names *ExplicitPaths_Names) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (names *ExplicitPaths_Names) SetParent(parent types.Entity) { names.parent = parent }

func (names *ExplicitPaths_Names) GetParent() types.Entity { return names.parent }

func (names *ExplicitPaths_Names) GetParentYangName() string { return "explicit-paths" }

// ExplicitPaths_Names_Name
// IP explicit path configured for a particular
// path name
type ExplicitPaths_Names_Name struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Path name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    PathName interface{}

    // Status of the path. The type is IepStatus.
    Status interface{}

    // List of IP addresses configured in the explicit path. The type is slice of
    // ExplicitPaths_Names_Name_Address.
    Address []ExplicitPaths_Names_Name_Address
}

func (name *ExplicitPaths_Names_Name) GetFilter() yfilter.YFilter { return name.YFilter }

func (name *ExplicitPaths_Names_Name) SetFilter(yf yfilter.YFilter) { name.YFilter = yf }

func (name *ExplicitPaths_Names_Name) GetGoName(yname string) string {
    if yname == "path-name" { return "PathName" }
    if yname == "status" { return "Status" }
    if yname == "address" { return "Address" }
    return ""
}

func (name *ExplicitPaths_Names_Name) GetSegmentPath() string {
    return "name" + "[path-name='" + fmt.Sprintf("%v", name.PathName) + "']"
}

func (name *ExplicitPaths_Names_Name) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range name.Address {
            if name.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ExplicitPaths_Names_Name_Address{}
        name.Address = append(name.Address, child)
        return &name.Address[len(name.Address)-1]
    }
    return nil
}

func (name *ExplicitPaths_Names_Name) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range name.Address {
        children[name.Address[i].GetSegmentPath()] = &name.Address[i]
    }
    return children
}

func (name *ExplicitPaths_Names_Name) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-name"] = name.PathName
    leafs["status"] = name.Status
    return leafs
}

func (name *ExplicitPaths_Names_Name) GetBundleName() string { return "cisco_ios_xr" }

func (name *ExplicitPaths_Names_Name) GetYangName() string { return "name" }

func (name *ExplicitPaths_Names_Name) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (name *ExplicitPaths_Names_Name) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (name *ExplicitPaths_Names_Name) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (name *ExplicitPaths_Names_Name) SetParent(parent types.Entity) { name.parent = parent }

func (name *ExplicitPaths_Names_Name) GetParent() types.Entity { return name.parent }

func (name *ExplicitPaths_Names_Name) GetParentYangName() string { return "names" }

// ExplicitPaths_Names_Name_Address
// List of IP addresses configured in the explicit
// path
type ExplicitPaths_Names_Name_Address struct {
    parent types.Entity
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

func (address *ExplicitPaths_Names_Name_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *ExplicitPaths_Names_Name_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *ExplicitPaths_Names_Name_Address) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "if-index" { return "IfIndex" }
    if yname == "address-type" { return "AddressType" }
    if yname == "hop-type" { return "HopType" }
    if yname == "address" { return "Address" }
    if yname == "mpls-label" { return "MplsLabel" }
    return ""
}

func (address *ExplicitPaths_Names_Name_Address) GetSegmentPath() string {
    return "address"
}

func (address *ExplicitPaths_Names_Name_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *ExplicitPaths_Names_Name_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *ExplicitPaths_Names_Name_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = address.Index
    leafs["if-index"] = address.IfIndex
    leafs["address-type"] = address.AddressType
    leafs["hop-type"] = address.HopType
    leafs["address"] = address.Address
    leafs["mpls-label"] = address.MplsLabel
    return leafs
}

func (address *ExplicitPaths_Names_Name_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *ExplicitPaths_Names_Name_Address) GetYangName() string { return "address" }

func (address *ExplicitPaths_Names_Name_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *ExplicitPaths_Names_Name_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *ExplicitPaths_Names_Name_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *ExplicitPaths_Names_Name_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *ExplicitPaths_Names_Name_Address) GetParent() types.Entity { return address.parent }

func (address *ExplicitPaths_Names_Name_Address) GetParentYangName() string { return "name" }

