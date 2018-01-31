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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about configured limits and the current values.
    Limits RoutingPolicy_Limits

    // Information about configured route policies.
    Policies RoutingPolicy_Policies

    // Information about configured sets.
    Sets RoutingPolicy_Sets
}

func (routingPolicy *RoutingPolicy) GetFilter() yfilter.YFilter { return routingPolicy.YFilter }

func (routingPolicy *RoutingPolicy) SetFilter(yf yfilter.YFilter) { routingPolicy.YFilter = yf }

func (routingPolicy *RoutingPolicy) GetGoName(yname string) string {
    if yname == "limits" { return "Limits" }
    if yname == "policies" { return "Policies" }
    if yname == "sets" { return "Sets" }
    return ""
}

func (routingPolicy *RoutingPolicy) GetSegmentPath() string {
    return "Cisco-IOS-XR-policy-repository-oper:routing-policy"
}

func (routingPolicy *RoutingPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "limits" {
        return &routingPolicy.Limits
    }
    if childYangName == "policies" {
        return &routingPolicy.Policies
    }
    if childYangName == "sets" {
        return &routingPolicy.Sets
    }
    return nil
}

func (routingPolicy *RoutingPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["limits"] = &routingPolicy.Limits
    children["policies"] = &routingPolicy.Policies
    children["sets"] = &routingPolicy.Sets
    return children
}

func (routingPolicy *RoutingPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
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

func (routingPolicy *RoutingPolicy) GetParentYangName() string { return "Cisco-IOS-XR-policy-repository-oper" }

// RoutingPolicy_Limits
// Information about configured limits and the
// current values
type RoutingPolicy_Limits struct {
    parent types.Entity
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

func (limits *RoutingPolicy_Limits) GetFilter() yfilter.YFilter { return limits.YFilter }

func (limits *RoutingPolicy_Limits) SetFilter(yf yfilter.YFilter) { limits.YFilter = yf }

func (limits *RoutingPolicy_Limits) GetGoName(yname string) string {
    if yname == "maximum-lines-of-policy" { return "MaximumLinesOfPolicy" }
    if yname == "current-lines-of-policy-limit" { return "CurrentLinesOfPolicyLimit" }
    if yname == "current-lines-of-policy-used" { return "CurrentLinesOfPolicyUsed" }
    if yname == "maximum-number-of-policies" { return "MaximumNumberOfPolicies" }
    if yname == "current-number-of-policies-limit" { return "CurrentNumberOfPoliciesLimit" }
    if yname == "current-number-of-policies-used" { return "CurrentNumberOfPoliciesUsed" }
    if yname == "compiled-policies-length" { return "CompiledPoliciesLength" }
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
    leafs["current-lines-of-policy-limit"] = limits.CurrentLinesOfPolicyLimit
    leafs["current-lines-of-policy-used"] = limits.CurrentLinesOfPolicyUsed
    leafs["maximum-number-of-policies"] = limits.MaximumNumberOfPolicies
    leafs["current-number-of-policies-limit"] = limits.CurrentNumberOfPoliciesLimit
    leafs["current-number-of-policies-used"] = limits.CurrentNumberOfPoliciesUsed
    leafs["compiled-policies-length"] = limits.CompiledPoliciesLength
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

// RoutingPolicy_Policies
// Information about configured route policies
type RoutingPolicy_Policies struct {
    parent types.Entity
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

func (policies *RoutingPolicy_Policies) GetFilter() yfilter.YFilter { return policies.YFilter }

func (policies *RoutingPolicy_Policies) SetFilter(yf yfilter.YFilter) { policies.YFilter = yf }

func (policies *RoutingPolicy_Policies) GetGoName(yname string) string {
    if yname == "route-policies" { return "RoutePolicies" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (policies *RoutingPolicy_Policies) GetSegmentPath() string {
    return "policies"
}

func (policies *RoutingPolicy_Policies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-policies" {
        return &policies.RoutePolicies
    }
    if childYangName == "unused" {
        return &policies.Unused
    }
    if childYangName == "inactive" {
        return &policies.Inactive
    }
    if childYangName == "active" {
        return &policies.Active
    }
    return nil
}

func (policies *RoutingPolicy_Policies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["route-policies"] = &policies.RoutePolicies
    children["unused"] = &policies.Unused
    children["inactive"] = &policies.Inactive
    children["active"] = &policies.Active
    return children
}

func (policies *RoutingPolicy_Policies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policies *RoutingPolicy_Policies) GetBundleName() string { return "cisco_ios_xr" }

func (policies *RoutingPolicy_Policies) GetYangName() string { return "policies" }

func (policies *RoutingPolicy_Policies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policies *RoutingPolicy_Policies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policies *RoutingPolicy_Policies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policies *RoutingPolicy_Policies) SetParent(parent types.Entity) { policies.parent = parent }

func (policies *RoutingPolicy_Policies) GetParent() types.Entity { return policies.parent }

func (policies *RoutingPolicy_Policies) GetParentYangName() string { return "routing-policy" }

// RoutingPolicy_Policies_RoutePolicies
// Information about individual policies
type RoutingPolicy_Policies_RoutePolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual policy. The type is slice of
    // RoutingPolicy_Policies_RoutePolicies_RoutePolicy.
    RoutePolicy []RoutingPolicy_Policies_RoutePolicies_RoutePolicy
}

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetFilter() yfilter.YFilter { return routePolicies.YFilter }

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) SetFilter(yf yfilter.YFilter) { routePolicies.YFilter = yf }

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetGoName(yname string) string {
    if yname == "route-policy" { return "RoutePolicy" }
    return ""
}

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetSegmentPath() string {
    return "route-policies"
}

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-policy" {
        for _, c := range routePolicies.RoutePolicy {
            if routePolicies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Policies_RoutePolicies_RoutePolicy{}
        routePolicies.RoutePolicy = append(routePolicies.RoutePolicy, child)
        return &routePolicies.RoutePolicy[len(routePolicies.RoutePolicy)-1]
    }
    return nil
}

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routePolicies.RoutePolicy {
        children[routePolicies.RoutePolicy[i].GetSegmentPath()] = &routePolicies.RoutePolicy[i]
    }
    return children
}

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetBundleName() string { return "cisco_ios_xr" }

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetYangName() string { return "route-policies" }

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) SetParent(parent types.Entity) { routePolicies.parent = parent }

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetParent() types.Entity { return routePolicies.parent }

func (routePolicies *RoutingPolicy_Policies_RoutePolicies) GetParentYangName() string { return "policies" }

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy
// Information about an individual policy
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy struct {
    parent types.Entity
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

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetFilter() yfilter.YFilter { return routePolicy.YFilter }

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) SetFilter(yf yfilter.YFilter) { routePolicy.YFilter = yf }

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "policy-uses" { return "PolicyUses" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetSegmentPath() string {
    return "route-policy" + "[route-policy-name='" + fmt.Sprintf("%v", routePolicy.RoutePolicyName) + "']"
}

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-uses" {
        return &routePolicy.PolicyUses
    }
    if childYangName == "used-by" {
        return &routePolicy.UsedBy
    }
    if childYangName == "attached" {
        return &routePolicy.Attached
    }
    return nil
}

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["policy-uses"] = &routePolicy.PolicyUses
    children["used-by"] = &routePolicy.UsedBy
    children["attached"] = &routePolicy.Attached
    return children
}

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = routePolicy.RoutePolicyName
    return leafs
}

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetYangName() string { return "route-policy" }

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) SetParent(parent types.Entity) { routePolicy.parent = parent }

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetParent() types.Entity { return routePolicy.parent }

