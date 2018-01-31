// This module contains a collection of YANG definitions
// for Cisco IOS-XR segment-routing-ms package operational data.
// 
// This module contains definitions
// for the following management objects:
//   srms: Segment Routing Mapping Server operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package segment_routing_ms_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package segment_routing_ms_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-segment-routing-ms-oper srms}", reflect.TypeOf(Srms{}))
    ydk.RegisterEntity("Cisco-IOS-XR-segment-routing-ms-oper:srms", reflect.TypeOf(Srms{}))
}

// SrmsMiAfEB represents Srms mi af e b
type SrmsMiAfEB string

const (
    // None
    SrmsMiAfEB_none SrmsMiAfEB = "none"

    // IPv4
    SrmsMiAfEB_ipv4 SrmsMiAfEB = "ipv4"

    // IPv6
    SrmsMiAfEB_ipv6 SrmsMiAfEB = "ipv6"
)

// SrmsMiFlagEB represents Srms mi flag e b
type SrmsMiFlagEB string

const (
    // False
    SrmsMiFlagEB_false SrmsMiFlagEB = "false"

    // True
    SrmsMiFlagEB_true SrmsMiFlagEB = "true"
)

// SrmsMiSrcEB represents Srms mi src e b
type SrmsMiSrcEB string

const (
    // None
    SrmsMiSrcEB_none SrmsMiSrcEB = "none"

    // Local
    SrmsMiSrcEB_local SrmsMiSrcEB = "local"

    // Remote
    SrmsMiSrcEB_remote SrmsMiSrcEB = "remote"
)

// Srms
// Segment Routing Mapping Server operational data
type Srms struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP prefix to SID mappings.
    Mapping Srms_Mapping

    // Policy operational data.
    Policy Srms_Policy
}

func (srms *Srms) GetFilter() yfilter.YFilter { return srms.YFilter }

func (srms *Srms) SetFilter(yf yfilter.YFilter) { srms.YFilter = yf }

func (srms *Srms) GetGoName(yname string) string {
    if yname == "mapping" { return "Mapping" }
    if yname == "policy" { return "Policy" }
    return ""
}

func (srms *Srms) GetSegmentPath() string {
    return "Cisco-IOS-XR-segment-routing-ms-oper:srms"
}

func (srms *Srms) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mapping" {
        return &srms.Mapping
    }
    if childYangName == "policy" {
        return &srms.Policy
    }
    return nil
}

func (srms *Srms) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mapping"] = &srms.Mapping
    children["policy"] = &srms.Policy
    return children
}

func (srms *Srms) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (srms *Srms) GetBundleName() string { return "cisco_ios_xr" }

func (srms *Srms) GetYangName() string { return "srms" }

func (srms *Srms) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srms *Srms) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srms *Srms) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srms *Srms) SetParent(parent types.Entity) { srms.parent = parent }

func (srms *Srms) GetParent() types.Entity { return srms.parent }

func (srms *Srms) GetParentYangName() string { return "Cisco-IOS-XR-segment-routing-ms-oper" }

// Srms_Mapping
// IP prefix to SID mappings
type Srms_Mapping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 prefix to SID mappings.
    MappingIpv4 Srms_Mapping_MappingIpv4

    // IPv6 prefix to SID mappings.
    MappingIpv6 Srms_Mapping_MappingIpv6
}

func (mapping *Srms_Mapping) GetFilter() yfilter.YFilter { return mapping.YFilter }

func (mapping *Srms_Mapping) SetFilter(yf yfilter.YFilter) { mapping.YFilter = yf }

func (mapping *Srms_Mapping) GetGoName(yname string) string {
    if yname == "mapping-ipv4" { return "MappingIpv4" }
    if yname == "mapping-ipv6" { return "MappingIpv6" }
    return ""
}

func (mapping *Srms_Mapping) GetSegmentPath() string {
    return "mapping"
}

func (mapping *Srms_Mapping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mapping-ipv4" {
        return &mapping.MappingIpv4
    }
    if childYangName == "mapping-ipv6" {
        return &mapping.MappingIpv6
    }
    return nil
}

func (mapping *Srms_Mapping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mapping-ipv4"] = &mapping.MappingIpv4
    children["mapping-ipv6"] = &mapping.MappingIpv6
    return children
}

func (mapping *Srms_Mapping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mapping *Srms_Mapping) GetBundleName() string { return "cisco_ios_xr" }

func (mapping *Srms_Mapping) GetYangName() string { return "mapping" }

func (mapping *Srms_Mapping) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mapping *Srms_Mapping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mapping *Srms_Mapping) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mapping *Srms_Mapping) SetParent(parent types.Entity) { mapping.parent = parent }

func (mapping *Srms_Mapping) GetParent() types.Entity { return mapping.parent }

func (mapping *Srms_Mapping) GetParentYangName() string { return "srms" }

