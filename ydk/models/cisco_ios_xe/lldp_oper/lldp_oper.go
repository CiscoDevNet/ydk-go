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
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of LLDP entries. The type is slice of LldpEntries_LldpEntry.
    LldpEntry []LldpEntries_LldpEntry
}

func (lldpEntries *LldpEntries) GetFilter() yfilter.YFilter { return lldpEntries.YFilter }

func (lldpEntries *LldpEntries) SetFilter(yf yfilter.YFilter) { lldpEntries.YFilter = yf }

func (lldpEntries *LldpEntries) GetGoName(yname string) string {
    if yname == "lldp-entry" { return "LldpEntry" }
    return ""
}

func (lldpEntries *LldpEntries) GetSegmentPath() string {
    return "Cisco-IOS-XE-lldp-oper:lldp-entries"
}

func (lldpEntries *LldpEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-entry" {
        for _, c := range lldpEntries.LldpEntry {
            if lldpEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LldpEntries_LldpEntry{}
        lldpEntries.LldpEntry = append(lldpEntries.LldpEntry, child)
        return &lldpEntries.LldpEntry[len(lldpEntries.LldpEntry)-1]
    }
    return nil
}

func (lldpEntries *LldpEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lldpEntries.LldpEntry {
        children[lldpEntries.LldpEntry[i].GetSegmentPath()] = &lldpEntries.LldpEntry[i]
    }
    return children
}

func (lldpEntries *LldpEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lldpEntries *LldpEntries) GetBundleName() string { return "cisco_ios_xe" }

func (lldpEntries *LldpEntries) GetYangName() string { return "lldp-entries" }

func (lldpEntries *LldpEntries) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpEntries *LldpEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpEntries *LldpEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpEntries *LldpEntries) SetParent(parent types.Entity) { lldpEntries.parent = parent }

func (lldpEntries *LldpEntries) GetParent() types.Entity { return lldpEntries.parent }

func (lldpEntries *LldpEntries) GetParentYangName() string { return "Cisco-IOS-XE-lldp-oper" }

// LldpEntries_LldpEntry
// The list of LLDP entries
type LldpEntries_LldpEntry struct {
    parent types.Entity
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

func (lldpEntry *LldpEntries_LldpEntry) GetFilter() yfilter.YFilter { return lldpEntry.YFilter }

func (lldpEntry *LldpEntries_LldpEntry) SetFilter(yf yfilter.YFilter) { lldpEntry.YFilter = yf }

func (lldpEntry *LldpEntries_LldpEntry) GetGoName(yname string) string {
    if yname == "device-id" { return "DeviceId" }
    if yname == "local-interface" { return "LocalInterface" }
    if yname == "connecting-interface" { return "ConnectingInterface" }
    if yname == "ttl" { return "Ttl" }
    if yname == "capabilities" { return "Capabilities" }
    return ""
}

func (lldpEntry *LldpEntries_LldpEntry) GetSegmentPath() string {
    return "lldp-entry" + "[device-id='" + fmt.Sprintf("%v", lldpEntry.DeviceId) + "']" + "[local-interface='" + fmt.Sprintf("%v", lldpEntry.LocalInterface) + "']" + "[connecting-interface='" + fmt.Sprintf("%v", lldpEntry.ConnectingInterface) + "']"
}

func (lldpEntry *LldpEntries_LldpEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "capabilities" {
        return &lldpEntry.Capabilities
    }
    return nil
}

func (lldpEntry *LldpEntries_LldpEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["capabilities"] = &lldpEntry.Capabilities
    return children
}

func (lldpEntry *LldpEntries_LldpEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-id"] = lldpEntry.DeviceId
    leafs["local-interface"] = lldpEntry.LocalInterface
    leafs["connecting-interface"] = lldpEntry.ConnectingInterface
    leafs["ttl"] = lldpEntry.Ttl
    return leafs
}

func (lldpEntry *LldpEntries_LldpEntry) GetBundleName() string { return "cisco_ios_xe" }

func (lldpEntry *LldpEntries_LldpEntry) GetYangName() string { return "lldp-entry" }

func (lldpEntry *LldpEntries_LldpEntry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpEntry *LldpEntries_LldpEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpEntry *LldpEntries_LldpEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpEntry *LldpEntries_LldpEntry) SetParent(parent types.Entity) { lldpEntry.parent = parent }

func (lldpEntry *LldpEntries_LldpEntry) GetParent() types.Entity { return lldpEntry.parent }

func (lldpEntry *LldpEntries_LldpEntry) GetParentYangName() string { return "lldp-entries" }

// LldpEntries_LldpEntry_Capabilities
// LLD device capabilities
type LldpEntries_LldpEntry_Capabilities struct {
    parent types.Entity
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

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetFilter() yfilter.YFilter { return capabilities.YFilter }

func (capabilities *LldpEntries_LldpEntry_Capabilities) SetFilter(yf yfilter.YFilter) { capabilities.YFilter = yf }

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetGoName(yname string) string {
    if yname == "repeater" { return "Repeater" }
    if yname == "bridge" { return "Bridge" }
    if yname == "access-point" { return "AccessPoint" }
    if yname == "router" { return "Router" }
    if yname == "telephone" { return "Telephone" }
    if yname == "docsis" { return "Docsis" }
    if yname == "station" { return "Station" }
    if yname == "other" { return "Other" }
    return ""
}

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetSegmentPath() string {
    return "capabilities"
}

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["repeater"] = capabilities.Repeater
    leafs["bridge"] = capabilities.Bridge
    leafs["access-point"] = capabilities.AccessPoint
    leafs["router"] = capabilities.Router
    leafs["telephone"] = capabilities.Telephone
    leafs["docsis"] = capabilities.Docsis
    leafs["station"] = capabilities.Station
    leafs["other"] = capabilities.Other
    return leafs
}

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetBundleName() string { return "cisco_ios_xe" }

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetYangName() string { return "capabilities" }

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (capabilities *LldpEntries_LldpEntry_Capabilities) SetParent(parent types.Entity) { capabilities.parent = parent }

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetParent() types.Entity { return capabilities.parent }

func (capabilities *LldpEntries_LldpEntry_Capabilities) GetParentYangName() string { return "lldp-entry" }