func (routePolicy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy) GetParentYangName() string { return "route-policies" }

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses
// Information about which policies and sets
// this policy uses
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses struct {
    parent types.Entity
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

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetFilter() yfilter.YFilter { return policyUses.YFilter }

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) SetFilter(yf yfilter.YFilter) { policyUses.YFilter = yf }

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetGoName(yname string) string {
    if yname == "directly-used-policies" { return "DirectlyUsedPolicies" }
    if yname == "all-used-sets" { return "AllUsedSets" }
    if yname == "directly-used-sets" { return "DirectlyUsedSets" }
    if yname == "all-used-policies" { return "AllUsedPolicies" }
    return ""
}

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetSegmentPath() string {
    return "policy-uses"
}

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "directly-used-policies" {
        return &policyUses.DirectlyUsedPolicies
    }
    if childYangName == "all-used-sets" {
        return &policyUses.AllUsedSets
    }
    if childYangName == "directly-used-sets" {
        return &policyUses.DirectlyUsedSets
    }
    if childYangName == "all-used-policies" {
        return &policyUses.AllUsedPolicies
    }
    return nil
}

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["directly-used-policies"] = &policyUses.DirectlyUsedPolicies
    children["all-used-sets"] = &policyUses.AllUsedSets
    children["directly-used-sets"] = &policyUses.DirectlyUsedSets
    children["all-used-policies"] = &policyUses.AllUsedPolicies
    return children
}

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetBundleName() string { return "cisco_ios_xr" }

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetYangName() string { return "policy-uses" }

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) SetParent(parent types.Entity) { policyUses.parent = parent }

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetParent() types.Entity { return policyUses.parent }

func (policyUses *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses) GetParentYangName() string { return "route-policy" }

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies
// Policies that this policy uses directly
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetFilter() yfilter.YFilter { return directlyUsedPolicies.YFilter }

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) SetFilter(yf yfilter.YFilter) { directlyUsedPolicies.YFilter = yf }

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetSegmentPath() string {
    return "directly-used-policies"
}

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = directlyUsedPolicies.Object
    return leafs
}

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetBundleName() string { return "cisco_ios_xr" }

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetYangName() string { return "directly-used-policies" }

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) SetParent(parent types.Entity) { directlyUsedPolicies.parent = parent }

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetParent() types.Entity { return directlyUsedPolicies.parent }

func (directlyUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedPolicies) GetParentYangName() string { return "policy-uses" }

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets
// Sets used by this policy, or by policies
// that it uses
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of sets in several domains. The type is slice of
    // RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets.
    Sets []RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets
}

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetFilter() yfilter.YFilter { return allUsedSets.YFilter }

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) SetFilter(yf yfilter.YFilter) { allUsedSets.YFilter = yf }

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    return ""
}

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetSegmentPath() string {
    return "all-used-sets"
}

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        for _, c := range allUsedSets.Sets {
            if allUsedSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets{}
        allUsedSets.Sets = append(allUsedSets.Sets, child)
        return &allUsedSets.Sets[len(allUsedSets.Sets)-1]
    }
    return nil
}

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range allUsedSets.Sets {
        children[allUsedSets.Sets[i].GetSegmentPath()] = &allUsedSets.Sets[i]
    }
    return children
}

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetBundleName() string { return "cisco_ios_xr" }

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetYangName() string { return "all-used-sets" }

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) SetParent(parent types.Entity) { allUsedSets.parent = parent }

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetParent() types.Entity { return allUsedSets.parent }

func (allUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets) GetParentYangName() string { return "policy-uses" }

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets
// List of sets in several domains
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain of sets. The type is string.
    SetDomain interface{}

    // Names of sets in this domain. The type is slice of string.
    SetName []interface{}
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetGoName(yname string) string {
    if yname == "set-domain" { return "SetDomain" }
    if yname == "set-name" { return "SetName" }
    return ""
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-domain"] = sets.SetDomain
    leafs["set-name"] = sets.SetName
    return leafs
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedSets_Sets) GetParentYangName() string { return "all-used-sets" }

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets
// Sets that this policy uses directly
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of sets in several domains. The type is slice of
    // RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets.
    Sets []RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets
}

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetFilter() yfilter.YFilter { return directlyUsedSets.YFilter }

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) SetFilter(yf yfilter.YFilter) { directlyUsedSets.YFilter = yf }

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    return ""
}

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetSegmentPath() string {
    return "directly-used-sets"
}

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        for _, c := range directlyUsedSets.Sets {
            if directlyUsedSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets{}
        directlyUsedSets.Sets = append(directlyUsedSets.Sets, child)
        return &directlyUsedSets.Sets[len(directlyUsedSets.Sets)-1]
    }
    return nil
}

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range directlyUsedSets.Sets {
        children[directlyUsedSets.Sets[i].GetSegmentPath()] = &directlyUsedSets.Sets[i]
    }
    return children
}

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetBundleName() string { return "cisco_ios_xr" }

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetYangName() string { return "directly-used-sets" }

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) SetParent(parent types.Entity) { directlyUsedSets.parent = parent }

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetParent() types.Entity { return directlyUsedSets.parent }

func (directlyUsedSets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets) GetParentYangName() string { return "policy-uses" }

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets
// List of sets in several domains
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain of sets. The type is string.
    SetDomain interface{}

    // Names of sets in this domain. The type is slice of string.
    SetName []interface{}
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetGoName(yname string) string {
    if yname == "set-domain" { return "SetDomain" }
    if yname == "set-name" { return "SetName" }
    return ""
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-domain"] = sets.SetDomain
    leafs["set-name"] = sets.SetName
    return leafs
}

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_DirectlyUsedSets_Sets) GetParentYangName() string { return "directly-used-sets" }

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies
// Policies used by this policy, or by policies
// that it uses
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetFilter() yfilter.YFilter { return allUsedPolicies.YFilter }

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) SetFilter(yf yfilter.YFilter) { allUsedPolicies.YFilter = yf }

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetSegmentPath() string {
    return "all-used-policies"
}

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = allUsedPolicies.Object
    return leafs
}

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetBundleName() string { return "cisco_ios_xr" }

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetYangName() string { return "all-used-policies" }

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) SetParent(parent types.Entity) { allUsedPolicies.parent = parent }

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetParent() types.Entity { return allUsedPolicies.parent }

func (allUsedPolicies *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_PolicyUses_AllUsedPolicies) GetParentYangName() string { return "policy-uses" }

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference.
    Reference []RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy) GetParentYangName() string { return "route-policy" }

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding.
    Binding []RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding
}

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached) GetParentYangName() string { return "route-policy" }

// RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding
// bindings list
type RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Policies_RoutePolicies_RoutePolicy_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Policies_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Policies_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Policies_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Policies_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Policies_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Policies_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Policies_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Policies_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Policies_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Policies_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Policies_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Policies_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Policies_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Policies_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Policies_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Policies_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Policies_Unused) GetParentYangName() string { return "policies" }

// RoutingPolicy_Policies_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Policies_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Policies_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Policies_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Policies_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Policies_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Policies_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Policies_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Policies_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Policies_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Policies_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Policies_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Policies_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Policies_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Policies_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Policies_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Policies_Inactive) GetParentYangName() string { return "policies" }

// RoutingPolicy_Policies_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Policies_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Policies_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Policies_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Policies_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Policies_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Policies_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Policies_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Policies_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Policies_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Policies_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Policies_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Policies_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Policies_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Policies_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Policies_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Policies_Active) GetParentYangName() string { return "policies" }

