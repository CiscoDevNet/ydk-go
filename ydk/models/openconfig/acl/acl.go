// This module defines configuration and operational state
// data for network access control lists (i.e., filters, rules,
// etc.).  ACLs are organized into ACL sets, with each set
// containing one or more ACL entries.  ACL sets are identified
// by a unique name, while each entry within a set is assigned
// a sequence-id that determines the order in which the ACL
// rules are applied to a packet.
// 
// Individual ACL rules specify match criteria based on fields in
// the packet, along with an action that defines how matching
// packets should be handled. Entries have a type that indicates
// the type of match criteria, e.g., MAC layer, IPv4, IPv6, etc.
package acl

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package acl"))
    ydk.RegisterEntity("{http://openconfig.net/yang/acl acl}", reflect.TypeOf(Acl{}))
    ydk.RegisterEntity("openconfig-acl:acl", reflect.TypeOf(Acl{}))
}

type ACLL2 struct {
}

func (id ACLL2) String() string {
	return "openconfig-acl:ACL_L2"
}

type ACLIPV6 struct {
}

func (id ACLIPV6) String() string {
	return "openconfig-acl:ACL_IPV6"
}

type LOGNONE struct {
}

func (id LOGNONE) String() string {
	return "openconfig-acl:LOG_NONE"
}

type ACLIPV4 struct {
}

func (id ACLIPV4) String() string {
	return "openconfig-acl:ACL_IPV4"
}

type DROP struct {
}

func (id DROP) String() string {
	return "openconfig-acl:DROP"
}

type ACCEPT struct {
}

func (id ACCEPT) String() string {
	return "openconfig-acl:ACCEPT"
}

type FORWARDINGACTION struct {
}

func (id FORWARDINGACTION) String() string {
	return "openconfig-acl:FORWARDING_ACTION"
}

type INTERFACEONLY struct {
}

func (id INTERFACEONLY) String() string {
	return "openconfig-acl:INTERFACE_ONLY"
}

type ACLMIXED struct {
}

func (id ACLMIXED) String() string {
	return "openconfig-acl:ACL_MIXED"
}

type INTERFACEAGGREGATE struct {
}

func (id INTERFACEAGGREGATE) String() string {
	return "openconfig-acl:INTERFACE_AGGREGATE"
}

type REJECT struct {
}

func (id REJECT) String() string {
	return "openconfig-acl:REJECT"
}

type LOGSYSLOG struct {
}

func (id LOGSYSLOG) String() string {
	return "openconfig-acl:LOG_SYSLOG"
}

type AGGREGATEONLY struct {
}

func (id AGGREGATEONLY) String() string {
	return "openconfig-acl:AGGREGATE_ONLY"
}

type ACLTYPE struct {
}

func (id ACLTYPE) String() string {
	return "openconfig-acl:ACL_TYPE"
}

type ACLCOUNTERCAPABILITY struct {
}

func (id ACLCOUNTERCAPABILITY) String() string {
	return "openconfig-acl:ACL_COUNTER_CAPABILITY"
}

type LOGACTION struct {
}

func (id LOGACTION) String() string {
	return "openconfig-acl:LOG_ACTION"
}

// Acl
// Top level enclosing container for ACL model config
// and operational state data
type Acl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global config data for ACLs.
    Config Acl_Config

    // Global operational state data for ACLs.
    State Acl_State

    // Access list entries variables enclosing container.
    AclSets Acl_AclSets

    // Enclosing container for the list of interfaces on which ACLs are set.
    Interfaces Acl_Interfaces
}

func (acl *Acl) GetEntityData() *types.CommonEntityData {
    acl.EntityData.YFilter = acl.YFilter
    acl.EntityData.YangName = "acl"
    acl.EntityData.BundleName = "openconfig"
    acl.EntityData.ParentYangName = "openconfig-acl"
    acl.EntityData.SegmentPath = "openconfig-acl:acl"
    acl.EntityData.AbsolutePath = acl.EntityData.SegmentPath
    acl.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    acl.EntityData.NamespaceTable = openconfig.GetNamespaces()
    acl.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    acl.EntityData.Children = types.NewOrderedMap()
    acl.EntityData.Children.Append("config", types.YChild{"Config", &acl.Config})
    acl.EntityData.Children.Append("state", types.YChild{"State", &acl.State})
    acl.EntityData.Children.Append("acl-sets", types.YChild{"AclSets", &acl.AclSets})
    acl.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &acl.Interfaces})
    acl.EntityData.Leafs = types.NewOrderedMap()

    acl.EntityData.YListKeys = []string {}

    return &(acl.EntityData)
}

// Acl_Config
// Global config data for ACLs
type Acl_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (config *Acl_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "acl"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_State
// Global operational state data for ACLs
type Acl_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // System reported indication of how ACL counters are reported by the target.
    // The type is one of the following:
    // INTERFACEONLYINTERFACEAGGREGATEAGGREGATEONLY.
    CounterCapability interface{}
}

func (state *Acl_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "acl"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("counter-capability", types.YLeaf{"CounterCapability", state.CounterCapability})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_AclSets
// Access list entries variables enclosing container
type Acl_AclSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of ACL sets, each comprising of a list of ACL entries. The type is
    // slice of Acl_AclSets_AclSet.
    AclSet []*Acl_AclSets_AclSet
}

func (aclSets *Acl_AclSets) GetEntityData() *types.CommonEntityData {
    aclSets.EntityData.YFilter = aclSets.YFilter
    aclSets.EntityData.YangName = "acl-sets"
    aclSets.EntityData.BundleName = "openconfig"
    aclSets.EntityData.ParentYangName = "acl"
    aclSets.EntityData.SegmentPath = "acl-sets"
    aclSets.EntityData.AbsolutePath = "openconfig-acl:acl/" + aclSets.EntityData.SegmentPath
    aclSets.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aclSets.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aclSets.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aclSets.EntityData.Children = types.NewOrderedMap()
    aclSets.EntityData.Children.Append("acl-set", types.YChild{"AclSet", nil})
    for i := range aclSets.AclSet {
        aclSets.EntityData.Children.Append(types.GetSegmentPath(aclSets.AclSet[i]), types.YChild{"AclSet", aclSets.AclSet[i]})
    }
    aclSets.EntityData.Leafs = types.NewOrderedMap()

    aclSets.EntityData.YListKeys = []string {}

    return &(aclSets.EntityData)
}

// Acl_AclSets_AclSet
// List of ACL sets, each comprising of a list of ACL
// entries
type Acl_AclSets_AclSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the name list key. The type is
    // string. Refers to acl.Acl_AclSets_AclSet_Config_Name
    Name interface{}

    // This attribute is a key. Reference to the type list key. The type is one of
    // the following: ACLL2ACLIPV6ACLIPV4ACLMIXED.
    Type interface{}

    // Access list config.
    Config Acl_AclSets_AclSet_Config

    // Access list state information.
    State Acl_AclSets_AclSet_State

    // Access list entries container.
    AclEntries Acl_AclSets_AclSet_AclEntries
}

