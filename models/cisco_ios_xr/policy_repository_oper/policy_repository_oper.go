// This module contains a collection of YANG definitions
// for Cisco IOS-XR policy-repository package operational data.
// 
// This module contains definitions
// for the following management objects:
//   routing-policy: Routing policy operational data
//   routing-policy-shadow: routing policy shadow
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package policy_repository_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package policy_repository_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-policy-repository-oper routing-policy}", reflect.TypeOf(RoutingPolicy{}))
    ydk.RegisterEntity("Cisco-IOS-XR-policy-repository-oper:routing-policy", reflect.TypeOf(RoutingPolicy{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-policy-repository-oper routing-policy-shadow}", reflect.TypeOf(RoutingPolicyShadow{}))
    ydk.RegisterEntity("Cisco-IOS-XR-policy-repository-oper:routing-policy-shadow", reflect.TypeOf(RoutingPolicyShadow{}))
}

// Group represents BGP Neighbor Group Type
type Group string

const (
    // Address Family Group
    Group_address_family_group Group = "address-family-group"

    // Session Group
    Group_session_group Group = "session-group"

    // Neighbor Group
    Group_neighbor_group Group = "neighbor-group"

    // Neighbor
    Group_neighbor Group = "neighbor"

    // Error Group
    Group_error_group Group = "error-group"
)

// AttachPointDirection represents Attach Point Direction
type AttachPointDirection string

const (
    // Attach Point Direction IN
    AttachPointDirection_in AttachPointDirection = "in"

    // Attach Point Direction OUT
    AttachPointDirection_out AttachPointDirection = "out"
)

// SubAddressFamily represents Sub Address Family
type SubAddressFamily string

const (
    // Unicast
    SubAddressFamily_unicast SubAddressFamily = "unicast"

    // Multicast
    SubAddressFamily_multicast SubAddressFamily = "multicast"

    // Label
    SubAddressFamily_label SubAddressFamily = "label"

    // Tunnel
    SubAddressFamily_tunnel SubAddressFamily = "tunnel"

    // VPN
    SubAddressFamily_vpn SubAddressFamily = "vpn"

    // MDT
    SubAddressFamily_mdt SubAddressFamily = "mdt"

    // VPLS
    SubAddressFamily_vpls SubAddressFamily = "vpls"

    // RTConstraint
    SubAddressFamily_rt_constraint SubAddressFamily = "rt-constraint"

    // MVPN
    SubAddressFamily_mvpn SubAddressFamily = "mvpn"

    // FLOW
    SubAddressFamily_flow SubAddressFamily = "flow"

    // VPN Multicast
    SubAddressFamily_vpn_mcast SubAddressFamily = "vpn-mcast"

    // EVPN
    SubAddressFamily_evpn SubAddressFamily = "evpn"

    // No SAFI
    SubAddressFamily_saf_none SubAddressFamily = "saf-none"

    // Unknown
    SubAddressFamily_saf_unknown SubAddressFamily = "saf-unknown"
)

// AddressFamily represents Address Family
type AddressFamily string

const (
    // IPv4 Address Family
    AddressFamily_ipv4 AddressFamily = "ipv4"

    // IPv6 Address Family
    AddressFamily_ipv6 AddressFamily = "ipv6"

    // L2VPN Address Family
    AddressFamily_l2vpn AddressFamily = "l2vpn"

    // LINKSTATE Address Family
    AddressFamily_ls AddressFamily = "ls"

    // No Address Family
    AddressFamily_af_none AddressFamily = "af-none"

    // Unknown Address Family
    AddressFamily_af_unknown AddressFamily = "af-unknown"
)

// ObjectStatus represents Whether an RPL object is used/referenced
type ObjectStatus string

const (
    // The object is in use
    ObjectStatus_active ObjectStatus = "active"

    // The object is referenced by another object, but
    // not used
    ObjectStatus_inactive ObjectStatus = "inactive"

    // The object is not used or referenced
    ObjectStatus_unused ObjectStatus = "unused"
)

// RoutingPolicy
// Routing policy operational data
type RoutingPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about configured limits and the current values.
    Limits RoutingPolicy_Limits

    // Information about configured route policies.
    Policies RoutingPolicy_Policies

    // Information about configured sets.
    Sets RoutingPolicy_Sets
}

func (routingPolicy *RoutingPolicy) GetEntityData() *types.CommonEntityData {
    routingPolicy.EntityData.YFilter = routingPolicy.YFilter
    routingPolicy.EntityData.YangName = "routing-policy"
    routingPolicy.EntityData.BundleName = "cisco_ios_xr"
    routingPolicy.EntityData.ParentYangName = "Cisco-IOS-XR-policy-repository-oper"
    routingPolicy.EntityData.SegmentPath = "Cisco-IOS-XR-policy-repository-oper:routing-policy"
    routingPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingPolicy.EntityData.Children = types.NewOrderedMap()
    routingPolicy.EntityData.Children.Append("limits", types.YChild{"Limits", &routingPolicy.Limits})
    routingPolicy.EntityData.Children.Append("policies", types.YChild{"Policies", &routingPolicy.Policies})
    routingPolicy.EntityData.Children.Append("sets", types.YChild{"Sets", &routingPolicy.Sets})
    routingPolicy.EntityData.Leafs = types.NewOrderedMap()

    routingPolicy.EntityData.YListKeys = []string {}

    return &(routingPolicy.EntityData)
}

// RoutingPolicy_Limits
// Information about configured limits and the
// current values
type RoutingPolicy_Limits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum lines of configuration allowable for all policies and sets. The
    // type is interface{} with range: 0..4294967295.
    MaximumLinesOfPolicy interface{}

    // Number of lines of configuration for policies/sets currently allowed. The
    // type is interface{} with range: 0..4294967295.
    CurrentLinesOfPolicyLimit interface{}

    // Current number of lines configured for all policies and sets. The type is
    // interface{} with range: 0..4294967295.
    CurrentLinesOfPolicyUsed interface{}

    // Maximum number of policies allowable. The type is interface{} with range:
    // 0..4294967295.
    MaximumNumberOfPolicies interface{}

    // Number of policies currently allowed. The type is interface{} with range:
    // 0..4294967295.
    CurrentNumberOfPoliciesLimit interface{}

    // Current number of policies configured. The type is interface{} with range:
    // 0..4294967295.
    CurrentNumberOfPoliciesUsed interface{}

    // The total compiled length of all policies. The type is interface{} with
    // range: 0..4294967295.
    CompiledPoliciesLength interface{}
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
    limits.EntityData.Leafs.Append("current-lines-of-policy-limit", types.YLeaf{"CurrentLinesOfPolicyLimit", limits.CurrentLinesOfPolicyLimit})
    limits.EntityData.Leafs.Append("current-lines-of-policy-used", types.YLeaf{"CurrentLinesOfPolicyUsed", limits.CurrentLinesOfPolicyUsed})
    limits.EntityData.Leafs.Append("maximum-number-of-policies", types.YLeaf{"MaximumNumberOfPolicies", limits.MaximumNumberOfPolicies})
    limits.EntityData.Leafs.Append("current-number-of-policies-limit", types.YLeaf{"CurrentNumberOfPoliciesLimit", limits.CurrentNumberOfPoliciesLimit})
    limits.EntityData.Leafs.Append("current-number-of-policies-used", types.YLeaf{"CurrentNumberOfPoliciesUsed", limits.CurrentNumberOfPoliciesUsed})
    limits.EntityData.Leafs.Append("compiled-policies-length", types.YLeaf{"CompiledPoliciesLength", limits.CompiledPoliciesLength})

    limits.EntityData.YListKeys = []string {}

    return &(limits.EntityData)
}

// RoutingPolicy_Policies
// Information about configured route policies
type RoutingPolicy_Policies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual policies.
    RoutePolicies RoutingPolicy_Policies_RoutePolicies

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Policies_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Policies_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Policies_Active
}

