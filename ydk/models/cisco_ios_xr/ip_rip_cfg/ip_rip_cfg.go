// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-rip package configuration.
// 
// This module contains definitions
// for the following management objects:
//   rip: RIP configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_rip_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_rip_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-rip-cfg rip}", reflect.TypeOf(Rip{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-rip-cfg:rip", reflect.TypeOf(Rip{}))
}

// RipAuthMode represents Rip auth mode
type RipAuthMode string

const (
    // Text mode
    RipAuthMode_text RipAuthMode = "text"

    // MD5 mode
    RipAuthMode_md5 RipAuthMode = "md5"
)

// IsisRedistRoute represents Isis redist route
type IsisRedistRoute string

const (
    // Level-1 routes only
    IsisRedistRoute_level1 IsisRedistRoute = "level1"

    // Level-1 routes only
    IsisRedistRoute_level2 IsisRedistRoute = "level2"

    // Level-1 and level-2 routes
    IsisRedistRoute_level1_and2 IsisRedistRoute = "level1-and2"
)

// DefaultInformationOption represents Default information option
type DefaultInformationOption string

const (
    // Always
    DefaultInformationOption_always DefaultInformationOption = "always"

    // Use route-policy
    DefaultInformationOption_policy DefaultInformationOption = "policy"
)

// BgpRedistRoute represents Bgp redist route
type BgpRedistRoute string

const (
    // All routes
    BgpRedistRoute_all BgpRedistRoute = "all"

    // Internal routes only
    BgpRedistRoute_internal BgpRedistRoute = "internal"

    // External routes only
    BgpRedistRoute_external BgpRedistRoute = "external"

    // Local routes only
    BgpRedistRoute_local BgpRedistRoute = "local"
)

// RipExtCommunity represents Rip ext community
type RipExtCommunity string

const (
    // AS:nn format
    RipExtCommunity_as RipExtCommunity = "as"

    // IPV4Address:nn format
    RipExtCommunity_ipv4_address RipExtCommunity = "ipv4-address"

    // 4-byte ASN format
    RipExtCommunity_four_byte_as RipExtCommunity = "four-byte-as"
)

// DefaultRedistRoute represents Default redist route
type DefaultRedistRoute string

const (
    // All routes
    DefaultRedistRoute_all DefaultRedistRoute = "all"
)

// Rip
// RIP configuration
type Rip struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RIP configuration for Default VRF.
    DefaultVrf Rip_DefaultVrf

    // VRF related RIP configuration.
    Vrfs Rip_Vrfs
}

func (rip *Rip) GetFilter() yfilter.YFilter { return rip.YFilter }

func (rip *Rip) SetFilter(yf yfilter.YFilter) { rip.YFilter = yf }

func (rip *Rip) GetGoName(yname string) string {
    if yname == "default-vrf" { return "DefaultVrf" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (rip *Rip) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-rip-cfg:rip"
}

func (rip *Rip) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-vrf" {
        return &rip.DefaultVrf
    }
    if childYangName == "vrfs" {
        return &rip.Vrfs
    }
    return nil
}

func (rip *Rip) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-vrf"] = &rip.DefaultVrf
    children["vrfs"] = &rip.Vrfs
    return children
}

func (rip *Rip) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rip *Rip) GetBundleName() string { return "cisco_ios_xr" }

func (rip *Rip) GetYangName() string { return "rip" }

func (rip *Rip) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rip *Rip) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rip *Rip) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rip *Rip) SetParent(parent types.Entity) { rip.parent = parent }

func (rip *Rip) GetParent() types.Entity { return rip.parent }

func (rip *Rip) GetParentYangName() string { return "Cisco-IOS-XR-ip-rip-cfg" }

// Rip_DefaultVrf
// RIP configuration for Default VRF
type Rip_DefaultVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Starts RIP configuration for Default VRF. The type is interface{}.
    Enable interface{}

    // Send RIP v2 output packets to broadcast address. The type is interface{}.
    BroadcastForV2 interface{}

    // Administrative distance. The type is interface{} with range: 0..255. The
    // default value is 120.
    Distance interface{}

    // Default metric of redistributed routes. The type is interface{} with range:
    // 0..16.
    DefaultMetric interface{}

    // Inter-packet delay for RIP updates. The type is interface{} with range:
    // 8..50. Units are millisecond.
    OutputDelay interface{}

    // Enable automatic network number summarization. The type is interface{}.
    AutoSummary interface{}

    // Route Policy for outbound routing updates. The type is string.
    PolicyOut interface{}

    // Disable validation of source address of routing updates. The type is
    // interface{}.
    ValidateSourceDisable interface{}

    // Maximum number of paths allowed per route. The type is interface{} with
    // range: 1..128. The default value is 4.
    MaximumPaths interface{}

    // Enable Cisco Non Stop Forwarding. The type is interface{}.
    Nsf interface{}

    // Route Policy for inbbound routing updates. The type is string.
    PolicyIn interface{}

    // Controls default information origination.
    DefaultInformation Rip_DefaultVrf_DefaultInformation

    // Redistribute information from another routing protocol.
    Redistribution Rip_DefaultVrf_Redistribution

    // Table of IP specific administrative distances.
    IpDistances Rip_DefaultVrf_IpDistances

    // Table of RIP interfaces.
    Interfaces Rip_DefaultVrf_Interfaces

    // Configure RIP Neighbors.
    Neighbors Rip_DefaultVrf_Neighbors

    // Various routing timers.
    Timers Rip_DefaultVrf_Timers
}

func (defaultVrf *Rip_DefaultVrf) GetFilter() yfilter.YFilter { return defaultVrf.YFilter }

func (defaultVrf *Rip_DefaultVrf) SetFilter(yf yfilter.YFilter) { defaultVrf.YFilter = yf }

func (defaultVrf *Rip_DefaultVrf) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "broadcast-for-v2" { return "BroadcastForV2" }
    if yname == "distance" { return "Distance" }
    if yname == "default-metric" { return "DefaultMetric" }
    if yname == "output-delay" { return "OutputDelay" }
    if yname == "auto-summary" { return "AutoSummary" }
    if yname == "policy-out" { return "PolicyOut" }
    if yname == "validate-source-disable" { return "ValidateSourceDisable" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    if yname == "nsf" { return "Nsf" }
    if yname == "policy-in" { return "PolicyIn" }
    if yname == "default-information" { return "DefaultInformation" }
    if yname == "redistribution" { return "Redistribution" }
    if yname == "ip-distances" { return "IpDistances" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "neighbors" { return "Neighbors" }
    if yname == "timers" { return "Timers" }
    return ""
}

func (defaultVrf *Rip_DefaultVrf) GetSegmentPath() string {
    return "default-vrf"
}

func (defaultVrf *Rip_DefaultVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-information" {
        return &defaultVrf.DefaultInformation
    }
    if childYangName == "redistribution" {
        return &defaultVrf.Redistribution
    }
    if childYangName == "ip-distances" {
        return &defaultVrf.IpDistances
    }
    if childYangName == "interfaces" {
        return &defaultVrf.Interfaces
    }
    if childYangName == "neighbors" {
        return &defaultVrf.Neighbors
    }
    if childYangName == "timers" {
        return &defaultVrf.Timers
    }
    return nil
}

func (defaultVrf *Rip_DefaultVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-information"] = &defaultVrf.DefaultInformation
    children["redistribution"] = &defaultVrf.Redistribution
    children["ip-distances"] = &defaultVrf.IpDistances
    children["interfaces"] = &defaultVrf.Interfaces
    children["neighbors"] = &defaultVrf.Neighbors
    children["timers"] = &defaultVrf.Timers
    return children
}

func (defaultVrf *Rip_DefaultVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = defaultVrf.Enable
    leafs["broadcast-for-v2"] = defaultVrf.BroadcastForV2
    leafs["distance"] = defaultVrf.Distance
    leafs["default-metric"] = defaultVrf.DefaultMetric
    leafs["output-delay"] = defaultVrf.OutputDelay
    leafs["auto-summary"] = defaultVrf.AutoSummary
    leafs["policy-out"] = defaultVrf.PolicyOut
    leafs["validate-source-disable"] = defaultVrf.ValidateSourceDisable
    leafs["maximum-paths"] = defaultVrf.MaximumPaths
    leafs["nsf"] = defaultVrf.Nsf
    leafs["policy-in"] = defaultVrf.PolicyIn
    return leafs
}

func (defaultVrf *Rip_DefaultVrf) GetBundleName() string { return "cisco_ios_xr" }

func (defaultVrf *Rip_DefaultVrf) GetYangName() string { return "default-vrf" }

func (defaultVrf *Rip_DefaultVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultVrf *Rip_DefaultVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultVrf *Rip_DefaultVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultVrf *Rip_DefaultVrf) SetParent(parent types.Entity) { defaultVrf.parent = parent }

func (defaultVrf *Rip_DefaultVrf) GetParent() types.Entity { return defaultVrf.parent }

func (defaultVrf *Rip_DefaultVrf) GetParentYangName() string { return "rip" }

