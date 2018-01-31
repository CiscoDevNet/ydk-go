// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-icmp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   icmp: IP ICMP configuration data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // IP Protocol Type.
    Ipv6 Icmp_Ipv6

    // IP Protocol Type.
    Ipv4 Icmp_Ipv4
}

func (icmp *Icmp) GetFilter() yfilter.YFilter { return icmp.YFilter }

func (icmp *Icmp) SetFilter(yf yfilter.YFilter) { icmp.YFilter = yf }

func (icmp *Icmp) GetGoName(yname string) string {
    if yname == "ipv6" { return "Ipv6" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (icmp *Icmp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-icmp-cfg:icmp"
}

func (icmp *Icmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &icmp.Ipv6
    }
    if childYangName == "ipv4" {
        return &icmp.Ipv4
    }
    return nil
}

func (icmp *Icmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &icmp.Ipv6
    children["ipv4"] = &icmp.Ipv4
    return children
}

func (icmp *Icmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (icmp *Icmp) GetBundleName() string { return "cisco_ios_xr" }

func (icmp *Icmp) GetYangName() string { return "icmp" }

func (icmp *Icmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (icmp *Icmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (icmp *Icmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (icmp *Icmp) SetParent(parent types.Entity) { icmp.parent = parent }

func (icmp *Icmp) GetParent() types.Entity { return icmp.parent }

func (icmp *Icmp) GetParentYangName() string { return "Cisco-IOS-XR-ip-icmp-cfg" }

// Icmp_Ipv6
// IP Protocol Type
type Icmp_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set ratelimit of ICMP packets.
    RateLimit Icmp_Ipv6_RateLimit

    // IP ICMP Source Address Selection Policy.
    Source Icmp_Ipv6_Source
}

func (ipv6 *Icmp_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Icmp_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Icmp_Ipv6) GetGoName(yname string) string {
    if yname == "rate-limit" { return "RateLimit" }
    if yname == "source" { return "Source" }
    return ""
}

func (ipv6 *Icmp_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Icmp_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rate-limit" {
        return &ipv6.RateLimit
    }
    if childYangName == "source" {
        return &ipv6.Source
    }
    return nil
}

func (ipv6 *Icmp_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rate-limit"] = &ipv6.RateLimit
    children["source"] = &ipv6.Source
    return children
}

func (ipv6 *Icmp_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *Icmp_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Icmp_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Icmp_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Icmp_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Icmp_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Icmp_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Icmp_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Icmp_Ipv6) GetParentYangName() string { return "icmp" }

// Icmp_Ipv6_RateLimit
// Set ratelimit of ICMP packets
type Icmp_Ipv6_RateLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set unreachable ICMP packets ratelimit.
    Unreachable Icmp_Ipv6_RateLimit_Unreachable
}

func (rateLimit *Icmp_Ipv6_RateLimit) GetFilter() yfilter.YFilter { return rateLimit.YFilter }

func (rateLimit *Icmp_Ipv6_RateLimit) SetFilter(yf yfilter.YFilter) { rateLimit.YFilter = yf }

func (rateLimit *Icmp_Ipv6_RateLimit) GetGoName(yname string) string {
    if yname == "unreachable" { return "Unreachable" }
    return ""
}

func (rateLimit *Icmp_Ipv6_RateLimit) GetSegmentPath() string {
    return "rate-limit"
}

func (rateLimit *Icmp_Ipv6_RateLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unreachable" {
        return &rateLimit.Unreachable
    }
    return nil
}

func (rateLimit *Icmp_Ipv6_RateLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["unreachable"] = &rateLimit.Unreachable
    return children
}

func (rateLimit *Icmp_Ipv6_RateLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rateLimit *Icmp_Ipv6_RateLimit) GetBundleName() string { return "cisco_ios_xr" }

func (rateLimit *Icmp_Ipv6_RateLimit) GetYangName() string { return "rate-limit" }

func (rateLimit *Icmp_Ipv6_RateLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rateLimit *Icmp_Ipv6_RateLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rateLimit *Icmp_Ipv6_RateLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rateLimit *Icmp_Ipv6_RateLimit) SetParent(parent types.Entity) { rateLimit.parent = parent }

func (rateLimit *Icmp_Ipv6_RateLimit) GetParent() types.Entity { return rateLimit.parent }

func (rateLimit *Icmp_Ipv6_RateLimit) GetParentYangName() string { return "ipv6" }

// Icmp_Ipv6_RateLimit_Unreachable
// Set unreachable ICMP packets ratelimit
type Icmp_Ipv6_RateLimit_Unreachable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rate Limit of Unreachable ICMP packets. The type is interface{} with range:
    // 0..4294967295.
    Rate interface{}

    // Rate Limit of Unreachable DF packets. The type is interface{} with range:
    // 0..4294967295.
    Fragmentation interface{}
}

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetFilter() yfilter.YFilter { return unreachable.YFilter }

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) SetFilter(yf yfilter.YFilter) { unreachable.YFilter = yf }

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetGoName(yname string) string {
    if yname == "rate" { return "Rate" }
    if yname == "fragmentation" { return "Fragmentation" }
    return ""
}

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetSegmentPath() string {
    return "unreachable"
}

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rate"] = unreachable.Rate
    leafs["fragmentation"] = unreachable.Fragmentation
    return leafs
}

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetBundleName() string { return "cisco_ios_xr" }

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetYangName() string { return "unreachable" }

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) SetParent(parent types.Entity) { unreachable.parent = parent }

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetParent() types.Entity { return unreachable.parent }

