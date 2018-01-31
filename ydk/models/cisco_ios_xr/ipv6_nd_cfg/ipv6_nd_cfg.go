// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-nd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ipv6-neighbor: IPv6 neighbor or neighbor discovery
//     configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_nd_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_nd_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-nd-cfg ipv6-neighbor}", reflect.TypeOf(Ipv6Neighbor{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-nd-cfg:ipv6-neighbor", reflect.TypeOf(Ipv6Neighbor{}))
}

// Ipv6ndMonth represents Ipv6nd month
type Ipv6ndMonth string

const (
    // January
    Ipv6ndMonth_january Ipv6ndMonth = "january"

    // February
    Ipv6ndMonth_february Ipv6ndMonth = "february"

    // March
    Ipv6ndMonth_march Ipv6ndMonth = "march"

    // April
    Ipv6ndMonth_april Ipv6ndMonth = "april"

    // May
    Ipv6ndMonth_may Ipv6ndMonth = "may"

    // June
    Ipv6ndMonth_june Ipv6ndMonth = "june"

    // July
    Ipv6ndMonth_july Ipv6ndMonth = "july"

    // August
    Ipv6ndMonth_august Ipv6ndMonth = "august"

    // September
    Ipv6ndMonth_september Ipv6ndMonth = "september"

    // October
    Ipv6ndMonth_october Ipv6ndMonth = "october"

    // November
    Ipv6ndMonth_november Ipv6ndMonth = "november"

    // December
    Ipv6ndMonth_december Ipv6ndMonth = "december"
)

// Ipv6NdRouterPref represents Ipv6 nd router pref
type Ipv6NdRouterPref string

const (
    // High preference
    Ipv6NdRouterPref_high Ipv6NdRouterPref = "high"

    // Medium preference
    Ipv6NdRouterPref_medium Ipv6NdRouterPref = "medium"

    // Low preference
    Ipv6NdRouterPref_low Ipv6NdRouterPref = "low"
)

// Ipv6srpEncapsulation represents Ipv6srp encapsulation
type Ipv6srpEncapsulation string

const (
    // Encapsulation type SRP, prefer side A
    Ipv6srpEncapsulation_srpa Ipv6srpEncapsulation = "srpa"

    // Encapsulation type SRP, prefer side B
    Ipv6srpEncapsulation_srpb Ipv6srpEncapsulation = "srpb"
)

// Ipv6Neighbor
// IPv6 neighbor or neighbor discovery configuration
type Ipv6Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set lifetime for stale neighbor. The type is interface{} with range:
    // 1..43200. Units are second.
    ScavengeTimeout interface{}

    // IPv6 neighbors.
    Neighbors Ipv6Neighbor_Neighbors
}

func (ipv6Neighbor *Ipv6Neighbor) GetFilter() yfilter.YFilter { return ipv6Neighbor.YFilter }

func (ipv6Neighbor *Ipv6Neighbor) SetFilter(yf yfilter.YFilter) { ipv6Neighbor.YFilter = yf }

func (ipv6Neighbor *Ipv6Neighbor) GetGoName(yname string) string {
    if yname == "scavenge-timeout" { return "ScavengeTimeout" }
    if yname == "neighbors" { return "Neighbors" }
    return ""
}

func (ipv6Neighbor *Ipv6Neighbor) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-nd-cfg:ipv6-neighbor"
}

func (ipv6Neighbor *Ipv6Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbors" {
        return &ipv6Neighbor.Neighbors
    }
    return nil
}

func (ipv6Neighbor *Ipv6Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbors"] = &ipv6Neighbor.Neighbors
    return children
}

func (ipv6Neighbor *Ipv6Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["scavenge-timeout"] = ipv6Neighbor.ScavengeTimeout
    return leafs
}

func (ipv6Neighbor *Ipv6Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Neighbor *Ipv6Neighbor) GetYangName() string { return "ipv6-neighbor" }

func (ipv6Neighbor *Ipv6Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Neighbor *Ipv6Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Neighbor *Ipv6Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Neighbor *Ipv6Neighbor) SetParent(parent types.Entity) { ipv6Neighbor.parent = parent }

func (ipv6Neighbor *Ipv6Neighbor) GetParent() types.Entity { return ipv6Neighbor.parent }

func (ipv6Neighbor *Ipv6Neighbor) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-nd-cfg" }

// Ipv6Neighbor_Neighbors
// IPv6 neighbors
type Ipv6Neighbor_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 neighbor configuration. The type is slice of
    // Ipv6Neighbor_Neighbors_Neighbor.
    Neighbor []Ipv6Neighbor_Neighbors_Neighbor
}

func (neighbors *Ipv6Neighbor_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Ipv6Neighbor_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Ipv6Neighbor_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *Ipv6Neighbor_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Ipv6Neighbor_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6Neighbor_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *Ipv6Neighbor_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *Ipv6Neighbor_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Ipv6Neighbor_Neighbors) GetBundleName() string { return "cisco_ios_xr" }

func (neighbors *Ipv6Neighbor_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Ipv6Neighbor_Neighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbors *Ipv6Neighbor_Neighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbors *Ipv6Neighbor_Neighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbors *Ipv6Neighbor_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Ipv6Neighbor_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Ipv6Neighbor_Neighbors) GetParentYangName() string { return "ipv6-neighbor" }

// Ipv6Neighbor_Neighbors_Neighbor
// IPv6 neighbor configuration
type Ipv6Neighbor_Neighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // IPv6 address zone. The type is string. The default value is 0.
    Zone interface{}

    // 48-bit hardware address H.H.H. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}. This attribute is mandatory.
    MacAddress interface{}

    // Encapsulation type only if interface type is SRP. The type is
    // Ipv6srpEncapsulation.
    Encapsulation interface{}
}

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "zone" { return "Zone" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "encapsulation" { return "Encapsulation" }
    return ""
}

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[neighbor-address='" + fmt.Sprintf("%v", neighbor.NeighborAddress) + "']" + "[interface-name='" + fmt.Sprintf("%v", neighbor.InterfaceName) + "']"
}

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-address"] = neighbor.NeighborAddress
    leafs["interface-name"] = neighbor.InterfaceName
    leafs["zone"] = neighbor.Zone
    leafs["mac-address"] = neighbor.MacAddress
    leafs["encapsulation"] = neighbor.Encapsulation
    return leafs
}

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

