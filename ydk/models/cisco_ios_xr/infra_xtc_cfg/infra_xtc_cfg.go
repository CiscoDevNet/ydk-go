// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-xtc package configuration.
// 
// This module contains definitions
// for the following management objects:
//   pce: PCE configuration data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_xtc_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_xtc_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-xtc-cfg pce}", reflect.TypeOf(Pce{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-xtc-cfg:pce", reflect.TypeOf(Pce{}))
}

// PceExplicitPathHop represents Pce explicit path hop
type PceExplicitPathHop string

const (
    // Address
    PceExplicitPathHop_address PceExplicitPathHop = "address"

    // SID Node
    PceExplicitPathHop_sid_node PceExplicitPathHop = "sid-node"

    // SID Adjacency
    PceExplicitPathHop_sid_adjancency PceExplicitPathHop = "sid-adjancency"

    // Binding SID
    PceExplicitPathHop_binding_sid PceExplicitPathHop = "binding-sid"
)

// PceDisjointPath represents Pce disjoint path
type PceDisjointPath string

const (
    // Link
    PceDisjointPath_link PceDisjointPath = "link"

    // Node
    PceDisjointPath_node PceDisjointPath = "node"

    // SRLG
    PceDisjointPath_srlg PceDisjointPath = "srlg"
)

// Pce
// PCE configuration data
type Pce struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address of PCE server. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerAddress interface{}

    // MD5 password. The type is string with pattern: (!.+)|([^!].+).
    Password interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Path computation client configuration.
    PccAddresses Pce_PccAddresses

    // PCE logging configuration.
    Logging Pce_Logging

    // PCE backoff configuration.
    Backoff Pce_Backoff

    // Standby PCE configuration.
    StateSyncs Pce_StateSyncs

    // PCE segment-routing configuration.
    SegmentRouting Pce_SegmentRouting

    // PCE Timers configuration.
    Timers Pce_Timers

    // NETCONF configuration.
    Netconf Pce_Netconf

    // Disjoint path configuration.
    DisjointPath Pce_DisjointPath

    // Explicit paths.
    ExplicitPaths Pce_ExplicitPaths
}

func (pce *Pce) GetFilter() yfilter.YFilter { return pce.YFilter }

func (pce *Pce) SetFilter(yf yfilter.YFilter) { pce.YFilter = yf }

func (pce *Pce) GetGoName(yname string) string {
    if yname == "server-address" { return "ServerAddress" }
    if yname == "password" { return "Password" }
    if yname == "enable" { return "Enable" }
    if yname == "pcc-addresses" { return "PccAddresses" }
    if yname == "logging" { return "Logging" }
    if yname == "backoff" { return "Backoff" }
    if yname == "state-syncs" { return "StateSyncs" }
    if yname == "segment-routing" { return "SegmentRouting" }
    if yname == "timers" { return "Timers" }
    if yname == "netconf" { return "Netconf" }
    if yname == "disjoint-path" { return "DisjointPath" }
    if yname == "explicit-paths" { return "ExplicitPaths" }
    return ""
}

func (pce *Pce) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-xtc-cfg:pce"
}

func (pce *Pce) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pcc-addresses" {
        return &pce.PccAddresses
    }
    if childYangName == "logging" {
        return &pce.Logging
    }
    if childYangName == "backoff" {
        return &pce.Backoff
    }
    if childYangName == "state-syncs" {
        return &pce.StateSyncs
    }
    if childYangName == "segment-routing" {
        return &pce.SegmentRouting
    }
    if childYangName == "timers" {
        return &pce.Timers
    }
    if childYangName == "netconf" {
        return &pce.Netconf
    }
    if childYangName == "disjoint-path" {
        return &pce.DisjointPath
    }
    if childYangName == "explicit-paths" {
        return &pce.ExplicitPaths
    }
    return nil
}

func (pce *Pce) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pcc-addresses"] = &pce.PccAddresses
    children["logging"] = &pce.Logging
    children["backoff"] = &pce.Backoff
    children["state-syncs"] = &pce.StateSyncs
    children["segment-routing"] = &pce.SegmentRouting
    children["timers"] = &pce.Timers
    children["netconf"] = &pce.Netconf
    children["disjoint-path"] = &pce.DisjointPath
    children["explicit-paths"] = &pce.ExplicitPaths
    return children
}

func (pce *Pce) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-address"] = pce.ServerAddress
    leafs["password"] = pce.Password
    leafs["enable"] = pce.Enable
    return leafs
}

func (pce *Pce) GetBundleName() string { return "cisco_ios_xr" }

func (pce *Pce) GetYangName() string { return "pce" }

func (pce *Pce) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pce *Pce) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pce *Pce) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pce *Pce) SetParent(parent types.Entity) { pce.parent = parent }

func (pce *Pce) GetParent() types.Entity { return pce.parent }

func (pce *Pce) GetParentYangName() string { return "Cisco-IOS-XR-infra-xtc-cfg" }