func (unreachable *Icmp_Ipv6_RateLimit_Unreachable) GetParentYangName() string { return "rate-limit" }

// Icmp_Ipv6_Source
// IP ICMP Source Address Selection Policy
type Icmp_Ipv6_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Source Address Policy. The type is SourcePolicy.
    SourceAddressPolicy interface{}
}

func (source *Icmp_Ipv6_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *Icmp_Ipv6_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *Icmp_Ipv6_Source) GetGoName(yname string) string {
    if yname == "source-address-policy" { return "SourceAddressPolicy" }
    return ""
}

func (source *Icmp_Ipv6_Source) GetSegmentPath() string {
    return "source"
}

func (source *Icmp_Ipv6_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (source *Icmp_Ipv6_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (source *Icmp_Ipv6_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address-policy"] = source.SourceAddressPolicy
    return leafs
}

func (source *Icmp_Ipv6_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *Icmp_Ipv6_Source) GetYangName() string { return "source" }

func (source *Icmp_Ipv6_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *Icmp_Ipv6_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *Icmp_Ipv6_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *Icmp_Ipv6_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *Icmp_Ipv6_Source) GetParent() types.Entity { return source.parent }

func (source *Icmp_Ipv6_Source) GetParentYangName() string { return "ipv6" }

// Icmp_Ipv4
// IP Protocol Type
type Icmp_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set ratelimit of ICMP packets.
    RateLimit Icmp_Ipv4_RateLimit

    // IP ICMP Source Address Selection Policy.
    Source Icmp_Ipv4_Source
}

func (ipv4 *Icmp_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Icmp_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Icmp_Ipv4) GetGoName(yname string) string {
    if yname == "rate-limit" { return "RateLimit" }
    if yname == "source" { return "Source" }
    return ""
}

func (ipv4 *Icmp_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Icmp_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rate-limit" {
        return &ipv4.RateLimit
    }
    if childYangName == "source" {
        return &ipv4.Source
    }
    return nil
}

func (ipv4 *Icmp_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rate-limit"] = &ipv4.RateLimit
    children["source"] = &ipv4.Source
    return children
}

func (ipv4 *Icmp_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *Icmp_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Icmp_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Icmp_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Icmp_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Icmp_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Icmp_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Icmp_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Icmp_Ipv4) GetParentYangName() string { return "icmp" }

// Icmp_Ipv4_RateLimit
// Set ratelimit of ICMP packets
type Icmp_Ipv4_RateLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set unreachable ICMP packets ratelimit.
    Unreachable Icmp_Ipv4_RateLimit_Unreachable
}

