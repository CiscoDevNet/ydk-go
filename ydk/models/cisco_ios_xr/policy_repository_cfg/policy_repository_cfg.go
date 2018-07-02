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
    EntityData types.CommonEntityData
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

func (routingPolicy *RoutingPolicy) GetEntityData() *types.CommonEntityData {
    routingPolicy.EntityData.YFilter = routingPolicy.YFilter
    routingPolicy.EntityData.YangName = "routing-policy"
    routingPolicy.EntityData.BundleName = "cisco_ios_xr"
    routingPolicy.EntityData.ParentYangName = "Cisco-IOS-XR-policy-repository-cfg"
    routingPolicy.EntityData.SegmentPath = "Cisco-IOS-XR-policy-repository-cfg:routing-policy"
    routingPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingPolicy.EntityData.Children = types.NewOrderedMap()
    routingPolicy.EntityData.Children.Append("route-policies", types.YChild{"RoutePolicies", &routingPolicy.RoutePolicies})
    routingPolicy.EntityData.Children.Append("sets", types.YChild{"Sets", &routingPolicy.Sets})
    routingPolicy.EntityData.Children.Append("limits", types.YChild{"Limits", &routingPolicy.Limits})
    routingPolicy.EntityData.Leafs = types.NewOrderedMap()
    routingPolicy.EntityData.Leafs.Append("set-exit-as-abort", types.YLeaf{"SetExitAsAbort", routingPolicy.SetExitAsAbort})
    routingPolicy.EntityData.Leafs.Append("editor", types.YLeaf{"Editor", routingPolicy.Editor})

    routingPolicy.EntityData.YListKeys = []string {}

    return &(routingPolicy.EntityData)
}

// RoutingPolicy_RoutePolicies
// All configured policies
type RoutingPolicy_RoutePolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual policy. The type is slice of
    // RoutingPolicy_RoutePolicies_RoutePolicy.
    RoutePolicy []*RoutingPolicy_RoutePolicies_RoutePolicy
}

func (routePolicies *RoutingPolicy_RoutePolicies) GetEntityData() *types.CommonEntityData {
    routePolicies.EntityData.YFilter = routePolicies.YFilter
    routePolicies.EntityData.YangName = "route-policies"
    routePolicies.EntityData.BundleName = "cisco_ios_xr"
    routePolicies.EntityData.ParentYangName = "routing-policy"
    routePolicies.EntityData.SegmentPath = "route-policies"
    routePolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routePolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routePolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routePolicies.EntityData.Children = types.NewOrderedMap()
    routePolicies.EntityData.Children.Append("route-policy", types.YChild{"RoutePolicy", nil})
    for i := range routePolicies.RoutePolicy {
        routePolicies.EntityData.Children.Append(types.GetSegmentPath(routePolicies.RoutePolicy[i]), types.YChild{"RoutePolicy", routePolicies.RoutePolicy[i]})
    }
    routePolicies.EntityData.Leafs = types.NewOrderedMap()

    routePolicies.EntityData.YListKeys = []string {}

    return &(routePolicies.EntityData)
}

// RoutingPolicy_RoutePolicies_RoutePolicy
// Information about an individual policy
type RoutingPolicy_RoutePolicies_RoutePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Route policy name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    RoutePolicyName interface{}

    // policy statements. The type is string. This attribute is mandatory.
    RplRoutePolicy interface{}
}

func (routePolicy *RoutingPolicy_RoutePolicies_RoutePolicy) GetEntityData() *types.CommonEntityData {
    routePolicy.EntityData.YFilter = routePolicy.YFilter
    routePolicy.EntityData.YangName = "route-policy"
    routePolicy.EntityData.BundleName = "cisco_ios_xr"
    routePolicy.EntityData.ParentYangName = "route-policies"
    routePolicy.EntityData.SegmentPath = "route-policy" + types.AddKeyToken(routePolicy.RoutePolicyName, "route-policy-name")
    routePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routePolicy.EntityData.Children = types.NewOrderedMap()
    routePolicy.EntityData.Leafs = types.NewOrderedMap()
    routePolicy.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", routePolicy.RoutePolicyName})
    routePolicy.EntityData.Leafs.Append("rpl-route-policy", types.YLeaf{"RplRoutePolicy", routePolicy.RplRoutePolicy})

    routePolicy.EntityData.YListKeys = []string {"RoutePolicyName"}

    return &(routePolicy.EntityData)
}

// RoutingPolicy_Sets
// All configured sets
type RoutingPolicy_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about Prefix sets.
    PrefixSets RoutingPolicy_Sets_PrefixSets

    // Information about Large Community sets.
    LargeCommunitySets RoutingPolicy_Sets_LargeCommunitySets

    // Information about Mac sets.
    MacSets RoutingPolicy_Sets_MacSets

    // Information about Opaque sets.
    ExtendedCommunityOpaqueSets RoutingPolicy_Sets_ExtendedCommunityOpaqueSets

    // Information about OSPF Area sets.
    OspfAreaSets RoutingPolicy_Sets_OspfAreaSets

    // Information about Cost sets.
    ExtendedCommunityCostSets RoutingPolicy_Sets_ExtendedCommunityCostSets

    // Information about SOO sets.
    ExtendedCommunitySooSets RoutingPolicy_Sets_ExtendedCommunitySooSets

    // Information about Esi sets.
    EsiSets RoutingPolicy_Sets_EsiSets

    // Information about SegNH sets.
    ExtendedCommunitySegNhSets RoutingPolicy_Sets_ExtendedCommunitySegNhSets

    // Information about RD sets.
    RdSets RoutingPolicy_Sets_RdSets

    // Information about PolicyGlobal sets.
    PolicyGlobalSetTable RoutingPolicy_Sets_PolicyGlobalSetTable

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

