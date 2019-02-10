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
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set lifetime for stale neighbor. The type is interface{} with range:
    // 1..43200. Units are second.
    ScavengeTimeout interface{}

    // IPv6 neighbors.
    Neighbors Ipv6Neighbor_Neighbors
}

func (ipv6Neighbor *Ipv6Neighbor) GetEntityData() *types.CommonEntityData {
    ipv6Neighbor.EntityData.YFilter = ipv6Neighbor.YFilter
    ipv6Neighbor.EntityData.YangName = "ipv6-neighbor"
    ipv6Neighbor.EntityData.BundleName = "cisco_ios_xr"
    ipv6Neighbor.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-nd-cfg"
    ipv6Neighbor.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-nd-cfg:ipv6-neighbor"
    ipv6Neighbor.EntityData.AbsolutePath = ipv6Neighbor.EntityData.SegmentPath
    ipv6Neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Neighbor.EntityData.Children = types.NewOrderedMap()
    ipv6Neighbor.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &ipv6Neighbor.Neighbors})
    ipv6Neighbor.EntityData.Leafs = types.NewOrderedMap()
    ipv6Neighbor.EntityData.Leafs.Append("scavenge-timeout", types.YLeaf{"ScavengeTimeout", ipv6Neighbor.ScavengeTimeout})

    ipv6Neighbor.EntityData.YListKeys = []string {}

    return &(ipv6Neighbor.EntityData)
}

// Ipv6Neighbor_Neighbors
// IPv6 neighbors
type Ipv6Neighbor_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 neighbor configuration. The type is slice of
    // Ipv6Neighbor_Neighbors_Neighbor.
    Neighbor []*Ipv6Neighbor_Neighbors_Neighbor
}

func (neighbors *Ipv6Neighbor_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "ipv6-neighbor"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-cfg:ipv6-neighbor/" + neighbors.EntityData.SegmentPath
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Ipv6Neighbor_Neighbors_Neighbor
// IPv6 neighbor configuration
type Ipv6Neighbor_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
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

func (neighbor *Ipv6Neighbor_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.NeighborAddress, "neighbor-address") + types.AddKeyToken(neighbor.InterfaceName, "interface-name")
    neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-nd-cfg:ipv6-neighbor/neighbors/" + neighbor.EntityData.SegmentPath
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", neighbor.NeighborAddress})
    neighbor.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", neighbor.InterfaceName})
    neighbor.EntityData.Leafs.Append("zone", types.YLeaf{"Zone", neighbor.Zone})
    neighbor.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", neighbor.MacAddress})
    neighbor.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", neighbor.Encapsulation})

    neighbor.EntityData.YListKeys = []string {"NeighborAddress", "InterfaceName"}

    return &(neighbor.EntityData)
}