func (policies *RoutingPolicy_Policies) GetEntityData() *types.CommonEntityData {
    policies.EntityData.YFilter = policies.YFilter
    policies.EntityData.YangName = "policies"
    policies.EntityData.BundleName = "cisco_ios_xr"
    policies.EntityData.ParentYangName = "routing-policy"
    policies.EntityData.SegmentPath = "policies"
    policies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policies.EntityData.Children = types.NewOrderedMap()
    policies.EntityData.Children.Append("route-policies", types.YChild{"RoutePolicies", &policies.RoutePolicies})
    policies.EntityData.Children.Append("unused", types.YChild{"Unused", &policies.Unused})
    policies.EntityData.Children.Append("inactive", types.YChild{"Inactive", &policies.Inactive})
    policies.EntityData.Children.Append("active", types.YChild{"Active", &policies.Active})
    policies.EntityData.Leafs = types.NewOrderedMap()

    policies.EntityData.YListKeys = []string {}

    return &(policies.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies
// Information about individual policies
type RoutingPolicy_Policies_RoutePolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual policy. The type is slice of
    // RoutingPolicy_Policies_RoutePolicies_RoutePolicy.
    RoutePolicy []*RoutingPolicy_Policies_RoutePolicies_RoutePolicy
}

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetEntityData() *types.CommonEntityData {
    routePolicies.EntityData.YFilter = routePolicies.YFilter
    routePolicies.EntityData.YangName = "route-policies"
    routePolicies.EntityData.BundleName = "cisco_ios_xr"
    routePolicies.EntityData.ParentYangName = "policies"
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

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy
// Information about an individual policy
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Route policy name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    RoutePolicyName interface{}

    // Information about which policies and sets this policy uses.
    PolicyUses RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached
}

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetEntityData() *types.CommonEntityData {
    routePolicy.EntityData.YFilter = routePolicy.YFilter
    routePolicy.EntityData.YangName = "route-policy"
    routePolicy.EntityData.BundleName = "cisco_ios_xr"
    routePolicy.EntityData.ParentYangName = "route-policies"
    routePolicy.EntityData.SegmentPath = "route-policy" + types.AddKeyToken(routePolicy.RoutePolicyName, "route-policy-name")
    routePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routePolicy.EntityData.Children = types.NewOrderedMap()
    routePolicy.EntityData.Children.Append("policy-uses", types.YChild{"PolicyUses", &routePolicy.PolicyUses})
    routePolicy.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &routePolicy.UsedBy})
    routePolicy.EntityData.Children.Append("attached", types.YChild{"Attached", &routePolicy.Attached})
    routePolicy.EntityData.Leafs = types.NewOrderedMap()
    routePolicy.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", routePolicy.RoutePolicyName})

    routePolicy.EntityData.YListKeys = []string {"RoutePolicyName"}

    return &(routePolicy.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses
// Information about which policies and sets
// this policy uses
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policies that this policy uses directly.
    DirectlyUsedPolicies RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies

    // Sets used by this policy, or by policies that it uses.
    AllUsedSets RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets

    // Sets that this policy uses directly.
    DirectlyUsedSets RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets

    // Policies used by this policy, or by policies that it uses.
    AllUsedPolicies RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies
}

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetEntityData() *types.CommonEntityData {
    policyUses.EntityData.YFilter = policyUses.YFilter
    policyUses.EntityData.YangName = "policy-uses"
    policyUses.EntityData.BundleName = "cisco_ios_xr"
    policyUses.EntityData.ParentYangName = "route-policy"
    policyUses.EntityData.SegmentPath = "policy-uses"
    policyUses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyUses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyUses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyUses.EntityData.Children = types.NewOrderedMap()
    policyUses.EntityData.Children.Append("directly-used-policies", types.YChild{"DirectlyUsedPolicies", &policyUses.DirectlyUsedPolicies})
    policyUses.EntityData.Children.Append("all-used-sets", types.YChild{"AllUsedSets", &policyUses.AllUsedSets})
    policyUses.EntityData.Children.Append("directly-used-sets", types.YChild{"DirectlyUsedSets", &policyUses.DirectlyUsedSets})
    policyUses.EntityData.Children.Append("all-used-policies", types.YChild{"AllUsedPolicies", &policyUses.AllUsedPolicies})
    policyUses.EntityData.Leafs = types.NewOrderedMap()

    policyUses.EntityData.YListKeys = []string {}

    return &(policyUses.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies
// Policies that this policy uses directly
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetEntityData() *types.CommonEntityData {
    directlyUsedPolicies.EntityData.YFilter = directlyUsedPolicies.YFilter
    directlyUsedPolicies.EntityData.YangName = "directly-used-policies"
    directlyUsedPolicies.EntityData.BundleName = "cisco_ios_xr"
    directlyUsedPolicies.EntityData.ParentYangName = "policy-uses"
    directlyUsedPolicies.EntityData.SegmentPath = "directly-used-policies"
    directlyUsedPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    directlyUsedPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    directlyUsedPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    directlyUsedPolicies.EntityData.Children = types.NewOrderedMap()
    directlyUsedPolicies.EntityData.Leafs = types.NewOrderedMap()
    directlyUsedPolicies.EntityData.Leafs.Append("object", types.YLeaf{"Object", directlyUsedPolicies.Object})

    directlyUsedPolicies.EntityData.YListKeys = []string {}

    return &(directlyUsedPolicies.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets
// Sets used by this policy, or by policies
// that it uses
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of sets in several domains. The type is slice of
    // RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets.
    Sets []*RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets
}

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetEntityData() *types.CommonEntityData {
    allUsedSets.EntityData.YFilter = allUsedSets.YFilter
    allUsedSets.EntityData.YangName = "all-used-sets"
    allUsedSets.EntityData.BundleName = "cisco_ios_xr"
    allUsedSets.EntityData.ParentYangName = "policy-uses"
    allUsedSets.EntityData.SegmentPath = "all-used-sets"
    allUsedSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allUsedSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allUsedSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allUsedSets.EntityData.Children = types.NewOrderedMap()
    allUsedSets.EntityData.Children.Append("sets", types.YChild{"Sets", nil})
    for i := range allUsedSets.Sets {
        allUsedSets.EntityData.Children.Append(types.GetSegmentPath(allUsedSets.Sets[i]), types.YChild{"Sets", allUsedSets.Sets[i]})
    }
    allUsedSets.EntityData.Leafs = types.NewOrderedMap()

    allUsedSets.EntityData.YListKeys = []string {}

    return &(allUsedSets.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets
// List of sets in several domains
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain of sets. The type is string.
    SetDomain interface{}

    // Names of sets in this domain. The type is slice of string.
    SetName []interface{}
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "all-used-sets"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Leafs = types.NewOrderedMap()
    sets.EntityData.Leafs.Append("set-domain", types.YLeaf{"SetDomain", sets.SetDomain})
    sets.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", sets.SetName})

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets
// Sets that this policy uses directly
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of sets in several domains. The type is slice of
    // RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets.
    Sets []*RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets
}

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetEntityData() *types.CommonEntityData {
    directlyUsedSets.EntityData.YFilter = directlyUsedSets.YFilter
    directlyUsedSets.EntityData.YangName = "directly-used-sets"
    directlyUsedSets.EntityData.BundleName = "cisco_ios_xr"
    directlyUsedSets.EntityData.ParentYangName = "policy-uses"
    directlyUsedSets.EntityData.SegmentPath = "directly-used-sets"
    directlyUsedSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    directlyUsedSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    directlyUsedSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    directlyUsedSets.EntityData.Children = types.NewOrderedMap()
    directlyUsedSets.EntityData.Children.Append("sets", types.YChild{"Sets", nil})
    for i := range directlyUsedSets.Sets {
        directlyUsedSets.EntityData.Children.Append(types.GetSegmentPath(directlyUsedSets.Sets[i]), types.YChild{"Sets", directlyUsedSets.Sets[i]})
    }
    directlyUsedSets.EntityData.Leafs = types.NewOrderedMap()

    directlyUsedSets.EntityData.YListKeys = []string {}

    return &(directlyUsedSets.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets
// List of sets in several domains
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain of sets. The type is string.
    SetDomain interface{}

    // Names of sets in this domain. The type is slice of string.
    SetName []interface{}
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "directly-used-sets"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Leafs = types.NewOrderedMap()
    sets.EntityData.Leafs.Append("set-domain", types.YLeaf{"SetDomain", sets.SetDomain})
    sets.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", sets.SetName})

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies
// Policies used by this policy, or by policies
// that it uses
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetEntityData() *types.CommonEntityData {
    allUsedPolicies.EntityData.YFilter = allUsedPolicies.YFilter
    allUsedPolicies.EntityData.YangName = "all-used-policies"
    allUsedPolicies.EntityData.BundleName = "cisco_ios_xr"
    allUsedPolicies.EntityData.ParentYangName = "policy-uses"
    allUsedPolicies.EntityData.SegmentPath = "all-used-policies"
    allUsedPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allUsedPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allUsedPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allUsedPolicies.EntityData.Children = types.NewOrderedMap()
    allUsedPolicies.EntityData.Leafs = types.NewOrderedMap()
    allUsedPolicies.EntityData.Leafs.Append("object", types.YLeaf{"Object", allUsedPolicies.Object})

    allUsedPolicies.EntityData.YListKeys = []string {}

    return &(allUsedPolicies.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference.
    Reference []*RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "route-policy"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding.
    Binding []*RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding
}

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "route-policy"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding
// bindings list
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Policies_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Policies_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Policies_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "policies"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Policies_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Policies_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Policies_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "policies"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Policies_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Policies_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Policies_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "policies"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets
// Information about configured sets
type RoutingPolicy_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about Etag sets.
    Etag RoutingPolicy_Sets_Etag

    // Information about OSPF Area sets.
    OspfArea RoutingPolicy_Sets_OspfArea

    // Information about Extended Community Opaque sets.
    ExtendedCommunityOpaque RoutingPolicy_Sets_ExtendedCommunityOpaque

    // Information about Extended Community SegNH sets.
    ExtendedCommunitySegNh RoutingPolicy_Sets_ExtendedCommunitySegNh

    // Information about Extended Community SOO sets.
    ExtendedCommunitySoo RoutingPolicy_Sets_ExtendedCommunitySoo

    // Information about Tag sets.
    Tag RoutingPolicy_Sets_Tag

    // Information about AS Path sets.
    Prefix RoutingPolicy_Sets_Prefix

    // Information about Community sets.
    Community RoutingPolicy_Sets_Community

    // Information about AS Path sets.
    AsPath RoutingPolicy_Sets_AsPath

    // Information about Large Community sets.
    LargeCommunity RoutingPolicy_Sets_LargeCommunity

    // Information about Esi sets.
    Esi RoutingPolicy_Sets_Esi

    // Information about Extended Community Bandwidth sets.
    ExtendedCommunityBandwidth RoutingPolicy_Sets_ExtendedCommunityBandwidth

    // Information about Extended Community RT sets.
    ExtendedCommunityRt RoutingPolicy_Sets_ExtendedCommunityRt

    // Information about RD sets.
    Rd RoutingPolicy_Sets_Rd

    // Information about Mac sets.
    Mac RoutingPolicy_Sets_Mac

    // Information about Extended Community Cost sets.
    ExtendedCommunityCost RoutingPolicy_Sets_ExtendedCommunityCost
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
    sets.EntityData.Children.Append("etag", types.YChild{"Etag", &sets.Etag})
    sets.EntityData.Children.Append("ospf-area", types.YChild{"OspfArea", &sets.OspfArea})
    sets.EntityData.Children.Append("extended-community-opaque", types.YChild{"ExtendedCommunityOpaque", &sets.ExtendedCommunityOpaque})
    sets.EntityData.Children.Append("extended-community-seg-nh", types.YChild{"ExtendedCommunitySegNh", &sets.ExtendedCommunitySegNh})
    sets.EntityData.Children.Append("extended-community-soo", types.YChild{"ExtendedCommunitySoo", &sets.ExtendedCommunitySoo})
    sets.EntityData.Children.Append("tag", types.YChild{"Tag", &sets.Tag})
    sets.EntityData.Children.Append("prefix", types.YChild{"Prefix", &sets.Prefix})
    sets.EntityData.Children.Append("community", types.YChild{"Community", &sets.Community})
    sets.EntityData.Children.Append("as-path", types.YChild{"AsPath", &sets.AsPath})
    sets.EntityData.Children.Append("large-community", types.YChild{"LargeCommunity", &sets.LargeCommunity})
    sets.EntityData.Children.Append("esi", types.YChild{"Esi", &sets.Esi})
    sets.EntityData.Children.Append("extended-community-bandwidth", types.YChild{"ExtendedCommunityBandwidth", &sets.ExtendedCommunityBandwidth})
    sets.EntityData.Children.Append("extended-community-rt", types.YChild{"ExtendedCommunityRt", &sets.ExtendedCommunityRt})
    sets.EntityData.Children.Append("rd", types.YChild{"Rd", &sets.Rd})
    sets.EntityData.Children.Append("mac", types.YChild{"Mac", &sets.Mac})
    sets.EntityData.Children.Append("extended-community-cost", types.YChild{"ExtendedCommunityCost", &sets.ExtendedCommunityCost})
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_Etag
// Information about Etag sets
type RoutingPolicy_Sets_Etag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Etag_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_Etag_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_Etag_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_Etag_Active
}

func (etag *RoutingPolicy_Sets_Etag) GetEntityData() *types.CommonEntityData {
    etag.EntityData.YFilter = etag.YFilter
    etag.EntityData.YangName = "etag"
    etag.EntityData.BundleName = "cisco_ios_xr"
    etag.EntityData.ParentYangName = "sets"
    etag.EntityData.SegmentPath = "etag"
    etag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    etag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    etag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    etag.EntityData.Children = types.NewOrderedMap()
    etag.EntityData.Children.Append("sets", types.YChild{"Sets", &etag.Sets})
    etag.EntityData.Children.Append("unused", types.YChild{"Unused", &etag.Unused})
    etag.EntityData.Children.Append("inactive", types.YChild{"Inactive", &etag.Inactive})
    etag.EntityData.Children.Append("active", types.YChild{"Active", &etag.Active})
    etag.EntityData.Leafs = types.NewOrderedMap()

    etag.EntityData.YListKeys = []string {}

    return &(etag.EntityData)
}

// RoutingPolicy_Sets_Etag_Sets
// Information about individual sets
type RoutingPolicy_Sets_Etag_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Etag_Sets_Set.
    Set []*RoutingPolicy_Sets_Etag_Sets_Set
}

func (sets *RoutingPolicy_Sets_Etag_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "etag"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_Etag_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Etag_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Etag_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Etag_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_Etag_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Etag_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Etag_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Etag_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_Etag_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Etag_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Etag_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "etag"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_Etag_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Etag_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "etag"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_Etag_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Etag_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Etag_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "etag"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_OspfArea
// Information about OSPF Area sets
type RoutingPolicy_Sets_OspfArea struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_OspfArea_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_OspfArea_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_OspfArea_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_OspfArea_Active
}

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetEntityData() *types.CommonEntityData {
    ospfArea.EntityData.YFilter = ospfArea.YFilter
    ospfArea.EntityData.YangName = "ospf-area"
    ospfArea.EntityData.BundleName = "cisco_ios_xr"
    ospfArea.EntityData.ParentYangName = "sets"
    ospfArea.EntityData.SegmentPath = "ospf-area"
    ospfArea.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfArea.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfArea.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfArea.EntityData.Children = types.NewOrderedMap()
    ospfArea.EntityData.Children.Append("sets", types.YChild{"Sets", &ospfArea.Sets})
    ospfArea.EntityData.Children.Append("unused", types.YChild{"Unused", &ospfArea.Unused})
    ospfArea.EntityData.Children.Append("inactive", types.YChild{"Inactive", &ospfArea.Inactive})
    ospfArea.EntityData.Children.Append("active", types.YChild{"Active", &ospfArea.Active})
    ospfArea.EntityData.Leafs = types.NewOrderedMap()

    ospfArea.EntityData.YListKeys = []string {}

    return &(ospfArea.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Sets
// Information about individual sets
type RoutingPolicy_Sets_OspfArea_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_OspfArea_Sets_Set.
    Set []*RoutingPolicy_Sets_OspfArea_Sets_Set
}

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "ospf-area"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_OspfArea_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_OspfArea_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_OspfArea_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_OspfArea_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "ospf-area"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_OspfArea_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "ospf-area"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_OspfArea_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_OspfArea_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "ospf-area"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque
// Information about Extended Community Opaque
// sets
type RoutingPolicy_Sets_ExtendedCommunityOpaque struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_ExtendedCommunityOpaque_Active
}

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetEntityData() *types.CommonEntityData {
    extendedCommunityOpaque.EntityData.YFilter = extendedCommunityOpaque.YFilter
    extendedCommunityOpaque.EntityData.YangName = "extended-community-opaque"
    extendedCommunityOpaque.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityOpaque.EntityData.ParentYangName = "sets"
    extendedCommunityOpaque.EntityData.SegmentPath = "extended-community-opaque"
    extendedCommunityOpaque.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityOpaque.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityOpaque.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityOpaque.EntityData.Children = types.NewOrderedMap()
    extendedCommunityOpaque.EntityData.Children.Append("sets", types.YChild{"Sets", &extendedCommunityOpaque.Sets})
    extendedCommunityOpaque.EntityData.Children.Append("unused", types.YChild{"Unused", &extendedCommunityOpaque.Unused})
    extendedCommunityOpaque.EntityData.Children.Append("inactive", types.YChild{"Inactive", &extendedCommunityOpaque.Inactive})
    extendedCommunityOpaque.EntityData.Children.Append("active", types.YChild{"Active", &extendedCommunityOpaque.Active})
    extendedCommunityOpaque.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunityOpaque.EntityData.YListKeys = []string {}

    return &(extendedCommunityOpaque.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set.
    Set []*RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "extended-community-opaque"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "extended-community-opaque"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "extended-community-opaque"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "extended-community-opaque"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh
// Information about Extended Community SegNH sets
type RoutingPolicy_Sets_ExtendedCommunitySegNh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_ExtendedCommunitySegNh_Active
}

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetEntityData() *types.CommonEntityData {
    extendedCommunitySegNh.EntityData.YFilter = extendedCommunitySegNh.YFilter
    extendedCommunitySegNh.EntityData.YangName = "extended-community-seg-nh"
    extendedCommunitySegNh.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunitySegNh.EntityData.ParentYangName = "sets"
    extendedCommunitySegNh.EntityData.SegmentPath = "extended-community-seg-nh"
    extendedCommunitySegNh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunitySegNh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunitySegNh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunitySegNh.EntityData.Children = types.NewOrderedMap()
    extendedCommunitySegNh.EntityData.Children.Append("sets", types.YChild{"Sets", &extendedCommunitySegNh.Sets})
    extendedCommunitySegNh.EntityData.Children.Append("unused", types.YChild{"Unused", &extendedCommunitySegNh.Unused})
    extendedCommunitySegNh.EntityData.Children.Append("inactive", types.YChild{"Inactive", &extendedCommunitySegNh.Inactive})
    extendedCommunitySegNh.EntityData.Children.Append("active", types.YChild{"Active", &extendedCommunitySegNh.Active})
    extendedCommunitySegNh.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunitySegNh.EntityData.YListKeys = []string {}

    return &(extendedCommunitySegNh.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set.
    Set []*RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "extended-community-seg-nh"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "extended-community-seg-nh"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "extended-community-seg-nh"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "extended-community-seg-nh"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo
// Information about Extended Community SOO sets
type RoutingPolicy_Sets_ExtendedCommunitySoo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_ExtendedCommunitySoo_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_ExtendedCommunitySoo_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_ExtendedCommunitySoo_Active
}

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetEntityData() *types.CommonEntityData {
    extendedCommunitySoo.EntityData.YFilter = extendedCommunitySoo.YFilter
    extendedCommunitySoo.EntityData.YangName = "extended-community-soo"
    extendedCommunitySoo.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunitySoo.EntityData.ParentYangName = "sets"
    extendedCommunitySoo.EntityData.SegmentPath = "extended-community-soo"
    extendedCommunitySoo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunitySoo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunitySoo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunitySoo.EntityData.Children = types.NewOrderedMap()
    extendedCommunitySoo.EntityData.Children.Append("sets", types.YChild{"Sets", &extendedCommunitySoo.Sets})
    extendedCommunitySoo.EntityData.Children.Append("unused", types.YChild{"Unused", &extendedCommunitySoo.Unused})
    extendedCommunitySoo.EntityData.Children.Append("inactive", types.YChild{"Inactive", &extendedCommunitySoo.Inactive})
    extendedCommunitySoo.EntityData.Children.Append("active", types.YChild{"Active", &extendedCommunitySoo.Active})
    extendedCommunitySoo.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunitySoo.EntityData.YListKeys = []string {}

    return &(extendedCommunitySoo.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set.
    Set []*RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "extended-community-soo"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_ExtendedCommunitySoo_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "extended-community-soo"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "extended-community-soo"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_ExtendedCommunitySoo_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "extended-community-soo"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_Tag
// Information about Tag sets
type RoutingPolicy_Sets_Tag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Tag_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_Tag_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_Tag_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_Tag_Active
}

func (tag *RoutingPolicy_Sets_Tag) GetEntityData() *types.CommonEntityData {
    tag.EntityData.YFilter = tag.YFilter
    tag.EntityData.YangName = "tag"
    tag.EntityData.BundleName = "cisco_ios_xr"
    tag.EntityData.ParentYangName = "sets"
    tag.EntityData.SegmentPath = "tag"
    tag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tag.EntityData.Children = types.NewOrderedMap()
    tag.EntityData.Children.Append("sets", types.YChild{"Sets", &tag.Sets})
    tag.EntityData.Children.Append("unused", types.YChild{"Unused", &tag.Unused})
    tag.EntityData.Children.Append("inactive", types.YChild{"Inactive", &tag.Inactive})
    tag.EntityData.Children.Append("active", types.YChild{"Active", &tag.Active})
    tag.EntityData.Leafs = types.NewOrderedMap()

    tag.EntityData.YListKeys = []string {}

    return &(tag.EntityData)
}

// RoutingPolicy_Sets_Tag_Sets
// Information about individual sets
type RoutingPolicy_Sets_Tag_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Tag_Sets_Set.
    Set []*RoutingPolicy_Sets_Tag_Sets_Set
}

func (sets *RoutingPolicy_Sets_Tag_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "tag"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_Tag_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Tag_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Tag_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Tag_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_Tag_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Tag_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Tag_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Tag_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_Tag_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Tag_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Tag_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "tag"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_Tag_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Tag_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "tag"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_Tag_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Tag_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Tag_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "tag"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_Prefix
// Information about AS Path sets
type RoutingPolicy_Sets_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Prefix_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_Prefix_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_Prefix_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_Prefix_Active
}

func (prefix *RoutingPolicy_Sets_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "sets"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Children.Append("sets", types.YChild{"Sets", &prefix.Sets})
    prefix.EntityData.Children.Append("unused", types.YChild{"Unused", &prefix.Unused})
    prefix.EntityData.Children.Append("inactive", types.YChild{"Inactive", &prefix.Inactive})
    prefix.EntityData.Children.Append("active", types.YChild{"Active", &prefix.Active})
    prefix.EntityData.Leafs = types.NewOrderedMap()

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// RoutingPolicy_Sets_Prefix_Sets
// Information about individual sets
type RoutingPolicy_Sets_Prefix_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Prefix_Sets_Set.
    Set []*RoutingPolicy_Sets_Prefix_Sets_Set
}

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "prefix"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_Prefix_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Prefix_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Prefix_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Prefix_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Prefix_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_Prefix_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Prefix_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "prefix"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_Prefix_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Prefix_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "prefix"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_Prefix_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Prefix_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Prefix_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "prefix"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_Community
// Information about Community sets
type RoutingPolicy_Sets_Community struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Community_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_Community_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_Community_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_Community_Active
}

func (community *RoutingPolicy_Sets_Community) GetEntityData() *types.CommonEntityData {
    community.EntityData.YFilter = community.YFilter
    community.EntityData.YangName = "community"
    community.EntityData.BundleName = "cisco_ios_xr"
    community.EntityData.ParentYangName = "sets"
    community.EntityData.SegmentPath = "community"
    community.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    community.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    community.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    community.EntityData.Children = types.NewOrderedMap()
    community.EntityData.Children.Append("sets", types.YChild{"Sets", &community.Sets})
    community.EntityData.Children.Append("unused", types.YChild{"Unused", &community.Unused})
    community.EntityData.Children.Append("inactive", types.YChild{"Inactive", &community.Inactive})
    community.EntityData.Children.Append("active", types.YChild{"Active", &community.Active})
    community.EntityData.Leafs = types.NewOrderedMap()

    community.EntityData.YListKeys = []string {}

    return &(community.EntityData)
}

// RoutingPolicy_Sets_Community_Sets
// Information about individual sets
type RoutingPolicy_Sets_Community_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Community_Sets_Set.
    Set []*RoutingPolicy_Sets_Community_Sets_Set
}

func (sets *RoutingPolicy_Sets_Community_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "community"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_Community_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Community_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Community_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Community_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_Community_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Community_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Community_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Community_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_Community_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Community_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Community_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "community"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_Community_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Community_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "community"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_Community_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Community_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Community_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "community"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_AsPath
// Information about AS Path sets
type RoutingPolicy_Sets_AsPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_AsPath_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_AsPath_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_AsPath_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_AsPath_Active
}

func (asPath *RoutingPolicy_Sets_AsPath) GetEntityData() *types.CommonEntityData {
    asPath.EntityData.YFilter = asPath.YFilter
    asPath.EntityData.YangName = "as-path"
    asPath.EntityData.BundleName = "cisco_ios_xr"
    asPath.EntityData.ParentYangName = "sets"
    asPath.EntityData.SegmentPath = "as-path"
    asPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asPath.EntityData.Children = types.NewOrderedMap()
    asPath.EntityData.Children.Append("sets", types.YChild{"Sets", &asPath.Sets})
    asPath.EntityData.Children.Append("unused", types.YChild{"Unused", &asPath.Unused})
    asPath.EntityData.Children.Append("inactive", types.YChild{"Inactive", &asPath.Inactive})
    asPath.EntityData.Children.Append("active", types.YChild{"Active", &asPath.Active})
    asPath.EntityData.Leafs = types.NewOrderedMap()

    asPath.EntityData.YListKeys = []string {}

    return &(asPath.EntityData)
}

// RoutingPolicy_Sets_AsPath_Sets
// Information about individual sets
type RoutingPolicy_Sets_AsPath_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_AsPath_Sets_Set.
    Set []*RoutingPolicy_Sets_AsPath_Sets_Set
}

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "as-path"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_AsPath_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_AsPath_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_AsPath_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_AsPath_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_AsPath_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_AsPath_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_AsPath_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "as-path"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_AsPath_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_AsPath_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "as-path"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_AsPath_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_AsPath_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_AsPath_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "as-path"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity
// Information about Large Community sets
type RoutingPolicy_Sets_LargeCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_LargeCommunity_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_LargeCommunity_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_LargeCommunity_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_LargeCommunity_Active
}

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetEntityData() *types.CommonEntityData {
    largeCommunity.EntityData.YFilter = largeCommunity.YFilter
    largeCommunity.EntityData.YangName = "large-community"
    largeCommunity.EntityData.BundleName = "cisco_ios_xr"
    largeCommunity.EntityData.ParentYangName = "sets"
    largeCommunity.EntityData.SegmentPath = "large-community"
    largeCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    largeCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    largeCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    largeCommunity.EntityData.Children = types.NewOrderedMap()
    largeCommunity.EntityData.Children.Append("sets", types.YChild{"Sets", &largeCommunity.Sets})
    largeCommunity.EntityData.Children.Append("unused", types.YChild{"Unused", &largeCommunity.Unused})
    largeCommunity.EntityData.Children.Append("inactive", types.YChild{"Inactive", &largeCommunity.Inactive})
    largeCommunity.EntityData.Children.Append("active", types.YChild{"Active", &largeCommunity.Active})
    largeCommunity.EntityData.Leafs = types.NewOrderedMap()

    largeCommunity.EntityData.YListKeys = []string {}

    return &(largeCommunity.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Sets
// Information about individual sets
type RoutingPolicy_Sets_LargeCommunity_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_LargeCommunity_Sets_Set.
    Set []*RoutingPolicy_Sets_LargeCommunity_Sets_Set
}

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "large-community"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_LargeCommunity_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_LargeCommunity_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "large-community"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_LargeCommunity_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "large-community"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_LargeCommunity_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "large-community"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_Esi
// Information about Esi sets
type RoutingPolicy_Sets_Esi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Esi_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_Esi_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_Esi_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_Esi_Active
}

func (esi *RoutingPolicy_Sets_Esi) GetEntityData() *types.CommonEntityData {
    esi.EntityData.YFilter = esi.YFilter
    esi.EntityData.YangName = "esi"
    esi.EntityData.BundleName = "cisco_ios_xr"
    esi.EntityData.ParentYangName = "sets"
    esi.EntityData.SegmentPath = "esi"
    esi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esi.EntityData.Children = types.NewOrderedMap()
    esi.EntityData.Children.Append("sets", types.YChild{"Sets", &esi.Sets})
    esi.EntityData.Children.Append("unused", types.YChild{"Unused", &esi.Unused})
    esi.EntityData.Children.Append("inactive", types.YChild{"Inactive", &esi.Inactive})
    esi.EntityData.Children.Append("active", types.YChild{"Active", &esi.Active})
    esi.EntityData.Leafs = types.NewOrderedMap()

    esi.EntityData.YListKeys = []string {}

    return &(esi.EntityData)
}

// RoutingPolicy_Sets_Esi_Sets
// Information about individual sets
type RoutingPolicy_Sets_Esi_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Esi_Sets_Set.
    Set []*RoutingPolicy_Sets_Esi_Sets_Set
}

func (sets *RoutingPolicy_Sets_Esi_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "esi"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_Esi_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Esi_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Esi_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Esi_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_Esi_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Esi_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Esi_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Esi_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_Esi_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Esi_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Esi_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "esi"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_Esi_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Esi_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "esi"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_Esi_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Esi_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Esi_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "esi"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth
// Information about Extended Community Bandwidth
// sets
type RoutingPolicy_Sets_ExtendedCommunityBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive
}

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetEntityData() *types.CommonEntityData {
    extendedCommunityBandwidth.EntityData.YFilter = extendedCommunityBandwidth.YFilter
    extendedCommunityBandwidth.EntityData.YangName = "extended-community-bandwidth"
    extendedCommunityBandwidth.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityBandwidth.EntityData.ParentYangName = "sets"
    extendedCommunityBandwidth.EntityData.SegmentPath = "extended-community-bandwidth"
    extendedCommunityBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityBandwidth.EntityData.Children = types.NewOrderedMap()
    extendedCommunityBandwidth.EntityData.Children.Append("sets", types.YChild{"Sets", &extendedCommunityBandwidth.Sets})
    extendedCommunityBandwidth.EntityData.Children.Append("unused", types.YChild{"Unused", &extendedCommunityBandwidth.Unused})
    extendedCommunityBandwidth.EntityData.Children.Append("inactive", types.YChild{"Inactive", &extendedCommunityBandwidth.Inactive})
    extendedCommunityBandwidth.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunityBandwidth.EntityData.YListKeys = []string {}

    return &(extendedCommunityBandwidth.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set.
    Set []*RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "extended-community-bandwidth"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "extended-community-bandwidth"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "extended-community-bandwidth"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt
// Information about Extended Community RT sets
type RoutingPolicy_Sets_ExtendedCommunityRt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_ExtendedCommunityRt_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_ExtendedCommunityRt_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_ExtendedCommunityRt_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_ExtendedCommunityRt_Active
}

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetEntityData() *types.CommonEntityData {
    extendedCommunityRt.EntityData.YFilter = extendedCommunityRt.YFilter
    extendedCommunityRt.EntityData.YangName = "extended-community-rt"
    extendedCommunityRt.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityRt.EntityData.ParentYangName = "sets"
    extendedCommunityRt.EntityData.SegmentPath = "extended-community-rt"
    extendedCommunityRt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityRt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityRt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityRt.EntityData.Children = types.NewOrderedMap()
    extendedCommunityRt.EntityData.Children.Append("sets", types.YChild{"Sets", &extendedCommunityRt.Sets})
    extendedCommunityRt.EntityData.Children.Append("unused", types.YChild{"Unused", &extendedCommunityRt.Unused})
    extendedCommunityRt.EntityData.Children.Append("inactive", types.YChild{"Inactive", &extendedCommunityRt.Inactive})
    extendedCommunityRt.EntityData.Children.Append("active", types.YChild{"Active", &extendedCommunityRt.Active})
    extendedCommunityRt.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunityRt.EntityData.YListKeys = []string {}

    return &(extendedCommunityRt.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set.
    Set []*RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "extended-community-rt"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_ExtendedCommunityRt_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "extended-community-rt"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_ExtendedCommunityRt_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "extended-community-rt"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_ExtendedCommunityRt_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "extended-community-rt"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_Rd
// Information about RD sets
type RoutingPolicy_Sets_Rd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Rd_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_Rd_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_Rd_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_Rd_Active
}

func (rd *RoutingPolicy_Sets_Rd) GetEntityData() *types.CommonEntityData {
    rd.EntityData.YFilter = rd.YFilter
    rd.EntityData.YangName = "rd"
    rd.EntityData.BundleName = "cisco_ios_xr"
    rd.EntityData.ParentYangName = "sets"
    rd.EntityData.SegmentPath = "rd"
    rd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rd.EntityData.Children = types.NewOrderedMap()
    rd.EntityData.Children.Append("sets", types.YChild{"Sets", &rd.Sets})
    rd.EntityData.Children.Append("unused", types.YChild{"Unused", &rd.Unused})
    rd.EntityData.Children.Append("inactive", types.YChild{"Inactive", &rd.Inactive})
    rd.EntityData.Children.Append("active", types.YChild{"Active", &rd.Active})
    rd.EntityData.Leafs = types.NewOrderedMap()

    rd.EntityData.YListKeys = []string {}

    return &(rd.EntityData)
}

// RoutingPolicy_Sets_Rd_Sets
// Information about individual sets
type RoutingPolicy_Sets_Rd_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Rd_Sets_Set.
    Set []*RoutingPolicy_Sets_Rd_Sets_Set
}

func (sets *RoutingPolicy_Sets_Rd_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "rd"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_Rd_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Rd_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Rd_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Rd_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_Rd_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Rd_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Rd_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Rd_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_Rd_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Rd_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Rd_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "rd"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_Rd_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Rd_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "rd"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_Rd_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Rd_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Rd_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "rd"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_Mac
// Information about Mac sets
type RoutingPolicy_Sets_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Mac_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_Mac_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_Mac_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_Mac_Active
}

func (mac *RoutingPolicy_Sets_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xr"
    mac.EntityData.ParentYangName = "sets"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Children.Append("sets", types.YChild{"Sets", &mac.Sets})
    mac.EntityData.Children.Append("unused", types.YChild{"Unused", &mac.Unused})
    mac.EntityData.Children.Append("inactive", types.YChild{"Inactive", &mac.Inactive})
    mac.EntityData.Children.Append("active", types.YChild{"Active", &mac.Active})
    mac.EntityData.Leafs = types.NewOrderedMap()

    mac.EntityData.YListKeys = []string {}

    return &(mac.EntityData)
}

// RoutingPolicy_Sets_Mac_Sets
// Information about individual sets
type RoutingPolicy_Sets_Mac_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Mac_Sets_Set.
    Set []*RoutingPolicy_Sets_Mac_Sets_Set
}

func (sets *RoutingPolicy_Sets_Mac_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "mac"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_Mac_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Mac_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Mac_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Mac_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_Mac_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Mac_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Mac_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Mac_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_Mac_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Mac_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Mac_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "mac"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_Mac_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Mac_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "mac"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_Mac_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Mac_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Mac_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "mac"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost
// Information about Extended Community Cost sets
type RoutingPolicy_Sets_ExtendedCommunityCost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_ExtendedCommunityCost_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_ExtendedCommunityCost_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_ExtendedCommunityCost_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicy_Sets_ExtendedCommunityCost_Active
}

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetEntityData() *types.CommonEntityData {
    extendedCommunityCost.EntityData.YFilter = extendedCommunityCost.YFilter
    extendedCommunityCost.EntityData.YangName = "extended-community-cost"
    extendedCommunityCost.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityCost.EntityData.ParentYangName = "sets"
    extendedCommunityCost.EntityData.SegmentPath = "extended-community-cost"
    extendedCommunityCost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityCost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityCost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityCost.EntityData.Children = types.NewOrderedMap()
    extendedCommunityCost.EntityData.Children.Append("sets", types.YChild{"Sets", &extendedCommunityCost.Sets})
    extendedCommunityCost.EntityData.Children.Append("unused", types.YChild{"Unused", &extendedCommunityCost.Unused})
    extendedCommunityCost.EntityData.Children.Append("inactive", types.YChild{"Inactive", &extendedCommunityCost.Inactive})
    extendedCommunityCost.EntityData.Children.Append("active", types.YChild{"Active", &extendedCommunityCost.Active})
    extendedCommunityCost.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunityCost.EntityData.YListKeys = []string {}

    return &(extendedCommunityCost.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set.
    Set []*RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "extended-community-cost"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_ExtendedCommunityCost_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "extended-community-cost"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_ExtendedCommunityCost_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "extended-community-cost"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_ExtendedCommunityCost_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "extended-community-cost"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow
// routing policy shadow
type RoutingPolicyShadow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about configured limits and the current values.
    Limits RoutingPolicyShadow_Limits

    // Information about configured route policies.
    Policies RoutingPolicyShadow_Policies

    // Information about configured sets.
    Sets RoutingPolicyShadow_Sets
}

func (routingPolicyShadow *RoutingPolicyShadow) GetEntityData() *types.CommonEntityData {
    routingPolicyShadow.EntityData.YFilter = routingPolicyShadow.YFilter
    routingPolicyShadow.EntityData.YangName = "routing-policy-shadow"
    routingPolicyShadow.EntityData.BundleName = "cisco_ios_xr"
    routingPolicyShadow.EntityData.ParentYangName = "Cisco-IOS-XR-policy-repository-oper"
    routingPolicyShadow.EntityData.SegmentPath = "Cisco-IOS-XR-policy-repository-oper:routing-policy-shadow"
    routingPolicyShadow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routingPolicyShadow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routingPolicyShadow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routingPolicyShadow.EntityData.Children = types.NewOrderedMap()
    routingPolicyShadow.EntityData.Children.Append("limits", types.YChild{"Limits", &routingPolicyShadow.Limits})
    routingPolicyShadow.EntityData.Children.Append("policies", types.YChild{"Policies", &routingPolicyShadow.Policies})
    routingPolicyShadow.EntityData.Children.Append("sets", types.YChild{"Sets", &routingPolicyShadow.Sets})
    routingPolicyShadow.EntityData.Leafs = types.NewOrderedMap()

    routingPolicyShadow.EntityData.YListKeys = []string {}

    return &(routingPolicyShadow.EntityData)
}

// RoutingPolicyShadow_Limits
// Information about configured limits and the
// current values
type RoutingPolicyShadow_Limits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum lines of configuration allowable for all policies and sets. The
    // type is interface{} with range: 0..4294967295.
    MaximumLinesOfPolicy interface{}

    // Number of lines of configuration for policies/sets currently allowed. The
    // type is interface{} with range: 0..4294967295.
    CurrentLinesOfPolicyLimit interface{}

    // Current number of lines configured for all policies and sets. The type is
    // interface{} with range: 0..4294967295.
    CurrentLinesOfPolicyUsed interface{}

    // Maximum number of policies allowable. The type is interface{} with range:
    // 0..4294967295.
    MaximumNumberOfPolicies interface{}

    // Number of policies currently allowed. The type is interface{} with range:
    // 0..4294967295.
    CurrentNumberOfPoliciesLimit interface{}

    // Current number of policies configured. The type is interface{} with range:
    // 0..4294967295.
    CurrentNumberOfPoliciesUsed interface{}

    // The total compiled length of all policies. The type is interface{} with
    // range: 0..4294967295.
    CompiledPoliciesLength interface{}
}

func (limits *RoutingPolicyShadow_Limits) GetEntityData() *types.CommonEntityData {
    limits.EntityData.YFilter = limits.YFilter
    limits.EntityData.YangName = "limits"
    limits.EntityData.BundleName = "cisco_ios_xr"
    limits.EntityData.ParentYangName = "routing-policy-shadow"
    limits.EntityData.SegmentPath = "limits"
    limits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    limits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    limits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    limits.EntityData.Children = types.NewOrderedMap()
    limits.EntityData.Leafs = types.NewOrderedMap()
    limits.EntityData.Leafs.Append("maximum-lines-of-policy", types.YLeaf{"MaximumLinesOfPolicy", limits.MaximumLinesOfPolicy})
    limits.EntityData.Leafs.Append("current-lines-of-policy-limit", types.YLeaf{"CurrentLinesOfPolicyLimit", limits.CurrentLinesOfPolicyLimit})
    limits.EntityData.Leafs.Append("current-lines-of-policy-used", types.YLeaf{"CurrentLinesOfPolicyUsed", limits.CurrentLinesOfPolicyUsed})
    limits.EntityData.Leafs.Append("maximum-number-of-policies", types.YLeaf{"MaximumNumberOfPolicies", limits.MaximumNumberOfPolicies})
    limits.EntityData.Leafs.Append("current-number-of-policies-limit", types.YLeaf{"CurrentNumberOfPoliciesLimit", limits.CurrentNumberOfPoliciesLimit})
    limits.EntityData.Leafs.Append("current-number-of-policies-used", types.YLeaf{"CurrentNumberOfPoliciesUsed", limits.CurrentNumberOfPoliciesUsed})
    limits.EntityData.Leafs.Append("compiled-policies-length", types.YLeaf{"CompiledPoliciesLength", limits.CompiledPoliciesLength})

    limits.EntityData.YListKeys = []string {}

    return &(limits.EntityData)
}

// RoutingPolicyShadow_Policies
// Information about configured route policies
type RoutingPolicyShadow_Policies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual policies.
    RoutePolicies RoutingPolicyShadow_Policies_RoutePolicies

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Policies_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Policies_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Policies_Active
}

func (policies *RoutingPolicyShadow_Policies) GetEntityData() *types.CommonEntityData {
    policies.EntityData.YFilter = policies.YFilter
    policies.EntityData.YangName = "policies"
    policies.EntityData.BundleName = "cisco_ios_xr"
    policies.EntityData.ParentYangName = "routing-policy-shadow"
    policies.EntityData.SegmentPath = "policies"
    policies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policies.EntityData.Children = types.NewOrderedMap()
    policies.EntityData.Children.Append("route-policies", types.YChild{"RoutePolicies", &policies.RoutePolicies})
    policies.EntityData.Children.Append("unused", types.YChild{"Unused", &policies.Unused})
    policies.EntityData.Children.Append("inactive", types.YChild{"Inactive", &policies.Inactive})
    policies.EntityData.Children.Append("active", types.YChild{"Active", &policies.Active})
    policies.EntityData.Leafs = types.NewOrderedMap()

    policies.EntityData.YListKeys = []string {}

    return &(policies.EntityData)
}

// RoutingPolicyShadow_Policies_RoutePolicies
// Information about individual policies
type RoutingPolicyShadow_Policies_RoutePolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual policy. The type is slice of
    // RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy.
    RoutePolicy []*RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy
}

