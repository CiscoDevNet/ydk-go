// This module contains a collection of YANG definitions
// for Cisco IOS-XR policy-repository package operational data.
// 
// This module contains definitions
// for the following management objects:
//   routing-policy: Routing policy operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

    routingPolicy.EntityData.Children = make(map[string]types.YChild)
    routingPolicy.EntityData.Children["limits"] = types.YChild{"Limits", &routingPolicy.Limits}
    routingPolicy.EntityData.Children["policies"] = types.YChild{"Policies", &routingPolicy.Policies}
    routingPolicy.EntityData.Children["sets"] = types.YChild{"Sets", &routingPolicy.Sets}
    routingPolicy.EntityData.Leafs = make(map[string]types.YLeaf)
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

    limits.EntityData.Children = make(map[string]types.YChild)
    limits.EntityData.Leafs = make(map[string]types.YLeaf)
    limits.EntityData.Leafs["maximum-lines-of-policy"] = types.YLeaf{"MaximumLinesOfPolicy", limits.MaximumLinesOfPolicy}
    limits.EntityData.Leafs["current-lines-of-policy-limit"] = types.YLeaf{"CurrentLinesOfPolicyLimit", limits.CurrentLinesOfPolicyLimit}
    limits.EntityData.Leafs["current-lines-of-policy-used"] = types.YLeaf{"CurrentLinesOfPolicyUsed", limits.CurrentLinesOfPolicyUsed}
    limits.EntityData.Leafs["maximum-number-of-policies"] = types.YLeaf{"MaximumNumberOfPolicies", limits.MaximumNumberOfPolicies}
    limits.EntityData.Leafs["current-number-of-policies-limit"] = types.YLeaf{"CurrentNumberOfPoliciesLimit", limits.CurrentNumberOfPoliciesLimit}
    limits.EntityData.Leafs["current-number-of-policies-used"] = types.YLeaf{"CurrentNumberOfPoliciesUsed", limits.CurrentNumberOfPoliciesUsed}
    limits.EntityData.Leafs["compiled-policies-length"] = types.YLeaf{"CompiledPoliciesLength", limits.CompiledPoliciesLength}
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

    policies.EntityData.Children = make(map[string]types.YChild)
    policies.EntityData.Children["route-policies"] = types.YChild{"RoutePolicies", &policies.RoutePolicies}
    policies.EntityData.Children["unused"] = types.YChild{"Unused", &policies.Unused}
    policies.EntityData.Children["inactive"] = types.YChild{"Inactive", &policies.Inactive}
    policies.EntityData.Children["active"] = types.YChild{"Active", &policies.Active}
    policies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policies.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies
// Information about individual policies
type RoutingPolicy_Policies_RoutePolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual policy. The type is slice of
    // RoutingPolicy_Policies_RoutePolicies_RoutePolicy.
    RoutePolicy []RoutingPolicy_Policies_RoutePolicies_RoutePolicy
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

    routePolicies.EntityData.Children = make(map[string]types.YChild)
    routePolicies.EntityData.Children["route-policy"] = types.YChild{"RoutePolicy", nil}
    for i := range routePolicies.RoutePolicy {
        routePolicies.EntityData.Children[types.GetSegmentPath(&routePolicies.RoutePolicy[i])] = types.YChild{"RoutePolicy", &routePolicies.RoutePolicy[i]}
    }
    routePolicies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routePolicies.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy
