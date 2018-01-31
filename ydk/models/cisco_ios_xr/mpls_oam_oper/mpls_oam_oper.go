// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-oam package operational data.
// 
// This module contains definitions
// for the following management objects:
//   mpls-oam: MPLS OAM operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package mpls_oam_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_oam_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mpls-oam-oper mpls-oam}", reflect.TypeOf(MplsOam{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mpls-oam-oper:mpls-oam", reflect.TypeOf(MplsOam{}))
}

// LspvBagInterfaceState represents LSPV interface state
type LspvBagInterfaceState string

const (
    // Not ready
    LspvBagInterfaceState_not_ready LspvBagInterfaceState = "not-ready"

    // Admin down
    LspvBagInterfaceState_admin_down LspvBagInterfaceState = "admin-down"

    // Down
    LspvBagInterfaceState_down LspvBagInterfaceState = "down"

    // UP
    LspvBagInterfaceState_up LspvBagInterfaceState = "up"

    // Shutdown
    LspvBagInterfaceState_shutdown LspvBagInterfaceState = "shutdown"

    // Error disable
    LspvBagInterfaceState_error_disable LspvBagInterfaceState = "error-disable"

    // Immediate down
    LspvBagInterfaceState_down_immediate LspvBagInterfaceState = "down-immediate"

    // Immediate admin
    LspvBagInterfaceState_admin_immediate LspvBagInterfaceState = "admin-immediate"

    // Graceful down
    LspvBagInterfaceState_graceful_down LspvBagInterfaceState = "graceful-down"

    // Begin shutdown
    LspvBagInterfaceState_begin_shutdown LspvBagInterfaceState = "begin-shutdown"

    // End shutdown
    LspvBagInterfaceState_end_shutdown LspvBagInterfaceState = "end-shutdown"

    // Begin error disable
    LspvBagInterfaceState_begin_error_disable LspvBagInterfaceState = "begin-error-disable"

    // End error disable
    LspvBagInterfaceState_end_error_disable LspvBagInterfaceState = "end-error-disable"

    // Begin graceful down
    LspvBagInterfaceState_begin_graceful_down LspvBagInterfaceState = "begin-graceful-down"

    // Reset
    LspvBagInterfaceState_reset LspvBagInterfaceState = "reset"

    // Operational
    LspvBagInterfaceState_operational LspvBagInterfaceState = "operational"

    // Not operational
    LspvBagInterfaceState_not_operational LspvBagInterfaceState = "not-operational"

    // Unknown
    LspvBagInterfaceState_not_known LspvBagInterfaceState = "not-known"
)

// MplsOam
// MPLS OAM operational data
type MplsOam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS OAM interface operational data.
    Interface MplsOam_Interface

    // LSPV packet counters operational data.
    Packet MplsOam_Packet

    // LSPV global counters operational data.
    Global MplsOam_Global
}

func (mplsOam *MplsOam) GetFilter() yfilter.YFilter { return mplsOam.YFilter }

func (mplsOam *MplsOam) SetFilter(yf yfilter.YFilter) { mplsOam.YFilter = yf }

func (mplsOam *MplsOam) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "packet" { return "Packet" }
    if yname == "global" { return "Global" }
    return ""
}

func (mplsOam *MplsOam) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-oam-oper:mpls-oam"
}

func (mplsOam *MplsOam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        return &mplsOam.Interface
    }
    if childYangName == "packet" {
        return &mplsOam.Packet
    }
    if childYangName == "global" {
        return &mplsOam.Global
    }
    return nil
}

func (mplsOam *MplsOam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface"] = &mplsOam.Interface
    children["packet"] = &mplsOam.Packet
    children["global"] = &mplsOam.Global
    return children
}

func (mplsOam *MplsOam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsOam *MplsOam) GetBundleName() string { return "cisco_ios_xr" }

func (mplsOam *MplsOam) GetYangName() string { return "mpls-oam" }

func (mplsOam *MplsOam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsOam *MplsOam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsOam *MplsOam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsOam *MplsOam) SetParent(parent types.Entity) { mplsOam.parent = parent }

func (mplsOam *MplsOam) GetParent() types.Entity { return mplsOam.parent }

func (mplsOam *MplsOam) GetParentYangName() string { return "Cisco-IOS-XR-mpls-oam-oper" }

// MplsOam_Interface
// MPLS OAM interface operational data
type MplsOam_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS OAM interface detail data.
    Briefs MplsOam_Interface_Briefs

    // MPLS OAM interface detail data.
    Details MplsOam_Interface_Details
}

func (self *MplsOam_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *MplsOam_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *MplsOam_Interface) GetGoName(yname string) string {
    if yname == "briefs" { return "Briefs" }
    if yname == "details" { return "Details" }
    return ""
}

func (self *MplsOam_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *MplsOam_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "briefs" {
        return &self.Briefs
    }
    if childYangName == "details" {
        return &self.Details
    }
    return nil
}

func (self *MplsOam_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["briefs"] = &self.Briefs
    children["details"] = &self.Details
    return children
}

func (self *MplsOam_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (self *MplsOam_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *MplsOam_Interface) GetYangName() string { return "interface" }

func (self *MplsOam_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *MplsOam_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *MplsOam_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *MplsOam_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *MplsOam_Interface) GetParent() types.Entity { return self.parent }

func (self *MplsOam_Interface) GetParentYangName() string { return "mpls-oam" }

// MplsOam_Interface_Briefs
// MPLS OAM interface detail data
type MplsOam_Interface_Briefs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS OAM interface operational data. The type is slice of
    // MplsOam_Interface_Briefs_Brief.
    Brief []MplsOam_Interface_Briefs_Brief
}

func (briefs *MplsOam_Interface_Briefs) GetFilter() yfilter.YFilter { return briefs.YFilter }

func (briefs *MplsOam_Interface_Briefs) SetFilter(yf yfilter.YFilter) { briefs.YFilter = yf }

func (briefs *MplsOam_Interface_Briefs) GetGoName(yname string) string {
    if yname == "brief" { return "Brief" }
    return ""
}

func (briefs *MplsOam_Interface_Briefs) GetSegmentPath() string {
    return "briefs"
}

func (briefs *MplsOam_Interface_Briefs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief" {
        for _, c := range briefs.Brief {
            if briefs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsOam_Interface_Briefs_Brief{}
        briefs.Brief = append(briefs.Brief, child)
        return &briefs.Brief[len(briefs.Brief)-1]
    }
    return nil
}

func (briefs *MplsOam_Interface_Briefs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range briefs.Brief {
        children[briefs.Brief[i].GetSegmentPath()] = &briefs.Brief[i]
    }
    return children
}

func (briefs *MplsOam_Interface_Briefs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (briefs *MplsOam_Interface_Briefs) GetBundleName() string { return "cisco_ios_xr" }

func (briefs *MplsOam_Interface_Briefs) GetYangName() string { return "briefs" }

func (briefs *MplsOam_Interface_Briefs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefs *MplsOam_Interface_Briefs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefs *MplsOam_Interface_Briefs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefs *MplsOam_Interface_Briefs) SetParent(parent types.Entity) { briefs.parent = parent }

func (briefs *MplsOam_Interface_Briefs) GetParent() types.Entity { return briefs.parent }

func (briefs *MplsOam_Interface_Briefs) GetParentYangName() string { return "interface" }

// MplsOam_Interface_Briefs_Brief
// MPLS OAM interface operational data
type MplsOam_Interface_Briefs_Brief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceNameXr interface{}

    // Interface state. The type is LspvBagInterfaceState.
    State interface{}

    // Interface MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Prefix length (IPv4). The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // Prefix length (IPv6). The type is interface{} with range: 0..4294967295.
    PrefixLengthV6 interface{}

    // Primary interface address (IPv4). The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrimaryAddress interface{}

    // Primary interface address (IPv6). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrimaryAddressV6 interface{}
}

func (brief *MplsOam_Interface_Briefs_Brief) GetFilter() yfilter.YFilter { return brief.YFilter }

func (brief *MplsOam_Interface_Briefs_Brief) SetFilter(yf yfilter.YFilter) { brief.YFilter = yf }

func (brief *MplsOam_Interface_Briefs_Brief) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "state" { return "State" }
    if yname == "mtu" { return "Mtu" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix-length-v6" { return "PrefixLengthV6" }
    if yname == "primary-address" { return "PrimaryAddress" }
    if yname == "primary-address-v6" { return "PrimaryAddressV6" }
    return ""
}

func (brief *MplsOam_Interface_Briefs_Brief) GetSegmentPath() string {
    return "brief" + "[interface-name='" + fmt.Sprintf("%v", brief.InterfaceName) + "']"
}

func (brief *MplsOam_Interface_Briefs_Brief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (brief *MplsOam_Interface_Briefs_Brief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (brief *MplsOam_Interface_Briefs_Brief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = brief.InterfaceName
    leafs["interface-name-xr"] = brief.InterfaceNameXr
    leafs["state"] = brief.State
    leafs["mtu"] = brief.Mtu
    leafs["prefix-length"] = brief.PrefixLength
    leafs["prefix-length-v6"] = brief.PrefixLengthV6
    leafs["primary-address"] = brief.PrimaryAddress
    leafs["primary-address-v6"] = brief.PrimaryAddressV6
    return leafs
}

func (brief *MplsOam_Interface_Briefs_Brief) GetBundleName() string { return "cisco_ios_xr" }

func (brief *MplsOam_Interface_Briefs_Brief) GetYangName() string { return "brief" }

func (brief *MplsOam_Interface_Briefs_Brief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (brief *MplsOam_Interface_Briefs_Brief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (brief *MplsOam_Interface_Briefs_Brief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (brief *MplsOam_Interface_Briefs_Brief) SetParent(parent types.Entity) { brief.parent = parent }

func (brief *MplsOam_Interface_Briefs_Brief) GetParent() types.Entity { return brief.parent }

func (brief *MplsOam_Interface_Briefs_Brief) GetParentYangName() string { return "briefs" }

// MplsOam_Interface_Details
// MPLS OAM interface detail data
type MplsOam_Interface_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS OAM interface operational data. The type is slice of
    // MplsOam_Interface_Details_Detail.
    Detail []MplsOam_Interface_Details_Detail
}

func (details *MplsOam_Interface_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *MplsOam_Interface_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *MplsOam_Interface_Details) GetGoName(yname string) string {
    if yname == "detail" { return "Detail" }
    return ""
}

func (details *MplsOam_Interface_Details) GetSegmentPath() string {
    return "details"
}

func (details *MplsOam_Interface_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        for _, c := range details.Detail {
            if details.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsOam_Interface_Details_Detail{}
        details.Detail = append(details.Detail, child)
        return &details.Detail[len(details.Detail)-1]
    }
    return nil
}

func (details *MplsOam_Interface_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range details.Detail {
        children[details.Detail[i].GetSegmentPath()] = &details.Detail[i]
    }
    return children
}

func (details *MplsOam_Interface_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (details *MplsOam_Interface_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *MplsOam_Interface_Details) GetYangName() string { return "details" }

func (details *MplsOam_Interface_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *MplsOam_Interface_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *MplsOam_Interface_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *MplsOam_Interface_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *MplsOam_Interface_Details) GetParent() types.Entity { return details.parent }

func (details *MplsOam_Interface_Details) GetParentYangName() string { return "interface" }

// MplsOam_Interface_Details_Detail
// MPLS OAM interface operational data
type MplsOam_Interface_Details_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface brief.
    InterfaceBrief MplsOam_Interface_Details_Detail_InterfaceBrief

    // Packet statistics.
    PacketStatistics MplsOam_Interface_Details_Detail_PacketStatistics
}

func (detail *MplsOam_Interface_Details_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *MplsOam_Interface_Details_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *MplsOam_Interface_Details_Detail) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-brief" { return "InterfaceBrief" }
    if yname == "packet-statistics" { return "PacketStatistics" }
    return ""
}

func (detail *MplsOam_Interface_Details_Detail) GetSegmentPath() string {
    return "detail" + "[interface-name='" + fmt.Sprintf("%v", detail.InterfaceName) + "']"
}

func (detail *MplsOam_Interface_Details_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-brief" {
        return &detail.InterfaceBrief
    }
    if childYangName == "packet-statistics" {
        return &detail.PacketStatistics
    }
    return nil
}

func (detail *MplsOam_Interface_Details_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-brief"] = &detail.InterfaceBrief
    children["packet-statistics"] = &detail.PacketStatistics
    return children
}

func (detail *MplsOam_Interface_Details_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = detail.InterfaceName
    return leafs
}

func (detail *MplsOam_Interface_Details_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *MplsOam_Interface_Details_Detail) GetYangName() string { return "detail" }

func (detail *MplsOam_Interface_Details_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *MplsOam_Interface_Details_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *MplsOam_Interface_Details_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *MplsOam_Interface_Details_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *MplsOam_Interface_Details_Detail) GetParent() types.Entity { return detail.parent }

func (detail *MplsOam_Interface_Details_Detail) GetParentYangName() string { return "details" }

// MplsOam_Interface_Details_Detail_InterfaceBrief
// Interface brief
type MplsOam_Interface_Details_Detail_InterfaceBrief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceNameXr interface{}

    // Interface state. The type is LspvBagInterfaceState.
    State interface{}

    // Interface MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Prefix length (IPv4). The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // Prefix length (IPv6). The type is interface{} with range: 0..4294967295.
    PrefixLengthV6 interface{}

    // Primary interface address (IPv4). The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrimaryAddress interface{}

    // Primary interface address (IPv6). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrimaryAddressV6 interface{}
}

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetFilter() yfilter.YFilter { return interfaceBrief.YFilter }

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) SetFilter(yf yfilter.YFilter) { interfaceBrief.YFilter = yf }

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetGoName(yname string) string {
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "state" { return "State" }
    if yname == "mtu" { return "Mtu" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix-length-v6" { return "PrefixLengthV6" }
    if yname == "primary-address" { return "PrimaryAddress" }
    if yname == "primary-address-v6" { return "PrimaryAddressV6" }
    return ""
}

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetSegmentPath() string {
    return "interface-brief"
}

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name-xr"] = interfaceBrief.InterfaceNameXr
    leafs["state"] = interfaceBrief.State
    leafs["mtu"] = interfaceBrief.Mtu
    leafs["prefix-length"] = interfaceBrief.PrefixLength
    leafs["prefix-length-v6"] = interfaceBrief.PrefixLengthV6
    leafs["primary-address"] = interfaceBrief.PrimaryAddress
    leafs["primary-address-v6"] = interfaceBrief.PrimaryAddressV6
    return leafs
}

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetYangName() string { return "interface-brief" }

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) SetParent(parent types.Entity) { interfaceBrief.parent = parent }

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetParent() types.Entity { return interfaceBrief.parent }

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetParentYangName() string { return "detail" }

