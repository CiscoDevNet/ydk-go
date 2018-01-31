// This module contains a collection of YANG definitions
// for Cisco IOS-XR lpts-pa package operational data.
// 
// This module contains definitions
// for the following management objects:
//   lpts-pa: lpts pre-ifib data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package lpts_pa_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lpts_pa_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lpts-pa-oper lpts-pa}", reflect.TypeOf(LptsPa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lpts-pa-oper:lpts-pa", reflect.TypeOf(LptsPa{}))
}

// LptsPa
// lpts pre-ifib data
type LptsPa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lpts pa bindings.
    EntryXr LptsPa_EntryXr

    // lpts pa clients.
    Entries LptsPa_Entries
}

func (lptsPa *LptsPa) GetFilter() yfilter.YFilter { return lptsPa.YFilter }

func (lptsPa *LptsPa) SetFilter(yf yfilter.YFilter) { lptsPa.YFilter = yf }

func (lptsPa *LptsPa) GetGoName(yname string) string {
    if yname == "entry-xr" { return "EntryXr" }
    if yname == "entries" { return "Entries" }
    return ""
}

func (lptsPa *LptsPa) GetSegmentPath() string {
    return "Cisco-IOS-XR-lpts-pa-oper:lpts-pa"
}

func (lptsPa *LptsPa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entry-xr" {
        return &lptsPa.EntryXr
    }
    if childYangName == "entries" {
        return &lptsPa.Entries
    }
    return nil
}

func (lptsPa *LptsPa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["entry-xr"] = &lptsPa.EntryXr
    children["entries"] = &lptsPa.Entries
    return children
}

func (lptsPa *LptsPa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lptsPa *LptsPa) GetBundleName() string { return "cisco_ios_xr" }

func (lptsPa *LptsPa) GetYangName() string { return "lpts-pa" }

func (lptsPa *LptsPa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lptsPa *LptsPa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lptsPa *LptsPa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lptsPa *LptsPa) SetParent(parent types.Entity) { lptsPa.parent = parent }

func (lptsPa *LptsPa) GetParent() types.Entity { return lptsPa.parent }

func (lptsPa *LptsPa) GetParentYangName() string { return "Cisco-IOS-XR-lpts-pa-oper" }

// LptsPa_EntryXr
// lpts pa bindings
type LptsPa_EntryXr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Data for single PA Binding. The type is slice of LptsPa_EntryXr_Entry.
    Entry []LptsPa_EntryXr_Entry
}

func (entryXr *LptsPa_EntryXr) GetFilter() yfilter.YFilter { return entryXr.YFilter }

func (entryXr *LptsPa_EntryXr) SetFilter(yf yfilter.YFilter) { entryXr.YFilter = yf }

func (entryXr *LptsPa_EntryXr) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (entryXr *LptsPa_EntryXr) GetSegmentPath() string {
    return "entry-xr"
}

func (entryXr *LptsPa_EntryXr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entry" {
        for _, c := range entryXr.Entry {
            if entryXr.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsPa_EntryXr_Entry{}
        entryXr.Entry = append(entryXr.Entry, child)
        return &entryXr.Entry[len(entryXr.Entry)-1]
    }
    return nil
}

func (entryXr *LptsPa_EntryXr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entryXr.Entry {
        children[entryXr.Entry[i].GetSegmentPath()] = &entryXr.Entry[i]
    }
    return children
}

func (entryXr *LptsPa_EntryXr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entryXr *LptsPa_EntryXr) GetBundleName() string { return "cisco_ios_xr" }

func (entryXr *LptsPa_EntryXr) GetYangName() string { return "entry-xr" }

func (entryXr *LptsPa_EntryXr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entryXr *LptsPa_EntryXr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entryXr *LptsPa_EntryXr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entryXr *LptsPa_EntryXr) SetParent(parent types.Entity) { entryXr.parent = parent }

func (entryXr *LptsPa_EntryXr) GetParent() types.Entity { return entryXr.parent }

func (entryXr *LptsPa_EntryXr) GetParentYangName() string { return "lpts-pa" }