// Rip_DefaultVrf_DefaultInformation
// Controls default information origination
// This type is a presence type.
type Rip_DefaultVrf_DefaultInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route policy name. The type is string.
    RoutePolicyName interface{}

    // Origination option. The type is DefaultInformationOption. This attribute is
    // mandatory.
    Option interface{}
}

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetFilter() yfilter.YFilter { return defaultInformation.YFilter }

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) SetFilter(yf yfilter.YFilter) { defaultInformation.YFilter = yf }

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "option" { return "Option" }
    return ""
}

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetSegmentPath() string {
    return "default-information"
}

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = defaultInformation.RoutePolicyName
    leafs["option"] = defaultInformation.Option
    return leafs
}

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetBundleName() string { return "cisco_ios_xr" }

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetYangName() string { return "default-information" }

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) SetParent(parent types.Entity) { defaultInformation.parent = parent }

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetParent() types.Entity { return defaultInformation.parent }

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetParentYangName() string { return "default-vrf" }

// Rip_DefaultVrf_Redistribution
// Redistribute information from another routing
// protocol
type Rip_DefaultVrf_Redistribution struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redistribute connected routes.
    Connected Rip_DefaultVrf_Redistribution_Connected

    // Redistribute BGP routes.
    Bgps Rip_DefaultVrf_Redistribution_Bgps

    // Redistribute IS-IS routes.
    Isises Rip_DefaultVrf_Redistribution_Isises

    // Redistribute EIGRP routes.
    EigrpS Rip_DefaultVrf_Redistribution_EigrpS

    // Redistribute static routes.
    Static Rip_DefaultVrf_Redistribution_Static

    // Redistribute OSPF routes.
    Ospfs Rip_DefaultVrf_Redistribution_Ospfs
}

func (redistribution *Rip_DefaultVrf_Redistribution) GetFilter() yfilter.YFilter { return redistribution.YFilter }

func (redistribution *Rip_DefaultVrf_Redistribution) SetFilter(yf yfilter.YFilter) { redistribution.YFilter = yf }

func (redistribution *Rip_DefaultVrf_Redistribution) GetGoName(yname string) string {
    if yname == "connected" { return "Connected" }
    if yname == "bgps" { return "Bgps" }
    if yname == "isises" { return "Isises" }
    if yname == "eigrp-s" { return "EigrpS" }
    if yname == "static" { return "Static" }
    if yname == "ospfs" { return "Ospfs" }
    return ""
}

func (redistribution *Rip_DefaultVrf_Redistribution) GetSegmentPath() string {
    return "redistribution"
}

func (redistribution *Rip_DefaultVrf_Redistribution) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "connected" {
        return &redistribution.Connected
    }
    if childYangName == "bgps" {
        return &redistribution.Bgps
    }
    if childYangName == "isises" {
        return &redistribution.Isises
    }
    if childYangName == "eigrp-s" {
        return &redistribution.EigrpS
    }
    if childYangName == "static" {
        return &redistribution.Static
    }
    if childYangName == "ospfs" {
        return &redistribution.Ospfs
    }
    return nil
}

func (redistribution *Rip_DefaultVrf_Redistribution) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["connected"] = &redistribution.Connected
    children["bgps"] = &redistribution.Bgps
    children["isises"] = &redistribution.Isises
    children["eigrp-s"] = &redistribution.EigrpS
    children["static"] = &redistribution.Static
    children["ospfs"] = &redistribution.Ospfs
    return children
}

func (redistribution *Rip_DefaultVrf_Redistribution) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (redistribution *Rip_DefaultVrf_Redistribution) GetBundleName() string { return "cisco_ios_xr" }

func (redistribution *Rip_DefaultVrf_Redistribution) GetYangName() string { return "redistribution" }

func (redistribution *Rip_DefaultVrf_Redistribution) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redistribution *Rip_DefaultVrf_Redistribution) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redistribution *Rip_DefaultVrf_Redistribution) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redistribution *Rip_DefaultVrf_Redistribution) SetParent(parent types.Entity) { redistribution.parent = parent }

func (redistribution *Rip_DefaultVrf_Redistribution) GetParent() types.Entity { return redistribution.parent }

func (redistribution *Rip_DefaultVrf_Redistribution) GetParentYangName() string { return "default-vrf" }

// Rip_DefaultVrf_Redistribution_Connected
// Redistribute connected routes
// This type is a presence type.
type Rip_DefaultVrf_Redistribution_Connected struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is DefaultRedistRoute.
    RouteType interface{}
}

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetFilter() yfilter.YFilter { return connected.YFilter }

func (connected *Rip_DefaultVrf_Redistribution_Connected) SetFilter(yf yfilter.YFilter) { connected.YFilter = yf }

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "route-type" { return "RouteType" }
    return ""
}

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetSegmentPath() string {
    return "connected"
}

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = connected.RoutePolicyName
    leafs["route-type"] = connected.RouteType
    return leafs
}

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetBundleName() string { return "cisco_ios_xr" }

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetYangName() string { return "connected" }

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (connected *Rip_DefaultVrf_Redistribution_Connected) SetParent(parent types.Entity) { connected.parent = parent }

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetParent() types.Entity { return connected.parent }

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetParentYangName() string { return "redistribution" }

// Rip_DefaultVrf_Redistribution_Bgps
// Redistribute BGP routes
type Rip_DefaultVrf_Redistribution_Bgps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Autonomous system number. The type is slice of
    // Rip_DefaultVrf_Redistribution_Bgps_Bgp.
    Bgp []Rip_DefaultVrf_Redistribution_Bgps_Bgp
}

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetFilter() yfilter.YFilter { return bgps.YFilter }

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) SetFilter(yf yfilter.YFilter) { bgps.YFilter = yf }

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetGoName(yname string) string {
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetSegmentPath() string {
    return "bgps"
}

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp" {
        for _, c := range bgps.Bgp {
            if bgps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_DefaultVrf_Redistribution_Bgps_Bgp{}
        bgps.Bgp = append(bgps.Bgp, child)
        return &bgps.Bgp[len(bgps.Bgp)-1]
    }
    return nil
}

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgps.Bgp {
        children[bgps.Bgp[i].GetSegmentPath()] = &bgps.Bgp[i]
    }
    return children
}

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetBundleName() string { return "cisco_ios_xr" }

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetYangName() string { return "bgps" }

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) SetParent(parent types.Entity) { bgps.parent = parent }

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetParent() types.Entity { return bgps.parent }

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetParentYangName() string { return "redistribution" }

// Rip_DefaultVrf_Redistribution_Bgps_Bgp
// Autonomous system number
type Rip_DefaultVrf_Redistribution_Bgps_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Higher 16 bits of 4-byte BGP AS number. The type
    // is interface{} with range: 0..65535.
    Asnxx interface{}

    // This attribute is a key. 2-byte or 4-byte BGP AS number. The type is
    // interface{} with range: 0..4294967295.
    Asnyy interface{}

    // Route Policy name. The type is string.
    Policy interface{}

    // Route type. The type is BgpRedistRoute.
    Type interface{}
}

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetGoName(yname string) string {
    if yname == "asnxx" { return "Asnxx" }
    if yname == "asnyy" { return "Asnyy" }
    if yname == "policy" { return "Policy" }
    if yname == "type" { return "Type" }
    return ""
}

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetSegmentPath() string {
    return "bgp" + "[asnxx='" + fmt.Sprintf("%v", bgp.Asnxx) + "']" + "[asnyy='" + fmt.Sprintf("%v", bgp.Asnyy) + "']"
}

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["asnxx"] = bgp.Asnxx
    leafs["asnyy"] = bgp.Asnyy
    leafs["policy"] = bgp.Policy
    leafs["type"] = bgp.Type
    return leafs
}

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetYangName() string { return "bgp" }

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetParentYangName() string { return "bgps" }

// Rip_DefaultVrf_Redistribution_Isises
// Redistribute IS-IS routes
type Rip_DefaultVrf_Redistribution_Isises struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redistribute IS-IS routes. The type is slice of
    // Rip_DefaultVrf_Redistribution_Isises_Isis.
    Isis []Rip_DefaultVrf_Redistribution_Isises_Isis
}

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetFilter() yfilter.YFilter { return isises.YFilter }

func (isises *Rip_DefaultVrf_Redistribution_Isises) SetFilter(yf yfilter.YFilter) { isises.YFilter = yf }

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetGoName(yname string) string {
    if yname == "isis" { return "Isis" }
    return ""
}

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetSegmentPath() string {
    return "isises"
}

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        for _, c := range isises.Isis {
            if isises.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_DefaultVrf_Redistribution_Isises_Isis{}
        isises.Isis = append(isises.Isis, child)
        return &isises.Isis[len(isises.Isis)-1]
    }
    return nil
}

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range isises.Isis {
        children[isises.Isis[i].GetSegmentPath()] = &isises.Isis[i]
    }
    return children
}

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetBundleName() string { return "cisco_ios_xr" }

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetYangName() string { return "isises" }

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isises *Rip_DefaultVrf_Redistribution_Isises) SetParent(parent types.Entity) { isises.parent = parent }

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetParent() types.Entity { return isises.parent }

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetParentYangName() string { return "redistribution" }

// Rip_DefaultVrf_Redistribution_Isises_Isis
// Redistribute IS-IS routes
type Rip_DefaultVrf_Redistribution_Isises_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IS-IS instance name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    IsisName interface{}

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is IsisRedistRoute.
    RouteType interface{}
}

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetGoName(yname string) string {
    if yname == "isis-name" { return "IsisName" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "route-type" { return "RouteType" }
    return ""
}

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetSegmentPath() string {
    return "isis" + "[isis-name='" + fmt.Sprintf("%v", isis.IsisName) + "']"
}

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["isis-name"] = isis.IsisName
    leafs["route-policy-name"] = isis.RoutePolicyName
    leafs["route-type"] = isis.RouteType
    return leafs
}

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetYangName() string { return "isis" }

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetParent() types.Entity { return isis.parent }

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetParentYangName() string { return "isises" }