// MplsOam_Interface_Details_Detail_PacketStatistics
// Packet statistics
type MplsOam_Interface_Details_Detail_PacketStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet reception counts.
    Received MplsOam_Interface_Details_Detail_PacketStatistics_Received

    // Packet transmit counts.
    Sent MplsOam_Interface_Details_Detail_PacketStatistics_Sent

    // Working Request Packet transmit counts.
    WorkingReqSent MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent

    // Working Reply Packet transmit counts.
    WorkingRepSent MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent

    // Protect Request Packet transmit counts.
    ProtectReqSent MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent

    // Protect Reply Packet transmit counts.
    ProtectRepSent MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent
}

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetFilter() yfilter.YFilter { return packetStatistics.YFilter }

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) SetFilter(yf yfilter.YFilter) { packetStatistics.YFilter = yf }

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetGoName(yname string) string {
    if yname == "received" { return "Received" }
    if yname == "sent" { return "Sent" }
    if yname == "working-req-sent" { return "WorkingReqSent" }
    if yname == "working-rep-sent" { return "WorkingRepSent" }
    if yname == "protect-req-sent" { return "ProtectReqSent" }
    if yname == "protect-rep-sent" { return "ProtectRepSent" }
    return ""
}

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetSegmentPath() string {
    return "packet-statistics"
}

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "received" {
        return &packetStatistics.Received
    }
    if childYangName == "sent" {
        return &packetStatistics.Sent
    }
    if childYangName == "working-req-sent" {
        return &packetStatistics.WorkingReqSent
    }
    if childYangName == "working-rep-sent" {
        return &packetStatistics.WorkingRepSent
    }
    if childYangName == "protect-req-sent" {
        return &packetStatistics.ProtectReqSent
    }
    if childYangName == "protect-rep-sent" {
        return &packetStatistics.ProtectRepSent
    }
    return nil
}

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["received"] = &packetStatistics.Received
    children["sent"] = &packetStatistics.Sent
    children["working-req-sent"] = &packetStatistics.WorkingReqSent
    children["working-rep-sent"] = &packetStatistics.WorkingRepSent
    children["protect-req-sent"] = &packetStatistics.ProtectReqSent
    children["protect-rep-sent"] = &packetStatistics.ProtectRepSent
    return children
}

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetYangName() string { return "packet-statistics" }

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) SetParent(parent types.Entity) { packetStatistics.parent = parent }

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetParent() types.Entity { return packetStatistics.parent }

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetParentYangName() string { return "detail" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received
// Packet reception counts
type MplsOam_Interface_Details_Detail_PacketStatistics_Received struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received good request.
    ReceivedGoodRequest MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest

    // Received good reply.
    ReceivedGoodReply MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply

    // Received unknown packets.
    ReceivedUnknown MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown

    // IP header error.
    ReceivedErrorIpHeader MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader

    // UDP header error.
    ReceivedErrorUdpHeader MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader

    // RUNT error.
    ReceivedErrorRunt MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt

    // Dropped queue full.
    ReceivedErrorQueueFull MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull

    // General error.
    ReceivedErrorGeneral MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral

    // Error no Interfaces.
    ReceivedErrorNoInterface MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface

    // Error no memory.
    ReceivedErrorNoMemory MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory

    // Protect Protocol Received good request.
    ProtectProtocolReceivedGoodRequest MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest

    // Protect Protocol Received good reply.
    ProtectProtocolReceivedGoodReply MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply

    // Received Reqeust with BFD TLV.
    ReceivedGoodBfdRequest MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest

    // Received Reply with BFD TLV.
    ReceivedGoodBfdReply MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply
}

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetFilter() yfilter.YFilter { return received.YFilter }

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) SetFilter(yf yfilter.YFilter) { received.YFilter = yf }

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetGoName(yname string) string {
    if yname == "received-good-request" { return "ReceivedGoodRequest" }
    if yname == "received-good-reply" { return "ReceivedGoodReply" }
    if yname == "received-unknown" { return "ReceivedUnknown" }
    if yname == "received-error-ip-header" { return "ReceivedErrorIpHeader" }
    if yname == "received-error-udp-header" { return "ReceivedErrorUdpHeader" }
    if yname == "received-error-runt" { return "ReceivedErrorRunt" }
    if yname == "received-error-queue-full" { return "ReceivedErrorQueueFull" }
    if yname == "received-error-general" { return "ReceivedErrorGeneral" }
    if yname == "received-error-no-interface" { return "ReceivedErrorNoInterface" }
    if yname == "received-error-no-memory" { return "ReceivedErrorNoMemory" }
    if yname == "protect-protocol-received-good-request" { return "ProtectProtocolReceivedGoodRequest" }
    if yname == "protect-protocol-received-good-reply" { return "ProtectProtocolReceivedGoodReply" }
    if yname == "received-good-bfd-request" { return "ReceivedGoodBfdRequest" }
    if yname == "received-good-bfd-reply" { return "ReceivedGoodBfdReply" }
    return ""
}

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetSegmentPath() string {
    return "received"
}

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "received-good-request" {
        return &received.ReceivedGoodRequest
    }
    if childYangName == "received-good-reply" {
        return &received.ReceivedGoodReply
    }
    if childYangName == "received-unknown" {
        return &received.ReceivedUnknown
    }
    if childYangName == "received-error-ip-header" {
        return &received.ReceivedErrorIpHeader
    }
    if childYangName == "received-error-udp-header" {
        return &received.ReceivedErrorUdpHeader
    }
    if childYangName == "received-error-runt" {
        return &received.ReceivedErrorRunt
    }
    if childYangName == "received-error-queue-full" {
        return &received.ReceivedErrorQueueFull
    }
    if childYangName == "received-error-general" {
        return &received.ReceivedErrorGeneral
    }
    if childYangName == "received-error-no-interface" {
        return &received.ReceivedErrorNoInterface
    }
    if childYangName == "received-error-no-memory" {
        return &received.ReceivedErrorNoMemory
    }
    if childYangName == "protect-protocol-received-good-request" {
        return &received.ProtectProtocolReceivedGoodRequest
    }
    if childYangName == "protect-protocol-received-good-reply" {
        return &received.ProtectProtocolReceivedGoodReply
    }
    if childYangName == "received-good-bfd-request" {
        return &received.ReceivedGoodBfdRequest
    }
    if childYangName == "received-good-bfd-reply" {
        return &received.ReceivedGoodBfdReply
    }
    return nil
}

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["received-good-request"] = &received.ReceivedGoodRequest
    children["received-good-reply"] = &received.ReceivedGoodReply
    children["received-unknown"] = &received.ReceivedUnknown
    children["received-error-ip-header"] = &received.ReceivedErrorIpHeader
    children["received-error-udp-header"] = &received.ReceivedErrorUdpHeader
    children["received-error-runt"] = &received.ReceivedErrorRunt
    children["received-error-queue-full"] = &received.ReceivedErrorQueueFull
    children["received-error-general"] = &received.ReceivedErrorGeneral
    children["received-error-no-interface"] = &received.ReceivedErrorNoInterface
    children["received-error-no-memory"] = &received.ReceivedErrorNoMemory
    children["protect-protocol-received-good-request"] = &received.ProtectProtocolReceivedGoodRequest
    children["protect-protocol-received-good-reply"] = &received.ProtectProtocolReceivedGoodReply
    children["received-good-bfd-request"] = &received.ReceivedGoodBfdRequest
    children["received-good-bfd-reply"] = &received.ReceivedGoodBfdReply
    return children
}

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetBundleName() string { return "cisco_ios_xr" }

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetYangName() string { return "received" }

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) SetParent(parent types.Entity) { received.parent = parent }

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetParent() types.Entity { return received.parent }

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetParentYangName() string { return "packet-statistics" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest
// Received good request
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetFilter() yfilter.YFilter { return receivedGoodRequest.YFilter }

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) SetFilter(yf yfilter.YFilter) { receivedGoodRequest.YFilter = yf }

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetSegmentPath() string {
    return "received-good-request"
}

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedGoodRequest.Packets
    leafs["bytes"] = receivedGoodRequest.Bytes
    return leafs
}

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetBundleName() string { return "cisco_ios_xr" }

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetYangName() string { return "received-good-request" }

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) SetParent(parent types.Entity) { receivedGoodRequest.parent = parent }

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetParent() types.Entity { return receivedGoodRequest.parent }

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply
// Received good reply
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetFilter() yfilter.YFilter { return receivedGoodReply.YFilter }

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) SetFilter(yf yfilter.YFilter) { receivedGoodReply.YFilter = yf }

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetSegmentPath() string {
    return "received-good-reply"
}

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedGoodReply.Packets
    leafs["bytes"] = receivedGoodReply.Bytes
    return leafs
}

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetBundleName() string { return "cisco_ios_xr" }

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetYangName() string { return "received-good-reply" }

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) SetParent(parent types.Entity) { receivedGoodReply.parent = parent }

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetParent() types.Entity { return receivedGoodReply.parent }

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown
// Received unknown packets
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetFilter() yfilter.YFilter { return receivedUnknown.YFilter }

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) SetFilter(yf yfilter.YFilter) { receivedUnknown.YFilter = yf }

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetSegmentPath() string {
    return "received-unknown"
}

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedUnknown.Packets
    leafs["bytes"] = receivedUnknown.Bytes
    return leafs
}

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetBundleName() string { return "cisco_ios_xr" }

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetYangName() string { return "received-unknown" }

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) SetParent(parent types.Entity) { receivedUnknown.parent = parent }

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetParent() types.Entity { return receivedUnknown.parent }

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader
// IP header error
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetFilter() yfilter.YFilter { return receivedErrorIpHeader.YFilter }

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) SetFilter(yf yfilter.YFilter) { receivedErrorIpHeader.YFilter = yf }

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetSegmentPath() string {
    return "received-error-ip-header"
}

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorIpHeader.Packets
    leafs["bytes"] = receivedErrorIpHeader.Bytes
    return leafs
}

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetYangName() string { return "received-error-ip-header" }

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) SetParent(parent types.Entity) { receivedErrorIpHeader.parent = parent }

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetParent() types.Entity { return receivedErrorIpHeader.parent }

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader
// UDP header error
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetFilter() yfilter.YFilter { return receivedErrorUdpHeader.YFilter }

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) SetFilter(yf yfilter.YFilter) { receivedErrorUdpHeader.YFilter = yf }

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetSegmentPath() string {
    return "received-error-udp-header"
}

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorUdpHeader.Packets
    leafs["bytes"] = receivedErrorUdpHeader.Bytes
    return leafs
}

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetYangName() string { return "received-error-udp-header" }

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) SetParent(parent types.Entity) { receivedErrorUdpHeader.parent = parent }

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetParent() types.Entity { return receivedErrorUdpHeader.parent }

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt
// RUNT error
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetFilter() yfilter.YFilter { return receivedErrorRunt.YFilter }

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) SetFilter(yf yfilter.YFilter) { receivedErrorRunt.YFilter = yf }

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetSegmentPath() string {
    return "received-error-runt"
}

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorRunt.Packets
    leafs["bytes"] = receivedErrorRunt.Bytes
    return leafs
}

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetYangName() string { return "received-error-runt" }

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) SetParent(parent types.Entity) { receivedErrorRunt.parent = parent }

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetParent() types.Entity { return receivedErrorRunt.parent }

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull
// Dropped queue full
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetFilter() yfilter.YFilter { return receivedErrorQueueFull.YFilter }

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) SetFilter(yf yfilter.YFilter) { receivedErrorQueueFull.YFilter = yf }

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetSegmentPath() string {
    return "received-error-queue-full"
}

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorQueueFull.Packets
    leafs["bytes"] = receivedErrorQueueFull.Bytes
    return leafs
}

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetYangName() string { return "received-error-queue-full" }

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) SetParent(parent types.Entity) { receivedErrorQueueFull.parent = parent }

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetParent() types.Entity { return receivedErrorQueueFull.parent }

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral
// General error
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetFilter() yfilter.YFilter { return receivedErrorGeneral.YFilter }

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) SetFilter(yf yfilter.YFilter) { receivedErrorGeneral.YFilter = yf }

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetSegmentPath() string {
    return "received-error-general"
}

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorGeneral.Packets
    leafs["bytes"] = receivedErrorGeneral.Bytes
    return leafs
}

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetYangName() string { return "received-error-general" }

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) SetParent(parent types.Entity) { receivedErrorGeneral.parent = parent }

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetParent() types.Entity { return receivedErrorGeneral.parent }

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface
// Error no Interfaces
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetFilter() yfilter.YFilter { return receivedErrorNoInterface.YFilter }

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) SetFilter(yf yfilter.YFilter) { receivedErrorNoInterface.YFilter = yf }

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetSegmentPath() string {
    return "received-error-no-interface"
}

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorNoInterface.Packets
    leafs["bytes"] = receivedErrorNoInterface.Bytes
    return leafs
}

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetYangName() string { return "received-error-no-interface" }

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) SetParent(parent types.Entity) { receivedErrorNoInterface.parent = parent }

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetParent() types.Entity { return receivedErrorNoInterface.parent }

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory
// Error no memory
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetFilter() yfilter.YFilter { return receivedErrorNoMemory.YFilter }

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) SetFilter(yf yfilter.YFilter) { receivedErrorNoMemory.YFilter = yf }

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetSegmentPath() string {
    return "received-error-no-memory"
}

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorNoMemory.Packets
    leafs["bytes"] = receivedErrorNoMemory.Bytes
    return leafs
}

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetYangName() string { return "received-error-no-memory" }

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) SetParent(parent types.Entity) { receivedErrorNoMemory.parent = parent }

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetParent() types.Entity { return receivedErrorNoMemory.parent }

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest
// Protect Protocol Received good request
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetFilter() yfilter.YFilter { return protectProtocolReceivedGoodRequest.YFilter }

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) SetFilter(yf yfilter.YFilter) { protectProtocolReceivedGoodRequest.YFilter = yf }

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetSegmentPath() string {
    return "protect-protocol-received-good-request"
}

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = protectProtocolReceivedGoodRequest.Packets
    leafs["bytes"] = protectProtocolReceivedGoodRequest.Bytes
    return leafs
}

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetBundleName() string { return "cisco_ios_xr" }

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetYangName() string { return "protect-protocol-received-good-request" }

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) SetParent(parent types.Entity) { protectProtocolReceivedGoodRequest.parent = parent }

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetParent() types.Entity { return protectProtocolReceivedGoodRequest.parent }

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply
// Protect Protocol Received good reply
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetFilter() yfilter.YFilter { return protectProtocolReceivedGoodReply.YFilter }

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) SetFilter(yf yfilter.YFilter) { protectProtocolReceivedGoodReply.YFilter = yf }

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetSegmentPath() string {
    return "protect-protocol-received-good-reply"
}

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = protectProtocolReceivedGoodReply.Packets
    leafs["bytes"] = protectProtocolReceivedGoodReply.Bytes
    return leafs
}

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetBundleName() string { return "cisco_ios_xr" }

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetYangName() string { return "protect-protocol-received-good-reply" }

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) SetParent(parent types.Entity) { protectProtocolReceivedGoodReply.parent = parent }

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetParent() types.Entity { return protectProtocolReceivedGoodReply.parent }

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest
// Received Reqeust with BFD TLV
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetFilter() yfilter.YFilter { return receivedGoodBfdRequest.YFilter }

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) SetFilter(yf yfilter.YFilter) { receivedGoodBfdRequest.YFilter = yf }

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetSegmentPath() string {
    return "received-good-bfd-request"
}

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedGoodBfdRequest.Packets
    leafs["bytes"] = receivedGoodBfdRequest.Bytes
    return leafs
}

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetBundleName() string { return "cisco_ios_xr" }

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetYangName() string { return "received-good-bfd-request" }

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) SetParent(parent types.Entity) { receivedGoodBfdRequest.parent = parent }

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetParent() types.Entity { return receivedGoodBfdRequest.parent }

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply
// Received Reply with BFD TLV
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetFilter() yfilter.YFilter { return receivedGoodBfdReply.YFilter }

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) SetFilter(yf yfilter.YFilter) { receivedGoodBfdReply.YFilter = yf }

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetSegmentPath() string {
    return "received-good-bfd-reply"
}

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedGoodBfdReply.Packets
    leafs["bytes"] = receivedGoodBfdReply.Bytes
    return leafs
}

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetBundleName() string { return "cisco_ios_xr" }

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetYangName() string { return "received-good-bfd-reply" }

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) SetParent(parent types.Entity) { receivedGoodBfdReply.parent = parent }

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetParent() types.Entity { return receivedGoodBfdReply.parent }

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetParentYangName() string { return "received" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Sent
// Packet transmit counts
type MplsOam_Interface_Details_Detail_PacketStatistics_Sent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply
}

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetFilter() yfilter.YFilter { return sent.YFilter }

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) SetFilter(yf yfilter.YFilter) { sent.YFilter = yf }

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetGoName(yname string) string {
    if yname == "transmit-good" { return "TransmitGood" }
    if yname == "transmit-drop" { return "TransmitDrop" }
    if yname == "transmit-bfd-good" { return "TransmitBfdGood" }
    if yname == "bfd-no-reply" { return "BfdNoReply" }
    return ""
}

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetSegmentPath() string {
    return "sent"
}

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit-good" {
        return &sent.TransmitGood
    }
    if childYangName == "transmit-drop" {
        return &sent.TransmitDrop
    }
    if childYangName == "transmit-bfd-good" {
        return &sent.TransmitBfdGood
    }
    if childYangName == "bfd-no-reply" {
        return &sent.BfdNoReply
    }
    return nil
}

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit-good"] = &sent.TransmitGood
    children["transmit-drop"] = &sent.TransmitDrop
    children["transmit-bfd-good"] = &sent.TransmitBfdGood
    children["bfd-no-reply"] = &sent.BfdNoReply
    return children
}

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetBundleName() string { return "cisco_ios_xr" }

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetYangName() string { return "sent" }

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) SetParent(parent types.Entity) { sent.parent = parent }

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetParent() types.Entity { return sent.parent }

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetParentYangName() string { return "packet-statistics" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood
// Transmit good packets
type MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetFilter() yfilter.YFilter { return transmitGood.YFilter }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) SetFilter(yf yfilter.YFilter) { transmitGood.YFilter = yf }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetSegmentPath() string {
    return "transmit-good"
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitGood.Packets
    leafs["bytes"] = transmitGood.Bytes
    return leafs
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetYangName() string { return "transmit-good" }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) SetParent(parent types.Entity) { transmitGood.parent = parent }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetParent() types.Entity { return transmitGood.parent }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetParentYangName() string { return "sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop
// Transmit drop packets
type MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetFilter() yfilter.YFilter { return transmitDrop.YFilter }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) SetFilter(yf yfilter.YFilter) { transmitDrop.YFilter = yf }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetSegmentPath() string {
    return "transmit-drop"
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitDrop.Packets
    leafs["bytes"] = transmitDrop.Bytes
    return leafs
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetBundleName() string { return "cisco_ios_xr" }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetYangName() string { return "transmit-drop" }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) SetParent(parent types.Entity) { transmitDrop.parent = parent }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetParent() types.Entity { return transmitDrop.parent }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetParentYangName() string { return "sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetFilter() yfilter.YFilter { return transmitBfdGood.YFilter }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) SetFilter(yf yfilter.YFilter) { transmitBfdGood.YFilter = yf }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetSegmentPath() string {
    return "transmit-bfd-good"
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitBfdGood.Packets
    leafs["bytes"] = transmitBfdGood.Bytes
    return leafs
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetYangName() string { return "transmit-bfd-good" }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) SetParent(parent types.Entity) { transmitBfdGood.parent = parent }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetParent() types.Entity { return transmitBfdGood.parent }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetParentYangName() string { return "sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetFilter() yfilter.YFilter { return bfdNoReply.YFilter }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) SetFilter(yf yfilter.YFilter) { bfdNoReply.YFilter = yf }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetSegmentPath() string {
    return "bfd-no-reply"
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = bfdNoReply.Packets
    leafs["bytes"] = bfdNoReply.Bytes
    return leafs
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetBundleName() string { return "cisco_ios_xr" }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetYangName() string { return "bfd-no-reply" }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) SetParent(parent types.Entity) { bfdNoReply.parent = parent }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetParent() types.Entity { return bfdNoReply.parent }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetParentYangName() string { return "sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent
// Working Request Packet transmit counts
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply
}

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetFilter() yfilter.YFilter { return workingReqSent.YFilter }

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) SetFilter(yf yfilter.YFilter) { workingReqSent.YFilter = yf }

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetGoName(yname string) string {
    if yname == "transmit-good" { return "TransmitGood" }
    if yname == "transmit-drop" { return "TransmitDrop" }
    if yname == "transmit-bfd-good" { return "TransmitBfdGood" }
    if yname == "bfd-no-reply" { return "BfdNoReply" }
    return ""
}

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetSegmentPath() string {
    return "working-req-sent"
}

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit-good" {
        return &workingReqSent.TransmitGood
    }
    if childYangName == "transmit-drop" {
        return &workingReqSent.TransmitDrop
    }
    if childYangName == "transmit-bfd-good" {
        return &workingReqSent.TransmitBfdGood
    }
    if childYangName == "bfd-no-reply" {
        return &workingReqSent.BfdNoReply
    }
    return nil
}

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit-good"] = &workingReqSent.TransmitGood
    children["transmit-drop"] = &workingReqSent.TransmitDrop
    children["transmit-bfd-good"] = &workingReqSent.TransmitBfdGood
    children["bfd-no-reply"] = &workingReqSent.BfdNoReply
    return children
}

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetBundleName() string { return "cisco_ios_xr" }

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetYangName() string { return "working-req-sent" }

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) SetParent(parent types.Entity) { workingReqSent.parent = parent }

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetParent() types.Entity { return workingReqSent.parent }

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetParentYangName() string { return "packet-statistics" }

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood
// Transmit good packets
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetFilter() yfilter.YFilter { return transmitGood.YFilter }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) SetFilter(yf yfilter.YFilter) { transmitGood.YFilter = yf }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetSegmentPath() string {
    return "transmit-good"
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitGood.Packets
    leafs["bytes"] = transmitGood.Bytes
    return leafs
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetYangName() string { return "transmit-good" }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) SetParent(parent types.Entity) { transmitGood.parent = parent }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetParent() types.Entity { return transmitGood.parent }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetParentYangName() string { return "working-req-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop
// Transmit drop packets
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetFilter() yfilter.YFilter { return transmitDrop.YFilter }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) SetFilter(yf yfilter.YFilter) { transmitDrop.YFilter = yf }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetSegmentPath() string {
    return "transmit-drop"
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitDrop.Packets
    leafs["bytes"] = transmitDrop.Bytes
    return leafs
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetBundleName() string { return "cisco_ios_xr" }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetYangName() string { return "transmit-drop" }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) SetParent(parent types.Entity) { transmitDrop.parent = parent }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetParent() types.Entity { return transmitDrop.parent }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetParentYangName() string { return "working-req-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetFilter() yfilter.YFilter { return transmitBfdGood.YFilter }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) SetFilter(yf yfilter.YFilter) { transmitBfdGood.YFilter = yf }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetSegmentPath() string {
    return "transmit-bfd-good"
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitBfdGood.Packets
    leafs["bytes"] = transmitBfdGood.Bytes
    return leafs
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetYangName() string { return "transmit-bfd-good" }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) SetParent(parent types.Entity) { transmitBfdGood.parent = parent }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetParent() types.Entity { return transmitBfdGood.parent }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetParentYangName() string { return "working-req-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetFilter() yfilter.YFilter { return bfdNoReply.YFilter }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) SetFilter(yf yfilter.YFilter) { bfdNoReply.YFilter = yf }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetSegmentPath() string {
    return "bfd-no-reply"
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = bfdNoReply.Packets
    leafs["bytes"] = bfdNoReply.Bytes
    return leafs
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetBundleName() string { return "cisco_ios_xr" }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetYangName() string { return "bfd-no-reply" }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) SetParent(parent types.Entity) { bfdNoReply.parent = parent }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetParent() types.Entity { return bfdNoReply.parent }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetParentYangName() string { return "working-req-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent
// Working Reply Packet transmit counts
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply
}

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetFilter() yfilter.YFilter { return workingRepSent.YFilter }

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) SetFilter(yf yfilter.YFilter) { workingRepSent.YFilter = yf }

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetGoName(yname string) string {
    if yname == "transmit-good" { return "TransmitGood" }
    if yname == "transmit-drop" { return "TransmitDrop" }
    if yname == "transmit-bfd-good" { return "TransmitBfdGood" }
    if yname == "bfd-no-reply" { return "BfdNoReply" }
    return ""
}

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetSegmentPath() string {
    return "working-rep-sent"
}

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit-good" {
        return &workingRepSent.TransmitGood
    }
    if childYangName == "transmit-drop" {
        return &workingRepSent.TransmitDrop
    }
    if childYangName == "transmit-bfd-good" {
        return &workingRepSent.TransmitBfdGood
    }
    if childYangName == "bfd-no-reply" {
        return &workingRepSent.BfdNoReply
    }
    return nil
}

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit-good"] = &workingRepSent.TransmitGood
    children["transmit-drop"] = &workingRepSent.TransmitDrop
    children["transmit-bfd-good"] = &workingRepSent.TransmitBfdGood
    children["bfd-no-reply"] = &workingRepSent.BfdNoReply
    return children
}

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetBundleName() string { return "cisco_ios_xr" }

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetYangName() string { return "working-rep-sent" }

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) SetParent(parent types.Entity) { workingRepSent.parent = parent }

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetParent() types.Entity { return workingRepSent.parent }

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetParentYangName() string { return "packet-statistics" }

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood
// Transmit good packets
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetFilter() yfilter.YFilter { return transmitGood.YFilter }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) SetFilter(yf yfilter.YFilter) { transmitGood.YFilter = yf }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetSegmentPath() string {
    return "transmit-good"
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitGood.Packets
    leafs["bytes"] = transmitGood.Bytes
    return leafs
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetYangName() string { return "transmit-good" }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) SetParent(parent types.Entity) { transmitGood.parent = parent }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetParent() types.Entity { return transmitGood.parent }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetParentYangName() string { return "working-rep-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop
// Transmit drop packets
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetFilter() yfilter.YFilter { return transmitDrop.YFilter }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) SetFilter(yf yfilter.YFilter) { transmitDrop.YFilter = yf }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetSegmentPath() string {
    return "transmit-drop"
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitDrop.Packets
    leafs["bytes"] = transmitDrop.Bytes
    return leafs
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetBundleName() string { return "cisco_ios_xr" }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetYangName() string { return "transmit-drop" }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) SetParent(parent types.Entity) { transmitDrop.parent = parent }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetParent() types.Entity { return transmitDrop.parent }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetParentYangName() string { return "working-rep-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetFilter() yfilter.YFilter { return transmitBfdGood.YFilter }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) SetFilter(yf yfilter.YFilter) { transmitBfdGood.YFilter = yf }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetSegmentPath() string {
    return "transmit-bfd-good"
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitBfdGood.Packets
    leafs["bytes"] = transmitBfdGood.Bytes
    return leafs
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetYangName() string { return "transmit-bfd-good" }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) SetParent(parent types.Entity) { transmitBfdGood.parent = parent }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetParent() types.Entity { return transmitBfdGood.parent }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetParentYangName() string { return "working-rep-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetFilter() yfilter.YFilter { return bfdNoReply.YFilter }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) SetFilter(yf yfilter.YFilter) { bfdNoReply.YFilter = yf }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetSegmentPath() string {
    return "bfd-no-reply"
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = bfdNoReply.Packets
    leafs["bytes"] = bfdNoReply.Bytes
    return leafs
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetBundleName() string { return "cisco_ios_xr" }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetYangName() string { return "bfd-no-reply" }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) SetParent(parent types.Entity) { bfdNoReply.parent = parent }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetParent() types.Entity { return bfdNoReply.parent }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetParentYangName() string { return "working-rep-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent
// Protect Request Packet transmit counts
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply
}

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetFilter() yfilter.YFilter { return protectReqSent.YFilter }

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) SetFilter(yf yfilter.YFilter) { protectReqSent.YFilter = yf }

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetGoName(yname string) string {
    if yname == "transmit-good" { return "TransmitGood" }
    if yname == "transmit-drop" { return "TransmitDrop" }
    if yname == "transmit-bfd-good" { return "TransmitBfdGood" }
    if yname == "bfd-no-reply" { return "BfdNoReply" }
    return ""
}

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetSegmentPath() string {
    return "protect-req-sent"
}

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit-good" {
        return &protectReqSent.TransmitGood
    }
    if childYangName == "transmit-drop" {
        return &protectReqSent.TransmitDrop
    }
    if childYangName == "transmit-bfd-good" {
        return &protectReqSent.TransmitBfdGood
    }
    if childYangName == "bfd-no-reply" {
        return &protectReqSent.BfdNoReply
    }
    return nil
}

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit-good"] = &protectReqSent.TransmitGood
    children["transmit-drop"] = &protectReqSent.TransmitDrop
    children["transmit-bfd-good"] = &protectReqSent.TransmitBfdGood
    children["bfd-no-reply"] = &protectReqSent.BfdNoReply
    return children
}

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetBundleName() string { return "cisco_ios_xr" }

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetYangName() string { return "protect-req-sent" }

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) SetParent(parent types.Entity) { protectReqSent.parent = parent }

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetParent() types.Entity { return protectReqSent.parent }

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetParentYangName() string { return "packet-statistics" }

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood
// Transmit good packets
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetFilter() yfilter.YFilter { return transmitGood.YFilter }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) SetFilter(yf yfilter.YFilter) { transmitGood.YFilter = yf }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetSegmentPath() string {
    return "transmit-good"
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitGood.Packets
    leafs["bytes"] = transmitGood.Bytes
    return leafs
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetYangName() string { return "transmit-good" }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) SetParent(parent types.Entity) { transmitGood.parent = parent }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetParent() types.Entity { return transmitGood.parent }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetParentYangName() string { return "protect-req-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop
// Transmit drop packets
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetFilter() yfilter.YFilter { return transmitDrop.YFilter }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) SetFilter(yf yfilter.YFilter) { transmitDrop.YFilter = yf }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetSegmentPath() string {
    return "transmit-drop"
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitDrop.Packets
    leafs["bytes"] = transmitDrop.Bytes
    return leafs
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetBundleName() string { return "cisco_ios_xr" }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetYangName() string { return "transmit-drop" }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) SetParent(parent types.Entity) { transmitDrop.parent = parent }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetParent() types.Entity { return transmitDrop.parent }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetParentYangName() string { return "protect-req-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetFilter() yfilter.YFilter { return transmitBfdGood.YFilter }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) SetFilter(yf yfilter.YFilter) { transmitBfdGood.YFilter = yf }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetSegmentPath() string {
    return "transmit-bfd-good"
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitBfdGood.Packets
    leafs["bytes"] = transmitBfdGood.Bytes
    return leafs
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetYangName() string { return "transmit-bfd-good" }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) SetParent(parent types.Entity) { transmitBfdGood.parent = parent }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetParent() types.Entity { return transmitBfdGood.parent }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetParentYangName() string { return "protect-req-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetFilter() yfilter.YFilter { return bfdNoReply.YFilter }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) SetFilter(yf yfilter.YFilter) { bfdNoReply.YFilter = yf }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetSegmentPath() string {
    return "bfd-no-reply"
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = bfdNoReply.Packets
    leafs["bytes"] = bfdNoReply.Bytes
    return leafs
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetBundleName() string { return "cisco_ios_xr" }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetYangName() string { return "bfd-no-reply" }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) SetParent(parent types.Entity) { bfdNoReply.parent = parent }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetParent() types.Entity { return bfdNoReply.parent }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetParentYangName() string { return "protect-req-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent
// Protect Reply Packet transmit counts
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply
}

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetFilter() yfilter.YFilter { return protectRepSent.YFilter }

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) SetFilter(yf yfilter.YFilter) { protectRepSent.YFilter = yf }

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetGoName(yname string) string {
    if yname == "transmit-good" { return "TransmitGood" }
    if yname == "transmit-drop" { return "TransmitDrop" }
    if yname == "transmit-bfd-good" { return "TransmitBfdGood" }
    if yname == "bfd-no-reply" { return "BfdNoReply" }
    return ""
}

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetSegmentPath() string {
    return "protect-rep-sent"
}

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit-good" {
        return &protectRepSent.TransmitGood
    }
    if childYangName == "transmit-drop" {
        return &protectRepSent.TransmitDrop
    }
    if childYangName == "transmit-bfd-good" {
        return &protectRepSent.TransmitBfdGood
    }
    if childYangName == "bfd-no-reply" {
        return &protectRepSent.BfdNoReply
    }
    return nil
}

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit-good"] = &protectRepSent.TransmitGood
    children["transmit-drop"] = &protectRepSent.TransmitDrop
    children["transmit-bfd-good"] = &protectRepSent.TransmitBfdGood
    children["bfd-no-reply"] = &protectRepSent.BfdNoReply
    return children
}

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetBundleName() string { return "cisco_ios_xr" }

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetYangName() string { return "protect-rep-sent" }

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) SetParent(parent types.Entity) { protectRepSent.parent = parent }

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetParent() types.Entity { return protectRepSent.parent }

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetParentYangName() string { return "packet-statistics" }

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood
// Transmit good packets
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetFilter() yfilter.YFilter { return transmitGood.YFilter }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) SetFilter(yf yfilter.YFilter) { transmitGood.YFilter = yf }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetSegmentPath() string {
    return "transmit-good"
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitGood.Packets
    leafs["bytes"] = transmitGood.Bytes
    return leafs
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetYangName() string { return "transmit-good" }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) SetParent(parent types.Entity) { transmitGood.parent = parent }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetParent() types.Entity { return transmitGood.parent }

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetParentYangName() string { return "protect-rep-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop
// Transmit drop packets
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetFilter() yfilter.YFilter { return transmitDrop.YFilter }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) SetFilter(yf yfilter.YFilter) { transmitDrop.YFilter = yf }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetSegmentPath() string {
    return "transmit-drop"
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitDrop.Packets
    leafs["bytes"] = transmitDrop.Bytes
    return leafs
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetBundleName() string { return "cisco_ios_xr" }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetYangName() string { return "transmit-drop" }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) SetParent(parent types.Entity) { transmitDrop.parent = parent }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetParent() types.Entity { return transmitDrop.parent }

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetParentYangName() string { return "protect-rep-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetFilter() yfilter.YFilter { return transmitBfdGood.YFilter }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) SetFilter(yf yfilter.YFilter) { transmitBfdGood.YFilter = yf }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetSegmentPath() string {
    return "transmit-bfd-good"
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitBfdGood.Packets
    leafs["bytes"] = transmitBfdGood.Bytes
    return leafs
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetYangName() string { return "transmit-bfd-good" }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) SetParent(parent types.Entity) { transmitBfdGood.parent = parent }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetParent() types.Entity { return transmitBfdGood.parent }

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetParentYangName() string { return "protect-rep-sent" }

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetFilter() yfilter.YFilter { return bfdNoReply.YFilter }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) SetFilter(yf yfilter.YFilter) { bfdNoReply.YFilter = yf }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetSegmentPath() string {
    return "bfd-no-reply"
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = bfdNoReply.Packets
    leafs["bytes"] = bfdNoReply.Bytes
    return leafs
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetBundleName() string { return "cisco_ios_xr" }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetYangName() string { return "bfd-no-reply" }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) SetParent(parent types.Entity) { bfdNoReply.parent = parent }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetParent() types.Entity { return bfdNoReply.parent }

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetParentYangName() string { return "protect-rep-sent" }

