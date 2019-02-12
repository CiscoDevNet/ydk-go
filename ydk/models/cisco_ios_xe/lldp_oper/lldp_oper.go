// This module contains a collection of YANG definitions
// for monitoring of LLDP state information.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package lldp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lldp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-lldp-oper lldp-entries}", reflect.TypeOf(LldpEntries{}))
    ydk.RegisterEntity("Cisco-IOS-XE-lldp-oper:lldp-entries", reflect.TypeOf(LldpEntries{}))
}

// LldpEntries
// Data nodes for lldp entries
type LldpEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of LLDP entries. The type is slice of LldpEntries_LldpEntry.
    LldpEntry []*LldpEntries_LldpEntry
}

func (lldpEntries *LldpEntries) GetEntityData() *types.CommonEntityData {
    lldpEntries.EntityData.YFilter = lldpEntries.YFilter
    lldpEntries.EntityData.YangName = "lldp-entries"
    lldpEntries.EntityData.BundleName = "cisco_ios_xe"
    lldpEntries.EntityData.ParentYangName = "Cisco-IOS-XE-lldp-oper"
    lldpEntries.EntityData.SegmentPath = "Cisco-IOS-XE-lldp-oper:lldp-entries"
    lldpEntries.EntityData.AbsolutePath = lldpEntries.EntityData.SegmentPath
    lldpEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpEntries.EntityData.Children = types.NewOrderedMap()
    lldpEntries.EntityData.Children.Append("lldp-entry", types.YChild{"LldpEntry", nil})
    for i := range lldpEntries.LldpEntry {
        lldpEntries.EntityData.Children.Append(types.GetSegmentPath(lldpEntries.LldpEntry[i]), types.YChild{"LldpEntry", lldpEntries.LldpEntry[i]})
    }
    lldpEntries.EntityData.Leafs = types.NewOrderedMap()

    lldpEntries.EntityData.YListKeys = []string {}

    return &(lldpEntries.EntityData)
}

// LldpEntries_LldpEntry
// The list of LLDP entries
type LldpEntries_LldpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Device ID of the link. The type is string.
    DeviceId interface{}

    // This attribute is a key. Name of the local interface on the current device.
    // The type is string.
    LocalInterface interface{}

    // This attribute is a key. Name of the connected interface to
    // 'local-interface'. The type is string.
    ConnectingInterface interface{}

    // TTL denoting hold-time of this link entry. The type is interface{} with
    // range: 0..4294967295.
    Ttl interface{}

    // LLD device capabilities.
    Capabilities LldpEntries_LldpEntry_Capabilities
}

func (lldpEntry *LldpEntries_LldpEntry) GetEntityData() *types.CommonEntityData {
    lldpEntry.EntityData.YFilter = lldpEntry.YFilter
    lldpEntry.EntityData.YangName = "lldp-entry"
    lldpEntry.EntityData.BundleName = "cisco_ios_xe"
    lldpEntry.EntityData.ParentYangName = "lldp-entries"
    lldpEntry.EntityData.SegmentPath = "lldp-entry" + types.AddKeyToken(lldpEntry.DeviceId, "device-id") + types.AddKeyToken(lldpEntry.LocalInterface, "local-interface") + types.AddKeyToken(lldpEntry.ConnectingInterface, "connecting-interface")
    lldpEntry.EntityData.AbsolutePath = "Cisco-IOS-XE-lldp-oper:lldp-entries/" + lldpEntry.EntityData.SegmentPath
    lldpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpEntry.EntityData.Children = types.NewOrderedMap()
    lldpEntry.EntityData.Children.Append("capabilities", types.YChild{"Capabilities", &lldpEntry.Capabilities})
    lldpEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpEntry.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", lldpEntry.DeviceId})
    lldpEntry.EntityData.Leafs.Append("local-interface", types.YLeaf{"LocalInterface", lldpEntry.LocalInterface})
    lldpEntry.EntityData.Leafs.Append("connecting-interface", types.YLeaf{"ConnectingInterface", lldpEntry.ConnectingInterface})
    lldpEntry.EntityData.Leafs.Append("ttl", types.YLeaf{"Ttl", lldpEntry.Ttl})

    lldpEntry.EntityData.YListKeys = []string {"DeviceId", "LocalInterface", "ConnectingInterface"}

    return &(lldpEntry.EntityData)
}

// LldpEntries_LldpEntry_Capabilities
// LLD device capabilities
type LldpEntries_LldpEntry_Capabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Repeater. The type is interface{}.
    Repeater interface{}

    // Bridge. The type is interface{}.
    Bridge interface{}

    // Access point. The type is interface{}.
    AccessPoint interface{}

    // Router. The type is interface{}.
    Router interface{}

    // Phone. The type is interface{}.
    Telephone interface{}

    // Docsis. The type is interface{}.
    Docsis interface{}

    // Station. The type is interface{}.
    Station interface{}

    // Other. The type is interface{}.
    Other interface{}
}

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetEntityData() *types.CommonEntityData {
    capabilities.EntityData.YFilter = capabilities.YFilter
    capabilities.EntityData.YangName = "capabilities"
    capabilities.EntityData.BundleName = "cisco_ios_xe"
    capabilities.EntityData.ParentYangName = "lldp-entry"
    capabilities.EntityData.SegmentPath = "capabilities"
    capabilities.EntityData.AbsolutePath = "Cisco-IOS-XE-lldp-oper:lldp-entries/lldp-entry/" + capabilities.EntityData.SegmentPath
    capabilities.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    capabilities.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    capabilities.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    capabilities.EntityData.Children = types.NewOrderedMap()
    capabilities.EntityData.Leafs = types.NewOrderedMap()
    capabilities.EntityData.Leafs.Append("repeater", types.YLeaf{"Repeater", capabilities.Repeater})
    capabilities.EntityData.Leafs.Append("bridge", types.YLeaf{"Bridge", capabilities.Bridge})
    capabilities.EntityData.Leafs.Append("access-point", types.YLeaf{"AccessPoint", capabilities.AccessPoint})
    capabilities.EntityData.Leafs.Append("router", types.YLeaf{"Router", capabilities.Router})
    capabilities.EntityData.Leafs.Append("telephone", types.YLeaf{"Telephone", capabilities.Telephone})
    capabilities.EntityData.Leafs.Append("docsis", types.YLeaf{"Docsis", capabilities.Docsis})
    capabilities.EntityData.Leafs.Append("station", types.YLeaf{"Station", capabilities.Station})
    capabilities.EntityData.Leafs.Append("other", types.YLeaf{"Other", capabilities.Other})

    capabilities.EntityData.YListKeys = []string {}

    return &(capabilities.EntityData)
}