// Rip_DefaultVrf_Redistribution_EigrpS
// Redistribute EIGRP routes
type Rip_DefaultVrf_Redistribution_EigrpS struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redistribute EIGRP routes. The type is slice of
    // Rip_DefaultVrf_Redistribution_EigrpS_Eigrp.
    Eigrp []Rip_DefaultVrf_Redistribution_EigrpS_Eigrp
}

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetFilter() yfilter.YFilter { return eigrpS.YFilter }

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) SetFilter(yf yfilter.YFilter) { eigrpS.YFilter = yf }

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetGoName(yname string) string {
    if yname == "eigrp" { return "Eigrp" }
    return ""
}

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetSegmentPath() string {
    return "eigrp-s"
}

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "eigrp" {
        for _, c := range eigrpS.Eigrp {
            if eigrpS.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_DefaultVrf_Redistribution_EigrpS_Eigrp{}
        eigrpS.Eigrp = append(eigrpS.Eigrp, child)
        return &eigrpS.Eigrp[len(eigrpS.Eigrp)-1]
    }
    return nil
}

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range eigrpS.Eigrp {
        children[eigrpS.Eigrp[i].GetSegmentPath()] = &eigrpS.Eigrp[i]
    }
    return children
}

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetBundleName() string { return "cisco_ios_xr" }

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetYangName() string { return "eigrp-s" }

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) SetParent(parent types.Entity) { eigrpS.parent = parent }

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetParent() types.Entity { return eigrpS.parent }

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetParentYangName() string { return "redistribution" }

// Rip_DefaultVrf_Redistribution_EigrpS_Eigrp
// Redistribute EIGRP routes
type Rip_DefaultVrf_Redistribution_EigrpS_Eigrp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Autonomous system number. The type is interface{}
    // with range: 1..65535.
    As interface{}

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is DefaultRedistRoute.
    RouteType interface{}
}

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetFilter() yfilter.YFilter { return eigrp.YFilter }

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) SetFilter(yf yfilter.YFilter) { eigrp.YFilter = yf }

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "route-type" { return "RouteType" }
    return ""
}

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetSegmentPath() string {
    return "eigrp" + "[as='" + fmt.Sprintf("%v", eigrp.As) + "']"
}

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = eigrp.As
    leafs["route-policy-name"] = eigrp.RoutePolicyName
    leafs["route-type"] = eigrp.RouteType
    return leafs
}

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetBundleName() string { return "cisco_ios_xr" }

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetYangName() string { return "eigrp" }

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) SetParent(parent types.Entity) { eigrp.parent = parent }

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetParent() types.Entity { return eigrp.parent }

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetParentYangName() string { return "eigrp-s" }

// Rip_DefaultVrf_Redistribution_Static
// Redistribute static routes
// This type is a presence type.
type Rip_DefaultVrf_Redistribution_Static struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is DefaultRedistRoute.
    RouteType interface{}
}

func (static *Rip_DefaultVrf_Redistribution_Static) GetFilter() yfilter.YFilter { return static.YFilter }

func (static *Rip_DefaultVrf_Redistribution_Static) SetFilter(yf yfilter.YFilter) { static.YFilter = yf }

func (static *Rip_DefaultVrf_Redistribution_Static) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "route-type" { return "RouteType" }
    return ""
}

func (static *Rip_DefaultVrf_Redistribution_Static) GetSegmentPath() string {
    return "static"
}

func (static *Rip_DefaultVrf_Redistribution_Static) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (static *Rip_DefaultVrf_Redistribution_Static) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (static *Rip_DefaultVrf_Redistribution_Static) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = static.RoutePolicyName
    leafs["route-type"] = static.RouteType
    return leafs
}

func (static *Rip_DefaultVrf_Redistribution_Static) GetBundleName() string { return "cisco_ios_xr" }

func (static *Rip_DefaultVrf_Redistribution_Static) GetYangName() string { return "static" }

func (static *Rip_DefaultVrf_Redistribution_Static) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (static *Rip_DefaultVrf_Redistribution_Static) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (static *Rip_DefaultVrf_Redistribution_Static) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (static *Rip_DefaultVrf_Redistribution_Static) SetParent(parent types.Entity) { static.parent = parent }

func (static *Rip_DefaultVrf_Redistribution_Static) GetParent() types.Entity { return static.parent }

func (static *Rip_DefaultVrf_Redistribution_Static) GetParentYangName() string { return "redistribution" }

// Rip_DefaultVrf_Redistribution_Ospfs
// Redistribute OSPF routes
type Rip_DefaultVrf_Redistribution_Ospfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redistribute OSPF routes. The type is slice of
    // Rip_DefaultVrf_Redistribution_Ospfs_Ospf.
    Ospf []Rip_DefaultVrf_Redistribution_Ospfs_Ospf
}

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetFilter() yfilter.YFilter { return ospfs.YFilter }

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) SetFilter(yf yfilter.YFilter) { ospfs.YFilter = yf }

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetGoName(yname string) string {
    if yname == "ospf" { return "Ospf" }
    return ""
}

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetSegmentPath() string {
    return "ospfs"
}

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospf" {
        for _, c := range ospfs.Ospf {
            if ospfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_DefaultVrf_Redistribution_Ospfs_Ospf{}
        ospfs.Ospf = append(ospfs.Ospf, child)
        return &ospfs.Ospf[len(ospfs.Ospf)-1]
    }
    return nil
}

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfs.Ospf {
        children[ospfs.Ospf[i].GetSegmentPath()] = &ospfs.Ospf[i]
    }
    return children
}

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetBundleName() string { return "cisco_ios_xr" }

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetYangName() string { return "ospfs" }

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) SetParent(parent types.Entity) { ospfs.parent = parent }

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetParent() types.Entity { return ospfs.parent }

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetParentYangName() string { return "redistribution" }

// Rip_DefaultVrf_Redistribution_Ospfs_Ospf
// Redistribute OSPF routes
type Rip_DefaultVrf_Redistribution_Ospfs_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Process ID for the OSPF instance. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    OspfName interface{}

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Internal routes. The type is bool.
    Internal interface{}

    // External routes. The type is bool.
    External interface{}

    // External route type. The type is interface{} with range: 0..2.
    ExternalType interface{}

    // NSSA External routes. The type is bool.
    NssaExternal interface{}

    // NSSA External route type. The type is interface{} with range: 0..2.
    NssaExternalType interface{}
}

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetGoName(yname string) string {
    if yname == "ospf-name" { return "OspfName" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "internal" { return "Internal" }
    if yname == "external" { return "External" }
    if yname == "external-type" { return "ExternalType" }
    if yname == "nssa-external" { return "NssaExternal" }
    if yname == "nssa-external-type" { return "NssaExternalType" }
    return ""
}

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetSegmentPath() string {
    return "ospf" + "[ospf-name='" + fmt.Sprintf("%v", ospf.OspfName) + "']"
}

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospf-name"] = ospf.OspfName
    leafs["route-policy-name"] = ospf.RoutePolicyName
    leafs["internal"] = ospf.Internal
    leafs["external"] = ospf.External
    leafs["external-type"] = ospf.ExternalType
    leafs["nssa-external"] = ospf.NssaExternal
    leafs["nssa-external-type"] = ospf.NssaExternalType
    return leafs
}

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetYangName() string { return "ospf" }

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetParentYangName() string { return "ospfs" }

// Rip_DefaultVrf_IpDistances
// Table of IP specific administrative distances
type Rip_DefaultVrf_IpDistances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP specific administrative distance. The type is slice of
    // Rip_DefaultVrf_IpDistances_IpDistance.
    IpDistance []Rip_DefaultVrf_IpDistances_IpDistance
}

func (ipDistances *Rip_DefaultVrf_IpDistances) GetFilter() yfilter.YFilter { return ipDistances.YFilter }

func (ipDistances *Rip_DefaultVrf_IpDistances) SetFilter(yf yfilter.YFilter) { ipDistances.YFilter = yf }

func (ipDistances *Rip_DefaultVrf_IpDistances) GetGoName(yname string) string {
    if yname == "ip-distance" { return "IpDistance" }
    return ""
}

func (ipDistances *Rip_DefaultVrf_IpDistances) GetSegmentPath() string {
    return "ip-distances"
}

func (ipDistances *Rip_DefaultVrf_IpDistances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ip-distance" {
        for _, c := range ipDistances.IpDistance {
            if ipDistances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_DefaultVrf_IpDistances_IpDistance{}
        ipDistances.IpDistance = append(ipDistances.IpDistance, child)
        return &ipDistances.IpDistance[len(ipDistances.IpDistance)-1]
    }
    return nil
}

func (ipDistances *Rip_DefaultVrf_IpDistances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipDistances.IpDistance {
        children[ipDistances.IpDistance[i].GetSegmentPath()] = &ipDistances.IpDistance[i]
    }
    return children
}

func (ipDistances *Rip_DefaultVrf_IpDistances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipDistances *Rip_DefaultVrf_IpDistances) GetBundleName() string { return "cisco_ios_xr" }

