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

// LacpPeriodType represents LACP messages
type LacpPeriodType string

const (
    // Send LACP packets every second
    LacpPeriodType_FAST LacpPeriodType = "FAST"

    // Send LACP packets every 30 seconds
    LacpPeriodType_SLOW LacpPeriodType = "SLOW"
)

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

// Lacp
// Configuration and operational state data for LACP protocol
// operation on the aggregate interface
type Lacp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for LACP.
    Config Lacp_Config

    // Operational state data for LACP.
    State Lacp_State

    // Enclosing container for the list of LACP-enabled interfaces.
    Interfaces Lacp_Interfaces
}

func (lacp *Lacp) GetFilter() yfilter.YFilter { return lacp.YFilter }

func (lacp *Lacp) SetFilter(yf yfilter.YFilter) { lacp.YFilter = yf }

func (lacp *Lacp) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (lacp *Lacp) GetSegmentPath() string {
    return "openconfig-lacp:lacp"
}

func (lacp *Lacp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &lacp.Config
    }
    if childYangName == "state" {
        return &lacp.State
    }
    if childYangName == "interfaces" {
        return &lacp.Interfaces
    }
    return nil
}

func (lacp *Lacp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &lacp.Config
    children["state"] = &lacp.State
    children["interfaces"] = &lacp.Interfaces
    return children
}

func (lacp *Lacp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lacp *Lacp) GetBundleName() string { return "openconfig" }

func (lacp *Lacp) GetYangName() string { return "lacp" }

func (lacp *Lacp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (lacp *Lacp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (lacp *Lacp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (lacp *Lacp) SetParent(parent types.Entity) { lacp.parent = parent }

func (lacp *Lacp) GetParent() types.Entity { return lacp.parent }

func (lacp *Lacp) GetParentYangName() string { return "openconfig-lacp" }

// Lacp_Config
// Configuration data for LACP
type Lacp_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sytem priority used by the node on this LAG interface. Lower value is
    // higher priority for determining which node is the controlling system. The
    // type is interface{} with range: 0..65535.
    SystemPriority interface{}
}

func (config *Lacp_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Lacp_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Lacp_Config) GetGoName(yname string) string {
    if yname == "system-priority" { return "SystemPriority" }
    return ""
}

func (config *Lacp_Config) GetSegmentPath() string {
    return "config"
}

func (config *Lacp_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Lacp_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Lacp_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-priority"] = config.SystemPriority
    return leafs
}

func (config *Lacp_Config) GetBundleName() string { return "openconfig" }

func (config *Lacp_Config) GetYangName() string { return "config" }

func (config *Lacp_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Lacp_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Lacp_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Lacp_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Lacp_Config) GetParent() types.Entity { return config.parent }

func (config *Lacp_Config) GetParentYangName() string { return "lacp" }

// Lacp_State
// Operational state data for LACP
type Lacp_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sytem priority used by the node on this LAG interface. Lower value is
    // higher priority for determining which node is the controlling system. The
    // type is interface{} with range: 0..65535.
    SystemPriority interface{}
}

func (state *Lacp_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Lacp_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Lacp_State) GetGoName(yname string) string {
    if yname == "system-priority" { return "SystemPriority" }
    return ""
}

func (state *Lacp_State) GetSegmentPath() string {
    return "state"
}

func (state *Lacp_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Lacp_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Lacp_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-priority"] = state.SystemPriority
    return leafs
}

func (state *Lacp_State) GetBundleName() string { return "openconfig" }

func (state *Lacp_State) GetYangName() string { return "state" }

func (state *Lacp_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Lacp_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Lacp_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Lacp_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Lacp_State) GetParent() types.Entity { return state.parent }

func (state *Lacp_State) GetParentYangName() string { return "lacp" }

// Lacp_Interfaces
// Enclosing container for the list of LACP-enabled
// interfaces
type Lacp_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of aggregate interfaces managed by LACP. The type is slice of
    // Lacp_Interfaces_Interface.
    Interface []Lacp_Interfaces_Interface
}

func (interfaces *Lacp_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Lacp_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Lacp_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Lacp_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Lacp_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lacp_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Lacp_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Lacp_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Lacp_Interfaces) GetBundleName() string { return "openconfig" }

