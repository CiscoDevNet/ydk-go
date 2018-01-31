// This module defines configuration and state variables for VLANs,
// in addition to VLAN parameters associated with interfaces
package vlan

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package vlan"))
    ydk.RegisterEntity("{http://openconfig.net/yang/vlan vlans}", reflect.TypeOf(Vlans{}))
    ydk.RegisterEntity("openconfig-vlan:vlans", reflect.TypeOf(Vlans{}))
}

// Vlans
// Container for VLAN configuration and state
// variables
type Vlans struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configured VLANs keyed by id. The type is slice of Vlans_Vlan.
    Vlan []Vlans_Vlan
}

func (vlans *Vlans) GetFilter() yfilter.YFilter { return vlans.YFilter }

func (vlans *Vlans) SetFilter(yf yfilter.YFilter) { vlans.YFilter = yf }

func (vlans *Vlans) GetGoName(yname string) string {
    if yname == "vlan" { return "Vlan" }
    return ""
}

func (vlans *Vlans) GetSegmentPath() string {
    return "openconfig-vlan:vlans"
}

func (vlans *Vlans) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vlan" {
        for _, c := range vlans.Vlan {
            if vlans.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlans_Vlan{}
        vlans.Vlan = append(vlans.Vlan, child)
        return &vlans.Vlan[len(vlans.Vlan)-1]
    }
    return nil
}

func (vlans *Vlans) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vlans.Vlan {
        children[vlans.Vlan[i].GetSegmentPath()] = &vlans.Vlan[i]
    }
    return children
}

func (vlans *Vlans) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vlans *Vlans) GetBundleName() string { return "openconfig" }

func (vlans *Vlans) GetYangName() string { return "vlans" }

func (vlans *Vlans) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (vlans *Vlans) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (vlans *Vlans) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (vlans *Vlans) SetParent(parent types.Entity) { vlans.parent = parent }

func (vlans *Vlans) GetParent() types.Entity { return vlans.parent }

func (vlans *Vlans) GetParentYangName() string { return "openconfig-vlan" }

// Vlans_Vlan
// Configured VLANs keyed by id
type Vlans_Vlan struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. references the configured vlan-id. The type is
    // string with range: 1..4094. Refers to vlan.Vlans_Vlan_Config_VlanId
    VlanId interface{}

    // Configuration parameters for VLANs.
    Config Vlans_Vlan_Config

    // State variables for VLANs.
    State Vlans_Vlan_State

    // Enclosing container for list of member interfaces.
    Members Vlans_Vlan_Members
}

func (vlan *Vlans_Vlan) GetFilter() yfilter.YFilter { return vlan.YFilter }

func (vlan *Vlans_Vlan) SetFilter(yf yfilter.YFilter) { vlan.YFilter = yf }

func (vlan *Vlans_Vlan) GetGoName(yname string) string {
    if yname == "vlan-id" { return "VlanId" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "members" { return "Members" }
    return ""
}

func (vlan *Vlans_Vlan) GetSegmentPath() string {
    return "vlan" + "[vlan-id='" + fmt.Sprintf("%v", vlan.VlanId) + "']"
}

func (vlan *Vlans_Vlan) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &vlan.Config
    }
    if childYangName == "state" {
        return &vlan.State
    }
    if childYangName == "members" {
        return &vlan.Members
    }
    return nil
}

func (vlan *Vlans_Vlan) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &vlan.Config
    children["state"] = &vlan.State
    children["members"] = &vlan.Members
    return children
}

func (vlan *Vlans_Vlan) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlan-id"] = vlan.VlanId
    return leafs
}

func (vlan *Vlans_Vlan) GetBundleName() string { return "openconfig" }

func (vlan *Vlans_Vlan) GetYangName() string { return "vlan" }

func (vlan *Vlans_Vlan) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (vlan *Vlans_Vlan) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (vlan *Vlans_Vlan) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (vlan *Vlans_Vlan) SetParent(parent types.Entity) { vlan.parent = parent }

func (vlan *Vlans_Vlan) GetParent() types.Entity { return vlan.parent }

func (vlan *Vlans_Vlan) GetParentYangName() string { return "vlans" }

// Vlans_Vlan_Config
// Configuration parameters for VLANs
type Vlans_Vlan_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface VLAN id. The type is interface{} with range: 1..4094.
    VlanId interface{}

    // Interface VLAN name. The type is string.
    Name interface{}

    // Admin state of the VLAN. The type is Status. The default value is ACTIVE.
    Status interface{}

    // Optionally set the tag protocol identifier field (TPID) that is accepted on
    // the VLAN. The type is one of the following:
    // TPID0X9200TPID0x8A88TPID0x8100TPID0x9100. The default value is
    // oc-vlan-types:TPID_0x8100.
    Tpid interface{}
}