func (ipDistances *Rip_DefaultVrf_IpDistances) GetYangName() string { return "ip-distances" }

func (ipDistances *Rip_DefaultVrf_IpDistances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipDistances *Rip_DefaultVrf_IpDistances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipDistances *Rip_DefaultVrf_IpDistances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipDistances *Rip_DefaultVrf_IpDistances) SetParent(parent types.Entity) { ipDistances.parent = parent }

func (ipDistances *Rip_DefaultVrf_IpDistances) GetParent() types.Entity { return ipDistances.parent }

func (ipDistances *Rip_DefaultVrf_IpDistances) GetParentYangName() string { return "default-vrf" }

// Rip_DefaultVrf_IpDistances_IpDistance
// IP specific administrative distance
type Rip_DefaultVrf_IpDistances_IpDistance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP Source address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IP address mask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netmask interface{}

    // Administrative distance. The type is interface{} with range: 0..255. This
    // attribute is mandatory.
    Distance interface{}
}

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetFilter() yfilter.YFilter { return ipDistance.YFilter }

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) SetFilter(yf yfilter.YFilter) { ipDistance.YFilter = yf }

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "netmask" { return "Netmask" }
    if yname == "distance" { return "Distance" }
    return ""
}

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetSegmentPath() string {
    return "ip-distance" + "[address='" + fmt.Sprintf("%v", ipDistance.Address) + "']" + "[netmask='" + fmt.Sprintf("%v", ipDistance.Netmask) + "']"
}

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = ipDistance.Address
    leafs["netmask"] = ipDistance.Netmask
    leafs["distance"] = ipDistance.Distance
    return leafs
}

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetBundleName() string { return "cisco_ios_xr" }

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetYangName() string { return "ip-distance" }

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) SetParent(parent types.Entity) { ipDistance.parent = parent }

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetParent() types.Entity { return ipDistance.parent }

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetParentYangName() string { return "ip-distances" }

// Rip_DefaultVrf_Interfaces
// Table of RIP interfaces
type Rip_DefaultVrf_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RIP interface name. The type is slice of
    // Rip_DefaultVrf_Interfaces_Interface.
    Interface []Rip_DefaultVrf_Interfaces_Interface
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Rip_DefaultVrf_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Rip_DefaultVrf_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_DefaultVrf_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Rip_DefaultVrf_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Rip_DefaultVrf_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Rip_DefaultVrf_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Rip_DefaultVrf_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Rip_DefaultVrf_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Rip_DefaultVrf_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Rip_DefaultVrf_Interfaces) GetParentYangName() string { return "default-vrf" }

// Rip_DefaultVrf_Interfaces_Interface
// RIP interface name
type Rip_DefaultVrf_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Send RIP v2 output packets to broadcast address. The type is interface{}.
    BroadcastForV2 interface{}

    // Enable poison reverse. The type is interface{}.
    PoisonReverse interface{}

    // Suppress routing updates on this interface. The type is interface{}.
    Passive interface{}

    // Starts RIP interface configuration. The type is interface{}.
    Enable interface{}

    // Route Policy for outbound routing updates. The type is string.
    PolicyOut interface{}

    // Accept RIP updates with metric 0. The type is interface{}.
    AcceptMetricZero interface{}

    // Route Policy for inbound routing updates. The type is string.
    PolicyIn interface{}

    // Disable split horizon. The type is interface{}.
    SplitHorizonDisable interface{}

    // Authentication keychain and mode.
    Authentication Rip_DefaultVrf_Interfaces_Interface_Authentication

    // SOO community for prefixes learned over this interface.
    SiteOfOrigin Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin

    // RIP versions supported for receiving advertisements.
    ReceiveVersion Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion

    // RIP versions supported for sending advertisements.
    SendVersion Rip_DefaultVrf_Interfaces_Interface_SendVersion
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Rip_DefaultVrf_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "broadcast-for-v2" { return "BroadcastForV2" }
    if yname == "poison-reverse" { return "PoisonReverse" }
    if yname == "passive" { return "Passive" }
    if yname == "enable" { return "Enable" }
    if yname == "policy-out" { return "PolicyOut" }
    if yname == "accept-metric-zero" { return "AcceptMetricZero" }
    if yname == "policy-in" { return "PolicyIn" }
    if yname == "split-horizon-disable" { return "SplitHorizonDisable" }
    if yname == "authentication" { return "Authentication" }
    if yname == "site-of-origin" { return "SiteOfOrigin" }
    if yname == "receive-version" { return "ReceiveVersion" }
    if yname == "send-version" { return "SendVersion" }
    return ""
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authentication" {
        return &self.Authentication
    }
    if childYangName == "site-of-origin" {
        return &self.SiteOfOrigin
    }
    if childYangName == "receive-version" {
        return &self.ReceiveVersion
    }
    if childYangName == "send-version" {
        return &self.SendVersion
    }
    return nil
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["authentication"] = &self.Authentication
    children["site-of-origin"] = &self.SiteOfOrigin
    children["receive-version"] = &self.ReceiveVersion
    children["send-version"] = &self.SendVersion
    return children
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["broadcast-for-v2"] = self.BroadcastForV2
    leafs["poison-reverse"] = self.PoisonReverse
    leafs["passive"] = self.Passive
    leafs["enable"] = self.Enable
    leafs["policy-out"] = self.PolicyOut
    leafs["accept-metric-zero"] = self.AcceptMetricZero
    leafs["policy-in"] = self.PolicyIn
    leafs["split-horizon-disable"] = self.SplitHorizonDisable
    return leafs
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Rip_DefaultVrf_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Rip_DefaultVrf_Interfaces_Interface_Authentication
// Authentication keychain and mode
// This type is a presence type.
type Rip_DefaultVrf_Interfaces_Interface_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of keychain. The type is string. This attribute is mandatory.
    Keychain interface{}

    // Authentication mode. The type is RipAuthMode. This attribute is mandatory.
    Mode interface{}
}

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetGoName(yname string) string {
    if yname == "keychain" { return "Keychain" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["keychain"] = authentication.Keychain
    leafs["mode"] = authentication.Mode
    return leafs
}

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetYangName() string { return "authentication" }

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetParentYangName() string { return "interface" }

// Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin
// SOO community for prefixes learned over this
// interface
type Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Extended community type. The type is RipExtCommunity.
    Type interface{}

    // AS Number for AS:nn format. The type is interface{} with range: 0..65535.
    AsXx interface{}

    // 32 bit value for AS:nn format. The type is interface{} with range:
    // 0..4294967295.
    AsYy interface{}

    // AS Number Index. The type is interface{} with range: 0..4294967295.
    AsIndex interface{}

    // IPV4 address for IPV4Address:nn format. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // 16bit value for IPV4Address:nn format. The type is interface{} with range:
    // 0..65535.
    AddressIndex interface{}
}

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetFilter() yfilter.YFilter { return siteOfOrigin.YFilter }

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) SetFilter(yf yfilter.YFilter) { siteOfOrigin.YFilter = yf }

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "as-xx" { return "AsXx" }
    if yname == "as-yy" { return "AsYy" }
    if yname == "as-index" { return "AsIndex" }
    if yname == "address" { return "Address" }
    if yname == "address-index" { return "AddressIndex" }
    return ""
}

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetSegmentPath() string {
    return "site-of-origin"
}

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = siteOfOrigin.Type
    leafs["as-xx"] = siteOfOrigin.AsXx
    leafs["as-yy"] = siteOfOrigin.AsYy
    leafs["as-index"] = siteOfOrigin.AsIndex
    leafs["address"] = siteOfOrigin.Address
    leafs["address-index"] = siteOfOrigin.AddressIndex
    return leafs
}

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetBundleName() string { return "cisco_ios_xr" }

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetYangName() string { return "site-of-origin" }

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) SetParent(parent types.Entity) { siteOfOrigin.parent = parent }

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetParent() types.Entity { return siteOfOrigin.parent }

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetParentYangName() string { return "interface" }

// Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion
// RIP versions supported for receiving
// advertisements
type Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Support RIP version 1. The type is bool.
    Version1 interface{}

    // Support RIP version 2. The type is bool. The default value is true.
    Version2 interface{}
}

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetFilter() yfilter.YFilter { return receiveVersion.YFilter }

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) SetFilter(yf yfilter.YFilter) { receiveVersion.YFilter = yf }

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetGoName(yname string) string {
    if yname == "version1" { return "Version1" }
    if yname == "version2" { return "Version2" }
    return ""
}

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetSegmentPath() string {
    return "receive-version"
}

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version1"] = receiveVersion.Version1
    leafs["version2"] = receiveVersion.Version2
    return leafs
}

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetBundleName() string { return "cisco_ios_xr" }

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetYangName() string { return "receive-version" }

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) SetParent(parent types.Entity) { receiveVersion.parent = parent }

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetParent() types.Entity { return receiveVersion.parent }

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetParentYangName() string { return "interface" }

