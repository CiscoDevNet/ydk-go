// This module describes configuration and operational state
// data for Link Aggregation Control Protocol (LACP) for
// managing aggregate interfaces.   It works in conjunction with
// the OpenConfig interfaces and aggregate interfaces models.
package lacp

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lacp"))
    ydk.RegisterEntity("{http://openconfig.net/yang/lacp lacp}", reflect.TypeOf(Lacp{}))
    ydk.RegisterEntity("openconfig-lacp:lacp", reflect.TypeOf(Lacp{}))
}

// LacpActivityType represents interface in the aggregate
type LacpActivityType string

const (
    // Interface is an active member, i.e., will detect and
    // maintain aggregates
    LacpActivityType_ACTIVE LacpActivityType = "ACTIVE"

    // Interface is a passive member, i.e., it participates
    // with an active partner
    LacpActivityType_PASSIVE LacpActivityType = "PASSIVE"
)

// LacpTimeoutType represents Type of timeout used, short or long, by LACP participants
type LacpTimeoutType string

const (
    // Participant wishes to use long timeouts to detect
    // status of the aggregate, i.e., will expect less frequent
    // transmissions. Long timeout is 90 seconds.
    LacpTimeoutType_LONG LacpTimeoutType = "LONG"

    // Participant wishes to use short timeouts, i.e., expects
    // frequent transmissions to aggressively detect status
    // changes. Short timeout is 3 seconds.
    LacpTimeoutType_SHORT LacpTimeoutType = "SHORT"
)

// LacpSynchronizationType represents Indicates LACP synchronization state of participant
type LacpSynchronizationType string

const (
    // Participant is in sync with the system id and key
    // transmitted
    LacpSynchronizationType_IN_SYNC LacpSynchronizationType = "IN_SYNC"

    // Participant is not in sync with the system id and key
    // transmitted
    LacpSynchronizationType_OUT_SYNC LacpSynchronizationType = "OUT_SYNC"
)

// LacpPeriodType represents LACP messages
type LacpPeriodType string

const (
    // Send LACP packets every second
    LacpPeriodType_FAST LacpPeriodType = "FAST"

    // Send LACP packets every 30 seconds
    LacpPeriodType_SLOW LacpPeriodType = "SLOW"
)

// Lacp
// Configuration and operational state data for LACP protocol
// operation on the aggregate interface
type Lacp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for LACP.
    Config Lacp_Config

    // Operational state data for LACP.
    State Lacp_State

    // Enclosing container for the list of LACP-enabled interfaces.
    Interfaces Lacp_Interfaces
}

func (lacp *Lacp) GetEntityData() *types.CommonEntityData {
    lacp.EntityData.YFilter = lacp.YFilter
    lacp.EntityData.YangName = "lacp"
    lacp.EntityData.BundleName = "openconfig"
    lacp.EntityData.ParentYangName = "openconfig-lacp"
    lacp.EntityData.SegmentPath = "openconfig-lacp:lacp"
    lacp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    lacp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    lacp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    lacp.EntityData.Children = make(map[string]types.YChild)
    lacp.EntityData.Children["config"] = types.YChild{"Config", &lacp.Config}
    lacp.EntityData.Children["state"] = types.YChild{"State", &lacp.State}
    lacp.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &lacp.Interfaces}
    lacp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lacp.EntityData)
}

// Lacp_Config
// Configuration data for LACP
type Lacp_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sytem priority used by the node on this LAG interface. Lower value is
    // higher priority for determining which node is the controlling system. The
    // type is interface{} with range: 0..65535.
    SystemPriority interface{}
}

func (config *Lacp_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "lacp"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["system-priority"] = types.YLeaf{"SystemPriority", config.SystemPriority}
    return &(config.EntityData)
}

// Lacp_State
// Operational state data for LACP
type Lacp_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sytem priority used by the node on this LAG interface. Lower value is
    // higher priority for determining which node is the controlling system. The
    // type is interface{} with range: 0..65535.
    SystemPriority interface{}
}

func (state *Lacp_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "lacp"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["system-priority"] = types.YLeaf{"SystemPriority", state.SystemPriority}
    return &(state.EntityData)
}