func (aclSet *Acl_AclSets_AclSet) GetEntityData() *types.CommonEntityData {
    aclSet.EntityData.YFilter = aclSet.YFilter
    aclSet.EntityData.YangName = "acl-set"
    aclSet.EntityData.BundleName = "openconfig"
    aclSet.EntityData.ParentYangName = "acl-sets"
    aclSet.EntityData.SegmentPath = "acl-set" + types.AddKeyToken(aclSet.Name, "name") + types.AddKeyToken(aclSet.Type, "type")
    aclSet.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/" + aclSet.EntityData.SegmentPath
    aclSet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aclSet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aclSet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aclSet.EntityData.Children = types.NewOrderedMap()
    aclSet.EntityData.Children.Append("config", types.YChild{"Config", &aclSet.Config})
    aclSet.EntityData.Children.Append("state", types.YChild{"State", &aclSet.State})
    aclSet.EntityData.Children.Append("acl-entries", types.YChild{"AclEntries", &aclSet.AclEntries})
    aclSet.EntityData.Leafs = types.NewOrderedMap()
    aclSet.EntityData.Leafs.Append("name", types.YLeaf{"Name", aclSet.Name})
    aclSet.EntityData.Leafs.Append("type", types.YLeaf{"Type", aclSet.Type})

    aclSet.EntityData.YListKeys = []string {"Name", "Type"}

    return &(aclSet.EntityData)
}

// Acl_AclSets_AclSet_Config
// Access list config
type Acl_AclSets_AclSet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the access-list set. The type is string.
    Name interface{}

    // The type determines the fields allowed in the ACL entries belonging to the
    // ACL set (e.g., IPv4, IPv6, etc.). The type is one of the following:
    // ACLL2ACLIPV6ACLIPV4ACLMIXED.
    Type interface{}

    // Description, or comment, for the ACL set. The type is string.
    Description interface{}
}

func (config *Acl_AclSets_AclSet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "acl-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("type", types.YLeaf{"Type", config.Type})
    config.EntityData.Leafs.Append("description", types.YLeaf{"Description", config.Description})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_AclSets_AclSet_State
// Access list state information
type Acl_AclSets_AclSet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the access-list set. The type is string.
    Name interface{}

    // The type determines the fields allowed in the ACL entries belonging to the
    // ACL set (e.g., IPv4, IPv6, etc.). The type is one of the following:
    // ACLL2ACLIPV6ACLIPV4ACLMIXED.
    Type interface{}

    // Description, or comment, for the ACL set. The type is string.
    Description interface{}
}

func (state *Acl_AclSets_AclSet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "acl-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("type", types.YLeaf{"Type", state.Type})
    state.EntityData.Leafs.Append("description", types.YLeaf{"Description", state.Description})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_AclSets_AclSet_AclEntries
// Access list entries container
type Acl_AclSets_AclSet_AclEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of ACL entries comprising an ACL set. The type is slice of
    // Acl_AclSets_AclSet_AclEntries_AclEntry.
    AclEntry []*Acl_AclSets_AclSet_AclEntries_AclEntry
}

func (aclEntries *Acl_AclSets_AclSet_AclEntries) GetEntityData() *types.CommonEntityData {
    aclEntries.EntityData.YFilter = aclEntries.YFilter
    aclEntries.EntityData.YangName = "acl-entries"
    aclEntries.EntityData.BundleName = "openconfig"
    aclEntries.EntityData.ParentYangName = "acl-set"
    aclEntries.EntityData.SegmentPath = "acl-entries"
    aclEntries.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/" + aclEntries.EntityData.SegmentPath
    aclEntries.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aclEntries.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aclEntries.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aclEntries.EntityData.Children = types.NewOrderedMap()
    aclEntries.EntityData.Children.Append("acl-entry", types.YChild{"AclEntry", nil})
    for i := range aclEntries.AclEntry {
        aclEntries.EntityData.Children.Append(types.GetSegmentPath(aclEntries.AclEntry[i]), types.YChild{"AclEntry", aclEntries.AclEntry[i]})
    }
    aclEntries.EntityData.Leafs = types.NewOrderedMap()

    aclEntries.EntityData.YListKeys = []string {}

    return &(aclEntries.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry
// List of ACL entries comprising an ACL set
type Acl_AclSets_AclSet_AclEntries_AclEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. references the list key. The type is string with
    // range: 0..4294967295. Refers to
    // acl.Acl_AclSets_AclSet_AclEntries_AclEntry_Config_SequenceId
    SequenceId interface{}

    // Access list entries config.
    Config Acl_AclSets_AclSet_AclEntries_AclEntry_Config

    // State information for ACL entries.
    State Acl_AclSets_AclSet_AclEntries_AclEntry_State

    // Ethernet header fields.
    L2 Acl_AclSets_AclSet_AclEntries_AclEntry_L2

    // Top level container for IPv4 match field data.
    Ipv4 Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4

    // Top-level container for IPv6 match field data.
    Ipv6 Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6

    // Transport fields container.
    Transport Acl_AclSets_AclSet_AclEntries_AclEntry_Transport

    // Input interface container.
    InputInterface Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface

    // Enclosing container for list of ACL actions associated with an entry.
    Actions Acl_AclSets_AclSet_AclEntries_AclEntry_Actions
}

func (aclEntry *Acl_AclSets_AclSet_AclEntries_AclEntry) GetEntityData() *types.CommonEntityData {
    aclEntry.EntityData.YFilter = aclEntry.YFilter
    aclEntry.EntityData.YangName = "acl-entry"
    aclEntry.EntityData.BundleName = "openconfig"
    aclEntry.EntityData.ParentYangName = "acl-entries"
    aclEntry.EntityData.SegmentPath = "acl-entry" + types.AddKeyToken(aclEntry.SequenceId, "sequence-id")
    aclEntry.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/" + aclEntry.EntityData.SegmentPath
    aclEntry.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aclEntry.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aclEntry.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aclEntry.EntityData.Children = types.NewOrderedMap()
    aclEntry.EntityData.Children.Append("config", types.YChild{"Config", &aclEntry.Config})
    aclEntry.EntityData.Children.Append("state", types.YChild{"State", &aclEntry.State})
    aclEntry.EntityData.Children.Append("l2", types.YChild{"L2", &aclEntry.L2})
    aclEntry.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &aclEntry.Ipv4})
    aclEntry.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &aclEntry.Ipv6})
    aclEntry.EntityData.Children.Append("transport", types.YChild{"Transport", &aclEntry.Transport})
    aclEntry.EntityData.Children.Append("input-interface", types.YChild{"InputInterface", &aclEntry.InputInterface})
    aclEntry.EntityData.Children.Append("actions", types.YChild{"Actions", &aclEntry.Actions})
    aclEntry.EntityData.Leafs = types.NewOrderedMap()
    aclEntry.EntityData.Leafs.Append("sequence-id", types.YLeaf{"SequenceId", aclEntry.SequenceId})

    aclEntry.EntityData.YListKeys = []string {"SequenceId"}

    return &(aclEntry.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_Config
// Access list entries config
type Acl_AclSets_AclSet_AclEntries_AclEntry_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The sequence id determines the order in which ACL entries are applied.  The
    // sequence id must be unique for each entry in an ACL set.  Target devices
    // should apply the ACL entry rules in the order determined by sequence id,
    // rather than the relying only on order in the list. The type is interface{}
    // with range: 0..4294967295.
    SequenceId interface{}

    // A user-defined description, or comment, for this Access List Entry. The
    // type is string.
    Description interface{}
}