// Srms_Mapping_MappingIpv4
// IPv4 prefix to SID mappings
type Srms_Mapping_MappingIpv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP prefix to SID mapping item. It's not possible to list all of the IP
    // prefix to SID mappings, as the set of valid prefixes could be very large.
    // Instead, SID map information must be retrieved individually for each prefix
    // of interest. The type is slice of Srms_Mapping_MappingIpv4_MappingMi.
    MappingMi []Srms_Mapping_MappingIpv4_MappingMi
}

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetFilter() yfilter.YFilter { return mappingIpv4.YFilter }

func (mappingIpv4 *Srms_Mapping_MappingIpv4) SetFilter(yf yfilter.YFilter) { mappingIpv4.YFilter = yf }

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetGoName(yname string) string {
    if yname == "mapping-mi" { return "MappingMi" }
    return ""
}

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetSegmentPath() string {
    return "mapping-ipv4"
}

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mapping-mi" {
        for _, c := range mappingIpv4.MappingMi {
            if mappingIpv4.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srms_Mapping_MappingIpv4_MappingMi{}
        mappingIpv4.MappingMi = append(mappingIpv4.MappingMi, child)
        return &mappingIpv4.MappingMi[len(mappingIpv4.MappingMi)-1]
    }
    return nil
}

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mappingIpv4.MappingMi {
        children[mappingIpv4.MappingMi[i].GetSegmentPath()] = &mappingIpv4.MappingMi[i]
    }
    return children
}

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetBundleName() string { return "cisco_ios_xr" }

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetYangName() string { return "mapping-ipv4" }

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mappingIpv4 *Srms_Mapping_MappingIpv4) SetParent(parent types.Entity) { mappingIpv4.parent = parent }

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetParent() types.Entity { return mappingIpv4.parent }

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetParentYangName() string { return "mapping" }

// Srms_Mapping_MappingIpv4_MappingMi
// IP prefix to SID mapping item. It's not possible
// to list all of the IP prefix to SID mappings, as
// the set of valid prefixes could be very large.
// Instead, SID map information must be retrieved
// individually for each prefix of interest.
type Srms_Mapping_MappingIpv4_MappingMi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Ip interface{}

    // Prefix. The type is interface{} with range: -2147483648..2147483647.
    Prefix interface{}

    // src. The type is SrmsMiSrcEB.
    Src interface{}

    // Router ID. The type is string with length: 0..30.
    Router interface{}

    // Area (OSPF) or Level (ISIS). The type is string with length: 0..30.
    Area interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixXr interface{}

    // Starting SID. The type is interface{} with range: 0..4294967295.
    SidStart interface{}

    // SID range. The type is interface{} with range: 0..4294967295.
    SidCount interface{}

    // Last IP Prefix. The type is string with length: 0..50.
    LastPrefix interface{}

    // Last SID Index. The type is interface{} with range: 0..4294967295.
    LastSidIndex interface{}

    // Attached flag. The type is SrmsMiFlagEB.
    FlagAttached interface{}

    // addr.
    Addr Srms_Mapping_MappingIpv4_MappingMi_Addr
}

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetFilter() yfilter.YFilter { return mappingMi.YFilter }

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) SetFilter(yf yfilter.YFilter) { mappingMi.YFilter = yf }

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix" { return "Prefix" }
    if yname == "src" { return "Src" }
    if yname == "router" { return "Router" }
    if yname == "area" { return "Area" }
    if yname == "prefix-xr" { return "PrefixXr" }
    if yname == "sid-start" { return "SidStart" }
    if yname == "sid-count" { return "SidCount" }
    if yname == "last-prefix" { return "LastPrefix" }
    if yname == "last-sid-index" { return "LastSidIndex" }
    if yname == "flag-attached" { return "FlagAttached" }
    if yname == "addr" { return "Addr" }
    return ""
}

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetSegmentPath() string {
    return "mapping-mi"
}

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addr" {
        return &mappingMi.Addr
    }
    return nil
}

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addr"] = &mappingMi.Addr
    return children
}

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = mappingMi.Ip
    leafs["prefix"] = mappingMi.Prefix
    leafs["src"] = mappingMi.Src
    leafs["router"] = mappingMi.Router
    leafs["area"] = mappingMi.Area
    leafs["prefix-xr"] = mappingMi.PrefixXr
    leafs["sid-start"] = mappingMi.SidStart
    leafs["sid-count"] = mappingMi.SidCount
    leafs["last-prefix"] = mappingMi.LastPrefix
    leafs["last-sid-index"] = mappingMi.LastSidIndex
    leafs["flag-attached"] = mappingMi.FlagAttached
    return leafs
}

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetBundleName() string { return "cisco_ios_xr" }

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetYangName() string { return "mapping-mi" }

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) SetParent(parent types.Entity) { mappingMi.parent = parent }

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetParent() types.Entity { return mappingMi.parent }

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetParentYangName() string { return "mapping-ipv4" }

// Srms_Mapping_MappingIpv4_MappingMi_Addr
// addr
type Srms_Mapping_MappingIpv4_MappingMi_Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AF. The type is SrmsMiAfEB.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetFilter() yfilter.YFilter { return addr.YFilter }

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) SetFilter(yf yfilter.YFilter) { addr.YFilter = yf }

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetSegmentPath() string {
    return "addr"
}

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af"] = addr.Af
    leafs["ipv4"] = addr.Ipv4
    leafs["ipv6"] = addr.Ipv6
    return leafs
}

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetBundleName() string { return "cisco_ios_xr" }

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetYangName() string { return "addr" }

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) SetParent(parent types.Entity) { addr.parent = parent }

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetParent() types.Entity { return addr.parent }

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetParentYangName() string { return "mapping-mi" }