// RoutingPolicy_Sets
// Information about configured sets
type RoutingPolicy_Sets struct {
    parent types.Entity
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

func (sets *RoutingPolicy_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets) GetGoName(yname string) string {
    if yname == "etag" { return "Etag" }
    if yname == "ospf-area" { return "OspfArea" }
    if yname == "extended-community-opaque" { return "ExtendedCommunityOpaque" }
    if yname == "extended-community-seg-nh" { return "ExtendedCommunitySegNh" }
    if yname == "extended-community-soo" { return "ExtendedCommunitySoo" }
    if yname == "tag" { return "Tag" }
    if yname == "prefix" { return "Prefix" }
    if yname == "community" { return "Community" }
    if yname == "as-path" { return "AsPath" }
    if yname == "large-community" { return "LargeCommunity" }
    if yname == "esi" { return "Esi" }
    if yname == "extended-community-bandwidth" { return "ExtendedCommunityBandwidth" }
    if yname == "extended-community-rt" { return "ExtendedCommunityRt" }
    if yname == "rd" { return "Rd" }
    if yname == "mac" { return "Mac" }
    if yname == "extended-community-cost" { return "ExtendedCommunityCost" }
    return ""
}

func (sets *RoutingPolicy_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "etag" {
        return &sets.Etag
    }
    if childYangName == "ospf-area" {
        return &sets.OspfArea
    }
    if childYangName == "extended-community-opaque" {
        return &sets.ExtendedCommunityOpaque
    }
    if childYangName == "extended-community-seg-nh" {
        return &sets.ExtendedCommunitySegNh
    }
    if childYangName == "extended-community-soo" {
        return &sets.ExtendedCommunitySoo
    }
    if childYangName == "tag" {
        return &sets.Tag
    }
    if childYangName == "prefix" {
        return &sets.Prefix
    }
    if childYangName == "community" {
        return &sets.Community
    }
    if childYangName == "as-path" {
        return &sets.AsPath
    }
    if childYangName == "large-community" {
        return &sets.LargeCommunity
    }
    if childYangName == "esi" {
        return &sets.Esi
    }
    if childYangName == "extended-community-bandwidth" {
        return &sets.ExtendedCommunityBandwidth
    }
    if childYangName == "extended-community-rt" {
        return &sets.ExtendedCommunityRt
    }
    if childYangName == "rd" {
        return &sets.Rd
    }
    if childYangName == "mac" {
        return &sets.Mac
    }
    if childYangName == "extended-community-cost" {
        return &sets.ExtendedCommunityCost
    }
    return nil
}

func (sets *RoutingPolicy_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["etag"] = &sets.Etag
    children["ospf-area"] = &sets.OspfArea
    children["extended-community-opaque"] = &sets.ExtendedCommunityOpaque
    children["extended-community-seg-nh"] = &sets.ExtendedCommunitySegNh
    children["extended-community-soo"] = &sets.ExtendedCommunitySoo
    children["tag"] = &sets.Tag
    children["prefix"] = &sets.Prefix
    children["community"] = &sets.Community
    children["as-path"] = &sets.AsPath
    children["large-community"] = &sets.LargeCommunity
    children["esi"] = &sets.Esi
    children["extended-community-bandwidth"] = &sets.ExtendedCommunityBandwidth
    children["extended-community-rt"] = &sets.ExtendedCommunityRt
    children["rd"] = &sets.Rd
    children["mac"] = &sets.Mac
    children["extended-community-cost"] = &sets.ExtendedCommunityCost
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

// RoutingPolicy_Sets_Etag
// Information about Etag sets
type RoutingPolicy_Sets_Etag struct {
    parent types.Entity
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

func (etag *RoutingPolicy_Sets_Etag) GetFilter() yfilter.YFilter { return etag.YFilter }

func (etag *RoutingPolicy_Sets_Etag) SetFilter(yf yfilter.YFilter) { etag.YFilter = yf }

func (etag *RoutingPolicy_Sets_Etag) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (etag *RoutingPolicy_Sets_Etag) GetSegmentPath() string {
    return "etag"
}

func (etag *RoutingPolicy_Sets_Etag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &etag.Sets
    }
    if childYangName == "unused" {
        return &etag.Unused
    }
    if childYangName == "inactive" {
        return &etag.Inactive
    }
    if childYangName == "active" {
        return &etag.Active
    }
    return nil
}

func (etag *RoutingPolicy_Sets_Etag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &etag.Sets
    children["unused"] = &etag.Unused
    children["inactive"] = &etag.Inactive
    children["active"] = &etag.Active
    return children
}

func (etag *RoutingPolicy_Sets_Etag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (etag *RoutingPolicy_Sets_Etag) GetBundleName() string { return "cisco_ios_xr" }

func (etag *RoutingPolicy_Sets_Etag) GetYangName() string { return "etag" }

func (etag *RoutingPolicy_Sets_Etag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (etag *RoutingPolicy_Sets_Etag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (etag *RoutingPolicy_Sets_Etag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (etag *RoutingPolicy_Sets_Etag) SetParent(parent types.Entity) { etag.parent = parent }

func (etag *RoutingPolicy_Sets_Etag) GetParent() types.Entity { return etag.parent }

func (etag *RoutingPolicy_Sets_Etag) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Etag_Sets
// Information about individual sets
type RoutingPolicy_Sets_Etag_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Etag_Sets_Set.
    Set []RoutingPolicy_Sets_Etag_Sets_Set
}

func (sets *RoutingPolicy_Sets_Etag_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_Etag_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_Etag_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_Etag_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_Etag_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Etag_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_Etag_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_Etag_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_Etag_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_Etag_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_Etag_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_Etag_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_Etag_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_Etag_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_Etag_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_Etag_Sets) GetParentYangName() string { return "etag" }

// RoutingPolicy_Sets_Etag_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Etag_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Etag_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Etag_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_Etag_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_Etag_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_Etag_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Etag_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Etag_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_Etag_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_Etag_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Etag_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_Etag_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_Etag_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_Etag_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Etag_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Etag_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_Etag_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_Etag_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_Etag_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_Etag_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_Etag_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_Etag_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_Etag_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_Etag_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_Etag_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_Etag_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_Etag_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_Etag_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_Etag_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_Etag_Unused) GetParentYangName() string { return "etag" }

// RoutingPolicy_Sets_Etag_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Etag_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_Etag_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_Etag_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_Etag_Inactive) GetParentYangName() string { return "etag" }

// RoutingPolicy_Sets_Etag_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Etag_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Etag_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_Etag_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_Etag_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_Etag_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_Etag_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_Etag_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_Etag_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_Etag_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_Etag_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_Etag_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_Etag_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_Etag_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_Etag_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_Etag_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_Etag_Active) GetParentYangName() string { return "etag" }

// RoutingPolicy_Sets_OspfArea
// Information about OSPF Area sets
type RoutingPolicy_Sets_OspfArea struct {
    parent types.Entity
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

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetFilter() yfilter.YFilter { return ospfArea.YFilter }

func (ospfArea *RoutingPolicy_Sets_OspfArea) SetFilter(yf yfilter.YFilter) { ospfArea.YFilter = yf }

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetSegmentPath() string {
    return "ospf-area"
}

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &ospfArea.Sets
    }
    if childYangName == "unused" {
        return &ospfArea.Unused
    }
    if childYangName == "inactive" {
        return &ospfArea.Inactive
    }
    if childYangName == "active" {
        return &ospfArea.Active
    }
    return nil
}

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &ospfArea.Sets
    children["unused"] = &ospfArea.Unused
    children["inactive"] = &ospfArea.Inactive
    children["active"] = &ospfArea.Active
    return children
}

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetBundleName() string { return "cisco_ios_xr" }

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetYangName() string { return "ospf-area" }

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfArea *RoutingPolicy_Sets_OspfArea) SetParent(parent types.Entity) { ospfArea.parent = parent }

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetParent() types.Entity { return ospfArea.parent }

func (ospfArea *RoutingPolicy_Sets_OspfArea) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_OspfArea_Sets
// Information about individual sets
type RoutingPolicy_Sets_OspfArea_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_OspfArea_Sets_Set.
    Set []RoutingPolicy_Sets_OspfArea_Sets_Set
}

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_OspfArea_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_OspfArea_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_OspfArea_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_OspfArea_Sets) GetParentYangName() string { return "ospf-area" }

// RoutingPolicy_Sets_OspfArea_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_OspfArea_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_OspfArea_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_OspfArea_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_OspfArea_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_OspfArea_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_OspfArea_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_OspfArea_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_OspfArea_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_OspfArea_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_OspfArea_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_OspfArea_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_OspfArea_Unused) GetParentYangName() string { return "ospf-area" }

// RoutingPolicy_Sets_OspfArea_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_OspfArea_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_OspfArea_Inactive) GetParentYangName() string { return "ospf-area" }

// RoutingPolicy_Sets_OspfArea_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_OspfArea_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_OspfArea_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_OspfArea_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_OspfArea_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_OspfArea_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_OspfArea_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_OspfArea_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_OspfArea_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_OspfArea_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_OspfArea_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_OspfArea_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_OspfArea_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_OspfArea_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_OspfArea_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_OspfArea_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_OspfArea_Active) GetParentYangName() string { return "ospf-area" }

// RoutingPolicy_Sets_ExtendedCommunityOpaque
// Information about Extended Community Opaque
// sets
type RoutingPolicy_Sets_ExtendedCommunityOpaque struct {
    parent types.Entity
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

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetFilter() yfilter.YFilter { return extendedCommunityOpaque.YFilter }

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) SetFilter(yf yfilter.YFilter) { extendedCommunityOpaque.YFilter = yf }

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetSegmentPath() string {
    return "extended-community-opaque"
}

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &extendedCommunityOpaque.Sets
    }
    if childYangName == "unused" {
        return &extendedCommunityOpaque.Unused
    }
    if childYangName == "inactive" {
        return &extendedCommunityOpaque.Inactive
    }
    if childYangName == "active" {
        return &extendedCommunityOpaque.Active
    }
    return nil
}

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &extendedCommunityOpaque.Sets
    children["unused"] = &extendedCommunityOpaque.Unused
    children["inactive"] = &extendedCommunityOpaque.Inactive
    children["active"] = &extendedCommunityOpaque.Active
    return children
}

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetYangName() string { return "extended-community-opaque" }

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) SetParent(parent types.Entity) { extendedCommunityOpaque.parent = parent }

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetParent() types.Entity { return extendedCommunityOpaque.parent }

