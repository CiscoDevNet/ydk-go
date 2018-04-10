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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Predefined sets of attributes used in policy match statements.
    DefinedSets RoutingPolicy_DefinedSets

    // Enclosing container for the list of top-level policy  definitions.
    PolicyDefinitions RoutingPolicy_PolicyDefinitions
}

func (routingPolicy *RoutingPolicy) GetEntityData() *types.CommonEntityData {
    routingPolicy.EntityData.YFilter = routingPolicy.YFilter
    routingPolicy.EntityData.YangName = "routing-policy"
    routingPolicy.EntityData.BundleName = "openconfig"
    routingPolicy.EntityData.ParentYangName = "openconfig-routing-policy"
    routingPolicy.EntityData.SegmentPath = "openconfig-routing-policy:routing-policy"
    routingPolicy.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routingPolicy.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routingPolicy.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routingPolicy.EntityData.Children = make(map[string]types.YChild)
    routingPolicy.EntityData.Children["defined-sets"] = types.YChild{"DefinedSets", &routingPolicy.DefinedSets}
    routingPolicy.EntityData.Children["policy-definitions"] = types.YChild{"PolicyDefinitions", &routingPolicy.PolicyDefinitions}
    routingPolicy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routingPolicy.EntityData)
}

// RoutingPolicy_DefinedSets
// Predefined sets of attributes used in policy match
// statements
type RoutingPolicy_DefinedSets struct {
    EntityData types.CommonEntityData
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

func (definedSets *RoutingPolicy_DefinedSets) GetEntityData() *types.CommonEntityData {
    definedSets.EntityData.YFilter = definedSets.YFilter
    definedSets.EntityData.YangName = "defined-sets"
    definedSets.EntityData.BundleName = "openconfig"
    definedSets.EntityData.ParentYangName = "routing-policy"
    definedSets.EntityData.SegmentPath = "defined-sets"
    definedSets.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    definedSets.EntityData.NamespaceTable = openconfig.GetNamespaces()
    definedSets.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    definedSets.EntityData.Children = make(map[string]types.YChild)
    definedSets.EntityData.Children["prefix-sets"] = types.YChild{"PrefixSets", &definedSets.PrefixSets}
    definedSets.EntityData.Children["neighbor-sets"] = types.YChild{"NeighborSets", &definedSets.NeighborSets}
    definedSets.EntityData.Children["tag-sets"] = types.YChild{"TagSets", &definedSets.TagSets}
    definedSets.EntityData.Children["openconfig-bgp-policy:bgp-defined-sets"] = types.YChild{"BgpDefinedSets", &definedSets.BgpDefinedSets}
    definedSets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(definedSets.EntityData)
}

// RoutingPolicy_DefinedSets_PrefixSets
// Enclosing container 
type RoutingPolicy_DefinedSets_PrefixSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of the defined prefix sets. The type is slice of
    // RoutingPolicy_DefinedSets_PrefixSets_PrefixSet.
    PrefixSet []RoutingPolicy_DefinedSets_PrefixSets_PrefixSet
}

func (prefixSets *RoutingPolicy_DefinedSets_PrefixSets) GetEntityData() *types.CommonEntityData {
    prefixSets.EntityData.YFilter = prefixSets.YFilter
    prefixSets.EntityData.YangName = "prefix-sets"
    prefixSets.EntityData.BundleName = "openconfig"
    prefixSets.EntityData.ParentYangName = "defined-sets"
    prefixSets.EntityData.SegmentPath = "prefix-sets"
    prefixSets.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixSets.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixSets.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixSets.EntityData.Children = make(map[string]types.YChild)
    prefixSets.EntityData.Children["prefix-set"] = types.YChild{"PrefixSet", nil}
    for i := range prefixSets.PrefixSet {
        prefixSets.EntityData.Children[types.GetSegmentPath(&prefixSets.PrefixSet[i])] = types.YChild{"PrefixSet", &prefixSets.PrefixSet[i]}
    }
    prefixSets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixSets.EntityData)
}

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet
// List of the defined prefix sets
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet struct {
    EntityData types.CommonEntityData
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

func (prefixSet *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet) GetEntityData() *types.CommonEntityData {
    prefixSet.EntityData.YFilter = prefixSet.YFilter
    prefixSet.EntityData.YangName = "prefix-set"
    prefixSet.EntityData.BundleName = "openconfig"
    prefixSet.EntityData.ParentYangName = "prefix-sets"
    prefixSet.EntityData.SegmentPath = "prefix-set" + "[prefix-set-name='" + fmt.Sprintf("%v", prefixSet.PrefixSetName) + "']"
    prefixSet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixSet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixSet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixSet.EntityData.Children = make(map[string]types.YChild)
    prefixSet.EntityData.Children["config"] = types.YChild{"Config", &prefixSet.Config}
    prefixSet.EntityData.Children["state"] = types.YChild{"State", &prefixSet.State}
    prefixSet.EntityData.Children["prefixes"] = types.YChild{"Prefixes", &prefixSet.Prefixes}
    prefixSet.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixSet.EntityData.Leafs["prefix-set-name"] = types.YLeaf{"PrefixSetName", prefixSet.PrefixSetName}
    return &(prefixSet.EntityData)
}

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config
// Configuration data for prefix sets
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name / label of the prefix set -- this is used to reference the set in
    // match conditions. The type is string.
    PrefixSetName interface{}
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["prefix-set-name"] = types.YLeaf{"PrefixSetName", config.PrefixSetName}
    return &(config.EntityData)
}

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State
// Operational state data 
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name / label of the prefix set -- this is used to reference the set in
    // match conditions. The type is string.
    PrefixSetName interface{}
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["prefix-set-name"] = types.YLeaf{"PrefixSetName", state.PrefixSetName}
    return &(state.EntityData)
}

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes
// Enclosing container for the list of prefixes in a policy
// prefix list
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of prefixes in the prefix set. The type is slice of
    // RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix.
    Prefix []RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix
}

func (prefixes *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes) GetEntityData() *types.CommonEntityData {
    prefixes.EntityData.YFilter = prefixes.YFilter
    prefixes.EntityData.YangName = "prefixes"
    prefixes.EntityData.BundleName = "openconfig"
    prefixes.EntityData.ParentYangName = "prefix-set"
    prefixes.EntityData.SegmentPath = "prefixes"
    prefixes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixes.EntityData.Children = make(map[string]types.YChild)
    prefixes.EntityData.Children["prefix"] = types.YChild{"Prefix", nil}
    for i := range prefixes.Prefix {
        prefixes.EntityData.Children[types.GetSegmentPath(&prefixes.Prefix[i])] = types.YChild{"Prefix", &prefixes.Prefix[i]}
    }
    prefixes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixes.EntityData)
}

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix
// List of prefixes in the prefix set
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the ip-prefix list key. The type is
    // one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    IpPrefix interface{}

    // This attribute is a key. Reference to the masklength-range list key. The
    // type is string with pattern: b'^([0-9]+\\.\\.[0-9]+)|exact$'. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config_MasklengthRange
    MasklengthRange interface{}

    // Configuration data for prefix definition.
    Config RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config

    // Operational state data for prefix definition.
    State RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State
}

func (prefix *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "openconfig"
    prefix.EntityData.ParentYangName = "prefixes"
    prefix.EntityData.SegmentPath = "prefix" + "[ip-prefix='" + fmt.Sprintf("%v", prefix.IpPrefix) + "']" + "[masklength-range='" + fmt.Sprintf("%v", prefix.MasklengthRange) + "']"
    prefix.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefix.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefix.EntityData.Children = make(map[string]types.YChild)
    prefix.EntityData.Children["config"] = types.YChild{"Config", &prefix.Config}
    prefix.EntityData.Children["state"] = types.YChild{"State", &prefix.State}
    prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    prefix.EntityData.Leafs["ip-prefix"] = types.YLeaf{"IpPrefix", prefix.IpPrefix}
    prefix.EntityData.Leafs["masklength-range"] = types.YLeaf{"MasklengthRange", prefix.MasklengthRange}
    return &(prefix.EntityData)
}

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config
// Configuration data for prefix definition
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The prefix member in CIDR notation -- while the prefix may be either IPv4
    // or IPv6, most implementations require all members of the prefix set to be
    // the same address family.  Mixing address types in the same prefix set is
    // likely to cause an error. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))'
    // This attribute is mandatory., or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'
    // This attribute is mandatory..
    IpPrefix interface{}

    // Defines a range for the masklength, or 'exact' if the prefix has an exact
    // length.  Example: 10.3.192.0/21 through 10.3.192.0/24 would be expressed as
    // prefix: 10.3.192.0/21, masklength-range: 21..24.  Example: 10.3.192.0/21
    // would be expressed as prefix: 10.3.192.0/21, masklength-range: exact. The
    // type is string with pattern: b'^([0-9]+\\.\\.[0-9]+)|exact$'.
    MasklengthRange interface{}
}

func (config *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["ip-prefix"] = types.YLeaf{"IpPrefix", config.IpPrefix}
    config.EntityData.Leafs["masklength-range"] = types.YLeaf{"MasklengthRange", config.MasklengthRange}
    return &(config.EntityData)
}

// RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State
// Operational state data for prefix definition
type RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The prefix member in CIDR notation -- while the prefix may be either IPv4
    // or IPv6, most implementations require all members of the prefix set to be
    // the same address family.  Mixing address types in the same prefix set is
    // likely to cause an error. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))'
    // This attribute is mandatory., or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'
    // This attribute is mandatory..
    IpPrefix interface{}

    // Defines a range for the masklength, or 'exact' if the prefix has an exact
    // length.  Example: 10.3.192.0/21 through 10.3.192.0/24 would be expressed as
    // prefix: 10.3.192.0/21, masklength-range: 21..24.  Example: 10.3.192.0/21
    // would be expressed as prefix: 10.3.192.0/21, masklength-range: exact. The
    // type is string with pattern: b'^([0-9]+\\.\\.[0-9]+)|exact$'.
    MasklengthRange interface{}
}

func (state *RoutingPolicy_DefinedSets_PrefixSets_PrefixSet_Prefixes_Prefix_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["ip-prefix"] = types.YLeaf{"IpPrefix", state.IpPrefix}
    state.EntityData.Leafs["masklength-range"] = types.YLeaf{"MasklengthRange", state.MasklengthRange}
    return &(state.EntityData)
}

// RoutingPolicy_DefinedSets_NeighborSets
// Enclosing container for the list of neighbor set
// definitions
type RoutingPolicy_DefinedSets_NeighborSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of defined neighbor sets for use in policies. The type is slice of
    // RoutingPolicy_DefinedSets_NeighborSets_NeighborSet.
    NeighborSet []RoutingPolicy_DefinedSets_NeighborSets_NeighborSet
}