// MplsOam_Packet
// LSPV packet counters operational data
type MplsOam_Packet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet reception counts.
    Received MplsOam_Packet_Received

    // Packet transmit counts.
    Sent MplsOam_Packet_Sent

    // Working Request Packet transmit counts.
    WorkingReqSent MplsOam_Packet_WorkingReqSent

    // Working Reply Packet transmit counts.
    WorkingRepSent MplsOam_Packet_WorkingRepSent

    // Protect Request Packet transmit counts.
    ProtectReqSent MplsOam_Packet_ProtectReqSent

    // Protect Reply Packet transmit counts.
    ProtectRepSent MplsOam_Packet_ProtectRepSent
}

func (packet *MplsOam_Packet) GetFilter() yfilter.YFilter { return packet.YFilter }

func (packet *MplsOam_Packet) SetFilter(yf yfilter.YFilter) { packet.YFilter = yf }

func (packet *MplsOam_Packet) GetGoName(yname string) string {
    if yname == "received" { return "Received" }
    if yname == "sent" { return "Sent" }
    if yname == "working-req-sent" { return "WorkingReqSent" }
    if yname == "working-rep-sent" { return "WorkingRepSent" }
    if yname == "protect-req-sent" { return "ProtectReqSent" }
    if yname == "protect-rep-sent" { return "ProtectRepSent" }
    return ""
}