// Srms_Mapping_MappingIpv6
// IPv6 prefix to SID mappings
type Srms_Mapping_MappingIpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP prefix to SID mapping item. It's not possible to list all of the IP
    // prefix to SID mappings, as the set of valid prefixes could be very large.
    // Instead, SID map information must be retrieved individually for each prefix
    // of interest. The type is slice of Srms_Mapping_MappingIpv6_MappingMi.
    MappingMi []Srms_Mapping_MappingIpv6_MappingMi
}

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetFilter() yfilter.YFilter { return mappingIpv6.YFilter }

func (mappingIpv6 *Srms_Mapping_MappingIpv6) SetFilter(yf yfilter.YFilter) { mappingIpv6.YFilter = yf }

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetGoName(yname string) string {
    if yname == "mapping-mi" { return "MappingMi" }
    return ""
}

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetSegmentPath() string {
    return "mapping-ipv6"
}

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mapping-mi" {
        for _, c := range mappingIpv6.MappingMi {
            if mappingIpv6.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srms_Mapping_MappingIpv6_MappingMi{}
        mappingIpv6.MappingMi = append(mappingIpv6.MappingMi, child)
        return &mappingIpv6.MappingMi[len(mappingIpv6.MappingMi)-1]
    }
    return nil
}

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mappingIpv6.MappingMi {
        children[mappingIpv6.MappingMi[i].GetSegmentPath()] = &mappingIpv6.MappingMi[i]
    }
    return children
}

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetBundleName() string { return "cisco_ios_xr" }

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetYangName() string { return "mapping-ipv6" }

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mappingIpv6 *Srms_Mapping_MappingIpv6) SetParent(parent types.Entity) { mappingIpv6.parent = parent }

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetParent() types.Entity { return mappingIpv6.parent }

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetParentYangName() string { return "mapping" }

// Srms_Mapping_MappingIpv6_MappingMi
// IP prefix to SID mapping item. It's not possible
// to list all of the IP prefix to SID mappings, as
// the set of valid prefixes could be very large.
// Instead, SID map information must be retrieved
// individually for each prefix of interest.
type Srms_Mapping_MappingIpv6_MappingMi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Ip interface{}

    // Prefix. The type is interface{} with range: -2147483648..2147483647.
    Prefix interface{}

    // src. The type is SrmsMiSrcEB.
    Src interface{}

    // Router ID. The type is string with length: 0..30.
    Router interface{}

    // Area (OSPF) or Level (ISIS). The type is string with length: 0..30.
    Area interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixXr interface{}

    // Starting SID. The type is interface{} with range: 0..4294967295.
    SidStart interface{}

    // SID range. The type is interface{} with range: 0..4294967295.
    SidCount interface{}

    // Last IP Prefix. The type is string with length: 0..50.
    LastPrefix interface{}

    // Last SID Index. The type is interface{} with range: 0..4294967295.
    LastSidIndex interface{}

    // Attached flag. The type is SrmsMiFlagEB.
    FlagAttached interface{}

    // addr.
    Addr Srms_Mapping_MappingIpv6_MappingMi_Addr
}

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetFilter() yfilter.YFilter { return mappingMi.YFilter }

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) SetFilter(yf yfilter.YFilter) { mappingMi.YFilter = yf }

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix" { return "Prefix" }
    if yname == "src" { return "Src" }
    if yname == "router" { return "Router" }
    if yname == "area" { return "Area" }
    if yname == "prefix-xr" { return "PrefixXr" }
    if yname == "sid-start" { return "SidStart" }
    if yname == "sid-count" { return "SidCount" }
    if yname == "last-prefix" { return "LastPrefix" }
    if yname == "last-sid-index" { return "LastSidIndex" }
    if yname == "flag-attached" { return "FlagAttached" }
    if yname == "addr" { return "Addr" }
    return ""
}

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetSegmentPath() string {
    return "mapping-mi"
}

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addr" {
        return &mappingMi.Addr
    }
    return nil
}

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addr"] = &mappingMi.Addr
    return children
}

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = mappingMi.Ip
    leafs["prefix"] = mappingMi.Prefix
    leafs["src"] = mappingMi.Src
    leafs["router"] = mappingMi.Router
    leafs["area"] = mappingMi.Area
    leafs["prefix-xr"] = mappingMi.PrefixXr
    leafs["sid-start"] = mappingMi.SidStart
    leafs["sid-count"] = mappingMi.SidCount
    leafs["last-prefix"] = mappingMi.LastPrefix
    leafs["last-sid-index"] = mappingMi.LastSidIndex
    leafs["flag-attached"] = mappingMi.FlagAttached
    return leafs
}

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetBundleName() string { return "cisco_ios_xr" }

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetYangName() string { return "mapping-mi" }

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) SetParent(parent types.Entity) { mappingMi.parent = parent }

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetParent() types.Entity { return mappingMi.parent }

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetParentYangName() string { return "mapping-ipv6" }

