// This module contains a collection of YANG definitions
// for Cisco IOS-XR policy-repository package configuration.
// 
// This module contains definitions
// for the following management objects:
//   routing-policy: Routing policy configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package policy_repository_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package policy_repository_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-policy-repository-cfg routing-policy}", reflect.TypeOf(RoutingPolicy{}))
    ydk.RegisterEntity("Cisco-IOS-XR-policy-repository-cfg:routing-policy", reflect.TypeOf(RoutingPolicy{}))
}

// RoutingPolicy
// Routing policy configuration
type RoutingPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set exit under RPL config to abort. The type is interface{}.
    SetExitAsAbort interface{}

    // 'emacs' or 'vim' or 'nano'. The type is string.
    Editor interface{}

    // All configured policies.
    RoutePolicies RoutingPolicy_RoutePolicies

    // All configured sets.
    Sets RoutingPolicy_Sets

    // Limits for Routing Policy.
    Limits RoutingPolicy_Limits
}

func (routingPolicy *RoutingPolicy) GetFilter() yfilter.YFilter { return routingPolicy.YFilter }

func (routingPolicy *RoutingPolicy) SetFilter(yf yfilter.YFilter) { routingPolicy.YFilter = yf }

func (routingPolicy *RoutingPolicy) GetGoName(yname string) string {
    if yname == "set-exit-as-abort" { return "SetExitAsAbort" }
    if yname == "editor" { return "Editor" }
    if yname == "route-policies" { return "RoutePolicies" }
    if yname == "sets" { return "Sets" }
    if yname == "limits" { return "Limits" }
    return ""
}

func (routingPolicy *RoutingPolicy) GetSegmentPath() string {
    return "Cisco-IOS-XR-policy-repository-cfg:routing-policy"
}

func (routingPolicy *RoutingPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-policies" {
        return &routingPolicy.RoutePolicies
    }
    if childYangName == "sets" {
        return &routingPolicy.Sets
    }
    if childYangName == "limits" {
        return &routingPolicy.Limits
    }
    return nil
}

func (routingPolicy *RoutingPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["route-policies"] = &routingPolicy.RoutePolicies
    children["sets"] = &routingPolicy.Sets
    children["limits"] = &routingPolicy.Limits
    return children
}

func (routingPolicy *RoutingPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-exit-as-abort"] = routingPolicy.SetExitAsAbort
    leafs["editor"] = routingPolicy.Editor
    return leafs
}

func (routingPolicy *RoutingPolicy) GetBundleName() string { return "cisco_ios_xr" }

func (routingPolicy *RoutingPolicy) GetYangName() string { return "routing-policy" }

func (routingPolicy *RoutingPolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routingPolicy *RoutingPolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routingPolicy *RoutingPolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routingPolicy *RoutingPolicy) SetParent(parent types.Entity) { routingPolicy.parent = parent }

func (routingPolicy *RoutingPolicy) GetParent() types.Entity { return routingPolicy.parent }

func (routingPolicy *RoutingPolicy) GetParentYangName() string { return "Cisco-IOS-XR-policy-repository-cfg" }

// RoutingPolicy_RoutePolicies
// All configured policies
type RoutingPolicy_RoutePolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual policy. The type is slice of
    // RoutingPolicy_RoutePolicies_RoutePolicy.
    RoutePolicy []RoutingPolicy_RoutePolicies_RoutePolicy
}

func (routePolicies *RoutingPolicy_RoutePolicies) GetFilter() yfilter.YFilter { return routePolicies.YFilter }

func (routePolicies *RoutingPolicy_RoutePolicies) SetFilter(yf yfilter.YFilter) { routePolicies.YFilter = yf }

func (routePolicies *RoutingPolicy_RoutePolicies) GetGoName(yname string) string {
    if yname == "route-policy" { return "RoutePolicy" }
    return ""
}

func (routePolicies *RoutingPolicy_RoutePolicies) GetSegmentPath() string {
    return "route-policies"
}

func (routePolicies *RoutingPolicy_RoutePolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-policy" {
        for _, c := range routePolicies.RoutePolicy {
            if routePolicies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_RoutePolicies_RoutePolicy{}
        routePolicies.RoutePolicy = append(routePolicies.RoutePolicy, child)
        return &routePolicies.RoutePolicy[len(routePolicies.RoutePolicy)-1]
    }
    return nil
}

func (routePolicies *RoutingPolicy_RoutePolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routePolicies.RoutePolicy {
        children[routePolicies.RoutePolicy[i].GetSegmentPath()] = &routePolicies.RoutePolicy[i]
    }
    return children
}

func (routePolicies *RoutingPolicy_RoutePolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routePolicies *RoutingPolicy_RoutePolicies) GetBundleName() string { return "cisco_ios_xr" }

func (routePolicies *RoutingPolicy_RoutePolicies) GetYangName() string { return "route-policies" }

func (routePolicies *RoutingPolicy_RoutePolicies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routePolicies *RoutingPolicy_RoutePolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routePolicies *RoutingPolicy_RoutePolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routePolicies *RoutingPolicy_RoutePolicies) SetParent(parent types.Entity) { routePolicies.parent = parent }

func (routePolicies *RoutingPolicy_RoutePolicies) GetParent() types.Entity { return routePolicies.parent }

func (routePolicies *RoutingPolicy_RoutePolicies) GetParentYangName() string { return "routing-policy" }

// RoutingPolicy_RoutePolicies_RoutePolicy
// Information about an individual policy
type RoutingPolicy_RoutePolicies_RoutePolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Route policy name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    RoutePolicyName interface{}

    // policy statements. The type is string. This attribute is mandatory.
    RplRoutePolicy interface{}
}

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetFilter() yfilter.YFilter { return routePolicy.YFilter }

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) SetFilter(yf yfilter.YFilter) { routePolicy.YFilter = yf }

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "rpl-route-policy" { return "RplRoutePolicy" }
    return ""
}

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetSegmentPath() string {
    return "route-policy" + "[route-policy-name='" + fmt.Sprintf("%v", routePolicy.RoutePolicyName) + "']"
}

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = routePolicy.RoutePolicyName
    leafs["rpl-route-policy"] = routePolicy.RplRoutePolicy
    return leafs
}

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetYangName() string { return "route-policy" }

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) SetParent(parent types.Entity) { routePolicy.parent = parent }

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetParent() types.Entity { return routePolicy.parent }

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetParentYangName() string { return "route-policies" }

// RoutingPolicy_Sets
// All configured sets
type RoutingPolicy_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about Etag sets.
    PrependEtagSets RoutingPolicy_Sets_PrependEtagSets

    // Information about Prefix sets.
    PrefixSets RoutingPolicy_Sets_PrefixSets

    // Information about Large Community sets.
    LargeCommunitySets RoutingPolicy_Sets_LargeCommunitySets

    // Information about Large Community sets.
    PrependLargeCommunitySets RoutingPolicy_Sets_PrependLargeCommunitySets

    // Information about Etag sets.
    AppendEtagSets RoutingPolicy_Sets_AppendEtagSets

    // Information about Etag sets.
    RemoveEtagSets RoutingPolicy_Sets_RemoveEtagSets

    // Information about Large Community sets.
    RemoveLargeCommunitySets RoutingPolicy_Sets_RemoveLargeCommunitySets

    // Information about Mac sets.
    MacSets RoutingPolicy_Sets_MacSets

    // Information about Opaque sets.
    ExtendedCommunityOpaqueSets RoutingPolicy_Sets_ExtendedCommunityOpaqueSets

    // Information about Mac sets.
    PrependMacSets RoutingPolicy_Sets_PrependMacSets

    // Information about OSPF Area sets.
    OspfAreaSets RoutingPolicy_Sets_OspfAreaSets

    // Information about Mac sets.
    AppendMacSets RoutingPolicy_Sets_AppendMacSets

    // Information about Cost sets.
    ExtendedCommunityCostSets RoutingPolicy_Sets_ExtendedCommunityCostSets

    // Information about Mac sets.
    RemoveMacSets RoutingPolicy_Sets_RemoveMacSets

    // Information about SOO sets.
    ExtendedCommunitySooSets RoutingPolicy_Sets_ExtendedCommunitySooSets

    // Information about Esi sets.
    EsiSets RoutingPolicy_Sets_EsiSets

    // Information about Esi sets.
    PrependEsiSets RoutingPolicy_Sets_PrependEsiSets

    // Information about Esi sets.
    AppendEsiSets RoutingPolicy_Sets_AppendEsiSets

    // Information about Esi sets.
    RemoveEsiSets RoutingPolicy_Sets_RemoveEsiSets

    // Information about SegNH sets.
    ExtendedCommunitySegNhSets RoutingPolicy_Sets_ExtendedCommunitySegNhSets

    // Information about RD sets.
    RdSets RoutingPolicy_Sets_RdSets

    // Information about PolicyGlobal sets.
    PolicyGlobalSetTable RoutingPolicy_Sets_PolicyGlobalSetTable

    // Information about Large Community sets.
    AppendLargeCommunitySets RoutingPolicy_Sets_AppendLargeCommunitySets

    // Information about Bandwidth sets.
    ExtendedCommunityBandwidthSets RoutingPolicy_Sets_ExtendedCommunityBandwidthSets

    // Information about Community sets.
    CommunitySets RoutingPolicy_Sets_CommunitySets

    // Information about AS Path sets.
    AsPathSets RoutingPolicy_Sets_AsPathSets

    // Information about Tag sets.
    TagSets RoutingPolicy_Sets_TagSets

    // Information about Etag sets.
    EtagSets RoutingPolicy_Sets_EtagSets

    // Information about RT sets.
    ExtendedCommunityRtSets RoutingPolicy_Sets_ExtendedCommunityRtSets
}

func (sets *RoutingPolicy_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets) GetGoName(yname string) string {
    if yname == "prepend-etag-sets" { return "PrependEtagSets" }
    if yname == "prefix-sets" { return "PrefixSets" }
    if yname == "large-community-sets" { return "LargeCommunitySets" }
    if yname == "prepend-large-community-sets" { return "PrependLargeCommunitySets" }
    if yname == "append-etag-sets" { return "AppendEtagSets" }
    if yname == "remove-etag-sets" { return "RemoveEtagSets" }
    if yname == "remove-large-community-sets" { return "RemoveLargeCommunitySets" }
    if yname == "mac-sets" { return "MacSets" }
    if yname == "extended-community-opaque-sets" { return "ExtendedCommunityOpaqueSets" }
    if yname == "prepend-mac-sets" { return "PrependMacSets" }
    if yname == "ospf-area-sets" { return "OspfAreaSets" }
    if yname == "append-mac-sets" { return "AppendMacSets" }
    if yname == "extended-community-cost-sets" { return "ExtendedCommunityCostSets" }
    if yname == "remove-mac-sets" { return "RemoveMacSets" }
    if yname == "extended-community-soo-sets" { return "ExtendedCommunitySooSets" }
    if yname == "esi-sets" { return "EsiSets" }
    if yname == "prepend-esi-sets" { return "PrependEsiSets" }
    if yname == "append-esi-sets" { return "AppendEsiSets" }
    if yname == "remove-esi-sets" { return "RemoveEsiSets" }
    if yname == "extended-community-seg-nh-sets" { return "ExtendedCommunitySegNhSets" }
    if yname == "rd-sets" { return "RdSets" }
    if yname == "policy-global-set-table" { return "PolicyGlobalSetTable" }
    if yname == "append-large-community-sets" { return "AppendLargeCommunitySets" }
    if yname == "extended-community-bandwidth-sets" { return "ExtendedCommunityBandwidthSets" }
    if yname == "community-sets" { return "CommunitySets" }
    if yname == "as-path-sets" { return "AsPathSets" }
    if yname == "tag-sets" { return "TagSets" }
    if yname == "etag-sets" { return "EtagSets" }
    if yname == "extended-community-rt-sets" { return "ExtendedCommunityRtSets" }
    return ""
}