func (packet *MplsOam_Packet) GetSegmentPath() string {
    return "packet"
}

func (packet *MplsOam_Packet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "received" {
        return &packet.Received
    }
    if childYangName == "sent" {
        return &packet.Sent
    }
    if childYangName == "working-req-sent" {
        return &packet.WorkingReqSent
    }
    if childYangName == "working-rep-sent" {
        return &packet.WorkingRepSent
    }
    if childYangName == "protect-req-sent" {
        return &packet.ProtectReqSent
    }
    if childYangName == "protect-rep-sent" {
        return &packet.ProtectRepSent
    }
    return nil
}

func (packet *MplsOam_Packet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["received"] = &packet.Received
    children["sent"] = &packet.Sent
    children["working-req-sent"] = &packet.WorkingReqSent
    children["working-rep-sent"] = &packet.WorkingRepSent
    children["protect-req-sent"] = &packet.ProtectReqSent
    children["protect-rep-sent"] = &packet.ProtectRepSent
    return children
}

func (packet *MplsOam_Packet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (packet *MplsOam_Packet) GetBundleName() string { return "cisco_ios_xr" }

func (packet *MplsOam_Packet) GetYangName() string { return "packet" }

func (packet *MplsOam_Packet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packet *MplsOam_Packet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packet *MplsOam_Packet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packet *MplsOam_Packet) SetParent(parent types.Entity) { packet.parent = parent }

func (packet *MplsOam_Packet) GetParent() types.Entity { return packet.parent }

func (packet *MplsOam_Packet) GetParentYangName() string { return "mpls-oam" }

// MplsOam_Packet_Received
// Packet reception counts
type MplsOam_Packet_Received struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Received good request.
    ReceivedGoodRequest MplsOam_Packet_Received_ReceivedGoodRequest

    // Received good reply.
    ReceivedGoodReply MplsOam_Packet_Received_ReceivedGoodReply

    // Received unknown packets.
    ReceivedUnknown MplsOam_Packet_Received_ReceivedUnknown

    // IP header error.
    ReceivedErrorIpHeader MplsOam_Packet_Received_ReceivedErrorIpHeader

    // UDP header error.
    ReceivedErrorUdpHeader MplsOam_Packet_Received_ReceivedErrorUdpHeader

    // RUNT error.
    ReceivedErrorRunt MplsOam_Packet_Received_ReceivedErrorRunt

    // Dropped queue full.
    ReceivedErrorQueueFull MplsOam_Packet_Received_ReceivedErrorQueueFull

    // General error.
    ReceivedErrorGeneral MplsOam_Packet_Received_ReceivedErrorGeneral

    // Error no Interfaces.
    ReceivedErrorNoInterface MplsOam_Packet_Received_ReceivedErrorNoInterface

    // Error no memory.
    ReceivedErrorNoMemory MplsOam_Packet_Received_ReceivedErrorNoMemory

    // Protect Protocol Received good request.
    ProtectProtocolReceivedGoodRequest MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest

    // Protect Protocol Received good reply.
    ProtectProtocolReceivedGoodReply MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply

    // Received Reqeust with BFD TLV.
    ReceivedGoodBfdRequest MplsOam_Packet_Received_ReceivedGoodBfdRequest

    // Received Reply with BFD TLV.
    ReceivedGoodBfdReply MplsOam_Packet_Received_ReceivedGoodBfdReply
}