func (config *Acl_AclSets_AclSet_AclEntries_AclEntry_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "acl-entry"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("sequence-id", types.YLeaf{"SequenceId", config.SequenceId})
    config.EntityData.Leafs.Append("description", types.YLeaf{"Description", config.Description})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_State
// State information for ACL entries
type Acl_AclSets_AclSet_AclEntries_AclEntry_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The sequence id determines the order in which ACL entries are applied.  The
    // sequence id must be unique for each entry in an ACL set.  Target devices
    // should apply the ACL entry rules in the order determined by sequence id,
    // rather than the relying only on order in the list. The type is interface{}
    // with range: 0..4294967295.
    SequenceId interface{}

    // A user-defined description, or comment, for this Access List Entry. The
    // type is string.
    Description interface{}

    // Count of the number of packets matching the current ACL entry.  An
    // implementation should provide this counter on a per-interface per-ACL-entry
    // if possible.  If an implementation only supports ACL counters per entry
    // (i.e., not broken out per interface), then the value should be equal to the
    // aggregate count across all interfaces.  An implementation that provides
    // counters per entry per interface is not required to also provide an
    // aggregate count, e.g., per entry -- the user is expected to be able
    // implement the required aggregation if such a count is needed. The type is
    // interface{} with range: 0..18446744073709551615.
    MatchedPackets interface{}

    // Count of the number of octets (bytes) matching the current ACL entry.  An
    // implementation should provide this counter on a per-interface per-ACL-entry
    // if possible.  If an implementation only supports ACL counters per entry
    // (i.e., not broken out per interface), then the value should be equal to the
    // aggregate count across all interfaces.  An implementation that provides
    // counters per entry per interface is not required to also provide an
    // aggregate count, e.g., per entry -- the user is expected to be able
    // implement the required aggregation if such a count is needed. The type is
    // interface{} with range: 0..18446744073709551615.
    MatchedOctets interface{}
}

func (state *Acl_AclSets_AclSet_AclEntries_AclEntry_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "acl-entry"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("sequence-id", types.YLeaf{"SequenceId", state.SequenceId})
    state.EntityData.Leafs.Append("description", types.YLeaf{"Description", state.Description})
    state.EntityData.Leafs.Append("matched-packets", types.YLeaf{"MatchedPackets", state.MatchedPackets})
    state.EntityData.Leafs.Append("matched-octets", types.YLeaf{"MatchedOctets", state.MatchedOctets})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_L2
// Ethernet header fields
type Acl_AclSets_AclSet_AclEntries_AclEntry_L2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data.
    Config Acl_AclSets_AclSet_AclEntries_AclEntry_L2_Config

    // State Information.
    State Acl_AclSets_AclSet_AclEntries_AclEntry_L2_State
}

func (l2 *Acl_AclSets_AclSet_AclEntries_AclEntry_L2) GetEntityData() *types.CommonEntityData {
    l2.EntityData.YFilter = l2.YFilter
    l2.EntityData.YangName = "l2"
    l2.EntityData.BundleName = "openconfig"
    l2.EntityData.ParentYangName = "acl-entry"
    l2.EntityData.SegmentPath = "l2"
    l2.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/" + l2.EntityData.SegmentPath
    l2.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l2.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l2.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l2.EntityData.Children = types.NewOrderedMap()
    l2.EntityData.Children.Append("config", types.YChild{"Config", &l2.Config})
    l2.EntityData.Children.Append("state", types.YChild{"State", &l2.State})
    l2.EntityData.Leafs = types.NewOrderedMap()

    l2.EntityData.YListKeys = []string {}

    return &(l2.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_L2_Config
// Configuration data
type Acl_AclSets_AclSet_AclEntries_AclEntry_L2_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source IEEE 802 MAC address. The type is string with pattern:
    // ^[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}$.
    SourceMac interface{}

    // Source IEEE 802 MAC address mask. The type is string with pattern:
    // ^[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}$.
    SourceMacMask interface{}

    // Destination IEEE 802 MAC address. The type is string with pattern:
    // ^[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}$.
    DestinationMac interface{}

    // Destination IEEE 802 MAC address mask. The type is string with pattern:
    // ^[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}$.
    DestinationMacMask interface{}

    // Ethertype field to match in Ethernet packets. The type is one of the
    // following types: int with range: 1..65535, or :go:struct:`ETHERTYPE
    // <ydk/models/packet_match_types/ETHERTYPE>`.
    Ethertype interface{}
}

func (config *Acl_AclSets_AclSet_AclEntries_AclEntry_L2_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "l2"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/l2/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("source-mac", types.YLeaf{"SourceMac", config.SourceMac})
    config.EntityData.Leafs.Append("source-mac-mask", types.YLeaf{"SourceMacMask", config.SourceMacMask})
    config.EntityData.Leafs.Append("destination-mac", types.YLeaf{"DestinationMac", config.DestinationMac})
    config.EntityData.Leafs.Append("destination-mac-mask", types.YLeaf{"DestinationMacMask", config.DestinationMacMask})
    config.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", config.Ethertype})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_L2_State
// State Information.
type Acl_AclSets_AclSet_AclEntries_AclEntry_L2_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source IEEE 802 MAC address. The type is string with pattern:
    // ^[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}$.
    SourceMac interface{}

    // Source IEEE 802 MAC address mask. The type is string with pattern:
    // ^[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}$.
    SourceMacMask interface{}

    // Destination IEEE 802 MAC address. The type is string with pattern:
    // ^[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}$.
    DestinationMac interface{}

    // Destination IEEE 802 MAC address mask. The type is string with pattern:
    // ^[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}$.
    DestinationMacMask interface{}

    // Ethertype field to match in Ethernet packets. The type is one of the
    // following types: int with range: 1..65535, or :go:struct:`ETHERTYPE
    // <ydk/models/packet_match_types/ETHERTYPE>`.
    Ethertype interface{}
}

func (state *Acl_AclSets_AclSet_AclEntries_AclEntry_L2_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "l2"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/l2/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("source-mac", types.YLeaf{"SourceMac", state.SourceMac})
    state.EntityData.Leafs.Append("source-mac-mask", types.YLeaf{"SourceMacMask", state.SourceMacMask})
    state.EntityData.Leafs.Append("destination-mac", types.YLeaf{"DestinationMac", state.DestinationMac})
    state.EntityData.Leafs.Append("destination-mac-mask", types.YLeaf{"DestinationMacMask", state.DestinationMacMask})
    state.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", state.Ethertype})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4
// Top level container for IPv4 match field data
type Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for IPv4 match fields.
    Config Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_Config

    // State information for IPv4 match fields.
    State Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_State
}

func (ipv4 *Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "openconfig"
    ipv4.EntityData.ParentYangName = "acl-entry"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("config", types.YChild{"Config", &ipv4.Config})
    ipv4.EntityData.Children.Append("state", types.YChild{"State", &ipv4.State})
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_Config
// Configuration data for IPv4 match fields
type Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source IPv4 address prefix. The type is string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))$.
    SourceAddress interface{}

    // Destination IPv4 address prefix. The type is string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))$.
    DestinationAddress interface{}

    // Value of diffserv codepoint. The type is interface{} with range: 0..63.
    Dscp interface{}

    // The protocol carried in the IP packet, expressed either as its IP protocol
    // number, or by a defined identity. The type is one of the following types:
    // int with range: 0..254, or :go:struct:`IPPROTOCOL
    // <ydk/models/packet_match_types/IPPROTOCOL>`.
    Protocol interface{}

    // The IP packet's hop limit -- known as TTL (in hops) in IPv4 packets, and
    // hop limit in IPv6. The type is interface{} with range: 0..255.
    HopLimit interface{}
}

func (config *Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ipv4"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", config.SourceAddress})
    config.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", config.DestinationAddress})
    config.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", config.Dscp})
    config.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", config.Protocol})
    config.EntityData.Leafs.Append("hop-limit", types.YLeaf{"HopLimit", config.HopLimit})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_State
