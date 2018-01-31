// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-arp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   arp: ARP configuraiton
//   iedge-cfg: iedge cfg
//   arpgmp: arpgmp
//   arp-redundancy: arp redundancy
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_arp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_arp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-arp-cfg arp}", reflect.TypeOf(Arp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-arp-cfg:arp", reflect.TypeOf(Arp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-arp-cfg iedge-cfg}", reflect.TypeOf(IedgeCfg{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-arp-cfg:iedge-cfg", reflect.TypeOf(IedgeCfg{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-arp-cfg arpgmp}", reflect.TypeOf(Arpgmp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-arp-cfg:arpgmp", reflect.TypeOf(Arpgmp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-arp-cfg arp-redundancy}", reflect.TypeOf(ArpRedundancy{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-arp-cfg:arp-redundancy", reflect.TypeOf(ArpRedundancy{}))
}

// ArpEntry represents Arp entry
type ArpEntry string

const (
    // Static ARP entry type
    ArpEntry_static ArpEntry = "static"

    // Alias ARP entry type
    ArpEntry_alias ArpEntry = "alias"
)

// ArpEncap represents Arp encap
type ArpEncap string

const (
    // Encapsulation type ARPA
    ArpEncap_arpa ArpEncap = "arpa"

    // Encapsulation type SRP
    ArpEncap_srp ArpEncap = "srp"

    // Encapsulation type SRPA
    ArpEncap_srpa ArpEncap = "srpa"

    // Encapsulation type SRPB
    ArpEncap_srpb ArpEncap = "srpb"
)

// Arp
// ARP configuraiton
type Arp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure maximum number of safe ARP entries per line card. The type is
    // interface{} with range: 1..256000.
    MaxEntries interface{}

    // Configure inner cos values for arp packets. The type is interface{} with
    // range: 0..7.
    InnerCos interface{}

    // Configure outer cos values for arp packets. The type is interface{} with
    // range: 0..7.
    OuterCos interface{}
}

func (arp *Arp) GetFilter() yfilter.YFilter { return arp.YFilter }

func (arp *Arp) SetFilter(yf yfilter.YFilter) { arp.YFilter = yf }

func (arp *Arp) GetGoName(yname string) string {
    if yname == "max-entries" { return "MaxEntries" }
    if yname == "inner-cos" { return "InnerCos" }
    if yname == "outer-cos" { return "OuterCos" }
    return ""
}

func (arp *Arp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-arp-cfg:arp"
}

