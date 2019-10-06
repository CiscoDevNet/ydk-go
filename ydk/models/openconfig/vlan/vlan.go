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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured VLANs keyed by id. The type is slice of Vlans_Vlan.
    Vlan []*Vlans_Vlan
}

func (vlans *Vlans) GetEntityData() *types.CommonEntityData {
    vlans.EntityData.YFilter = vlans.YFilter
    vlans.EntityData.YangName = "vlans"
    vlans.EntityData.BundleName = "openconfig"
    vlans.EntityData.ParentYangName = "openconfig-vlan"
    vlans.EntityData.SegmentPath = "openconfig-vlan:vlans"
    vlans.EntityData.AbsolutePath = vlans.EntityData.SegmentPath
    vlans.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    vlans.EntityData.NamespaceTable = openconfig.GetNamespaces()
    vlans.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    vlans.EntityData.Children = types.NewOrderedMap()
    vlans.EntityData.Children.Append("vlan", types.YChild{"Vlan", nil})
    for i := range vlans.Vlan {
        vlans.EntityData.Children.Append(types.GetSegmentPath(vlans.Vlan[i]), types.YChild{"Vlan", vlans.Vlan[i]})
    }
    vlans.EntityData.Leafs = types.NewOrderedMap()

    vlans.EntityData.YListKeys = []string {}

    return &(vlans.EntityData)
}

// Vlans_Vlan
// Configured VLANs keyed by id
type Vlans_Vlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (vlan *Vlans_Vlan) GetEntityData() *types.CommonEntityData {
    vlan.EntityData.YFilter = vlan.YFilter
    vlan.EntityData.YangName = "vlan"
    vlan.EntityData.BundleName = "openconfig"
    vlan.EntityData.ParentYangName = "vlans"
    vlan.EntityData.SegmentPath = "vlan" + types.AddKeyToken(vlan.VlanId, "vlan-id")
    vlan.EntityData.AbsolutePath = "openconfig-vlan:vlans/" + vlan.EntityData.SegmentPath
    vlan.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    vlan.EntityData.NamespaceTable = openconfig.GetNamespaces()
    vlan.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    vlan.EntityData.Children = types.NewOrderedMap()
    vlan.EntityData.Children.Append("config", types.YChild{"Config", &vlan.Config})
    vlan.EntityData.Children.Append("state", types.YChild{"State", &vlan.State})
    vlan.EntityData.Children.Append("members", types.YChild{"Members", &vlan.Members})
    vlan.EntityData.Leafs = types.NewOrderedMap()
    vlan.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", vlan.VlanId})

    vlan.EntityData.YListKeys = []string {"VlanId"}

    return &(vlan.EntityData)
}

// Vlans_Vlan_Config
// Configuration parameters for VLANs
type Vlans_Vlan_Config struct {
    EntityData types.CommonEntityData
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

func (config *Vlans_Vlan_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "vlan"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-vlan:vlans/vlan/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", config.VlanId})
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("status", types.YLeaf{"Status", config.Status})
    config.EntityData.Leafs.Append("tpid", types.YLeaf{"Tpid", config.Tpid})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (state *Vlans_Vlan_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "vlan"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-vlan:vlans/vlan/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", state.VlanId})
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("status", types.YLeaf{"Status", state.Status})
    state.EntityData.Leafs.Append("tpid", types.YLeaf{"Tpid", state.Tpid})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of references to interfaces / subinterfaces associated with the VLAN.
    // The type is slice of Vlans_Vlan_Members_Member.
    Member []*Vlans_Vlan_Members_Member
}

func (members *Vlans_Vlan_Members) GetEntityData() *types.CommonEntityData {
    members.EntityData.YFilter = members.YFilter
    members.EntityData.YangName = "members"
    members.EntityData.BundleName = "openconfig"
    members.EntityData.ParentYangName = "vlan"
    members.EntityData.SegmentPath = "members"
    members.EntityData.AbsolutePath = "openconfig-vlan:vlans/vlan/" + members.EntityData.SegmentPath
    members.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    members.EntityData.NamespaceTable = openconfig.GetNamespaces()
    members.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    members.EntityData.Children = types.NewOrderedMap()
    members.EntityData.Children.Append("member", types.YChild{"Member", nil})
    for i := range members.Member {
        types.SetYListKey(members.Member[i], i)
        members.EntityData.Children.Append(types.GetSegmentPath(members.Member[i]), types.YChild{"Member", members.Member[i]})
    }
    members.EntityData.Leafs = types.NewOrderedMap()

    members.EntityData.YListKeys = []string {}

    return &(members.EntityData)
}

// Vlans_Vlan_Members_Member
// List of references to interfaces / subinterfaces
// associated with the VLAN.
type Vlans_Vlan_Members_Member struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Reference to an interface or subinterface.
    InterfaceRef Vlans_Vlan_Members_Member_InterfaceRef
}

func (member *Vlans_Vlan_Members_Member) GetEntityData() *types.CommonEntityData {
    member.EntityData.YFilter = member.YFilter
    member.EntityData.YangName = "member"
    member.EntityData.BundleName = "openconfig"
    member.EntityData.ParentYangName = "members"
    member.EntityData.SegmentPath = "member" + types.AddNoKeyToken(member)
    member.EntityData.AbsolutePath = "openconfig-vlan:vlans/vlan/members/" + member.EntityData.SegmentPath
    member.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    member.EntityData.NamespaceTable = openconfig.GetNamespaces()
    member.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    member.EntityData.Children = types.NewOrderedMap()
    member.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &member.InterfaceRef})
    member.EntityData.Leafs = types.NewOrderedMap()

    member.EntityData.YListKeys = []string {}

    return &(member.EntityData)
}

// Vlans_Vlan_Members_Member_InterfaceRef
// Reference to an interface or subinterface
type Vlans_Vlan_Members_Member_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational state for interface-ref.
    State Vlans_Vlan_Members_Member_InterfaceRef_State
}

func (interfaceRef *Vlans_Vlan_Members_Member_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "openconfig"
    interfaceRef.EntityData.ParentYangName = "member"
    interfaceRef.EntityData.SegmentPath = "interface-ref"
    interfaceRef.EntityData.AbsolutePath = "openconfig-vlan:vlans/vlan/members/member/" + interfaceRef.EntityData.SegmentPath
    interfaceRef.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceRef.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceRef.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceRef.EntityData.Children = types.NewOrderedMap()
    interfaceRef.EntityData.Children.Append("state", types.YChild{"State", &interfaceRef.State})
    interfaceRef.EntityData.Leafs = types.NewOrderedMap()

    interfaceRef.EntityData.YListKeys = []string {}

    return &(interfaceRef.EntityData)
}

// Vlans_Vlan_Members_Member_InterfaceRef_State
// Operational state for interface-ref
type Vlans_Vlan_Members_Member_InterfaceRef_State struct {
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

func (state *Vlans_Vlan_Members_Member_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface-ref"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-vlan:vlans/vlan/members/member/interface-ref/" + state.EntityData.SegmentPath
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