// Pce_PccAddresses
// Path computation client configuration
type Pce_PccAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path computation client address. The type is slice of
    // Pce_PccAddresses_PccAddress.
    PccAddress []Pce_PccAddresses_PccAddress
}

func (pccAddresses *Pce_PccAddresses) GetFilter() yfilter.YFilter { return pccAddresses.YFilter }

func (pccAddresses *Pce_PccAddresses) SetFilter(yf yfilter.YFilter) { pccAddresses.YFilter = yf }

func (pccAddresses *Pce_PccAddresses) GetGoName(yname string) string {
    if yname == "pcc-address" { return "PccAddress" }
    return ""
}

func (pccAddresses *Pce_PccAddresses) GetSegmentPath() string {
    return "pcc-addresses"
}

func (pccAddresses *Pce_PccAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pcc-address" {
        for _, c := range pccAddresses.PccAddress {
            if pccAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_PccAddresses_PccAddress{}
        pccAddresses.PccAddress = append(pccAddresses.PccAddress, child)
        return &pccAddresses.PccAddress[len(pccAddresses.PccAddress)-1]
    }
    return nil
}

func (pccAddresses *Pce_PccAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pccAddresses.PccAddress {
        children[pccAddresses.PccAddress[i].GetSegmentPath()] = &pccAddresses.PccAddress[i]
    }
    return children
}

func (pccAddresses *Pce_PccAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pccAddresses *Pce_PccAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (pccAddresses *Pce_PccAddresses) GetYangName() string { return "pcc-addresses" }

func (pccAddresses *Pce_PccAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pccAddresses *Pce_PccAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pccAddresses *Pce_PccAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pccAddresses *Pce_PccAddresses) SetParent(parent types.Entity) { pccAddresses.parent = parent }

func (pccAddresses *Pce_PccAddresses) GetParent() types.Entity { return pccAddresses.parent }

func (pccAddresses *Pce_PccAddresses) GetParentYangName() string { return "pce" }

// Pce_PccAddresses_PccAddress
// Path computation client address
type Pce_PccAddresses_PccAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // MPLS label switched path.
    LspNames Pce_PccAddresses_PccAddress_LspNames
}

func (pccAddress *Pce_PccAddresses_PccAddress) GetFilter() yfilter.YFilter { return pccAddress.YFilter }

func (pccAddress *Pce_PccAddresses_PccAddress) SetFilter(yf yfilter.YFilter) { pccAddress.YFilter = yf }

func (pccAddress *Pce_PccAddresses_PccAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "enable" { return "Enable" }
    if yname == "lsp-names" { return "LspNames" }
    return ""
}

func (pccAddress *Pce_PccAddresses_PccAddress) GetSegmentPath() string {
    return "pcc-address" + "[address='" + fmt.Sprintf("%v", pccAddress.Address) + "']"
}

func (pccAddress *Pce_PccAddresses_PccAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-names" {
        return &pccAddress.LspNames
    }
    return nil
}

func (pccAddress *Pce_PccAddresses_PccAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lsp-names"] = &pccAddress.LspNames
    return children
}

func (pccAddress *Pce_PccAddresses_PccAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = pccAddress.Address
    leafs["enable"] = pccAddress.Enable
    return leafs
}

func (pccAddress *Pce_PccAddresses_PccAddress) GetBundleName() string { return "cisco_ios_xr" }

func (pccAddress *Pce_PccAddresses_PccAddress) GetYangName() string { return "pcc-address" }

func (pccAddress *Pce_PccAddresses_PccAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pccAddress *Pce_PccAddresses_PccAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pccAddress *Pce_PccAddresses_PccAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pccAddress *Pce_PccAddresses_PccAddress) SetParent(parent types.Entity) { pccAddress.parent = parent }

func (pccAddress *Pce_PccAddresses_PccAddress) GetParent() types.Entity { return pccAddress.parent }

func (pccAddress *Pce_PccAddresses_PccAddress) GetParentYangName() string { return "pcc-addresses" }

// Pce_PccAddresses_PccAddress_LspNames
// MPLS label switched path
type Pce_PccAddresses_PccAddress_LspNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS label switched path. The type is slice of
    // Pce_PccAddresses_PccAddress_LspNames_LspName.
    LspName []Pce_PccAddresses_PccAddress_LspNames_LspName
}

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetFilter() yfilter.YFilter { return lspNames.YFilter }

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) SetFilter(yf yfilter.YFilter) { lspNames.YFilter = yf }

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetGoName(yname string) string {
    if yname == "lsp-name" { return "LspName" }
    return ""
}

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetSegmentPath() string {
    return "lsp-names"
}

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-name" {
        for _, c := range lspNames.LspName {
            if lspNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_PccAddresses_PccAddress_LspNames_LspName{}
        lspNames.LspName = append(lspNames.LspName, child)
        return &lspNames.LspName[len(lspNames.LspName)-1]
    }
    return nil
}

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lspNames.LspName {
        children[lspNames.LspName[i].GetSegmentPath()] = &lspNames.LspName[i]
    }
    return children
}

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetBundleName() string { return "cisco_ios_xr" }

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetYangName() string { return "lsp-names" }

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) SetParent(parent types.Entity) { lspNames.parent = parent }

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetParent() types.Entity { return lspNames.parent }

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetParentYangName() string { return "pcc-address" }