// Lacp_Interfaces
// Enclosing container for the list of LACP-enabled
// interfaces
type Lacp_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of aggregate interfaces managed by LACP. The type is slice of
    // Lacp_Interfaces_Interface_.
    Interface_ []Lacp_Interfaces_Interface
}

func (interfaces *Lacp_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "openconfig"
    interfaces.EntityData.ParentYangName = "lacp"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaces.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Lacp_Interfaces_Interface
// List of aggregate interfaces managed by LACP
type Lacp_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the list key. The type is string.
    // Refers to lacp.Lacp_Interfaces_Interface_Config_Name
    Name interface{}

    // Configuration data for each LACP aggregate interface.
    Config Lacp_Interfaces_Interface_Config

    // Operational state data for each LACP aggregate interface.
    State Lacp_Interfaces_Interface_State

    // Enclosing container for the list of members interfaces of the aggregate.
    // This list is considered operational state only so is labeled config false
    // and has no config container.
    Members Lacp_Interfaces_Interface_Members
}

func (self *Lacp_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "openconfig"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
    self.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    self.EntityData.NamespaceTable = openconfig.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["config"] = types.YChild{"Config", &self.Config}
    self.EntityData.Children["state"] = types.YChild{"State", &self.State}
    self.EntityData.Children["members"] = types.YChild{"Members", &self.Members}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["name"] = types.YLeaf{"Name", self.Name}
    return &(self.EntityData)
}

// Lacp_Interfaces_Interface_Config
// Configuration data for each LACP aggregate interface
type Lacp_Interfaces_Interface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the interface on which LACP should be configured.   The type
    // of the target interface must be ieee8023adLag. The type is string. Refers
    // to interfaces.Interfaces_Interface_Name
    Name interface{}

    // Set the period between LACP messages -- uses the lacp-period-type
    // enumeration. The type is LacpPeriodType. The default value is SLOW.
    Interval interface{}

    // ACTIVE is to initiate the transmission of LACP packets. PASSIVE is to wait
    // for peer to initiate the transmission of LACP packets. The type is
    // LacpActivityType. The default value is ACTIVE.
    LacpMode interface{}

    // The MAC address portion of the node's System ID. This is combined with the
    // system priority to construct the 8-octet system-id. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SystemIdMac interface{}

    // Sytem priority used by the node on this LAG interface. Lower value is
    // higher priority for determining which node is the controlling system. The
    // type is interface{} with range: 0..65535.
    SystemPriority interface{}
}

func (config *Lacp_Interfaces_Interface_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["name"] = types.YLeaf{"Name", config.Name}
    config.EntityData.Leafs["interval"] = types.YLeaf{"Interval", config.Interval}
    config.EntityData.Leafs["lacp-mode"] = types.YLeaf{"LacpMode", config.LacpMode}
    config.EntityData.Leafs["system-id-mac"] = types.YLeaf{"SystemIdMac", config.SystemIdMac}
    config.EntityData.Leafs["system-priority"] = types.YLeaf{"SystemPriority", config.SystemPriority}
    return &(config.EntityData)
}

// Lacp_Interfaces_Interface_State
// Operational state data for each LACP aggregate
// interface
type Lacp_Interfaces_Interface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the interface on which LACP should be configured.   The type
    // of the target interface must be ieee8023adLag. The type is string. Refers
    // to interfaces.Interfaces_Interface_Name
    Name interface{}

    // Set the period between LACP messages -- uses the lacp-period-type
    // enumeration. The type is LacpPeriodType. The default value is SLOW.
    Interval interface{}

    // ACTIVE is to initiate the transmission of LACP packets. PASSIVE is to wait
    // for peer to initiate the transmission of LACP packets. The type is
    // LacpActivityType. The default value is ACTIVE.
    LacpMode interface{}

    // The MAC address portion of the node's System ID. This is combined with the
    // system priority to construct the 8-octet system-id. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SystemIdMac interface{}

    // Sytem priority used by the node on this LAG interface. Lower value is
    // higher priority for determining which node is the controlling system. The
    // type is interface{} with range: 0..65535.
    SystemPriority interface{}
}

