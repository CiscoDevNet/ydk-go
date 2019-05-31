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
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
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

func (arp *Arp) GetEntityData() *types.CommonEntityData {
    arp.EntityData.YFilter = arp.YFilter
    arp.EntityData.YangName = "arp"
    arp.EntityData.BundleName = "cisco_ios_xr"
    arp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-arp-cfg"
    arp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-arp-cfg:arp"
    arp.EntityData.AbsolutePath = arp.EntityData.SegmentPath
    arp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    arp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    arp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    arp.EntityData.Children = types.NewOrderedMap()
    arp.EntityData.Leafs = types.NewOrderedMap()
    arp.EntityData.Leafs.Append("max-entries", types.YLeaf{"MaxEntries", arp.MaxEntries})
    arp.EntityData.Leafs.Append("inner-cos", types.YLeaf{"InnerCos", arp.InnerCos})
    arp.EntityData.Leafs.Append("outer-cos", types.YLeaf{"OuterCos", arp.OuterCos})

    arp.EntityData.YListKeys = []string {}

    return &(arp.EntityData)
}

// IedgeCfg
// iedge cfg
type IedgeCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ARP Subscriber Enable Unconditional Proxy ARP. The type is interface{}.
    SubscriberUncondProxy interface{}

    // ARP Subscriber Scale Mode Configuration. The type is interface{}.
    SubscriberScaleMode interface{}
}

func (iedgeCfg *IedgeCfg) GetEntityData() *types.CommonEntityData {
    iedgeCfg.EntityData.YFilter = iedgeCfg.YFilter
    iedgeCfg.EntityData.YangName = "iedge-cfg"
    iedgeCfg.EntityData.BundleName = "cisco_ios_xr"
    iedgeCfg.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-arp-cfg"
    iedgeCfg.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-arp-cfg:iedge-cfg"
    iedgeCfg.EntityData.AbsolutePath = iedgeCfg.EntityData.SegmentPath
    iedgeCfg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iedgeCfg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iedgeCfg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iedgeCfg.EntityData.Children = types.NewOrderedMap()
    iedgeCfg.EntityData.Leafs = types.NewOrderedMap()
    iedgeCfg.EntityData.Leafs.Append("subscriber-uncond-proxy", types.YLeaf{"SubscriberUncondProxy", iedgeCfg.SubscriberUncondProxy})
    iedgeCfg.EntityData.Leafs.Append("subscriber-scale-mode", types.YLeaf{"SubscriberScaleMode", iedgeCfg.SubscriberScaleMode})

    iedgeCfg.EntityData.YListKeys = []string {}

    return &(iedgeCfg.EntityData)
}

// Arpgmp
// arpgmp
type Arpgmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per VRF configuration, for the default VRF use 'default'. The type is slice
    // of Arpgmp_Vrf.
    Vrf []*Arpgmp_Vrf
}

func (arpgmp *Arpgmp) GetEntityData() *types.CommonEntityData {
    arpgmp.EntityData.YFilter = arpgmp.YFilter
    arpgmp.EntityData.YangName = "arpgmp"
    arpgmp.EntityData.BundleName = "cisco_ios_xr"
    arpgmp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-arp-cfg"
    arpgmp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-arp-cfg:arpgmp"
    arpgmp.EntityData.AbsolutePath = arpgmp.EntityData.SegmentPath
    arpgmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    arpgmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    arpgmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    arpgmp.EntityData.Children = types.NewOrderedMap()
    arpgmp.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range arpgmp.Vrf {
        arpgmp.EntityData.Children.Append(types.GetSegmentPath(arpgmp.Vrf[i]), types.YChild{"Vrf", arpgmp.Vrf[i]})
    }
    arpgmp.EntityData.Leafs = types.NewOrderedMap()

    arpgmp.EntityData.YListKeys = []string {}

    return &(arpgmp.EntityData)
}

// Arpgmp_Vrf
// Per VRF configuration, for the default VRF use
// 'default'
type Arpgmp_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // ARP static and alias entry configuration.
    Entries Arpgmp_Vrf_Entries
}

func (vrf *Arpgmp_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "arpgmp"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-arp-cfg:arpgmp/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("entries", types.YChild{"Entries", &vrf.Entries})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Arpgmp_Vrf_Entries
// ARP static and alias entry configuration
type Arpgmp_Vrf_Entries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ARP static and alias entry configuration item. The type is slice of
    // Arpgmp_Vrf_Entries_Entry.
    Entry []*Arpgmp_Vrf_Entries_Entry
}