func (arp *Arp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (arp *Arp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (arp *Arp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-entries"] = arp.MaxEntries
    leafs["inner-cos"] = arp.InnerCos
    leafs["outer-cos"] = arp.OuterCos
    return leafs
}

func (arp *Arp) GetBundleName() string { return "cisco_ios_xr" }

func (arp *Arp) GetYangName() string { return "arp" }

func (arp *Arp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (arp *Arp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (arp *Arp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (arp *Arp) SetParent(parent types.Entity) { arp.parent = parent }

func (arp *Arp) GetParent() types.Entity { return arp.parent }

func (arp *Arp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-arp-cfg" }

// IedgeCfg
// iedge cfg
type IedgeCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ARP Subscriber Enable Unconditional Proxy ARP. The type is interface{}.
    SubscriberUncondProxy interface{}

    // ARP Subscriber Scale Mode Configuration. The type is interface{}.
    SubscriberScaleMode interface{}
}

func (iedgeCfg *IedgeCfg) GetFilter() yfilter.YFilter { return iedgeCfg.YFilter }

func (iedgeCfg *IedgeCfg) SetFilter(yf yfilter.YFilter) { iedgeCfg.YFilter = yf }

func (iedgeCfg *IedgeCfg) GetGoName(yname string) string {
    if yname == "subscriber-uncond-proxy" { return "SubscriberUncondProxy" }
    if yname == "subscriber-scale-mode" { return "SubscriberScaleMode" }
    return ""
}

func (iedgeCfg *IedgeCfg) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-arp-cfg:iedge-cfg"
}

func (iedgeCfg *IedgeCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (iedgeCfg *IedgeCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (iedgeCfg *IedgeCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscriber-uncond-proxy"] = iedgeCfg.SubscriberUncondProxy
    leafs["subscriber-scale-mode"] = iedgeCfg.SubscriberScaleMode
    return leafs
}

func (iedgeCfg *IedgeCfg) GetBundleName() string { return "cisco_ios_xr" }

func (iedgeCfg *IedgeCfg) GetYangName() string { return "iedge-cfg" }

func (iedgeCfg *IedgeCfg) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (iedgeCfg *IedgeCfg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (iedgeCfg *IedgeCfg) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (iedgeCfg *IedgeCfg) SetParent(parent types.Entity) { iedgeCfg.parent = parent }

func (iedgeCfg *IedgeCfg) GetParent() types.Entity { return iedgeCfg.parent }

func (iedgeCfg *IedgeCfg) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-arp-cfg" }

// Arpgmp
// arpgmp
type Arpgmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per VRF configuration, for the default VRF use 'default'. The type is slice
    // of Arpgmp_Vrf.
    Vrf []Arpgmp_Vrf
}

func (arpgmp *Arpgmp) GetFilter() yfilter.YFilter { return arpgmp.YFilter }

func (arpgmp *Arpgmp) SetFilter(yf yfilter.YFilter) { arpgmp.YFilter = yf }

func (arpgmp *Arpgmp) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (arpgmp *Arpgmp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-arp-cfg:arpgmp"
}

func (arpgmp *Arpgmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range arpgmp.Vrf {
            if arpgmp.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Arpgmp_Vrf{}
        arpgmp.Vrf = append(arpgmp.Vrf, child)
        return &arpgmp.Vrf[len(arpgmp.Vrf)-1]
    }
    return nil
}

func (arpgmp *Arpgmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range arpgmp.Vrf {
        children[arpgmp.Vrf[i].GetSegmentPath()] = &arpgmp.Vrf[i]
    }
    return children
}

func (arpgmp *Arpgmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (arpgmp *Arpgmp) GetBundleName() string { return "cisco_ios_xr" }

func (arpgmp *Arpgmp) GetYangName() string { return "arpgmp" }

func (arpgmp *Arpgmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (arpgmp *Arpgmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (arpgmp *Arpgmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (arpgmp *Arpgmp) SetParent(parent types.Entity) { arpgmp.parent = parent }

func (arpgmp *Arpgmp) GetParent() types.Entity { return arpgmp.parent }

func (arpgmp *Arpgmp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-arp-cfg" }

// Arpgmp_Vrf
// Per VRF configuration, for the default VRF use
// 'default'
type Arpgmp_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // ARP static and alias entry configuration.
    Entries Arpgmp_Vrf_Entries
}

func (vrf *Arpgmp_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Arpgmp_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Arpgmp_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "entries" { return "Entries" }
    return ""
}

func (vrf *Arpgmp_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Arpgmp_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entries" {
        return &vrf.Entries
    }
    return nil
}

func (vrf *Arpgmp_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["entries"] = &vrf.Entries
    return children
}

func (vrf *Arpgmp_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Arpgmp_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Arpgmp_Vrf) GetYangName() string { return "vrf" }

func (vrf *Arpgmp_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Arpgmp_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Arpgmp_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Arpgmp_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Arpgmp_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Arpgmp_Vrf) GetParentYangName() string { return "arpgmp" }

// Arpgmp_Vrf_Entries
// ARP static and alias entry configuration
type Arpgmp_Vrf_Entries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ARP static and alias entry configuration item. The type is slice of
    // Arpgmp_Vrf_Entries_Entry.
    Entry []Arpgmp_Vrf_Entries_Entry
}

func (entries *Arpgmp_Vrf_Entries) GetFilter() yfilter.YFilter { return entries.YFilter }

func (entries *Arpgmp_Vrf_Entries) SetFilter(yf yfilter.YFilter) { entries.YFilter = yf }

func (entries *Arpgmp_Vrf_Entries) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (entries *Arpgmp_Vrf_Entries) GetSegmentPath() string {
    return "entries"
}

func (entries *Arpgmp_Vrf_Entries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entry" {
        for _, c := range entries.Entry {
            if entries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Arpgmp_Vrf_Entries_Entry{}
        entries.Entry = append(entries.Entry, child)
        return &entries.Entry[len(entries.Entry)-1]
    }
    return nil
}

func (entries *Arpgmp_Vrf_Entries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entries.Entry {
        children[entries.Entry[i].GetSegmentPath()] = &entries.Entry[i]
    }
    return children
}

func (entries *Arpgmp_Vrf_Entries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entries *Arpgmp_Vrf_Entries) GetBundleName() string { return "cisco_ios_xr" }

func (entries *Arpgmp_Vrf_Entries) GetYangName() string { return "entries" }

func (entries *Arpgmp_Vrf_Entries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entries *Arpgmp_Vrf_Entries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entries *Arpgmp_Vrf_Entries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entries *Arpgmp_Vrf_Entries) SetParent(parent types.Entity) { entries.parent = parent }

func (entries *Arpgmp_Vrf_Entries) GetParent() types.Entity { return entries.parent }

func (entries *Arpgmp_Vrf_Entries) GetParentYangName() string { return "vrf" }

// Arpgmp_Vrf_Entries_Entry
// ARP static and alias entry configuration item
type Arpgmp_Vrf_Entries_Entry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Encapsulation type. The type is ArpEncap.
    Encapsulation interface{}

    // Entry type. The type is ArpEntry.
    EntryType interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}
}

func (entry *Arpgmp_Vrf_Entries_Entry) GetFilter() yfilter.YFilter { return entry.YFilter }

func (entry *Arpgmp_Vrf_Entries_Entry) SetFilter(yf yfilter.YFilter) { entry.YFilter = yf }

func (entry *Arpgmp_Vrf_Entries_Entry) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "entry-type" { return "EntryType" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (entry *Arpgmp_Vrf_Entries_Entry) GetSegmentPath() string {
    return "entry" + "[address='" + fmt.Sprintf("%v", entry.Address) + "']"
}

func (entry *Arpgmp_Vrf_Entries_Entry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entry *Arpgmp_Vrf_Entries_Entry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entry *Arpgmp_Vrf_Entries_Entry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = entry.Address
    leafs["mac-address"] = entry.MacAddress
    leafs["encapsulation"] = entry.Encapsulation
    leafs["entry-type"] = entry.EntryType
    leafs["interface"] = entry.Interface
    return leafs
}

func (entry *Arpgmp_Vrf_Entries_Entry) GetBundleName() string { return "cisco_ios_xr" }

func (entry *Arpgmp_Vrf_Entries_Entry) GetYangName() string { return "entry" }

func (entry *Arpgmp_Vrf_Entries_Entry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entry *Arpgmp_Vrf_Entries_Entry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entry *Arpgmp_Vrf_Entries_Entry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entry *Arpgmp_Vrf_Entries_Entry) SetParent(parent types.Entity) { entry.parent = parent }

func (entry *Arpgmp_Vrf_Entries_Entry) GetParent() types.Entity { return entry.parent }

func (entry *Arpgmp_Vrf_Entries_Entry) GetParentYangName() string { return "entries" }

// ArpRedundancy
// arp redundancy
type ArpRedundancy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure parameter for ARP Geo redundancy.
    Redundancy ArpRedundancy_Redundancy
}

func (arpRedundancy *ArpRedundancy) GetFilter() yfilter.YFilter { return arpRedundancy.YFilter }

func (arpRedundancy *ArpRedundancy) SetFilter(yf yfilter.YFilter) { arpRedundancy.YFilter = yf }

func (arpRedundancy *ArpRedundancy) GetGoName(yname string) string {
    if yname == "redundancy" { return "Redundancy" }
    return ""
}

func (arpRedundancy *ArpRedundancy) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-arp-cfg:arp-redundancy"
}

func (arpRedundancy *ArpRedundancy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "redundancy" {
        return &arpRedundancy.Redundancy
    }
    return nil
}

func (arpRedundancy *ArpRedundancy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["redundancy"] = &arpRedundancy.Redundancy
    return children
}

func (arpRedundancy *ArpRedundancy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (arpRedundancy *ArpRedundancy) GetBundleName() string { return "cisco_ios_xr" }

func (arpRedundancy *ArpRedundancy) GetYangName() string { return "arp-redundancy" }

func (arpRedundancy *ArpRedundancy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (arpRedundancy *ArpRedundancy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (arpRedundancy *ArpRedundancy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (arpRedundancy *ArpRedundancy) SetParent(parent types.Entity) { arpRedundancy.parent = parent }

func (arpRedundancy *ArpRedundancy) GetParent() types.Entity { return arpRedundancy.parent }

func (arpRedundancy *ArpRedundancy) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-arp-cfg" }

// ArpRedundancy_Redundancy
// Configure parameter for ARP Geo redundancy
// This type is a presence type.
type ArpRedundancy_Redundancy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Configure parameter for ARP Geo redundancy. Deletion of this object
    // also causes deletion of all associated objects under ArpRedundancy. The
    // type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Table of Group.
    Groups ArpRedundancy_Redundancy_Groups
}

func (redundancy *ArpRedundancy_Redundancy) GetFilter() yfilter.YFilter { return redundancy.YFilter }

func (redundancy *ArpRedundancy_Redundancy) SetFilter(yf yfilter.YFilter) { redundancy.YFilter = yf }

func (redundancy *ArpRedundancy_Redundancy) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "groups" { return "Groups" }
    return ""
}

func (redundancy *ArpRedundancy_Redundancy) GetSegmentPath() string {
    return "redundancy"
}

func (redundancy *ArpRedundancy_Redundancy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groups" {
        return &redundancy.Groups
    }
    return nil
}

func (redundancy *ArpRedundancy_Redundancy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["groups"] = &redundancy.Groups
    return children
}

func (redundancy *ArpRedundancy_Redundancy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = redundancy.Enable
    return leafs
}

func (redundancy *ArpRedundancy_Redundancy) GetBundleName() string { return "cisco_ios_xr" }

func (redundancy *ArpRedundancy_Redundancy) GetYangName() string { return "redundancy" }

func (redundancy *ArpRedundancy_Redundancy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redundancy *ArpRedundancy_Redundancy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redundancy *ArpRedundancy_Redundancy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redundancy *ArpRedundancy_Redundancy) SetParent(parent types.Entity) { redundancy.parent = parent }

func (redundancy *ArpRedundancy_Redundancy) GetParent() types.Entity { return redundancy.parent }

func (redundancy *ArpRedundancy_Redundancy) GetParentYangName() string { return "arp-redundancy" }

// ArpRedundancy_Redundancy_Groups
// Table of Group
type ArpRedundancy_Redundancy_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of ArpRedundancy_Redundancy_Groups_Group.
    Group []ArpRedundancy_Redundancy_Groups_Group
}

func (groups *ArpRedundancy_Redundancy_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *ArpRedundancy_Redundancy_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *ArpRedundancy_Redundancy_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *ArpRedundancy_Redundancy_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *ArpRedundancy_Redundancy_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ArpRedundancy_Redundancy_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *ArpRedundancy_Redundancy_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *ArpRedundancy_Redundancy_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *ArpRedundancy_Redundancy_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *ArpRedundancy_Redundancy_Groups) GetYangName() string { return "groups" }

func (groups *ArpRedundancy_Redundancy_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *ArpRedundancy_Redundancy_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *ArpRedundancy_Redundancy_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *ArpRedundancy_Redundancy_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *ArpRedundancy_Redundancy_Groups) GetParent() types.Entity { return groups.parent }

func (groups *ArpRedundancy_Redundancy_Groups) GetParentYangName() string { return "redundancy" }

// ArpRedundancy_Redundancy_Groups_Group
// None
type ArpRedundancy_Redundancy_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group ID. The type is interface{} with range:
    // 1..32.
    GroupId interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Table of Peer.
    Peers ArpRedundancy_Redundancy_Groups_Group_Peers

    // List of Interfaces for this Group.
    InterfaceList ArpRedundancy_Redundancy_Groups_Group_InterfaceList
}

func (group *ArpRedundancy_Redundancy_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *ArpRedundancy_Redundancy_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *ArpRedundancy_Redundancy_Groups_Group) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "peers" { return "Peers" }
    if yname == "interface-list" { return "InterfaceList" }
    return ""
}

func (group *ArpRedundancy_Redundancy_Groups_Group) GetSegmentPath() string {
    return "group" + "[group-id='" + fmt.Sprintf("%v", group.GroupId) + "']"
}

func (group *ArpRedundancy_Redundancy_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peers" {
        return &group.Peers
    }
    if childYangName == "interface-list" {
        return &group.InterfaceList
    }
    return nil
}

func (group *ArpRedundancy_Redundancy_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peers"] = &group.Peers
    children["interface-list"] = &group.InterfaceList
    return children
}

func (group *ArpRedundancy_Redundancy_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = group.GroupId
    leafs["source-interface"] = group.SourceInterface
    return leafs
}

func (group *ArpRedundancy_Redundancy_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *ArpRedundancy_Redundancy_Groups_Group) GetYangName() string { return "group" }

func (group *ArpRedundancy_Redundancy_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *ArpRedundancy_Redundancy_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *ArpRedundancy_Redundancy_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *ArpRedundancy_Redundancy_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *ArpRedundancy_Redundancy_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *ArpRedundancy_Redundancy_Groups_Group) GetParentYangName() string { return "groups" }

// ArpRedundancy_Redundancy_Groups_Group_Peers
// Table of Peer
type ArpRedundancy_Redundancy_Groups_Group_Peers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of
    // ArpRedundancy_Redundancy_Groups_Group_Peers_Peer.
    Peer []ArpRedundancy_Redundancy_Groups_Group_Peers_Peer
}

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetFilter() yfilter.YFilter { return peers.YFilter }

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) SetFilter(yf yfilter.YFilter) { peers.YFilter = yf }

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetGoName(yname string) string {
    if yname == "peer" { return "Peer" }
    return ""
}

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetSegmentPath() string {
    return "peers"
}

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer" {
        for _, c := range peers.Peer {
            if peers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ArpRedundancy_Redundancy_Groups_Group_Peers_Peer{}
        peers.Peer = append(peers.Peer, child)
        return &peers.Peer[len(peers.Peer)-1]
    }
    return nil
}

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peers.Peer {
        children[peers.Peer[i].GetSegmentPath()] = &peers.Peer[i]
    }
    return children
}

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetBundleName() string { return "cisco_ios_xr" }

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetYangName() string { return "peers" }

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) SetParent(parent types.Entity) { peers.parent = parent }

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetParent() types.Entity { return peers.parent }

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetParentYangName() string { return "group" }

