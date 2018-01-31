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
    parent types.Entity
    YFilter yfilter.YFilter

    // Set maximum depth for route recursion check. The type is interface{} with
    // range: 5..16.
    MaxRecursionDepth interface{}

    // RIB address family configuration.
    Af Rib_Af
}

func (rib *Rib) GetFilter() yfilter.YFilter { return rib.YFilter }

func (rib *Rib) SetFilter(yf yfilter.YFilter) { rib.YFilter = yf }

func (rib *Rib) GetGoName(yname string) string {
    if yname == "max-recursion-depth" { return "MaxRecursionDepth" }
    if yname == "af" { return "Af" }
    return ""
}

func (rib *Rib) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-rib-cfg:rib"
}

func (rib *Rib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        return &rib.Af
    }
    return nil
}

func (rib *Rib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["af"] = &rib.Af
    return children
}

func (rib *Rib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-recursion-depth"] = rib.MaxRecursionDepth
    return leafs
}

func (rib *Rib) GetBundleName() string { return "cisco_ios_xr" }

func (rib *Rib) GetYangName() string { return "rib" }

func (rib *Rib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rib *Rib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rib *Rib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rib *Rib) SetParent(parent types.Entity) { rib.parent = parent }

func (rib *Rib) GetParent() types.Entity { return rib.parent }

func (rib *Rib) GetParentYangName() string { return "Cisco-IOS-XR-ip-rib-cfg" }

// Rib_Af
// RIB address family configuration
type Rib_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 configuration.
    Ipv4 Rib_Af_Ipv4

    // IPv6 configuration.
    Ipv6 Rib_Af_Ipv6
}

func (af *Rib_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *Rib_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *Rib_Af) GetGoName(yname string) string {
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (af *Rib_Af) GetSegmentPath() string {
    return "af"
}

func (af *Rib_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4" {
        return &af.Ipv4
    }
    if childYangName == "ipv6" {
        return &af.Ipv6
    }
    return nil
}

func (af *Rib_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4"] = &af.Ipv4
    children["ipv6"] = &af.Ipv6
    return children
}

func (af *Rib_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (af *Rib_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *Rib_Af) GetYangName() string { return "af" }

func (af *Rib_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *Rib_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *Rib_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *Rib_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *Rib_Af) GetParent() types.Entity { return af.parent }

func (af *Rib_Af) GetParentYangName() string { return "rib" }

// Rib_Af_Ipv4
// IPv4 configuration
type Rib_Af_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable next-hop dampening. The type is interface{}.
    NextHopDampeningDisable interface{}

    // Redistribution history related configs.
    RedistributionHistory Rib_Af_Ipv4_RedistributionHistory
}

func (ipv4 *Rib_Af_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Rib_Af_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Rib_Af_Ipv4) GetGoName(yname string) string {
    if yname == "next-hop-dampening-disable" { return "NextHopDampeningDisable" }
    if yname == "redistribution-history" { return "RedistributionHistory" }
    return ""
}

func (ipv4 *Rib_Af_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Rib_Af_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "redistribution-history" {
        return &ipv4.RedistributionHistory
    }
    return nil
}

func (ipv4 *Rib_Af_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["redistribution-history"] = &ipv4.RedistributionHistory
    return children
}

func (ipv4 *Rib_Af_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop-dampening-disable"] = ipv4.NextHopDampeningDisable
    return leafs
}

func (ipv4 *Rib_Af_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Rib_Af_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Rib_Af_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Rib_Af_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Rib_Af_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Rib_Af_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Rib_Af_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Rib_Af_Ipv4) GetParentYangName() string { return "af" }

// Rib_Af_Ipv4_RedistributionHistory
// Redistribution history related configs
type Rib_Af_Ipv4_RedistributionHistory struct {
    parent types.Entity
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

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetFilter() yfilter.YFilter { return redistributionHistory.YFilter }

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) SetFilter(yf yfilter.YFilter) { redistributionHistory.YFilter = yf }

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetGoName(yname string) string {
    if yname == "bcdl-client" { return "BcdlClient" }
    if yname == "protocol-client" { return "ProtocolClient" }
    if yname == "keep" { return "Keep" }
    return ""
}

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetSegmentPath() string {
    return "redistribution-history"
}

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "keep" {
        return &redistributionHistory.Keep
    }
    return nil
}

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["keep"] = &redistributionHistory.Keep
    return children
}

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bcdl-client"] = redistributionHistory.BcdlClient
    leafs["protocol-client"] = redistributionHistory.ProtocolClient
    return leafs
}

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetBundleName() string { return "cisco_ios_xr" }

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetYangName() string { return "redistribution-history" }

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) SetParent(parent types.Entity) { redistributionHistory.parent = parent }

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetParent() types.Entity { return redistributionHistory.parent }

