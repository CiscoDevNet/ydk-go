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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // lpts pa bindings.
    EntryXr LptsPa_EntryXr

    // lpts pa clients.
    Entries LptsPa_Entries
}

func (lptsPa *LptsPa) GetEntityData() *types.CommonEntityData {
    lptsPa.EntityData.YFilter = lptsPa.YFilter
    lptsPa.EntityData.YangName = "lpts-pa"
    lptsPa.EntityData.BundleName = "cisco_ios_xr"
    lptsPa.EntityData.ParentYangName = "Cisco-IOS-XR-lpts-pa-oper"
    lptsPa.EntityData.SegmentPath = "Cisco-IOS-XR-lpts-pa-oper:lpts-pa"
    lptsPa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lptsPa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lptsPa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lptsPa.EntityData.Children = make(map[string]types.YChild)
    lptsPa.EntityData.Children["entry-xr"] = types.YChild{"EntryXr", &lptsPa.EntryXr}
    lptsPa.EntityData.Children["entries"] = types.YChild{"Entries", &lptsPa.Entries}
    lptsPa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lptsPa.EntityData)
}

// LptsPa_EntryXr
// lpts pa bindings
type LptsPa_EntryXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data for single PA Binding. The type is slice of LptsPa_EntryXr_Entry.
    Entry []LptsPa_EntryXr_Entry
}

func (entryXr *LptsPa_EntryXr) GetEntityData() *types.CommonEntityData {
    entryXr.EntityData.YFilter = entryXr.YFilter
    entryXr.EntityData.YangName = "entry-xr"
    entryXr.EntityData.BundleName = "cisco_ios_xr"
    entryXr.EntityData.ParentYangName = "lpts-pa"
    entryXr.EntityData.SegmentPath = "entry-xr"
    entryXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entryXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entryXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entryXr.EntityData.Children = make(map[string]types.YChild)
    entryXr.EntityData.Children["entry"] = types.YChild{"Entry", nil}
    for i := range entryXr.Entry {
        entryXr.EntityData.Children[types.GetSegmentPath(&entryXr.Entry[i])] = types.YChild{"Entry", &entryXr.Entry[i]}
    }
    entryXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entryXr.EntityData)
}

// LptsPa_EntryXr_Entry
// Data for single PA Binding
type LptsPa_EntryXr_Entry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Single Binding entry. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    LocalIp interface{}

    // Remote address. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
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

func (entry *LptsPa_EntryXr_Entry) GetEntityData() *types.CommonEntityData {
    entry.EntityData.YFilter = entry.YFilter
    entry.EntityData.YangName = "entry"
    entry.EntityData.BundleName = "cisco_ios_xr"
    entry.EntityData.ParentYangName = "entry-xr"
    entry.EntityData.SegmentPath = "entry" + "[entry='" + fmt.Sprintf("%v", entry.Entry) + "']"
    entry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entry.EntityData.Children = make(map[string]types.YChild)
    entry.EntityData.Children["ctime"] = types.YChild{"Ctime", &entry.Ctime}
    entry.EntityData.Children["utime"] = types.YChild{"Utime", &entry.Utime}
    entry.EntityData.Leafs = make(map[string]types.YLeaf)
    entry.EntityData.Leafs["entry"] = types.YLeaf{"Entry", entry.Entry}
    entry.EntityData.Leafs["location"] = types.YLeaf{"Location", entry.Location}
    entry.EntityData.Leafs["client-id"] = types.YLeaf{"ClientId", entry.ClientId}
    entry.EntityData.Leafs["vid"] = types.YLeaf{"Vid", entry.Vid}
    entry.EntityData.Leafs["cookie"] = types.YLeaf{"Cookie", entry.Cookie}
    entry.EntityData.Leafs["l3protocol"] = types.YLeaf{"L3Protocol", entry.L3Protocol}
    entry.EntityData.Leafs["l4protocol"] = types.YLeaf{"L4Protocol", entry.L4Protocol}
    entry.EntityData.Leafs["smask"] = types.YLeaf{"Smask", entry.Smask}
    entry.EntityData.Leafs["ifs"] = types.YLeaf{"Ifs", entry.Ifs}
    entry.EntityData.Leafs["ptype"] = types.YLeaf{"Ptype", entry.Ptype}
    entry.EntityData.Leafs["local-ip"] = types.YLeaf{"LocalIp", entry.LocalIp}
    entry.EntityData.Leafs["remote-ip"] = types.YLeaf{"RemoteIp", entry.RemoteIp}
    entry.EntityData.Leafs["local-len"] = types.YLeaf{"LocalLen", entry.LocalLen}
    entry.EntityData.Leafs["remote-len"] = types.YLeaf{"RemoteLen", entry.RemoteLen}
    entry.EntityData.Leafs["local-port"] = types.YLeaf{"LocalPort", entry.LocalPort}
    entry.EntityData.Leafs["remote-port"] = types.YLeaf{"RemotePort", entry.RemotePort}
    entry.EntityData.Leafs["packet-misc"] = types.YLeaf{"PacketMisc", entry.PacketMisc}
    entry.EntityData.Leafs["scope"] = types.YLeaf{"Scope", entry.Scope}
    entry.EntityData.Leafs["client-flags"] = types.YLeaf{"ClientFlags", entry.ClientFlags}
    entry.EntityData.Leafs["min-ttl"] = types.YLeaf{"MinTtl", entry.MinTtl}
    entry.EntityData.Leafs["lazy-bindq-delay"] = types.YLeaf{"LazyBindqDelay", entry.LazyBindqDelay}
    entry.EntityData.Leafs["ptq-delay"] = types.YLeaf{"PtqDelay", entry.PtqDelay}
    return &(entry.EntityData)
}