func (routePolicies *RoutingPolicyShadow_Policies_RoutePolicies) GetEntityData() *types.CommonEntityData {
    routePolicies.EntityData.YFilter = routePolicies.YFilter
    routePolicies.EntityData.YangName = "route-policies"
    routePolicies.EntityData.BundleName = "cisco_ios_xr"
    routePolicies.EntityData.ParentYangName = "policies"
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

// RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy
// Information about an individual policy
type RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Route policy name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    RoutePolicyName interface{}

    // Information about which policies and sets this policy uses.
    PolicyUses RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_Attached
}

func (routePolicy *RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy) GetEntityData() *types.CommonEntityData {
    routePolicy.EntityData.YFilter = routePolicy.YFilter
    routePolicy.EntityData.YangName = "route-policy"
    routePolicy.EntityData.BundleName = "cisco_ios_xr"
    routePolicy.EntityData.ParentYangName = "route-policies"
    routePolicy.EntityData.SegmentPath = "route-policy" + types.AddKeyToken(routePolicy.RoutePolicyName, "route-policy-name")
    routePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routePolicy.EntityData.Children = types.NewOrderedMap()
    routePolicy.EntityData.Children.Append("policy-uses", types.YChild{"PolicyUses", &routePolicy.PolicyUses})
    routePolicy.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &routePolicy.UsedBy})
    routePolicy.EntityData.Children.Append("attached", types.YChild{"Attached", &routePolicy.Attached})
    routePolicy.EntityData.Leafs = types.NewOrderedMap()
    routePolicy.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", routePolicy.RoutePolicyName})

    routePolicy.EntityData.YListKeys = []string {"RoutePolicyName"}

    return &(routePolicy.EntityData)
}

// RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses
// Information about which policies and sets
// this policy uses
type RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policies that this policy uses directly.
    DirectlyUsedPolicies RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies

    // Sets used by this policy, or by policies that it uses.
    AllUsedSets RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets

    // Sets that this policy uses directly.
    DirectlyUsedSets RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets

    // Policies used by this policy, or by policies that it uses.
    AllUsedPolicies RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies
}

func (policyUses *RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetEntityData() *types.CommonEntityData {
    policyUses.EntityData.YFilter = policyUses.YFilter
    policyUses.EntityData.YangName = "policy-uses"
    policyUses.EntityData.BundleName = "cisco_ios_xr"
    policyUses.EntityData.ParentYangName = "route-policy"
    policyUses.EntityData.SegmentPath = "policy-uses"
    policyUses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyUses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyUses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyUses.EntityData.Children = types.NewOrderedMap()
    policyUses.EntityData.Children.Append("directly-used-policies", types.YChild{"DirectlyUsedPolicies", &policyUses.DirectlyUsedPolicies})
    policyUses.EntityData.Children.Append("all-used-sets", types.YChild{"AllUsedSets", &policyUses.AllUsedSets})
    policyUses.EntityData.Children.Append("directly-used-sets", types.YChild{"DirectlyUsedSets", &policyUses.DirectlyUsedSets})
    policyUses.EntityData.Children.Append("all-used-policies", types.YChild{"AllUsedPolicies", &policyUses.AllUsedPolicies})
    policyUses.EntityData.Leafs = types.NewOrderedMap()

    policyUses.EntityData.YListKeys = []string {}

    return &(policyUses.EntityData)
}

// RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies
// Policies that this policy uses directly
type RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (directlyUsedPolicies *RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetEntityData() *types.CommonEntityData {
    directlyUsedPolicies.EntityData.YFilter = directlyUsedPolicies.YFilter
    directlyUsedPolicies.EntityData.YangName = "directly-used-policies"
    directlyUsedPolicies.EntityData.BundleName = "cisco_ios_xr"
    directlyUsedPolicies.EntityData.ParentYangName = "policy-uses"
    directlyUsedPolicies.EntityData.SegmentPath = "directly-used-policies"
    directlyUsedPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    directlyUsedPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    directlyUsedPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    directlyUsedPolicies.EntityData.Children = types.NewOrderedMap()
    directlyUsedPolicies.EntityData.Leafs = types.NewOrderedMap()
    directlyUsedPolicies.EntityData.Leafs.Append("object", types.YLeaf{"Object", directlyUsedPolicies.Object})

    directlyUsedPolicies.EntityData.YListKeys = []string {}

    return &(directlyUsedPolicies.EntityData)
}

// RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets
// Sets used by this policy, or by policies
// that it uses
type RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of sets in several domains. The type is slice of
    // RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets.
    Sets []*RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets
}

func (allUsedSets *RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetEntityData() *types.CommonEntityData {
    allUsedSets.EntityData.YFilter = allUsedSets.YFilter
    allUsedSets.EntityData.YangName = "all-used-sets"
    allUsedSets.EntityData.BundleName = "cisco_ios_xr"
    allUsedSets.EntityData.ParentYangName = "policy-uses"
    allUsedSets.EntityData.SegmentPath = "all-used-sets"
    allUsedSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allUsedSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allUsedSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allUsedSets.EntityData.Children = types.NewOrderedMap()
    allUsedSets.EntityData.Children.Append("sets", types.YChild{"Sets", nil})
    for i := range allUsedSets.Sets {
        allUsedSets.EntityData.Children.Append(types.GetSegmentPath(allUsedSets.Sets[i]), types.YChild{"Sets", allUsedSets.Sets[i]})
    }
    allUsedSets.EntityData.Leafs = types.NewOrderedMap()

    allUsedSets.EntityData.YListKeys = []string {}

    return &(allUsedSets.EntityData)
}

// RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets
// List of sets in several domains
type RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain of sets. The type is string.
    SetDomain interface{}

    // Names of sets in this domain. The type is slice of string.
    SetName []interface{}
}

func (sets *RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "all-used-sets"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Leafs = types.NewOrderedMap()
    sets.EntityData.Leafs.Append("set-domain", types.YLeaf{"SetDomain", sets.SetDomain})
    sets.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", sets.SetName})

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets
// Sets that this policy uses directly
type RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of sets in several domains. The type is slice of
    // RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets.
    Sets []*RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets
}

func (directlyUsedSets *RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetEntityData() *types.CommonEntityData {
    directlyUsedSets.EntityData.YFilter = directlyUsedSets.YFilter
    directlyUsedSets.EntityData.YangName = "directly-used-sets"
    directlyUsedSets.EntityData.BundleName = "cisco_ios_xr"
    directlyUsedSets.EntityData.ParentYangName = "policy-uses"
    directlyUsedSets.EntityData.SegmentPath = "directly-used-sets"
    directlyUsedSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    directlyUsedSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    directlyUsedSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    directlyUsedSets.EntityData.Children = types.NewOrderedMap()
    directlyUsedSets.EntityData.Children.Append("sets", types.YChild{"Sets", nil})
    for i := range directlyUsedSets.Sets {
        directlyUsedSets.EntityData.Children.Append(types.GetSegmentPath(directlyUsedSets.Sets[i]), types.YChild{"Sets", directlyUsedSets.Sets[i]})
    }
    directlyUsedSets.EntityData.Leafs = types.NewOrderedMap()

    directlyUsedSets.EntityData.YListKeys = []string {}

    return &(directlyUsedSets.EntityData)
}

// RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets
// List of sets in several domains
type RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain of sets. The type is string.
    SetDomain interface{}

    // Names of sets in this domain. The type is slice of string.
    SetName []interface{}
}

func (sets *RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "directly-used-sets"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Leafs = types.NewOrderedMap()
    sets.EntityData.Leafs.Append("set-domain", types.YLeaf{"SetDomain", sets.SetDomain})
    sets.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", sets.SetName})

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies
// Policies used by this policy, or by policies
// that it uses
type RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (allUsedPolicies *RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetEntityData() *types.CommonEntityData {
    allUsedPolicies.EntityData.YFilter = allUsedPolicies.YFilter
    allUsedPolicies.EntityData.YangName = "all-used-policies"
    allUsedPolicies.EntityData.BundleName = "cisco_ios_xr"
    allUsedPolicies.EntityData.ParentYangName = "policy-uses"
    allUsedPolicies.EntityData.SegmentPath = "all-used-policies"
    allUsedPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allUsedPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allUsedPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allUsedPolicies.EntityData.Children = types.NewOrderedMap()
    allUsedPolicies.EntityData.Leafs = types.NewOrderedMap()
    allUsedPolicies.EntityData.Leafs.Append("object", types.YLeaf{"Object", allUsedPolicies.Object})

    allUsedPolicies.EntityData.YListKeys = []string {}

    return &(allUsedPolicies.EntityData)
}

// RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "route-policy"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_Attached_Binding.
    Binding []*RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_Attached_Binding
}

func (attached *RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "route-policy"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_Attached_Binding
// bindings list
type RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Policies_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Policies_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Policies_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "policies"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Policies_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Policies_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Policies_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "policies"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Policies_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Policies_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Policies_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "policies"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets
// Information about configured sets
type RoutingPolicyShadow_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about Etag sets.
    Etag RoutingPolicyShadow_Sets_Etag

    // Information about OSPF Area sets.
    OspfArea RoutingPolicyShadow_Sets_OspfArea

    // Information about Extended Community Opaque sets.
    ExtendedCommunityOpaque RoutingPolicyShadow_Sets_ExtendedCommunityOpaque

    // Information about Extended Community SegNH sets.
    ExtendedCommunitySegNh RoutingPolicyShadow_Sets_ExtendedCommunitySegNh

    // Information about Extended Community SOO sets.
    ExtendedCommunitySoo RoutingPolicyShadow_Sets_ExtendedCommunitySoo

    // Information about Tag sets.
    Tag RoutingPolicyShadow_Sets_Tag

    // Information about AS Path sets.
    Prefix RoutingPolicyShadow_Sets_Prefix

    // Information about Community sets.
    Community RoutingPolicyShadow_Sets_Community

    // Information about AS Path sets.
    AsPath RoutingPolicyShadow_Sets_AsPath

    // Information about Large Community sets.
    LargeCommunity RoutingPolicyShadow_Sets_LargeCommunity

    // Information about Esi sets.
    Esi RoutingPolicyShadow_Sets_Esi

    // Information about Extended Community Bandwidth sets.
    ExtendedCommunityBandwidth RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth

    // Information about Extended Community RT sets.
    ExtendedCommunityRt RoutingPolicyShadow_Sets_ExtendedCommunityRt

    // Information about RD sets.
    Rd RoutingPolicyShadow_Sets_Rd

    // Information about Mac sets.
    Mac RoutingPolicyShadow_Sets_Mac

    // Information about Extended Community Cost sets.
    ExtendedCommunityCost RoutingPolicyShadow_Sets_ExtendedCommunityCost
}

func (sets *RoutingPolicyShadow_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "routing-policy-shadow"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("etag", types.YChild{"Etag", &sets.Etag})
    sets.EntityData.Children.Append("ospf-area", types.YChild{"OspfArea", &sets.OspfArea})
    sets.EntityData.Children.Append("extended-community-opaque", types.YChild{"ExtendedCommunityOpaque", &sets.ExtendedCommunityOpaque})
    sets.EntityData.Children.Append("extended-community-seg-nh", types.YChild{"ExtendedCommunitySegNh", &sets.ExtendedCommunitySegNh})
    sets.EntityData.Children.Append("extended-community-soo", types.YChild{"ExtendedCommunitySoo", &sets.ExtendedCommunitySoo})
    sets.EntityData.Children.Append("tag", types.YChild{"Tag", &sets.Tag})
    sets.EntityData.Children.Append("prefix", types.YChild{"Prefix", &sets.Prefix})
    sets.EntityData.Children.Append("community", types.YChild{"Community", &sets.Community})
    sets.EntityData.Children.Append("as-path", types.YChild{"AsPath", &sets.AsPath})
    sets.EntityData.Children.Append("large-community", types.YChild{"LargeCommunity", &sets.LargeCommunity})
    sets.EntityData.Children.Append("esi", types.YChild{"Esi", &sets.Esi})
    sets.EntityData.Children.Append("extended-community-bandwidth", types.YChild{"ExtendedCommunityBandwidth", &sets.ExtendedCommunityBandwidth})
    sets.EntityData.Children.Append("extended-community-rt", types.YChild{"ExtendedCommunityRt", &sets.ExtendedCommunityRt})
    sets.EntityData.Children.Append("rd", types.YChild{"Rd", &sets.Rd})
    sets.EntityData.Children.Append("mac", types.YChild{"Mac", &sets.Mac})
    sets.EntityData.Children.Append("extended-community-cost", types.YChild{"ExtendedCommunityCost", &sets.ExtendedCommunityCost})
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_Etag
// Information about Etag sets
type RoutingPolicyShadow_Sets_Etag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_Etag_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_Etag_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_Etag_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_Etag_Active
}

func (etag *RoutingPolicyShadow_Sets_Etag) GetEntityData() *types.CommonEntityData {
    etag.EntityData.YFilter = etag.YFilter
    etag.EntityData.YangName = "etag"
    etag.EntityData.BundleName = "cisco_ios_xr"
    etag.EntityData.ParentYangName = "sets"
    etag.EntityData.SegmentPath = "etag"
    etag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    etag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    etag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    etag.EntityData.Children = types.NewOrderedMap()
    etag.EntityData.Children.Append("sets", types.YChild{"Sets", &etag.Sets})
    etag.EntityData.Children.Append("unused", types.YChild{"Unused", &etag.Unused})
    etag.EntityData.Children.Append("inactive", types.YChild{"Inactive", &etag.Inactive})
    etag.EntityData.Children.Append("active", types.YChild{"Active", &etag.Active})
    etag.EntityData.Leafs = types.NewOrderedMap()

    etag.EntityData.YListKeys = []string {}

    return &(etag.EntityData)
}

// RoutingPolicyShadow_Sets_Etag_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_Etag_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_Etag_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_Etag_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_Etag_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "etag"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_Etag_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_Etag_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_Etag_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_Etag_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_Etag_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_Etag_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_Etag_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_Etag_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_Etag_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_Etag_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_Etag_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_Etag_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_Etag_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_Etag_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_Etag_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_Etag_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_Etag_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_Etag_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_Etag_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_Etag_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_Etag_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_Etag_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_Etag_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_Etag_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "etag"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_Etag_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_Etag_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_Etag_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "etag"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_Etag_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_Etag_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_Etag_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "etag"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_OspfArea
// Information about OSPF Area sets
type RoutingPolicyShadow_Sets_OspfArea struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_OspfArea_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_OspfArea_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_OspfArea_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_OspfArea_Active
}

func (ospfArea *RoutingPolicyShadow_Sets_OspfArea) GetEntityData() *types.CommonEntityData {
    ospfArea.EntityData.YFilter = ospfArea.YFilter
    ospfArea.EntityData.YangName = "ospf-area"
    ospfArea.EntityData.BundleName = "cisco_ios_xr"
    ospfArea.EntityData.ParentYangName = "sets"
    ospfArea.EntityData.SegmentPath = "ospf-area"
    ospfArea.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfArea.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfArea.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfArea.EntityData.Children = types.NewOrderedMap()
    ospfArea.EntityData.Children.Append("sets", types.YChild{"Sets", &ospfArea.Sets})
    ospfArea.EntityData.Children.Append("unused", types.YChild{"Unused", &ospfArea.Unused})
    ospfArea.EntityData.Children.Append("inactive", types.YChild{"Inactive", &ospfArea.Inactive})
    ospfArea.EntityData.Children.Append("active", types.YChild{"Active", &ospfArea.Active})
    ospfArea.EntityData.Leafs = types.NewOrderedMap()

    ospfArea.EntityData.YListKeys = []string {}

    return &(ospfArea.EntityData)
}

// RoutingPolicyShadow_Sets_OspfArea_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_OspfArea_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_OspfArea_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_OspfArea_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_OspfArea_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "ospf-area"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_OspfArea_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_OspfArea_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_OspfArea_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_OspfArea_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_OspfArea_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_OspfArea_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_OspfArea_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_OspfArea_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_OspfArea_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_OspfArea_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_OspfArea_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_OspfArea_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_OspfArea_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_OspfArea_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_OspfArea_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_OspfArea_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_OspfArea_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_OspfArea_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_OspfArea_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_OspfArea_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_OspfArea_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_OspfArea_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_OspfArea_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "ospf-area"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_OspfArea_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_OspfArea_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_OspfArea_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "ospf-area"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_OspfArea_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_OspfArea_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_OspfArea_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "ospf-area"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityOpaque
// Information about Extended Community Opaque
// sets
type RoutingPolicyShadow_Sets_ExtendedCommunityOpaque struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Active
}

