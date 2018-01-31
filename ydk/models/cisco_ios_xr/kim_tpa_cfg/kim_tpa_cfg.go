// This module contains a collection of YANG definitions
// for Cisco IOS-XR kim-tpa package configuration.
// 
// This module contains definitions
// for the following management objects:
//   tpa: tpa configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package kim_tpa_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package kim_tpa_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-kim-tpa-cfg tpa}", reflect.TypeOf(Tpa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-kim-tpa-cfg:tpa", reflect.TypeOf(Tpa{}))
}

// Tpa
// tpa configuration commands
type Tpa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF container.
    VrfNames Tpa_VrfNames

    // Third party app logging.
    Logging Tpa_Logging

    // Statistics.
    Statistics Tpa_Statistics
}

func (tpa *Tpa) GetFilter() yfilter.YFilter { return tpa.YFilter }

func (tpa *Tpa) SetFilter(yf yfilter.YFilter) { tpa.YFilter = yf }

func (tpa *Tpa) GetGoName(yname string) string {
    if yname == "vrf-names" { return "VrfNames" }
    if yname == "logging" { return "Logging" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (tpa *Tpa) GetSegmentPath() string {
    return "Cisco-IOS-XR-kim-tpa-cfg:tpa"
}

func (tpa *Tpa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-names" {
        return &tpa.VrfNames
    }
    if childYangName == "logging" {
        return &tpa.Logging
    }
    if childYangName == "statistics" {
        return &tpa.Statistics
    }
    return nil
}

func (tpa *Tpa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrf-names"] = &tpa.VrfNames
    children["logging"] = &tpa.Logging
    children["statistics"] = &tpa.Statistics
    return children
}

func (tpa *Tpa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tpa *Tpa) GetBundleName() string { return "cisco_ios_xr" }

func (tpa *Tpa) GetYangName() string { return "tpa" }

func (tpa *Tpa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tpa *Tpa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tpa *Tpa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tpa *Tpa) SetParent(parent types.Entity) { tpa.parent = parent }

func (tpa *Tpa) GetParent() types.Entity { return tpa.parent }

func (tpa *Tpa) GetParentYangName() string { return "Cisco-IOS-XR-kim-tpa-cfg" }

// Tpa_VrfNames
// VRF container
type Tpa_VrfNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is slice of Tpa_VrfNames_VrfName.
    VrfName []Tpa_VrfNames_VrfName
}

func (vrfNames *Tpa_VrfNames) GetFilter() yfilter.YFilter { return vrfNames.YFilter }

func (vrfNames *Tpa_VrfNames) SetFilter(yf yfilter.YFilter) { vrfNames.YFilter = yf }

func (vrfNames *Tpa_VrfNames) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (vrfNames *Tpa_VrfNames) GetSegmentPath() string {
    return "vrf-names"
}

func (vrfNames *Tpa_VrfNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-name" {
        for _, c := range vrfNames.VrfName {
            if vrfNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Tpa_VrfNames_VrfName{}
        vrfNames.VrfName = append(vrfNames.VrfName, child)
        return &vrfNames.VrfName[len(vrfNames.VrfName)-1]
    }
    return nil
}

func (vrfNames *Tpa_VrfNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfNames.VrfName {
        children[vrfNames.VrfName[i].GetSegmentPath()] = &vrfNames.VrfName[i]
    }
    return children
}

func (vrfNames *Tpa_VrfNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfNames *Tpa_VrfNames) GetBundleName() string { return "cisco_ios_xr" }

func (vrfNames *Tpa_VrfNames) GetYangName() string { return "vrf-names" }

func (vrfNames *Tpa_VrfNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfNames *Tpa_VrfNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfNames *Tpa_VrfNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfNames *Tpa_VrfNames) SetParent(parent types.Entity) { vrfNames.parent = parent }

func (vrfNames *Tpa_VrfNames) GetParent() types.Entity { return vrfNames.parent }