// Rip_DefaultVrf_Interfaces_Interface_SendVersion
// RIP versions supported for sending
// advertisements
type Rip_DefaultVrf_Interfaces_Interface_SendVersion struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Support RIP version 1. The type is bool.
    Version1 interface{}

    // Support RIP version 2. The type is bool. The default value is true.
    Version2 interface{}
}

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetFilter() yfilter.YFilter { return sendVersion.YFilter }

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) SetFilter(yf yfilter.YFilter) { sendVersion.YFilter = yf }

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetGoName(yname string) string {
    if yname == "version1" { return "Version1" }
    if yname == "version2" { return "Version2" }
    return ""
}

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetSegmentPath() string {
    return "send-version"
}

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version1"] = sendVersion.Version1
    leafs["version2"] = sendVersion.Version2
    return leafs
}

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetBundleName() string { return "cisco_ios_xr" }

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetYangName() string { return "send-version" }

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) SetParent(parent types.Entity) { sendVersion.parent = parent }

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetParent() types.Entity { return sendVersion.parent }

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetParentYangName() string { return "interface" }

// Rip_DefaultVrf_Neighbors
// Configure RIP Neighbors
type Rip_DefaultVrf_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor address. The type is slice of Rip_DefaultVrf_Neighbors_Neighbor.
    Neighbor []Rip_DefaultVrf_Neighbors_Neighbor
}

func (neighbors *Rip_DefaultVrf_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Rip_DefaultVrf_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Rip_DefaultVrf_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *Rip_DefaultVrf_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Rip_DefaultVrf_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_DefaultVrf_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *Rip_DefaultVrf_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *Rip_DefaultVrf_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Rip_DefaultVrf_Neighbors) GetBundleName() string { return "cisco_ios_xr" }

func (neighbors *Rip_DefaultVrf_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Rip_DefaultVrf_Neighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbors *Rip_DefaultVrf_Neighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbors *Rip_DefaultVrf_Neighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbors *Rip_DefaultVrf_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Rip_DefaultVrf_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Rip_DefaultVrf_Neighbors) GetParentYangName() string { return "default-vrf" }

// Rip_DefaultVrf_Neighbors_Neighbor
// Neighbor address
type Rip_DefaultVrf_Neighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}
}

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-address" { return "NeighborAddress" }
    return ""
}

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[neighbor-address='" + fmt.Sprintf("%v", neighbor.NeighborAddress) + "']"
}

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-address"] = neighbor.NeighborAddress
    return leafs
}

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// Rip_DefaultVrf_Timers
// Various routing timers
// This type is a presence type.
type Rip_DefaultVrf_Timers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interval between updates. The type is interface{} with range: 5..50000.
    // This attribute is mandatory.
    UpdateTimer interface{}

    // Invalid. The type is interface{} with range: 15..200000. This attribute is
    // mandatory.
    InvalidTimer interface{}

    // Holddown. The type is interface{} with range: 15..200000. This attribute is
    // mandatory.
    HolddownTimer interface{}

    // Flush. The type is interface{} with range: 16..250000. This attribute is
    // mandatory.
    FlushTimer interface{}
}

func (timers *Rip_DefaultVrf_Timers) GetFilter() yfilter.YFilter { return timers.YFilter }

func (timers *Rip_DefaultVrf_Timers) SetFilter(yf yfilter.YFilter) { timers.YFilter = yf }

func (timers *Rip_DefaultVrf_Timers) GetGoName(yname string) string {
    if yname == "update-timer" { return "UpdateTimer" }
    if yname == "invalid-timer" { return "InvalidTimer" }
    if yname == "holddown-timer" { return "HolddownTimer" }
    if yname == "flush-timer" { return "FlushTimer" }
    return ""
}

func (timers *Rip_DefaultVrf_Timers) GetSegmentPath() string {
    return "timers"
}

func (timers *Rip_DefaultVrf_Timers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timers *Rip_DefaultVrf_Timers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timers *Rip_DefaultVrf_Timers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["update-timer"] = timers.UpdateTimer
    leafs["invalid-timer"] = timers.InvalidTimer
    leafs["holddown-timer"] = timers.HolddownTimer
    leafs["flush-timer"] = timers.FlushTimer
    return leafs
}

func (timers *Rip_DefaultVrf_Timers) GetBundleName() string { return "cisco_ios_xr" }

func (timers *Rip_DefaultVrf_Timers) GetYangName() string { return "timers" }

func (timers *Rip_DefaultVrf_Timers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timers *Rip_DefaultVrf_Timers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timers *Rip_DefaultVrf_Timers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timers *Rip_DefaultVrf_Timers) SetParent(parent types.Entity) { timers.parent = parent }

func (timers *Rip_DefaultVrf_Timers) GetParent() types.Entity { return timers.parent }

func (timers *Rip_DefaultVrf_Timers) GetParentYangName() string { return "default-vrf" }

// Rip_Vrfs
// VRF related RIP configuration
type Rip_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RIP configuration for a particular VRF. The type is slice of Rip_Vrfs_Vrf.
    Vrf []Rip_Vrfs_Vrf
}

func (vrfs *Rip_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Rip_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Rip_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Rip_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Rip_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Rip_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Rip_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Rip_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Rip_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Rip_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Rip_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Rip_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Rip_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Rip_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Rip_Vrfs) GetParentYangName() string { return "rip" }

// Rip_Vrfs_Vrf
// RIP configuration for a particular VRF
type Rip_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Starts RIP configuration for a particular VRF. The type is interface{}.
    Enable interface{}

    // Send RIP v2 output packets to broadcast address. The type is interface{}.
    BroadcastForV2 interface{}

    // Administrative distance. The type is interface{} with range: 0..255. The
    // default value is 120.
    Distance interface{}

    // Default metric of redistributed routes. The type is interface{} with range:
    // 0..16.
    DefaultMetric interface{}

    // Inter-packet delay for RIP updates. The type is interface{} with range:
    // 8..50. Units are millisecond.
    OutputDelay interface{}

    // Enable automatic network number summarization. The type is interface{}.
    AutoSummary interface{}

    // Route Policy for outbound routing updates. The type is string.
    PolicyOut interface{}

    // Disable validation of source address of routing updates. The type is
    // interface{}.
    ValidateSourceDisable interface{}

    // Maximum number of paths allowed per route. The type is interface{} with
    // range: 1..128. The default value is 4.
    MaximumPaths interface{}

    // Enable Cisco Non Stop Forwarding. The type is interface{}.
    Nsf interface{}

    // Route Policy for inbbound routing updates. The type is string.
    PolicyIn interface{}

    // Controls default information origination.
    DefaultInformation Rip_Vrfs_Vrf_DefaultInformation

    // Redistribute information from another routing protocol.
    Redistribution Rip_Vrfs_Vrf_Redistribution

    // Table of IP specific administrative distances.
    IpDistances Rip_Vrfs_Vrf_IpDistances

    // Table of RIP interfaces.
    Interfaces Rip_Vrfs_Vrf_Interfaces

    // Configure RIP Neighbors.
    Neighbors Rip_Vrfs_Vrf_Neighbors

    // Various routing timers.
    Timers Rip_Vrfs_Vrf_Timers
}

func (vrf *Rip_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Rip_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Rip_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "enable" { return "Enable" }
    if yname == "broadcast-for-v2" { return "BroadcastForV2" }
    if yname == "distance" { return "Distance" }
    if yname == "default-metric" { return "DefaultMetric" }
    if yname == "output-delay" { return "OutputDelay" }
    if yname == "auto-summary" { return "AutoSummary" }
    if yname == "policy-out" { return "PolicyOut" }
    if yname == "validate-source-disable" { return "ValidateSourceDisable" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    if yname == "nsf" { return "Nsf" }
    if yname == "policy-in" { return "PolicyIn" }
    if yname == "default-information" { return "DefaultInformation" }
    if yname == "redistribution" { return "Redistribution" }
    if yname == "ip-distances" { return "IpDistances" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "neighbors" { return "Neighbors" }
    if yname == "timers" { return "Timers" }
    return ""
}

func (vrf *Rip_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Rip_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-information" {
        return &vrf.DefaultInformation
    }
    if childYangName == "redistribution" {
        return &vrf.Redistribution
    }
    if childYangName == "ip-distances" {
        return &vrf.IpDistances
    }
    if childYangName == "interfaces" {
        return &vrf.Interfaces
    }
    if childYangName == "neighbors" {
        return &vrf.Neighbors
    }
    if childYangName == "timers" {
        return &vrf.Timers
    }
    return nil
}

func (vrf *Rip_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-information"] = &vrf.DefaultInformation
    children["redistribution"] = &vrf.Redistribution
    children["ip-distances"] = &vrf.IpDistances
    children["interfaces"] = &vrf.Interfaces
    children["neighbors"] = &vrf.Neighbors
    children["timers"] = &vrf.Timers
    return children
}

func (vrf *Rip_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["enable"] = vrf.Enable
    leafs["broadcast-for-v2"] = vrf.BroadcastForV2
    leafs["distance"] = vrf.Distance
    leafs["default-metric"] = vrf.DefaultMetric
    leafs["output-delay"] = vrf.OutputDelay
    leafs["auto-summary"] = vrf.AutoSummary
    leafs["policy-out"] = vrf.PolicyOut
    leafs["validate-source-disable"] = vrf.ValidateSourceDisable
    leafs["maximum-paths"] = vrf.MaximumPaths
    leafs["nsf"] = vrf.Nsf
    leafs["policy-in"] = vrf.PolicyIn
    return leafs
}

func (vrf *Rip_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Rip_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Rip_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Rip_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Rip_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Rip_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Rip_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Rip_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Rip_Vrfs_Vrf_DefaultInformation
// Controls default information origination
// This type is a presence type.
type Rip_Vrfs_Vrf_DefaultInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route policy name. The type is string.
    RoutePolicyName interface{}

    // Origination option. The type is DefaultInformationOption. This attribute is
    // mandatory.
    Option interface{}
}

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetFilter() yfilter.YFilter { return defaultInformation.YFilter }

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) SetFilter(yf yfilter.YFilter) { defaultInformation.YFilter = yf }

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "option" { return "Option" }
    return ""
}

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetSegmentPath() string {
    return "default-information"
}

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = defaultInformation.RoutePolicyName
    leafs["option"] = defaultInformation.Option
    return leafs
}

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetBundleName() string { return "cisco_ios_xr" }

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetYangName() string { return "default-information" }

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) SetParent(parent types.Entity) { defaultInformation.parent = parent }

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetParent() types.Entity { return defaultInformation.parent }

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetParentYangName() string { return "vrf" }