func (entries *Arpgmp_Vrf_Entries) GetEntityData() *types.CommonEntityData {
    entries.EntityData.YFilter = entries.YFilter
    entries.EntityData.YangName = "entries"
    entries.EntityData.BundleName = "cisco_ios_xr"
    entries.EntityData.ParentYangName = "vrf"
    entries.EntityData.SegmentPath = "entries"
    entries.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-arp-cfg:arpgmp/vrf/" + entries.EntityData.SegmentPath
    entries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entries.EntityData.Children = types.NewOrderedMap()
    entries.EntityData.Children.Append("entry", types.YChild{"Entry", nil})
    for i := range entries.Entry {
        entries.EntityData.Children.Append(types.GetSegmentPath(entries.Entry[i]), types.YChild{"Entry", entries.Entry[i]})
    }
    entries.EntityData.Leafs = types.NewOrderedMap()

    entries.EntityData.YListKeys = []string {}

    return &(entries.EntityData)
}

// Arpgmp_Vrf_Entries_Entry
// ARP static and alias entry configuration item
type Arpgmp_Vrf_Entries_Entry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IP Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Encapsulation type. The type is ArpEncap.
    Encapsulation interface{}

    // Entry type. The type is ArpEntry.
    EntryType interface{}

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    Interface interface{}
}

func (entry *Arpgmp_Vrf_Entries_Entry) GetEntityData() *types.CommonEntityData {
    entry.EntityData.YFilter = entry.YFilter
    entry.EntityData.YangName = "entry"
    entry.EntityData.BundleName = "cisco_ios_xr"
    entry.EntityData.ParentYangName = "entries"
    entry.EntityData.SegmentPath = "entry" + types.AddKeyToken(entry.Address, "address")
    entry.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-arp-cfg:arpgmp/vrf/entries/" + entry.EntityData.SegmentPath
    entry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entry.EntityData.Children = types.NewOrderedMap()
    entry.EntityData.Leafs = types.NewOrderedMap()
    entry.EntityData.Leafs.Append("address", types.YLeaf{"Address", entry.Address})
    entry.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", entry.MacAddress})
    entry.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", entry.Encapsulation})
    entry.EntityData.Leafs.Append("entry-type", types.YLeaf{"EntryType", entry.EntryType})
    entry.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", entry.Interface})

    entry.EntityData.YListKeys = []string {"Address"}

    return &(entry.EntityData)
}

// ArpRedundancy
// arp redundancy
type ArpRedundancy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure parameter for ARP Geo redundancy.
    Redundancy ArpRedundancy_Redundancy
}

func (arpRedundancy *ArpRedundancy) GetEntityData() *types.CommonEntityData {
    arpRedundancy.EntityData.YFilter = arpRedundancy.YFilter
    arpRedundancy.EntityData.YangName = "arp-redundancy"
    arpRedundancy.EntityData.BundleName = "cisco_ios_xr"
    arpRedundancy.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-arp-cfg"
    arpRedundancy.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-arp-cfg:arp-redundancy"
    arpRedundancy.EntityData.AbsolutePath = arpRedundancy.EntityData.SegmentPath
    arpRedundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    arpRedundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    arpRedundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    arpRedundancy.EntityData.Children = types.NewOrderedMap()
    arpRedundancy.EntityData.Children.Append("redundancy", types.YChild{"Redundancy", &arpRedundancy.Redundancy})
    arpRedundancy.EntityData.Leafs = types.NewOrderedMap()

    arpRedundancy.EntityData.YListKeys = []string {}

    return &(arpRedundancy.EntityData)
}

// ArpRedundancy_Redundancy
// Configure parameter for ARP Geo redundancy
// This type is a presence type.
type ArpRedundancy_Redundancy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable Configure parameter for ARP Geo redundancy. Deletion of this object
    // also causes deletion of all associated objects under ArpRedundancy. The
    // type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Table of Group.
    Groups ArpRedundancy_Redundancy_Groups
}

func (redundancy *ArpRedundancy_Redundancy) GetEntityData() *types.CommonEntityData {
    redundancy.EntityData.YFilter = redundancy.YFilter
    redundancy.EntityData.YangName = "redundancy"
    redundancy.EntityData.BundleName = "cisco_ios_xr"
    redundancy.EntityData.ParentYangName = "arp-redundancy"
    redundancy.EntityData.SegmentPath = "redundancy"
    redundancy.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-arp-cfg:arp-redundancy/" + redundancy.EntityData.SegmentPath
    redundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redundancy.EntityData.Children = types.NewOrderedMap()
    redundancy.EntityData.Children.Append("groups", types.YChild{"Groups", &redundancy.Groups})
    redundancy.EntityData.Leafs = types.NewOrderedMap()
    redundancy.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", redundancy.Enable})

    redundancy.EntityData.YListKeys = []string {}

    return &(redundancy.EntityData)
}