// Srms_Mapping_MappingIpv6_MappingMi_Addr
// addr
type Srms_Mapping_MappingIpv6_MappingMi_Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AF. The type is SrmsMiAfEB.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetFilter() yfilter.YFilter { return addr.YFilter }

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) SetFilter(yf yfilter.YFilter) { addr.YFilter = yf }

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetSegmentPath() string {
    return "addr"
}

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af"] = addr.Af
    leafs["ipv4"] = addr.Ipv4
    leafs["ipv6"] = addr.Ipv6
    return leafs
}

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetBundleName() string { return "cisco_ios_xr" }

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetYangName() string { return "addr" }

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) SetParent(parent types.Entity) { addr.parent = parent }

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetParent() types.Entity { return addr.parent }

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetParentYangName() string { return "mapping-mi" }

// Srms_Policy
// Policy operational data
type Srms_Policy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 policy operational data.
    PolicyIpv4 Srms_Policy_PolicyIpv4

    // IPv6 policy operational data.
    PolicyIpv6 Srms_Policy_PolicyIpv6
}

func (policy *Srms_Policy) GetFilter() yfilter.YFilter { return policy.YFilter }

func (policy *Srms_Policy) SetFilter(yf yfilter.YFilter) { policy.YFilter = yf }

func (policy *Srms_Policy) GetGoName(yname string) string {
    if yname == "policy-ipv4" { return "PolicyIpv4" }
    if yname == "policy-ipv6" { return "PolicyIpv6" }
    return ""
}

func (policy *Srms_Policy) GetSegmentPath() string {
    return "policy"
}

func (policy *Srms_Policy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-ipv4" {
        return &policy.PolicyIpv4
    }
    if childYangName == "policy-ipv6" {
        return &policy.PolicyIpv6
    }
    return nil
}

func (policy *Srms_Policy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["policy-ipv4"] = &policy.PolicyIpv4
    children["policy-ipv6"] = &policy.PolicyIpv6
    return children
}

func (policy *Srms_Policy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policy *Srms_Policy) GetBundleName() string { return "cisco_ios_xr" }

func (policy *Srms_Policy) GetYangName() string { return "policy" }

func (policy *Srms_Policy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policy *Srms_Policy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policy *Srms_Policy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policy *Srms_Policy) SetParent(parent types.Entity) { policy.parent = parent }

func (policy *Srms_Policy) GetParent() types.Entity { return policy.parent }

func (policy *Srms_Policy) GetParentYangName() string { return "srms" }

// Srms_Policy_PolicyIpv4
// IPv4 policy operational data
type Srms_Policy_PolicyIpv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 backup policy operational data.
    PolicyIpv4Backup Srms_Policy_PolicyIpv4_PolicyIpv4Backup

    // IPv4 active policy operational data.
    PolicyIpv4Active Srms_Policy_PolicyIpv4_PolicyIpv4Active
}

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetFilter() yfilter.YFilter { return policyIpv4.YFilter }

func (policyIpv4 *Srms_Policy_PolicyIpv4) SetFilter(yf yfilter.YFilter) { policyIpv4.YFilter = yf }

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetGoName(yname string) string {
    if yname == "policy-ipv4-backup" { return "PolicyIpv4Backup" }
    if yname == "policy-ipv4-active" { return "PolicyIpv4Active" }
    return ""
}

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetSegmentPath() string {
    return "policy-ipv4"
}

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-ipv4-backup" {
        return &policyIpv4.PolicyIpv4Backup
    }
    if childYangName == "policy-ipv4-active" {
        return &policyIpv4.PolicyIpv4Active
    }
    return nil
}

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["policy-ipv4-backup"] = &policyIpv4.PolicyIpv4Backup
    children["policy-ipv4-active"] = &policyIpv4.PolicyIpv4Active
    return children
}

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetBundleName() string { return "cisco_ios_xr" }

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetYangName() string { return "policy-ipv4" }

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyIpv4 *Srms_Policy_PolicyIpv4) SetParent(parent types.Entity) { policyIpv4.parent = parent }

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetParent() types.Entity { return policyIpv4.parent }

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetParentYangName() string { return "policy" }