func (sets *RoutingPolicy_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prepend-etag-sets" {
        return &sets.PrependEtagSets
    }
    if childYangName == "prefix-sets" {
        return &sets.PrefixSets
    }
    if childYangName == "large-community-sets" {
        return &sets.LargeCommunitySets
    }
    if childYangName == "prepend-large-community-sets" {
        return &sets.PrependLargeCommunitySets
    }
    if childYangName == "append-etag-sets" {
        return &sets.AppendEtagSets
    }
    if childYangName == "remove-etag-sets" {
        return &sets.RemoveEtagSets
    }
    if childYangName == "remove-large-community-sets" {
        return &sets.RemoveLargeCommunitySets
    }
    if childYangName == "mac-sets" {
        return &sets.MacSets
    }
    if childYangName == "extended-community-opaque-sets" {
        return &sets.ExtendedCommunityOpaqueSets
    }
    if childYangName == "prepend-mac-sets" {
        return &sets.PrependMacSets
    }
    if childYangName == "ospf-area-sets" {
        return &sets.OspfAreaSets
    }
    if childYangName == "append-mac-sets" {
        return &sets.AppendMacSets
    }
    if childYangName == "extended-community-cost-sets" {
        return &sets.ExtendedCommunityCostSets
    }
    if childYangName == "remove-mac-sets" {
        return &sets.RemoveMacSets
    }
    if childYangName == "extended-community-soo-sets" {
        return &sets.ExtendedCommunitySooSets
    }
    if childYangName == "esi-sets" {
        return &sets.EsiSets
    }
    if childYangName == "prepend-esi-sets" {
        return &sets.PrependEsiSets
    }
    if childYangName == "append-esi-sets" {
        return &sets.AppendEsiSets
    }
    if childYangName == "remove-esi-sets" {
        return &sets.RemoveEsiSets
    }
    if childYangName == "extended-community-seg-nh-sets" {
        return &sets.ExtendedCommunitySegNhSets
    }
    if childYangName == "rd-sets" {
        return &sets.RdSets
    }
    if childYangName == "policy-global-set-table" {
        return &sets.PolicyGlobalSetTable
    }
    if childYangName == "append-large-community-sets" {
        return &sets.AppendLargeCommunitySets
    }
    if childYangName == "extended-community-bandwidth-sets" {
        return &sets.ExtendedCommunityBandwidthSets
    }
    if childYangName == "community-sets" {
        return &sets.CommunitySets
    }
    if childYangName == "as-path-sets" {
        return &sets.AsPathSets
    }
    if childYangName == "tag-sets" {
        return &sets.TagSets
    }
    if childYangName == "etag-sets" {
        return &sets.EtagSets
    }
    if childYangName == "extended-community-rt-sets" {
        return &sets.ExtendedCommunityRtSets
    }
    return nil
}

func (sets *RoutingPolicy_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prepend-etag-sets"] = &sets.PrependEtagSets
    children["prefix-sets"] = &sets.PrefixSets
    children["large-community-sets"] = &sets.LargeCommunitySets
    children["prepend-large-community-sets"] = &sets.PrependLargeCommunitySets
    children["append-etag-sets"] = &sets.AppendEtagSets
    children["remove-etag-sets"] = &sets.RemoveEtagSets
    children["remove-large-community-sets"] = &sets.RemoveLargeCommunitySets
    children["mac-sets"] = &sets.MacSets
    children["extended-community-opaque-sets"] = &sets.ExtendedCommunityOpaqueSets
    children["prepend-mac-sets"] = &sets.PrependMacSets
    children["ospf-area-sets"] = &sets.OspfAreaSets
    children["append-mac-sets"] = &sets.AppendMacSets
    children["extended-community-cost-sets"] = &sets.ExtendedCommunityCostSets
    children["remove-mac-sets"] = &sets.RemoveMacSets
    children["extended-community-soo-sets"] = &sets.ExtendedCommunitySooSets
    children["esi-sets"] = &sets.EsiSets
    children["prepend-esi-sets"] = &sets.PrependEsiSets
    children["append-esi-sets"] = &sets.AppendEsiSets
    children["remove-esi-sets"] = &sets.RemoveEsiSets
    children["extended-community-seg-nh-sets"] = &sets.ExtendedCommunitySegNhSets
    children["rd-sets"] = &sets.RdSets
    children["policy-global-set-table"] = &sets.PolicyGlobalSetTable
    children["append-large-community-sets"] = &sets.AppendLargeCommunitySets
    children["extended-community-bandwidth-sets"] = &sets.ExtendedCommunityBandwidthSets
    children["community-sets"] = &sets.CommunitySets
    children["as-path-sets"] = &sets.AsPathSets
    children["tag-sets"] = &sets.TagSets
    children["etag-sets"] = &sets.EtagSets
    children["extended-community-rt-sets"] = &sets.ExtendedCommunityRtSets
    return children
}

func (sets *RoutingPolicy_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets) GetParentYangName() string { return "routing-policy" }

// RoutingPolicy_Sets_PrependEtagSets
// Information about Etag sets
type RoutingPolicy_Sets_PrependEtagSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prepend the entries to the existing set. The type is slice of
    // RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet.
    PrependEtagSet []RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet
}

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) GetFilter() yfilter.YFilter { return prependEtagSets.YFilter }

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) SetFilter(yf yfilter.YFilter) { prependEtagSets.YFilter = yf }

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) GetGoName(yname string) string {
    if yname == "prepend-etag-set" { return "PrependEtagSet" }
    return ""
}

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) GetSegmentPath() string {
    return "prepend-etag-sets"
}

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prepend-etag-set" {
        for _, c := range prependEtagSets.PrependEtagSet {
            if prependEtagSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet{}
        prependEtagSets.PrependEtagSet = append(prependEtagSets.PrependEtagSet, child)
        return &prependEtagSets.PrependEtagSet[len(prependEtagSets.PrependEtagSet)-1]
    }
    return nil
}

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prependEtagSets.PrependEtagSet {
        children[prependEtagSets.PrependEtagSet[i].GetSegmentPath()] = &prependEtagSets.PrependEtagSet[i]
    }
    return children
}

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) GetBundleName() string { return "cisco_ios_xr" }

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) GetYangName() string { return "prepend-etag-sets" }

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) SetParent(parent types.Entity) { prependEtagSets.parent = parent }

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) GetParent() types.Entity { return prependEtagSets.parent }

func (prependEtagSets *RoutingPolicy_Sets_PrependEtagSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet
// Prepend the entries to the existing set
type RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Etag Set. The type is string. This attribute is mandatory.
    EtagSetAsText interface{}
}

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) GetFilter() yfilter.YFilter { return prependEtagSet.YFilter }

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) SetFilter(yf yfilter.YFilter) { prependEtagSet.YFilter = yf }

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "etag-set-as-text" { return "EtagSetAsText" }
    return ""
}

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) GetSegmentPath() string {
    return "prepend-etag-set" + "[set-name='" + fmt.Sprintf("%v", prependEtagSet.SetName) + "']"
}

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = prependEtagSet.SetName
    leafs["etag-set-as-text"] = prependEtagSet.EtagSetAsText
    return leafs
}

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) GetBundleName() string { return "cisco_ios_xr" }

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) GetYangName() string { return "prepend-etag-set" }

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) SetParent(parent types.Entity) { prependEtagSet.parent = parent }

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) GetParent() types.Entity { return prependEtagSet.parent }

func (prependEtagSet *RoutingPolicy_Sets_PrependEtagSets_PrependEtagSet) GetParentYangName() string { return "prepend-etag-sets" }

// RoutingPolicy_Sets_PrefixSets
// Information about Prefix sets
type RoutingPolicy_Sets_PrefixSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_PrefixSets_PrefixSet.
    PrefixSet []RoutingPolicy_Sets_PrefixSets_PrefixSet
}

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetFilter() yfilter.YFilter { return prefixSets.YFilter }

func (prefixSets *RoutingPolicy_Sets_PrefixSets) SetFilter(yf yfilter.YFilter) { prefixSets.YFilter = yf }

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetGoName(yname string) string {
    if yname == "prefix-set" { return "PrefixSet" }
    return ""
}

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetSegmentPath() string {
    return "prefix-sets"
}

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-set" {
        for _, c := range prefixSets.PrefixSet {
            if prefixSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_PrefixSets_PrefixSet{}
        prefixSets.PrefixSet = append(prefixSets.PrefixSet, child)
        return &prefixSets.PrefixSet[len(prefixSets.PrefixSet)-1]
    }
    return nil
}

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixSets.PrefixSet {
        children[prefixSets.PrefixSet[i].GetSegmentPath()] = &prefixSets.PrefixSet[i]
    }
    return children
}

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetBundleName() string { return "cisco_ios_xr" }

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetYangName() string { return "prefix-sets" }

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixSets *RoutingPolicy_Sets_PrefixSets) SetParent(parent types.Entity) { prefixSets.parent = parent }

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetParent() types.Entity { return prefixSets.parent }

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_PrefixSets_PrefixSet
// Information about an individual set
type RoutingPolicy_Sets_PrefixSets_PrefixSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // prefix statements. The type is string. This attribute is mandatory.
    RplPrefixSet interface{}
}

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetFilter() yfilter.YFilter { return prefixSet.YFilter }

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) SetFilter(yf yfilter.YFilter) { prefixSet.YFilter = yf }

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "rpl-prefix-set" { return "RplPrefixSet" }
    return ""
}

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetSegmentPath() string {
    return "prefix-set" + "[set-name='" + fmt.Sprintf("%v", prefixSet.SetName) + "']"
}

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = prefixSet.SetName
    leafs["rpl-prefix-set"] = prefixSet.RplPrefixSet
    return leafs
}

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetBundleName() string { return "cisco_ios_xr" }

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetYangName() string { return "prefix-set" }

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) SetParent(parent types.Entity) { prefixSet.parent = parent }

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetParent() types.Entity { return prefixSet.parent }

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetParentYangName() string { return "prefix-sets" }

// RoutingPolicy_Sets_LargeCommunitySets
// Information about Large Community sets
type RoutingPolicy_Sets_LargeCommunitySets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet.
    LargeCommunitySet []RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet
}

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetFilter() yfilter.YFilter { return largeCommunitySets.YFilter }

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) SetFilter(yf yfilter.YFilter) { largeCommunitySets.YFilter = yf }

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetGoName(yname string) string {
    if yname == "large-community-set" { return "LargeCommunitySet" }
    return ""
}

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetSegmentPath() string {
    return "large-community-sets"
}

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "large-community-set" {
        for _, c := range largeCommunitySets.LargeCommunitySet {
            if largeCommunitySets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet{}
        largeCommunitySets.LargeCommunitySet = append(largeCommunitySets.LargeCommunitySet, child)
        return &largeCommunitySets.LargeCommunitySet[len(largeCommunitySets.LargeCommunitySet)-1]
    }
    return nil
}

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range largeCommunitySets.LargeCommunitySet {
        children[largeCommunitySets.LargeCommunitySet[i].GetSegmentPath()] = &largeCommunitySets.LargeCommunitySet[i]
    }
    return children
}

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetBundleName() string { return "cisco_ios_xr" }

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetYangName() string { return "large-community-sets" }

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) SetParent(parent types.Entity) { largeCommunitySets.parent = parent }

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetParent() types.Entity { return largeCommunitySets.parent }

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet
// Information about an individual set
type RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Large Community Set. The type is string. This attribute is mandatory.
    LargeCommunitySetAsText interface{}
}

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetFilter() yfilter.YFilter { return largeCommunitySet.YFilter }

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) SetFilter(yf yfilter.YFilter) { largeCommunitySet.YFilter = yf }

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "large-community-set-as-text" { return "LargeCommunitySetAsText" }
    return ""
}

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetSegmentPath() string {
    return "large-community-set" + "[set-name='" + fmt.Sprintf("%v", largeCommunitySet.SetName) + "']"
}

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = largeCommunitySet.SetName
    leafs["large-community-set-as-text"] = largeCommunitySet.LargeCommunitySetAsText
    return leafs
}

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetBundleName() string { return "cisco_ios_xr" }

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetYangName() string { return "large-community-set" }

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) SetParent(parent types.Entity) { largeCommunitySet.parent = parent }

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetParent() types.Entity { return largeCommunitySet.parent }

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetParentYangName() string { return "large-community-sets" }