func (vrfNames *Tpa_VrfNames) GetParentYangName() string { return "tpa" }

// Tpa_VrfNames_VrfName
// VRF name
type Tpa_VrfNames_VrfName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // EastWest container.
    EastWestNames Tpa_VrfNames_VrfName_EastWestNames

    // Address family.
    AddressFamily Tpa_VrfNames_VrfName_AddressFamily
}

func (vrfName *Tpa_VrfNames_VrfName) GetFilter() yfilter.YFilter { return vrfName.YFilter }

func (vrfName *Tpa_VrfNames_VrfName) SetFilter(yf yfilter.YFilter) { vrfName.YFilter = yf }

func (vrfName *Tpa_VrfNames_VrfName) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "east-west-names" { return "EastWestNames" }
    if yname == "address-family" { return "AddressFamily" }
    return ""
}

func (vrfName *Tpa_VrfNames_VrfName) GetSegmentPath() string {
    return "vrf-name" + "[vrf-name='" + fmt.Sprintf("%v", vrfName.VrfName) + "']"
}

func (vrfName *Tpa_VrfNames_VrfName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "east-west-names" {
        return &vrfName.EastWestNames
    }
    if childYangName == "address-family" {
        return &vrfName.AddressFamily
    }
    return nil
}

func (vrfName *Tpa_VrfNames_VrfName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["east-west-names"] = &vrfName.EastWestNames
    children["address-family"] = &vrfName.AddressFamily
    return children
}

func (vrfName *Tpa_VrfNames_VrfName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrfName.VrfName
    return leafs
}

func (vrfName *Tpa_VrfNames_VrfName) GetBundleName() string { return "cisco_ios_xr" }

func (vrfName *Tpa_VrfNames_VrfName) GetYangName() string { return "vrf-name" }

func (vrfName *Tpa_VrfNames_VrfName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfName *Tpa_VrfNames_VrfName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfName *Tpa_VrfNames_VrfName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfName *Tpa_VrfNames_VrfName) SetParent(parent types.Entity) { vrfName.parent = parent }

func (vrfName *Tpa_VrfNames_VrfName) GetParent() types.Entity { return vrfName.parent }

func (vrfName *Tpa_VrfNames_VrfName) GetParentYangName() string { return "vrf-names" }

// Tpa_VrfNames_VrfName_EastWestNames
// EastWest container
type Tpa_VrfNames_VrfName_EastWestNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // East West interface. The type is slice of
    // Tpa_VrfNames_VrfName_EastWestNames_EastWestName.
    EastWestName []Tpa_VrfNames_VrfName_EastWestNames_EastWestName
}

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetFilter() yfilter.YFilter { return eastWestNames.YFilter }

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) SetFilter(yf yfilter.YFilter) { eastWestNames.YFilter = yf }

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetGoName(yname string) string {
    if yname == "east-west-name" { return "EastWestName" }
    return ""
}

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetSegmentPath() string {
    return "east-west-names"
}

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "east-west-name" {
        for _, c := range eastWestNames.EastWestName {
            if eastWestNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Tpa_VrfNames_VrfName_EastWestNames_EastWestName{}
        eastWestNames.EastWestName = append(eastWestNames.EastWestName, child)
        return &eastWestNames.EastWestName[len(eastWestNames.EastWestName)-1]
    }
    return nil
}

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range eastWestNames.EastWestName {
        children[eastWestNames.EastWestName[i].GetSegmentPath()] = &eastWestNames.EastWestName[i]
    }
    return children
}

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetBundleName() string { return "cisco_ios_xr" }

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetYangName() string { return "east-west-names" }

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) SetParent(parent types.Entity) { eastWestNames.parent = parent }

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetParent() types.Entity { return eastWestNames.parent }

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetParentYangName() string { return "vrf-name" }

