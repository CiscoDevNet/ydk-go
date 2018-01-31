// This module describes a YANG model for routing policy
// configuration. It is a limited subset of all of the policy
// configuration parameters available in the variety of vendor
// implementations, but supports widely used constructs for managing
// how routes are imported, exported, and modified across different
// routing protocols.  This module is intended to be used in
// conjunction with routing protocol configuration models (e.g.,
// BGP) defined in other modules.
// 
// Route policy expression:
// 
// Policies are expressed as a set of top-level policy definitions,
// each of which consists of a sequence of policy statements. Policy
// statements consist of simple condition-action tuples. Conditions
// may include mutiple match or comparison operations, and similarly
// actions may be multitude of changes to route attributes or a
// final disposition of accepting or rejecting the route.
// 
// Route policy evaluation:
// 
// Policy definitions are referenced in routing protocol
// configurations using import and export configuration statements.
// The arguments are members of an ordered list of named policy
// definitions which comprise a policy chain, and optionally, an
// explicit default policy action (i.e., reject or accept).
// 
// Evaluation of each policy definition proceeds by evaluating its
// corresponding individual policy statements in order.  When a
// condition statement in a policy statement is satisfied, the
// corresponding action statement is executed.  If the action
// statement has either accept-route or reject-route actions, policy
// evaluation of the current policy definition stops, and no further
// policy definitions in the chain are evaluated.
// 
// If the condition is not satisfied, then evaluation proceeds to
// the next policy statement.  If none of the policy statement
// conditions are satisfied, then evaluation of the current policy
// definition stops, and the next policy definition in the chain is
// evaluated.  When the end of the policy chain is reached, the
// default route disposition action is performed (i.e., reject-route
// unless an an alternate default action is specified for the
// chain).
// 
// Policy 'subroutines' (or nested policies) are supported by
// allowing policy statement conditions to reference another policy
// definition which applies conditions and actions from the
// referenced policy before returning to the calling policy
// statement and resuming evaluation.  If the called policy
// results in an accept-route (either explicit or by default), then
// the subroutine returns an effective true value to the calling
// policy.  Similarly, a reject-route action returns false.  If the
// subroutine returns true, the calling policy continues to evaluate
// the remaining conditions (using a modified route if the
// subroutine performed any changes to the route).
package routing_policy

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package routing_policy"))
    ydk.RegisterEntity("{http://openconfig.net/yang/routing-policy routing-policy}", reflect.TypeOf(RoutingPolicy{}))
    ydk.RegisterEntity("openconfig-routing-policy:routing-policy", reflect.TypeOf(RoutingPolicy{}))
}

// DefaultPolicyType represents a policy chain
type DefaultPolicyType string

const (
    // default policy to accept the route
    DefaultPolicyType_ACCEPT_ROUTE DefaultPolicyType = "ACCEPT_ROUTE"

    // default policy to reject the route
    DefaultPolicyType_REJECT_ROUTE DefaultPolicyType = "REJECT_ROUTE"
)

// RoutingPolicy
// Top-level container for all routing policy configuration
type RoutingPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Predefined sets of attributes used in policy match statements.
    DefinedSets RoutingPolicy_DefinedSets

    // Enclosing container for the list of top-level policy  definitions.
    PolicyDefinitions RoutingPolicy_PolicyDefinitions
}

func (routingPolicy *RoutingPolicy) GetFilter() yfilter.YFilter { return routingPolicy.YFilter }

func (routingPolicy *RoutingPolicy) SetFilter(yf yfilter.YFilter) { routingPolicy.YFilter = yf }

func (routingPolicy *RoutingPolicy) GetGoName(yname string) string {
    if yname == "defined-sets" { return "DefinedSets" }
    if yname == "policy-definitions" { return "PolicyDefinitions" }
    return ""
}

func (routingPolicy *RoutingPolicy) GetSegmentPath() string {
    return "openconfig-routing-policy:routing-policy"
}

func (routingPolicy *RoutingPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "defined-sets" {
        return &routingPolicy.DefinedSets
    }
    if childYangName == "policy-definitions" {
        return &routingPolicy.PolicyDefinitions
    }
    return nil
}

func (routingPolicy *RoutingPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["defined-sets"] = &routingPolicy.DefinedSets
    children["policy-definitions"] = &routingPolicy.PolicyDefinitions
    return children
}

func (routingPolicy *RoutingPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routingPolicy *RoutingPolicy) GetBundleName() string { return "openconfig" }

func (routingPolicy *RoutingPolicy) GetYangName() string { return "routing-policy" }

func (routingPolicy *RoutingPolicy) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routingPolicy *RoutingPolicy) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routingPolicy *RoutingPolicy) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routingPolicy *RoutingPolicy) SetParent(parent types.Entity) { routingPolicy.parent = parent }

func (routingPolicy *RoutingPolicy) GetParent() types.Entity { return routingPolicy.parent }

func (routingPolicy *RoutingPolicy) GetParentYangName() string { return "openconfig-routing-policy" }

// RoutingPolicy_DefinedSets
// Predefined sets of attributes used in policy match
// statements
type RoutingPolicy_DefinedSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enclosing container .
    PrefixSets RoutingPolicy_DefinedSets_PrefixSets

    // Enclosing container for the list of neighbor set definitions.
    NeighborSets RoutingPolicy_DefinedSets_NeighborSets

    // Enclosing container for the list of tag sets.
    TagSets RoutingPolicy_DefinedSets_TagSets

    // BGP-related set definitions for policy match conditions.
    BgpDefinedSets RoutingPolicy_DefinedSets_BgpDefinedSets
}

func (definedSets *RoutingPolicy_DefinedSets) GetFilter() yfilter.YFilter { return definedSets.YFilter }

func (definedSets *RoutingPolicy_DefinedSets) SetFilter(yf yfilter.YFilter) { definedSets.YFilter = yf }

func (definedSets *RoutingPolicy_DefinedSets) GetGoName(yname string) string {
    if yname == "prefix-sets" { return "PrefixSets" }
    if yname == "neighbor-sets" { return "NeighborSets" }
    if yname == "tag-sets" { return "TagSets" }
    if yname == "openconfig-bgp-policy:bgp-defined-sets" { return "BgpDefinedSets" }
    return ""
}

func (definedSets *RoutingPolicy_DefinedSets) GetSegmentPath() string {
    return "defined-sets"
}

func (definedSets *RoutingPolicy_DefinedSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-sets" {
        return &definedSets.PrefixSets
    }
    if childYangName == "neighbor-sets" {
        return &definedSets.NeighborSets
    }
    if childYangName == "tag-sets" {
        return &definedSets.TagSets
    }
    if childYangName == "openconfig-bgp-policy:bgp-defined-sets" {
        return &definedSets.BgpDefinedSets
    }
    return nil
}

func (definedSets *RoutingPolicy_DefinedSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-sets"] = &definedSets.PrefixSets
    children["neighbor-sets"] = &definedSets.NeighborSets
    children["tag-sets"] = &definedSets.TagSets
    children["openconfig-bgp-policy:bgp-defined-sets"] = &definedSets.BgpDefinedSets
    return children
}