// LptsPa_EntryXr_Entry
// Data for single PA Binding
type LptsPa_EntryXr_Entry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Single Binding entry. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Entry interface{}

    // Rack/slot/instance. The type is interface{} with range: 0..4294967295.
    Location interface{}

    // Client ID. The type is interface{} with range: 0..4294967295.
    ClientId interface{}

    // VR/VRF ID. The type is interface{} with range: 0..4294967295.
    Vid interface{}

    // Cookie. The type is interface{} with range: 0..4294967295.
    Cookie interface{}

    // Layer 3 protocol. The type is interface{} with range: 0..4294967295.
    L3Protocol interface{}

    // Layer 4 protocol. The type is interface{} with range: 0..4294967295.
    L4Protocol interface{}

    // Filter operation. The type is interface{} with range: 0..4294967295.
    Smask interface{}

    // Ifhandle. The type is interface{} with range: 0..4294967295.
    Ifs interface{}

    // Packet type. The type is interface{} with range: 0..4294967295.
    Ptype interface{}

    // Local address. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    LocalIp interface{}

    // Remote address. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    RemoteIp interface{}

    // Local address length. The type is interface{} with range: 0..255.
    LocalLen interface{}

    // Remote address length. The type is interface{} with range: 0..255.
    RemoteLen interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Remote port. The type is interface{} with range: 0..65535.
    RemotePort interface{}

    // L5 info. The type is interface{} with range: 0..4294967295.
    PacketMisc interface{}

    // Scope. The type is interface{} with range: 0..4294967295.
    Scope interface{}

    // Client flags. The type is interface{} with range: 0..4294967295.
    ClientFlags interface{}

    // Minimum TTL. The type is interface{} with range: 0..255.
    MinTtl interface{}

    // lazy binding queue delay. The type is interface{} with range:
    // 0..4294967295.
    LazyBindqDelay interface{}

    // pending transactions queue delay. The type is interface{} with range:
    // 0..4294967295.
    PtqDelay interface{}

    // Creation Time.
    Ctime LptsPa_EntryXr_Entry_Ctime

    // Update Time.
    Utime LptsPa_EntryXr_Entry_Utime
}

func (entry *LptsPa_EntryXr_Entry) GetFilter() yfilter.YFilter { return entry.YFilter }

func (entry *LptsPa_EntryXr_Entry) SetFilter(yf yfilter.YFilter) { entry.YFilter = yf }

func (entry *LptsPa_EntryXr_Entry) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    if yname == "location" { return "Location" }
    if yname == "client-id" { return "ClientId" }
    if yname == "vid" { return "Vid" }
    if yname == "cookie" { return "Cookie" }
    if yname == "l3protocol" { return "L3Protocol" }
    if yname == "l4protocol" { return "L4Protocol" }
    if yname == "smask" { return "Smask" }
    if yname == "ifs" { return "Ifs" }
    if yname == "ptype" { return "Ptype" }
    if yname == "local-ip" { return "LocalIp" }
    if yname == "remote-ip" { return "RemoteIp" }
    if yname == "local-len" { return "LocalLen" }
    if yname == "remote-len" { return "RemoteLen" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "remote-port" { return "RemotePort" }
    if yname == "packet-misc" { return "PacketMisc" }
    if yname == "scope" { return "Scope" }
    if yname == "client-flags" { return "ClientFlags" }
    if yname == "min-ttl" { return "MinTtl" }
    if yname == "lazy-bindq-delay" { return "LazyBindqDelay" }
    if yname == "ptq-delay" { return "PtqDelay" }
    if yname == "ctime" { return "Ctime" }
    if yname == "utime" { return "Utime" }
    return ""
}

func (entry *LptsPa_EntryXr_Entry) GetSegmentPath() string {
    return "entry" + "[entry='" + fmt.Sprintf("%v", entry.Entry) + "']"
}

func (entry *LptsPa_EntryXr_Entry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ctime" {
        return &entry.Ctime
    }
    if childYangName == "utime" {
        return &entry.Utime
    }
    return nil
}

func (entry *LptsPa_EntryXr_Entry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ctime"] = &entry.Ctime
    children["utime"] = &entry.Utime
    return children
}