func (extendedCommunityOpaque *RoutingPolicy_Sets_ExtendedCommunityOpaque) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set.
    Set []RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets) GetParentYangName() string { return "extended-community-opaque" }

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_ExtendedCommunityOpaque_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_ExtendedCommunityOpaque_Unused) GetParentYangName() string { return "extended-community-opaque" }

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityOpaque_Inactive) GetParentYangName() string { return "extended-community-opaque" }

// RoutingPolicy_Sets_ExtendedCommunityOpaque_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_ExtendedCommunityOpaque_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_ExtendedCommunityOpaque_Active) GetParentYangName() string { return "extended-community-opaque" }

// RoutingPolicy_Sets_ExtendedCommunitySegNh
// Information about Extended Community SegNH sets
type RoutingPolicy_Sets_ExtendedCommunitySegNh struct {
    parent types.Entity
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

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetFilter() yfilter.YFilter { return extendedCommunitySegNh.YFilter }

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) SetFilter(yf yfilter.YFilter) { extendedCommunitySegNh.YFilter = yf }

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetSegmentPath() string {
    return "extended-community-seg-nh"
}

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &extendedCommunitySegNh.Sets
    }
    if childYangName == "unused" {
        return &extendedCommunitySegNh.Unused
    }
    if childYangName == "inactive" {
        return &extendedCommunitySegNh.Inactive
    }
    if childYangName == "active" {
        return &extendedCommunitySegNh.Active
    }
    return nil
}

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &extendedCommunitySegNh.Sets
    children["unused"] = &extendedCommunitySegNh.Unused
    children["inactive"] = &extendedCommunitySegNh.Inactive
    children["active"] = &extendedCommunitySegNh.Active
    return children
}

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetYangName() string { return "extended-community-seg-nh" }

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) SetParent(parent types.Entity) { extendedCommunitySegNh.parent = parent }

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetParent() types.Entity { return extendedCommunitySegNh.parent }

func (extendedCommunitySegNh *RoutingPolicy_Sets_ExtendedCommunitySegNh) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set.
    Set []RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets) GetParentYangName() string { return "extended-community-seg-nh" }

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySegNh_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySegNh_Unused) GetParentYangName() string { return "extended-community-seg-nh" }

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySegNh_Inactive) GetParentYangName() string { return "extended-community-seg-nh" }

// RoutingPolicy_Sets_ExtendedCommunitySegNh_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_ExtendedCommunitySegNh_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_ExtendedCommunitySegNh_Active) GetParentYangName() string { return "extended-community-seg-nh" }

// RoutingPolicy_Sets_ExtendedCommunitySoo
// Information about Extended Community SOO sets
type RoutingPolicy_Sets_ExtendedCommunitySoo struct {
    parent types.Entity
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

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetFilter() yfilter.YFilter { return extendedCommunitySoo.YFilter }

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) SetFilter(yf yfilter.YFilter) { extendedCommunitySoo.YFilter = yf }

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetSegmentPath() string {
    return "extended-community-soo"
}

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &extendedCommunitySoo.Sets
    }
    if childYangName == "unused" {
        return &extendedCommunitySoo.Unused
    }
    if childYangName == "inactive" {
        return &extendedCommunitySoo.Inactive
    }
    if childYangName == "active" {
        return &extendedCommunitySoo.Active
    }
    return nil
}

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &extendedCommunitySoo.Sets
    children["unused"] = &extendedCommunitySoo.Unused
    children["inactive"] = &extendedCommunitySoo.Inactive
    children["active"] = &extendedCommunitySoo.Active
    return children
}

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetYangName() string { return "extended-community-soo" }

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) SetParent(parent types.Entity) { extendedCommunitySoo.parent = parent }

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetParent() types.Entity { return extendedCommunitySoo.parent }

func (extendedCommunitySoo *RoutingPolicy_Sets_ExtendedCommunitySoo) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set.
    Set []RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets) GetParentYangName() string { return "extended-community-soo" }

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_ExtendedCommunitySoo_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_ExtendedCommunitySoo_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_ExtendedCommunitySoo_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_ExtendedCommunitySoo_Unused) GetParentYangName() string { return "extended-community-soo" }

// RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_ExtendedCommunitySoo_Inactive) GetParentYangName() string { return "extended-community-soo" }

// RoutingPolicy_Sets_ExtendedCommunitySoo_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_ExtendedCommunitySoo_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_ExtendedCommunitySoo_Active) GetParentYangName() string { return "extended-community-soo" }

// RoutingPolicy_Sets_Tag
// Information about Tag sets
type RoutingPolicy_Sets_Tag struct {
    parent types.Entity
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

func (tag *RoutingPolicy_Sets_Tag) GetFilter() yfilter.YFilter { return tag.YFilter }

func (tag *RoutingPolicy_Sets_Tag) SetFilter(yf yfilter.YFilter) { tag.YFilter = yf }

func (tag *RoutingPolicy_Sets_Tag) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (tag *RoutingPolicy_Sets_Tag) GetSegmentPath() string {
    return "tag"
}

func (tag *RoutingPolicy_Sets_Tag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &tag.Sets
    }
    if childYangName == "unused" {
        return &tag.Unused
    }
    if childYangName == "inactive" {
        return &tag.Inactive
    }
    if childYangName == "active" {
        return &tag.Active
    }
    return nil
}

func (tag *RoutingPolicy_Sets_Tag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &tag.Sets
    children["unused"] = &tag.Unused
    children["inactive"] = &tag.Inactive
    children["active"] = &tag.Active
    return children
}

func (tag *RoutingPolicy_Sets_Tag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tag *RoutingPolicy_Sets_Tag) GetBundleName() string { return "cisco_ios_xr" }

func (tag *RoutingPolicy_Sets_Tag) GetYangName() string { return "tag" }

func (tag *RoutingPolicy_Sets_Tag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tag *RoutingPolicy_Sets_Tag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tag *RoutingPolicy_Sets_Tag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tag *RoutingPolicy_Sets_Tag) SetParent(parent types.Entity) { tag.parent = parent }

func (tag *RoutingPolicy_Sets_Tag) GetParent() types.Entity { return tag.parent }

func (tag *RoutingPolicy_Sets_Tag) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Tag_Sets
// Information about individual sets
type RoutingPolicy_Sets_Tag_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Tag_Sets_Set.
    Set []RoutingPolicy_Sets_Tag_Sets_Set
}

func (sets *RoutingPolicy_Sets_Tag_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_Tag_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_Tag_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_Tag_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_Tag_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Tag_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_Tag_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_Tag_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_Tag_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_Tag_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_Tag_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_Tag_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_Tag_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_Tag_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_Tag_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_Tag_Sets) GetParentYangName() string { return "tag" }

// RoutingPolicy_Sets_Tag_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Tag_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Tag_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Tag_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_Tag_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_Tag_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_Tag_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Tag_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Tag_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_Tag_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_Tag_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Tag_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_Tag_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_Tag_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_Tag_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Tag_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Tag_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_Tag_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_Tag_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_Tag_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_Tag_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_Tag_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_Tag_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_Tag_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_Tag_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_Tag_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_Tag_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_Tag_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_Tag_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_Tag_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_Tag_Unused) GetParentYangName() string { return "tag" }

// RoutingPolicy_Sets_Tag_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Tag_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_Tag_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_Tag_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_Tag_Inactive) GetParentYangName() string { return "tag" }

// RoutingPolicy_Sets_Tag_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Tag_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Tag_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_Tag_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_Tag_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_Tag_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_Tag_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_Tag_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_Tag_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_Tag_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_Tag_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_Tag_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_Tag_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_Tag_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_Tag_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_Tag_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_Tag_Active) GetParentYangName() string { return "tag" }