// Information about an individual policy
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Route policy name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    routePolicy.EntityData.SegmentPath = "route-policy" + "[route-policy-name='" + fmt.Sprintf("%v", routePolicy.RoutePolicyName) + "']"
    routePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routePolicy.EntityData.Children = make(map[string]types.YChild)
    routePolicy.EntityData.Children["policy-uses"] = types.YChild{"PolicyUses", &routePolicy.PolicyUses}
    routePolicy.EntityData.Children["used-by"] = types.YChild{"UsedBy", &routePolicy.UsedBy}
    routePolicy.EntityData.Children["attached"] = types.YChild{"Attached", &routePolicy.Attached}
    routePolicy.EntityData.Leafs = make(map[string]types.YLeaf)
    routePolicy.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", routePolicy.RoutePolicyName}
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

    policyUses.EntityData.Children = make(map[string]types.YChild)
    policyUses.EntityData.Children["directly-used-policies"] = types.YChild{"DirectlyUsedPolicies", &policyUses.DirectlyUsedPolicies}
    policyUses.EntityData.Children["all-used-sets"] = types.YChild{"AllUsedSets", &policyUses.AllUsedSets}
    policyUses.EntityData.Children["directly-used-sets"] = types.YChild{"DirectlyUsedSets", &policyUses.DirectlyUsedSets}
    policyUses.EntityData.Children["all-used-policies"] = types.YChild{"AllUsedPolicies", &policyUses.AllUsedPolicies}
    policyUses.EntityData.Leafs = make(map[string]types.YLeaf)
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

    directlyUsedPolicies.EntityData.Children = make(map[string]types.YChild)
    directlyUsedPolicies.EntityData.Leafs = make(map[string]types.YLeaf)
    directlyUsedPolicies.EntityData.Leafs["object"] = types.YLeaf{"Object", directlyUsedPolicies.Object}
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
    Sets []RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets
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

    allUsedSets.EntityData.Children = make(map[string]types.YChild)
    allUsedSets.EntityData.Children["sets"] = types.YChild{"Sets", nil}
    for i := range allUsedSets.Sets {
        allUsedSets.EntityData.Children[types.GetSegmentPath(&allUsedSets.Sets[i])] = types.YChild{"Sets", &allUsedSets.Sets[i]}
    }
    allUsedSets.EntityData.Leafs = make(map[string]types.YLeaf)
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

    sets.EntityData.Children = make(map[string]types.YChild)
    sets.EntityData.Leafs = make(map[string]types.YLeaf)
    sets.EntityData.Leafs["set-domain"] = types.YLeaf{"SetDomain", sets.SetDomain}
    sets.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", sets.SetName}
    return &(sets.EntityData)
}

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets
// Sets that this policy uses directly
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of sets in several domains. The type is slice of
    // RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets.
    Sets []RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets
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

    directlyUsedSets.EntityData.Children = make(map[string]types.YChild)
    directlyUsedSets.EntityData.Children["sets"] = types.YChild{"Sets", nil}
    for i := range directlyUsedSets.Sets {
        directlyUsedSets.EntityData.Children[types.GetSegmentPath(&directlyUsedSets.Sets[i])] = types.YChild{"Sets", &directlyUsedSets.Sets[i]}
    }
    directlyUsedSets.EntityData.Leafs = make(map[string]types.YLeaf)
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

    sets.EntityData.Children = make(map[string]types.YChild)
    sets.EntityData.Leafs = make(map[string]types.YLeaf)
    sets.EntityData.Leafs["set-domain"] = types.YLeaf{"SetDomain", sets.SetDomain}
    sets.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", sets.SetName}
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

    allUsedPolicies.EntityData.Children = make(map[string]types.YChild)
    allUsedPolicies.EntityData.Leafs = make(map[string]types.YLeaf)
    allUsedPolicies.EntityData.Leafs["object"] = types.YLeaf{"Object", allUsedPolicies.Object}
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
    Reference []RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference
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

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
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

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
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
    Binding []RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding
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

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
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

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
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

    sets.EntityData.Children = make(map[string]types.YChild)
    sets.EntityData.Children["etag"] = types.YChild{"Etag", &sets.Etag}
    sets.EntityData.Children["ospf-area"] = types.YChild{"OspfArea", &sets.OspfArea}
    sets.EntityData.Children["extended-community-opaque"] = types.YChild{"ExtendedCommunityOpaque", &sets.ExtendedCommunityOpaque}
    sets.EntityData.Children["extended-community-seg-nh"] = types.YChild{"ExtendedCommunitySegNh", &sets.ExtendedCommunitySegNh}
    sets.EntityData.Children["extended-community-soo"] = types.YChild{"ExtendedCommunitySoo", &sets.ExtendedCommunitySoo}
    sets.EntityData.Children["tag"] = types.YChild{"Tag", &sets.Tag}
    sets.EntityData.Children["prefix"] = types.YChild{"Prefix", &sets.Prefix}
    sets.EntityData.Children["community"] = types.YChild{"Community", &sets.Community}
    sets.EntityData.Children["as-path"] = types.YChild{"AsPath", &sets.AsPath}
    sets.EntityData.Children["large-community"] = types.YChild{"LargeCommunity", &sets.LargeCommunity}
    sets.EntityData.Children["esi"] = types.YChild{"Esi", &sets.Esi}
    sets.EntityData.Children["extended-community-bandwidth"] = types.YChild{"ExtendedCommunityBandwidth", &sets.ExtendedCommunityBandwidth}
    sets.EntityData.Children["extended-community-rt"] = types.YChild{"ExtendedCommunityRt", &sets.ExtendedCommunityRt}
    sets.EntityData.Children["rd"] = types.YChild{"Rd", &sets.Rd}
    sets.EntityData.Children["mac"] = types.YChild{"Mac", &sets.Mac}
    sets.EntityData.Children["extended-community-cost"] = types.YChild{"ExtendedCommunityCost", &sets.ExtendedCommunityCost}
    sets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets.EntityData)
}