func (redistributionHistory *Rib_Af_Ipv4_RedistributionHistory) GetParentYangName() string { return "ipv4" }

// Rib_Af_Ipv4_RedistributionHistory_Keep
// Retain redistribution history after disconnect.
type Rib_Af_Ipv4_RedistributionHistory_Keep struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable retain BCDL history. The type is interface{}.
    Bcdl interface{}
}

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetFilter() yfilter.YFilter { return keep.YFilter }

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) SetFilter(yf yfilter.YFilter) { keep.YFilter = yf }

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetGoName(yname string) string {
    if yname == "bcdl" { return "Bcdl" }
    return ""
}

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetSegmentPath() string {
    return "keep"
}

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bcdl"] = keep.Bcdl
    return leafs
}

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetBundleName() string { return "cisco_ios_xr" }

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetYangName() string { return "keep" }

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) SetParent(parent types.Entity) { keep.parent = parent }

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetParent() types.Entity { return keep.parent }

func (keep *Rib_Af_Ipv4_RedistributionHistory_Keep) GetParentYangName() string { return "redistribution-history" }

// Rib_Af_Ipv6
// IPv6 configuration
type Rib_Af_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable next-hop dampening. The type is interface{}.
    NextHopDampeningDisable interface{}

    // Redistribution history related configs.
    RedistributionHistory Rib_Af_Ipv6_RedistributionHistory
}

func (ipv6 *Rib_Af_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Rib_Af_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Rib_Af_Ipv6) GetGoName(yname string) string {
    if yname == "next-hop-dampening-disable" { return "NextHopDampeningDisable" }
    if yname == "redistribution-history" { return "RedistributionHistory" }
    return ""
}

func (ipv6 *Rib_Af_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Rib_Af_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "redistribution-history" {
        return &ipv6.RedistributionHistory
    }
    return nil
}

func (ipv6 *Rib_Af_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["redistribution-history"] = &ipv6.RedistributionHistory
    return children
}

func (ipv6 *Rib_Af_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop-dampening-disable"] = ipv6.NextHopDampeningDisable
    return leafs
}

func (ipv6 *Rib_Af_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Rib_Af_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Rib_Af_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Rib_Af_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Rib_Af_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Rib_Af_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Rib_Af_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Rib_Af_Ipv6) GetParentYangName() string { return "af" }

// Rib_Af_Ipv6_RedistributionHistory
// Redistribution history related configs
type Rib_Af_Ipv6_RedistributionHistory struct {
    parent types.Entity
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

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetFilter() yfilter.YFilter { return redistributionHistory.YFilter }

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) SetFilter(yf yfilter.YFilter) { redistributionHistory.YFilter = yf }

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetGoName(yname string) string {
    if yname == "bcdl-client" { return "BcdlClient" }
    if yname == "protocol-client" { return "ProtocolClient" }
    if yname == "keep" { return "Keep" }
    return ""
}

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetSegmentPath() string {
    return "redistribution-history"
}

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "keep" {
        return &redistributionHistory.Keep
    }
    return nil
}

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["keep"] = &redistributionHistory.Keep
    return children
}

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bcdl-client"] = redistributionHistory.BcdlClient
    leafs["protocol-client"] = redistributionHistory.ProtocolClient
    return leafs
}

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetBundleName() string { return "cisco_ios_xr" }

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetYangName() string { return "redistribution-history" }

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) SetParent(parent types.Entity) { redistributionHistory.parent = parent }

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetParent() types.Entity { return redistributionHistory.parent }

func (redistributionHistory *Rib_Af_Ipv6_RedistributionHistory) GetParentYangName() string { return "ipv6" }

// Rib_Af_Ipv6_RedistributionHistory_Keep
// Retain redistribution history after disconnect.
type Rib_Af_Ipv6_RedistributionHistory_Keep struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable retain BCDL history. The type is interface{}.
    Bcdl interface{}
}

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetFilter() yfilter.YFilter { return keep.YFilter }

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) SetFilter(yf yfilter.YFilter) { keep.YFilter = yf }

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetGoName(yname string) string {
    if yname == "bcdl" { return "Bcdl" }
    return ""
}

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetSegmentPath() string {
    return "keep"
}

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bcdl"] = keep.Bcdl
    return leafs
}

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetBundleName() string { return "cisco_ios_xr" }

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetYangName() string { return "keep" }

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) SetParent(parent types.Entity) { keep.parent = parent }

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetParent() types.Entity { return keep.parent }

func (keep *Rib_Af_Ipv6_RedistributionHistory_Keep) GetParentYangName() string { return "redistribution-history" }