// Pce_PccAddresses_PccAddress_LspNames_LspName
// MPLS label switched path
type Pce_PccAddresses_PccAddress_LspNames_LspName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LSP name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Undelegate LSP. The type is interface{}.
    Undelegate interface{}

    // Explicit-path name. The type is string.
    ExplicitPathName interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // RSVP-TE configuration.
    RsvpTe Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe
}

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetFilter() yfilter.YFilter { return lspName.YFilter }

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) SetFilter(yf yfilter.YFilter) { lspName.YFilter = yf }

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "undelegate" { return "Undelegate" }
    if yname == "explicit-path-name" { return "ExplicitPathName" }
    if yname == "enable" { return "Enable" }
    if yname == "rsvp-te" { return "RsvpTe" }
    return ""
}

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetSegmentPath() string {
    return "lsp-name" + "[name='" + fmt.Sprintf("%v", lspName.Name) + "']"
}

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rsvp-te" {
        return &lspName.RsvpTe
    }
    return nil
}

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rsvp-te"] = &lspName.RsvpTe
    return children
}

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = lspName.Name
    leafs["undelegate"] = lspName.Undelegate
    leafs["explicit-path-name"] = lspName.ExplicitPathName
    leafs["enable"] = lspName.Enable
    return leafs
}

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetBundleName() string { return "cisco_ios_xr" }

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetYangName() string { return "lsp-name" }

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) SetParent(parent types.Entity) { lspName.parent = parent }

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetParent() types.Entity { return lspName.parent }

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetParentYangName() string { return "lsp-names" }

// Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe
// RSVP-TE configuration
type Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable fast protection. The type is interface{}.
    FastProtect interface{}

    // Bandwidth configuration. The type is interface{} with range:
    // -2147483648..2147483647. Units are kbit/s.
    Bandwidth interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // LSP Affinity.
    Affinity Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity

    // Tunnel Setup and Hold Priorities.
    Priority Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority
}

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetFilter() yfilter.YFilter { return rsvpTe.YFilter }

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) SetFilter(yf yfilter.YFilter) { rsvpTe.YFilter = yf }

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetGoName(yname string) string {
    if yname == "fast-protect" { return "FastProtect" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "enable" { return "Enable" }
    if yname == "affinity" { return "Affinity" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetSegmentPath() string {
    return "rsvp-te"
}

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "affinity" {
        return &rsvpTe.Affinity
    }
    if childYangName == "priority" {
        return &rsvpTe.Priority
    }
    return nil
}

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["affinity"] = &rsvpTe.Affinity
    children["priority"] = &rsvpTe.Priority
    return children
}

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fast-protect"] = rsvpTe.FastProtect
    leafs["bandwidth"] = rsvpTe.Bandwidth
    leafs["enable"] = rsvpTe.Enable
    return leafs
}

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetBundleName() string { return "cisco_ios_xr" }

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetYangName() string { return "rsvp-te" }

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) SetParent(parent types.Entity) { rsvpTe.parent = parent }

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetParent() types.Entity { return rsvpTe.parent }

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetParentYangName() string { return "lsp-name" }

// Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity
// LSP Affinity
type Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Include-any affinity value. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    IncludeAny interface{}

    // Include-all affinity value. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    IncludeAll interface{}

    // Exclude-any affinity value. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    ExcludeAny interface{}
}

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetFilter() yfilter.YFilter { return affinity.YFilter }

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) SetFilter(yf yfilter.YFilter) { affinity.YFilter = yf }

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetGoName(yname string) string {
    if yname == "include-any" { return "IncludeAny" }
    if yname == "include-all" { return "IncludeAll" }
    if yname == "exclude-any" { return "ExcludeAny" }
    return ""
}

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetSegmentPath() string {
    return "affinity"
}

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["include-any"] = affinity.IncludeAny
    leafs["include-all"] = affinity.IncludeAll
    leafs["exclude-any"] = affinity.ExcludeAny
    return leafs
}

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetBundleName() string { return "cisco_ios_xr" }

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetYangName() string { return "affinity" }

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) SetParent(parent types.Entity) { affinity.parent = parent }

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetParent() types.Entity { return affinity.parent }

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetParentYangName() string { return "rsvp-te" }

// Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority
// Tunnel Setup and Hold Priorities
// This type is a presence type.
type Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Setup Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    SetupPriority interface{}

    // Hold Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    HoldPriority interface{}
}

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetFilter() yfilter.YFilter { return priority.YFilter }

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) SetFilter(yf yfilter.YFilter) { priority.YFilter = yf }

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetGoName(yname string) string {
    if yname == "setup-priority" { return "SetupPriority" }
    if yname == "hold-priority" { return "HoldPriority" }
    return ""
}

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetSegmentPath() string {
    return "priority"
}

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["setup-priority"] = priority.SetupPriority
    leafs["hold-priority"] = priority.HoldPriority
    return leafs
}

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetBundleName() string { return "cisco_ios_xr" }

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetYangName() string { return "priority" }

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) SetParent(parent types.Entity) { priority.parent = parent }

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetParent() types.Entity { return priority.parent }

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetParentYangName() string { return "rsvp-te" }