// RoutingPolicy_Sets_Etag
// Information about Etag sets
type RoutingPolicy_Sets_Etag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Etag_Sets_

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

    etag.EntityData.Children = make(map[string]types.YChild)
    etag.EntityData.Children["sets"] = types.YChild{"Sets", &etag.Sets}
    etag.EntityData.Children["unused"] = types.YChild{"Unused", &etag.Unused}
    etag.EntityData.Children["inactive"] = types.YChild{"Inactive", &etag.Inactive}
    etag.EntityData.Children["active"] = types.YChild{"Active", &etag.Active}
    etag.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(etag.EntityData)
}

// RoutingPolicy_Sets_Etag_Sets_
// Information about individual sets
type RoutingPolicy_Sets_Etag_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Etag_Sets__Set.
    Set []RoutingPolicy_Sets_Etag_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_Etag_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "etag"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_Etag_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_Etag_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Etag_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Etag_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_Etag_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_Etag_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Etag_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Etag_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Etag_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Etag_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Etag_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Etag_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_Etag_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Etag_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Etag_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Etag_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Etag_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Etag_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Etag_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Etag_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_Etag_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_OspfArea
// Information about OSPF Area sets
type RoutingPolicy_Sets_OspfArea struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_OspfArea_Sets_

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

    ospfArea.EntityData.Children = make(map[string]types.YChild)
    ospfArea.EntityData.Children["sets"] = types.YChild{"Sets", &ospfArea.Sets}
    ospfArea.EntityData.Children["unused"] = types.YChild{"Unused", &ospfArea.Unused}
    ospfArea.EntityData.Children["inactive"] = types.YChild{"Inactive", &ospfArea.Inactive}
    ospfArea.EntityData.Children["active"] = types.YChild{"Active", &ospfArea.Active}
    ospfArea.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfArea.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Sets_
// Information about individual sets
type RoutingPolicy_Sets_OspfArea_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_OspfArea_Sets__Set.
    Set []RoutingPolicy_Sets_OspfArea_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_OspfArea_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "ospf-area"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_OspfArea_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_OspfArea_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_OspfArea_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_OspfArea_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_OspfArea_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_OspfArea_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_OspfArea_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_OspfArea_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_OspfArea_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_OspfArea_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_OspfArea_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_OspfArea_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_OspfArea_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_OspfArea_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_OspfArea_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_OspfArea_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque
// Information about Extended Community Opaque
// sets
type RoutingPolicy_Sets_ExtendedCommunityOpaque struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_

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

    extendedCommunityOpaque.EntityData.Children = make(map[string]types.YChild)
    extendedCommunityOpaque.EntityData.Children["sets"] = types.YChild{"Sets", &extendedCommunityOpaque.Sets}
    extendedCommunityOpaque.EntityData.Children["unused"] = types.YChild{"Unused", &extendedCommunityOpaque.Unused}
    extendedCommunityOpaque.EntityData.Children["inactive"] = types.YChild{"Inactive", &extendedCommunityOpaque.Inactive}
    extendedCommunityOpaque.EntityData.Children["active"] = types.YChild{"Active", &extendedCommunityOpaque.Active}
    extendedCommunityOpaque.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(extendedCommunityOpaque.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set.
    Set []RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "extended-community-opaque"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh
// Information about Extended Community SegNH sets
type RoutingPolicy_Sets_ExtendedCommunitySegNh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_

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

    extendedCommunitySegNh.EntityData.Children = make(map[string]types.YChild)
    extendedCommunitySegNh.EntityData.Children["sets"] = types.YChild{"Sets", &extendedCommunitySegNh.Sets}
    extendedCommunitySegNh.EntityData.Children["unused"] = types.YChild{"Unused", &extendedCommunitySegNh.Unused}
    extendedCommunitySegNh.EntityData.Children["inactive"] = types.YChild{"Inactive", &extendedCommunitySegNh.Inactive}
    extendedCommunitySegNh.EntityData.Children["active"] = types.YChild{"Active", &extendedCommunitySegNh.Active}
    extendedCommunitySegNh.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(extendedCommunitySegNh.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set.
    Set []RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "extended-community-seg-nh"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo
// Information about Extended Community SOO sets
type RoutingPolicy_Sets_ExtendedCommunitySoo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_

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

    extendedCommunitySoo.EntityData.Children = make(map[string]types.YChild)
    extendedCommunitySoo.EntityData.Children["sets"] = types.YChild{"Sets", &extendedCommunitySoo.Sets}
    extendedCommunitySoo.EntityData.Children["unused"] = types.YChild{"Unused", &extendedCommunitySoo.Unused}
    extendedCommunitySoo.EntityData.Children["inactive"] = types.YChild{"Inactive", &extendedCommunitySoo.Inactive}
    extendedCommunitySoo.EntityData.Children["active"] = types.YChild{"Active", &extendedCommunitySoo.Active}
    extendedCommunitySoo.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(extendedCommunitySoo.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set.
    Set []RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "extended-community-soo"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_Tag
// Information about Tag sets
type RoutingPolicy_Sets_Tag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Tag_Sets_

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

    tag.EntityData.Children = make(map[string]types.YChild)
    tag.EntityData.Children["sets"] = types.YChild{"Sets", &tag.Sets}
    tag.EntityData.Children["unused"] = types.YChild{"Unused", &tag.Unused}
    tag.EntityData.Children["inactive"] = types.YChild{"Inactive", &tag.Inactive}
    tag.EntityData.Children["active"] = types.YChild{"Active", &tag.Active}
    tag.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tag.EntityData)
}

// RoutingPolicy_Sets_Tag_Sets_
// Information about individual sets
type RoutingPolicy_Sets_Tag_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Tag_Sets__Set.
    Set []RoutingPolicy_Sets_Tag_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_Tag_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "tag"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_Tag_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_Tag_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Tag_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Tag_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_Tag_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_Tag_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Tag_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Tag_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Tag_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Tag_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Tag_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Tag_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_Tag_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Tag_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Tag_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Tag_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Tag_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Tag_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Tag_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Tag_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_Tag_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_Prefix
// Information about AS Path sets
type RoutingPolicy_Sets_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Prefix_Sets_

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

    prefix.EntityData.Children = make(map[string]types.YChild)
    prefix.EntityData.Children["sets"] = types.YChild{"Sets", &prefix.Sets}
    prefix.EntityData.Children["unused"] = types.YChild{"Unused", &prefix.Unused}
    prefix.EntityData.Children["inactive"] = types.YChild{"Inactive", &prefix.Inactive}
    prefix.EntityData.Children["active"] = types.YChild{"Active", &prefix.Active}
    prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefix.EntityData)
}