func (extendedCommunityOpaque *RoutingPolicyShadow_Sets_ExtendedCommunityOpaque) GetEntityData() *types.CommonEntityData {
    extendedCommunityOpaque.EntityData.YFilter = extendedCommunityOpaque.YFilter
    extendedCommunityOpaque.EntityData.YangName = "extended-community-opaque"
    extendedCommunityOpaque.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityOpaque.EntityData.ParentYangName = "sets"
    extendedCommunityOpaque.EntityData.SegmentPath = "extended-community-opaque"
    extendedCommunityOpaque.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityOpaque.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityOpaque.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityOpaque.EntityData.Children = types.NewOrderedMap()
    extendedCommunityOpaque.EntityData.Children.Append("sets", types.YChild{"Sets", &extendedCommunityOpaque.Sets})
    extendedCommunityOpaque.EntityData.Children.Append("unused", types.YChild{"Unused", &extendedCommunityOpaque.Unused})
    extendedCommunityOpaque.EntityData.Children.Append("inactive", types.YChild{"Inactive", &extendedCommunityOpaque.Inactive})
    extendedCommunityOpaque.EntityData.Children.Append("active", types.YChild{"Active", &extendedCommunityOpaque.Active})
    extendedCommunityOpaque.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunityOpaque.EntityData.YListKeys = []string {}

    return &(extendedCommunityOpaque.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "extended-community-opaque"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "extended-community-opaque"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "extended-community-opaque"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_ExtendedCommunityOpaque_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "extended-community-opaque"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySegNh
// Information about Extended Community SegNH sets
type RoutingPolicyShadow_Sets_ExtendedCommunitySegNh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Active
}

func (extendedCommunitySegNh *RoutingPolicyShadow_Sets_ExtendedCommunitySegNh) GetEntityData() *types.CommonEntityData {
    extendedCommunitySegNh.EntityData.YFilter = extendedCommunitySegNh.YFilter
    extendedCommunitySegNh.EntityData.YangName = "extended-community-seg-nh"
    extendedCommunitySegNh.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunitySegNh.EntityData.ParentYangName = "sets"
    extendedCommunitySegNh.EntityData.SegmentPath = "extended-community-seg-nh"
    extendedCommunitySegNh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunitySegNh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunitySegNh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunitySegNh.EntityData.Children = types.NewOrderedMap()
    extendedCommunitySegNh.EntityData.Children.Append("sets", types.YChild{"Sets", &extendedCommunitySegNh.Sets})
    extendedCommunitySegNh.EntityData.Children.Append("unused", types.YChild{"Unused", &extendedCommunitySegNh.Unused})
    extendedCommunitySegNh.EntityData.Children.Append("inactive", types.YChild{"Inactive", &extendedCommunitySegNh.Inactive})
    extendedCommunitySegNh.EntityData.Children.Append("active", types.YChild{"Active", &extendedCommunitySegNh.Active})
    extendedCommunitySegNh.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunitySegNh.EntityData.YListKeys = []string {}

    return &(extendedCommunitySegNh.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "extended-community-seg-nh"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "extended-community-seg-nh"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "extended-community-seg-nh"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_ExtendedCommunitySegNh_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "extended-community-seg-nh"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySoo
// Information about Extended Community SOO sets
type RoutingPolicyShadow_Sets_ExtendedCommunitySoo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Active
}

func (extendedCommunitySoo *RoutingPolicyShadow_Sets_ExtendedCommunitySoo) GetEntityData() *types.CommonEntityData {
    extendedCommunitySoo.EntityData.YFilter = extendedCommunitySoo.YFilter
    extendedCommunitySoo.EntityData.YangName = "extended-community-soo"
    extendedCommunitySoo.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunitySoo.EntityData.ParentYangName = "sets"
    extendedCommunitySoo.EntityData.SegmentPath = "extended-community-soo"
    extendedCommunitySoo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunitySoo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunitySoo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunitySoo.EntityData.Children = types.NewOrderedMap()
    extendedCommunitySoo.EntityData.Children.Append("sets", types.YChild{"Sets", &extendedCommunitySoo.Sets})
    extendedCommunitySoo.EntityData.Children.Append("unused", types.YChild{"Unused", &extendedCommunitySoo.Unused})
    extendedCommunitySoo.EntityData.Children.Append("inactive", types.YChild{"Inactive", &extendedCommunitySoo.Inactive})
    extendedCommunitySoo.EntityData.Children.Append("active", types.YChild{"Active", &extendedCommunitySoo.Active})
    extendedCommunitySoo.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunitySoo.EntityData.YListKeys = []string {}

    return &(extendedCommunitySoo.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "extended-community-soo"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "extended-community-soo"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "extended-community-soo"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_ExtendedCommunitySoo_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "extended-community-soo"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_Tag
// Information about Tag sets
type RoutingPolicyShadow_Sets_Tag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_Tag_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_Tag_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_Tag_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_Tag_Active
}

func (tag *RoutingPolicyShadow_Sets_Tag) GetEntityData() *types.CommonEntityData {
    tag.EntityData.YFilter = tag.YFilter
    tag.EntityData.YangName = "tag"
    tag.EntityData.BundleName = "cisco_ios_xr"
    tag.EntityData.ParentYangName = "sets"
    tag.EntityData.SegmentPath = "tag"
    tag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tag.EntityData.Children = types.NewOrderedMap()
    tag.EntityData.Children.Append("sets", types.YChild{"Sets", &tag.Sets})
    tag.EntityData.Children.Append("unused", types.YChild{"Unused", &tag.Unused})
    tag.EntityData.Children.Append("inactive", types.YChild{"Inactive", &tag.Inactive})
    tag.EntityData.Children.Append("active", types.YChild{"Active", &tag.Active})
    tag.EntityData.Leafs = types.NewOrderedMap()

    tag.EntityData.YListKeys = []string {}

    return &(tag.EntityData)
}

// RoutingPolicyShadow_Sets_Tag_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_Tag_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_Tag_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_Tag_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_Tag_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "tag"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_Tag_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_Tag_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_Tag_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_Tag_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_Tag_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_Tag_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_Tag_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_Tag_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_Tag_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_Tag_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_Tag_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_Tag_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_Tag_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_Tag_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_Tag_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_Tag_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_Tag_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_Tag_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_Tag_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_Tag_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_Tag_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_Tag_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_Tag_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_Tag_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "tag"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_Tag_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_Tag_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_Tag_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "tag"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_Tag_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_Tag_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_Tag_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "tag"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_Prefix
// Information about AS Path sets
type RoutingPolicyShadow_Sets_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_Prefix_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_Prefix_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_Prefix_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_Prefix_Active
}

func (prefix *RoutingPolicyShadow_Sets_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "sets"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Children.Append("sets", types.YChild{"Sets", &prefix.Sets})
    prefix.EntityData.Children.Append("unused", types.YChild{"Unused", &prefix.Unused})
    prefix.EntityData.Children.Append("inactive", types.YChild{"Inactive", &prefix.Inactive})
    prefix.EntityData.Children.Append("active", types.YChild{"Active", &prefix.Active})
    prefix.EntityData.Leafs = types.NewOrderedMap()

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// RoutingPolicyShadow_Sets_Prefix_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_Prefix_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_Prefix_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_Prefix_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_Prefix_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "prefix"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_Prefix_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_Prefix_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_Prefix_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_Prefix_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_Prefix_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_Prefix_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_Prefix_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_Prefix_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_Prefix_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_Prefix_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_Prefix_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_Prefix_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_Prefix_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_Prefix_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_Prefix_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_Prefix_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_Prefix_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_Prefix_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_Prefix_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_Prefix_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_Prefix_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_Prefix_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_Prefix_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_Prefix_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "prefix"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_Prefix_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_Prefix_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_Prefix_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "prefix"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_Prefix_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_Prefix_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_Prefix_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "prefix"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_Community
// Information about Community sets
type RoutingPolicyShadow_Sets_Community struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_Community_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_Community_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_Community_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_Community_Active
}

func (community *RoutingPolicyShadow_Sets_Community) GetEntityData() *types.CommonEntityData {
    community.EntityData.YFilter = community.YFilter
    community.EntityData.YangName = "community"
    community.EntityData.BundleName = "cisco_ios_xr"
    community.EntityData.ParentYangName = "sets"
    community.EntityData.SegmentPath = "community"
    community.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    community.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    community.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    community.EntityData.Children = types.NewOrderedMap()
    community.EntityData.Children.Append("sets", types.YChild{"Sets", &community.Sets})
    community.EntityData.Children.Append("unused", types.YChild{"Unused", &community.Unused})
    community.EntityData.Children.Append("inactive", types.YChild{"Inactive", &community.Inactive})
    community.EntityData.Children.Append("active", types.YChild{"Active", &community.Active})
    community.EntityData.Leafs = types.NewOrderedMap()

    community.EntityData.YListKeys = []string {}

    return &(community.EntityData)
}

// RoutingPolicyShadow_Sets_Community_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_Community_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_Community_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_Community_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_Community_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "community"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_Community_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_Community_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_Community_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_Community_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_Community_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_Community_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_Community_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_Community_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_Community_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_Community_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_Community_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_Community_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_Community_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_Community_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_Community_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_Community_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_Community_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_Community_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_Community_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_Community_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_Community_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_Community_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_Community_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_Community_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "community"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_Community_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_Community_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_Community_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "community"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_Community_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_Community_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_Community_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "community"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_AsPath
// Information about AS Path sets
type RoutingPolicyShadow_Sets_AsPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_AsPath_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_AsPath_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_AsPath_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_AsPath_Active
}

func (asPath *RoutingPolicyShadow_Sets_AsPath) GetEntityData() *types.CommonEntityData {
    asPath.EntityData.YFilter = asPath.YFilter
    asPath.EntityData.YangName = "as-path"
    asPath.EntityData.BundleName = "cisco_ios_xr"
    asPath.EntityData.ParentYangName = "sets"
    asPath.EntityData.SegmentPath = "as-path"
    asPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asPath.EntityData.Children = types.NewOrderedMap()
    asPath.EntityData.Children.Append("sets", types.YChild{"Sets", &asPath.Sets})
    asPath.EntityData.Children.Append("unused", types.YChild{"Unused", &asPath.Unused})
    asPath.EntityData.Children.Append("inactive", types.YChild{"Inactive", &asPath.Inactive})
    asPath.EntityData.Children.Append("active", types.YChild{"Active", &asPath.Active})
    asPath.EntityData.Leafs = types.NewOrderedMap()

    asPath.EntityData.YListKeys = []string {}

    return &(asPath.EntityData)
}

// RoutingPolicyShadow_Sets_AsPath_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_AsPath_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_AsPath_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_AsPath_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_AsPath_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "as-path"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_AsPath_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_AsPath_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_AsPath_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_AsPath_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_AsPath_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_AsPath_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_AsPath_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_AsPath_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_AsPath_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_AsPath_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_AsPath_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_AsPath_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_AsPath_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_AsPath_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_AsPath_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_AsPath_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_AsPath_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_AsPath_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_AsPath_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_AsPath_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_AsPath_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_AsPath_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_AsPath_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_AsPath_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "as-path"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_AsPath_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_AsPath_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_AsPath_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "as-path"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_AsPath_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_AsPath_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_AsPath_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "as-path"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_LargeCommunity
// Information about Large Community sets
type RoutingPolicyShadow_Sets_LargeCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_LargeCommunity_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_LargeCommunity_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_LargeCommunity_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_LargeCommunity_Active
}