func (interfaces *Lacp_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Lacp_Interfaces) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaces *Lacp_Interfaces) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaces *Lacp_Interfaces) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaces *Lacp_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Lacp_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Lacp_Interfaces) GetParentYangName() string { return "lacp" }

// Lacp_Interfaces_Interface
// List of aggregate interfaces managed by LACP
type Lacp_Interfaces_Interface struct {
    parent types.Entity
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

func (self *Lacp_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Lacp_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Lacp_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "members" { return "Members" }
    return ""
}

func (self *Lacp_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
}

func (self *Lacp_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &self.Config
    }
    if childYangName == "state" {
        return &self.State
    }
    if childYangName == "members" {
        return &self.Members
    }
    return nil
}

func (self *Lacp_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &self.Config
    children["state"] = &self.State
    children["members"] = &self.Members
    return children
}

func (self *Lacp_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    return leafs
}

func (self *Lacp_Interfaces_Interface) GetBundleName() string { return "openconfig" }

func (self *Lacp_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Lacp_Interfaces_Interface) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (self *Lacp_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (self *Lacp_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (self *Lacp_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Lacp_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Lacp_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Lacp_Interfaces_Interface_Config
// Configuration data for each LACP aggregate interface
type Lacp_Interfaces_Interface_Config struct {
    parent types.Entity
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
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemIdMac interface{}

    // Sytem priority used by the node on this LAG interface. Lower value is
    // higher priority for determining which node is the controlling system. The
    // type is interface{} with range: 0..65535.
    SystemPriority interface{}
}

func (config *Lacp_Interfaces_Interface_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Lacp_Interfaces_Interface_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Lacp_Interfaces_Interface_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "interval" { return "Interval" }
    if yname == "lacp-mode" { return "LacpMode" }
    if yname == "system-id-mac" { return "SystemIdMac" }
    if yname == "system-priority" { return "SystemPriority" }
    return ""
}

func (config *Lacp_Interfaces_Interface_Config) GetSegmentPath() string {
    return "config"
}

func (config *Lacp_Interfaces_Interface_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Lacp_Interfaces_Interface_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Lacp_Interfaces_Interface_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    leafs["interval"] = config.Interval
    leafs["lacp-mode"] = config.LacpMode
    leafs["system-id-mac"] = config.SystemIdMac
    leafs["system-priority"] = config.SystemPriority
    return leafs
}

func (config *Lacp_Interfaces_Interface_Config) GetBundleName() string { return "openconfig" }

func (config *Lacp_Interfaces_Interface_Config) GetYangName() string { return "config" }

func (config *Lacp_Interfaces_Interface_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Lacp_Interfaces_Interface_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Lacp_Interfaces_Interface_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Lacp_Interfaces_Interface_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Lacp_Interfaces_Interface_Config) GetParent() types.Entity { return config.parent }

func (config *Lacp_Interfaces_Interface_Config) GetParentYangName() string { return "interface" }

// Lacp_Interfaces_Interface_State
// Operational state data for each LACP aggregate
// interface
type Lacp_Interfaces_Interface_State struct {
    parent types.Entity
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
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemIdMac interface{}

    // Sytem priority used by the node on this LAG interface. Lower value is
    // higher priority for determining which node is the controlling system. The
    // type is interface{} with range: 0..65535.
    SystemPriority interface{}
}

func (state *Lacp_Interfaces_Interface_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Lacp_Interfaces_Interface_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Lacp_Interfaces_Interface_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "interval" { return "Interval" }
    if yname == "lacp-mode" { return "LacpMode" }
    if yname == "system-id-mac" { return "SystemIdMac" }
    if yname == "system-priority" { return "SystemPriority" }
    return ""
}

func (state *Lacp_Interfaces_Interface_State) GetSegmentPath() string {
    return "state"
}

func (state *Lacp_Interfaces_Interface_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Lacp_Interfaces_Interface_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Lacp_Interfaces_Interface_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    leafs["interval"] = state.Interval
    leafs["lacp-mode"] = state.LacpMode
    leafs["system-id-mac"] = state.SystemIdMac
    leafs["system-priority"] = state.SystemPriority
    return leafs
}

func (state *Lacp_Interfaces_Interface_State) GetBundleName() string { return "openconfig" }

func (state *Lacp_Interfaces_Interface_State) GetYangName() string { return "state" }

func (state *Lacp_Interfaces_Interface_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Lacp_Interfaces_Interface_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Lacp_Interfaces_Interface_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Lacp_Interfaces_Interface_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Lacp_Interfaces_Interface_State) GetParent() types.Entity { return state.parent }

func (state *Lacp_Interfaces_Interface_State) GetParentYangName() string { return "interface" }

// Lacp_Interfaces_Interface_Members
// Enclosing container for the list of members interfaces of
// the aggregate. This list is considered operational state
// only so is labeled config false and has no config container
type Lacp_Interfaces_Interface_Members struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of member interfaces and their associated status for a LACP-controlled
    // aggregate interface.  Member list is not configurable here -- each
    // interface indicates items its participation in the LAG. The type is slice
    // of Lacp_Interfaces_Interface_Members_Member.
    Member []Lacp_Interfaces_Interface_Members_Member
}

func (members *Lacp_Interfaces_Interface_Members) GetFilter() yfilter.YFilter { return members.YFilter }

func (members *Lacp_Interfaces_Interface_Members) SetFilter(yf yfilter.YFilter) { members.YFilter = yf }

func (members *Lacp_Interfaces_Interface_Members) GetGoName(yname string) string {
    if yname == "member" { return "Member" }
    return ""
}

func (members *Lacp_Interfaces_Interface_Members) GetSegmentPath() string {
    return "members"
}

func (members *Lacp_Interfaces_Interface_Members) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "member" {
        for _, c := range members.Member {
            if members.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lacp_Interfaces_Interface_Members_Member{}
        members.Member = append(members.Member, child)
        return &members.Member[len(members.Member)-1]
    }
    return nil
}

func (members *Lacp_Interfaces_Interface_Members) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range members.Member {
        children[members.Member[i].GetSegmentPath()] = &members.Member[i]
    }
    return children
}

func (members *Lacp_Interfaces_Interface_Members) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (members *Lacp_Interfaces_Interface_Members) GetBundleName() string { return "openconfig" }

func (members *Lacp_Interfaces_Interface_Members) GetYangName() string { return "members" }

func (members *Lacp_Interfaces_Interface_Members) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (members *Lacp_Interfaces_Interface_Members) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (members *Lacp_Interfaces_Interface_Members) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (members *Lacp_Interfaces_Interface_Members) SetParent(parent types.Entity) { members.parent = parent }

func (members *Lacp_Interfaces_Interface_Members) GetParent() types.Entity { return members.parent }

func (members *Lacp_Interfaces_Interface_Members) GetParentYangName() string { return "interface" }

// Lacp_Interfaces_Interface_Members_Member
// List of member interfaces and their associated status for
// a LACP-controlled aggregate interface.  Member list is not
// configurable here -- each interface indicates items
// its participation in the LAG.
type Lacp_Interfaces_Interface_Members_Member struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to aggregate member interface. The type
    // is string. Refers to
    // lacp.Lacp_Interfaces_Interface_Members_Member_State_Interface
    Interface interface{}

    // Operational state data for aggregate members.
    State Lacp_Interfaces_Interface_Members_Member_State
}