// Srms_Policy_PolicyIpv4_PolicyIpv4Backup
// IPv4 backup policy operational data
type Srms_Policy_PolicyIpv4_PolicyIpv4Backup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mapping Item. The type is slice of
    // Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi.
    PolicyMi []Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi
}

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetFilter() yfilter.YFilter { return policyIpv4Backup.YFilter }

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) SetFilter(yf yfilter.YFilter) { policyIpv4Backup.YFilter = yf }

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetGoName(yname string) string {
    if yname == "policy-mi" { return "PolicyMi" }
    return ""
}

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetSegmentPath() string {
    return "policy-ipv4-backup"
}

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-mi" {
        for _, c := range policyIpv4Backup.PolicyMi {
            if policyIpv4Backup.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi{}
        policyIpv4Backup.PolicyMi = append(policyIpv4Backup.PolicyMi, child)
        return &policyIpv4Backup.PolicyMi[len(policyIpv4Backup.PolicyMi)-1]
    }
    return nil
}

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policyIpv4Backup.PolicyMi {
        children[policyIpv4Backup.PolicyMi[i].GetSegmentPath()] = &policyIpv4Backup.PolicyMi[i]
    }
    return children
}

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetBundleName() string { return "cisco_ios_xr" }

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetYangName() string { return "policy-ipv4-backup" }

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) SetParent(parent types.Entity) { policyIpv4Backup.parent = parent }

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetParent() types.Entity { return policyIpv4Backup.parent }

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetParentYangName() string { return "policy-ipv4" }

// Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi
// Mapping Item
type Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Mapping Item ID (0, 1, 2, ...). The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    MiId interface{}

    // src. The type is SrmsMiSrcEB.
    Src interface{}

    // Router ID. The type is string with length: 0..30.
    Router interface{}

    // Area (OSPF) or Level (ISIS). The type is string with length: 0..30.
    Area interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixXr interface{}

    // Starting SID. The type is interface{} with range: 0..4294967295.
    SidStart interface{}

    // SID range. The type is interface{} with range: 0..4294967295.
    SidCount interface{}

    // Last IP Prefix. The type is string with length: 0..50.
    LastPrefix interface{}

    // Last SID Index. The type is interface{} with range: 0..4294967295.
    LastSidIndex interface{}

    // Attached flag. The type is SrmsMiFlagEB.
    FlagAttached interface{}

    // addr.
    Addr Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetFilter() yfilter.YFilter { return policyMi.YFilter }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) SetFilter(yf yfilter.YFilter) { policyMi.YFilter = yf }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetGoName(yname string) string {
    if yname == "mi-id" { return "MiId" }
    if yname == "src" { return "Src" }
    if yname == "router" { return "Router" }
    if yname == "area" { return "Area" }
    if yname == "prefix-xr" { return "PrefixXr" }
    if yname == "sid-start" { return "SidStart" }
    if yname == "sid-count" { return "SidCount" }
    if yname == "last-prefix" { return "LastPrefix" }
    if yname == "last-sid-index" { return "LastSidIndex" }
    if yname == "flag-attached" { return "FlagAttached" }
    if yname == "addr" { return "Addr" }
    return ""
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetSegmentPath() string {
    return "policy-mi" + "[mi-id='" + fmt.Sprintf("%v", policyMi.MiId) + "']"
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addr" {
        return &policyMi.Addr
    }
    return nil
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addr"] = &policyMi.Addr
    return children
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mi-id"] = policyMi.MiId
    leafs["src"] = policyMi.Src
    leafs["router"] = policyMi.Router
    leafs["area"] = policyMi.Area
    leafs["prefix-xr"] = policyMi.PrefixXr
    leafs["sid-start"] = policyMi.SidStart
    leafs["sid-count"] = policyMi.SidCount
    leafs["last-prefix"] = policyMi.LastPrefix
    leafs["last-sid-index"] = policyMi.LastSidIndex
    leafs["flag-attached"] = policyMi.FlagAttached
    return leafs
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetBundleName() string { return "cisco_ios_xr" }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetYangName() string { return "policy-mi" }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) SetParent(parent types.Entity) { policyMi.parent = parent }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetParent() types.Entity { return policyMi.parent }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetParentYangName() string { return "policy-ipv4-backup" }

// Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr
// addr
type Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AF. The type is SrmsMiAfEB.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetFilter() yfilter.YFilter { return addr.YFilter }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) SetFilter(yf yfilter.YFilter) { addr.YFilter = yf }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetSegmentPath() string {
    return "addr"
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af"] = addr.Af
    leafs["ipv4"] = addr.Ipv4
    leafs["ipv6"] = addr.Ipv6
    return leafs
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetBundleName() string { return "cisco_ios_xr" }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetYangName() string { return "addr" }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) SetParent(parent types.Entity) { addr.parent = parent }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetParent() types.Entity { return addr.parent }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetParentYangName() string { return "policy-mi" }

// Srms_Policy_PolicyIpv4_PolicyIpv4Active
// IPv4 active policy operational data
type Srms_Policy_PolicyIpv4_PolicyIpv4Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mapping Item. The type is slice of
    // Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi.
    PolicyMi []Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi
}

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetFilter() yfilter.YFilter { return policyIpv4Active.YFilter }

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) SetFilter(yf yfilter.YFilter) { policyIpv4Active.YFilter = yf }

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetGoName(yname string) string {
    if yname == "policy-mi" { return "PolicyMi" }
    return ""
}

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetSegmentPath() string {
    return "policy-ipv4-active"
}

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-mi" {
        for _, c := range policyIpv4Active.PolicyMi {
            if policyIpv4Active.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi{}
        policyIpv4Active.PolicyMi = append(policyIpv4Active.PolicyMi, child)
        return &policyIpv4Active.PolicyMi[len(policyIpv4Active.PolicyMi)-1]
    }
    return nil
}

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policyIpv4Active.PolicyMi {
        children[policyIpv4Active.PolicyMi[i].GetSegmentPath()] = &policyIpv4Active.PolicyMi[i]
    }
    return children
}

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetBundleName() string { return "cisco_ios_xr" }

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetYangName() string { return "policy-ipv4-active" }

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) SetParent(parent types.Entity) { policyIpv4Active.parent = parent }

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetParent() types.Entity { return policyIpv4Active.parent }

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetParentYangName() string { return "policy-ipv4" }

// Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi
// Mapping Item
type Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Mapping Item ID (0, 1, 2, ...). The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    MiId interface{}

    // src. The type is SrmsMiSrcEB.
    Src interface{}

    // Router ID. The type is string with length: 0..30.
    Router interface{}

    // Area (OSPF) or Level (ISIS). The type is string with length: 0..30.
    Area interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixXr interface{}

    // Starting SID. The type is interface{} with range: 0..4294967295.
    SidStart interface{}

    // SID range. The type is interface{} with range: 0..4294967295.
    SidCount interface{}

    // Last IP Prefix. The type is string with length: 0..50.
    LastPrefix interface{}

    // Last SID Index. The type is interface{} with range: 0..4294967295.
    LastSidIndex interface{}

    // Attached flag. The type is SrmsMiFlagEB.
    FlagAttached interface{}

    // addr.
    Addr Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetFilter() yfilter.YFilter { return policyMi.YFilter }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) SetFilter(yf yfilter.YFilter) { policyMi.YFilter = yf }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetGoName(yname string) string {
    if yname == "mi-id" { return "MiId" }
    if yname == "src" { return "Src" }
    if yname == "router" { return "Router" }
    if yname == "area" { return "Area" }
    if yname == "prefix-xr" { return "PrefixXr" }
    if yname == "sid-start" { return "SidStart" }
    if yname == "sid-count" { return "SidCount" }
    if yname == "last-prefix" { return "LastPrefix" }
    if yname == "last-sid-index" { return "LastSidIndex" }
    if yname == "flag-attached" { return "FlagAttached" }
    if yname == "addr" { return "Addr" }
    return ""
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetSegmentPath() string {
    return "policy-mi" + "[mi-id='" + fmt.Sprintf("%v", policyMi.MiId) + "']"
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addr" {
        return &policyMi.Addr
    }
    return nil
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addr"] = &policyMi.Addr
    return children
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mi-id"] = policyMi.MiId
    leafs["src"] = policyMi.Src
    leafs["router"] = policyMi.Router
    leafs["area"] = policyMi.Area
    leafs["prefix-xr"] = policyMi.PrefixXr
    leafs["sid-start"] = policyMi.SidStart
    leafs["sid-count"] = policyMi.SidCount
    leafs["last-prefix"] = policyMi.LastPrefix
    leafs["last-sid-index"] = policyMi.LastSidIndex
    leafs["flag-attached"] = policyMi.FlagAttached
    return leafs
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetBundleName() string { return "cisco_ios_xr" }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetYangName() string { return "policy-mi" }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) SetParent(parent types.Entity) { policyMi.parent = parent }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetParent() types.Entity { return policyMi.parent }

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetParentYangName() string { return "policy-ipv4-active" }

// Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr
// addr
type Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AF. The type is SrmsMiAfEB.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetFilter() yfilter.YFilter { return addr.YFilter }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) SetFilter(yf yfilter.YFilter) { addr.YFilter = yf }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetSegmentPath() string {
    return "addr"
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af"] = addr.Af
    leafs["ipv4"] = addr.Ipv4
    leafs["ipv6"] = addr.Ipv6
    return leafs
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetBundleName() string { return "cisco_ios_xr" }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetYangName() string { return "addr" }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) SetParent(parent types.Entity) { addr.parent = parent }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetParent() types.Entity { return addr.parent }

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetParentYangName() string { return "policy-mi" }

// Srms_Policy_PolicyIpv6
// IPv6 policy operational data
type Srms_Policy_PolicyIpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 backup policy operational data.
    PolicyIpv6Backup Srms_Policy_PolicyIpv6_PolicyIpv6Backup

    // IPv6 active policy operational data.
    PolicyIpv6Active Srms_Policy_PolicyIpv6_PolicyIpv6Active
}

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetFilter() yfilter.YFilter { return policyIpv6.YFilter }

func (policyIpv6 *Srms_Policy_PolicyIpv6) SetFilter(yf yfilter.YFilter) { policyIpv6.YFilter = yf }

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetGoName(yname string) string {
    if yname == "policy-ipv6-backup" { return "PolicyIpv6Backup" }
    if yname == "policy-ipv6-active" { return "PolicyIpv6Active" }
    return ""
}

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetSegmentPath() string {
    return "policy-ipv6"
}

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-ipv6-backup" {
        return &policyIpv6.PolicyIpv6Backup
    }
    if childYangName == "policy-ipv6-active" {
        return &policyIpv6.PolicyIpv6Active
    }
    return nil
}

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["policy-ipv6-backup"] = &policyIpv6.PolicyIpv6Backup
    children["policy-ipv6-active"] = &policyIpv6.PolicyIpv6Active
    return children
}

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetBundleName() string { return "cisco_ios_xr" }

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetYangName() string { return "policy-ipv6" }

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyIpv6 *Srms_Policy_PolicyIpv6) SetParent(parent types.Entity) { policyIpv6.parent = parent }

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetParent() types.Entity { return policyIpv6.parent }

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetParentYangName() string { return "policy" }