func (config *Vlans_Vlan_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Vlans_Vlan_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Vlans_Vlan_Config) GetGoName(yname string) string {
    if yname == "vlan-id" { return "VlanId" }
    if yname == "name" { return "Name" }
    if yname == "status" { return "Status" }
    if yname == "tpid" { return "Tpid" }
    return ""
}

func (config *Vlans_Vlan_Config) GetSegmentPath() string {
    return "config"
}

func (config *Vlans_Vlan_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Vlans_Vlan_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Vlans_Vlan_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlan-id"] = config.VlanId
    leafs["name"] = config.Name
    leafs["status"] = config.Status
    leafs["tpid"] = config.Tpid
    return leafs
}

func (config *Vlans_Vlan_Config) GetBundleName() string { return "openconfig" }

func (config *Vlans_Vlan_Config) GetYangName() string { return "config" }

func (config *Vlans_Vlan_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Vlans_Vlan_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Vlans_Vlan_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Vlans_Vlan_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Vlans_Vlan_Config) GetParent() types.Entity { return config.parent }

func (config *Vlans_Vlan_Config) GetParentYangName() string { return "vlan" }

// Vlans_Vlan_Config_Status represents Admin state of the VLAN
type Vlans_Vlan_Config_Status string

const (
    // VLAN is active
    Vlans_Vlan_Config_Status_ACTIVE Vlans_Vlan_Config_Status = "ACTIVE"

    // VLAN is inactive / suspended
    Vlans_Vlan_Config_Status_SUSPENDED Vlans_Vlan_Config_Status = "SUSPENDED"
)

// Vlans_Vlan_State
// State variables for VLANs
type Vlans_Vlan_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface VLAN id. The type is interface{} with range: 1..4094.
    VlanId interface{}

    // Interface VLAN name. The type is string.
    Name interface{}

    // Admin state of the VLAN. The type is Status. The default value is ACTIVE.
    Status interface{}

    // Optionally set the tag protocol identifier field (TPID) that is accepted on
    // the VLAN. The type is one of the following:
    // TPID0X9200TPID0x8A88TPID0x8100TPID0x9100. The default value is
    // oc-vlan-types:TPID_0x8100.
    Tpid interface{}
}

func (state *Vlans_Vlan_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Vlans_Vlan_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Vlans_Vlan_State) GetGoName(yname string) string {
    if yname == "vlan-id" { return "VlanId" }
    if yname == "name" { return "Name" }
    if yname == "status" { return "Status" }
    if yname == "tpid" { return "Tpid" }
    return ""
}

func (state *Vlans_Vlan_State) GetSegmentPath() string {
    return "state"
}

func (state *Vlans_Vlan_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Vlans_Vlan_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Vlans_Vlan_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlan-id"] = state.VlanId
    leafs["name"] = state.Name
    leafs["status"] = state.Status
    leafs["tpid"] = state.Tpid
    return leafs
}

func (state *Vlans_Vlan_State) GetBundleName() string { return "openconfig" }

func (state *Vlans_Vlan_State) GetYangName() string { return "state" }

func (state *Vlans_Vlan_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Vlans_Vlan_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Vlans_Vlan_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Vlans_Vlan_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Vlans_Vlan_State) GetParent() types.Entity { return state.parent }

func (state *Vlans_Vlan_State) GetParentYangName() string { return "vlan" }

// Vlans_Vlan_State_Status represents Admin state of the VLAN
type Vlans_Vlan_State_Status string

const (
    // VLAN is active
    Vlans_Vlan_State_Status_ACTIVE Vlans_Vlan_State_Status = "ACTIVE"

    // VLAN is inactive / suspended
    Vlans_Vlan_State_Status_SUSPENDED Vlans_Vlan_State_Status = "SUSPENDED"
)

// Vlans_Vlan_Members
// Enclosing container for list of member interfaces
type Vlans_Vlan_Members struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of references to interfaces / subinterfaces associated with the VLAN.
    // The type is slice of Vlans_Vlan_Members_Member.
    Member []Vlans_Vlan_Members_Member
}

func (members *Vlans_Vlan_Members) GetFilter() yfilter.YFilter { return members.YFilter }