// State information for IPv4 match fields
type Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source IPv4 address prefix. The type is string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))$.
    SourceAddress interface{}

    // Destination IPv4 address prefix. The type is string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))$.
    DestinationAddress interface{}

    // Value of diffserv codepoint. The type is interface{} with range: 0..63.
    Dscp interface{}

    // The protocol carried in the IP packet, expressed either as its IP protocol
    // number, or by a defined identity. The type is one of the following types:
    // int with range: 0..254, or :go:struct:`IPPROTOCOL
    // <ydk/models/packet_match_types/IPPROTOCOL>`.
    Protocol interface{}

    // The IP packet's hop limit -- known as TTL (in hops) in IPv4 packets, and
    // hop limit in IPv6. The type is interface{} with range: 0..255.
    HopLimit interface{}
}

func (state *Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ipv4"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", state.SourceAddress})
    state.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", state.DestinationAddress})
    state.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", state.Dscp})
    state.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", state.Protocol})
    state.EntityData.Leafs.Append("hop-limit", types.YLeaf{"HopLimit", state.HopLimit})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6
// Top-level container for IPv6 match field data
type Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for IPv6 match fields.
    Config Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_Config

    // Operational state data for IPv6 match fields.
    State Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_State
}

func (ipv6 *Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "openconfig"
    ipv6.EntityData.ParentYangName = "acl-entry"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/" + ipv6.EntityData.SegmentPath
    ipv6.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("config", types.YChild{"Config", &ipv6.Config})
    ipv6.EntityData.Children.Append("state", types.YChild{"State", &ipv6.State})
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_Config
// Configuration data for IPv6 match fields
type Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source IPv6 address prefix. The type is string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9])$.
    SourceAddress interface{}

    // Source IPv6 Flow label. The type is interface{} with range: 0..1048575.
    SourceFlowLabel interface{}

    // Destination IPv6 address prefix. The type is string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9])$.
    DestinationAddress interface{}

    // Destination IPv6 Flow label. The type is interface{} with range:
    // 0..1048575.
    DestinationFlowLabel interface{}

    // Value of diffserv codepoint. The type is interface{} with range: 0..63.
    Dscp interface{}

    // The protocol carried in the IP packet, expressed either as its IP protocol
    // number, or by a defined identity. The type is one of the following types:
    // int with range: 0..254, or :go:struct:`IPPROTOCOL
    // <ydk/models/packet_match_types/IPPROTOCOL>`.
    Protocol interface{}

    // The IP packet's hop limit -- known as TTL (in hops) in IPv4 packets, and
    // hop limit in IPv6. The type is interface{} with range: 0..255.
    HopLimit interface{}
}

func (config *Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ipv6"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", config.SourceAddress})
    config.EntityData.Leafs.Append("source-flow-label", types.YLeaf{"SourceFlowLabel", config.SourceFlowLabel})
    config.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", config.DestinationAddress})
    config.EntityData.Leafs.Append("destination-flow-label", types.YLeaf{"DestinationFlowLabel", config.DestinationFlowLabel})
    config.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", config.Dscp})
    config.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", config.Protocol})
    config.EntityData.Leafs.Append("hop-limit", types.YLeaf{"HopLimit", config.HopLimit})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_State
// Operational state data for IPv6 match fields
type Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source IPv6 address prefix. The type is string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9])$.
    SourceAddress interface{}

    // Source IPv6 Flow label. The type is interface{} with range: 0..1048575.
    SourceFlowLabel interface{}

    // Destination IPv6 address prefix. The type is string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9])$.
    DestinationAddress interface{}

    // Destination IPv6 Flow label. The type is interface{} with range:
    // 0..1048575.
    DestinationFlowLabel interface{}

    // Value of diffserv codepoint. The type is interface{} with range: 0..63.
    Dscp interface{}

    // The protocol carried in the IP packet, expressed either as its IP protocol
    // number, or by a defined identity. The type is one of the following types:
    // int with range: 0..254, or :go:struct:`IPPROTOCOL
    // <ydk/models/packet_match_types/IPPROTOCOL>`.
    Protocol interface{}

    // The IP packet's hop limit -- known as TTL (in hops) in IPv4 packets, and
    // hop limit in IPv6. The type is interface{} with range: 0..255.
    HopLimit interface{}
}

func (state *Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ipv6"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", state.SourceAddress})
    state.EntityData.Leafs.Append("source-flow-label", types.YLeaf{"SourceFlowLabel", state.SourceFlowLabel})
    state.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", state.DestinationAddress})
    state.EntityData.Leafs.Append("destination-flow-label", types.YLeaf{"DestinationFlowLabel", state.DestinationFlowLabel})
    state.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", state.Dscp})
    state.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", state.Protocol})
    state.EntityData.Leafs.Append("hop-limit", types.YLeaf{"HopLimit", state.HopLimit})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_Transport
// Transport fields container
type Acl_AclSets_AclSet_AclEntries_AclEntry_Transport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data.
    Config Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config

    // State data.
    State Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_State
}