// Tpa_VrfNames_VrfName_EastWestNames_EastWestName
// East West interface
type Tpa_VrfNames_VrfName_EastWestNames_EastWestName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    EastWestName interface{}

    // VRF name. The type is string.
    Vrf interface{}

    // Interface. The type is string.
    Interface interface{}
}

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetFilter() yfilter.YFilter { return eastWestName.YFilter }

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) SetFilter(yf yfilter.YFilter) { eastWestName.YFilter = yf }

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetGoName(yname string) string {
    if yname == "east-west-name" { return "EastWestName" }
    if yname == "vrf" { return "Vrf" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetSegmentPath() string {
    return "east-west-name" + "[east-west-name='" + fmt.Sprintf("%v", eastWestName.EastWestName) + "']"
}

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["east-west-name"] = eastWestName.EastWestName
    leafs["vrf"] = eastWestName.Vrf
    leafs["interface"] = eastWestName.Interface
    return leafs
}

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetBundleName() string { return "cisco_ios_xr" }

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetYangName() string { return "east-west-name" }

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) SetParent(parent types.Entity) { eastWestName.parent = parent }

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetParent() types.Entity { return eastWestName.parent }

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetParentYangName() string { return "east-west-names" }

// Tpa_VrfNames_VrfName_AddressFamily
// Address family
type Tpa_VrfNames_VrfName_AddressFamily struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 configuration.
    Ipv6 Tpa_VrfNames_VrfName_AddressFamily_Ipv6

    // IPv4 configuration.
    Ipv4 Tpa_VrfNames_VrfName_AddressFamily_Ipv4
}

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetFilter() yfilter.YFilter { return addressFamily.YFilter }

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) SetFilter(yf yfilter.YFilter) { addressFamily.YFilter = yf }

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetGoName(yname string) string {
    if yname == "ipv6" { return "Ipv6" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetSegmentPath() string {
    return "address-family"
}

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &addressFamily.Ipv6
    }
    if childYangName == "ipv4" {
        return &addressFamily.Ipv4
    }
    return nil
}

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &addressFamily.Ipv6
    children["ipv4"] = &addressFamily.Ipv4
    return children
}

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetBundleName() string { return "cisco_ios_xr" }

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetYangName() string { return "address-family" }

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) SetParent(parent types.Entity) { addressFamily.parent = parent }

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetParent() types.Entity { return addressFamily.parent }

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetParentYangName() string { return "vrf-name" }

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6
// IPv6 configuration
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default interface used for routing. The type is string.
    DefaultRoute interface{}

    // Interface name for source address. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    UpdateSource interface{}

    // TPA Cli to configure LPTS entries.
    LptsAllowEntries Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries
}

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetGoName(yname string) string {
    if yname == "default-route" { return "DefaultRoute" }
    if yname == "update-source" { return "UpdateSource" }
    if yname == "lpts-allow-entries" { return "LptsAllowEntries" }
    return ""
}

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lpts-allow-entries" {
        return &ipv6.LptsAllowEntries
    }
    return nil
}

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lpts-allow-entries"] = &ipv6.LptsAllowEntries
    return children
}

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["default-route"] = ipv6.DefaultRoute
    leafs["update-source"] = ipv6.UpdateSource
    return leafs
}

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetParentYangName() string { return "address-family" }

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries
// TPA Cli to configure LPTS entries
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TPA Cli to configure LPTS entry. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry.
    LptsAllowEntry []Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry
}

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) GetFilter() yfilter.YFilter { return lptsAllowEntries.YFilter }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) SetFilter(yf yfilter.YFilter) { lptsAllowEntries.YFilter = yf }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) GetGoName(yname string) string {
    if yname == "lpts-allow-entry" { return "LptsAllowEntry" }
    return ""
}

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) GetSegmentPath() string {
    return "lpts-allow-entries"
}

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lpts-allow-entry" {
        for _, c := range lptsAllowEntries.LptsAllowEntry {
            if lptsAllowEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry{}
        lptsAllowEntries.LptsAllowEntry = append(lptsAllowEntries.LptsAllowEntry, child)
        return &lptsAllowEntries.LptsAllowEntry[len(lptsAllowEntries.LptsAllowEntry)-1]
    }
    return nil
}

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lptsAllowEntries.LptsAllowEntry {
        children[lptsAllowEntries.LptsAllowEntry[i].GetSegmentPath()] = &lptsAllowEntries.LptsAllowEntry[i]
    }
    return children
}

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) GetBundleName() string { return "cisco_ios_xr" }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) GetYangName() string { return "lpts-allow-entries" }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) SetParent(parent types.Entity) { lptsAllowEntries.parent = parent }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) GetParent() types.Entity { return lptsAllowEntries.parent }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries) GetParentYangName() string { return "ipv6" }

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry
// TPA Cli to configure LPTS entry
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    InterfaceName interface{}

    // This attribute is a key. remote address. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    RemoteAddr interface{}

    // This attribute is a key. local address. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    LocalAddr interface{}

    // This attribute is a key. remote port. The type is interface{} with range:
    // -2147483648..2147483647.
    RemotePort interface{}

    // This attribute is a key. local port. The type is interface{} with range:
    // -2147483648..2147483647.
    LocalPort interface{}

    // This attribute is a key. L4 protocol. The type is interface{} with range:
    // -2147483648..2147483647.
    Protocol interface{}

    // Interface name for allow command. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceNameXr interface{}

    // IPv4/6 remote-address prefix to match. The type is string.
    RemoteAddrXr interface{}

    // IPv4/6 local-address prefix to match. The type is string.
    LocalAddrXr interface{}

    // remote port value. The type is interface{} with range:
    // -2147483648..2147483647.
    RemotePortXr interface{}

    // local port value. The type is interface{} with range:
    // -2147483648..2147483647.
    LocalPortXr interface{}

    // L4 protocol value. The type is interface{} with range:
    // -2147483648..2147483647.
    ProtocolXr interface{}
}

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) GetFilter() yfilter.YFilter { return lptsAllowEntry.YFilter }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) SetFilter(yf yfilter.YFilter) { lptsAllowEntry.YFilter = yf }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "remote-addr" { return "RemoteAddr" }
    if yname == "local-addr" { return "LocalAddr" }
    if yname == "remote-port" { return "RemotePort" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "protocol" { return "Protocol" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "remote-addr-xr" { return "RemoteAddrXr" }
    if yname == "local-addr-xr" { return "LocalAddrXr" }
    if yname == "remote-port-xr" { return "RemotePortXr" }
    if yname == "local-port-xr" { return "LocalPortXr" }
    if yname == "protocol-xr" { return "ProtocolXr" }
    return ""
}

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) GetSegmentPath() string {
    return "lpts-allow-entry" + "[interface-name='" + fmt.Sprintf("%v", lptsAllowEntry.InterfaceName) + "']" + "[remote-addr='" + fmt.Sprintf("%v", lptsAllowEntry.RemoteAddr) + "']" + "[local-addr='" + fmt.Sprintf("%v", lptsAllowEntry.LocalAddr) + "']" + "[remote-port='" + fmt.Sprintf("%v", lptsAllowEntry.RemotePort) + "']" + "[local-port='" + fmt.Sprintf("%v", lptsAllowEntry.LocalPort) + "']" + "[protocol='" + fmt.Sprintf("%v", lptsAllowEntry.Protocol) + "']"
}

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = lptsAllowEntry.InterfaceName
    leafs["remote-addr"] = lptsAllowEntry.RemoteAddr
    leafs["local-addr"] = lptsAllowEntry.LocalAddr
    leafs["remote-port"] = lptsAllowEntry.RemotePort
    leafs["local-port"] = lptsAllowEntry.LocalPort
    leafs["protocol"] = lptsAllowEntry.Protocol
    leafs["interface-name-xr"] = lptsAllowEntry.InterfaceNameXr
    leafs["remote-addr-xr"] = lptsAllowEntry.RemoteAddrXr
    leafs["local-addr-xr"] = lptsAllowEntry.LocalAddrXr
    leafs["remote-port-xr"] = lptsAllowEntry.RemotePortXr
    leafs["local-port-xr"] = lptsAllowEntry.LocalPortXr
    leafs["protocol-xr"] = lptsAllowEntry.ProtocolXr
    return leafs
}

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) GetYangName() string { return "lpts-allow-entry" }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) SetParent(parent types.Entity) { lptsAllowEntry.parent = parent }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) GetParent() types.Entity { return lptsAllowEntry.parent }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_LptsAllowEntries_LptsAllowEntry) GetParentYangName() string { return "lpts-allow-entries" }

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4
// IPv4 configuration
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default interface used for routing. The type is string.
    DefaultRoute interface{}

    // Interface name for source address. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    UpdateSource interface{}

    // TPA Cli to configure LPTS entries.
    LptsAllowEntries Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries
}

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetGoName(yname string) string {
    if yname == "default-route" { return "DefaultRoute" }
    if yname == "update-source" { return "UpdateSource" }
    if yname == "lpts-allow-entries" { return "LptsAllowEntries" }
    return ""
}

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lpts-allow-entries" {
        return &ipv4.LptsAllowEntries
    }
    return nil
}

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lpts-allow-entries"] = &ipv4.LptsAllowEntries
    return children
}

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["default-route"] = ipv4.DefaultRoute
    leafs["update-source"] = ipv4.UpdateSource
    return leafs
}

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetParentYangName() string { return "address-family" }

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries
// TPA Cli to configure LPTS entries
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TPA Cli to configure LPTS entry. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry.
    LptsAllowEntry []Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry
}

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) GetFilter() yfilter.YFilter { return lptsAllowEntries.YFilter }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) SetFilter(yf yfilter.YFilter) { lptsAllowEntries.YFilter = yf }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) GetGoName(yname string) string {
    if yname == "lpts-allow-entry" { return "LptsAllowEntry" }
    return ""
}

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) GetSegmentPath() string {
    return "lpts-allow-entries"
}

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lpts-allow-entry" {
        for _, c := range lptsAllowEntries.LptsAllowEntry {
            if lptsAllowEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry{}
        lptsAllowEntries.LptsAllowEntry = append(lptsAllowEntries.LptsAllowEntry, child)
        return &lptsAllowEntries.LptsAllowEntry[len(lptsAllowEntries.LptsAllowEntry)-1]
    }
    return nil
}

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lptsAllowEntries.LptsAllowEntry {
        children[lptsAllowEntries.LptsAllowEntry[i].GetSegmentPath()] = &lptsAllowEntries.LptsAllowEntry[i]
    }
    return children
}

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) GetBundleName() string { return "cisco_ios_xr" }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) GetYangName() string { return "lpts-allow-entries" }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) SetParent(parent types.Entity) { lptsAllowEntries.parent = parent }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) GetParent() types.Entity { return lptsAllowEntries.parent }