func (largeCommunity *RoutingPolicyShadow_Sets_LargeCommunity) GetEntityData() *types.CommonEntityData {
    largeCommunity.EntityData.YFilter = largeCommunity.YFilter
    largeCommunity.EntityData.YangName = "large-community"
    largeCommunity.EntityData.BundleName = "cisco_ios_xr"
    largeCommunity.EntityData.ParentYangName = "sets"
    largeCommunity.EntityData.SegmentPath = "large-community"
    largeCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    largeCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    largeCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    largeCommunity.EntityData.Children = types.NewOrderedMap()
    largeCommunity.EntityData.Children.Append("sets", types.YChild{"Sets", &largeCommunity.Sets})
    largeCommunity.EntityData.Children.Append("unused", types.YChild{"Unused", &largeCommunity.Unused})
    largeCommunity.EntityData.Children.Append("inactive", types.YChild{"Inactive", &largeCommunity.Inactive})
    largeCommunity.EntityData.Children.Append("active", types.YChild{"Active", &largeCommunity.Active})
    largeCommunity.EntityData.Leafs = types.NewOrderedMap()

    largeCommunity.EntityData.YListKeys = []string {}

    return &(largeCommunity.EntityData)
}

// RoutingPolicyShadow_Sets_LargeCommunity_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_LargeCommunity_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_LargeCommunity_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "large-community"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_LargeCommunity_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_LargeCommunity_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_LargeCommunity_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "large-community"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_LargeCommunity_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_LargeCommunity_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_LargeCommunity_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "large-community"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_LargeCommunity_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_LargeCommunity_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_LargeCommunity_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "large-community"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_Esi
// Information about Esi sets
type RoutingPolicyShadow_Sets_Esi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_Esi_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_Esi_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_Esi_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_Esi_Active
}