func (transport *Acl_AclSets_AclSet_AclEntries_AclEntry_Transport) GetEntityData() *types.CommonEntityData {
    transport.EntityData.YFilter = transport.YFilter
    transport.EntityData.YangName = "transport"
    transport.EntityData.BundleName = "openconfig"
    transport.EntityData.ParentYangName = "acl-entry"
    transport.EntityData.SegmentPath = "transport"
    transport.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/" + transport.EntityData.SegmentPath
    transport.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    transport.EntityData.NamespaceTable = openconfig.GetNamespaces()
    transport.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    transport.EntityData.Children = types.NewOrderedMap()
    transport.EntityData.Children.Append("config", types.YChild{"Config", &transport.Config})
    transport.EntityData.Children.Append("state", types.YChild{"State", &transport.State})
    transport.EntityData.Leafs = types.NewOrderedMap()

    transport.EntityData.YListKeys = []string {}

    return &(transport.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config
// Configuration data
type Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source port or range. The type is one of the following types: string with
    // pattern:
    // ^(6[0-5][0-5][0-3][0-5]|[0-5]?[0-9]?[0-9]?[0-9]?[0-9]?)\.\.(6[0-5][0-5][0-3][0-5]|[0-5]?[0-9]?[0-9]?[0-9]?[0-9]?)$,
    // or int with range: 0..65535, or enumeration PortNumRange.
    SourcePort interface{}

    // Destination port or range. The type is one of the following types: string
    // with pattern:
    // ^(6[0-5][0-5][0-3][0-5]|[0-5]?[0-9]?[0-9]?[0-9]?[0-9]?)\.\.(6[0-5][0-5][0-3][0-5]|[0-5]?[0-9]?[0-9]?[0-9]?[0-9]?)$,
    // or int with range: 0..65535, or enumeration PortNumRange.
    DestinationPort interface{}

    // List of TCP flags to match. The type is slice of [u'TCPACK', u'TCPSYN',
    // u'TCPECE', u'TCPFIN', u'TCPRST', u'TCPURG', u'TCPPSH', u'TCPCWR'].
    TcpFlags []interface{}
}

func (config *Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "transport"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/transport/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("source-port", types.YLeaf{"SourcePort", config.SourcePort})
    config.EntityData.Leafs.Append("destination-port", types.YLeaf{"DestinationPort", config.DestinationPort})
    config.EntityData.Leafs.Append("tcp-flags", types.YLeaf{"TcpFlags", config.TcpFlags})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_State
// State data
type Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source port or range. The type is one of the following types: string with
    // pattern:
    // ^(6[0-5][0-5][0-3][0-5]|[0-5]?[0-9]?[0-9]?[0-9]?[0-9]?)\.\.(6[0-5][0-5][0-3][0-5]|[0-5]?[0-9]?[0-9]?[0-9]?[0-9]?)$,
    // or int with range: 0..65535, or enumeration PortNumRange.
    SourcePort interface{}

    // Destination port or range. The type is one of the following types: string
    // with pattern:
    // ^(6[0-5][0-5][0-3][0-5]|[0-5]?[0-9]?[0-9]?[0-9]?[0-9]?)\.\.(6[0-5][0-5][0-3][0-5]|[0-5]?[0-9]?[0-9]?[0-9]?[0-9]?)$,
    // or int with range: 0..65535, or enumeration PortNumRange.
    DestinationPort interface{}

    // List of TCP flags to match. The type is slice of [u'TCPACK', u'TCPSYN',
    // u'TCPECE', u'TCPFIN', u'TCPRST', u'TCPURG', u'TCPPSH', u'TCPCWR'].
    TcpFlags []interface{}
}

func (state *Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "transport"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/transport/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("source-port", types.YLeaf{"SourcePort", state.SourcePort})
    state.EntityData.Leafs.Append("destination-port", types.YLeaf{"DestinationPort", state.DestinationPort})
    state.EntityData.Leafs.Append("tcp-flags", types.YLeaf{"TcpFlags", state.TcpFlags})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface
// Input interface container
type Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config data.
    Config Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_Config

    // State information.
    State Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_State

    // Reference to an interface or subinterface.
    InterfaceRef Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_InterfaceRef
}

func (inputInterface *Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface) GetEntityData() *types.CommonEntityData {
    inputInterface.EntityData.YFilter = inputInterface.YFilter
    inputInterface.EntityData.YangName = "input-interface"
    inputInterface.EntityData.BundleName = "openconfig"
    inputInterface.EntityData.ParentYangName = "acl-entry"
    inputInterface.EntityData.SegmentPath = "input-interface"
    inputInterface.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/" + inputInterface.EntityData.SegmentPath
    inputInterface.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    inputInterface.EntityData.NamespaceTable = openconfig.GetNamespaces()
    inputInterface.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    inputInterface.EntityData.Children = types.NewOrderedMap()
    inputInterface.EntityData.Children.Append("config", types.YChild{"Config", &inputInterface.Config})
    inputInterface.EntityData.Children.Append("state", types.YChild{"State", &inputInterface.State})
    inputInterface.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &inputInterface.InterfaceRef})
    inputInterface.EntityData.Leafs = types.NewOrderedMap()

    inputInterface.EntityData.YListKeys = []string {}

    return &(inputInterface.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_Config
// Config data
type Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (config *Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "input-interface"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_State
// State information
type Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (state *Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "input-interface"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_InterfaceRef
// Reference to an interface or subinterface
type Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_InterfaceRef_Config

    // Operational state for interface-ref.
    State Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_InterfaceRef_State
}

func (interfaceRef *Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "openconfig"
    interfaceRef.EntityData.ParentYangName = "input-interface"
    interfaceRef.EntityData.SegmentPath = "interface-ref"
    interfaceRef.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/" + interfaceRef.EntityData.SegmentPath
    interfaceRef.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceRef.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceRef.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceRef.EntityData.Children = types.NewOrderedMap()
    interfaceRef.EntityData.Children.Append("config", types.YChild{"Config", &interfaceRef.Config})
    interfaceRef.EntityData.Children.Append("state", types.YChild{"State", &interfaceRef.State})
    interfaceRef.EntityData.Leafs = types.NewOrderedMap()

    interfaceRef.EntityData.YListKeys = []string {}

    return &(interfaceRef.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_InterfaceRef_Config
// Configured reference to interface / subinterface
type Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_InterfaceRef_Config struct {
    EntityData types.CommonEntityData
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

func (config *Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_InterfaceRef_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface-ref"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", config.Interface})
    config.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", config.Subinterface})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_InterfaceRef_State
// Operational state for interface-ref
type Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_InterfaceRef_State struct {
    EntityData types.CommonEntityData
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

func (state *Acl_AclSets_AclSet_AclEntries_AclEntry_InputInterface_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface-ref"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", state.Interface})
    state.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", state.Subinterface})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_Actions
// Enclosing container for list of ACL actions associated
// with an entry
type Acl_AclSets_AclSet_AclEntries_AclEntry_Actions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config data for ACL actions.
    Config Acl_AclSets_AclSet_AclEntries_AclEntry_Actions_Config

    // State information for ACL actions.
    State Acl_AclSets_AclSet_AclEntries_AclEntry_Actions_State
}

func (actions *Acl_AclSets_AclSet_AclEntries_AclEntry_Actions) GetEntityData() *types.CommonEntityData {
    actions.EntityData.YFilter = actions.YFilter
    actions.EntityData.YangName = "actions"
    actions.EntityData.BundleName = "openconfig"
    actions.EntityData.ParentYangName = "acl-entry"
    actions.EntityData.SegmentPath = "actions"
    actions.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/" + actions.EntityData.SegmentPath
    actions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    actions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    actions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    actions.EntityData.Children = types.NewOrderedMap()
    actions.EntityData.Children.Append("config", types.YChild{"Config", &actions.Config})
    actions.EntityData.Children.Append("state", types.YChild{"State", &actions.State})
    actions.EntityData.Leafs = types.NewOrderedMap()

    actions.EntityData.YListKeys = []string {}

    return &(actions.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_Actions_Config
// Config data for ACL actions
type Acl_AclSets_AclSet_AclEntries_AclEntry_Actions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifies the forwarding action.  One forwarding action must be specified
    // for each ACL entry. The type is one of the following: DROPACCEPTREJECT.
    // This attribute is mandatory.
    ForwardingAction interface{}

    // Specifies the log action and destination for matched packets.  The default
    // is not to log the packet. The type is one of the following:
    // LOGNONELOGSYSLOG. The default value is LOG_NONE.
    LogAction interface{}
}

func (config *Acl_AclSets_AclSet_AclEntries_AclEntry_Actions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "actions"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/actions/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("forwarding-action", types.YLeaf{"ForwardingAction", config.ForwardingAction})
    config.EntityData.Leafs.Append("log-action", types.YLeaf{"LogAction", config.LogAction})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_AclSets_AclSet_AclEntries_AclEntry_Actions_State
// State information for ACL actions
type Acl_AclSets_AclSet_AclEntries_AclEntry_Actions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifies the forwarding action.  One forwarding action must be specified
    // for each ACL entry. The type is one of the following: DROPACCEPTREJECT.
    // This attribute is mandatory.
    ForwardingAction interface{}

    // Specifies the log action and destination for matched packets.  The default
    // is not to log the packet. The type is one of the following:
    // LOGNONELOGSYSLOG. The default value is LOG_NONE.
    LogAction interface{}
}

func (state *Acl_AclSets_AclSet_AclEntries_AclEntry_Actions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "actions"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry/actions/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("forwarding-action", types.YLeaf{"ForwardingAction", state.ForwardingAction})
    state.EntityData.Leafs.Append("log-action", types.YLeaf{"LogAction", state.LogAction})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_Interfaces
// Enclosing container for the list of interfaces on which
// ACLs are set
type Acl_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of interfaces on which ACLs are set. The type is slice of
    // Acl_Interfaces_Interface.
    Interface []*Acl_Interfaces_Interface
}

func (interfaces *Acl_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "openconfig"
    interfaces.EntityData.ParentYangName = "acl"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "openconfig-acl:acl/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaces.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Acl_Interfaces_Interface
// List of interfaces on which ACLs are set
type Acl_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the interface id list key. The type
    // is string. Refers to acl.Acl_Interfaces_Interface_Config_Id
    Id interface{}

    // Configuration for ACL per-interface data.
    Config Acl_Interfaces_Interface_Config

    // Operational state for ACL per-interface data.
    State Acl_Interfaces_Interface_State

    // Reference to an interface or subinterface.
    InterfaceRef Acl_Interfaces_Interface_InterfaceRef

    // Enclosing container the list of ingress ACLs on the interface.
    IngressAclSets Acl_Interfaces_Interface_IngressAclSets

    // Enclosing container the list of egress ACLs on the interface.
    EgressAclSets Acl_Interfaces_Interface_EgressAclSets
}

func (self *Acl_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "openconfig"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Id, "id")
    self.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    self.EntityData.NamespaceTable = openconfig.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("config", types.YChild{"Config", &self.Config})
    self.EntityData.Children.Append("state", types.YChild{"State", &self.State})
    self.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &self.InterfaceRef})
    self.EntityData.Children.Append("ingress-acl-sets", types.YChild{"IngressAclSets", &self.IngressAclSets})
    self.EntityData.Children.Append("egress-acl-sets", types.YChild{"EgressAclSets", &self.EgressAclSets})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("id", types.YLeaf{"Id", self.Id})

    self.EntityData.YListKeys = []string {"Id"}

    return &(self.EntityData)
}

// Acl_Interfaces_Interface_Config
// Configuration for ACL per-interface data
type Acl_Interfaces_Interface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // User-defined identifier for the interface -- a common convention could be
    // '<if name>.<subif index>'. The type is string.
    Id interface{}
}