// RoutingPolicy_Sets_PrependLargeCommunitySets
// Information about Large Community sets
type RoutingPolicy_Sets_PrependLargeCommunitySets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prepend the entries to the existing set. The type is slice of
    // RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet.
    PrependLargeCommunitySet []RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet
}

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) GetFilter() yfilter.YFilter { return prependLargeCommunitySets.YFilter }

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) SetFilter(yf yfilter.YFilter) { prependLargeCommunitySets.YFilter = yf }

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) GetGoName(yname string) string {
    if yname == "prepend-large-community-set" { return "PrependLargeCommunitySet" }
    return ""
}

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) GetSegmentPath() string {
    return "prepend-large-community-sets"
}

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prepend-large-community-set" {
        for _, c := range prependLargeCommunitySets.PrependLargeCommunitySet {
            if prependLargeCommunitySets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet{}
        prependLargeCommunitySets.PrependLargeCommunitySet = append(prependLargeCommunitySets.PrependLargeCommunitySet, child)
        return &prependLargeCommunitySets.PrependLargeCommunitySet[len(prependLargeCommunitySets.PrependLargeCommunitySet)-1]
    }
    return nil
}

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prependLargeCommunitySets.PrependLargeCommunitySet {
        children[prependLargeCommunitySets.PrependLargeCommunitySet[i].GetSegmentPath()] = &prependLargeCommunitySets.PrependLargeCommunitySet[i]
    }
    return children
}

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) GetBundleName() string { return "cisco_ios_xr" }

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) GetYangName() string { return "prepend-large-community-sets" }

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) SetParent(parent types.Entity) { prependLargeCommunitySets.parent = parent }

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) GetParent() types.Entity { return prependLargeCommunitySets.parent }

func (prependLargeCommunitySets *RoutingPolicy_Sets_PrependLargeCommunitySets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet
// Prepend the entries to the existing set
type RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Large Community Set. The type is string. This attribute is mandatory.
    LargeCommunitySetAsText interface{}
}

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) GetFilter() yfilter.YFilter { return prependLargeCommunitySet.YFilter }

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) SetFilter(yf yfilter.YFilter) { prependLargeCommunitySet.YFilter = yf }

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "large-community-set-as-text" { return "LargeCommunitySetAsText" }
    return ""
}

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) GetSegmentPath() string {
    return "prepend-large-community-set" + "[set-name='" + fmt.Sprintf("%v", prependLargeCommunitySet.SetName) + "']"
}

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = prependLargeCommunitySet.SetName
    leafs["large-community-set-as-text"] = prependLargeCommunitySet.LargeCommunitySetAsText
    return leafs
}

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) GetBundleName() string { return "cisco_ios_xr" }

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) GetYangName() string { return "prepend-large-community-set" }

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) SetParent(parent types.Entity) { prependLargeCommunitySet.parent = parent }

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) GetParent() types.Entity { return prependLargeCommunitySet.parent }

func (prependLargeCommunitySet *RoutingPolicy_Sets_PrependLargeCommunitySets_PrependLargeCommunitySet) GetParentYangName() string { return "prepend-large-community-sets" }

// RoutingPolicy_Sets_AppendEtagSets
// Information about Etag sets
type RoutingPolicy_Sets_AppendEtagSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Append the entries to the existing set. The type is slice of
    // RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet.
    AppendEtagSet []RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet
}

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) GetFilter() yfilter.YFilter { return appendEtagSets.YFilter }

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) SetFilter(yf yfilter.YFilter) { appendEtagSets.YFilter = yf }

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) GetGoName(yname string) string {
    if yname == "append-etag-set" { return "AppendEtagSet" }
    return ""
}

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) GetSegmentPath() string {
    return "append-etag-sets"
}

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "append-etag-set" {
        for _, c := range appendEtagSets.AppendEtagSet {
            if appendEtagSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet{}
        appendEtagSets.AppendEtagSet = append(appendEtagSets.AppendEtagSet, child)
        return &appendEtagSets.AppendEtagSet[len(appendEtagSets.AppendEtagSet)-1]
    }
    return nil
}

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range appendEtagSets.AppendEtagSet {
        children[appendEtagSets.AppendEtagSet[i].GetSegmentPath()] = &appendEtagSets.AppendEtagSet[i]
    }
    return children
}

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) GetBundleName() string { return "cisco_ios_xr" }

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) GetYangName() string { return "append-etag-sets" }

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) SetParent(parent types.Entity) { appendEtagSets.parent = parent }

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) GetParent() types.Entity { return appendEtagSets.parent }

func (appendEtagSets *RoutingPolicy_Sets_AppendEtagSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet
// Append the entries to the existing set
type RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Etag Set. The type is string. This attribute is mandatory.
    EtagSetAsText interface{}
}

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) GetFilter() yfilter.YFilter { return appendEtagSet.YFilter }

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) SetFilter(yf yfilter.YFilter) { appendEtagSet.YFilter = yf }

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "etag-set-as-text" { return "EtagSetAsText" }
    return ""
}

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) GetSegmentPath() string {
    return "append-etag-set" + "[set-name='" + fmt.Sprintf("%v", appendEtagSet.SetName) + "']"
}

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = appendEtagSet.SetName
    leafs["etag-set-as-text"] = appendEtagSet.EtagSetAsText
    return leafs
}

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) GetBundleName() string { return "cisco_ios_xr" }

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) GetYangName() string { return "append-etag-set" }

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) SetParent(parent types.Entity) { appendEtagSet.parent = parent }

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) GetParent() types.Entity { return appendEtagSet.parent }

func (appendEtagSet *RoutingPolicy_Sets_AppendEtagSets_AppendEtagSet) GetParentYangName() string { return "append-etag-sets" }

// RoutingPolicy_Sets_RemoveEtagSets
// Information about Etag sets
type RoutingPolicy_Sets_RemoveEtagSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Remove the entries from the existing set. The type is slice of
    // RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet.
    RemoveEtagSet []RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet
}

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) GetFilter() yfilter.YFilter { return removeEtagSets.YFilter }

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) SetFilter(yf yfilter.YFilter) { removeEtagSets.YFilter = yf }

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) GetGoName(yname string) string {
    if yname == "remove-etag-set" { return "RemoveEtagSet" }
    return ""
}

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) GetSegmentPath() string {
    return "remove-etag-sets"
}

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remove-etag-set" {
        for _, c := range removeEtagSets.RemoveEtagSet {
            if removeEtagSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet{}
        removeEtagSets.RemoveEtagSet = append(removeEtagSets.RemoveEtagSet, child)
        return &removeEtagSets.RemoveEtagSet[len(removeEtagSets.RemoveEtagSet)-1]
    }
    return nil
}

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range removeEtagSets.RemoveEtagSet {
        children[removeEtagSets.RemoveEtagSet[i].GetSegmentPath()] = &removeEtagSets.RemoveEtagSet[i]
    }
    return children
}

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) GetBundleName() string { return "cisco_ios_xr" }

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) GetYangName() string { return "remove-etag-sets" }

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) SetParent(parent types.Entity) { removeEtagSets.parent = parent }

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) GetParent() types.Entity { return removeEtagSets.parent }

func (removeEtagSets *RoutingPolicy_Sets_RemoveEtagSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet
// Remove the entries from the existing set
type RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Etag Set. The type is string. This attribute is mandatory.
    EtagSetAsText interface{}
}

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) GetFilter() yfilter.YFilter { return removeEtagSet.YFilter }

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) SetFilter(yf yfilter.YFilter) { removeEtagSet.YFilter = yf }

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "etag-set-as-text" { return "EtagSetAsText" }
    return ""
}

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) GetSegmentPath() string {
    return "remove-etag-set" + "[set-name='" + fmt.Sprintf("%v", removeEtagSet.SetName) + "']"
}

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = removeEtagSet.SetName
    leafs["etag-set-as-text"] = removeEtagSet.EtagSetAsText
    return leafs
}

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) GetBundleName() string { return "cisco_ios_xr" }

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) GetYangName() string { return "remove-etag-set" }

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) SetParent(parent types.Entity) { removeEtagSet.parent = parent }

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) GetParent() types.Entity { return removeEtagSet.parent }

func (removeEtagSet *RoutingPolicy_Sets_RemoveEtagSets_RemoveEtagSet) GetParentYangName() string { return "remove-etag-sets" }

// RoutingPolicy_Sets_RemoveLargeCommunitySets
// Information about Large Community sets
type RoutingPolicy_Sets_RemoveLargeCommunitySets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Remove the entries from the existing set. The type is slice of
    // RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet.
    RemoveLargeCommunitySet []RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet
}

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) GetFilter() yfilter.YFilter { return removeLargeCommunitySets.YFilter }

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) SetFilter(yf yfilter.YFilter) { removeLargeCommunitySets.YFilter = yf }

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) GetGoName(yname string) string {
    if yname == "remove-large-community-set" { return "RemoveLargeCommunitySet" }
    return ""
}

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) GetSegmentPath() string {
    return "remove-large-community-sets"
}

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remove-large-community-set" {
        for _, c := range removeLargeCommunitySets.RemoveLargeCommunitySet {
            if removeLargeCommunitySets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet{}
        removeLargeCommunitySets.RemoveLargeCommunitySet = append(removeLargeCommunitySets.RemoveLargeCommunitySet, child)
        return &removeLargeCommunitySets.RemoveLargeCommunitySet[len(removeLargeCommunitySets.RemoveLargeCommunitySet)-1]
    }
    return nil
}

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range removeLargeCommunitySets.RemoveLargeCommunitySet {
        children[removeLargeCommunitySets.RemoveLargeCommunitySet[i].GetSegmentPath()] = &removeLargeCommunitySets.RemoveLargeCommunitySet[i]
    }
    return children
}

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) GetBundleName() string { return "cisco_ios_xr" }

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) GetYangName() string { return "remove-large-community-sets" }

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) SetParent(parent types.Entity) { removeLargeCommunitySets.parent = parent }

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) GetParent() types.Entity { return removeLargeCommunitySets.parent }

func (removeLargeCommunitySets *RoutingPolicy_Sets_RemoveLargeCommunitySets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet
// Remove the entries from the existing set
type RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Large Community Set. The type is string. This attribute is mandatory.
    LargeCommunitySetAsText interface{}
}

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) GetFilter() yfilter.YFilter { return removeLargeCommunitySet.YFilter }

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) SetFilter(yf yfilter.YFilter) { removeLargeCommunitySet.YFilter = yf }

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "large-community-set-as-text" { return "LargeCommunitySetAsText" }
    return ""
}

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) GetSegmentPath() string {
    return "remove-large-community-set" + "[set-name='" + fmt.Sprintf("%v", removeLargeCommunitySet.SetName) + "']"
}

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = removeLargeCommunitySet.SetName
    leafs["large-community-set-as-text"] = removeLargeCommunitySet.LargeCommunitySetAsText
    return leafs
}

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) GetBundleName() string { return "cisco_ios_xr" }

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) GetYangName() string { return "remove-large-community-set" }

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) SetParent(parent types.Entity) { removeLargeCommunitySet.parent = parent }

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) GetParent() types.Entity { return removeLargeCommunitySet.parent }

func (removeLargeCommunitySet *RoutingPolicy_Sets_RemoveLargeCommunitySets_RemoveLargeCommunitySet) GetParentYangName() string { return "remove-large-community-sets" }

// RoutingPolicy_Sets_MacSets
// Information about Mac sets
type RoutingPolicy_Sets_MacSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_MacSets_MacSet.
    MacSet []RoutingPolicy_Sets_MacSets_MacSet
}

func (macSets *RoutingPolicy_Sets_MacSets) GetFilter() yfilter.YFilter { return macSets.YFilter }

func (macSets *RoutingPolicy_Sets_MacSets) SetFilter(yf yfilter.YFilter) { macSets.YFilter = yf }

func (macSets *RoutingPolicy_Sets_MacSets) GetGoName(yname string) string {
    if yname == "mac-set" { return "MacSet" }
    return ""
}

func (macSets *RoutingPolicy_Sets_MacSets) GetSegmentPath() string {
    return "mac-sets"
}

func (macSets *RoutingPolicy_Sets_MacSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-set" {
        for _, c := range macSets.MacSet {
            if macSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_MacSets_MacSet{}
        macSets.MacSet = append(macSets.MacSet, child)
        return &macSets.MacSet[len(macSets.MacSet)-1]
    }
    return nil
}