func (lptsAllowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries) GetParentYangName() string { return "ipv4" }

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry
// TPA Cli to configure LPTS entry
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    InterfaceName interface{}

    // This attribute is a key. remote address. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    RemoteAddr interface{}

    // This attribute is a key. local address. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    LocalAddr interface{}

    // This attribute is a key. remote port. The type is interface{} with range:
    // -2147483648..2147483647.
    RemotePort interface{}

    // This attribute is a key. local port. The type is interface{} with range:
    // -2147483648..2147483647.
    LocalPort interface{}

    // This attribute is a key. L4 protocol. The type is interface{} with range:
    // -2147483648..2147483647.
    Protocol interface{}

    // Interface name for allow command. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceNameXr interface{}

    // IPv4/6 remote-address prefix to match. The type is string.
    RemoteAddrXr interface{}

    // IPv4/6 local-address prefix to match. The type is string.
    LocalAddrXr interface{}

    // remote port value. The type is interface{} with range:
    // -2147483648..2147483647.
    RemotePortXr interface{}

    // local port value. The type is interface{} with range:
    // -2147483648..2147483647.
    LocalPortXr interface{}

    // L4 protocol value. The type is interface{} with range:
    // -2147483648..2147483647.
    ProtocolXr interface{}
}

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) GetFilter() yfilter.YFilter { return lptsAllowEntry.YFilter }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) SetFilter(yf yfilter.YFilter) { lptsAllowEntry.YFilter = yf }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "remote-addr" { return "RemoteAddr" }
    if yname == "local-addr" { return "LocalAddr" }
    if yname == "remote-port" { return "RemotePort" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "protocol" { return "Protocol" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "remote-addr-xr" { return "RemoteAddrXr" }
    if yname == "local-addr-xr" { return "LocalAddrXr" }
    if yname == "remote-port-xr" { return "RemotePortXr" }
    if yname == "local-port-xr" { return "LocalPortXr" }
    if yname == "protocol-xr" { return "ProtocolXr" }
    return ""
}

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) GetSegmentPath() string {
    return "lpts-allow-entry" + "[interface-name='" + fmt.Sprintf("%v", lptsAllowEntry.InterfaceName) + "']" + "[remote-addr='" + fmt.Sprintf("%v", lptsAllowEntry.RemoteAddr) + "']" + "[local-addr='" + fmt.Sprintf("%v", lptsAllowEntry.LocalAddr) + "']" + "[remote-port='" + fmt.Sprintf("%v", lptsAllowEntry.RemotePort) + "']" + "[local-port='" + fmt.Sprintf("%v", lptsAllowEntry.LocalPort) + "']" + "[protocol='" + fmt.Sprintf("%v", lptsAllowEntry.Protocol) + "']"
}

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = lptsAllowEntry.InterfaceName
    leafs["remote-addr"] = lptsAllowEntry.RemoteAddr
    leafs["local-addr"] = lptsAllowEntry.LocalAddr
    leafs["remote-port"] = lptsAllowEntry.RemotePort
    leafs["local-port"] = lptsAllowEntry.LocalPort
    leafs["protocol"] = lptsAllowEntry.Protocol
    leafs["interface-name-xr"] = lptsAllowEntry.InterfaceNameXr
    leafs["remote-addr-xr"] = lptsAllowEntry.RemoteAddrXr
    leafs["local-addr-xr"] = lptsAllowEntry.LocalAddrXr
    leafs["remote-port-xr"] = lptsAllowEntry.RemotePortXr
    leafs["local-port-xr"] = lptsAllowEntry.LocalPortXr
    leafs["protocol-xr"] = lptsAllowEntry.ProtocolXr
    return leafs
}

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) GetYangName() string { return "lpts-allow-entry" }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) SetParent(parent types.Entity) { lptsAllowEntry.parent = parent }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) GetParent() types.Entity { return lptsAllowEntry.parent }