func (config *Acl_Interfaces_Interface_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("id", types.YLeaf{"Id", config.Id})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_Interfaces_Interface_State
// Operational state for ACL per-interface data
type Acl_Interfaces_Interface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // User-defined identifier for the interface -- a common convention could be
    // '<if name>.<subif index>'. The type is string.
    Id interface{}
}

func (state *Acl_Interfaces_Interface_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("id", types.YLeaf{"Id", state.Id})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_Interfaces_Interface_InterfaceRef
// Reference to an interface or subinterface
type Acl_Interfaces_Interface_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Acl_Interfaces_Interface_InterfaceRef_Config

    // Operational state for interface-ref.
    State Acl_Interfaces_Interface_InterfaceRef_State
}

func (interfaceRef *Acl_Interfaces_Interface_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "openconfig"
    interfaceRef.EntityData.ParentYangName = "interface"
    interfaceRef.EntityData.SegmentPath = "interface-ref"
    interfaceRef.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/" + interfaceRef.EntityData.SegmentPath
    interfaceRef.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceRef.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceRef.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceRef.EntityData.Children = types.NewOrderedMap()
    interfaceRef.EntityData.Children.Append("config", types.YChild{"Config", &interfaceRef.Config})
    interfaceRef.EntityData.Children.Append("state", types.YChild{"State", &interfaceRef.State})
    interfaceRef.EntityData.Leafs = types.NewOrderedMap()

    interfaceRef.EntityData.YListKeys = []string {}

    return &(interfaceRef.EntityData)
}

// Acl_Interfaces_Interface_InterfaceRef_Config
// Configured reference to interface / subinterface
type Acl_Interfaces_Interface_InterfaceRef_Config struct {
    EntityData types.CommonEntityData
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

func (config *Acl_Interfaces_Interface_InterfaceRef_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface-ref"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/interface-ref/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", config.Interface})
    config.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", config.Subinterface})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_Interfaces_Interface_InterfaceRef_State
// Operational state for interface-ref
type Acl_Interfaces_Interface_InterfaceRef_State struct {
    EntityData types.CommonEntityData
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

func (state *Acl_Interfaces_Interface_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface-ref"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/interface-ref/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", state.Interface})
    state.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", state.Subinterface})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_Interfaces_Interface_IngressAclSets
// Enclosing container the list of ingress ACLs on the
// interface
type Acl_Interfaces_Interface_IngressAclSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of ingress ACLs on the interface. The type is slice of
    // Acl_Interfaces_Interface_IngressAclSets_IngressAclSet.
    IngressAclSet []*Acl_Interfaces_Interface_IngressAclSets_IngressAclSet
}

func (ingressAclSets *Acl_Interfaces_Interface_IngressAclSets) GetEntityData() *types.CommonEntityData {
    ingressAclSets.EntityData.YFilter = ingressAclSets.YFilter
    ingressAclSets.EntityData.YangName = "ingress-acl-sets"
    ingressAclSets.EntityData.BundleName = "openconfig"
    ingressAclSets.EntityData.ParentYangName = "interface"
    ingressAclSets.EntityData.SegmentPath = "ingress-acl-sets"
    ingressAclSets.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/" + ingressAclSets.EntityData.SegmentPath
    ingressAclSets.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ingressAclSets.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ingressAclSets.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ingressAclSets.EntityData.Children = types.NewOrderedMap()
    ingressAclSets.EntityData.Children.Append("ingress-acl-set", types.YChild{"IngressAclSet", nil})
    for i := range ingressAclSets.IngressAclSet {
        ingressAclSets.EntityData.Children.Append(types.GetSegmentPath(ingressAclSets.IngressAclSet[i]), types.YChild{"IngressAclSet", ingressAclSets.IngressAclSet[i]})
    }
    ingressAclSets.EntityData.Leafs = types.NewOrderedMap()

    ingressAclSets.EntityData.YListKeys = []string {}

    return &(ingressAclSets.EntityData)
}

// Acl_Interfaces_Interface_IngressAclSets_IngressAclSet
// List of ingress ACLs on the interface
type Acl_Interfaces_Interface_IngressAclSets_IngressAclSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to set name list key. The type is
    // string. Refers to
    // acl.Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_Config_SetName
    SetName interface{}

    // This attribute is a key. Reference to type list key. The type is one of the
    // following: ACLL2ACLIPV6ACLIPV4ACLMIXED.
    Type interface{}

    // Configuration data .
    Config Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_Config

    // Operational state data for interface ingress ACLs.
    State Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_State

    // Enclosing container for list of references to ACLs.
    AclEntries Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries
}

func (ingressAclSet *Acl_Interfaces_Interface_IngressAclSets_IngressAclSet) GetEntityData() *types.CommonEntityData {
    ingressAclSet.EntityData.YFilter = ingressAclSet.YFilter
    ingressAclSet.EntityData.YangName = "ingress-acl-set"
    ingressAclSet.EntityData.BundleName = "openconfig"
    ingressAclSet.EntityData.ParentYangName = "ingress-acl-sets"
    ingressAclSet.EntityData.SegmentPath = "ingress-acl-set" + types.AddKeyToken(ingressAclSet.SetName, "set-name") + types.AddKeyToken(ingressAclSet.Type, "type")
    ingressAclSet.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/ingress-acl-sets/" + ingressAclSet.EntityData.SegmentPath
    ingressAclSet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ingressAclSet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ingressAclSet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ingressAclSet.EntityData.Children = types.NewOrderedMap()
    ingressAclSet.EntityData.Children.Append("config", types.YChild{"Config", &ingressAclSet.Config})
    ingressAclSet.EntityData.Children.Append("state", types.YChild{"State", &ingressAclSet.State})
    ingressAclSet.EntityData.Children.Append("acl-entries", types.YChild{"AclEntries", &ingressAclSet.AclEntries})
    ingressAclSet.EntityData.Leafs = types.NewOrderedMap()
    ingressAclSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", ingressAclSet.SetName})
    ingressAclSet.EntityData.Leafs.Append("type", types.YLeaf{"Type", ingressAclSet.Type})

    ingressAclSet.EntityData.YListKeys = []string {"SetName", "Type"}

    return &(ingressAclSet.EntityData)
}

// Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_Config
// Configuration data 
type Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the ACL set name applied on ingress. The type is string.
    // Refers to acl.Acl_AclSets_AclSet_Config_Name
    SetName interface{}

    // Reference to the ACL set type applied on ingress. The type is one of the
    // following: ACLL2ACLIPV6ACLIPV4ACLMIXED.
    Type interface{}
}

func (config *Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ingress-acl-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", config.SetName})
    config.EntityData.Leafs.Append("type", types.YLeaf{"Type", config.Type})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_State
// Operational state data for interface ingress ACLs
type Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the ACL set name applied on ingress. The type is string.
    // Refers to acl.Acl_AclSets_AclSet_Config_Name
    SetName interface{}

    // Reference to the ACL set type applied on ingress. The type is one of the
    // following: ACLL2ACLIPV6ACLIPV4ACLMIXED.
    Type interface{}
}

func (state *Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ingress-acl-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", state.SetName})
    state.EntityData.Leafs.Append("type", types.YLeaf{"Type", state.Type})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries
// Enclosing container for list of references to ACLs
type Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of ACL entries assigned to an interface. The type is slice of
    // Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries_AclEntry.
    AclEntry []*Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries_AclEntry
}

func (aclEntries *Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries) GetEntityData() *types.CommonEntityData {
    aclEntries.EntityData.YFilter = aclEntries.YFilter
    aclEntries.EntityData.YangName = "acl-entries"
    aclEntries.EntityData.BundleName = "openconfig"
    aclEntries.EntityData.ParentYangName = "ingress-acl-set"
    aclEntries.EntityData.SegmentPath = "acl-entries"
    aclEntries.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/" + aclEntries.EntityData.SegmentPath
    aclEntries.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aclEntries.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aclEntries.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aclEntries.EntityData.Children = types.NewOrderedMap()
    aclEntries.EntityData.Children.Append("acl-entry", types.YChild{"AclEntry", nil})
    for i := range aclEntries.AclEntry {
        aclEntries.EntityData.Children.Append(types.GetSegmentPath(aclEntries.AclEntry[i]), types.YChild{"AclEntry", aclEntries.AclEntry[i]})
    }
    aclEntries.EntityData.Leafs = types.NewOrderedMap()

    aclEntries.EntityData.YListKeys = []string {}

    return &(aclEntries.EntityData)
}

// Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries_AclEntry
// List of ACL entries assigned to an interface
type Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries_AclEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to per-interface acl entry key. The type
    // is string with range: 0..4294967295. Refers to
    // acl.Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries_AclEntry_State_SequenceId
    SequenceId interface{}

    // Operational state data for per-interface ACL entries.
    State Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries_AclEntry_State
}

func (aclEntry *Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries_AclEntry) GetEntityData() *types.CommonEntityData {
    aclEntry.EntityData.YFilter = aclEntry.YFilter
    aclEntry.EntityData.YangName = "acl-entry"
    aclEntry.EntityData.BundleName = "openconfig"
    aclEntry.EntityData.ParentYangName = "acl-entries"
    aclEntry.EntityData.SegmentPath = "acl-entry" + types.AddKeyToken(aclEntry.SequenceId, "sequence-id")
    aclEntry.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/" + aclEntry.EntityData.SegmentPath
    aclEntry.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aclEntry.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aclEntry.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aclEntry.EntityData.Children = types.NewOrderedMap()
    aclEntry.EntityData.Children.Append("state", types.YChild{"State", &aclEntry.State})
    aclEntry.EntityData.Leafs = types.NewOrderedMap()
    aclEntry.EntityData.Leafs.Append("sequence-id", types.YLeaf{"SequenceId", aclEntry.SequenceId})

    aclEntry.EntityData.YListKeys = []string {"SequenceId"}

    return &(aclEntry.EntityData)
}

// Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries_AclEntry_State
// Operational state data for per-interface ACL entries
type Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries_AclEntry_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to an entry in the ACL set applied to an interface. The type is
    // string with range: 0..4294967295. Refers to
    // acl.Acl_AclSets_AclSet_AclEntries_AclEntry_SequenceId
    SequenceId interface{}

    // Count of the number of packets matching the current ACL entry.  An
    // implementation should provide this counter on a per-interface per-ACL-entry
    // if possible.  If an implementation only supports ACL counters per entry
    // (i.e., not broken out per interface), then the value should be equal to the
    // aggregate count across all interfaces.  An implementation that provides
    // counters per entry per interface is not required to also provide an
    // aggregate count, e.g., per entry -- the user is expected to be able
    // implement the required aggregation if such a count is needed. The type is
    // interface{} with range: 0..18446744073709551615.
    MatchedPackets interface{}

    // Count of the number of octets (bytes) matching the current ACL entry.  An
    // implementation should provide this counter on a per-interface per-ACL-entry
    // if possible.  If an implementation only supports ACL counters per entry
    // (i.e., not broken out per interface), then the value should be equal to the
    // aggregate count across all interfaces.  An implementation that provides
    // counters per entry per interface is not required to also provide an
    // aggregate count, e.g., per entry -- the user is expected to be able
    // implement the required aggregation if such a count is needed. The type is
    // interface{} with range: 0..18446744073709551615.
    MatchedOctets interface{}
}

func (state *Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries_AclEntry_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "acl-entry"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/acl-entry/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("sequence-id", types.YLeaf{"SequenceId", state.SequenceId})
    state.EntityData.Leafs.Append("matched-packets", types.YLeaf{"MatchedPackets", state.MatchedPackets})
    state.EntityData.Leafs.Append("matched-octets", types.YLeaf{"MatchedOctets", state.MatchedOctets})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_Interfaces_Interface_EgressAclSets
// Enclosing container the list of egress ACLs on the
// interface
type Acl_Interfaces_Interface_EgressAclSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of egress ACLs on the interface. The type is slice of
    // Acl_Interfaces_Interface_EgressAclSets_EgressAclSet.
    EgressAclSet []*Acl_Interfaces_Interface_EgressAclSets_EgressAclSet
}

func (egressAclSets *Acl_Interfaces_Interface_EgressAclSets) GetEntityData() *types.CommonEntityData {
    egressAclSets.EntityData.YFilter = egressAclSets.YFilter
    egressAclSets.EntityData.YangName = "egress-acl-sets"
    egressAclSets.EntityData.BundleName = "openconfig"
    egressAclSets.EntityData.ParentYangName = "interface"
    egressAclSets.EntityData.SegmentPath = "egress-acl-sets"
    egressAclSets.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/" + egressAclSets.EntityData.SegmentPath
    egressAclSets.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    egressAclSets.EntityData.NamespaceTable = openconfig.GetNamespaces()
    egressAclSets.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    egressAclSets.EntityData.Children = types.NewOrderedMap()
    egressAclSets.EntityData.Children.Append("egress-acl-set", types.YChild{"EgressAclSet", nil})
    for i := range egressAclSets.EgressAclSet {
        egressAclSets.EntityData.Children.Append(types.GetSegmentPath(egressAclSets.EgressAclSet[i]), types.YChild{"EgressAclSet", egressAclSets.EgressAclSet[i]})
    }
    egressAclSets.EntityData.Leafs = types.NewOrderedMap()

    egressAclSets.EntityData.YListKeys = []string {}

    return &(egressAclSets.EntityData)
}

// Acl_Interfaces_Interface_EgressAclSets_EgressAclSet
// List of egress ACLs on the interface
type Acl_Interfaces_Interface_EgressAclSets_EgressAclSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to set name list key. The type is
    // string. Refers to
    // acl.Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_Config_SetName
    SetName interface{}

    // This attribute is a key. Reference to type list key. The type is one of the
    // following: ACLL2ACLIPV6ACLIPV4ACLMIXED.
    Type interface{}

    // Configuration data .
    Config Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_Config

    // Operational state data for interface egress ACLs.
    State Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_State

    // Enclosing container for list of references to ACLs.
    AclEntries Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries
}