func (member *Lacp_Interfaces_Interface_Members_Member) GetFilter() yfilter.YFilter { return member.YFilter }

func (member *Lacp_Interfaces_Interface_Members_Member) SetFilter(yf yfilter.YFilter) { member.YFilter = yf }

func (member *Lacp_Interfaces_Interface_Members_Member) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "state" { return "State" }
    return ""
}

func (member *Lacp_Interfaces_Interface_Members_Member) GetSegmentPath() string {
    return "member" + "[interface='" + fmt.Sprintf("%v", member.Interface) + "']"
}

func (member *Lacp_Interfaces_Interface_Members_Member) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state" {
        return &member.State
    }
    return nil
}

func (member *Lacp_Interfaces_Interface_Members_Member) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["state"] = &member.State
    return children
}

func (member *Lacp_Interfaces_Interface_Members_Member) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = member.Interface
    return leafs
}

func (member *Lacp_Interfaces_Interface_Members_Member) GetBundleName() string { return "openconfig" }

func (member *Lacp_Interfaces_Interface_Members_Member) GetYangName() string { return "member" }

func (member *Lacp_Interfaces_Interface_Members_Member) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (member *Lacp_Interfaces_Interface_Members_Member) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (member *Lacp_Interfaces_Interface_Members_Member) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (member *Lacp_Interfaces_Interface_Members_Member) SetParent(parent types.Entity) { member.parent = parent }