func (entry *LptsPa_EntryXr_Entry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = entry.Entry
    leafs["location"] = entry.Location
    leafs["client-id"] = entry.ClientId
    leafs["vid"] = entry.Vid
    leafs["cookie"] = entry.Cookie
    leafs["l3protocol"] = entry.L3Protocol
    leafs["l4protocol"] = entry.L4Protocol
    leafs["smask"] = entry.Smask
    leafs["ifs"] = entry.Ifs
    leafs["ptype"] = entry.Ptype
    leafs["local-ip"] = entry.LocalIp
    leafs["remote-ip"] = entry.RemoteIp
    leafs["local-len"] = entry.LocalLen
    leafs["remote-len"] = entry.RemoteLen
    leafs["local-port"] = entry.LocalPort
    leafs["remote-port"] = entry.RemotePort
    leafs["packet-misc"] = entry.PacketMisc
    leafs["scope"] = entry.Scope
    leafs["client-flags"] = entry.ClientFlags
    leafs["min-ttl"] = entry.MinTtl
    leafs["lazy-bindq-delay"] = entry.LazyBindqDelay
    leafs["ptq-delay"] = entry.PtqDelay
    return leafs
}

func (entry *LptsPa_EntryXr_Entry) GetBundleName() string { return "cisco_ios_xr" }

func (entry *LptsPa_EntryXr_Entry) GetYangName() string { return "entry" }

func (entry *LptsPa_EntryXr_Entry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entry *LptsPa_EntryXr_Entry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entry *LptsPa_EntryXr_Entry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entry *LptsPa_EntryXr_Entry) SetParent(parent types.Entity) { entry.parent = parent }

func (entry *LptsPa_EntryXr_Entry) GetParent() types.Entity { return entry.parent }

func (entry *LptsPa_EntryXr_Entry) GetParentYangName() string { return "entry-xr" }

// LptsPa_EntryXr_Entry_Ctime
// Creation Time
type LptsPa_EntryXr_Entry_Ctime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Sec. The type is interface{} with range: 0..4294967295.
    TvSec interface{}

    // Time Nanosec. The type is interface{} with range: 0..4294967295.
    TvNsec interface{}
}

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetFilter() yfilter.YFilter { return ctime.YFilter }

func (ctime *LptsPa_EntryXr_Entry_Ctime) SetFilter(yf yfilter.YFilter) { ctime.YFilter = yf }

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetGoName(yname string) string {
    if yname == "tv-sec" { return "TvSec" }
    if yname == "tv-nsec" { return "TvNsec" }
    return ""
}

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetSegmentPath() string {
    return "ctime"
}

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tv-sec"] = ctime.TvSec
    leafs["tv-nsec"] = ctime.TvNsec
    return leafs
}

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetBundleName() string { return "cisco_ios_xr" }

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetYangName() string { return "ctime" }

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ctime *LptsPa_EntryXr_Entry_Ctime) SetParent(parent types.Entity) { ctime.parent = parent }

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetParent() types.Entity { return ctime.parent }

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetParentYangName() string { return "entry" }

// LptsPa_EntryXr_Entry_Utime
// Update Time
type LptsPa_EntryXr_Entry_Utime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Sec. The type is interface{} with range: 0..4294967295.
    TvSec interface{}

    // Time Nanosec. The type is interface{} with range: 0..4294967295.
    TvNsec interface{}
}

func (utime *LptsPa_EntryXr_Entry_Utime) GetFilter() yfilter.YFilter { return utime.YFilter }

func (utime *LptsPa_EntryXr_Entry_Utime) SetFilter(yf yfilter.YFilter) { utime.YFilter = yf }

func (utime *LptsPa_EntryXr_Entry_Utime) GetGoName(yname string) string {
    if yname == "tv-sec" { return "TvSec" }
    if yname == "tv-nsec" { return "TvNsec" }
    return ""
}

func (utime *LptsPa_EntryXr_Entry_Utime) GetSegmentPath() string {
    return "utime"
}