func (received *MplsOam_Packet_Received) GetFilter() yfilter.YFilter { return received.YFilter }

func (received *MplsOam_Packet_Received) SetFilter(yf yfilter.YFilter) { received.YFilter = yf }

func (received *MplsOam_Packet_Received) GetGoName(yname string) string {
    if yname == "received-good-request" { return "ReceivedGoodRequest" }
    if yname == "received-good-reply" { return "ReceivedGoodReply" }
    if yname == "received-unknown" { return "ReceivedUnknown" }
    if yname == "received-error-ip-header" { return "ReceivedErrorIpHeader" }
    if yname == "received-error-udp-header" { return "ReceivedErrorUdpHeader" }
    if yname == "received-error-runt" { return "ReceivedErrorRunt" }
    if yname == "received-error-queue-full" { return "ReceivedErrorQueueFull" }
    if yname == "received-error-general" { return "ReceivedErrorGeneral" }
    if yname == "received-error-no-interface" { return "ReceivedErrorNoInterface" }
    if yname == "received-error-no-memory" { return "ReceivedErrorNoMemory" }
    if yname == "protect-protocol-received-good-request" { return "ProtectProtocolReceivedGoodRequest" }
    if yname == "protect-protocol-received-good-reply" { return "ProtectProtocolReceivedGoodReply" }
    if yname == "received-good-bfd-request" { return "ReceivedGoodBfdRequest" }
    if yname == "received-good-bfd-reply" { return "ReceivedGoodBfdReply" }
    return ""
}

func (received *MplsOam_Packet_Received) GetSegmentPath() string {
    return "received"
}

func (received *MplsOam_Packet_Received) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "received-good-request" {
        return &received.ReceivedGoodRequest
    }
    if childYangName == "received-good-reply" {
        return &received.ReceivedGoodReply
    }
    if childYangName == "received-unknown" {
        return &received.ReceivedUnknown
    }
    if childYangName == "received-error-ip-header" {
        return &received.ReceivedErrorIpHeader
    }
    if childYangName == "received-error-udp-header" {
        return &received.ReceivedErrorUdpHeader
    }
    if childYangName == "received-error-runt" {
        return &received.ReceivedErrorRunt
    }
    if childYangName == "received-error-queue-full" {
        return &received.ReceivedErrorQueueFull
    }
    if childYangName == "received-error-general" {
        return &received.ReceivedErrorGeneral
    }
    if childYangName == "received-error-no-interface" {
        return &received.ReceivedErrorNoInterface
    }
    if childYangName == "received-error-no-memory" {
        return &received.ReceivedErrorNoMemory
    }
    if childYangName == "protect-protocol-received-good-request" {
        return &received.ProtectProtocolReceivedGoodRequest
    }
    if childYangName == "protect-protocol-received-good-reply" {
        return &received.ProtectProtocolReceivedGoodReply
    }
    if childYangName == "received-good-bfd-request" {
        return &received.ReceivedGoodBfdRequest
    }
    if childYangName == "received-good-bfd-reply" {
        return &received.ReceivedGoodBfdReply
    }
    return nil
}

func (received *MplsOam_Packet_Received) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["received-good-request"] = &received.ReceivedGoodRequest
    children["received-good-reply"] = &received.ReceivedGoodReply
    children["received-unknown"] = &received.ReceivedUnknown
    children["received-error-ip-header"] = &received.ReceivedErrorIpHeader
    children["received-error-udp-header"] = &received.ReceivedErrorUdpHeader
    children["received-error-runt"] = &received.ReceivedErrorRunt
    children["received-error-queue-full"] = &received.ReceivedErrorQueueFull
    children["received-error-general"] = &received.ReceivedErrorGeneral
    children["received-error-no-interface"] = &received.ReceivedErrorNoInterface
    children["received-error-no-memory"] = &received.ReceivedErrorNoMemory
    children["protect-protocol-received-good-request"] = &received.ProtectProtocolReceivedGoodRequest
    children["protect-protocol-received-good-reply"] = &received.ProtectProtocolReceivedGoodReply
    children["received-good-bfd-request"] = &received.ReceivedGoodBfdRequest
    children["received-good-bfd-reply"] = &received.ReceivedGoodBfdReply
    return children
}

func (received *MplsOam_Packet_Received) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (received *MplsOam_Packet_Received) GetBundleName() string { return "cisco_ios_xr" }

func (received *MplsOam_Packet_Received) GetYangName() string { return "received" }

func (received *MplsOam_Packet_Received) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (received *MplsOam_Packet_Received) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (received *MplsOam_Packet_Received) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (received *MplsOam_Packet_Received) SetParent(parent types.Entity) { received.parent = parent }

func (received *MplsOam_Packet_Received) GetParent() types.Entity { return received.parent }

func (received *MplsOam_Packet_Received) GetParentYangName() string { return "packet" }

// MplsOam_Packet_Received_ReceivedGoodRequest
// Received good request
type MplsOam_Packet_Received_ReceivedGoodRequest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetFilter() yfilter.YFilter { return receivedGoodRequest.YFilter }

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) SetFilter(yf yfilter.YFilter) { receivedGoodRequest.YFilter = yf }

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetSegmentPath() string {
    return "received-good-request"
}

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedGoodRequest.Packets
    leafs["bytes"] = receivedGoodRequest.Bytes
    return leafs
}

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetBundleName() string { return "cisco_ios_xr" }

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetYangName() string { return "received-good-request" }

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) SetParent(parent types.Entity) { receivedGoodRequest.parent = parent }

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetParent() types.Entity { return receivedGoodRequest.parent }

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetParentYangName() string { return "received" }

// MplsOam_Packet_Received_ReceivedGoodReply
// Received good reply
type MplsOam_Packet_Received_ReceivedGoodReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetFilter() yfilter.YFilter { return receivedGoodReply.YFilter }

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) SetFilter(yf yfilter.YFilter) { receivedGoodReply.YFilter = yf }

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetSegmentPath() string {
    return "received-good-reply"
}

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedGoodReply.Packets
    leafs["bytes"] = receivedGoodReply.Bytes
    return leafs
}

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetBundleName() string { return "cisco_ios_xr" }

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetYangName() string { return "received-good-reply" }

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) SetParent(parent types.Entity) { receivedGoodReply.parent = parent }

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetParent() types.Entity { return receivedGoodReply.parent }

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetParentYangName() string { return "received" }

// MplsOam_Packet_Received_ReceivedUnknown
// Received unknown packets
type MplsOam_Packet_Received_ReceivedUnknown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetFilter() yfilter.YFilter { return receivedUnknown.YFilter }

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) SetFilter(yf yfilter.YFilter) { receivedUnknown.YFilter = yf }

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetSegmentPath() string {
    return "received-unknown"
}

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedUnknown.Packets
    leafs["bytes"] = receivedUnknown.Bytes
    return leafs
}

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetBundleName() string { return "cisco_ios_xr" }

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetYangName() string { return "received-unknown" }

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) SetParent(parent types.Entity) { receivedUnknown.parent = parent }

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetParent() types.Entity { return receivedUnknown.parent }

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetParentYangName() string { return "received" }

// MplsOam_Packet_Received_ReceivedErrorIpHeader
// IP header error
type MplsOam_Packet_Received_ReceivedErrorIpHeader struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetFilter() yfilter.YFilter { return receivedErrorIpHeader.YFilter }

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) SetFilter(yf yfilter.YFilter) { receivedErrorIpHeader.YFilter = yf }

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetSegmentPath() string {
    return "received-error-ip-header"
}

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorIpHeader.Packets
    leafs["bytes"] = receivedErrorIpHeader.Bytes
    return leafs
}

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetYangName() string { return "received-error-ip-header" }

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) SetParent(parent types.Entity) { receivedErrorIpHeader.parent = parent }

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetParent() types.Entity { return receivedErrorIpHeader.parent }

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetParentYangName() string { return "received" }

// MplsOam_Packet_Received_ReceivedErrorUdpHeader
// UDP header error
type MplsOam_Packet_Received_ReceivedErrorUdpHeader struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetFilter() yfilter.YFilter { return receivedErrorUdpHeader.YFilter }

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) SetFilter(yf yfilter.YFilter) { receivedErrorUdpHeader.YFilter = yf }

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetSegmentPath() string {
    return "received-error-udp-header"
}

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorUdpHeader.Packets
    leafs["bytes"] = receivedErrorUdpHeader.Bytes
    return leafs
}

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetYangName() string { return "received-error-udp-header" }

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) SetParent(parent types.Entity) { receivedErrorUdpHeader.parent = parent }

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetParent() types.Entity { return receivedErrorUdpHeader.parent }

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetParentYangName() string { return "received" }

// MplsOam_Packet_Received_ReceivedErrorRunt
// RUNT error
type MplsOam_Packet_Received_ReceivedErrorRunt struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetFilter() yfilter.YFilter { return receivedErrorRunt.YFilter }

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) SetFilter(yf yfilter.YFilter) { receivedErrorRunt.YFilter = yf }

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetSegmentPath() string {
    return "received-error-runt"
}

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorRunt.Packets
    leafs["bytes"] = receivedErrorRunt.Bytes
    return leafs
}

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetYangName() string { return "received-error-runt" }

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) SetParent(parent types.Entity) { receivedErrorRunt.parent = parent }

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetParent() types.Entity { return receivedErrorRunt.parent }

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetParentYangName() string { return "received" }

// MplsOam_Packet_Received_ReceivedErrorQueueFull
// Dropped queue full
type MplsOam_Packet_Received_ReceivedErrorQueueFull struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetFilter() yfilter.YFilter { return receivedErrorQueueFull.YFilter }

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) SetFilter(yf yfilter.YFilter) { receivedErrorQueueFull.YFilter = yf }

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetSegmentPath() string {
    return "received-error-queue-full"
}

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorQueueFull.Packets
    leafs["bytes"] = receivedErrorQueueFull.Bytes
    return leafs
}

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetYangName() string { return "received-error-queue-full" }

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) SetParent(parent types.Entity) { receivedErrorQueueFull.parent = parent }

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetParent() types.Entity { return receivedErrorQueueFull.parent }

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetParentYangName() string { return "received" }

// MplsOam_Packet_Received_ReceivedErrorGeneral
// General error
type MplsOam_Packet_Received_ReceivedErrorGeneral struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetFilter() yfilter.YFilter { return receivedErrorGeneral.YFilter }

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) SetFilter(yf yfilter.YFilter) { receivedErrorGeneral.YFilter = yf }

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetSegmentPath() string {
    return "received-error-general"
}

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorGeneral.Packets
    leafs["bytes"] = receivedErrorGeneral.Bytes
    return leafs
}

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetYangName() string { return "received-error-general" }

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) SetParent(parent types.Entity) { receivedErrorGeneral.parent = parent }

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetParent() types.Entity { return receivedErrorGeneral.parent }

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetParentYangName() string { return "received" }

// MplsOam_Packet_Received_ReceivedErrorNoInterface
// Error no Interfaces
type MplsOam_Packet_Received_ReceivedErrorNoInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetFilter() yfilter.YFilter { return receivedErrorNoInterface.YFilter }

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) SetFilter(yf yfilter.YFilter) { receivedErrorNoInterface.YFilter = yf }

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetSegmentPath() string {
    return "received-error-no-interface"
}

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorNoInterface.Packets
    leafs["bytes"] = receivedErrorNoInterface.Bytes
    return leafs
}

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetYangName() string { return "received-error-no-interface" }

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) SetParent(parent types.Entity) { receivedErrorNoInterface.parent = parent }

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetParent() types.Entity { return receivedErrorNoInterface.parent }

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetParentYangName() string { return "received" }

// MplsOam_Packet_Received_ReceivedErrorNoMemory
// Error no memory
type MplsOam_Packet_Received_ReceivedErrorNoMemory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetFilter() yfilter.YFilter { return receivedErrorNoMemory.YFilter }

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) SetFilter(yf yfilter.YFilter) { receivedErrorNoMemory.YFilter = yf }

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetSegmentPath() string {
    return "received-error-no-memory"
}

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedErrorNoMemory.Packets
    leafs["bytes"] = receivedErrorNoMemory.Bytes
    return leafs
}

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetBundleName() string { return "cisco_ios_xr" }

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetYangName() string { return "received-error-no-memory" }

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) SetParent(parent types.Entity) { receivedErrorNoMemory.parent = parent }

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetParent() types.Entity { return receivedErrorNoMemory.parent }

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetParentYangName() string { return "received" }

// MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest
// Protect Protocol Received good request
type MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetFilter() yfilter.YFilter { return protectProtocolReceivedGoodRequest.YFilter }

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) SetFilter(yf yfilter.YFilter) { protectProtocolReceivedGoodRequest.YFilter = yf }

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetSegmentPath() string {
    return "protect-protocol-received-good-request"
}

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = protectProtocolReceivedGoodRequest.Packets
    leafs["bytes"] = protectProtocolReceivedGoodRequest.Bytes
    return leafs
}

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetBundleName() string { return "cisco_ios_xr" }

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetYangName() string { return "protect-protocol-received-good-request" }

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) SetParent(parent types.Entity) { protectProtocolReceivedGoodRequest.parent = parent }

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetParent() types.Entity { return protectProtocolReceivedGoodRequest.parent }

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetParentYangName() string { return "received" }

// MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply
// Protect Protocol Received good reply
type MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetFilter() yfilter.YFilter { return protectProtocolReceivedGoodReply.YFilter }

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) SetFilter(yf yfilter.YFilter) { protectProtocolReceivedGoodReply.YFilter = yf }

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetSegmentPath() string {
    return "protect-protocol-received-good-reply"
}

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = protectProtocolReceivedGoodReply.Packets
    leafs["bytes"] = protectProtocolReceivedGoodReply.Bytes
    return leafs
}

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetBundleName() string { return "cisco_ios_xr" }

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetYangName() string { return "protect-protocol-received-good-reply" }

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) SetParent(parent types.Entity) { protectProtocolReceivedGoodReply.parent = parent }

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetParent() types.Entity { return protectProtocolReceivedGoodReply.parent }

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetParentYangName() string { return "received" }

// MplsOam_Packet_Received_ReceivedGoodBfdRequest
// Received Reqeust with BFD TLV
type MplsOam_Packet_Received_ReceivedGoodBfdRequest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetFilter() yfilter.YFilter { return receivedGoodBfdRequest.YFilter }

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) SetFilter(yf yfilter.YFilter) { receivedGoodBfdRequest.YFilter = yf }

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetSegmentPath() string {
    return "received-good-bfd-request"
}

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedGoodBfdRequest.Packets
    leafs["bytes"] = receivedGoodBfdRequest.Bytes
    return leafs
}

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetBundleName() string { return "cisco_ios_xr" }

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetYangName() string { return "received-good-bfd-request" }

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) SetParent(parent types.Entity) { receivedGoodBfdRequest.parent = parent }

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetParent() types.Entity { return receivedGoodBfdRequest.parent }

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetParentYangName() string { return "received" }

// MplsOam_Packet_Received_ReceivedGoodBfdReply
// Received Reply with BFD TLV
type MplsOam_Packet_Received_ReceivedGoodBfdReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetFilter() yfilter.YFilter { return receivedGoodBfdReply.YFilter }

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) SetFilter(yf yfilter.YFilter) { receivedGoodBfdReply.YFilter = yf }

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetSegmentPath() string {
    return "received-good-bfd-reply"
}

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = receivedGoodBfdReply.Packets
    leafs["bytes"] = receivedGoodBfdReply.Bytes
    return leafs
}

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetBundleName() string { return "cisco_ios_xr" }

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetYangName() string { return "received-good-bfd-reply" }

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) SetParent(parent types.Entity) { receivedGoodBfdReply.parent = parent }

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetParent() types.Entity { return receivedGoodBfdReply.parent }

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetParentYangName() string { return "received" }

// MplsOam_Packet_Sent
// Packet transmit counts
type MplsOam_Packet_Sent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Packet_Sent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Packet_Sent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Packet_Sent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Packet_Sent_BfdNoReply
}

func (sent *MplsOam_Packet_Sent) GetFilter() yfilter.YFilter { return sent.YFilter }

func (sent *MplsOam_Packet_Sent) SetFilter(yf yfilter.YFilter) { sent.YFilter = yf }

func (sent *MplsOam_Packet_Sent) GetGoName(yname string) string {
    if yname == "transmit-good" { return "TransmitGood" }
    if yname == "transmit-drop" { return "TransmitDrop" }
    if yname == "transmit-bfd-good" { return "TransmitBfdGood" }
    if yname == "bfd-no-reply" { return "BfdNoReply" }
    return ""
}

func (sent *MplsOam_Packet_Sent) GetSegmentPath() string {
    return "sent"
}

func (sent *MplsOam_Packet_Sent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit-good" {
        return &sent.TransmitGood
    }
    if childYangName == "transmit-drop" {
        return &sent.TransmitDrop
    }
    if childYangName == "transmit-bfd-good" {
        return &sent.TransmitBfdGood
    }
    if childYangName == "bfd-no-reply" {
        return &sent.BfdNoReply
    }
    return nil
}

func (sent *MplsOam_Packet_Sent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit-good"] = &sent.TransmitGood
    children["transmit-drop"] = &sent.TransmitDrop
    children["transmit-bfd-good"] = &sent.TransmitBfdGood
    children["bfd-no-reply"] = &sent.BfdNoReply
    return children
}

func (sent *MplsOam_Packet_Sent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sent *MplsOam_Packet_Sent) GetBundleName() string { return "cisco_ios_xr" }

func (sent *MplsOam_Packet_Sent) GetYangName() string { return "sent" }

func (sent *MplsOam_Packet_Sent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sent *MplsOam_Packet_Sent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sent *MplsOam_Packet_Sent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sent *MplsOam_Packet_Sent) SetParent(parent types.Entity) { sent.parent = parent }

func (sent *MplsOam_Packet_Sent) GetParent() types.Entity { return sent.parent }

func (sent *MplsOam_Packet_Sent) GetParentYangName() string { return "packet" }

// MplsOam_Packet_Sent_TransmitGood
// Transmit good packets
type MplsOam_Packet_Sent_TransmitGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetFilter() yfilter.YFilter { return transmitGood.YFilter }

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) SetFilter(yf yfilter.YFilter) { transmitGood.YFilter = yf }

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetSegmentPath() string {
    return "transmit-good"
}

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitGood.Packets
    leafs["bytes"] = transmitGood.Bytes
    return leafs
}

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetYangName() string { return "transmit-good" }

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) SetParent(parent types.Entity) { transmitGood.parent = parent }

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetParent() types.Entity { return transmitGood.parent }

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetParentYangName() string { return "sent" }

// MplsOam_Packet_Sent_TransmitDrop
// Transmit drop packets
type MplsOam_Packet_Sent_TransmitDrop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetFilter() yfilter.YFilter { return transmitDrop.YFilter }

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) SetFilter(yf yfilter.YFilter) { transmitDrop.YFilter = yf }

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetSegmentPath() string {
    return "transmit-drop"
}

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitDrop.Packets
    leafs["bytes"] = transmitDrop.Bytes
    return leafs
}

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetBundleName() string { return "cisco_ios_xr" }

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetYangName() string { return "transmit-drop" }

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) SetParent(parent types.Entity) { transmitDrop.parent = parent }

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetParent() types.Entity { return transmitDrop.parent }

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetParentYangName() string { return "sent" }

// MplsOam_Packet_Sent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Packet_Sent_TransmitBfdGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetFilter() yfilter.YFilter { return transmitBfdGood.YFilter }

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) SetFilter(yf yfilter.YFilter) { transmitBfdGood.YFilter = yf }

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetSegmentPath() string {
    return "transmit-bfd-good"
}

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitBfdGood.Packets
    leafs["bytes"] = transmitBfdGood.Bytes
    return leafs
}

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetYangName() string { return "transmit-bfd-good" }

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) SetParent(parent types.Entity) { transmitBfdGood.parent = parent }

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetParent() types.Entity { return transmitBfdGood.parent }

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetParentYangName() string { return "sent" }

// MplsOam_Packet_Sent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Packet_Sent_BfdNoReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetFilter() yfilter.YFilter { return bfdNoReply.YFilter }

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) SetFilter(yf yfilter.YFilter) { bfdNoReply.YFilter = yf }

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetSegmentPath() string {
    return "bfd-no-reply"
}

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = bfdNoReply.Packets
    leafs["bytes"] = bfdNoReply.Bytes
    return leafs
}

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetBundleName() string { return "cisco_ios_xr" }

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetYangName() string { return "bfd-no-reply" }

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) SetParent(parent types.Entity) { bfdNoReply.parent = parent }

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetParent() types.Entity { return bfdNoReply.parent }

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetParentYangName() string { return "sent" }

// MplsOam_Packet_WorkingReqSent
// Working Request Packet transmit counts
type MplsOam_Packet_WorkingReqSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Packet_WorkingReqSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Packet_WorkingReqSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Packet_WorkingReqSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Packet_WorkingReqSent_BfdNoReply
}

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetFilter() yfilter.YFilter { return workingReqSent.YFilter }

func (workingReqSent *MplsOam_Packet_WorkingReqSent) SetFilter(yf yfilter.YFilter) { workingReqSent.YFilter = yf }

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetGoName(yname string) string {
    if yname == "transmit-good" { return "TransmitGood" }
    if yname == "transmit-drop" { return "TransmitDrop" }
    if yname == "transmit-bfd-good" { return "TransmitBfdGood" }
    if yname == "bfd-no-reply" { return "BfdNoReply" }
    return ""
}

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetSegmentPath() string {
    return "working-req-sent"
}

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit-good" {
        return &workingReqSent.TransmitGood
    }
    if childYangName == "transmit-drop" {
        return &workingReqSent.TransmitDrop
    }
    if childYangName == "transmit-bfd-good" {
        return &workingReqSent.TransmitBfdGood
    }
    if childYangName == "bfd-no-reply" {
        return &workingReqSent.BfdNoReply
    }
    return nil
}

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit-good"] = &workingReqSent.TransmitGood
    children["transmit-drop"] = &workingReqSent.TransmitDrop
    children["transmit-bfd-good"] = &workingReqSent.TransmitBfdGood
    children["bfd-no-reply"] = &workingReqSent.BfdNoReply
    return children
}

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetBundleName() string { return "cisco_ios_xr" }

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetYangName() string { return "working-req-sent" }

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (workingReqSent *MplsOam_Packet_WorkingReqSent) SetParent(parent types.Entity) { workingReqSent.parent = parent }

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetParent() types.Entity { return workingReqSent.parent }

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetParentYangName() string { return "packet" }

// MplsOam_Packet_WorkingReqSent_TransmitGood
// Transmit good packets
type MplsOam_Packet_WorkingReqSent_TransmitGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetFilter() yfilter.YFilter { return transmitGood.YFilter }

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) SetFilter(yf yfilter.YFilter) { transmitGood.YFilter = yf }

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetSegmentPath() string {
    return "transmit-good"
}

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitGood.Packets
    leafs["bytes"] = transmitGood.Bytes
    return leafs
}

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetYangName() string { return "transmit-good" }

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) SetParent(parent types.Entity) { transmitGood.parent = parent }

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetParent() types.Entity { return transmitGood.parent }

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetParentYangName() string { return "working-req-sent" }

// MplsOam_Packet_WorkingReqSent_TransmitDrop
// Transmit drop packets
type MplsOam_Packet_WorkingReqSent_TransmitDrop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetFilter() yfilter.YFilter { return transmitDrop.YFilter }

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) SetFilter(yf yfilter.YFilter) { transmitDrop.YFilter = yf }

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetSegmentPath() string {
    return "transmit-drop"
}

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitDrop.Packets
    leafs["bytes"] = transmitDrop.Bytes
    return leafs
}

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetBundleName() string { return "cisco_ios_xr" }

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetYangName() string { return "transmit-drop" }

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) SetParent(parent types.Entity) { transmitDrop.parent = parent }

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetParent() types.Entity { return transmitDrop.parent }

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetParentYangName() string { return "working-req-sent" }

// MplsOam_Packet_WorkingReqSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Packet_WorkingReqSent_TransmitBfdGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetFilter() yfilter.YFilter { return transmitBfdGood.YFilter }

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) SetFilter(yf yfilter.YFilter) { transmitBfdGood.YFilter = yf }

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetSegmentPath() string {
    return "transmit-bfd-good"
}

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitBfdGood.Packets
    leafs["bytes"] = transmitBfdGood.Bytes
    return leafs
}

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetYangName() string { return "transmit-bfd-good" }

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) SetParent(parent types.Entity) { transmitBfdGood.parent = parent }

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetParent() types.Entity { return transmitBfdGood.parent }

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetParentYangName() string { return "working-req-sent" }

// MplsOam_Packet_WorkingReqSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Packet_WorkingReqSent_BfdNoReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetFilter() yfilter.YFilter { return bfdNoReply.YFilter }

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) SetFilter(yf yfilter.YFilter) { bfdNoReply.YFilter = yf }

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetSegmentPath() string {
    return "bfd-no-reply"
}

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = bfdNoReply.Packets
    leafs["bytes"] = bfdNoReply.Bytes
    return leafs
}

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetBundleName() string { return "cisco_ios_xr" }

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetYangName() string { return "bfd-no-reply" }

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) SetParent(parent types.Entity) { bfdNoReply.parent = parent }

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetParent() types.Entity { return bfdNoReply.parent }

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetParentYangName() string { return "working-req-sent" }