func (macSets *RoutingPolicy_Sets_MacSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macSets.MacSet {
        children[macSets.MacSet[i].GetSegmentPath()] = &macSets.MacSet[i]
    }
    return children
}

func (macSets *RoutingPolicy_Sets_MacSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macSets *RoutingPolicy_Sets_MacSets) GetBundleName() string { return "cisco_ios_xr" }

func (macSets *RoutingPolicy_Sets_MacSets) GetYangName() string { return "mac-sets" }

func (macSets *RoutingPolicy_Sets_MacSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macSets *RoutingPolicy_Sets_MacSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macSets *RoutingPolicy_Sets_MacSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macSets *RoutingPolicy_Sets_MacSets) SetParent(parent types.Entity) { macSets.parent = parent }

func (macSets *RoutingPolicy_Sets_MacSets) GetParent() types.Entity { return macSets.parent }

func (macSets *RoutingPolicy_Sets_MacSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_MacSets_MacSet
// Information about an individual set
type RoutingPolicy_Sets_MacSets_MacSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Mac Set. The type is string. This attribute is mandatory.
    MacSetAsText interface{}
}

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetFilter() yfilter.YFilter { return macSet.YFilter }

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) SetFilter(yf yfilter.YFilter) { macSet.YFilter = yf }

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "mac-set-as-text" { return "MacSetAsText" }
    return ""
}

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetSegmentPath() string {
    return "mac-set" + "[set-name='" + fmt.Sprintf("%v", macSet.SetName) + "']"
}

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = macSet.SetName
    leafs["mac-set-as-text"] = macSet.MacSetAsText
    return leafs
}

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetBundleName() string { return "cisco_ios_xr" }

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetYangName() string { return "mac-set" }

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) SetParent(parent types.Entity) { macSet.parent = parent }

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetParent() types.Entity { return macSet.parent }

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetParentYangName() string { return "mac-sets" }

// RoutingPolicy_Sets_ExtendedCommunityOpaqueSets
// Information about Opaque sets
type RoutingPolicy_Sets_ExtendedCommunityOpaqueSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet.
    ExtendedCommunityOpaqueSet []RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet
}

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetFilter() yfilter.YFilter { return extendedCommunityOpaqueSets.YFilter }

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) SetFilter(yf yfilter.YFilter) { extendedCommunityOpaqueSets.YFilter = yf }

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetGoName(yname string) string {
    if yname == "extended-community-opaque-set" { return "ExtendedCommunityOpaqueSet" }
    return ""
}

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetSegmentPath() string {
    return "extended-community-opaque-sets"
}

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "extended-community-opaque-set" {
        for _, c := range extendedCommunityOpaqueSets.ExtendedCommunityOpaqueSet {
            if extendedCommunityOpaqueSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet{}
        extendedCommunityOpaqueSets.ExtendedCommunityOpaqueSet = append(extendedCommunityOpaqueSets.ExtendedCommunityOpaqueSet, child)
        return &extendedCommunityOpaqueSets.ExtendedCommunityOpaqueSet[len(extendedCommunityOpaqueSets.ExtendedCommunityOpaqueSet)-1]
    }
    return nil
}

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extendedCommunityOpaqueSets.ExtendedCommunityOpaqueSet {
        children[extendedCommunityOpaqueSets.ExtendedCommunityOpaqueSet[i].GetSegmentPath()] = &extendedCommunityOpaqueSets.ExtendedCommunityOpaqueSet[i]
    }
    return children
}

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetYangName() string { return "extended-community-opaque-sets" }

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) SetParent(parent types.Entity) { extendedCommunityOpaqueSets.parent = parent }

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetParent() types.Entity { return extendedCommunityOpaqueSets.parent }

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Extended Community Opaque Set. The type is string. This attribute is
    // mandatory.
    RplExtendedCommunityOpaqueSet interface{}
}

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetFilter() yfilter.YFilter { return extendedCommunityOpaqueSet.YFilter }

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) SetFilter(yf yfilter.YFilter) { extendedCommunityOpaqueSet.YFilter = yf }

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "rpl-extended-community-opaque-set" { return "RplExtendedCommunityOpaqueSet" }
    return ""
}

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetSegmentPath() string {
    return "extended-community-opaque-set" + "[set-name='" + fmt.Sprintf("%v", extendedCommunityOpaqueSet.SetName) + "']"
}

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = extendedCommunityOpaqueSet.SetName
    leafs["rpl-extended-community-opaque-set"] = extendedCommunityOpaqueSet.RplExtendedCommunityOpaqueSet
    return leafs
}

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetYangName() string { return "extended-community-opaque-set" }

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) SetParent(parent types.Entity) { extendedCommunityOpaqueSet.parent = parent }

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetParent() types.Entity { return extendedCommunityOpaqueSet.parent }

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetParentYangName() string { return "extended-community-opaque-sets" }

// RoutingPolicy_Sets_PrependMacSets
// Information about Mac sets
type RoutingPolicy_Sets_PrependMacSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prepend the entries to the existing set. The type is slice of
    // RoutingPolicy_Sets_PrependMacSets_PrependMacSet.
    PrependMacSet []RoutingPolicy_Sets_PrependMacSets_PrependMacSet
}

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) GetFilter() yfilter.YFilter { return prependMacSets.YFilter }

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) SetFilter(yf yfilter.YFilter) { prependMacSets.YFilter = yf }

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) GetGoName(yname string) string {
    if yname == "prepend-mac-set" { return "PrependMacSet" }
    return ""
}

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) GetSegmentPath() string {
    return "prepend-mac-sets"
}

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prepend-mac-set" {
        for _, c := range prependMacSets.PrependMacSet {
            if prependMacSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_PrependMacSets_PrependMacSet{}
        prependMacSets.PrependMacSet = append(prependMacSets.PrependMacSet, child)
        return &prependMacSets.PrependMacSet[len(prependMacSets.PrependMacSet)-1]
    }
    return nil
}

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prependMacSets.PrependMacSet {
        children[prependMacSets.PrependMacSet[i].GetSegmentPath()] = &prependMacSets.PrependMacSet[i]
    }
    return children
}

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) GetBundleName() string { return "cisco_ios_xr" }

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) GetYangName() string { return "prepend-mac-sets" }

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) SetParent(parent types.Entity) { prependMacSets.parent = parent }

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) GetParent() types.Entity { return prependMacSets.parent }

func (prependMacSets *RoutingPolicy_Sets_PrependMacSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_PrependMacSets_PrependMacSet
// Prepend the entries to the existing set
type RoutingPolicy_Sets_PrependMacSets_PrependMacSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Mac Set. The type is string. This attribute is mandatory.
    MacSetAsText interface{}
}

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) GetFilter() yfilter.YFilter { return prependMacSet.YFilter }

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) SetFilter(yf yfilter.YFilter) { prependMacSet.YFilter = yf }

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "mac-set-as-text" { return "MacSetAsText" }
    return ""
}

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) GetSegmentPath() string {
    return "prepend-mac-set" + "[set-name='" + fmt.Sprintf("%v", prependMacSet.SetName) + "']"
}

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = prependMacSet.SetName
    leafs["mac-set-as-text"] = prependMacSet.MacSetAsText
    return leafs
}

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) GetBundleName() string { return "cisco_ios_xr" }

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) GetYangName() string { return "prepend-mac-set" }

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) SetParent(parent types.Entity) { prependMacSet.parent = parent }

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) GetParent() types.Entity { return prependMacSet.parent }

func (prependMacSet *RoutingPolicy_Sets_PrependMacSets_PrependMacSet) GetParentYangName() string { return "prepend-mac-sets" }

// RoutingPolicy_Sets_OspfAreaSets
// Information about OSPF Area sets
type RoutingPolicy_Sets_OspfAreaSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual OSPF area set. Usage: OSPF area set allows
    // to define named set of area numbers        which can be referenced in the
    // route-policy. Area sets      may be used during redistribution of the ospf
    // protocol.  Example: ospf-area-set EXAMPLE      1,                          
    // 192.168.1.255                                  end-set                     
    // Syntax: OSPF area number can be entered as 32 bit number or in          the
    // ip address format. See example.                     Semantic: Area numbers
    // listed in the set will be searched for             a match. In the example
    // these are areas 1 and                  192.168.1.255.                      
    // . The type is slice of RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet.
    OspfAreaSet []RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet
}

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetFilter() yfilter.YFilter { return ospfAreaSets.YFilter }

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) SetFilter(yf yfilter.YFilter) { ospfAreaSets.YFilter = yf }

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetGoName(yname string) string {
    if yname == "ospf-area-set" { return "OspfAreaSet" }
    return ""
}

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetSegmentPath() string {
    return "ospf-area-sets"
}

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospf-area-set" {
        for _, c := range ospfAreaSets.OspfAreaSet {
            if ospfAreaSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet{}
        ospfAreaSets.OspfAreaSet = append(ospfAreaSets.OspfAreaSet, child)
        return &ospfAreaSets.OspfAreaSet[len(ospfAreaSets.OspfAreaSet)-1]
    }
    return nil
}

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfAreaSets.OspfAreaSet {
        children[ospfAreaSets.OspfAreaSet[i].GetSegmentPath()] = &ospfAreaSets.OspfAreaSet[i]
    }
    return children
}

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetBundleName() string { return "cisco_ios_xr" }

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetYangName() string { return "ospf-area-sets" }

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) SetParent(parent types.Entity) { ospfAreaSets.parent = parent }

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetParent() types.Entity { return ospfAreaSets.parent }

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet
// Information about an individual OSPF area set.
// Usage: OSPF area set allows to define named
// set of area numbers        which can be
// referenced in the route-policy. Area sets     
// may be used during redistribution of the ospf
// protocol.  Example: ospf-area-set EXAMPLE     
// 1,                                            
// 192.168.1.255                                 
// end-set                                       
// Syntax: OSPF area number can be entered as 32
// bit number or in          the ip address
// format. See example.                    
// Semantic: Area numbers listed in the set will
// be searched for             a match. In the
// example these are areas 1 and                 
// 192.168.1.255.                                
type RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // OSPF Area Set. The type is string. This attribute is mandatory.
    RplospfAreaSet interface{}
}

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetFilter() yfilter.YFilter { return ospfAreaSet.YFilter }

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) SetFilter(yf yfilter.YFilter) { ospfAreaSet.YFilter = yf }

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "rplospf-area-set" { return "RplospfAreaSet" }
    return ""
}

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetSegmentPath() string {
    return "ospf-area-set" + "[set-name='" + fmt.Sprintf("%v", ospfAreaSet.SetName) + "']"
}

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = ospfAreaSet.SetName
    leafs["rplospf-area-set"] = ospfAreaSet.RplospfAreaSet
    return leafs
}

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetBundleName() string { return "cisco_ios_xr" }

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetYangName() string { return "ospf-area-set" }

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) SetParent(parent types.Entity) { ospfAreaSet.parent = parent }

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetParent() types.Entity { return ospfAreaSet.parent }

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetParentYangName() string { return "ospf-area-sets" }

// RoutingPolicy_Sets_AppendMacSets
// Information about Mac sets
type RoutingPolicy_Sets_AppendMacSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Append the entries to the existing set. The type is slice of
    // RoutingPolicy_Sets_AppendMacSets_AppendMacSet.
    AppendMacSet []RoutingPolicy_Sets_AppendMacSets_AppendMacSet
}

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) GetFilter() yfilter.YFilter { return appendMacSets.YFilter }

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) SetFilter(yf yfilter.YFilter) { appendMacSets.YFilter = yf }

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) GetGoName(yname string) string {
    if yname == "append-mac-set" { return "AppendMacSet" }
    return ""
}

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) GetSegmentPath() string {
    return "append-mac-sets"
}

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "append-mac-set" {
        for _, c := range appendMacSets.AppendMacSet {
            if appendMacSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_AppendMacSets_AppendMacSet{}
        appendMacSets.AppendMacSet = append(appendMacSets.AppendMacSet, child)
        return &appendMacSets.AppendMacSet[len(appendMacSets.AppendMacSet)-1]
    }
    return nil
}

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range appendMacSets.AppendMacSet {
        children[appendMacSets.AppendMacSet[i].GetSegmentPath()] = &appendMacSets.AppendMacSet[i]
    }
    return children
}

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) GetBundleName() string { return "cisco_ios_xr" }

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) GetYangName() string { return "append-mac-sets" }

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) SetParent(parent types.Entity) { appendMacSets.parent = parent }

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) GetParent() types.Entity { return appendMacSets.parent }