// RoutingPolicy_Sets_Prefix
// Information about AS Path sets
type RoutingPolicy_Sets_Prefix struct {
    parent types.Entity
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

func (prefix *RoutingPolicy_Sets_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *RoutingPolicy_Sets_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *RoutingPolicy_Sets_Prefix) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (prefix *RoutingPolicy_Sets_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *RoutingPolicy_Sets_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &prefix.Sets
    }
    if childYangName == "unused" {
        return &prefix.Unused
    }
    if childYangName == "inactive" {
        return &prefix.Inactive
    }
    if childYangName == "active" {
        return &prefix.Active
    }
    return nil
}

func (prefix *RoutingPolicy_Sets_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &prefix.Sets
    children["unused"] = &prefix.Unused
    children["inactive"] = &prefix.Inactive
    children["active"] = &prefix.Active
    return children
}

func (prefix *RoutingPolicy_Sets_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefix *RoutingPolicy_Sets_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *RoutingPolicy_Sets_Prefix) GetYangName() string { return "prefix" }

func (prefix *RoutingPolicy_Sets_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *RoutingPolicy_Sets_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *RoutingPolicy_Sets_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *RoutingPolicy_Sets_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *RoutingPolicy_Sets_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *RoutingPolicy_Sets_Prefix) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Prefix_Sets
// Information about individual sets
type RoutingPolicy_Sets_Prefix_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Prefix_Sets_Set.
    Set []RoutingPolicy_Sets_Prefix_Sets_Set
}

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_Prefix_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Prefix_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_Prefix_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_Prefix_Sets) GetParentYangName() string { return "prefix" }

// RoutingPolicy_Sets_Prefix_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Prefix_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Prefix_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_Prefix_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_Prefix_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_Prefix_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Prefix_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_Prefix_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_Prefix_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_Prefix_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Prefix_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_Prefix_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_Prefix_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_Prefix_Unused) GetParentYangName() string { return "prefix" }

// RoutingPolicy_Sets_Prefix_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Prefix_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_Prefix_Inactive) GetParentYangName() string { return "prefix" }

// RoutingPolicy_Sets_Prefix_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Prefix_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Prefix_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_Prefix_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_Prefix_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_Prefix_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_Prefix_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_Prefix_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_Prefix_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_Prefix_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_Prefix_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_Prefix_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_Prefix_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_Prefix_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_Prefix_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_Prefix_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_Prefix_Active) GetParentYangName() string { return "prefix" }

// RoutingPolicy_Sets_Community
// Information about Community sets
type RoutingPolicy_Sets_Community struct {
    parent types.Entity
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

func (community *RoutingPolicy_Sets_Community) GetFilter() yfilter.YFilter { return community.YFilter }

func (community *RoutingPolicy_Sets_Community) SetFilter(yf yfilter.YFilter) { community.YFilter = yf }

func (community *RoutingPolicy_Sets_Community) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (community *RoutingPolicy_Sets_Community) GetSegmentPath() string {
    return "community"
}

func (community *RoutingPolicy_Sets_Community) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &community.Sets
    }
    if childYangName == "unused" {
        return &community.Unused
    }
    if childYangName == "inactive" {
        return &community.Inactive
    }
    if childYangName == "active" {
        return &community.Active
    }
    return nil
}

func (community *RoutingPolicy_Sets_Community) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &community.Sets
    children["unused"] = &community.Unused
    children["inactive"] = &community.Inactive
    children["active"] = &community.Active
    return children
}

func (community *RoutingPolicy_Sets_Community) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (community *RoutingPolicy_Sets_Community) GetBundleName() string { return "cisco_ios_xr" }

func (community *RoutingPolicy_Sets_Community) GetYangName() string { return "community" }

func (community *RoutingPolicy_Sets_Community) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (community *RoutingPolicy_Sets_Community) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (community *RoutingPolicy_Sets_Community) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (community *RoutingPolicy_Sets_Community) SetParent(parent types.Entity) { community.parent = parent }

func (community *RoutingPolicy_Sets_Community) GetParent() types.Entity { return community.parent }

func (community *RoutingPolicy_Sets_Community) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Community_Sets
// Information about individual sets
type RoutingPolicy_Sets_Community_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Community_Sets_Set.
    Set []RoutingPolicy_Sets_Community_Sets_Set
}

func (sets *RoutingPolicy_Sets_Community_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_Community_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_Community_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_Community_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_Community_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Community_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_Community_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_Community_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_Community_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_Community_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_Community_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_Community_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_Community_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_Community_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_Community_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_Community_Sets) GetParentYangName() string { return "community" }

// RoutingPolicy_Sets_Community_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Community_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Community_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Community_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_Community_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_Community_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_Community_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Community_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Community_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_Community_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_Community_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_Community_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Community_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_Community_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_Community_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_Community_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Community_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Community_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_Community_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_Community_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_Community_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_Community_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_Community_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_Community_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_Community_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_Community_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_Community_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_Community_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_Community_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_Community_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_Community_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_Community_Unused) GetParentYangName() string { return "community" }

// RoutingPolicy_Sets_Community_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Community_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_Community_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_Community_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_Community_Inactive) GetParentYangName() string { return "community" }

// RoutingPolicy_Sets_Community_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Community_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Community_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_Community_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_Community_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_Community_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_Community_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_Community_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_Community_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_Community_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_Community_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_Community_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_Community_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_Community_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_Community_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_Community_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_Community_Active) GetParentYangName() string { return "community" }

// RoutingPolicy_Sets_AsPath
// Information about AS Path sets
type RoutingPolicy_Sets_AsPath struct {
    parent types.Entity
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

func (asPath *RoutingPolicy_Sets_AsPath) GetFilter() yfilter.YFilter { return asPath.YFilter }

func (asPath *RoutingPolicy_Sets_AsPath) SetFilter(yf yfilter.YFilter) { asPath.YFilter = yf }

func (asPath *RoutingPolicy_Sets_AsPath) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (asPath *RoutingPolicy_Sets_AsPath) GetSegmentPath() string {
    return "as-path"
}

func (asPath *RoutingPolicy_Sets_AsPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &asPath.Sets
    }
    if childYangName == "unused" {
        return &asPath.Unused
    }
    if childYangName == "inactive" {
        return &asPath.Inactive
    }
    if childYangName == "active" {
        return &asPath.Active
    }
    return nil
}

func (asPath *RoutingPolicy_Sets_AsPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &asPath.Sets
    children["unused"] = &asPath.Unused
    children["inactive"] = &asPath.Inactive
    children["active"] = &asPath.Active
    return children
}

func (asPath *RoutingPolicy_Sets_AsPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asPath *RoutingPolicy_Sets_AsPath) GetBundleName() string { return "cisco_ios_xr" }

func (asPath *RoutingPolicy_Sets_AsPath) GetYangName() string { return "as-path" }

func (asPath *RoutingPolicy_Sets_AsPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asPath *RoutingPolicy_Sets_AsPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asPath *RoutingPolicy_Sets_AsPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asPath *RoutingPolicy_Sets_AsPath) SetParent(parent types.Entity) { asPath.parent = parent }

func (asPath *RoutingPolicy_Sets_AsPath) GetParent() types.Entity { return asPath.parent }

func (asPath *RoutingPolicy_Sets_AsPath) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_AsPath_Sets
// Information about individual sets
type RoutingPolicy_Sets_AsPath_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_AsPath_Sets_Set.
    Set []RoutingPolicy_Sets_AsPath_Sets_Set
}

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_AsPath_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_AsPath_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_AsPath_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_AsPath_Sets) GetParentYangName() string { return "as-path" }

// RoutingPolicy_Sets_AsPath_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_AsPath_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_AsPath_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_AsPath_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_AsPath_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_AsPath_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_AsPath_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_AsPath_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_AsPath_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_AsPath_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_AsPath_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_AsPath_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_AsPath_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_AsPath_Unused) GetParentYangName() string { return "as-path" }

// RoutingPolicy_Sets_AsPath_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_AsPath_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_AsPath_Inactive) GetParentYangName() string { return "as-path" }

// RoutingPolicy_Sets_AsPath_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_AsPath_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_AsPath_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_AsPath_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_AsPath_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_AsPath_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_AsPath_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_AsPath_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_AsPath_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_AsPath_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_AsPath_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_AsPath_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_AsPath_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_AsPath_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_AsPath_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_AsPath_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_AsPath_Active) GetParentYangName() string { return "as-path" }