func (sets *RoutingPolicy_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "routing-policy"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("prefix-sets", types.YChild{"PrefixSets", &sets.PrefixSets})
    sets.EntityData.Children.Append("large-community-sets", types.YChild{"LargeCommunitySets", &sets.LargeCommunitySets})
    sets.EntityData.Children.Append("mac-sets", types.YChild{"MacSets", &sets.MacSets})
    sets.EntityData.Children.Append("extended-community-opaque-sets", types.YChild{"ExtendedCommunityOpaqueSets", &sets.ExtendedCommunityOpaqueSets})
    sets.EntityData.Children.Append("ospf-area-sets", types.YChild{"OspfAreaSets", &sets.OspfAreaSets})
    sets.EntityData.Children.Append("extended-community-cost-sets", types.YChild{"ExtendedCommunityCostSets", &sets.ExtendedCommunityCostSets})
    sets.EntityData.Children.Append("extended-community-soo-sets", types.YChild{"ExtendedCommunitySooSets", &sets.ExtendedCommunitySooSets})
    sets.EntityData.Children.Append("esi-sets", types.YChild{"EsiSets", &sets.EsiSets})
    sets.EntityData.Children.Append("extended-community-seg-nh-sets", types.YChild{"ExtendedCommunitySegNhSets", &sets.ExtendedCommunitySegNhSets})
    sets.EntityData.Children.Append("rd-sets", types.YChild{"RdSets", &sets.RdSets})
    sets.EntityData.Children.Append("policy-global-set-table", types.YChild{"PolicyGlobalSetTable", &sets.PolicyGlobalSetTable})
    sets.EntityData.Children.Append("extended-community-bandwidth-sets", types.YChild{"ExtendedCommunityBandwidthSets", &sets.ExtendedCommunityBandwidthSets})
    sets.EntityData.Children.Append("community-sets", types.YChild{"CommunitySets", &sets.CommunitySets})
    sets.EntityData.Children.Append("as-path-sets", types.YChild{"AsPathSets", &sets.AsPathSets})
    sets.EntityData.Children.Append("tag-sets", types.YChild{"TagSets", &sets.TagSets})
    sets.EntityData.Children.Append("etag-sets", types.YChild{"EtagSets", &sets.EtagSets})
    sets.EntityData.Children.Append("extended-community-rt-sets", types.YChild{"ExtendedCommunityRtSets", &sets.ExtendedCommunityRtSets})
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_PrefixSets
// Information about Prefix sets
type RoutingPolicy_Sets_PrefixSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_PrefixSets_PrefixSet.
    PrefixSet []*RoutingPolicy_Sets_PrefixSets_PrefixSet
}

func (prefixSets *RoutingPolicy_Sets_PrefixSets) GetEntityData() *types.CommonEntityData {
    prefixSets.EntityData.YFilter = prefixSets.YFilter
    prefixSets.EntityData.YangName = "prefix-sets"
    prefixSets.EntityData.BundleName = "cisco_ios_xr"
    prefixSets.EntityData.ParentYangName = "sets"
    prefixSets.EntityData.SegmentPath = "prefix-sets"
    prefixSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixSets.EntityData.Children = types.NewOrderedMap()
    prefixSets.EntityData.Children.Append("prefix-set", types.YChild{"PrefixSet", nil})
    for i := range prefixSets.PrefixSet {
        prefixSets.EntityData.Children.Append(types.GetSegmentPath(prefixSets.PrefixSet[i]), types.YChild{"PrefixSet", prefixSets.PrefixSet[i]})
    }
    prefixSets.EntityData.Leafs = types.NewOrderedMap()

    prefixSets.EntityData.YListKeys = []string {}

    return &(prefixSets.EntityData)
}

// RoutingPolicy_Sets_PrefixSets_PrefixSet
// Information about an individual set
type RoutingPolicy_Sets_PrefixSets_PrefixSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // prefix statements. The type is string. This attribute is mandatory.
    RplPrefixSet interface{}
}

func (prefixSet *RoutingPolicy_Sets_PrefixSets_PrefixSet) GetEntityData() *types.CommonEntityData {
    prefixSet.EntityData.YFilter = prefixSet.YFilter
    prefixSet.EntityData.YangName = "prefix-set"
    prefixSet.EntityData.BundleName = "cisco_ios_xr"
    prefixSet.EntityData.ParentYangName = "prefix-sets"
    prefixSet.EntityData.SegmentPath = "prefix-set" + types.AddKeyToken(prefixSet.SetName, "set-name")
    prefixSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixSet.EntityData.Children = types.NewOrderedMap()
    prefixSet.EntityData.Leafs = types.NewOrderedMap()
    prefixSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", prefixSet.SetName})
    prefixSet.EntityData.Leafs.Append("rpl-prefix-set", types.YLeaf{"RplPrefixSet", prefixSet.RplPrefixSet})

    prefixSet.EntityData.YListKeys = []string {"SetName"}

    return &(prefixSet.EntityData)
}

// RoutingPolicy_Sets_LargeCommunitySets
// Information about Large Community sets
type RoutingPolicy_Sets_LargeCommunitySets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet.
    LargeCommunitySet []*RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet
}

func (largeCommunitySets *RoutingPolicy_Sets_LargeCommunitySets) GetEntityData() *types.CommonEntityData {
    largeCommunitySets.EntityData.YFilter = largeCommunitySets.YFilter
    largeCommunitySets.EntityData.YangName = "large-community-sets"
    largeCommunitySets.EntityData.BundleName = "cisco_ios_xr"
    largeCommunitySets.EntityData.ParentYangName = "sets"
    largeCommunitySets.EntityData.SegmentPath = "large-community-sets"
    largeCommunitySets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    largeCommunitySets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    largeCommunitySets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    largeCommunitySets.EntityData.Children = types.NewOrderedMap()
    largeCommunitySets.EntityData.Children.Append("large-community-set", types.YChild{"LargeCommunitySet", nil})
    for i := range largeCommunitySets.LargeCommunitySet {
        largeCommunitySets.EntityData.Children.Append(types.GetSegmentPath(largeCommunitySets.LargeCommunitySet[i]), types.YChild{"LargeCommunitySet", largeCommunitySets.LargeCommunitySet[i]})
    }
    largeCommunitySets.EntityData.Leafs = types.NewOrderedMap()

    largeCommunitySets.EntityData.YListKeys = []string {}

    return &(largeCommunitySets.EntityData)
}

// RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet
// Information about an individual set
type RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Large Community Set. The type is string. This attribute is mandatory.
    LargeCommunitySetAsText interface{}
}

func (largeCommunitySet *RoutingPolicy_Sets_LargeCommunitySets_LargeCommunitySet) GetEntityData() *types.CommonEntityData {
    largeCommunitySet.EntityData.YFilter = largeCommunitySet.YFilter
    largeCommunitySet.EntityData.YangName = "large-community-set"
    largeCommunitySet.EntityData.BundleName = "cisco_ios_xr"
    largeCommunitySet.EntityData.ParentYangName = "large-community-sets"
    largeCommunitySet.EntityData.SegmentPath = "large-community-set" + types.AddKeyToken(largeCommunitySet.SetName, "set-name")
    largeCommunitySet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    largeCommunitySet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    largeCommunitySet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    largeCommunitySet.EntityData.Children = types.NewOrderedMap()
    largeCommunitySet.EntityData.Leafs = types.NewOrderedMap()
    largeCommunitySet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", largeCommunitySet.SetName})
    largeCommunitySet.EntityData.Leafs.Append("large-community-set-as-text", types.YLeaf{"LargeCommunitySetAsText", largeCommunitySet.LargeCommunitySetAsText})

    largeCommunitySet.EntityData.YListKeys = []string {"SetName"}

    return &(largeCommunitySet.EntityData)
}

// RoutingPolicy_Sets_MacSets
// Information about Mac sets
type RoutingPolicy_Sets_MacSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_MacSets_MacSet.
    MacSet []*RoutingPolicy_Sets_MacSets_MacSet
}