func (appendMacSets *RoutingPolicy_Sets_AppendMacSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_AppendMacSets_AppendMacSet
// Append the entries to the existing set
type RoutingPolicy_Sets_AppendMacSets_AppendMacSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Mac Set. The type is string. This attribute is mandatory.
    MacSetAsText interface{}
}

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) GetFilter() yfilter.YFilter { return appendMacSet.YFilter }

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) SetFilter(yf yfilter.YFilter) { appendMacSet.YFilter = yf }

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "mac-set-as-text" { return "MacSetAsText" }
    return ""
}

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) GetSegmentPath() string {
    return "append-mac-set" + "[set-name='" + fmt.Sprintf("%v", appendMacSet.SetName) + "']"
}

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = appendMacSet.SetName
    leafs["mac-set-as-text"] = appendMacSet.MacSetAsText
    return leafs
}

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) GetBundleName() string { return "cisco_ios_xr" }

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) GetYangName() string { return "append-mac-set" }

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) SetParent(parent types.Entity) { appendMacSet.parent = parent }

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) GetParent() types.Entity { return appendMacSet.parent }

func (appendMacSet *RoutingPolicy_Sets_AppendMacSets_AppendMacSet) GetParentYangName() string { return "append-mac-sets" }

// RoutingPolicy_Sets_ExtendedCommunityCostSets
// Information about Cost sets
type RoutingPolicy_Sets_ExtendedCommunityCostSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet.
    ExtendedCommunityCostSet []RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet
}

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetFilter() yfilter.YFilter { return extendedCommunityCostSets.YFilter }

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) SetFilter(yf yfilter.YFilter) { extendedCommunityCostSets.YFilter = yf }

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetGoName(yname string) string {
    if yname == "extended-community-cost-set" { return "ExtendedCommunityCostSet" }
    return ""
}

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetSegmentPath() string {
    return "extended-community-cost-sets"
}

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "extended-community-cost-set" {
        for _, c := range extendedCommunityCostSets.ExtendedCommunityCostSet {
            if extendedCommunityCostSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet{}
        extendedCommunityCostSets.ExtendedCommunityCostSet = append(extendedCommunityCostSets.ExtendedCommunityCostSet, child)
        return &extendedCommunityCostSets.ExtendedCommunityCostSet[len(extendedCommunityCostSets.ExtendedCommunityCostSet)-1]
    }
    return nil
}

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extendedCommunityCostSets.ExtendedCommunityCostSet {
        children[extendedCommunityCostSets.ExtendedCommunityCostSet[i].GetSegmentPath()] = &extendedCommunityCostSets.ExtendedCommunityCostSet[i]
    }
    return children
}

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetYangName() string { return "extended-community-cost-sets" }

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) SetParent(parent types.Entity) { extendedCommunityCostSets.parent = parent }

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetParent() types.Entity { return extendedCommunityCostSets.parent }

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Extended Community Cost Set. The type is string. This attribute is
    // mandatory.
    RplExtendedCommunityCostSet interface{}
}

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetFilter() yfilter.YFilter { return extendedCommunityCostSet.YFilter }

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) SetFilter(yf yfilter.YFilter) { extendedCommunityCostSet.YFilter = yf }

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "rpl-extended-community-cost-set" { return "RplExtendedCommunityCostSet" }
    return ""
}

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetSegmentPath() string {
    return "extended-community-cost-set" + "[set-name='" + fmt.Sprintf("%v", extendedCommunityCostSet.SetName) + "']"
}

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = extendedCommunityCostSet.SetName
    leafs["rpl-extended-community-cost-set"] = extendedCommunityCostSet.RplExtendedCommunityCostSet
    return leafs
}

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetYangName() string { return "extended-community-cost-set" }

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) SetParent(parent types.Entity) { extendedCommunityCostSet.parent = parent }

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetParent() types.Entity { return extendedCommunityCostSet.parent }

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetParentYangName() string { return "extended-community-cost-sets" }

// RoutingPolicy_Sets_RemoveMacSets
// Information about Mac sets
type RoutingPolicy_Sets_RemoveMacSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Remove the entries from the existing set. The type is slice of
    // RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet.
    RemoveMacSet []RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet
}

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) GetFilter() yfilter.YFilter { return removeMacSets.YFilter }

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) SetFilter(yf yfilter.YFilter) { removeMacSets.YFilter = yf }

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) GetGoName(yname string) string {
    if yname == "remove-mac-set" { return "RemoveMacSet" }
    return ""
}

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) GetSegmentPath() string {
    return "remove-mac-sets"
}

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remove-mac-set" {
        for _, c := range removeMacSets.RemoveMacSet {
            if removeMacSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet{}
        removeMacSets.RemoveMacSet = append(removeMacSets.RemoveMacSet, child)
        return &removeMacSets.RemoveMacSet[len(removeMacSets.RemoveMacSet)-1]
    }
    return nil
}

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range removeMacSets.RemoveMacSet {
        children[removeMacSets.RemoveMacSet[i].GetSegmentPath()] = &removeMacSets.RemoveMacSet[i]
    }
    return children
}

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) GetBundleName() string { return "cisco_ios_xr" }

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) GetYangName() string { return "remove-mac-sets" }

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) SetParent(parent types.Entity) { removeMacSets.parent = parent }

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) GetParent() types.Entity { return removeMacSets.parent }

func (removeMacSets *RoutingPolicy_Sets_RemoveMacSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet
// Remove the entries from the existing set
type RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Mac Set. The type is string. This attribute is mandatory.
    MacSetAsText interface{}
}

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) GetFilter() yfilter.YFilter { return removeMacSet.YFilter }

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) SetFilter(yf yfilter.YFilter) { removeMacSet.YFilter = yf }

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "mac-set-as-text" { return "MacSetAsText" }
    return ""
}

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) GetSegmentPath() string {
    return "remove-mac-set" + "[set-name='" + fmt.Sprintf("%v", removeMacSet.SetName) + "']"
}

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = removeMacSet.SetName
    leafs["mac-set-as-text"] = removeMacSet.MacSetAsText
    return leafs
}

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) GetBundleName() string { return "cisco_ios_xr" }

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) GetYangName() string { return "remove-mac-set" }

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) SetParent(parent types.Entity) { removeMacSet.parent = parent }

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) GetParent() types.Entity { return removeMacSet.parent }

func (removeMacSet *RoutingPolicy_Sets_RemoveMacSets_RemoveMacSet) GetParentYangName() string { return "remove-mac-sets" }

// RoutingPolicy_Sets_ExtendedCommunitySooSets
// Information about SOO sets
type RoutingPolicy_Sets_ExtendedCommunitySooSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet.
    ExtendedCommunitySooSet []RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet
}

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetFilter() yfilter.YFilter { return extendedCommunitySooSets.YFilter }

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) SetFilter(yf yfilter.YFilter) { extendedCommunitySooSets.YFilter = yf }

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetGoName(yname string) string {
    if yname == "extended-community-soo-set" { return "ExtendedCommunitySooSet" }
    return ""
}

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetSegmentPath() string {
    return "extended-community-soo-sets"
}

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "extended-community-soo-set" {
        for _, c := range extendedCommunitySooSets.ExtendedCommunitySooSet {
            if extendedCommunitySooSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet{}
        extendedCommunitySooSets.ExtendedCommunitySooSet = append(extendedCommunitySooSets.ExtendedCommunitySooSet, child)
        return &extendedCommunitySooSets.ExtendedCommunitySooSet[len(extendedCommunitySooSets.ExtendedCommunitySooSet)-1]
    }
    return nil
}

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extendedCommunitySooSets.ExtendedCommunitySooSet {
        children[extendedCommunitySooSets.ExtendedCommunitySooSet[i].GetSegmentPath()] = &extendedCommunitySooSets.ExtendedCommunitySooSet[i]
    }
    return children
}

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetYangName() string { return "extended-community-soo-sets" }

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) SetParent(parent types.Entity) { extendedCommunitySooSets.parent = parent }

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetParent() types.Entity { return extendedCommunitySooSets.parent }

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Extended Community SOO Set. The type is string. This attribute is
    // mandatory.
    RplExtendedCommunitySooSet interface{}
}

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetFilter() yfilter.YFilter { return extendedCommunitySooSet.YFilter }

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) SetFilter(yf yfilter.YFilter) { extendedCommunitySooSet.YFilter = yf }

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "rpl-extended-community-soo-set" { return "RplExtendedCommunitySooSet" }
    return ""
}

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetSegmentPath() string {
    return "extended-community-soo-set" + "[set-name='" + fmt.Sprintf("%v", extendedCommunitySooSet.SetName) + "']"
}

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = extendedCommunitySooSet.SetName
    leafs["rpl-extended-community-soo-set"] = extendedCommunitySooSet.RplExtendedCommunitySooSet
    return leafs
}

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetYangName() string { return "extended-community-soo-set" }

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) SetParent(parent types.Entity) { extendedCommunitySooSet.parent = parent }

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetParent() types.Entity { return extendedCommunitySooSet.parent }

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetParentYangName() string { return "extended-community-soo-sets" }

// RoutingPolicy_Sets_EsiSets
// Information about Esi sets
type RoutingPolicy_Sets_EsiSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_EsiSets_EsiSet.
    EsiSet []RoutingPolicy_Sets_EsiSets_EsiSet
}

func (esiSets *RoutingPolicy_Sets_EsiSets) GetFilter() yfilter.YFilter { return esiSets.YFilter }

func (esiSets *RoutingPolicy_Sets_EsiSets) SetFilter(yf yfilter.YFilter) { esiSets.YFilter = yf }

func (esiSets *RoutingPolicy_Sets_EsiSets) GetGoName(yname string) string {
    if yname == "esi-set" { return "EsiSet" }
    return ""
}

func (esiSets *RoutingPolicy_Sets_EsiSets) GetSegmentPath() string {
    return "esi-sets"
}

func (esiSets *RoutingPolicy_Sets_EsiSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "esi-set" {
        for _, c := range esiSets.EsiSet {
            if esiSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_EsiSets_EsiSet{}
        esiSets.EsiSet = append(esiSets.EsiSet, child)
        return &esiSets.EsiSet[len(esiSets.EsiSet)-1]
    }
    return nil
}

func (esiSets *RoutingPolicy_Sets_EsiSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range esiSets.EsiSet {
        children[esiSets.EsiSet[i].GetSegmentPath()] = &esiSets.EsiSet[i]
    }
    return children
}

func (esiSets *RoutingPolicy_Sets_EsiSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (esiSets *RoutingPolicy_Sets_EsiSets) GetBundleName() string { return "cisco_ios_xr" }

func (esiSets *RoutingPolicy_Sets_EsiSets) GetYangName() string { return "esi-sets" }

func (esiSets *RoutingPolicy_Sets_EsiSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esiSets *RoutingPolicy_Sets_EsiSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esiSets *RoutingPolicy_Sets_EsiSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esiSets *RoutingPolicy_Sets_EsiSets) SetParent(parent types.Entity) { esiSets.parent = parent }

func (esiSets *RoutingPolicy_Sets_EsiSets) GetParent() types.Entity { return esiSets.parent }

func (esiSets *RoutingPolicy_Sets_EsiSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_EsiSets_EsiSet
// Information about an individual set
type RoutingPolicy_Sets_EsiSets_EsiSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Esi Set. The type is string. This attribute is mandatory.
    EsiSetAsText interface{}
}

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetFilter() yfilter.YFilter { return esiSet.YFilter }

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) SetFilter(yf yfilter.YFilter) { esiSet.YFilter = yf }

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "esi-set-as-text" { return "EsiSetAsText" }
    return ""
}

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetSegmentPath() string {
    return "esi-set" + "[set-name='" + fmt.Sprintf("%v", esiSet.SetName) + "']"
}

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = esiSet.SetName
    leafs["esi-set-as-text"] = esiSet.EsiSetAsText
    return leafs
}

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetBundleName() string { return "cisco_ios_xr" }

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetYangName() string { return "esi-set" }

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) SetParent(parent types.Entity) { esiSet.parent = parent }

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetParent() types.Entity { return esiSet.parent }

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetParentYangName() string { return "esi-sets" }