func (lptsAllowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_LptsAllowEntries_LptsAllowEntry) GetParentYangName() string { return "lpts-allow-entries" }

// Tpa_Logging
// Third party app logging
type Tpa_Logging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // KIM logging.
    Kim Tpa_Logging_Kim
}

func (logging *Tpa_Logging) GetFilter() yfilter.YFilter { return logging.YFilter }

func (logging *Tpa_Logging) SetFilter(yf yfilter.YFilter) { logging.YFilter = yf }

func (logging *Tpa_Logging) GetGoName(yname string) string {
    if yname == "kim" { return "Kim" }
    return ""
}

func (logging *Tpa_Logging) GetSegmentPath() string {
    return "logging"
}

func (logging *Tpa_Logging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "kim" {
        return &logging.Kim
    }
    return nil
}

func (logging *Tpa_Logging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["kim"] = &logging.Kim
    return children
}

func (logging *Tpa_Logging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (logging *Tpa_Logging) GetBundleName() string { return "cisco_ios_xr" }

func (logging *Tpa_Logging) GetYangName() string { return "logging" }

func (logging *Tpa_Logging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logging *Tpa_Logging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logging *Tpa_Logging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logging *Tpa_Logging) SetParent(parent types.Entity) { logging.parent = parent }

func (logging *Tpa_Logging) GetParent() types.Entity { return logging.parent }

func (logging *Tpa_Logging) GetParentYangName() string { return "tpa" }

// Tpa_Logging_Kim
// KIM logging
type Tpa_Logging_Kim struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // How many log rotation files to keep. The type is interface{} with range:
    // -2147483648..2147483647.
    RotationMax interface{}

    // Size in Kilobytes of the log file. The type is interface{} with range:
    // -2147483648..2147483647. Units are kilobyte.
    FileSizeMaxKb interface{}
}

