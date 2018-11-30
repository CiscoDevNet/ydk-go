// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-rt-check package configuration.
// 
// This module contains definitions
// for the following management objects:
//   rcc: RCC (Route Consistency Checker) configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_rt_check_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_rt_check_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rt-check-cfg rcc}", reflect.TypeOf(Rcc{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rt-check-cfg:rcc", reflect.TypeOf(Rcc{}))
}

// Rcc
// RCC (Route Consistency Checker) configuration
type Rcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RCC/LCC configuration for IPv6.
    Ipv6 Rcc_Ipv6

    // RCC/LCC configuration for IPv4.
    Ipv4 Rcc_Ipv4
}

func (rcc *Rcc) GetEntityData() *types.CommonEntityData {
    rcc.EntityData.YFilter = rcc.YFilter
    rcc.EntityData.YangName = "rcc"
    rcc.EntityData.BundleName = "cisco_ios_xr"
    rcc.EntityData.ParentYangName = "Cisco-IOS-XR-infra-rt-check-cfg"
    rcc.EntityData.SegmentPath = "Cisco-IOS-XR-infra-rt-check-cfg:rcc"
    rcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rcc.EntityData.Children = types.NewOrderedMap()
    rcc.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &rcc.Ipv6})
    rcc.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &rcc.Ipv4})
    rcc.EntityData.Leafs = types.NewOrderedMap()

    rcc.EntityData.YListKeys = []string {}

    return &(rcc.EntityData)
}

// Rcc_Ipv6
// RCC/LCC configuration for IPv6
type Rcc_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4/IPv6 LCC (Label Consistency Checker) configuration.
    Lcc Rcc_Ipv6_Lcc

    // RCC configuration for unicast.
    Unicast Rcc_Ipv6_Unicast

    // RCC configuration for multicast.
    Multicast Rcc_Ipv6_Multicast
}

func (ipv6 *Rcc_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "rcc"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("lcc", types.YChild{"Lcc", &ipv6.Lcc})
    ipv6.EntityData.Children.Append("unicast", types.YChild{"Unicast", &ipv6.Unicast})
    ipv6.EntityData.Children.Append("multicast", types.YChild{"Multicast", &ipv6.Multicast})
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Rcc_Ipv6_Lcc
// IPv4/IPv6 LCC (Label Consistency Checker)
// configuration
type Rcc_Ipv6_Lcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Period of check in milliseconds. The type is interface{} with range:
    // 100..600000. Units are millisecond.
    Period interface{}

    // Enable RCC/LCC. The type is interface{}.
    Enable interface{}
}

func (lcc *Rcc_Ipv6_Lcc) GetEntityData() *types.CommonEntityData {
    lcc.EntityData.YFilter = lcc.YFilter
    lcc.EntityData.YangName = "lcc"
    lcc.EntityData.BundleName = "cisco_ios_xr"
    lcc.EntityData.ParentYangName = "ipv6"
    lcc.EntityData.SegmentPath = "lcc"
    lcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcc.EntityData.Children = types.NewOrderedMap()
    lcc.EntityData.Leafs = types.NewOrderedMap()
    lcc.EntityData.Leafs.Append("period", types.YLeaf{"Period", lcc.Period})
    lcc.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", lcc.Enable})

    lcc.EntityData.YListKeys = []string {}

    return &(lcc.EntityData)
}

// Rcc_Ipv6_Unicast
// RCC configuration for unicast
type Rcc_Ipv6_Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Period of check in milliseconds. The type is interface{} with range:
    // 100..600000. Units are millisecond.
    Period interface{}

    // Enable RCC/LCC. The type is interface{}.
    Enable interface{}
}

func (unicast *Rcc_Ipv6_Unicast) GetEntityData() *types.CommonEntityData {
    unicast.EntityData.YFilter = unicast.YFilter
    unicast.EntityData.YangName = "unicast"
    unicast.EntityData.BundleName = "cisco_ios_xr"
    unicast.EntityData.ParentYangName = "ipv6"
    unicast.EntityData.SegmentPath = "unicast"
    unicast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unicast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unicast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unicast.EntityData.Children = types.NewOrderedMap()
    unicast.EntityData.Leafs = types.NewOrderedMap()
    unicast.EntityData.Leafs.Append("period", types.YLeaf{"Period", unicast.Period})
    unicast.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", unicast.Enable})

    unicast.EntityData.YListKeys = []string {}

    return &(unicast.EntityData)
}

// Rcc_Ipv6_Multicast
// RCC configuration for multicast
type Rcc_Ipv6_Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Period of check in milliseconds. The type is interface{} with range:
    // 100..600000. Units are millisecond.
    Period interface{}

    // Enable RCC/LCC. The type is interface{}.
    Enable interface{}
}