// RoutingPolicy_Sets_PrependEsiSets
// Information about Esi sets
type RoutingPolicy_Sets_PrependEsiSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prepend the entries to the existing set. The type is slice of
    // RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet.
    PrependEsiSet []RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet
}

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) GetFilter() yfilter.YFilter { return prependEsiSets.YFilter }

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) SetFilter(yf yfilter.YFilter) { prependEsiSets.YFilter = yf }

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) GetGoName(yname string) string {
    if yname == "prepend-esi-set" { return "PrependEsiSet" }
    return ""
}

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) GetSegmentPath() string {
    return "prepend-esi-sets"
}

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prepend-esi-set" {
        for _, c := range prependEsiSets.PrependEsiSet {
            if prependEsiSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet{}
        prependEsiSets.PrependEsiSet = append(prependEsiSets.PrependEsiSet, child)
        return &prependEsiSets.PrependEsiSet[len(prependEsiSets.PrependEsiSet)-1]
    }
    return nil
}

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prependEsiSets.PrependEsiSet {
        children[prependEsiSets.PrependEsiSet[i].GetSegmentPath()] = &prependEsiSets.PrependEsiSet[i]
    }
    return children
}

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) GetBundleName() string { return "cisco_ios_xr" }

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) GetYangName() string { return "prepend-esi-sets" }

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) SetParent(parent types.Entity) { prependEsiSets.parent = parent }

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) GetParent() types.Entity { return prependEsiSets.parent }

func (prependEsiSets *RoutingPolicy_Sets_PrependEsiSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet
// Prepend the entries to the existing set
type RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Esi Set. The type is string. This attribute is mandatory.
    EsiSetAsText interface{}
}

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) GetFilter() yfilter.YFilter { return prependEsiSet.YFilter }

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) SetFilter(yf yfilter.YFilter) { prependEsiSet.YFilter = yf }

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "esi-set-as-text" { return "EsiSetAsText" }
    return ""
}

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) GetSegmentPath() string {
    return "prepend-esi-set" + "[set-name='" + fmt.Sprintf("%v", prependEsiSet.SetName) + "']"
}

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = prependEsiSet.SetName
    leafs["esi-set-as-text"] = prependEsiSet.EsiSetAsText
    return leafs
}

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) GetBundleName() string { return "cisco_ios_xr" }

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) GetYangName() string { return "prepend-esi-set" }

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) SetParent(parent types.Entity) { prependEsiSet.parent = parent }

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) GetParent() types.Entity { return prependEsiSet.parent }

func (prependEsiSet *RoutingPolicy_Sets_PrependEsiSets_PrependEsiSet) GetParentYangName() string { return "prepend-esi-sets" }

// RoutingPolicy_Sets_AppendEsiSets
// Information about Esi sets
type RoutingPolicy_Sets_AppendEsiSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Append the entries to the existing set. The type is slice of
    // RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet.
    AppendEsiSet []RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet
}

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) GetFilter() yfilter.YFilter { return appendEsiSets.YFilter }

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) SetFilter(yf yfilter.YFilter) { appendEsiSets.YFilter = yf }

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) GetGoName(yname string) string {
    if yname == "append-esi-set" { return "AppendEsiSet" }
    return ""
}

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) GetSegmentPath() string {
    return "append-esi-sets"
}

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "append-esi-set" {
        for _, c := range appendEsiSets.AppendEsiSet {
            if appendEsiSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet{}
        appendEsiSets.AppendEsiSet = append(appendEsiSets.AppendEsiSet, child)
        return &appendEsiSets.AppendEsiSet[len(appendEsiSets.AppendEsiSet)-1]
    }
    return nil
}

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range appendEsiSets.AppendEsiSet {
        children[appendEsiSets.AppendEsiSet[i].GetSegmentPath()] = &appendEsiSets.AppendEsiSet[i]
    }
    return children
}

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) GetBundleName() string { return "cisco_ios_xr" }

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) GetYangName() string { return "append-esi-sets" }

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) SetParent(parent types.Entity) { appendEsiSets.parent = parent }

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) GetParent() types.Entity { return appendEsiSets.parent }

func (appendEsiSets *RoutingPolicy_Sets_AppendEsiSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet
// Append the entries to the existing set
type RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Esi Set. The type is string. This attribute is mandatory.
    EsiSetAsText interface{}
}

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) GetFilter() yfilter.YFilter { return appendEsiSet.YFilter }

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) SetFilter(yf yfilter.YFilter) { appendEsiSet.YFilter = yf }

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "esi-set-as-text" { return "EsiSetAsText" }
    return ""
}

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) GetSegmentPath() string {
    return "append-esi-set" + "[set-name='" + fmt.Sprintf("%v", appendEsiSet.SetName) + "']"
}

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = appendEsiSet.SetName
    leafs["esi-set-as-text"] = appendEsiSet.EsiSetAsText
    return leafs
}

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) GetBundleName() string { return "cisco_ios_xr" }

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) GetYangName() string { return "append-esi-set" }

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) SetParent(parent types.Entity) { appendEsiSet.parent = parent }

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) GetParent() types.Entity { return appendEsiSet.parent }

func (appendEsiSet *RoutingPolicy_Sets_AppendEsiSets_AppendEsiSet) GetParentYangName() string { return "append-esi-sets" }

// RoutingPolicy_Sets_RemoveEsiSets
// Information about Esi sets
type RoutingPolicy_Sets_RemoveEsiSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Remove the entries from the existing set. The type is slice of
    // RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet.
    RemoveEsiSet []RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet
}

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) GetFilter() yfilter.YFilter { return removeEsiSets.YFilter }

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) SetFilter(yf yfilter.YFilter) { removeEsiSets.YFilter = yf }

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) GetGoName(yname string) string {
    if yname == "remove-esi-set" { return "RemoveEsiSet" }
    return ""
}

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) GetSegmentPath() string {
    return "remove-esi-sets"
}

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remove-esi-set" {
        for _, c := range removeEsiSets.RemoveEsiSet {
            if removeEsiSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet{}
        removeEsiSets.RemoveEsiSet = append(removeEsiSets.RemoveEsiSet, child)
        return &removeEsiSets.RemoveEsiSet[len(removeEsiSets.RemoveEsiSet)-1]
    }
    return nil
}

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range removeEsiSets.RemoveEsiSet {
        children[removeEsiSets.RemoveEsiSet[i].GetSegmentPath()] = &removeEsiSets.RemoveEsiSet[i]
    }
    return children
}

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) GetBundleName() string { return "cisco_ios_xr" }

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) GetYangName() string { return "remove-esi-sets" }

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) SetParent(parent types.Entity) { removeEsiSets.parent = parent }

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) GetParent() types.Entity { return removeEsiSets.parent }

func (removeEsiSets *RoutingPolicy_Sets_RemoveEsiSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet
// Remove the entries from the existing set
type RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Esi Set. The type is string. This attribute is mandatory.
    EsiSetAsText interface{}
}

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) GetFilter() yfilter.YFilter { return removeEsiSet.YFilter }

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) SetFilter(yf yfilter.YFilter) { removeEsiSet.YFilter = yf }

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "esi-set-as-text" { return "EsiSetAsText" }
    return ""
}

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) GetSegmentPath() string {
    return "remove-esi-set" + "[set-name='" + fmt.Sprintf("%v", removeEsiSet.SetName) + "']"
}

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = removeEsiSet.SetName
    leafs["esi-set-as-text"] = removeEsiSet.EsiSetAsText
    return leafs
}

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) GetBundleName() string { return "cisco_ios_xr" }

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) GetYangName() string { return "remove-esi-set" }

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) SetParent(parent types.Entity) { removeEsiSet.parent = parent }

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) GetParent() types.Entity { return removeEsiSet.parent }

func (removeEsiSet *RoutingPolicy_Sets_RemoveEsiSets_RemoveEsiSet) GetParentYangName() string { return "remove-esi-sets" }

// RoutingPolicy_Sets_ExtendedCommunitySegNhSets
// Information about SegNH sets
type RoutingPolicy_Sets_ExtendedCommunitySegNhSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet.
    ExtendedCommunitySegNhSet []RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet
}

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetFilter() yfilter.YFilter { return extendedCommunitySegNhSets.YFilter }

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) SetFilter(yf yfilter.YFilter) { extendedCommunitySegNhSets.YFilter = yf }

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetGoName(yname string) string {
    if yname == "extended-community-seg-nh-set" { return "ExtendedCommunitySegNhSet" }
    return ""
}

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetSegmentPath() string {
    return "extended-community-seg-nh-sets"
}

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "extended-community-seg-nh-set" {
        for _, c := range extendedCommunitySegNhSets.ExtendedCommunitySegNhSet {
            if extendedCommunitySegNhSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet{}
        extendedCommunitySegNhSets.ExtendedCommunitySegNhSet = append(extendedCommunitySegNhSets.ExtendedCommunitySegNhSet, child)
        return &extendedCommunitySegNhSets.ExtendedCommunitySegNhSet[len(extendedCommunitySegNhSets.ExtendedCommunitySegNhSet)-1]
    }
    return nil
}

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extendedCommunitySegNhSets.ExtendedCommunitySegNhSet {
        children[extendedCommunitySegNhSets.ExtendedCommunitySegNhSet[i].GetSegmentPath()] = &extendedCommunitySegNhSets.ExtendedCommunitySegNhSet[i]
    }
    return children
}

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetYangName() string { return "extended-community-seg-nh-sets" }

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) SetParent(parent types.Entity) { extendedCommunitySegNhSets.parent = parent }

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetParent() types.Entity { return extendedCommunitySegNhSets.parent }

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Extended Community SegNH Set. The type is string. This attribute is
    // mandatory.
    RplExtendedCommunitySegNhSet interface{}
}

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetFilter() yfilter.YFilter { return extendedCommunitySegNhSet.YFilter }

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) SetFilter(yf yfilter.YFilter) { extendedCommunitySegNhSet.YFilter = yf }

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "rpl-extended-community-seg-nh-set" { return "RplExtendedCommunitySegNhSet" }
    return ""
}

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetSegmentPath() string {
    return "extended-community-seg-nh-set" + "[set-name='" + fmt.Sprintf("%v", extendedCommunitySegNhSet.SetName) + "']"
}

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = extendedCommunitySegNhSet.SetName
    leafs["rpl-extended-community-seg-nh-set"] = extendedCommunitySegNhSet.RplExtendedCommunitySegNhSet
    return leafs
}

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetYangName() string { return "extended-community-seg-nh-set" }

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) SetParent(parent types.Entity) { extendedCommunitySegNhSet.parent = parent }

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetParent() types.Entity { return extendedCommunitySegNhSet.parent }

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetParentYangName() string { return "extended-community-seg-nh-sets" }

// RoutingPolicy_Sets_RdSets
// Information about RD sets
type RoutingPolicy_Sets_RdSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_RdSets_RdSet.
    RdSet []RoutingPolicy_Sets_RdSets_RdSet
}

func (rdSets *RoutingPolicy_Sets_RdSets) GetFilter() yfilter.YFilter { return rdSets.YFilter }

func (rdSets *RoutingPolicy_Sets_RdSets) SetFilter(yf yfilter.YFilter) { rdSets.YFilter = yf }

func (rdSets *RoutingPolicy_Sets_RdSets) GetGoName(yname string) string {
    if yname == "rd-set" { return "RdSet" }
    return ""
}

func (rdSets *RoutingPolicy_Sets_RdSets) GetSegmentPath() string {
    return "rd-sets"
}

func (rdSets *RoutingPolicy_Sets_RdSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rd-set" {
        for _, c := range rdSets.RdSet {
            if rdSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_RdSets_RdSet{}
        rdSets.RdSet = append(rdSets.RdSet, child)
        return &rdSets.RdSet[len(rdSets.RdSet)-1]
    }
    return nil
}