// RoutingPolicy_Sets_LargeCommunity
// Information about Large Community sets
type RoutingPolicy_Sets_LargeCommunity struct {
    parent types.Entity
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

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetFilter() yfilter.YFilter { return largeCommunity.YFilter }

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) SetFilter(yf yfilter.YFilter) { largeCommunity.YFilter = yf }

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetSegmentPath() string {
    return "large-community"
}

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &largeCommunity.Sets
    }
    if childYangName == "unused" {
        return &largeCommunity.Unused
    }
    if childYangName == "inactive" {
        return &largeCommunity.Inactive
    }
    if childYangName == "active" {
        return &largeCommunity.Active
    }
    return nil
}

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &largeCommunity.Sets
    children["unused"] = &largeCommunity.Unused
    children["inactive"] = &largeCommunity.Inactive
    children["active"] = &largeCommunity.Active
    return children
}

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetYangName() string { return "large-community" }

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) SetParent(parent types.Entity) { largeCommunity.parent = parent }

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetParent() types.Entity { return largeCommunity.parent }

func (largeCommunity *RoutingPolicy_Sets_LargeCommunity) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_LargeCommunity_Sets
// Information about individual sets
type RoutingPolicy_Sets_LargeCommunity_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_LargeCommunity_Sets_Set.
    Set []RoutingPolicy_Sets_LargeCommunity_Sets_Set
}

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_LargeCommunity_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_LargeCommunity_Sets) GetParentYangName() string { return "large-community" }

// RoutingPolicy_Sets_LargeCommunity_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_LargeCommunity_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_LargeCommunity_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_LargeCommunity_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_LargeCommunity_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_LargeCommunity_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_LargeCommunity_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_LargeCommunity_Unused) GetParentYangName() string { return "large-community" }

// RoutingPolicy_Sets_LargeCommunity_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_LargeCommunity_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_LargeCommunity_Inactive) GetParentYangName() string { return "large-community" }

// RoutingPolicy_Sets_LargeCommunity_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_LargeCommunity_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_LargeCommunity_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_LargeCommunity_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_LargeCommunity_Active) GetParentYangName() string { return "large-community" }

// RoutingPolicy_Sets_Esi
// Information about Esi sets
type RoutingPolicy_Sets_Esi struct {
    parent types.Entity
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

func (esi *RoutingPolicy_Sets_Esi) GetFilter() yfilter.YFilter { return esi.YFilter }

func (esi *RoutingPolicy_Sets_Esi) SetFilter(yf yfilter.YFilter) { esi.YFilter = yf }

func (esi *RoutingPolicy_Sets_Esi) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (esi *RoutingPolicy_Sets_Esi) GetSegmentPath() string {
    return "esi"
}

func (esi *RoutingPolicy_Sets_Esi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &esi.Sets
    }
    if childYangName == "unused" {
        return &esi.Unused
    }
    if childYangName == "inactive" {
        return &esi.Inactive
    }
    if childYangName == "active" {
        return &esi.Active
    }
    return nil
}

func (esi *RoutingPolicy_Sets_Esi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &esi.Sets
    children["unused"] = &esi.Unused
    children["inactive"] = &esi.Inactive
    children["active"] = &esi.Active
    return children
}

func (esi *RoutingPolicy_Sets_Esi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (esi *RoutingPolicy_Sets_Esi) GetBundleName() string { return "cisco_ios_xr" }

func (esi *RoutingPolicy_Sets_Esi) GetYangName() string { return "esi" }

func (esi *RoutingPolicy_Sets_Esi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esi *RoutingPolicy_Sets_Esi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esi *RoutingPolicy_Sets_Esi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esi *RoutingPolicy_Sets_Esi) SetParent(parent types.Entity) { esi.parent = parent }

func (esi *RoutingPolicy_Sets_Esi) GetParent() types.Entity { return esi.parent }

func (esi *RoutingPolicy_Sets_Esi) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Esi_Sets
// Information about individual sets
type RoutingPolicy_Sets_Esi_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Esi_Sets_Set.
    Set []RoutingPolicy_Sets_Esi_Sets_Set
}

func (sets *RoutingPolicy_Sets_Esi_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_Esi_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_Esi_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_Esi_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_Esi_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Esi_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_Esi_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_Esi_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_Esi_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_Esi_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_Esi_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_Esi_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_Esi_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_Esi_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_Esi_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_Esi_Sets) GetParentYangName() string { return "esi" }

// RoutingPolicy_Sets_Esi_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Esi_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Esi_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Esi_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_Esi_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_Esi_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_Esi_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Esi_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Esi_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_Esi_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_Esi_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Esi_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_Esi_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_Esi_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_Esi_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Esi_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Esi_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_Esi_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_Esi_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_Esi_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_Esi_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_Esi_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_Esi_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_Esi_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_Esi_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_Esi_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_Esi_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_Esi_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_Esi_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_Esi_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_Esi_Unused) GetParentYangName() string { return "esi" }

// RoutingPolicy_Sets_Esi_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Esi_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_Esi_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_Esi_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_Esi_Inactive) GetParentYangName() string { return "esi" }

// RoutingPolicy_Sets_Esi_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Esi_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Esi_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_Esi_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_Esi_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_Esi_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_Esi_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_Esi_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_Esi_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_Esi_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_Esi_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_Esi_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_Esi_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_Esi_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_Esi_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_Esi_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_Esi_Active) GetParentYangName() string { return "esi" }

// RoutingPolicy_Sets_ExtendedCommunityBandwidth
// Information about Extended Community Bandwidth
// sets
type RoutingPolicy_Sets_ExtendedCommunityBandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about individual sets.
    Sets RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets

    // All objects of a given type that are not referenced at all.
    Unused RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused

    // All objects of a given type that are not attached to a protocol.
    Inactive RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive
}

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetFilter() yfilter.YFilter { return extendedCommunityBandwidth.YFilter }

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) SetFilter(yf yfilter.YFilter) { extendedCommunityBandwidth.YFilter = yf }

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    return ""
}

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetSegmentPath() string {
    return "extended-community-bandwidth"
}

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &extendedCommunityBandwidth.Sets
    }
    if childYangName == "unused" {
        return &extendedCommunityBandwidth.Unused
    }
    if childYangName == "inactive" {
        return &extendedCommunityBandwidth.Inactive
    }
    return nil
}

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &extendedCommunityBandwidth.Sets
    children["unused"] = &extendedCommunityBandwidth.Unused
    children["inactive"] = &extendedCommunityBandwidth.Inactive
    return children
}

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetYangName() string { return "extended-community-bandwidth" }

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) SetParent(parent types.Entity) { extendedCommunityBandwidth.parent = parent }

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetParent() types.Entity { return extendedCommunityBandwidth.parent }

func (extendedCommunityBandwidth *RoutingPolicy_Sets_ExtendedCommunityBandwidth) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set.
    Set []RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets) GetParentYangName() string { return "extended-community-bandwidth" }

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Unused) GetParentYangName() string { return "extended-community-bandwidth" }

// RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityBandwidth_Inactive) GetParentYangName() string { return "extended-community-bandwidth" }

// RoutingPolicy_Sets_ExtendedCommunityRt
// Information about Extended Community RT sets
type RoutingPolicy_Sets_ExtendedCommunityRt struct {
    parent types.Entity
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

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetFilter() yfilter.YFilter { return extendedCommunityRt.YFilter }

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) SetFilter(yf yfilter.YFilter) { extendedCommunityRt.YFilter = yf }

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetSegmentPath() string {
    return "extended-community-rt"
}

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &extendedCommunityRt.Sets
    }
    if childYangName == "unused" {
        return &extendedCommunityRt.Unused
    }
    if childYangName == "inactive" {
        return &extendedCommunityRt.Inactive
    }
    if childYangName == "active" {
        return &extendedCommunityRt.Active
    }
    return nil
}

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &extendedCommunityRt.Sets
    children["unused"] = &extendedCommunityRt.Unused
    children["inactive"] = &extendedCommunityRt.Inactive
    children["active"] = &extendedCommunityRt.Active
    return children
}

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetYangName() string { return "extended-community-rt" }

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) SetParent(parent types.Entity) { extendedCommunityRt.parent = parent }

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetParent() types.Entity { return extendedCommunityRt.parent }

func (extendedCommunityRt *RoutingPolicy_Sets_ExtendedCommunityRt) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set.
    Set []RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_ExtendedCommunityRt_Sets) GetParentYangName() string { return "extended-community-rt" }

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_ExtendedCommunityRt_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_ExtendedCommunityRt_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_ExtendedCommunityRt_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_ExtendedCommunityRt_Unused) GetParentYangName() string { return "extended-community-rt" }