func (macSets *RoutingPolicy_Sets_MacSets) GetEntityData() *types.CommonEntityData {
    macSets.EntityData.YFilter = macSets.YFilter
    macSets.EntityData.YangName = "mac-sets"
    macSets.EntityData.BundleName = "cisco_ios_xr"
    macSets.EntityData.ParentYangName = "sets"
    macSets.EntityData.SegmentPath = "mac-sets"
    macSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macSets.EntityData.Children = types.NewOrderedMap()
    macSets.EntityData.Children.Append("mac-set", types.YChild{"MacSet", nil})
    for i := range macSets.MacSet {
        macSets.EntityData.Children.Append(types.GetSegmentPath(macSets.MacSet[i]), types.YChild{"MacSet", macSets.MacSet[i]})
    }
    macSets.EntityData.Leafs = types.NewOrderedMap()

    macSets.EntityData.YListKeys = []string {}

    return &(macSets.EntityData)
}

// RoutingPolicy_Sets_MacSets_MacSet
// Information about an individual set
type RoutingPolicy_Sets_MacSets_MacSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Mac Set. The type is string. This attribute is mandatory.
    MacSetAsText interface{}
}

func (macSet *RoutingPolicy_Sets_MacSets_MacSet) GetEntityData() *types.CommonEntityData {
    macSet.EntityData.YFilter = macSet.YFilter
    macSet.EntityData.YangName = "mac-set"
    macSet.EntityData.BundleName = "cisco_ios_xr"
    macSet.EntityData.ParentYangName = "mac-sets"
    macSet.EntityData.SegmentPath = "mac-set" + types.AddKeyToken(macSet.SetName, "set-name")
    macSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macSet.EntityData.Children = types.NewOrderedMap()
    macSet.EntityData.Leafs = types.NewOrderedMap()
    macSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", macSet.SetName})
    macSet.EntityData.Leafs.Append("mac-set-as-text", types.YLeaf{"MacSetAsText", macSet.MacSetAsText})

    macSet.EntityData.YListKeys = []string {"SetName"}

    return &(macSet.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaqueSets
// Information about Opaque sets
type RoutingPolicy_Sets_ExtendedCommunityOpaqueSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet.
    ExtendedCommunityOpaqueSet []*RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet
}

func (extendedCommunityOpaqueSets *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets) GetEntityData() *types.CommonEntityData {
    extendedCommunityOpaqueSets.EntityData.YFilter = extendedCommunityOpaqueSets.YFilter
    extendedCommunityOpaqueSets.EntityData.YangName = "extended-community-opaque-sets"
    extendedCommunityOpaqueSets.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityOpaqueSets.EntityData.ParentYangName = "sets"
    extendedCommunityOpaqueSets.EntityData.SegmentPath = "extended-community-opaque-sets"
    extendedCommunityOpaqueSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityOpaqueSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityOpaqueSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityOpaqueSets.EntityData.Children = types.NewOrderedMap()
    extendedCommunityOpaqueSets.EntityData.Children.Append("extended-community-opaque-set", types.YChild{"ExtendedCommunityOpaqueSet", nil})
    for i := range extendedCommunityOpaqueSets.ExtendedCommunityOpaqueSet {
        extendedCommunityOpaqueSets.EntityData.Children.Append(types.GetSegmentPath(extendedCommunityOpaqueSets.ExtendedCommunityOpaqueSet[i]), types.YChild{"ExtendedCommunityOpaqueSet", extendedCommunityOpaqueSets.ExtendedCommunityOpaqueSet[i]})
    }
    extendedCommunityOpaqueSets.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunityOpaqueSets.EntityData.YListKeys = []string {}

    return &(extendedCommunityOpaqueSets.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Extended Community Opaque Set. The type is string. This attribute is
    // mandatory.
    RplExtendedCommunityOpaqueSet interface{}
}

func (extendedCommunityOpaqueSet *RoutingPolicy_Sets_ExtendedCommunityOpaqueSets_ExtendedCommunityOpaqueSet) GetEntityData() *types.CommonEntityData {
    extendedCommunityOpaqueSet.EntityData.YFilter = extendedCommunityOpaqueSet.YFilter
    extendedCommunityOpaqueSet.EntityData.YangName = "extended-community-opaque-set"
    extendedCommunityOpaqueSet.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityOpaqueSet.EntityData.ParentYangName = "extended-community-opaque-sets"
    extendedCommunityOpaqueSet.EntityData.SegmentPath = "extended-community-opaque-set" + types.AddKeyToken(extendedCommunityOpaqueSet.SetName, "set-name")
    extendedCommunityOpaqueSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityOpaqueSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityOpaqueSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityOpaqueSet.EntityData.Children = types.NewOrderedMap()
    extendedCommunityOpaqueSet.EntityData.Leafs = types.NewOrderedMap()
    extendedCommunityOpaqueSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", extendedCommunityOpaqueSet.SetName})
    extendedCommunityOpaqueSet.EntityData.Leafs.Append("rpl-extended-community-opaque-set", types.YLeaf{"RplExtendedCommunityOpaqueSet", extendedCommunityOpaqueSet.RplExtendedCommunityOpaqueSet})

    extendedCommunityOpaqueSet.EntityData.YListKeys = []string {"SetName"}

    return &(extendedCommunityOpaqueSet.EntityData)
}

// RoutingPolicy_Sets_OspfAreaSets
// Information about OSPF Area sets
type RoutingPolicy_Sets_OspfAreaSets struct {
    EntityData types.CommonEntityData
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
    OspfAreaSet []*RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet
}

func (ospfAreaSets *RoutingPolicy_Sets_OspfAreaSets) GetEntityData() *types.CommonEntityData {
    ospfAreaSets.EntityData.YFilter = ospfAreaSets.YFilter
    ospfAreaSets.EntityData.YangName = "ospf-area-sets"
    ospfAreaSets.EntityData.BundleName = "cisco_ios_xr"
    ospfAreaSets.EntityData.ParentYangName = "sets"
    ospfAreaSets.EntityData.SegmentPath = "ospf-area-sets"
    ospfAreaSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfAreaSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfAreaSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfAreaSets.EntityData.Children = types.NewOrderedMap()
    ospfAreaSets.EntityData.Children.Append("ospf-area-set", types.YChild{"OspfAreaSet", nil})
    for i := range ospfAreaSets.OspfAreaSet {
        ospfAreaSets.EntityData.Children.Append(types.GetSegmentPath(ospfAreaSets.OspfAreaSet[i]), types.YChild{"OspfAreaSet", ospfAreaSets.OspfAreaSet[i]})
    }
    ospfAreaSets.EntityData.Leafs = types.NewOrderedMap()

    ospfAreaSets.EntityData.YListKeys = []string {}

    return &(ospfAreaSets.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // OSPF Area Set. The type is string. This attribute is mandatory.
    RplospfAreaSet interface{}
}

func (ospfAreaSet *RoutingPolicy_Sets_OspfAreaSets_OspfAreaSet) GetEntityData() *types.CommonEntityData {
    ospfAreaSet.EntityData.YFilter = ospfAreaSet.YFilter
    ospfAreaSet.EntityData.YangName = "ospf-area-set"
    ospfAreaSet.EntityData.BundleName = "cisco_ios_xr"
    ospfAreaSet.EntityData.ParentYangName = "ospf-area-sets"
    ospfAreaSet.EntityData.SegmentPath = "ospf-area-set" + types.AddKeyToken(ospfAreaSet.SetName, "set-name")
    ospfAreaSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfAreaSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfAreaSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfAreaSet.EntityData.Children = types.NewOrderedMap()
    ospfAreaSet.EntityData.Leafs = types.NewOrderedMap()
    ospfAreaSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", ospfAreaSet.SetName})
    ospfAreaSet.EntityData.Leafs.Append("rplospf-area-set", types.YLeaf{"RplospfAreaSet", ospfAreaSet.RplospfAreaSet})

    ospfAreaSet.EntityData.YListKeys = []string {"SetName"}

    return &(ospfAreaSet.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCostSets
// Information about Cost sets
type RoutingPolicy_Sets_ExtendedCommunityCostSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet.
    ExtendedCommunityCostSet []*RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet
}

func (extendedCommunityCostSets *RoutingPolicy_Sets_ExtendedCommunityCostSets) GetEntityData() *types.CommonEntityData {
    extendedCommunityCostSets.EntityData.YFilter = extendedCommunityCostSets.YFilter
    extendedCommunityCostSets.EntityData.YangName = "extended-community-cost-sets"
    extendedCommunityCostSets.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityCostSets.EntityData.ParentYangName = "sets"
    extendedCommunityCostSets.EntityData.SegmentPath = "extended-community-cost-sets"
    extendedCommunityCostSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityCostSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityCostSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityCostSets.EntityData.Children = types.NewOrderedMap()
    extendedCommunityCostSets.EntityData.Children.Append("extended-community-cost-set", types.YChild{"ExtendedCommunityCostSet", nil})
    for i := range extendedCommunityCostSets.ExtendedCommunityCostSet {
        extendedCommunityCostSets.EntityData.Children.Append(types.GetSegmentPath(extendedCommunityCostSets.ExtendedCommunityCostSet[i]), types.YChild{"ExtendedCommunityCostSet", extendedCommunityCostSets.ExtendedCommunityCostSet[i]})
    }
    extendedCommunityCostSets.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunityCostSets.EntityData.YListKeys = []string {}

    return &(extendedCommunityCostSets.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Extended Community Cost Set. The type is string. This attribute is
    // mandatory.
    RplExtendedCommunityCostSet interface{}
}

func (extendedCommunityCostSet *RoutingPolicy_Sets_ExtendedCommunityCostSets_ExtendedCommunityCostSet) GetEntityData() *types.CommonEntityData {
    extendedCommunityCostSet.EntityData.YFilter = extendedCommunityCostSet.YFilter
    extendedCommunityCostSet.EntityData.YangName = "extended-community-cost-set"
    extendedCommunityCostSet.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityCostSet.EntityData.ParentYangName = "extended-community-cost-sets"
    extendedCommunityCostSet.EntityData.SegmentPath = "extended-community-cost-set" + types.AddKeyToken(extendedCommunityCostSet.SetName, "set-name")
    extendedCommunityCostSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityCostSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityCostSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityCostSet.EntityData.Children = types.NewOrderedMap()
    extendedCommunityCostSet.EntityData.Leafs = types.NewOrderedMap()
    extendedCommunityCostSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", extendedCommunityCostSet.SetName})
    extendedCommunityCostSet.EntityData.Leafs.Append("rpl-extended-community-cost-set", types.YLeaf{"RplExtendedCommunityCostSet", extendedCommunityCostSet.RplExtendedCommunityCostSet})

    extendedCommunityCostSet.EntityData.YListKeys = []string {"SetName"}

    return &(extendedCommunityCostSet.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySooSets
// Information about SOO sets
type RoutingPolicy_Sets_ExtendedCommunitySooSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet.
    ExtendedCommunitySooSet []*RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet
}

func (extendedCommunitySooSets *RoutingPolicy_Sets_ExtendedCommunitySooSets) GetEntityData() *types.CommonEntityData {
    extendedCommunitySooSets.EntityData.YFilter = extendedCommunitySooSets.YFilter
    extendedCommunitySooSets.EntityData.YangName = "extended-community-soo-sets"
    extendedCommunitySooSets.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunitySooSets.EntityData.ParentYangName = "sets"
    extendedCommunitySooSets.EntityData.SegmentPath = "extended-community-soo-sets"
    extendedCommunitySooSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunitySooSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunitySooSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunitySooSets.EntityData.Children = types.NewOrderedMap()
    extendedCommunitySooSets.EntityData.Children.Append("extended-community-soo-set", types.YChild{"ExtendedCommunitySooSet", nil})
    for i := range extendedCommunitySooSets.ExtendedCommunitySooSet {
        extendedCommunitySooSets.EntityData.Children.Append(types.GetSegmentPath(extendedCommunitySooSets.ExtendedCommunitySooSet[i]), types.YChild{"ExtendedCommunitySooSet", extendedCommunitySooSets.ExtendedCommunitySooSet[i]})
    }
    extendedCommunitySooSets.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunitySooSets.EntityData.YListKeys = []string {}

    return &(extendedCommunitySooSets.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Extended Community SOO Set. The type is string. This attribute is
    // mandatory.
    RplExtendedCommunitySooSet interface{}
}

func (extendedCommunitySooSet *RoutingPolicy_Sets_ExtendedCommunitySooSets_ExtendedCommunitySooSet) GetEntityData() *types.CommonEntityData {
    extendedCommunitySooSet.EntityData.YFilter = extendedCommunitySooSet.YFilter
    extendedCommunitySooSet.EntityData.YangName = "extended-community-soo-set"
    extendedCommunitySooSet.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunitySooSet.EntityData.ParentYangName = "extended-community-soo-sets"
    extendedCommunitySooSet.EntityData.SegmentPath = "extended-community-soo-set" + types.AddKeyToken(extendedCommunitySooSet.SetName, "set-name")
    extendedCommunitySooSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunitySooSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunitySooSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunitySooSet.EntityData.Children = types.NewOrderedMap()
    extendedCommunitySooSet.EntityData.Leafs = types.NewOrderedMap()
    extendedCommunitySooSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", extendedCommunitySooSet.SetName})
    extendedCommunitySooSet.EntityData.Leafs.Append("rpl-extended-community-soo-set", types.YLeaf{"RplExtendedCommunitySooSet", extendedCommunitySooSet.RplExtendedCommunitySooSet})

    extendedCommunitySooSet.EntityData.YListKeys = []string {"SetName"}

    return &(extendedCommunitySooSet.EntityData)
}

// RoutingPolicy_Sets_EsiSets
// Information about Esi sets
type RoutingPolicy_Sets_EsiSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_EsiSets_EsiSet.
    EsiSet []*RoutingPolicy_Sets_EsiSets_EsiSet
}

func (esiSets *RoutingPolicy_Sets_EsiSets) GetEntityData() *types.CommonEntityData {
    esiSets.EntityData.YFilter = esiSets.YFilter
    esiSets.EntityData.YangName = "esi-sets"
    esiSets.EntityData.BundleName = "cisco_ios_xr"
    esiSets.EntityData.ParentYangName = "sets"
    esiSets.EntityData.SegmentPath = "esi-sets"
    esiSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esiSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esiSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esiSets.EntityData.Children = types.NewOrderedMap()
    esiSets.EntityData.Children.Append("esi-set", types.YChild{"EsiSet", nil})
    for i := range esiSets.EsiSet {
        esiSets.EntityData.Children.Append(types.GetSegmentPath(esiSets.EsiSet[i]), types.YChild{"EsiSet", esiSets.EsiSet[i]})
    }
    esiSets.EntityData.Leafs = types.NewOrderedMap()

    esiSets.EntityData.YListKeys = []string {}

    return &(esiSets.EntityData)
}

// RoutingPolicy_Sets_EsiSets_EsiSet
// Information about an individual set
type RoutingPolicy_Sets_EsiSets_EsiSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Esi Set. The type is string. This attribute is mandatory.
    EsiSetAsText interface{}
}

func (esiSet *RoutingPolicy_Sets_EsiSets_EsiSet) GetEntityData() *types.CommonEntityData {
    esiSet.EntityData.YFilter = esiSet.YFilter
    esiSet.EntityData.YangName = "esi-set"
    esiSet.EntityData.BundleName = "cisco_ios_xr"
    esiSet.EntityData.ParentYangName = "esi-sets"
    esiSet.EntityData.SegmentPath = "esi-set" + types.AddKeyToken(esiSet.SetName, "set-name")
    esiSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esiSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esiSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esiSet.EntityData.Children = types.NewOrderedMap()
    esiSet.EntityData.Leafs = types.NewOrderedMap()
    esiSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", esiSet.SetName})
    esiSet.EntityData.Leafs.Append("esi-set-as-text", types.YLeaf{"EsiSetAsText", esiSet.EsiSetAsText})

    esiSet.EntityData.YListKeys = []string {"SetName"}

    return &(esiSet.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNhSets
// Information about SegNH sets
type RoutingPolicy_Sets_ExtendedCommunitySegNhSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet.
    ExtendedCommunitySegNhSet []*RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet
}

func (extendedCommunitySegNhSets *RoutingPolicy_Sets_ExtendedCommunitySegNhSets) GetEntityData() *types.CommonEntityData {
    extendedCommunitySegNhSets.EntityData.YFilter = extendedCommunitySegNhSets.YFilter
    extendedCommunitySegNhSets.EntityData.YangName = "extended-community-seg-nh-sets"
    extendedCommunitySegNhSets.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunitySegNhSets.EntityData.ParentYangName = "sets"
    extendedCommunitySegNhSets.EntityData.SegmentPath = "extended-community-seg-nh-sets"
    extendedCommunitySegNhSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunitySegNhSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunitySegNhSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunitySegNhSets.EntityData.Children = types.NewOrderedMap()
    extendedCommunitySegNhSets.EntityData.Children.Append("extended-community-seg-nh-set", types.YChild{"ExtendedCommunitySegNhSet", nil})
    for i := range extendedCommunitySegNhSets.ExtendedCommunitySegNhSet {
        extendedCommunitySegNhSets.EntityData.Children.Append(types.GetSegmentPath(extendedCommunitySegNhSets.ExtendedCommunitySegNhSet[i]), types.YChild{"ExtendedCommunitySegNhSet", extendedCommunitySegNhSets.ExtendedCommunitySegNhSet[i]})
    }
    extendedCommunitySegNhSets.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunitySegNhSets.EntityData.YListKeys = []string {}

    return &(extendedCommunitySegNhSets.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Extended Community SegNH Set. The type is string. This attribute is
    // mandatory.
    RplExtendedCommunitySegNhSet interface{}
}

func (extendedCommunitySegNhSet *RoutingPolicy_Sets_ExtendedCommunitySegNhSets_ExtendedCommunitySegNhSet) GetEntityData() *types.CommonEntityData {
    extendedCommunitySegNhSet.EntityData.YFilter = extendedCommunitySegNhSet.YFilter
    extendedCommunitySegNhSet.EntityData.YangName = "extended-community-seg-nh-set"
    extendedCommunitySegNhSet.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunitySegNhSet.EntityData.ParentYangName = "extended-community-seg-nh-sets"
    extendedCommunitySegNhSet.EntityData.SegmentPath = "extended-community-seg-nh-set" + types.AddKeyToken(extendedCommunitySegNhSet.SetName, "set-name")
    extendedCommunitySegNhSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunitySegNhSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunitySegNhSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunitySegNhSet.EntityData.Children = types.NewOrderedMap()
    extendedCommunitySegNhSet.EntityData.Leafs = types.NewOrderedMap()
    extendedCommunitySegNhSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", extendedCommunitySegNhSet.SetName})
    extendedCommunitySegNhSet.EntityData.Leafs.Append("rpl-extended-community-seg-nh-set", types.YLeaf{"RplExtendedCommunitySegNhSet", extendedCommunitySegNhSet.RplExtendedCommunitySegNhSet})

    extendedCommunitySegNhSet.EntityData.YListKeys = []string {"SetName"}

    return &(extendedCommunitySegNhSet.EntityData)
}

// RoutingPolicy_Sets_RdSets
// Information about RD sets
type RoutingPolicy_Sets_RdSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_RdSets_RdSet.
    RdSet []*RoutingPolicy_Sets_RdSets_RdSet
}

func (rdSets *RoutingPolicy_Sets_RdSets) GetEntityData() *types.CommonEntityData {
    rdSets.EntityData.YFilter = rdSets.YFilter
    rdSets.EntityData.YangName = "rd-sets"
    rdSets.EntityData.BundleName = "cisco_ios_xr"
    rdSets.EntityData.ParentYangName = "sets"
    rdSets.EntityData.SegmentPath = "rd-sets"
    rdSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rdSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rdSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rdSets.EntityData.Children = types.NewOrderedMap()
    rdSets.EntityData.Children.Append("rd-set", types.YChild{"RdSet", nil})
    for i := range rdSets.RdSet {
        rdSets.EntityData.Children.Append(types.GetSegmentPath(rdSets.RdSet[i]), types.YChild{"RdSet", rdSets.RdSet[i]})
    }
    rdSets.EntityData.Leafs = types.NewOrderedMap()

    rdSets.EntityData.YListKeys = []string {}

    return &(rdSets.EntityData)
}

// RoutingPolicy_Sets_RdSets_RdSet
// Information about an individual set
type RoutingPolicy_Sets_RdSets_RdSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // RD Set. The type is string. This attribute is mandatory.
    RplrdSet interface{}
}

func (rdSet *RoutingPolicy_Sets_RdSets_RdSet) GetEntityData() *types.CommonEntityData {
    rdSet.EntityData.YFilter = rdSet.YFilter
    rdSet.EntityData.YangName = "rd-set"
    rdSet.EntityData.BundleName = "cisco_ios_xr"
    rdSet.EntityData.ParentYangName = "rd-sets"
    rdSet.EntityData.SegmentPath = "rd-set" + types.AddKeyToken(rdSet.SetName, "set-name")
    rdSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rdSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rdSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rdSet.EntityData.Children = types.NewOrderedMap()
    rdSet.EntityData.Leafs = types.NewOrderedMap()
    rdSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", rdSet.SetName})
    rdSet.EntityData.Leafs.Append("rplrd-set", types.YLeaf{"RplrdSet", rdSet.RplrdSet})

    rdSet.EntityData.YListKeys = []string {"SetName"}

    return &(rdSet.EntityData)
}

// RoutingPolicy_Sets_PolicyGlobalSetTable
// Information about PolicyGlobal sets
type RoutingPolicy_Sets_PolicyGlobalSetTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is string.
    PolicyGlobalSet interface{}
}

func (policyGlobalSetTable *RoutingPolicy_Sets_PolicyGlobalSetTable) GetEntityData() *types.CommonEntityData {
    policyGlobalSetTable.EntityData.YFilter = policyGlobalSetTable.YFilter
    policyGlobalSetTable.EntityData.YangName = "policy-global-set-table"
    policyGlobalSetTable.EntityData.BundleName = "cisco_ios_xr"
    policyGlobalSetTable.EntityData.ParentYangName = "sets"
    policyGlobalSetTable.EntityData.SegmentPath = "policy-global-set-table"
    policyGlobalSetTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyGlobalSetTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyGlobalSetTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyGlobalSetTable.EntityData.Children = types.NewOrderedMap()
    policyGlobalSetTable.EntityData.Leafs = types.NewOrderedMap()
    policyGlobalSetTable.EntityData.Leafs.Append("policy-global-set", types.YLeaf{"PolicyGlobalSet", policyGlobalSetTable.PolicyGlobalSet})

    policyGlobalSetTable.EntityData.YListKeys = []string {}

    return &(policyGlobalSetTable.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidthSets
// Information about Bandwidth sets
type RoutingPolicy_Sets_ExtendedCommunityBandwidthSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet.
    ExtendedCommunityBandwidthSet []*RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet
}

func (extendedCommunityBandwidthSets *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets) GetEntityData() *types.CommonEntityData {
    extendedCommunityBandwidthSets.EntityData.YFilter = extendedCommunityBandwidthSets.YFilter
    extendedCommunityBandwidthSets.EntityData.YangName = "extended-community-bandwidth-sets"
    extendedCommunityBandwidthSets.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityBandwidthSets.EntityData.ParentYangName = "sets"
    extendedCommunityBandwidthSets.EntityData.SegmentPath = "extended-community-bandwidth-sets"
    extendedCommunityBandwidthSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityBandwidthSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityBandwidthSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityBandwidthSets.EntityData.Children = types.NewOrderedMap()
    extendedCommunityBandwidthSets.EntityData.Children.Append("extended-community-bandwidth-set", types.YChild{"ExtendedCommunityBandwidthSet", nil})
    for i := range extendedCommunityBandwidthSets.ExtendedCommunityBandwidthSet {
        extendedCommunityBandwidthSets.EntityData.Children.Append(types.GetSegmentPath(extendedCommunityBandwidthSets.ExtendedCommunityBandwidthSet[i]), types.YChild{"ExtendedCommunityBandwidthSet", extendedCommunityBandwidthSets.ExtendedCommunityBandwidthSet[i]})
    }
    extendedCommunityBandwidthSets.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunityBandwidthSets.EntityData.YListKeys = []string {}

    return &(extendedCommunityBandwidthSets.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Extended Community Bandwidth Set. The type is string. This attribute is
    // mandatory.
    RplExtendedCommunityBandwidthSet interface{}
}

func (extendedCommunityBandwidthSet *RoutingPolicy_Sets_ExtendedCommunityBandwidthSets_ExtendedCommunityBandwidthSet) GetEntityData() *types.CommonEntityData {
    extendedCommunityBandwidthSet.EntityData.YFilter = extendedCommunityBandwidthSet.YFilter
    extendedCommunityBandwidthSet.EntityData.YangName = "extended-community-bandwidth-set"
    extendedCommunityBandwidthSet.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityBandwidthSet.EntityData.ParentYangName = "extended-community-bandwidth-sets"
    extendedCommunityBandwidthSet.EntityData.SegmentPath = "extended-community-bandwidth-set" + types.AddKeyToken(extendedCommunityBandwidthSet.SetName, "set-name")
    extendedCommunityBandwidthSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityBandwidthSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityBandwidthSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityBandwidthSet.EntityData.Children = types.NewOrderedMap()
    extendedCommunityBandwidthSet.EntityData.Leafs = types.NewOrderedMap()
    extendedCommunityBandwidthSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", extendedCommunityBandwidthSet.SetName})
    extendedCommunityBandwidthSet.EntityData.Leafs.Append("rpl-extended-community-bandwidth-set", types.YLeaf{"RplExtendedCommunityBandwidthSet", extendedCommunityBandwidthSet.RplExtendedCommunityBandwidthSet})

    extendedCommunityBandwidthSet.EntityData.YListKeys = []string {"SetName"}

    return &(extendedCommunityBandwidthSet.EntityData)
}

// RoutingPolicy_Sets_CommunitySets
// Information about Community sets
type RoutingPolicy_Sets_CommunitySets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_CommunitySets_CommunitySet.
    CommunitySet []*RoutingPolicy_Sets_CommunitySets_CommunitySet
}

func (communitySets *RoutingPolicy_Sets_CommunitySets) GetEntityData() *types.CommonEntityData {
    communitySets.EntityData.YFilter = communitySets.YFilter
    communitySets.EntityData.YangName = "community-sets"
    communitySets.EntityData.BundleName = "cisco_ios_xr"
    communitySets.EntityData.ParentYangName = "sets"
    communitySets.EntityData.SegmentPath = "community-sets"
    communitySets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    communitySets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    communitySets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    communitySets.EntityData.Children = types.NewOrderedMap()
    communitySets.EntityData.Children.Append("community-set", types.YChild{"CommunitySet", nil})
    for i := range communitySets.CommunitySet {
        communitySets.EntityData.Children.Append(types.GetSegmentPath(communitySets.CommunitySet[i]), types.YChild{"CommunitySet", communitySets.CommunitySet[i]})
    }
    communitySets.EntityData.Leafs = types.NewOrderedMap()

    communitySets.EntityData.YListKeys = []string {}

    return &(communitySets.EntityData)
}

// RoutingPolicy_Sets_CommunitySets_CommunitySet
// Information about an individual set
type RoutingPolicy_Sets_CommunitySets_CommunitySet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Community Set. The type is string. This attribute is mandatory.
    RplCommunitySet interface{}
}

func (communitySet *RoutingPolicy_Sets_CommunitySets_CommunitySet) GetEntityData() *types.CommonEntityData {
    communitySet.EntityData.YFilter = communitySet.YFilter
    communitySet.EntityData.YangName = "community-set"
    communitySet.EntityData.BundleName = "cisco_ios_xr"
    communitySet.EntityData.ParentYangName = "community-sets"
    communitySet.EntityData.SegmentPath = "community-set" + types.AddKeyToken(communitySet.SetName, "set-name")
    communitySet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    communitySet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    communitySet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    communitySet.EntityData.Children = types.NewOrderedMap()
    communitySet.EntityData.Leafs = types.NewOrderedMap()
    communitySet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", communitySet.SetName})
    communitySet.EntityData.Leafs.Append("rpl-community-set", types.YLeaf{"RplCommunitySet", communitySet.RplCommunitySet})

    communitySet.EntityData.YListKeys = []string {"SetName"}

    return &(communitySet.EntityData)
}

// RoutingPolicy_Sets_AsPathSets
// Information about AS Path sets
type RoutingPolicy_Sets_AsPathSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_AsPathSets_AsPathSet.
    AsPathSet []*RoutingPolicy_Sets_AsPathSets_AsPathSet
}

func (asPathSets *RoutingPolicy_Sets_AsPathSets) GetEntityData() *types.CommonEntityData {
    asPathSets.EntityData.YFilter = asPathSets.YFilter
    asPathSets.EntityData.YangName = "as-path-sets"
    asPathSets.EntityData.BundleName = "cisco_ios_xr"
    asPathSets.EntityData.ParentYangName = "sets"
    asPathSets.EntityData.SegmentPath = "as-path-sets"
    asPathSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asPathSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asPathSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asPathSets.EntityData.Children = types.NewOrderedMap()
    asPathSets.EntityData.Children.Append("as-path-set", types.YChild{"AsPathSet", nil})
    for i := range asPathSets.AsPathSet {
        asPathSets.EntityData.Children.Append(types.GetSegmentPath(asPathSets.AsPathSet[i]), types.YChild{"AsPathSet", asPathSets.AsPathSet[i]})
    }
    asPathSets.EntityData.Leafs = types.NewOrderedMap()

    asPathSets.EntityData.YListKeys = []string {}

    return &(asPathSets.EntityData)
}

// RoutingPolicy_Sets_AsPathSets_AsPathSet
// Information about an individual set
type RoutingPolicy_Sets_AsPathSets_AsPathSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // ASPath Set. The type is string. This attribute is mandatory.
    RplasPathSet interface{}
}

func (asPathSet *RoutingPolicy_Sets_AsPathSets_AsPathSet) GetEntityData() *types.CommonEntityData {
    asPathSet.EntityData.YFilter = asPathSet.YFilter
    asPathSet.EntityData.YangName = "as-path-set"
    asPathSet.EntityData.BundleName = "cisco_ios_xr"
    asPathSet.EntityData.ParentYangName = "as-path-sets"
    asPathSet.EntityData.SegmentPath = "as-path-set" + types.AddKeyToken(asPathSet.SetName, "set-name")
    asPathSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asPathSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asPathSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asPathSet.EntityData.Children = types.NewOrderedMap()
    asPathSet.EntityData.Leafs = types.NewOrderedMap()
    asPathSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", asPathSet.SetName})
    asPathSet.EntityData.Leafs.Append("rplas-path-set", types.YLeaf{"RplasPathSet", asPathSet.RplasPathSet})

    asPathSet.EntityData.YListKeys = []string {"SetName"}

    return &(asPathSet.EntityData)
}

// RoutingPolicy_Sets_TagSets
// Information about Tag sets
type RoutingPolicy_Sets_TagSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_TagSets_TagSet.
    TagSet []*RoutingPolicy_Sets_TagSets_TagSet
}

func (tagSets *RoutingPolicy_Sets_TagSets) GetEntityData() *types.CommonEntityData {
    tagSets.EntityData.YFilter = tagSets.YFilter
    tagSets.EntityData.YangName = "tag-sets"
    tagSets.EntityData.BundleName = "cisco_ios_xr"
    tagSets.EntityData.ParentYangName = "sets"
    tagSets.EntityData.SegmentPath = "tag-sets"
    tagSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tagSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tagSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tagSets.EntityData.Children = types.NewOrderedMap()
    tagSets.EntityData.Children.Append("tag-set", types.YChild{"TagSet", nil})
    for i := range tagSets.TagSet {
        tagSets.EntityData.Children.Append(types.GetSegmentPath(tagSets.TagSet[i]), types.YChild{"TagSet", tagSets.TagSet[i]})
    }
    tagSets.EntityData.Leafs = types.NewOrderedMap()

    tagSets.EntityData.YListKeys = []string {}

    return &(tagSets.EntityData)
}

// RoutingPolicy_Sets_TagSets_TagSet
// Information about an individual set
type RoutingPolicy_Sets_TagSets_TagSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Tag Set. The type is string. This attribute is mandatory.
    RplTagSet interface{}
}

