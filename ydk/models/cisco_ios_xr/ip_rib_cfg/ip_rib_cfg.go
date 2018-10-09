// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-rib package configuration.
// 
// This module contains definitions
// for the following management objects:
//   rib: RIB configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-infra-rsi-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ip_rib_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_rib_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-rib-cfg rib}", reflect.TypeOf(Rib{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-rib-cfg:rib", reflect.TypeOf(Rib{}))
}

// Rib
// RIB configuration.
type Rib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set maximum depth for route recursion check. The type is interface{} with
    // range: 5..16.
    MaxRecursionDepth interface{}

    // RIB address family configuration.
    Af Rib_Af
}

func (rib *Rib) GetEntityData() *types.CommonEntityData {
    rib.EntityData.YFilter = rib.YFilter
    rib.EntityData.YangName = "rib"
    rib.EntityData.BundleName = "cisco_ios_xr"
    rib.EntityData.ParentYangName = "Cisco-IOS-XR-ip-rib-cfg"
    rib.EntityData.SegmentPath = "Cisco-IOS-XR-ip-rib-cfg:rib"
    rib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rib.EntityData.Children = types.NewOrderedMap()
    rib.EntityData.Children.Append("af", types.YChild{"Af", &rib.Af})
    rib.EntityData.Leafs = types.NewOrderedMap()
    rib.EntityData.Leafs.Append("max-recursion-depth", types.YLeaf{"MaxRecursionDepth", rib.MaxRecursionDepth})

    rib.EntityData.YListKeys = []string {}

    return &(rib.EntityData)
}

// Rib_Af
// RIB address family configuration
type Rib_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 configuration.
    Ipv4 Rib_Af_Ipv4

    // IPv6 configuration.
    Ipv6 Rib_Af_Ipv6
}

func (af *Rib_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "rib"
    af.EntityData.SegmentPath = "af"
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &af.Ipv4})
    af.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &af.Ipv6})
    af.EntityData.Leafs = types.NewOrderedMap()

    af.EntityData.YListKeys = []string {}

    return &(af.EntityData)
}

// Rib_Af_Ipv4
// IPv4 configuration
type Rib_Af_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable next-hop dampening. The type is interface{}.
    NextHopDampeningDisable interface{}

    // Redistribution history related configs.
    RedistributionHistory Rib_Af_Ipv4_RedistributionHistory
}

func (ipv4 *Rib_Af_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "af"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("redistribution-history", types.YChild{"RedistributionHistory", &ipv4.RedistributionHistory})
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("next-hop-dampening-disable", types.YLeaf{"NextHopDampeningDisable", ipv4.NextHopDampeningDisable})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Rib_Af_Ipv4_RedistributionHistory
// Redistribution history related configs
type Rib_Af_Ipv4_RedistributionHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum BCDL redistribution history size. The type is interface{} with
    // range: 10..2000000.
    BcdlClient interface{}

    // Maximum protocol redistribution history size. The type is interface{} with
    // range: 10..250000.
    ProtocolClient interface{}

    // Retain redistribution history after disconnect.
    Keep Rib_Af_Ipv4_RedistributionHistory_Keep
}

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetEntityData() *types.CommonEntityData {
    redistributionHistory.EntityData.YFilter = redistributionHistory.YFilter
    redistributionHistory.EntityData.YangName = "redistribution-history"
    redistributionHistory.EntityData.BundleName = "cisco_ios_xr"
    redistributionHistory.EntityData.ParentYangName = "ipv4"
    redistributionHistory.EntityData.SegmentPath = "redistribution-history"
    redistributionHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributionHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributionHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributionHistory.EntityData.Children = types.NewOrderedMap()
    redistributionHistory.EntityData.Children.Append("keep", types.YChild{"Keep", &redistributionHistory.Keep})
    redistributionHistory.EntityData.Leafs = types.NewOrderedMap()
    redistributionHistory.EntityData.Leafs.Append("bcdl-client", types.YLeaf{"BcdlClient", redistributionHistory.BcdlClient})
    redistributionHistory.EntityData.Leafs.Append("protocol-client", types.YLeaf{"ProtocolClient", redistributionHistory.ProtocolClient})

    redistributionHistory.EntityData.YListKeys = []string {}

    return &(redistributionHistory.EntityData)
}