func (egressAclSet *Acl_Interfaces_Interface_EgressAclSets_EgressAclSet) GetEntityData() *types.CommonEntityData {
    egressAclSet.EntityData.YFilter = egressAclSet.YFilter
    egressAclSet.EntityData.YangName = "egress-acl-set"
    egressAclSet.EntityData.BundleName = "openconfig"
    egressAclSet.EntityData.ParentYangName = "egress-acl-sets"
    egressAclSet.EntityData.SegmentPath = "egress-acl-set" + types.AddKeyToken(egressAclSet.SetName, "set-name") + types.AddKeyToken(egressAclSet.Type, "type")
    egressAclSet.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/egress-acl-sets/" + egressAclSet.EntityData.SegmentPath
    egressAclSet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    egressAclSet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    egressAclSet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    egressAclSet.EntityData.Children = types.NewOrderedMap()
    egressAclSet.EntityData.Children.Append("config", types.YChild{"Config", &egressAclSet.Config})
    egressAclSet.EntityData.Children.Append("state", types.YChild{"State", &egressAclSet.State})
    egressAclSet.EntityData.Children.Append("acl-entries", types.YChild{"AclEntries", &egressAclSet.AclEntries})
    egressAclSet.EntityData.Leafs = types.NewOrderedMap()
    egressAclSet.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", egressAclSet.SetName})
    egressAclSet.EntityData.Leafs.Append("type", types.YLeaf{"Type", egressAclSet.Type})

    egressAclSet.EntityData.YListKeys = []string {"SetName", "Type"}

    return &(egressAclSet.EntityData)
}

// Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_Config
// Configuration data 
type Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the ACL set name applied on egress. The type is string. Refers
    // to acl.Acl_AclSets_AclSet_Config_Name
    SetName interface{}

    // Reference to the ACL set type applied on egress. The type is one of the
    // following: ACLL2ACLIPV6ACLIPV4ACLMIXED.
    Type interface{}
}

func (config *Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "egress-acl-set"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/egress-acl-sets/egress-acl-set/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", config.SetName})
    config.EntityData.Leafs.Append("type", types.YLeaf{"Type", config.Type})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_State
// Operational state data for interface egress ACLs
type Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the ACL set name applied on egress. The type is string. Refers
    // to acl.Acl_AclSets_AclSet_Config_Name
    SetName interface{}

    // Reference to the ACL set type applied on egress. The type is one of the
    // following: ACLL2ACLIPV6ACLIPV4ACLMIXED.
    Type interface{}
}

func (state *Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "egress-acl-set"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/egress-acl-sets/egress-acl-set/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("set-name", types.YLeaf{"SetName", state.SetName})
    state.EntityData.Leafs.Append("type", types.YLeaf{"Type", state.Type})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries
// Enclosing container for list of references to ACLs
type Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of ACL entries assigned to an interface. The type is slice of
    // Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries_AclEntry.
    AclEntry []*Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries_AclEntry
}

func (aclEntries *Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries) GetEntityData() *types.CommonEntityData {
    aclEntries.EntityData.YFilter = aclEntries.YFilter
    aclEntries.EntityData.YangName = "acl-entries"
    aclEntries.EntityData.BundleName = "openconfig"
    aclEntries.EntityData.ParentYangName = "egress-acl-set"
    aclEntries.EntityData.SegmentPath = "acl-entries"
    aclEntries.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/egress-acl-sets/egress-acl-set/" + aclEntries.EntityData.SegmentPath
    aclEntries.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aclEntries.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aclEntries.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aclEntries.EntityData.Children = types.NewOrderedMap()
    aclEntries.EntityData.Children.Append("acl-entry", types.YChild{"AclEntry", nil})
    for i := range aclEntries.AclEntry {
        aclEntries.EntityData.Children.Append(types.GetSegmentPath(aclEntries.AclEntry[i]), types.YChild{"AclEntry", aclEntries.AclEntry[i]})
    }
    aclEntries.EntityData.Leafs = types.NewOrderedMap()

    aclEntries.EntityData.YListKeys = []string {}

    return &(aclEntries.EntityData)
}

// Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries_AclEntry
// List of ACL entries assigned to an interface
type Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries_AclEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to per-interface acl entry key. The type
    // is string with range: 0..4294967295. Refers to
    // acl.Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries_AclEntry_State_SequenceId
    SequenceId interface{}

    // Operational state data for per-interface ACL entries.
    State Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries_AclEntry_State
}

func (aclEntry *Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries_AclEntry) GetEntityData() *types.CommonEntityData {
    aclEntry.EntityData.YFilter = aclEntry.YFilter
    aclEntry.EntityData.YangName = "acl-entry"
    aclEntry.EntityData.BundleName = "openconfig"
    aclEntry.EntityData.ParentYangName = "acl-entries"
    aclEntry.EntityData.SegmentPath = "acl-entry" + types.AddKeyToken(aclEntry.SequenceId, "sequence-id")
    aclEntry.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/" + aclEntry.EntityData.SegmentPath
    aclEntry.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aclEntry.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aclEntry.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aclEntry.EntityData.Children = types.NewOrderedMap()
    aclEntry.EntityData.Children.Append("state", types.YChild{"State", &aclEntry.State})
    aclEntry.EntityData.Leafs = types.NewOrderedMap()
    aclEntry.EntityData.Leafs.Append("sequence-id", types.YLeaf{"SequenceId", aclEntry.SequenceId})

    aclEntry.EntityData.YListKeys = []string {"SequenceId"}

    return &(aclEntry.EntityData)
}

// Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries_AclEntry_State
// Operational state data for per-interface ACL entries
type Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries_AclEntry_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to an entry in the ACL set applied to an interface. The type is
    // string with range: 0..4294967295. Refers to
    // acl.Acl_AclSets_AclSet_AclEntries_AclEntry_SequenceId
    SequenceId interface{}

    // Count of the number of packets matching the current ACL entry.  An
    // implementation should provide this counter on a per-interface per-ACL-entry
    // if possible.  If an implementation only supports ACL counters per entry
    // (i.e., not broken out per interface), then the value should be equal to the
    // aggregate count across all interfaces.  An implementation that provides
    // counters per entry per interface is not required to also provide an
    // aggregate count, e.g., per entry -- the user is expected to be able
    // implement the required aggregation if such a count is needed. The type is
    // interface{} with range: 0..18446744073709551615.
    MatchedPackets interface{}

    // Count of the number of octets (bytes) matching the current ACL entry.  An
    // implementation should provide this counter on a per-interface per-ACL-entry
    // if possible.  If an implementation only supports ACL counters per entry
    // (i.e., not broken out per interface), then the value should be equal to the
    // aggregate count across all interfaces.  An implementation that provides
    // counters per entry per interface is not required to also provide an
    // aggregate count, e.g., per entry -- the user is expected to be able
    // implement the required aggregation if such a count is needed. The type is
    // interface{} with range: 0..18446744073709551615.
    MatchedOctets interface{}
}

func (state *Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries_AclEntry_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "acl-entry"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-acl:acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("sequence-id", types.YLeaf{"SequenceId", state.SequenceId})
    state.EntityData.Leafs.Append("matched-packets", types.YLeaf{"MatchedPackets", state.MatchedPackets})
    state.EntityData.Leafs.Append("matched-octets", types.YLeaf{"MatchedOctets", state.MatchedOctets})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