func (members *Vlans_Vlan_Members) SetFilter(yf yfilter.YFilter) { members.YFilter = yf }

func (members *Vlans_Vlan_Members) GetGoName(yname string) string {
    if yname == "member" { return "Member" }
    return ""
}

func (members *Vlans_Vlan_Members) GetSegmentPath() string {
    return "members"
}

func (members *Vlans_Vlan_Members) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "member" {
        for _, c := range members.Member {
            if members.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vlans_Vlan_Members_Member{}
        members.Member = append(members.Member, child)
        return &members.Member[len(members.Member)-1]
    }
    return nil
}

func (members *Vlans_Vlan_Members) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range members.Member {
        children[members.Member[i].GetSegmentPath()] = &members.Member[i]
    }
    return children
}

func (members *Vlans_Vlan_Members) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (members *Vlans_Vlan_Members) GetBundleName() string { return "openconfig" }

func (members *Vlans_Vlan_Members) GetYangName() string { return "members" }

func (members *Vlans_Vlan_Members) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (members *Vlans_Vlan_Members) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (members *Vlans_Vlan_Members) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (members *Vlans_Vlan_Members) SetParent(parent types.Entity) { members.parent = parent }

func (members *Vlans_Vlan_Members) GetParent() types.Entity { return members.parent }

func (members *Vlans_Vlan_Members) GetParentYangName() string { return "vlan" }

// Vlans_Vlan_Members_Member
// List of references to interfaces / subinterfaces
// associated with the VLAN.
type Vlans_Vlan_Members_Member struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reference to an interface or subinterface.
    InterfaceRef Vlans_Vlan_Members_Member_InterfaceRef
}

func (member *Vlans_Vlan_Members_Member) GetFilter() yfilter.YFilter { return member.YFilter }

func (member *Vlans_Vlan_Members_Member) SetFilter(yf yfilter.YFilter) { member.YFilter = yf }

func (member *Vlans_Vlan_Members_Member) GetGoName(yname string) string {
    if yname == "interface-ref" { return "InterfaceRef" }
    return ""
}

func (member *Vlans_Vlan_Members_Member) GetSegmentPath() string {
    return "member"
}

func (member *Vlans_Vlan_Members_Member) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-ref" {
        return &member.InterfaceRef
    }
    return nil
}

func (member *Vlans_Vlan_Members_Member) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-ref"] = &member.InterfaceRef
    return children
}

func (member *Vlans_Vlan_Members_Member) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (member *Vlans_Vlan_Members_Member) GetBundleName() string { return "openconfig" }

func (member *Vlans_Vlan_Members_Member) GetYangName() string { return "member" }

func (member *Vlans_Vlan_Members_Member) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (member *Vlans_Vlan_Members_Member) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (member *Vlans_Vlan_Members_Member) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (member *Vlans_Vlan_Members_Member) SetParent(parent types.Entity) { member.parent = parent }

func (member *Vlans_Vlan_Members_Member) GetParent() types.Entity { return member.parent }

func (member *Vlans_Vlan_Members_Member) GetParentYangName() string { return "members" }

// Vlans_Vlan_Members_Member_InterfaceRef
// Reference to an interface or subinterface
type Vlans_Vlan_Members_Member_InterfaceRef struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational state for interface-ref.
    State Vlans_Vlan_Members_Member_InterfaceRef_State
}

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetFilter() yfilter.YFilter { return interfaceRef.YFilter }

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) SetFilter(yf yfilter.YFilter) { interfaceRef.YFilter = yf }

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetGoName(yname string) string {
    if yname == "state" { return "State" }
    return ""
}

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetSegmentPath() string {
    return "interface-ref"
}

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state" {
        return &interfaceRef.State
    }
    return nil
}

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["state"] = &interfaceRef.State
    return children
}

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetBundleName() string { return "openconfig" }

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetYangName() string { return "interface-ref" }

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) SetParent(parent types.Entity) { interfaceRef.parent = parent }

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetParent() types.Entity { return interfaceRef.parent }

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetParentYangName() string { return "member" }

// Vlans_Vlan_Members_Member_InterfaceRef_State
// Operational state for interface-ref
type Vlans_Vlan_Members_Member_InterfaceRef_State struct {
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

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetSegmentPath() string {
    return "state"
}

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = state.Interface
    leafs["subinterface"] = state.Subinterface
    return leafs
}

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetBundleName() string { return "openconfig" }

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetYangName() string { return "state" }

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetParent() types.Entity { return state.parent }

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetParentYangName() string { return "interface-ref" }