// ArpRedundancy_Redundancy_Groups_Group_Peers_Peer
// None
type ArpRedundancy_Redundancy_Groups_Group_Peers_Peer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IPv4 address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrefixString interface{}
}

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetFilter() yfilter.YFilter { return peer.YFilter }

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) SetFilter(yf yfilter.YFilter) { peer.YFilter = yf }

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetGoName(yname string) string {
    if yname == "prefix-string" { return "PrefixString" }
    return ""
}

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetSegmentPath() string {
    return "peer" + "[prefix-string='" + fmt.Sprintf("%v", peer.PrefixString) + "']"
}

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-string"] = peer.PrefixString
    return leafs
}

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetBundleName() string { return "cisco_ios_xr" }

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetYangName() string { return "peer" }

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) SetParent(parent types.Entity) { peer.parent = parent }

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetParent() types.Entity { return peer.parent }

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetParentYangName() string { return "peers" }

// ArpRedundancy_Redundancy_Groups_Group_InterfaceList
// List of Interfaces for this Group
// This type is a presence type.
type ArpRedundancy_Redundancy_Groups_Group_InterfaceList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable List of Interfaces for this Group. Deletion of this object also
    // causes deletion of all associated objects under InterfaceList. The type is
    // interface{}. This attribute is mandatory.
    Enable interface{}

    // Table of Interface.
    Interfaces ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces
}

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetFilter() yfilter.YFilter { return interfaceList.YFilter }

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) SetFilter(yf yfilter.YFilter) { interfaceList.YFilter = yf }

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetSegmentPath() string {
    return "interface-list"
}

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &interfaceList.Interfaces
    }
    return nil
}

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &interfaceList.Interfaces
    return children
}

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = interfaceList.Enable
    return leafs
}

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetYangName() string { return "interface-list" }

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) SetParent(parent types.Entity) { interfaceList.parent = parent }

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetParent() types.Entity { return interfaceList.parent }

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetParentYangName() string { return "group" }

// ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces
// Table of Interface
type ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface for this Group. The type is slice of
    // ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface.
    Interface []ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface
}

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetParentYangName() string { return "interface-list" }

// ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface
// Interface for this Group
type ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface Id for the interface. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    InterfaceId interface{}
}

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-id" { return "InterfaceId" }
    return ""
}

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-id"] = self.InterfaceId
    return leafs
}

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

