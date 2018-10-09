// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-icmp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   icmp: IP ICMP configuration data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ip_icmp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_icmp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-icmp-cfg icmp}", reflect.TypeOf(Icmp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-icmp-cfg:icmp", reflect.TypeOf(Icmp{}))
}

// SourcePolicy represents Source policy
type SourcePolicy string

const (
    // Set Source Selection Policy to vrf
    SourcePolicy_vrf SourcePolicy = "vrf"

    // Set Source Selection Policy to rfc
    SourcePolicy_rfc SourcePolicy = "rfc"
)

// Icmp
// IP ICMP configuration data
type Icmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Protocol Type.
    Ipv6 Icmp_Ipv6

    // IP Protocol Type.
    Ipv4 Icmp_Ipv4
}

func (icmp *Icmp) GetEntityData() *types.CommonEntityData {
    icmp.EntityData.YFilter = icmp.YFilter
    icmp.EntityData.YangName = "icmp"
    icmp.EntityData.BundleName = "cisco_ios_xr"
    icmp.EntityData.ParentYangName = "Cisco-IOS-XR-ip-icmp-cfg"
    icmp.EntityData.SegmentPath = "Cisco-IOS-XR-ip-icmp-cfg:icmp"
    icmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmp.EntityData.Children = types.NewOrderedMap()
    icmp.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &icmp.Ipv6})
    icmp.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &icmp.Ipv4})
    icmp.EntityData.Leafs = types.NewOrderedMap()

    icmp.EntityData.YListKeys = []string {}

    return &(icmp.EntityData)
}

// Icmp_Ipv6
// IP Protocol Type
type Icmp_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set ratelimit of ICMP packets.
    RateLimit Icmp_Ipv6_RateLimit

    // IP ICMP Source Address Selection Policy.
    Source Icmp_Ipv6_Source
}

func (ipv6 *Icmp_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "icmp"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("rate-limit", types.YChild{"RateLimit", &ipv6.RateLimit})
    ipv6.EntityData.Children.Append("source", types.YChild{"Source", &ipv6.Source})
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Icmp_Ipv6_RateLimit
// Set ratelimit of ICMP packets
type Icmp_Ipv6_RateLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set unreachable ICMP packets ratelimit.
    Unreachable Icmp_Ipv6_RateLimit_Unreachable
}

func (rateLimit *Icmp_Ipv6_RateLimit) GetEntityData() *types.CommonEntityData {
    rateLimit.EntityData.YFilter = rateLimit.YFilter
    rateLimit.EntityData.YangName = "rate-limit"
    rateLimit.EntityData.BundleName = "cisco_ios_xr"
    rateLimit.EntityData.ParentYangName = "ipv6"
    rateLimit.EntityData.SegmentPath = "rate-limit"
    rateLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rateLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rateLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rateLimit.EntityData.Children = types.NewOrderedMap()
    rateLimit.EntityData.Children.Append("unreachable", types.YChild{"Unreachable", &rateLimit.Unreachable})
    rateLimit.EntityData.Leafs = types.NewOrderedMap()

    rateLimit.EntityData.YListKeys = []string {}

    return &(rateLimit.EntityData)
}

// Icmp_Ipv6_RateLimit_Unreachable
// Set unreachable ICMP packets ratelimit
type Icmp_Ipv6_RateLimit_Unreachable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rate Limit of Unreachable ICMP packets. The type is interface{} with range:
    // 0..4294967295.
    Rate interface{}

    // Rate Limit of Unreachable DF packets. The type is interface{} with range:
    // 0..4294967295.
    Fragmentation interface{}
}

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetEntityData() *types.CommonEntityData {
    unreachable.EntityData.YFilter = unreachable.YFilter
    unreachable.EntityData.YangName = "unreachable"
    unreachable.EntityData.BundleName = "cisco_ios_xr"
    unreachable.EntityData.ParentYangName = "rate-limit"
    unreachable.EntityData.SegmentPath = "unreachable"
    unreachable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unreachable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unreachable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unreachable.EntityData.Children = types.NewOrderedMap()
    unreachable.EntityData.Leafs = types.NewOrderedMap()
    unreachable.EntityData.Leafs.Append("rate", types.YLeaf{"Rate", unreachable.Rate})
    unreachable.EntityData.Leafs.Append("fragmentation", types.YLeaf{"Fragmentation", unreachable.Fragmentation})

    unreachable.EntityData.YListKeys = []string {}

    return &(unreachable.EntityData)
}