func (definedSets *RoutingPolicy_DefinedSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (definedSets *RoutingPolicy_DefinedSets) GetBundleName() string { return "openconfig" }

func (definedSets *RoutingPolicy_DefinedSets) GetYangName() string { return "defined-sets" }

func (definedSets *RoutingPolicy_DefinedSets) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (definedSets *RoutingPolicy_DefinedSets) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (definedSets *RoutingPolicy_DefinedSets) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (definedSets *RoutingPolicy_DefinedSets) SetParent(parent types.Entity) { definedSets.parent = parent }

func (definedSets *RoutingPolicy_DefinedSets) GetParent() types.Entity { return definedSets.parent }

func (definedSets *RoutingPolicy_DefinedSets) GetParentYangName() string { return "routing-policy" }

// RoutingPolicy_DefinedSets_PrefixSets
// Enclosing container 
type RoutingPolicy_DefinedSets_PrefixSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of the defined prefix sets. The type is slice of
    // RoutingPolicy_DefinedSets_PrefixSets_PrefixSet.
    PrefixSet []RoutingPolicy_DefinedSets_PrefixSets_PrefixSet
}

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetFilter() yfilter.YFilter { return prefixSets.YFilter }

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) SetFilter(yf yfilter.YFilter) { prefixSets.YFilter = yf }

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetGoName(yname string) string {
    if yname == "prefix-set" { return "PrefixSet" }
    return ""
}

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetSegmentPath() string {
    return "prefix-sets"
}

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-set" {
        for _, c := range prefixSets.PrefixSet {
            if prefixSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_DefinedSets_PrefixSets_PrefixSet{}
        prefixSets.PrefixSet = append(prefixSets.PrefixSet, child)
        return &prefixSets.PrefixSet[len(prefixSets.PrefixSet)-1]
    }
    return nil
}

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixSets.PrefixSet {
        children[prefixSets.PrefixSet[i].GetSegmentPath()] = &prefixSets.PrefixSet[i]
    }
    return children
}

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetBundleName() string { return "openconfig" }

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetYangName() string { return "prefix-sets" }

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) SetParent(parent types.Entity) { prefixSets.parent = parent }

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetParent() types.Entity { return prefixSets.parent }

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetParentYangName() string { return "defined-sets" }

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet
// List of the defined prefix sets
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to prefix name list key. The type is
    // string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config_PrefixSetName
    PrefixSetName interface{}

    // Configuration data for prefix sets.
    Config RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config

    // Operational state data .
    State RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State

    // Enclosing container for the list of prefixes in a policy prefix list.
    Prefixes RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes
}

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetFilter() yfilter.YFilter { return prefixSet.YFilter }

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) SetFilter(yf yfilter.YFilter) { prefixSet.YFilter = yf }

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetGoName(yname string) string {
    if yname == "prefix-set-name" { return "PrefixSetName" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "prefixes" { return "Prefixes" }
    return ""
}

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetSegmentPath() string {
    return "prefix-set" + "[prefix-set-name='" + fmt.Sprintf("%v", prefixSet.PrefixSetName) + "']"
}

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixSet.Config
    }
    if childYangName == "state" {
        return &prefixSet.State
    }
    if childYangName == "prefixes" {
        return &prefixSet.Prefixes
    }
    return nil
}

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixSet.Config
    children["state"] = &prefixSet.State
    children["prefixes"] = &prefixSet.Prefixes
    return children
}

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-set-name"] = prefixSet.PrefixSetName
    return leafs
}

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetBundleName() string { return "openconfig" }

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetYangName() string { return "prefix-set" }

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) SetParent(parent types.Entity) { prefixSet.parent = parent }

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetParent() types.Entity { return prefixSet.parent }

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetParentYangName() string { return "prefix-sets" }

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config
// Configuration data for prefix sets
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name / label of the prefix set -- this is used to reference the set in
    // match conditions. The type is string.
    PrefixSetName interface{}
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetGoName(yname string) string {
    if yname == "prefix-set-name" { return "PrefixSetName" }
    return ""
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-set-name"] = config.PrefixSetName
    return leafs
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetParentYangName() string { return "prefix-set" }

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State
// Operational state data 
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name / label of the prefix set -- this is used to reference the set in
    // match conditions. The type is string.
    PrefixSetName interface{}
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetGoName(yname string) string {
    if yname == "prefix-set-name" { return "PrefixSetName" }
    return ""
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-set-name"] = state.PrefixSetName
    return leafs
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetParentYangName() string { return "prefix-set" }

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes
// Enclosing container for the list of prefixes in a policy
// prefix list
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of prefixes in the prefix set. The type is slice of
    // RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix.
    Prefix []RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix
}

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetFilter() yfilter.YFilter { return prefixes.YFilter }

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) SetFilter(yf yfilter.YFilter) { prefixes.YFilter = yf }

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetSegmentPath() string {
    return "prefixes"
}

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        for _, c := range prefixes.Prefix {
            if prefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix{}
        prefixes.Prefix = append(prefixes.Prefix, child)
        return &prefixes.Prefix[len(prefixes.Prefix)-1]
    }
    return nil
}

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixes.Prefix {
        children[prefixes.Prefix[i].GetSegmentPath()] = &prefixes.Prefix[i]
    }
    return children
}

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetBundleName() string { return "openconfig" }

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetYangName() string { return "prefixes" }

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) SetParent(parent types.Entity) { prefixes.parent = parent }

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetParent() types.Entity { return prefixes.parent }

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetParentYangName() string { return "prefix-set" }

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix
// List of prefixes in the prefix set
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the ip-prefix list key. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    IpPrefix interface{}

    // This attribute is a key. Reference to the masklength-range list key. The
    // type is string with pattern: ^([0-9]+\.\.[0-9]+)|exact$. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config_MasklengthRange
    MasklengthRange interface{}

    // Configuration data for prefix definition.
    Config RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config

    // Operational state data for prefix definition.
    State RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State
}

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetGoName(yname string) string {
    if yname == "ip-prefix" { return "IpPrefix" }
    if yname == "masklength-range" { return "MasklengthRange" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetSegmentPath() string {
    return "prefix" + "[ip-prefix='" + fmt.Sprintf("%v", prefix.IpPrefix) + "']" + "[masklength-range='" + fmt.Sprintf("%v", prefix.MasklengthRange) + "']"
}

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefix.Config
    }
    if childYangName == "state" {
        return &prefix.State
    }
    return nil
}

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefix.Config
    children["state"] = &prefix.State
    return children
}

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-prefix"] = prefix.IpPrefix
    leafs["masklength-range"] = prefix.MasklengthRange
    return leafs
}

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetBundleName() string { return "openconfig" }

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetYangName() string { return "prefix" }

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetParentYangName() string { return "prefixes" }

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config
// Configuration data for prefix definition
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The prefix member in CIDR notation -- while the prefix may be either IPv4
    // or IPv6, most implementations require all members of the prefix set to be
    // the same address family.  Mixing address types in the same prefix set is
    // likely to cause an error. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))
    // This attribute is mandatory..
    IpPrefix interface{}

    // Defines a range for the masklength, or 'exact' if the prefix has an exact
    // length.  Example: 10.3.192.0/21 through 10.3.192.0/24 would be expressed as
    // prefix: 10.3.192.0/21, masklength-range: 21..24.  Example: 10.3.192.0/21
    // would be expressed as prefix: 10.3.192.0/21, masklength-range: exact. The
    // type is string with pattern: ^([0-9]+\.\.[0-9]+)|exact$.
    MasklengthRange interface{}
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetGoName(yname string) string {
    if yname == "ip-prefix" { return "IpPrefix" }
    if yname == "masklength-range" { return "MasklengthRange" }
    return ""
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-prefix"] = config.IpPrefix
    leafs["masklength-range"] = config.MasklengthRange
    return leafs
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetParentYangName() string { return "prefix" }

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State
// Operational state data for prefix definition
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The prefix member in CIDR notation -- while the prefix may be either IPv4
    // or IPv6, most implementations require all members of the prefix set to be
    // the same address family.  Mixing address types in the same prefix set is
    // likely to cause an error. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))
    // This attribute is mandatory..
    IpPrefix interface{}

    // Defines a range for the masklength, or 'exact' if the prefix has an exact
    // length.  Example: 10.3.192.0/21 through 10.3.192.0/24 would be expressed as
    // prefix: 10.3.192.0/21, masklength-range: 21..24.  Example: 10.3.192.0/21
    // would be expressed as prefix: 10.3.192.0/21, masklength-range: exact. The
    // type is string with pattern: ^([0-9]+\.\.[0-9]+)|exact$.
    MasklengthRange interface{}
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetGoName(yname string) string {
    if yname == "ip-prefix" { return "IpPrefix" }
    if yname == "masklength-range" { return "MasklengthRange" }
    return ""
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-prefix"] = state.IpPrefix
    leafs["masklength-range"] = state.MasklengthRange
    return leafs
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetParentYangName() string { return "prefix" }

// RoutingPolicy_DefinedSets_NeighborSets
// Enclosing container for the list of neighbor set
// definitions
type RoutingPolicy_DefinedSets_NeighborSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of defined neighbor sets for use in policies. The type is slice of
    // RoutingPolicy_DefinedSets_NeighborSets_NeighborSet.
    NeighborSet []RoutingPolicy_DefinedSets_NeighborSets_NeighborSet
}

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetFilter() yfilter.YFilter { return neighborSets.YFilter }

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) SetFilter(yf yfilter.YFilter) { neighborSets.YFilter = yf }

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetGoName(yname string) string {
    if yname == "neighbor-set" { return "NeighborSet" }
    return ""
}

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetSegmentPath() string {
    return "neighbor-sets"
}

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor-set" {
        for _, c := range neighborSets.NeighborSet {
            if neighborSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_DefinedSets_NeighborSets_NeighborSet{}
        neighborSets.NeighborSet = append(neighborSets.NeighborSet, child)
        return &neighborSets.NeighborSet[len(neighborSets.NeighborSet)-1]
    }
    return nil
}

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighborSets.NeighborSet {
        children[neighborSets.NeighborSet[i].GetSegmentPath()] = &neighborSets.NeighborSet[i]
    }
    return children
}

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetBundleName() string { return "openconfig" }

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetYangName() string { return "neighbor-sets" }

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) SetParent(parent types.Entity) { neighborSets.parent = parent }

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetParent() types.Entity { return neighborSets.parent }

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetParentYangName() string { return "defined-sets" }

// RoutingPolicy_DefinedSets_NeighborSets_NeighborSet
// List of defined neighbor sets for use in policies.
type RoutingPolicy_DefinedSets_NeighborSets_NeighborSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the neighbor set name list key. The
    // type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config_NeighborSetName
    NeighborSetName interface{}

    // Configuration data for neighbor sets.
    Config RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config

    // Operational state data for neighbor sets.
    State RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State
}

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetFilter() yfilter.YFilter { return neighborSet.YFilter }

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) SetFilter(yf yfilter.YFilter) { neighborSet.YFilter = yf }

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetGoName(yname string) string {
    if yname == "neighbor-set-name" { return "NeighborSetName" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetSegmentPath() string {
    return "neighbor-set" + "[neighbor-set-name='" + fmt.Sprintf("%v", neighborSet.NeighborSetName) + "']"
}

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &neighborSet.Config
    }
    if childYangName == "state" {
        return &neighborSet.State
    }
    return nil
}

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &neighborSet.Config
    children["state"] = &neighborSet.State
    return children
}

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-set-name"] = neighborSet.NeighborSetName
    return leafs
}

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetBundleName() string { return "openconfig" }

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetYangName() string { return "neighbor-set" }

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) SetParent(parent types.Entity) { neighborSet.parent = parent }

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetParent() types.Entity { return neighborSet.parent }

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetParentYangName() string { return "neighbor-sets" }

// RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config
// Configuration data for neighbor sets.
type RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name / label of the neighbor set -- this is used to reference the set in
    // match conditions. The type is string.
    NeighborSetName interface{}

    // List of IP addresses in the neighbor set. The type is one of the following
    // types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address []interface{}
}

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetGoName(yname string) string {
    if yname == "neighbor-set-name" { return "NeighborSetName" }
    if yname == "address" { return "Address" }
    return ""
}

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-set-name"] = config.NeighborSetName
    leafs["address"] = config.Address
    return leafs
}

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetParentYangName() string { return "neighbor-set" }

// RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State
// Operational state data for neighbor sets.
type RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name / label of the neighbor set -- this is used to reference the set in
    // match conditions. The type is string.
    NeighborSetName interface{}

    // List of IP addresses in the neighbor set. The type is one of the following
    // types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address []interface{}
}

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetGoName(yname string) string {
    if yname == "neighbor-set-name" { return "NeighborSetName" }
    if yname == "address" { return "Address" }
    return ""
}

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-set-name"] = state.NeighborSetName
    leafs["address"] = state.Address
    return leafs
}

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetParentYangName() string { return "neighbor-set" }

// RoutingPolicy_DefinedSets_TagSets
// Enclosing container for the list of tag sets.
type RoutingPolicy_DefinedSets_TagSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of tag set definitions. The type is slice of
    // RoutingPolicy_DefinedSets_TagSets_TagSet.
    TagSet []RoutingPolicy_DefinedSets_TagSets_TagSet
}

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetFilter() yfilter.YFilter { return tagSets.YFilter }

func (tagSets *RoutingPolicy_DefinedSets_TagSets) SetFilter(yf yfilter.YFilter) { tagSets.YFilter = yf }

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetGoName(yname string) string {
    if yname == "tag-set" { return "TagSet" }
    return ""
}

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetSegmentPath() string {
    return "tag-sets"
}

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tag-set" {
        for _, c := range tagSets.TagSet {
            if tagSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_DefinedSets_TagSets_TagSet{}
        tagSets.TagSet = append(tagSets.TagSet, child)
        return &tagSets.TagSet[len(tagSets.TagSet)-1]
    }
    return nil
}

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tagSets.TagSet {
        children[tagSets.TagSet[i].GetSegmentPath()] = &tagSets.TagSet[i]
    }
    return children
}

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetBundleName() string { return "openconfig" }

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetYangName() string { return "tag-sets" }

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (tagSets *RoutingPolicy_DefinedSets_TagSets) SetParent(parent types.Entity) { tagSets.parent = parent }

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetParent() types.Entity { return tagSets.parent }

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetParentYangName() string { return "defined-sets" }

// RoutingPolicy_DefinedSets_TagSets_TagSet
// List of tag set definitions.
type RoutingPolicy_DefinedSets_TagSets_TagSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the tag set name list key. The type
    // is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_TagSets_TagSet_Config_TagSetName
    TagSetName interface{}

    // Configuration data for tag sets.
    Config RoutingPolicy_DefinedSets_TagSets_TagSet_Config

    // Operational state data for tag sets.
    State RoutingPolicy_DefinedSets_TagSets_TagSet_State
}

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetFilter() yfilter.YFilter { return tagSet.YFilter }

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) SetFilter(yf yfilter.YFilter) { tagSet.YFilter = yf }

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetGoName(yname string) string {
    if yname == "tag-set-name" { return "TagSetName" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetSegmentPath() string {
    return "tag-set" + "[tag-set-name='" + fmt.Sprintf("%v", tagSet.TagSetName) + "']"
}

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &tagSet.Config
    }
    if childYangName == "state" {
        return &tagSet.State
    }
    return nil
}

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &tagSet.Config
    children["state"] = &tagSet.State
    return children
}

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tag-set-name"] = tagSet.TagSetName
    return leafs
}

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetBundleName() string { return "openconfig" }

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetYangName() string { return "tag-set" }

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) SetParent(parent types.Entity) { tagSet.parent = parent }

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetParent() types.Entity { return tagSet.parent }

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetParentYangName() string { return "tag-sets" }

// RoutingPolicy_DefinedSets_TagSets_TagSet_Config
// Configuration data for tag sets
type RoutingPolicy_DefinedSets_TagSets_TagSet_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name / label of the tag set -- this is used to reference the set in match
    // conditions. The type is string.
    TagSetName interface{}

    // Value of the tag set member. The type is one of the following types: slice
    // of int with range: 0..4294967295, or slice of string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TagValue []interface{}
}

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetGoName(yname string) string {
    if yname == "tag-set-name" { return "TagSetName" }
    if yname == "tag-value" { return "TagValue" }
    return ""
}

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tag-set-name"] = config.TagSetName
    leafs["tag-value"] = config.TagValue
    return leafs
}

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetParentYangName() string { return "tag-set" }