func (multicast *Rcc_Ipv6_Multicast) GetEntityData() *types.CommonEntityData {
    multicast.EntityData.YFilter = multicast.YFilter
    multicast.EntityData.YangName = "multicast"
    multicast.EntityData.BundleName = "cisco_ios_xr"
    multicast.EntityData.ParentYangName = "ipv6"
    multicast.EntityData.SegmentPath = "multicast"
    multicast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicast.EntityData.Children = types.NewOrderedMap()
    multicast.EntityData.Leafs = types.NewOrderedMap()
    multicast.EntityData.Leafs.Append("period", types.YLeaf{"Period", multicast.Period})
    multicast.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", multicast.Enable})

    multicast.EntityData.YListKeys = []string {}

    return &(multicast.EntityData)
}

// Rcc_Ipv4
// RCC/LCC configuration for IPv4
type Rcc_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4/IPv6 LCC (Label Consistency Checker) configuration.
    Lcc Rcc_Ipv4_Lcc

    // RCC configuration for unicast.
    Unicast Rcc_Ipv4_Unicast

    // RCC configuration for multicast.
    Multicast Rcc_Ipv4_Multicast
}

func (ipv4 *Rcc_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "rcc"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("lcc", types.YChild{"Lcc", &ipv4.Lcc})
    ipv4.EntityData.Children.Append("unicast", types.YChild{"Unicast", &ipv4.Unicast})
    ipv4.EntityData.Children.Append("multicast", types.YChild{"Multicast", &ipv4.Multicast})
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Rcc_Ipv4_Lcc
// IPv4/IPv6 LCC (Label Consistency Checker)
// configuration
type Rcc_Ipv4_Lcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Period of check in milliseconds. The type is interface{} with range:
    // 100..600000. Units are millisecond.
    Period interface{}

    // Enable RCC/LCC. The type is interface{}.
    Enable interface{}
}

func (lcc *Rcc_Ipv4_Lcc) GetEntityData() *types.CommonEntityData {
    lcc.EntityData.YFilter = lcc.YFilter
    lcc.EntityData.YangName = "lcc"
    lcc.EntityData.BundleName = "cisco_ios_xr"
    lcc.EntityData.ParentYangName = "ipv4"
    lcc.EntityData.SegmentPath = "lcc"
    lcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcc.EntityData.Children = types.NewOrderedMap()
    lcc.EntityData.Leafs = types.NewOrderedMap()
    lcc.EntityData.Leafs.Append("period", types.YLeaf{"Period", lcc.Period})
    lcc.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", lcc.Enable})

    lcc.EntityData.YListKeys = []string {}

    return &(lcc.EntityData)
}

// Rcc_Ipv4_Unicast
// RCC configuration for unicast
type Rcc_Ipv4_Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Period of check in milliseconds. The type is interface{} with range:
    // 100..600000. Units are millisecond.
    Period interface{}

    // Enable RCC/LCC. The type is interface{}.
    Enable interface{}
}

func (unicast *Rcc_Ipv4_Unicast) GetEntityData() *types.CommonEntityData {
    unicast.EntityData.YFilter = unicast.YFilter
    unicast.EntityData.YangName = "unicast"
    unicast.EntityData.BundleName = "cisco_ios_xr"
    unicast.EntityData.ParentYangName = "ipv4"
    unicast.EntityData.SegmentPath = "unicast"
    unicast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unicast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unicast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unicast.EntityData.Children = types.NewOrderedMap()
    unicast.EntityData.Leafs = types.NewOrderedMap()
    unicast.EntityData.Leafs.Append("period", types.YLeaf{"Period", unicast.Period})
    unicast.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", unicast.Enable})

    unicast.EntityData.YListKeys = []string {}

    return &(unicast.EntityData)
}

// Rcc_Ipv4_Multicast
// RCC configuration for multicast
type Rcc_Ipv4_Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Period of check in milliseconds. The type is interface{} with range:
    // 100..600000. Units are millisecond.
    Period interface{}

    // Enable RCC/LCC. The type is interface{}.
    Enable interface{}
}

func (multicast *Rcc_Ipv4_Multicast) GetEntityData() *types.CommonEntityData {
    multicast.EntityData.YFilter = multicast.YFilter
    multicast.EntityData.YangName = "multicast"
    multicast.EntityData.BundleName = "cisco_ios_xr"
    multicast.EntityData.ParentYangName = "ipv4"
    multicast.EntityData.SegmentPath = "multicast"
    multicast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicast.EntityData.Children = types.NewOrderedMap()
    multicast.EntityData.Leafs = types.NewOrderedMap()
    multicast.EntityData.Leafs.Append("period", types.YLeaf{"Period", multicast.Period})
    multicast.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", multicast.Enable})

    multicast.EntityData.YListKeys = []string {}

    return &(multicast.EntityData)
}