// Rib_Af_Ipv4_RedistributionHistory_Keep
// Retain redistribution history after disconnect.
type Rib_Af_Ipv4_RedistributionHistory_Keep struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable retain BCDL history. The type is interface{}.
    Bcdl interface{}
}

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetEntityData() *types.CommonEntityData {
    keep.EntityData.YFilter = keep.YFilter
    keep.EntityData.YangName = "keep"
    keep.EntityData.BundleName = "cisco_ios_xr"
    keep.EntityData.ParentYangName = "redistribution-history"
    keep.EntityData.SegmentPath = "keep"
    keep.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keep.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keep.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keep.EntityData.Children = types.NewOrderedMap()
    keep.EntityData.Leafs = types.NewOrderedMap()
    keep.EntityData.Leafs.Append("bcdl", types.YLeaf{"Bcdl", keep.Bcdl})

    keep.EntityData.YListKeys = []string {}

    return &(keep.EntityData)
}

// Rib_Af_Ipv6
// IPv6 configuration
type Rib_Af_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable next-hop dampening. The type is interface{}.
    NextHopDampeningDisable interface{}

    // Redistribution history related configs.
    RedistributionHistory Rib_Af_Ipv6_RedistributionHistory
}

func (ipv6 *Rib_Af_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "af"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("redistribution-history", types.YChild{"RedistributionHistory", &ipv6.RedistributionHistory})
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("next-hop-dampening-disable", types.YLeaf{"NextHopDampeningDisable", ipv6.NextHopDampeningDisable})

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Rib_Af_Ipv6_RedistributionHistory
// Redistribution history related configs
type Rib_Af_Ipv6_RedistributionHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum BCDL redistribution history size. The type is interface{} with
    // range: 10..2000000.
    BcdlClient interface{}

    // Maximum protocol redistribution history size. The type is interface{} with
    // range: 10..250000.
    ProtocolClient interface{}

    // Retain redistribution history after disconnect.
    Keep Rib_Af_Ipv6_RedistributionHistory_Keep
}

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetEntityData() *types.CommonEntityData {
    redistributionHistory.EntityData.YFilter = redistributionHistory.YFilter
    redistributionHistory.EntityData.YangName = "redistribution-history"
    redistributionHistory.EntityData.BundleName = "cisco_ios_xr"
    redistributionHistory.EntityData.ParentYangName = "ipv6"
    redistributionHistory.EntityData.SegmentPath = "redistribution-history"
    redistributionHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributionHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributionHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributionHistory.EntityData.Children = types.NewOrderedMap()
    redistributionHistory.EntityData.Children.Append("keep", types.YChild{"Keep", &redistributionHistory.Keep})
    redistributionHistory.EntityData.Leafs = types.NewOrderedMap()
    redistributionHistory.EntityData.Leafs.Append("bcdl-client", types.YLeaf{"BcdlClient", redistributionHistory.BcdlClient})
    redistributionHistory.EntityData.Leafs.Append("protocol-client", types.YLeaf{"ProtocolClient", redistributionHistory.ProtocolClient})

    redistributionHistory.EntityData.YListKeys = []string {}

    return &(redistributionHistory.EntityData)
}

// Rib_Af_Ipv6_RedistributionHistory_Keep
// Retain redistribution history after disconnect.
type Rib_Af_Ipv6_RedistributionHistory_Keep struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable retain BCDL history. The type is interface{}.
    Bcdl interface{}
}

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetEntityData() *types.CommonEntityData {
    keep.EntityData.YFilter = keep.YFilter
    keep.EntityData.YangName = "keep"
    keep.EntityData.BundleName = "cisco_ios_xr"
    keep.EntityData.ParentYangName = "redistribution-history"
    keep.EntityData.SegmentPath = "keep"
    keep.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keep.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keep.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keep.EntityData.Children = types.NewOrderedMap()
    keep.EntityData.Leafs = types.NewOrderedMap()
    keep.EntityData.Leafs.Append("bcdl", types.YLeaf{"Bcdl", keep.Bcdl})

    keep.EntityData.YListKeys = []string {}

    return &(keep.EntityData)
}