func (neighborSets *RoutingPolicy_DefinedSets_NeighborSets) GetEntityData() *types.CommonEntityData {
    neighborSets.EntityData.YFilter = neighborSets.YFilter
    neighborSets.EntityData.YangName = "neighbor-sets"
    neighborSets.EntityData.BundleName = "openconfig"
    neighborSets.EntityData.ParentYangName = "defined-sets"
    neighborSets.EntityData.SegmentPath = "neighbor-sets"
    neighborSets.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighborSets.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighborSets.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighborSets.EntityData.Children = make(map[string]types.YChild)
    neighborSets.EntityData.Children["neighbor-set"] = types.YChild{"NeighborSet", nil}
    for i := range neighborSets.NeighborSet {
        neighborSets.EntityData.Children[types.GetSegmentPath(&neighborSets.NeighborSet[i])] = types.YChild{"NeighborSet", &neighborSets.NeighborSet[i]}
    }
    neighborSets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(neighborSets.EntityData)
}

// RoutingPolicy_DefinedSets_NeighborSets_NeighborSet
// List of defined neighbor sets for use in policies.
type RoutingPolicy_DefinedSets_NeighborSets_NeighborSet struct {
    EntityData types.CommonEntityData
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

func (neighborSet *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet) GetEntityData() *types.CommonEntityData {
    neighborSet.EntityData.YFilter = neighborSet.YFilter
    neighborSet.EntityData.YangName = "neighbor-set"
    neighborSet.EntityData.BundleName = "openconfig"
    neighborSet.EntityData.ParentYangName = "neighbor-sets"
    neighborSet.EntityData.SegmentPath = "neighbor-set" + "[neighbor-set-name='" + fmt.Sprintf("%v", neighborSet.NeighborSetName) + "']"
    neighborSet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighborSet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighborSet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighborSet.EntityData.Children = make(map[string]types.YChild)
    neighborSet.EntityData.Children["config"] = types.YChild{"Config", &neighborSet.Config}
    neighborSet.EntityData.Children["state"] = types.YChild{"State", &neighborSet.State}
    neighborSet.EntityData.Leafs = make(map[string]types.YLeaf)
    neighborSet.EntityData.Leafs["neighbor-set-name"] = types.YLeaf{"NeighborSetName", neighborSet.NeighborSetName}
    return &(neighborSet.EntityData)
}

// RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config
// Configuration data for neighbor sets.
type RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name / label of the neighbor set -- this is used to reference the set in
    // match conditions. The type is string.
    NeighborSetName interface{}

    // List of IP addresses in the neighbor set. The type is one of the following
    // types: slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address []interface{}
}

func (config *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "neighbor-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["neighbor-set-name"] = types.YLeaf{"NeighborSetName", config.NeighborSetName}
    config.EntityData.Leafs["address"] = types.YLeaf{"Address", config.Address}
    return &(config.EntityData)
}

// RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State
// Operational state data for neighbor sets.
type RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name / label of the neighbor set -- this is used to reference the set in
    // match conditions. The type is string.
    NeighborSetName interface{}

    // List of IP addresses in the neighbor set. The type is one of the following
    // types: slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address []interface{}
}

func (state *RoutingPolicy_DefinedSets_NeighborSets_NeighborSet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "neighbor-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["neighbor-set-name"] = types.YLeaf{"NeighborSetName", state.NeighborSetName}
    state.EntityData.Leafs["address"] = types.YLeaf{"Address", state.Address}
    return &(state.EntityData)
}

// RoutingPolicy_DefinedSets_TagSets
// Enclosing container for the list of tag sets.
type RoutingPolicy_DefinedSets_TagSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of tag set definitions. The type is slice of
    // RoutingPolicy_DefinedSets_TagSets_TagSet.
    TagSet []RoutingPolicy_DefinedSets_TagSets_TagSet
}

func (tagSets *RoutingPolicy_DefinedSets_TagSets) GetEntityData() *types.CommonEntityData {
    tagSets.EntityData.YFilter = tagSets.YFilter
    tagSets.EntityData.YangName = "tag-sets"
    tagSets.EntityData.BundleName = "openconfig"
    tagSets.EntityData.ParentYangName = "defined-sets"
    tagSets.EntityData.SegmentPath = "tag-sets"
    tagSets.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    tagSets.EntityData.NamespaceTable = openconfig.GetNamespaces()
    tagSets.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    tagSets.EntityData.Children = make(map[string]types.YChild)
    tagSets.EntityData.Children["tag-set"] = types.YChild{"TagSet", nil}
    for i := range tagSets.TagSet {
        tagSets.EntityData.Children[types.GetSegmentPath(&tagSets.TagSet[i])] = types.YChild{"TagSet", &tagSets.TagSet[i]}
    }
    tagSets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tagSets.EntityData)
}

// RoutingPolicy_DefinedSets_TagSets_TagSet
// List of tag set definitions.
type RoutingPolicy_DefinedSets_TagSets_TagSet struct {
    EntityData types.CommonEntityData
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

func (tagSet *RoutingPolicy_DefinedSets_TagSets_TagSet) GetEntityData() *types.CommonEntityData {
    tagSet.EntityData.YFilter = tagSet.YFilter
    tagSet.EntityData.YangName = "tag-set"
    tagSet.EntityData.BundleName = "openconfig"
    tagSet.EntityData.ParentYangName = "tag-sets"
    tagSet.EntityData.SegmentPath = "tag-set" + "[tag-set-name='" + fmt.Sprintf("%v", tagSet.TagSetName) + "']"
    tagSet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    tagSet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    tagSet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    tagSet.EntityData.Children = make(map[string]types.YChild)
    tagSet.EntityData.Children["config"] = types.YChild{"Config", &tagSet.Config}
    tagSet.EntityData.Children["state"] = types.YChild{"State", &tagSet.State}
    tagSet.EntityData.Leafs = make(map[string]types.YLeaf)
    tagSet.EntityData.Leafs["tag-set-name"] = types.YLeaf{"TagSetName", tagSet.TagSetName}
    return &(tagSet.EntityData)
}

// RoutingPolicy_DefinedSets_TagSets_TagSet_Config
// Configuration data for tag sets
type RoutingPolicy_DefinedSets_TagSets_TagSet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name / label of the tag set -- this is used to reference the set in match
    // conditions. The type is string.
    TagSetName interface{}

    // Value of the tag set member. The type is one of the following types: slice
    // of int with range: 0..4294967295, or slice of string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    TagValue []interface{}
}

func (config *RoutingPolicy_DefinedSets_TagSets_TagSet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "tag-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["tag-set-name"] = types.YLeaf{"TagSetName", config.TagSetName}
    config.EntityData.Leafs["tag-value"] = types.YLeaf{"TagValue", config.TagValue}
    return &(config.EntityData)
}

// RoutingPolicy_DefinedSets_TagSets_TagSet_State
// Operational state data for tag sets
type RoutingPolicy_DefinedSets_TagSets_TagSet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name / label of the tag set -- this is used to reference the set in match
    // conditions. The type is string.
    TagSetName interface{}

    // Value of the tag set member. The type is one of the following types: slice
    // of int with range: 0..4294967295, or slice of string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    TagValue []interface{}
}

func (state *RoutingPolicy_DefinedSets_TagSets_TagSet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "tag-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["tag-set-name"] = types.YLeaf{"TagSetName", state.TagSetName}
    state.EntityData.Leafs["tag-value"] = types.YLeaf{"TagValue", state.TagValue}
    return &(state.EntityData)
}

// RoutingPolicy_DefinedSets_BgpDefinedSets
// BGP-related set definitions for policy match conditions
type RoutingPolicy_DefinedSets_BgpDefinedSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enclosing container for list of defined BGP community sets.
    CommunitySets RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets

    // Enclosing container for list of extended BGP community sets.
    ExtCommunitySets RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets

    // Enclosing container for list of define AS path sets.
    AsPathSets RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets
}

func (bgpDefinedSets *RoutingPolicy_DefinedSets_BgpDefinedSets) GetEntityData() *types.CommonEntityData {
    bgpDefinedSets.EntityData.YFilter = bgpDefinedSets.YFilter
    bgpDefinedSets.EntityData.YangName = "bgp-defined-sets"
    bgpDefinedSets.EntityData.BundleName = "openconfig"
    bgpDefinedSets.EntityData.ParentYangName = "defined-sets"
    bgpDefinedSets.EntityData.SegmentPath = "openconfig-bgp-policy:bgp-defined-sets"
    bgpDefinedSets.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    bgpDefinedSets.EntityData.NamespaceTable = openconfig.GetNamespaces()
    bgpDefinedSets.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    bgpDefinedSets.EntityData.Children = make(map[string]types.YChild)
    bgpDefinedSets.EntityData.Children["community-sets"] = types.YChild{"CommunitySets", &bgpDefinedSets.CommunitySets}
    bgpDefinedSets.EntityData.Children["ext-community-sets"] = types.YChild{"ExtCommunitySets", &bgpDefinedSets.ExtCommunitySets}
    bgpDefinedSets.EntityData.Children["as-path-sets"] = types.YChild{"AsPathSets", &bgpDefinedSets.AsPathSets}
    bgpDefinedSets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpDefinedSets.EntityData)
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets
// Enclosing container for list of defined BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of defined BGP community sets. The type is slice of
    // RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet.
    CommunitySet []RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet
}