// MplsOam_Packet_WorkingRepSent
// Working Reply Packet transmit counts
type MplsOam_Packet_WorkingRepSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Packet_WorkingRepSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Packet_WorkingRepSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Packet_WorkingRepSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Packet_WorkingRepSent_BfdNoReply
}

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetFilter() yfilter.YFilter { return workingRepSent.YFilter }

func (workingRepSent *MplsOam_Packet_WorkingRepSent) SetFilter(yf yfilter.YFilter) { workingRepSent.YFilter = yf }

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetGoName(yname string) string {
    if yname == "transmit-good" { return "TransmitGood" }
    if yname == "transmit-drop" { return "TransmitDrop" }
    if yname == "transmit-bfd-good" { return "TransmitBfdGood" }
    if yname == "bfd-no-reply" { return "BfdNoReply" }
    return ""
}

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetSegmentPath() string {
    return "working-rep-sent"
}

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit-good" {
        return &workingRepSent.TransmitGood
    }
    if childYangName == "transmit-drop" {
        return &workingRepSent.TransmitDrop
    }
    if childYangName == "transmit-bfd-good" {
        return &workingRepSent.TransmitBfdGood
    }
    if childYangName == "bfd-no-reply" {
        return &workingRepSent.BfdNoReply
    }
    return nil
}

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit-good"] = &workingRepSent.TransmitGood
    children["transmit-drop"] = &workingRepSent.TransmitDrop
    children["transmit-bfd-good"] = &workingRepSent.TransmitBfdGood
    children["bfd-no-reply"] = &workingRepSent.BfdNoReply
    return children
}

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetBundleName() string { return "cisco_ios_xr" }

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetYangName() string { return "working-rep-sent" }

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (workingRepSent *MplsOam_Packet_WorkingRepSent) SetParent(parent types.Entity) { workingRepSent.parent = parent }

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetParent() types.Entity { return workingRepSent.parent }

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetParentYangName() string { return "packet" }

// MplsOam_Packet_WorkingRepSent_TransmitGood
// Transmit good packets
type MplsOam_Packet_WorkingRepSent_TransmitGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetFilter() yfilter.YFilter { return transmitGood.YFilter }

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) SetFilter(yf yfilter.YFilter) { transmitGood.YFilter = yf }

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetSegmentPath() string {
    return "transmit-good"
}

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitGood.Packets
    leafs["bytes"] = transmitGood.Bytes
    return leafs
}

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetYangName() string { return "transmit-good" }

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) SetParent(parent types.Entity) { transmitGood.parent = parent }

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetParent() types.Entity { return transmitGood.parent }

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetParentYangName() string { return "working-rep-sent" }

// MplsOam_Packet_WorkingRepSent_TransmitDrop
// Transmit drop packets
type MplsOam_Packet_WorkingRepSent_TransmitDrop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetFilter() yfilter.YFilter { return transmitDrop.YFilter }

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) SetFilter(yf yfilter.YFilter) { transmitDrop.YFilter = yf }

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetSegmentPath() string {
    return "transmit-drop"
}

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitDrop.Packets
    leafs["bytes"] = transmitDrop.Bytes
    return leafs
}

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetBundleName() string { return "cisco_ios_xr" }

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetYangName() string { return "transmit-drop" }

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) SetParent(parent types.Entity) { transmitDrop.parent = parent }

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetParent() types.Entity { return transmitDrop.parent }

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetParentYangName() string { return "working-rep-sent" }

// MplsOam_Packet_WorkingRepSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Packet_WorkingRepSent_TransmitBfdGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetFilter() yfilter.YFilter { return transmitBfdGood.YFilter }

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) SetFilter(yf yfilter.YFilter) { transmitBfdGood.YFilter = yf }

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetSegmentPath() string {
    return "transmit-bfd-good"
}

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitBfdGood.Packets
    leafs["bytes"] = transmitBfdGood.Bytes
    return leafs
}

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetYangName() string { return "transmit-bfd-good" }

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) SetParent(parent types.Entity) { transmitBfdGood.parent = parent }

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetParent() types.Entity { return transmitBfdGood.parent }

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetParentYangName() string { return "working-rep-sent" }

// MplsOam_Packet_WorkingRepSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Packet_WorkingRepSent_BfdNoReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetFilter() yfilter.YFilter { return bfdNoReply.YFilter }

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) SetFilter(yf yfilter.YFilter) { bfdNoReply.YFilter = yf }

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetSegmentPath() string {
    return "bfd-no-reply"
}

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = bfdNoReply.Packets
    leafs["bytes"] = bfdNoReply.Bytes
    return leafs
}

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetBundleName() string { return "cisco_ios_xr" }

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetYangName() string { return "bfd-no-reply" }

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) SetParent(parent types.Entity) { bfdNoReply.parent = parent }

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetParent() types.Entity { return bfdNoReply.parent }

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetParentYangName() string { return "working-rep-sent" }

// MplsOam_Packet_ProtectReqSent
// Protect Request Packet transmit counts
type MplsOam_Packet_ProtectReqSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Packet_ProtectReqSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Packet_ProtectReqSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Packet_ProtectReqSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Packet_ProtectReqSent_BfdNoReply
}

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetFilter() yfilter.YFilter { return protectReqSent.YFilter }

func (protectReqSent *MplsOam_Packet_ProtectReqSent) SetFilter(yf yfilter.YFilter) { protectReqSent.YFilter = yf }

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetGoName(yname string) string {
    if yname == "transmit-good" { return "TransmitGood" }
    if yname == "transmit-drop" { return "TransmitDrop" }
    if yname == "transmit-bfd-good" { return "TransmitBfdGood" }
    if yname == "bfd-no-reply" { return "BfdNoReply" }
    return ""
}

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetSegmentPath() string {
    return "protect-req-sent"
}

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit-good" {
        return &protectReqSent.TransmitGood
    }
    if childYangName == "transmit-drop" {
        return &protectReqSent.TransmitDrop
    }
    if childYangName == "transmit-bfd-good" {
        return &protectReqSent.TransmitBfdGood
    }
    if childYangName == "bfd-no-reply" {
        return &protectReqSent.BfdNoReply
    }
    return nil
}

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit-good"] = &protectReqSent.TransmitGood
    children["transmit-drop"] = &protectReqSent.TransmitDrop
    children["transmit-bfd-good"] = &protectReqSent.TransmitBfdGood
    children["bfd-no-reply"] = &protectReqSent.BfdNoReply
    return children
}

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetBundleName() string { return "cisco_ios_xr" }

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetYangName() string { return "protect-req-sent" }

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protectReqSent *MplsOam_Packet_ProtectReqSent) SetParent(parent types.Entity) { protectReqSent.parent = parent }

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetParent() types.Entity { return protectReqSent.parent }

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetParentYangName() string { return "packet" }

// MplsOam_Packet_ProtectReqSent_TransmitGood
// Transmit good packets
type MplsOam_Packet_ProtectReqSent_TransmitGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetFilter() yfilter.YFilter { return transmitGood.YFilter }

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) SetFilter(yf yfilter.YFilter) { transmitGood.YFilter = yf }

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetSegmentPath() string {
    return "transmit-good"
}

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitGood.Packets
    leafs["bytes"] = transmitGood.Bytes
    return leafs
}

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetYangName() string { return "transmit-good" }

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) SetParent(parent types.Entity) { transmitGood.parent = parent }

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetParent() types.Entity { return transmitGood.parent }

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetParentYangName() string { return "protect-req-sent" }

// MplsOam_Packet_ProtectReqSent_TransmitDrop
// Transmit drop packets
type MplsOam_Packet_ProtectReqSent_TransmitDrop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetFilter() yfilter.YFilter { return transmitDrop.YFilter }

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) SetFilter(yf yfilter.YFilter) { transmitDrop.YFilter = yf }

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetSegmentPath() string {
    return "transmit-drop"
}

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitDrop.Packets
    leafs["bytes"] = transmitDrop.Bytes
    return leafs
}

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetBundleName() string { return "cisco_ios_xr" }

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetYangName() string { return "transmit-drop" }

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) SetParent(parent types.Entity) { transmitDrop.parent = parent }

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetParent() types.Entity { return transmitDrop.parent }

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetParentYangName() string { return "protect-req-sent" }

// MplsOam_Packet_ProtectReqSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Packet_ProtectReqSent_TransmitBfdGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetFilter() yfilter.YFilter { return transmitBfdGood.YFilter }

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) SetFilter(yf yfilter.YFilter) { transmitBfdGood.YFilter = yf }

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetSegmentPath() string {
    return "transmit-bfd-good"
}

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitBfdGood.Packets
    leafs["bytes"] = transmitBfdGood.Bytes
    return leafs
}

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetYangName() string { return "transmit-bfd-good" }

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) SetParent(parent types.Entity) { transmitBfdGood.parent = parent }

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetParent() types.Entity { return transmitBfdGood.parent }

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetParentYangName() string { return "protect-req-sent" }

// MplsOam_Packet_ProtectReqSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Packet_ProtectReqSent_BfdNoReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetFilter() yfilter.YFilter { return bfdNoReply.YFilter }

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) SetFilter(yf yfilter.YFilter) { bfdNoReply.YFilter = yf }

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetSegmentPath() string {
    return "bfd-no-reply"
}

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = bfdNoReply.Packets
    leafs["bytes"] = bfdNoReply.Bytes
    return leafs
}

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetBundleName() string { return "cisco_ios_xr" }

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetYangName() string { return "bfd-no-reply" }

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) SetParent(parent types.Entity) { bfdNoReply.parent = parent }

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetParent() types.Entity { return bfdNoReply.parent }

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetParentYangName() string { return "protect-req-sent" }

// MplsOam_Packet_ProtectRepSent
// Protect Reply Packet transmit counts
type MplsOam_Packet_ProtectRepSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Packet_ProtectRepSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Packet_ProtectRepSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Packet_ProtectRepSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Packet_ProtectRepSent_BfdNoReply
}

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetFilter() yfilter.YFilter { return protectRepSent.YFilter }

func (protectRepSent *MplsOam_Packet_ProtectRepSent) SetFilter(yf yfilter.YFilter) { protectRepSent.YFilter = yf }

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetGoName(yname string) string {
    if yname == "transmit-good" { return "TransmitGood" }
    if yname == "transmit-drop" { return "TransmitDrop" }
    if yname == "transmit-bfd-good" { return "TransmitBfdGood" }
    if yname == "bfd-no-reply" { return "BfdNoReply" }
    return ""
}

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetSegmentPath() string {
    return "protect-rep-sent"
}

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmit-good" {
        return &protectRepSent.TransmitGood
    }
    if childYangName == "transmit-drop" {
        return &protectRepSent.TransmitDrop
    }
    if childYangName == "transmit-bfd-good" {
        return &protectRepSent.TransmitBfdGood
    }
    if childYangName == "bfd-no-reply" {
        return &protectRepSent.BfdNoReply
    }
    return nil
}

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmit-good"] = &protectRepSent.TransmitGood
    children["transmit-drop"] = &protectRepSent.TransmitDrop
    children["transmit-bfd-good"] = &protectRepSent.TransmitBfdGood
    children["bfd-no-reply"] = &protectRepSent.BfdNoReply
    return children
}

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetBundleName() string { return "cisco_ios_xr" }

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetYangName() string { return "protect-rep-sent" }

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protectRepSent *MplsOam_Packet_ProtectRepSent) SetParent(parent types.Entity) { protectRepSent.parent = parent }

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetParent() types.Entity { return protectRepSent.parent }

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetParentYangName() string { return "packet" }

// MplsOam_Packet_ProtectRepSent_TransmitGood
// Transmit good packets
type MplsOam_Packet_ProtectRepSent_TransmitGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetFilter() yfilter.YFilter { return transmitGood.YFilter }

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) SetFilter(yf yfilter.YFilter) { transmitGood.YFilter = yf }

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetSegmentPath() string {
    return "transmit-good"
}

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitGood.Packets
    leafs["bytes"] = transmitGood.Bytes
    return leafs
}

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetYangName() string { return "transmit-good" }

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) SetParent(parent types.Entity) { transmitGood.parent = parent }

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetParent() types.Entity { return transmitGood.parent }

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetParentYangName() string { return "protect-rep-sent" }

// MplsOam_Packet_ProtectRepSent_TransmitDrop
// Transmit drop packets
type MplsOam_Packet_ProtectRepSent_TransmitDrop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetFilter() yfilter.YFilter { return transmitDrop.YFilter }

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) SetFilter(yf yfilter.YFilter) { transmitDrop.YFilter = yf }

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetSegmentPath() string {
    return "transmit-drop"
}

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitDrop.Packets
    leafs["bytes"] = transmitDrop.Bytes
    return leafs
}

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetBundleName() string { return "cisco_ios_xr" }

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetYangName() string { return "transmit-drop" }

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) SetParent(parent types.Entity) { transmitDrop.parent = parent }

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetParent() types.Entity { return transmitDrop.parent }

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetParentYangName() string { return "protect-rep-sent" }

// MplsOam_Packet_ProtectRepSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Packet_ProtectRepSent_TransmitBfdGood struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetFilter() yfilter.YFilter { return transmitBfdGood.YFilter }

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) SetFilter(yf yfilter.YFilter) { transmitBfdGood.YFilter = yf }

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetSegmentPath() string {
    return "transmit-bfd-good"
}

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = transmitBfdGood.Packets
    leafs["bytes"] = transmitBfdGood.Bytes
    return leafs
}

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetBundleName() string { return "cisco_ios_xr" }

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetYangName() string { return "transmit-bfd-good" }

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) SetParent(parent types.Entity) { transmitBfdGood.parent = parent }

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetParent() types.Entity { return transmitBfdGood.parent }

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetParentYangName() string { return "protect-rep-sent" }