func (member *Lacp_Interfaces_Interface_Members_Member) GetParent() types.Entity { return member.parent }

func (member *Lacp_Interfaces_Interface_Members_Member) GetParentYangName() string { return "members" }

// Lacp_Interfaces_Interface_Members_Member_State
// Operational state data for aggregate members
type Lacp_Interfaces_Interface_Members_Member_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reference to interface member of the LACP aggregate. The type is string.
    // Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

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
    // The type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemId interface{}

    // Current operational value of the key for the aggregate interface. The type
    // is interface{} with range: 0..65535.
    OperKey interface{}

    // MAC address representing the protocol partner's interface system ID. The
    // type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PartnerId interface{}

    // Operational value of the protocol partner's key. The type is interface{}
    // with range: 0..65535.
    PartnerKey interface{}

    // LACP protocol counters.
    Counters Lacp_Interfaces_Interface_Members_Member_State_Counters
}

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Lacp_Interfaces_Interface_Members_Member_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "activity" { return "Activity" }
    if yname == "timeout" { return "Timeout" }
    if yname == "synchronization" { return "Synchronization" }
    if yname == "aggregatable" { return "Aggregatable" }
    if yname == "collecting" { return "Collecting" }
    if yname == "distributing" { return "Distributing" }
    if yname == "system-id" { return "SystemId" }
    if yname == "oper-key" { return "OperKey" }
    if yname == "partner-id" { return "PartnerId" }
    if yname == "partner-key" { return "PartnerKey" }
    if yname == "counters" { return "Counters" }
    return ""
}

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetSegmentPath() string {
    return "state"
}

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counters" {
        return &state.Counters
    }
    return nil
}

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counters"] = &state.Counters
    return children
}

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = state.Interface
    leafs["activity"] = state.Activity
    leafs["timeout"] = state.Timeout
    leafs["synchronization"] = state.Synchronization
    leafs["aggregatable"] = state.Aggregatable
    leafs["collecting"] = state.Collecting
    leafs["distributing"] = state.Distributing
    leafs["system-id"] = state.SystemId
    leafs["oper-key"] = state.OperKey
    leafs["partner-id"] = state.PartnerId
    leafs["partner-key"] = state.PartnerKey
    return leafs
}

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetBundleName() string { return "openconfig" }

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetYangName() string { return "state" }

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Lacp_Interfaces_Interface_Members_Member_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetParent() types.Entity { return state.parent }

func (state *Lacp_Interfaces_Interface_Members_Member_State) GetParentYangName() string { return "member" }

// Lacp_Interfaces_Interface_Members_Member_State_Counters
// LACP protocol counters
type Lacp_Interfaces_Interface_Members_Member_State_Counters struct {
    parent types.Entity
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

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetGoName(yname string) string {
    if yname == "lacp-in-pkts" { return "LacpInPkts" }
    if yname == "lacp-out-pkts" { return "LacpOutPkts" }
    if yname == "lacp-rx-errors" { return "LacpRxErrors" }
    if yname == "lacp-tx-errors" { return "LacpTxErrors" }
    if yname == "lacp-unknown-errors" { return "LacpUnknownErrors" }
    if yname == "lacp-errors" { return "LacpErrors" }
    return ""
}

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lacp-in-pkts"] = counters.LacpInPkts
    leafs["lacp-out-pkts"] = counters.LacpOutPkts
    leafs["lacp-rx-errors"] = counters.LacpRxErrors
    leafs["lacp-tx-errors"] = counters.LacpTxErrors
    leafs["lacp-unknown-errors"] = counters.LacpUnknownErrors
    leafs["lacp-errors"] = counters.LacpErrors
    return leafs
}

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetBundleName() string { return "openconfig" }

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetYangName() string { return "counters" }

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetParent() types.Entity { return counters.parent }

func (counters *Lacp_Interfaces_Interface_Members_Member_State_Counters) GetParentYangName() string { return "state" }