func (communitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets) GetEntityData() *types.CommonEntityData {
    communitySets.EntityData.YFilter = communitySets.YFilter
    communitySets.EntityData.YangName = "community-sets"
    communitySets.EntityData.BundleName = "openconfig"
    communitySets.EntityData.ParentYangName = "bgp-defined-sets"
    communitySets.EntityData.SegmentPath = "community-sets"
    communitySets.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    communitySets.EntityData.NamespaceTable = openconfig.GetNamespaces()
    communitySets.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    communitySets.EntityData.Children = make(map[string]types.YChild)
    communitySets.EntityData.Children["community-set"] = types.YChild{"CommunitySet", nil}
    for i := range communitySets.CommunitySet {
        communitySets.EntityData.Children[types.GetSegmentPath(&communitySets.CommunitySet[i])] = types.YChild{"CommunitySet", &communitySets.CommunitySet[i]}
    }
    communitySets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(communitySets.EntityData)
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet
// List of defined BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet struct {
    EntityData types.CommonEntityData
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

func (communitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet) GetEntityData() *types.CommonEntityData {
    communitySet.EntityData.YFilter = communitySet.YFilter
    communitySet.EntityData.YangName = "community-set"
    communitySet.EntityData.BundleName = "openconfig"
    communitySet.EntityData.ParentYangName = "community-sets"
    communitySet.EntityData.SegmentPath = "community-set" + "[community-set-name='" + fmt.Sprintf("%v", communitySet.CommunitySetName) + "']"
    communitySet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    communitySet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    communitySet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    communitySet.EntityData.Children = make(map[string]types.YChild)
    communitySet.EntityData.Children["config"] = types.YChild{"Config", &communitySet.Config}
    communitySet.EntityData.Children["state"] = types.YChild{"State", &communitySet.State}
    communitySet.EntityData.Leafs = make(map[string]types.YLeaf)
    communitySet.EntityData.Leafs["community-set-name"] = types.YLeaf{"CommunitySetName", communitySet.CommunitySetName}
    return &(communitySet.EntityData)
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config
// Configuration data for BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name / label of the community set -- this is used to reference the set in
    // match conditions. The type is string. This attribute is mandatory.
    CommunitySetName interface{}

    // members of the community set. The type is one of the following types: slice
    // of int with range: 65536..4294901759, or slice of string with pattern:
    // b'([0-9]+:[0-9]+)'., or slice of string, or slice of  
    // :go:struct:`BGPWELLKNOWNSTDCOMMUNITY
    // <ydk/models/bgp_types/BGPWELLKNOWNSTDCOMMUNITY>`.
    CommunityMember []interface{}
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "community-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["community-set-name"] = types.YLeaf{"CommunitySetName", config.CommunitySetName}
    config.EntityData.Leafs["community-member"] = types.YLeaf{"CommunityMember", config.CommunityMember}
    return &(config.EntityData)
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State
// Operational state data for BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name / label of the community set -- this is used to reference the set in
    // match conditions. The type is string. This attribute is mandatory.
    CommunitySetName interface{}

    // members of the community set. The type is one of the following types: slice
    // of int with range: 65536..4294901759, or slice of string with pattern:
    // b'([0-9]+:[0-9]+)'., or slice of string, or slice of  
    // :go:struct:`BGPWELLKNOWNSTDCOMMUNITY
    // <ydk/models/bgp_types/BGPWELLKNOWNSTDCOMMUNITY>`.
    CommunityMember []interface{}
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "community-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["community-set-name"] = types.YLeaf{"CommunitySetName", state.CommunitySetName}
    state.EntityData.Leafs["community-member"] = types.YLeaf{"CommunityMember", state.CommunityMember}
    return &(state.EntityData)
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets
// Enclosing container for list of extended BGP community
// sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of defined extended BGP community sets. The type is slice of
    // RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet.
    ExtCommunitySet []RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet
}

func (extCommunitySets *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets) GetEntityData() *types.CommonEntityData {
    extCommunitySets.EntityData.YFilter = extCommunitySets.YFilter
    extCommunitySets.EntityData.YangName = "ext-community-sets"
    extCommunitySets.EntityData.BundleName = "openconfig"
    extCommunitySets.EntityData.ParentYangName = "bgp-defined-sets"
    extCommunitySets.EntityData.SegmentPath = "ext-community-sets"
    extCommunitySets.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    extCommunitySets.EntityData.NamespaceTable = openconfig.GetNamespaces()
    extCommunitySets.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    extCommunitySets.EntityData.Children = make(map[string]types.YChild)
    extCommunitySets.EntityData.Children["ext-community-set"] = types.YChild{"ExtCommunitySet", nil}
    for i := range extCommunitySets.ExtCommunitySet {
        extCommunitySets.EntityData.Children[types.GetSegmentPath(&extCommunitySets.ExtCommunitySet[i])] = types.YChild{"ExtCommunitySet", &extCommunitySets.ExtCommunitySet[i]}
    }
    extCommunitySets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(extCommunitySets.EntityData)
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet
// List of defined extended BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet struct {
    EntityData types.CommonEntityData
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

func (extCommunitySet *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet) GetEntityData() *types.CommonEntityData {
    extCommunitySet.EntityData.YFilter = extCommunitySet.YFilter
    extCommunitySet.EntityData.YangName = "ext-community-set"
    extCommunitySet.EntityData.BundleName = "openconfig"
    extCommunitySet.EntityData.ParentYangName = "ext-community-sets"
    extCommunitySet.EntityData.SegmentPath = "ext-community-set" + "[ext-community-set-name='" + fmt.Sprintf("%v", extCommunitySet.ExtCommunitySetName) + "']"
    extCommunitySet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    extCommunitySet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    extCommunitySet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    extCommunitySet.EntityData.Children = make(map[string]types.YChild)
    extCommunitySet.EntityData.Children["config"] = types.YChild{"Config", &extCommunitySet.Config}
    extCommunitySet.EntityData.Children["state"] = types.YChild{"State", &extCommunitySet.State}
    extCommunitySet.EntityData.Leafs = make(map[string]types.YLeaf)
    extCommunitySet.EntityData.Leafs["ext-community-set-name"] = types.YLeaf{"ExtCommunitySetName", extCommunitySet.ExtCommunitySetName}
    return &(extCommunitySet.EntityData)
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config
// Configuration data for extended BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name / label of the extended community set -- this is used to reference the
    // set in match conditions. The type is string.
    ExtCommunitySetName interface{}

    // members of the extended community set. The type is one of the following
    // types: slice of string with pattern:
    // b'(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9])',
    // or slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9])',
    // or slice of string with pattern:
    // b'route\\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9])',
    // or slice of string with pattern:
    // b'route\\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9])',
    // or slice of string with pattern:
    // b'route\\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9])',
    // or slice of string with pattern:
    // b'route\\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9])'.,
    // or slice of string.
    ExtCommunityMember []interface{}
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ext-community-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["ext-community-set-name"] = types.YLeaf{"ExtCommunitySetName", config.ExtCommunitySetName}
    config.EntityData.Leafs["ext-community-member"] = types.YLeaf{"ExtCommunityMember", config.ExtCommunityMember}
    return &(config.EntityData)
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State
// Operational state data for extended BGP community sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name / label of the extended community set -- this is used to reference the
    // set in match conditions. The type is string.
    ExtCommunitySetName interface{}

    // members of the extended community set. The type is one of the following
    // types: slice of string with pattern:
    // b'(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9])',
    // or slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9])',
    // or slice of string with pattern:
    // b'route\\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9])',
    // or slice of string with pattern:
    // b'route\\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9])',
    // or slice of string with pattern:
    // b'route\\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9])',
    // or slice of string with pattern:
    // b'route\\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9])'.,
    // or slice of string.
    ExtCommunityMember []interface{}
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ext-community-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["ext-community-set-name"] = types.YLeaf{"ExtCommunitySetName", state.ExtCommunitySetName}
    state.EntityData.Leafs["ext-community-member"] = types.YLeaf{"ExtCommunityMember", state.ExtCommunityMember}
    return &(state.EntityData)
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets
// Enclosing container for list of define AS path sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of defined AS path sets. The type is slice of
    // RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet.
    AsPathSet []RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet
}

func (asPathSets *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets) GetEntityData() *types.CommonEntityData {
    asPathSets.EntityData.YFilter = asPathSets.YFilter
    asPathSets.EntityData.YangName = "as-path-sets"
    asPathSets.EntityData.BundleName = "openconfig"
    asPathSets.EntityData.ParentYangName = "bgp-defined-sets"
    asPathSets.EntityData.SegmentPath = "as-path-sets"
    asPathSets.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    asPathSets.EntityData.NamespaceTable = openconfig.GetNamespaces()
    asPathSets.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    asPathSets.EntityData.Children = make(map[string]types.YChild)
    asPathSets.EntityData.Children["as-path-set"] = types.YChild{"AsPathSet", nil}
    for i := range asPathSets.AsPathSet {
        asPathSets.EntityData.Children[types.GetSegmentPath(&asPathSets.AsPathSet[i])] = types.YChild{"AsPathSet", &asPathSets.AsPathSet[i]}
    }
    asPathSets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asPathSets.EntityData)
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet
// List of defined AS path sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet struct {
    EntityData types.CommonEntityData
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

func (asPathSet *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet) GetEntityData() *types.CommonEntityData {
    asPathSet.EntityData.YFilter = asPathSet.YFilter
    asPathSet.EntityData.YangName = "as-path-set"
    asPathSet.EntityData.BundleName = "openconfig"
    asPathSet.EntityData.ParentYangName = "as-path-sets"
    asPathSet.EntityData.SegmentPath = "as-path-set" + "[as-path-set-name='" + fmt.Sprintf("%v", asPathSet.AsPathSetName) + "']"
    asPathSet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    asPathSet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    asPathSet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    asPathSet.EntityData.Children = make(map[string]types.YChild)
    asPathSet.EntityData.Children["config"] = types.YChild{"Config", &asPathSet.Config}
    asPathSet.EntityData.Children["state"] = types.YChild{"State", &asPathSet.State}
    asPathSet.EntityData.Leafs = make(map[string]types.YLeaf)
    asPathSet.EntityData.Leafs["as-path-set-name"] = types.YLeaf{"AsPathSetName", asPathSet.AsPathSetName}
    return &(asPathSet.EntityData)
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config
// Configuration data for AS path sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name of the AS path set -- this is used to reference the set in match
    // conditions. The type is string.
    AsPathSetName interface{}

    // AS path expression -- list of ASes in the set. The type is slice of string.
    AsPathSetMember []interface{}
}

func (config *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "as-path-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["as-path-set-name"] = types.YLeaf{"AsPathSetName", config.AsPathSetName}
    config.EntityData.Leafs["as-path-set-member"] = types.YLeaf{"AsPathSetMember", config.AsPathSetMember}
    return &(config.EntityData)
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State
// Operational state data for AS path sets
type RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name of the AS path set -- this is used to reference the set in match
    // conditions. The type is string.
    AsPathSetName interface{}

    // AS path expression -- list of ASes in the set. The type is slice of string.
    AsPathSetMember []interface{}
}

func (state *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "as-path-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["as-path-set-name"] = types.YLeaf{"AsPathSetName", state.AsPathSetName}
    state.EntityData.Leafs["as-path-set-member"] = types.YLeaf{"AsPathSetMember", state.AsPathSetMember}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions
// Enclosing container for the list of top-level policy
//  definitions
type RoutingPolicy_PolicyDefinitions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of top-level policy definitions, keyed by unique name.  These policy
    // definitions are expected to be referenced (by name) in policy chains
    // specified in import or export configuration statements. The type is slice
    // of RoutingPolicy_PolicyDefinitions_PolicyDefinition.
    PolicyDefinition []RoutingPolicy_PolicyDefinitions_PolicyDefinition
}

func (policyDefinitions *RoutingPolicy_PolicyDefinitions) GetEntityData() *types.CommonEntityData {
    policyDefinitions.EntityData.YFilter = policyDefinitions.YFilter
    policyDefinitions.EntityData.YangName = "policy-definitions"
    policyDefinitions.EntityData.BundleName = "openconfig"
    policyDefinitions.EntityData.ParentYangName = "routing-policy"
    policyDefinitions.EntityData.SegmentPath = "policy-definitions"
    policyDefinitions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    policyDefinitions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    policyDefinitions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    policyDefinitions.EntityData.Children = make(map[string]types.YChild)
    policyDefinitions.EntityData.Children["policy-definition"] = types.YChild{"PolicyDefinition", nil}
    for i := range policyDefinitions.PolicyDefinition {
        policyDefinitions.EntityData.Children[types.GetSegmentPath(&policyDefinitions.PolicyDefinition[i])] = types.YChild{"PolicyDefinition", &policyDefinitions.PolicyDefinition[i]}
    }
    policyDefinitions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policyDefinitions.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition
// List of top-level policy definitions, keyed by unique
// name.  These policy definitions are expected to be
// referenced (by name) in policy chains specified in import
// or export configuration statements.
type RoutingPolicy_PolicyDefinitions_PolicyDefinition struct {
    EntityData types.CommonEntityData
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

func (policyDefinition *RoutingPolicy_PolicyDefinitions_PolicyDefinition) GetEntityData() *types.CommonEntityData {
    policyDefinition.EntityData.YFilter = policyDefinition.YFilter
    policyDefinition.EntityData.YangName = "policy-definition"
    policyDefinition.EntityData.BundleName = "openconfig"
    policyDefinition.EntityData.ParentYangName = "policy-definitions"
    policyDefinition.EntityData.SegmentPath = "policy-definition" + "[name='" + fmt.Sprintf("%v", policyDefinition.Name) + "']"
    policyDefinition.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    policyDefinition.EntityData.NamespaceTable = openconfig.GetNamespaces()
    policyDefinition.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    policyDefinition.EntityData.Children = make(map[string]types.YChild)
    policyDefinition.EntityData.Children["config"] = types.YChild{"Config", &policyDefinition.Config}
    policyDefinition.EntityData.Children["state"] = types.YChild{"State", &policyDefinition.State}
    policyDefinition.EntityData.Children["statements"] = types.YChild{"Statements", &policyDefinition.Statements}
    policyDefinition.EntityData.Leafs = make(map[string]types.YLeaf)
    policyDefinition.EntityData.Leafs["name"] = types.YLeaf{"Name", policyDefinition.Name}
    return &(policyDefinition.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config
// Configuration data for policy defintions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the top-level policy definition -- this name is used in references
    // to the current policy. The type is string.
    Name interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "policy-definition"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["name"] = types.YLeaf{"Name", config.Name}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_State
// Operational state data for policy definitions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the top-level policy definition -- this name is used in references
    // to the current policy. The type is string.
    Name interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "policy-definition"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["name"] = types.YLeaf{"Name", state.Name}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements
// Enclosing container for policy statements
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy statements group conditions and actions within a policy definition. 
    // They are evaluated in the order specified (see the description of policy
    // evaluation at the top of this module. The type is slice of
    // RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement.
    Statement []RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement
}

func (statements *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements) GetEntityData() *types.CommonEntityData {
    statements.EntityData.YFilter = statements.YFilter
    statements.EntityData.YangName = "statements"
    statements.EntityData.BundleName = "openconfig"
    statements.EntityData.ParentYangName = "policy-definition"
    statements.EntityData.SegmentPath = "statements"
    statements.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    statements.EntityData.NamespaceTable = openconfig.GetNamespaces()
    statements.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    statements.EntityData.Children = make(map[string]types.YChild)
    statements.EntityData.Children["statement"] = types.YChild{"Statement", nil}
    for i := range statements.Statement {
        statements.EntityData.Children[types.GetSegmentPath(&statements.Statement[i])] = types.YChild{"Statement", &statements.Statement[i]}
    }
    statements.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statements.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement
// Policy statements group conditions and actions
// within a policy definition.  They are evaluated in
// the order specified (see the description of policy
// evaluation at the top of this module.
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement struct {
    EntityData types.CommonEntityData
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

func (statement *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement) GetEntityData() *types.CommonEntityData {
    statement.EntityData.YFilter = statement.YFilter
    statement.EntityData.YangName = "statement"
    statement.EntityData.BundleName = "openconfig"
    statement.EntityData.ParentYangName = "statements"
    statement.EntityData.SegmentPath = "statement" + "[name='" + fmt.Sprintf("%v", statement.Name) + "']"
    statement.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    statement.EntityData.NamespaceTable = openconfig.GetNamespaces()
    statement.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    statement.EntityData.Children = make(map[string]types.YChild)
    statement.EntityData.Children["config"] = types.YChild{"Config", &statement.Config}
    statement.EntityData.Children["state"] = types.YChild{"State", &statement.State}
    statement.EntityData.Children["conditions"] = types.YChild{"Conditions", &statement.Conditions}
    statement.EntityData.Children["actions"] = types.YChild{"Actions", &statement.Actions}
    statement.EntityData.Leafs = make(map[string]types.YLeaf)
    statement.EntityData.Leafs["name"] = types.YLeaf{"Name", statement.Name}
    return &(statement.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config
// Configuration data for policy statements
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name of the policy statement. The type is string.
    Name interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "statement"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["name"] = types.YLeaf{"Name", config.Name}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State
// Operational state data for policy statements
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name of the policy statement. The type is string.
    Name interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "statement"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["name"] = types.YLeaf{"Name", state.Name}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions
// Condition statements for the current policy statement
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions struct {
    EntityData types.CommonEntityData
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

    // Match conditions relating to the IS-IS protocol.
    IsisConditions RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IsisConditions
}

func (conditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions) GetEntityData() *types.CommonEntityData {
    conditions.EntityData.YFilter = conditions.YFilter
    conditions.EntityData.YangName = "conditions"
    conditions.EntityData.BundleName = "openconfig"
    conditions.EntityData.ParentYangName = "statement"
    conditions.EntityData.SegmentPath = "conditions"
    conditions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    conditions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    conditions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    conditions.EntityData.Children = make(map[string]types.YChild)
    conditions.EntityData.Children["config"] = types.YChild{"Config", &conditions.Config}
    conditions.EntityData.Children["state"] = types.YChild{"State", &conditions.State}
    conditions.EntityData.Children["match-interface"] = types.YChild{"MatchInterface", &conditions.MatchInterface}
    conditions.EntityData.Children["match-prefix-set"] = types.YChild{"MatchPrefixSet", &conditions.MatchPrefixSet}
    conditions.EntityData.Children["match-neighbor-set"] = types.YChild{"MatchNeighborSet", &conditions.MatchNeighborSet}
    conditions.EntityData.Children["match-tag-set"] = types.YChild{"MatchTagSet", &conditions.MatchTagSet}
    conditions.EntityData.Children["igp-conditions"] = types.YChild{"IgpConditions", &conditions.IgpConditions}
    conditions.EntityData.Children["openconfig-bgp-policy:bgp-conditions"] = types.YChild{"BgpConditions", &conditions.BgpConditions}
    conditions.EntityData.Children["openconfig-isis-policy:isis-conditions"] = types.YChild{"IsisConditions", &conditions.IsisConditions}
    conditions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(conditions.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config
// Configuration data for policy conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config struct {
    EntityData types.CommonEntityData
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
    // BGPISISOSPFOSPF3STATICDIRECTLYCONNECTEDLOCALAGGREGATE.
    InstallProtocolEq interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "conditions"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["call-policy"] = types.YLeaf{"CallPolicy", config.CallPolicy}
    config.EntityData.Leafs["install-protocol-eq"] = types.YLeaf{"InstallProtocolEq", config.InstallProtocolEq}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State
// Operational state data for policy conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State struct {
    EntityData types.CommonEntityData
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
    // BGPISISOSPFOSPF3STATICDIRECTLYCONNECTEDLOCALAGGREGATE.
    InstallProtocolEq interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "conditions"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["call-policy"] = types.YLeaf{"CallPolicy", state.CallPolicy}
    state.EntityData.Leafs["install-protocol-eq"] = types.YLeaf{"InstallProtocolEq", state.InstallProtocolEq}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface
// Top-level container for interface match conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for interface match conditions.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config

    // Operational state data for interface match conditions.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State
}

func (matchInterface *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface) GetEntityData() *types.CommonEntityData {
    matchInterface.EntityData.YFilter = matchInterface.YFilter
    matchInterface.EntityData.YangName = "match-interface"
    matchInterface.EntityData.BundleName = "openconfig"
    matchInterface.EntityData.ParentYangName = "conditions"
    matchInterface.EntityData.SegmentPath = "match-interface"
    matchInterface.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    matchInterface.EntityData.NamespaceTable = openconfig.GetNamespaces()
    matchInterface.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    matchInterface.EntityData.Children = make(map[string]types.YChild)
    matchInterface.EntityData.Children["config"] = types.YChild{"Config", &matchInterface.Config}
    matchInterface.EntityData.Children["state"] = types.YChild{"State", &matchInterface.State}
    matchInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(matchInterface.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config
// Configuration data for interface match conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface_ interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "match-interface"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", config.Interface_}
    config.EntityData.Leafs["subinterface"] = types.YLeaf{"Subinterface", config.Subinterface}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State
// Operational state data for interface match conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface_ interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchInterface_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "match-interface"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", state.Interface_}
    state.EntityData.Leafs["subinterface"] = types.YLeaf{"Subinterface", state.Subinterface}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet
// Match a referenced prefix-set according to the logic
// defined in the match-set-options leaf
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for a prefix-set condition.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config

    // Operational state data for a prefix-set condition.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State
}

func (matchPrefixSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet) GetEntityData() *types.CommonEntityData {
    matchPrefixSet.EntityData.YFilter = matchPrefixSet.YFilter
    matchPrefixSet.EntityData.YangName = "match-prefix-set"
    matchPrefixSet.EntityData.BundleName = "openconfig"
    matchPrefixSet.EntityData.ParentYangName = "conditions"
    matchPrefixSet.EntityData.SegmentPath = "match-prefix-set"
    matchPrefixSet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    matchPrefixSet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    matchPrefixSet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    matchPrefixSet.EntityData.Children = make(map[string]types.YChild)
    matchPrefixSet.EntityData.Children["config"] = types.YChild{"Config", &matchPrefixSet.Config}
    matchPrefixSet.EntityData.Children["state"] = types.YChild{"State", &matchPrefixSet.State}
    matchPrefixSet.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(matchPrefixSet.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config
// Configuration data for a prefix-set condition
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config struct {
    EntityData types.CommonEntityData
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

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "match-prefix-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["prefix-set"] = types.YLeaf{"PrefixSet", config.PrefixSet}
    config.EntityData.Leafs["match-set-options"] = types.YLeaf{"MatchSetOptions", config.MatchSetOptions}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State
// Operational state data for a prefix-set condition
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State struct {
    EntityData types.CommonEntityData
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

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchPrefixSet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "match-prefix-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["prefix-set"] = types.YLeaf{"PrefixSet", state.PrefixSet}
    state.EntityData.Leafs["match-set-options"] = types.YLeaf{"MatchSetOptions", state.MatchSetOptions}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet
// Match a referenced neighbor set according to the logic
// defined in the match-set-options-leaf
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data .
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config

    // Operational state data .
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State
}

func (matchNeighborSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet) GetEntityData() *types.CommonEntityData {
    matchNeighborSet.EntityData.YFilter = matchNeighborSet.YFilter
    matchNeighborSet.EntityData.YangName = "match-neighbor-set"
    matchNeighborSet.EntityData.BundleName = "openconfig"
    matchNeighborSet.EntityData.ParentYangName = "conditions"
    matchNeighborSet.EntityData.SegmentPath = "match-neighbor-set"
    matchNeighborSet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    matchNeighborSet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    matchNeighborSet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    matchNeighborSet.EntityData.Children = make(map[string]types.YChild)
    matchNeighborSet.EntityData.Children["config"] = types.YChild{"Config", &matchNeighborSet.Config}
    matchNeighborSet.EntityData.Children["state"] = types.YChild{"State", &matchNeighborSet.State}
    matchNeighborSet.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(matchNeighborSet.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config
// Configuration data 
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config struct {
    EntityData types.CommonEntityData
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

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "match-neighbor-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["neighbor-set"] = types.YLeaf{"NeighborSet", config.NeighborSet}
    config.EntityData.Leafs["match-set-options"] = types.YLeaf{"MatchSetOptions", config.MatchSetOptions}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State
// Operational state data 
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State struct {
    EntityData types.CommonEntityData
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

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchNeighborSet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "match-neighbor-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["neighbor-set"] = types.YLeaf{"NeighborSet", state.NeighborSet}
    state.EntityData.Leafs["match-set-options"] = types.YLeaf{"MatchSetOptions", state.MatchSetOptions}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet
// Match a referenced tag set according to the logic defined
// in the match-options-set leaf
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for tag-set conditions.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config

    // Operational state data tag-set conditions.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State
}

func (matchTagSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet) GetEntityData() *types.CommonEntityData {
    matchTagSet.EntityData.YFilter = matchTagSet.YFilter
    matchTagSet.EntityData.YangName = "match-tag-set"
    matchTagSet.EntityData.BundleName = "openconfig"
    matchTagSet.EntityData.ParentYangName = "conditions"
    matchTagSet.EntityData.SegmentPath = "match-tag-set"
    matchTagSet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    matchTagSet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    matchTagSet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    matchTagSet.EntityData.Children = make(map[string]types.YChild)
    matchTagSet.EntityData.Children["config"] = types.YChild{"Config", &matchTagSet.Config}
    matchTagSet.EntityData.Children["state"] = types.YChild{"State", &matchTagSet.State}
    matchTagSet.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(matchTagSet.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config
// Configuration data for tag-set conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config struct {
    EntityData types.CommonEntityData
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

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "match-tag-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["tag-set"] = types.YLeaf{"TagSet", config.TagSet}
    config.EntityData.Leafs["match-set-options"] = types.YLeaf{"MatchSetOptions", config.MatchSetOptions}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State
// Operational state data tag-set conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State struct {
    EntityData types.CommonEntityData
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

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_MatchTagSet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "match-tag-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["tag-set"] = types.YLeaf{"TagSet", state.TagSet}
    state.EntityData.Leafs["match-set-options"] = types.YLeaf{"MatchSetOptions", state.MatchSetOptions}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions
// Policy conditions for IGP attributes
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (igpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IgpConditions) GetEntityData() *types.CommonEntityData {
    igpConditions.EntityData.YFilter = igpConditions.YFilter
    igpConditions.EntityData.YangName = "igp-conditions"
    igpConditions.EntityData.BundleName = "openconfig"
    igpConditions.EntityData.ParentYangName = "conditions"
    igpConditions.EntityData.SegmentPath = "igp-conditions"
    igpConditions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    igpConditions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    igpConditions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    igpConditions.EntityData.Children = make(map[string]types.YChild)
    igpConditions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(igpConditions.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions
// Top-level container 
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions struct {
    EntityData types.CommonEntityData
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

func (bgpConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions) GetEntityData() *types.CommonEntityData {
    bgpConditions.EntityData.YFilter = bgpConditions.YFilter
    bgpConditions.EntityData.YangName = "bgp-conditions"
    bgpConditions.EntityData.BundleName = "openconfig"
    bgpConditions.EntityData.ParentYangName = "conditions"
    bgpConditions.EntityData.SegmentPath = "openconfig-bgp-policy:bgp-conditions"
    bgpConditions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    bgpConditions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    bgpConditions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    bgpConditions.EntityData.Children = make(map[string]types.YChild)
    bgpConditions.EntityData.Children["config"] = types.YChild{"Config", &bgpConditions.Config}
    bgpConditions.EntityData.Children["state"] = types.YChild{"State", &bgpConditions.State}
    bgpConditions.EntityData.Children["community-count"] = types.YChild{"CommunityCount", &bgpConditions.CommunityCount}
    bgpConditions.EntityData.Children["as-path-length"] = types.YChild{"AsPathLength", &bgpConditions.AsPathLength}
    bgpConditions.EntityData.Children["match-community-set"] = types.YChild{"MatchCommunitySet", &bgpConditions.MatchCommunitySet}
    bgpConditions.EntityData.Children["match-ext-community-set"] = types.YChild{"MatchExtCommunitySet", &bgpConditions.MatchExtCommunitySet}
    bgpConditions.EntityData.Children["match-as-path-set"] = types.YChild{"MatchAsPathSet", &bgpConditions.MatchAsPathSet}
    bgpConditions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpConditions.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config
// Configuration data for BGP-specific policy conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Condition to check if the received MED value is equal to the specified
    // value. The type is interface{} with range: 0..4294967295.
    MedEq interface{}

    // Condition to check if the route origin is equal to the specified value. The
    // type is BgpOriginAttrType.
    OriginEq interface{}

    // List of next hop addresses to check for in the route update. The type is
    // one of the following types: slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHopIn []interface{}

    // List of address families which the NLRI may be within. The type is slice of
    // ['IPV4UNICAST', 'IPV6UNICAST', 'IPV4LABELEDUNICAST', 'IPV6LABELEDUNICAST',
    // 'L3VPNIPV4UNICAST', 'L3VPNIPV6UNICAST', 'L3VPNIPV4MULTICAST',
    // 'L3VPNIPV6MULTICAST', 'L2VPNVPLS', 'L2VPNEVPN'].
    AfiSafiIn []interface{}

    // Condition to check if the local pref attribute is equal to the specified
    // value. The type is interface{} with range: 0..4294967295.
    LocalPrefEq interface{}

    // Condition to check the route type in the route update. The type is
    // RouteType.
    RouteType interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "bgp-conditions"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["med-eq"] = types.YLeaf{"MedEq", config.MedEq}
    config.EntityData.Leafs["origin-eq"] = types.YLeaf{"OriginEq", config.OriginEq}
    config.EntityData.Leafs["next-hop-in"] = types.YLeaf{"NextHopIn", config.NextHopIn}
    config.EntityData.Leafs["afi-safi-in"] = types.YLeaf{"AfiSafiIn", config.AfiSafiIn}
    config.EntityData.Leafs["local-pref-eq"] = types.YLeaf{"LocalPrefEq", config.LocalPrefEq}
    config.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", config.RouteType}
    return &(config.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Condition to check if the received MED value is equal to the specified
    // value. The type is interface{} with range: 0..4294967295.
    MedEq interface{}

    // Condition to check if the route origin is equal to the specified value. The
    // type is BgpOriginAttrType.
    OriginEq interface{}

    // List of next hop addresses to check for in the route update. The type is
    // one of the following types: slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHopIn []interface{}

    // List of address families which the NLRI may be within. The type is slice of
    // ['IPV4UNICAST', 'IPV6UNICAST', 'IPV4LABELEDUNICAST', 'IPV6LABELEDUNICAST',
    // 'L3VPNIPV4UNICAST', 'L3VPNIPV6UNICAST', 'L3VPNIPV4MULTICAST',
    // 'L3VPNIPV6MULTICAST', 'L2VPNVPLS', 'L2VPNEVPN'].
    AfiSafiIn []interface{}

    // Condition to check if the local pref attribute is equal to the specified
    // value. The type is interface{} with range: 0..4294967295.
    LocalPrefEq interface{}

    // Condition to check the route type in the route update. The type is
    // RouteType.
    RouteType interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "bgp-conditions"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["med-eq"] = types.YLeaf{"MedEq", state.MedEq}
    state.EntityData.Leafs["origin-eq"] = types.YLeaf{"OriginEq", state.OriginEq}
    state.EntityData.Leafs["next-hop-in"] = types.YLeaf{"NextHopIn", state.NextHopIn}
    state.EntityData.Leafs["afi-safi-in"] = types.YLeaf{"AfiSafiIn", state.AfiSafiIn}
    state.EntityData.Leafs["local-pref-eq"] = types.YLeaf{"LocalPrefEq", state.LocalPrefEq}
    state.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", state.RouteType}
    return &(state.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for community count condition.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config

    // Operational state data for community count condition.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State
}

func (communityCount *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount) GetEntityData() *types.CommonEntityData {
    communityCount.EntityData.YFilter = communityCount.YFilter
    communityCount.EntityData.YangName = "community-count"
    communityCount.EntityData.BundleName = "openconfig"
    communityCount.EntityData.ParentYangName = "bgp-conditions"
    communityCount.EntityData.SegmentPath = "community-count"
    communityCount.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    communityCount.EntityData.NamespaceTable = openconfig.GetNamespaces()
    communityCount.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    communityCount.EntityData.Children = make(map[string]types.YChild)
    communityCount.EntityData.Children["config"] = types.YChild{"Config", &communityCount.Config}
    communityCount.EntityData.Children["state"] = types.YChild{"State", &communityCount.State}
    communityCount.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(communityCount.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config
// Configuration data for community count condition
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type of comparison to be performed. The type is one of the following:
    // ATTRIBUTEEQATTRIBUTEGEATTRIBUTELE.
    Operator interface{}

    // value to compare with the community count. The type is interface{} with
    // range: 0..4294967295.
    Value interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "community-count"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["operator"] = types.YLeaf{"Operator", config.Operator}
    config.EntityData.Leafs["value"] = types.YLeaf{"Value", config.Value}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State
// Operational state data for community count condition
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type of comparison to be performed. The type is one of the following:
    // ATTRIBUTEEQATTRIBUTEGEATTRIBUTELE.
    Operator interface{}

    // value to compare with the community count. The type is interface{} with
    // range: 0..4294967295.
    Value interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_CommunityCount_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "community-count"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["operator"] = types.YLeaf{"Operator", state.Operator}
    state.EntityData.Leafs["value"] = types.YLeaf{"Value", state.Value}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength
// Value and comparison operations for conditions based on the
// length of the AS path in the route update
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for AS path length condition.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config

    // Operational state data for AS path length condition.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State
}

func (asPathLength *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength) GetEntityData() *types.CommonEntityData {
    asPathLength.EntityData.YFilter = asPathLength.YFilter
    asPathLength.EntityData.YangName = "as-path-length"
    asPathLength.EntityData.BundleName = "openconfig"
    asPathLength.EntityData.ParentYangName = "bgp-conditions"
    asPathLength.EntityData.SegmentPath = "as-path-length"
    asPathLength.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    asPathLength.EntityData.NamespaceTable = openconfig.GetNamespaces()
    asPathLength.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    asPathLength.EntityData.Children = make(map[string]types.YChild)
    asPathLength.EntityData.Children["config"] = types.YChild{"Config", &asPathLength.Config}
    asPathLength.EntityData.Children["state"] = types.YChild{"State", &asPathLength.State}
    asPathLength.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asPathLength.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config
// Configuration data for AS path length condition
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type of comparison to be performed. The type is one of the following:
    // ATTRIBUTEEQATTRIBUTEGEATTRIBUTELE.
    Operator interface{}

    // value to compare with the community count. The type is interface{} with
    // range: 0..4294967295.
    Value interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "as-path-length"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["operator"] = types.YLeaf{"Operator", config.Operator}
    config.EntityData.Leafs["value"] = types.YLeaf{"Value", config.Value}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State
// Operational state data for AS path length condition
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type of comparison to be performed. The type is one of the following:
    // ATTRIBUTEEQATTRIBUTEGEATTRIBUTELE.
    Operator interface{}

    // value to compare with the community count. The type is interface{} with
    // range: 0..4294967295.
    Value interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_AsPathLength_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "as-path-length"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["operator"] = types.YLeaf{"Operator", state.Operator}
    state.EntityData.Leafs["value"] = types.YLeaf{"Value", state.Value}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet
// Top-level container for match conditions on communities.
// Match a referenced community-set according to the logic
// defined in the match-set-options leaf
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for match conditions on communities.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config

    // Operational state data .
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State
}

func (matchCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet) GetEntityData() *types.CommonEntityData {
    matchCommunitySet.EntityData.YFilter = matchCommunitySet.YFilter
    matchCommunitySet.EntityData.YangName = "match-community-set"
    matchCommunitySet.EntityData.BundleName = "openconfig"
    matchCommunitySet.EntityData.ParentYangName = "bgp-conditions"
    matchCommunitySet.EntityData.SegmentPath = "match-community-set"
    matchCommunitySet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    matchCommunitySet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    matchCommunitySet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    matchCommunitySet.EntityData.Children = make(map[string]types.YChild)
    matchCommunitySet.EntityData.Children["config"] = types.YChild{"Config", &matchCommunitySet.Config}
    matchCommunitySet.EntityData.Children["state"] = types.YChild{"State", &matchCommunitySet.State}
    matchCommunitySet.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(matchCommunitySet.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config
// Configuration data for match conditions on communities
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // References a defined community set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_CommunitySetName
    CommunitySet interface{}

    // Optional parameter that governs the behaviour of the match operation. The
    // type is MatchSetOptionsType.
    MatchSetOptions interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "match-community-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["community-set"] = types.YLeaf{"CommunitySet", config.CommunitySet}
    config.EntityData.Leafs["match-set-options"] = types.YLeaf{"MatchSetOptions", config.MatchSetOptions}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State
// Operational state data 
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // References a defined community set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_CommunitySetName
    CommunitySet interface{}

    // Optional parameter that governs the behaviour of the match operation. The
    // type is MatchSetOptionsType.
    MatchSetOptions interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchCommunitySet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "match-community-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["community-set"] = types.YLeaf{"CommunitySet", state.CommunitySet}
    state.EntityData.Leafs["match-set-options"] = types.YLeaf{"MatchSetOptions", state.MatchSetOptions}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet
// Match a referenced extended community-set according to the
// logic defined in the match-set-options leaf
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for match conditions on extended communities.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config

    // Operational state data for match conditions on extended communities.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State
}

func (matchExtCommunitySet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet) GetEntityData() *types.CommonEntityData {
    matchExtCommunitySet.EntityData.YFilter = matchExtCommunitySet.YFilter
    matchExtCommunitySet.EntityData.YangName = "match-ext-community-set"
    matchExtCommunitySet.EntityData.BundleName = "openconfig"
    matchExtCommunitySet.EntityData.ParentYangName = "bgp-conditions"
    matchExtCommunitySet.EntityData.SegmentPath = "match-ext-community-set"
    matchExtCommunitySet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    matchExtCommunitySet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    matchExtCommunitySet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    matchExtCommunitySet.EntityData.Children = make(map[string]types.YChild)
    matchExtCommunitySet.EntityData.Children["config"] = types.YChild{"Config", &matchExtCommunitySet.Config}
    matchExtCommunitySet.EntityData.Children["state"] = types.YChild{"State", &matchExtCommunitySet.State}
    matchExtCommunitySet.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(matchExtCommunitySet.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config
// Configuration data for match conditions on extended
// communities
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // References a defined extended community set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_ExtCommunitySetName
    ExtCommunitySet interface{}

    // Optional parameter that governs the behaviour of the match operation. The
    // type is MatchSetOptionsType.
    MatchSetOptions interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "match-ext-community-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["ext-community-set"] = types.YLeaf{"ExtCommunitySet", config.ExtCommunitySet}
    config.EntityData.Leafs["match-set-options"] = types.YLeaf{"MatchSetOptions", config.MatchSetOptions}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State
// Operational state data for match conditions on extended
// communities
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // References a defined extended community set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_ExtCommunitySetName
    ExtCommunitySet interface{}

    // Optional parameter that governs the behaviour of the match operation. The
    // type is MatchSetOptionsType.
    MatchSetOptions interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchExtCommunitySet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "match-ext-community-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["ext-community-set"] = types.YLeaf{"ExtCommunitySet", state.ExtCommunitySet}
    state.EntityData.Leafs["match-set-options"] = types.YLeaf{"MatchSetOptions", state.MatchSetOptions}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet
// Match a referenced as-path set according to the logic
// defined in the match-set-options leaf
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for match conditions on AS path set.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config

    // Operational state data for match conditions on AS path set.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State
}

func (matchAsPathSet *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet) GetEntityData() *types.CommonEntityData {
    matchAsPathSet.EntityData.YFilter = matchAsPathSet.YFilter
    matchAsPathSet.EntityData.YangName = "match-as-path-set"
    matchAsPathSet.EntityData.BundleName = "openconfig"
    matchAsPathSet.EntityData.ParentYangName = "bgp-conditions"
    matchAsPathSet.EntityData.SegmentPath = "match-as-path-set"
    matchAsPathSet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    matchAsPathSet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    matchAsPathSet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    matchAsPathSet.EntityData.Children = make(map[string]types.YChild)
    matchAsPathSet.EntityData.Children["config"] = types.YChild{"Config", &matchAsPathSet.Config}
    matchAsPathSet.EntityData.Children["state"] = types.YChild{"State", &matchAsPathSet.State}
    matchAsPathSet.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(matchAsPathSet.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config
// Configuration data for match conditions on AS path set
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // References a defined AS path set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_AsPathSetName
    AsPathSet interface{}

    // Optional parameter that governs the behaviour of the match operation. The
    // type is MatchSetOptionsType.
    MatchSetOptions interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "match-as-path-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["as-path-set"] = types.YLeaf{"AsPathSet", config.AsPathSet}
    config.EntityData.Leafs["match-set-options"] = types.YLeaf{"MatchSetOptions", config.MatchSetOptions}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State
// Operational state data for match conditions on AS
// path set
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // References a defined AS path set. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSets_AsPathSet_AsPathSetName
    AsPathSet interface{}

    // Optional parameter that governs the behaviour of the match operation. The
    // type is MatchSetOptionsType.
    MatchSetOptions interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_BgpConditions_MatchAsPathSet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "match-as-path-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["as-path-set"] = types.YLeaf{"AsPathSet", state.AsPathSet}
    state.EntityData.Leafs["match-set-options"] = types.YLeaf{"MatchSetOptions", state.MatchSetOptions}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IsisConditions
// Match conditions relating to the IS-IS protocol
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IsisConditions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to IS-IS match conditions.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IsisConditions_Config

    // Operational state parameters relating to IS-IS match conditions.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IsisConditions_State
}

func (isisConditions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IsisConditions) GetEntityData() *types.CommonEntityData {
    isisConditions.EntityData.YFilter = isisConditions.YFilter
    isisConditions.EntityData.YangName = "isis-conditions"
    isisConditions.EntityData.BundleName = "openconfig"
    isisConditions.EntityData.ParentYangName = "conditions"
    isisConditions.EntityData.SegmentPath = "openconfig-isis-policy:isis-conditions"
    isisConditions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    isisConditions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    isisConditions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    isisConditions.EntityData.Children = make(map[string]types.YChild)
    isisConditions.EntityData.Children["config"] = types.YChild{"Config", &isisConditions.Config}
    isisConditions.EntityData.Children["state"] = types.YChild{"State", &isisConditions.State}
    isisConditions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(isisConditions.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IsisConditions_Config
// Configuration parameters relating to IS-IS match
// conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IsisConditions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Match the level that the IS-IS prefix is within. This can be used in the
    // case that import or export policies refer to an IS-IS instance that has
    // multiple levels configured within it. The type is interface{} with range:
    // 1..2.
    LevelEq interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IsisConditions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "isis-conditions"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["level-eq"] = types.YLeaf{"LevelEq", config.LevelEq}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IsisConditions_State
// Operational state parameters relating to IS-IS match
// conditions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IsisConditions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Match the level that the IS-IS prefix is within. This can be used in the
    // case that import or export policies refer to an IS-IS instance that has
    // multiple levels configured within it. The type is interface{} with range:
    // 1..2.
    LevelEq interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Conditions_IsisConditions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "isis-conditions"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["level-eq"] = types.YLeaf{"LevelEq", state.LevelEq}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions
// Top-level container for policy action statements
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for policy actions.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config

    // Operational state data for policy actions.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State

    // Actions to set IGP route attributes; these actions apply to multiple IGPs.
    IgpActions RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions

    // Top-level container for BGP-specific actions.
    BgpActions RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions

    // Actions that can be performed by IS-IS within a policy.
    IsisActions RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IsisActions
}

func (actions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions) GetEntityData() *types.CommonEntityData {
    actions.EntityData.YFilter = actions.YFilter
    actions.EntityData.YangName = "actions"
    actions.EntityData.BundleName = "openconfig"
    actions.EntityData.ParentYangName = "statement"
    actions.EntityData.SegmentPath = "actions"
    actions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    actions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    actions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    actions.EntityData.Children = make(map[string]types.YChild)
    actions.EntityData.Children["config"] = types.YChild{"Config", &actions.Config}
    actions.EntityData.Children["state"] = types.YChild{"State", &actions.State}
    actions.EntityData.Children["igp-actions"] = types.YChild{"IgpActions", &actions.IgpActions}
    actions.EntityData.Children["openconfig-bgp-policy:bgp-actions"] = types.YChild{"BgpActions", &actions.BgpActions}
    actions.EntityData.Children["openconfig-isis-policy:isis-actions"] = types.YChild{"IsisActions", &actions.IsisActions}
    actions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(actions.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config
// Configuration data for policy actions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // accepts the route into the routing table. The type is interface{}.
    AcceptRoute interface{}

    // rejects the route. The type is interface{}.
    RejectRoute interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "actions"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["accept-route"] = types.YLeaf{"AcceptRoute", config.AcceptRoute}
    config.EntityData.Leafs["reject-route"] = types.YLeaf{"RejectRoute", config.RejectRoute}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State
// Operational state data for policy actions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // accepts the route into the routing table. The type is interface{}.
    AcceptRoute interface{}

    // rejects the route. The type is interface{}.
    RejectRoute interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "actions"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["accept-route"] = types.YLeaf{"AcceptRoute", state.AcceptRoute}
    state.EntityData.Leafs["reject-route"] = types.YLeaf{"RejectRoute", state.RejectRoute}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions
// Actions to set IGP route attributes; these actions
// apply to multiple IGPs
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data .
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config

    // Operational state data .
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State
}

func (igpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions) GetEntityData() *types.CommonEntityData {
    igpActions.EntityData.YFilter = igpActions.YFilter
    igpActions.EntityData.YangName = "igp-actions"
    igpActions.EntityData.BundleName = "openconfig"
    igpActions.EntityData.ParentYangName = "actions"
    igpActions.EntityData.SegmentPath = "igp-actions"
    igpActions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    igpActions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    igpActions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    igpActions.EntityData.Children = make(map[string]types.YChild)
    igpActions.EntityData.Children["config"] = types.YChild{"Config", &igpActions.Config}
    igpActions.EntityData.Children["state"] = types.YChild{"State", &igpActions.State}
    igpActions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(igpActions.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config
// Configuration data 
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the tag value for OSPF or IS-IS routes. The type is one of the
    // following types: int with range: 0..4294967295, or string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    SetTag interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "igp-actions"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["set-tag"] = types.YLeaf{"SetTag", config.SetTag}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State
// Operational state data 
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the tag value for OSPF or IS-IS routes. The type is one of the
    // following types: int with range: 0..4294967295, or string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    SetTag interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IgpActions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "igp-actions"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["set-tag"] = types.YLeaf{"SetTag", state.SetTag}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions
// Top-level container for BGP-specific actions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions struct {
    EntityData types.CommonEntityData
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

func (bgpActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions) GetEntityData() *types.CommonEntityData {
    bgpActions.EntityData.YFilter = bgpActions.YFilter
    bgpActions.EntityData.YangName = "bgp-actions"
    bgpActions.EntityData.BundleName = "openconfig"
    bgpActions.EntityData.ParentYangName = "actions"
    bgpActions.EntityData.SegmentPath = "openconfig-bgp-policy:bgp-actions"
    bgpActions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    bgpActions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    bgpActions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    bgpActions.EntityData.Children = make(map[string]types.YChild)
    bgpActions.EntityData.Children["config"] = types.YChild{"Config", &bgpActions.Config}
    bgpActions.EntityData.Children["state"] = types.YChild{"State", &bgpActions.State}
    bgpActions.EntityData.Children["set-as-path-prepend"] = types.YChild{"SetAsPathPrepend", &bgpActions.SetAsPathPrepend}
    bgpActions.EntityData.Children["set-community"] = types.YChild{"SetCommunity", &bgpActions.SetCommunity}
    bgpActions.EntityData.Children["set-ext-community"] = types.YChild{"SetExtCommunity", &bgpActions.SetExtCommunity}
    bgpActions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpActions.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config
// Configuration data for BGP-specific actions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // set the origin attribute to the specified value. The type is
    // BgpOriginAttrType.
    SetRouteOrigin interface{}

    // set the local pref attribute on the route update. The type is interface{}
    // with range: 0..4294967295.
    SetLocalPref interface{}

    // set the next-hop attribute in the route update. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.,
    // or enumeration BgpNextHopType.
    SetNextHop interface{}

    // set the med metric attribute in the route update. The type is one of the
    // following types: int with range: 0..4294967295, or string with pattern:
    // b'^[+-][0-9]+', or enumeration BgpSetMedType.
    SetMed interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "bgp-actions"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["set-route-origin"] = types.YLeaf{"SetRouteOrigin", config.SetRouteOrigin}
    config.EntityData.Leafs["set-local-pref"] = types.YLeaf{"SetLocalPref", config.SetLocalPref}
    config.EntityData.Leafs["set-next-hop"] = types.YLeaf{"SetNextHop", config.SetNextHop}
    config.EntityData.Leafs["set-med"] = types.YLeaf{"SetMed", config.SetMed}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State
// Operational state data for BGP-specific actions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // set the origin attribute to the specified value. The type is
    // BgpOriginAttrType.
    SetRouteOrigin interface{}

    // set the local pref attribute on the route update. The type is interface{}
    // with range: 0..4294967295.
    SetLocalPref interface{}

    // set the next-hop attribute in the route update. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.,
    // or enumeration BgpNextHopType.
    SetNextHop interface{}

    // set the med metric attribute in the route update. The type is one of the
    // following types: int with range: 0..4294967295, or string with pattern:
    // b'^[+-][0-9]+', or enumeration BgpSetMedType.
    SetMed interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "bgp-actions"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["set-route-origin"] = types.YLeaf{"SetRouteOrigin", state.SetRouteOrigin}
    state.EntityData.Leafs["set-local-pref"] = types.YLeaf{"SetLocalPref", state.SetLocalPref}
    state.EntityData.Leafs["set-next-hop"] = types.YLeaf{"SetNextHop", state.SetNextHop}
    state.EntityData.Leafs["set-med"] = types.YLeaf{"SetMed", state.SetMed}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend
// action to prepend local AS number to the AS-path a
// specified number of times
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for the AS path prepend action.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config

    // Operational state data for the AS path prepend action.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State
}

func (setAsPathPrepend *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend) GetEntityData() *types.CommonEntityData {
    setAsPathPrepend.EntityData.YFilter = setAsPathPrepend.YFilter
    setAsPathPrepend.EntityData.YangName = "set-as-path-prepend"
    setAsPathPrepend.EntityData.BundleName = "openconfig"
    setAsPathPrepend.EntityData.ParentYangName = "bgp-actions"
    setAsPathPrepend.EntityData.SegmentPath = "set-as-path-prepend"
    setAsPathPrepend.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    setAsPathPrepend.EntityData.NamespaceTable = openconfig.GetNamespaces()
    setAsPathPrepend.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    setAsPathPrepend.EntityData.Children = make(map[string]types.YChild)
    setAsPathPrepend.EntityData.Children["config"] = types.YChild{"Config", &setAsPathPrepend.Config}
    setAsPathPrepend.EntityData.Children["state"] = types.YChild{"State", &setAsPathPrepend.State}
    setAsPathPrepend.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(setAsPathPrepend.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config
// Configuration data for the AS path prepend action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of times to prepend the local AS number to the AS path.  The value
    // should be between 1 and the maximum supported by the implementation. The
    // type is interface{} with range: 1..255.
    RepeatN interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "set-as-path-prepend"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["repeat-n"] = types.YLeaf{"RepeatN", config.RepeatN}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State
// Operational state data for the AS path prepend action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of times to prepend the local AS number to the AS path.  The value
    // should be between 1 and the maximum supported by the implementation. The
    // type is interface{} with range: 1..255.
    RepeatN interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetAsPathPrepend_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "set-as-path-prepend"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["repeat-n"] = types.YLeaf{"RepeatN", state.RepeatN}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity
// Action to set the community attributes of the route, along
// with options to modify how the community is modified.
// Communities may be set using an inline list OR
// reference to an existing defined set (not both).
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity struct {
    EntityData types.CommonEntityData
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

func (setCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity) GetEntityData() *types.CommonEntityData {
    setCommunity.EntityData.YFilter = setCommunity.YFilter
    setCommunity.EntityData.YangName = "set-community"
    setCommunity.EntityData.BundleName = "openconfig"
    setCommunity.EntityData.ParentYangName = "bgp-actions"
    setCommunity.EntityData.SegmentPath = "set-community"
    setCommunity.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    setCommunity.EntityData.NamespaceTable = openconfig.GetNamespaces()
    setCommunity.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    setCommunity.EntityData.Children = make(map[string]types.YChild)
    setCommunity.EntityData.Children["config"] = types.YChild{"Config", &setCommunity.Config}
    setCommunity.EntityData.Children["state"] = types.YChild{"State", &setCommunity.State}
    setCommunity.EntityData.Children["inline"] = types.YChild{"Inline", &setCommunity.Inline}
    setCommunity.EntityData.Children["reference"] = types.YChild{"Reference", &setCommunity.Reference}
    setCommunity.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(setCommunity.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config
// Configuration data for the set-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates the method used to specify the extended communities for the
    // set-ext-community action. The type is Method.
    Method interface{}

    // Options for modifying the community attribute with the specified values. 
    // These options apply to both methods of setting the community attribute. The
    // type is BgpSetCommunityOptionType.
    Options interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "set-community"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["method"] = types.YLeaf{"Method", config.Method}
    config.EntityData.Leafs["options"] = types.YLeaf{"Options", config.Options}
    return &(config.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates the method used to specify the extended communities for the
    // set-ext-community action. The type is Method.
    Method interface{}

    // Options for modifying the community attribute with the specified values. 
    // These options apply to both methods of setting the community attribute. The
    // type is BgpSetCommunityOptionType.
    Options interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "set-community"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["method"] = types.YLeaf{"Method", state.Method}
    state.EntityData.Leafs["options"] = types.YLeaf{"Options", state.Options}
    return &(state.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data or inline specification of set-community action.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config

    // Operational state data or inline specification of set-community action.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline) GetEntityData() *types.CommonEntityData {
    inline.EntityData.YFilter = inline.YFilter
    inline.EntityData.YangName = "inline"
    inline.EntityData.BundleName = "openconfig"
    inline.EntityData.ParentYangName = "set-community"
    inline.EntityData.SegmentPath = "inline"
    inline.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    inline.EntityData.NamespaceTable = openconfig.GetNamespaces()
    inline.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    inline.EntityData.Children = make(map[string]types.YChild)
    inline.EntityData.Children["config"] = types.YChild{"Config", &inline.Config}
    inline.EntityData.Children["state"] = types.YChild{"State", &inline.State}
    inline.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(inline.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config
// Configuration data or inline specification of set-community
// action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the community values for the update inline with a list. The type is one
    // of the following types: slice of int with range: 65536..4294901759, or
    // slice of string with pattern: b'([0-9]+:[0-9]+)'., or slice of  
    // :go:struct:`BGPWELLKNOWNSTDCOMMUNITY
    // <ydk/models/bgp_types/BGPWELLKNOWNSTDCOMMUNITY>`.
    Communities []interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "inline"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["communities"] = types.YLeaf{"Communities", config.Communities}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State
// Operational state data or inline specification of
// set-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the community values for the update inline with a list. The type is one
    // of the following types: slice of int with range: 65536..4294901759, or
    // slice of string with pattern: b'([0-9]+:[0-9]+)'., or slice of  
    // :go:struct:`BGPWELLKNOWNSTDCOMMUNITY
    // <ydk/models/bgp_types/BGPWELLKNOWNSTDCOMMUNITY>`.
    Communities []interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Inline_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "inline"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["communities"] = types.YLeaf{"Communities", state.Communities}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference
// Provide a reference to a defined community set for the
// set-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for referening a community-set in the set-community
    // action.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config

    // Operational state data for referening a community-set in the set-community
    // action.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "openconfig"
    reference.EntityData.ParentYangName = "set-community"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    reference.EntityData.NamespaceTable = openconfig.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Children["config"] = types.YChild{"Config", &reference.Config}
    reference.EntityData.Children["state"] = types.YChild{"State", &reference.State}
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reference.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config
// Configuration data for referening a community-set in the
// set-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // References a defined community set by name. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_CommunitySetName
    CommunitySetRef interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "reference"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["community-set-ref"] = types.YLeaf{"CommunitySetRef", config.CommunitySetRef}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State
// Operational state data for referening a community-set
// in the set-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // References a defined community set by name. The type is string. Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySets_CommunitySet_CommunitySetName
    CommunitySetRef interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetCommunity_Reference_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "reference"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["community-set-ref"] = types.YLeaf{"CommunitySetRef", state.CommunitySetRef}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity
// Action to set the extended community attributes of the
// route, along with options to modify how the community is
// modified. Extended communities may be set using an inline
// list OR a reference to an existing defined set (but not
// both).
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity struct {
    EntityData types.CommonEntityData
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

func (setExtCommunity *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity) GetEntityData() *types.CommonEntityData {
    setExtCommunity.EntityData.YFilter = setExtCommunity.YFilter
    setExtCommunity.EntityData.YangName = "set-ext-community"
    setExtCommunity.EntityData.BundleName = "openconfig"
    setExtCommunity.EntityData.ParentYangName = "bgp-actions"
    setExtCommunity.EntityData.SegmentPath = "set-ext-community"
    setExtCommunity.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    setExtCommunity.EntityData.NamespaceTable = openconfig.GetNamespaces()
    setExtCommunity.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    setExtCommunity.EntityData.Children = make(map[string]types.YChild)
    setExtCommunity.EntityData.Children["config"] = types.YChild{"Config", &setExtCommunity.Config}
    setExtCommunity.EntityData.Children["state"] = types.YChild{"State", &setExtCommunity.State}
    setExtCommunity.EntityData.Children["inline"] = types.YChild{"Inline", &setExtCommunity.Inline}
    setExtCommunity.EntityData.Children["reference"] = types.YChild{"Reference", &setExtCommunity.Reference}
    setExtCommunity.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(setExtCommunity.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config
// Configuration data for the set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates the method used to specify the extended communities for the
    // set-ext-community action. The type is Method.
    Method interface{}

    // Options for modifying the community attribute with the specified values. 
    // These options apply to both methods of setting the community attribute. The
    // type is BgpSetCommunityOptionType.
    Options interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "set-ext-community"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["method"] = types.YLeaf{"Method", config.Method}
    config.EntityData.Leafs["options"] = types.YLeaf{"Options", config.Options}
    return &(config.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates the method used to specify the extended communities for the
    // set-ext-community action. The type is Method.
    Method interface{}

    // Options for modifying the community attribute with the specified values. 
    // These options apply to both methods of setting the community attribute. The
    // type is BgpSetCommunityOptionType.
    Options interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "set-ext-community"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["method"] = types.YLeaf{"Method", state.Method}
    state.EntityData.Leafs["options"] = types.YLeaf{"Options", state.Options}
    return &(state.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data or inline specification of set-ext-community action.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config

    // Operational state data or inline specification of set-ext-community action.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State
}

func (inline *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline) GetEntityData() *types.CommonEntityData {
    inline.EntityData.YFilter = inline.YFilter
    inline.EntityData.YangName = "inline"
    inline.EntityData.BundleName = "openconfig"
    inline.EntityData.ParentYangName = "set-ext-community"
    inline.EntityData.SegmentPath = "inline"
    inline.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    inline.EntityData.NamespaceTable = openconfig.GetNamespaces()
    inline.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    inline.EntityData.Children = make(map[string]types.YChild)
    inline.EntityData.Children["config"] = types.YChild{"Config", &inline.Config}
    inline.EntityData.Children["state"] = types.YChild{"State", &inline.State}
    inline.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(inline.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config
// Configuration data or inline specification of
// set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the extended community values for the update inline with a list. The
    // type is one of the following types: slice of string with pattern:
    // b'(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9])',
    // or slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9])',
    // or slice of string with pattern:
    // b'route\\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9])',
    // or slice of string with pattern:
    // b'route\\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9])',
    // or slice of string with pattern:
    // b'route\\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9])',
    // or slice of string with pattern:
    // b'route\\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9])'.,
    // or slice of   :go:struct:`BGPWELLKNOWNSTDCOMMUNITY
    // <ydk/models/bgp_types/BGPWELLKNOWNSTDCOMMUNITY>`.
    Communities []interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "inline"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["communities"] = types.YLeaf{"Communities", config.Communities}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State
// Operational state data or inline specification of
// set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the extended community values for the update inline with a list. The
    // type is one of the following types: slice of string with pattern:
    // b'(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9])',
    // or slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9])',
    // or slice of string with pattern:
    // b'route\\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9])',
    // or slice of string with pattern:
    // b'route\\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9])',
    // or slice of string with pattern:
    // b'route\\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9])',
    // or slice of string with pattern:
    // b'route\\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9])'.,
    // or slice of   :go:struct:`BGPWELLKNOWNSTDCOMMUNITY
    // <ydk/models/bgp_types/BGPWELLKNOWNSTDCOMMUNITY>`.
    Communities []interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Inline_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "inline"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["communities"] = types.YLeaf{"Communities", state.Communities}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference
// Provide a reference to an extended community set for the
// set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for referening an extended community-set in the
    // set-ext-community action.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config

    // Operational state data for referening an extended community-set in the
    // set-ext-community action.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State
}

func (reference *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "openconfig"
    reference.EntityData.ParentYangName = "set-ext-community"
    reference.EntityData.SegmentPath = "reference"
    reference.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    reference.EntityData.NamespaceTable = openconfig.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Children["config"] = types.YChild{"Config", &reference.Config}
    reference.EntityData.Children["state"] = types.YChild{"State", &reference.State}
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reference.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config
// Configuration data for referening an extended
// community-set in the set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // References a defined extended community set by name. The type is string.
    // Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_ExtCommunitySetName
    ExtCommunitySetRef interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "reference"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["ext-community-set-ref"] = types.YLeaf{"ExtCommunitySetRef", config.ExtCommunitySetRef}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State
// Operational state data for referening an extended
// community-set in the set-ext-community action
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // References a defined extended community set by name. The type is string.
    // Refers to
    // routing_policy.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySets_ExtCommunitySet_ExtCommunitySetName
    ExtCommunitySetRef interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_BgpActions_SetExtCommunity_Reference_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "reference"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["ext-community-set-ref"] = types.YLeaf{"ExtCommunitySetRef", state.ExtCommunitySetRef}
    return &(state.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IsisActions
// Actions that can be performed by IS-IS within a policy
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IsisActions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to IS-IS actions.
    Config RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IsisActions_Config

    // Operational state associated with IS-IS actions.
    State RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IsisActions_State
}

func (isisActions *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IsisActions) GetEntityData() *types.CommonEntityData {
    isisActions.EntityData.YFilter = isisActions.YFilter
    isisActions.EntityData.YangName = "isis-actions"
    isisActions.EntityData.BundleName = "openconfig"
    isisActions.EntityData.ParentYangName = "actions"
    isisActions.EntityData.SegmentPath = "openconfig-isis-policy:isis-actions"
    isisActions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    isisActions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    isisActions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    isisActions.EntityData.Children = make(map[string]types.YChild)
    isisActions.EntityData.Children["config"] = types.YChild{"Config", &isisActions.Config}
    isisActions.EntityData.Children["state"] = types.YChild{"State", &isisActions.State}
    isisActions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(isisActions.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IsisActions_Config
// Configuration parameters relating to IS-IS actions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IsisActions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the level that a prefix is to be imported into. The type is interface{}
    // with range: 1..2.
    SetLevel interface{}

    // Set the type of metric that is to be specified when the set metric leaf is
    // specified. The type is interface{} with range: 1..2.
    SetMetricType interface{}

    // Set the metric of the IS-IS prefix. The type is interface{} with range:
    // 1..16777215.
    SetMetric interface{}
}

func (config *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IsisActions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "isis-actions"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["set-level"] = types.YLeaf{"SetLevel", config.SetLevel}
    config.EntityData.Leafs["set-metric-type"] = types.YLeaf{"SetMetricType", config.SetMetricType}
    config.EntityData.Leafs["set-metric"] = types.YLeaf{"SetMetric", config.SetMetric}
    return &(config.EntityData)
}

// RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IsisActions_State
// Operational state associated with IS-IS actions
type RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IsisActions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the level that a prefix is to be imported into. The type is interface{}
    // with range: 1..2.
    SetLevel interface{}

    // Set the type of metric that is to be specified when the set metric leaf is
    // specified. The type is interface{} with range: 1..2.
    SetMetricType interface{}

    // Set the metric of the IS-IS prefix. The type is interface{} with range:
    // 1..16777215.
    SetMetric interface{}
}

func (state *RoutingPolicy_PolicyDefinitions_PolicyDefinition_Statements_Statement_Actions_IsisActions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "isis-actions"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["set-level"] = types.YLeaf{"SetLevel", state.SetLevel}
    state.EntityData.Leafs["set-metric-type"] = types.YLeaf{"SetMetricType", state.SetMetricType}
    state.EntityData.Leafs["set-metric"] = types.YLeaf{"SetMetric", state.SetMetric}
    return &(state.EntityData)
}