// RoutingPolicy_Sets_Prefix_Sets_
// Information about individual sets
type RoutingPolicy_Sets_Prefix_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Prefix_Sets__Set.
    Set []RoutingPolicy_Sets_Prefix_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_Prefix_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "prefix"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_Prefix_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_Prefix_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Prefix_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Prefix_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_Prefix_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_Prefix_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Prefix_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Prefix_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Prefix_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Prefix_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Prefix_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Prefix_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_Prefix_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Prefix_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Prefix_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Prefix_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Prefix_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Prefix_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Prefix_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Prefix_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_Prefix_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_Community
// Information about Community sets
type RoutingPolicy_Sets_Community struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Community_Sets_

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

    community.EntityData.Children = make(map[string]types.YChild)
    community.EntityData.Children["sets"] = types.YChild{"Sets", &community.Sets}
    community.EntityData.Children["unused"] = types.YChild{"Unused", &community.Unused}
    community.EntityData.Children["inactive"] = types.YChild{"Inactive", &community.Inactive}
    community.EntityData.Children["active"] = types.YChild{"Active", &community.Active}
    community.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(community.EntityData)
}

// RoutingPolicy_Sets_Community_Sets_
// Information about individual sets
type RoutingPolicy_Sets_Community_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Community_Sets__Set.
    Set []RoutingPolicy_Sets_Community_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_Community_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "community"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_Community_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_Community_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Community_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Community_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_Community_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_Community_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Community_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Community_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Community_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Community_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Community_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Community_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_Community_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Community_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Community_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Community_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Community_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Community_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Community_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Community_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_Community_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_AsPath
// Information about AS Path sets
type RoutingPolicy_Sets_AsPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_AsPath_Sets_

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

    asPath.EntityData.Children = make(map[string]types.YChild)
    asPath.EntityData.Children["sets"] = types.YChild{"Sets", &asPath.Sets}
    asPath.EntityData.Children["unused"] = types.YChild{"Unused", &asPath.Unused}
    asPath.EntityData.Children["inactive"] = types.YChild{"Inactive", &asPath.Inactive}
    asPath.EntityData.Children["active"] = types.YChild{"Active", &asPath.Active}
    asPath.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asPath.EntityData)
}

// RoutingPolicy_Sets_AsPath_Sets_
// Information about individual sets
type RoutingPolicy_Sets_AsPath_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_AsPath_Sets__Set.
    Set []RoutingPolicy_Sets_AsPath_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_AsPath_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "as-path"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_AsPath_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_AsPath_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_AsPath_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_AsPath_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_AsPath_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_AsPath_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_AsPath_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_AsPath_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_AsPath_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_AsPath_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_AsPath_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_AsPath_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_AsPath_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_AsPath_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_AsPath_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_AsPath_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_AsPath_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_AsPath_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_AsPath_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_AsPath_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_AsPath_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity
// Information about Large Community sets
type RoutingPolicy_Sets_LargeCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_LargeCommunity_Sets_

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

    largeCommunity.EntityData.Children = make(map[string]types.YChild)
    largeCommunity.EntityData.Children["sets"] = types.YChild{"Sets", &largeCommunity.Sets}
    largeCommunity.EntityData.Children["unused"] = types.YChild{"Unused", &largeCommunity.Unused}
    largeCommunity.EntityData.Children["inactive"] = types.YChild{"Inactive", &largeCommunity.Inactive}
    largeCommunity.EntityData.Children["active"] = types.YChild{"Active", &largeCommunity.Active}
    largeCommunity.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(largeCommunity.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Sets_
// Information about individual sets
type RoutingPolicy_Sets_LargeCommunity_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_LargeCommunity_Sets__Set.
    Set []RoutingPolicy_Sets_LargeCommunity_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_LargeCommunity_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "large-community"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_LargeCommunity_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_LargeCommunity_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_LargeCommunity_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_LargeCommunity_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_LargeCommunity_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_LargeCommunity_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_LargeCommunity_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_LargeCommunity_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_LargeCommunity_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_LargeCommunity_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_LargeCommunity_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_LargeCommunity_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_LargeCommunity_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_Esi
// Information about Esi sets
type RoutingPolicy_Sets_Esi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Esi_Sets_

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

    esi.EntityData.Children = make(map[string]types.YChild)
    esi.EntityData.Children["sets"] = types.YChild{"Sets", &esi.Sets}
    esi.EntityData.Children["unused"] = types.YChild{"Unused", &esi.Unused}
    esi.EntityData.Children["inactive"] = types.YChild{"Inactive", &esi.Inactive}
    esi.EntityData.Children["active"] = types.YChild{"Active", &esi.Active}
    esi.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(esi.EntityData)
}

// RoutingPolicy_Sets_Esi_Sets_
// Information about individual sets
type RoutingPolicy_Sets_Esi_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Esi_Sets__Set.
    Set []RoutingPolicy_Sets_Esi_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_Esi_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "esi"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_Esi_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_Esi_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Esi_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Esi_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_Esi_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_Esi_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Esi_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Esi_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Esi_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Esi_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Esi_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Esi_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_Esi_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Esi_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Esi_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Esi_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Esi_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Esi_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Esi_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Esi_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_Esi_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth
// Information about Extended Community Bandwidth
// sets
type RoutingPolicy_Sets_ExtendedCommunityBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_

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

    extendedCommunityBandwidth.EntityData.Children = make(map[string]types.YChild)
    extendedCommunityBandwidth.EntityData.Children["sets"] = types.YChild{"Sets", &extendedCommunityBandwidth.Sets}
    extendedCommunityBandwidth.EntityData.Children["unused"] = types.YChild{"Unused", &extendedCommunityBandwidth.Unused}
    extendedCommunityBandwidth.EntityData.Children["inactive"] = types.YChild{"Inactive", &extendedCommunityBandwidth.Inactive}
    extendedCommunityBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(extendedCommunityBandwidth.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set.
    Set []RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "extended-community-bandwidth"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
    return &(inactive.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt
// Information about Extended Community RT sets
type RoutingPolicy_Sets_ExtendedCommunityRt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_ExtendedCommunityRt_Sets_

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

    extendedCommunityRt.EntityData.Children = make(map[string]types.YChild)
    extendedCommunityRt.EntityData.Children["sets"] = types.YChild{"Sets", &extendedCommunityRt.Sets}
    extendedCommunityRt.EntityData.Children["unused"] = types.YChild{"Unused", &extendedCommunityRt.Unused}
    extendedCommunityRt.EntityData.Children["inactive"] = types.YChild{"Inactive", &extendedCommunityRt.Inactive}
    extendedCommunityRt.EntityData.Children["active"] = types.YChild{"Active", &extendedCommunityRt.Active}
    extendedCommunityRt.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(extendedCommunityRt.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets_
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set.
    Set []RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "extended-community-rt"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_Rd
// Information about RD sets
type RoutingPolicy_Sets_Rd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Rd_Sets_

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

    rd.EntityData.Children = make(map[string]types.YChild)
    rd.EntityData.Children["sets"] = types.YChild{"Sets", &rd.Sets}
    rd.EntityData.Children["unused"] = types.YChild{"Unused", &rd.Unused}
    rd.EntityData.Children["inactive"] = types.YChild{"Inactive", &rd.Inactive}
    rd.EntityData.Children["active"] = types.YChild{"Active", &rd.Active}
    rd.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rd.EntityData)
}

// RoutingPolicy_Sets_Rd_Sets_
// Information about individual sets
type RoutingPolicy_Sets_Rd_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Rd_Sets__Set.
    Set []RoutingPolicy_Sets_Rd_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_Rd_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "rd"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_Rd_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_Rd_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Rd_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Rd_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_Rd_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_Rd_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Rd_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Rd_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Rd_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Rd_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Rd_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Rd_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_Rd_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Rd_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Rd_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Rd_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Rd_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Rd_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Rd_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Rd_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_Rd_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_Mac
// Information about Mac sets
type RoutingPolicy_Sets_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_Mac_Sets_

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

    mac.EntityData.Children = make(map[string]types.YChild)
    mac.EntityData.Children["sets"] = types.YChild{"Sets", &mac.Sets}
    mac.EntityData.Children["unused"] = types.YChild{"Unused", &mac.Unused}
    mac.EntityData.Children["inactive"] = types.YChild{"Inactive", &mac.Inactive}
    mac.EntityData.Children["active"] = types.YChild{"Active", &mac.Active}
    mac.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mac.EntityData)
}