// RoutingPolicy_Sets_ExtendedCommunityRt_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_ExtendedCommunityRt_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityRt_Inactive) GetParentYangName() string { return "extended-community-rt" }

// RoutingPolicy_Sets_ExtendedCommunityRt_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_ExtendedCommunityRt_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_ExtendedCommunityRt_Active) GetParentYangName() string { return "extended-community-rt" }

// RoutingPolicy_Sets_Rd
// Information about RD sets
type RoutingPolicy_Sets_Rd struct {
    parent types.Entity
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

func (rd *RoutingPolicy_Sets_Rd) GetFilter() yfilter.YFilter { return rd.YFilter }

func (rd *RoutingPolicy_Sets_Rd) SetFilter(yf yfilter.YFilter) { rd.YFilter = yf }

func (rd *RoutingPolicy_Sets_Rd) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (rd *RoutingPolicy_Sets_Rd) GetSegmentPath() string {
    return "rd"
}

func (rd *RoutingPolicy_Sets_Rd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &rd.Sets
    }
    if childYangName == "unused" {
        return &rd.Unused
    }
    if childYangName == "inactive" {
        return &rd.Inactive
    }
    if childYangName == "active" {
        return &rd.Active
    }
    return nil
}

func (rd *RoutingPolicy_Sets_Rd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &rd.Sets
    children["unused"] = &rd.Unused
    children["inactive"] = &rd.Inactive
    children["active"] = &rd.Active
    return children
}

func (rd *RoutingPolicy_Sets_Rd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rd *RoutingPolicy_Sets_Rd) GetBundleName() string { return "cisco_ios_xr" }

func (rd *RoutingPolicy_Sets_Rd) GetYangName() string { return "rd" }

func (rd *RoutingPolicy_Sets_Rd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rd *RoutingPolicy_Sets_Rd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rd *RoutingPolicy_Sets_Rd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rd *RoutingPolicy_Sets_Rd) SetParent(parent types.Entity) { rd.parent = parent }

func (rd *RoutingPolicy_Sets_Rd) GetParent() types.Entity { return rd.parent }

func (rd *RoutingPolicy_Sets_Rd) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Rd_Sets
// Information about individual sets
type RoutingPolicy_Sets_Rd_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Rd_Sets_Set.
    Set []RoutingPolicy_Sets_Rd_Sets_Set
}

func (sets *RoutingPolicy_Sets_Rd_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_Rd_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_Rd_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_Rd_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_Rd_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Rd_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_Rd_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_Rd_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_Rd_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_Rd_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_Rd_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_Rd_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_Rd_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_Rd_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_Rd_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_Rd_Sets) GetParentYangName() string { return "rd" }

// RoutingPolicy_Sets_Rd_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Rd_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Rd_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Rd_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_Rd_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_Rd_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_Rd_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Rd_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Rd_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_Rd_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_Rd_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Rd_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_Rd_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_Rd_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_Rd_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Rd_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Rd_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_Rd_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_Rd_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_Rd_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_Rd_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_Rd_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_Rd_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_Rd_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_Rd_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_Rd_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_Rd_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_Rd_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_Rd_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_Rd_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_Rd_Unused) GetParentYangName() string { return "rd" }

// RoutingPolicy_Sets_Rd_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Rd_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_Rd_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_Rd_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_Rd_Inactive) GetParentYangName() string { return "rd" }

// RoutingPolicy_Sets_Rd_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Rd_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Rd_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_Rd_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_Rd_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_Rd_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_Rd_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_Rd_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_Rd_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_Rd_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_Rd_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_Rd_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_Rd_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_Rd_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_Rd_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_Rd_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_Rd_Active) GetParentYangName() string { return "rd" }

// RoutingPolicy_Sets_Mac
// Information about Mac sets
type RoutingPolicy_Sets_Mac struct {
    parent types.Entity
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

func (mac *RoutingPolicy_Sets_Mac) GetFilter() yfilter.YFilter { return mac.YFilter }

func (mac *RoutingPolicy_Sets_Mac) SetFilter(yf yfilter.YFilter) { mac.YFilter = yf }

func (mac *RoutingPolicy_Sets_Mac) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (mac *RoutingPolicy_Sets_Mac) GetSegmentPath() string {
    return "mac"
}

func (mac *RoutingPolicy_Sets_Mac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &mac.Sets
    }
    if childYangName == "unused" {
        return &mac.Unused
    }
    if childYangName == "inactive" {
        return &mac.Inactive
    }
    if childYangName == "active" {
        return &mac.Active
    }
    return nil
}

func (mac *RoutingPolicy_Sets_Mac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &mac.Sets
    children["unused"] = &mac.Unused
    children["inactive"] = &mac.Inactive
    children["active"] = &mac.Active
    return children
}

func (mac *RoutingPolicy_Sets_Mac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mac *RoutingPolicy_Sets_Mac) GetBundleName() string { return "cisco_ios_xr" }

func (mac *RoutingPolicy_Sets_Mac) GetYangName() string { return "mac" }

func (mac *RoutingPolicy_Sets_Mac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mac *RoutingPolicy_Sets_Mac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mac *RoutingPolicy_Sets_Mac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mac *RoutingPolicy_Sets_Mac) SetParent(parent types.Entity) { mac.parent = parent }

func (mac *RoutingPolicy_Sets_Mac) GetParent() types.Entity { return mac.parent }

func (mac *RoutingPolicy_Sets_Mac) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Mac_Sets
// Information about individual sets
type RoutingPolicy_Sets_Mac_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_Mac_Sets_Set.
    Set []RoutingPolicy_Sets_Mac_Sets_Set
}

func (sets *RoutingPolicy_Sets_Mac_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_Mac_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_Mac_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_Mac_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_Mac_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Mac_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_Mac_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_Mac_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_Mac_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_Mac_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_Mac_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_Mac_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_Mac_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_Mac_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_Mac_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_Mac_Sets) GetParentYangName() string { return "mac" }

// RoutingPolicy_Sets_Mac_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_Mac_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_Mac_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_Mac_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_Mac_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_Mac_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_Mac_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_Mac_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_Mac_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_Mac_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_Mac_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_Mac_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_Mac_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_Mac_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_Mac_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_Mac_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_Mac_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_Mac_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_Mac_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_Mac_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_Mac_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_Mac_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_Mac_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_Mac_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_Mac_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_Mac_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_Mac_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_Mac_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_Mac_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_Mac_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_Mac_Unused) GetParentYangName() string { return "mac" }

// RoutingPolicy_Sets_Mac_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_Mac_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_Mac_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_Mac_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_Mac_Inactive) GetParentYangName() string { return "mac" }

// RoutingPolicy_Sets_Mac_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_Mac_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_Mac_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_Mac_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_Mac_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_Mac_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_Mac_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_Mac_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_Mac_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_Mac_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_Mac_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_Mac_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_Mac_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_Mac_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_Mac_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_Mac_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_Mac_Active) GetParentYangName() string { return "mac" }

// RoutingPolicy_Sets_ExtendedCommunityCost
// Information about Extended Community Cost sets
type RoutingPolicy_Sets_ExtendedCommunityCost struct {
    parent types.Entity
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

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetFilter() yfilter.YFilter { return extendedCommunityCost.YFilter }

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) SetFilter(yf yfilter.YFilter) { extendedCommunityCost.YFilter = yf }

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetGoName(yname string) string {
    if yname == "sets" { return "Sets" }
    if yname == "unused" { return "Unused" }
    if yname == "inactive" { return "Inactive" }
    if yname == "active" { return "Active" }
    return ""
}

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetSegmentPath() string {
    return "extended-community-cost"
}

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sets" {
        return &extendedCommunityCost.Sets
    }
    if childYangName == "unused" {
        return &extendedCommunityCost.Unused
    }
    if childYangName == "inactive" {
        return &extendedCommunityCost.Inactive
    }
    if childYangName == "active" {
        return &extendedCommunityCost.Active
    }
    return nil
}

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sets"] = &extendedCommunityCost.Sets
    children["unused"] = &extendedCommunityCost.Unused
    children["inactive"] = &extendedCommunityCost.Inactive
    children["active"] = &extendedCommunityCost.Active
    return children
}

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetBundleName() string { return "cisco_ios_xr" }

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetYangName() string { return "extended-community-cost" }

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) SetParent(parent types.Entity) { extendedCommunityCost.parent = parent }

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetParent() types.Entity { return extendedCommunityCost.parent }