// Pce_Logging
// PCE logging configuration
type Pce_Logging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Logging NO-PATH configuration. The type is interface{}.
    NoPath interface{}

    // Logging fallback configuration. The type is interface{}.
    Fallback interface{}
}

func (logging *Pce_Logging) GetFilter() yfilter.YFilter { return logging.YFilter }

func (logging *Pce_Logging) SetFilter(yf yfilter.YFilter) { logging.YFilter = yf }

func (logging *Pce_Logging) GetGoName(yname string) string {
    if yname == "no-path" { return "NoPath" }
    if yname == "fallback" { return "Fallback" }
    return ""
}

func (logging *Pce_Logging) GetSegmentPath() string {
    return "logging"
}

func (logging *Pce_Logging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (logging *Pce_Logging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (logging *Pce_Logging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["no-path"] = logging.NoPath
    leafs["fallback"] = logging.Fallback
    return leafs
}

func (logging *Pce_Logging) GetBundleName() string { return "cisco_ios_xr" }

func (logging *Pce_Logging) GetYangName() string { return "logging" }

func (logging *Pce_Logging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logging *Pce_Logging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logging *Pce_Logging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logging *Pce_Logging) SetParent(parent types.Entity) { logging.parent = parent }

func (logging *Pce_Logging) GetParent() types.Entity { return logging.parent }

func (logging *Pce_Logging) GetParentYangName() string { return "pce" }

// Pce_Backoff
// PCE backoff configuration
type Pce_Backoff struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Backoff common ratio configuration. The type is interface{} with range:
    // 0..255. The default value is 2.
    Ratio interface{}

    // Backoff threshold configuration. The type is interface{} with range:
    // 0..3600. The default value is 0.
    Threshold interface{}

    // Backoff common difference configuration. The type is interface{} with
    // range: 0..255. The default value is 2.
    Difference interface{}
}

func (backoff *Pce_Backoff) GetFilter() yfilter.YFilter { return backoff.YFilter }

func (backoff *Pce_Backoff) SetFilter(yf yfilter.YFilter) { backoff.YFilter = yf }

func (backoff *Pce_Backoff) GetGoName(yname string) string {
    if yname == "ratio" { return "Ratio" }
    if yname == "threshold" { return "Threshold" }
    if yname == "difference" { return "Difference" }
    return ""
}

func (backoff *Pce_Backoff) GetSegmentPath() string {
    return "backoff"
}

func (backoff *Pce_Backoff) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (backoff *Pce_Backoff) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (backoff *Pce_Backoff) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ratio"] = backoff.Ratio
    leafs["threshold"] = backoff.Threshold
    leafs["difference"] = backoff.Difference
    return leafs
}

func (backoff *Pce_Backoff) GetBundleName() string { return "cisco_ios_xr" }

func (backoff *Pce_Backoff) GetYangName() string { return "backoff" }

func (backoff *Pce_Backoff) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backoff *Pce_Backoff) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backoff *Pce_Backoff) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backoff *Pce_Backoff) SetParent(parent types.Entity) { backoff.parent = parent }

func (backoff *Pce_Backoff) GetParent() types.Entity { return backoff.parent }

func (backoff *Pce_Backoff) GetParentYangName() string { return "pce" }

// Pce_StateSyncs
// Standby PCE configuration
type Pce_StateSyncs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Standby PCE ipv4 address. The type is slice of Pce_StateSyncs_StateSync.
    StateSync []Pce_StateSyncs_StateSync
}

func (stateSyncs *Pce_StateSyncs) GetFilter() yfilter.YFilter { return stateSyncs.YFilter }

func (stateSyncs *Pce_StateSyncs) SetFilter(yf yfilter.YFilter) { stateSyncs.YFilter = yf }

func (stateSyncs *Pce_StateSyncs) GetGoName(yname string) string {
    if yname == "state-sync" { return "StateSync" }
    return ""
}

func (stateSyncs *Pce_StateSyncs) GetSegmentPath() string {
    return "state-syncs"
}

func (stateSyncs *Pce_StateSyncs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state-sync" {
        for _, c := range stateSyncs.StateSync {
            if stateSyncs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_StateSyncs_StateSync{}
        stateSyncs.StateSync = append(stateSyncs.StateSync, child)
        return &stateSyncs.StateSync[len(stateSyncs.StateSync)-1]
    }
    return nil
}

func (stateSyncs *Pce_StateSyncs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stateSyncs.StateSync {
        children[stateSyncs.StateSync[i].GetSegmentPath()] = &stateSyncs.StateSync[i]
    }
    return children
}

func (stateSyncs *Pce_StateSyncs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stateSyncs *Pce_StateSyncs) GetBundleName() string { return "cisco_ios_xr" }