func (kim *Tpa_Logging_Kim) GetFilter() yfilter.YFilter { return kim.YFilter }

func (kim *Tpa_Logging_Kim) SetFilter(yf yfilter.YFilter) { kim.YFilter = yf }

func (kim *Tpa_Logging_Kim) GetGoName(yname string) string {
    if yname == "rotation-max" { return "RotationMax" }
    if yname == "file-size-max-kb" { return "FileSizeMaxKb" }
    return ""
}

func (kim *Tpa_Logging_Kim) GetSegmentPath() string {
    return "kim"
}

func (kim *Tpa_Logging_Kim) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (kim *Tpa_Logging_Kim) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (kim *Tpa_Logging_Kim) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rotation-max"] = kim.RotationMax
    leafs["file-size-max-kb"] = kim.FileSizeMaxKb
    return leafs
}

func (kim *Tpa_Logging_Kim) GetBundleName() string { return "cisco_ios_xr" }

func (kim *Tpa_Logging_Kim) GetYangName() string { return "kim" }

func (kim *Tpa_Logging_Kim) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (kim *Tpa_Logging_Kim) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (kim *Tpa_Logging_Kim) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (kim *Tpa_Logging_Kim) SetParent(parent types.Entity) { kim.parent = parent }

func (kim *Tpa_Logging_Kim) GetParent() types.Entity { return kim.parent }