// Rip_Vrfs_Vrf_Redistribution
// Redistribute information from another routing
// protocol
type Rip_Vrfs_Vrf_Redistribution struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redistribute connected routes.
    Connected Rip_Vrfs_Vrf_Redistribution_Connected

    // Redistribute BGP routes.
    Bgps Rip_Vrfs_Vrf_Redistribution_Bgps

    // Redistribute IS-IS routes.
    Isises Rip_Vrfs_Vrf_Redistribution_Isises

    // Redistribute EIGRP routes.
    EigrpS Rip_Vrfs_Vrf_Redistribution_EigrpS

    // Redistribute static routes.
    Static Rip_Vrfs_Vrf_Redistribution_Static

    // Redistribute OSPF routes.
    Ospfs Rip_Vrfs_Vrf_Redistribution_Ospfs
}

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetFilter() yfilter.YFilter { return redistribution.YFilter }

func (redistribution *Rip_Vrfs_Vrf_Redistribution) SetFilter(yf yfilter.YFilter) { redistribution.YFilter = yf }

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetGoName(yname string) string {
    if yname == "connected" { return "Connected" }
    if yname == "bgps" { return "Bgps" }
    if yname == "isises" { return "Isises" }
    if yname == "eigrp-s" { return "EigrpS" }
    if yname == "static" { return "Static" }
    if yname == "ospfs" { return "Ospfs" }
    return ""
}

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetSegmentPath() string {
    return "redistribution"
}

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "connected" {
        return &redistribution.Connected
    }
    if childYangName == "bgps" {
        return &redistribution.Bgps
    }
    if childYangName == "isises" {
        return &redistribution.Isises
    }
    if childYangName == "eigrp-s" {
        return &redistribution.EigrpS
    }
    if childYangName == "static" {
        return &redistribution.Static
    }
    if childYangName == "ospfs" {
        return &redistribution.Ospfs
    }
    return nil
}

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["connected"] = &redistribution.Connected
    children["bgps"] = &redistribution.Bgps
    children["isises"] = &redistribution.Isises
    children["eigrp-s"] = &redistribution.EigrpS
    children["static"] = &redistribution.Static
    children["ospfs"] = &redistribution.Ospfs
    return children
}

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetBundleName() string { return "cisco_ios_xr" }

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetYangName() string { return "redistribution" }

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redistribution *Rip_Vrfs_Vrf_Redistribution) SetParent(parent types.Entity) { redistribution.parent = parent }

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetParent() types.Entity { return redistribution.parent }

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetParentYangName() string { return "vrf" }

// Rip_Vrfs_Vrf_Redistribution_Connected
// Redistribute connected routes
// This type is a presence type.
type Rip_Vrfs_Vrf_Redistribution_Connected struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is DefaultRedistRoute.
    RouteType interface{}
}

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetFilter() yfilter.YFilter { return connected.YFilter }

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) SetFilter(yf yfilter.YFilter) { connected.YFilter = yf }

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "route-type" { return "RouteType" }
    return ""
}

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetSegmentPath() string {
    return "connected"
}

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = connected.RoutePolicyName
    leafs["route-type"] = connected.RouteType
    return leafs
}

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetBundleName() string { return "cisco_ios_xr" }

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetYangName() string { return "connected" }

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) SetParent(parent types.Entity) { connected.parent = parent }

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetParent() types.Entity { return connected.parent }

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetParentYangName() string { return "redistribution" }

// Rip_Vrfs_Vrf_Redistribution_Bgps
// Redistribute BGP routes
type Rip_Vrfs_Vrf_Redistribution_Bgps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Autonomous system number. The type is slice of
    // Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp.
    Bgp []Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp
}

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetFilter() yfilter.YFilter { return bgps.YFilter }

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) SetFilter(yf yfilter.YFilter) { bgps.YFilter = yf }

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetGoName(yname string) string {
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetSegmentPath() string {
    return "bgps"
}

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp" {
        for _, c := range bgps.Bgp {
            if bgps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp{}
        bgps.Bgp = append(bgps.Bgp, child)
        return &bgps.Bgp[len(bgps.Bgp)-1]
    }
    return nil
}

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgps.Bgp {
        children[bgps.Bgp[i].GetSegmentPath()] = &bgps.Bgp[i]
    }
    return children
}

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetBundleName() string { return "cisco_ios_xr" }

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetYangName() string { return "bgps" }

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) SetParent(parent types.Entity) { bgps.parent = parent }

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetParent() types.Entity { return bgps.parent }

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetParentYangName() string { return "redistribution" }

// Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp
// Autonomous system number
type Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Higher 16 bits of 4-byte BGP AS number. The type
    // is interface{} with range: 0..65535.
    Asnxx interface{}

    // This attribute is a key. 2-byte or 4-byte BGP AS number. The type is
    // interface{} with range: 0..4294967295.
    Asnyy interface{}

    // Route Policy name. The type is string.
    Policy interface{}

    // Route type. The type is BgpRedistRoute.
    Type interface{}
}

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetGoName(yname string) string {
    if yname == "asnxx" { return "Asnxx" }
    if yname == "asnyy" { return "Asnyy" }
    if yname == "policy" { return "Policy" }
    if yname == "type" { return "Type" }
    return ""
}

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetSegmentPath() string {
    return "bgp" + "[asnxx='" + fmt.Sprintf("%v", bgp.Asnxx) + "']" + "[asnyy='" + fmt.Sprintf("%v", bgp.Asnyy) + "']"
}

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["asnxx"] = bgp.Asnxx
    leafs["asnyy"] = bgp.Asnyy
    leafs["policy"] = bgp.Policy
    leafs["type"] = bgp.Type
    return leafs
}

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetYangName() string { return "bgp" }

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetParentYangName() string { return "bgps" }

// Rip_Vrfs_Vrf_Redistribution_Isises
// Redistribute IS-IS routes
type Rip_Vrfs_Vrf_Redistribution_Isises struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redistribute IS-IS routes. The type is slice of
    // Rip_Vrfs_Vrf_Redistribution_Isises_Isis.
    Isis []Rip_Vrfs_Vrf_Redistribution_Isises_Isis
}

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetFilter() yfilter.YFilter { return isises.YFilter }

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) SetFilter(yf yfilter.YFilter) { isises.YFilter = yf }

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetGoName(yname string) string {
    if yname == "isis" { return "Isis" }
    return ""
}

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetSegmentPath() string {
    return "isises"
}

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        for _, c := range isises.Isis {
            if isises.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf_Redistribution_Isises_Isis{}
        isises.Isis = append(isises.Isis, child)
        return &isises.Isis[len(isises.Isis)-1]
    }
    return nil
}

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range isises.Isis {
        children[isises.Isis[i].GetSegmentPath()] = &isises.Isis[i]
    }
    return children
}

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetBundleName() string { return "cisco_ios_xr" }

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetYangName() string { return "isises" }

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) SetParent(parent types.Entity) { isises.parent = parent }

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetParent() types.Entity { return isises.parent }

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetParentYangName() string { return "redistribution" }

// Rip_Vrfs_Vrf_Redistribution_Isises_Isis
// Redistribute IS-IS routes
type Rip_Vrfs_Vrf_Redistribution_Isises_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IS-IS instance name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    IsisName interface{}

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is IsisRedistRoute.
    RouteType interface{}
}

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetGoName(yname string) string {
    if yname == "isis-name" { return "IsisName" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "route-type" { return "RouteType" }
    return ""
}

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetSegmentPath() string {
    return "isis" + "[isis-name='" + fmt.Sprintf("%v", isis.IsisName) + "']"
}

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["isis-name"] = isis.IsisName
    leafs["route-policy-name"] = isis.RoutePolicyName
    leafs["route-type"] = isis.RouteType
    return leafs
}

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetYangName() string { return "isis" }

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetParent() types.Entity { return isis.parent }

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetParentYangName() string { return "isises" }

// Rip_Vrfs_Vrf_Redistribution_EigrpS
// Redistribute EIGRP routes
type Rip_Vrfs_Vrf_Redistribution_EigrpS struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redistribute EIGRP routes. The type is slice of
    // Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp.
    Eigrp []Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp
}

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetFilter() yfilter.YFilter { return eigrpS.YFilter }

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) SetFilter(yf yfilter.YFilter) { eigrpS.YFilter = yf }

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetGoName(yname string) string {
    if yname == "eigrp" { return "Eigrp" }
    return ""
}

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetSegmentPath() string {
    return "eigrp-s"
}

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "eigrp" {
        for _, c := range eigrpS.Eigrp {
            if eigrpS.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp{}
        eigrpS.Eigrp = append(eigrpS.Eigrp, child)
        return &eigrpS.Eigrp[len(eigrpS.Eigrp)-1]
    }
    return nil
}

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range eigrpS.Eigrp {
        children[eigrpS.Eigrp[i].GetSegmentPath()] = &eigrpS.Eigrp[i]
    }
    return children
}

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetBundleName() string { return "cisco_ios_xr" }

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetYangName() string { return "eigrp-s" }

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) SetParent(parent types.Entity) { eigrpS.parent = parent }

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetParent() types.Entity { return eigrpS.parent }

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetParentYangName() string { return "redistribution" }

// Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp
// Redistribute EIGRP routes
type Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Autonomous system number. The type is interface{}
    // with range: 1..65535.
    As interface{}

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is DefaultRedistRoute.
    RouteType interface{}
}

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetFilter() yfilter.YFilter { return eigrp.YFilter }

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) SetFilter(yf yfilter.YFilter) { eigrp.YFilter = yf }

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "route-type" { return "RouteType" }
    return ""
}

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetSegmentPath() string {
    return "eigrp" + "[as='" + fmt.Sprintf("%v", eigrp.As) + "']"
}

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = eigrp.As
    leafs["route-policy-name"] = eigrp.RoutePolicyName
    leafs["route-type"] = eigrp.RouteType
    return leafs
}

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetBundleName() string { return "cisco_ios_xr" }

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetYangName() string { return "eigrp" }

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) SetParent(parent types.Entity) { eigrp.parent = parent }

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetParent() types.Entity { return eigrp.parent }

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetParentYangName() string { return "eigrp-s" }

// Rip_Vrfs_Vrf_Redistribution_Static
// Redistribute static routes
// This type is a presence type.
type Rip_Vrfs_Vrf_Redistribution_Static struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is DefaultRedistRoute.
    RouteType interface{}
}

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetFilter() yfilter.YFilter { return static.YFilter }

func (static *Rip_Vrfs_Vrf_Redistribution_Static) SetFilter(yf yfilter.YFilter) { static.YFilter = yf }

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "route-type" { return "RouteType" }
    return ""
}

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetSegmentPath() string {
    return "static"
}

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = static.RoutePolicyName
    leafs["route-type"] = static.RouteType
    return leafs
}

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetBundleName() string { return "cisco_ios_xr" }

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetYangName() string { return "static" }

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (static *Rip_Vrfs_Vrf_Redistribution_Static) SetParent(parent types.Entity) { static.parent = parent }

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetParent() types.Entity { return static.parent }

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetParentYangName() string { return "redistribution" }

// Rip_Vrfs_Vrf_Redistribution_Ospfs
// Redistribute OSPF routes
type Rip_Vrfs_Vrf_Redistribution_Ospfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redistribute OSPF routes. The type is slice of
    // Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf.
    Ospf []Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf
}

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetFilter() yfilter.YFilter { return ospfs.YFilter }

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) SetFilter(yf yfilter.YFilter) { ospfs.YFilter = yf }

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetGoName(yname string) string {
    if yname == "ospf" { return "Ospf" }
    return ""
}

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetSegmentPath() string {
    return "ospfs"
}

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospf" {
        for _, c := range ospfs.Ospf {
            if ospfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf{}
        ospfs.Ospf = append(ospfs.Ospf, child)
        return &ospfs.Ospf[len(ospfs.Ospf)-1]
    }
    return nil
}

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfs.Ospf {
        children[ospfs.Ospf[i].GetSegmentPath()] = &ospfs.Ospf[i]
    }
    return children
}

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetBundleName() string { return "cisco_ios_xr" }

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetYangName() string { return "ospfs" }

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) SetParent(parent types.Entity) { ospfs.parent = parent }

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetParent() types.Entity { return ospfs.parent }

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetParentYangName() string { return "redistribution" }

// Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf
// Redistribute OSPF routes
type Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Process ID for the OSPF instance. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    OspfName interface{}

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Internal routes. The type is bool.
    Internal interface{}

    // External routes. The type is bool.
    External interface{}

    // External route type. The type is interface{} with range: 0..2.
    ExternalType interface{}

    // NSSA External routes. The type is bool.
    NssaExternal interface{}

    // NSSA External route type. The type is interface{} with range: 0..2.
    NssaExternalType interface{}
}

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetGoName(yname string) string {
    if yname == "ospf-name" { return "OspfName" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "internal" { return "Internal" }
    if yname == "external" { return "External" }
    if yname == "external-type" { return "ExternalType" }
    if yname == "nssa-external" { return "NssaExternal" }
    if yname == "nssa-external-type" { return "NssaExternalType" }
    return ""
}

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetSegmentPath() string {
    return "ospf" + "[ospf-name='" + fmt.Sprintf("%v", ospf.OspfName) + "']"
}

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospf-name"] = ospf.OspfName
    leafs["route-policy-name"] = ospf.RoutePolicyName
    leafs["internal"] = ospf.Internal
    leafs["external"] = ospf.External
    leafs["external-type"] = ospf.ExternalType
    leafs["nssa-external"] = ospf.NssaExternal
    leafs["nssa-external-type"] = ospf.NssaExternalType
    return leafs
}

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetYangName() string { return "ospf" }

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetParentYangName() string { return "ospfs" }

// Rip_Vrfs_Vrf_IpDistances
// Table of IP specific administrative distances
type Rip_Vrfs_Vrf_IpDistances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP specific administrative distance. The type is slice of
    // Rip_Vrfs_Vrf_IpDistances_IpDistance.
    IpDistance []Rip_Vrfs_Vrf_IpDistances_IpDistance
}

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetFilter() yfilter.YFilter { return ipDistances.YFilter }

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) SetFilter(yf yfilter.YFilter) { ipDistances.YFilter = yf }

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetGoName(yname string) string {
    if yname == "ip-distance" { return "IpDistance" }
    return ""
}

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetSegmentPath() string {
    return "ip-distances"
}

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ip-distance" {
        for _, c := range ipDistances.IpDistance {
            if ipDistances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf_IpDistances_IpDistance{}
        ipDistances.IpDistance = append(ipDistances.IpDistance, child)
        return &ipDistances.IpDistance[len(ipDistances.IpDistance)-1]
    }
    return nil
}

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipDistances.IpDistance {
        children[ipDistances.IpDistance[i].GetSegmentPath()] = &ipDistances.IpDistance[i]
    }
    return children
}

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetBundleName() string { return "cisco_ios_xr" }

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetYangName() string { return "ip-distances" }

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) SetParent(parent types.Entity) { ipDistances.parent = parent }

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetParent() types.Entity { return ipDistances.parent }

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetParentYangName() string { return "vrf" }

// Rip_Vrfs_Vrf_IpDistances_IpDistance
// IP specific administrative distance
type Rip_Vrfs_Vrf_IpDistances_IpDistance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP Source address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IP address mask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netmask interface{}

    // Administrative distance. The type is interface{} with range: 0..255. This
    // attribute is mandatory.
    Distance interface{}
}

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetFilter() yfilter.YFilter { return ipDistance.YFilter }

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) SetFilter(yf yfilter.YFilter) { ipDistance.YFilter = yf }

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "netmask" { return "Netmask" }
    if yname == "distance" { return "Distance" }
    return ""
}

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetSegmentPath() string {
    return "ip-distance" + "[address='" + fmt.Sprintf("%v", ipDistance.Address) + "']" + "[netmask='" + fmt.Sprintf("%v", ipDistance.Netmask) + "']"
}

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = ipDistance.Address
    leafs["netmask"] = ipDistance.Netmask
    leafs["distance"] = ipDistance.Distance
    return leafs
}

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetBundleName() string { return "cisco_ios_xr" }

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetYangName() string { return "ip-distance" }

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) SetParent(parent types.Entity) { ipDistance.parent = parent }

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetParent() types.Entity { return ipDistance.parent }

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetParentYangName() string { return "ip-distances" }

// Rip_Vrfs_Vrf_Interfaces
// Table of RIP interfaces
type Rip_Vrfs_Vrf_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RIP interface name. The type is slice of Rip_Vrfs_Vrf_Interfaces_Interface.
    Interface []Rip_Vrfs_Vrf_Interfaces_Interface
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetParentYangName() string { return "vrf" }

