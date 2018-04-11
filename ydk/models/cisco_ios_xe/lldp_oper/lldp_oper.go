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
    LldpEntry []LldpEntries_LldpEntry
}

func (lldpEntries *LldpEntries) GetEntityData() *types.CommonEntityData {
    lldpEntries.EntityData.YFilter = lldpEntries.YFilter
    lldpEntries.EntityData.YangName = "lldp-entries"
    lldpEntries.EntityData.BundleName = "cisco_ios_xe"
    lldpEntries.EntityData.ParentYangName = "Cisco-IOS-XE-lldp-oper"
    lldpEntries.EntityData.SegmentPath = "Cisco-IOS-XE-lldp-oper:lldp-entries"
    lldpEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpEntries.EntityData.Children = make(map[string]types.YChild)
    lldpEntries.EntityData.Children["lldp-entry"] = types.YChild{"LldpEntry", nil}
    for i := range lldpEntries.LldpEntry {
        lldpEntries.EntityData.Children[types.GetSegmentPath(&lldpEntries.LldpEntry[i])] = types.YChild{"LldpEntry", &lldpEntries.LldpEntry[i]}
    }
    lldpEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lldpEntries.EntityData)
}

// LldpEntries_LldpEntry
// The list of LLDP entries
type LldpEntries_LldpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    lldpEntry.EntityData.SegmentPath = "lldp-entry" + "[device-id='" + fmt.Sprintf("%v", lldpEntry.DeviceId) + "']" + "[local-interface='" + fmt.Sprintf("%v", lldpEntry.LocalInterface) + "']" + "[connecting-interface='" + fmt.Sprintf("%v", lldpEntry.ConnectingInterface) + "']"
    lldpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpEntry.EntityData.Children = make(map[string]types.YChild)
    lldpEntry.EntityData.Children["capabilities"] = types.YChild{"Capabilities", &lldpEntry.Capabilities}
    lldpEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    lldpEntry.EntityData.Leafs["device-id"] = types.YLeaf{"DeviceId", lldpEntry.DeviceId}
    lldpEntry.EntityData.Leafs["local-interface"] = types.YLeaf{"LocalInterface", lldpEntry.LocalInterface}
    lldpEntry.EntityData.Leafs["connecting-interface"] = types.YLeaf{"ConnectingInterface", lldpEntry.ConnectingInterface}
    lldpEntry.EntityData.Leafs["ttl"] = types.YLeaf{"Ttl", lldpEntry.Ttl}
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
    capabilities.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    capabilities.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    capabilities.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    capabilities.EntityData.Children = make(map[string]types.YChild)
    capabilities.EntityData.Leafs = make(map[string]types.YLeaf)
    capabilities.EntityData.Leafs["repeater"] = types.YLeaf{"Repeater", capabilities.Repeater}
    capabilities.EntityData.Leafs["bridge"] = types.YLeaf{"Bridge", capabilities.Bridge}
    capabilities.EntityData.Leafs["access-point"] = types.YLeaf{"AccessPoint", capabilities.AccessPoint}
    capabilities.EntityData.Leafs["router"] = types.YLeaf{"Router", capabilities.Router}
    capabilities.EntityData.Leafs["telephone"] = types.YLeaf{"Telephone", capabilities.Telephone}
    capabilities.EntityData.Leafs["docsis"] = types.YLeaf{"Docsis", capabilities.Docsis}
    capabilities.EntityData.Leafs["station"] = types.YLeaf{"Station", capabilities.Station}
    capabilities.EntityData.Leafs["other"] = types.YLeaf{"Other", capabilities.Other}
    return &(capabilities.EntityData)
}