func (tagSet *RoutingPolicy_Sets_TagSets_TagSet) GetEntityData() *types.CommonEntityData {
    tagSet.EntityData.YFilter = tagSet.YFilter
    tagSet.EntityData.YangName = "tag-set"
    tagSet.EntityData.BundleName = "cisco_ios_xr"
    tagSet.EntityData.ParentYangName = "tag-sets"
    tagSet.EntityData.SegmentPath = "tag-set" + types.AddKeyToken(tagSet.SetName, "set-name")
    tagSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tagSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tagSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tagSet.EntityData.Children = types.NewOrderedMap()
    tagSet.EntityData.Leafs = types.NewOrderedMap()
    tagSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", tagSet.SetName})
    tagSet.EntityData.Leafs.Append("rpl-tag-set", types.YLeaf{"RplTagSet", tagSet.RplTagSet})

    tagSet.EntityData.YListKeys = []string {"SetName"}

    return &(tagSet.EntityData)
}

// RoutingPolicy_Sets_EtagSets
// Information about Etag sets
type RoutingPolicy_Sets_EtagSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_EtagSets_EtagSet.
    EtagSet []*RoutingPolicy_Sets_EtagSets_EtagSet
}

func (etagSets *RoutingPolicy_Sets_EtagSets) GetEntityData() *types.CommonEntityData {
    etagSets.EntityData.YFilter = etagSets.YFilter
    etagSets.EntityData.YangName = "etag-sets"
    etagSets.EntityData.BundleName = "cisco_ios_xr"
    etagSets.EntityData.ParentYangName = "sets"
    etagSets.EntityData.SegmentPath = "etag-sets"
    etagSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    etagSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    etagSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    etagSets.EntityData.Children = types.NewOrderedMap()
    etagSets.EntityData.Children.Append("etag-set", types.YChild{"EtagSet", nil})
    for i := range etagSets.EtagSet {
        etagSets.EntityData.Children.Append(types.GetSegmentPath(etagSets.EtagSet[i]), types.YChild{"EtagSet", etagSets.EtagSet[i]})
    }
    etagSets.EntityData.Leafs = types.NewOrderedMap()

    etagSets.EntityData.YListKeys = []string {}

    return &(etagSets.EntityData)
}