func (esi *RoutingPolicyShadow_Sets_Esi) GetEntityData() *types.CommonEntityData {
    esi.EntityData.YFilter = esi.YFilter
    esi.EntityData.YangName = "esi"
    esi.EntityData.BundleName = "cisco_ios_xr"
    esi.EntityData.ParentYangName = "sets"
    esi.EntityData.SegmentPath = "esi"
    esi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esi.EntityData.Children = types.NewOrderedMap()
    esi.EntityData.Children.Append("sets", types.YChild{"Sets", &esi.Sets})
    esi.EntityData.Children.Append("unused", types.YChild{"Unused", &esi.Unused})
    esi.EntityData.Children.Append("inactive", types.YChild{"Inactive", &esi.Inactive})
    esi.EntityData.Children.Append("active", types.YChild{"Active", &esi.Active})
    esi.EntityData.Leafs = types.NewOrderedMap()

    esi.EntityData.YListKeys = []string {}

    return &(esi.EntityData)
}

// RoutingPolicyShadow_Sets_Esi_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_Esi_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_Esi_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_Esi_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_Esi_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "esi"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_Esi_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_Esi_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_Esi_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_Esi_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_Esi_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_Esi_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_Esi_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_Esi_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_Esi_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_Esi_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_Esi_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_Esi_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_Esi_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_Esi_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_Esi_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_Esi_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_Esi_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_Esi_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_Esi_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_Esi_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_Esi_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_Esi_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_Esi_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_Esi_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "esi"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_Esi_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_Esi_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_Esi_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "esi"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_Esi_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_Esi_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_Esi_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "esi"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth
// Information about Extended Community Bandwidth
// sets
type RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Inactive
}