// MplsOam_Packet_ProtectRepSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Packet_ProtectRepSent_BfdNoReply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetFilter() yfilter.YFilter { return bfdNoReply.YFilter }

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) SetFilter(yf yfilter.YFilter) { bfdNoReply.YFilter = yf }

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetGoName(yname string) string {
    if yname == "packets" { return "Packets" }
    if yname == "bytes" { return "Bytes" }
    return ""
}

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetSegmentPath() string {
    return "bfd-no-reply"
}

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets"] = bfdNoReply.Packets
    leafs["bytes"] = bfdNoReply.Bytes
    return leafs
}

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetBundleName() string { return "cisco_ios_xr" }

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetYangName() string { return "bfd-no-reply" }

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) SetParent(parent types.Entity) { bfdNoReply.parent = parent }

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetParent() types.Entity { return bfdNoReply.parent }

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetParentYangName() string { return "protect-rep-sent" }

// MplsOam_Global
// LSPV global counters operational data
type MplsOam_Global struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of clients. The type is interface{} with range: 0..4294967295.
    TotalClients interface{}

    // Message statistics.
    MessageStatistics MplsOam_Global_MessageStatistics

    // Collaborator statistics.
    CollaboratorStatistics MplsOam_Global_CollaboratorStatistics
}

func (global *MplsOam_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *MplsOam_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *MplsOam_Global) GetGoName(yname string) string {
    if yname == "total-clients" { return "TotalClients" }
    if yname == "message-statistics" { return "MessageStatistics" }
    if yname == "collaborator-statistics" { return "CollaboratorStatistics" }
    return ""
}

func (global *MplsOam_Global) GetSegmentPath() string {
    return "global"
}

func (global *MplsOam_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "message-statistics" {
        return &global.MessageStatistics
    }
    if childYangName == "collaborator-statistics" {
        return &global.CollaboratorStatistics
    }
    return nil
}

func (global *MplsOam_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["message-statistics"] = &global.MessageStatistics
    children["collaborator-statistics"] = &global.CollaboratorStatistics
    return children
}

func (global *MplsOam_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-clients"] = global.TotalClients
    return leafs
}

func (global *MplsOam_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *MplsOam_Global) GetYangName() string { return "global" }

func (global *MplsOam_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *MplsOam_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *MplsOam_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *MplsOam_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *MplsOam_Global) GetParent() types.Entity { return global.parent }

func (global *MplsOam_Global) GetParentYangName() string { return "mpls-oam" }

// MplsOam_Global_MessageStatistics
// Message statistics
type MplsOam_Global_MessageStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Message register count. The type is interface{} with range: 0..4294967295.
    RegisterMessages interface{}

    // Message unregister count. The type is interface{} with range:
    // 0..4294967295.
    UnregisterMessages interface{}

    // Message echo submit count. The type is interface{} with range:
    // 0..4294967295.
    EchoSubmitMessages interface{}

    // Message echo cancel count. The type is interface{} with range:
    // 0..4294967295.
    EchoCancelMessages interface{}

    // Message get results count. The type is interface{} with range:
    // 0..4294967295.
    GetResultMessages interface{}

    // Message get configiuration count. The type is interface{} with range:
    // 0..4294967295.
    GetConfigMessages interface{}

    // Message get response count. The type is interface{} with range:
    // 0..4294967295.
    GetResponseMessages interface{}

    // Message property response count. The type is interface{} with range:
    // 0..4294967295.
    PropertyResponseMessages interface{}

    // Message property request count. The type is interface{} with range:
    // 0..4294967295.
    PropertyRequestMessages interface{}

    // Message property block count. The type is interface{} with range:
    // 0..4294967295.
    PropertyBlockMessages interface{}

    // Message thread request count. The type is interface{} with range:
    // 0..4294967295.
    ThreadRequestMessages interface{}
}

func (messageStatistics *MplsOam_Global_MessageStatistics) GetFilter() yfilter.YFilter { return messageStatistics.YFilter }

func (messageStatistics *MplsOam_Global_MessageStatistics) SetFilter(yf yfilter.YFilter) { messageStatistics.YFilter = yf }

func (messageStatistics *MplsOam_Global_MessageStatistics) GetGoName(yname string) string {
    if yname == "register-messages" { return "RegisterMessages" }
    if yname == "unregister-messages" { return "UnregisterMessages" }
    if yname == "echo-submit-messages" { return "EchoSubmitMessages" }
    if yname == "echo-cancel-messages" { return "EchoCancelMessages" }
    if yname == "get-result-messages" { return "GetResultMessages" }
    if yname == "get-config-messages" { return "GetConfigMessages" }
    if yname == "get-response-messages" { return "GetResponseMessages" }
    if yname == "property-response-messages" { return "PropertyResponseMessages" }
    if yname == "property-request-messages" { return "PropertyRequestMessages" }
    if yname == "property-block-messages" { return "PropertyBlockMessages" }
    if yname == "thread-request-messages" { return "ThreadRequestMessages" }
    return ""
}

func (messageStatistics *MplsOam_Global_MessageStatistics) GetSegmentPath() string {
    return "message-statistics"
}

func (messageStatistics *MplsOam_Global_MessageStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (messageStatistics *MplsOam_Global_MessageStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (messageStatistics *MplsOam_Global_MessageStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["register-messages"] = messageStatistics.RegisterMessages
    leafs["unregister-messages"] = messageStatistics.UnregisterMessages
    leafs["echo-submit-messages"] = messageStatistics.EchoSubmitMessages
    leafs["echo-cancel-messages"] = messageStatistics.EchoCancelMessages
    leafs["get-result-messages"] = messageStatistics.GetResultMessages
    leafs["get-config-messages"] = messageStatistics.GetConfigMessages
    leafs["get-response-messages"] = messageStatistics.GetResponseMessages
    leafs["property-response-messages"] = messageStatistics.PropertyResponseMessages
    leafs["property-request-messages"] = messageStatistics.PropertyRequestMessages
    leafs["property-block-messages"] = messageStatistics.PropertyBlockMessages
    leafs["thread-request-messages"] = messageStatistics.ThreadRequestMessages
    return leafs
}

func (messageStatistics *MplsOam_Global_MessageStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (messageStatistics *MplsOam_Global_MessageStatistics) GetYangName() string { return "message-statistics" }

func (messageStatistics *MplsOam_Global_MessageStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (messageStatistics *MplsOam_Global_MessageStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (messageStatistics *MplsOam_Global_MessageStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (messageStatistics *MplsOam_Global_MessageStatistics) SetParent(parent types.Entity) { messageStatistics.parent = parent }

func (messageStatistics *MplsOam_Global_MessageStatistics) GetParent() types.Entity { return messageStatistics.parent }

func (messageStatistics *MplsOam_Global_MessageStatistics) GetParentYangName() string { return "global" }

// MplsOam_Global_CollaboratorStatistics
// Collaborator statistics
type MplsOam_Global_CollaboratorStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collaborator IPARM counts.
    CollaboratorIParm MplsOam_Global_CollaboratorStatistics_CollaboratorIParm

    // Collaborator IM counts.
    CollaboratorIm MplsOam_Global_CollaboratorStatistics_CollaboratorIm

    // Collaborator NetIO counts.
    CollaboratorNetIo MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo

    // Collaborator RIB counts.
    CollaboratorRib MplsOam_Global_CollaboratorStatistics_CollaboratorRib
}

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetFilter() yfilter.YFilter { return collaboratorStatistics.YFilter }

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) SetFilter(yf yfilter.YFilter) { collaboratorStatistics.YFilter = yf }

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetGoName(yname string) string {
    if yname == "collaborator-i-parm" { return "CollaboratorIParm" }
    if yname == "collaborator-im" { return "CollaboratorIm" }
    if yname == "collaborator-net-io" { return "CollaboratorNetIo" }
    if yname == "collaborator-rib" { return "CollaboratorRib" }
    return ""
}

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetSegmentPath() string {
    return "collaborator-statistics"
}

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "collaborator-i-parm" {
        return &collaboratorStatistics.CollaboratorIParm
    }
    if childYangName == "collaborator-im" {
        return &collaboratorStatistics.CollaboratorIm
    }
    if childYangName == "collaborator-net-io" {
        return &collaboratorStatistics.CollaboratorNetIo
    }
    if childYangName == "collaborator-rib" {
        return &collaboratorStatistics.CollaboratorRib
    }
    return nil
}

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["collaborator-i-parm"] = &collaboratorStatistics.CollaboratorIParm
    children["collaborator-im"] = &collaboratorStatistics.CollaboratorIm
    children["collaborator-net-io"] = &collaboratorStatistics.CollaboratorNetIo
    children["collaborator-rib"] = &collaboratorStatistics.CollaboratorRib
    return children
}

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetYangName() string { return "collaborator-statistics" }

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) SetParent(parent types.Entity) { collaboratorStatistics.parent = parent }

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetParent() types.Entity { return collaboratorStatistics.parent }

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetParentYangName() string { return "global" }

// MplsOam_Global_CollaboratorStatistics_CollaboratorIParm
// Collaborator IPARM counts
type MplsOam_Global_CollaboratorStatistics_CollaboratorIParm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collaborator up counter. The type is interface{} with range: 0..4294967295.
    Ups interface{}

    // Collaborator down counter. The type is interface{} with range:
    // 0..4294967295.
    Downs interface{}
}

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetFilter() yfilter.YFilter { return collaboratorIParm.YFilter }

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) SetFilter(yf yfilter.YFilter) { collaboratorIParm.YFilter = yf }

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetGoName(yname string) string {
    if yname == "ups" { return "Ups" }
    if yname == "downs" { return "Downs" }
    return ""
}

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetSegmentPath() string {
    return "collaborator-i-parm"
}

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ups"] = collaboratorIParm.Ups
    leafs["downs"] = collaboratorIParm.Downs
    return leafs
}

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetBundleName() string { return "cisco_ios_xr" }

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetYangName() string { return "collaborator-i-parm" }

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) SetParent(parent types.Entity) { collaboratorIParm.parent = parent }

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetParent() types.Entity { return collaboratorIParm.parent }

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetParentYangName() string { return "collaborator-statistics" }

// MplsOam_Global_CollaboratorStatistics_CollaboratorIm
// Collaborator IM counts
type MplsOam_Global_CollaboratorStatistics_CollaboratorIm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collaborator up counter. The type is interface{} with range: 0..4294967295.
    Ups interface{}

    // Collaborator down counter. The type is interface{} with range:
    // 0..4294967295.
    Downs interface{}
}

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetFilter() yfilter.YFilter { return collaboratorIm.YFilter }

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) SetFilter(yf yfilter.YFilter) { collaboratorIm.YFilter = yf }

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetGoName(yname string) string {
    if yname == "ups" { return "Ups" }
    if yname == "downs" { return "Downs" }
    return ""
}

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetSegmentPath() string {
    return "collaborator-im"
}

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ups"] = collaboratorIm.Ups
    leafs["downs"] = collaboratorIm.Downs
    return leafs
}

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetBundleName() string { return "cisco_ios_xr" }

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetYangName() string { return "collaborator-im" }

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) SetParent(parent types.Entity) { collaboratorIm.parent = parent }

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetParent() types.Entity { return collaboratorIm.parent }

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetParentYangName() string { return "collaborator-statistics" }

// MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo
// Collaborator NetIO counts
type MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collaborator up counter. The type is interface{} with range: 0..4294967295.
    Ups interface{}

    // Collaborator down counter. The type is interface{} with range:
    // 0..4294967295.
    Downs interface{}
}

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetFilter() yfilter.YFilter { return collaboratorNetIo.YFilter }

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) SetFilter(yf yfilter.YFilter) { collaboratorNetIo.YFilter = yf }

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetGoName(yname string) string {
    if yname == "ups" { return "Ups" }
    if yname == "downs" { return "Downs" }
    return ""
}

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetSegmentPath() string {
    return "collaborator-net-io"
}

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ups"] = collaboratorNetIo.Ups
    leafs["downs"] = collaboratorNetIo.Downs
    return leafs
}

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetBundleName() string { return "cisco_ios_xr" }

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetYangName() string { return "collaborator-net-io" }

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) SetParent(parent types.Entity) { collaboratorNetIo.parent = parent }

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetParent() types.Entity { return collaboratorNetIo.parent }

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetParentYangName() string { return "collaborator-statistics" }

// MplsOam_Global_CollaboratorStatistics_CollaboratorRib
// Collaborator RIB counts
type MplsOam_Global_CollaboratorStatistics_CollaboratorRib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collaborator up counter. The type is interface{} with range: 0..4294967295.
    Ups interface{}

    // Collaborator down counter. The type is interface{} with range:
    // 0..4294967295.
    Downs interface{}
}

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetFilter() yfilter.YFilter { return collaboratorRib.YFilter }

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) SetFilter(yf yfilter.YFilter) { collaboratorRib.YFilter = yf }

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetGoName(yname string) string {
    if yname == "ups" { return "Ups" }
    if yname == "downs" { return "Downs" }
    return ""
}

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetSegmentPath() string {
    return "collaborator-rib"
}

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ups"] = collaboratorRib.Ups
    leafs["downs"] = collaboratorRib.Downs
    return leafs
}

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetBundleName() string { return "cisco_ios_xr" }

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetYangName() string { return "collaborator-rib" }

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) SetParent(parent types.Entity) { collaboratorRib.parent = parent }

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetParent() types.Entity { return collaboratorRib.parent }

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetParentYangName() string { return "collaborator-statistics" }