func (rateLimit *Icmp_Ipv4_RateLimit) GetFilter() yfilter.YFilter { return rateLimit.YFilter }

func (rateLimit *Icmp_Ipv4_RateLimit) SetFilter(yf yfilter.YFilter) { rateLimit.YFilter = yf }

func (rateLimit *Icmp_Ipv4_RateLimit) GetGoName(yname string) string {
    if yname == "unreachable" { return "Unreachable" }
    return ""
}

func (rateLimit *Icmp_Ipv4_RateLimit) GetSegmentPath() string {
    return "rate-limit"
}

func (rateLimit *Icmp_Ipv4_RateLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unreachable" {
        return &rateLimit.Unreachable
    }
    return nil
}

func (rateLimit *Icmp_Ipv4_RateLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["unreachable"] = &rateLimit.Unreachable
    return children
}

func (rateLimit *Icmp_Ipv4_RateLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rateLimit *Icmp_Ipv4_RateLimit) GetBundleName() string { return "cisco_ios_xr" }

func (rateLimit *Icmp_Ipv4_RateLimit) GetYangName() string { return "rate-limit" }

func (rateLimit *Icmp_Ipv4_RateLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rateLimit *Icmp_Ipv4_RateLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rateLimit *Icmp_Ipv4_RateLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rateLimit *Icmp_Ipv4_RateLimit) SetParent(parent types.Entity) { rateLimit.parent = parent }

func (rateLimit *Icmp_Ipv4_RateLimit) GetParent() types.Entity { return rateLimit.parent }

func (rateLimit *Icmp_Ipv4_RateLimit) GetParentYangName() string { return "ipv4" }

// Icmp_Ipv4_RateLimit_Unreachable
// Set unreachable ICMP packets ratelimit
type Icmp_Ipv4_RateLimit_Unreachable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rate Limit of Unreachable ICMP packets. The type is interface{} with range:
    // 0..4294967295.
    Rate interface{}

    // Rate Limit of Unreachable DF packets. The type is interface{} with range:
    // 0..4294967295.
    Fragmentation interface{}
}

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetFilter() yfilter.YFilter { return unreachable.YFilter }

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) SetFilter(yf yfilter.YFilter) { unreachable.YFilter = yf }

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetGoName(yname string) string {
    if yname == "rate" { return "Rate" }
    if yname == "fragmentation" { return "Fragmentation" }
    return ""
}

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetSegmentPath() string {
    return "unreachable"
}

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rate"] = unreachable.Rate
    leafs["fragmentation"] = unreachable.Fragmentation
    return leafs
}

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetBundleName() string { return "cisco_ios_xr" }

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetYangName() string { return "unreachable" }

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) SetParent(parent types.Entity) { unreachable.parent = parent }

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetParent() types.Entity { return unreachable.parent }

func (unreachable *Icmp_Ipv4_RateLimit_Unreachable) GetParentYangName() string { return "rate-limit" }

// Icmp_Ipv4_Source
// IP ICMP Source Address Selection Policy
type Icmp_Ipv4_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Source Address Policy. The type is SourcePolicy.
    SourceAddressPolicy interface{}
}

func (source *Icmp_Ipv4_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *Icmp_Ipv4_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *Icmp_Ipv4_Source) GetGoName(yname string) string {
    if yname == "source-address-policy" { return "SourceAddressPolicy" }
    return ""
}

func (source *Icmp_Ipv4_Source) GetSegmentPath() string {
    return "source"
}

func (source *Icmp_Ipv4_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (source *Icmp_Ipv4_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (source *Icmp_Ipv4_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address-policy"] = source.SourceAddressPolicy
    return leafs
}

func (source *Icmp_Ipv4_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *Icmp_Ipv4_Source) GetYangName() string { return "source" }

func (source *Icmp_Ipv4_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *Icmp_Ipv4_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *Icmp_Ipv4_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *Icmp_Ipv4_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *Icmp_Ipv4_Source) GetParent() types.Entity { return source.parent }

func (source *Icmp_Ipv4_Source) GetParentYangName() string { return "ipv4" }