func (extendedCommunityBandwidth *RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth) GetEntityData() *types.CommonEntityData {
    extendedCommunityBandwidth.EntityData.YFilter = extendedCommunityBandwidth.YFilter
    extendedCommunityBandwidth.EntityData.YangName = "extended-community-bandwidth"
    extendedCommunityBandwidth.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityBandwidth.EntityData.ParentYangName = "sets"
    extendedCommunityBandwidth.EntityData.SegmentPath = "extended-community-bandwidth"
    extendedCommunityBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityBandwidth.EntityData.Children = types.NewOrderedMap()
    extendedCommunityBandwidth.EntityData.Children.Append("sets", types.YChild{"Sets", &extendedCommunityBandwidth.Sets})
    extendedCommunityBandwidth.EntityData.Children.Append("unused", types.YChild{"Unused", &extendedCommunityBandwidth.Unused})
    extendedCommunityBandwidth.EntityData.Children.Append("inactive", types.YChild{"Inactive", &extendedCommunityBandwidth.Inactive})
    extendedCommunityBandwidth.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunityBandwidth.EntityData.YListKeys = []string {}

    return &(extendedCommunityBandwidth.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "extended-community-bandwidth"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "extended-community-bandwidth"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_ExtendedCommunityBandwidth_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "extended-community-bandwidth"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityRt
// Information about Extended Community RT sets
type RoutingPolicyShadow_Sets_ExtendedCommunityRt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_ExtendedCommunityRt_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_ExtendedCommunityRt_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_ExtendedCommunityRt_Active
}

func (extendedCommunityRt *RoutingPolicyShadow_Sets_ExtendedCommunityRt) GetEntityData() *types.CommonEntityData {
    extendedCommunityRt.EntityData.YFilter = extendedCommunityRt.YFilter
    extendedCommunityRt.EntityData.YangName = "extended-community-rt"
    extendedCommunityRt.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityRt.EntityData.ParentYangName = "sets"
    extendedCommunityRt.EntityData.SegmentPath = "extended-community-rt"
    extendedCommunityRt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityRt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityRt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityRt.EntityData.Children = types.NewOrderedMap()
    extendedCommunityRt.EntityData.Children.Append("sets", types.YChild{"Sets", &extendedCommunityRt.Sets})
    extendedCommunityRt.EntityData.Children.Append("unused", types.YChild{"Unused", &extendedCommunityRt.Unused})
    extendedCommunityRt.EntityData.Children.Append("inactive", types.YChild{"Inactive", &extendedCommunityRt.Inactive})
    extendedCommunityRt.EntityData.Children.Append("active", types.YChild{"Active", &extendedCommunityRt.Active})
    extendedCommunityRt.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunityRt.EntityData.YListKeys = []string {}

    return &(extendedCommunityRt.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "extended-community-rt"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityRt_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_ExtendedCommunityRt_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_ExtendedCommunityRt_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "extended-community-rt"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityRt_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_ExtendedCommunityRt_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_ExtendedCommunityRt_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "extended-community-rt"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityRt_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_ExtendedCommunityRt_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_ExtendedCommunityRt_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "extended-community-rt"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_Rd
// Information about RD sets
type RoutingPolicyShadow_Sets_Rd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_Rd_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_Rd_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_Rd_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_Rd_Active
}

func (rd *RoutingPolicyShadow_Sets_Rd) GetEntityData() *types.CommonEntityData {
    rd.EntityData.YFilter = rd.YFilter
    rd.EntityData.YangName = "rd"
    rd.EntityData.BundleName = "cisco_ios_xr"
    rd.EntityData.ParentYangName = "sets"
    rd.EntityData.SegmentPath = "rd"
    rd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rd.EntityData.Children = types.NewOrderedMap()
    rd.EntityData.Children.Append("sets", types.YChild{"Sets", &rd.Sets})
    rd.EntityData.Children.Append("unused", types.YChild{"Unused", &rd.Unused})
    rd.EntityData.Children.Append("inactive", types.YChild{"Inactive", &rd.Inactive})
    rd.EntityData.Children.Append("active", types.YChild{"Active", &rd.Active})
    rd.EntityData.Leafs = types.NewOrderedMap()

    rd.EntityData.YListKeys = []string {}

    return &(rd.EntityData)
}

// RoutingPolicyShadow_Sets_Rd_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_Rd_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_Rd_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_Rd_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_Rd_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "rd"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_Rd_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_Rd_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_Rd_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_Rd_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_Rd_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_Rd_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_Rd_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_Rd_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_Rd_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_Rd_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_Rd_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_Rd_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_Rd_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_Rd_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_Rd_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_Rd_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_Rd_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_Rd_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_Rd_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_Rd_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_Rd_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_Rd_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_Rd_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_Rd_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "rd"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_Rd_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_Rd_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_Rd_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "rd"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_Rd_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_Rd_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_Rd_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "rd"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_Mac
// Information about Mac sets
type RoutingPolicyShadow_Sets_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_Mac_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_Mac_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_Mac_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_Mac_Active
}

func (mac *RoutingPolicyShadow_Sets_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xr"
    mac.EntityData.ParentYangName = "sets"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Children.Append("sets", types.YChild{"Sets", &mac.Sets})
    mac.EntityData.Children.Append("unused", types.YChild{"Unused", &mac.Unused})
    mac.EntityData.Children.Append("inactive", types.YChild{"Inactive", &mac.Inactive})
    mac.EntityData.Children.Append("active", types.YChild{"Active", &mac.Active})
    mac.EntityData.Leafs = types.NewOrderedMap()

    mac.EntityData.YListKeys = []string {}

    return &(mac.EntityData)
}

// RoutingPolicyShadow_Sets_Mac_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_Mac_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_Mac_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_Mac_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_Mac_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "mac"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_Mac_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_Mac_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_Mac_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_Mac_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_Mac_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_Mac_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_Mac_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_Mac_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_Mac_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_Mac_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_Mac_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_Mac_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_Mac_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_Mac_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_Mac_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_Mac_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_Mac_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_Mac_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_Mac_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_Mac_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_Mac_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_Mac_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_Mac_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_Mac_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "mac"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_Mac_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_Mac_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_Mac_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "mac"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_Mac_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_Mac_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_Mac_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "mac"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityCost
// Information about Extended Community Cost sets
type RoutingPolicyShadow_Sets_ExtendedCommunityCost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicyShadow_Sets_ExtendedCommunityCost_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicyShadow_Sets_ExtendedCommunityCost_Inactive

    // All objects of a given type that are attached to a protocol.
    Active RoutingPolicyShadow_Sets_ExtendedCommunityCost_Active
}

func (extendedCommunityCost *RoutingPolicyShadow_Sets_ExtendedCommunityCost) GetEntityData() *types.CommonEntityData {
    extendedCommunityCost.EntityData.YFilter = extendedCommunityCost.YFilter
    extendedCommunityCost.EntityData.YangName = "extended-community-cost"
    extendedCommunityCost.EntityData.BundleName = "cisco_ios_xr"
    extendedCommunityCost.EntityData.ParentYangName = "sets"
    extendedCommunityCost.EntityData.SegmentPath = "extended-community-cost"
    extendedCommunityCost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedCommunityCost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedCommunityCost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedCommunityCost.EntityData.Children = types.NewOrderedMap()
    extendedCommunityCost.EntityData.Children.Append("sets", types.YChild{"Sets", &extendedCommunityCost.Sets})
    extendedCommunityCost.EntityData.Children.Append("unused", types.YChild{"Unused", &extendedCommunityCost.Unused})
    extendedCommunityCost.EntityData.Children.Append("inactive", types.YChild{"Inactive", &extendedCommunityCost.Inactive})
    extendedCommunityCost.EntityData.Children.Append("active", types.YChild{"Active", &extendedCommunityCost.Active})
    extendedCommunityCost.EntityData.Leafs = types.NewOrderedMap()

    extendedCommunityCost.EntityData.YListKeys = []string {}

    return &(extendedCommunityCost.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets
// Information about individual sets
type RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set.
    Set []*RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set
}

func (sets *RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets) GetEntityData() *types.CommonEntityData {
    sets.EntityData.YFilter = sets.YFilter
    sets.EntityData.YangName = "sets"
    sets.EntityData.BundleName = "cisco_ios_xr"
    sets.EntityData.ParentYangName = "extended-community-cost"
    sets.EntityData.SegmentPath = "sets"
    sets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets.EntityData.Children = types.NewOrderedMap()
    sets.EntityData.Children.Append("set", types.YChild{"Set", nil})
    for i := range sets.Set {
        sets.EntityData.Children.Append(types.GetSegmentPath(sets.Set[i]), types.YChild{"Set", sets.Set[i]})
    }
    sets.EntityData.Leafs = types.NewOrderedMap()

    sets.EntityData.YListKeys = []string {}

    return &(sets.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set
// Information about an individual set
type RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_Attached
}

func (set *RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + types.AddKeyToken(set.SetName, "set-name")
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Children.Append("used-by", types.YChild{"UsedBy", &set.UsedBy})
    set.EntityData.Children.Append("attached", types.YChild{"Attached", &set.Attached})
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", set.SetName})

    set.EntityData.YListKeys = []string {"SetName"}

    return &(set.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference.
    Reference []*RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = types.NewOrderedMap()
    usedBy.EntityData.Children.Append("reference", types.YChild{"Reference", nil})
    for i := range usedBy.Reference {
        usedBy.EntityData.Children.Append(types.GetSegmentPath(usedBy.Reference[i]), types.YChild{"Reference", usedBy.Reference[i]})
    }
    usedBy.EntityData.Leafs = types.NewOrderedMap()

    usedBy.EntityData.YListKeys = []string {}

    return &(usedBy.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = types.NewOrderedMap()
    reference.EntityData.Leafs = types.NewOrderedMap()
    reference.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", reference.RoutePolicyName})
    reference.EntityData.Leafs.Append("used-directly", types.YLeaf{"UsedDirectly", reference.UsedDirectly})
    reference.EntityData.Leafs.Append("status", types.YLeaf{"Status", reference.Status})

    reference.EntityData.YListKeys = []string {}

    return &(reference.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding.
    Binding []*RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = types.NewOrderedMap()
    attached.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range attached.Binding {
        attached.EntityData.Children.Append(types.GetSegmentPath(attached.Binding[i]), types.YChild{"Binding", attached.Binding[i]})
    }
    attached.EntityData.Leafs = types.NewOrderedMap()

    attached.EntityData.YListKeys = []string {}

    return &(attached.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol to which policy attached. The type is string.
    Protocol interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Protocol instance. The type is string.
    ProtoInstance interface{}

    // Address Family Identifier. The type is AddressFamily.
    AfName interface{}

    // Subsequent Address Family Identifier. The type is SubAddressFamily.
    SafName interface{}

    // Neighbor IP Address. The type is string.
    NeighborAddress interface{}

    // Neighbor IP Address Family. The type is AddressFamily.
    NeighborAfName interface{}

    // Neighbor Group Name. The type is string.
    GroupName interface{}

    // Direction In or Out. The type is AttachPointDirection.
    Direction interface{}

    // Neighbor Group . The type is Group.
    Group interface{}

    // Source Protocol to redistribute, Source Protocol can be one of the
    // following values{all, connected, local, static, bgp, rip, isis, ospf
    // ,ospfv3, eigrp, unknown }. The type is string.
    SourceProtocol interface{}

    // Aggregate IP address or Network IP Address in IPv4 or IPv6 Format. The type
    // is string.
    AggregateNetworkAddress interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Instance. The type is string.
    Instance interface{}

    // OSPF Area ID in Decimal Integer Format. The type is string.
    AreaId interface{}

    // ISIS Propogate From Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateFrom interface{}

    // ISIS Propogate To Level. The type is interface{} with range:
    // -2147483648..2147483647.
    PropogateTo interface{}

    // Policy that uses object in question. The type is string.
    RoutePolicyName interface{}

    // The attached policy that (maybe indirectly) uses the object in question.
    // The type is string.
    AttachedPolicy interface{}

    // Name of attach point where policy is attached. The type is string.
    AttachPoint interface{}
}

func (binding *RoutingPolicyShadow_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", binding.Protocol})
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("proto-instance", types.YLeaf{"ProtoInstance", binding.ProtoInstance})
    binding.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", binding.AfName})
    binding.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", binding.SafName})
    binding.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", binding.NeighborAddress})
    binding.EntityData.Leafs.Append("neighbor-af-name", types.YLeaf{"NeighborAfName", binding.NeighborAfName})
    binding.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", binding.GroupName})
    binding.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", binding.Direction})
    binding.EntityData.Leafs.Append("group", types.YLeaf{"Group", binding.Group})
    binding.EntityData.Leafs.Append("source-protocol", types.YLeaf{"SourceProtocol", binding.SourceProtocol})
    binding.EntityData.Leafs.Append("aggregate-network-address", types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress})
    binding.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", binding.InterfaceName})
    binding.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", binding.Instance})
    binding.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", binding.AreaId})
    binding.EntityData.Leafs.Append("propogate-from", types.YLeaf{"PropogateFrom", binding.PropogateFrom})
    binding.EntityData.Leafs.Append("propogate-to", types.YLeaf{"PropogateTo", binding.PropogateTo})
    binding.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", binding.RoutePolicyName})
    binding.EntityData.Leafs.Append("attached-policy", types.YLeaf{"AttachedPolicy", binding.AttachedPolicy})
    binding.EntityData.Leafs.Append("attach-point", types.YLeaf{"AttachPoint", binding.AttachPoint})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityCost_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicyShadow_Sets_ExtendedCommunityCost_Unused struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicyShadow_Sets_ExtendedCommunityCost_Unused) GetEntityData() *types.CommonEntityData {
    unused.EntityData.YFilter = unused.YFilter
    unused.EntityData.YangName = "unused"
    unused.EntityData.BundleName = "cisco_ios_xr"
    unused.EntityData.ParentYangName = "extended-community-cost"
    unused.EntityData.SegmentPath = "unused"
    unused.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unused.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unused.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unused.EntityData.Children = types.NewOrderedMap()
    unused.EntityData.Leafs = types.NewOrderedMap()
    unused.EntityData.Leafs.Append("object", types.YLeaf{"Object", unused.Object})

    unused.EntityData.YListKeys = []string {}

    return &(unused.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityCost_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicyShadow_Sets_ExtendedCommunityCost_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicyShadow_Sets_ExtendedCommunityCost_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "extended-community-cost"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("object", types.YLeaf{"Object", inactive.Object})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// RoutingPolicyShadow_Sets_ExtendedCommunityCost_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicyShadow_Sets_ExtendedCommunityCost_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicyShadow_Sets_ExtendedCommunityCost_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "extended-community-cost"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("object", types.YLeaf{"Object", active.Object})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