func (stateSyncs *Pce_StateSyncs) GetYangName() string { return "state-syncs" }

func (stateSyncs *Pce_StateSyncs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stateSyncs *Pce_StateSyncs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stateSyncs *Pce_StateSyncs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stateSyncs *Pce_StateSyncs) SetParent(parent types.Entity) { stateSyncs.parent = parent }

func (stateSyncs *Pce_StateSyncs) GetParent() types.Entity { return stateSyncs.parent }

func (stateSyncs *Pce_StateSyncs) GetParentYangName() string { return "pce" }

// Pce_StateSyncs_StateSync
// Standby PCE ipv4 address
type Pce_StateSyncs_StateSync struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (stateSync *Pce_StateSyncs_StateSync) GetFilter() yfilter.YFilter { return stateSync.YFilter }

func (stateSync *Pce_StateSyncs_StateSync) SetFilter(yf yfilter.YFilter) { stateSync.YFilter = yf }

func (stateSync *Pce_StateSyncs_StateSync) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (stateSync *Pce_StateSyncs_StateSync) GetSegmentPath() string {
    return "state-sync" + "[address='" + fmt.Sprintf("%v", stateSync.Address) + "']"
}

func (stateSync *Pce_StateSyncs_StateSync) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stateSync *Pce_StateSyncs_StateSync) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stateSync *Pce_StateSyncs_StateSync) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = stateSync.Address
    return leafs
}

func (stateSync *Pce_StateSyncs_StateSync) GetBundleName() string { return "cisco_ios_xr" }

func (stateSync *Pce_StateSyncs_StateSync) GetYangName() string { return "state-sync" }

func (stateSync *Pce_StateSyncs_StateSync) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stateSync *Pce_StateSyncs_StateSync) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stateSync *Pce_StateSyncs_StateSync) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stateSync *Pce_StateSyncs_StateSync) SetParent(parent types.Entity) { stateSync.parent = parent }

func (stateSync *Pce_StateSyncs_StateSync) GetParent() types.Entity { return stateSync.parent }

func (stateSync *Pce_StateSyncs_StateSync) GetParentYangName() string { return "state-syncs" }

// Pce_SegmentRouting
// PCE segment-routing configuration
type Pce_SegmentRouting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Use te-latency algorithm configuration. The type is interface{}.
    TeLatency interface{}

    // Use strict-sid-only configuration. The type is interface{}.
    StrictSidOnly interface{}
}

func (segmentRouting *Pce_SegmentRouting) GetFilter() yfilter.YFilter { return segmentRouting.YFilter }

func (segmentRouting *Pce_SegmentRouting) SetFilter(yf yfilter.YFilter) { segmentRouting.YFilter = yf }

func (segmentRouting *Pce_SegmentRouting) GetGoName(yname string) string {
    if yname == "te-latency" { return "TeLatency" }
    if yname == "strict-sid-only" { return "StrictSidOnly" }
    return ""
}

func (segmentRouting *Pce_SegmentRouting) GetSegmentPath() string {
    return "segment-routing"
}

