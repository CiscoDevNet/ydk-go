// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1k-mxp-lldp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   lldp-snoop-data: Information related to LLDP Snoop
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs1k_mxp_lldp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs1k_mxp_lldp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs1k-mxp-lldp-oper lldp-snoop-data}", reflect.TypeOf(LldpSnoopData{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs1k-mxp-lldp-oper:lldp-snoop-data", reflect.TypeOf(LldpSnoopData{}))
}

// LldpL3AddrProtocol represents Lldp l3 addr protocol
type LldpL3AddrProtocol string

const (
    // IPv4
    LldpL3AddrProtocol_ipv4 LldpL3AddrProtocol = "ipv4"

    // IPv6
    LldpL3AddrProtocol_ipv6 LldpL3AddrProtocol = "ipv6"
)

// LldpSnoopData
// Information related to LLDP Snoop
type LldpSnoopData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NCS1K LLDP Neighbor brief info.
    LldpNeighborBrief LldpSnoopData_LldpNeighborBrief

    // Ethernet controller snoop data.
    EthernetControllerNames LldpSnoopData_EthernetControllerNames
}

func (lldpSnoopData *LldpSnoopData) GetFilter() yfilter.YFilter { return lldpSnoopData.YFilter }

func (lldpSnoopData *LldpSnoopData) SetFilter(yf yfilter.YFilter) { lldpSnoopData.YFilter = yf }

func (lldpSnoopData *LldpSnoopData) GetGoName(yname string) string {
    if yname == "lldp-neighbor-brief" { return "LldpNeighborBrief" }
    if yname == "ethernet-controller-names" { return "EthernetControllerNames" }
    return ""
}

func (lldpSnoopData *LldpSnoopData) GetSegmentPath() string {
    return "Cisco-IOS-XR-ncs1k-mxp-lldp-oper:lldp-snoop-data"
}

func (lldpSnoopData *LldpSnoopData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-neighbor-brief" {
        return &lldpSnoopData.LldpNeighborBrief
    }
    if childYangName == "ethernet-controller-names" {
        return &lldpSnoopData.EthernetControllerNames
    }
    return nil
}

func (lldpSnoopData *LldpSnoopData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lldp-neighbor-brief"] = &lldpSnoopData.LldpNeighborBrief
    children["ethernet-controller-names"] = &lldpSnoopData.EthernetControllerNames
    return children
}

func (lldpSnoopData *LldpSnoopData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lldpSnoopData *LldpSnoopData) GetBundleName() string { return "cisco_ios_xr" }

func (lldpSnoopData *LldpSnoopData) GetYangName() string { return "lldp-snoop-data" }

func (lldpSnoopData *LldpSnoopData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpSnoopData *LldpSnoopData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpSnoopData *LldpSnoopData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpSnoopData *LldpSnoopData) SetParent(parent types.Entity) { lldpSnoopData.parent = parent }

func (lldpSnoopData *LldpSnoopData) GetParent() types.Entity { return lldpSnoopData.parent }

func (lldpSnoopData *LldpSnoopData) GetParentYangName() string { return "Cisco-IOS-XR-ncs1k-mxp-lldp-oper" }

// LldpSnoopData_LldpNeighborBrief
// NCS1K LLDP Neighbor brief info
type LldpSnoopData_LldpNeighborBrief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of active neighbors entries. The type is interface{} with range:
    // 0..65535.
    NumberOfEntries interface{}

    // List of LLDP neighbors.
    Neighbours LldpSnoopData_LldpNeighborBrief_Neighbours
}

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) GetFilter() yfilter.YFilter { return lldpNeighborBrief.YFilter }

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) SetFilter(yf yfilter.YFilter) { lldpNeighborBrief.YFilter = yf }

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) GetGoName(yname string) string {
    if yname == "number-of-entries" { return "NumberOfEntries" }
    if yname == "neighbours" { return "Neighbours" }
    return ""
}

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) GetSegmentPath() string {
    return "lldp-neighbor-brief"
}

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbours" {
        return &lldpNeighborBrief.Neighbours
    }
    return nil
}

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbours"] = &lldpNeighborBrief.Neighbours
    return children
}

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number-of-entries"] = lldpNeighborBrief.NumberOfEntries
    return leafs
}

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) GetBundleName() string { return "cisco_ios_xr" }

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) GetYangName() string { return "lldp-neighbor-brief" }

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) SetParent(parent types.Entity) { lldpNeighborBrief.parent = parent }

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) GetParent() types.Entity { return lldpNeighborBrief.parent }