func (state *Lacp_Interfaces_Interface_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["name"] = types.YLeaf{"Name", state.Name}
    state.EntityData.Leafs["interval"] = types.YLeaf{"Interval", state.Interval}
    state.EntityData.Leafs["lacp-mode"] = types.YLeaf{"LacpMode", state.LacpMode}
    state.EntityData.Leafs["system-id-mac"] = types.YLeaf{"SystemIdMac", state.SystemIdMac}
    state.EntityData.Leafs["system-priority"] = types.YLeaf{"SystemPriority", state.SystemPriority}
    return &(state.EntityData)
}

// Lacp_Interfaces_Interface_Members
// Enclosing container for the list of members interfaces of
// the aggregate. This list is considered operational state
// only so is labeled config false and has no config container
type Lacp_Interfaces_Interface_Members struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of member interfaces and their associated status for a LACP-controlled
    // aggregate interface.  Member list is not configurable here -- each
    // interface indicates items its participation in the LAG. The type is slice
    // of Lacp_Interfaces_Interface_Members_Member.
    Member []Lacp_Interfaces_Interface_Members_Member
}

func (members *Lacp_Interfaces_Interface_Members) GetEntityData() *types.CommonEntityData {
    members.EntityData.YFilter = members.YFilter
    members.EntityData.YangName = "members"
    members.EntityData.BundleName = "openconfig"
    members.EntityData.ParentYangName = "interface"
    members.EntityData.SegmentPath = "members"
    members.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    members.EntityData.NamespaceTable = openconfig.GetNamespaces()
    members.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    members.EntityData.Children = make(map[string]types.YChild)
    members.EntityData.Children["member"] = types.YChild{"Member", nil}
    for i := range members.Member {
        members.EntityData.Children[types.GetSegmentPath(&members.Member[i])] = types.YChild{"Member", &members.Member[i]}
    }
    members.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(members.EntityData)
}

// Lacp_Interfaces_Interface_Members_Member
// List of member interfaces and their associated status for
// a LACP-controlled aggregate interface.  Member list is not
// configurable here -- each interface indicates items
// its participation in the LAG.
type Lacp_Interfaces_Interface_Members_Member struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to aggregate member interface. The type
    // is string. Refers to
    // lacp.Lacp_Interfaces_Interface_Members_Member_State_Interface_
    Interface_ interface{}

    // Operational state data for aggregate members.
    State Lacp_Interfaces_Interface_Members_Member_State
}

func (member *Lacp_Interfaces_Interface_Members_Member) GetEntityData() *types.CommonEntityData {
    member.EntityData.YFilter = member.YFilter
    member.EntityData.YangName = "member"
    member.EntityData.BundleName = "openconfig"
    member.EntityData.ParentYangName = "members"
    member.EntityData.SegmentPath = "member" + "[interface='" + fmt.Sprintf("%v", member.Interface_) + "']"
    member.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    member.EntityData.NamespaceTable = openconfig.GetNamespaces()
    member.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    member.EntityData.Children = make(map[string]types.YChild)
    member.EntityData.Children["state"] = types.YChild{"State", &member.State}
    member.EntityData.Leafs = make(map[string]types.YLeaf)
    member.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", member.Interface_}
    return &(member.EntityData)
}