func (segmentRouting *Pce_SegmentRouting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (segmentRouting *Pce_SegmentRouting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (segmentRouting *Pce_SegmentRouting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["te-latency"] = segmentRouting.TeLatency
    leafs["strict-sid-only"] = segmentRouting.StrictSidOnly
    return leafs
}

func (segmentRouting *Pce_SegmentRouting) GetBundleName() string { return "cisco_ios_xr" }

func (segmentRouting *Pce_SegmentRouting) GetYangName() string { return "segment-routing" }

func (segmentRouting *Pce_SegmentRouting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (segmentRouting *Pce_SegmentRouting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (segmentRouting *Pce_SegmentRouting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (segmentRouting *Pce_SegmentRouting) SetParent(parent types.Entity) { segmentRouting.parent = parent }

func (segmentRouting *Pce_SegmentRouting) GetParent() types.Entity { return segmentRouting.parent }

func (segmentRouting *Pce_SegmentRouting) GetParentYangName() string { return "pce" }

// Pce_Timers
// PCE Timers configuration
type Pce_Timers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Topology reoptimization timer configuration. The type is interface{} with
    // range: 10..3600. Units are second. The default value is 60.
    ReoptimizationTimer interface{}

    // Keepalive interval in seconds, zero to disable. The type is interface{}
    // with range: 0..255. Units are second. The default value is 30.
    Keepalive interface{}

    // Minimum acceptable peer proposed keepalive interval. The type is
    // interface{} with range: 0..255. Units are second. The default value is 20.
    MinimumPeerKeepalive interface{}
}

func (timers *Pce_Timers) GetFilter() yfilter.YFilter { return timers.YFilter }

func (timers *Pce_Timers) SetFilter(yf yfilter.YFilter) { timers.YFilter = yf }

func (timers *Pce_Timers) GetGoName(yname string) string {
    if yname == "reoptimization-timer" { return "ReoptimizationTimer" }
    if yname == "keepalive" { return "Keepalive" }
    if yname == "minimum-peer-keepalive" { return "MinimumPeerKeepalive" }
    return ""
}

func (timers *Pce_Timers) GetSegmentPath() string {
    return "timers"
}

func (timers *Pce_Timers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timers *Pce_Timers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timers *Pce_Timers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reoptimization-timer"] = timers.ReoptimizationTimer
    leafs["keepalive"] = timers.Keepalive
    leafs["minimum-peer-keepalive"] = timers.MinimumPeerKeepalive
    return leafs
}

func (timers *Pce_Timers) GetBundleName() string { return "cisco_ios_xr" }

func (timers *Pce_Timers) GetYangName() string { return "timers" }

func (timers *Pce_Timers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timers *Pce_Timers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timers *Pce_Timers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timers *Pce_Timers) SetParent(parent types.Entity) { timers.parent = parent }

func (timers *Pce_Timers) GetParent() types.Entity { return timers.parent }

func (timers *Pce_Timers) GetParentYangName() string { return "pce" }

// Pce_Netconf
// NETCONF configuration
type Pce_Netconf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NETCONF SSH configuration.
    NetconfSsh Pce_Netconf_NetconfSsh
}

func (netconf *Pce_Netconf) GetFilter() yfilter.YFilter { return netconf.YFilter }

func (netconf *Pce_Netconf) SetFilter(yf yfilter.YFilter) { netconf.YFilter = yf }

func (netconf *Pce_Netconf) GetGoName(yname string) string {
    if yname == "netconf-ssh" { return "NetconfSsh" }
    return ""
}

func (netconf *Pce_Netconf) GetSegmentPath() string {
    return "netconf"
}

func (netconf *Pce_Netconf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "netconf-ssh" {
        return &netconf.NetconfSsh
    }
    return nil
}

func (netconf *Pce_Netconf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["netconf-ssh"] = &netconf.NetconfSsh
    return children
}

func (netconf *Pce_Netconf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (netconf *Pce_Netconf) GetBundleName() string { return "cisco_ios_xr" }

func (netconf *Pce_Netconf) GetYangName() string { return "netconf" }

func (netconf *Pce_Netconf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (netconf *Pce_Netconf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (netconf *Pce_Netconf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (netconf *Pce_Netconf) SetParent(parent types.Entity) { netconf.parent = parent }

func (netconf *Pce_Netconf) GetParent() types.Entity { return netconf.parent }

func (netconf *Pce_Netconf) GetParentYangName() string { return "pce" }

// Pce_Netconf_NetconfSsh
// NETCONF SSH configuration
type Pce_Netconf_NetconfSsh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Password to use for NETCONF SSH connections. The type is string with
    // pattern: (!.+)|([^!].+).
    NetconfSshPassword interface{}

    // User name to use for NETCONF SSH connections. The type is string.
    NetconfSshUser interface{}
}

func (netconfSsh *Pce_Netconf_NetconfSsh) GetFilter() yfilter.YFilter { return netconfSsh.YFilter }

func (netconfSsh *Pce_Netconf_NetconfSsh) SetFilter(yf yfilter.YFilter) { netconfSsh.YFilter = yf }

func (netconfSsh *Pce_Netconf_NetconfSsh) GetGoName(yname string) string {
    if yname == "netconf-ssh-password" { return "NetconfSshPassword" }
    if yname == "netconf-ssh-user" { return "NetconfSshUser" }
    return ""
}

func (netconfSsh *Pce_Netconf_NetconfSsh) GetSegmentPath() string {
    return "netconf-ssh"
}

func (netconfSsh *Pce_Netconf_NetconfSsh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (netconfSsh *Pce_Netconf_NetconfSsh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (netconfSsh *Pce_Netconf_NetconfSsh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["netconf-ssh-password"] = netconfSsh.NetconfSshPassword
    leafs["netconf-ssh-user"] = netconfSsh.NetconfSshUser
    return leafs
}

func (netconfSsh *Pce_Netconf_NetconfSsh) GetBundleName() string { return "cisco_ios_xr" }

func (netconfSsh *Pce_Netconf_NetconfSsh) GetYangName() string { return "netconf-ssh" }

func (netconfSsh *Pce_Netconf_NetconfSsh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (netconfSsh *Pce_Netconf_NetconfSsh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (netconfSsh *Pce_Netconf_NetconfSsh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (netconfSsh *Pce_Netconf_NetconfSsh) SetParent(parent types.Entity) { netconfSsh.parent = parent }

func (netconfSsh *Pce_Netconf_NetconfSsh) GetParent() types.Entity { return netconfSsh.parent }

func (netconfSsh *Pce_Netconf_NetconfSsh) GetParentYangName() string { return "netconf" }

// Pce_DisjointPath
// Disjoint path configuration
type Pce_DisjointPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Association configuration.
    Groups Pce_DisjointPath_Groups
}

func (disjointPath *Pce_DisjointPath) GetFilter() yfilter.YFilter { return disjointPath.YFilter }

func (disjointPath *Pce_DisjointPath) SetFilter(yf yfilter.YFilter) { disjointPath.YFilter = yf }

func (disjointPath *Pce_DisjointPath) GetGoName(yname string) string {
    if yname == "groups" { return "Groups" }
    return ""
}

func (disjointPath *Pce_DisjointPath) GetSegmentPath() string {
    return "disjoint-path"
}

func (disjointPath *Pce_DisjointPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groups" {
        return &disjointPath.Groups
    }
    return nil
}

func (disjointPath *Pce_DisjointPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["groups"] = &disjointPath.Groups
    return children
}

func (disjointPath *Pce_DisjointPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (disjointPath *Pce_DisjointPath) GetBundleName() string { return "cisco_ios_xr" }

func (disjointPath *Pce_DisjointPath) GetYangName() string { return "disjoint-path" }

func (disjointPath *Pce_DisjointPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (disjointPath *Pce_DisjointPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (disjointPath *Pce_DisjointPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (disjointPath *Pce_DisjointPath) SetParent(parent types.Entity) { disjointPath.parent = parent }

func (disjointPath *Pce_DisjointPath) GetParent() types.Entity { return disjointPath.parent }

func (disjointPath *Pce_DisjointPath) GetParentYangName() string { return "pce" }

// Pce_DisjointPath_Groups
// Association configuration
type Pce_DisjointPath_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Association Group Configuration. The type is slice of
    // Pce_DisjointPath_Groups_Group.
    Group []Pce_DisjointPath_Groups_Group
}

func (groups *Pce_DisjointPath_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *Pce_DisjointPath_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *Pce_DisjointPath_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *Pce_DisjointPath_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *Pce_DisjointPath_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_DisjointPath_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *Pce_DisjointPath_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *Pce_DisjointPath_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *Pce_DisjointPath_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *Pce_DisjointPath_Groups) GetYangName() string { return "groups" }

func (groups *Pce_DisjointPath_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *Pce_DisjointPath_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *Pce_DisjointPath_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *Pce_DisjointPath_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *Pce_DisjointPath_Groups) GetParent() types.Entity { return groups.parent }

func (groups *Pce_DisjointPath_Groups) GetParentYangName() string { return "disjoint-path" }

// Pce_DisjointPath_Groups_Group
// Association Group Configuration
type Pce_DisjointPath_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group ID. The type is interface{} with range:
    // 1..65535.
    GroupId interface{}

    // This attribute is a key. Disjoiness type. The type is PceDisjointPath.
    DpType interface{}

    // This attribute is a key. Sub group ID, 0 to unset. The type is interface{}
    // with range: 0..65535.
    SubId interface{}

    // Disable Fallback. The type is interface{}.
    Strict interface{}
}

func (group *Pce_DisjointPath_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *Pce_DisjointPath_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *Pce_DisjointPath_Groups_Group) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "dp-type" { return "DpType" }
    if yname == "sub-id" { return "SubId" }
    if yname == "strict" { return "Strict" }
    return ""
}

func (group *Pce_DisjointPath_Groups_Group) GetSegmentPath() string {
    return "group" + "[group-id='" + fmt.Sprintf("%v", group.GroupId) + "']" + "[dp-type='" + fmt.Sprintf("%v", group.DpType) + "']" + "[sub-id='" + fmt.Sprintf("%v", group.SubId) + "']"
}

func (group *Pce_DisjointPath_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (group *Pce_DisjointPath_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (group *Pce_DisjointPath_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = group.GroupId
    leafs["dp-type"] = group.DpType
    leafs["sub-id"] = group.SubId
    leafs["strict"] = group.Strict
    return leafs
}

func (group *Pce_DisjointPath_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *Pce_DisjointPath_Groups_Group) GetYangName() string { return "group" }

func (group *Pce_DisjointPath_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *Pce_DisjointPath_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *Pce_DisjointPath_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *Pce_DisjointPath_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *Pce_DisjointPath_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *Pce_DisjointPath_Groups_Group) GetParentYangName() string { return "groups" }

// Pce_ExplicitPaths
// Explicit paths
type Pce_ExplicitPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Explicit-path configuration. The type is slice of
    // Pce_ExplicitPaths_ExplicitPath.
    ExplicitPath []Pce_ExplicitPaths_ExplicitPath
}

func (explicitPaths *Pce_ExplicitPaths) GetFilter() yfilter.YFilter { return explicitPaths.YFilter }

func (explicitPaths *Pce_ExplicitPaths) SetFilter(yf yfilter.YFilter) { explicitPaths.YFilter = yf }

func (explicitPaths *Pce_ExplicitPaths) GetGoName(yname string) string {
    if yname == "explicit-path" { return "ExplicitPath" }
    return ""
}

func (explicitPaths *Pce_ExplicitPaths) GetSegmentPath() string {
    return "explicit-paths"
}

func (explicitPaths *Pce_ExplicitPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "explicit-path" {
        for _, c := range explicitPaths.ExplicitPath {
            if explicitPaths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_ExplicitPaths_ExplicitPath{}
        explicitPaths.ExplicitPath = append(explicitPaths.ExplicitPath, child)
        return &explicitPaths.ExplicitPath[len(explicitPaths.ExplicitPath)-1]
    }
    return nil
}

func (explicitPaths *Pce_ExplicitPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range explicitPaths.ExplicitPath {
        children[explicitPaths.ExplicitPath[i].GetSegmentPath()] = &explicitPaths.ExplicitPath[i]
    }
    return children
}

func (explicitPaths *Pce_ExplicitPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (explicitPaths *Pce_ExplicitPaths) GetBundleName() string { return "cisco_ios_xr" }

func (explicitPaths *Pce_ExplicitPaths) GetYangName() string { return "explicit-paths" }

func (explicitPaths *Pce_ExplicitPaths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitPaths *Pce_ExplicitPaths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitPaths *Pce_ExplicitPaths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitPaths *Pce_ExplicitPaths) SetParent(parent types.Entity) { explicitPaths.parent = parent }

func (explicitPaths *Pce_ExplicitPaths) GetParent() types.Entity { return explicitPaths.parent }

func (explicitPaths *Pce_ExplicitPaths) GetParentYangName() string { return "pce" }

// Pce_ExplicitPaths_ExplicitPath
// Explicit-path configuration
type Pce_ExplicitPaths_ExplicitPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Explicit-path name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Path Hops.
    PathHops Pce_ExplicitPaths_ExplicitPath_PathHops
}

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetFilter() yfilter.YFilter { return explicitPath.YFilter }

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) SetFilter(yf yfilter.YFilter) { explicitPath.YFilter = yf }

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "enable" { return "Enable" }
    if yname == "path-hops" { return "PathHops" }
    return ""
}

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetSegmentPath() string {
    return "explicit-path" + "[name='" + fmt.Sprintf("%v", explicitPath.Name) + "']"
}

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-hops" {
        return &explicitPath.PathHops
    }
    return nil
}

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-hops"] = &explicitPath.PathHops
    return children
}

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = explicitPath.Name
    leafs["enable"] = explicitPath.Enable
    return leafs
}

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetBundleName() string { return "cisco_ios_xr" }

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetYangName() string { return "explicit-path" }

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) SetParent(parent types.Entity) { explicitPath.parent = parent }

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetParent() types.Entity { return explicitPath.parent }

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetParentYangName() string { return "explicit-paths" }

// Pce_ExplicitPaths_ExplicitPath_PathHops
// Path Hops
type Pce_ExplicitPaths_ExplicitPath_PathHops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Explicit path hop configuration. The type is slice of
    // Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop.
    PathHop []Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop
}

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetFilter() yfilter.YFilter { return pathHops.YFilter }

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) SetFilter(yf yfilter.YFilter) { pathHops.YFilter = yf }

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetGoName(yname string) string {
    if yname == "path-hop" { return "PathHop" }
    return ""
}

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetSegmentPath() string {
    return "path-hops"
}

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-hop" {
        for _, c := range pathHops.PathHop {
            if pathHops.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop{}
        pathHops.PathHop = append(pathHops.PathHop, child)
        return &pathHops.PathHop[len(pathHops.PathHop)-1]
    }
    return nil
}

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pathHops.PathHop {
        children[pathHops.PathHop[i].GetSegmentPath()] = &pathHops.PathHop[i]
    }
    return children
}

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetBundleName() string { return "cisco_ios_xr" }

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetYangName() string { return "path-hops" }

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) SetParent(parent types.Entity) { pathHops.parent = parent }

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetParent() types.Entity { return pathHops.parent }

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetParentYangName() string { return "explicit-path" }

// Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop
// Explicit path hop configuration
type Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Hop Index. The type is interface{} with range:
    // 1..65535.
    Index interface{}

    // Path hop type. The type is PceExplicitPathHop.
    HopType interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // The default value is 0.0.0.0.
    Address interface{}

    // Remote IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // The default value is 0.0.0.0.
    RemoteAddress interface{}

    // MPLS Label. The type is interface{} with range: 0..1048575. The default
    // value is 0.
    MplsLabel interface{}
}

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetFilter() yfilter.YFilter { return pathHop.YFilter }

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) SetFilter(yf yfilter.YFilter) { pathHop.YFilter = yf }

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "hop-type" { return "HopType" }
    if yname == "address" { return "Address" }
    if yname == "remote-address" { return "RemoteAddress" }
    if yname == "mpls-label" { return "MplsLabel" }
    return ""
}

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetSegmentPath() string {
    return "path-hop" + "[index='" + fmt.Sprintf("%v", pathHop.Index) + "']"
}

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = pathHop.Index
    leafs["hop-type"] = pathHop.HopType
    leafs["address"] = pathHop.Address
    leafs["remote-address"] = pathHop.RemoteAddress
    leafs["mpls-label"] = pathHop.MplsLabel
    return leafs
}

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetBundleName() string { return "cisco_ios_xr" }

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetYangName() string { return "path-hop" }

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) SetParent(parent types.Entity) { pathHop.parent = parent }

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetParent() types.Entity { return pathHop.parent }

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetParentYangName() string { return "path-hops" }