func (utime *LptsPa_EntryXr_Entry_Utime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (utime *LptsPa_EntryXr_Entry_Utime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (utime *LptsPa_EntryXr_Entry_Utime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tv-sec"] = utime.TvSec
    leafs["tv-nsec"] = utime.TvNsec
    return leafs
}

func (utime *LptsPa_EntryXr_Entry_Utime) GetBundleName() string { return "cisco_ios_xr" }

func (utime *LptsPa_EntryXr_Entry_Utime) GetYangName() string { return "utime" }

func (utime *LptsPa_EntryXr_Entry_Utime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (utime *LptsPa_EntryXr_Entry_Utime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (utime *LptsPa_EntryXr_Entry_Utime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (utime *LptsPa_EntryXr_Entry_Utime) SetParent(parent types.Entity) { utime.parent = parent }

func (utime *LptsPa_EntryXr_Entry_Utime) GetParent() types.Entity { return utime.parent }

func (utime *LptsPa_EntryXr_Entry_Utime) GetParentYangName() string { return "entry" }

// LptsPa_Entries
// lpts pa clients
type LptsPa_Entries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Data for single PA Client. The type is slice of LptsPa_Entries_Entry.
    Entry []LptsPa_Entries_Entry
}

func (entries *LptsPa_Entries) GetFilter() yfilter.YFilter { return entries.YFilter }

func (entries *LptsPa_Entries) SetFilter(yf yfilter.YFilter) { entries.YFilter = yf }

func (entries *LptsPa_Entries) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (entries *LptsPa_Entries) GetSegmentPath() string {
    return "entries"
}

func (entries *LptsPa_Entries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entry" {
        for _, c := range entries.Entry {
            if entries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LptsPa_Entries_Entry{}
        entries.Entry = append(entries.Entry, child)
        return &entries.Entry[len(entries.Entry)-1]
    }
    return nil
}

func (entries *LptsPa_Entries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entries.Entry {
        children[entries.Entry[i].GetSegmentPath()] = &entries.Entry[i]
    }
    return children
}

func (entries *LptsPa_Entries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entries *LptsPa_Entries) GetBundleName() string { return "cisco_ios_xr" }

func (entries *LptsPa_Entries) GetYangName() string { return "entries" }

func (entries *LptsPa_Entries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entries *LptsPa_Entries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entries *LptsPa_Entries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entries *LptsPa_Entries) SetParent(parent types.Entity) { entries.parent = parent }

func (entries *LptsPa_Entries) GetParent() types.Entity { return entries.parent }

func (entries *LptsPa_Entries) GetParentYangName() string { return "lpts-pa" }

// LptsPa_Entries_Entry
// Data for single PA Client
type LptsPa_Entries_Entry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Single Client entry. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Entry interface{}

    // Client flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // Open flags. The type is interface{} with range: 0..4294967295.
    OpenFlags interface{}

    // Rack/slot/instance. The type is interface{} with range: 0..4294967295.
    Location interface{}

    // Client ID. The type is interface{} with range: 0..4294967295.
    ClientId interface{}

    // Transaction statisitics. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Times interface{}
}

func (entry *LptsPa_Entries_Entry) GetFilter() yfilter.YFilter { return entry.YFilter }

func (entry *LptsPa_Entries_Entry) SetFilter(yf yfilter.YFilter) { entry.YFilter = yf }

func (entry *LptsPa_Entries_Entry) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    if yname == "flags" { return "Flags" }
    if yname == "open-flags" { return "OpenFlags" }
    if yname == "location" { return "Location" }
    if yname == "client-id" { return "ClientId" }
    if yname == "times" { return "Times" }
    return ""
}

func (entry *LptsPa_Entries_Entry) GetSegmentPath() string {
    return "entry" + "[entry='" + fmt.Sprintf("%v", entry.Entry) + "']"
}

func (entry *LptsPa_Entries_Entry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entry *LptsPa_Entries_Entry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entry *LptsPa_Entries_Entry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = entry.Entry
    leafs["flags"] = entry.Flags
    leafs["open-flags"] = entry.OpenFlags
    leafs["location"] = entry.Location
    leafs["client-id"] = entry.ClientId
    leafs["times"] = entry.Times
    return leafs
}

func (entry *LptsPa_Entries_Entry) GetBundleName() string { return "cisco_ios_xr" }

func (entry *LptsPa_Entries_Entry) GetYangName() string { return "entry" }

func (entry *LptsPa_Entries_Entry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entry *LptsPa_Entries_Entry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entry *LptsPa_Entries_Entry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entry *LptsPa_Entries_Entry) SetParent(parent types.Entity) { entry.parent = parent }

func (entry *LptsPa_Entries_Entry) GetParent() types.Entity { return entry.parent }

func (entry *LptsPa_Entries_Entry) GetParentYangName() string { return "entries" }