func (lldpNeighborBrief *LldpSnoopData_LldpNeighborBrief) GetParentYangName() string { return "lldp-snoop-data" }

// LldpSnoopData_LldpNeighborBrief_Neighbours
// List of LLDP neighbors
type LldpSnoopData_LldpNeighborBrief_Neighbours struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lldp neighbor brief entry. The type is slice of
    // LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry.
    LldpNeighborBriefEntry []LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry
}

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) GetFilter() yfilter.YFilter { return neighbours.YFilter }

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) SetFilter(yf yfilter.YFilter) { neighbours.YFilter = yf }

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) GetGoName(yname string) string {
    if yname == "lldp-neighbor-brief-entry" { return "LldpNeighborBriefEntry" }
    return ""
}

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) GetSegmentPath() string {
    return "neighbours"
}

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-neighbor-brief-entry" {
        for _, c := range neighbours.LldpNeighborBriefEntry {
            if neighbours.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry{}
        neighbours.LldpNeighborBriefEntry = append(neighbours.LldpNeighborBriefEntry, child)
        return &neighbours.LldpNeighborBriefEntry[len(neighbours.LldpNeighborBriefEntry)-1]
    }
    return nil
}

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbours.LldpNeighborBriefEntry {
        children[neighbours.LldpNeighborBriefEntry[i].GetSegmentPath()] = &neighbours.LldpNeighborBriefEntry[i]
    }
    return children
}

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) GetBundleName() string { return "cisco_ios_xr" }

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) GetYangName() string { return "neighbours" }

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) SetParent(parent types.Entity) { neighbours.parent = parent }

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) GetParent() types.Entity { return neighbours.parent }

func (neighbours *LldpSnoopData_LldpNeighborBrief_Neighbours) GetParentYangName() string { return "lldp-neighbor-brief" }

// LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry
// lldp neighbor brief entry
type LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Chassis id. The type is string.
    ChassisId interface{}

    // Outgoing port identifier. The type is string.
    PortIdDetail interface{}

    // System Name. The type is string.
    SystemName interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // Receive Interface. The type is string.
    RecvIntf interface{}

    // Remaining hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}
}

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) GetFilter() yfilter.YFilter { return lldpNeighborBriefEntry.YFilter }

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) SetFilter(yf yfilter.YFilter) { lldpNeighborBriefEntry.YFilter = yf }

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) GetGoName(yname string) string {
    if yname == "chassis-id" { return "ChassisId" }
    if yname == "port-id-detail" { return "PortIdDetail" }
    if yname == "system-name" { return "SystemName" }
    if yname == "enabled-capabilities" { return "EnabledCapabilities" }
    if yname == "recv-intf" { return "RecvIntf" }
    if yname == "hold-time" { return "HoldTime" }
    return ""
}

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) GetSegmentPath() string {
    return "lldp-neighbor-brief-entry"
}

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["chassis-id"] = lldpNeighborBriefEntry.ChassisId
    leafs["port-id-detail"] = lldpNeighborBriefEntry.PortIdDetail
    leafs["system-name"] = lldpNeighborBriefEntry.SystemName
    leafs["enabled-capabilities"] = lldpNeighborBriefEntry.EnabledCapabilities
    leafs["recv-intf"] = lldpNeighborBriefEntry.RecvIntf
    leafs["hold-time"] = lldpNeighborBriefEntry.HoldTime
    return leafs
}

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) GetYangName() string { return "lldp-neighbor-brief-entry" }

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) SetParent(parent types.Entity) { lldpNeighborBriefEntry.parent = parent }

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) GetParent() types.Entity { return lldpNeighborBriefEntry.parent }

func (lldpNeighborBriefEntry *LldpSnoopData_LldpNeighborBrief_Neighbours_LldpNeighborBriefEntry) GetParentYangName() string { return "neighbours" }