// Icmp_Ipv6_Source
// IP ICMP Source Address Selection Policy
type Icmp_Ipv6_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Source Address Policy. The type is SourcePolicy.
    SourceAddressPolicy interface{}
}

func (source *Icmp_Ipv6_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "ipv6"
    source.EntityData.SegmentPath = "source"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = types.NewOrderedMap()
    source.EntityData.Leafs = types.NewOrderedMap()
    source.EntityData.Leafs.Append("source-address-policy", types.YLeaf{"SourceAddressPolicy", source.SourceAddressPolicy})

    source.EntityData.YListKeys = []string {}

    return &(source.EntityData)
}

// Icmp_Ipv4
// IP Protocol Type
type Icmp_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set ratelimit of ICMP packets.
    RateLimit Icmp_Ipv4_RateLimit

    // IP ICMP Source Address Selection Policy.
    Source Icmp_Ipv4_Source
}

func (ipv4 *Icmp_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "icmp"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("rate-limit", types.YChild{"RateLimit", &ipv4.RateLimit})
    ipv4.EntityData.Children.Append("source", types.YChild{"Source", &ipv4.Source})
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Icmp_Ipv4_RateLimit
// Set ratelimit of ICMP packets
type Icmp_Ipv4_RateLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set unreachable ICMP packets ratelimit.
    Unreachable Icmp_Ipv4_RateLimit_Unreachable
}

func (rateLimit *Icmp_Ipv4_RateLimit) GetEntityData() *types.CommonEntityData {
    rateLimit.EntityData.YFilter = rateLimit.YFilter
    rateLimit.EntityData.YangName = "rate-limit"
    rateLimit.EntityData.BundleName = "cisco_ios_xr"
    rateLimit.EntityData.ParentYangName = "ipv4"
    rateLimit.EntityData.SegmentPath = "rate-limit"
    rateLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rateLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rateLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rateLimit.EntityData.Children = types.NewOrderedMap()
    rateLimit.EntityData.Children.Append("unreachable", types.YChild{"Unreachable", &rateLimit.Unreachable})
    rateLimit.EntityData.Leafs = types.NewOrderedMap()

    rateLimit.EntityData.YListKeys = []string {}

    return &(rateLimit.EntityData)
}

// Icmp_Ipv4_RateLimit_Unreachable
// Set unreachable ICMP packets ratelimit
type Icmp_Ipv4_RateLimit_Unreachable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rate Limit of Unreachable ICMP packets. The type is interface{} with range:
    // 0..4294967295.
    Rate interface{}

    // Rate Limit of Unreachable DF packets. The type is interface{} with range:
    // 0..4294967295.
    Fragmentation interface{}
}

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetEntityData() *types.CommonEntityData {
    unreachable.EntityData.YFilter = unreachable.YFilter
    unreachable.EntityData.YangName = "unreachable"
    unreachable.EntityData.BundleName = "cisco_ios_xr"
    unreachable.EntityData.ParentYangName = "rate-limit"
    unreachable.EntityData.SegmentPath = "unreachable"
    unreachable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unreachable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unreachable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unreachable.EntityData.Children = types.NewOrderedMap()
    unreachable.EntityData.Leafs = types.NewOrderedMap()
    unreachable.EntityData.Leafs.Append("rate", types.YLeaf{"Rate", unreachable.Rate})
    unreachable.EntityData.Leafs.Append("fragmentation", types.YLeaf{"Fragmentation", unreachable.Fragmentation})

    unreachable.EntityData.YListKeys = []string {}

    return &(unreachable.EntityData)
}

// Icmp_Ipv4_Source
// IP ICMP Source Address Selection Policy
type Icmp_Ipv4_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Source Address Policy. The type is SourcePolicy.
    SourceAddressPolicy interface{}
}

func (source *Icmp_Ipv4_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "ipv4"
    source.EntityData.SegmentPath = "source"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = types.NewOrderedMap()
    source.EntityData.Leafs = types.NewOrderedMap()
    source.EntityData.Leafs.Append("source-address-policy", types.YLeaf{"SourceAddressPolicy", source.SourceAddressPolicy})

    source.EntityData.YListKeys = []string {}

    return &(source.EntityData)
}