// Rip_Vrfs_Vrf_Interfaces_Interface
// RIP interface name
type Rip_Vrfs_Vrf_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Send RIP v2 output packets to broadcast address. The type is interface{}.
    BroadcastForV2 interface{}

    // Enable poison reverse. The type is interface{}.
    PoisonReverse interface{}

    // Suppress routing updates on this interface. The type is interface{}.
    Passive interface{}

    // Starts RIP interface configuration. The type is interface{}.
    Enable interface{}

    // Route Policy for outbound routing updates. The type is string.
    PolicyOut interface{}

    // Accept RIP updates with metric 0. The type is interface{}.
    AcceptMetricZero interface{}

    // Route Policy for inbound routing updates. The type is string.
    PolicyIn interface{}

    // Disable split horizon. The type is interface{}.
    SplitHorizonDisable interface{}

    // Authentication keychain and mode.
    Authentication Rip_Vrfs_Vrf_Interfaces_Interface_Authentication

    // SOO community for prefixes learned over this interface.
    SiteOfOrigin Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin

    // RIP versions supported for receiving advertisements.
    ReceiveVersion Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion

    // RIP versions supported for sending advertisements.
    SendVersion Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "broadcast-for-v2" { return "BroadcastForV2" }
    if yname == "poison-reverse" { return "PoisonReverse" }
    if yname == "passive" { return "Passive" }
    if yname == "enable" { return "Enable" }
    if yname == "policy-out" { return "PolicyOut" }
    if yname == "accept-metric-zero" { return "AcceptMetricZero" }
    if yname == "policy-in" { return "PolicyIn" }
    if yname == "split-horizon-disable" { return "SplitHorizonDisable" }
    if yname == "authentication" { return "Authentication" }
    if yname == "site-of-origin" { return "SiteOfOrigin" }
    if yname == "receive-version" { return "ReceiveVersion" }
    if yname == "send-version" { return "SendVersion" }
    return ""
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authentication" {
        return &self.Authentication
    }
    if childYangName == "site-of-origin" {
        return &self.SiteOfOrigin
    }
    if childYangName == "receive-version" {
        return &self.ReceiveVersion
    }
    if childYangName == "send-version" {
        return &self.SendVersion
    }
    return nil
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["authentication"] = &self.Authentication
    children["site-of-origin"] = &self.SiteOfOrigin
    children["receive-version"] = &self.ReceiveVersion
    children["send-version"] = &self.SendVersion
    return children
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["broadcast-for-v2"] = self.BroadcastForV2
    leafs["poison-reverse"] = self.PoisonReverse
    leafs["passive"] = self.Passive
    leafs["enable"] = self.Enable
    leafs["policy-out"] = self.PolicyOut
    leafs["accept-metric-zero"] = self.AcceptMetricZero
    leafs["policy-in"] = self.PolicyIn
    leafs["split-horizon-disable"] = self.SplitHorizonDisable
    return leafs
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Rip_Vrfs_Vrf_Interfaces_Interface_Authentication
// Authentication keychain and mode
// This type is a presence type.
type Rip_Vrfs_Vrf_Interfaces_Interface_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of keychain. The type is string. This attribute is mandatory.
    Keychain interface{}

    // Authentication mode. The type is RipAuthMode. This attribute is mandatory.
    Mode interface{}
}

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetGoName(yname string) string {
    if yname == "keychain" { return "Keychain" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["keychain"] = authentication.Keychain
    leafs["mode"] = authentication.Mode
    return leafs
}

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetYangName() string { return "authentication" }

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetParentYangName() string { return "interface" }

// Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin
// SOO community for prefixes learned over this
// interface
type Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Extended community type. The type is RipExtCommunity.
    Type interface{}

    // AS Number for AS:nn format. The type is interface{} with range: 0..65535.
    AsXx interface{}

    // 32 bit value for AS:nn format. The type is interface{} with range:
    // 0..4294967295.
    AsYy interface{}

    // AS Number Index. The type is interface{} with range: 0..4294967295.
    AsIndex interface{}

    // IPV4 address for IPV4Address:nn format. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // 16bit value for IPV4Address:nn format. The type is interface{} with range:
    // 0..65535.
    AddressIndex interface{}
}

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetFilter() yfilter.YFilter { return siteOfOrigin.YFilter }

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) SetFilter(yf yfilter.YFilter) { siteOfOrigin.YFilter = yf }

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "as-xx" { return "AsXx" }
    if yname == "as-yy" { return "AsYy" }
    if yname == "as-index" { return "AsIndex" }
    if yname == "address" { return "Address" }
    if yname == "address-index" { return "AddressIndex" }
    return ""
}

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetSegmentPath() string {
    return "site-of-origin"
}

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = siteOfOrigin.Type
    leafs["as-xx"] = siteOfOrigin.AsXx
    leafs["as-yy"] = siteOfOrigin.AsYy
    leafs["as-index"] = siteOfOrigin.AsIndex
    leafs["address"] = siteOfOrigin.Address
    leafs["address-index"] = siteOfOrigin.AddressIndex
    return leafs
}

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetBundleName() string { return "cisco_ios_xr" }

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetYangName() string { return "site-of-origin" }

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) SetParent(parent types.Entity) { siteOfOrigin.parent = parent }

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetParent() types.Entity { return siteOfOrigin.parent }

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetParentYangName() string { return "interface" }

// Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion
// RIP versions supported for receiving
// advertisements
type Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Support RIP version 1. The type is bool.
    Version1 interface{}

    // Support RIP version 2. The type is bool. The default value is true.
    Version2 interface{}
}

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetFilter() yfilter.YFilter { return receiveVersion.YFilter }

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) SetFilter(yf yfilter.YFilter) { receiveVersion.YFilter = yf }

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetGoName(yname string) string {
    if yname == "version1" { return "Version1" }
    if yname == "version2" { return "Version2" }
    return ""
}

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetSegmentPath() string {
    return "receive-version"
}

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version1"] = receiveVersion.Version1
    leafs["version2"] = receiveVersion.Version2
    return leafs
}

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetBundleName() string { return "cisco_ios_xr" }

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetYangName() string { return "receive-version" }

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) SetParent(parent types.Entity) { receiveVersion.parent = parent }

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetParent() types.Entity { return receiveVersion.parent }

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetParentYangName() string { return "interface" }

// Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion
// RIP versions supported for sending
// advertisements
type Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Support RIP version 1. The type is bool.
    Version1 interface{}

    // Support RIP version 2. The type is bool. The default value is true.
    Version2 interface{}
}

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetFilter() yfilter.YFilter { return sendVersion.YFilter }

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) SetFilter(yf yfilter.YFilter) { sendVersion.YFilter = yf }

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetGoName(yname string) string {
    if yname == "version1" { return "Version1" }
    if yname == "version2" { return "Version2" }
    return ""
}

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetSegmentPath() string {
    return "send-version"
}

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version1"] = sendVersion.Version1
    leafs["version2"] = sendVersion.Version2
    return leafs
}

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetBundleName() string { return "cisco_ios_xr" }

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetYangName() string { return "send-version" }

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) SetParent(parent types.Entity) { sendVersion.parent = parent }

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetParent() types.Entity { return sendVersion.parent }

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetParentYangName() string { return "interface" }

// Rip_Vrfs_Vrf_Neighbors
// Configure RIP Neighbors
type Rip_Vrfs_Vrf_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor address. The type is slice of Rip_Vrfs_Vrf_Neighbors_Neighbor.
    Neighbor []Rip_Vrfs_Vrf_Neighbors_Neighbor
}

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Rip_Vrfs_Vrf_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetBundleName() string { return "cisco_ios_xr" }

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbors *Rip_Vrfs_Vrf_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetParentYangName() string { return "vrf" }

// Rip_Vrfs_Vrf_Neighbors_Neighbor
// Neighbor address
type Rip_Vrfs_Vrf_Neighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}
}

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-address" { return "NeighborAddress" }
    return ""
}

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[neighbor-address='" + fmt.Sprintf("%v", neighbor.NeighborAddress) + "']"
}

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-address"] = neighbor.NeighborAddress
    return leafs
}

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// Rip_Vrfs_Vrf_Timers
// Various routing timers
// This type is a presence type.
type Rip_Vrfs_Vrf_Timers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interval between updates. The type is interface{} with range: 5..50000.
    // This attribute is mandatory.
    UpdateTimer interface{}

    // Invalid. The type is interface{} with range: 15..200000. This attribute is
    // mandatory.
    InvalidTimer interface{}

    // Holddown. The type is interface{} with range: 15..200000. This attribute is
    // mandatory.
    HolddownTimer interface{}

    // Flush. The type is interface{} with range: 16..250000. This attribute is
    // mandatory.
    FlushTimer interface{}
}

func (timers *Rip_Vrfs_Vrf_Timers) GetFilter() yfilter.YFilter { return timers.YFilter }

func (timers *Rip_Vrfs_Vrf_Timers) SetFilter(yf yfilter.YFilter) { timers.YFilter = yf }

func (timers *Rip_Vrfs_Vrf_Timers) GetGoName(yname string) string {
    if yname == "update-timer" { return "UpdateTimer" }
    if yname == "invalid-timer" { return "InvalidTimer" }
    if yname == "holddown-timer" { return "HolddownTimer" }
    if yname == "flush-timer" { return "FlushTimer" }
    return ""
}

func (timers *Rip_Vrfs_Vrf_Timers) GetSegmentPath() string {
    return "timers"
}

func (timers *Rip_Vrfs_Vrf_Timers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timers *Rip_Vrfs_Vrf_Timers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timers *Rip_Vrfs_Vrf_Timers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["update-timer"] = timers.UpdateTimer
    leafs["invalid-timer"] = timers.InvalidTimer
    leafs["holddown-timer"] = timers.HolddownTimer
    leafs["flush-timer"] = timers.FlushTimer
    return leafs
}

func (timers *Rip_Vrfs_Vrf_Timers) GetBundleName() string { return "cisco_ios_xr" }

func (timers *Rip_Vrfs_Vrf_Timers) GetYangName() string { return "timers" }

func (timers *Rip_Vrfs_Vrf_Timers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timers *Rip_Vrfs_Vrf_Timers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timers *Rip_Vrfs_Vrf_Timers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timers *Rip_Vrfs_Vrf_Timers) SetParent(parent types.Entity) { timers.parent = parent }

func (timers *Rip_Vrfs_Vrf_Timers) GetParent() types.Entity { return timers.parent }

func (timers *Rip_Vrfs_Vrf_Timers) GetParentYangName() string { return "vrf" }