func (kim *Tpa_Logging_Kim) GetParentYangName() string { return "logging" }

// Tpa_Statistics
// Statistics
type Tpa_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // How many interface events to record. The type is interface{} with range:
    // -2147483648..2147483647.
    MaxIntfEvents interface{}

    // How many LPTS events to record. The type is interface{} with range:
    // -2147483648..2147483647.
    MaxLptsEvents interface{}

    // Statistics update frequency into Linux. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    StatisticsUpdateFrequency interface{}
}

func (statistics *Tpa_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Tpa_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Tpa_Statistics) GetGoName(yname string) string {
    if yname == "max-intf-events" { return "MaxIntfEvents" }
    if yname == "max-lpts-events" { return "MaxLptsEvents" }
    if yname == "statistics-update-frequency" { return "StatisticsUpdateFrequency" }
    return ""
}

func (statistics *Tpa_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Tpa_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Tpa_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Tpa_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-intf-events"] = statistics.MaxIntfEvents
    leafs["max-lpts-events"] = statistics.MaxLptsEvents
    leafs["statistics-update-frequency"] = statistics.StatisticsUpdateFrequency
    return leafs
}

func (statistics *Tpa_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Tpa_Statistics) GetYangName() string { return "statistics" }

func (statistics *Tpa_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Tpa_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Tpa_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Tpa_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Tpa_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Tpa_Statistics) GetParentYangName() string { return "tpa" }