// LldpSnoopData_EthernetControllerNames
// Ethernet controller snoop data
type LldpSnoopData_EthernetControllerNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // port Name. The type is slice of
    // LldpSnoopData_EthernetControllerNames_EthernetControllerName.
    EthernetControllerName []LldpSnoopData_EthernetControllerNames_EthernetControllerName
}

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetFilter() yfilter.YFilter { return ethernetControllerNames.YFilter }

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) SetFilter(yf yfilter.YFilter) { ethernetControllerNames.YFilter = yf }

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetGoName(yname string) string {
    if yname == "ethernet-controller-name" { return "EthernetControllerName" }
    return ""
}

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetSegmentPath() string {
    return "ethernet-controller-names"
}

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-controller-name" {
        for _, c := range ethernetControllerNames.EthernetControllerName {
            if ethernetControllerNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LldpSnoopData_EthernetControllerNames_EthernetControllerName{}
        ethernetControllerNames.EthernetControllerName = append(ethernetControllerNames.EthernetControllerName, child)
        return &ethernetControllerNames.EthernetControllerName[len(ethernetControllerNames.EthernetControllerName)-1]
    }
    return nil
}

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetControllerNames.EthernetControllerName {
        children[ethernetControllerNames.EthernetControllerName[i].GetSegmentPath()] = &ethernetControllerNames.EthernetControllerName[i]
    }
    return children
}

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetYangName() string { return "ethernet-controller-names" }

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) SetParent(parent types.Entity) { ethernetControllerNames.parent = parent }

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetParent() types.Entity { return ethernetControllerNames.parent }

func (ethernetControllerNames *LldpSnoopData_EthernetControllerNames) GetParentYangName() string { return "lldp-snoop-data" }

// LldpSnoopData_EthernetControllerNames_EthernetControllerName
// port Name
type LldpSnoopData_EthernetControllerNames_EthernetControllerName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Name interface{}

    // Mac address of the neighbor. The type is string.
    SourceMac interface{}

    // Chassis id. The type is string.
    ChassisId interface{}

    // Outgoing port identifier. The type is string.
    PortIdDetail interface{}

    // Remaining hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}

    // Port Description. The type is string.
    PortDescription interface{}

    // System Name. The type is string.
    SystemName interface{}

    // System Description. The type is string.
    SystemDescription interface{}

    // System Capabilities. The type is string.
    SystemCapabilities interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // LldpNeighbor. The type is string with length: 0..40.
    LldpNeighbor interface{}

    // LLDP Packet Drop Enabled. The type is bool.
    DropEnabled interface{}

    // Received LLDP Packets count. The type is interface{} with range:
    // 0..18446744073709551615.
    RxLldpPkts interface{}

    // Management Addresses.
    NetworkAddresses LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses
}

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetFilter() yfilter.YFilter { return ethernetControllerName.YFilter }

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) SetFilter(yf yfilter.YFilter) { ethernetControllerName.YFilter = yf }

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "source-mac" { return "SourceMac" }
    if yname == "chassis-id" { return "ChassisId" }
    if yname == "port-id-detail" { return "PortIdDetail" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "port-description" { return "PortDescription" }
    if yname == "system-name" { return "SystemName" }
    if yname == "system-description" { return "SystemDescription" }
    if yname == "system-capabilities" { return "SystemCapabilities" }
    if yname == "enabled-capabilities" { return "EnabledCapabilities" }
    if yname == "lldp-neighbor" { return "LldpNeighbor" }
    if yname == "drop-enabled" { return "DropEnabled" }
    if yname == "rx-lldp-pkts" { return "RxLldpPkts" }
    if yname == "network-addresses" { return "NetworkAddresses" }
    return ""
}

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetSegmentPath() string {
    return "ethernet-controller-name" + "[name='" + fmt.Sprintf("%v", ethernetControllerName.Name) + "']"
}

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-addresses" {
        return &ethernetControllerName.NetworkAddresses
    }
    return nil
}

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["network-addresses"] = &ethernetControllerName.NetworkAddresses
    return children
}

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = ethernetControllerName.Name
    leafs["source-mac"] = ethernetControllerName.SourceMac
    leafs["chassis-id"] = ethernetControllerName.ChassisId
    leafs["port-id-detail"] = ethernetControllerName.PortIdDetail
    leafs["hold-time"] = ethernetControllerName.HoldTime
    leafs["port-description"] = ethernetControllerName.PortDescription
    leafs["system-name"] = ethernetControllerName.SystemName
    leafs["system-description"] = ethernetControllerName.SystemDescription
    leafs["system-capabilities"] = ethernetControllerName.SystemCapabilities
    leafs["enabled-capabilities"] = ethernetControllerName.EnabledCapabilities
    leafs["lldp-neighbor"] = ethernetControllerName.LldpNeighbor
    leafs["drop-enabled"] = ethernetControllerName.DropEnabled
    leafs["rx-lldp-pkts"] = ethernetControllerName.RxLldpPkts
    return leafs
}

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetYangName() string { return "ethernet-controller-name" }

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) SetParent(parent types.Entity) { ethernetControllerName.parent = parent }

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetParent() types.Entity { return ethernetControllerName.parent }