// RoutingPolicy_Sets_Mac_Sets_
// Information about individual sets
type RoutingPolicy_Sets_Mac_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Mac_Sets__Set.
    Set []RoutingPolicy_Sets_Mac_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_Mac_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "mac"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_Mac_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_Mac_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Mac_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Mac_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_Mac_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_Mac_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Mac_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Mac_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Mac_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Mac_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_Mac_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Mac_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_Mac_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_Mac_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Mac_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Mac_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Mac_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Mac_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_Mac_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Mac_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_Mac_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost
// Information about Extended Community Cost sets
type RoutingPolicy_Sets_ExtendedCommunityCost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_ExtendedCommunityCost_Sets_

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

    extendedCommunityCost.EntityData.Children = make(map[string]types.YChild)
    extendedCommunityCost.EntityData.Children["sets"] = types.YChild{"Sets", &extendedCommunityCost.Sets}
    extendedCommunityCost.EntityData.Children["unused"] = types.YChild{"Unused", &extendedCommunityCost.Unused}
    extendedCommunityCost.EntityData.Children["inactive"] = types.YChild{"Inactive", &extendedCommunityCost.Inactive}
    extendedCommunityCost.EntityData.Children["active"] = types.YChild{"Active", &extendedCommunityCost.Active}
    extendedCommunityCost.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(extendedCommunityCost.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets_
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set.
    Set []RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set
}