// RoutingPolicy_DefinedSets_TagSets_TagSet_State
// Operational state data for tag sets
type RoutingPolicy_DefinedSets_TagSets_TagSet_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name / label of the tag set -- this is used to reference the set in match
    // conditions. The type is string.
    TagSetName interface{}

    // Value of the tag set member. The type is one of the following types: slice
    // of int with range: 0..4294967295, or slice of string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TagValue []interface{}
}

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetGoName(yname string) string {
    if yname == "tag-set-name" { return "TagSetName" }
    if yname == "tag-value" { return "TagValue" }
    return ""
}

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tag-set-name"] = state.TagSetName
    leafs["tag-value"] = state.TagValue
    return leafs
}

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetParentYangName() string { return "tag-set" }

// RoutingPolicy_DefinedSets_BgpDefinedSets
// BGP-related set definitions for policy match conditions
type RoutingPolicy_DefinedSets_BgpDefinedSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enclosing container for list of defined BGP community sets.
    CommunitySets RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets

    // Enclosing container for list of extended BGP community sets.
    ExtCommunitySets RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets

    // Enclosing container for list of define AS path sets.
    AsPathSets RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets
}

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetFilter() yfilter.YFilter { return bgpDefinedSets.YFilter }

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) SetFilter(yf yfilter.YFilter) { bgpDefinedSets.YFilter = yf }

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetGoName(yname string) string {
    if yname == "community-sets" { return "CommunitySets" }
    if yname == "ext-community-sets" { return "ExtCommunitySets" }
    if yname == "as-path-sets" { return "AsPathSets" }
    return ""
}

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetSegmentPath() string {
    return "openconfig-bgp-policy:bgp-defined-sets"
}

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "community-sets" {
        return &bgpDefinedSets.CommunitySets
    }
    if childYangName == "ext-community-sets" {
        return &bgpDefinedSets.ExtCommunitySets
    }
    if childYangName == "as-path-sets" {
        return &bgpDefinedSets.AsPathSets
    }
    return nil
}

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["community-sets"] = &bgpDefinedSets.CommunitySets
    children["ext-community-sets"] = &bgpDefinedSets.ExtCommunitySets
    children["as-path-sets"] = &bgpDefinedSets.AsPathSets
    return children
}

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetBundleName() string { return "openconfig" }

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetYangName() string { return "bgp-defined-sets" }

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) SetParent(parent types.Entity) { bgpDefinedSets.parent = parent }

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetParent() types.Entity { return bgpDefinedSets.parent }

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetParentYangName() string { return "defined-sets" }

// RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets
// Enclosing container for list of defined BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of defined BGP community sets. The type is slice of
    // RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet.
    CommunitySet []RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet
}

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetFilter() yfilter.YFilter { return communitySets.YFilter }

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) SetFilter(yf yfilter.YFilter) { communitySets.YFilter = yf }

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetGoName(yname string) string {
    if yname == "community-set" { return "CommunitySet" }
    return ""
}

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetSegmentPath() string {
    return "community-sets"
}

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "community-set" {
        for _, c := range communitySets.CommunitySet {
            if communitySets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet{}
        communitySets.CommunitySet = append(communitySets.CommunitySet, child)
        return &communitySets.CommunitySet[len(communitySets.CommunitySet)-1]
    }
    return nil
}

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range communitySets.CommunitySet {
        children[communitySets.CommunitySet[i].GetSegmentPath()] = &communitySets.CommunitySet[i]
    }
    return children
}

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetBundleName() string { return "openconfig" }

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetYangName() string { return "community-sets" }

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) SetParent(parent types.Entity) { communitySets.parent = parent }

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetParent() types.Entity { return communitySets.parent }

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetParentYangName() string { return "bgp-defined-sets" }

// RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet
// List of defined BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to list key. The type is string. Refers
    // to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config_CommunitySetName
    CommunitySetName interface{}

    // Configuration data for BGP community sets.
    Config RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config

    // Operational state data for BGP community sets.
    State RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State
}

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetFilter() yfilter.YFilter { return communitySet.YFilter }

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) SetFilter(yf yfilter.YFilter) { communitySet.YFilter = yf }

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetGoName(yname string) string {
    if yname == "community-set-name" { return "CommunitySetName" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetSegmentPath() string {
    return "community-set" + "[community-set-name='" + fmt.Sprintf("%v", communitySet.CommunitySetName) + "']"
}

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &communitySet.Config
    }
    if childYangName == "state" {
        return &communitySet.State
    }
    return nil
}

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &communitySet.Config
    children["state"] = &communitySet.State
    return children
}

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-set-name"] = communitySet.CommunitySetName
    return leafs
}

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetBundleName() string { return "openconfig" }

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetYangName() string { return "community-set" }

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) SetParent(parent types.Entity) { communitySet.parent = parent }

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetParent() types.Entity { return communitySet.parent }

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetParentYangName() string { return "community-sets" }

// RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config
// Configuration data for BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name / label of the community set -- this is used to reference the set in
    // match conditions. The type is string. This attribute is mandatory.
    CommunitySetName interface{}

    // members of the community set. The type is one of the following types: slice
    // of int with range: 65536..4294901759, or slice of string with pattern:
    // ([0-9]+:[0-9]+)., or slice of string, or slice of  
    // :go:struct:`BGPWELLKNOWNSTDCOMMUNITY
    // <ydk/models/bgp_types/BGPWELLKNOWNSTDCOMMUNITY>`.
    CommunityMember []interface{}
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetGoName(yname string) string {
    if yname == "community-set-name" { return "CommunitySetName" }
    if yname == "community-member" { return "CommunityMember" }
    return ""
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-set-name"] = config.CommunitySetName
    leafs["community-member"] = config.CommunityMember
    return leafs
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetParentYangName() string { return "community-set" }

// RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State
// Operational state data for BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name / label of the community set -- this is used to reference the set in
    // match conditions. The type is string. This attribute is mandatory.
    CommunitySetName interface{}

    // members of the community set. The type is one of the following types: slice
    // of int with range: 65536..4294901759, or slice of string with pattern:
    // ([0-9]+:[0-9]+)., or slice of string, or slice of  
    // :go:struct:`BGPWELLKNOWNSTDCOMMUNITY
    // <ydk/models/bgp_types/BGPWELLKNOWNSTDCOMMUNITY>`.
    CommunityMember []interface{}
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetGoName(yname string) string {
    if yname == "community-set-name" { return "CommunitySetName" }
    if yname == "community-member" { return "CommunityMember" }
    return ""
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-set-name"] = state.CommunitySetName
    leafs["community-member"] = state.CommunityMember
    return leafs
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetParentYangName() string { return "community-set" }

// RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets
// Enclosing container for list of extended BGP community
// sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of defined extended BGP community sets. The type is slice of
    // RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet.
    ExtCommunitySet []RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet
}

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetFilter() yfilter.YFilter { return extCommunitySets.YFilter }

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) SetFilter(yf yfilter.YFilter) { extCommunitySets.YFilter = yf }

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetGoName(yname string) string {
    if yname == "ext-community-set" { return "ExtCommunitySet" }
    return ""
}

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetSegmentPath() string {
    return "ext-community-sets"
}

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext-community-set" {
        for _, c := range extCommunitySets.ExtCommunitySet {
            if extCommunitySets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet{}
        extCommunitySets.ExtCommunitySet = append(extCommunitySets.ExtCommunitySet, child)
        return &extCommunitySets.ExtCommunitySet[len(extCommunitySets.ExtCommunitySet)-1]
    }
    return nil
}

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extCommunitySets.ExtCommunitySet {
        children[extCommunitySets.ExtCommunitySet[i].GetSegmentPath()] = &extCommunitySets.ExtCommunitySet[i]
    }
    return children
}

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetBundleName() string { return "openconfig" }

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetYangName() string { return "ext-community-sets" }

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) SetParent(parent types.Entity) { extCommunitySets.parent = parent }

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetParent() types.Entity { return extCommunitySets.parent }

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetParentYangName() string { return "bgp-defined-sets" }

// RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet
// List of defined extended BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to list key. The type is string. Refers
    // to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config_ExtCommunitySetName
    ExtCommunitySetName interface{}

    // Configuration data for extended BGP community sets.
    Config RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config

    // Operational state data for extended BGP community sets.
    State RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State
}

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetFilter() yfilter.YFilter { return extCommunitySet.YFilter }

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) SetFilter(yf yfilter.YFilter) { extCommunitySet.YFilter = yf }

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetGoName(yname string) string {
    if yname == "ext-community-set-name" { return "ExtCommunitySetName" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetSegmentPath() string {
    return "ext-community-set" + "[ext-community-set-name='" + fmt.Sprintf("%v", extCommunitySet.ExtCommunitySetName) + "']"
}

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &extCommunitySet.Config
    }
    if childYangName == "state" {
        return &extCommunitySet.State
    }
    return nil
}

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &extCommunitySet.Config
    children["state"] = &extCommunitySet.State
    return children
}

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ext-community-set-name"] = extCommunitySet.ExtCommunitySetName
    return leafs
}

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetBundleName() string { return "openconfig" }

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetYangName() string { return "ext-community-set" }

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) SetParent(parent types.Entity) { extCommunitySet.parent = parent }

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetParent() types.Entity { return extCommunitySet.parent }

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetParentYangName() string { return "ext-community-sets" }

// RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config
// Configuration data for extended BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name / label of the extended community set -- this is used to reference the
    // set in match conditions. The type is string.
    ExtCommunitySetName interface{}

    // members of the extended community set. The type is one of the following
    // types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).,
    // or slice of string.
    ExtCommunityMember []interface{}
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetGoName(yname string) string {
    if yname == "ext-community-set-name" { return "ExtCommunitySetName" }
    if yname == "ext-community-member" { return "ExtCommunityMember" }
    return ""
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ext-community-set-name"] = config.ExtCommunitySetName
    leafs["ext-community-member"] = config.ExtCommunityMember
    return leafs
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetParentYangName() string { return "ext-community-set" }

// RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State
// Operational state data for extended BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name / label of the extended community set -- this is used to reference the
    // set in match conditions. The type is string.
    ExtCommunitySetName interface{}

    // members of the extended community set. The type is one of the following
    // types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).,
    // or slice of string.
    ExtCommunityMember []interface{}
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetGoName(yname string) string {
    if yname == "ext-community-set-name" { return "ExtCommunitySetName" }
    if yname == "ext-community-member" { return "ExtCommunityMember" }
    return ""
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ext-community-set-name"] = state.ExtCommunitySetName
    leafs["ext-community-member"] = state.ExtCommunityMember
    return leafs
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetParentYangName() string { return "ext-community-set" }

// RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets
// Enclosing container for list of define AS path sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of defined AS path sets. The type is slice of
    // RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet.
    AsPathSet []RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet
}

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetFilter() yfilter.YFilter { return asPathSets.YFilter }

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) SetFilter(yf yfilter.YFilter) { asPathSets.YFilter = yf }

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetGoName(yname string) string {
    if yname == "as-path-set" { return "AsPathSet" }
    return ""
}

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetSegmentPath() string {
    return "as-path-sets"
}

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "as-path-set" {
        for _, c := range asPathSets.AsPathSet {
            if asPathSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet{}
        asPathSets.AsPathSet = append(asPathSets.AsPathSet, child)
        return &asPathSets.AsPathSet[len(asPathSets.AsPathSet)-1]
    }
    return nil
}

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asPathSets.AsPathSet {
        children[asPathSets.AsPathSet[i].GetSegmentPath()] = &asPathSets.AsPathSet[i]
    }
    return children
}

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetBundleName() string { return "openconfig" }

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetYangName() string { return "as-path-sets" }

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) SetParent(parent types.Entity) { asPathSets.parent = parent }

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetParent() types.Entity { return asPathSets.parent }

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetParentYangName() string { return "bgp-defined-sets" }

// RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet
// List of defined AS path sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to list key. The type is string. Refers
    // to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config_AsPathSetName
    AsPathSetName interface{}

    // Configuration data for AS path sets.
    Config RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config

    // Operational state data for AS path sets.
    State RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State
}

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetFilter() yfilter.YFilter { return asPathSet.YFilter }

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) SetFilter(yf yfilter.YFilter) { asPathSet.YFilter = yf }

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetGoName(yname string) string {
    if yname == "as-path-set-name" { return "AsPathSetName" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetSegmentPath() string {
    return "as-path-set" + "[as-path-set-name='" + fmt.Sprintf("%v", asPathSet.AsPathSetName) + "']"
}

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &asPathSet.Config
    }
    if childYangName == "state" {
        return &asPathSet.State
    }
    return nil
}

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &asPathSet.Config
    children["state"] = &asPathSet.State
    return children
}

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-path-set-name"] = asPathSet.AsPathSetName
    return leafs
}

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetBundleName() string { return "openconfig" }

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetYangName() string { return "as-path-set" }

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) SetParent(parent types.Entity) { asPathSet.parent = parent }

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetParent() types.Entity { return asPathSet.parent }

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetParentYangName() string { return "as-path-sets" }

// RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config
// Configuration data for AS path sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name of the AS path set -- this is used to reference the set in match
    // conditions. The type is string.
    AsPathSetName interface{}

    // AS path expression -- list of ASes in the set. The type is slice of string.
    AsPathSetMember []interface{}
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetGoName(yname string) string {
    if yname == "as-path-set-name" { return "AsPathSetName" }
    if yname == "as-path-set-member" { return "AsPathSetMember" }
    return ""
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-path-set-name"] = config.AsPathSetName
    leafs["as-path-set-member"] = config.AsPathSetMember
    return leafs
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetParentYangName() string { return "as-path-set" }

// RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State
// Operational state data for AS path sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name of the AS path set -- this is used to reference the set in match
    // conditions. The type is string.
    AsPathSetName interface{}

    // AS path expression -- list of ASes in the set. The type is slice of string.
    AsPathSetMember []interface{}
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetGoName(yname string) string {
    if yname == "as-path-set-name" { return "AsPathSetName" }
    if yname == "as-path-set-member" { return "AsPathSetMember" }
    return ""
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-path-set-name"] = state.AsPathSetName
    leafs["as-path-set-member"] = state.AsPathSetMember
    return leafs
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetParentYangName() string { return "as-path-set" }

// RoutingPolicy_PolicyDefinitions
// Enclosing container for the list of top-level policy
//  definitions
type RoutingPolicy_PolicyDefinitions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of top-level policy definitions, keyed by unique name.  These policy
    // definitions are expected to be referenced (by name) in policy chains
    // specified in import or export configuration statements. The type is slice
    // of RoutingPolicy_PolicyDefinitions_PolicyDefinition.
    PolicyDefinition []RoutingPolicy_PolicyDefinitions_PolicyDefinition
}

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetFilter() yfilter.YFilter { return policyDefinitions.YFilter }

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) SetFilter(yf yfilter.YFilter) { policyDefinitions.YFilter = yf }

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetGoName(yname string) string {
    if yname == "policy-definition" { return "PolicyDefinition" }
    return ""
}

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetSegmentPath() string {
    return "policy-definitions"
}

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-definition" {
        for _, c := range policyDefinitions.PolicyDefinition {
            if policyDefinitions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_PolicyDefinitions_PolicyDefinition{}
        policyDefinitions.PolicyDefinition = append(policyDefinitions.PolicyDefinition, child)
        return &policyDefinitions.PolicyDefinition[len(policyDefinitions.PolicyDefinition)-1]
    }
    return nil
}

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policyDefinitions.PolicyDefinition {
        children[policyDefinitions.PolicyDefinition[i].GetSegmentPath()] = &policyDefinitions.PolicyDefinition[i]
    }
    return children
}

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetBundleName() string { return "openconfig" }

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetYangName() string { return "policy-definitions" }

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) SetParent(parent types.Entity) { policyDefinitions.parent = parent }

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetParent() types.Entity { return policyDefinitions.parent }

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetParentYangName() string { return "routing-policy" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition
// List of top-level policy definitions, keyed by unique
// name.  These policy definitions are expected to be
// referenced (by name) in policy chains specified in import
// or export configuration statements.
type RoutingPolicy_PolicyDefinitions_PolicyDefinition struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the list key. The type is string.
    // Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config_Name
    Name interface{}

    // Configuration data for policy defintions.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config

    // Operational state data for policy definitions.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_State

    // Enclosing container for policy statements.
    Statements RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements
}

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetFilter() yfilter.YFilter { return policyDefinition.YFilter }

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) SetFilter(yf yfilter.YFilter) { policyDefinition.YFilter = yf }

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "statements" { return "Statements" }
    return ""
}

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetSegmentPath() string {
    return "policy-definition" + "[name='" + fmt.Sprintf("%v", policyDefinition.Name) + "']"
}

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &policyDefinition.Config
    }
    if childYangName == "state" {
        return &policyDefinition.State
    }
    if childYangName == "statements" {
        return &policyDefinition.Statements
    }
    return nil
}

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &policyDefinition.Config
    children["state"] = &policyDefinition.State
    children["statements"] = &policyDefinition.Statements
    return children
}

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = policyDefinition.Name
    return leafs
}

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetBundleName() string { return "openconfig" }

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetYangName() string { return "policy-definition" }

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) SetParent(parent types.Entity) { policyDefinition.parent = parent }

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetParent() types.Entity { return policyDefinition.parent }

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetParentYangName() string { return "policy-definitions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config
// Configuration data for policy defintions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the top-level policy definition -- this name is used in references
    // to the current policy. The type is string.
    Name interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetParentYangName() string { return "policy-definition" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_State
// Operational state data for policy definitions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the top-level policy definition -- this name is used in references
    // to the current policy. The type is string.
    Name interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetParentYangName() string { return "policy-definition" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements
// Enclosing container for policy statements
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy statements group conditions and actions within a policy definition. 
    // They are evaluated in the order specified (see the description of policy
    // evaluation at the top of this module. The type is slice of
    // RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement.
    Statement []RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement
}

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetFilter() yfilter.YFilter { return statements.YFilter }

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) SetFilter(yf yfilter.YFilter) { statements.YFilter = yf }

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetGoName(yname string) string {
    if yname == "statement" { return "Statement" }
    return ""
}

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetSegmentPath() string {
    return "statements"
}

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statement" {
        for _, c := range statements.Statement {
            if statements.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement{}
        statements.Statement = append(statements.Statement, child)
        return &statements.Statement[len(statements.Statement)-1]
    }
    return nil
}

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statements.Statement {
        children[statements.Statement[i].GetSegmentPath()] = &statements.Statement[i]
    }
    return children
}

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetBundleName() string { return "openconfig" }

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetYangName() string { return "statements" }

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) SetParent(parent types.Entity) { statements.parent = parent }

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetParent() types.Entity { return statements.parent }

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetParentYangName() string { return "policy-definition" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement
// Policy statements group conditions and actions
// within a policy definition.  They are evaluated in
// the order specified (see the description of policy
// evaluation at the top of this module.
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to list key. The type is string. Refers
    // to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config_Name
    Name interface{}

    // Configuration data for policy statements.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config

    // Operational state data for policy statements.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State

    // Condition statements for the current policy statement.
    Conditions RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions

    // Top-level container for policy action statements.
    Actions RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions
}

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetFilter() yfilter.YFilter { return statement.YFilter }

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) SetFilter(yf yfilter.YFilter) { statement.YFilter = yf }

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "conditions" { return "Conditions" }
    if yname == "actions" { return "Actions" }
    return ""
}

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetSegmentPath() string {
    return "statement" + "[name='" + fmt.Sprintf("%v", statement.Name) + "']"
}

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &statement.Config
    }
    if childYangName == "state" {
        return &statement.State
    }
    if childYangName == "conditions" {
        return &statement.Conditions
    }
    if childYangName == "actions" {
        return &statement.Actions
    }
    return nil
}

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &statement.Config
    children["state"] = &statement.State
    children["conditions"] = &statement.Conditions
    children["actions"] = &statement.Actions
    return children
}

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = statement.Name
    return leafs
}

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetBundleName() string { return "openconfig" }

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetYangName() string { return "statement" }

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) SetParent(parent types.Entity) { statement.parent = parent }

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetParent() types.Entity { return statement.parent }

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetParentYangName() string { return "statements" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config
// Configuration data for policy statements
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name of the policy statement. The type is string.
    Name interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetParentYangName() string { return "statement" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State
// Operational state data for policy statements
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name of the policy statement. The type is string.
    Name interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetParentYangName() string { return "statement" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions
// Condition statements for the current policy statement
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for policy conditions.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config

    // Operational state data for policy conditions.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State

    // Top-level container for interface match conditions.
    MatchInterface RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface

    // Match a referenced prefix-set according to the logic defined in the
    // match-set-options leaf.
    MatchPrefixSet RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet

    // Match a referenced neighbor set according to the logic defined in the
    // match-set-options-leaf.
    MatchNeighborSet RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet

    // Match a referenced tag set according to the logic defined in the
    // match-options-set leaf.
    MatchTagSet RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet

    // Policy conditions for IGP attributes.
    IgpConditions RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions

    // Top-level container .
    BgpConditions RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions
}

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetFilter() yfilter.YFilter { return conditions.YFilter }

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) SetFilter(yf yfilter.YFilter) { conditions.YFilter = yf }

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "match-interface" { return "MatchInterface" }
    if yname == "match-prefix-set" { return "MatchPrefixSet" }
    if yname == "match-neighbor-set" { return "MatchNeighborSet" }
    if yname == "match-tag-set" { return "MatchTagSet" }
    if yname == "igp-conditions" { return "IgpConditions" }
    if yname == "openconfig-bgp-policy:bgp-conditions" { return "BgpConditions" }
    return ""
}

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetSegmentPath() string {
    return "conditions"
}

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &conditions.Config
    }
    if childYangName == "state" {
        return &conditions.State
    }
    if childYangName == "match-interface" {
        return &conditions.MatchInterface
    }
    if childYangName == "match-prefix-set" {
        return &conditions.MatchPrefixSet
    }
    if childYangName == "match-neighbor-set" {
        return &conditions.MatchNeighborSet
    }
    if childYangName == "match-tag-set" {
        return &conditions.MatchTagSet
    }
    if childYangName == "igp-conditions" {
        return &conditions.IgpConditions
    }
    if childYangName == "openconfig-bgp-policy:bgp-conditions" {
        return &conditions.BgpConditions
    }
    return nil
}

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &conditions.Config
    children["state"] = &conditions.State
    children["match-interface"] = &conditions.MatchInterface
    children["match-prefix-set"] = &conditions.MatchPrefixSet
    children["match-neighbor-set"] = &conditions.MatchNeighborSet
    children["match-tag-set"] = &conditions.MatchTagSet
    children["igp-conditions"] = &conditions.IgpConditions
    children["openconfig-bgp-policy:bgp-conditions"] = &conditions.BgpConditions
    return children
}

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetBundleName() string { return "openconfig" }

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetYangName() string { return "conditions" }

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) SetParent(parent types.Entity) { conditions.parent = parent }

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetParent() types.Entity { return conditions.parent }

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetParentYangName() string { return "statement" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config
// Configuration data for policy conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Applies the statements from the specified policy definition and then
    // returns control the current policy statement. Note that the called policy
    // may itself call other policies (subject to implementation limitations).
    // This is intended to provide a policy 'subroutine' capability.  The called
    // policy should contain an explicit or a default route disposition that
    // returns an effective true (accept-route) or false (reject-route), otherwise
    // the behavior may be ambiguous and implementation dependent. The type is
    // string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    CallPolicy interface{}

    // Condition to check the protocol / method used to install the route into the
    // local routing table. The type is one of the following:
    // BGPISISLOCALAGGREGATESTATICDIRECTLYCONNECTEDOSPFOSPF3.
    InstallProtocolEq interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetGoName(yname string) string {
    if yname == "call-policy" { return "CallPolicy" }
    if yname == "install-protocol-eq" { return "InstallProtocolEq" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["call-policy"] = config.CallPolicy
    leafs["install-protocol-eq"] = config.InstallProtocolEq
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetParentYangName() string { return "conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State
// Operational state data for policy conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Applies the statements from the specified policy definition and then
    // returns control the current policy statement. Note that the called policy
    // may itself call other policies (subject to implementation limitations).
    // This is intended to provide a policy 'subroutine' capability.  The called
    // policy should contain an explicit or a default route disposition that
    // returns an effective true (accept-route) or false (reject-route), otherwise
    // the behavior may be ambiguous and implementation dependent. The type is
    // string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    CallPolicy interface{}

    // Condition to check the protocol / method used to install the route into the
    // local routing table. The type is one of the following:
    // BGPISISLOCALAGGREGATESTATICDIRECTLYCONNECTEDOSPFOSPF3.
    InstallProtocolEq interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetGoName(yname string) string {
    if yname == "call-policy" { return "CallPolicy" }
    if yname == "install-protocol-eq" { return "InstallProtocolEq" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["call-policy"] = state.CallPolicy
    leafs["install-protocol-eq"] = state.InstallProtocolEq
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetParentYangName() string { return "conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface
// Top-level container for interface match conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for interface match conditions.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config

    // Operational state data for interface match conditions.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State
}

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetFilter() yfilter.YFilter { return matchInterface.YFilter }

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) SetFilter(yf yfilter.YFilter) { matchInterface.YFilter = yf }

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetSegmentPath() string {
    return "match-interface"
}

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &matchInterface.Config
    }
    if childYangName == "state" {
        return &matchInterface.State
    }
    return nil
}

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &matchInterface.Config
    children["state"] = &matchInterface.State
    return children
}

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetBundleName() string { return "openconfig" }

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetYangName() string { return "match-interface" }

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) SetParent(parent types.Entity) { matchInterface.parent = parent }

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetParent() types.Entity { return matchInterface.parent }

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetParentYangName() string { return "conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config
// Configuration data for interface match conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = config.Interface
    leafs["subinterface"] = config.Subinterface
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetParentYangName() string { return "match-interface" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State
// Operational state data for interface match conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = state.Interface
    leafs["subinterface"] = state.Subinterface
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetParentYangName() string { return "match-interface" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet
// Match a referenced prefix-set according to the logic
// defined in the match-set-options leaf
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for a prefix-set condition.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config

    // Operational state data for a prefix-set condition.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State
}

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetFilter() yfilter.YFilter { return matchPrefixSet.YFilter }

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) SetFilter(yf yfilter.YFilter) { matchPrefixSet.YFilter = yf }

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetSegmentPath() string {
    return "match-prefix-set"
}

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &matchPrefixSet.Config
    }
    if childYangName == "state" {
        return &matchPrefixSet.State
    }
    return nil
}

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &matchPrefixSet.Config
    children["state"] = &matchPrefixSet.State
    return children
}

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetBundleName() string { return "openconfig" }

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetYangName() string { return "match-prefix-set" }

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) SetParent(parent types.Entity) { matchPrefixSet.parent = parent }

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetParent() types.Entity { return matchPrefixSet.parent }

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetParentYangName() string { return "conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config
// Configuration data for a prefix-set condition
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined prefix set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_PrefixSetName
    PrefixSet interface{}

    // Optional parameter that governs the behaviour of the match operation.  This
    // leaf only supports matching on ANY member of the set or inverting the
    // match.  Matching on ALL is not supported). The type is
    // MatchSetOptionsRestrictedType.
    MatchSetOptions interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetGoName(yname string) string {
    if yname == "prefix-set" { return "PrefixSet" }
    if yname == "match-set-options" { return "MatchSetOptions" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-set"] = config.PrefixSet
    leafs["match-set-options"] = config.MatchSetOptions
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetParentYangName() string { return "match-prefix-set" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State
// Operational state data for a prefix-set condition
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined prefix set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_PrefixSetName
    PrefixSet interface{}

    // Optional parameter that governs the behaviour of the match operation.  This
    // leaf only supports matching on ANY member of the set or inverting the
    // match.  Matching on ALL is not supported). The type is
    // MatchSetOptionsRestrictedType.
    MatchSetOptions interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetGoName(yname string) string {
    if yname == "prefix-set" { return "PrefixSet" }
    if yname == "match-set-options" { return "MatchSetOptions" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-set"] = state.PrefixSet
    leafs["match-set-options"] = state.MatchSetOptions
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetParentYangName() string { return "match-prefix-set" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet
// Match a referenced neighbor set according to the logic
// defined in the match-set-options-leaf
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data .
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config

    // Operational state data .
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State
}

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetFilter() yfilter.YFilter { return matchNeighborSet.YFilter }

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) SetFilter(yf yfilter.YFilter) { matchNeighborSet.YFilter = yf }

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetSegmentPath() string {
    return "match-neighbor-set"
}

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &matchNeighborSet.Config
    }
    if childYangName == "state" {
        return &matchNeighborSet.State
    }
    return nil
}

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &matchNeighborSet.Config
    children["state"] = &matchNeighborSet.State
    return children
}

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetBundleName() string { return "openconfig" }

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetYangName() string { return "match-neighbor-set" }

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) SetParent(parent types.Entity) { matchNeighborSet.parent = parent }

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetParent() types.Entity { return matchNeighborSet.parent }

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetParentYangName() string { return "conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config
// Configuration data 
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined neighbor set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_NeighborSetName
    NeighborSet interface{}

    // Optional parameter that governs the behaviour of the match operation.  This
    // leaf only supports matching on ANY member of the set or inverting the
    // match.  Matching on ALL is not supported). The type is
    // MatchSetOptionsRestrictedType.
    MatchSetOptions interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetGoName(yname string) string {
    if yname == "neighbor-set" { return "NeighborSet" }
    if yname == "match-set-options" { return "MatchSetOptions" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-set"] = config.NeighborSet
    leafs["match-set-options"] = config.MatchSetOptions
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetParentYangName() string { return "match-neighbor-set" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State
// Operational state data 
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined neighbor set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_NeighborSetName
    NeighborSet interface{}

    // Optional parameter that governs the behaviour of the match operation.  This
    // leaf only supports matching on ANY member of the set or inverting the
    // match.  Matching on ALL is not supported). The type is
    // MatchSetOptionsRestrictedType.
    MatchSetOptions interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetGoName(yname string) string {
    if yname == "neighbor-set" { return "NeighborSet" }
    if yname == "match-set-options" { return "MatchSetOptions" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-set"] = state.NeighborSet
    leafs["match-set-options"] = state.MatchSetOptions
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetParentYangName() string { return "match-neighbor-set" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet
// Match a referenced tag set according to the logic defined
// in the match-options-set leaf
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for tag-set conditions.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config

    // Operational state data tag-set conditions.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State
}

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetFilter() yfilter.YFilter { return matchTagSet.YFilter }

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) SetFilter(yf yfilter.YFilter) { matchTagSet.YFilter = yf }

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetSegmentPath() string {
    return "match-tag-set"
}

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &matchTagSet.Config
    }
    if childYangName == "state" {
        return &matchTagSet.State
    }
    return nil
}

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &matchTagSet.Config
    children["state"] = &matchTagSet.State
    return children
}

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetBundleName() string { return "openconfig" }

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetYangName() string { return "match-tag-set" }

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) SetParent(parent types.Entity) { matchTagSet.parent = parent }

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetParent() types.Entity { return matchTagSet.parent }

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetParentYangName() string { return "conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config
// Configuration data for tag-set conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined tag set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_TagSets_TagSet_TagSetName
    TagSet interface{}

    // Optional parameter that governs the behaviour of the match operation.  This
    // leaf only supports matching on ANY member of the set or inverting the
    // match.  Matching on ALL is not supported). The type is
    // MatchSetOptionsRestrictedType.
    MatchSetOptions interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetGoName(yname string) string {
    if yname == "tag-set" { return "TagSet" }
    if yname == "match-set-options" { return "MatchSetOptions" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tag-set"] = config.TagSet
    leafs["match-set-options"] = config.MatchSetOptions
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetParentYangName() string { return "match-tag-set" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State
// Operational state data tag-set conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined tag set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_TagSets_TagSet_TagSetName
    TagSet interface{}

    // Optional parameter that governs the behaviour of the match operation.  This
    // leaf only supports matching on ANY member of the set or inverting the
    // match.  Matching on ALL is not supported). The type is
    // MatchSetOptionsRestrictedType.
    MatchSetOptions interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetGoName(yname string) string {
    if yname == "tag-set" { return "TagSet" }
    if yname == "match-set-options" { return "MatchSetOptions" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tag-set"] = state.TagSet
    leafs["match-set-options"] = state.MatchSetOptions
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetParentYangName() string { return "match-tag-set" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions
// Policy conditions for IGP attributes
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetFilter() yfilter.YFilter { return igpConditions.YFilter }

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) SetFilter(yf yfilter.YFilter) { igpConditions.YFilter = yf }

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetGoName(yname string) string {
    return ""
}

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetSegmentPath() string {
    return "igp-conditions"
}

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetBundleName() string { return "openconfig" }

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetYangName() string { return "igp-conditions" }

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) SetParent(parent types.Entity) { igpConditions.parent = parent }

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetParent() types.Entity { return igpConditions.parent }

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetParentYangName() string { return "conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions
// Top-level container 
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for BGP-specific policy conditions.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config

    // Operational state data for BGP-specific policy conditions.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State

    // Value and comparison operations for conditions based on the number of
    // communities in the route update.
    CommunityCount RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount

    // Value and comparison operations for conditions based on the length of the
    // AS path in the route update.
    AsPathLength RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength

    // Top-level container for match conditions on communities. Match a referenced
    // community-set according to the logic defined in the match-set-options leaf.
    MatchCommunitySet RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet

    // Match a referenced extended community-set according to the logic defined in
    // the match-set-options leaf.
    MatchExtCommunitySet RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet

    // Match a referenced as-path set according to the logic defined in the
    // match-set-options leaf.
    MatchAsPathSet RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet
}

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetFilter() yfilter.YFilter { return bgpConditions.YFilter }

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) SetFilter(yf yfilter.YFilter) { bgpConditions.YFilter = yf }

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "community-count" { return "CommunityCount" }
    if yname == "as-path-length" { return "AsPathLength" }
    if yname == "match-community-set" { return "MatchCommunitySet" }
    if yname == "match-ext-community-set" { return "MatchExtCommunitySet" }
    if yname == "match-as-path-set" { return "MatchAsPathSet" }
    return ""
}

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetSegmentPath() string {
    return "openconfig-bgp-policy:bgp-conditions"
}

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &bgpConditions.Config
    }
    if childYangName == "state" {
        return &bgpConditions.State
    }
    if childYangName == "community-count" {
        return &bgpConditions.CommunityCount
    }
    if childYangName == "as-path-length" {
        return &bgpConditions.AsPathLength
    }
    if childYangName == "match-community-set" {
        return &bgpConditions.MatchCommunitySet
    }
    if childYangName == "match-ext-community-set" {
        return &bgpConditions.MatchExtCommunitySet
    }
    if childYangName == "match-as-path-set" {
        return &bgpConditions.MatchAsPathSet
    }
    return nil
}

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &bgpConditions.Config
    children["state"] = &bgpConditions.State
    children["community-count"] = &bgpConditions.CommunityCount
    children["as-path-length"] = &bgpConditions.AsPathLength
    children["match-community-set"] = &bgpConditions.MatchCommunitySet
    children["match-ext-community-set"] = &bgpConditions.MatchExtCommunitySet
    children["match-as-path-set"] = &bgpConditions.MatchAsPathSet
    return children
}

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetBundleName() string { return "openconfig" }

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetYangName() string { return "bgp-conditions" }

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) SetParent(parent types.Entity) { bgpConditions.parent = parent }

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetParent() types.Entity { return bgpConditions.parent }

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetParentYangName() string { return "conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config
// Configuration data for BGP-specific policy conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Condition to check if the received MED value is equal to the specified
    // value. The type is interface{} with range: 0..4294967295.
    MedEq interface{}

    // Condition to check if the route origin is equal to the specified value. The
    // type is BgpOriginAttrType.
    OriginEq interface{}

    // List of next hop addresses to check for in the route update. The type is
    // one of the following types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHopIn []interface{}

    // List of address families which the NLRI may be within. The type is slice of
    // [u'L2VPNEVPN', u'L2VPNVPLS', u'IPV4UNICAST', u'L3VPNIPV6MULTICAST',
    // u'L3VPNIPV6UNICAST', u'L3VPNIPV4UNICAST', u'L3VPNIPV4MULTICAST',
    // u'IPV4LABELEDUNICAST', u'IPV6UNICAST', u'IPV6LABELEDUNICAST'].
    AfiSafiIn []interface{}

    // Condition to check if the local pref attribute is equal to the specified
    // value. The type is interface{} with range: 0..4294967295.
    LocalPrefEq interface{}

    // Condition to check the route type in the route update. The type is
    // RouteType.
    RouteType interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetGoName(yname string) string {
    if yname == "med-eq" { return "MedEq" }
    if yname == "origin-eq" { return "OriginEq" }
    if yname == "next-hop-in" { return "NextHopIn" }
    if yname == "afi-safi-in" { return "AfiSafiIn" }
    if yname == "local-pref-eq" { return "LocalPrefEq" }
    if yname == "route-type" { return "RouteType" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["med-eq"] = config.MedEq
    leafs["origin-eq"] = config.OriginEq
    leafs["next-hop-in"] = config.NextHopIn
    leafs["afi-safi-in"] = config.AfiSafiIn
    leafs["local-pref-eq"] = config.LocalPrefEq
    leafs["route-type"] = config.RouteType
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetParentYangName() string { return "bgp-conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config_RouteType represents Condition to check the route type in the route update
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config_RouteType string

const (
    // route type is internal
    RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config_RouteType_INTERNAL RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config_RouteType = "INTERNAL"

    // route type is external
    RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config_RouteType_EXTERNAL RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config_RouteType = "EXTERNAL"
)

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State
// Operational state data for BGP-specific policy
// conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Condition to check if the received MED value is equal to the specified
    // value. The type is interface{} with range: 0..4294967295.
    MedEq interface{}

    // Condition to check if the route origin is equal to the specified value. The
    // type is BgpOriginAttrType.
    OriginEq interface{}

    // List of next hop addresses to check for in the route update. The type is
    // one of the following types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHopIn []interface{}

    // List of address families which the NLRI may be within. The type is slice of
    // [u'L2VPNEVPN', u'L2VPNVPLS', u'IPV4UNICAST', u'L3VPNIPV6MULTICAST',
    // u'L3VPNIPV6UNICAST', u'L3VPNIPV4UNICAST', u'L3VPNIPV4MULTICAST',
    // u'IPV4LABELEDUNICAST', u'IPV6UNICAST', u'IPV6LABELEDUNICAST'].
    AfiSafiIn []interface{}

    // Condition to check if the local pref attribute is equal to the specified
    // value. The type is interface{} with range: 0..4294967295.
    LocalPrefEq interface{}

    // Condition to check the route type in the route update. The type is
    // RouteType.
    RouteType interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetGoName(yname string) string {
    if yname == "med-eq" { return "MedEq" }
    if yname == "origin-eq" { return "OriginEq" }
    if yname == "next-hop-in" { return "NextHopIn" }
    if yname == "afi-safi-in" { return "AfiSafiIn" }
    if yname == "local-pref-eq" { return "LocalPrefEq" }
    if yname == "route-type" { return "RouteType" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["med-eq"] = state.MedEq
    leafs["origin-eq"] = state.OriginEq
    leafs["next-hop-in"] = state.NextHopIn
    leafs["afi-safi-in"] = state.AfiSafiIn
    leafs["local-pref-eq"] = state.LocalPrefEq
    leafs["route-type"] = state.RouteType
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetParentYangName() string { return "bgp-conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State_RouteType represents Condition to check the route type in the route update
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State_RouteType string

const (
    // route type is internal
    RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State_RouteType_INTERNAL RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State_RouteType = "INTERNAL"

    // route type is external
    RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State_RouteType_EXTERNAL RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State_RouteType = "EXTERNAL"
)

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount
// Value and comparison operations for conditions based on the
// number of communities in the route update
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for community count condition.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config

    // Operational state data for community count condition.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State
}

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetFilter() yfilter.YFilter { return communityCount.YFilter }

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) SetFilter(yf yfilter.YFilter) { communityCount.YFilter = yf }

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetSegmentPath() string {
    return "community-count"
}

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &communityCount.Config
    }
    if childYangName == "state" {
        return &communityCount.State
    }
    return nil
}

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &communityCount.Config
    children["state"] = &communityCount.State
    return children
}

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetBundleName() string { return "openconfig" }

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetYangName() string { return "community-count" }

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) SetParent(parent types.Entity) { communityCount.parent = parent }

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetParent() types.Entity { return communityCount.parent }

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetParentYangName() string { return "bgp-conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config
// Configuration data for community count condition
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type of comparison to be performed. The type is one of the following:
    // ATTRIBUTEGEATTRIBUTEEQATTRIBUTELE.
    Operator interface{}

    // value to compare with the community count. The type is interface{} with
    // range: 0..4294967295.
    Value interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = config.Operator
    leafs["value"] = config.Value
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetParentYangName() string { return "community-count" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State
// Operational state data for community count condition
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type of comparison to be performed. The type is one of the following:
    // ATTRIBUTEGEATTRIBUTEEQATTRIBUTELE.
    Operator interface{}

    // value to compare with the community count. The type is interface{} with
    // range: 0..4294967295.
    Value interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = state.Operator
    leafs["value"] = state.Value
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetParentYangName() string { return "community-count" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength
// Value and comparison operations for conditions based on the
// length of the AS path in the route update
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for AS path length condition.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config

    // Operational state data for AS path length condition.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State
}

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetFilter() yfilter.YFilter { return asPathLength.YFilter }

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) SetFilter(yf yfilter.YFilter) { asPathLength.YFilter = yf }

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetSegmentPath() string {
    return "as-path-length"
}

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &asPathLength.Config
    }
    if childYangName == "state" {
        return &asPathLength.State
    }
    return nil
}

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &asPathLength.Config
    children["state"] = &asPathLength.State
    return children
}

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetBundleName() string { return "openconfig" }

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetYangName() string { return "as-path-length" }

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) SetParent(parent types.Entity) { asPathLength.parent = parent }

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetParent() types.Entity { return asPathLength.parent }

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetParentYangName() string { return "bgp-conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config
// Configuration data for AS path length condition
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type of comparison to be performed. The type is one of the following:
    // ATTRIBUTEGEATTRIBUTEEQATTRIBUTELE.
    Operator interface{}

    // value to compare with the community count. The type is interface{} with
    // range: 0..4294967295.
    Value interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = config.Operator
    leafs["value"] = config.Value
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetParentYangName() string { return "as-path-length" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State
// Operational state data for AS path length condition
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type of comparison to be performed. The type is one of the following:
    // ATTRIBUTEGEATTRIBUTEEQATTRIBUTELE.
    Operator interface{}

    // value to compare with the community count. The type is interface{} with
    // range: 0..4294967295.
    Value interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = state.Operator
    leafs["value"] = state.Value
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetParentYangName() string { return "as-path-length" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet
// Top-level container for match conditions on communities.
// Match a referenced community-set according to the logic
// defined in the match-set-options leaf
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for match conditions on communities.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config

    // Operational state data .
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State
}

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetFilter() yfilter.YFilter { return matchCommunitySet.YFilter }

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) SetFilter(yf yfilter.YFilter) { matchCommunitySet.YFilter = yf }

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetSegmentPath() string {
    return "match-community-set"
}

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &matchCommunitySet.Config
    }
    if childYangName == "state" {
        return &matchCommunitySet.State
    }
    return nil
}

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &matchCommunitySet.Config
    children["state"] = &matchCommunitySet.State
    return children
}

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetBundleName() string { return "openconfig" }

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetYangName() string { return "match-community-set" }

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) SetParent(parent types.Entity) { matchCommunitySet.parent = parent }

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetParent() types.Entity { return matchCommunitySet.parent }

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetParentYangName() string { return "bgp-conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config
// Configuration data for match conditions on communities
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined community set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_CommunitySetName
    CommunitySet interface{}

    // Optional parameter that governs the behaviour of the match operation. The
    // type is MatchSetOptionsType.
    MatchSetOptions interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetGoName(yname string) string {
    if yname == "community-set" { return "CommunitySet" }
    if yname == "match-set-options" { return "MatchSetOptions" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-set"] = config.CommunitySet
    leafs["match-set-options"] = config.MatchSetOptions
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetParentYangName() string { return "match-community-set" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State
// Operational state data 
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined community set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_CommunitySetName
    CommunitySet interface{}

    // Optional parameter that governs the behaviour of the match operation. The
    // type is MatchSetOptionsType.
    MatchSetOptions interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetGoName(yname string) string {
    if yname == "community-set" { return "CommunitySet" }
    if yname == "match-set-options" { return "MatchSetOptions" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-set"] = state.CommunitySet
    leafs["match-set-options"] = state.MatchSetOptions
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetParentYangName() string { return "match-community-set" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet
// Match a referenced extended community-set according to the
// logic defined in the match-set-options leaf
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for match conditions on extended communities.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config

    // Operational state data for match conditions on extended communities.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State
}

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetFilter() yfilter.YFilter { return matchExtCommunitySet.YFilter }

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) SetFilter(yf yfilter.YFilter) { matchExtCommunitySet.YFilter = yf }

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetSegmentPath() string {
    return "match-ext-community-set"
}

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &matchExtCommunitySet.Config
    }
    if childYangName == "state" {
        return &matchExtCommunitySet.State
    }
    return nil
}

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &matchExtCommunitySet.Config
    children["state"] = &matchExtCommunitySet.State
    return children
}

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetBundleName() string { return "openconfig" }

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetYangName() string { return "match-ext-community-set" }

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) SetParent(parent types.Entity) { matchExtCommunitySet.parent = parent }

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetParent() types.Entity { return matchExtCommunitySet.parent }

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetParentYangName() string { return "bgp-conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config
// Configuration data for match conditions on extended
// communities
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined extended community set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_ExtCommunitySetName
    ExtCommunitySet interface{}

    // Optional parameter that governs the behaviour of the match operation. The
    // type is MatchSetOptionsType.
    MatchSetOptions interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetGoName(yname string) string {
    if yname == "ext-community-set" { return "ExtCommunitySet" }
    if yname == "match-set-options" { return "MatchSetOptions" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ext-community-set"] = config.ExtCommunitySet
    leafs["match-set-options"] = config.MatchSetOptions
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetParentYangName() string { return "match-ext-community-set" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State
// Operational state data for match conditions on extended
// communities
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined extended community set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_ExtCommunitySetName
    ExtCommunitySet interface{}

    // Optional parameter that governs the behaviour of the match operation. The
    // type is MatchSetOptionsType.
    MatchSetOptions interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetGoName(yname string) string {
    if yname == "ext-community-set" { return "ExtCommunitySet" }
    if yname == "match-set-options" { return "MatchSetOptions" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ext-community-set"] = state.ExtCommunitySet
    leafs["match-set-options"] = state.MatchSetOptions
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetParentYangName() string { return "match-ext-community-set" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet
// Match a referenced as-path set according to the logic
// defined in the match-set-options leaf
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for match conditions on AS path set.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config

    // Operational state data for match conditions on AS path set.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State
}

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetFilter() yfilter.YFilter { return matchAsPathSet.YFilter }

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) SetFilter(yf yfilter.YFilter) { matchAsPathSet.YFilter = yf }

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetSegmentPath() string {
    return "match-as-path-set"
}

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &matchAsPathSet.Config
    }
    if childYangName == "state" {
        return &matchAsPathSet.State
    }
    return nil
}

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &matchAsPathSet.Config
    children["state"] = &matchAsPathSet.State
    return children
}

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetBundleName() string { return "openconfig" }

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetYangName() string { return "match-as-path-set" }

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) SetParent(parent types.Entity) { matchAsPathSet.parent = parent }

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetParent() types.Entity { return matchAsPathSet.parent }

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetParentYangName() string { return "bgp-conditions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config
// Configuration data for match conditions on AS path set
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined AS path set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_AsPathSetName
    AsPathSet interface{}

    // Optional parameter that governs the behaviour of the match operation. The
    // type is MatchSetOptionsType.
    MatchSetOptions interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetGoName(yname string) string {
    if yname == "as-path-set" { return "AsPathSet" }
    if yname == "match-set-options" { return "MatchSetOptions" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-path-set"] = config.AsPathSet
    leafs["match-set-options"] = config.MatchSetOptions
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetParentYangName() string { return "match-as-path-set" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State
// Operational state data for match conditions on AS
// path set
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined AS path set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_AsPathSetName
    AsPathSet interface{}

    // Optional parameter that governs the behaviour of the match operation. The
    // type is MatchSetOptionsType.
    MatchSetOptions interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetGoName(yname string) string {
    if yname == "as-path-set" { return "AsPathSet" }
    if yname == "match-set-options" { return "MatchSetOptions" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-path-set"] = state.AsPathSet
    leafs["match-set-options"] = state.MatchSetOptions
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetParentYangName() string { return "match-as-path-set" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions
// Top-level container for policy action statements
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for policy actions.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config

    // Operational state data for policy actions.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State

    // Actions to set IGP route attributes; these actions apply to multiple IGPs.
    IgpActions RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions

    // Top-level container for BGP-specific actions.
    BgpActions RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions
}

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetFilter() yfilter.YFilter { return actions.YFilter }

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) SetFilter(yf yfilter.YFilter) { actions.YFilter = yf }

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "igp-actions" { return "IgpActions" }
    if yname == "openconfig-bgp-policy:bgp-actions" { return "BgpActions" }
    return ""
}

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetSegmentPath() string {
    return "actions"
}

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &actions.Config
    }
    if childYangName == "state" {
        return &actions.State
    }
    if childYangName == "igp-actions" {
        return &actions.IgpActions
    }
    if childYangName == "openconfig-bgp-policy:bgp-actions" {
        return &actions.BgpActions
    }
    return nil
}

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &actions.Config
    children["state"] = &actions.State
    children["igp-actions"] = &actions.IgpActions
    children["openconfig-bgp-policy:bgp-actions"] = &actions.BgpActions
    return children
}

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetBundleName() string { return "openconfig" }

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetYangName() string { return "actions" }

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) SetParent(parent types.Entity) { actions.parent = parent }

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetParent() types.Entity { return actions.parent }

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetParentYangName() string { return "statement" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config
// Configuration data for policy actions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // accepts the route into the routing table. The type is interface{}.
    AcceptRoute interface{}

    // rejects the route. The type is interface{}.
    RejectRoute interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetGoName(yname string) string {
    if yname == "accept-route" { return "AcceptRoute" }
    if yname == "reject-route" { return "RejectRoute" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["accept-route"] = config.AcceptRoute
    leafs["reject-route"] = config.RejectRoute
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetParentYangName() string { return "actions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State
// Operational state data for policy actions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // accepts the route into the routing table. The type is interface{}.
    AcceptRoute interface{}

    // rejects the route. The type is interface{}.
    RejectRoute interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetGoName(yname string) string {
    if yname == "accept-route" { return "AcceptRoute" }
    if yname == "reject-route" { return "RejectRoute" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["accept-route"] = state.AcceptRoute
    leafs["reject-route"] = state.RejectRoute
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetParentYangName() string { return "actions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions
// Actions to set IGP route attributes; these actions
// apply to multiple IGPs
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data .
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config

    // Operational state data .
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State
}

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetFilter() yfilter.YFilter { return igpActions.YFilter }

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) SetFilter(yf yfilter.YFilter) { igpActions.YFilter = yf }

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetSegmentPath() string {
    return "igp-actions"
}

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &igpActions.Config
    }
    if childYangName == "state" {
        return &igpActions.State
    }
    return nil
}

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &igpActions.Config
    children["state"] = &igpActions.State
    return children
}

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetBundleName() string { return "openconfig" }

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetYangName() string { return "igp-actions" }

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) SetParent(parent types.Entity) { igpActions.parent = parent }

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetParent() types.Entity { return igpActions.parent }

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetParentYangName() string { return "actions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config
// Configuration data 
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the tag value for OSPF or IS-IS routes. The type is one of the
    // following types: int with range: 0..4294967295, or string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    SetTag interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetGoName(yname string) string {
    if yname == "set-tag" { return "SetTag" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-tag"] = config.SetTag
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetParentYangName() string { return "igp-actions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State
// Operational state data 
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the tag value for OSPF or IS-IS routes. The type is one of the
    // following types: int with range: 0..4294967295, or string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    SetTag interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetGoName(yname string) string {
    if yname == "set-tag" { return "SetTag" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-tag"] = state.SetTag
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetParentYangName() string { return "igp-actions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions
// Top-level container for BGP-specific actions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for BGP-specific actions.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config

    // Operational state data for BGP-specific actions.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State

    // action to prepend local AS number to the AS-path a specified number of
    // times.
    SetAsPathPrepend RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend

    // Action to set the community attributes of the route, along with options to
    // modify how the community is modified. Communities may be set using an
    // inline list OR reference to an existing defined set (not both).
    SetCommunity RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity

    // Action to set the extended community attributes of the route, along with
    // options to modify how the community is modified. Extended communities may
    // be set using an inline list OR a reference to an existing defined set (but
    // not both).
    SetExtCommunity RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity
}

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetFilter() yfilter.YFilter { return bgpActions.YFilter }

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) SetFilter(yf yfilter.YFilter) { bgpActions.YFilter = yf }

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "set-as-path-prepend" { return "SetAsPathPrepend" }
    if yname == "set-community" { return "SetCommunity" }
    if yname == "set-ext-community" { return "SetExtCommunity" }
    return ""
}

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetSegmentPath() string {
    return "openconfig-bgp-policy:bgp-actions"
}

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &bgpActions.Config
    }
    if childYangName == "state" {
        return &bgpActions.State
    }
    if childYangName == "set-as-path-prepend" {
        return &bgpActions.SetAsPathPrepend
    }
    if childYangName == "set-community" {
        return &bgpActions.SetCommunity
    }
    if childYangName == "set-ext-community" {
        return &bgpActions.SetExtCommunity
    }
    return nil
}

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &bgpActions.Config
    children["state"] = &bgpActions.State
    children["set-as-path-prepend"] = &bgpActions.SetAsPathPrepend
    children["set-community"] = &bgpActions.SetCommunity
    children["set-ext-community"] = &bgpActions.SetExtCommunity
    return children
}

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetBundleName() string { return "openconfig" }

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetYangName() string { return "bgp-actions" }

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) SetParent(parent types.Entity) { bgpActions.parent = parent }

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetParent() types.Entity { return bgpActions.parent }

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetParentYangName() string { return "actions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config
// Configuration data for BGP-specific actions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // set the origin attribute to the specified value. The type is
    // BgpOriginAttrType.
    SetRouteOrigin interface{}

    // set the local pref attribute on the route update. The type is interface{}
    // with range: 0..4294967295.
    SetLocalPref interface{}

    // set the next-hop attribute in the route update. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.,
    // or enumeration BgpNextHopType.
    SetNextHop interface{}

    // set the med metric attribute in the route update. The type is one of the
    // following types: int with range: 0..4294967295, or string with pattern:
    // ^[+-][0-9]+, or enumeration BgpSetMedType.
    SetMed interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetGoName(yname string) string {
    if yname == "set-route-origin" { return "SetRouteOrigin" }
    if yname == "set-local-pref" { return "SetLocalPref" }
    if yname == "set-next-hop" { return "SetNextHop" }
    if yname == "set-med" { return "SetMed" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-route-origin"] = config.SetRouteOrigin
    leafs["set-local-pref"] = config.SetLocalPref
    leafs["set-next-hop"] = config.SetNextHop
    leafs["set-med"] = config.SetMed
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetParentYangName() string { return "bgp-actions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State
// Operational state data for BGP-specific actions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // set the origin attribute to the specified value. The type is
    // BgpOriginAttrType.
    SetRouteOrigin interface{}

    // set the local pref attribute on the route update. The type is interface{}
    // with range: 0..4294967295.
    SetLocalPref interface{}

    // set the next-hop attribute in the route update. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.,
    // or enumeration BgpNextHopType.
    SetNextHop interface{}

    // set the med metric attribute in the route update. The type is one of the
    // following types: int with range: 0..4294967295, or string with pattern:
    // ^[+-][0-9]+, or enumeration BgpSetMedType.
    SetMed interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetGoName(yname string) string {
    if yname == "set-route-origin" { return "SetRouteOrigin" }
    if yname == "set-local-pref" { return "SetLocalPref" }
    if yname == "set-next-hop" { return "SetNextHop" }
    if yname == "set-med" { return "SetMed" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-route-origin"] = state.SetRouteOrigin
    leafs["set-local-pref"] = state.SetLocalPref
    leafs["set-next-hop"] = state.SetNextHop
    leafs["set-med"] = state.SetMed
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetParentYangName() string { return "bgp-actions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend
// action to prepend local AS number to the AS-path a
// specified number of times
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for the AS path prepend action.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config

    // Operational state data for the AS path prepend action.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State
}

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetFilter() yfilter.YFilter { return setAsPathPrepend.YFilter }

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) SetFilter(yf yfilter.YFilter) { setAsPathPrepend.YFilter = yf }

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetSegmentPath() string {
    return "set-as-path-prepend"
}

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &setAsPathPrepend.Config
    }
    if childYangName == "state" {
        return &setAsPathPrepend.State
    }
    return nil
}

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &setAsPathPrepend.Config
    children["state"] = &setAsPathPrepend.State
    return children
}

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetBundleName() string { return "openconfig" }

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetYangName() string { return "set-as-path-prepend" }

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) SetParent(parent types.Entity) { setAsPathPrepend.parent = parent }

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetParent() types.Entity { return setAsPathPrepend.parent }

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetParentYangName() string { return "bgp-actions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config
// Configuration data for the AS path prepend action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of times to prepend the local AS number to the AS path.  The value
    // should be between 1 and the maximum supported by the implementation. The
    // type is interface{} with range: 1..255.
    RepeatN interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetGoName(yname string) string {
    if yname == "repeat-n" { return "RepeatN" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["repeat-n"] = config.RepeatN
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetParentYangName() string { return "set-as-path-prepend" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State
// Operational state data for the AS path prepend action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of times to prepend the local AS number to the AS path.  The value
    // should be between 1 and the maximum supported by the implementation. The
    // type is interface{} with range: 1..255.
    RepeatN interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetGoName(yname string) string {
    if yname == "repeat-n" { return "RepeatN" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["repeat-n"] = state.RepeatN
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetParentYangName() string { return "set-as-path-prepend" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity
// Action to set the community attributes of the route, along
// with options to modify how the community is modified.
// Communities may be set using an inline list OR
// reference to an existing defined set (not both).
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for the set-community action.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config

    // Operational state data for the set-community action.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State

    // Set the community values for the action inline with a list.
    Inline RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline

    // Provide a reference to a defined community set for the set-community
    // action.
    Reference RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference
}

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetFilter() yfilter.YFilter { return setCommunity.YFilter }

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) SetFilter(yf yfilter.YFilter) { setCommunity.YFilter = yf }

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "inline" { return "Inline" }
    if yname == "reference" { return "Reference" }
    return ""
}

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetSegmentPath() string {
    return "set-community"
}

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &setCommunity.Config
    }
    if childYangName == "state" {
        return &setCommunity.State
    }
    if childYangName == "inline" {
        return &setCommunity.Inline
    }
    if childYangName == "reference" {
        return &setCommunity.Reference
    }
    return nil
}

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &setCommunity.Config
    children["state"] = &setCommunity.State
    children["inline"] = &setCommunity.Inline
    children["reference"] = &setCommunity.Reference
    return children
}

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetBundleName() string { return "openconfig" }

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetYangName() string { return "set-community" }

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) SetParent(parent types.Entity) { setCommunity.parent = parent }

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetParent() types.Entity { return setCommunity.parent }

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetParentYangName() string { return "bgp-actions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config
// Configuration data for the set-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates the method used to specify the extended communities for the
    // set-ext-community action. The type is Method.
    Method interface{}

    // Options for modifying the community attribute with the specified values. 
    // These options apply to both methods of setting the community attribute. The
    // type is BgpSetCommunityOptionType.
    Options interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetGoName(yname string) string {
    if yname == "method" { return "Method" }
    if yname == "options" { return "Options" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method"] = config.Method
    leafs["options"] = config.Options
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetParentYangName() string { return "set-community" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config_Method represents communities for the set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config_Method string

const (
    // The extended communities are specified inline as a
    // list
    RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config_Method_INLINE RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config_Method = "INLINE"

    // The extended communities are specified by referencing a
    // defined ext-community set
    RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config_Method_REFERENCE RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config_Method = "REFERENCE"
)

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State
// Operational state data for the set-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates the method used to specify the extended communities for the
    // set-ext-community action. The type is Method.
    Method interface{}

    // Options for modifying the community attribute with the specified values. 
    // These options apply to both methods of setting the community attribute. The
    // type is BgpSetCommunityOptionType.
    Options interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetGoName(yname string) string {
    if yname == "method" { return "Method" }
    if yname == "options" { return "Options" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method"] = state.Method
    leafs["options"] = state.Options
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetParentYangName() string { return "set-community" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State_Method represents communities for the set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State_Method string

const (
    // The extended communities are specified inline as a
    // list
    RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State_Method_INLINE RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State_Method = "INLINE"

    // The extended communities are specified by referencing a
    // defined ext-community set
    RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State_Method_REFERENCE RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State_Method = "REFERENCE"
)

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline
// Set the community values for the action inline with
// a list.
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data or inline specification of set-community action.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config

    // Operational state data or inline specification of set-community action.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetFilter() yfilter.YFilter { return inline.YFilter }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) SetFilter(yf yfilter.YFilter) { inline.YFilter = yf }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetSegmentPath() string {
    return "inline"
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &inline.Config
    }
    if childYangName == "state" {
        return &inline.State
    }
    return nil
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &inline.Config
    children["state"] = &inline.State
    return children
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetBundleName() string { return "openconfig" }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetYangName() string { return "inline" }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) SetParent(parent types.Entity) { inline.parent = parent }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetParent() types.Entity { return inline.parent }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetParentYangName() string { return "set-community" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config
// Configuration data or inline specification of set-community
// action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the community values for the update inline with a list. The type is one
    // of the following types: slice of int with range: 65536..4294901759, or
    // slice of string with pattern: ([0-9]+:[0-9]+)., or slice of  
    // :go:struct:`BGPWELLKNOWNSTDCOMMUNITY
    // <ydk/models/bgp_types/BGPWELLKNOWNSTDCOMMUNITY>`.
    Communities []interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetGoName(yname string) string {
    if yname == "communities" { return "Communities" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["communities"] = config.Communities
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetParentYangName() string { return "inline" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State
// Operational state data or inline specification of
// set-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the community values for the update inline with a list. The type is one
    // of the following types: slice of int with range: 65536..4294901759, or
    // slice of string with pattern: ([0-9]+:[0-9]+)., or slice of  
    // :go:struct:`BGPWELLKNOWNSTDCOMMUNITY
    // <ydk/models/bgp_types/BGPWELLKNOWNSTDCOMMUNITY>`.
    Communities []interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetGoName(yname string) string {
    if yname == "communities" { return "Communities" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["communities"] = state.Communities
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetParentYangName() string { return "inline" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference
// Provide a reference to a defined community set for the
// set-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for referening a community-set in the set-community
    // action.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config

    // Operational state data for referening a community-set in the set-community
    // action.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &reference.Config
    }
    if childYangName == "state" {
        return &reference.State
    }
    return nil
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &reference.Config
    children["state"] = &reference.State
    return children
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetBundleName() string { return "openconfig" }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetParentYangName() string { return "set-community" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config
// Configuration data for referening a community-set in the
// set-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined community set by name. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_CommunitySetName
    CommunitySetRef interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetGoName(yname string) string {
    if yname == "community-set-ref" { return "CommunitySetRef" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-set-ref"] = config.CommunitySetRef
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetParentYangName() string { return "reference" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State
// Operational state data for referening a community-set
// in the set-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined community set by name. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_CommunitySetName
    CommunitySetRef interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetGoName(yname string) string {
    if yname == "community-set-ref" { return "CommunitySetRef" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-set-ref"] = state.CommunitySetRef
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetParentYangName() string { return "reference" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity
// Action to set the extended community attributes of the
// route, along with options to modify how the community is
// modified. Extended communities may be set using an inline
// list OR a reference to an existing defined set (but not
// both).
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for the set-ext-community action.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config

    // Operational state data for the set-ext-community action.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State

    // Set the extended community values for the action inline with a list.
    Inline RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline

    // Provide a reference to an extended community set for the set-ext-community
    // action.
    Reference RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference
}

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetFilter() yfilter.YFilter { return setExtCommunity.YFilter }

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) SetFilter(yf yfilter.YFilter) { setExtCommunity.YFilter = yf }

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "inline" { return "Inline" }
    if yname == "reference" { return "Reference" }
    return ""
}

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetSegmentPath() string {
    return "set-ext-community"
}

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &setExtCommunity.Config
    }
    if childYangName == "state" {
        return &setExtCommunity.State
    }
    if childYangName == "inline" {
        return &setExtCommunity.Inline
    }
    if childYangName == "reference" {
        return &setExtCommunity.Reference
    }
    return nil
}

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &setExtCommunity.Config
    children["state"] = &setExtCommunity.State
    children["inline"] = &setExtCommunity.Inline
    children["reference"] = &setExtCommunity.Reference
    return children
}

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetBundleName() string { return "openconfig" }

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetYangName() string { return "set-ext-community" }

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) SetParent(parent types.Entity) { setExtCommunity.parent = parent }

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetParent() types.Entity { return setExtCommunity.parent }

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetParentYangName() string { return "bgp-actions" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config
// Configuration data for the set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates the method used to specify the extended communities for the
    // set-ext-community action. The type is Method.
    Method interface{}

    // Options for modifying the community attribute with the specified values. 
    // These options apply to both methods of setting the community attribute. The
    // type is BgpSetCommunityOptionType.
    Options interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetGoName(yname string) string {
    if yname == "method" { return "Method" }
    if yname == "options" { return "Options" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method"] = config.Method
    leafs["options"] = config.Options
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetParentYangName() string { return "set-ext-community" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config_Method represents communities for the set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config_Method string

const (
    // The extended communities are specified inline as a
    // list
    RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config_Method_INLINE RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config_Method = "INLINE"

    // The extended communities are specified by referencing a
    // defined ext-community set
    RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config_Method_REFERENCE RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config_Method = "REFERENCE"
)

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State
// Operational state data for the set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates the method used to specify the extended communities for the
    // set-ext-community action. The type is Method.
    Method interface{}

    // Options for modifying the community attribute with the specified values. 
    // These options apply to both methods of setting the community attribute. The
    // type is BgpSetCommunityOptionType.
    Options interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetGoName(yname string) string {
    if yname == "method" { return "Method" }
    if yname == "options" { return "Options" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method"] = state.Method
    leafs["options"] = state.Options
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetParentYangName() string { return "set-ext-community" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State_Method represents communities for the set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State_Method string

const (
    // The extended communities are specified inline as a
    // list
    RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State_Method_INLINE RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State_Method = "INLINE"

    // The extended communities are specified by referencing a
    // defined ext-community set
    RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State_Method_REFERENCE RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State_Method = "REFERENCE"
)

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline
// Set the extended community values for the action inline with
// a list.
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data or inline specification of set-ext-community action.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config

    // Operational state data or inline specification of set-ext-community action.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetFilter() yfilter.YFilter { return inline.YFilter }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) SetFilter(yf yfilter.YFilter) { inline.YFilter = yf }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetSegmentPath() string {
    return "inline"
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &inline.Config
    }
    if childYangName == "state" {
        return &inline.State
    }
    return nil
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &inline.Config
    children["state"] = &inline.State
    return children
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetBundleName() string { return "openconfig" }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetYangName() string { return "inline" }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) SetParent(parent types.Entity) { inline.parent = parent }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetParent() types.Entity { return inline.parent }

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetParentYangName() string { return "set-ext-community" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config
// Configuration data or inline specification of
// set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the extended community values for the update inline with a list. The
    // type is one of the following types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).,
    // or slice of   :go:struct:`BGPWELLKNOWNSTDCOMMUNITY
    // <ydk/models/bgp_types/BGPWELLKNOWNSTDCOMMUNITY>`.
    Communities []interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetGoName(yname string) string {
    if yname == "communities" { return "Communities" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["communities"] = config.Communities
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetParentYangName() string { return "inline" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State
// Operational state data or inline specification of
// set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the extended community values for the update inline with a list. The
    // type is one of the following types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).,
    // or slice of   :go:struct:`BGPWELLKNOWNSTDCOMMUNITY
    // <ydk/models/bgp_types/BGPWELLKNOWNSTDCOMMUNITY>`.
    Communities []interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetGoName(yname string) string {
    if yname == "communities" { return "Communities" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["communities"] = state.Communities
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetParentYangName() string { return "inline" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference
// Provide a reference to an extended community set for the
// set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for referening an extended community-set in the
    // set-ext-community action.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config

    // Operational state data for referening an extended community-set in the
    // set-ext-community action.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetSegmentPath() string {
    return "reference"
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &reference.Config
    }
    if childYangName == "state" {
        return &reference.State
    }
    return nil
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &reference.Config
    children["state"] = &reference.State
    return children
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetBundleName() string { return "openconfig" }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetYangName() string { return "reference" }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetParent() types.Entity { return reference.parent }

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetParentYangName() string { return "set-ext-community" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config
// Configuration data for referening an extended
// community-set in the set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined extended community set by name. The type is string.
    // Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_ExtCommunitySetName
    ExtCommunitySetRef interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetGoName(yname string) string {
    if yname == "ext-community-set-ref" { return "ExtCommunitySetRef" }
    return ""
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetSegmentPath() string {
    return "config"
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ext-community-set-ref"] = config.ExtCommunitySetRef
    return leafs
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetBundleName() string { return "openconfig" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetYangName() string { return "config" }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetParent() types.Entity { return config.parent }

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetParentYangName() string { return "reference" }

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State
// Operational state data for referening an extended
// community-set in the set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References a defined extended community set by name. The type is string.
    // Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_ExtCommunitySetName
    ExtCommunitySetRef interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetGoName(yname string) string {
    if yname == "ext-community-set-ref" { return "ExtCommunitySetRef" }
    return ""
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetSegmentPath() string {
    return "state"
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ext-community-set-ref"] = state.ExtCommunitySetRef
    return leafs
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetBundleName() string { return "openconfig" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetYangName() string { return "state" }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetParent() types.Entity { return state.parent }

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetParentYangName() string { return "reference" }