// Lacp_Interfaces_Interface_Members_Member_State
// Operational state data for aggregate members
type Lacp_Interfaces_Interface_Members_Member_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to interface member of the LACP aggregate. The type is string.
    // Refers to interfaces.Interfaces_Interface_Name
    Interface_ interface{}

    // Indicates participant is active or passive. The type is LacpActivityType.
    Activity interface{}

    // The timeout type (short or long) used by the participant. The type is
    // LacpTimeoutType.
    Timeout interface{}

    // Indicates whether the participant is in-sync or out-of-sync. The type is
    // LacpSynchronizationType.
    Synchronization interface{}

    // A true value indicates that the participant will allow the link to be used
    // as part of the aggregate. A false value indicates the link should be used
    // as an individual link. The type is bool.
    Aggregatable interface{}

    // If true, the participant is collecting incoming frames on the link,
    // otherwise false. The type is bool.
    Collecting interface{}

    // When true, the participant is distributing outgoing frames; when false,
    // distribution is disabled. The type is bool.
    Distributing interface{}

    // MAC address that defines the local system ID for the aggregate interface.
    // The type is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SystemId interface{}

    // Current operational value of the key for the aggregate interface. The type
    // is interface{} with range: 0..65535.
    OperKey interface{}

    // MAC address representing the protocol partner's interface system ID. The
    // type is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    PartnerId interface{}

    // Operational value of the protocol partner's key. The type is interface{}
    // with range: 0..65535.
    PartnerKey interface{}

    // LACP protocol counters.
    Counters Lacp_Interfaces_Interface_Members_Member_State_Counters
}

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "member"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Children["counters"] = types.YChild{"Counters", &state.Counters}
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", state.Interface_}
    state.EntityData.Leafs["activity"] = types.YLeaf{"Activity", state.Activity}
    state.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", state.Timeout}
    state.EntityData.Leafs["synchronization"] = types.YLeaf{"Synchronization", state.Synchronization}
    state.EntityData.Leafs["aggregatable"] = types.YLeaf{"Aggregatable", state.Aggregatable}
    state.EntityData.Leafs["collecting"] = types.YLeaf{"Collecting", state.Collecting}
    state.EntityData.Leafs["distributing"] = types.YLeaf{"Distributing", state.Distributing}
    state.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", state.SystemId}
    state.EntityData.Leafs["oper-key"] = types.YLeaf{"OperKey", state.OperKey}
    state.EntityData.Leafs["partner-id"] = types.YLeaf{"PartnerId", state.PartnerId}
    state.EntityData.Leafs["partner-key"] = types.YLeaf{"PartnerKey", state.PartnerKey}
    return &(state.EntityData)
}

// Lacp_Interfaces_Interface_Members_Member_State_Counters
// LACP protocol counters
type Lacp_Interfaces_Interface_Members_Member_State_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of LACPDUs received. The type is interface{} with range:
    // 0..18446744073709551615.
    LacpInPkts interface{}

    // Number of LACPDUs transmitted. The type is interface{} with range:
    // 0..18446744073709551615.
    LacpOutPkts interface{}

    // Number of LACPDU receive packet errors. The type is interface{} with range:
    // 0..18446744073709551615.
    LacpRxErrors interface{}

    // Number of LACPDU transmit packet errors. The type is interface{} with
    // range: 0..18446744073709551615.
    LacpTxErrors interface{}

    // Number of LACPDU unknown packet errors. The type is interface{} with range:
    // 0..18446744073709551615.
    LacpUnknownErrors interface{}

    // Number of LACPDU illegal packet errors. The type is interface{} with range:
    // 0..18446744073709551615.
    LacpErrors interface{}
}

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "openconfig"
    counters.EntityData.ParentYangName = "state"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    counters.EntityData.NamespaceTable = openconfig.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    counters.EntityData.Children = make(map[string]types.YChild)
    counters.EntityData.Leafs = make(map[string]types.YLeaf)
    counters.EntityData.Leafs["lacp-in-pkts"] = types.YLeaf{"LacpInPkts", counters.LacpInPkts}
    counters.EntityData.Leafs["lacp-out-pkts"] = types.YLeaf{"LacpOutPkts", counters.LacpOutPkts}
    counters.EntityData.Leafs["lacp-rx-errors"] = types.YLeaf{"LacpRxErrors", counters.LacpRxErrors}
    counters.EntityData.Leafs["lacp-tx-errors"] = types.YLeaf{"LacpTxErrors", counters.LacpTxErrors}
    counters.EntityData.Leafs["lacp-unknown-errors"] = types.YLeaf{"LacpUnknownErrors", counters.LacpUnknownErrors}
    counters.EntityData.Leafs["lacp-errors"] = types.YLeaf{"LacpErrors", counters.LacpErrors}
    return &(counters.EntityData)
}