func (sets_ *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_) GetEntityData() *types.CommonEntityData {
    sets_.EntityData.YFilter = sets_.YFilter
    sets_.EntityData.YangName = "sets"
    sets_.EntityData.BundleName = "cisco_ios_xr"
    sets_.EntityData.ParentYangName = "extended-community-cost"
    sets_.EntityData.SegmentPath = "sets"
    sets_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sets_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sets_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sets_.EntityData.Children = make(map[string]types.YChild)
    sets_.EntityData.Children["set"] = types.YChild{"Set", nil}
    for i := range sets_.Set {
        sets_.EntityData.Children[types.GetSegmentPath(&sets_.Set[i])] = types.YChild{"Set", &sets_.Set[i]}
    }
    sets_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sets_.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "sets"
    set.EntityData.SegmentPath = "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = make(map[string]types.YChild)
    set.EntityData.Children["used-by"] = types.YChild{"UsedBy", &set.UsedBy}
    set.EntityData.Children["attached"] = types.YChild{"Attached", &set.Attached}
    set.EntityData.Leafs = make(map[string]types.YLeaf)
    set.EntityData.Leafs["set-name"] = types.YLeaf{"SetName", set.SetName}
    return &(set.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_UsedBy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_UsedBy) GetEntityData() *types.CommonEntityData {
    usedBy.EntityData.YFilter = usedBy.YFilter
    usedBy.EntityData.YangName = "used-by"
    usedBy.EntityData.BundleName = "cisco_ios_xr"
    usedBy.EntityData.ParentYangName = "set"
    usedBy.EntityData.SegmentPath = "used-by"
    usedBy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usedBy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usedBy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usedBy.EntityData.Children = make(map[string]types.YChild)
    usedBy.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range usedBy.Reference {
        usedBy.EntityData.Children[types.GetSegmentPath(&usedBy.Reference[i])] = types.YChild{"Reference", &usedBy.Reference[i]}
    }
    usedBy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usedBy.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_UsedBy_Reference struct {
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

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_UsedBy_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "used-by"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", reference.RoutePolicyName}
    reference.EntityData.Leafs["used-directly"] = types.YLeaf{"UsedDirectly", reference.UsedDirectly}
    reference.EntityData.Leafs["status"] = types.YLeaf{"Status", reference.Status}
    return &(reference.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_Attached struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_Attached) GetEntityData() *types.CommonEntityData {
    attached.EntityData.YFilter = attached.YFilter
    attached.EntityData.YangName = "attached"
    attached.EntityData.BundleName = "cisco_ios_xr"
    attached.EntityData.ParentYangName = "set"
    attached.EntityData.SegmentPath = "attached"
    attached.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attached.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attached.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attached.EntityData.Children = make(map[string]types.YChild)
    attached.EntityData.Children["binding"] = types.YChild{"Binding", nil}
    for i := range attached.Binding {
        attached.EntityData.Children[types.GetSegmentPath(&attached.Binding[i])] = types.YChild{"Binding", &attached.Binding[i]}
    }
    attached.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attached.EntityData)
}

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_Attached_Binding struct {
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

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets__Set_Attached_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "attached"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", binding.Protocol}
    binding.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", binding.VrfName}
    binding.EntityData.Leafs["proto-instance"] = types.YLeaf{"ProtoInstance", binding.ProtoInstance}
    binding.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", binding.AfName}
    binding.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", binding.SafName}
    binding.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", binding.NeighborAddress}
    binding.EntityData.Leafs["neighbor-af-name"] = types.YLeaf{"NeighborAfName", binding.NeighborAfName}
    binding.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", binding.GroupName}
    binding.EntityData.Leafs["direction"] = types.YLeaf{"Direction", binding.Direction}
    binding.EntityData.Leafs["group"] = types.YLeaf{"Group", binding.Group}
    binding.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", binding.SourceProtocol}
    binding.EntityData.Leafs["aggregate-network-address"] = types.YLeaf{"AggregateNetworkAddress", binding.AggregateNetworkAddress}
    binding.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", binding.InterfaceName}
    binding.EntityData.Leafs["instance"] = types.YLeaf{"Instance", binding.Instance}
    binding.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", binding.AreaId}
    binding.EntityData.Leafs["propogate-from"] = types.YLeaf{"PropogateFrom", binding.PropogateFrom}
    binding.EntityData.Leafs["propogate-to"] = types.YLeaf{"PropogateTo", binding.PropogateTo}
    binding.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", binding.RoutePolicyName}
    binding.EntityData.Leafs["attached-policy"] = types.YLeaf{"AttachedPolicy", binding.AttachedPolicy}
    binding.EntityData.Leafs["attach-point"] = types.YLeaf{"AttachPoint", binding.AttachPoint}
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

    unused.EntityData.Children = make(map[string]types.YChild)
    unused.EntityData.Leafs = make(map[string]types.YLeaf)
    unused.EntityData.Leafs["object"] = types.YLeaf{"Object", unused.Object}
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

    inactive.EntityData.Children = make(map[string]types.YChild)
    inactive.EntityData.Leafs = make(map[string]types.YLeaf)
    inactive.EntityData.Leafs["object"] = types.YLeaf{"Object", inactive.Object}
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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    active.EntityData.Leafs["object"] = types.YLeaf{"Object", active.Object}
    return &(active.EntityData)
}