// LptsPa_EntryXr_Entry_Ctime
// Creation Time
type LptsPa_EntryXr_Entry_Ctime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Sec. The type is interface{} with range: 0..4294967295.
    TvSec interface{}

    // Time Nanosec. The type is interface{} with range: 0..4294967295.
    TvNsec interface{}
}

func (ctime *LptsPa_EntryXr_Entry_Ctime) GetEntityData() *types.CommonEntityData {
    ctime.EntityData.YFilter = ctime.YFilter
    ctime.EntityData.YangName = "ctime"
    ctime.EntityData.BundleName = "cisco_ios_xr"
    ctime.EntityData.ParentYangName = "entry"
    ctime.EntityData.SegmentPath = "ctime"
    ctime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ctime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ctime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ctime.EntityData.Children = make(map[string]types.YChild)
    ctime.EntityData.Leafs = make(map[string]types.YLeaf)
    ctime.EntityData.Leafs["tv-sec"] = types.YLeaf{"TvSec", ctime.TvSec}
    ctime.EntityData.Leafs["tv-nsec"] = types.YLeaf{"TvNsec", ctime.TvNsec}
    return &(ctime.EntityData)
}

// LptsPa_EntryXr_Entry_Utime
// Update Time
type LptsPa_EntryXr_Entry_Utime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Sec. The type is interface{} with range: 0..4294967295.
    TvSec interface{}

    // Time Nanosec. The type is interface{} with range: 0..4294967295.
    TvNsec interface{}
}

func (utime *LptsPa_EntryXr_Entry_Utime) GetEntityData() *types.CommonEntityData {
    utime.EntityData.YFilter = utime.YFilter
    utime.EntityData.YangName = "utime"
    utime.EntityData.BundleName = "cisco_ios_xr"
    utime.EntityData.ParentYangName = "entry"
    utime.EntityData.SegmentPath = "utime"
    utime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utime.EntityData.Children = make(map[string]types.YChild)
    utime.EntityData.Leafs = make(map[string]types.YLeaf)
    utime.EntityData.Leafs["tv-sec"] = types.YLeaf{"TvSec", utime.TvSec}
    utime.EntityData.Leafs["tv-nsec"] = types.YLeaf{"TvNsec", utime.TvNsec}
    return &(utime.EntityData)
}

// LptsPa_Entries
// lpts pa clients
type LptsPa_Entries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data for single PA Client. The type is slice of LptsPa_Entries_Entry.
    Entry []LptsPa_Entries_Entry
}

func (entries *LptsPa_Entries) GetEntityData() *types.CommonEntityData {
    entries.EntityData.YFilter = entries.YFilter
    entries.EntityData.YangName = "entries"
    entries.EntityData.BundleName = "cisco_ios_xr"
    entries.EntityData.ParentYangName = "lpts-pa"
    entries.EntityData.SegmentPath = "entries"
    entries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entries.EntityData.Children = make(map[string]types.YChild)
    entries.EntityData.Children["entry"] = types.YChild{"Entry", nil}
    for i := range entries.Entry {
        entries.EntityData.Children[types.GetSegmentPath(&entries.Entry[i])] = types.YChild{"Entry", &entries.Entry[i]}
    }
    entries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entries.EntityData)
}

// LptsPa_Entries_Entry
// Data for single PA Client
type LptsPa_Entries_Entry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Single Client entry. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Times interface{}
}

func (entry *LptsPa_Entries_Entry) GetEntityData() *types.CommonEntityData {
    entry.EntityData.YFilter = entry.YFilter
    entry.EntityData.YangName = "entry"
    entry.EntityData.BundleName = "cisco_ios_xr"
    entry.EntityData.ParentYangName = "entries"
    entry.EntityData.SegmentPath = "entry" + "[entry='" + fmt.Sprintf("%v", entry.Entry) + "']"
    entry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entry.EntityData.Children = make(map[string]types.YChild)
    entry.EntityData.Leafs = make(map[string]types.YLeaf)
    entry.EntityData.Leafs["entry"] = types.YLeaf{"Entry", entry.Entry}
    entry.EntityData.Leafs["flags"] = types.YLeaf{"Flags", entry.Flags}
    entry.EntityData.Leafs["open-flags"] = types.YLeaf{"OpenFlags", entry.OpenFlags}
    entry.EntityData.Leafs["location"] = types.YLeaf{"Location", entry.Location}
    entry.EntityData.Leafs["client-id"] = types.YLeaf{"ClientId", entry.ClientId}
    entry.EntityData.Leafs["times"] = types.YLeaf{"Times", entry.Times}
    return &(entry.EntityData)
}