func (extendedCommunityCost *RoutingPolicy_Sets_ExtendedCommunityCost) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets
// Information about individual sets
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an individual set. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set.
    Set []RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetFilter() yfilter.YFilter { return sets.YFilter }

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) SetFilter(yf yfilter.YFilter) { sets.YFilter = yf }

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetGoName(yname string) string {
    if yname == "set" { return "Set" }
    return ""
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetSegmentPath() string {
    return "sets"
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        for _, c := range sets.Set {
            if sets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set{}
        sets.Set = append(sets.Set, child)
        return &sets.Set[len(sets.Set)-1]
    }
    return nil
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sets.Set {
        children[sets.Set[i].GetSegmentPath()] = &sets.Set[i]
    }
    return children
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetBundleName() string { return "cisco_ios_xr" }

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetYangName() string { return "sets" }

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) SetParent(parent types.Entity) { sets.parent = parent }

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetParent() types.Entity { return sets.parent }

func (sets *RoutingPolicy_Sets_ExtendedCommunityCost_Sets) GetParentYangName() string { return "extended-community-cost" }

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set
// Information about an individual set
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SetName interface{}

    // Policies that use this object, directly or indirectly.
    UsedBy RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy

    // Information about where this policy or set is attached.
    Attached RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached
}

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetGoName(yname string) string {
    if yname == "set-name" { return "SetName" }
    if yname == "used-by" { return "UsedBy" }
    if yname == "attached" { return "Attached" }
    return ""
}

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetSegmentPath() string {
    return "set" + "[set-name='" + fmt.Sprintf("%v", set.SetName) + "']"
}

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "used-by" {
        return &set.UsedBy
    }
    if childYangName == "attached" {
        return &set.Attached
    }
    return nil
}

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["used-by"] = &set.UsedBy
    children["attached"] = &set.Attached
    return children
}

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-name"] = set.SetName
    return leafs
}

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetYangName() string { return "set" }

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetParent() types.Entity { return set.parent }

func (set *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set) GetParentYangName() string { return "sets" }

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy
// Policies that use this object, directly or
// indirectly
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about policies referring to this object. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference.
    Reference []RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetFilter() yfilter.YFilter { return usedBy.YFilter }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) SetFilter(yf yfilter.YFilter) { usedBy.YFilter = yf }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetSegmentPath() string {
    return "used-by"
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range usedBy.Reference {
            if usedBy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference{}
        usedBy.Reference = append(usedBy.Reference, child)
        return &usedBy.Reference[len(usedBy.Reference)-1]
    }
    return nil
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usedBy.Reference {
        children[usedBy.Reference[i].GetSegmentPath()] = &usedBy.Reference[i]
    }
    return children
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetBundleName() string { return "cisco_ios_xr" }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetYangName() string { return "used-by" }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) SetParent(parent types.Entity) { usedBy.parent = parent }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetParent() types.Entity { return usedBy.parent }

func (usedBy *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference
// Information about policies referring to this
// object
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of policy. The type is string.
    RoutePolicyName interface{}

    // Whether the policy uses this object directly or indirectly. The type is
    // bool.
    UsedDirectly interface{}

    // Active, Inactive, or Unused. The type is ObjectStatus.
    Status interface{}
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "used-directly" { return "UsedDirectly" }
    if yname == "status" { return "Status" }
    return ""
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = reference.RoutePolicyName
    leafs["used-directly"] = reference.UsedDirectly
    leafs["status"] = reference.Status
    return leafs
}

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_UsedBy_Reference) GetParentYangName() string { return "used-by" }

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached
// Information about where this policy or set is
// attached
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bindings list. The type is slice of
    // RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding.
    Binding []RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetFilter() yfilter.YFilter { return attached.YFilter }

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) SetFilter(yf yfilter.YFilter) { attached.YFilter = yf }

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetSegmentPath() string {
    return "attached"
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range attached.Binding {
            if attached.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding{}
        attached.Binding = append(attached.Binding, child)
        return &attached.Binding[len(attached.Binding)-1]
    }
    return nil
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attached.Binding {
        children[attached.Binding[i].GetSegmentPath()] = &attached.Binding[i]
    }
    return children
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetBundleName() string { return "cisco_ios_xr" }

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetYangName() string { return "attached" }

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) SetParent(parent types.Entity) { attached.parent = parent }

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetParent() types.Entity { return attached.parent }

func (attached *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached) GetParentYangName() string { return "set" }

// RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding
// bindings list
type RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding struct {
    parent types.Entity
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

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetGoName(yname string) string {
    if yname == "protocol" { return "Protocol" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "proto-instance" { return "ProtoInstance" }
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "neighbor-af-name" { return "NeighborAfName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "direction" { return "Direction" }
    if yname == "group" { return "Group" }
    if yname == "source-protocol" { return "SourceProtocol" }
    if yname == "aggregate-network-address" { return "AggregateNetworkAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    if yname == "area-id" { return "AreaId" }
    if yname == "propogate-from" { return "PropogateFrom" }
    if yname == "propogate-to" { return "PropogateTo" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "attached-policy" { return "AttachedPolicy" }
    if yname == "attach-point" { return "AttachPoint" }
    return ""
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol"] = binding.Protocol
    leafs["vrf-name"] = binding.VrfName
    leafs["proto-instance"] = binding.ProtoInstance
    leafs["af-name"] = binding.AfName
    leafs["saf-name"] = binding.SafName
    leafs["neighbor-address"] = binding.NeighborAddress
    leafs["neighbor-af-name"] = binding.NeighborAfName
    leafs["group-name"] = binding.GroupName
    leafs["direction"] = binding.Direction
    leafs["group"] = binding.Group
    leafs["source-protocol"] = binding.SourceProtocol
    leafs["aggregate-network-address"] = binding.AggregateNetworkAddress
    leafs["interface-name"] = binding.InterfaceName
    leafs["instance"] = binding.Instance
    leafs["area-id"] = binding.AreaId
    leafs["propogate-from"] = binding.PropogateFrom
    leafs["propogate-to"] = binding.PropogateTo
    leafs["route-policy-name"] = binding.RoutePolicyName
    leafs["attached-policy"] = binding.AttachedPolicy
    leafs["attach-point"] = binding.AttachPoint
    return leafs
}

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetYangName() string { return "binding" }

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetParent() types.Entity { return binding.parent }

func (binding *RoutingPolicy_Sets_ExtendedCommunityCost_Sets_Set_Attached_Binding) GetParentYangName() string { return "attached" }

// RoutingPolicy_Sets_ExtendedCommunityCost_Unused
// All objects of a given type that are not
// referenced at all
type RoutingPolicy_Sets_ExtendedCommunityCost_Unused struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetFilter() yfilter.YFilter { return unused.YFilter }

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) SetFilter(yf yfilter.YFilter) { unused.YFilter = yf }

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetSegmentPath() string {
    return "unused"
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = unused.Object
    return leafs
}

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetBundleName() string { return "cisco_ios_xr" }

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetYangName() string { return "unused" }

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) SetParent(parent types.Entity) { unused.parent = parent }

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetParent() types.Entity { return unused.parent }

func (unused *RoutingPolicy_Sets_ExtendedCommunityCost_Unused) GetParentYangName() string { return "extended-community-cost" }

// RoutingPolicy_Sets_ExtendedCommunityCost_Inactive
// All objects of a given type that are not
// attached to a protocol
type RoutingPolicy_Sets_ExtendedCommunityCost_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = inactive.Object
    return leafs
}

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetYangName() string { return "inactive" }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *RoutingPolicy_Sets_ExtendedCommunityCost_Inactive) GetParentYangName() string { return "extended-community-cost" }

// RoutingPolicy_Sets_ExtendedCommunityCost_Active
// All objects of a given type that are attached to
// a protocol
type RoutingPolicy_Sets_ExtendedCommunityCost_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy objects. The type is slice of string.
    Object []interface{}
}

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetSegmentPath() string {
    return "active"
}

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = active.Object
    return leafs
}

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetYangName() string { return "active" }

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetParent() types.Entity { return active.parent }

func (active *RoutingPolicy_Sets_ExtendedCommunityCost_Active) GetParentYangName() string { return "extended-community-cost" }