func (rdSets *RoutingPolicy_Sets_RdSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rdSets.RdSet {
        children[rdSets.RdSet[i].GetSegmentPath()] = &rdSets.RdSet[i]
    }
    return children
}

func (rdSets *RoutingPolicy_Sets_RdSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rdSets *RoutingPolicy_Sets_RdSets) GetBundleName() string { return "cisco_ios_xr" }

func (rdSets *RoutingPolicy_Sets_RdSets) GetYangName() string { return "rd-sets" }

func (rdSets *RoutingPolicy_Sets_RdSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rdSets *RoutingPolicy_Sets_RdSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rdSets *RoutingPolicy_Sets_RdSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rdSets *RoutingPolicy_Sets_RdSets) SetParent(parent types.Entity) { rdSets.parent = parent }

func (rdSets *RoutingPolicy_Sets_RdSets) GetParent() types.Entity { return rdSets.parent }

func (rdSets *RoutingPolicy_Sets_RdSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_RdSets_RdSet
// Information about an individual set
type RoutingPolicy_Sets_RdSets_RdSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // RD Set. The type is string. This attribute is mandatory.
    RplrdSet interface{}
}

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetFilter() yfilter.YFilter { return rdSet.YFilter }

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) SetFilter(yf yfilter.YFilter) { rdSet.YFilter = yf }

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "rplrd-set" { return "RplrdSet" }
    return ""
}

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetSegmentPath() string {
    return "rd-set" + "[set-name='" + fmt.Sprintf("%v", rdSet.SetName) + "']"
}

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = rdSet.SetName
    leafs["rplrd-set"] = rdSet.RplrdSet
    return leafs
}

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetBundleName() string { return "cisco_ios_xr" }

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetYangName() string { return "rd-set" }

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) SetParent(parent types.Entity) { rdSet.parent = parent }

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetParent() types.Entity { return rdSet.parent }

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetParentYangName() string { return "rd-sets" }

// RoutingPolicy_Sets_PolicyGlobalSetTable
// Information about PolicyGlobal sets
type RoutingPolicy_Sets_PolicyGlobalSetTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is string.
    PolicyGlobalSet interface{}
}

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetFilter() yfilter.YFilter { return policyGlobalSetTable.YFilter }

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) SetFilter(yf yfilter.YFilter) { policyGlobalSetTable.YFilter = yf }

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetGoName(yname string) string {
    if yname == "policy-global-set" { return "PolicyGlobalSet" }
    return ""
}

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetSegmentPath() string {
    return "policy-global-set-table"
}

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-global-set"] = policyGlobalSetTable.PolicyGlobalSet
    return leafs
}

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetBundleName() string { return "cisco_ios_xr" }

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetYangName() string { return "policy-global-set-table" }

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) SetParent(parent types.Entity) { policyGlobalSetTable.parent = parent }

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetParent() types.Entity { return policyGlobalSetTable.parent }

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_AppendLargeCommunitySets
// Information about Large Community sets
type RoutingPolicy_Sets_AppendLargeCommunitySets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Append the entries to the existing set. The type is slice of
    // RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet.
    AppendLargeCommunitySet []RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet
}

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) GetFilter() yfilter.YFilter { return appendLargeCommunitySets.YFilter }

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) SetFilter(yf yfilter.YFilter) { appendLargeCommunitySets.YFilter = yf }

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) GetGoName(yname string) string {
    if yname == "append-large-community-set" { return "AppendLargeCommunitySet" }
    return ""
}

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) GetSegmentPath() string {
    return "append-large-community-sets"
}

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "append-large-community-set" {
        for _, c := range appendLargeCommunitySets.AppendLargeCommunitySet {
            if appendLargeCommunitySets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet{}
        appendLargeCommunitySets.AppendLargeCommunitySet = append(appendLargeCommunitySets.AppendLargeCommunitySet, child)
        return &appendLargeCommunitySets.AppendLargeCommunitySet[len(appendLargeCommunitySets.AppendLargeCommunitySet)-1]
    }
    return nil
}

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range appendLargeCommunitySets.AppendLargeCommunitySet {
        children[appendLargeCommunitySets.AppendLargeCommunitySet[i].GetSegmentPath()] = &appendLargeCommunitySets.AppendLargeCommunitySet[i]
    }
    return children
}

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) GetBundleName() string { return "cisco_ios_xr" }

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) GetYangName() string { return "append-large-community-sets" }

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) SetParent(parent types.Entity) { appendLargeCommunitySets.parent = parent }

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) GetParent() types.Entity { return appendLargeCommunitySets.parent }

func (appendLargeCommunitySets *RoutingPolicy_Sets_AppendLargeCommunitySets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet
// Append the entries to the existing set
type RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Large Community Set. The type is string. This attribute is mandatory.
    LargeCommunitySetAsText interface{}
}

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) GetFilter() yfilter.YFilter { return appendLargeCommunitySet.YFilter }

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) SetFilter(yf yfilter.YFilter) { appendLargeCommunitySet.YFilter = yf }

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "large-community-set-as-text" { return "LargeCommunitySetAsText" }
    return ""
}

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) GetSegmentPath() string {
    return "append-large-community-set" + "[set-name='" + fmt.Sprintf("%v", appendLargeCommunitySet.SetName) + "']"
}

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = appendLargeCommunitySet.SetName
    leafs["large-community-set-as-text"] = appendLargeCommunitySet.LargeCommunitySetAsText
    return leafs
}

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) GetBundleName() string { return "cisco_ios_xr" }

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) GetYangName() string { return "append-large-community-set" }

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) SetParent(parent types.Entity) { appendLargeCommunitySet.parent = parent }

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) GetParent() types.Entity { return appendLargeCommunitySet.parent }

func (appendLargeCommunitySet *RoutingPolicy_Sets_AppendLargeCommunitySets_AppendLargeCommunitySet) GetParentYangName() string { return "append-large-community-sets" }

// RoutingPolicy_Sets_ExtendedCommunityBandwidthSets
// Information about Bandwidth sets
type RoutingPolicy_Sets_ExtendedCommunityBandwidthSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet.
    ExtendedCommunityBandwidthSet []RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet
}

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetFilter() yfilter.YFilter { return extendedCommunityBandwidthSets.YFilter }

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) SetFilter(yf yfilter.YFilter) { extendedCommunityBandwidthSets.YFilter = yf }

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetGoName(yname string) string {
    if yname == "extended-community-bandwidth-set" { return "ExtendedCommunityBandwidthSet" }
    return ""
}

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetSegmentPath() string {
    return "extended-community-bandwidth-sets"
}

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "extended-community-bandwidth-set" {
        for _, c := range extendedCommunityBandwidthSets.ExtendedCommunityBandwidthSet {
            if extendedCommunityBandwidthSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet{}
        extendedCommunityBandwidthSets.ExtendedCommunityBandwidthSet = append(extendedCommunityBandwidthSets.ExtendedCommunityBandwidthSet, child)
        return &extendedCommunityBandwidthSets.ExtendedCommunityBandwidthSet[len(extendedCommunityBandwidthSets.ExtendedCommunityBandwidthSet)-1]
    }
    return nil
}

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extendedCommunityBandwidthSets.ExtendedCommunityBandwidthSet {
        children[extendedCommunityBandwidthSets.ExtendedCommunityBandwidthSet[i].GetSegmentPath()] = &extendedCommunityBandwidthSets.ExtendedCommunityBandwidthSet[i]
    }
    return children
}

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetYangName() string { return "extended-community-bandwidth-sets" }

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) SetParent(parent types.Entity) { extendedCommunityBandwidthSets.parent = parent }

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetParent() types.Entity { return extendedCommunityBandwidthSets.parent }

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Extended Community Bandwidth Set. The type is string. This attribute is
    // mandatory.
    RplExtendedCommunityBandwidthSet interface{}
}

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetFilter() yfilter.YFilter { return extendedCommunityBandwidthSet.YFilter }

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) SetFilter(yf yfilter.YFilter) { extendedCommunityBandwidthSet.YFilter = yf }

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "rpl-extended-community-bandwidth-set" { return "RplExtendedCommunityBandwidthSet" }
    return ""
}

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetSegmentPath() string {
    return "extended-community-bandwidth-set" + "[set-name='" + fmt.Sprintf("%v", extendedCommunityBandwidthSet.SetName) + "']"
}

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = extendedCommunityBandwidthSet.SetName
    leafs["rpl-extended-community-bandwidth-set"] = extendedCommunityBandwidthSet.RplExtendedCommunityBandwidthSet
    return leafs
}

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetYangName() string { return "extended-community-bandwidth-set" }

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) SetParent(parent types.Entity) { extendedCommunityBandwidthSet.parent = parent }

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetParent() types.Entity { return extendedCommunityBandwidthSet.parent }

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetParentYangName() string { return "extended-community-bandwidth-sets" }

// RoutingPolicy_Sets_CommunitySets
// Information about Community sets
type RoutingPolicy_Sets_CommunitySets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_CommunitySets_CommunitySet.
    CommunitySet []RoutingPolicy_Sets_CommunitySets_CommunitySet
}

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetFilter() yfilter.YFilter { return communitySets.YFilter }

func (communitySets *RoutingPolicy_Sets_CommunitySets) SetFilter(yf yfilter.YFilter) { communitySets.YFilter = yf }

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetGoName(yname string) string {
    if yname == "community-set" { return "CommunitySet" }
    return ""
}

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetSegmentPath() string {
    return "community-sets"
}

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "community-set" {
        for _, c := range communitySets.CommunitySet {
            if communitySets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_CommunitySets_CommunitySet{}
        communitySets.CommunitySet = append(communitySets.CommunitySet, child)
        return &communitySets.CommunitySet[len(communitySets.CommunitySet)-1]
    }
    return nil
}

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range communitySets.CommunitySet {
        children[communitySets.CommunitySet[i].GetSegmentPath()] = &communitySets.CommunitySet[i]
    }
    return children
}

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetBundleName() string { return "cisco_ios_xr" }

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetYangName() string { return "community-sets" }

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (communitySets *RoutingPolicy_Sets_CommunitySets) SetParent(parent types.Entity) { communitySets.parent = parent }

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetParent() types.Entity { return communitySets.parent }

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_CommunitySets_CommunitySet
// Information about an individual set
type RoutingPolicy_Sets_CommunitySets_CommunitySet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Community Set. The type is string. This attribute is mandatory.
    RplCommunitySet interface{}
}

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetFilter() yfilter.YFilter { return communitySet.YFilter }

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) SetFilter(yf yfilter.YFilter) { communitySet.YFilter = yf }

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "rpl-community-set" { return "RplCommunitySet" }
    return ""
}

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetSegmentPath() string {
    return "community-set" + "[set-name='" + fmt.Sprintf("%v", communitySet.SetName) + "']"
}

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = communitySet.SetName
    leafs["rpl-community-set"] = communitySet.RplCommunitySet
    return leafs
}

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetBundleName() string { return "cisco_ios_xr" }

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetYangName() string { return "community-set" }

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) SetParent(parent types.Entity) { communitySet.parent = parent }

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetParent() types.Entity { return communitySet.parent }

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetParentYangName() string { return "community-sets" }

// RoutingPolicy_Sets_AsPathSets
// Information about AS Path sets
type RoutingPolicy_Sets_AsPathSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_AsPathSets_AsPathSet.
    AsPathSet []RoutingPolicy_Sets_AsPathSets_AsPathSet
}

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetFilter() yfilter.YFilter { return asPathSets.YFilter }