// RoutingPolicy_Sets_EtagSets_EtagSet
// Information about an individual set
type RoutingPolicy_Sets_EtagSets_EtagSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Etag Set. The type is string. This attribute is mandatory.
    EtagSetAsText interface{}
}

func (etagSet *RoutingPolicy_Sets_EtagSets_EtagSet) GetEntityData() *types.CommonEntityData {
    etagSet.EntityData.YFilter = etagSet.YFilter
    etagSet.EntityData.YangName = "etag-set"
    etagSet.EntityData.BundleName = "cisco_ios_xr"
    etagSet.EntityData.ParentYangName = "etag-sets"
    etagSet.EntityData.SegmentPath = "etag-set" + types.AddKeyToken(etagSet.SetName, "set-name")
    etagSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    etagSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    etagSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    etagSet.EntityData.Children = types.NewOrderedMap()
    etagSet.EntityData.Leafs = types.NewOrderedMap()
    etagSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", etagSet.SetName})
    etagSet.EntityData.Leafs.Append("etag-set-as-text", types.YLeaf{"EtagSetAsText", etagSet.EtagSetAsText})

    etagSet.EntityData.YListKeys = []string {"SetName"}

    return &(etagSet.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRtSets
// Information about RT sets
type RoutingPolicy_Sets_ExtendedCommunityRtSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet.
    ExtendedCommunityRtSet []*RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet
}

func (extendedCommunityRtSets *RoutingPolicy_Sets_ExtendedCommunityRtSets) GetEntityData() *types.CommonEntityData {
    extendedCommunityRtSets.EntityData.YFilter = extendedCommunityRtSets.YFilter
    extendedCommunityRtSets.EntityData.YangName = "extended-community-rt-sets"
    extendedCommunityRtSets.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityRtSets.EntityData.ParentYangName = "sets"
    extendedCommunityRtSets.EntityData.SegmentPath = "extended-community-rt-sets"
    extendedCommunityRtSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityRtSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityRtSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityRtSets.EntityData.Children = types.NewOrderedMap()
    extendedCommunityRtSets.EntityData.Children.Append("extended-community-rt-set", types.YChild{"ExtendedCommunityRtSet", nil})
    for i := range extendedCommunityRtSets.ExtendedCommunityRtSet {
        extendedCommunityRtSets.EntityData.Children.Append(types.GetSegmentPath(extendedCommunityRtSets.ExtendedCommunityRtSet[i]), types.YChild{"ExtendedCommunityRtSet", extendedCommunityRtSets.ExtendedCommunityRtSet[i]})
    }
    extendedCommunityRtSets.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunityRtSets.EntityData.YListKeys = []string {}

    return &(extendedCommunityRtSets.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Extended Community RT Set. The type is string. This attribute is mandatory.
    RplExtendedCommunityRtSet interface{}
}

func (extendedCommunityRtSet *RoutingPolicy_Sets_ExtendedCommunityRtSets_ExtendedCommunityRtSet) GetEntityData() *types.CommonEntityData {
    extendedCommunityRtSet.EntityData.YFilter = extendedCommunityRtSet.YFilter
    extendedCommunityRtSet.EntityData.YangName = "extended-community-rt-set"
    extendedCommunityRtSet.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityRtSet.EntityData.ParentYangName = "extended-community-rt-sets"
    extendedCommunityRtSet.EntityData.SegmentPath = "extended-community-rt-set" + types.AddKeyToken(extendedCommunityRtSet.SetName, "set-name")
    extendedCommunityRtSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityRtSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityRtSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityRtSet.EntityData.Children = types.NewOrderedMap()
    extendedCommunityRtSet.EntityData.Leafs = types.NewOrderedMap()
    extendedCommunityRtSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", extendedCommunityRtSet.SetName})
    extendedCommunityRtSet.EntityData.Leafs.Append("rpl-extended-community-rt-set", types.YLeaf{"RplExtendedCommunityRtSet", extendedCommunityRtSet.RplExtendedCommunityRtSet})

    extendedCommunityRtSet.EntityData.YListKeys = []string {"SetName"}

    return &(extendedCommunityRtSet.EntityData)
}

// RoutingPolicy_Limits
// Limits for Routing Policy
type RoutingPolicy_Limits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of lines of policy configuration that may be configured in
    // total. The type is interface{} with range: 0..4294967295. The default value
    // is 131072.
    MaximumLinesOfPolicy interface{}

    // Maximum number of policies that may be configured. The type is interface{}
    // with range: 0..4294967295. The default value is 5000.
    MaximumNumberOfPolicies interface{}
}

func (limits *RoutingPolicy_Limits) GetEntityData() *types.CommonEntityData {
    limits.EntityData.YFilter = limits.YFilter
    limits.EntityData.YangName = "limits"
    limits.EntityData.BundleName = "cisco_ios_xr"
    limits.EntityData.ParentYangName = "routing-policy"
    limits.EntityData.SegmentPath = "limits"
    limits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    limits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    limits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    limits.EntityData.Children = types.NewOrderedMap()
    limits.EntityData.Leafs = types.NewOrderedMap()
    limits.EntityData.Leafs.Append("maximum-lines-of-policy", types.YLeaf{"MaximumLinesOfPolicy", limits.MaximumLinesOfPolicy})
    limits.EntityData.Leafs.Append("maximum-number-of-policies", types.YLeaf{"MaximumNumberOfPolicies", limits.MaximumNumberOfPolicies})

    limits.EntityData.YListKeys = []string {}

    return &(limits.EntityData)
}