// Srms_Policy_PolicyIpv6_PolicyIpv6Backup
// IPv6 backup policy operational data
type Srms_Policy_PolicyIpv6_PolicyIpv6Backup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mapping Item. The type is slice of
    // Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi.
    PolicyMi []Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi
}

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetFilter() yfilter.YFilter { return policyIpv6Backup.YFilter }

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) SetFilter(yf yfilter.YFilter) { policyIpv6Backup.YFilter = yf }

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetGoName(yname string) string {
    if yname == "policy-mi" { return "PolicyMi" }
    return ""
}

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetSegmentPath() string {
    return "policy-ipv6-backup"
}

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-mi" {
        for _, c := range policyIpv6Backup.PolicyMi {
            if policyIpv6Backup.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi{}
        policyIpv6Backup.PolicyMi = append(policyIpv6Backup.PolicyMi, child)
        return &policyIpv6Backup.PolicyMi[len(policyIpv6Backup.PolicyMi)-1]
    }
    return nil
}

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policyIpv6Backup.PolicyMi {
        children[policyIpv6Backup.PolicyMi[i].GetSegmentPath()] = &policyIpv6Backup.PolicyMi[i]
    }
    return children
}

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetBundleName() string { return "cisco_ios_xr" }

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetYangName() string { return "policy-ipv6-backup" }

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) SetParent(parent types.Entity) { policyIpv6Backup.parent = parent }

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetParent() types.Entity { return policyIpv6Backup.parent }

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetParentYangName() string { return "policy-ipv6" }

// Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi
// Mapping Item
type Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Mapping Item ID (0, 1, 2, ...). The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    MiId interface{}

    // src. The type is SrmsMiSrcEB.
    Src interface{}

    // Router ID. The type is string with length: 0..30.
    Router interface{}

    // Area (OSPF) or Level (ISIS). The type is string with length: 0..30.
    Area interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixXr interface{}

    // Starting SID. The type is interface{} with range: 0..4294967295.
    SidStart interface{}

    // SID range. The type is interface{} with range: 0..4294967295.
    SidCount interface{}

    // Last IP Prefix. The type is string with length: 0..50.
    LastPrefix interface{}

    // Last SID Index. The type is interface{} with range: 0..4294967295.
    LastSidIndex interface{}

    // Attached flag. The type is SrmsMiFlagEB.
    FlagAttached interface{}

    // addr.
    Addr Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetFilter() yfilter.YFilter { return policyMi.YFilter }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) SetFilter(yf yfilter.YFilter) { policyMi.YFilter = yf }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetGoName(yname string) string {
    if yname == "mi-id" { return "MiId" }
    if yname == "src" { return "Src" }
    if yname == "router" { return "Router" }
    if yname == "area" { return "Area" }
    if yname == "prefix-xr" { return "PrefixXr" }
    if yname == "sid-start" { return "SidStart" }
    if yname == "sid-count" { return "SidCount" }
    if yname == "last-prefix" { return "LastPrefix" }
    if yname == "last-sid-index" { return "LastSidIndex" }
    if yname == "flag-attached" { return "FlagAttached" }
    if yname == "addr" { return "Addr" }
    return ""
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetSegmentPath() string {
    return "policy-mi" + "[mi-id='" + fmt.Sprintf("%v", policyMi.MiId) + "']"
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addr" {
        return &policyMi.Addr
    }
    return nil
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addr"] = &policyMi.Addr
    return children
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mi-id"] = policyMi.MiId
    leafs["src"] = policyMi.Src
    leafs["router"] = policyMi.Router
    leafs["area"] = policyMi.Area
    leafs["prefix-xr"] = policyMi.PrefixXr
    leafs["sid-start"] = policyMi.SidStart
    leafs["sid-count"] = policyMi.SidCount
    leafs["last-prefix"] = policyMi.LastPrefix
    leafs["last-sid-index"] = policyMi.LastSidIndex
    leafs["flag-attached"] = policyMi.FlagAttached
    return leafs
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetBundleName() string { return "cisco_ios_xr" }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetYangName() string { return "policy-mi" }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) SetParent(parent types.Entity) { policyMi.parent = parent }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetParent() types.Entity { return policyMi.parent }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetParentYangName() string { return "policy-ipv6-backup" }

// Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr
// addr
type Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AF. The type is SrmsMiAfEB.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetFilter() yfilter.YFilter { return addr.YFilter }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) SetFilter(yf yfilter.YFilter) { addr.YFilter = yf }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetSegmentPath() string {
    return "addr"
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af"] = addr.Af
    leafs["ipv4"] = addr.Ipv4
    leafs["ipv6"] = addr.Ipv6
    return leafs
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetBundleName() string { return "cisco_ios_xr" }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetYangName() string { return "addr" }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) SetParent(parent types.Entity) { addr.parent = parent }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetParent() types.Entity { return addr.parent }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetParentYangName() string { return "policy-mi" }