// ArpRedundancy_Redundancy_Groups
// Table of Group
type ArpRedundancy_Redundancy_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of ArpRedundancy_Redundancy_Groups_Group.
    Group []*ArpRedundancy_Redundancy_Groups_Group
}

func (groups *ArpRedundancy_Redundancy_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "redundancy"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-arp-cfg:arp-redundancy/redundancy/" + groups.EntityData.SegmentPath
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// ArpRedundancy_Redundancy_Groups_Group
// None
type ArpRedundancy_Redundancy_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Group ID. The type is interface{} with range:
    // 1..32.
    GroupId interface{}

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    SourceInterface interface{}

    // Table of Peer.
    Peers ArpRedundancy_Redundancy_Groups_Group_Peers

    // List of Interfaces for this Group.
    InterfaceList ArpRedundancy_Redundancy_Groups_Group_InterfaceList
}

func (group *ArpRedundancy_Redundancy_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.GroupId, "group-id")
    group.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-arp-cfg:arp-redundancy/redundancy/groups/" + group.EntityData.SegmentPath
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Children.Append("peers", types.YChild{"Peers", &group.Peers})
    group.EntityData.Children.Append("interface-list", types.YChild{"InterfaceList", &group.InterfaceList})
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", group.GroupId})
    group.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", group.SourceInterface})

    group.EntityData.YListKeys = []string {"GroupId"}

    return &(group.EntityData)
}

// ArpRedundancy_Redundancy_Groups_Group_Peers
// Table of Peer
type ArpRedundancy_Redundancy_Groups_Group_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of
    // ArpRedundancy_Redundancy_Groups_Group_Peers_Peer.
    Peer []*ArpRedundancy_Redundancy_Groups_Group_Peers_Peer
}

func (peers *ArpRedundancy_Redundancy_Groups_Group_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "group"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-arp-cfg:arp-redundancy/redundancy/groups/group/" + peers.EntityData.SegmentPath
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// ArpRedundancy_Redundancy_Groups_Group_Peers_Peer
// None
type ArpRedundancy_Redundancy_Groups_Group_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor IPv4 address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PrefixString interface{}
}

func (peer *ArpRedundancy_Redundancy_Groups_Group_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.PrefixString, "prefix-string")
    peer.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-arp-cfg:arp-redundancy/redundancy/groups/group/peers/" + peer.EntityData.SegmentPath
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("prefix-string", types.YLeaf{"PrefixString", peer.PrefixString})

    peer.EntityData.YListKeys = []string {"PrefixString"}

    return &(peer.EntityData)
}

// ArpRedundancy_Redundancy_Groups_Group_InterfaceList
// List of Interfaces for this Group
// This type is a presence type.
type ArpRedundancy_Redundancy_Groups_Group_InterfaceList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable List of Interfaces for this Group. Deletion of this object also
    // causes deletion of all associated objects under InterfaceList. The type is
    // interface{}. This attribute is mandatory.
    Enable interface{}

    // Table of Interface.
    Interfaces ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces
}

func (interfaceList *ArpRedundancy_Redundancy_Groups_Group_InterfaceList) GetEntityData() *types.CommonEntityData {
    interfaceList.EntityData.YFilter = interfaceList.YFilter
    interfaceList.EntityData.YangName = "interface-list"
    interfaceList.EntityData.BundleName = "cisco_ios_xr"
    interfaceList.EntityData.ParentYangName = "group"
    interfaceList.EntityData.SegmentPath = "interface-list"
    interfaceList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-arp-cfg:arp-redundancy/redundancy/groups/group/" + interfaceList.EntityData.SegmentPath
    interfaceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceList.EntityData.Children = types.NewOrderedMap()
    interfaceList.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &interfaceList.Interfaces})
    interfaceList.EntityData.Leafs = types.NewOrderedMap()
    interfaceList.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", interfaceList.Enable})

    interfaceList.EntityData.YListKeys = []string {}

    return &(interfaceList.EntityData)
}

// ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces
// Table of Interface
type ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface for this Group. The type is slice of
    // ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface.
    Interface []*ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface
}

func (interfaces *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "interface-list"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-arp-cfg:arp-redundancy/redundancy/groups/group/interface-list/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface
// Interface for this Group
type ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Interface Id for the interface. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    InterfaceId interface{}
}

func (self *ArpRedundancy_Redundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-arp-cfg:arp-redundancy/redundancy/groups/group/interface-list/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", self.InterfaceId})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