func (ethernetControllerName *LldpSnoopData_EthernetControllerNames_EthernetControllerName) GetParentYangName() string { return "ethernet-controller-names" }

// LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses
// Management Addresses
type LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lldp addr entry. The type is slice of
    // LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry.
    LldpAddrEntry []LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry
}

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetFilter() yfilter.YFilter { return networkAddresses.YFilter }

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) SetFilter(yf yfilter.YFilter) { networkAddresses.YFilter = yf }

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetGoName(yname string) string {
    if yname == "lldp-addr-entry" { return "LldpAddrEntry" }
    return ""
}

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetSegmentPath() string {
    return "network-addresses"
}

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-addr-entry" {
        for _, c := range networkAddresses.LldpAddrEntry {
            if networkAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry{}
        networkAddresses.LldpAddrEntry = append(networkAddresses.LldpAddrEntry, child)
        return &networkAddresses.LldpAddrEntry[len(networkAddresses.LldpAddrEntry)-1]
    }
    return nil
}

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networkAddresses.LldpAddrEntry {
        children[networkAddresses.LldpAddrEntry[i].GetSegmentPath()] = &networkAddresses.LldpAddrEntry[i]
    }
    return children
}

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetYangName() string { return "network-addresses" }

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) SetParent(parent types.Entity) { networkAddresses.parent = parent }

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetParent() types.Entity { return networkAddresses.parent }

func (networkAddresses *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses) GetParentYangName() string { return "ethernet-controller-name" }

// LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry
// lldp addr entry
type LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MA sub type. The type is interface{} with range: 0..255.
    MaSubtype interface{}

    // Interface num. The type is interface{} with range: 0..4294967295.
    IfNum interface{}

    // Network layer address.
    Address LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address
}

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetFilter() yfilter.YFilter { return lldpAddrEntry.YFilter }

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) SetFilter(yf yfilter.YFilter) { lldpAddrEntry.YFilter = yf }

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetGoName(yname string) string {
    if yname == "ma-subtype" { return "MaSubtype" }
    if yname == "if-num" { return "IfNum" }
    if yname == "address" { return "Address" }
    return ""
}

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetSegmentPath() string {
    return "lldp-addr-entry"
}

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &lldpAddrEntry.Address
    }
    return nil
}

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &lldpAddrEntry.Address
    return children
}

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ma-subtype"] = lldpAddrEntry.MaSubtype
    leafs["if-num"] = lldpAddrEntry.IfNum
    return leafs
}

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetYangName() string { return "lldp-addr-entry" }

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) SetParent(parent types.Entity) { lldpAddrEntry.parent = parent }

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetParent() types.Entity { return lldpAddrEntry.parent }

func (lldpAddrEntry *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry) GetParentYangName() string { return "network-addresses" }

// LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address
// Network layer address
type LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AddressType. The type is LldpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetGoName(yname string) string {
    if yname == "address-type" { return "AddressType" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetSegmentPath() string {
    return "address"
}

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-type"] = address.AddressType
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetYangName() string { return "address" }

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetParent() types.Entity { return address.parent }

func (address *LldpSnoopData_EthernetControllerNames_EthernetControllerName_NetworkAddresses_LldpAddrEntry_Address) GetParentYangName() string { return "lldp-addr-entry" }