// Srms_Policy_PolicyIpv6_PolicyIpv6Active
// IPv6 active policy operational data
type Srms_Policy_PolicyIpv6_PolicyIpv6Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mapping Item. The type is slice of
    // Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi.
    PolicyMi []Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi
}

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetFilter() yfilter.YFilter { return policyIpv6Active.YFilter }

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) SetFilter(yf yfilter.YFilter) { policyIpv6Active.YFilter = yf }

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetGoName(yname string) string {
    if yname == "policy-mi" { return "PolicyMi" }
    return ""
}

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetSegmentPath() string {
    return "policy-ipv6-active"
}

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-mi" {
        for _, c := range policyIpv6Active.PolicyMi {
            if policyIpv6Active.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi{}
        policyIpv6Active.PolicyMi = append(policyIpv6Active.PolicyMi, child)
        return &policyIpv6Active.PolicyMi[len(policyIpv6Active.PolicyMi)-1]
    }
    return nil
}

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policyIpv6Active.PolicyMi {
        children[policyIpv6Active.PolicyMi[i].GetSegmentPath()] = &policyIpv6Active.PolicyMi[i]
    }
    return children
}

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetBundleName() string { return "cisco_ios_xr" }

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetYangName() string { return "policy-ipv6-active" }

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) SetParent(parent types.Entity) { policyIpv6Active.parent = parent }

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetParent() types.Entity { return policyIpv6Active.parent }

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetParentYangName() string { return "policy-ipv6" }

// Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi
// Mapping Item
type Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Mapping Item ID (0, 1, 2, ...). The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    MiId interface{}

    // src. The type is SrmsMiSrcEB.
    Src interface{}

    // Router ID. The type is string with length: 0..30.
    Router interface{}

    // Area (OSPF) or Level (ISIS). The type is string with length: 0..30.
    Area interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixXr interface{}

    // Starting SID. The type is interface{} with range: 0..4294967295.
    SidStart interface{}

    // SID range. The type is interface{} with range: 0..4294967295.
    SidCount interface{}

    // Last IP Prefix. The type is string with length: 0..50.
    LastPrefix interface{}

    // Last SID Index. The type is interface{} with range: 0..4294967295.
    LastSidIndex interface{}

    // Attached flag. The type is SrmsMiFlagEB.
    FlagAttached interface{}

    // addr.
    Addr Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetFilter() yfilter.YFilter { return policyMi.YFilter }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) SetFilter(yf yfilter.YFilter) { policyMi.YFilter = yf }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetGoName(yname string) string {
    if yname == "mi-id" { return "MiId" }
    if yname == "src" { return "Src" }
    if yname == "router" { return "Router" }
    if yname == "area" { return "Area" }
    if yname == "prefix-xr" { return "PrefixXr" }
    if yname == "sid-start" { return "SidStart" }
    if yname == "sid-count" { return "SidCount" }
    if yname == "last-prefix" { return "LastPrefix" }
    if yname == "last-sid-index" { return "LastSidIndex" }
    if yname == "flag-attached" { return "FlagAttached" }
    if yname == "addr" { return "Addr" }
    return ""
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetSegmentPath() string {
    return "policy-mi" + "[mi-id='" + fmt.Sprintf("%v", policyMi.MiId) + "']"
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addr" {
        return &policyMi.Addr
    }
    return nil
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addr"] = &policyMi.Addr
    return children
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mi-id"] = policyMi.MiId
    leafs["src"] = policyMi.Src
    leafs["router"] = policyMi.Router
    leafs["area"] = policyMi.Area
    leafs["prefix-xr"] = policyMi.PrefixXr
    leafs["sid-start"] = policyMi.SidStart
    leafs["sid-count"] = policyMi.SidCount
    leafs["last-prefix"] = policyMi.LastPrefix
    leafs["last-sid-index"] = policyMi.LastSidIndex
    leafs["flag-attached"] = policyMi.FlagAttached
    return leafs
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetBundleName() string { return "cisco_ios_xr" }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetYangName() string { return "policy-mi" }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) SetParent(parent types.Entity) { policyMi.parent = parent }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetParent() types.Entity { return policyMi.parent }

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetParentYangName() string { return "policy-ipv6-active" }

// Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr
// addr
type Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AF. The type is SrmsMiAfEB.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetFilter() yfilter.YFilter { return addr.YFilter }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) SetFilter(yf yfilter.YFilter) { addr.YFilter = yf }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetSegmentPath() string {
    return "addr"
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af"] = addr.Af
    leafs["ipv4"] = addr.Ipv4
    leafs["ipv6"] = addr.Ipv6
    return leafs
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetBundleName() string { return "cisco_ios_xr" }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetYangName() string { return "addr" }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) SetParent(parent types.Entity) { addr.parent = parent }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetParent() types.Entity { return addr.parent }

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetParentYangName() string { return "policy-mi" }