func (asPathSets *RoutingPolicy_Sets_AsPathSets) SetFilter(yf yfilter.YFilter) { asPathSets.YFilter = yf }

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetGoName(yname string) string {
    if yname == "as-path-set" { return "AsPathSet" }
    return ""
}

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetSegmentPath() string {
    return "as-path-sets"
}

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "as-path-set" {
        for _, c := range asPathSets.AsPathSet {
            if asPathSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_AsPathSets_AsPathSet{}
        asPathSets.AsPathSet = append(asPathSets.AsPathSet, child)
        return &asPathSets.AsPathSet[len(asPathSets.AsPathSet)-1]
    }
    return nil
}

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asPathSets.AsPathSet {
        children[asPathSets.AsPathSet[i].GetSegmentPath()] = &asPathSets.AsPathSet[i]
    }
    return children
}

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetBundleName() string { return "cisco_ios_xr" }

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetYangName() string { return "as-path-sets" }

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asPathSets *RoutingPolicy_Sets_AsPathSets) SetParent(parent types.Entity) { asPathSets.parent = parent }

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetParent() types.Entity { return asPathSets.parent }

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_AsPathSets_AsPathSet
// Information about an individual set
type RoutingPolicy_Sets_AsPathSets_AsPathSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // ASPath Set. The type is string. This attribute is mandatory.
    RplasPathSet interface{}
}

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetFilter() yfilter.YFilter { return asPathSet.YFilter }

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) SetFilter(yf yfilter.YFilter) { asPathSet.YFilter = yf }

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "rplas-path-set" { return "RplasPathSet" }
    return ""
}

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetSegmentPath() string {
    return "as-path-set" + "[set-name='" + fmt.Sprintf("%v", asPathSet.SetName) + "']"
}

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = asPathSet.SetName
    leafs["rplas-path-set"] = asPathSet.RplasPathSet
    return leafs
}

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetBundleName() string { return "cisco_ios_xr" }

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetYangName() string { return "as-path-set" }

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) SetParent(parent types.Entity) { asPathSet.parent = parent }

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetParent() types.Entity { return asPathSet.parent }

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetParentYangName() string { return "as-path-sets" }

// RoutingPolicy_Sets_TagSets
// Information about Tag sets
type RoutingPolicy_Sets_TagSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_TagSets_TagSet.
    TagSet []RoutingPolicy_Sets_TagSets_TagSet
}

func (tagSets *RoutingPolicy_Sets_TagSets) GetFilter() yfilter.YFilter { return tagSets.YFilter }

func (tagSets *RoutingPolicy_Sets_TagSets) SetFilter(yf yfilter.YFilter) { tagSets.YFilter = yf }

func (tagSets *RoutingPolicy_Sets_TagSets) GetGoName(yname string) string {
    if yname == "tag-set" { return "TagSet" }
    return ""
}

func (tagSets *RoutingPolicy_Sets_TagSets) GetSegmentPath() string {
    return "tag-sets"
}

func (tagSets *RoutingPolicy_Sets_TagSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tag-set" {
        for _, c := range tagSets.TagSet {
            if tagSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_TagSets_TagSet{}
        tagSets.TagSet = append(tagSets.TagSet, child)
        return &tagSets.TagSet[len(tagSets.TagSet)-1]
    }
    return nil
}

func (tagSets *RoutingPolicy_Sets_TagSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tagSets.TagSet {
        children[tagSets.TagSet[i].GetSegmentPath()] = &tagSets.TagSet[i]
    }
    return children
}

func (tagSets *RoutingPolicy_Sets_TagSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tagSets *RoutingPolicy_Sets_TagSets) GetBundleName() string { return "cisco_ios_xr" }

func (tagSets *RoutingPolicy_Sets_TagSets) GetYangName() string { return "tag-sets" }

func (tagSets *RoutingPolicy_Sets_TagSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tagSets *RoutingPolicy_Sets_TagSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tagSets *RoutingPolicy_Sets_TagSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tagSets *RoutingPolicy_Sets_TagSets) SetParent(parent types.Entity) { tagSets.parent = parent }

func (tagSets *RoutingPolicy_Sets_TagSets) GetParent() types.Entity { return tagSets.parent }

func (tagSets *RoutingPolicy_Sets_TagSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_TagSets_TagSet
// Information about an individual set
type RoutingPolicy_Sets_TagSets_TagSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Tag Set. The type is string. This attribute is mandatory.
    RplTagSet interface{}
}

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetFilter() yfilter.YFilter { return tagSet.YFilter }

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) SetFilter(yf yfilter.YFilter) { tagSet.YFilter = yf }

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "rpl-tag-set" { return "RplTagSet" }
    return ""
}

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetSegmentPath() string {
    return "tag-set" + "[set-name='" + fmt.Sprintf("%v", tagSet.SetName) + "']"
}

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = tagSet.SetName
    leafs["rpl-tag-set"] = tagSet.RplTagSet
    return leafs
}

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetBundleName() string { return "cisco_ios_xr" }

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetYangName() string { return "tag-set" }

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) SetParent(parent types.Entity) { tagSet.parent = parent }

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetParent() types.Entity { return tagSet.parent }

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetParentYangName() string { return "tag-sets" }

// RoutingPolicy_Sets_EtagSets
// Information about Etag sets
type RoutingPolicy_Sets_EtagSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_EtagSets_EtagSet.
    EtagSet []RoutingPolicy_Sets_EtagSets_EtagSet
}

func (etagSets *RoutingPolicy_Sets_EtagSets) GetFilter() yfilter.YFilter { return etagSets.YFilter }

func (etagSets *RoutingPolicy_Sets_EtagSets) SetFilter(yf yfilter.YFilter) { etagSets.YFilter = yf }

func (etagSets *RoutingPolicy_Sets_EtagSets) GetGoName(yname string) string {
    if yname == "etag-set" { return "EtagSet" }
    return ""
}

func (etagSets *RoutingPolicy_Sets_EtagSets) GetSegmentPath() string {
    return "etag-sets"
}

func (etagSets *RoutingPolicy_Sets_EtagSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "etag-set" {
        for _, c := range etagSets.EtagSet {
            if etagSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_EtagSets_EtagSet{}
        etagSets.EtagSet = append(etagSets.EtagSet, child)
        return &etagSets.EtagSet[len(etagSets.EtagSet)-1]
    }
    return nil
}

func (etagSets *RoutingPolicy_Sets_EtagSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range etagSets.EtagSet {
        children[etagSets.EtagSet[i].GetSegmentPath()] = &etagSets.EtagSet[i]
    }
    return children
}

func (etagSets *RoutingPolicy_Sets_EtagSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (etagSets *RoutingPolicy_Sets_EtagSets) GetBundleName() string { return "cisco_ios_xr" }

func (etagSets *RoutingPolicy_Sets_EtagSets) GetYangName() string { return "etag-sets" }

func (etagSets *RoutingPolicy_Sets_EtagSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (etagSets *RoutingPolicy_Sets_EtagSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (etagSets *RoutingPolicy_Sets_EtagSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (etagSets *RoutingPolicy_Sets_EtagSets) SetParent(parent types.Entity) { etagSets.parent = parent }

func (etagSets *RoutingPolicy_Sets_EtagSets) GetParent() types.Entity { return etagSets.parent }

func (etagSets *RoutingPolicy_Sets_EtagSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_EtagSets_EtagSet
// Information about an individual set
type RoutingPolicy_Sets_EtagSets_EtagSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Etag Set. The type is string. This attribute is mandatory.
    EtagSetAsText interface{}
}

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetFilter() yfilter.YFilter { return etagSet.YFilter }

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) SetFilter(yf yfilter.YFilter) { etagSet.YFilter = yf }

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "etag-set-as-text" { return "EtagSetAsText" }
    return ""
}

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetSegmentPath() string {
    return "etag-set" + "[set-name='" + fmt.Sprintf("%v", etagSet.SetName) + "']"
}

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = etagSet.SetName
    leafs["etag-set-as-text"] = etagSet.EtagSetAsText
    return leafs
}

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetBundleName() string { return "cisco_ios_xr" }

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetYangName() string { return "etag-set" }

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) SetParent(parent types.Entity) { etagSet.parent = parent }

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetParent() types.Entity { return etagSet.parent }

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetParentYangName() string { return "etag-sets" }

// RoutingPolicy_Sets_ExtendedCommunityRtSets
// Information about RT sets
type RoutingPolicy_Sets_ExtendedCommunityRtSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet.
    ExtendedCommunityRtSet []RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet
}

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetFilter() yfilter.YFilter { return extendedCommunityRtSets.YFilter }

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) SetFilter(yf yfilter.YFilter) { extendedCommunityRtSets.YFilter = yf }

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetGoName(yname string) string {
    if yname == "extended-community-rt-set" { return "ExtendedCommunityRtSet" }
    return ""
}

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetSegmentPath() string {
    return "extended-community-rt-sets"
}

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "extended-community-rt-set" {
        for _, c := range extendedCommunityRtSets.ExtendedCommunityRtSet {
            if extendedCommunityRtSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet{}
        extendedCommunityRtSets.ExtendedCommunityRtSet = append(extendedCommunityRtSets.ExtendedCommunityRtSet, child)
        return &extendedCommunityRtSets.ExtendedCommunityRtSet[len(extendedCommunityRtSets.ExtendedCommunityRtSet)-1]
    }
    return nil
}

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extendedCommunityRtSets.ExtendedCommunityRtSet {
        children[extendedCommunityRtSets.ExtendedCommunityRtSet[i].GetSegmentPath()] = &extendedCommunityRtSets.ExtendedCommunityRtSet[i]
    }
    return children
}

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetYangName() string { return "extended-community-rt-sets" }

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) SetParent(parent types.Entity) { extendedCommunityRtSets.parent = parent }

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetParent() types.Entity { return extendedCommunityRtSets.parent }

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Extended Community RT Set. The type is string. This attribute is mandatory.
    RplExtendedCommunityRtSet interface{}
}

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetFilter() yfilter.YFilter { return extendedCommunityRtSet.YFilter }

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) SetFilter(yf yfilter.YFilter) { extendedCommunityRtSet.YFilter = yf }

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "rpl-extended-community-rt-set" { return "RplExtendedCommunityRtSet" }
    return ""
}

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetSegmentPath() string {
    return "extended-community-rt-set" + "[set-name='" + fmt.Sprintf("%v", extendedCommunityRtSet.SetName) + "']"
}

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = extendedCommunityRtSet.SetName
    leafs["rpl-extended-community-rt-set"] = extendedCommunityRtSet.RplExtendedCommunityRtSet
    return leafs
}

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetYangName() string { return "extended-community-rt-set" }

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) SetParent(parent types.Entity) { extendedCommunityRtSet.parent = parent }

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetParent() types.Entity { return extendedCommunityRtSet.parent }

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetParentYangName() string { return "extended-community-rt-sets" }

// RoutingPolicy_Limits
// Limits for Routing Policy
type RoutingPolicy_Limits struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of lines of policy configuration that may be configured in
    // total. The type is interface{} with range: -2147483648..2147483647. The
    // default value is 131072.
    MaximumLinesOfPolicy interface{}

    // Maximum number of policies that may be configured. The type is interface{}
    // with range: -2147483648..2147483647. The default value is 5000.
    MaximumNumberOfPolicies interface{}
}

func (limits *RoutingPolicy_Limits) GetFilter() yfilter.YFilter { return limits.YFilter }

func (limits *RoutingPolicy_Limits) SetFilter(yf yfilter.YFilter) { limits.YFilter = yf }

func (limits *RoutingPolicy_Limits) GetGoName(yname string) string {
    if yname == "maximum-lines-of-policy" { return "MaximumLinesOfPolicy" }
    if yname == "maximum-number-of-policies" { return "MaximumNumberOfPolicies" }
    return ""
}

func (limits *RoutingPolicy_Limits) GetSegmentPath() string {
    return "limits"
}

func (limits *RoutingPolicy_Limits) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (limits *RoutingPolicy_Limits) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (limits *RoutingPolicy_Limits) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-lines-of-policy"] = limits.MaximumLinesOfPolicy
    leafs["maximum-number-of-policies"] = limits.MaximumNumberOfPolicies
    return leafs
}

func (limits *RoutingPolicy_Limits) GetBundleName() string { return "cisco_ios_xr" }

func (limits *RoutingPolicy_Limits) GetYangName() string { return "limits" }

func (limits *RoutingPolicy_Limits) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (limits *RoutingPolicy_Limits) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (limits *RoutingPolicy_Limits) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (limits *RoutingPolicy_Limits) SetParent(parent types.Entity) { limits.parent = parent }

func (limits *RoutingPolicy_Limits) GetParent() types.Entity { return limits.parent }

func (limits *RoutingPolicy_Limits) GetParentYangName() string { return "routing-policy" }

