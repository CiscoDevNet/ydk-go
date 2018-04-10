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
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

    rib.EntityData.Children = make(map[string]types.YChild)
    rib.EntityData.Children["af"] = types.YChild{"Af", &rib.Af}
    rib.EntityData.Leafs = make(map[string]types.YLeaf)
    rib.EntityData.Leafs["max-recursion-depth"] = types.YLeaf{"MaxRecursionDepth", rib.MaxRecursionDepth}
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

    af.EntityData.Children = make(map[string]types.YChild)
    af.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &af.Ipv4}
    af.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &af.Ipv6}
    af.EntityData.Leafs = make(map[string]types.YLeaf)
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

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["redistribution-history"] = types.YChild{"RedistributionHistory", &ipv4.RedistributionHistory}
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4.EntityData.Leafs["next-hop-dampening-disable"] = types.YLeaf{"NextHopDampeningDisable", ipv4.NextHopDampeningDisable}
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

    redistributionHistory.EntityData.Children = make(map[string]types.YChild)
    redistributionHistory.EntityData.Children["keep"] = types.YChild{"Keep", &redistributionHistory.Keep}
    redistributionHistory.EntityData.Leafs = make(map[string]types.YLeaf)
    redistributionHistory.EntityData.Leafs["bcdl-client"] = types.YLeaf{"BcdlClient", redistributionHistory.BcdlClient}
    redistributionHistory.EntityData.Leafs["protocol-client"] = types.YLeaf{"ProtocolClient", redistributionHistory.ProtocolClient}
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

    keep.EntityData.Children = make(map[string]types.YChild)
    keep.EntityData.Leafs = make(map[string]types.YLeaf)
    keep.EntityData.Leafs["bcdl"] = types.YLeaf{"Bcdl", keep.Bcdl}
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

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["redistribution-history"] = types.YChild{"RedistributionHistory", &ipv6.RedistributionHistory}
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6.EntityData.Leafs["next-hop-dampening-disable"] = types.YLeaf{"NextHopDampeningDisable", ipv6.NextHopDampeningDisable}
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

    redistributionHistory.EntityData.Children = make(map[string]types.YChild)
    redistributionHistory.EntityData.Children["keep"] = types.YChild{"Keep", &redistributionHistory.Keep}
    redistributionHistory.EntityData.Leafs = make(map[string]types.YLeaf)
    redistributionHistory.EntityData.Leafs["bcdl-client"] = types.YLeaf{"BcdlClient", redistributionHistory.BcdlClient}
    redistributionHistory.EntityData.Leafs["protocol-client"] = types.YLeaf{"ProtocolClient", redistributionHistory.ProtocolClient}
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

    keep.EntityData.Children = make(map[string]types.YChild)
    keep.EntityData.Leafs = make(map[string]types.YLeaf)
    keep.EntityData.Leafs["bcdl"] = types.YLeaf{"Bcdl", keep.Bcdl}
    return &(keep.EntityData)
}

